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

// ContainerChangeResponseItem ContainerChangeResponseItem
//
// change item in response to ContainerChanges operation
// swagger:model ContainerChangeResponseItem
type ContainerChangeResponseItem struct {

	// Kind of change
	// Required: true
	Kind uint8 `json:"Kind"`

	// Path to file that has changed
	// Required: true
	Path string `json:"Path"`
}

// Validate validates this container change response item
func (m *ContainerChangeResponseItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKind(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePath(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var containerChangeResponseItemTypeKindPropEnum []interface{}

func init() {
	var res []uint8
	if err := json.Unmarshal([]byte(`[0,1,2]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		containerChangeResponseItemTypeKindPropEnum = append(containerChangeResponseItemTypeKindPropEnum, v)
	}
}

// prop value enum
func (m *ContainerChangeResponseItem) validateKindEnum(path, location string, value uint8) error {
	if err := validate.Enum(path, location, value, containerChangeResponseItemTypeKindPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ContainerChangeResponseItem) validateKind(formats strfmt.Registry) error {

	if err := validate.Required("Kind", "body", uint8(m.Kind)); err != nil {
		return err
	}

	// value enum
	if err := m.validateKindEnum("Kind", "body", m.Kind); err != nil {
		return err
	}

	return nil
}

func (m *ContainerChangeResponseItem) validatePath(formats strfmt.Registry) error {

	if err := validate.RequiredString("Path", "body", string(m.Path)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContainerChangeResponseItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContainerChangeResponseItem) UnmarshalBinary(b []byte) error {
	var res ContainerChangeResponseItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
