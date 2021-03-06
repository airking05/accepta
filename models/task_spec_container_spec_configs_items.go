// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TaskSpecContainerSpecConfigsItems task spec container spec configs items
// swagger:model taskSpecContainerSpecConfigsItems
type TaskSpecContainerSpecConfigsItems struct {

	// ConfigID represents the ID of the specific config that we're referencing.
	ConfigID string `json:"ConfigID,omitempty"`

	// ConfigName is the name of the config that this references, but this is just provided for
	// lookup/display purposes. The config in the reference will be identified by its ID.
	//
	ConfigName string `json:"ConfigName,omitempty"`

	// file
	File *TaskSpecContainerSpecConfigsItemsFile `json:"File,omitempty"`
}

// Validate validates this task spec container spec configs items
func (m *TaskSpecContainerSpecConfigsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFile(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskSpecContainerSpecConfigsItems) validateFile(formats strfmt.Registry) error {

	if swag.IsZero(m.File) { // not required
		return nil
	}

	if m.File != nil {

		if err := m.File.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("File")
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TaskSpecContainerSpecConfigsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskSpecContainerSpecConfigsItems) UnmarshalBinary(b []byte) error {
	var res TaskSpecContainerSpecConfigsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
