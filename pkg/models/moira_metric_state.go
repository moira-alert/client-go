// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MoiraMetricState moira metric state
//
// swagger:model moira.MetricState
type MoiraMetricState struct {

	// DeletedButKept controls whether the metric is shown to the user if the trigger has ttlState = Del
	// and the metric is in Maintenance. The metric remains in the database
	// Example: false
	DeletedButKept bool `json:"deleted_but_kept,omitempty"`

	// event timestamp
	// Example: 1590741878
	EventTimestamp int64 `json:"event_timestamp,omitempty"`

	// maintenance
	// Example: 0
	Maintenance int64 `json:"maintenance,omitempty"`

	// maintenance info
	MaintenanceInfo *MoiraMaintenanceInfo `json:"maintenance_info,omitempty"`

	// state
	// Example: OK
	State string `json:"state,omitempty"`

	// suppressed
	// Example: false
	Suppressed bool `json:"suppressed,omitempty"`

	// suppressed state
	SuppressedState string `json:"suppressed_state,omitempty"`

	// timestamp
	// Example: 1590741878
	Timestamp int64 `json:"timestamp,omitempty"`

	// value
	// Example: 70
	Value *float64 `json:"value,omitempty"`

	// values
	Values map[string]float64 `json:"values,omitempty"`
}

// Validate validates this moira metric state
func (m *MoiraMetricState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMaintenanceInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MoiraMetricState) validateMaintenanceInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.MaintenanceInfo) { // not required
		return nil
	}

	if m.MaintenanceInfo != nil {
		if err := m.MaintenanceInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("maintenance_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("maintenance_info")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this moira metric state based on the context it is used
func (m *MoiraMetricState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMaintenanceInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MoiraMetricState) contextValidateMaintenanceInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.MaintenanceInfo != nil {

		if swag.IsZero(m.MaintenanceInfo) { // not required
			return nil
		}

		if err := m.MaintenanceInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("maintenance_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("maintenance_info")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MoiraMetricState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MoiraMetricState) UnmarshalBinary(b []byte) error {
	var res MoiraMetricState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
