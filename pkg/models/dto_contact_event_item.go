// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DtoContactEventItem dto contact event item
//
// swagger:model dto.ContactEventItem
type DtoContactEventItem struct {

	// metric
	Metric string `json:"metric,omitempty"`

	// old state
	OldState string `json:"old_state,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// timestamp
	Timestamp int64 `json:"timestamp,omitempty"`

	// trigger id
	TriggerID string `json:"trigger_id,omitempty"`
}

// Validate validates this dto contact event item
func (m *DtoContactEventItem) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this dto contact event item based on context it is used
func (m *DtoContactEventItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DtoContactEventItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DtoContactEventItem) UnmarshalBinary(b []byte) error {
	var res DtoContactEventItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
