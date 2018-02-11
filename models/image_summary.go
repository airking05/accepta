// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ImageSummary image summary
// swagger:model ImageSummary
type ImageSummary struct {

	// containers
	// Required: true
	Containers int64 `json:"Containers"`

	// created
	// Required: true
	Created int64 `json:"Created"`

	// Id
	// Required: true
	ID string `json:"Id"`

	// labels
	// Required: true
	Labels map[string]string `json:"Labels"`

	// parent Id
	// Required: true
	ParentID string `json:"ParentId"`

	// repo digests
	// Required: true
	RepoDigests []string `json:"RepoDigests"`

	// repo tags
	// Required: true
	RepoTags []string `json:"RepoTags"`

	// shared size
	// Required: true
	SharedSize int64 `json:"SharedSize"`

	// size
	// Required: true
	Size int64 `json:"Size"`

	// virtual size
	// Required: true
	VirtualSize int64 `json:"VirtualSize"`
}

// Validate validates this image summary
func (m *ImageSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContainers(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateParentID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRepoDigests(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRepoTags(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSharedSize(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVirtualSize(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImageSummary) validateContainers(formats strfmt.Registry) error {

	if err := validate.Required("Containers", "body", int64(m.Containers)); err != nil {
		return err
	}

	return nil
}

func (m *ImageSummary) validateCreated(formats strfmt.Registry) error {

	if err := validate.Required("Created", "body", int64(m.Created)); err != nil {
		return err
	}

	return nil
}

func (m *ImageSummary) validateID(formats strfmt.Registry) error {

	if err := validate.RequiredString("Id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *ImageSummary) validateLabels(formats strfmt.Registry) error {

	return nil
}

func (m *ImageSummary) validateParentID(formats strfmt.Registry) error {

	if err := validate.RequiredString("ParentId", "body", string(m.ParentID)); err != nil {
		return err
	}

	return nil
}

func (m *ImageSummary) validateRepoDigests(formats strfmt.Registry) error {

	if err := validate.Required("RepoDigests", "body", m.RepoDigests); err != nil {
		return err
	}

	return nil
}

func (m *ImageSummary) validateRepoTags(formats strfmt.Registry) error {

	if err := validate.Required("RepoTags", "body", m.RepoTags); err != nil {
		return err
	}

	return nil
}

func (m *ImageSummary) validateSharedSize(formats strfmt.Registry) error {

	if err := validate.Required("SharedSize", "body", int64(m.SharedSize)); err != nil {
		return err
	}

	return nil
}

func (m *ImageSummary) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("Size", "body", int64(m.Size)); err != nil {
		return err
	}

	return nil
}

func (m *ImageSummary) validateVirtualSize(formats strfmt.Registry) error {

	if err := validate.Required("VirtualSize", "body", int64(m.VirtualSize)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ImageSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageSummary) UnmarshalBinary(b []byte) error {
	var res ImageSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
