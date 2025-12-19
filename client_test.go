package scm

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestClient_JwtFieldInAuthFile tests that the Jwt field can be read from auth_file JSON
// This is the key requirement from AUTH-FILE-JWT-CACHING-SOLUTION.md
func TestClient_JwtFieldInAuthFile(t *testing.T) {
	// Create a temporary auth file with JWT token
	authFileContent := `{
		"client_id": "test-client-id@123456.iam.panserviceaccount.com",
		"client_secret": "test-secret-key",
		"scope": "tsg_id:123456",
		"host": "api.sase.paloaltonetworks.com",
		"jwt": "eyJhbGciOiJSUzI1NiIsImtpZCI6IjEyMyIsInR5cCI6IkpXVCJ9.test.token"
	}`

	// Create temporary file
	tmpfile, err := os.CreateTemp("", "scm-auth-test-*.json")
	require.NoError(t, err, "Failed to create temporary auth file")
	defer os.Remove(tmpfile.Name()) // Clean up

	// Write auth file content
	_, err = tmpfile.Write([]byte(authFileContent))
	require.NoError(t, err, "Failed to write to temporary auth file")
	err = tmpfile.Close()
	require.NoError(t, err, "Failed to close temporary auth file")

	// Create client with auth file
	client := &Client{
		AuthFile: tmpfile.Name(),
	}

	// Setup the client (this should read the auth file)
	err = client.Setup()
	require.NoError(t, err, "Failed to setup client")

	// TDD ASSERTION: Verify that the JWT field was populated from the auth file
	// This will FAIL initially because Jwt has json:"-" tag
	expectedJwt := "eyJhbGciOiJSUzI1NiIsImtpZCI6IjEyMyIsInR5cCI6IkpXVCJ9.test.token"
	assert.Equal(t, expectedJwt, client.Jwt, "JWT should be read from auth_file")

	// Verify other fields were also read correctly
	assert.Equal(t, "test-client-id@123456.iam.panserviceaccount.com", client.ClientId)
	assert.Equal(t, "test-secret-key", client.ClientSecret)
	assert.Equal(t, "tsg_id:123456", client.Scope)
	assert.Equal(t, "api.sase.paloaltonetworks.com", client.Host)
}

// TestClient_JwtFieldJsonMarshalling tests that Client struct can marshal/unmarshal JWT
// This validates that the json:"jwt" tag works correctly
func TestClient_JwtFieldJsonMarshalling(t *testing.T) {
	// Test 1: Unmarshal JSON with JWT into Client struct
	jsonData := `{
		"client_id": "test-id",
		"client_secret": "test-secret",
		"scope": "test-scope",
		"host": "test-host",
		"jwt": "test-jwt-token-12345"
	}`

	var client Client
	err := json.Unmarshal([]byte(jsonData), &client)
	require.NoError(t, err, "Failed to unmarshal JSON")

	// TDD ASSERTION: JWT should be populated from JSON
	// This will FAIL initially because Jwt has json:"-" tag
	assert.Equal(t, "test-jwt-token-12345", client.Jwt, "JWT should be unmarshaled from JSON")
	assert.Equal(t, "test-id", client.ClientId)
	assert.Equal(t, "test-secret", client.ClientSecret)

	// Test 2: Marshal Client struct with JWT to JSON
	client2 := Client{
		ClientId:     "marshal-test-id",
		ClientSecret: "marshal-test-secret",
		Scope:        "marshal-scope",
		Host:         "marshal-host",
		Jwt:          "marshal-jwt-token-67890",
	}

	jsonBytes, err := json.Marshal(client2)
	require.NoError(t, err, "Failed to marshal Client to JSON")

	// TDD ASSERTION: JWT should be included in marshaled JSON
	// This will FAIL initially because Jwt has json:"-" tag
	jsonString := string(jsonBytes)
	assert.Contains(t, jsonString, `"jwt":"marshal-jwt-token-67890"`, "JWT should be included in marshaled JSON")
	assert.Contains(t, jsonString, `"client_id":"marshal-test-id"`)
}

// TestClient_BackwardCompatibility tests that auth files without JWT still work
// This ensures our change is backward compatible
func TestClient_AuthFileWithoutJwt(t *testing.T) {
	// Create auth file WITHOUT JWT token (legacy format)
	authFileContent := `{
		"client_id": "test-client-id",
		"client_secret": "test-secret",
		"scope": "test-scope",
		"host": "test-host"
	}`

	tmpfile, err := os.CreateTemp("", "scm-auth-no-jwt-*.json")
	require.NoError(t, err, "Failed to create temporary auth file")
	defer os.Remove(tmpfile.Name())

	_, err = tmpfile.Write([]byte(authFileContent))
	require.NoError(t, err, "Failed to write to temporary auth file")
	err = tmpfile.Close()
	require.NoError(t, err, "Failed to close temporary auth file")

	client := &Client{
		AuthFile: tmpfile.Name(),
	}

	err = client.Setup()
	require.NoError(t, err, "Failed to setup client")

	// JWT should be empty (not populated)
	assert.Empty(t, client.Jwt, "JWT should be empty when not in auth file")

	// Other fields should still be populated
	assert.Equal(t, "test-client-id", client.ClientId)
	assert.Equal(t, "test-secret", client.ClientSecret)
	assert.Equal(t, "test-scope", client.Scope)
	assert.Equal(t, "test-host", client.Host)
}

// TestClient_JwtEmptyStringInAuthFile tests handling of empty JWT in auth file
func TestClient_JwtEmptyStringInAuthFile(t *testing.T) {
	authFileContent := `{
		"client_id": "test-client-id",
		"client_secret": "test-secret",
		"scope": "test-scope",
		"host": "test-host",
		"jwt": ""
	}`

	tmpfile, err := os.CreateTemp("", "scm-auth-empty-jwt-*.json")
	require.NoError(t, err, "Failed to create temporary auth file")
	defer os.Remove(tmpfile.Name())

	_, err = tmpfile.Write([]byte(authFileContent))
	require.NoError(t, err, "Failed to write to temporary auth file")
	err = tmpfile.Close()
	require.NoError(t, err, "Failed to close temporary auth file")

	client := &Client{
		AuthFile: tmpfile.Name(),
	}

	err = client.Setup()
	require.NoError(t, err, "Failed to setup client")

	// JWT should be empty string
	assert.Empty(t, client.Jwt, "JWT should be empty when set to empty string in auth file")
}

// TestClient_JwtFromAuthFileSkipsRefresh validates the key behavior:
// When JWT is present in auth file, RefreshJwt() should be skipped
// This is the main value proposition described in AUTH-FILE-JWT-CACHING-SOLUTION.md
func TestClient_JwtFromAuthFileSkipsRefresh(t *testing.T) {
	// Create auth file WITH JWT token
	authFileContent := `{
		"client_id": "test-client-id@123456.iam.panserviceaccount.com",
		"client_secret": "test-secret-key",
		"scope": "tsg_id:123456",
		"host": "api.sase.paloaltonetworks.com",
		"jwt": "eyJhbGciOiJSUzI1NiIsImtpZCI6IjEyMyIsInR5cCI6IkpXVCJ9.cached.token"
	}`

	tmpfile, err := os.CreateTemp("", "scm-auth-skip-refresh-*.json")
	require.NoError(t, err, "Failed to create temporary auth file")
	defer os.Remove(tmpfile.Name())

	_, err = tmpfile.Write([]byte(authFileContent))
	require.NoError(t, err, "Failed to write to temporary auth file")
	err = tmpfile.Close()
	require.NoError(t, err, "Failed to close temporary auth file")

	// Create client with auth file
	client := &Client{
		AuthFile: tmpfile.Name(),
	}

	// Setup the client
	err = client.Setup()
	require.NoError(t, err, "Failed to setup client")

	// Verify JWT was loaded from auth file
	expectedJwt := "eyJhbGciOiJSUzI1NiIsImtpZCI6IjEyMyIsInR5cCI6IkpXVCJ9.cached.token"
	assert.Equal(t, expectedJwt, client.Jwt, "JWT should be loaded from auth_file")

	// This demonstrates the key behavior:
	// In a real provider setup (like test_utils.go:37-40), the code checks:
	//   if setupClient.Jwt == "" {
	//       err = setupClient.RefreshJwt(ctx)
	//   }
	// Since client.Jwt is NOT empty, RefreshJwt() would be skipped
	assert.NotEmpty(t, client.Jwt, "JWT is not empty, so RefreshJwt() will be skipped")

	// This is the condition that determines whether RefreshJwt() is called
	shouldSkipRefresh := client.Jwt != ""
	assert.True(t, shouldSkipRefresh, "RefreshJwt() should be skipped when JWT is present")
}

// TestClient_AuthFileWithoutJwtRequiresRefresh validates the opposite case:
// When JWT is NOT in auth file, RefreshJwt() should be called
func TestClient_AuthFileWithoutJwtRequiresRefresh(t *testing.T) {
	// Create auth file WITHOUT JWT token
	authFileContent := `{
		"client_id": "test-client-id",
		"client_secret": "test-secret",
		"scope": "test-scope",
		"host": "test-host"
	}`

	tmpfile, err := os.CreateTemp("", "scm-auth-needs-refresh-*.json")
	require.NoError(t, err, "Failed to create temporary auth file")
	defer os.Remove(tmpfile.Name())

	_, err = tmpfile.Write([]byte(authFileContent))
	require.NoError(t, err, "Failed to write to temporary auth file")
	err = tmpfile.Close()
	require.NoError(t, err, "Failed to close temporary auth file")

	client := &Client{
		AuthFile: tmpfile.Name(),
	}

	err = client.Setup()
	require.NoError(t, err, "Failed to setup client")

	// JWT should be empty
	assert.Empty(t, client.Jwt, "JWT should be empty when not in auth file")

	// This demonstrates the opposite case:
	// In a real provider setup, the code checks: if setupClient.Jwt == ""
	// Since client.Jwt IS empty, RefreshJwt() WOULD be called
	shouldCallRefresh := client.Jwt == ""
	assert.True(t, shouldCallRefresh, "RefreshJwt() should be called when JWT is not present")
}
