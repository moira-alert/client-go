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
	"github.com/go-openapi/swag"
)

// NewGetContactsNoisinessParams creates a new GetContactsNoisinessParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetContactsNoisinessParams() *GetContactsNoisinessParams {
	return &GetContactsNoisinessParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetContactsNoisinessParamsWithTimeout creates a new GetContactsNoisinessParams object
// with the ability to set a timeout on a request.
func NewGetContactsNoisinessParamsWithTimeout(timeout time.Duration) *GetContactsNoisinessParams {
	return &GetContactsNoisinessParams{
		timeout: timeout,
	}
}

// NewGetContactsNoisinessParamsWithContext creates a new GetContactsNoisinessParams object
// with the ability to set a context for a request.
func NewGetContactsNoisinessParamsWithContext(ctx context.Context) *GetContactsNoisinessParams {
	return &GetContactsNoisinessParams{
		Context: ctx,
	}
}

// NewGetContactsNoisinessParamsWithHTTPClient creates a new GetContactsNoisinessParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetContactsNoisinessParamsWithHTTPClient(client *http.Client) *GetContactsNoisinessParams {
	return &GetContactsNoisinessParams{
		HTTPClient: client,
	}
}

/*
GetContactsNoisinessParams contains all the parameters to send to the API endpoint

	for the get contacts noisiness operation.

	Typically these are written to a http.Request.
*/
type GetContactsNoisinessParams struct {

	/* From.

	   Start time of the time range

	   Default: "-3hours"
	*/
	From *string

	/* P.

	   Defines the number of the displayed page. E.g, p=2 would display the 2nd page
	*/
	P *int64

	/* Size.

	   Number of items to be displayed on one page. if size = -1 then all events returned

	   Default: 100
	*/
	Size *int64

	/* Sort.

	   String to set sort order (by events_count). On empty - no order, asc - ascending, desc - descending

	   Default: "desc"
	*/
	Sort *string

	/* To.

	   End time of the time range

	   Default: "now"
	*/
	To *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get contacts noisiness params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetContactsNoisinessParams) WithDefaults() *GetContactsNoisinessParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get contacts noisiness params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetContactsNoisinessParams) SetDefaults() {
	var (
		fromDefault = string("-3hours")

		pDefault = int64(0)

		sizeDefault = int64(100)

		sortDefault = string("desc")

		toDefault = string("now")
	)

	val := GetContactsNoisinessParams{
		From: &fromDefault,
		P:    &pDefault,
		Size: &sizeDefault,
		Sort: &sortDefault,
		To:   &toDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get contacts noisiness params
func (o *GetContactsNoisinessParams) WithTimeout(timeout time.Duration) *GetContactsNoisinessParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get contacts noisiness params
func (o *GetContactsNoisinessParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get contacts noisiness params
func (o *GetContactsNoisinessParams) WithContext(ctx context.Context) *GetContactsNoisinessParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get contacts noisiness params
func (o *GetContactsNoisinessParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get contacts noisiness params
func (o *GetContactsNoisinessParams) WithHTTPClient(client *http.Client) *GetContactsNoisinessParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get contacts noisiness params
func (o *GetContactsNoisinessParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFrom adds the from to the get contacts noisiness params
func (o *GetContactsNoisinessParams) WithFrom(from *string) *GetContactsNoisinessParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the get contacts noisiness params
func (o *GetContactsNoisinessParams) SetFrom(from *string) {
	o.From = from
}

// WithP adds the p to the get contacts noisiness params
func (o *GetContactsNoisinessParams) WithP(p *int64) *GetContactsNoisinessParams {
	o.SetP(p)
	return o
}

// SetP adds the p to the get contacts noisiness params
func (o *GetContactsNoisinessParams) SetP(p *int64) {
	o.P = p
}

// WithSize adds the size to the get contacts noisiness params
func (o *GetContactsNoisinessParams) WithSize(size *int64) *GetContactsNoisinessParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get contacts noisiness params
func (o *GetContactsNoisinessParams) SetSize(size *int64) {
	o.Size = size
}

// WithSort adds the sort to the get contacts noisiness params
func (o *GetContactsNoisinessParams) WithSort(sort *string) *GetContactsNoisinessParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get contacts noisiness params
func (o *GetContactsNoisinessParams) SetSort(sort *string) {
	o.Sort = sort
}

// WithTo adds the to to the get contacts noisiness params
func (o *GetContactsNoisinessParams) WithTo(to *string) *GetContactsNoisinessParams {
	o.SetTo(to)
	return o
}

// SetTo adds the to to the get contacts noisiness params
func (o *GetContactsNoisinessParams) SetTo(to *string) {
	o.To = to
}

// WriteToRequest writes these params to a swagger request
func (o *GetContactsNoisinessParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.From != nil {

		// query param from
		var qrFrom string

		if o.From != nil {
			qrFrom = *o.From
		}
		qFrom := qrFrom
		if qFrom != "" {

			if err := r.SetQueryParam("from", qFrom); err != nil {
				return err
			}
		}
	}

	if o.P != nil {

		// query param p
		var qrP int64

		if o.P != nil {
			qrP = *o.P
		}
		qP := swag.FormatInt64(qrP)
		if qP != "" {

			if err := r.SetQueryParam("p", qP); err != nil {
				return err
			}
		}
	}

	if o.Size != nil {

		// query param size
		var qrSize int64

		if o.Size != nil {
			qrSize = *o.Size
		}
		qSize := swag.FormatInt64(qrSize)
		if qSize != "" {

			if err := r.SetQueryParam("size", qSize); err != nil {
				return err
			}
		}
	}

	if o.Sort != nil {

		// query param sort
		var qrSort string

		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {

			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}
	}

	if o.To != nil {

		// query param to
		var qrTo string

		if o.To != nil {
			qrTo = *o.To
		}
		qTo := qrTo
		if qTo != "" {

			if err := r.SetQueryParam("to", qTo); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
