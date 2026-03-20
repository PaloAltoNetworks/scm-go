package errors_test

import (
	"context"
	"fmt"
	"log"

	scmErrors "github.com/paloaltonetworks/scm-go/errors"
)

// Example_nameConflict demonstrates handling name uniqueness errors
func Example_nameConflict() {
	// Simulate an API error
	var err error = scmErrors.NewNameNotUniqueError("web-server")

	// Check for name conflict
	if scmErrors.IsNameNotUnique(err) {
		nameErr, ok := scmErrors.AsNameNotUnique(err)
		if ok {
			fmt.Printf("Name conflict: %s\n", nameErr.ObjectName)
			fmt.Printf("Error code: %s\n", nameErr.ErrorCode())
			fmt.Printf("HTTP status: %d\n", nameErr.HTTPStatusCode())
		}
	}

	// Output:
	// Name conflict: web-server
	// Error code: E016
	// HTTP status: 409
}

// Example_objectNotFound demonstrates handling object not found errors
func Example_objectNotFound() {
	// Simulate an object not found error
	var err error = scmErrors.NewObjectNotPresentError("12345", "web-server")

	// Check various error types
	if scmErrors.IsObjectNotPresent(err) {
		objErr, ok := scmErrors.AsObjectNotPresent(err)
		if ok {
			fmt.Printf("Object not found: %s (ID: %s)\n", objErr.ObjectName, objErr.ObjectID)
		}
	}

	// Output:
	// Object not found: web-server (ID: 12345)
}

// Example_rateLimiting demonstrates handling rate limit errors
func Example_rateLimiting() {
	// Simulate a rate limit error
	var err error = scmErrors.NewTooManyRequestsError(60, "Too many requests")

	// Check for rate limiting
	if scmErrors.IsTooManyRequests(err) {
		rateErr, ok := scmErrors.AsTooManyRequests(err)
		if ok {
			fmt.Printf("Rate limited - retry after %d seconds\n", rateErr.RetryAfter)
		}
	}

	// Output:
	// Rate limited - retry after 60 seconds
}

// Example_comprehensiveErrorHandling shows a complete error handling pattern
func Example_comprehensiveErrorHandling() {
	// Simulate various errors
	testErrors := []error{
		scmErrors.NewNameNotUniqueError("web-server"),
		scmErrors.NewObjectNotPresentError("123", "db-server"),
		scmErrors.NewAuthorizationError("no permission"),
		scmErrors.NewInternalServerError("server error"),
	}

	for _, err := range testErrors {
		handleError(err)
	}

	// Output:
	// Name conflict: web-server
	// Object not found: db-server
	// Permission denied
	// Server error occurred
}

func handleError(err error) {
	switch {
	case scmErrors.IsNameNotUnique(err):
		nameErr, _ := scmErrors.AsNameNotUnique(err)
		fmt.Printf("Name conflict: %s\n", nameErr.ObjectName)

	case scmErrors.IsObjectNotPresent(err):
		objErr, _ := scmErrors.AsObjectNotPresent(err)
		fmt.Printf("Object not found: %s\n", objErr.ObjectName)

	case scmErrors.IsAuthorizationError(err):
		fmt.Println("Permission denied")

	case scmErrors.IsServerError(err):
		fmt.Println("Server error occurred")

	default:
		fmt.Printf("Unknown error: %v\n", err)
	}
}

// Example_authentication demonstrates handling authentication errors
func Example_authentication() {
	errors := []error{
		scmErrors.NewNotAuthenticatedError("not authenticated"),
		scmErrors.NewInvalidCredentialError("bad credentials"),
		scmErrors.NewKeyExpiredError("key expired"),
	}

	for _, err := range errors {
		if scmErrors.IsAuthenticationError(err) {
			switch {
			case scmErrors.IsNotAuthenticated(err):
				fmt.Println("Not authenticated - please login")
			case scmErrors.IsInvalidCredential(err):
				fmt.Println("Invalid credentials - check client ID/secret")
			case scmErrors.IsKeyExpired(err):
				fmt.Println("API key expired - please renew")
			}
		}
	}

	// Output:
	// Not authenticated - please login
	// Invalid credentials - check client ID/secret
	// API key expired - please renew
}

// Example_referenceNotZero demonstrates handling deletion errors
func Example_referenceNotZero() {
	// Simulate a reference error
	var err error = scmErrors.NewReferenceNotZeroError("my-security-rule", 5)

	if scmErrors.IsReferenceNotZero(err) {
		refErr, ok := scmErrors.AsReferenceNotZero(err)
		if ok {
			fmt.Printf("Cannot delete '%s': %d reference(s) exist\n",
				refErr.ObjectName, refErr.ReferenceCount)
		}
	}

	// Output:
	// Cannot delete 'my-security-rule': 5 reference(s) exist
}

// Example_errorInterface demonstrates using the ScmError interface
func Example_errorInterface() {
	var err scmErrors.ScmError = scmErrors.NewNameNotUniqueError("web-server")

	// Access interface methods
	fmt.Printf("Is SCM error: %v\n", err.IsScmError())
	fmt.Printf("HTTP status: %d\n", err.HTTPStatusCode())
	fmt.Printf("Error code: %s\n", err.ErrorCode())
	fmt.Printf("Message: %s\n", err.ErrorMessage())

	// Output:
	// Is SCM error: true
	// HTTP status: 409
	// Error code: E016
	// Message: Name 'web-server' is not unique
}

// Example_realWorld demonstrates realistic API usage
func Example_realWorld() {
	ctx := context.Background()

	// Simulated API call result
	err := simulateAPICall(ctx, "web-server")

	// Handle the error
	if err != nil {
		switch {
		case scmErrors.IsNameNotUnique(err):
			log.Println("Name already exists - trying with suffix")
			err = simulateAPICall(ctx, "web-server-2")
			if err != nil {
				log.Printf("Retry failed: %v", err)
			}

		case scmErrors.IsAuthorizationError(err):
			log.Println("Permission denied - check credentials")

		case scmErrors.IsTooManyRequests(err):
			if rateErr, ok := scmErrors.AsTooManyRequests(err); ok {
				log.Printf("Rate limited - wait %d seconds", rateErr.RetryAfter)
			}

		default:
			log.Printf("Unexpected error: %v", err)
		}
	}
}

func simulateAPICall(ctx context.Context, name string) error {
	// This would be your actual API call
	// For example purposes, return a name conflict
	if name == "web-server" {
		return scmErrors.NewNameNotUniqueError(name)
	}
	return nil
}
