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

// MoiraTriggerCheck moira trigger check
//
// swagger:model moira.TriggerCheck
type MoiraTriggerCheck struct {

	// alone metrics
	// Example: {"t1":true}
	AloneMetrics map[string]bool `json:"alone_metrics,omitempty"`

	// created at
	CreatedAt *int64 `json:"created_at,omitempty"`

	// created by
	CreatedBy string `json:"created_by,omitempty"`

	// desc
	// Example: check the size of /var/log
	Desc *string `json:"desc,omitempty"`

	// error value
	// Example: 1000
	ErrorValue *float64 `json:"error_value,omitempty"`

	// expression
	Expression *string `json:"expression,omitempty"`

	// highlights
	Highlights map[string]string `json:"highlights,omitempty"`

	// id
	// Example: 292516ed-4924-4154-a62c-ebe312431fce
	ID string `json:"id,omitempty"`

	// is remote
	// Example: false
	IsRemote bool `json:"is_remote,omitempty"`

	// last check
	LastCheck *MoiraCheckData `json:"last_check,omitempty"`

	// mute new metrics
	// Example: false
	MuteNewMetrics bool `json:"mute_new_metrics,omitempty"`

	// name
	// Example: Not enough disk space left
	Name string `json:"name,omitempty"`

	// patterns
	// Example: [""]
	Patterns []string `json:"patterns"`

	// python expression
	PythonExpression *string `json:"python_expression,omitempty"`

	// sched
	Sched struct {
		MoiraScheduleData
	} `json:"sched,omitempty"`

	// tags
	// Example: ["server","disk"]
	Tags []string `json:"tags"`

	// targets
	// Example: ["devOps.my_server.hdd.freespace_mbytes"]
	Targets []string `json:"targets"`

	// throttling
	// Example: 0
	Throttling int64 `json:"throttling,omitempty"`

	// trigger type
	// Example: rising
	TriggerType string `json:"trigger_type,omitempty"`

	// ttl
	// Example: 600
	TTL int64 `json:"ttl,omitempty"`

	// ttl state
	// Example: NODATA
	TTLState *string `json:"ttl_state,omitempty"`

	// updated at
	UpdatedAt *int64 `json:"updated_at,omitempty"`

	// updated by
	UpdatedBy string `json:"updated_by,omitempty"`

	// warn value
	// Example: 5000
	WarnValue *float64 `json:"warn_value,omitempty"`
}

// Validate validates this moira trigger check
func (m *MoiraTriggerCheck) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastCheck(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSched(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MoiraTriggerCheck) validateLastCheck(formats strfmt.Registry) error {
	if swag.IsZero(m.LastCheck) { // not required
		return nil
	}

	if m.LastCheck != nil {
		if err := m.LastCheck.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("last_check")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("last_check")
			}
			return err
		}
	}

	return nil
}

func (m *MoiraTriggerCheck) validateSched(formats strfmt.Registry) error {
	if swag.IsZero(m.Sched) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this moira trigger check based on the context it is used
func (m *MoiraTriggerCheck) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLastCheck(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSched(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MoiraTriggerCheck) contextValidateLastCheck(ctx context.Context, formats strfmt.Registry) error {

	if m.LastCheck != nil {

		if swag.IsZero(m.LastCheck) { // not required
			return nil
		}

		if err := m.LastCheck.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("last_check")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("last_check")
			}
			return err
		}
	}

	return nil
}

func (m *MoiraTriggerCheck) contextValidateSched(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *MoiraTriggerCheck) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MoiraTriggerCheck) UnmarshalBinary(b []byte) error {
	var res MoiraTriggerCheck
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
