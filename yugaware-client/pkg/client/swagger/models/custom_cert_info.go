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

// CustomCertInfo custom cert info
//
// swagger:model CustomCertInfo
type CustomCertInfo struct {

	// client cert path
	// Required: true
	ClientCertPath *string `json:"clientCertPath"`

	// client key path
	// Required: true
	ClientKeyPath *string `json:"clientKeyPath"`

	// node cert path
	// Required: true
	NodeCertPath *string `json:"nodeCertPath"`

	// node key path
	// Required: true
	NodeKeyPath *string `json:"nodeKeyPath"`

	// root cert path
	// Required: true
	RootCertPath *string `json:"rootCertPath"`
}

// Validate validates this custom cert info
func (m *CustomCertInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientCertPath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientKeyPath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeCertPath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeKeyPath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRootCertPath(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomCertInfo) validateClientCertPath(formats strfmt.Registry) error {

	if err := validate.Required("clientCertPath", "body", m.ClientCertPath); err != nil {
		return err
	}

	return nil
}

func (m *CustomCertInfo) validateClientKeyPath(formats strfmt.Registry) error {

	if err := validate.Required("clientKeyPath", "body", m.ClientKeyPath); err != nil {
		return err
	}

	return nil
}

func (m *CustomCertInfo) validateNodeCertPath(formats strfmt.Registry) error {

	if err := validate.Required("nodeCertPath", "body", m.NodeCertPath); err != nil {
		return err
	}

	return nil
}

func (m *CustomCertInfo) validateNodeKeyPath(formats strfmt.Registry) error {

	if err := validate.Required("nodeKeyPath", "body", m.NodeKeyPath); err != nil {
		return err
	}

	return nil
}

func (m *CustomCertInfo) validateRootCertPath(formats strfmt.Registry) error {

	if err := validate.Required("rootCertPath", "body", m.RootCertPath); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this custom cert info based on context it is used
func (m *CustomCertInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustomCertInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomCertInfo) UnmarshalBinary(b []byte) error {
	var res CustomCertInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
