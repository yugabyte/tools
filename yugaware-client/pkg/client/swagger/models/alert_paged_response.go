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

// AlertPagedResponse alert paged response
//
// swagger:model AlertPagedResponse
type AlertPagedResponse struct {

	// entities
	// Required: true
	Entities []*Alert `json:"entities"`

	// has next
	// Required: true
	HasNext *bool `json:"hasNext"`

	// has prev
	// Required: true
	HasPrev *bool `json:"hasPrev"`

	// total count
	// Required: true
	TotalCount *int32 `json:"totalCount"`
}

// Validate validates this alert paged response
func (m *AlertPagedResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHasNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHasPrev(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertPagedResponse) validateEntities(formats strfmt.Registry) error {

	if err := validate.Required("entities", "body", m.Entities); err != nil {
		return err
	}

	for i := 0; i < len(m.Entities); i++ {
		if swag.IsZero(m.Entities[i]) { // not required
			continue
		}

		if m.Entities[i] != nil {
			if err := m.Entities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("entities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AlertPagedResponse) validateHasNext(formats strfmt.Registry) error {

	if err := validate.Required("hasNext", "body", m.HasNext); err != nil {
		return err
	}

	return nil
}

func (m *AlertPagedResponse) validateHasPrev(formats strfmt.Registry) error {

	if err := validate.Required("hasPrev", "body", m.HasPrev); err != nil {
		return err
	}

	return nil
}

func (m *AlertPagedResponse) validateTotalCount(formats strfmt.Registry) error {

	if err := validate.Required("totalCount", "body", m.TotalCount); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this alert paged response based on the context it is used
func (m *AlertPagedResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEntities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertPagedResponse) contextValidateEntities(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Entities); i++ {

		if m.Entities[i] != nil {
			if err := m.Entities[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("entities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AlertPagedResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlertPagedResponse) UnmarshalBinary(b []byte) error {
	var res AlertPagedResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
