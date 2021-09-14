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

// BulkImportParams Bulk import parameters
//
// swagger:model BulkImportParams
type BulkImportParams struct {

	// Amazon Resource Name (ARN) of the CMK
	CmkArn string `json:"cmkArn,omitempty"`

	// Communication ports
	CommunicationPorts *CommunicationPorts `json:"communicationPorts,omitempty"`

	// Device information
	DeviceInfo *DeviceInfo `json:"deviceInfo,omitempty"`

	// Encryption at rest configation
	EncryptionAtRestConfig *EncryptionAtRestConfig `json:"encryptionAtRestConfig,omitempty"`

	// Error message
	ErrorString string `json:"errorString,omitempty"`

	// Expected universe version
	ExpectedUniverseVersion int32 `json:"expectedUniverseVersion,omitempty"`

	// Extra dependencies
	ExtraDependencies *ExtraDependencies `json:"extraDependencies,omitempty"`

	// Whether this task has been tried before
	FirstTry bool `json:"firstTry,omitempty"`

	// Instance count
	InstanceCount int32 `json:"instanceCount,omitempty"`

	// Key space
	Keyspace string `json:"keyspace,omitempty"`

	// Node details
	// Unique: true
	NodeDetailsSet []*NodeDetails `json:"nodeDetailsSet"`

	// Node exporter user
	NodeExporterUser string `json:"nodeExporterUser,omitempty"`

	// S3 bucket URL
	// Required: true
	S3Bucket *string `json:"s3Bucket"`

	// The source universe's sync replication relationships
	// Read Only: true
	SourceAsyncReplicationRelationships []*AsyncReplicationConfig `json:"sourceAsyncReplicationRelationships"`

	// Is SSE
	Sse bool `json:"sse,omitempty"`

	// Table name
	TableName string `json:"tableName,omitempty"`

	// Table UUID
	// Format: uuid
	TableUUID strfmt.UUID `json:"tableUUID,omitempty"`

	// The target universe's async replication relationships
	// Read Only: true
	TargetAsyncReplicationRelationships []*AsyncReplicationConfig `json:"targetAsyncReplicationRelationships"`

	// Associated universe UUID
	// Format: uuid
	UniverseUUID strfmt.UUID `json:"universeUUID,omitempty"`

	// Previous software version
	YbPrevSoftwareVersion string `json:"ybPrevSoftwareVersion,omitempty"`
}

// Validate validates this bulk import params
func (m *BulkImportParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommunicationPorts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEncryptionAtRestConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtraDependencies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeDetailsSet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateS3Bucket(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceAsyncReplicationRelationships(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTableUUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetAsyncReplicationRelationships(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUniverseUUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BulkImportParams) validateCommunicationPorts(formats strfmt.Registry) error {
	if swag.IsZero(m.CommunicationPorts) { // not required
		return nil
	}

	if m.CommunicationPorts != nil {
		if err := m.CommunicationPorts.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("communicationPorts")
			}
			return err
		}
	}

	return nil
}

func (m *BulkImportParams) validateDeviceInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.DeviceInfo) { // not required
		return nil
	}

	if m.DeviceInfo != nil {
		if err := m.DeviceInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deviceInfo")
			}
			return err
		}
	}

	return nil
}

func (m *BulkImportParams) validateEncryptionAtRestConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.EncryptionAtRestConfig) { // not required
		return nil
	}

	if m.EncryptionAtRestConfig != nil {
		if err := m.EncryptionAtRestConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("encryptionAtRestConfig")
			}
			return err
		}
	}

	return nil
}

func (m *BulkImportParams) validateExtraDependencies(formats strfmt.Registry) error {
	if swag.IsZero(m.ExtraDependencies) { // not required
		return nil
	}

	if m.ExtraDependencies != nil {
		if err := m.ExtraDependencies.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("extraDependencies")
			}
			return err
		}
	}

	return nil
}

func (m *BulkImportParams) validateNodeDetailsSet(formats strfmt.Registry) error {
	if swag.IsZero(m.NodeDetailsSet) { // not required
		return nil
	}

	if err := validate.UniqueItems("nodeDetailsSet", "body", m.NodeDetailsSet); err != nil {
		return err
	}

	for i := 0; i < len(m.NodeDetailsSet); i++ {
		if swag.IsZero(m.NodeDetailsSet[i]) { // not required
			continue
		}

		if m.NodeDetailsSet[i] != nil {
			if err := m.NodeDetailsSet[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nodeDetailsSet" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BulkImportParams) validateS3Bucket(formats strfmt.Registry) error {

	if err := validate.Required("s3Bucket", "body", m.S3Bucket); err != nil {
		return err
	}

	return nil
}

func (m *BulkImportParams) validateSourceAsyncReplicationRelationships(formats strfmt.Registry) error {
	if swag.IsZero(m.SourceAsyncReplicationRelationships) { // not required
		return nil
	}

	for i := 0; i < len(m.SourceAsyncReplicationRelationships); i++ {
		if swag.IsZero(m.SourceAsyncReplicationRelationships[i]) { // not required
			continue
		}

		if m.SourceAsyncReplicationRelationships[i] != nil {
			if err := m.SourceAsyncReplicationRelationships[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sourceAsyncReplicationRelationships" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BulkImportParams) validateTableUUID(formats strfmt.Registry) error {
	if swag.IsZero(m.TableUUID) { // not required
		return nil
	}

	if err := validate.FormatOf("tableUUID", "body", "uuid", m.TableUUID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BulkImportParams) validateTargetAsyncReplicationRelationships(formats strfmt.Registry) error {
	if swag.IsZero(m.TargetAsyncReplicationRelationships) { // not required
		return nil
	}

	for i := 0; i < len(m.TargetAsyncReplicationRelationships); i++ {
		if swag.IsZero(m.TargetAsyncReplicationRelationships[i]) { // not required
			continue
		}

		if m.TargetAsyncReplicationRelationships[i] != nil {
			if err := m.TargetAsyncReplicationRelationships[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("targetAsyncReplicationRelationships" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BulkImportParams) validateUniverseUUID(formats strfmt.Registry) error {
	if swag.IsZero(m.UniverseUUID) { // not required
		return nil
	}

	if err := validate.FormatOf("universeUUID", "body", "uuid", m.UniverseUUID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this bulk import params based on the context it is used
func (m *BulkImportParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCommunicationPorts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeviceInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEncryptionAtRestConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExtraDependencies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNodeDetailsSet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSourceAsyncReplicationRelationships(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTargetAsyncReplicationRelationships(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BulkImportParams) contextValidateCommunicationPorts(ctx context.Context, formats strfmt.Registry) error {

	if m.CommunicationPorts != nil {
		if err := m.CommunicationPorts.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("communicationPorts")
			}
			return err
		}
	}

	return nil
}

func (m *BulkImportParams) contextValidateDeviceInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.DeviceInfo != nil {
		if err := m.DeviceInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deviceInfo")
			}
			return err
		}
	}

	return nil
}

func (m *BulkImportParams) contextValidateEncryptionAtRestConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.EncryptionAtRestConfig != nil {
		if err := m.EncryptionAtRestConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("encryptionAtRestConfig")
			}
			return err
		}
	}

	return nil
}

func (m *BulkImportParams) contextValidateExtraDependencies(ctx context.Context, formats strfmt.Registry) error {

	if m.ExtraDependencies != nil {
		if err := m.ExtraDependencies.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("extraDependencies")
			}
			return err
		}
	}

	return nil
}

func (m *BulkImportParams) contextValidateNodeDetailsSet(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NodeDetailsSet); i++ {

		if m.NodeDetailsSet[i] != nil {
			if err := m.NodeDetailsSet[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nodeDetailsSet" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BulkImportParams) contextValidateSourceAsyncReplicationRelationships(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "sourceAsyncReplicationRelationships", "body", []*AsyncReplicationConfig(m.SourceAsyncReplicationRelationships)); err != nil {
		return err
	}

	for i := 0; i < len(m.SourceAsyncReplicationRelationships); i++ {

		if m.SourceAsyncReplicationRelationships[i] != nil {
			if err := m.SourceAsyncReplicationRelationships[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sourceAsyncReplicationRelationships" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BulkImportParams) contextValidateTargetAsyncReplicationRelationships(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "targetAsyncReplicationRelationships", "body", []*AsyncReplicationConfig(m.TargetAsyncReplicationRelationships)); err != nil {
		return err
	}

	for i := 0; i < len(m.TargetAsyncReplicationRelationships); i++ {

		if m.TargetAsyncReplicationRelationships[i] != nil {
			if err := m.TargetAsyncReplicationRelationships[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("targetAsyncReplicationRelationships" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BulkImportParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BulkImportParams) UnmarshalBinary(b []byte) error {
	var res BulkImportParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
