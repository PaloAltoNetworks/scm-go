package quarantineddevices

// This code is automatically generated.
// Manual changes will be overwritten upon SDK generation.

import (
	"context"
	"net/url"

	"github.com/paloaltonetworks/scm-go/api"
	gmuDWVF "github.com/paloaltonetworks/scm-go/netsec/schemas/quarantined/devices"
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

ShortName: ooAAxrN
Parent chains:
* CreateInput

Args:

Param Request (gmuDWVF.Config): the Request param.
*/
type CreateInput struct {
	Request *gmuDWVF.Config `json:"request,omitempty"`
}

// Create creates the specified object.
//
// Method: post
// URI: /quarantined-devices
func (c *Client) Create(ctx context.Context, input CreateInput) error {
	// Variables.
	var err error
	path := "/quarantined-devices"
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
QuarantinedDevicesDeleteInput handles input for the QuarantinedDevicesDelete function.

ShortName: ooAAxrN
Parent chains:
* QuarantinedDevicesDeleteInput

Args:

Param HostId (string, required): the HostId param.
*/
type QuarantinedDevicesDeleteInput struct {
	HostId string `json:"host_id"`
}

// QuarantinedDevicesDelete performs the given operation.
//
// Method: delete
// URI: /quarantined-devices
func (c *Client) QuarantinedDevicesDelete(ctx context.Context, input QuarantinedDevicesDeleteInput) error {
	// Variables.
	var err error
	path := "/quarantined-devices"

	// Query parameter handling.
	uv := url.Values{}
	uv.Set("host_id", input.HostId)
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
QuarantinedDevicesGetInput handles input for the QuarantinedDevicesGet function.

ShortName: ooAAxrN
Parent chains:
* QuarantinedDevicesGetInput

Args:

Param HostId (string): the HostId param.

Param SerialNumber (string): the SerialNumber param.
*/
type QuarantinedDevicesGetInput struct {
	HostId       *string `json:"host_id,omitempty"`
	SerialNumber *string `json:"serial_number,omitempty"`
}

// QuarantinedDevicesGet performs the given operation.
//
// Method: get
// URI: /quarantined-devices
func (c *Client) QuarantinedDevicesGet(ctx context.Context, input QuarantinedDevicesGetInput) ([]gmuDWVF.Config, error) {
	// Variables.
	var err error
	var ans []gmuDWVF.Config
	path := "/quarantined-devices"

	// Query parameter handling.
	uv := url.Values{}
	if input.HostId != nil {
		uv.Set("host_id", *input.HostId)
	}
	if input.SerialNumber != nil {
		uv.Set("serial_number", *input.SerialNumber)
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
