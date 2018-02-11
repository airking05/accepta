// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SwarmInitParamsBody swarm init params body
// swagger:model swarmInitParamsBody
type SwarmInitParamsBody struct {

	// Externally reachable address advertised to other nodes. This can either be an address/port combination in the form `192.168.1.1:4567`, or an interface followed by a port number, like `eth0:4567`. If the port number is omitted, the port number from the listen address is used. If `AdvertiseAddr` is not specified, it will be automatically detected when possible.
	AdvertiseAddr string `json:"AdvertiseAddr,omitempty"`

	// Address or interface to use for data path traffic (format: `<ip|interface>`), for example,  `192.168.1.1`,
	// or an interface, like `eth0`. If `DataPathAddr` is unspecified, the same address as `AdvertiseAddr`
	// is used.
	//
	// The `DataPathAddr` specifies the address that global scope network drivers will publish towards other
	// nodes in order to reach the containers running on this node. Using this parameter it is possible to
	// separate the container data traffic from the management traffic of the cluster.
	//
	DataPathAddr string `json:"DataPathAddr,omitempty"`

	// Force creation of a new swarm.
	ForceNewCluster bool `json:"ForceNewCluster,omitempty"`

	// Listen address used for inter-manager communication, as well as determining the networking interface used for the VXLAN Tunnel Endpoint (VTEP). This can either be an address/port combination in the form `192.168.1.1:4567`, or an interface followed by a port number, like `eth0:4567`. If the port number is omitted, the default swarm listening port is used.
	ListenAddr string `json:"ListenAddr,omitempty"`

	// spec
	Spec *SwarmSpec `json:"Spec,omitempty"`
}

// Validate validates this swarm init params body
func (m *SwarmInitParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSpec(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SwarmInitParamsBody) validateSpec(formats strfmt.Registry) error {

	if swag.IsZero(m.Spec) { // not required
		return nil
	}

	if m.Spec != nil {

		if err := m.Spec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Spec")
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SwarmInitParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SwarmInitParamsBody) UnmarshalBinary(b []byte) error {
	var res SwarmInitParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}