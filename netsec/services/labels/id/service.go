package id

// This code is automatically generated.
// Manual changes will be overwritten upon SDK generation.

import (
	"context"
	"strings"

	"github.com/paloaltonetworks/scm-go/api"
	nodMrsO "github.com/paloaltonetworks/scm-go/netsec/schemas/labels/getbyid/response"
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
ReadInput handles input for the Read function.

ShortName: lhPcfTR
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
// URI: /labels/{id}
func (c *Client) Read(ctx context.Context, input ReadInput) (nodMrsO.Config, error) {
	// Variables.
	var err error
	var ans nodMrsO.Config
	path := "/labels/{id}"

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
