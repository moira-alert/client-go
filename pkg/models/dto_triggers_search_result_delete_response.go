// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DtoTriggersSearchResultDeleteResponse dto triggers search result delete response
//
// swagger:model dto.TriggersSearchResultDeleteResponse
type DtoTriggersSearchResultDeleteResponse struct {

	// pager id
	// Example: 292516ed-4924-4154-a62c-ebe312431fce
	PagerID string `json:"pager_id,omitempty"`
}

// Validate validates this dto triggers search result delete response
func (m *DtoTriggersSearchResultDeleteResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this dto triggers search result delete response based on context it is used
func (m *DtoTriggersSearchResultDeleteResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DtoTriggersSearchResultDeleteResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DtoTriggersSearchResultDeleteResponse) UnmarshalBinary(b []byte) error {
	var res DtoTriggersSearchResultDeleteResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}