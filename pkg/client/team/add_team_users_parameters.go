// Code generated by go-swagger; DO NOT EDIT.

package team

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

// NewAddTeamUsersParams creates a new AddTeamUsersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddTeamUsersParams() *AddTeamUsersParams {
	return &AddTeamUsersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddTeamUsersParamsWithTimeout creates a new AddTeamUsersParams object
// with the ability to set a timeout on a request.
func NewAddTeamUsersParamsWithTimeout(timeout time.Duration) *AddTeamUsersParams {
	return &AddTeamUsersParams{
		timeout: timeout,
	}
}

// NewAddTeamUsersParamsWithContext creates a new AddTeamUsersParams object
// with the ability to set a context for a request.
func NewAddTeamUsersParamsWithContext(ctx context.Context) *AddTeamUsersParams {
	return &AddTeamUsersParams{
		Context: ctx,
	}
}

// NewAddTeamUsersParamsWithHTTPClient creates a new AddTeamUsersParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddTeamUsersParamsWithHTTPClient(client *http.Client) *AddTeamUsersParams {
	return &AddTeamUsersParams{
		HTTPClient: client,
	}
}

/*
AddTeamUsersParams contains all the parameters to send to the API endpoint

	for the add team users operation.

	Typically these are written to a http.Request.
*/
type AddTeamUsersParams struct {

	/* TeamID.

	   ID of the team

	   Default: "bcba82f5-48cf-44c0-b7d6-e1d32c64a88c"
	*/
	TeamID string

	/* Usernames.

	   Usernames to add to the team
	*/
	Usernames *models.DtoTeamMembers

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add team users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddTeamUsersParams) WithDefaults() *AddTeamUsersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add team users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddTeamUsersParams) SetDefaults() {
	var (
		teamIDDefault = string("bcba82f5-48cf-44c0-b7d6-e1d32c64a88c")
	)

	val := AddTeamUsersParams{
		TeamID: teamIDDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the add team users params
func (o *AddTeamUsersParams) WithTimeout(timeout time.Duration) *AddTeamUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add team users params
func (o *AddTeamUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add team users params
func (o *AddTeamUsersParams) WithContext(ctx context.Context) *AddTeamUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add team users params
func (o *AddTeamUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add team users params
func (o *AddTeamUsersParams) WithHTTPClient(client *http.Client) *AddTeamUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add team users params
func (o *AddTeamUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTeamID adds the teamID to the add team users params
func (o *AddTeamUsersParams) WithTeamID(teamID string) *AddTeamUsersParams {
	o.SetTeamID(teamID)
	return o
}

// SetTeamID adds the teamId to the add team users params
func (o *AddTeamUsersParams) SetTeamID(teamID string) {
	o.TeamID = teamID
}

// WithUsernames adds the usernames to the add team users params
func (o *AddTeamUsersParams) WithUsernames(usernames *models.DtoTeamMembers) *AddTeamUsersParams {
	o.SetUsernames(usernames)
	return o
}

// SetUsernames adds the usernames to the add team users params
func (o *AddTeamUsersParams) SetUsernames(usernames *models.DtoTeamMembers) {
	o.Usernames = usernames
}

// WriteToRequest writes these params to a swagger request
func (o *AddTeamUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param teamID
	if err := r.SetPathParam("teamID", o.TeamID); err != nil {
		return err
	}
	if o.Usernames != nil {
		if err := r.SetBodyParam(o.Usernames); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}