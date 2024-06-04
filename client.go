package scm

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/paloaltonetworks/scm-go/api"
)

/*
Client is the client connection to the SCM API.

There are multiple ways to specify the client's parameters.  If overlapping
values are configured for the client, then this is the resolution order:

1. Non-empty values for the param (explicitly defined).
2. Environment variables
3. Taken from the JSON config file

This resolution happens during `Setup()`.

The following is supported:

Param | Environment Variable | JSON Key | Default
-------------------------------------------------
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
			b, err = ioutil.ReadFile(c.AuthFile)
		}

		if err != nil {
			return err
		}

		if err = json.Unmarshal(b, &json_client); err != nil {
			return err
		}
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
	// Ensure that this is atomic.
	nv := atomic.AddInt32(&c.jwtAtomic, 1)
	defer atomic.AddInt32(&c.jwtAtomic, -1)
	if nv != 1 {
		for {
			if atomic.LoadInt32(&c.jwtAtomic) == nv-1 {
				break
			}
		}
		return nil
	}

	var resp *http.Response
	var err error
	var body []byte

	c.Log(ctx, "", "refreshing jwt")

	if len(c.testData) != 0 {
		// Testing.
		resp = c.testData[c.testIndex%len(c.testData)]
		c.testIndex++
	} else {
		authClient := &http.Client{
			Transport: c.Transport,
			Timeout:   time.Duration(10 * time.Second),
		}

		uv := url.Values{}
		uv.Set("scope", c.Scope)
		uv.Set("grant_type", "client_credentials")

		uri := "https://auth.apps.paloaltonetworks.com/auth/v1/oauth2/access_token"

		req, err := http.NewRequestWithContext(ctx, http.MethodPost, uri, strings.NewReader(uv.Encode()))

		if err != nil {
			return err
		}

		// Add in headers.
		req.SetBasicAuth(c.ClientId, c.ClientSecret)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		for k, v := range c.Headers {
			req.Header.Set(k, v)
		}

		resp, err = authClient.Do(req)
	}

	if resp == nil {
		return fmt.Errorf("no response")
	} else if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var aa authResponse
	if err = json.Unmarshal(body, &aa); err != nil {
		return err
	} else if err = aa.Failed(resp.StatusCode, body); err != nil {
		return err
	}

	c.Jwt = aa.Jwt

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
	body, err = ioutil.ReadAll(resp.Body)
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
	case http.StatusTooManyRequests, http.StatusBadGateway, http.StatusServiceUnavailable, http.StatusGatewayTimeout:
		// Sleep for a bit and try again.
		// TODO(shinmog): When/if this is implemented, verify backoff logic with eng.
		time.Sleep(time.Duration(len(retry)+1) * time.Second)
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
