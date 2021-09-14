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

// AlertDestination alert destination
//
// swagger:model AlertDestination
type AlertDestination struct {

	// channels
	// Read Only: true
	Channels []strfmt.UUID `json:"channels"`

	// customer UUID
	// Required: true
	// Format: uuid
	CustomerUUID *strfmt.UUID `json:"customerUUID"`

	// default destination
	// Required: true
	DefaultDestination *bool `json:"defaultDestination"`

	// name
	// Required: true
	Name *string `json:"name"`

	// uuid
	// Required: true
	// Format: uuid
	UUID *strfmt.UUID `json:"uuid"`
}

// Validate validates this alert destination
func (m *AlertDestination) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChannels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomerUUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultDestination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertDestination) validateChannels(formats strfmt.Registry) error {
	if swag.IsZero(m.Channels) { // not required
		return nil
	}

	for i := 0; i < len(m.Channels); i++ {

		if err := validate.FormatOf("channels"+"."+strconv.Itoa(i), "body", "uuid", m.Channels[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *AlertDestination) validateCustomerUUID(formats strfmt.Registry) error {

	if err := validate.Required("customerUUID", "body", m.CustomerUUID); err != nil {
		return err
	}

	if err := validate.FormatOf("customerUUID", "body", "uuid", m.CustomerUUID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AlertDestination) validateDefaultDestination(formats strfmt.Registry) error {

	if err := validate.Required("defaultDestination", "body", m.DefaultDestination); err != nil {
		return err
	}

	return nil
}

func (m *AlertDestination) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *AlertDestination) validateUUID(formats strfmt.Registry) error {

	if err := validate.Required("uuid", "body", m.UUID); err != nil {
		return err
	}

	if err := validate.FormatOf("uuid", "body", "uuid", m.UUID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this alert destination based on the context it is used
func (m *AlertDestination) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChannels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertDestination) contextValidateChannels(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "channels", "body", []strfmt.UUID(m.Channels)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AlertDestination) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlertDestination) UnmarshalBinary(b []byte) error {
	var res AlertDestination
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
