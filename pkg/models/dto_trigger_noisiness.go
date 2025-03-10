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

// DtoTriggerNoisiness dto trigger noisiness
//
// swagger:model dto.TriggerNoisiness
type DtoTriggerNoisiness struct {

	// A list of targets that have only alone metrics
	// Example: {"t1":true}
	AloneMetrics map[string]bool `json:"alone_metrics,omitempty"`

	// Shows the exact cluster from where the metrics are fetched
	// Example: default
	ClusterID string `json:"cluster_id,omitempty"`

	// Datetime when the trigger was created
	CreatedAt *string `json:"created_at,omitempty"`

	// Username who created trigger
	CreatedBy string `json:"created_by,omitempty"`

	// Description string
	// Example: check the size of /var/log
	Desc *string `json:"desc,omitempty"`

	// ERROR threshold
	// Example: 1000
	ErrorValue *float64 `json:"error_value,omitempty"`

	// EventsCount for the trigger.
	EventsCount int64 `json:"events_count,omitempty"`

	// Used if you need more complex logic than provided by WARN/ERROR values
	Expression string `json:"expression,omitempty"`

	// Trigger unique ID
	// Example: 292516ed-4924-4154-a62c-ebe312431fce
	ID string `json:"id,omitempty"`

	// Shows if trigger is remote (graphite-backend) based or stored inside Moira-Redis DB
	//
	// Deprecated: Use TriggerSource field instead
	// Example: false
	IsRemote bool `json:"is_remote,omitempty"`

	// If true, first event NODATA → OK will be omitted
	// Example: false
	MuteNewMetrics bool `json:"mute_new_metrics,omitempty"`

	// Trigger name
	// Example: Not enough disk space left
	Name string `json:"name,omitempty"`

	// Graphite patterns for trigger
	// Example: [""]
	Patterns []string `json:"patterns"`

	// Determines when Moira should monitor trigger
	Sched struct {
		MoiraScheduleData
	} `json:"sched,omitempty"`

	// Set of tags to manipulate subscriptions
	// Example: ["server","disk"]
	Tags []string `json:"tags"`

	// Graphite-like targets: t1, t2, ...
	// Example: ["devOps.my_server.hdd.freespace_mbytes"]
	Targets []string `json:"targets"`

	// throttling
	// Example: 0
	Throttling int64 `json:"throttling,omitempty"`

	// Shows the type of source from where the metrics are fetched
	// Example: graphite_local
	TriggerSource string `json:"trigger_source,omitempty"`

	// Could be: rising, falling, expression
	// Example: rising
	TriggerType string `json:"trigger_type,omitempty"`

	// When there are no metrics for trigger, Moira will switch metric to TTLState state after TTL seconds
	// Example: 600
	TTL int64 `json:"ttl,omitempty"`

	// When there are no metrics for trigger, Moira will switch metric to TTLState state after TTL seconds
	// Example: NODATA
	TTLState *string `json:"ttl_state,omitempty"`

	// Datetime  when the trigger was updated
	UpdatedAt *string `json:"updated_at,omitempty"`

	// Username who updated trigger
	UpdatedBy string `json:"updated_by,omitempty"`

	// WARN threshold
	// Example: 500
	WarnValue *float64 `json:"warn_value,omitempty"`
}

// Validate validates this dto trigger noisiness
func (m *DtoTriggerNoisiness) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSched(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DtoTriggerNoisiness) validateSched(formats strfmt.Registry) error {
	if swag.IsZero(m.Sched) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this dto trigger noisiness based on the context it is used
func (m *DtoTriggerNoisiness) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSched(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DtoTriggerNoisiness) contextValidateSched(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *DtoTriggerNoisiness) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DtoTriggerNoisiness) UnmarshalBinary(b []byte) error {
	var res DtoTriggerNoisiness
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
