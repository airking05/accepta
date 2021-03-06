// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TaskSpecContainerSpecDNSConfig Specification for DNS related configurations in resolver configuration file (`resolv.conf`).
// swagger:model taskSpecContainerSpecDnsConfig
type TaskSpecContainerSpecDNSConfig struct {

	// The IP addresses of the name servers.
	Nameservers []string `json:"Nameservers"`

	// A list of internal resolver variables to be modified (e.g., `debug`, `ndots:3`, etc.).
	Options []string `json:"Options"`

	// A search list for host-name lookup.
	Search []string `json:"Search"`
}

// Validate validates this task spec container spec Dns config
func (m *TaskSpecContainerSpecDNSConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNameservers(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOptions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSearch(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskSpecContainerSpecDNSConfig) validateNameservers(formats strfmt.Registry) error {

	if swag.IsZero(m.Nameservers) { // not required
		return nil
	}

	return nil
}

func (m *TaskSpecContainerSpecDNSConfig) validateOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.Options) { // not required
		return nil
	}

	return nil
}

func (m *TaskSpecContainerSpecDNSConfig) validateSearch(formats strfmt.Registry) error {

	if swag.IsZero(m.Search) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TaskSpecContainerSpecDNSConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskSpecContainerSpecDNSConfig) UnmarshalBinary(b []byte) error {
	var res TaskSpecContainerSpecDNSConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
