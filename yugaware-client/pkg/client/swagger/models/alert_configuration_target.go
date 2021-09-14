// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AlertConfigurationTarget Alert target. Set to `all`, or specify one or more target UUIDs.
//
// swagger:model AlertConfigurationTarget
type AlertConfigurationTarget struct {

	// Alert applicable to all targets
	All bool `json:"all,omitempty"`

	// Alert target UUIDs
	// Unique: true
	Uuids []strfmt.UUID `json:"uuids"`
}

// Validate validates this alert configuration target
func (m *AlertConfigurationTarget) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUuids(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertConfigurationTarget) validateUuids(formats strfmt.Registry) error {
	if swag.IsZero(m.Uuids) { // not required
		return nil
	}

	if err := validate.UniqueItems("uuids", "body", m.Uuids); err != nil {
		return err
	}

	for i := 0; i < len(m.Uuids); i++ {

		if err := validate.FormatOf("uuids"+"."+strconv.Itoa(i), "body", "uuid", m.Uuids[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

// ContextValidate validates this alert configuration target based on context it is used
func (m *AlertConfigurationTarget) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AlertConfigurationTarget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlertConfigurationTarget) UnmarshalBinary(b []byte) error {
	var res AlertConfigurationTarget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
