// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIErrorInvalidRequestExample api error invalid request example
//
// swagger:model api.ErrorInvalidRequestExample
type APIErrorInvalidRequestExample struct {

	// error
	// Example: resource with the ID does not exist
	Error string `json:"error,omitempty"`

	// status
	// Example: Invalid request
	Status string `json:"status,omitempty"`
}

// Validate validates this api error invalid request example
func (m *APIErrorInvalidRequestExample) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this api error invalid request example based on context it is used
func (m *APIErrorInvalidRequestExample) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIErrorInvalidRequestExample) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIErrorInvalidRequestExample) UnmarshalBinary(b []byte) error {
	var res APIErrorInvalidRequestExample
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
