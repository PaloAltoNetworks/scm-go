// Package common
/*
Testing utilities for common
*/
package common

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"time"
)

// LoggingRoundTripper Custom HTTP client that logs requests as curl commands
type LoggingRoundTripper struct {
	Wrapped http.RoundTripper
}

func (lrt *LoggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	// Print the curl command equivalent
	printCurlCommand(req)
	// Execute the actual request
	return lrt.Wrapped.RoundTrip(req)
}

// maskCurlCommand masks sensitive information in curl commands
func maskCurlCommand(command string) string {
	// Mask Authorization headers
	authHeaderRegex := regexp.MustCompile(`(-H\s+['"]Authorization:\s*Bearer\s+)([^'"]+)(['"])`)
	command = authHeaderRegex.ReplaceAllString(command, `${1}****${3}`)

	// Mask X-Auth-Jwt headers
	xAuthJwtRegex := regexp.MustCompile(`(-H\s+['"]X-Auth-Jwt:\s*)([^'"]+)(['"])`)
	command = xAuthJwtRegex.ReplaceAllString(command, `${1}****${3}`)

	// Mask API keys in headers
	apiKeyRegex := regexp.MustCompile(`(-H\s+['"](?:X-)?API-Key:\s*)([^'"]+)(['"])`)
	command = apiKeyRegex.ReplaceAllString(command, `${1}****${3}`)

	// Mask tokens in URL parameters
	tokenParamRegex := regexp.MustCompile(`([?&](?:token|access_token|api_key)=)([^&\s]+)`)
	command = tokenParamRegex.ReplaceAllString(command, `${1}****`)

	// Mask basic auth in URLs
	basicAuthRegex := regexp.MustCompile(`(https?://)([^:]+):([^@]+)(@)`)
	command = basicAuthRegex.ReplaceAllString(command, `${1}${2}:****${4}`)

	return command
}

func printCurlCommand(req *http.Request) {
	fmt.Printf("\n=== CURL COMMAND EQUIVALENT ===\n")
	// Start building the curl command
	curlCmd := []string{"curl", "-X", req.Method}

	// Add headers
	for name, values := range req.Header {
		for _, value := range values {
			curlCmd = append(curlCmd, "-H", fmt.Sprintf("'%s: %s'", name, value))
		}
	}

	// Add body if present
	if req.Body != nil {
		bodyBytes, err := io.ReadAll(req.Body)
		if err == nil && len(bodyBytes) > 0 {
			// Restore the body for the actual request
			req.Body = io.NopCloser(bytes.NewReader(bodyBytes))
			// Add to curl command
			curlCmd = append(curlCmd, "-d", fmt.Sprintf("'%s'", string(bodyBytes)))
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
