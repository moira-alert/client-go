// Code generated by go-swagger; DO NOT EDIT.

package tag

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

// NewRemoveTagParams creates a new RemoveTagParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRemoveTagParams() *RemoveTagParams {
	return &RemoveTagParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveTagParamsWithTimeout creates a new RemoveTagParams object
// with the ability to set a timeout on a request.
func NewRemoveTagParamsWithTimeout(timeout time.Duration) *RemoveTagParams {
	return &RemoveTagParams{
		timeout: timeout,
	}
}

// NewRemoveTagParamsWithContext creates a new RemoveTagParams object
// with the ability to set a context for a request.
func NewRemoveTagParamsWithContext(ctx context.Context) *RemoveTagParams {
	return &RemoveTagParams{
		Context: ctx,
	}
}

// NewRemoveTagParamsWithHTTPClient creates a new RemoveTagParams object
// with the ability to set a custom HTTPClient for a request.
func NewRemoveTagParamsWithHTTPClient(client *http.Client) *RemoveTagParams {
	return &RemoveTagParams{
		HTTPClient: client,
	}
}

/*
RemoveTagParams contains all the parameters to send to the API endpoint

	for the remove tag operation.

	Typically these are written to a http.Request.
*/
type RemoveTagParams struct {

	/* Tag.

	   Name of the tag to remove

	   Default: "cpu"
	*/
	Tag string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the remove tag params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveTagParams) WithDefaults() *RemoveTagParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the remove tag params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveTagParams) SetDefaults() {
	var (
		tagDefault = string("cpu")
	)

	val := RemoveTagParams{
		Tag: tagDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the remove tag params
func (o *RemoveTagParams) WithTimeout(timeout time.Duration) *RemoveTagParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove tag params
func (o *RemoveTagParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove tag params
func (o *RemoveTagParams) WithContext(ctx context.Context) *RemoveTagParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove tag params
func (o *RemoveTagParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove tag params
func (o *RemoveTagParams) WithHTTPClient(client *http.Client) *RemoveTagParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove tag params
func (o *RemoveTagParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTag adds the tag to the remove tag params
func (o *RemoveTagParams) WithTag(tag string) *RemoveTagParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the remove tag params
func (o *RemoveTagParams) SetTag(tag string) {
	o.Tag = tag
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveTagParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param tag
	if err := r.SetPathParam("tag", o.Tag); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
