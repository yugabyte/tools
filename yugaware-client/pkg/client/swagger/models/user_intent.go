// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UserIntent user intent
//
// swagger:model UserIntent
type UserIntent struct {

	// access key code
	AccessKeyCode string `json:"accessKeyCode,omitempty"`

	// assign public IP
	AssignPublicIP bool `json:"assignPublicIP,omitempty"`

	// Whether to assign static public IP
	AssignStaticPublicIP bool `json:"assignStaticPublicIP,omitempty"`

	// aws arn string
	AwsArnString string `json:"awsArnString,omitempty"`

	// device info
	DeviceInfo *DeviceInfo `json:"deviceInfo,omitempty"`

	// enable client to node encrypt
	EnableClientToNodeEncrypt bool `json:"enableClientToNodeEncrypt,omitempty"`

	// enable exposing service
	// Enum: [NONE EXPOSED UNEXPOSED]
	EnableExposingService string `json:"enableExposingService,omitempty"`

	// enable IP v6
	EnableIPV6 bool `json:"enableIPV6,omitempty"`

	// enable node to node encrypt
	EnableNodeToNodeEncrypt bool `json:"enableNodeToNodeEncrypt,omitempty"`

	// enable volume encryption
	EnableVolumeEncryption bool `json:"enableVolumeEncryption,omitempty"`

	// enable y e d i s
	EnableYEDIS bool `json:"enableYEDIS,omitempty"`

	// enable y SQL
	EnableYSQL bool `json:"enableYSQL,omitempty"`

	// instance tags
	InstanceTags map[string]string `json:"instanceTags,omitempty"`

	// instance type
	InstanceType string `json:"instanceType,omitempty"`

	// master g flags
	MasterGFlags map[string]string `json:"masterGFlags,omitempty"`

	// num nodes
	NumNodes int32 `json:"numNodes,omitempty"`

	// preferred region
	// Format: uuid
	PreferredRegion strfmt.UUID `json:"preferredRegion,omitempty"`

	// provider
	Provider string `json:"provider,omitempty"`

	// provider type
	// Enum: [unknown aws gcp azu docker onprem kubernetes local other]
	ProviderType string `json:"providerType,omitempty"`

	// region list
	RegionList []strfmt.UUID `json:"regionList"`

	// replication factor
	ReplicationFactor int32 `json:"replicationFactor,omitempty"`

	// tserver g flags
	TserverGFlags map[string]string `json:"tserverGFlags,omitempty"`

	// universe name
	UniverseName string `json:"universeName,omitempty"`

	// use hostname
	UseHostname bool `json:"useHostname,omitempty"`

	// use systemd
	UseSystemd bool `json:"useSystemd,omitempty"`

	// use time sync
	UseTimeSync bool `json:"useTimeSync,omitempty"`

	// yb software version
	YbSoftwareVersion string `json:"ybSoftwareVersion,omitempty"`
}

// Validate validates this user intent
func (m *UserIntent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeviceInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnableExposingService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreferredRegion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProviderType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegionList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserIntent) validateDeviceInfo(formats strfmt.Registry) error {
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

var userIntentTypeEnableExposingServicePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NONE","EXPOSED","UNEXPOSED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userIntentTypeEnableExposingServicePropEnum = append(userIntentTypeEnableExposingServicePropEnum, v)
	}
}

const (

	// UserIntentEnableExposingServiceNONE captures enum value "NONE"
	UserIntentEnableExposingServiceNONE string = "NONE"

	// UserIntentEnableExposingServiceEXPOSED captures enum value "EXPOSED"
	UserIntentEnableExposingServiceEXPOSED string = "EXPOSED"

	// UserIntentEnableExposingServiceUNEXPOSED captures enum value "UNEXPOSED"
	UserIntentEnableExposingServiceUNEXPOSED string = "UNEXPOSED"
)

// prop value enum
func (m *UserIntent) validateEnableExposingServiceEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, userIntentTypeEnableExposingServicePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UserIntent) validateEnableExposingService(formats strfmt.Registry) error {
	if swag.IsZero(m.EnableExposingService) { // not required
		return nil
	}

	// value enum
	if err := m.validateEnableExposingServiceEnum("enableExposingService", "body", m.EnableExposingService); err != nil {
		return err
	}

	return nil
}

func (m *UserIntent) validatePreferredRegion(formats strfmt.Registry) error {
	if swag.IsZero(m.PreferredRegion) { // not required
		return nil
	}

	if err := validate.FormatOf("preferredRegion", "body", "uuid", m.PreferredRegion.String(), formats); err != nil {
		return err
	}

	return nil
}

var userIntentTypeProviderTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["unknown","aws","gcp","azu","docker","onprem","kubernetes","local","other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userIntentTypeProviderTypePropEnum = append(userIntentTypeProviderTypePropEnum, v)
	}
}

const (

	// UserIntentProviderTypeUnknown captures enum value "unknown"
	UserIntentProviderTypeUnknown string = "unknown"

	// UserIntentProviderTypeAws captures enum value "aws"
	UserIntentProviderTypeAws string = "aws"

	// UserIntentProviderTypeGcp captures enum value "gcp"
	UserIntentProviderTypeGcp string = "gcp"

	// UserIntentProviderTypeAzu captures enum value "azu"
	UserIntentProviderTypeAzu string = "azu"

	// UserIntentProviderTypeDocker captures enum value "docker"
	UserIntentProviderTypeDocker string = "docker"

	// UserIntentProviderTypeOnprem captures enum value "onprem"
	UserIntentProviderTypeOnprem string = "onprem"

	// UserIntentProviderTypeKubernetes captures enum value "kubernetes"
	UserIntentProviderTypeKubernetes string = "kubernetes"

	// UserIntentProviderTypeLocal captures enum value "local"
	UserIntentProviderTypeLocal string = "local"

	// UserIntentProviderTypeOther captures enum value "other"
	UserIntentProviderTypeOther string = "other"
)

// prop value enum
func (m *UserIntent) validateProviderTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, userIntentTypeProviderTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UserIntent) validateProviderType(formats strfmt.Registry) error {
	if swag.IsZero(m.ProviderType) { // not required
		return nil
	}

	// value enum
	if err := m.validateProviderTypeEnum("providerType", "body", m.ProviderType); err != nil {
		return err
	}

	return nil
}

func (m *UserIntent) validateRegionList(formats strfmt.Registry) error {
	if swag.IsZero(m.RegionList) { // not required
		return nil
	}

	for i := 0; i < len(m.RegionList); i++ {

		if err := validate.FormatOf("regionList"+"."+strconv.Itoa(i), "body", "uuid", m.RegionList[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

// ContextValidate validate this user intent based on the context it is used
func (m *UserIntent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDeviceInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserIntent) contextValidateDeviceInfo(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *UserIntent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserIntent) UnmarshalBinary(b []byte) error {
	var res UserIntent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
