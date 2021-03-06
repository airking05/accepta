// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SwarmSpecTaskDefaults Defaults for creating tasks in this cluster.
// swagger:model swarmSpecTaskDefaults
type SwarmSpecTaskDefaults struct {

	// log driver
	LogDriver *SwarmSpecTaskDefaultsLogDriver `json:"LogDriver,omitempty"`
}

// Validate validates this swarm spec task defaults
func (m *SwarmSpecTaskDefaults) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLogDriver(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SwarmSpecTaskDefaults) validateLogDriver(formats strfmt.Registry) error {

	if swag.IsZero(m.LogDriver) { // not required
		return nil
	}

	if m.LogDriver != nil {

		if err := m.LogDriver.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("LogDriver")
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SwarmSpecTaskDefaults) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SwarmSpecTaskDefaults) UnmarshalBinary(b []byte) error {
	var res SwarmSpecTaskDefaults
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
