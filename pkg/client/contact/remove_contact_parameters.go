// Code generated by go-swagger; DO NOT EDIT.

package contact

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

// NewRemoveContactParams creates a new RemoveContactParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRemoveContactParams() *RemoveContactParams {
	return &RemoveContactParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveContactParamsWithTimeout creates a new RemoveContactParams object
// with the ability to set a timeout on a request.
func NewRemoveContactParamsWithTimeout(timeout time.Duration) *RemoveContactParams {
	return &RemoveContactParams{
		timeout: timeout,
	}
}

// NewRemoveContactParamsWithContext creates a new RemoveContactParams object
// with the ability to set a context for a request.
func NewRemoveContactParamsWithContext(ctx context.Context) *RemoveContactParams {
	return &RemoveContactParams{
		Context: ctx,
	}
}

// NewRemoveContactParamsWithHTTPClient creates a new RemoveContactParams object
// with the ability to set a custom HTTPClient for a request.
func NewRemoveContactParamsWithHTTPClient(client *http.Client) *RemoveContactParams {
	return &RemoveContactParams{
		HTTPClient: client,
	}
}

/*
RemoveContactParams contains all the parameters to send to the API endpoint

	for the remove contact operation.

	Typically these are written to a http.Request.
*/
type RemoveContactParams struct {

	/* ContactID.

	   ID of the contact to remove

	   Default: "bcba82f5-48cf-44c0-b7d6-e1d32c64a88c"
	*/
	ContactID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the remove contact params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveContactParams) WithDefaults() *RemoveContactParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the remove contact params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveContactParams) SetDefaults() {
	var (
		contactIDDefault = string("bcba82f5-48cf-44c0-b7d6-e1d32c64a88c")
	)

	val := RemoveContactParams{
		ContactID: contactIDDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the remove contact params
func (o *RemoveContactParams) WithTimeout(timeout time.Duration) *RemoveContactParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove contact params
func (o *RemoveContactParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove contact params
func (o *RemoveContactParams) WithContext(ctx context.Context) *RemoveContactParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove contact params
func (o *RemoveContactParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove contact params
func (o *RemoveContactParams) WithHTTPClient(client *http.Client) *RemoveContactParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove contact params
func (o *RemoveContactParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContactID adds the contactID to the remove contact params
func (o *RemoveContactParams) WithContactID(contactID string) *RemoveContactParams {
	o.SetContactID(contactID)
	return o
}

// SetContactID adds the contactId to the remove contact params
func (o *RemoveContactParams) SetContactID(contactID string) {
	o.ContactID = contactID
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveContactParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param contactID
	if err := r.SetPathParam("contactID", o.ContactID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
