// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// DtoTypeOfProblem dto type of problem
//
// swagger:model dto.typeOfProblem
type DtoTypeOfProblem string

func NewDtoTypeOfProblem(value DtoTypeOfProblem) *DtoTypeOfProblem {
	return &value
}

// Pointer returns a pointer to a freshly-allocated DtoTypeOfProblem.
func (m DtoTypeOfProblem) Pointer() *DtoTypeOfProblem {
	return &m
}

const (

	// DtoTypeOfProblemWarn captures enum value "warn"
	DtoTypeOfProblemWarn DtoTypeOfProblem = "warn"

	// DtoTypeOfProblemBad captures enum value "bad"
	DtoTypeOfProblemBad DtoTypeOfProblem = "bad"
)

// for schema
var dtoTypeOfProblemEnum []interface{}

func init() {
	var res []DtoTypeOfProblem
	if err := json.Unmarshal([]byte(`["warn","bad"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dtoTypeOfProblemEnum = append(dtoTypeOfProblemEnum, v)
	}
}

func (m DtoTypeOfProblem) validateDtoTypeOfProblemEnum(path, location string, value DtoTypeOfProblem) error {
	if err := validate.EnumCase(path, location, value, dtoTypeOfProblemEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this dto type of problem
func (m DtoTypeOfProblem) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDtoTypeOfProblemEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this dto type of problem based on context it is used
func (m DtoTypeOfProblem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
