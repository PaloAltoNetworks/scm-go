package securityrules

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
	uewNibC "github.com/paloaltonetworks/scm-go/netsec/schemas/rule/based/move"
	uvXdTvM "github.com/paloaltonetworks/scm-go/netsec/schemas/security/rules"
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

ShortName: teqpOQC
Parent chains:
* CreateInput

Args:

Param Device (string): the Device param.

Param Folder (string): the Folder param.

Param Position (string, required): the Position param. String must be one of these: `"pre"`, `"post"`. Default: `"pre"`.

Param Request (uvXdTvM.Config): the Request param.

Param Snippet (string): the Snippet param.
*/
type CreateInput struct {
	Device   *string         `json:"device,omitempty"`
	Folder   *string         `json:"folder,omitempty"`
	Position string          `json:"position"`
	Request  *uvXdTvM.Config `json:"request,omitempty"`
	Snippet  *string         `json:"snippet,omitempty"`
}

// Create creates the specified object.
//
// Method: post
// URI: /security-rules
func (c *Client) Create(ctx context.Context, input CreateInput) (uvXdTvM.Config, error) {
	// Variables.
	var err error
	var ans uvXdTvM.Config
	path := "/security-rules"

	// Query parameter handling.
	uv := url.Values{}
	uv.Set("position", input.Position)
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

ShortName: teqpOQC
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
// URI: /security-rules/{id}
func (c *Client) Read(ctx context.Context, input ReadInput) (uvXdTvM.Config, error) {
	// Variables.
	var err error
	var ans uvXdTvM.Config
	path := "/security-rules/{id}"

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

/*
UpdateInput handles input for the Update function.

ShortName: teqpOQC
Parent chains:
* UpdateInput

Args:

Param Id (string, required): the Id param.

Param Request (uvXdTvM.Config): the Request param.
*/
type UpdateInput struct {
	Id      string          `json:"id"`
	Request *uvXdTvM.Config `json:"request,omitempty"`
}

// Update modifies the configuration of the given object.
//
// Method: put
// URI: /security-rules/{id}
func (c *Client) Update(ctx context.Context, input UpdateInput) (uvXdTvM.Config, error) {
	// Variables.
	var err error
	var ans uvXdTvM.Config
	path := "/security-rules/{id}"

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

ShortName: teqpOQC
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
// URI: /security-rules/{id}
func (c *Client) Delete(ctx context.Context, input DeleteInput) (uvXdTvM.Config, error) {
	// Variables.
	var err error
	var ans uvXdTvM.Config
	path := "/security-rules/{id}"

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

ShortName: teqpOQC
Parent chains:
* ListInput

Args:

Param Device (string): the Device param.

Param Folder (string): the Folder param.

Param Limit (int64): the Limit param. Default: `200`.

Param Name (string): the Name param.

Param Offset (int64): the Offset param. Default: `0`.

Param Position (string, required): the Position param. String must be one of these: `"pre"`, `"post"`. Default: `"pre"`.

Param Snippet (string): the Snippet param.
*/
type ListInput struct {
	Device   *string `json:"device,omitempty"`
	Folder   *string `json:"folder,omitempty"`
	Limit    *int64  `json:"limit,omitempty"`
	Name     *string `json:"name,omitempty"`
	Offset   *int64  `json:"offset,omitempty"`
	Position string  `json:"position"`
	Snippet  *string `json:"snippet,omitempty"`
}

/*
ListOutput handles output for the List function.

ShortName:
Parent chains:
* *Delayed*

Args:

Param Data ([]uvXdTvM.Config): the Data param.

Param Limit (int64): the Limit param. Default: `200`.

Param Offset (int64): the Offset param. Default: `0`.

Param Total (int64): the Total param.
*/
type ListOutput struct {
	Data   []uvXdTvM.Config `json:"data,omitempty"`
	Limit  *int64           `json:"limit,omitempty"`
	Offset *int64           `json:"offset,omitempty"`
	Total  *int64           `json:"total,omitempty"`
}

// List gets a list of objects back.
//
// Method: get
// URI: /security-rules
func (c *Client) List(ctx context.Context, input ListInput) (ListOutput, error) {
	// Variables.
	var err error
	var ans ListOutput
	path := "/security-rules"

	// Query parameter handling.
	uv := url.Values{}
	if input.Name != nil {
		uv.Set("name", *input.Name)
	}
	uv.Set("position", input.Position)
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
	var items map[string]uvXdTvM.Config
	maxLimit := int64(2000)
	everything := ListInput{
		Limit:    &maxLimit,
		Position: input.Position,
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

		items = make(map[string]uvXdTvM.Config)
		total := int(*ans.Total)
		numRetrievers := int(math.Ceil(float64(total) / float64(maxLimit)))
		responses := make(chan listResponse, numRetrievers)

		for i := 0; i < numRetrievers; i++ {
			offset := int64(int64(i) * maxLimit)
			ri := ListInput{
				Offset:   &offset,
				Limit:    &maxLimit,
				Name:     input.Name,
				Position: input.Position,
				Folder:   input.Folder,
				Snippet:  input.Snippet,
				Device:   input.Device,
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
	listing := make([]uvXdTvM.Config, 0, len(items))
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

/*
MoveInput handles input for the Move function.

ShortName: teqpOQC
Parent chains:
* MoveInput

Args:

Param Id (string, required): the Id param.

Param Request (uewNibC.Config): the Request param.
*/
type MoveInput struct {
	Id      string          `json:"id"`
	Request *uewNibC.Config `json:"request,omitempty"`
}

// Move creates the specified object.
//
// Method: post
// URI: /security-rules/{id}:move
func (c *Client) Move(ctx context.Context, input MoveInput) error {
	// Variables.
	var err error
	path := "/security-rules/{id}:move"

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
