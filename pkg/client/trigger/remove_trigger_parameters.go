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

// NewRemoveTriggerParams creates a new RemoveTriggerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRemoveTriggerParams() *RemoveTriggerParams {
	return &RemoveTriggerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveTriggerParamsWithTimeout creates a new RemoveTriggerParams object
// with the ability to set a timeout on a request.
func NewRemoveTriggerParamsWithTimeout(timeout time.Duration) *RemoveTriggerParams {
	return &RemoveTriggerParams{
		timeout: timeout,
	}
}

// NewRemoveTriggerParamsWithContext creates a new RemoveTriggerParams object
// with the ability to set a context for a request.
func NewRemoveTriggerParamsWithContext(ctx context.Context) *RemoveTriggerParams {
	return &RemoveTriggerParams{
		Context: ctx,
	}
}

// NewRemoveTriggerParamsWithHTTPClient creates a new RemoveTriggerParams object
// with the ability to set a custom HTTPClient for a request.
func NewRemoveTriggerParamsWithHTTPClient(client *http.Client) *RemoveTriggerParams {
	return &RemoveTriggerParams{
		HTTPClient: client,
	}
}

/*
RemoveTriggerParams contains all the parameters to send to the API endpoint

	for the remove trigger operation.

	Typically these are written to a http.Request.
*/
type RemoveTriggerParams struct {

	/* TriggerID.

	   Trigger ID

	   Default: "bcba82f5-48cf-44c0-b7d6-e1d32c64a88c"
	*/
	TriggerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the remove trigger params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveTriggerParams) WithDefaults() *RemoveTriggerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the remove trigger params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveTriggerParams) SetDefaults() {
	var (
		triggerIDDefault = string("bcba82f5-48cf-44c0-b7d6-e1d32c64a88c")
	)

	val := RemoveTriggerParams{
		TriggerID: triggerIDDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the remove trigger params
func (o *RemoveTriggerParams) WithTimeout(timeout time.Duration) *RemoveTriggerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove trigger params
func (o *RemoveTriggerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove trigger params
func (o *RemoveTriggerParams) WithContext(ctx context.Context) *RemoveTriggerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove trigger params
func (o *RemoveTriggerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove trigger params
func (o *RemoveTriggerParams) WithHTTPClient(client *http.Client) *RemoveTriggerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove trigger params
func (o *RemoveTriggerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTriggerID adds the triggerID to the remove trigger params
func (o *RemoveTriggerParams) WithTriggerID(triggerID string) *RemoveTriggerParams {
	o.SetTriggerID(triggerID)
	return o
}

// SetTriggerID adds the triggerId to the remove trigger params
func (o *RemoveTriggerParams) SetTriggerID(triggerID string) {
	o.TriggerID = triggerID
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveTriggerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param triggerID
	if err := r.SetPathParam("triggerID", o.TriggerID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
