// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DtoSubscription dto subscription
//
// swagger:model dto.Subscription
type DtoSubscription struct {

	// any tags
	// Example: false
	AnyTags bool `json:"any_tags,omitempty"`

	// contacts
	// Example: ["acd2db98-1659-4a2f-b227-52d71f6e3ba1"]
	Contacts []string `json:"contacts"`

	// enabled
	// Example: true
	Enabled bool `json:"enabled,omitempty"`

	// id
	// Example: 292516ed-4924-4154-a62c-ebe312431fce
	ID string `json:"id,omitempty"`

	// ignore recoverings
	// Example: false
	IgnoreRecoverings bool `json:"ignore_recoverings,omitempty"`

	// ignore warnings
	// Example: false
	IgnoreWarnings bool `json:"ignore_warnings,omitempty"`

	// plotting
	Plotting *MoiraPlottingData `json:"plotting,omitempty"`

	// sched
	Sched *MoiraScheduleData `json:"sched,omitempty"`

	// tags
	// Example: ["server","cpu"]
	Tags []string `json:"tags"`

	// team id
	// Example: 324516ed-4924-4154-a62c-eb124234fce
	TeamID string `json:"team_id,omitempty"`

	// throttling
	// Example: false
	Throttling bool `json:"throttling,omitempty"`

	// user
	User string `json:"user,omitempty"`
}

// Validate validates this dto subscription
func (m *DtoSubscription) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePlotting(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSched(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DtoSubscription) validatePlotting(formats strfmt.Registry) error {
	if swag.IsZero(m.Plotting) { // not required
		return nil
	}

	if m.Plotting != nil {
		if err := m.Plotting.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("plotting")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("plotting")
			}
			return err
		}
	}

	return nil
}

func (m *DtoSubscription) validateSched(formats strfmt.Registry) error {
	if swag.IsZero(m.Sched) { // not required
		return nil
	}

	if m.Sched != nil {
		if err := m.Sched.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sched")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sched")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this dto subscription based on the context it is used
func (m *DtoSubscription) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePlotting(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSched(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DtoSubscription) contextValidatePlotting(ctx context.Context, formats strfmt.Registry) error {

	if m.Plotting != nil {

		if swag.IsZero(m.Plotting) { // not required
			return nil
		}

		if err := m.Plotting.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("plotting")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("plotting")
			}
			return err
		}
	}

	return nil
}

func (m *DtoSubscription) contextValidateSched(ctx context.Context, formats strfmt.Registry) error {

	if m.Sched != nil {

		if swag.IsZero(m.Sched) { // not required
			return nil
		}

		if err := m.Sched.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sched")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sched")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DtoSubscription) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DtoSubscription) UnmarshalBinary(b []byte) error {
	var res DtoSubscription
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
