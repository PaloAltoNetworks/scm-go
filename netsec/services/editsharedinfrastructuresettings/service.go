package editsharedinfrastructuresettings

// This code is automatically generated.
// Manual changes will be overwritten upon SDK generation.

import (
	"context"

	"github.com/paloaltonetworks/scm-go/api"
	ctlHcHg "github.com/paloaltonetworks/scm-go/netsec/schemas/edit/shared/infrastructure/settings"
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
SharedInfrastructureSettingsPutInput handles input for the SharedInfrastructureSettingsPut function.

ShortName: xQJbYkW
Parent chains:
* SharedInfrastructureSettingsPutInput

Args:

Param Request (ctlHcHg.Config): the Request param.
*/
type SharedInfrastructureSettingsPutInput struct {
	Request *ctlHcHg.Config `json:"request,omitempty"`
}

// SharedInfrastructureSettingsPut performs the given operation.
//
// Method: put
// URI: /shared-infrastructure-settings
func (c *Client) SharedInfrastructureSettingsPut(ctx context.Context, input SharedInfrastructureSettingsPutInput) error {
	// Variables.
	var err error
	path := "/shared-infrastructure-settings"
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
