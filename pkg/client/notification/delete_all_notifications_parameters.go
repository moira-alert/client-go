// Code generated by go-swagger; DO NOT EDIT.

package notification

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

// NewDeleteAllNotificationsParams creates a new DeleteAllNotificationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAllNotificationsParams() *DeleteAllNotificationsParams {
	return &DeleteAllNotificationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAllNotificationsParamsWithTimeout creates a new DeleteAllNotificationsParams object
// with the ability to set a timeout on a request.
func NewDeleteAllNotificationsParamsWithTimeout(timeout time.Duration) *DeleteAllNotificationsParams {
	return &DeleteAllNotificationsParams{
		timeout: timeout,
	}
}

// NewDeleteAllNotificationsParamsWithContext creates a new DeleteAllNotificationsParams object
// with the ability to set a context for a request.
func NewDeleteAllNotificationsParamsWithContext(ctx context.Context) *DeleteAllNotificationsParams {
	return &DeleteAllNotificationsParams{
		Context: ctx,
	}
}

// NewDeleteAllNotificationsParamsWithHTTPClient creates a new DeleteAllNotificationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAllNotificationsParamsWithHTTPClient(client *http.Client) *DeleteAllNotificationsParams {
	return &DeleteAllNotificationsParams{
		HTTPClient: client,
	}
}

/*
DeleteAllNotificationsParams contains all the parameters to send to the API endpoint

	for the delete all notifications operation.

	Typically these are written to a http.Request.
*/
type DeleteAllNotificationsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete all notifications params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAllNotificationsParams) WithDefaults() *DeleteAllNotificationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete all notifications params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAllNotificationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete all notifications params
func (o *DeleteAllNotificationsParams) WithTimeout(timeout time.Duration) *DeleteAllNotificationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete all notifications params
func (o *DeleteAllNotificationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete all notifications params
func (o *DeleteAllNotificationsParams) WithContext(ctx context.Context) *DeleteAllNotificationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete all notifications params
func (o *DeleteAllNotificationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete all notifications params
func (o *DeleteAllNotificationsParams) WithHTTPClient(client *http.Client) *DeleteAllNotificationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete all notifications params
func (o *DeleteAllNotificationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAllNotificationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}