// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DtoTeamMembers dto team members
//
// swagger:model dto.TeamMembers
type DtoTeamMembers struct {

	// usernames
	// Example: ["anonymous"]
	Usernames []string `json:"usernames"`
}

// Validate validates this dto team members
func (m *DtoTeamMembers) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this dto team members based on context it is used
func (m *DtoTeamMembers) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DtoTeamMembers) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DtoTeamMembers) UnmarshalBinary(b []byte) error {
	var res DtoTeamMembers
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}