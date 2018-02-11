// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SwarmSpecCAConfigExternalCAsItems swarm spec c a config external c as items
// swagger:model swarmSpecCAConfigExternalCAsItems
type SwarmSpecCAConfigExternalCAsItems struct {

	// The root CA certificate (in PEM format) this external CA uses to issue TLS certificates (assumed to be to the current swarm root CA certificate if not provided).
	CACert string `json:"CACert,omitempty"`

	// An object with key/value pairs that are interpreted as protocol-specific options for the external CA driver.
	Options map[string]string `json:"Options,omitempty"`

	// Protocol for communication with the external CA (currently only `cfssl` is supported).
	Protocol *string `json:"Protocol,omitempty"`

	// URL where certificate signing requests should be sent.
	URL string `json:"URL,omitempty"`
}

// Validate validates this swarm spec c a config external c as items
func (m *SwarmSpecCAConfigExternalCAsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProtocol(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var swarmSpecCAConfigExternalCAsItemsTypeProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cfssl"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		swarmSpecCAConfigExternalCAsItemsTypeProtocolPropEnum = append(swarmSpecCAConfigExternalCAsItemsTypeProtocolPropEnum, v)
	}
}

const (

	// SwarmSpecCAConfigExternalCAsItemsProtocolCfssl captures enum value "cfssl"
	SwarmSpecCAConfigExternalCAsItemsProtocolCfssl string = "cfssl"
)

// prop value enum
func (m *SwarmSpecCAConfigExternalCAsItems) validateProtocolEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, swarmSpecCAConfigExternalCAsItemsTypeProtocolPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SwarmSpecCAConfigExternalCAsItems) validateProtocol(formats strfmt.Registry) error {

	if swag.IsZero(m.Protocol) { // not required
		return nil
	}

	// value enum
	if err := m.validateProtocolEnum("Protocol", "body", *m.Protocol); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SwarmSpecCAConfigExternalCAsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SwarmSpecCAConfigExternalCAsItems) UnmarshalBinary(b []byte) error {
	var res SwarmSpecCAConfigExternalCAsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}