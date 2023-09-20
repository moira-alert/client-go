// Code generated by go-swagger; DO NOT EDIT.

package trigger

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetUnusedTriggersParams creates a new GetUnusedTriggersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUnusedTriggersParams() *GetUnusedTriggersParams {
	return &GetUnusedTriggersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUnusedTriggersParamsWithTimeout creates a new GetUnusedTriggersParams object
// with the ability to set a timeout on a request.
func NewGetUnusedTriggersParamsWithTimeout(timeout time.Duration) *GetUnusedTriggersParams {
	return &GetUnusedTriggersParams{
		timeout: timeout,
	}
}

// NewGetUnusedTriggersParamsWithContext creates a new GetUnusedTriggersParams object
// with the ability to set a context for a request.
func NewGetUnusedTriggersParamsWithContext(ctx context.Context) *GetUnusedTriggersParams {
	return &GetUnusedTriggersParams{
		Context: ctx,
	}
}

// NewGetUnusedTriggersParamsWithHTTPClient creates a new GetUnusedTriggersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUnusedTriggersParamsWithHTTPClient(client *http.Client) *GetUnusedTriggersParams {
	return &GetUnusedTriggersParams{
		HTTPClient: client,
	}
}

/*
GetUnusedTriggersParams contains all the parameters to send to the API endpoint

	for the get unused triggers operation.

	Typically these are written to a http.Request.
*/
type GetUnusedTriggersParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get unused triggers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUnusedTriggersParams) WithDefaults() *GetUnusedTriggersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get unused triggers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUnusedTriggersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get unused triggers params
func (o *GetUnusedTriggersParams) WithTimeout(timeout time.Duration) *GetUnusedTriggersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get unused triggers params
func (o *GetUnusedTriggersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get unused triggers params
func (o *GetUnusedTriggersParams) WithContext(ctx context.Context) *GetUnusedTriggersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get unused triggers params
func (o *GetUnusedTriggersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get unused triggers params
func (o *GetUnusedTriggersParams) WithHTTPClient(client *http.Client) *GetUnusedTriggersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get unused triggers params
func (o *GetUnusedTriggersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetUnusedTriggersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
