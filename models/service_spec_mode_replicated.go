// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ServiceSpecModeReplicated service spec mode replicated
// swagger:model serviceSpecModeReplicated
type ServiceSpecModeReplicated struct {

	// replicas
	Replicas int64 `json:"Replicas,omitempty"`
}

// Validate validates this service spec mode replicated
func (m *ServiceSpecModeReplicated) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceSpecModeReplicated) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceSpecModeReplicated) UnmarshalBinary(b []byte) error {
	var res ServiceSpecModeReplicated
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
