// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/moira-alert/client-go/pkg/client/config"
	"github.com/moira-alert/client-go/pkg/client/contact"
	"github.com/moira-alert/client-go/pkg/client/event"
	"github.com/moira-alert/client-go/pkg/client/health"
	"github.com/moira-alert/client-go/pkg/client/notification"
	"github.com/moira-alert/client-go/pkg/client/pattern"
	"github.com/moira-alert/client-go/pkg/client/subscription"
	"github.com/moira-alert/client-go/pkg/client/tag"
	"github.com/moira-alert/client-go/pkg/client/team"
	"github.com/moira-alert/client-go/pkg/client/team_contact"
	"github.com/moira-alert/client-go/pkg/client/team_subscription"
	"github.com/moira-alert/client-go/pkg/client/trigger"
	"github.com/moira-alert/client-go/pkg/client/user"
)

// Default moira API HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new moira API HTTP client.
func NewHTTPClient(formats strfmt.Registry) *MoiraAPI {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new moira API HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *MoiraAPI {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new moira API client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *MoiraAPI {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(MoiraAPI)
	cli.Transport = transport
	cli.Config = config.New(transport, formats)
	cli.Contact = contact.New(transport, formats)
	cli.Event = event.New(transport, formats)
	cli.Health = health.New(transport, formats)
	cli.Notification = notification.New(transport, formats)
	cli.Pattern = pattern.New(transport, formats)
	cli.Subscription = subscription.New(transport, formats)
	cli.Tag = tag.New(transport, formats)
	cli.Team = team.New(transport, formats)
	cli.TeamContact = team_contact.New(transport, formats)
	cli.TeamSubscription = team_subscription.New(transport, formats)
	cli.Trigger = trigger.New(transport, formats)
	cli.User = user.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// MoiraAPI is a client for moira API
type MoiraAPI struct {
	Config config.ClientService

	Contact contact.ClientService

	Event event.ClientService

	Health health.ClientService

	Notification notification.ClientService

	Pattern pattern.ClientService

	Subscription subscription.ClientService

	Tag tag.ClientService

	Team team.ClientService

	TeamContact team_contact.ClientService

	TeamSubscription team_subscription.ClientService

	Trigger trigger.ClientService

	User user.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *MoiraAPI) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Config.SetTransport(transport)
	c.Contact.SetTransport(transport)
	c.Event.SetTransport(transport)
	c.Health.SetTransport(transport)
	c.Notification.SetTransport(transport)
	c.Pattern.SetTransport(transport)
	c.Subscription.SetTransport(transport)
	c.Tag.SetTransport(transport)
	c.Team.SetTransport(transport)
	c.TeamContact.SetTransport(transport)
	c.TeamSubscription.SetTransport(transport)
	c.Trigger.SetTransport(transport)
	c.User.SetTransport(transport)
}
