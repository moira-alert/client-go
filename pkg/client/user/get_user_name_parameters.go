// Code generated by go-swagger; DO NOT EDIT.

package user

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

// NewGetUserNameParams creates a new GetUserNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUserNameParams() *GetUserNameParams {
	return &GetUserNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserNameParamsWithTimeout creates a new GetUserNameParams object
// with the ability to set a timeout on a request.
func NewGetUserNameParamsWithTimeout(timeout time.Duration) *GetUserNameParams {
	return &GetUserNameParams{
		timeout: timeout,
	}
}

// NewGetUserNameParamsWithContext creates a new GetUserNameParams object
// with the ability to set a context for a request.
func NewGetUserNameParamsWithContext(ctx context.Context) *GetUserNameParams {
	return &GetUserNameParams{
		Context: ctx,
	}
}

// NewGetUserNameParamsWithHTTPClient creates a new GetUserNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUserNameParamsWithHTTPClient(client *http.Client) *GetUserNameParams {
	return &GetUserNameParams{
		HTTPClient: client,
	}
}

/*
GetUserNameParams contains all the parameters to send to the API endpoint

	for the get user name operation.

	Typically these are written to a http.Request.
*/
type GetUserNameParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get user name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserNameParams) WithDefaults() *GetUserNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get user name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get user name params
func (o *GetUserNameParams) WithTimeout(timeout time.Duration) *GetUserNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user name params
func (o *GetUserNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user name params
func (o *GetUserNameParams) WithContext(ctx context.Context) *GetUserNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user name params
func (o *GetUserNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user name params
func (o *GetUserNameParams) WithHTTPClient(client *http.Client) *GetUserNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user name params
func (o *GetUserNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
