package test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"

	scm "github.com/paloaltonetworks/scm-go"
)

// Config represents the auth configuration
type Config struct {
	ClientID     string    `json:"client_id"`
	ClientSecret string    `json:"client_secret"`
	Host         string    `json:"host"`
	AuthURL      string    `json:"auth_url"`
	Protocol     string    `json:"protocol"`
	Scope        string    `json:"scope"`
	Logging      string    `json:"logging"`
	JWT          string    `json:"jwt,omitempty"`
	JWTExpiresAt time.Time `json:"jwt_expires_at,omitempty"`
	JWTLifetime  int64     `json:"jwt_lifetime,omitempty"`
}

func saveConfigAtomic(path string, config *Config) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	tmpPath := path + ".tmp"
	if err := os.WriteFile(tmpPath, data, 0600); err != nil {
		return err
	}
	return os.Rename(tmpPath, path)
}

// TestJWTCaching_FreshToken tests loading a fresh JWT from auth file
func TestJWTCaching_FreshToken(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	// Create temporary auth file
	tmpDir := t.TempDir()
	authFile := filepath.Join(tmpDir, "auth-config.json")

	// Create config with fresh JWT (expires in 15 minutes)
	config := &Config{
		ClientID:     "test-client-id",
		ClientSecret: "test-client-secret",
		Host:         "api.test.paloaltonetworks.com",
		AuthURL:      "https://auth.test.paloaltonetworks.com/auth/v1/oauth2/access_token",
		Protocol:     "https",
		Scope:        "tsg_id:123456789",
		Logging:      "quiet",
		JWT:          "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.test_token",
		JWTExpiresAt: time.Now().Add(15 * time.Minute),
		JWTLifetime:  900,
	}

	if err := saveConfigAtomic(authFile, config); err != nil {
		t.Fatalf("Failed to save config: %v", err)
	}

	// Create client with auth file
	client := &scm.Client{
		AuthFile:         authFile,
		CheckEnvironment: false,
	}

	if err := client.Setup(); err != nil {
		t.Fatalf("Setup failed: %v", err)
	}

	// Verify JWT was loaded
	if client.Jwt == "" {
		t.Error("Expected JWT to be loaded from auth file")
	}

	if client.Jwt != config.JWT {
		t.Errorf("Expected JWT %s, got %s", config.JWT, client.Jwt)
	}

	// Verify expiration was loaded
	if client.JwtExpiresAt.IsZero() {
		t.Error("Expected JwtExpiresAt to be set")
	}

	// Verify token is not expired
	if time.Now().After(client.JwtExpiresAt) {
		t.Error("Token should not be expired")
	}

	// Verify lifetime was loaded
	if client.JwtLifetime != config.JWTLifetime {
		t.Errorf("Expected JwtLifetime %d, got %d", config.JWTLifetime, client.JwtLifetime)
	}
}

// TestJWTCaching_ExpiredToken tests handling of expired JWT
func TestJWTCaching_ExpiredToken(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	tmpDir := t.TempDir()
	authFile := filepath.Join(tmpDir, "auth-config.json")

	// Create config with expired JWT
	config := &Config{
		ClientID:     "test-client-id",
		ClientSecret: "test-client-secret",
		Host:         "api.test.paloaltonetworks.com",
		AuthURL:      "https://auth.test.paloaltonetworks.com/auth/v1/oauth2/access_token",
		Protocol:     "https",
		Scope:        "tsg_id:123456789",
		Logging:      "quiet",
		JWT:          "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.expired_token",
		JWTExpiresAt: time.Now().Add(-1 * time.Hour), // Expired 1 hour ago
		JWTLifetime:  900,
	}

	if err := saveConfigAtomic(authFile, config); err != nil {
		t.Fatalf("Failed to save config: %v", err)
	}

	client := &scm.Client{
		AuthFile:         authFile,
		CheckEnvironment: false,
	}

	if err := client.Setup(); err != nil {
		t.Fatalf("Setup failed: %v", err)
	}

	// Verify JWT was loaded even though expired
	if client.Jwt == "" {
		t.Error("Expected JWT to be loaded from auth file")
	}

	// Verify expiration detection
	if !time.Now().After(client.JwtExpiresAt) {
		t.Error("Token should be detected as expired")
	}
}

// TestJWTCaching_MissingJWT tests behavior when JWT fields are missing
func TestJWTCaching_MissingJWT(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	tmpDir := t.TempDir()
	authFile := filepath.Join(tmpDir, "auth-config.json")

	// Create config WITHOUT JWT fields
	config := &Config{
		ClientID:     "test-client-id",
		ClientSecret: "test-client-secret",
		Host:         "api.test.paloaltonetworks.com",
		AuthURL:      "https://auth.test.paloaltonetworks.com/auth/v1/oauth2/access_token",
		Protocol:     "https",
		Scope:        "tsg_id:123456789",
		Logging:      "quiet",
	}

	if err := saveConfigAtomic(authFile, config); err != nil {
		t.Fatalf("Failed to save config: %v", err)
	}

	client := &scm.Client{
		AuthFile:         authFile,
		CheckEnvironment: false,
	}

	if err := client.Setup(); err != nil {
		t.Fatalf("Setup failed: %v", err)
	}

	// Verify JWT is empty (will need to be fetched)
	if client.Jwt != "" {
		t.Error("Expected JWT to be empty when not in auth file")
	}

	// Verify expiration is zero
	if !client.JwtExpiresAt.IsZero() {
		t.Error("Expected JwtExpiresAt to be zero when not in auth file")
	}
}

// TestJWTCaching_ConcurrentProcesses tests multiple processes reading cached JWT
func TestJWTCaching_ConcurrentProcesses(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	tmpDir := t.TempDir()
	authFile := filepath.Join(tmpDir, "auth-config.json")

	// Create config with fresh JWT
	config := &Config{
		ClientID:     "test-client-id",
		ClientSecret: "test-client-secret",
		Host:         "api.test.paloaltonetworks.com",
		AuthURL:      "https://auth.test.paloaltonetworks.com/auth/v1/oauth2/access_token",
		Protocol:     "https",
		Scope:        "tsg_id:123456789",
		Logging:      "quiet",
		JWT:          "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.concurrent_test_token",
		JWTExpiresAt: time.Now().Add(15 * time.Minute),
		JWTLifetime:  900,
	}

	if err := saveConfigAtomic(authFile, config); err != nil {
		t.Fatalf("Failed to save config: %v", err)
	}

	// Launch 10 concurrent clients
	const numClients = 10
	var wg sync.WaitGroup
	errors := make(chan error, numClients)
	tokens := make(chan string, numClients)

	for i := 0; i < numClients; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			client := &scm.Client{
				AuthFile:         authFile,
				CheckEnvironment: false,
			}

			if err := client.Setup(); err != nil {
				errors <- err
				return
			}

			if client.Jwt == "" {
				return
			}

			tokens <- client.Jwt
		}(i)
	}

	wg.Wait()
	close(errors)
	close(tokens)

	// Check for errors
	var errs []error
	for err := range errors {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		t.Fatalf("Expected no errors, got %d errors: %v", len(errs), errs[0])
	}

	// Verify all clients got the same token
	tokenMap := make(map[string]int)
	for token := range tokens {
		tokenMap[token]++
	}

	if len(tokenMap) != 1 {
		t.Errorf("Expected all clients to get same token, got %d different tokens", len(tokenMap))
	}

	for token, count := range tokenMap {
		if count != numClients {
			t.Errorf("Expected %d clients with token %s, got %d", numClients, token, count)
		}
		if token != config.JWT {
			t.Errorf("Expected token %s, got %s", config.JWT, token)
		}
	}
}

// TestJWTCaching_InvalidAuthFile tests behavior with invalid auth file
func TestJWTCaching_InvalidAuthFile(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	tmpDir := t.TempDir()
	authFile := filepath.Join(tmpDir, "invalid-auth.json")

	// Write invalid JSON
	if err := os.WriteFile(authFile, []byte("invalid json{{{{{{"), 0600); err != nil {
		t.Fatalf("Failed to write invalid JSON: %v", err)
	}

	client := &scm.Client{
		AuthFile:         authFile,
		CheckEnvironment: false,
	}

	err := client.Setup()
	if err == nil {
		t.Error("Expected error with invalid JSON, got nil")
	}
}

// TestJWTCaching_MissingAuthFile tests behavior when auth file doesn't exist
func TestJWTCaching_MissingAuthFile(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	tmpDir := t.TempDir()
	authFile := filepath.Join(tmpDir, "nonexistent-auth.json")

	client := &scm.Client{
		AuthFile:         authFile,
		CheckEnvironment: false,
	}

	err := client.Setup()
	if err == nil {
		t.Error("Expected error with missing auth file, got nil")
	}
}

// TestJWTCaching_EmptyAuthFile tests behavior with empty auth file
func TestJWTCaching_EmptyAuthFile(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	tmpDir := t.TempDir()
	authFile := filepath.Join(tmpDir, "empty-auth.json")

	// Write empty file
	if err := os.WriteFile(authFile, []byte(""), 0600); err != nil {
		t.Fatalf("Failed to write empty file: %v", err)
	}

	client := &scm.Client{
		AuthFile:         authFile,
		CheckEnvironment: false,
	}

	err := client.Setup()
	if err == nil {
		t.Error("Expected error with empty auth file, got nil")
	}
}

// TestJWTCaching_PartialConfig tests behavior with partial config (missing required fields)
func TestJWTCaching_PartialConfig(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	tmpDir := t.TempDir()
	authFile := filepath.Join(tmpDir, "partial-auth.json")

	// Create config missing required fields
	partialConfig := map[string]interface{}{
		"host":     "api.test.paloaltonetworks.com",
		"protocol": "https",
		// Missing: client_id, client_secret, scope
	}

	data, _ := json.MarshalIndent(partialConfig, "", "  ")
	if err := os.WriteFile(authFile, data, 0600); err != nil {
		t.Fatalf("Failed to write partial config: %v", err)
	}

	client := &scm.Client{
		AuthFile:         authFile,
		CheckEnvironment: false,
	}

	err := client.Setup()
	if err == nil {
		t.Error("Expected error with partial config, got nil")
	}
}

// TestJWTCaching_TokenExpiringSoon tests detection of tokens expiring soon
func TestJWTCaching_TokenExpiringSoon(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	tmpDir := t.TempDir()
	authFile := filepath.Join(tmpDir, "auth-config.json")

	// Create config with token expiring in 30 seconds
	config := &Config{
		ClientID:     "test-client-id",
		ClientSecret: "test-client-secret",
		Host:         "api.test.paloaltonetworks.com",
		AuthURL:      "https://auth.test.paloaltonetworks.com/auth/v1/oauth2/access_token",
		Protocol:     "https",
		Scope:        "tsg_id:123456789",
		Logging:      "quiet",
		JWT:          "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.expiring_soon_token",
		JWTExpiresAt: time.Now().Add(30 * time.Second),
		JWTLifetime:  900,
	}

	if err := saveConfigAtomic(authFile, config); err != nil {
		t.Fatalf("Failed to save config: %v", err)
	}

	client := &scm.Client{
		AuthFile:         authFile,
		CheckEnvironment: false,
	}

	if err := client.Setup(); err != nil {
		t.Fatalf("Setup failed: %v", err)
	}

	// Token should be loaded and not yet expired
	if client.Jwt == "" {
		t.Error("Expected JWT to be loaded")
	}

	if time.Now().After(client.JwtExpiresAt) {
		t.Error("Token should not be expired yet")
	}

	// But it's expiring soon (within 60 seconds)
	if time.Now().After(client.JwtExpiresAt.Add(-60 * time.Second)) {
		t.Log("Token is expiring soon (within 60 seconds) - should be refreshed proactively")
	}
}
