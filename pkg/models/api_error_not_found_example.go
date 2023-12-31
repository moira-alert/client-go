// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIErrorNotFoundExample api error not found example
//
// swagger:model api.ErrorNotFoundExample
type APIErrorNotFoundExample struct {

	// error
	// Example: resource with ID '66741a8c-c2ba-4357-a2c9-ee78e0e7' does not exist
	Error string `json:"error,omitempty"`

	// status
	// Example: Resource not found
	Status string `json:"status,omitempty"`
}

// Validate validates this api error not found example
func (m *APIErrorNotFoundExample) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this api error not found example based on context it is used
func (m *APIErrorNotFoundExample) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIErrorNotFoundExample) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIErrorNotFoundExample) UnmarshalBinary(b []byte) error {
	var res APIErrorNotFoundExample
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
