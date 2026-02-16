package errors

// Helper functions for checking and extracting specific error types.
// These functions provide a type-safe way to inspect errors without
// needing to use type assertions directly.

// ============================================================================
// Type Check Helpers (Is* functions)
// ============================================================================

// IsScmError checks if the error is any SCM error type.
func IsScmError(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(ScmError)
	return ok
}

// IsClientError checks if the error is a client error (4xx).
func IsClientError(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(*ClientError)
	if ok {
		return true
	}
	// Check embedded types
	switch err.(type) {
	case *AuthenticationError, *AuthorizationError, *BadRequestError,
		*NotFoundError, *ConflictError, *MethodNotAllowedError,
		*RequestTimeoutError, *TooManyRequestsError, *SessionTimedOutError,
		*NotAuthenticatedError, *InvalidCredentialError, *KeyExpiredError,
		*InvalidObjectError, *MissingQueryParameterError, *InvalidQueryParameterError,
		*MalformedCommandError, *ObjectNotPresentError, *NameNotUniqueError,
		*ObjectNotUniqueError, *ReferenceNotZeroError:
		return true
	}
	return false
}

// IsServerError checks if the error is a server error (5xx).
func IsServerError(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(*ServerError)
	if ok {
		return true
	}
	// Check embedded types
	switch err.(type) {
	case *InternalServerError, *BadGatewayError, *ServiceUnavailableError, *GatewayTimeoutError:
		return true
	}
	return false
}

// IsAuthenticationError checks if the error is an authentication error (401).
func IsAuthenticationError(err error) bool {
	if err == nil {
		return false
	}
	switch err.(type) {
	case *AuthenticationError, *NotAuthenticatedError, *InvalidCredentialError, *KeyExpiredError:
		return true
	}
	return false
}

// IsNotAuthenticated checks if the error is NotAuthenticatedError.
func IsNotAuthenticated(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(*NotAuthenticatedError)
	return ok
}

// IsInvalidCredential checks if the error is InvalidCredentialError.
func IsInvalidCredential(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(*InvalidCredentialError)
	return ok
}

// IsKeyExpired checks if the error is KeyExpiredError.
func IsKeyExpired(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(*KeyExpiredError)
	return ok
}

// IsAuthorizationError checks if the error is AuthorizationError.
func IsAuthorizationError(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(*AuthorizationError)
	return ok
}

// IsBadRequest checks if the error is a bad request error (400).
func IsBadRequest(err error) bool {
	if err == nil {
		return false
	}
	switch err.(type) {
	case *BadRequestError, *InvalidObjectError, *MissingQueryParameterError,
		*InvalidQueryParameterError, *MalformedCommandError:
		return true
	}
	return false
}

// IsInvalidObject checks if the error is InvalidObjectError.
func IsInvalidObject(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(*InvalidObjectError)
	return ok
}

// IsMissingQueryParameter checks if the error is MissingQueryParameterError.
func IsMissingQueryParameter(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(*MissingQueryParameterError)
	return ok
}

// IsInvalidQueryParameter checks if the error is InvalidQueryParameterError.
func IsInvalidQueryParameter(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(*InvalidQueryParameterError)
	return ok
}

// IsMalformedCommand checks if the error is MalformedCommandError.
func IsMalformedCommand(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(*MalformedCommandError)
	return ok
}

// IsNotFound checks if the error is a not found error (404).
func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	switch err.(type) {
	case *NotFoundError, *ObjectNotPresentError:
		return true
	}
	return false
}

// IsObjectNotPresent checks if the error is ObjectNotPresentError.
func IsObjectNotPresent(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(*ObjectNotPresentError)
	return ok
}

// IsConflict checks if the error is a conflict error (409).
func IsConflict(err error) bool {
	if err == nil {
		return false
	}
	switch err.(type) {
	case *ConflictError, *NameNotUniqueError, *ObjectNotUniqueError, *ReferenceNotZeroError:
		return true
	}
	return false
}

// IsNameNotUnique checks if the error is NameNotUniqueError.
func IsNameNotUnique(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(*NameNotUniqueError)
	return ok
}

// IsObjectNotUnique checks if the error is ObjectNotUniqueError.
func IsObjectNotUnique(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(*ObjectNotUniqueError)
	return ok
}

// IsReferenceNotZero checks if the error is ReferenceNotZeroError.
func IsReferenceNotZero(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(*ReferenceNotZeroError)
	return ok
}

// IsMethodNotAllowed checks if the error is MethodNotAllowedError.
func IsMethodNotAllowed(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(*MethodNotAllowedError)
	return ok
}

// IsRequestTimeout checks if the error is RequestTimeoutError.
func IsRequestTimeout(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(*RequestTimeoutError)
	return ok
}

// IsTooManyRequests checks if the error is TooManyRequestsError (rate limit).
func IsTooManyRequests(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(*TooManyRequestsError)
	return ok
}

// IsSessionTimedOut checks if the error is SessionTimedOutError.
func IsSessionTimedOut(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(*SessionTimedOutError)
	return ok
}

// IsInternalServerError checks if the error is InternalServerError.
func IsInternalServerError(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(*InternalServerError)
	return ok
}

// IsBadGateway checks if the error is BadGatewayError.
func IsBadGateway(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(*BadGatewayError)
	return ok
}

// IsServiceUnavailable checks if the error is ServiceUnavailableError.
func IsServiceUnavailable(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(*ServiceUnavailableError)
	return ok
}

// IsGatewayTimeout checks if the error is GatewayTimeoutError.
func IsGatewayTimeout(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(*GatewayTimeoutError)
	return ok
}

// ============================================================================
// Type Extraction Helpers (As* functions)
// ============================================================================

// AsObjectNotPresent attempts to extract ObjectNotPresentError from the error.
// Returns the typed error and true if successful, nil and false otherwise.
func AsObjectNotPresent(err error) (*ObjectNotPresentError, bool) {
	if err == nil {
		return nil, false
	}
	e, ok := err.(*ObjectNotPresentError)
	return e, ok
}

// AsNameNotUnique attempts to extract NameNotUniqueError from the error.
// Returns the typed error and true if successful, nil and false otherwise.
func AsNameNotUnique(err error) (*NameNotUniqueError, bool) {
	if err == nil {
		return nil, false
	}
	e, ok := err.(*NameNotUniqueError)
	return e, ok
}

// AsObjectNotUnique attempts to extract ObjectNotUniqueError from the error.
// Returns the typed error and true if successful, nil and false otherwise.
func AsObjectNotUnique(err error) (*ObjectNotUniqueError, bool) {
	if err == nil {
		return nil, false
	}
	e, ok := err.(*ObjectNotUniqueError)
	return e, ok
}

// AsReferenceNotZero attempts to extract ReferenceNotZeroError from the error.
// Returns the typed error and true if successful, nil and false otherwise.
func AsReferenceNotZero(err error) (*ReferenceNotZeroError, bool) {
	if err == nil {
		return nil, false
	}
	e, ok := err.(*ReferenceNotZeroError)
	return e, ok
}

// AsTooManyRequests attempts to extract TooManyRequestsError from the error.
// Returns the typed error and true if successful, nil and false otherwise.
func AsTooManyRequests(err error) (*TooManyRequestsError, bool) {
	if err == nil {
		return nil, false
	}
	e, ok := err.(*TooManyRequestsError)
	return e, ok
}

// AsInvalidObject attempts to extract InvalidObjectError from the error.
// Returns the typed error and true if successful, nil and false otherwise.
func AsInvalidObject(err error) (*InvalidObjectError, bool) {
	if err == nil {
		return nil, false
	}
	e, ok := err.(*InvalidObjectError)
	return e, ok
}

// AsMissingQueryParameter attempts to extract MissingQueryParameterError from the error.
// Returns the typed error and true if successful, nil and false otherwise.
func AsMissingQueryParameter(err error) (*MissingQueryParameterError, bool) {
	if err == nil {
		return nil, false
	}
	e, ok := err.(*MissingQueryParameterError)
	return e, ok
}

// AsInvalidQueryParameter attempts to extract InvalidQueryParameterError from the error.
// Returns the typed error and true if successful, nil and false otherwise.
func AsInvalidQueryParameter(err error) (*InvalidQueryParameterError, bool) {
	if err == nil {
		return nil, false
	}
	e, ok := err.(*InvalidQueryParameterError)
	return e, ok
}

// AsMethodNotAllowed attempts to extract MethodNotAllowedError from the error.
// Returns the typed error and true if successful, nil and false otherwise.
func AsMethodNotAllowed(err error) (*MethodNotAllowedError, bool) {
	if err == nil {
		return nil, false
	}
	e, ok := err.(*MethodNotAllowedError)
	return e, ok
}
