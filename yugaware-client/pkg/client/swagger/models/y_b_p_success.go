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

// YBPSuccess y b p success
//
// swagger:model YBPSuccess
type YBPSuccess struct {

	// API response message.
	// Read Only: true
	Message string `json:"message,omitempty"`

	// API operation status. A value of true indicates the operation was successful.
	// Read Only: true
	Success *bool `json:"success,omitempty"`
}

// Validate validates this y b p success
func (m *YBPSuccess) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this y b p success based on the context it is used
func (m *YBPSuccess) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSuccess(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *YBPSuccess) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "message", "body", string(m.Message)); err != nil {
		return err
	}

	return nil
}

func (m *YBPSuccess) contextValidateSuccess(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "success", "body", m.Success); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *YBPSuccess) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *YBPSuccess) UnmarshalBinary(b []byte) error {
	var res YBPSuccess
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
