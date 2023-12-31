// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MoiraScheduleDataDay moira schedule data day
//
// swagger:model moira.ScheduleDataDay
type MoiraScheduleDataDay struct {

	// enabled
	// Example: true
	Enabled bool `json:"enabled,omitempty"`

	// name
	// Example: Mon
	Name string `json:"name,omitempty"`
}

// Validate validates this moira schedule data day
func (m *MoiraScheduleDataDay) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this moira schedule data day based on context it is used
func (m *MoiraScheduleDataDay) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MoiraScheduleDataDay) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MoiraScheduleDataDay) UnmarshalBinary(b []byte) error {
	var res MoiraScheduleDataDay
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
