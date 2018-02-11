// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PluginDevice plugin device
// swagger:model PluginDevice
type PluginDevice struct {

	// description
	// Required: true
	Description string `json:"Description"`

	// name
	// Required: true
	Name string `json:"Name"`

	// path
	// Required: true
	Path *string `json:"Path"`

	// settable
	// Required: true
	Settable []string `json:"Settable"`
}

// Validate validates this plugin device
func (m *PluginDevice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePath(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSettable(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PluginDevice) validateDescription(formats strfmt.Registry) error {

	if err := validate.RequiredString("Description", "body", string(m.Description)); err != nil {
		return err
	}

	return nil
}

func (m *PluginDevice) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("Name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *PluginDevice) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("Path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

func (m *PluginDevice) validateSettable(formats strfmt.Registry) error {

	if err := validate.Required("Settable", "body", m.Settable); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PluginDevice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PluginDevice) UnmarshalBinary(b []byte) error {
	var res PluginDevice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
