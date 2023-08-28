// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DtoSaveTeamResponse dto save team response
//
// swagger:model dto.SaveTeamResponse
type DtoSaveTeamResponse struct {

	// id
	// Example: d5d98eb3-ee18-4f75-9364-244f67e23b54
	ID string `json:"id,omitempty"`
}

// Validate validates this dto save team response
func (m *DtoSaveTeamResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this dto save team response based on context it is used
func (m *DtoSaveTeamResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DtoSaveTeamResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DtoSaveTeamResponse) UnmarshalBinary(b []byte) error {
	var res DtoSaveTeamResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}