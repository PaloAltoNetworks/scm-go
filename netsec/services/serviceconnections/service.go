package serviceconnections

// This code is automatically generated.
// Manual changes will be overwritten upon SDK generation.

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"strconv"
	"strings"

	"github.com/paloaltonetworks/scm-go/api"
	wugpput "github.com/paloaltonetworks/scm-go/netsec/schemas/service/connections"
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

/*
CreateInput handles input for the Create function.

ShortName: uIHLJPY
Parent chains:
* CreateInput

Args:

Param Folder (string, required): the Folder param. String can either be a specific string(`"Service Connections"`) or match this regex: `^[0-9a-zA-Z._\s-]{1,}$`. Default: `"Service Connections"`.

Param Request (wugpput.Config): the Request param.
*/
type CreateInput struct {
	Folder  string          `json:"folder"`
	Request *wugpput.Config `json:"request,omitempty"`
}

// Create creates the specified object.
//
// Method: post
// URI: /service-connections
func (c *Client) Create(ctx context.Context, input CreateInput) (wugpput.Config, error) {
	// Variables.
	var err error
	var ans wugpput.Config
	path := "/service-connections"

	// Query parameter handling.
	uv := url.Values{}
	uv.Set("folder", input.Folder)
	prefix, ok := Servers[c.client.GetHost()]
	if !ok {
		return ans, api.UnknownHostError
	}
	if prefix != "" {
		path = prefix + path
	}

	// Execute the command.
	_, err = c.client.Do(ctx, "POST", path, uv, input.Request, &ans)

	// Done.
	return ans, err
}

/*
ReadInput handles input for the Read function.

ShortName: uIHLJPY
Parent chains:
* ReadInput

Args:

Param Folder (string, required): the Folder param. String can either be a specific string(`"Service Connections"`) or match this regex: `^[0-9a-zA-Z._\s-]{1,}$`. Default: `"Service Connections"`.

Param Id (string, required): the Id param.
*/
type ReadInput struct {
	Folder string `json:"folder"`
	Id     string `json:"id"`
}

// Read returns the configuration of the specified object.
//
// Method: get
// URI: /service-connections/{id}
func (c *Client) Read(ctx context.Context, input ReadInput) (wugpput.Config, error) {
	// Variables.
	var err error
	var ans wugpput.Config
	path := "/service-connections/{id}"

	// Query parameter handling.
	uv := url.Values{}
	uv.Set("folder", input.Folder)

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
	_, err = c.client.Do(ctx, "GET", path, uv, nil, &ans)

	// Done.
	return ans, err
}

/*
UpdateInput handles input for the Update function.

ShortName: uIHLJPY
Parent chains:
* UpdateInput

Args:

Param Id (string, required): the Id param.

Param Request (wugpput.Config): the Request param.
*/
type UpdateInput struct {
	Id      string          `json:"id"`
	Request *wugpput.Config `json:"request,omitempty"`
}

// Update modifies the configuration of the given object.
//
// Method: put
// URI: /service-connections/{id}
func (c *Client) Update(ctx context.Context, input UpdateInput) (wugpput.Config, error) {
	// Variables.
	var err error
	var ans wugpput.Config
	path := "/service-connections/{id}"

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
	_, err = c.client.Do(ctx, "PUT", path, nil, input.Request, &ans)

	// Done.
	return ans, err
}

/*
DeleteInput handles input for the Delete function.

ShortName: uIHLJPY
Parent chains:
* DeleteInput

Args:

Param Id (string, required): the Id param.
*/
type DeleteInput struct {
	Id string `json:"id"`
}

// Delete removes the specified configuration.
//
// Method: delete
// URI: /service-connections/{id}
func (c *Client) Delete(ctx context.Context, input DeleteInput) (wugpput.Config, error) {
	// Variables.
	var err error
	var ans wugpput.Config
	path := "/service-connections/{id}"

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
	_, err = c.client.Do(ctx, "DELETE", path, nil, nil, &ans)

	// Done.
	return ans, err
}

/*
ListInput handles input for the List function.

ShortName: uIHLJPY
Parent chains:
* ListInput

Args:

Param Folder (string, required): the Folder param. String can either be a specific string(`"Service Connections"`) or match this regex: `^[0-9a-zA-Z._\s-]{1,}$`. Default: `"Service Connections"`.

Param Limit (int64): the Limit param. Default: `200`.

Param Name (string): the Name param.

Param Offset (int64): the Offset param. Default: `0`.
*/
type ListInput struct {
	Folder string  `json:"folder"`
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

Param Data ([]wugpput.Config): the Data param.

Param Limit (int64): the Limit param. Default: `200`.

Param Offset (int64): the Offset param. Default: `0`.

Param Total (int64): the Total param.
*/
type ListOutput struct {
	Data   []wugpput.Config `json:"data,omitempty"`
	Limit  *int64           `json:"limit,omitempty"`
	Offset *int64           `json:"offset,omitempty"`
	Total  *int64           `json:"total,omitempty"`
}

// List gets a list of objects back.
//
// Method: get
// URI: /service-connections
func (c *Client) List(ctx context.Context, input ListInput) (ListOutput, error) {
	// Variables.
	var err error
	var ans ListOutput
	path := "/service-connections"

	// Query parameter handling.
	uv := url.Values{}
	uv.Set("folder", input.Folder)
	if input.Limit != nil {
		uv.Set("limit", strconv.FormatInt(*input.Limit, 10))
	}
	if input.Offset != nil {
		uv.Set("offset", strconv.FormatInt(*input.Offset, 10))
	}
	if input.Name != nil {
		uv.Set("name", *input.Name)
	}
	prefix, ok := Servers[c.client.GetHost()]
	if !ok {
		return ans, api.UnknownHostError
	}
	if prefix != "" {
		path = prefix + path
	}

	// Optional: retrieve everything if limit is -1.
	if input.Limit != nil && *input.Limit == -1 {
		return c.listAll(ctx, input)
	}

	// Execute the command.
	_, err = c.client.Do(ctx, "GET", path, uv, nil, &ans)

	// Done.
	return ans, err
}

type listResponse struct {
	Output ListOutput
	Error  error
}

func (c *Client) listAll(ctx context.Context, input ListInput) (ListOutput, error) {
	var err error
	var ans ListOutput
	var items map[string]wugpput.Config
	maxLimit := int64(2000)
	everything := ListInput{
		Limit:  &maxLimit,
		Folder: input.Folder,
	}

	times := 0
	for {
		// Get the total number of things.
		ans, err = c.List(ctx, everything)
		if err != nil {
			return ans, err
		}
		if ans.Total == nil {
			return ans, fmt.Errorf("total was nil")
		} else if len(ans.Data) == int(*ans.Total) {
			return ans, nil
		}

		items = make(map[string]wugpput.Config)
		total := int(*ans.Total)
		numRetrievers := int(math.Ceil(float64(total) / float64(maxLimit)))
		responses := make(chan listResponse, numRetrievers)

		for i := 0; i < numRetrievers; i++ {
			offset := int64(int64(i) * maxLimit)
			ri := ListInput{
				Offset: &offset,
				Limit:  &maxLimit,
				Folder: input.Folder,
				Name:   input.Name,
			}
			go func() {
				rout, rerr := c.List(ctx, ri)
				responses <- listResponse{
					Output: rout,
					Error:  rerr,
				}
			}()
		}

		var totalChanged bool
		for i := 0; i < numRetrievers; i++ {
			resp := <-responses
			if resp.Error != nil {
				return resp.Output, resp.Error
			} else if totalChanged {
				continue
			}
			if resp.Output.Total == nil {
				return ListOutput{}, fmt.Errorf("total is nil")
			}
			if *resp.Output.Total != *ans.Total {
				totalChanged = true
				continue
			}
			for j := 0; j < len(resp.Output.Data); j++ {
				if resp.Output.Data[j].Id == nil {
					return ListOutput{}, fmt.Errorf("no ID present")
				}
				if _, ok := items[*resp.Output.Data[j].Id]; !ok {
					items[*resp.Output.Data[j].Id] = resp.Output.Data[j]
				}
			}
		}

		if !totalChanged && len(items) == total {
			break
		}

		times++
		if times >= 5 {
			return ListOutput{}, api.TooManyRetriesError
		}
	}

	endTotal := int64(len(items))
	listing := make([]wugpput.Config, 0, len(items))
	for key := range items {
		listing = append(listing, items[key])
	}

	negativeOne := int64(-1)
	return ListOutput{
		Data:  listing,
		Total: &endTotal,
		Limit: &negativeOne,
	}, nil
}
