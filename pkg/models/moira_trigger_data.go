// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MoiraTriggerData moira trigger data
//
// swagger:model moira.TriggerData
type MoiraTriggerData struct {

	// notifier trigger tags
	// Example: ["server","disk"]
	NotifierTriggerTags []string `json:"__notifier_trigger_tags"`

	// cluster id
	// Example: default
	ClusterID string `json:"cluster_id,omitempty"`

	// desc
	// Example: check the size of /var/log
	Desc string `json:"desc,omitempty"`

	// error value
	// Example: 1000
	ErrorValue float64 `json:"error_value,omitempty"`

	// id
	// Example: 292516ed-4924-4154-a62c-ebe312431fce
	ID string `json:"id,omitempty"`

	// is remote
	// Example: false
	IsRemote bool `json:"is_remote,omitempty"`

	// name
	// Example: Not enough disk space left
	Name string `json:"name,omitempty"`

	// targets
	// Example: ["devOps.my_server.hdd.freespace_mbytes"]
	Targets []string `json:"targets"`

	// trigger source
	// Example: graphite_local
	TriggerSource string `json:"trigger_source,omitempty"`

	// warn value
	// Example: 5000
	WarnValue float64 `json:"warn_value,omitempty"`
}

// Validate validates this moira trigger data
func (m *MoiraTriggerData) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this moira trigger data based on context it is used
func (m *MoiraTriggerData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MoiraTriggerData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MoiraTriggerData) UnmarshalBinary(b []byte) error {
	var res MoiraTriggerData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
