package certificatesimport

// This code is automatically generated.
// Manual changes will be overwritten upon SDK generation.

import (
	"context"
	"net/url"

	"github.com/paloaltonetworks/scm-go/api"
	wugpput "github.com/paloaltonetworks/scm-go/netsec/schemas/certificates/certimport"
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

ShortName: uIHLJPY
Parent chains:
* CreateInput

Args:

Param Device (string): the Device param.

Param Folder (string): the Folder param.

Param Request (wugpput.Config): the Request param.

Param Snippet (string): the Snippet param.
*/
type CreateInput struct {
	Device  *string         `json:"device,omitempty"`
	Folder  *string         `json:"folder,omitempty"`
	Request *wugpput.Config `json:"request,omitempty"`
	Snippet *string         `json:"snippet,omitempty"`
}

// Create creates the specified object.
//
// Method: post
// URI: /certificates:import
func (c *Client) Create(ctx context.Context, input CreateInput) error {
	// Variables.
	var err error
	path := "/certificates:import"

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
		return api.UnknownHostError
	}
	if prefix != "" {
		path = prefix + path
	}

	// Execute the command.
	_, err = c.client.Do(ctx, "POST", path, uv, input.Request, nil)

	// Done.
	return err
}
