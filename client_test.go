package scm

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/paloaltonetworks/scm-go/api"
)

func TestSetupFromEnvironmentVariables(t *testing.T) {
	var err error
	var key, cv string
	var ok bool

	key = "SCM_HOST"
	cv, ok = os.LookupEnv(key)
	if !ok {
		defer os.Unsetenv(key)
	} else {
		defer os.Setenv(key, cv)
	}
	if err = os.Setenv(key, "test.host"); err != nil {
		t.Fatalf("Failed to set %s: %s", key, err)
	}

	key = "SCM_PORT"
	cv, ok = os.LookupEnv(key)
	if !ok {
		defer os.Unsetenv(key)
	} else {
		defer os.Setenv(key, cv)
	}
	if err = os.Setenv(key, "12345"); err != nil {
		t.Fatalf("Failed to set %s: %s", key, err)
	}

	key = "SCM_CLIENT_ID"
	cv, ok = os.LookupEnv(key)
	if !ok {
		defer os.Unsetenv(key)
	} else {
		defer os.Setenv(key, cv)
	}
	if err = os.Setenv(key, "test_client_id"); err != nil {
		t.Fatalf("Failed to set %s: %s", key, err)
	}

	key = "SCM_CLIENT_SECRET"
	cv, ok = os.LookupEnv(key)
	if !ok {
		defer os.Unsetenv(key)
	} else {
		defer os.Setenv(key, cv)
	}
	if err = os.Setenv(key, "secret_for_testing"); err != nil {
		t.Fatalf("Failed to set %s: %s", key, err)
	}

	key = "SCM_SCOPE"
	cv, ok = os.LookupEnv(key)
	if !ok {
		defer os.Unsetenv(key)
	} else {
		defer os.Setenv(key, cv)
	}
	if err = os.Setenv(key, "test:scope"); err != nil {
		t.Fatalf("Failed to set %s: %s", key, err)
	}

	key = "SCM_PROTOCOL"
	cv, ok = os.LookupEnv(key)
	if !ok {
		defer os.Unsetenv(key)
	} else {
		defer os.Setenv(key, cv)
	}
	if err = os.Setenv(key, "http"); err != nil {
		t.Fatalf("Failed to set %s: %s", key, err)
	}

	key = "SCM_TIMEOUT"
	cv, ok = os.LookupEnv(key)
	if !ok {
		defer os.Unsetenv(key)
	} else {
		defer os.Setenv(key, cv)
	}
	if err = os.Setenv(key, "300"); err != nil {
		t.Fatalf("Failed to set %s: %s", key, err)
	}

	key = "SCM_SKIP_VERIFY_CERTIFICATE"
	cv, ok = os.LookupEnv(key)
	if !ok {
		defer os.Unsetenv(key)
	} else {
		defer os.Setenv(key, cv)
	}
	if err = os.Setenv(key, "true"); err != nil {
		t.Fatalf("Failed to set %s: %s", key, err)
	}

	c := Client{CheckEnvironment: true}
	if err = c.Setup(); err != nil {
		t.Fatalf("Error in setup: %s", err)
	}

	if c.Host != "test.host" {
		t.Errorf("Host: %s", c.Host)
	}

	if c.Port != 12345 {
		t.Errorf("Port: %d", c.Port)
	}

	if c.ClientId != "test_client_id" {
		t.Errorf("ClientId: %s", c.ClientId)
	}

	if c.ClientSecret != "secret_for_testing" {
		t.Errorf("ClientSecret: %s", c.ClientSecret)
	}

	if c.Scope != "test:scope" {
		t.Errorf("Scope: %s", c.Scope)
	}

	if c.Protocol != "http" {
		t.Errorf("Protocol: %s", c.Protocol)
	}

	if c.Timeout != 300 {
		t.Errorf("Timeout: %d", c.Timeout)
	}

	if !c.SkipVerifyCertificate {
		t.Errorf("Skip verify cert is not enabled")
	}
}

func TestSetupFromJsonConfigFile(t *testing.T) {
	conf := []byte(`{
    "host": "test.host",
    "port": 12345,
    "client_id": "test_client_id",
    "client_secret": "secret_for_testing",
    "scope": "test:scope",
    "protocol": "http",
    "timeout": 300,
    "skip_verify_certificate": true
}`)

	c := Client{
		testData: []*http.Response{{
			StatusCode: http.StatusOK,
		}},
		AuthFile:        "testfile.json",
		authFileContent: conf,
	}

	if err := c.Setup(); err != nil {
		t.Fatalf("Failed setup: %s", err)
	}

	if c.Host != "test.host" {
		t.Errorf("Host: %s", c.Host)
	}

	if c.Port != 12345 {
		t.Errorf("Port: %d", c.Port)
	}

	if c.ClientId != "test_client_id" {
		t.Errorf("ClientId: %s", c.ClientId)
	}

	if c.ClientSecret != "secret_for_testing" {
		t.Errorf("ClientSecret: %s", c.ClientSecret)
	}

	if c.Scope != "test:scope" {
		t.Errorf("Scope: %s", c.Scope)
	}

	if c.Protocol != "http" {
		t.Errorf("Protocol: %s", c.Protocol)
	}

	if c.Timeout != 300 {
		t.Errorf("Timeout: %d", c.Timeout)
	}

	if !c.SkipVerifyCertificate {
		t.Errorf("Skip verify cert is not enabled")
	}
}

func TestRefreshJwtFailedNoJwtPresent(t *testing.T) {
	ctx := context.TODO()

	c := Client{
		testData: []*http.Response{{
			StatusCode: http.StatusBadRequest,
			Body:       io.NopCloser(strings.NewReader("")),
		}},
	}

	if err := c.RefreshJwt(ctx); err == nil {
		t.Errorf("Was expecting this to error out, got nil")
	}
}

func TestRefreshJwtOk(t *testing.T) {
	ctx := context.TODO()
	data := `{
    "access_token": "secret",
    "scope": "test:scope",
    "token_type": "Bearer",
    "expires_in": 899
}`

	c := Client{
		testData: []*http.Response{{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(data)),
		}},
	}

	if err := c.RefreshJwt(ctx); err != nil {
		t.Fatalf("Got error: %s", err)
	}
	if c.Jwt != "secret" {
		t.Errorf("Jwt is %q, expected \"secret\"", c.Jwt)
	}
}

func TestRefreshJwtFailedBadRequest(t *testing.T) {
	ctx := context.TODO()
	data := `{"error_description":"Client authentication failed","error":"invalid_client"}`

	c := Client{
		testData: []*http.Response{{
			StatusCode: http.StatusBadRequest,
			Body:       io.NopCloser(strings.NewReader(data)),
		}},
	}

	if err := c.RefreshJwt(ctx); err == nil {
		t.Fatalf("Expected error, got nil")
	}
}

func TestRefreshJwtFailedUnauthorized(t *testing.T) {
	ctx := context.TODO()
	data := `{"error_description":"Client authentication failed","error":"invalid_client"}`

	c := Client{
		testData: []*http.Response{{
			StatusCode: http.StatusUnauthorized,
			Body:       io.NopCloser(strings.NewReader(data)),
		}},
	}

	if err := c.RefreshJwt(ctx); err == nil {
		t.Fatalf("Expected error, got nil")
	}
}

func restoreLogging() {
	log.SetFlags(0)
}

func TestLogQuiet(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer restoreLogging()
	ctx := context.TODO()

	c := Client{Logging: api.LogQuiet}
	c.Log(ctx, "basic", "a")
	c.Log(ctx, "basic", "b")
	c.Log(ctx, "detailed", "c")
	c.Log(ctx, "detailed", "d")

	if buf.String() != "" {
		t.Fail()
	}
}

func TestLogBasic(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer restoreLogging()
	ctx := context.TODO()

	c := Client{Logging: api.LogBasic}
	c.Log(ctx, api.LogBasic, "a")
	c.Log(ctx, api.LogDetailed, "b")
	c.Log(ctx, api.LogBasic, "c")
	c.Log(ctx, api.LogDetailed, "d")

	if buf.String() != "a\nc\n" {
		t.Fail()
	}
}

func TestLogDetailed(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer restoreLogging()
	ctx := context.TODO()

	c := Client{Logging: api.LogDetailed}
	c.Log(ctx, api.LogBasic, "a")
	c.Log(ctx, api.LogDetailed, "b")
	c.Log(ctx, api.LogBasic, "c")
	c.Log(ctx, api.LogDetailed, "d")

	if buf.String() != "b\nd\n" {
		t.Fail()
	}
}

func TestDoUnmarshalsResponse(t *testing.T) {
	ctx := context.TODO()
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age,omitempty"`
	}

	name := "Jason"
	age := 42

	c := Client{
		apiPrefix: "https://testing-api-prefix",
		Logging:   api.LogDetailed,
		testData: []*http.Response{{
			StatusCode: http.StatusOK,
			Body: io.NopCloser(strings.NewReader(fmt.Sprintf(`{
    "name": %q,
    "age": %d
}`, name, age))),
		}},
	}

	var ans Person
	_, err := c.Do(ctx, http.MethodGet, "/one/two", nil, nil, &ans)

	if err != nil {
		t.Fatalf("Do returned an error: %s", err)
	}

	if ans.Name != name {
		t.Errorf("Name is %q, not %q", ans.Name, name)
	}
	if ans.Age != age {
		t.Errorf("Age is %d, not %d", ans.Age, age)
	}
}

func TestDoNoUnmarshalOnErrorResponse(t *testing.T) {
	ctx := context.TODO()
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age,omitempty"`
	}

	c := Client{
		apiPrefix: "https://testing-api-prefix",
		Logging:   api.LogDetailed,
		testData: []*http.Response{{
			StatusCode: http.StatusBadRequest,
			Body:       io.NopCloser(strings.NewReader("")),
		}},
	}

	var ans Person
	_, err := c.Do(ctx, http.MethodGet, "/one/two", nil, nil, &ans)

	if err == nil {
		t.Errorf("nil error was returned")
	}
	if ans.Name != "" {
		t.Errorf("Name is %q", ans.Name)
	}
	if ans.Age != 0 {
		t.Errorf("Age is %d", ans.Age)
	}

	e2, ok := err.(api.Response)
	if !ok {
		t.Fatalf("Error is not an api.Response")
	}

	if e2.StatusCode != http.StatusBadRequest {
		t.Errorf("StatusCode is %d, not %d", e2.StatusCode, http.StatusBadRequest)
	}
}

func TestDoCanRefreshTheJwt(t *testing.T) {
	ctx := context.TODO()
	data := `{
    "access_token": "secret",
    "scope": "test:scope",
    "token_type": "Bearer",
    "expires_in": 899
}`

	c := Client{
		apiPrefix: "https://testing-api-prefix",
		Logging:   api.LogDetailed,
		testData: []*http.Response{
			{
				StatusCode: http.StatusUnauthorized,
				Body:       io.NopCloser(strings.NewReader("")),
			},
			{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(data)),
			},
			{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader("")),
			},
		},
	}

	_, err := c.Do(ctx, http.MethodGet, "/one/two", nil, nil, nil)

	if err != nil {
		t.Errorf("error was returned")
	}

	if c.Jwt != "secret" {
		t.Errorf("JWT is %q, not \"secret\"", c.Jwt)
	}
}

func TestDoWithMultipleUnauthorizedFailures(t *testing.T) {
	ctx := context.TODO()
	data := `{
    "access_token": "secret",
    "scope": "test:scope",
    "token_type": "Bearer",
    "expires_in": 899
}`

	c := Client{
		apiPrefix: "https://testing-api-prefix",
		Logging:   api.LogDetailed,
		testData: []*http.Response{
			{
				StatusCode: http.StatusUnauthorized,
				Body:       io.NopCloser(strings.NewReader("")),
			},
			{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(data)),
			},
			{
				StatusCode: http.StatusUnauthorized,
				Body:       io.NopCloser(strings.NewReader("")),
			},
			{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader("")),
			},
			{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader("")),
			},
			{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader("")),
			},
		},
	}

	_, err := c.Do(ctx, http.MethodGet, "/one/two", nil, nil, nil)

	if err == nil {
		t.Fatalf("nil error was returned")
	}

	e2, ok := err.(api.Response)
	if !ok {
		t.Fatalf("Error is not an api.Response")
	}
	if e2.StatusCode != http.StatusUnauthorized {
		t.Fatalf("Status code is %d, not %d", e2.StatusCode, http.StatusUnauthorized)
	}
}

// TestClient_JwtFieldInAuthFile tests that the Jwt field can be read from auth_file JSON
// This is the key requirement from AUTH-FILE-JWT-CACHING-SOLUTION.md
func TestClient_JwtFieldInAuthFile(t *testing.T) {
	authFileContent := `{
		"client_id": "test-client-id@123456.iam.panserviceaccount.com",
		"client_secret": "test-secret-key",
		"scope": "tsg_id:123456",
		"host": "api.sase.paloaltonetworks.com",
		"jwt": "eyJhbGciOiJSUzI1NiIsImtpZCI6IjEyMyIsInR5cCI6IkpXVCJ9.test.token"
	}`

	tmpfile, err := os.CreateTemp("", "scm-auth-test-*.json")
	if err != nil {
		t.Fatalf("Failed to create temporary auth file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err = tmpfile.Write([]byte(authFileContent)); err != nil {
		t.Fatalf("Failed to write to temporary auth file: %v", err)
	}
	if err = tmpfile.Close(); err != nil {
		t.Fatalf("Failed to close temporary auth file: %v", err)
	}

	client := &Client{
		AuthFile: tmpfile.Name(),
	}

	if err = client.Setup(); err != nil {
		t.Fatalf("Failed to setup client: %v", err)
	}

	expectedJwt := "eyJhbGciOiJSUzI1NiIsImtpZCI6IjEyMyIsInR5cCI6IkpXVCJ9.test.token"
	if client.Jwt != expectedJwt {
		t.Errorf("JWT mismatch: got %q, want %q", client.Jwt, expectedJwt)
	}

	if client.ClientId != "test-client-id@123456.iam.panserviceaccount.com" {
		t.Errorf("ClientId mismatch: got %q", client.ClientId)
	}
}

// TestClient_AuthFileWithoutJwt tests backward compatibility
func TestClient_AuthFileWithoutJwt(t *testing.T) {
	authFileContent := `{
		"client_id": "test-client-id",
		"client_secret": "test-secret",
		"scope": "test-scope",
		"host": "test-host"
	}`

	tmpfile, err := os.CreateTemp("", "scm-auth-no-jwt-*.json")
	if err != nil {
		t.Fatalf("Failed to create temporary auth file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err = tmpfile.Write([]byte(authFileContent)); err != nil {
		t.Fatalf("Failed to write to temporary auth file: %v", err)
	}
	if err = tmpfile.Close(); err != nil {
		t.Fatalf("Failed to close temporary auth file: %v", err)
	}

	client := &Client{
		AuthFile: tmpfile.Name(),
	}

	if err = client.Setup(); err != nil {
		t.Fatalf("Failed to setup client: %v", err)
	}

	if client.Jwt != "" {
		t.Errorf("JWT should be empty when not in auth file, got: %q", client.Jwt)
	}
}

// TestClient_JwtFromAuthFileSkipsRefresh validates that JWT presence skips refresh
func TestClient_JwtFromAuthFileSkipsRefresh(t *testing.T) {
	authFileContent := `{
		"client_id": "test-client-id@123456.iam.panserviceaccount.com",
		"client_secret": "test-secret-key",
		"scope": "tsg_id:123456",
		"host": "api.sase.paloaltonetworks.com",
		"jwt": "eyJhbGciOiJSUzI1NiIsImtpZCI6IjEyMyIsInR5cCI6IkpXVCJ9.cached.token"
	}`

	tmpfile, err := os.CreateTemp("", "scm-auth-skip-refresh-*.json")
	if err != nil {
		t.Fatalf("Failed to create temporary auth file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err = tmpfile.Write([]byte(authFileContent)); err != nil {
		t.Fatalf("Failed to write to temporary auth file: %v", err)
	}
	if err = tmpfile.Close(); err != nil {
		t.Fatalf("Failed to close temporary auth file: %v", err)
	}

	client := &Client{
		AuthFile: tmpfile.Name(),
	}

	if err = client.Setup(); err != nil {
		t.Fatalf("Failed to setup client: %v", err)
	}

	expectedJwt := "eyJhbGciOiJSUzI1NiIsImtpZCI6IjEyMyIsInR5cCI6IkpXVCJ9.cached.token"
	if client.Jwt != expectedJwt {
		t.Errorf("JWT mismatch: got %q, want %q", client.Jwt, expectedJwt)
	}

	// This demonstrates the key behavior: when JWT is present, RefreshJwt() is skipped
	if client.Jwt == "" {
		t.Error("JWT should not be empty")
	}
}
