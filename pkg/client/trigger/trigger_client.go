// Code generated by go-swagger; DO NOT EDIT.

package trigger

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new trigger API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new trigger API client with basic auth credentials.
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

// New creates a new trigger API client with a bearer token for authentication.
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
Client for trigger API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// This client is generated with a few options you might find useful for your swagger spec.
//
// Feel free to add you own set of options.

// WithAccept allows the client to force the Accept header
// to negotiate a specific Producer from the server.
//
// You may use this option to set arbitrary extensions to your MIME media type.
func WithAccept(mime string) ClientOption {
	return func(r *runtime.ClientOperation) {
		r.ProducesMediaTypes = []string{mime}
	}
}

// WithAcceptApplicationJSON sets the Accept header to "application/json".
func WithAcceptApplicationJSON(r *runtime.ClientOperation) {
	r.ProducesMediaTypes = []string{"application/json"}
}

// WithAcceptImagePng sets the Accept header to "image/png".
func WithAcceptImagePng(r *runtime.ClientOperation) {
	r.ProducesMediaTypes = []string{"image/png"}
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateTrigger(params *CreateTriggerParams, opts ...ClientOption) (*CreateTriggerOK, error)

	DeletePager(params *DeletePagerParams, opts ...ClientOption) (*DeletePagerOK, error)

	DeleteTriggerMetric(params *DeleteTriggerMetricParams, opts ...ClientOption) (*DeleteTriggerMetricOK, error)

	DeleteTriggerNodataMetrics(params *DeleteTriggerNodataMetricsParams, opts ...ClientOption) (*DeleteTriggerNodataMetricsOK, error)

	DeleteTriggerThrottling(params *DeleteTriggerThrottlingParams, opts ...ClientOption) (*DeleteTriggerThrottlingOK, error)

	GetAllTriggers(params *GetAllTriggersParams, opts ...ClientOption) (*GetAllTriggersOK, error)

	GetTrigger(params *GetTriggerParams, opts ...ClientOption) (*GetTriggerOK, error)

	GetTriggerDump(params *GetTriggerDumpParams, opts ...ClientOption) (*GetTriggerDumpOK, error)

	GetTriggerMetrics(params *GetTriggerMetricsParams, opts ...ClientOption) (*GetTriggerMetricsOK, error)

	GetTriggerState(params *GetTriggerStateParams, opts ...ClientOption) (*GetTriggerStateOK, error)

	GetTriggerThrottling(params *GetTriggerThrottlingParams, opts ...ClientOption) (*GetTriggerThrottlingOK, error)

	GetUnusedTriggers(params *GetUnusedTriggersParams, opts ...ClientOption) (*GetUnusedTriggersOK, error)

	RemoveTrigger(params *RemoveTriggerParams, opts ...ClientOption) (*RemoveTriggerOK, error)

	RenderTriggerMetrics(params *RenderTriggerMetricsParams, opts ...ClientOption) (*RenderTriggerMetricsOK, error)

	SearchTriggers(params *SearchTriggersParams, opts ...ClientOption) (*SearchTriggersOK, error)

	SetTriggerMaintenance(params *SetTriggerMaintenanceParams, opts ...ClientOption) (*SetTriggerMaintenanceOK, error)

	TriggerCheck(params *TriggerCheckParams, opts ...ClientOption) (*TriggerCheckOK, error)

	UpdateTrigger(params *UpdateTriggerParams, opts ...ClientOption) (*UpdateTriggerOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateTrigger creates a new trigger
*/
func (a *Client) CreateTrigger(params *CreateTriggerParams, opts ...ClientOption) (*CreateTriggerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateTriggerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "create-trigger",
		Method:             "PUT",
		PathPattern:        "/trigger",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateTriggerReader{formats: a.formats},
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
	success, ok := result.(*CreateTriggerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for create-trigger: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeletePager deletes triggers pager
*/
func (a *Client) DeletePager(params *DeletePagerParams, opts ...ClientOption) (*DeletePagerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePagerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delete-pager",
		Method:             "DELETE",
		PathPattern:        "/trigger/search/pager",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeletePagerReader{formats: a.formats},
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
	success, ok := result.(*DeletePagerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete-pager: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteTriggerMetric deletes metric from last check and all trigger pattern metrics
*/
func (a *Client) DeleteTriggerMetric(params *DeleteTriggerMetricParams, opts ...ClientOption) (*DeleteTriggerMetricOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTriggerMetricParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delete-trigger-metric",
		Method:             "DELETE",
		PathPattern:        "/trigger/{triggerID}/metrics",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteTriggerMetricReader{formats: a.formats},
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
	success, ok := result.(*DeleteTriggerMetricOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete-trigger-metric: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteTriggerNodataMetrics deletes all metrics from last data which are in n o d a t a state it also deletes all trigger patterns of those metrics
*/
func (a *Client) DeleteTriggerNodataMetrics(params *DeleteTriggerNodataMetricsParams, opts ...ClientOption) (*DeleteTriggerNodataMetricsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTriggerNodataMetricsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delete-trigger-nodata-metrics",
		Method:             "DELETE",
		PathPattern:        "/trigger/{triggerID}/metrics/nodata",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteTriggerNodataMetricsReader{formats: a.formats},
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
	success, ok := result.(*DeleteTriggerNodataMetricsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete-trigger-nodata-metrics: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteTriggerThrottling deletes throttling for a trigger
*/
func (a *Client) DeleteTriggerThrottling(params *DeleteTriggerThrottlingParams, opts ...ClientOption) (*DeleteTriggerThrottlingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTriggerThrottlingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delete-trigger-throttling",
		Method:             "DELETE",
		PathPattern:        "/trigger/{triggerID}/throttling",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteTriggerThrottlingReader{formats: a.formats},
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
	success, ok := result.(*DeleteTriggerThrottlingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete-trigger-throttling: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAllTriggers gets all triggers
*/
func (a *Client) GetAllTriggers(params *GetAllTriggersParams, opts ...ClientOption) (*GetAllTriggersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllTriggersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-all-triggers",
		Method:             "GET",
		PathPattern:        "/trigger",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAllTriggersReader{formats: a.formats},
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
	success, ok := result.(*GetAllTriggersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-all-triggers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetTrigger gets an existing trigger
*/
func (a *Client) GetTrigger(params *GetTriggerParams, opts ...ClientOption) (*GetTriggerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTriggerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-trigger",
		Method:             "GET",
		PathPattern:        "/trigger/{triggerID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTriggerReader{formats: a.formats},
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
	success, ok := result.(*GetTriggerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-trigger: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetTriggerDump gets trigger dump
*/
func (a *Client) GetTriggerDump(params *GetTriggerDumpParams, opts ...ClientOption) (*GetTriggerDumpOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTriggerDumpParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-trigger-dump",
		Method:             "GET",
		PathPattern:        "/trigger/{triggerID}/dump",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTriggerDumpReader{formats: a.formats},
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
	success, ok := result.(*GetTriggerDumpOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-trigger-dump: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetTriggerMetrics gets metrics associated with certain trigger
*/
func (a *Client) GetTriggerMetrics(params *GetTriggerMetricsParams, opts ...ClientOption) (*GetTriggerMetricsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTriggerMetricsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-trigger-metrics",
		Method:             "GET",
		PathPattern:        "/trigger/{triggerID}/metrics",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTriggerMetricsReader{formats: a.formats},
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
	success, ok := result.(*GetTriggerMetricsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-trigger-metrics: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetTriggerState gets the trigger state as at last check
*/
func (a *Client) GetTriggerState(params *GetTriggerStateParams, opts ...ClientOption) (*GetTriggerStateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTriggerStateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-trigger-state",
		Method:             "GET",
		PathPattern:        "/trigger/{triggerID}/state",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTriggerStateReader{formats: a.formats},
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
	success, ok := result.(*GetTriggerStateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-trigger-state: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetTriggerThrottling gets a trigger with its throttling i e its next allowed message time
*/
func (a *Client) GetTriggerThrottling(params *GetTriggerThrottlingParams, opts ...ClientOption) (*GetTriggerThrottlingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTriggerThrottlingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-trigger-throttling",
		Method:             "GET",
		PathPattern:        "/trigger/{triggerID}/throttling",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTriggerThrottlingReader{formats: a.formats},
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
	success, ok := result.(*GetTriggerThrottlingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-trigger-throttling: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetUnusedTriggers gets unused triggers
*/
func (a *Client) GetUnusedTriggers(params *GetUnusedTriggersParams, opts ...ClientOption) (*GetUnusedTriggersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUnusedTriggersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-unused-triggers",
		Method:             "GET",
		PathPattern:        "/trigger/unused",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetUnusedTriggersReader{formats: a.formats},
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
	success, ok := result.(*GetUnusedTriggersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-unused-triggers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RemoveTrigger removes trigger
*/
func (a *Client) RemoveTrigger(params *RemoveTriggerParams, opts ...ClientOption) (*RemoveTriggerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveTriggerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "remove-trigger",
		Method:             "DELETE",
		PathPattern:        "/trigger/{triggerID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RemoveTriggerReader{formats: a.formats},
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
	success, ok := result.(*RemoveTriggerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for remove-trigger: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RenderTriggerMetrics renders trigger metrics plot
*/
func (a *Client) RenderTriggerMetrics(params *RenderTriggerMetricsParams, opts ...ClientOption) (*RenderTriggerMetricsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRenderTriggerMetricsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "render-trigger-metrics",
		Method:             "GET",
		PathPattern:        "/trigger/{triggerID}/render",
		ProducesMediaTypes: []string{"image/png"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RenderTriggerMetricsReader{formats: a.formats},
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
	success, ok := result.(*RenderTriggerMetricsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for render-trigger-metrics: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	SearchTriggers searches triggers replaces the deprecated page path

	You can also add filtering by tags, for this purpose add query parameters tags[0]=test, tags[1]=test1 and so on

For example, `/api/trigger/search?tags[0]=test&tags[1]=test1`
*/
func (a *Client) SearchTriggers(params *SearchTriggersParams, opts ...ClientOption) (*SearchTriggersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchTriggersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "search-triggers",
		Method:             "GET",
		PathPattern:        "/trigger/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SearchTriggersReader{formats: a.formats},
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
	success, ok := result.(*SearchTriggersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for search-triggers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SetTriggerMaintenance sets metrics and the trigger itself to maintenance mode
*/
func (a *Client) SetTriggerMaintenance(params *SetTriggerMaintenanceParams, opts ...ClientOption) (*SetTriggerMaintenanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetTriggerMaintenanceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "set-trigger-maintenance",
		Method:             "PUT",
		PathPattern:        "/trigger/{triggerID}/setMaintenance",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SetTriggerMaintenanceReader{formats: a.formats},
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
	success, ok := result.(*SetTriggerMaintenanceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for set-trigger-maintenance: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
TriggerCheck validates trigger target
*/
func (a *Client) TriggerCheck(params *TriggerCheckParams, opts ...ClientOption) (*TriggerCheckOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTriggerCheckParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "trigger-check",
		Method:             "PUT",
		PathPattern:        "/trigger/check",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &TriggerCheckReader{formats: a.formats},
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
	success, ok := result.(*TriggerCheckOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for trigger-check: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateTrigger updates existing trigger
*/
func (a *Client) UpdateTrigger(params *UpdateTriggerParams, opts ...ClientOption) (*UpdateTriggerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateTriggerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "update-trigger",
		Method:             "PUT",
		PathPattern:        "/trigger/{triggerID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateTriggerReader{formats: a.formats},
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
	success, ok := result.(*UpdateTriggerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for update-trigger: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
