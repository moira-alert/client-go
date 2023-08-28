// Code generated by go-swagger; DO NOT EDIT.

package event

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

// NewDeleteAllEventsParams creates a new DeleteAllEventsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAllEventsParams() *DeleteAllEventsParams {
	return &DeleteAllEventsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAllEventsParamsWithTimeout creates a new DeleteAllEventsParams object
// with the ability to set a timeout on a request.
func NewDeleteAllEventsParamsWithTimeout(timeout time.Duration) *DeleteAllEventsParams {
	return &DeleteAllEventsParams{
		timeout: timeout,
	}
}

// NewDeleteAllEventsParamsWithContext creates a new DeleteAllEventsParams object
// with the ability to set a context for a request.
func NewDeleteAllEventsParamsWithContext(ctx context.Context) *DeleteAllEventsParams {
	return &DeleteAllEventsParams{
		Context: ctx,
	}
}

// NewDeleteAllEventsParamsWithHTTPClient creates a new DeleteAllEventsParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAllEventsParamsWithHTTPClient(client *http.Client) *DeleteAllEventsParams {
	return &DeleteAllEventsParams{
		HTTPClient: client,
	}
}

/*
DeleteAllEventsParams contains all the parameters to send to the API endpoint

	for the delete all events operation.

	Typically these are written to a http.Request.
*/
type DeleteAllEventsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete all events params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAllEventsParams) WithDefaults() *DeleteAllEventsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete all events params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAllEventsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete all events params
func (o *DeleteAllEventsParams) WithTimeout(timeout time.Duration) *DeleteAllEventsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete all events params
func (o *DeleteAllEventsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete all events params
func (o *DeleteAllEventsParams) WithContext(ctx context.Context) *DeleteAllEventsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete all events params
func (o *DeleteAllEventsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete all events params
func (o *DeleteAllEventsParams) WithHTTPClient(client *http.Client) *DeleteAllEventsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete all events params
func (o *DeleteAllEventsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAllEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}