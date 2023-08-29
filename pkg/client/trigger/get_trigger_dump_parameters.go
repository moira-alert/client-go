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

// NewGetTriggerDumpParams creates a new GetTriggerDumpParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTriggerDumpParams() *GetTriggerDumpParams {
	return &GetTriggerDumpParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTriggerDumpParamsWithTimeout creates a new GetTriggerDumpParams object
// with the ability to set a timeout on a request.
func NewGetTriggerDumpParamsWithTimeout(timeout time.Duration) *GetTriggerDumpParams {
	return &GetTriggerDumpParams{
		timeout: timeout,
	}
}

// NewGetTriggerDumpParamsWithContext creates a new GetTriggerDumpParams object
// with the ability to set a context for a request.
func NewGetTriggerDumpParamsWithContext(ctx context.Context) *GetTriggerDumpParams {
	return &GetTriggerDumpParams{
		Context: ctx,
	}
}

// NewGetTriggerDumpParamsWithHTTPClient creates a new GetTriggerDumpParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTriggerDumpParamsWithHTTPClient(client *http.Client) *GetTriggerDumpParams {
	return &GetTriggerDumpParams{
		HTTPClient: client,
	}
}

/*
GetTriggerDumpParams contains all the parameters to send to the API endpoint

	for the get trigger dump operation.

	Typically these are written to a http.Request.
*/
type GetTriggerDumpParams struct {

	/* TriggerID.

	   Trigger ID

	   Default: "bcba82f5-48cf-44c0-b7d6-e1d32c64a88c"
	*/
	TriggerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get trigger dump params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTriggerDumpParams) WithDefaults() *GetTriggerDumpParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get trigger dump params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTriggerDumpParams) SetDefaults() {
	var (
		triggerIDDefault = string("bcba82f5-48cf-44c0-b7d6-e1d32c64a88c")
	)

	val := GetTriggerDumpParams{
		TriggerID: triggerIDDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get trigger dump params
func (o *GetTriggerDumpParams) WithTimeout(timeout time.Duration) *GetTriggerDumpParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get trigger dump params
func (o *GetTriggerDumpParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get trigger dump params
func (o *GetTriggerDumpParams) WithContext(ctx context.Context) *GetTriggerDumpParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get trigger dump params
func (o *GetTriggerDumpParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get trigger dump params
func (o *GetTriggerDumpParams) WithHTTPClient(client *http.Client) *GetTriggerDumpParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get trigger dump params
func (o *GetTriggerDumpParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTriggerID adds the triggerID to the get trigger dump params
func (o *GetTriggerDumpParams) WithTriggerID(triggerID string) *GetTriggerDumpParams {
	o.SetTriggerID(triggerID)
	return o
}

// SetTriggerID adds the triggerId to the get trigger dump params
func (o *GetTriggerDumpParams) SetTriggerID(triggerID string) {
	o.TriggerID = triggerID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTriggerDumpParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
