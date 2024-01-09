package mfaservers

// This code is automatically generated.
// Manual changes will be overwritten upon SDK generation.

import (
	"context"
	"net/url"
	"strconv"
	"strings"

	"github.com/paloaltonetworks/scm-go/api"
	oPPPeKY "github.com/paloaltonetworks/scm-go/netsec/schemas/mfa/servers"
)

// Servers specficiation.
var (
	Servers = map[string]string{
		"api.sase.paloaltonetworks.com":   "/sse/config/v1",
		"api.strata.paloaltonetworks.com": "/config/identity/v1",
	}
)

// Client is the client for the namespace.
type Client struct {
	client api.Client
}

// NewClient returns a new client for this namespace.
func NewClient(client api.Client) *Client {
	return &Client{client: client}
}

/*
CreateInput handles input for the Create function.

ShortName: jhtSIUK
Parent chains:
* CreateInput

Args:

Param Device (string): the Device param.

Param Folder (string): the Folder param.

Param Request (oPPPeKY.Config): the Request param.

Param Snippet (string): the Snippet param.
*/
type CreateInput struct {
	Device  *string         `json:"device,omitempty"`
	Folder  *string         `json:"folder,omitempty"`
	Request *oPPPeKY.Config `json:"request,omitempty"`
	Snippet *string         `json:"snippet,omitempty"`
}

// Create creates the specified object.
//
// Method: post
// URI: /mfa-servers
func (c *Client) Create(ctx context.Context, input CreateInput) (oPPPeKY.Config, error) {
	// Variables.
	var err error
	var ans oPPPeKY.Config
	path := "/mfa-servers"

	// Query parameter handling.
	uv := url.Values{}
	if input.Folder != nil {
		uv.Set("folder", *input.Folder)
	}
	if input.Snippet != nil {
		uv.Set("snippet", *input.Snippet)
	}
	if input.Device != nil {
		uv.Set("device", *input.Device)
	}
	prefix, ok := Servers[c.client.GetHost()]
	if !ok {
		return ans, api.UnknownHostError
	}
	if prefix != "" {
		path = prefix + path
	}

	// Execute the command.
	_, err = c.client.Do(ctx, "POST", path, uv, input.Request, &ans)

	// Done.
	return ans, err
}

/*
ReadInput handles input for the Read function.

ShortName: jhtSIUK
Parent chains:
* ReadInput

Args:

Param Id (string, required): the Id param.
*/
type ReadInput struct {
	Id string `json:"id"`
}

// Read returns the configuration of the specified object.
//
// Method: get
// URI: /mfa-servers/{id}
func (c *Client) Read(ctx context.Context, input ReadInput) (oPPPeKY.Config, error) {
	// Variables.
	var err error
	var ans oPPPeKY.Config
	path := "/mfa-servers/{id}"

	// Path param handling.
	path = strings.ReplaceAll(path, "{id}", input.Id)
	prefix, ok := Servers[c.client.GetHost()]
	if !ok {
		return ans, api.UnknownHostError
	}
	if prefix != "" {
		path = prefix + path
	}

	// Execute the command.
	_, err = c.client.Do(ctx, "GET", path, nil, nil, &ans)

	// Done.
	return ans, err
}

/*
UpdateInput handles input for the Update function.

ShortName: jhtSIUK
Parent chains:
* UpdateInput

Args:

Param Id (string, required): the Id param.

Param Request (oPPPeKY.Config): the Request param.
*/
type UpdateInput struct {
	Id      string          `json:"id"`
	Request *oPPPeKY.Config `json:"request,omitempty"`
}

// Update modifies the configuration of the given object.
//
// Method: put
// URI: /mfa-servers/{id}
func (c *Client) Update(ctx context.Context, input UpdateInput) (oPPPeKY.Config, error) {
	// Variables.
	var err error
	var ans oPPPeKY.Config
	path := "/mfa-servers/{id}"

	// Path param handling.
	path = strings.ReplaceAll(path, "{id}", input.Id)
	prefix, ok := Servers[c.client.GetHost()]
	if !ok {
		return ans, api.UnknownHostError
	}
	if prefix != "" {
		path = prefix + path
	}

	// Execute the command.
	_, err = c.client.Do(ctx, "PUT", path, nil, input.Request, &ans)

	// Done.
	return ans, err
}

/*
DeleteInput handles input for the Delete function.

ShortName: jhtSIUK
Parent chains:
* DeleteInput

Args:

Param Id (string, required): the Id param.
*/
type DeleteInput struct {
	Id string `json:"id"`
}

// Delete removes the specified configuration.
//
// Method: delete
// URI: /mfa-servers/{id}
func (c *Client) Delete(ctx context.Context, input DeleteInput) (oPPPeKY.Config, error) {
	// Variables.
	var err error
	var ans oPPPeKY.Config
	path := "/mfa-servers/{id}"

	// Path param handling.
	path = strings.ReplaceAll(path, "{id}", input.Id)
	prefix, ok := Servers[c.client.GetHost()]
	if !ok {
		return ans, api.UnknownHostError
	}
	if prefix != "" {
		path = prefix + path
	}

	// Execute the command.
	_, err = c.client.Do(ctx, "DELETE", path, nil, nil, &ans)

	// Done.
	return ans, err
}

/*
MfaServersGetInput handles input for the MfaServersGet function.

ShortName: jhtSIUK
Parent chains:
* MfaServersGetInput

Args:

Param Device (string): the Device param.

Param Folder (string): the Folder param.

Param Limit (int64): the Limit param. Default: `200`.

Param Name (string): the Name param.

Param Offset (int64): the Offset param. Default: `0`.

Param Position (string, required): the Position param. String must be one of these: `"pre"`, `"post"`. Default: `"pre"`.

Param Snippet (string): the Snippet param.
*/
type MfaServersGetInput struct {
	Device   *string `json:"device,omitempty"`
	Folder   *string `json:"folder,omitempty"`
	Limit    *int64  `json:"limit,omitempty"`
	Name     *string `json:"name,omitempty"`
	Offset   *int64  `json:"offset,omitempty"`
	Position string  `json:"position"`
	Snippet  *string `json:"snippet,omitempty"`
}

// MfaServersGet performs the given operation.
//
// Method: get
// URI: /mfa-servers
func (c *Client) MfaServersGet(ctx context.Context, input MfaServersGetInput) ([]oPPPeKY.Config, error) {
	// Variables.
	var err error
	var ans []oPPPeKY.Config
	path := "/mfa-servers"

	// Query parameter handling.
	uv := url.Values{}
	if input.Name != nil {
		uv.Set("name", *input.Name)
	}
	uv.Set("position", input.Position)
	if input.Folder != nil {
		uv.Set("folder", *input.Folder)
	}
	if input.Snippet != nil {
		uv.Set("snippet", *input.Snippet)
	}
	if input.Device != nil {
		uv.Set("device", *input.Device)
	}
	if input.Limit != nil {
		uv.Set("limit", strconv.FormatInt(*input.Limit, 10))
	}
	if input.Offset != nil {
		uv.Set("offset", strconv.FormatInt(*input.Offset, 10))
	}
	prefix, ok := Servers[c.client.GetHost()]
	if !ok {
		return ans, api.UnknownHostError
	}
	if prefix != "" {
		path = prefix + path
	}

	// Execute the command.
	_, err = c.client.Do(ctx, "GET", path, uv, nil, &ans)

	// Done.
	return ans, err
}
