package move

// This code is automatically generated.
// Manual changes will be overwritten upon SDK generation.

import (
	"context"
	"strings"

	"github.com/paloaltonetworks/scm-go/api"
	ivVDSwf "github.com/paloaltonetworks/scm-go/netsec/schemas/folder/move/payload"
)

// Servers specficiation.
var (
	Servers = map[string]string{
		"api.strata.paloaltonetworks.com": "/config/setup/v1",
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
MovePutInput handles input for the MovePut function.

ShortName: uIHLJPY
Parent chains:
* MovePutInput

Args:

Param Id (string, required): the Id param.

Param Request (ivVDSwf.Config): the Request param.
*/
type MovePutInput struct {
	Id      string          `json:"id"`
	Request *ivVDSwf.Config `json:"request,omitempty"`
}

// MovePut performs the given operation.
//
// Method: put
// URI: /folders/{id}/move
func (c *Client) MovePut(ctx context.Context, input MovePutInput) error {
	// Variables.
	var err error
	path := "/folders/{id}/move"

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
	_, err = c.client.Do(ctx, "PUT", path, nil, input.Request, nil)

	// Done.
	return err
}
