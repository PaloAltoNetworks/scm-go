package idmove

// This code is automatically generated.
// Manual changes will be overwritten upon SDK generation.

import (
	"context"
	"strings"

	"github.com/paloaltonetworks/scm-go/api"
	urOXkAM "github.com/paloaltonetworks/scm-go/netsec/schemas/rule/based/move"
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

ShortName: vqtUUHF
Parent chains:
* CreateInput

Args:

Param Id (string, required): the Id param.

Param Request (urOXkAM.Config): the Request param.
*/
type CreateInput struct {
	Id      string          `json:"id"`
	Request *urOXkAM.Config `json:"request,omitempty"`
}

// Create creates the specified object.
//
// Method: post
// URI: /authentication-rules/{id}:move
func (c *Client) Create(ctx context.Context, input CreateInput) error {
	// Variables.
	var err error
	path := "/authentication-rules/{id}:move"

	// Path param handling.
	path = strings.ReplaceAll(path, "{id}", input.Id)
	prefix, ok := Servers[c.client.GetHost()]
	if !ok {
		return api.UnknownHostError
	}
	if prefix != "" {
		path = prefix + path
	}

	// Execute the command.
	_, err = c.client.Do(ctx, "POST", path, nil, input.Request, nil)

	// Done.
	return err
}
