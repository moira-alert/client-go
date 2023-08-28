// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DtoPatternData dto pattern data
//
// swagger:model dto.PatternData
type DtoPatternData struct {

	// metrics
	// Example: ["DevOps.my_server.hdd.freespace_mbytes"," DevOps.my_server.hdd.freespace_mbytes"," DevOps.my_server.db.*"]
	Metrics []string `json:"metrics"`

	// pattern
	// Example: Devops.my_server.*
	Pattern string `json:"pattern,omitempty"`

	// triggers
	Triggers []*DtoTriggerModel `json:"triggers"`
}

// Validate validates this dto pattern data
func (m *DtoPatternData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTriggers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DtoPatternData) validateTriggers(formats strfmt.Registry) error {
	if swag.IsZero(m.Triggers) { // not required
		return nil
	}

	for i := 0; i < len(m.Triggers); i++ {
		if swag.IsZero(m.Triggers[i]) { // not required
			continue
		}

		if m.Triggers[i] != nil {
			if err := m.Triggers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("triggers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("triggers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this dto pattern data based on the context it is used
func (m *DtoPatternData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTriggers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DtoPatternData) contextValidateTriggers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Triggers); i++ {

		if m.Triggers[i] != nil {

			if swag.IsZero(m.Triggers[i]) { // not required
				return nil
			}

			if err := m.Triggers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("triggers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("triggers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DtoPatternData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DtoPatternData) UnmarshalBinary(b []byte) error {
	var res DtoPatternData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}