package urlcategories

// This code is automatically generated.
// Manual changes will be overwritten upon SDK generation.

import (
	"context"
	"net/url"
	"strconv"

	"github.com/paloaltonetworks/scm-go/api"
	cSTLpMW "github.com/paloaltonetworks/scm-go/netsec/schemas/url/categories"
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

ShortName: dPHRIQI
Parent chains:
* CreateInput

Args:

Param Device (string): the Device param.

Param Folder (string): the Folder param.

Param Request (cSTLpMW.Config): the Request param.

Param Snippet (string): the Snippet param.
*/
type CreateInput struct {
	Device  *string         `json:"device,omitempty"`
	Folder  *string         `json:"folder,omitempty"`
	Request *cSTLpMW.Config `json:"request,omitempty"`
	Snippet *string         `json:"snippet,omitempty"`
}

// Create creates the specified object.
//
// Method: post
// URI: /url-categories
func (c *Client) Create(ctx context.Context, input CreateInput) error {
	// Variables.
	var err error
	path := "/url-categories"

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

/*
ListInput handles input for the List function.

ShortName: dPHRIQI
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

Param Data ([]cSTLpMW.Config): the Data param.

Param Limit (int64): the Limit param. Default: `200`.

Param Offset (int64): the Offset param. Default: `0`.

Param Total (int64): the Total param.
*/
type ListOutput struct {
	Data   []cSTLpMW.Config `json:"data,omitempty"`
	Limit  *int64           `json:"limit,omitempty"`
	Offset *int64           `json:"offset,omitempty"`
	Total  *int64           `json:"total,omitempty"`
}

// List gets a list of objects back.
//
// Method: get
// URI: /url-categories
func (c *Client) List(ctx context.Context, input ListInput) (ListOutput, error) {
	// Variables.
	var err error
	var ans ListOutput
	path := "/url-categories"

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
