# JWT Token Caching Tests

This directory contains comprehensive tests for JWT token caching functionality in the scm-go SDK.

## Test Files

### jwt_caching_test.go

Unit and integration tests for JWT token caching from auth files.

**Test Coverage:**

#### Positive Test Cases ✅

1. **TestJWTCaching_FreshToken**
   - Loads a fresh JWT token from auth file
   - Verifies token, expiration, and lifetime are correctly loaded
   - Ensures token is not marked as expired

2. **TestJWTCaching_ConcurrentProcesses**
   - Simulates 10 concurrent clients reading the same auth file
   - Verifies all clients receive the same cached token
   - Validates concurrent read access is safe

3. **TestJWTCaching_TokenExpiringSoon**
   - Tests detection of tokens expiring within 60 seconds
   - Validates tokens can be loaded even when expiring soon

#### Negative Test Cases ❌

1. **TestJWTCaching_ExpiredToken**
   - Loads an expired JWT token (expired 1 hour ago)
   - Verifies token is loaded but marked as expired
   - Tests expiration detection logic

2. **TestJWTCaching_MissingJWT**
   - Tests config file without JWT fields
   - Verifies client handles missing JWT gracefully
   - Token should be empty and will need to be fetched

3. **TestJWTCaching_InvalidAuthFile**
   - Tests behavior with malformed JSON in auth file
   - Expects Setup() to return an error

4. **TestJWTCaching_MissingAuthFile**
   - Tests behavior when auth file doesn't exist
   - Expects Setup() to return an error

5. **TestJWTCaching_EmptyAuthFile**
   - Tests behavior with empty auth file
   - Expects Setup() to return an error

6. **TestJWTCaching_PartialConfig**
   - Tests config missing required fields (client_id, client_secret, scope)
   - Expects Setup() to return an error

## Running the Tests

### Run all tests
```bash
go test -v ./auth/test/
```

### Run specific test
```bash
go test -v ./auth/test/ -run TestJWTCaching_FreshToken
```

### Run with short mode (skip integration tests)
```bash
go test -v -short ./auth/test/
```

### Run with race detector
```bash
go test -v -race ./auth/test/
```

## Test Requirements

These are **unit tests** that do not require:
- Live SCM API access
- Real credentials
- Network connectivity

The tests use:
- Temporary directories for auth files
- Mock JWT tokens
- Local file operations only

## Auth File Format

Tests validate this JSON structure:

```json
{
  "client_id": "your-client-id",
  "client_secret": "your-client-secret",
  "host": "api.strata.paloaltonetworks.com",
  "auth_url": "https://auth.apps.paloaltonetworks.com/auth/v1/oauth2/access_token",
  "protocol": "https",
  "scope": "tsg_id:1234567890",
  "logging": "quiet",
  "jwt": "eyJ0eXAi...",
  "jwt_expires_at": "2026-01-21T10:30:00Z",
  "jwt_lifetime": 900
}
```

## Integration Testing

For **integration tests** with live API:
1. Set up `scm-config.json` with real credentials
2. Run the actual token caching test programs in the examples directory
3. These tests only validate the **file handling and parsing** logic

## Expected Test Results

All tests should PASS:

```
=== RUN   TestJWTCaching_FreshToken
--- PASS: TestJWTCaching_FreshToken (0.00s)
=== RUN   TestJWTCaching_ExpiredToken
--- PASS: TestJWTCaching_ExpiredToken (0.00s)
=== RUN   TestJWTCaching_MissingJWT
--- PASS: TestJWTCaching_MissingJWT (0.00s)
=== RUN   TestJWTCaching_ConcurrentProcesses
--- PASS: TestJWTCaching_ConcurrentProcesses (0.01s)
=== RUN   TestJWTCaching_InvalidAuthFile
--- PASS: TestJWTCaching_InvalidAuthFile (0.00s)
=== RUN   TestJWTCaching_MissingAuthFile
--- PASS: TestJWTCaching_MissingAuthFile (0.00s)
=== RUN   TestJWTCaching_EmptyAuthFile
--- PASS: TestJWTCaching_EmptyAuthFile (0.00s)
=== RUN   TestJWTCaching_PartialConfig
--- PASS: TestJWTCaching_PartialConfig (0.00s)
=== RUN   TestJWTCaching_TokenExpiringSoon
--- PASS: TestJWTCaching_TokenExpiringSoon (0.00s)
PASS
```

## Continuous Integration

These tests are designed to run in CI/CD pipelines:
- No external dependencies
- Fast execution (< 1 second total)
- Deterministic results
- No flaky tests

Add to your CI workflow:
```yaml
- name: Run JWT Caching Tests
  run: go test -v -race ./auth/test/
```
