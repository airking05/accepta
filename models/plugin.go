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

// Plugin A plugin for the Engine API
// swagger:model Plugin
type Plugin struct {

	// config
	// Required: true
	Config PluginConfig `json:"Config"`

	// True if the plugin is running. False if the plugin is not running, only installed.
	// Required: true
	Enabled bool `json:"Enabled"`

	// Id
	ID string `json:"Id,omitempty"`

	// name
	// Required: true
	Name string `json:"Name"`

	// plugin remote reference used to push/pull the plugin
	PluginReference string `json:"PluginReference,omitempty"`

	// settings
	// Required: true
	Settings PluginSettings `json:"Settings"`
}

// Validate validates this plugin
func (m *Plugin) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSettings(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Plugin) validateConfig(formats strfmt.Registry) error {

	if err := m.Config.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Config")
		}
		return err
	}

	return nil
}

func (m *Plugin) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("Enabled", "body", bool(m.Enabled)); err != nil {
		return err
	}

	return nil
}

func (m *Plugin) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("Name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *Plugin) validateSettings(formats strfmt.Registry) error {

	if err := m.Settings.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Settings")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Plugin) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Plugin) UnmarshalBinary(b []byte) error {
	var res Plugin
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
