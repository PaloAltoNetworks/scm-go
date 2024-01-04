package scm

import (
	"net/http"
	"testing"
)

func TestOkAuth(t *testing.T) {
	ar := authResponse{
		Jwt:       "secret",
		Scope:     "test_scope",
		Type:      "test_type",
		ExpiresIn: 899,
	}

	if err := ar.Failed(http.StatusOK, nil); err != nil {
		t.Errorf("Auth failed for %d: %s", http.StatusOK, err)
	}
}

func TestFailedAuth(t *testing.T) {
	ar := authResponse{
		Err:     "invalid_client",
		ErrDesc: "Client authentication failed",
	}

	data := `{"error_description":"Client authentication failed","error":"invalid_client"}`
	if err := ar.Failed(http.StatusBadRequest, []byte(data)); err == nil {
		t.Errorf("Auth succeeded for %d", http.StatusBadRequest)
	}
}

func TestFailedAuth2(t *testing.T) {
	ar := authResponse{
		Err:     "invalid_client",
		ErrDesc: "Client authentication failed",
	}

	data := `{"error_description":"Client authentication failed","error":"invalid_client"}`
	if err := ar.Failed(http.StatusUnauthorized, []byte(data)); err == nil {
		t.Errorf("Auth succeeded for %d", http.StatusUnauthorized)
	}
}
