// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DtoMessageResponse dto message response
//
// swagger:model dto.MessageResponse
type DtoMessageResponse struct {

	// message
	// Example: tag deleted
	Message string `json:"message,omitempty"`
}

// Validate validates this dto message response
func (m *DtoMessageResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this dto message response based on context it is used
func (m *DtoMessageResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DtoMessageResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DtoMessageResponse) UnmarshalBinary(b []byte) error {
	var res DtoMessageResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
