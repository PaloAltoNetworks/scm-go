// Package common
/*
Testing utilities for common
*/
package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"testing"
	"time"
)

// LoggingRoundTripper Custom HTTP client that logs requests as curl commands.
// Request and response bodies are omitted by default. Set SCM_LOG_BODIES=1 to
// print them; secret-like fields are always redacted when bodies are printed.
type LoggingRoundTripper struct {
	Wrapped http.RoundTripper
}

func (lrt *LoggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	// Print the curl command equivalent
	printCurlCommand(req)
	// Execute the actual request
	resp, err := lrt.Wrapped.RoundTrip(req)
	// Print headers and body for error responses
	if resp != nil && resp.StatusCode >= 400 {
		fmt.Printf("=== API RESPONSE HEADERS ===\n")
		fmt.Printf("Status Code: %d\n", resp.StatusCode)
		fmt.Printf("X-Request-ID: %s\n", resp.Header.Get("X-Request-ID"))
		fmt.Printf("X-Trace-ID: %s\n", resp.Header.Get("X-Trace-ID"))
		if flowErr := resp.Header.Get("X-Request-Flow-Error"); flowErr != "" {
			fmt.Printf("X-Request-Flow-Error: %s\n", flowErr)
		}
		fmt.Printf("============================\n")
		if resp.Body != nil {
			bodyBytes, readErr := io.ReadAll(resp.Body)
			resp.Body.Close()
			resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			if readErr == nil && len(bodyBytes) > 0 {
				fmt.Printf("=== API RESPONSE BODY ===\n%s\n=========================\n", redactBody(bodyBytes))
			}
		}
	}
	return resp, err
}

// logBodiesEnabled reports whether request/response bodies may be printed.
func logBodiesEnabled() bool {
	switch strings.ToLower(os.Getenv("SCM_LOG_BODIES")) {
	case "1", "true", "yes":
		return true
	}
	return false
}

// sensitiveNameFragments are case-insensitive fragments of JSON field names and
// header names whose values must never be printed.
var sensitiveNameFragments = []string{
	"passphrase", "password", "passwd", "secret", "token", "credential",
	"private", "key_file", "keyfile", "psk", "pre_shared", "preshared",
	"authorization", "cookie", "api_key", "apikey", "auth_key", "jwt",
}

func isSensitiveName(name string) bool {
	n := strings.ToLower(name)
	if strings.HasSuffix(n, "key") {
		return true
	}
	for _, frag := range sensitiveNameFragments {
		if strings.Contains(n, frag) {
			return true
		}
	}
	return false
}

// redactValue walks decoded JSON and replaces values of sensitive fields.
func redactValue(v interface{}) interface{} {
	switch t := v.(type) {
	case map[string]interface{}:
		for k, val := range t {
			if isSensitiveName(k) {
				t[k] = "****"
			} else {
				t[k] = redactValue(val)
			}
		}
		return t
	case []interface{}:
		for i, val := range t {
			t[i] = redactValue(val)
		}
		return t
	}
	return v
}

// RedactJSON returns a printable form of an HTTP body with secret-like JSON
// fields replaced by "****". Non-JSON bodies are never printed because they
// cannot be redacted reliably.
func RedactJSON(body []byte) string {
	dec := json.NewDecoder(bytes.NewReader(body))
	dec.UseNumber()
	var v interface{}
	if err := dec.Decode(&v); err != nil {
		return fmt.Sprintf("[non-JSON body omitted, %d bytes]", len(body))
	}
	out, err := json.Marshal(redactValue(v))
	if err != nil {
		return fmt.Sprintf("[body omitted, %d bytes]", len(body))
	}
	return string(out)
}

// redactBody returns a printable, redacted form of an HTTP body. Bodies are
// omitted unless SCM_LOG_BODIES is set.
func redactBody(body []byte) string {
	if !logBodiesEnabled() {
		return fmt.Sprintf("[body omitted, %d bytes; set SCM_LOG_BODIES=1 to print redacted]", len(body))
	}
	return RedactJSON(body)
}

// maskCurlCommand masks sensitive information in curl commands
func maskCurlCommand(command string) string {
	// Mask Authorization headers
	authHeaderRegex := regexp.MustCompile("(-H\\s+['\"]Authorization:\\s*Bearer\\s+)([^'\"]+)(['\"])")
	command = authHeaderRegex.ReplaceAllString(command, "${1}****${3}")

	// Mask X-Auth-Jwt headers
	xAuthJwtRegex := regexp.MustCompile("(-H\\s+['\"]X-Auth-Jwt:\\s*)([^'\"]+)(['\"])")
	command = xAuthJwtRegex.ReplaceAllString(command, "${1}****${3}")

	// Mask API keys in headers
	apiKeyRegex := regexp.MustCompile("(-H\\s+['\"](?:X-)?API-Key:\\s*)([^'\"]+)(['\"])")
	command = apiKeyRegex.ReplaceAllString(command, "${1}****${3}")

	// Mask secrets in URL parameters
	tokenParamRegex := regexp.MustCompile("(?i)([?&](?:token|access_token|api_key|client_secret|secret|password|passphrase|key)=)([^&\\s']+)")
	command = tokenParamRegex.ReplaceAllString(command, "${1}****")

	// Mask basic auth in URLs
	basicAuthRegex := regexp.MustCompile("(https?://)([^:]+):([^@]+)(@)")
	command = basicAuthRegex.ReplaceAllString(command, "${1}${2}:****${4}")

	return command
}

func printCurlCommand(req *http.Request) {
	fmt.Printf("\n=== CURL COMMAND EQUIVALENT ===\n")
	// Start building the curl command
	curlCmd := []string{"curl", "-X", req.Method}

	// Add headers, masking any sensitive header value
	for name, values := range req.Header {
		for _, value := range values {
			if isSensitiveName(name) {
				value = "****"
			}
			curlCmd = append(curlCmd, "-H", fmt.Sprintf("'%s: %s'", name, value))
		}
	}

	// Add body if present
	if req.Body != nil {
		bodyBytes, err := io.ReadAll(req.Body)
		if err == nil && len(bodyBytes) > 0 {
			// Restore the body for the actual request
			req.Body = io.NopCloser(bytes.NewReader(bodyBytes))
			// Add to curl command (omitted unless SCM_LOG_BODIES is set, always redacted)
			curlCmd = append(curlCmd, "-d", fmt.Sprintf("'%s'", redactBody(bodyBytes)))
		}
	}

	// Add URL
	curlCmd = append(curlCmd, fmt.Sprintf("'%s'", req.URL.String()))

	// Print the command with masked sensitive information
	curlCommand := strings.Join(curlCmd, " ")
	maskedCommand := maskCurlCommand(curlCommand)
	fmt.Printf("%s\n", maskedCommand)
	fmt.Printf("================================\n\n")
}

func GetConfigPath() string {
	// Get the current file's directory
	_, filename, _, _ := runtime.Caller(0)
	// Get the project root (3 levels up: test -> common -> generated -> root)
	projectRoot := filepath.Dir(filepath.Dir(filename))
	// Join with config file path
	return filepath.Join(projectRoot, "config", "scm-config.json")
}

func StringPtr(s string) *string {
	return &s
}

func BoolPtr(b bool) *bool {
	return &b
}

func IntPtr(i int) *int {
	return &i
}

func Int8Ptr(i int8) *int8 {
	return &i
}

func Int16Ptr(i int16) *int16 {
	return &i
}

func Int32Ptr(i int32) *int32 {
	return &i
}

func Int64Ptr(i int64) *int64 {
	return &i
}

func UintPtr(u uint) *uint {
	return &u
}

func Uint8Ptr(u uint8) *uint8 {
	return &u
}

func Uint16Ptr(u uint16) *uint16 {
	return &u
}

func Uint32Ptr(u uint32) *uint32 {
	return &u
}

func Uint64Ptr(u uint64) *uint64 {
	return &u
}

func Float32Ptr(f float32) *float32 {
	return &f
}

func Float64Ptr(f float64) *float64 {
	return &f
}

func GenerateRandomString(length int) string {
	// Simple random string generator for testing
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))] // Actually random!
	}
	return string(result)
}

// ErrorDetail represents individual error details
type ErrorDetail struct {
	Type    string   `json:"type"`
	Message string   `json:"message"`
	Params  []string `json:"params"`
}

// ErrorInfo represents error information structure
type ErrorInfo struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details struct {
		ErrorType string        `json:"errorType"`
		Message   []string      `json:"message"`
		Errors    []ErrorDetail `json:"errors"`
	} `json:"details"`
}

// APIErrorResponse represents the complete error response structure
type APIErrorResponse struct {
	Errors    []ErrorInfo `json:"_errors"`
	RequestID string      `json:"_request_id"`
}

// LogRequestIDOnFailure registers a cleanup that logs the X-Request-Id, X-Trace-Id, and X-Request-Flow-Error headers if the test fails.
// Call it once after each Execute() call; captures the request ID by value so multiple calls
// in the same test each log their own ID independently.
func LogRequestIDOnFailure(t *testing.T, httpRes *http.Response) {
	t.Helper()
	if httpRes == nil {
		return
	}
	reqID := httpRes.Header.Get("X-Request-Id")
	traceID := httpRes.Header.Get("X-Trace-Id")
	flowError := httpRes.Header.Get("X-Request-Flow-Error")
	method, path := "", ""
	if httpRes.Request != nil {
		method = httpRes.Request.Method
		path = httpRes.Request.URL.Path
	}
	t.Cleanup(func() {
		if t.Failed() {
			if reqID != "" {
				t.Logf("X-Request-ID [%s %s]: %s", method, path, reqID)
			} else if traceID != "" {
				t.Logf("X-Trace-ID [%s %s]: %s", method, path, traceID)
			} else if flowError != "" {
				t.Logf("X-Request-Flow-Error [%s %s]: %s", method, path, flowError)
			}
		}
	})
}
