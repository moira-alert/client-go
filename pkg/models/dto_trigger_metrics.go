// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DtoTriggerMetrics dto trigger metrics
//
// swagger:model dto.TriggerMetrics
type DtoTriggerMetrics map[string]map[string][]MoiraMetricValue

// Validate validates this dto trigger metrics
func (m DtoTriggerMetrics) Validate(formats strfmt.Registry) error {
	var res []error

	for k := range m {

		for kk := range m[k] {

			if err := validate.Required(k+"."+kk, "body", m[k][kk]); err != nil {
				return err
			}

			for i := 0; i < len(m[k][kk]); i++ {

				if err := m[k][kk][i].Validate(formats); err != nil {
					if ve, ok := err.(*errors.Validation); ok {
						return ve.ValidateName(k + "." + kk + "." + strconv.Itoa(i))
					} else if ce, ok := err.(*errors.CompositeError); ok {
						return ce.ValidateName(k + "." + kk + "." + strconv.Itoa(i))
					}
					return err
				}

			}

		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this dto trigger metrics based on the context it is used
func (m DtoTriggerMetrics) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	for k := range m {

		for kk := range m[k] {

			for i := 0; i < len(m[k][kk]); i++ {

				if swag.IsZero(m[k][kk][i]) { // not required
					return nil
				}

				if err := m[k][kk][i].ContextValidate(ctx, formats); err != nil {
					if ve, ok := err.(*errors.Validation); ok {
						return ve.ValidateName(k + "." + kk + "." + strconv.Itoa(i))
					} else if ce, ok := err.(*errors.CompositeError); ok {
						return ce.ValidateName(k + "." + kk + "." + strconv.Itoa(i))
					}
					return err
				}

			}

		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}