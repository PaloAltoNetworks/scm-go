package jobs

// This code is automatically generated.
// Manual changes will be overwritten upon SDK generation.

import (
	"context"
	"strconv"
	"strings"

	"github.com/paloaltonetworks/scm-go/api"
	qFWAgJG "github.com/paloaltonetworks/scm-go/netsec/schemas/jobs"
)

// Servers specficiation.
var (
	Servers = map[string]string{
		"api.sase.paloaltonetworks.com":   "/sse/config/v1",
		"api.strata.paloaltonetworks.com": "/config/operations/v1",
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

ShortName: mIAatvm
Parent chains:
* ReadInput

Args:

Param Id (int64, required): the Id param.
*/
type ReadInput struct {
	Id int64 `json:"id"`
}

// Read returns the configuration of the specified object.
//
// Method: get
// URI: /jobs/{id}
func (c *Client) Read(ctx context.Context, input ReadInput) (qFWAgJG.Config, error) {
	// Variables.
	var err error
	var ans qFWAgJG.Config
	path := "/jobs/{id}"

	// Path param handling.
	path = strings.ReplaceAll(path, "{id}", strconv.FormatInt(input.Id, 10))
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
ListOutput handles output for the List function.

ShortName:
Parent chains:
* *Delayed*

Args:

Param Data ([]qFWAgJG.Config): the Data param.

Param Limit (int64): the Limit param. Default: `200`.

Param Offset (int64): the Offset param. Default: `0`.

Param Total (int64): the Total param.
*/
type ListOutput struct {
	Data   []qFWAgJG.Config `json:"data,omitempty"`
	Limit  *int64           `json:"limit,omitempty"`
	Offset *int64           `json:"offset,omitempty"`
	Total  *int64           `json:"total,omitempty"`
}

// List gets a list of objects back.
//
// Method: get
// URI: /jobs
func (c *Client) List(ctx context.Context) (ListOutput, error) {
	// Variables.
	var err error
	var ans ListOutput
	path := "/jobs"
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
