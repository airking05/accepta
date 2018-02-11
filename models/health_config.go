// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// HealthConfig A test to perform to check that the container is healthy.
// swagger:model HealthConfig
type HealthConfig struct {

	// The time to wait between checks in nanoseconds. It should be 0 or at least 1000000 (1 ms). 0 means inherit.
	Interval int64 `json:"Interval,omitempty"`

	// The number of consecutive failures needed to consider a container as unhealthy. 0 means inherit.
	Retries int64 `json:"Retries,omitempty"`

	// Start period for the container to initialize before starting health-retries countdown in nanoseconds. It should be 0 or at least 1000000 (1 ms). 0 means inherit.
	StartPeriod int64 `json:"StartPeriod,omitempty"`

	// The test to perform. Possible values are:
	//
	// - `[]` inherit healthcheck from image or parent image
	// - `["NONE"]` disable healthcheck
	// - `["CMD", args...]` exec arguments directly
	// - `["CMD-SHELL", command]` run command with system's default shell
	//
	Test []string `json:"Test"`

	// The time to wait before considering the check to have hung. It should be 0 or at least 1000000 (1 ms). 0 means inherit.
	Timeout int64 `json:"Timeout,omitempty"`
}

// Validate validates this health config
func (m *HealthConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTest(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HealthConfig) validateTest(formats strfmt.Registry) error {

	if swag.IsZero(m.Test) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HealthConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HealthConfig) UnmarshalBinary(b []byte) error {
	var res HealthConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
