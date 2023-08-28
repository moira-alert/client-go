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
	"github.com/go-openapi/swag"
)

// NewGetNotificationsParams creates a new GetNotificationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNotificationsParams() *GetNotificationsParams {
	return &GetNotificationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNotificationsParamsWithTimeout creates a new GetNotificationsParams object
// with the ability to set a timeout on a request.
func NewGetNotificationsParamsWithTimeout(timeout time.Duration) *GetNotificationsParams {
	return &GetNotificationsParams{
		timeout: timeout,
	}
}

// NewGetNotificationsParamsWithContext creates a new GetNotificationsParams object
// with the ability to set a context for a request.
func NewGetNotificationsParamsWithContext(ctx context.Context) *GetNotificationsParams {
	return &GetNotificationsParams{
		Context: ctx,
	}
}

// NewGetNotificationsParamsWithHTTPClient creates a new GetNotificationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNotificationsParamsWithHTTPClient(client *http.Client) *GetNotificationsParams {
	return &GetNotificationsParams{
		HTTPClient: client,
	}
}

/*
GetNotificationsParams contains all the parameters to send to the API endpoint

	for the get notifications operation.

	Typically these are written to a http.Request.
*/
type GetNotificationsParams struct {

	/* End.

	   Default Value: -1

	   Default: -1
	*/
	End *int64

	/* Start.

	   Default Value: 0
	*/
	Start *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get notifications params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNotificationsParams) WithDefaults() *GetNotificationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get notifications params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNotificationsParams) SetDefaults() {
	var (
		endDefault = int64(-1)

		startDefault = int64(0)
	)

	val := GetNotificationsParams{
		End:   &endDefault,
		Start: &startDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get notifications params
func (o *GetNotificationsParams) WithTimeout(timeout time.Duration) *GetNotificationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get notifications params
func (o *GetNotificationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get notifications params
func (o *GetNotificationsParams) WithContext(ctx context.Context) *GetNotificationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get notifications params
func (o *GetNotificationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get notifications params
func (o *GetNotificationsParams) WithHTTPClient(client *http.Client) *GetNotificationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get notifications params
func (o *GetNotificationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnd adds the end to the get notifications params
func (o *GetNotificationsParams) WithEnd(end *int64) *GetNotificationsParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the get notifications params
func (o *GetNotificationsParams) SetEnd(end *int64) {
	o.End = end
}

// WithStart adds the start to the get notifications params
func (o *GetNotificationsParams) WithStart(start *int64) *GetNotificationsParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get notifications params
func (o *GetNotificationsParams) SetStart(start *int64) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *GetNotificationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.End != nil {

		// query param end
		var qrEnd int64

		if o.End != nil {
			qrEnd = *o.End
		}
		qEnd := swag.FormatInt64(qrEnd)
		if qEnd != "" {

			if err := r.SetQueryParam("end", qEnd); err != nil {
				return err
			}
		}
	}

	if o.Start != nil {

		// query param start
		var qrStart int64

		if o.Start != nil {
			qrStart = *o.Start
		}
		qStart := swag.FormatInt64(qrStart)
		if qStart != "" {

			if err := r.SetQueryParam("start", qStart); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}