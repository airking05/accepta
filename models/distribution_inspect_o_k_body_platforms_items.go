// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DistributionInspectOKBodyPlatformsItems distribution inspect o k body platforms items
// swagger:model distributionInspectOKBodyPlatformsItems
type DistributionInspectOKBodyPlatformsItems struct {

	// architecture
	Architecture string `json:"Architecture,omitempty"`

	// features
	Features []string `json:"Features"`

	// o s
	OS string `json:"OS,omitempty"`

	// o s features
	OSFeatures []string `json:"OSFeatures"`

	// o s version
	OSVersion string `json:"OSVersion,omitempty"`

	// variant
	Variant string `json:"Variant,omitempty"`
}

// Validate validates this distribution inspect o k body platforms items
func (m *DistributionInspectOKBodyPlatformsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFeatures(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOSFeatures(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DistributionInspectOKBodyPlatformsItems) validateFeatures(formats strfmt.Registry) error {

	if swag.IsZero(m.Features) { // not required
		return nil
	}

	return nil
}

func (m *DistributionInspectOKBodyPlatformsItems) validateOSFeatures(formats strfmt.Registry) error {

	if swag.IsZero(m.OSFeatures) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DistributionInspectOKBodyPlatformsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DistributionInspectOKBodyPlatformsItems) UnmarshalBinary(b []byte) error {
	var res DistributionInspectOKBodyPlatformsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
