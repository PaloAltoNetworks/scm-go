package errors

import "fmt"

// Constructor functions for creating typed errors with proper defaults.
// These functions ensure errors are created with the correct status codes,
// error codes, and structure.

// ============================================================================
// Authentication Error Constructors (401)
// ============================================================================

// NewNotAuthenticatedError creates a new NotAuthenticatedError.
// Error Code: E003, HTTP Status: 401
func NewNotAuthenticatedError(message string) *NotAuthenticatedError {
	if message == "" {
		message = "Not authenticated"
	}
	return &NotAuthenticatedError{
		AuthenticationError: AuthenticationError{
			ClientError: ClientError{
				BaseError: BaseError{
					StatusCode: 401,
					Code:       "E003",
					Message:    message,
					Details:    make(map[string]interface{}),
				},
			},
		},
	}
}

// NewInvalidCredentialError creates a new InvalidCredentialError.
// Error Code: E004, HTTP Status: 401
func NewInvalidCredentialError(message string) *InvalidCredentialError {
	if message == "" {
		message = "Invalid credentials"
	}
	return &InvalidCredentialError{
		AuthenticationError: AuthenticationError{
			ClientError: ClientError{
				BaseError: BaseError{
					StatusCode: 401,
					Code:       "E004",
					Message:    message,
					Details:    make(map[string]interface{}),
				},
			},
		},
	}
}

// NewKeyExpiredError creates a new KeyExpiredError.
// Error Code: E013, HTTP Status: 401
func NewKeyExpiredError(message string) *KeyExpiredError {
	if message == "" {
		message = "API key expired"
	}
	return &KeyExpiredError{
		AuthenticationError: AuthenticationError{
			ClientError: ClientError{
				BaseError: BaseError{
					StatusCode: 401,
					Code:       "E013",
					Message:    message,
					Details:    make(map[string]interface{}),
				},
			},
		},
	}
}

// ============================================================================
// Authorization Error Constructors (403)
// ============================================================================

// NewAuthorizationError creates a new AuthorizationError.
// Error Code: E009, HTTP Status: 403
func NewAuthorizationError(message string) *AuthorizationError {
	if message == "" {
		message = "Forbidden - insufficient permissions"
	}
	return &AuthorizationError{
		ClientError: ClientError{
			BaseError: BaseError{
				StatusCode: 403,
				Code:       "E009",
				Message:    message,
				Details:    make(map[string]interface{}),
			},
		},
	}
}

// ============================================================================
// Bad Request Error Constructors (400)
// ============================================================================

// NewInvalidObjectError creates a new InvalidObjectError.
// Error Code: E023, HTTP Status: 400
func NewInvalidObjectError(objectName, message string) *InvalidObjectError {
	if message == "" {
		message = fmt.Sprintf("Invalid object: %s", objectName)
	}
	return &InvalidObjectError{
		BadRequestError: BadRequestError{
			ClientError: ClientError{
				BaseError: BaseError{
					StatusCode: 400,
					Code:       "E023",
					Message:    message,
					Details:    make(map[string]interface{}),
				},
			},
		},
		ObjectName: objectName,
	}
}

// NewMissingQueryParameterError creates a new MissingQueryParameterError.
// Error Code: E007, HTTP Status: 400
func NewMissingQueryParameterError(paramName string) *MissingQueryParameterError {
	message := fmt.Sprintf("Missing required query parameter: %s", paramName)
	return &MissingQueryParameterError{
		BadRequestError: BadRequestError{
			ClientError: ClientError{
				BaseError: BaseError{
					StatusCode: 400,
					Code:       "E007",
					Message:    message,
					Details:    make(map[string]interface{}),
				},
			},
		},
		ParameterName: paramName,
	}
}

// NewInvalidQueryParameterError creates a new InvalidQueryParameterError.
// Error Code: E024, HTTP Status: 400
func NewInvalidQueryParameterError(paramName, paramValue, message string) *InvalidQueryParameterError {
	if message == "" {
		message = fmt.Sprintf("Invalid query parameter '%s': %s", paramName, paramValue)
	}
	return &InvalidQueryParameterError{
		BadRequestError: BadRequestError{
			ClientError: ClientError{
				BaseError: BaseError{
					StatusCode: 400,
					Code:       "E024",
					Message:    message,
					Details:    make(map[string]interface{}),
				},
			},
		},
		ParameterName:  paramName,
		ParameterValue: paramValue,
	}
}

// NewMalformedCommandError creates a new MalformedCommandError.
// Error Code: E006, HTTP Status: 400
func NewMalformedCommandError(message string) *MalformedCommandError {
	if message == "" {
		message = "Malformed command"
	}
	return &MalformedCommandError{
		BadRequestError: BadRequestError{
			ClientError: ClientError{
				BaseError: BaseError{
					StatusCode: 400,
					Code:       "E006",
					Message:    message,
					Details:    make(map[string]interface{}),
				},
			},
		},
	}
}

// ============================================================================
// Not Found Error Constructors (404)
// ============================================================================

// NewObjectNotPresentError creates a new ObjectNotPresentError.
// Error Code: E005, HTTP Status: 404
func NewObjectNotPresentError(objectID, objectName string) *ObjectNotPresentError {
	message := fmt.Sprintf("Object not found: %s", objectName)
	if objectID != "" {
		message = fmt.Sprintf("Object not found: %s (ID: %s)", objectName, objectID)
	}
	return &ObjectNotPresentError{
		NotFoundError: NotFoundError{
			ClientError: ClientError{
				BaseError: BaseError{
					StatusCode: 404,
					Code:       "E005",
					Message:    message,
					Details:    make(map[string]interface{}),
				},
			},
		},
		ObjectID:   objectID,
		ObjectName: objectName,
	}
}

// ============================================================================
// Conflict Error Constructors (409)
// ============================================================================

// NewNameNotUniqueError creates a new NameNotUniqueError.
// Error Code: E016, HTTP Status: 409
func NewNameNotUniqueError(objectName string) *NameNotUniqueError {
	message := fmt.Sprintf("Name '%s' is not unique", objectName)
	return &NameNotUniqueError{
		ConflictError: ConflictError{
			ClientError: ClientError{
				BaseError: BaseError{
					StatusCode: 409,
					Code:       "E016",
					Message:    message,
					Details:    make(map[string]interface{}),
				},
			},
		},
		ObjectName: objectName,
	}
}

// NewObjectNotUniqueError creates a new ObjectNotUniqueError.
// Error Code: E017, HTTP Status: 409
func NewObjectNotUniqueError(objectID, objectName string) *ObjectNotUniqueError {
	message := fmt.Sprintf("Object '%s' is not unique", objectName)
	return &ObjectNotUniqueError{
		ConflictError: ConflictError{
			ClientError: ClientError{
				BaseError: BaseError{
					StatusCode: 409,
					Code:       "E017",
					Message:    message,
					Details:    make(map[string]interface{}),
				},
			},
		},
		ObjectID:   objectID,
		ObjectName: objectName,
	}
}

// NewReferenceNotZeroError creates a new ReferenceNotZeroError.
// Error Code: E018, HTTP Status: 409
func NewReferenceNotZeroError(objectName string, refCount int) *ReferenceNotZeroError {
	message := fmt.Sprintf("Cannot delete '%s': %d reference(s) exist", objectName, refCount)
	return &ReferenceNotZeroError{
		ConflictError: ConflictError{
			ClientError: ClientError{
				BaseError: BaseError{
					StatusCode: 409,
					Code:       "E018",
					Message:    message,
					Details:    make(map[string]interface{}),
				},
			},
		},
		ObjectName:     objectName,
		ReferenceCount: refCount,
	}
}

// ============================================================================
// Other Client Error Constructors
// ============================================================================

// NewMethodNotAllowedError creates a new MethodNotAllowedError.
// Error Code: E010, HTTP Status: 405
func NewMethodNotAllowedError(method string) *MethodNotAllowedError {
	message := fmt.Sprintf("Method not allowed: %s", method)
	return &MethodNotAllowedError{
		ClientError: ClientError{
			BaseError: BaseError{
				StatusCode: 405,
				Code:       "E010",
				Message:    message,
				Details:    make(map[string]interface{}),
			},
		},
		Method: method,
	}
}

// NewRequestTimeoutError creates a new RequestTimeoutError.
// Error Code: E011, HTTP Status: 408
func NewRequestTimeoutError(message string) *RequestTimeoutError {
	if message == "" {
		message = "Request timeout"
	}
	return &RequestTimeoutError{
		ClientError: ClientError{
			BaseError: BaseError{
				StatusCode: 408,
				Code:       "E011",
				Message:    message,
				Details:    make(map[string]interface{}),
			},
		},
	}
}

// NewTooManyRequestsError creates a new TooManyRequestsError.
// Error Code: E012, HTTP Status: 429
func NewTooManyRequestsError(retryAfter int, message string) *TooManyRequestsError {
	if message == "" {
		message = fmt.Sprintf("Too many requests - retry after %d seconds", retryAfter)
	}
	return &TooManyRequestsError{
		ClientError: ClientError{
			BaseError: BaseError{
				StatusCode: 429,
				Code:       "E012",
				Message:    message,
				Details:    make(map[string]interface{}),
			},
		},
		RetryAfter: retryAfter,
	}
}

// NewSessionTimedOutError creates a new SessionTimedOutError.
// Error Code: E019, HTTP Status: 401
func NewSessionTimedOutError(message string) *SessionTimedOutError {
	if message == "" {
		message = "Session timed out"
	}
	return &SessionTimedOutError{
		ClientError: ClientError{
			BaseError: BaseError{
				StatusCode: 401,
				Code:       "E019",
				Message:    message,
				Details:    make(map[string]interface{}),
			},
		},
	}
}

// ============================================================================
// Server Error Constructors (5xx)
// ============================================================================

// NewInternalServerError creates a new InternalServerError.
// Error Code: E020, HTTP Status: 500
func NewInternalServerError(message string) *InternalServerError {
	if message == "" {
		message = "Internal server error"
	}
	return &InternalServerError{
		ServerError: ServerError{
			BaseError: BaseError{
				StatusCode: 500,
				Code:       "E020",
				Message:    message,
				Details:    make(map[string]interface{}),
			},
		},
	}
}

// NewBadGatewayError creates a new BadGatewayError.
// Error Code: E021, HTTP Status: 502
func NewBadGatewayError(message string) *BadGatewayError {
	if message == "" {
		message = "Bad gateway"
	}
	return &BadGatewayError{
		ServerError: ServerError{
			BaseError: BaseError{
				StatusCode: 502,
				Code:       "E021",
				Message:    message,
				Details:    make(map[string]interface{}),
			},
		},
	}
}

// NewServiceUnavailableError creates a new ServiceUnavailableError.
// Error Code: E022, HTTP Status: 503
func NewServiceUnavailableError(message string) *ServiceUnavailableError {
	if message == "" {
		message = "Service unavailable"
	}
	return &ServiceUnavailableError{
		ServerError: ServerError{
			BaseError: BaseError{
				StatusCode: 503,
				Code:       "E022",
				Message:    message,
				Details:    make(map[string]interface{}),
			},
		},
	}
}

// NewGatewayTimeoutError creates a new GatewayTimeoutError.
// Error Code: E024, HTTP Status: 504
func NewGatewayTimeoutError(message string) *GatewayTimeoutError {
	if message == "" {
		message = "Gateway timeout"
	}
	return &GatewayTimeoutError{
		ServerError: ServerError{
			BaseError: BaseError{
				StatusCode: 504,
				Code:       "E024",
				Message:    message,
				Details:    make(map[string]interface{}),
			},
		},
	}
}
