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

// PVMInstanceAction p VM instance action
// swagger:model PVMInstanceAction
type PVMInstanceAction struct {

	// Name of the action to take; can be start, stop, hard-reboot, soft-reboot, immediate-shutdown, reset-state
	// Required: true
	// Enum: [start stop immediate-shutdown hard-reboot soft-reboot reset-state]
	Action *string `json:"action"`
}

// Validate validates this p VM instance action
func (m *PVMInstanceAction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var pVmInstanceActionTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["start","stop","immediate-shutdown","hard-reboot","soft-reboot","reset-state"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pVmInstanceActionTypeActionPropEnum = append(pVmInstanceActionTypeActionPropEnum, v)
	}
}

const (

	// PVMInstanceActionActionStart captures enum value "start"
	PVMInstanceActionActionStart string = "start"

	// PVMInstanceActionActionStop captures enum value "stop"
	PVMInstanceActionActionStop string = "stop"

	// PVMInstanceActionActionImmediateShutdown captures enum value "immediate-shutdown"
	PVMInstanceActionActionImmediateShutdown string = "immediate-shutdown"

	// PVMInstanceActionActionHardReboot captures enum value "hard-reboot"
	PVMInstanceActionActionHardReboot string = "hard-reboot"

	// PVMInstanceActionActionSoftReboot captures enum value "soft-reboot"
	PVMInstanceActionActionSoftReboot string = "soft-reboot"

	// PVMInstanceActionActionResetState captures enum value "reset-state"
	PVMInstanceActionActionResetState string = "reset-state"
)

// prop value enum
func (m *PVMInstanceAction) validateActionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, pVmInstanceActionTypeActionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PVMInstanceAction) validateAction(formats strfmt.Registry) error {

	if err := validate.Required("action", "body", m.Action); err != nil {
		return err
	}

	// value enum
	if err := m.validateActionEnum("action", "body", *m.Action); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PVMInstanceAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PVMInstanceAction) UnmarshalBinary(b []byte) error {
	var res PVMInstanceAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}