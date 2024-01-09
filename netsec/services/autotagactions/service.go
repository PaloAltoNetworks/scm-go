package autotagactions

// This code is automatically generated.
// Manual changes will be overwritten upon SDK generation.

import (
	"context"
	"net/url"
	"strconv"

	"github.com/paloaltonetworks/scm-go/api"
	tephihM "github.com/paloaltonetworks/scm-go/netsec/schemas/auto/tag/actions"
)

// Servers specficiation.
var (
	Servers = map[string]string{
		"api.sase.paloaltonetworks.com":   "/sse/config/v1",
		"api.strata.paloaltonetworks.com": "/config/objects/v1",
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

ShortName: dvnOhnM
Parent chains:
* CreateInput

Args:

Param Request (tephihM.Config): the Request param.
*/
type CreateInput struct {
	Request *tephihM.Config `json:"request,omitempty"`
}

// Create creates the specified object.
//
// Method: post
// URI: /auto-tag-actions
func (c *Client) Create(ctx context.Context, input CreateInput) error {
	// Variables.
	var err error
	path := "/auto-tag-actions"
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

/*
ListInput handles input for the List function.

ShortName: dvnOhnM
Parent chains:
* ListInput

Args:

Param Limit (int64): the Limit param. Default: `200`.

Param Name (string): the Name param.

Param Offset (int64): the Offset param. Default: `0`.
*/
type ListInput struct {
	Limit  *int64  `json:"limit,omitempty"`
	Name   *string `json:"name,omitempty"`
	Offset *int64  `json:"offset,omitempty"`
}

/*
ListOutput handles output for the List function.

ShortName:
Parent chains:
* *Delayed*

Args:

Param Data ([]tephihM.Config): the Data param.

Param Limit (int64): the Limit param. Default: `200`.

Param Offset (int64): the Offset param. Default: `0`.

Param Total (int64): the Total param.
*/
type ListOutput struct {
	Data   []tephihM.Config `json:"data,omitempty"`
	Limit  *int64           `json:"limit,omitempty"`
	Offset *int64           `json:"offset,omitempty"`
	Total  *int64           `json:"total,omitempty"`
}

// List gets a list of objects back.
//
// Method: get
// URI: /auto-tag-actions
func (c *Client) List(ctx context.Context, input ListInput) (ListOutput, error) {
	// Variables.
	var err error
	var ans ListOutput
	path := "/auto-tag-actions"

	// Query parameter handling.
	uv := url.Values{}
	if input.Name != nil {
		uv.Set("name", *input.Name)
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

/*
AutoTagActionsDeleteInput handles input for the AutoTagActionsDelete function.

ShortName: dvnOhnM
Parent chains:
* AutoTagActionsDeleteInput

Args:

Param Name (string, required): the Name param.
*/
type AutoTagActionsDeleteInput struct {
	Name string `json:"name"`
}

// AutoTagActionsDelete performs the given operation.
//
// Method: delete
// URI: /auto-tag-actions
func (c *Client) AutoTagActionsDelete(ctx context.Context, input AutoTagActionsDeleteInput) error {
	// Variables.
	var err error
	path := "/auto-tag-actions"

	// Query parameter handling.
	uv := url.Values{}
	uv.Set("name", input.Name)
	prefix, ok := Servers[c.client.GetHost()]
	if !ok {
		return api.UnknownHostError
	}
	if prefix != "" {
		path = prefix + path
	}

	// Execute the command.
	_, err = c.client.Do(ctx, "DELETE", path, uv, nil, nil)

	// Done.
	return err
}

/*
AutoTagActionsPutInput handles input for the AutoTagActionsPut function.

ShortName: dvnOhnM
Parent chains:
* AutoTagActionsPutInput

Args:

Param Request (tephihM.Config): the Request param.
*/
type AutoTagActionsPutInput struct {
	Request *tephihM.Config `json:"request,omitempty"`
}

// AutoTagActionsPut performs the given operation.
//
// Method: put
// URI: /auto-tag-actions
func (c *Client) AutoTagActionsPut(ctx context.Context, input AutoTagActionsPutInput) error {
	// Variables.
	var err error
	path := "/auto-tag-actions"
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
