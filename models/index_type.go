// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// IndexType IndexType
// swagger:model IndexType
type IndexType struct {

	// List of possible index levels
	Levels []*IndexLevel `json:"levels"`

	// Name of this index
	Name *string `json:"name,omitempty"`
}

// Validate validates this index type
func (m *IndexType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLevels(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IndexType) validateLevels(formats strfmt.Registry) error {

	if swag.IsZero(m.Levels) { // not required
		return nil
	}

	for i := 0; i < len(m.Levels); i++ {
		if swag.IsZero(m.Levels[i]) { // not required
			continue
		}

		if m.Levels[i] != nil {
			if err := m.Levels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("levels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IndexType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IndexType) UnmarshalBinary(b []byte) error {
	var res IndexType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
