// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DtoUser dto user
//
// swagger:model dto.User
type DtoUser struct {

	// auth enabled
	// Example: true
	AuthEnabled bool `json:"auth_enabled,omitempty"`

	// login
	// Example: john
	Login string `json:"login,omitempty"`

	// role
	// Example: user
	Role string `json:"role,omitempty"`
}

// Validate validates this dto user
func (m *DtoUser) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this dto user based on context it is used
func (m *DtoUser) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DtoUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DtoUser) UnmarshalBinary(b []byte) error {
	var res DtoUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
