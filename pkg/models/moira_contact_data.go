// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MoiraContactData moira contact data
//
// swagger:model moira.ContactData
type MoiraContactData struct {

	// id
	// Example: 1dd38765-c5be-418d-81fa-7a5f879c2315
	ID string `json:"id,omitempty"`

	// name
	// Example: Mail Alerts
	Name string `json:"name,omitempty"`

	// team
	Team string `json:"team,omitempty"`

	// type
	// Example: mail
	Type string `json:"type,omitempty"`

	// user
	User string `json:"user,omitempty"`

	// value
	// Example: devops@example.com
	Value string `json:"value,omitempty"`
}

// Validate validates this moira contact data
func (m *MoiraContactData) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this moira contact data based on context it is used
func (m *MoiraContactData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MoiraContactData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MoiraContactData) UnmarshalBinary(b []byte) error {
	var res MoiraContactData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
