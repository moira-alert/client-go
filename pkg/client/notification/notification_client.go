// Code generated by go-swagger; DO NOT EDIT.

package notification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new notification API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for notification API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteAllNotifications(params *DeleteAllNotificationsParams, opts ...ClientOption) (*DeleteAllNotificationsOK, error)

	GetNotifications(params *GetNotificationsParams, opts ...ClientOption) (*GetNotificationsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteAllNotifications deletes all notifications
*/
func (a *Client) DeleteAllNotifications(params *DeleteAllNotificationsParams, opts ...ClientOption) (*DeleteAllNotificationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAllNotificationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delete-all-notifications",
		Method:             "DELETE",
		PathPattern:        "/notification",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteAllNotificationsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteAllNotificationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete-all-notifications: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetNotifications gets a paginated list of notifications all notifications are fetched if end 1 and start 0
*/
func (a *Client) GetNotifications(params *GetNotificationsParams, opts ...ClientOption) (*GetNotificationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNotificationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-notifications",
		Method:             "GET",
		PathPattern:        "/notification",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetNotificationsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNotificationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-notifications: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
