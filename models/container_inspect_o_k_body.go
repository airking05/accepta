// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ContainerInspectOKBody container inspect o k body
// swagger:model containerInspectOKBody
type ContainerInspectOKBody struct {

	// app armor profile
	AppArmorProfile string `json:"AppArmorProfile,omitempty"`

	// The arguments to the command being run
	Args []string `json:"Args"`

	// config
	Config *ContainerConfig `json:"Config,omitempty"`

	// The time the container was created
	Created string `json:"Created,omitempty"`

	// driver
	Driver string `json:"Driver,omitempty"`

	// exec ids
	ExecIds string `json:"ExecIDs,omitempty"`

	// graph driver
	GraphDriver *GraphDriverData `json:"GraphDriver,omitempty"`

	// host config
	HostConfig *HostConfig `json:"HostConfig,omitempty"`

	// hostname path
	HostnamePath string `json:"HostnamePath,omitempty"`

	// hosts path
	HostsPath string `json:"HostsPath,omitempty"`

	// The ID of the container
	ID string `json:"Id,omitempty"`

	// The container's image
	Image string `json:"Image,omitempty"`

	// log path
	LogPath string `json:"LogPath,omitempty"`

	// mount label
	MountLabel string `json:"MountLabel,omitempty"`

	// mounts
	Mounts ContainerInspectOKBodyMounts `json:"Mounts"`

	// name
	Name string `json:"Name,omitempty"`

	// network settings
	NetworkSettings *NetworkSettings `json:"NetworkSettings,omitempty"`

	// TODO
	Node interface{} `json:"Node,omitempty"`

	// The path to the command being run
	Path string `json:"Path,omitempty"`

	// process label
	ProcessLabel string `json:"ProcessLabel,omitempty"`

	// resolv conf path
	ResolvConfPath string `json:"ResolvConfPath,omitempty"`

	// restart count
	RestartCount int64 `json:"RestartCount,omitempty"`

	// The total size of all the files in this container.
	SizeRootFs int64 `json:"SizeRootFs,omitempty"`

	// The size of files that have been created or changed by this container.
	SizeRw int64 `json:"SizeRw,omitempty"`

	// state
	State *ContainerInspectOKBodyState `json:"State,omitempty"`
}

// Validate validates this container inspect o k body
func (m *ContainerInspectOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArgs(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateConfig(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGraphDriver(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNetworkSettings(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContainerInspectOKBody) validateArgs(formats strfmt.Registry) error {

	if swag.IsZero(m.Args) { // not required
		return nil
	}

	return nil
}

func (m *ContainerInspectOKBody) validateConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.Config) { // not required
		return nil
	}

	if m.Config != nil {

		if err := m.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Config")
			}
			return err
		}

	}

	return nil
}

func (m *ContainerInspectOKBody) validateGraphDriver(formats strfmt.Registry) error {

	if swag.IsZero(m.GraphDriver) { // not required
		return nil
	}

	if m.GraphDriver != nil {

		if err := m.GraphDriver.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("GraphDriver")
			}
			return err
		}

	}

	return nil
}

func (m *ContainerInspectOKBody) validateNetworkSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.NetworkSettings) { // not required
		return nil
	}

	if m.NetworkSettings != nil {

		if err := m.NetworkSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("NetworkSettings")
			}
			return err
		}

	}

	return nil
}

func (m *ContainerInspectOKBody) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	if m.State != nil {

		if err := m.State.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("State")
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContainerInspectOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContainerInspectOKBody) UnmarshalBinary(b []byte) error {
	var res ContainerInspectOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
