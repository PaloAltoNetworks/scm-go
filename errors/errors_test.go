package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ============================================================================
// Constructor Tests
// ============================================================================

func TestNewNotAuthenticatedError(t *testing.T) {
	err := NewNotAuthenticatedError("test message")

	assert.Equal(t, 401, err.HTTPStatusCode())
	assert.Equal(t, "E003", err.ErrorCode())
	assert.Equal(t, "test message", err.ErrorMessage())
	assert.Equal(t, "[E003] test message", err.Error())
	assert.True(t, err.IsScmError())
}

func TestNewNotAuthenticatedError_DefaultMessage(t *testing.T) {
	err := NewNotAuthenticatedError("")

	assert.Equal(t, "Not authenticated", err.ErrorMessage())
}

func TestNewInvalidCredentialError(t *testing.T) {
	err := NewInvalidCredentialError("bad credentials")

	assert.Equal(t, 401, err.HTTPStatusCode())
	assert.Equal(t, "E004", err.ErrorCode())
	assert.Equal(t, "bad credentials", err.ErrorMessage())
}

func TestNewKeyExpiredError(t *testing.T) {
	err := NewKeyExpiredError("key expired")

	assert.Equal(t, 401, err.HTTPStatusCode())
	assert.Equal(t, "E013", err.ErrorCode())
	assert.Equal(t, "key expired", err.ErrorMessage())
}

func TestNewAuthorizationError(t *testing.T) {
	err := NewAuthorizationError("no permission")

	assert.Equal(t, 403, err.HTTPStatusCode())
	assert.Equal(t, "E009", err.ErrorCode())
	assert.Equal(t, "no permission", err.ErrorMessage())
}

func TestNewInvalidObjectError(t *testing.T) {
	err := NewInvalidObjectError("my-object", "validation failed")

	assert.Equal(t, 400, err.HTTPStatusCode())
	assert.Equal(t, "E023", err.ErrorCode())
	assert.Equal(t, "validation failed", err.ErrorMessage())
	assert.Equal(t, "my-object", err.ObjectName)
}

func TestNewMissingQueryParameterError(t *testing.T) {
	err := NewMissingQueryParameterError("folder")

	assert.Equal(t, 400, err.HTTPStatusCode())
	assert.Equal(t, "E007", err.ErrorCode())
	assert.Contains(t, err.ErrorMessage(), "folder")
	assert.Equal(t, "folder", err.ParameterName)
}

func TestNewInvalidQueryParameterError(t *testing.T) {
	err := NewInvalidQueryParameterError("limit", "abc", "must be integer")

	assert.Equal(t, 400, err.HTTPStatusCode())
	assert.Equal(t, "E024", err.ErrorCode())
	assert.Equal(t, "must be integer", err.ErrorMessage())
	assert.Equal(t, "limit", err.ParameterName)
	assert.Equal(t, "abc", err.ParameterValue)
}

func TestNewMalformedCommandError(t *testing.T) {
	err := NewMalformedCommandError("syntax error")

	assert.Equal(t, 400, err.HTTPStatusCode())
	assert.Equal(t, "E006", err.ErrorCode())
	assert.Equal(t, "syntax error", err.ErrorMessage())
}

func TestNewObjectNotPresentError(t *testing.T) {
	err := NewObjectNotPresentError("123", "web-server")

	assert.Equal(t, 404, err.HTTPStatusCode())
	assert.Equal(t, "E005", err.ErrorCode())
	assert.Contains(t, err.ErrorMessage(), "web-server")
	assert.Contains(t, err.ErrorMessage(), "123")
	assert.Equal(t, "123", err.ObjectID)
	assert.Equal(t, "web-server", err.ObjectName)
}

func TestNewObjectNotPresentError_NoID(t *testing.T) {
	err := NewObjectNotPresentError("", "web-server")

	assert.Contains(t, err.ErrorMessage(), "web-server")
	assert.NotContains(t, err.ErrorMessage(), "ID:")
}

func TestNewNameNotUniqueError(t *testing.T) {
	err := NewNameNotUniqueError("web-server")

	assert.Equal(t, 409, err.HTTPStatusCode())
	assert.Equal(t, "E016", err.ErrorCode())
	assert.Contains(t, err.ErrorMessage(), "web-server")
	assert.Contains(t, err.ErrorMessage(), "not unique")
	assert.Equal(t, "web-server", err.ObjectName)
}

func TestNewObjectNotUniqueError(t *testing.T) {
	err := NewObjectNotUniqueError("123", "web-server")

	assert.Equal(t, 409, err.HTTPStatusCode())
	assert.Equal(t, "E017", err.ErrorCode())
	assert.Contains(t, err.ErrorMessage(), "web-server")
	assert.Equal(t, "123", err.ObjectID)
	assert.Equal(t, "web-server", err.ObjectName)
}

func TestNewReferenceNotZeroError(t *testing.T) {
	err := NewReferenceNotZeroError("my-rule", 5)

	assert.Equal(t, 409, err.HTTPStatusCode())
	assert.Equal(t, "E018", err.ErrorCode())
	assert.Contains(t, err.ErrorMessage(), "my-rule")
	assert.Contains(t, err.ErrorMessage(), "5")
	assert.Equal(t, "my-rule", err.ObjectName)
	assert.Equal(t, 5, err.ReferenceCount)
}

func TestNewMethodNotAllowedError(t *testing.T) {
	err := NewMethodNotAllowedError("DELETE")

	assert.Equal(t, 405, err.HTTPStatusCode())
	assert.Equal(t, "E010", err.ErrorCode())
	assert.Contains(t, err.ErrorMessage(), "DELETE")
	assert.Equal(t, "DELETE", err.Method)
}

func TestNewRequestTimeoutError(t *testing.T) {
	err := NewRequestTimeoutError("timeout occurred")

	assert.Equal(t, 408, err.HTTPStatusCode())
	assert.Equal(t, "E011", err.ErrorCode())
	assert.Equal(t, "timeout occurred", err.ErrorMessage())
}

func TestNewTooManyRequestsError(t *testing.T) {
	err := NewTooManyRequestsError(60, "rate limited")

	assert.Equal(t, 429, err.HTTPStatusCode())
	assert.Equal(t, "E012", err.ErrorCode())
	assert.Equal(t, "rate limited", err.ErrorMessage())
	assert.Equal(t, 60, err.RetryAfter)
}

func TestNewSessionTimedOutError(t *testing.T) {
	err := NewSessionTimedOutError("session expired")

	assert.Equal(t, 401, err.HTTPStatusCode())
	assert.Equal(t, "E019", err.ErrorCode())
	assert.Equal(t, "session expired", err.ErrorMessage())
}

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("server crashed")

	assert.Equal(t, 500, err.HTTPStatusCode())
	assert.Equal(t, "E020", err.ErrorCode())
	assert.Equal(t, "server crashed", err.ErrorMessage())
}

func TestNewBadGatewayError(t *testing.T) {
	err := NewBadGatewayError("bad gateway")

	assert.Equal(t, 502, err.HTTPStatusCode())
	assert.Equal(t, "E021", err.ErrorCode())
	assert.Equal(t, "bad gateway", err.ErrorMessage())
}

func TestNewServiceUnavailableError(t *testing.T) {
	err := NewServiceUnavailableError("service down")

	assert.Equal(t, 503, err.HTTPStatusCode())
	assert.Equal(t, "E022", err.ErrorCode())
	assert.Equal(t, "service down", err.ErrorMessage())
}

func TestNewGatewayTimeoutError(t *testing.T) {
	err := NewGatewayTimeoutError("gateway timeout")

	assert.Equal(t, 504, err.HTTPStatusCode())
	assert.Equal(t, "E024", err.ErrorCode())
	assert.Equal(t, "gateway timeout", err.ErrorMessage())
}

// ============================================================================
// Helper Function Tests
// ============================================================================

func TestIsScmError(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected bool
	}{
		{"nil error", nil, false},
		{"NotAuthenticatedError", NewNotAuthenticatedError("test"), true},
		{"NameNotUniqueError", NewNameNotUniqueError("test"), true},
		{"InternalServerError", NewInternalServerError("test"), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, IsScmError(tt.err))
		})
	}
}

func TestIsClientError(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected bool
	}{
		{"nil", nil, false},
		{"NotAuthenticatedError", NewNotAuthenticatedError("test"), true},
		{"AuthorizationError", NewAuthorizationError("test"), true},
		{"NameNotUniqueError", NewNameNotUniqueError("test"), true},
		{"InternalServerError", NewInternalServerError("test"), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, IsClientError(tt.err))
		})
	}
}

func TestIsServerError(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected bool
	}{
		{"nil", nil, false},
		{"InternalServerError", NewInternalServerError("test"), true},
		{"BadGatewayError", NewBadGatewayError("test"), true},
		{"ServiceUnavailableError", NewServiceUnavailableError("test"), true},
		{"GatewayTimeoutError", NewGatewayTimeoutError("test"), true},
		{"NotAuthenticatedError", NewNotAuthenticatedError("test"), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, IsServerError(tt.err))
		})
	}
}

func TestIsAuthenticationError(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected bool
	}{
		{"nil", nil, false},
		{"NotAuthenticatedError", NewNotAuthenticatedError("test"), true},
		{"InvalidCredentialError", NewInvalidCredentialError("test"), true},
		{"KeyExpiredError", NewKeyExpiredError("test"), true},
		{"AuthorizationError", NewAuthorizationError("test"), false},
		{"NameNotUniqueError", NewNameNotUniqueError("test"), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, IsAuthenticationError(tt.err))
		})
	}
}

func TestIsNotAuthenticated(t *testing.T) {
	assert.True(t, IsNotAuthenticated(NewNotAuthenticatedError("test")))
	assert.False(t, IsNotAuthenticated(NewInvalidCredentialError("test")))
	assert.False(t, IsNotAuthenticated(nil))
}

func TestIsInvalidCredential(t *testing.T) {
	assert.True(t, IsInvalidCredential(NewInvalidCredentialError("test")))
	assert.False(t, IsInvalidCredential(NewNotAuthenticatedError("test")))
	assert.False(t, IsInvalidCredential(nil))
}

func TestIsKeyExpired(t *testing.T) {
	assert.True(t, IsKeyExpired(NewKeyExpiredError("test")))
	assert.False(t, IsKeyExpired(NewNotAuthenticatedError("test")))
	assert.False(t, IsKeyExpired(nil))
}

func TestIsAuthorizationError(t *testing.T) {
	assert.True(t, IsAuthorizationError(NewAuthorizationError("test")))
	assert.False(t, IsAuthorizationError(NewNotAuthenticatedError("test")))
	assert.False(t, IsAuthorizationError(nil))
}

func TestIsBadRequest(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected bool
	}{
		{"nil", nil, false},
		{"InvalidObjectError", NewInvalidObjectError("obj", "msg"), true},
		{"MissingQueryParameterError", NewMissingQueryParameterError("param"), true},
		{"InvalidQueryParameterError", NewInvalidQueryParameterError("p", "v", "m"), true},
		{"MalformedCommandError", NewMalformedCommandError("msg"), true},
		{"NotAuthenticatedError", NewNotAuthenticatedError("test"), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, IsBadRequest(tt.err))
		})
	}
}

func TestIsInvalidObject(t *testing.T) {
	assert.True(t, IsInvalidObject(NewInvalidObjectError("obj", "msg")))
	assert.False(t, IsInvalidObject(NewMissingQueryParameterError("param")))
	assert.False(t, IsInvalidObject(nil))
}

func TestIsMissingQueryParameter(t *testing.T) {
	assert.True(t, IsMissingQueryParameter(NewMissingQueryParameterError("param")))
	assert.False(t, IsMissingQueryParameter(NewInvalidObjectError("obj", "msg")))
	assert.False(t, IsMissingQueryParameter(nil))
}

func TestIsInvalidQueryParameter(t *testing.T) {
	assert.True(t, IsInvalidQueryParameter(NewInvalidQueryParameterError("p", "v", "m")))
	assert.False(t, IsInvalidQueryParameter(NewMissingQueryParameterError("param")))
	assert.False(t, IsInvalidQueryParameter(nil))
}

func TestIsMalformedCommand(t *testing.T) {
	assert.True(t, IsMalformedCommand(NewMalformedCommandError("msg")))
	assert.False(t, IsMalformedCommand(NewInvalidObjectError("obj", "msg")))
	assert.False(t, IsMalformedCommand(nil))
}

func TestIsNotFound(t *testing.T) {
	assert.True(t, IsNotFound(NewObjectNotPresentError("id", "name")))
	assert.False(t, IsNotFound(NewNameNotUniqueError("name")))
	assert.False(t, IsNotFound(nil))
}

func TestIsObjectNotPresent(t *testing.T) {
	assert.True(t, IsObjectNotPresent(NewObjectNotPresentError("id", "name")))
	assert.False(t, IsObjectNotPresent(NewNameNotUniqueError("name")))
	assert.False(t, IsObjectNotPresent(nil))
}

func TestIsConflict(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected bool
	}{
		{"nil", nil, false},
		{"NameNotUniqueError", NewNameNotUniqueError("name"), true},
		{"ObjectNotUniqueError", NewObjectNotUniqueError("id", "name"), true},
		{"ReferenceNotZeroError", NewReferenceNotZeroError("name", 5), true},
		{"ObjectNotPresentError", NewObjectNotPresentError("id", "name"), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, IsConflict(tt.err))
		})
	}
}

func TestIsNameNotUnique(t *testing.T) {
	assert.True(t, IsNameNotUnique(NewNameNotUniqueError("name")))
	assert.False(t, IsNameNotUnique(NewObjectNotUniqueError("id", "name")))
	assert.False(t, IsNameNotUnique(nil))
}

func TestIsObjectNotUnique(t *testing.T) {
	assert.True(t, IsObjectNotUnique(NewObjectNotUniqueError("id", "name")))
	assert.False(t, IsObjectNotUnique(NewNameNotUniqueError("name")))
	assert.False(t, IsObjectNotUnique(nil))
}

func TestIsReferenceNotZero(t *testing.T) {
	assert.True(t, IsReferenceNotZero(NewReferenceNotZeroError("name", 5)))
	assert.False(t, IsReferenceNotZero(NewNameNotUniqueError("name")))
	assert.False(t, IsReferenceNotZero(nil))
}

func TestIsMethodNotAllowed(t *testing.T) {
	assert.True(t, IsMethodNotAllowed(NewMethodNotAllowedError("DELETE")))
	assert.False(t, IsMethodNotAllowed(NewNotAuthenticatedError("test")))
	assert.False(t, IsMethodNotAllowed(nil))
}

func TestIsRequestTimeout(t *testing.T) {
	assert.True(t, IsRequestTimeout(NewRequestTimeoutError("msg")))
	assert.False(t, IsRequestTimeout(NewNotAuthenticatedError("test")))
	assert.False(t, IsRequestTimeout(nil))
}

func TestIsTooManyRequests(t *testing.T) {
	assert.True(t, IsTooManyRequests(NewTooManyRequestsError(60, "msg")))
	assert.False(t, IsTooManyRequests(NewNotAuthenticatedError("test")))
	assert.False(t, IsTooManyRequests(nil))
}

func TestIsSessionTimedOut(t *testing.T) {
	assert.True(t, IsSessionTimedOut(NewSessionTimedOutError("msg")))
	assert.False(t, IsSessionTimedOut(NewNotAuthenticatedError("test")))
	assert.False(t, IsSessionTimedOut(nil))
}

func TestIsInternalServerError(t *testing.T) {
	assert.True(t, IsInternalServerError(NewInternalServerError("msg")))
	assert.False(t, IsInternalServerError(NewBadGatewayError("msg")))
	assert.False(t, IsInternalServerError(nil))
}

func TestIsBadGateway(t *testing.T) {
	assert.True(t, IsBadGateway(NewBadGatewayError("msg")))
	assert.False(t, IsBadGateway(NewInternalServerError("msg")))
	assert.False(t, IsBadGateway(nil))
}

func TestIsServiceUnavailable(t *testing.T) {
	assert.True(t, IsServiceUnavailable(NewServiceUnavailableError("msg")))
	assert.False(t, IsServiceUnavailable(NewInternalServerError("msg")))
	assert.False(t, IsServiceUnavailable(nil))
}

func TestIsGatewayTimeout(t *testing.T) {
	assert.True(t, IsGatewayTimeout(NewGatewayTimeoutError("msg")))
	assert.False(t, IsGatewayTimeout(NewInternalServerError("msg")))
	assert.False(t, IsGatewayTimeout(nil))
}

// ============================================================================
// As* Helper Tests
// ============================================================================

func TestAsObjectNotPresent(t *testing.T) {
	err := NewObjectNotPresentError("123", "web-server")
	extracted, ok := AsObjectNotPresent(err)
	require.True(t, ok)
	assert.Equal(t, "123", extracted.ObjectID)
	assert.Equal(t, "web-server", extracted.ObjectName)

	_, ok = AsObjectNotPresent(NewNameNotUniqueError("name"))
	assert.False(t, ok)

	_, ok = AsObjectNotPresent(nil)
	assert.False(t, ok)
}

func TestAsNameNotUnique(t *testing.T) {
	err := NewNameNotUniqueError("web-server")
	extracted, ok := AsNameNotUnique(err)
	require.True(t, ok)
	assert.Equal(t, "web-server", extracted.ObjectName)

	_, ok = AsNameNotUnique(NewObjectNotPresentError("id", "name"))
	assert.False(t, ok)

	_, ok = AsNameNotUnique(nil)
	assert.False(t, ok)
}

func TestAsObjectNotUnique(t *testing.T) {
	err := NewObjectNotUniqueError("123", "web-server")
	extracted, ok := AsObjectNotUnique(err)
	require.True(t, ok)
	assert.Equal(t, "123", extracted.ObjectID)
	assert.Equal(t, "web-server", extracted.ObjectName)

	_, ok = AsObjectNotUnique(nil)
	assert.False(t, ok)
}

func TestAsReferenceNotZero(t *testing.T) {
	err := NewReferenceNotZeroError("my-rule", 5)
	extracted, ok := AsReferenceNotZero(err)
	require.True(t, ok)
	assert.Equal(t, "my-rule", extracted.ObjectName)
	assert.Equal(t, 5, extracted.ReferenceCount)

	_, ok = AsReferenceNotZero(nil)
	assert.False(t, ok)
}

func TestAsTooManyRequests(t *testing.T) {
	err := NewTooManyRequestsError(60, "rate limited")
	extracted, ok := AsTooManyRequests(err)
	require.True(t, ok)
	assert.Equal(t, 60, extracted.RetryAfter)

	_, ok = AsTooManyRequests(nil)
	assert.False(t, ok)
}

func TestAsInvalidObject(t *testing.T) {
	err := NewInvalidObjectError("my-object", "validation failed")
	extracted, ok := AsInvalidObject(err)
	require.True(t, ok)
	assert.Equal(t, "my-object", extracted.ObjectName)

	_, ok = AsInvalidObject(nil)
	assert.False(t, ok)
}

func TestAsMissingQueryParameter(t *testing.T) {
	err := NewMissingQueryParameterError("folder")
	extracted, ok := AsMissingQueryParameter(err)
	require.True(t, ok)
	assert.Equal(t, "folder", extracted.ParameterName)

	_, ok = AsMissingQueryParameter(nil)
	assert.False(t, ok)
}

func TestAsInvalidQueryParameter(t *testing.T) {
	err := NewInvalidQueryParameterError("limit", "abc", "must be integer")
	extracted, ok := AsInvalidQueryParameter(err)
	require.True(t, ok)
	assert.Equal(t, "limit", extracted.ParameterName)
	assert.Equal(t, "abc", extracted.ParameterValue)

	_, ok = AsInvalidQueryParameter(nil)
	assert.False(t, ok)
}

func TestAsMethodNotAllowed(t *testing.T) {
	err := NewMethodNotAllowedError("DELETE")
	extracted, ok := AsMethodNotAllowed(err)
	require.True(t, ok)
	assert.Equal(t, "DELETE", extracted.Method)

	_, ok = AsMethodNotAllowed(nil)
	assert.False(t, ok)
}

// ============================================================================
// Error Interface Tests
// ============================================================================

func TestBaseError_ErrorDetails(t *testing.T) {
	// Test with nil details
	err := &BaseError{
		StatusCode: 400,
		Code:       "E001",
		Message:    "test",
		Details:    nil,
	}
	details := err.ErrorDetails()
	assert.NotNil(t, details)
	assert.Empty(t, details)

	// Test with populated details
	err.Details = map[string]interface{}{
		"key1": "value1",
		"key2": 42,
	}
	details = err.ErrorDetails()
	assert.Equal(t, "value1", details["key1"])
	assert.Equal(t, 42, details["key2"])
}

func TestErrorMessage_WithoutCode(t *testing.T) {
	err := &BaseError{
		StatusCode: 400,
		Code:       "",
		Message:    "test message",
	}
	assert.Equal(t, "test message", err.Error())
}

func TestErrorMessage_WithCode(t *testing.T) {
	err := &BaseError{
		StatusCode: 400,
		Code:       "E001",
		Message:    "test message",
	}
	assert.Equal(t, "[E001] test message", err.Error())
}

// ============================================================================
// Real-World Usage Examples
// ============================================================================

func TestRealWorldUsageExample(t *testing.T) {
	// Simulate an API error
	var err error = NewNameNotUniqueError("web-server")

	// Check if it's a name conflict
	if IsNameNotUnique(err) {
		// Extract typed error
		nameErr, ok := AsNameNotUnique(err)
		require.True(t, ok)
		assert.Equal(t, "web-server", nameErr.ObjectName)
		assert.Equal(t, 409, nameErr.HTTPStatusCode())
		assert.Equal(t, "E016", nameErr.ErrorCode())
	} else {
		t.Fatal("Expected NameNotUniqueError")
	}
}

func TestRealWorldUsageExample_ObjectNotFound(t *testing.T) {
	// Simulate a not found error
	var err error = NewObjectNotPresentError("12345", "web-server")

	// Check various error types
	assert.False(t, IsNameNotUnique(err))
	assert.True(t, IsObjectNotPresent(err))
	assert.True(t, IsNotFound(err))
	assert.True(t, IsClientError(err))
	assert.False(t, IsServerError(err))

	// Extract object details
	objErr, ok := AsObjectNotPresent(err)
	require.True(t, ok)
	assert.Equal(t, "12345", objErr.ObjectID)
	assert.Equal(t, "web-server", objErr.ObjectName)
}

func TestRealWorldUsageExample_RateLimiting(t *testing.T) {
	// Simulate rate limiting
	var err error = NewTooManyRequestsError(60, "Too many requests")

	if IsTooManyRequests(err) {
		rateErr, ok := AsTooManyRequests(err)
		require.True(t, ok)
		// Application can use RetryAfter
		assert.Equal(t, 60, rateErr.RetryAfter)
	}
}
