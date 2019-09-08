// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Standard Standard
// swagger:model Standard
type Standard struct {

	// Limit value of the pollutant
	Limit *float64 `json:"limit,omitempty"`

	// Name of this standard
	Name *string `json:"name,omitempty"`

	// Pollutant measurement as percent of allowable limit
	Percent *float64 `json:"percent,omitempty"`

	// Pollutant described by this standard
	Pollutant *string `json:"pollutant,omitempty"`
}

// Validate validates this standard
func (m *Standard) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Standard) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Standard) UnmarshalBinary(b []byte) error {
	var res Standard
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}