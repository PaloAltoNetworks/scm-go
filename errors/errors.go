// Package errors provides typed error types for the SCM Go SDK.
//
// This package defines a comprehensive hierarchy of error types that match
// the Palo Alto Networks Strata Cloud Manager API error responses. All errors
// implement the ScmError interface and provide rich context about what went wrong.
//
// Error Hierarchy:
//
//	ScmError (interface)
//	├── ClientError (4xx errors)
//	│   ├── AuthenticationError (401)
//	│   │   ├── NotAuthenticatedError
//	│   │   ├── InvalidCredentialError
//	│   │   └── KeyExpiredError
//	│   ├── AuthorizationError (403)
//	│   ├── BadRequestError (400)
//	│   │   ├── InvalidObjectError
//	│   │   ├── MissingQueryParameterError
//	│   │   ├── InvalidQueryParameterError
//	│   │   └── MalformedCommandError
//	│   ├── NotFoundError (404)
//	│   │   └── ObjectNotPresentError
//	│   ├── ConflictError (409)
//	│   │   ├── NameNotUniqueError
//	│   │   ├── ObjectNotUniqueError
//	│   │   └── ReferenceNotZeroError
//	│   ├── MethodNotAllowedError (405)
//	│   ├── RequestTimeoutError (408)
//	│   ├── TooManyRequestsError (429)
//	│   └── SessionTimedOutError
//	└── ServerError (5xx errors)
//	    ├── InternalServerError (500)
//	    ├── BadGatewayError (502)
//	    ├── ServiceUnavailableError (503)
//	    └── GatewayTimeoutError (504)
//
// Usage Example:
//
//	import (
//	    scmErrors "github.com/paloaltonetworks/scm-go/errors"
//	)
//
//	resp, _, err := api.CreateAddresses(ctx).Addresses(addr).Execute()
//	if err != nil {
//	    // Check for specific error types
//	    if scmErrors.IsNameNotUnique(err) {
//	        nameErr := err.(*scmErrors.NameNotUniqueError)
//	        log.Printf("Name conflict: %s", nameErr.ObjectName)
//	        return handleNameConflict(nameErr)
//	    }
//
//	    if scmErrors.IsObjectNotPresent(err) {
//	        log.Printf("Object not found")
//	        return nil
//	    }
//
//	    // Generic error handling
//	    return err
//	}
package errors

import "fmt"

// ScmError is the base interface for all SCM SDK errors.
// All typed errors in this package implement this interface.
type ScmError interface {
	error
	// IsScmError identifies this as an SCM SDK error
	IsScmError() bool
	// HTTPStatusCode returns the HTTP status code if applicable
	HTTPStatusCode() int
	// ErrorCode returns the SCM-specific error code (e.g., "E003")
	ErrorCode() string
	// ErrorMessage returns the detailed error message
	ErrorMessage() string
	// ErrorDetails returns additional error context
	ErrorDetails() map[string]interface{}
}

// BaseError provides common functionality for all SCM errors.
// It implements the ScmError interface and can be embedded in specific error types.
type BaseError struct {
	StatusCode int
	Code       string
	Message    string
	Details    map[string]interface{}
}

// Error implements the error interface
func (e *BaseError) Error() string {
	if e.Code != "" {
		return fmt.Sprintf("[%s] %s", e.Code, e.Message)
	}
	return e.Message
}

// IsScmError identifies this as an SCM error
func (e *BaseError) IsScmError() bool { return true }

// HTTPStatusCode returns the HTTP status code
func (e *BaseError) HTTPStatusCode() int { return e.StatusCode }

// ErrorCode returns the SCM error code
func (e *BaseError) ErrorCode() string { return e.Code }

// ErrorMessage returns the error message
func (e *BaseError) ErrorMessage() string { return e.Message }

// ErrorDetails returns additional error details
func (e *BaseError) ErrorDetails() map[string]interface{} {
	if e.Details == nil {
		return make(map[string]interface{})
	}
	return e.Details
}

// ============================================================================
// Client Errors (4xx)
// ============================================================================

// ClientError represents a 4xx client error.
// This is the base type for all client-side errors.
type ClientError struct {
	BaseError
}

// ============================================================================
// Authentication Errors (401)
// ============================================================================

// AuthenticationError represents a 401 authentication error.
// This indicates the request lacks valid authentication credentials.
type AuthenticationError struct {
	ClientError
}

// NotAuthenticatedError indicates the request is not authenticated.
// Error Code: E003
type NotAuthenticatedError struct {
	AuthenticationError
}

// InvalidCredentialError indicates invalid credentials were provided.
// Error Code: E004
type InvalidCredentialError struct {
	AuthenticationError
}

// KeyExpiredError indicates the API key has expired.
// Error Code: E013
type KeyExpiredError struct {
	AuthenticationError
}

// ============================================================================
// Authorization Errors (403)
// ============================================================================

// AuthorizationError represents a 403 authorization error.
// This indicates the authenticated user lacks permission for the operation.
// Error Code: E009
type AuthorizationError struct {
	ClientError
}

// ============================================================================
// Bad Request Errors (400)
// ============================================================================

// BadRequestError represents a 400 bad request error.
// This is the base type for malformed or invalid requests.
type BadRequestError struct {
	ClientError
}

// InvalidObjectError indicates the object failed validation.
// Error Code: E023
type InvalidObjectError struct {
	BadRequestError
	ObjectName string
}

// MissingQueryParameterError indicates a required query parameter is missing.
// Error Code: E007
type MissingQueryParameterError struct {
	BadRequestError
	ParameterName string
}

// InvalidQueryParameterError indicates a query parameter has an invalid value.
// Error Code: E024
type InvalidQueryParameterError struct {
	BadRequestError
	ParameterName  string
	ParameterValue string
}

// MalformedCommandError indicates the command syntax is invalid.
// Error Code: E006
type MalformedCommandError struct {
	BadRequestError
}

// ============================================================================
// Not Found Errors (404)
// ============================================================================

// NotFoundError represents a 404 not found error.
// This is the base type for resource-not-found errors.
type NotFoundError struct {
	ClientError
}

// ObjectNotPresentError indicates the requested object doesn't exist.
// Error Code: E005
type ObjectNotPresentError struct {
	NotFoundError
	ObjectID   string
	ObjectName string
}

// ============================================================================
// Conflict Errors (409)
// ============================================================================

// ConflictError represents a 409 conflict error.
// This is the base type for resource conflicts.
type ConflictError struct {
	ClientError
}

// NameNotUniqueError indicates a name conflict (duplicate name).
// Error Code: E016
type NameNotUniqueError struct {
	ConflictError
	ObjectName string
}

// ObjectNotUniqueError indicates an object conflict.
// Error Code: E017
type ObjectNotUniqueError struct {
	ConflictError
	ObjectID   string
	ObjectName string
}

// ReferenceNotZeroError indicates the object has references preventing deletion.
// Error Code: E018
type ReferenceNotZeroError struct {
	ConflictError
	ObjectName     string
	ReferenceCount int
}

// ============================================================================
// Other Client Errors
// ============================================================================

// MethodNotAllowedError represents a 405 method not allowed error.
// Error Code: E010
type MethodNotAllowedError struct {
	ClientError
	Method string
}

// RequestTimeoutError represents a 408 request timeout error.
// Error Code: E011
type RequestTimeoutError struct {
	ClientError
}

// TooManyRequestsError represents a 429 too many requests error (rate limiting).
// Error Code: E012
type TooManyRequestsError struct {
	ClientError
	RetryAfter int // Seconds to wait before retrying
}

// SessionTimedOutError indicates the session has expired.
// Error Code: E019
type SessionTimedOutError struct {
	ClientError
}

// ============================================================================
// Server Errors (5xx)
// ============================================================================

// ServerError represents a 5xx server error.
// This is the base type for all server-side errors.
type ServerError struct {
	BaseError
}

// InternalServerError represents a 500 internal server error.
// Error Code: E020
type InternalServerError struct {
	ServerError
}

// BadGatewayError represents a 502 bad gateway error.
// Error Code: E021
type BadGatewayError struct {
	ServerError
}

// ServiceUnavailableError represents a 503 service unavailable error.
// Error Code: E022
type ServiceUnavailableError struct {
	ServerError
}

// GatewayTimeoutError represents a 504 gateway timeout error.
// Error Code: E024
type GatewayTimeoutError struct {
	ServerError
}
