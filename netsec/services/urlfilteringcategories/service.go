package urlfilteringcategories

// This code is automatically generated.
// Manual changes will be overwritten upon SDK generation.

import (
	"context"
	"net/url"
	"strconv"

	"github.com/paloaltonetworks/scm-go/api"
	gVcGYNO "github.com/paloaltonetworks/scm-go/netsec/schemas/url/filtering/categories"
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
ListInput handles input for the List function.

ShortName: wajgxii
Parent chains:
* ListInput

Args:

Param Device (string): the Device param.

Param Folder (string): the Folder param.

Param Limit (int64): the Limit param. Default: `200`.

Param Name (string): the Name param.

Param Offset (int64): the Offset param. Default: `0`.

Param Snippet (string): the Snippet param.
*/
type ListInput struct {
	Device  *string `json:"device,omitempty"`
	Folder  *string `json:"folder,omitempty"`
	Limit   *int64  `json:"limit,omitempty"`
	Name    *string `json:"name,omitempty"`
	Offset  *int64  `json:"offset,omitempty"`
	Snippet *string `json:"snippet,omitempty"`
}

/*
ListOutput handles output for the List function.

ShortName:
Parent chains:
* *Delayed*

Args:

Param Data ([]gVcGYNO.Config): the Data param.

Param Limit (int64): the Limit param. Default: `200`.

Param Offset (int64): the Offset param. Default: `0`.

Param Total (int64): the Total param.
*/
type ListOutput struct {
	Data   []gVcGYNO.Config `json:"data,omitempty"`
	Limit  *int64           `json:"limit,omitempty"`
	Offset *int64           `json:"offset,omitempty"`
	Total  *int64           `json:"total,omitempty"`
}

// List gets a list of objects back.
//
// Method: get
// URI: /url-filtering-categories
func (c *Client) List(ctx context.Context, input ListInput) (ListOutput, error) {
	// Variables.
	var err error
	var ans ListOutput
	path := "/url-filtering-categories"

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
