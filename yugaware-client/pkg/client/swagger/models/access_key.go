// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AccessKey Access key for the cloud provider. This helps to authenticate the user and get access to the provider.
//
// swagger:model AccessKey
type AccessKey struct {

	// id key
	// Required: true
	IDKey *AccessKeyID `json:"idKey"`

	// Cloud provider key information
	// Required: true
	KeyInfo *KeyInfo `json:"keyInfo"`
}

// Validate validates this access key
func (m *AccessKey) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIDKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeyInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccessKey) validateIDKey(formats strfmt.Registry) error {

	if err := validate.Required("idKey", "body", m.IDKey); err != nil {
		return err
	}

	if m.IDKey != nil {
		if err := m.IDKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("idKey")
			}
			return err
		}
	}

	return nil
}

func (m *AccessKey) validateKeyInfo(formats strfmt.Registry) error {

	if err := validate.Required("keyInfo", "body", m.KeyInfo); err != nil {
		return err
	}

	if m.KeyInfo != nil {
		if err := m.KeyInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("keyInfo")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this access key based on the context it is used
func (m *AccessKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIDKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKeyInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccessKey) contextValidateIDKey(ctx context.Context, formats strfmt.Registry) error {

	if m.IDKey != nil {
		if err := m.IDKey.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("idKey")
			}
			return err
		}
	}

	return nil
}

func (m *AccessKey) contextValidateKeyInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.KeyInfo != nil {
		if err := m.KeyInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("keyInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccessKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccessKey) UnmarshalBinary(b []byte) error {
	var res AccessKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
