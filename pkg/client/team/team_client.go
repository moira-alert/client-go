// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new team API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new team API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new team API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for team API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddTeamUsers(params *AddTeamUsersParams, opts ...ClientOption) (*AddTeamUsersOK, error)

	CreateTeam(params *CreateTeamParams, opts ...ClientOption) (*CreateTeamOK, error)

	DeleteTeam(params *DeleteTeamParams, opts ...ClientOption) (*DeleteTeamOK, error)

	DeleteTeamUser(params *DeleteTeamUserParams, opts ...ClientOption) (*DeleteTeamUserOK, error)

	GetAllTeams(params *GetAllTeamsParams, opts ...ClientOption) (*GetAllTeamsOK, error)

	GetAllTeamsForUser(params *GetAllTeamsForUserParams, opts ...ClientOption) (*GetAllTeamsForUserOK, error)

	GetTeam(params *GetTeamParams, opts ...ClientOption) (*GetTeamOK, error)

	GetTeamSettings(params *GetTeamSettingsParams, opts ...ClientOption) (*GetTeamSettingsOK, error)

	GetTeamUsers(params *GetTeamUsersParams, opts ...ClientOption) (*GetTeamUsersOK, error)

	SetTeamUsers(params *SetTeamUsersParams, opts ...ClientOption) (*SetTeamUsersOK, error)

	UpdateTeam(params *UpdateTeamParams, opts ...ClientOption) (*UpdateTeamOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AddTeamUsers adds users to a team
*/
func (a *Client) AddTeamUsers(params *AddTeamUsersParams, opts ...ClientOption) (*AddTeamUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddTeamUsersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "add-team-users",
		Method:             "POST",
		PathPattern:        "/teams/{teamID}/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddTeamUsersReader{formats: a.formats},
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
	success, ok := result.(*AddTeamUsersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for add-team-users: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateTeam creates a new team
*/
func (a *Client) CreateTeam(params *CreateTeamParams, opts ...ClientOption) (*CreateTeamOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateTeamParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "create-team",
		Method:             "POST",
		PathPattern:        "/teams",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateTeamReader{formats: a.formats},
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
	success, ok := result.(*CreateTeamOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for create-team: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteTeam deletes a team
*/
func (a *Client) DeleteTeam(params *DeleteTeamParams, opts ...ClientOption) (*DeleteTeamOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTeamParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delete-team",
		Method:             "DELETE",
		PathPattern:        "/teams/{teamID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteTeamReader{formats: a.formats},
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
	success, ok := result.(*DeleteTeamOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete-team: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteTeamUser deletes a user from a team
*/
func (a *Client) DeleteTeamUser(params *DeleteTeamUserParams, opts ...ClientOption) (*DeleteTeamUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTeamUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delete-team-user",
		Method:             "DELETE",
		PathPattern:        "/teams/{teamID}/users/{teamUserID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteTeamUserReader{formats: a.formats},
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
	success, ok := result.(*DeleteTeamUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete-team-user: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAllTeams gets all moira teams
*/
func (a *Client) GetAllTeams(params *GetAllTeamsParams, opts ...ClientOption) (*GetAllTeamsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllTeamsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-all-teams",
		Method:             "GET",
		PathPattern:        "/teams/all",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAllTeamsReader{formats: a.formats},
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
	success, ok := result.(*GetAllTeamsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-all-teams: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAllTeamsForUser gets all teams for user
*/
func (a *Client) GetAllTeamsForUser(params *GetAllTeamsForUserParams, opts ...ClientOption) (*GetAllTeamsForUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllTeamsForUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-all-teams-for-user",
		Method:             "GET",
		PathPattern:        "/teams",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAllTeamsForUserReader{formats: a.formats},
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
	success, ok := result.(*GetAllTeamsForUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-all-teams-for-user: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetTeam gets a team by ID
*/
func (a *Client) GetTeam(params *GetTeamParams, opts ...ClientOption) (*GetTeamOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTeamParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-team",
		Method:             "GET",
		PathPattern:        "/teams/{teamID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTeamReader{formats: a.formats},
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
	success, ok := result.(*GetTeamOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-team: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetTeamSettings gets team settings
*/
func (a *Client) GetTeamSettings(params *GetTeamSettingsParams, opts ...ClientOption) (*GetTeamSettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTeamSettingsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-team-settings",
		Method:             "GET",
		PathPattern:        "/teams/{teamID}/settings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTeamSettingsReader{formats: a.formats},
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
	success, ok := result.(*GetTeamSettingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-team-settings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetTeamUsers gets users of a team
*/
func (a *Client) GetTeamUsers(params *GetTeamUsersParams, opts ...ClientOption) (*GetTeamUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTeamUsersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-team-users",
		Method:             "GET",
		PathPattern:        "/teams/{teamID}/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTeamUsersReader{formats: a.formats},
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
	success, ok := result.(*GetTeamUsersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-team-users: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SetTeamUsers sets users of a team
*/
func (a *Client) SetTeamUsers(params *SetTeamUsersParams, opts ...ClientOption) (*SetTeamUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetTeamUsersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "set-team-users",
		Method:             "PUT",
		PathPattern:        "/teams/{teamID}/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SetTeamUsersReader{formats: a.formats},
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
	success, ok := result.(*SetTeamUsersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for set-team-users: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateTeam updates existing team
*/
func (a *Client) UpdateTeam(params *UpdateTeamParams, opts ...ClientOption) (*UpdateTeamOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateTeamParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "update-team",
		Method:             "PATCH",
		PathPattern:        "/teams/{teamID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateTeamReader{formats: a.formats},
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
	success, ok := result.(*UpdateTeamOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for update-team: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
