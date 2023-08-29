// Code generated by go-swagger; DO NOT EDIT.

package subscription

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

	"github.com/moira-alert/client-go/pkg/models"
)

// NewCreateSubscriptionParams creates a new CreateSubscriptionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateSubscriptionParams() *CreateSubscriptionParams {
	return &CreateSubscriptionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSubscriptionParamsWithTimeout creates a new CreateSubscriptionParams object
// with the ability to set a timeout on a request.
func NewCreateSubscriptionParamsWithTimeout(timeout time.Duration) *CreateSubscriptionParams {
	return &CreateSubscriptionParams{
		timeout: timeout,
	}
}

// NewCreateSubscriptionParamsWithContext creates a new CreateSubscriptionParams object
// with the ability to set a context for a request.
func NewCreateSubscriptionParamsWithContext(ctx context.Context) *CreateSubscriptionParams {
	return &CreateSubscriptionParams{
		Context: ctx,
	}
}

// NewCreateSubscriptionParamsWithHTTPClient creates a new CreateSubscriptionParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateSubscriptionParamsWithHTTPClient(client *http.Client) *CreateSubscriptionParams {
	return &CreateSubscriptionParams{
		HTTPClient: client,
	}
}

/*
CreateSubscriptionParams contains all the parameters to send to the API endpoint

	for the create subscription operation.

	Typically these are written to a http.Request.
*/
type CreateSubscriptionParams struct {

	/* Subscription.

	   Subscription data
	*/
	Subscription *models.DtoSubscription

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create subscription params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateSubscriptionParams) WithDefaults() *CreateSubscriptionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create subscription params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateSubscriptionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create subscription params
func (o *CreateSubscriptionParams) WithTimeout(timeout time.Duration) *CreateSubscriptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create subscription params
func (o *CreateSubscriptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create subscription params
func (o *CreateSubscriptionParams) WithContext(ctx context.Context) *CreateSubscriptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create subscription params
func (o *CreateSubscriptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create subscription params
func (o *CreateSubscriptionParams) WithHTTPClient(client *http.Client) *CreateSubscriptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create subscription params
func (o *CreateSubscriptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSubscription adds the subscription to the create subscription params
func (o *CreateSubscriptionParams) WithSubscription(subscription *models.DtoSubscription) *CreateSubscriptionParams {
	o.SetSubscription(subscription)
	return o
}

// SetSubscription adds the subscription to the create subscription params
func (o *CreateSubscriptionParams) SetSubscription(subscription *models.DtoSubscription) {
	o.Subscription = subscription
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSubscriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Subscription != nil {
		if err := r.SetBodyParam(o.Subscription); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
