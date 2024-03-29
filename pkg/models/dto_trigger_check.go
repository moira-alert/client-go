// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DtoTriggerCheck dto trigger check
//
// swagger:model dto.TriggerCheck
type DtoTriggerCheck struct {

	// event timestamp
	// Example: 1590741878
	EventTimestamp int64 `json:"event_timestamp,omitempty"`

	// LastSuccessfulCheckTimestamp - time of the last check of the trigger, during which there were no errors
	// Example: 1590741916
	LastSuccessfulCheckTimestamp int64 `json:"last_successful_check_timestamp,omitempty"`

	// maintenance
	// Example: 0
	Maintenance int64 `json:"maintenance,omitempty"`

	// maintenance info
	MaintenanceInfo *MoiraMaintenanceInfo `json:"maintenance_info,omitempty"`

	// metrics
	Metrics map[string]MoiraMetricState `json:"metrics,omitempty"`

	// MetricsToTargetRelation is a map that holds relation between metric names that was alone during last
	// check and targets that fetched this metric
	// 	{"t1": "metric.name.1", "t2": "metric.name.2"}
	// Example: {"t1":"metric.name.1","t2":"metric.name.2"}
	MetricsToTargetRelation map[string]string `json:"metrics_to_target_relation,omitempty"`

	// msg
	Msg string `json:"msg,omitempty"`

	// score
	// Example: 100
	Score int64 `json:"score,omitempty"`

	// state
	// Example: OK
	State string `json:"state,omitempty"`

	// suppressed
	// Example: true
	Suppressed bool `json:"suppressed,omitempty"`

	// suppressed state
	SuppressedState string `json:"suppressed_state,omitempty"`

	// Timestamp - time, which means when the checker last checked this trigger, this value stops updating if the trigger does not receive metrics
	// Example: 1590741916
	Timestamp int64 `json:"timestamp,omitempty"`

	// trigger id
	// Example: trigger_id
	TriggerID string `json:"trigger_id,omitempty"`
}

// Validate validates this dto trigger check
func (m *DtoTriggerCheck) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMaintenanceInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetrics(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DtoTriggerCheck) validateMaintenanceInfo(formats strfmt.Registry) error {
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

func (m *DtoTriggerCheck) validateMetrics(formats strfmt.Registry) error {
	if swag.IsZero(m.Metrics) { // not required
		return nil
	}

	for k := range m.Metrics {

		if err := validate.Required("metrics"+"."+k, "body", m.Metrics[k]); err != nil {
			return err
		}
		if val, ok := m.Metrics[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metrics" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("metrics" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this dto trigger check based on the context it is used
func (m *DtoTriggerCheck) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMaintenanceInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetrics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DtoTriggerCheck) contextValidateMaintenanceInfo(ctx context.Context, formats strfmt.Registry) error {

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

func (m *DtoTriggerCheck) contextValidateMetrics(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Metrics {

		if val, ok := m.Metrics[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DtoTriggerCheck) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DtoTriggerCheck) UnmarshalBinary(b []byte) error {
	var res DtoTriggerCheck
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
