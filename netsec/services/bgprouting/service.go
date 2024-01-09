package bgprouting

// This code is automatically generated.
// Manual changes will be overwritten upon SDK generation.

import (
	"context"

	"github.com/paloaltonetworks/scm-go/api"
	suxdMuj "github.com/paloaltonetworks/scm-go/netsec/schemas/bgp/routing"
)

// Servers specficiation.
var (
	Servers = map[string]string{
		"api.sase.paloaltonetworks.com":   "/sse/config/v1",
		"api.strata.paloaltonetworks.com": "/config/deployment/v1",
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

// BgpRoutingGet performs the given operation.
//
// Method: get
// URI: /bgp-routing
func (c *Client) BgpRoutingGet(ctx context.Context) (suxdMuj.Config, error) {
	// Variables.
	var err error
	var ans suxdMuj.Config
	path := "/bgp-routing"
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
BgpRoutingPutInput handles input for the BgpRoutingPut function.

ShortName: sdhSKaQ
Parent chains:
* BgpRoutingPutInput

Args:

Param Request (suxdMuj.Config): the Request param.
*/
type BgpRoutingPutInput struct {
	Request *suxdMuj.Config `json:"request,omitempty"`
}

// BgpRoutingPut performs the given operation.
//
// Method: put
// URI: /bgp-routing
func (c *Client) BgpRoutingPut(ctx context.Context, input BgpRoutingPutInput) error {
	// Variables.
	var err error
	path := "/bgp-routing"
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
