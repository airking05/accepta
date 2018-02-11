// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// PluginConfigInterfaceTypes plugin config interface types
// swagger:model pluginConfigInterfaceTypes
type PluginConfigInterfaceTypes []PluginInterfaceType

// Validate validates this plugin config interface types
func (m PluginConfigInterfaceTypes) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
