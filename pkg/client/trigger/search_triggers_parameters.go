// Code generated by go-swagger; DO NOT EDIT.

package trigger

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewSearchTriggersParams creates a new SearchTriggersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchTriggersParams() *SearchTriggersParams {
	return &SearchTriggersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchTriggersParamsWithTimeout creates a new SearchTriggersParams object
// with the ability to set a timeout on a request.
func NewSearchTriggersParamsWithTimeout(timeout time.Duration) *SearchTriggersParams {
	return &SearchTriggersParams{
		timeout: timeout,
	}
}

// NewSearchTriggersParamsWithContext creates a new SearchTriggersParams object
// with the ability to set a context for a request.
func NewSearchTriggersParamsWithContext(ctx context.Context) *SearchTriggersParams {
	return &SearchTriggersParams{
		Context: ctx,
	}
}

// NewSearchTriggersParamsWithHTTPClient creates a new SearchTriggersParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchTriggersParamsWithHTTPClient(client *http.Client) *SearchTriggersParams {
	return &SearchTriggersParams{
		HTTPClient: client,
	}
}

/*
SearchTriggersParams contains all the parameters to send to the API endpoint

	for the search triggers operation.

	Typically these are written to a http.Request.
*/
type SearchTriggersParams struct {

	/* CreatePager.

	   Create pager
	*/
	CreatePager *bool

	/* CreatedBy.

	   Created By

	   Default: "moira.team"
	*/
	CreatedBy *string

	/* OnlyProblems.

	   Only include problems
	*/
	OnlyProblems *bool

	/* P.

	   Page number
	*/
	P *int64

	/* PagerID.

	   Pager ID

	   Default: "bcba82f5-48cf-44c0-b7d6-e1d32c64a88c"
	*/
	PagerID *string

	/* Size.

	   Page size

	   Default: 10
	*/
	Size *int64

	/* Tags.

	   Tags
	*/
	Tags []string

	/* Text.

	   Search text

	   Default: "cpu"
	*/
	Text *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search triggers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchTriggersParams) WithDefaults() *SearchTriggersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search triggers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchTriggersParams) SetDefaults() {
	var (
		createPagerDefault = bool(false)

		createdByDefault = string("moira.team")

		onlyProblemsDefault = bool(false)

		pDefault = int64(0)

		pagerIDDefault = string("bcba82f5-48cf-44c0-b7d6-e1d32c64a88c")

		sizeDefault = int64(10)

		textDefault = string("cpu")
	)

	val := SearchTriggersParams{
		CreatePager:  &createPagerDefault,
		CreatedBy:    &createdByDefault,
		OnlyProblems: &onlyProblemsDefault,
		P:            &pDefault,
		PagerID:      &pagerIDDefault,
		Size:         &sizeDefault,
		Text:         &textDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the search triggers params
func (o *SearchTriggersParams) WithTimeout(timeout time.Duration) *SearchTriggersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search triggers params
func (o *SearchTriggersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search triggers params
func (o *SearchTriggersParams) WithContext(ctx context.Context) *SearchTriggersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search triggers params
func (o *SearchTriggersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search triggers params
func (o *SearchTriggersParams) WithHTTPClient(client *http.Client) *SearchTriggersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search triggers params
func (o *SearchTriggersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreatePager adds the createPager to the search triggers params
func (o *SearchTriggersParams) WithCreatePager(createPager *bool) *SearchTriggersParams {
	o.SetCreatePager(createPager)
	return o
}

// SetCreatePager adds the createPager to the search triggers params
func (o *SearchTriggersParams) SetCreatePager(createPager *bool) {
	o.CreatePager = createPager
}

// WithCreatedBy adds the createdBy to the search triggers params
func (o *SearchTriggersParams) WithCreatedBy(createdBy *string) *SearchTriggersParams {
	o.SetCreatedBy(createdBy)
	return o
}

// SetCreatedBy adds the createdBy to the search triggers params
func (o *SearchTriggersParams) SetCreatedBy(createdBy *string) {
	o.CreatedBy = createdBy
}

// WithOnlyProblems adds the onlyProblems to the search triggers params
func (o *SearchTriggersParams) WithOnlyProblems(onlyProblems *bool) *SearchTriggersParams {
	o.SetOnlyProblems(onlyProblems)
	return o
}

// SetOnlyProblems adds the onlyProblems to the search triggers params
func (o *SearchTriggersParams) SetOnlyProblems(onlyProblems *bool) {
	o.OnlyProblems = onlyProblems
}

// WithP adds the p to the search triggers params
func (o *SearchTriggersParams) WithP(p *int64) *SearchTriggersParams {
	o.SetP(p)
	return o
}

// SetP adds the p to the search triggers params
func (o *SearchTriggersParams) SetP(p *int64) {
	o.P = p
}

// WithPagerID adds the pagerID to the search triggers params
func (o *SearchTriggersParams) WithPagerID(pagerID *string) *SearchTriggersParams {
	o.SetPagerID(pagerID)
	return o
}

// SetPagerID adds the pagerId to the search triggers params
func (o *SearchTriggersParams) SetPagerID(pagerID *string) {
	o.PagerID = pagerID
}

// WithSize adds the size to the search triggers params
func (o *SearchTriggersParams) WithSize(size *int64) *SearchTriggersParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the search triggers params
func (o *SearchTriggersParams) SetSize(size *int64) {
	o.Size = size
}

// WithTags adds the tags to the search triggers params
func (o *SearchTriggersParams) WithTags(tags []string) *SearchTriggersParams {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the search triggers params
func (o *SearchTriggersParams) SetTags(tags []string) {
	o.Tags = tags
}

// WithText adds the text to the search triggers params
func (o *SearchTriggersParams) WithText(text *string) *SearchTriggersParams {
	o.SetText(text)
	return o
}

// SetText adds the text to the search triggers params
func (o *SearchTriggersParams) SetText(text *string) {
	o.Text = text
}

// WriteToRequest writes these params to a swagger request
func (o *SearchTriggersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreatePager != nil {

		// query param createPager
		var qrCreatePager bool

		if o.CreatePager != nil {
			qrCreatePager = *o.CreatePager
		}
		qCreatePager := swag.FormatBool(qrCreatePager)
		if qCreatePager != "" {

			if err := r.SetQueryParam("createPager", qCreatePager); err != nil {
				return err
			}
		}
	}

	if o.CreatedBy != nil {

		// query param createdBy
		var qrCreatedBy string

		if o.CreatedBy != nil {
			qrCreatedBy = *o.CreatedBy
		}
		qCreatedBy := qrCreatedBy
		if qCreatedBy != "" {

			if err := r.SetQueryParam("createdBy", qCreatedBy); err != nil {
				return err
			}
		}
	}

	if o.OnlyProblems != nil {

		// query param onlyProblems
		var qrOnlyProblems bool

		if o.OnlyProblems != nil {
			qrOnlyProblems = *o.OnlyProblems
		}
		qOnlyProblems := swag.FormatBool(qrOnlyProblems)
		if qOnlyProblems != "" {

			if err := r.SetQueryParam("onlyProblems", qOnlyProblems); err != nil {
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

	if o.PagerID != nil {

		// query param pagerID
		var qrPagerID string

		if o.PagerID != nil {
			qrPagerID = *o.PagerID
		}
		qPagerID := qrPagerID
		if qPagerID != "" {

			if err := r.SetQueryParam("pagerID", qPagerID); err != nil {
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

	if o.Tags != nil {
		// WAS NOT GENERATED because spec does not support our query param format !!!
		for i, tag := range o.Tags {
			if err := r.SetQueryParam("tags["+strconv.Itoa(i)+"]", tag); err != nil {
				return err
			}
		}
	}

	if o.Text != nil {

		// query param text
		var qrText string

		if o.Text != nil {
			qrText = *o.Text
		}
		qText := qrText
		if qText != "" {

			if err := r.SetQueryParam("text", qText); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSearchTriggers binds the parameter tags
func (o *SearchTriggersParams) bindParamTags(formats strfmt.Registry) []string {
	tagsIR := o.Tags

	var tagsIC []string
	for _, tagsIIR := range tagsIR { // explode []string

		tagsIIV := tagsIIR // string as string
		tagsIC = append(tagsIC, tagsIIV)
	}

	// items.CollectionFormat: "csv"
	tagsIS := swag.JoinByFormat(tagsIC, "csv")

	return tagsIS
}
