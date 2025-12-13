package scm

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/paloaltonetworks/scm-go/api"
	retry "github.com/sethvargo/go-retry"
)

/*
Client is the client connection to the SCM API.

There are multiple ways to specify the client's parameters.  If overlapping
values are configured for the client, then this is the resolution order:

1. Non-empty values for the param (explicitly defined).
2. Environment variables
3. Taken from the JSON config file

This resolution happens during Setup().

The following is supported:

Param | Environment Variable | JSON Key | Default
-------------------------------------------------
AuthUrl | SCM_AUTH_URL | auth_url | "https://auth.apps.paloaltonetworks.com/auth/v1/oauth2/access_token"
Host | SCM_HOST | host | "api.strata.paloaltonetworks.com"
Port | SCM_PORT | port | 0
ClientId | SCM_CLIENT_ID | client_id | ""
ClientSecret | SCM_CLIENT_SECRET | client_secret | ""
Scope | SCM_SCOPE | scope | ""
Protocol | SCM_PROTOCOL | protocol | "https"
Headers | SCM_HEADERS | headers | nil
Agent | - | agent | ""
SkipVerifyCertificate | SCM_SKIP_VERIFY_CERTIFICATE | skip_verify_certificate | false
Logging | SCM_LOGGING | logging | "quiet"
SkipLoggingTransport | - | skip_logging_transport | false
*/
type Client struct {
	AuthUrl      string            `json:"auth_url"`
	Host         string            `json:"host"`
	Port         int               `json:"port"`
	ClientId     string            `json:"client_id"`
	ClientSecret string            `json:"client_secret"`
	Scope        string            `json:"scope"`
	Protocol     string            `json:"protocol"`
	Headers      map[string]string `json:"headers"`
	Agent        string            `json:"agent"`

	AuthFile         string `json:"-"`
	CheckEnvironment bool   `json:"-"`

	SkipVerifyCertificate bool            `json:"skip_verify_certificate"`
	Transport             *http.Transport `json:"-"`

	SkipLoggingTransport bool       `json:"skip_logging_transport"`
	Logging              string     `json:"logging"`
	Logger               api.Logger `json:"-"`

	Jwt       string `json:"-"`
	jwtAtomic int32  `json:"-"`

	JwtExpiresAt time.Time `json:"-"` // The actual time the JWT will expire
	JwtLifetime  int64     `json:"-"` // The TTL received from the auth server (in seconds)

	apiPrefix string

	HttpClient *http.Client

	testData        []*http.Response
	testIndex       int
	authFileContent []byte
}

// Setup configures the HttpClient param according to the combination of locally
// defined params, environment variables, and the JSON config file.
func (c *Client) Setup() error {
	var err error

	// Load up the JSON config file.
	var json_client Client
	if c.AuthFile != "" {
		var b []byte
		if len(c.testData) != 0 {
			b, err = c.authFileContent, nil
		} else {
			b, err = os.ReadFile(c.AuthFile)
		}

		if err != nil {
			return err
		}

		if err = json.Unmarshal(b, &json_client); err != nil {
			return err
		}
	}

	// AuthUrl.
	if c.AuthUrl == "" {
		if val := os.Getenv("SCM_AUTH_URL"); c.CheckEnvironment && val != "" {
			c.AuthUrl = val
		} else if json_client.AuthUrl != "" {
			c.AuthUrl = json_client.AuthUrl
		}
	}
	if c.AuthUrl == "" {
		c.AuthUrl = "https://auth.apps.paloaltonetworks.com/auth/v1/oauth2/access_token"
	}
	if !strings.HasPrefix(c.AuthUrl, "http://") && !strings.HasPrefix(c.AuthUrl, "https://") {
		return fmt.Errorf("AuthUrl should start with http:// or https://")
	}

	// Host.
	if c.Host == "" {
		if val := os.Getenv("SCM_HOST"); c.CheckEnvironment && val != "" {
			c.Host = val
		} else if json_client.Host != "" {
			c.Host = json_client.Host
		}
	}
	if c.Host == "" {
		c.Host = "api.strata.paloaltonetworks.com"
	}

	// Port.
	if c.Port == 0 {
		if val := os.Getenv("SCM_PORT"); c.CheckEnvironment && val != "" {
			if ival, err := strconv.Atoi(val); err != nil {
				return fmt.Errorf("Failed to parse port env var as int: %s", err)
			} else {
				c.Port = ival
			}
		} else if json_client.Port != 0 {
			c.Port = json_client.Port
		}
	}
	if c.Port < 0 || c.Port > 65535 {
		return fmt.Errorf("Port is outside the valid port range: %d", c.Port)
	}

	// Client ID.
	if c.ClientId == "" {
		if val := os.Getenv("SCM_CLIENT_ID"); c.CheckEnvironment && val != "" {
			c.ClientId = val
		} else if json_client.ClientId != "" {
			c.ClientId = json_client.ClientId
		} else {
			return fmt.Errorf("ClientId must be specified")
		}
	}

	// Client secret.
	if c.ClientSecret == "" {
		if val := os.Getenv("SCM_CLIENT_SECRET"); c.CheckEnvironment && val != "" {
			c.ClientSecret = val
		} else if json_client.ClientSecret != "" {
			c.ClientSecret = json_client.ClientSecret
		} else {
			return fmt.Errorf("ClientSecret must be specified")
		}
	}

	// Scope.
	if c.Scope == "" {
		if val := os.Getenv("SCM_SCOPE"); c.CheckEnvironment && val != "" {
			c.Scope = val
		} else if json_client.Scope != "" {
			c.Scope = json_client.Scope
		} else {
			return fmt.Errorf("Scope must be specified")
		}
	}

	// Protocol.
	if c.Protocol == "" {
		if val := os.Getenv("SCM_PROTOCOL"); c.CheckEnvironment && val != "" {
			c.Protocol = val
		} else if json_client.Protocol != "" {
			c.Protocol = json_client.Protocol
		} else {
			c.Protocol = "https"
		}
	}

	// Headers.
	if len(c.Headers) == 0 {
		if val := os.Getenv("SCM_HEADERS"); c.CheckEnvironment && val != "" {
			if err := json.Unmarshal([]byte(val), &c.Headers); err != nil {
				return err
			}
		}
		if len(c.Headers) == 0 && len(json_client.Headers) > 0 {
			c.Headers = make(map[string]string)
			for k, v := range json_client.Headers {
				c.Headers[k] = v
			}
		}
	}

	// Skip verify certificate.
	if !c.SkipVerifyCertificate {
		if val := os.Getenv("SCM_SKIP_VERIFY_CERTIFICATE"); c.CheckEnvironment && val != "" {
			if vcb, err := strconv.ParseBool(val); err != nil {
				return err
			} else if vcb {
				c.SkipVerifyCertificate = vcb
			}
		}
		if !c.SkipVerifyCertificate && json_client.SkipVerifyCertificate {
			c.SkipVerifyCertificate = json_client.SkipVerifyCertificate
		}
	}

	// Logging.
	if c.Logging == "" {
		if val := os.Getenv("SCM_LOGGING"); c.CheckEnvironment && val != "" {
			c.Logging = val
		} else if json_client.Logging != "" {
			c.Logging = json_client.Logging
		} else {
			c.Logging = api.LogQuiet
		}
	}

	// Setup the https client.
	if c.Transport == nil {
		c.Transport = &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: c.SkipVerifyCertificate,
			},
		}
	}
	c.HttpClient = &http.Client{
		Transport: c.Transport,
	}

	// Attach logging transport.
	if !c.SkipLoggingTransport && !json_client.SkipLoggingTransport {
		c.HttpClient.Transport = api.NewTransport(c.HttpClient.Transport, c)
	}

	// Configure the uri prefix.
	if c.Port != 0 {
		c.apiPrefix = fmt.Sprintf("%s://%s:%d", c.Protocol, c.Host, c.Port)
	} else {
		c.apiPrefix = fmt.Sprintf("%s://%s", c.Protocol, c.Host)
	}

	return nil
}

// RefreshJwt refreshes the JWT necessary to interact with the API.
//
// This function is atomic (only one may be running at any given time).
func (c *Client) RefreshJwt(ctx context.Context) error {
	c.Log(ctx, "", "=== RefreshJwt() CALLED ===")

	// Ensure that this is atomic.
	nv := atomic.AddInt32(&c.jwtAtomic, 1)
	defer atomic.AddInt32(&c.jwtAtomic, -1)
	c.Log(ctx, "", fmt.Sprintf("RefreshJwt: Atomic counter value: %d", nv))

	if nv != 1 {
		c.Log(ctx, "", "RefreshJwt: Another refresh in progress, waiting...")
		for {
			if atomic.LoadInt32(&c.jwtAtomic) == nv-1 {
				break
			}
		}
		c.Log(ctx, "", "RefreshJwt: Wait completed, returning")
		return nil
	}

	var resp *http.Response
	var err error
	var body []byte

	c.Log(ctx, "", "RefreshJwt: Starting JWT refresh process")

	if len(c.testData) != 0 {
		// Testing.
		c.Log(ctx, "", "RefreshJwt: Using test data")
		resp = c.testData[c.testIndex%len(c.testData)]
		c.testIndex++
	} else {
		c.Log(ctx, "", "RefreshJwt: Creating auth client")
		authClient := &http.Client{
			Transport: c.Transport,
			Timeout:   time.Duration(30 * time.Second),
		}

		uv := url.Values{}
		uv.Set("scope", c.Scope)
		uv.Set("grant_type", "client_credentials")

		// Define backoff strategy
		backoff := retry.NewExponential(1 * time.Second)
		backoff = retry.WithCappedDuration(10*time.Second, backoff)
		backoff = retry.WithJitter(500*time.Millisecond, backoff)
		backoff = retry.WithMaxRetries(5, backoff)
		c.Log(ctx, "", "RefreshJwt: Configured exponential backoff strategy")

		var aa authResponse

		// Use retry.Do to wrap the request logic
		c.Log(ctx, "", "RefreshJwt: Starting retry.Do loop")
		err = retry.Do(ctx, backoff, func(ctx context.Context) error {
			c.Log(ctx, "", "RefreshJwt: Making auth request attempt")
			req, reqErr := http.NewRequestWithContext(ctx, http.MethodPost, c.AuthUrl, strings.NewReader(uv.Encode()))
			if reqErr != nil {
				c.Log(ctx, "", fmt.Sprintf("RefreshJwt: Failed to create request: %v", reqErr))
				// This is a permanent error (e.g., bad URL), so just return it.
				return reqErr
			}

			// Add in headers.
			req.SetBasicAuth(c.ClientId, c.ClientSecret)
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			for k, v := range c.Headers {
				req.Header.Set(k, v)
			}
			c.Log(ctx, "", "RefreshJwt: Headers set, making HTTP request")

			resp, doErr := authClient.Do(req)
			if doErr != nil {
				// Network error, timeout, etc. This is retryable.
				c.Log(ctx, api.LogBasic, fmt.Sprintf("RefreshJwt: Auth request failed, will retry: %s", doErr))
				return retry.RetryableError(doErr)
			}

			c.Log(ctx, "", fmt.Sprintf("RefreshJwt: Got response with status code: %d", resp.StatusCode))

			// Check for 5xx server errors, which are retryable
			if resp.StatusCode >= 500 {
				resp.Body.Close() // Must close body to avoid leaks
				c.Log(ctx, api.LogBasic, fmt.Sprintf("RefreshJwt: Auth request failed with status %d, will retry", resp.StatusCode))
				return retry.RetryableError(fmt.Errorf("auth server returned %d", resp.StatusCode))
			}

			// --- At this point, any error is considered permanent ---
			defer resp.Body.Close()
			body, readErr := io.ReadAll(resp.Body)
			if readErr != nil {
				c.Log(ctx, "", fmt.Sprintf("RefreshJwt: Failed to read response body: %v", readErr))
				// Failed to read body, permanent error
				return readErr
			}

			c.Log(ctx, "", fmt.Sprintf("RefreshJwt: Response body: %s", string(body)))

			if err := json.Unmarshal(body, &aa); err != nil {
				c.Log(ctx, "", fmt.Sprintf("RefreshJwt: Failed to unmarshal response: %v", err))
				// Failed to parse response, permanent error
				return err
			}

			if err := aa.Failed(resp.StatusCode, body); err != nil {
				c.Log(ctx, "", fmt.Sprintf("RefreshJwt: Auth response indicates failure: %v", err))
				// Auth failure (e.g., 401, 403) is a permanent error
				return err
			}

			// Success!
			c.Log(ctx, "", "RefreshJwt: Auth request successful!")
			return nil
		})

		if err != nil {
			c.Log(ctx, "", fmt.Sprintf("RefreshJwt: retry.Do failed after all attempts: %v", err))
			return err // Will be the last error (if retries failed) or the permanent error
		}

		c.Log(ctx, "", "RefreshJwt: Setting new JWT token")
		c.Jwt = aa.Jwt
		// Set the lifetime and calculated expiry time (new lines)
		c.JwtLifetime = int64(aa.ExpiresIn)
		c.Log(ctx, "", fmt.Sprintf("RefreshJwt: Expires In: %d", c.JwtLifetime))
		// Set expiry to the current time plus the token lifetime, minus a 60 second buffer
		c.JwtExpiresAt = time.Now().Add(time.Duration(c.JwtLifetime)*time.Second - 60*time.Second)
		c.Log(ctx, "", "JWT Expiry Set - ExpiresAt: "+c.JwtExpiresAt.Format("15:04:05"))

		c.Log(ctx, "", "=== RefreshJwt() COMPLETED SUCCESSFULLY ===")
		return nil
	}

	// This part is for the testData block
	c.Log(ctx, "", "RefreshJwt: Processing test data response")
	if resp == nil {
		c.Log(ctx, "", "RefreshJwt: No response received")
		return fmt.Errorf("no response")
	} else if err != nil {
		c.Log(ctx, "", fmt.Sprintf("RefreshJwt: Error in test data: %v", err))
		return err
	}

	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		c.Log(ctx, "", fmt.Sprintf("RefreshJwt: Failed to read test response body: %v", err))
		return err
	}

	var aa authResponse
	if err = json.Unmarshal(body, &aa); err != nil {
		c.Log(ctx, "", fmt.Sprintf("RefreshJwt: Failed to unmarshal test response: %v", err))
		return err
	} else if err = aa.Failed(resp.StatusCode, body); err != nil {
		c.Log(ctx, "", fmt.Sprintf("RefreshJwt: Test auth response failed: %v", err))
		return err
	}

	c.Jwt = aa.Jwt
	// Set the lifetime and calculated expiry time (new lines)
	c.JwtLifetime = int64(aa.ExpiresIn)
	c.JwtExpiresAt = time.Now().Add(time.Duration(c.JwtLifetime)*time.Second - 60*time.Second)
	return nil
}

// LoggingIsSetTo checks if the logging is configured as the given string.
func (c *Client) LoggingIsSetTo(v string) bool {
	return c.Logging == v
}

// Log outputs API actions.
func (c *Client) Log(ctx context.Context, level, msg string) {
	if c.Logging == api.LogQuiet {
		return
	}

	if level == "" || c.Logging == level {
		if c.Logger == nil {
			log.Printf(msg)
		} else {
			c.Logger(ctx, msg)
		}
	}
}

/*
Do performs the given API request.

Param method should be one of the http.Method constants.

Param path should be a slice of path parts that will be joined together with
the base apiPrefix to create the final API endpoint.

Param queryParams are the query params that should be appended to the API URL.

Param input is an interface that can be passed in to json.Marshal() to send to
the API.

Param output is a pointer to a struct that will be filled with json.Unmarshal().

This function returns the content of the body from the API call and any errors
that may have been present.  If this function got all the way to invoking the
API and getting a response, then the error passed back will be a `api.ErrorResponse`
if an error was detected.
*/
func (c *Client) Do(ctx context.Context, method string, path string, queryParams url.Values, input, output interface{}, retry ...error) ([]byte, error) {
	if c.apiPrefix == "" {
		return nil, fmt.Errorf("Setup() has not been invoked yet")
	} else if len(retry) > 5 {
		return nil, retry[len(retry)-1]
	}

	// Refresh token if it expires or empty
	if c.Jwt != "" && time.Now().After(c.JwtExpiresAt) {
		c.Log(ctx, api.LogBasic, "JWT expired or near expiry, attempting proactive refresh.")
		// The Jwt check is atomic, so it's safe to call here.
		if err := c.RefreshJwt(ctx); err != nil {
			return nil, fmt.Errorf("failed to proactively refresh JWT: %w", err)
		}
	}

	var err error
	var body, data []byte
	var resp *http.Response
	var qp string

	// Convert input into JSON.
	if input != nil {
		data, err = json.Marshal(input)
		if err != nil {
			return nil, err
		}
	}

	if len(queryParams) > 0 {
		qp = fmt.Sprintf("?%s", queryParams.Encode())
	}

	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	uri := fmt.Sprintf("%s%s%s", c.apiPrefix, path, qp)
	c.Log(ctx, api.LogBasic, fmt.Sprintf("[%s] %s", method, uri))

	if len(c.testData) != 0 {
		// Testing.
		resp = c.testData[c.testIndex%len(c.testData)]
		c.testIndex++
	} else {
		req, err := http.NewRequestWithContext(ctx, method, uri, strings.NewReader(string(data)))
		if err != nil {
			return nil, err
		}

		// Configure headers.
		req.Header.Set("Content-Type", "application/json")
		if c.Agent != "" {
			req.Header.Set("User-Agent", c.Agent)
		}
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Jwt))
		req.Header.Set("Accept", "application/json")
		req.Header.Set("x-auth-jwt", c.Jwt)
		for k, v := range c.Headers {
			req.Header.Set(k, v)
		}

		resp, err = c.HttpClient.Do(req)
	}

	if err != nil {
		return nil, err
	} else if resp == nil {
		return nil, fmt.Errorf("no response received")
	}

	// Read the body content.
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Discover if an error occurred.
	stat := api.NewResponse(resp.StatusCode, body)

	/*
	   2023/01/26 02:15:45 [HTTP 404] API_I00013 Your configuration is not valid. Please review the error message for more details. - map[errorType:Object Not Present message:Failed to find obj-uuid for command get]

	   {
	       "_errors":[
	           {
	               "code":"API_I00013",
	               "message":"Your configuration is not valid. Please review the error message for more details.",
	               "details":{
	                   "errorType":"Object Not Present",
	                   "message":"Failed to find obj-uuid for command get",
	               },
	           },
	       ],
	       "_request_id":"93cdac8f-5bfe-4438-8ae3-3744114223a7",
	   }
	*/
	switch resp.StatusCode {
	case http.StatusOK, http.StatusCreated, http.StatusAccepted:
	case http.StatusNotFound:
		return body, api.ObjectNotFoundError
	case http.StatusUnauthorized:
		if len(retry) > 0 {
			lastErr, ok := retry[len(retry)-1].(api.Response)
			if ok && lastErr.StatusCode == http.StatusUnauthorized {
				// Getting 401s back-to-back, so just stop.
				return body, stat
			}
		}

		// First auth failure, so refresh the JWT then retry the operation.
		if err = c.RefreshJwt(ctx); err != nil {
			return nil, err
		}
		return c.Do(ctx, method, path, queryParams, input, output, append(retry, stat)...)
	case http.StatusTooManyRequests, http.StatusInternalServerError, http.StatusBadGateway, http.StatusServiceUnavailable, http.StatusGatewayTimeout:
		// When these errors are encountered, we should be sleeping and then retrying again:
		// https://pan.dev/prisma-cloud/api/cspm/api-errors/#reattempting-requests-that-fail-due-to-a-server-error
		// TODO(shinmog): When/if this is implemented, verify backoff logic with eng.

		// Only sleep if we're not running tests.
		if len(c.testData) == 0 {
			time.Sleep(time.Duration(len(retry)+1*2) * time.Second)
		}
		return c.Do(ctx, method, path, queryParams, input, output, append(retry, stat)...)
	default:
		return body, stat
	}

	// Optional: unmarshal the output.
	if output != nil {
		if err = json.Unmarshal(body, output); err != nil {
			return body, err
		}
	}

	// Done.
	return body, nil
}

// GetHost returns the Host property.
func (c *Client) GetHost() string { return c.Host }

// authResponse represents the response from the auth endpoint
type authResponse struct {
	Jwt       string `json:"access_token"`
	TokenType string `json:"token_type"`
	ExpiresIn int    `json:"expires_in"`
	Scope     string `json:"scope"`
	Error     string `json:"error"`
	ErrorDesc string `json:"error_description"`
}

// Failed checks if the auth response indicates a failure
func (a *authResponse) Failed(statusCode int, body []byte) error {
	if statusCode >= 200 && statusCode < 300 {
		return nil
	}

	if a.Error != "" {
		return fmt.Errorf("auth error: %s - %s", a.Error, a.ErrorDesc)
	}

	return fmt.Errorf("auth failed with status %d: %s", statusCode, string(body))
}

// JWTRefreshTransport is a RoundTripper that automatically handles JWT token refresh
type JWTRefreshTransport struct {
	Wrapped     http.RoundTripper
	SetupClient *Client // This refers to the Client struct from client.go in this same repo
}

// RoundTrip implements http.RoundTripper interface
func (j *JWTRefreshTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Ensure we have a wrapped transport
	if j.Wrapped == nil {
		j.Wrapped = http.DefaultTransport
	}

	// Check if JWT needs refresh (similar logic to client.go Do method)
	if j.SetupClient.Jwt != "" && time.Now().After(j.SetupClient.JwtExpiresAt) {
		fmt.Printf("Refreshing tokens\n")
		// JWT expired or near expiry, attempt proactive refresh
		if err := j.SetupClient.RefreshJwt(req.Context()); err != nil {
			// If refresh fails, continue with existing token and let the request handle the auth failure
			return nil, err
		}
	}

	req.Header.Set("Authorization", "Bearer "+j.SetupClient.Jwt)
	req.Header.Set("x-auth-jwt", j.SetupClient.Jwt)

	// Execute the request
	resp, err := j.Wrapped.RoundTrip(req)

	return resp, err
}
