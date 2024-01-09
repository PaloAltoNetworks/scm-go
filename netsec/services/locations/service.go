package locations

// This code is automatically generated.
// Manual changes will be overwritten upon SDK generation.

import (
	"context"

	"github.com/paloaltonetworks/scm-go/api"
	aeWshcf "github.com/paloaltonetworks/scm-go/netsec/schemas/locations"
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

// LocationsGet performs the given operation.
//
// Method: get
// URI: /locations
func (c *Client) LocationsGet(ctx context.Context) ([]aeWshcf.Config, error) {
	// Variables.
	var err error
	var ans []aeWshcf.Config
	path := "/locations"
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
