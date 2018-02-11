// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ContainerPruneOKBody container prune o k body
// swagger:model containerPruneOKBody
type ContainerPruneOKBody struct {

	// Container IDs that were deleted
	ContainersDeleted []string `json:"ContainersDeleted"`

	// Disk space reclaimed in bytes
	SpaceReclaimed int64 `json:"SpaceReclaimed,omitempty"`
}

// Validate validates this container prune o k body
func (m *ContainerPruneOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContainersDeleted(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContainerPruneOKBody) validateContainersDeleted(formats strfmt.Registry) error {

	if swag.IsZero(m.ContainersDeleted) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContainerPruneOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContainerPruneOKBody) UnmarshalBinary(b []byte) error {
	var res ContainerPruneOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}