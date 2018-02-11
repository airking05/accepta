// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GenericResourcesItems generic resources items
// swagger:model genericResourcesItems
type GenericResourcesItems struct {

	// discrete resource spec
	DiscreteResourceSpec *GenericResourcesItemsDiscreteResourceSpec `json:"DiscreteResourceSpec,omitempty"`

	// named resource spec
	NamedResourceSpec *GenericResourcesItemsNamedResourceSpec `json:"NamedResourceSpec,omitempty"`
}

// Validate validates this generic resources items
func (m *GenericResourcesItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDiscreteResourceSpec(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNamedResourceSpec(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GenericResourcesItems) validateDiscreteResourceSpec(formats strfmt.Registry) error {

	if swag.IsZero(m.DiscreteResourceSpec) { // not required
		return nil
	}

	if m.DiscreteResourceSpec != nil {

		if err := m.DiscreteResourceSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DiscreteResourceSpec")
			}
			return err
		}

	}

	return nil
}

func (m *GenericResourcesItems) validateNamedResourceSpec(formats strfmt.Registry) error {

	if swag.IsZero(m.NamedResourceSpec) { // not required
		return nil
	}

	if m.NamedResourceSpec != nil {

		if err := m.NamedResourceSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("NamedResourceSpec")
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GenericResourcesItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GenericResourcesItems) UnmarshalBinary(b []byte) error {
	var res GenericResourcesItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
