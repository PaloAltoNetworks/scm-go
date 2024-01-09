package decryptionexclusions

// This code is automatically generated.
// Manual changes will be overwritten upon SDK generation.

import (
	"context"
	"net/url"
	"strconv"
	"strings"

	"github.com/paloaltonetworks/scm-go/api"
	xdEvbZX "github.com/paloaltonetworks/scm-go/netsec/schemas/decryption/exclusions"
)

// Servers specficiation.
var (
	Servers = map[string]string{
		"api.sase.paloaltonetworks.com":   "/sse/config/v1",
		"api.strata.paloaltonetworks.com": "/config/security/v1",
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

ShortName: wRodOhd
Parent chains:
* CreateInput

Args:

Param Device (string): the Device param.

Param Folder (string): the Folder param.

Param Request (xdEvbZX.Config): the Request param.

Param Snippet (string): the Snippet param.
*/
type CreateInput struct {
	Device  *string         `json:"device,omitempty"`
	Folder  *string         `json:"folder,omitempty"`
	Request *xdEvbZX.Config `json:"request,omitempty"`
	Snippet *string         `json:"snippet,omitempty"`
}

// Create creates the specified object.
//
// Method: post
// URI: /decryption-exclusions
func (c *Client) Create(ctx context.Context, input CreateInput) (xdEvbZX.Config, error) {
	// Variables.
	var err error
	var ans xdEvbZX.Config
	path := "/decryption-exclusions"

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

ShortName: wRodOhd
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
// URI: /decryption-exclusions/{id}
func (c *Client) Read(ctx context.Context, input ReadInput) (xdEvbZX.Config, error) {
	// Variables.
	var err error
	var ans xdEvbZX.Config
	path := "/decryption-exclusions/{id}"

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

ShortName: wRodOhd
Parent chains:
* UpdateInput

Args:

Param Id (string, required): the Id param.

Param Request (xdEvbZX.Config): the Request param.
*/
type UpdateInput struct {
	Id      string          `json:"id"`
	Request *xdEvbZX.Config `json:"request,omitempty"`
}

// Update modifies the configuration of the given object.
//
// Method: put
// URI: /decryption-exclusions/{id}
func (c *Client) Update(ctx context.Context, input UpdateInput) (xdEvbZX.Config, error) {
	// Variables.
	var err error
	var ans xdEvbZX.Config
	path := "/decryption-exclusions/{id}"

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

ShortName: wRodOhd
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
// URI: /decryption-exclusions/{id}
func (c *Client) Delete(ctx context.Context, input DeleteInput) (xdEvbZX.Config, error) {
	// Variables.
	var err error
	var ans xdEvbZX.Config
	path := "/decryption-exclusions/{id}"

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
DecryptionExclusionsGetInput handles input for the DecryptionExclusionsGet function.

ShortName: wRodOhd
Parent chains:
* DecryptionExclusionsGetInput

Args:

Param Device (string): the Device param.

Param Folder (string): the Folder param.

Param Limit (int64): the Limit param. Default: `200`.

Param Name (string): the Name param.

Param Offset (int64): the Offset param. Default: `0`.

Param Snippet (string): the Snippet param.
*/
type DecryptionExclusionsGetInput struct {
	Device  *string `json:"device,omitempty"`
	Folder  *string `json:"folder,omitempty"`
	Limit   *int64  `json:"limit,omitempty"`
	Name    *string `json:"name,omitempty"`
	Offset  *int64  `json:"offset,omitempty"`
	Snippet *string `json:"snippet,omitempty"`
}

// DecryptionExclusionsGet performs the given operation.
//
// Method: get
// URI: /decryption-exclusions
func (c *Client) DecryptionExclusionsGet(ctx context.Context, input DecryptionExclusionsGetInput) ([]xdEvbZX.Config, error) {
	// Variables.
	var err error
	var ans []xdEvbZX.Config
	path := "/decryption-exclusions"

	// Query parameter handling.
	uv := url.Values{}
	if input.Name != nil {
		uv.Set("name", *input.Name)
	}
	if input.Folder != nil {
		uv.Set("folder", *input.Folder)
	}
	if input.Snippet != nil {
		uv.Set("snippet", *input.Snippet)
	}
	if input.Device != nil {
		uv.Set("device", *input.Device)
	}
	if input.Offset != nil {
		uv.Set("offset", strconv.FormatInt(*input.Offset, 10))
	}
	if input.Limit != nil {
		uv.Set("limit", strconv.FormatInt(*input.Limit, 10))
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
