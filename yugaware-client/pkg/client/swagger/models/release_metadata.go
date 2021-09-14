// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ReleaseMetadata Yugabyte release metadata
//
// swagger:model ReleaseMetadata
type ReleaseMetadata struct {

	// Helm chart path
	ChartPath string `json:"chartPath,omitempty"`

	// Release file path
	FilePath string `json:"filePath,omitempty"`

	// GCS link and credentials
	Gcs *GCSLocation `json:"gcs,omitempty"`

	// HTTP link to the release
	HTTP *HTTPLocation `json:"http,omitempty"`

	// Release image tag
	ImageTag string `json:"imageTag,omitempty"`

	// Release notes
	Notes []string `json:"notes"`

	// S3 link and credentials
	S3 *S3Location `json:"s3,omitempty"`

	// Release state
	// Example: ACTIVE
	// Enum: [ACTIVE DISABLED DELETED]
	State string `json:"state,omitempty"`
}

// Validate validates this release metadata
func (m *ReleaseMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGcs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateS3(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReleaseMetadata) validateGcs(formats strfmt.Registry) error {
	if swag.IsZero(m.Gcs) { // not required
		return nil
	}

	if m.Gcs != nil {
		if err := m.Gcs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcs")
			}
			return err
		}
	}

	return nil
}

func (m *ReleaseMetadata) validateHTTP(formats strfmt.Registry) error {
	if swag.IsZero(m.HTTP) { // not required
		return nil
	}

	if m.HTTP != nil {
		if err := m.HTTP.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("http")
			}
			return err
		}
	}

	return nil
}

func (m *ReleaseMetadata) validateS3(formats strfmt.Registry) error {
	if swag.IsZero(m.S3) { // not required
		return nil
	}

	if m.S3 != nil {
		if err := m.S3.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("s3")
			}
			return err
		}
	}

	return nil
}

var releaseMetadataTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACTIVE","DISABLED","DELETED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		releaseMetadataTypeStatePropEnum = append(releaseMetadataTypeStatePropEnum, v)
	}
}

const (

	// ReleaseMetadataStateACTIVE captures enum value "ACTIVE"
	ReleaseMetadataStateACTIVE string = "ACTIVE"

	// ReleaseMetadataStateDISABLED captures enum value "DISABLED"
	ReleaseMetadataStateDISABLED string = "DISABLED"

	// ReleaseMetadataStateDELETED captures enum value "DELETED"
	ReleaseMetadataStateDELETED string = "DELETED"
)

// prop value enum
func (m *ReleaseMetadata) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, releaseMetadataTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ReleaseMetadata) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this release metadata based on the context it is used
func (m *ReleaseMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGcs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHTTP(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateS3(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReleaseMetadata) contextValidateGcs(ctx context.Context, formats strfmt.Registry) error {

	if m.Gcs != nil {
		if err := m.Gcs.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcs")
			}
			return err
		}
	}

	return nil
}

func (m *ReleaseMetadata) contextValidateHTTP(ctx context.Context, formats strfmt.Registry) error {

	if m.HTTP != nil {
		if err := m.HTTP.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("http")
			}
			return err
		}
	}

	return nil
}

func (m *ReleaseMetadata) contextValidateS3(ctx context.Context, formats strfmt.Registry) error {

	if m.S3 != nil {
		if err := m.S3.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("s3")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReleaseMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReleaseMetadata) UnmarshalBinary(b []byte) error {
	var res ReleaseMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
