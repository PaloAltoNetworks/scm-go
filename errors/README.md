# SCM Go SDK - Errors Package

Typed error handling for the Palo Alto Networks Strata Cloud Manager Go SDK.

## Overview

The `errors` package provides 24 typed error types that match the SCM API error responses. This enables type-safe error handling without needing to parse error strings.

## Features

- ✅ **24 Typed Errors**: Specific types for every SCM API error
- ✅ **Error Hierarchy**: Organized by HTTP status codes (4xx client, 5xx server)
- ✅ **Rich Context**: Errors include object names, IDs, and detailed messages
- ✅ **Helper Functions**: Easy type checking with `Is*()` and extraction with `As*()`
- ✅ **Backwards Compatible**: Purely additive - all existing code continues to work
- ✅ **Opt-In**: Use typed errors only where you need them

## Installation

```bash
go get github.com/paloaltonetworks/scm-go/errors
```

## Quick Start

```go
import (
    scmErrors "github.com/paloaltonetworks/scm-go/errors"
    "github.com/paloaltonetworks/scm-go/generated/objects"
)

resp, _, err := api.CreateAddresses(ctx).Addresses(addr).Execute()
if err != nil {
    // Check for specific error types
    if scmErrors.IsNameNotUnique(err) {
        log.Printf("Name conflict detected")
        return handleNameConflict(err)
    }

    if scmErrors.IsObjectNotPresent(err) {
        log.Printf("Object not found")
        return nil
    }

    // Generic error handling
    return err
}
```

## Error Hierarchy

```
ScmError (interface)
├── ClientError (4xx errors)
│   ├── AuthenticationError (401)
│   │   ├── NotAuthenticatedError (E003)
│   │   ├── InvalidCredentialError (E004)
│   │   └── KeyExpiredError (E013)
│   ├── AuthorizationError (403, E009)
│   ├── BadRequestError (400)
│   │   ├── InvalidObjectError (E023)
│   │   ├── MissingQueryParameterError (E007)
│   │   ├── InvalidQueryParameterError (E024)
│   │   └── MalformedCommandError (E006)
│   ├── NotFoundError (404)
│   │   └── ObjectNotPresentError (E005)
│   ├── ConflictError (409)
│   │   ├── NameNotUniqueError (E016)
│   │   ├── ObjectNotUniqueError (E017)
│   │   └── ReferenceNotZeroError (E018)
│   ├── MethodNotAllowedError (405, E010)
│   ├── RequestTimeoutError (408, E011)
│   ├── TooManyRequestsError (429, E012)
│   └── SessionTimedOutError (401, E019)
└── ServerError (5xx errors)
    ├── InternalServerError (500, E020)
    ├── BadGatewayError (502, E021)
    ├── ServiceUnavailableError (503, E022)
    └── GatewayTimeoutError (504, E024)
```

## Usage Examples

### Example 1: Handle Name Conflicts

```go
import scmErrors "github.com/paloaltonetworks/scm-go/errors"

addr := objects.Addresses{
    Name:      "web-server",
    Folder:    common.StringPtr("Texas"),
    IpNetmask: common.StringPtr("10.1.1.1/32"),
}

resp, _, err := api.CreateAddresses(ctx).Addresses(addr).Execute()
if err != nil {
    if scmErrors.IsNameNotUnique(err) {
        // Extract typed error for more details
        nameErr, ok := scmErrors.AsNameNotUnique(err)
        if ok {
            log.Printf("Name conflict: %s already exists", nameErr.ObjectName)
            // Handle by appending suffix, generating new name, etc.
            return handleNameConflict(nameErr.ObjectName)
        }
    }
    return err
}
```

### Example 2: Handle Object Not Found

```go
resp, _, err := api.GetAddressesByID(ctx, objectID).Execute()
if err != nil {
    if scmErrors.IsObjectNotPresent(err) {
        // Object doesn't exist - this may be expected
        log.Printf("Object %s not found - may have been deleted", objectID)
        return nil // Or handle gracefully
    }

    // Unexpected error
    return fmt.Errorf("failed to get address: %w", err)
}
```

### Example 3: Handle Rate Limiting

```go
resp, _, err := api.CreateAddresses(ctx).Addresses(addr).Execute()
if err != nil {
    if scmErrors.IsTooManyRequests(err) {
        rateErr, ok := scmErrors.AsTooManyRequests(err)
        if ok {
            // Use RetryAfter to implement backoff
            log.Printf("Rate limited - retry after %d seconds", rateErr.RetryAfter)
            time.Sleep(time.Duration(rateErr.RetryAfter) * time.Second)
            // Retry the request
            return retryRequest(ctx, api, addr)
        }
    }
    return err
}
```

### Example 4: Handle Authentication Errors

```go
resp, _, err := api.CreateAddresses(ctx).Addresses(addr).Execute()
if err != nil {
    if scmErrors.IsAuthenticationError(err) {
        // Could be NotAuthenticatedError, InvalidCredentialError, or KeyExpiredError
        if scmErrors.IsKeyExpired(err) {
            log.Printf("API key expired - please renew")
            return refreshAPIKey()
        }

        if scmErrors.IsInvalidCredential(err) {
            log.Printf("Invalid credentials - please check client ID/secret")
            return fmt.Errorf("authentication failed")
        }

        // Generic authentication error
        log.Printf("Not authenticated")
        return fmt.Errorf("authentication required")
    }
    return err
}
```

### Example 5: Handle Object Deletion with References

```go
_, err := api.DeleteAddressesByID(ctx, objectID).Execute()
if err != nil {
    if scmErrors.IsReferenceNotZero(err) {
        refErr, ok := scmErrors.AsReferenceNotZero(err)
        if ok {
            log.Printf("Cannot delete %s: %d reference(s) exist",
                refErr.ObjectName, refErr.ReferenceCount)

            // Handle by removing references first, or showing user the count
            return fmt.Errorf("delete blocked: %d references exist", refErr.ReferenceCount)
        }
    }
    return err
}
```

### Example 6: Comprehensive Error Handling

```go
func createAddress(ctx context.Context, api *objects.AddressesAPIService, addr objects.Addresses) error {
    resp, _, err := api.CreateAddresses(ctx).Addresses(addr).Execute()
    if err != nil {
        // Handle specific client errors
        if scmErrors.IsClientError(err) {
            switch {
            case scmErrors.IsNameNotUnique(err):
                return fmt.Errorf("name '%s' already exists", addr.Name)

            case scmErrors.IsInvalidObject(err):
                objErr, _ := scmErrors.AsInvalidObject(err)
                return fmt.Errorf("invalid object '%s': %s",
                    objErr.ObjectName, objErr.ErrorMessage())

            case scmErrors.IsAuthorizationError(err):
                return fmt.Errorf("permission denied")

            case scmErrors.IsTooManyRequests(err):
                rateErr, _ := scmErrors.AsTooManyRequests(err)
                return fmt.Errorf("rate limited - retry after %d seconds",
                    rateErr.RetryAfter)

            default:
                return fmt.Errorf("client error: %w", err)
            }
        }

        // Handle server errors
        if scmErrors.IsServerError(err) {
            if scmErrors.IsServiceUnavailable(err) {
                return fmt.Errorf("service temporarily unavailable - retry later")
            }
            return fmt.Errorf("server error: %w", err)
        }

        // Unknown error
        return fmt.Errorf("unexpected error: %w", err)
    }

    log.Printf("Created address: %s (ID: %s)", resp.Name, resp.Id)
    return nil
}
```

## Helper Functions

### Type Checking (`Is*` functions)

Check if an error is of a specific type:

```go
IsScmError(err)              // Any SCM error
IsClientError(err)           // Any 4xx error
IsServerError(err)           // Any 5xx error
IsAuthenticationError(err)   // Any 401 error
IsNotAuthenticated(err)      // E003
IsInvalidCredential(err)     // E004
IsKeyExpired(err)            // E013
IsAuthorizationError(err)    // E009
IsBadRequest(err)            // Any 400 error
IsInvalidObject(err)         // E023
IsMissingQueryParameter(err) // E007
IsInvalidQueryParameter(err) // E024
IsMalformedCommand(err)      // E006
IsNotFound(err)              // Any 404 error
IsObjectNotPresent(err)      // E005
IsConflict(err)              // Any 409 error
IsNameNotUnique(err)         // E016
IsObjectNotUnique(err)       // E017
IsReferenceNotZero(err)      // E018
IsMethodNotAllowed(err)      // E010
IsRequestTimeout(err)        // E011
IsTooManyRequests(err)       // E012
IsSessionTimedOut(err)       // E019
IsInternalServerError(err)   // E020
IsBadGateway(err)            // E021
IsServiceUnavailable(err)    // E022
IsGatewayTimeout(err)        // E024
```

### Type Extraction (`As*` functions)

Extract typed errors to access specific fields:

```go
// Returns (*ObjectNotPresentError, bool)
objErr, ok := AsObjectNotPresent(err)
if ok {
    fmt.Println(objErr.ObjectID)
    fmt.Println(objErr.ObjectName)
}

// Returns (*NameNotUniqueError, bool)
nameErr, ok := AsNameNotUnique(err)
if ok {
    fmt.Println(nameErr.ObjectName)
}

// Returns (*TooManyRequestsError, bool)
rateErr, ok := AsTooManyRequests(err)
if ok {
    fmt.Println(rateErr.RetryAfter)
}

// Returns (*InvalidObjectError, bool)
invalidErr, ok := AsInvalidObject(err)
if ok {
    fmt.Println(invalidErr.ObjectName)
}

// Returns (*ReferenceNotZeroError, bool)
refErr, ok := AsReferenceNotZero(err)
if ok {
    fmt.Println(refErr.ObjectName)
    fmt.Println(refErr.ReferenceCount)
}
```

## Constructor Functions

Create errors manually (useful for testing or custom error handling):

```go
err := scmErrors.NewNameNotUniqueError("web-server")
err := scmErrors.NewObjectNotPresentError("12345", "web-server")
err := scmErrors.NewTooManyRequestsError(60, "rate limited")
err := scmErrors.NewInvalidObjectError("my-object", "validation failed")
err := scmErrors.NewReferenceNotZeroError("my-rule", 5)
// ... constructors for all 24 error types
```

## Error Codes Reference

| Code | HTTP | Error Type | Description |
|------|------|------------|-------------|
| E003 | 401 | NotAuthenticatedError | Not authenticated |
| E004 | 401 | InvalidCredentialError | Invalid credentials |
| E005 | 404 | ObjectNotPresentError | Object not found |
| E006 | 400 | MalformedCommandError | Malformed command |
| E007 | 400 | MissingQueryParameterError | Missing query parameter |
| E009 | 403 | AuthorizationError | Forbidden |
| E010 | 405 | MethodNotAllowedError | Method not allowed |
| E011 | 408 | RequestTimeoutError | Request timeout |
| E012 | 429 | TooManyRequestsError | Too many requests |
| E013 | 401 | KeyExpiredError | API key expired |
| E016 | 409 | NameNotUniqueError | Name not unique |
| E017 | 409 | ObjectNotUniqueError | Object not unique |
| E018 | 409 | ReferenceNotZeroError | Object has references |
| E019 | 401 | SessionTimedOutError | Session timed out |
| E020 | 500 | InternalServerError | Internal server error |
| E021 | 502 | BadGatewayError | Bad gateway |
| E022 | 503 | ServiceUnavailableError | Service unavailable |
| E023 | 400 | InvalidObjectError | Invalid object |
| E024 | 400/504 | InvalidQueryParameterError/GatewayTimeoutError | Invalid query parameter / Gateway timeout |

## Backwards Compatibility

This package is **100% backwards compatible**:

- ✅ All existing code continues to work without changes
- ✅ Typed errors are **opt-in** via type assertions
- ✅ No breaking changes to existing APIs
- ✅ Purely additive enhancements

**Before (still works):**
```go
resp, _, err := api.CreateAddresses(ctx).Addresses(addr).Execute()
if err != nil {
    log.Printf("Error: %v", err)
    return err
}
```

**After (opt-in enhancement):**
```go
resp, _, err := api.CreateAddresses(ctx).Addresses(addr).Execute()
if err != nil {
    if scmErrors.IsNameNotUnique(err) {
        // Handle specifically
    }
    return err
}
```

## Testing

Run the tests:

```bash
cd errors
go test -v
```

Test coverage:

```bash
go test -cover
```

## Best Practices

1. **Check Specific Types First**: Handle specific errors before generic ones
2. **Use Is* for Checking**: Use `Is*()` functions instead of type assertions
3. **Use As* for Extraction**: Use `As*()` functions to extract typed errors safely
4. **Fail Open**: Always have a default case for unknown errors
5. **Log Context**: Use error details (object names, IDs) in log messages

## Contributing

When adding new error types:

1. Add the error type to `errors.go`
2. Add constructor to `constructors.go`
3. Add helpers to `helpers.go`
4. Add tests to `errors_test.go`
5. Update this README

## License

Copyright © 2024 Palo Alto Networks

See LICENSE file for details.
