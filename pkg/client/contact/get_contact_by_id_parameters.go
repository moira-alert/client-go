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

// NewGetContactByIDParams creates a new GetContactByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetContactByIDParams() *GetContactByIDParams {
	return &GetContactByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetContactByIDParamsWithTimeout creates a new GetContactByIDParams object
// with the ability to set a timeout on a request.
func NewGetContactByIDParamsWithTimeout(timeout time.Duration) *GetContactByIDParams {
	return &GetContactByIDParams{
		timeout: timeout,
	}
}

// NewGetContactByIDParamsWithContext creates a new GetContactByIDParams object
// with the ability to set a context for a request.
func NewGetContactByIDParamsWithContext(ctx context.Context) *GetContactByIDParams {
	return &GetContactByIDParams{
		Context: ctx,
	}
}

// NewGetContactByIDParamsWithHTTPClient creates a new GetContactByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetContactByIDParamsWithHTTPClient(client *http.Client) *GetContactByIDParams {
	return &GetContactByIDParams{
		HTTPClient: client,
	}
}

/*
GetContactByIDParams contains all the parameters to send to the API endpoint

	for the get contact by id operation.

	Typically these are written to a http.Request.
*/
type GetContactByIDParams struct {

	/* ContactID.

	   Contact ID

	   Default: "bcba82f5-48cf-44c0-b7d6-e1d32c64a88c"
	*/
	ContactID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get contact by id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetContactByIDParams) WithDefaults() *GetContactByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get contact by id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetContactByIDParams) SetDefaults() {
	var (
		contactIDDefault = string("bcba82f5-48cf-44c0-b7d6-e1d32c64a88c")
	)

	val := GetContactByIDParams{
		ContactID: contactIDDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get contact by id params
func (o *GetContactByIDParams) WithTimeout(timeout time.Duration) *GetContactByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get contact by id params
func (o *GetContactByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get contact by id params
func (o *GetContactByIDParams) WithContext(ctx context.Context) *GetContactByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get contact by id params
func (o *GetContactByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get contact by id params
func (o *GetContactByIDParams) WithHTTPClient(client *http.Client) *GetContactByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get contact by id params
func (o *GetContactByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContactID adds the contactID to the get contact by id params
func (o *GetContactByIDParams) WithContactID(contactID string) *GetContactByIDParams {
	o.SetContactID(contactID)
	return o
}

// SetContactID adds the contactId to the get contact by id params
func (o *GetContactByIDParams) SetContactID(contactID string) {
	o.ContactID = contactID
}

// WriteToRequest writes these params to a swagger request
func (o *GetContactByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
