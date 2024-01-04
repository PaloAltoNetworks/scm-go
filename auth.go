package scm

import (
	"fmt"
	"net/http"

	"github.com/paloaltonetworks/scm-go/api"
)

type authResponse struct {
	Jwt       string `json:"access_token"`
	Scope     string `json:"scope"`
	Type      string `json:"token_type"`
	ExpiresIn int    `json:"expires_in"`

	Err     string `json:"error"`
	ErrDesc string `json:"error_description"`
}

func (a authResponse) Failed(code int, body []byte) error {
	switch code {
	case http.StatusOK:
		if a.Jwt == "" {
			return api.Response{
				StatusCode: code,
				Errors: []api.Error{{
					Message: "Auth successful, but not JWT found in response.",
					Details: string(body),
				}},
			}
		}
		return nil
	case http.StatusBadRequest, http.StatusUnauthorized:
		// Auth API actually returns 400 right now, not 401.  Checking for
		// both right now ensures that if they change this it doesn't suddenly
		// break the SDK.
		if a.Err != "" || a.ErrDesc != "" {
			return api.Response{
				StatusCode: code,
				Errors: []api.Error{{
					Message: fmt.Sprintf("%s: %s", a.Err, a.ErrDesc),
				}},
			}
		}
		return api.Response{
			StatusCode: code,
			Errors: []api.Error{{
				Message: "Unauthorized",
				Details: string(body),
			}},
		}
	}

	return api.Response{
		StatusCode: code,
		Errors: []api.Error{{
			Message: "Unknown error, enable receive logging for more info",
			Details: string(body),
		}},
	}
}
