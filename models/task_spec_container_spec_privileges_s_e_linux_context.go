// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TaskSpecContainerSpecPrivilegesSELinuxContext SELinux labels of the container
// swagger:model taskSpecContainerSpecPrivilegesSELinuxContext
type TaskSpecContainerSpecPrivilegesSELinuxContext struct {

	// Disable SELinux
	Disable bool `json:"Disable,omitempty"`

	// SELinux level label
	Level string `json:"Level,omitempty"`

	// SELinux role label
	Role string `json:"Role,omitempty"`

	// SELinux type label
	Type string `json:"Type,omitempty"`

	// SELinux user label
	User string `json:"User,omitempty"`
}

// Validate validates this task spec container spec privileges s e linux context
func (m *TaskSpecContainerSpecPrivilegesSELinuxContext) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *TaskSpecContainerSpecPrivilegesSELinuxContext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskSpecContainerSpecPrivilegesSELinuxContext) UnmarshalBinary(b []byte) error {
	var res TaskSpecContainerSpecPrivilegesSELinuxContext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
