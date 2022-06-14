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

// NodeDetailsResp node details resp
//
// swagger:model NodeDetailsResp
type NodeDetailsResp struct {

	// allowed actions
	// Read Only: true
	// Unique: true
	AllowedActions []string `json:"allowedActions"`

	// The availability zone's UUID
	// Format: uuid
	AzUUID strfmt.UUID `json:"azUuid,omitempty"`

	// Node information, as reported by the cloud provider
	CloudInfo *CloudSpecificInfo `json:"cloudInfo,omitempty"`

	// True if cron jobs were properly configured for this node
	CronsActive bool `json:"cronsActive,omitempty"`

	// Disks are mounted by uuid
	DisksAreMountedByUUID bool `json:"disksAreMountedByUUID,omitempty"`

	// True if this node is a master
	IsMaster bool `json:"isMaster,omitempty"`

	// True if this node is a REDIS server
	IsRedisServer bool `json:"isRedisServer,omitempty"`

	// True if this node is a Tablet server
	IsTserver bool `json:"isTserver,omitempty"`

	// True if this node is a YCQL server
	IsYqlServer bool `json:"isYqlServer,omitempty"`

	// True if this node is a YSQL server
	IsYsqlServer bool `json:"isYsqlServer,omitempty"`

	// Machine image name
	MachineImage string `json:"machineImage,omitempty"`

	// Master HTTP port
	MasterHTTPPort int32 `json:"masterHttpPort,omitempty"`

	// Master RCP port
	MasterRPCPort int32 `json:"masterRpcPort,omitempty"`

	// Master state
	// Example: ToStart
	// Enum: [None ToStart Configured ToStop]
	MasterState string `json:"masterState,omitempty"`

	// Node exporter port
	NodeExporterPort int32 `json:"nodeExporterPort,omitempty"`

	// Node ID
	NodeIdx int32 `json:"nodeIdx,omitempty"`

	// Node name
	NodeName string `json:"nodeName,omitempty"`

	// Node UUID
	// Format: uuid
	NodeUUID strfmt.UUID `json:"nodeUuid,omitempty"`

	// UUID of the cluster to which this node belongs
	// Format: uuid
	PlacementUUID strfmt.UUID `json:"placementUuid,omitempty"`

	// REDIS HTTP port
	RedisServerHTTPPort int32 `json:"redisServerHttpPort,omitempty"`

	// REDIS RPC port
	RedisServerRPCPort int32 `json:"redisServerRpcPort,omitempty"`

	// Node state
	// Example: Provisioned
	// Enum: [ToBeAdded InstanceCreated ServerSetup ToJoinCluster Reprovisioning Provisioned SoftwareInstalled UpgradeSoftware UpdateGFlags Live Stopping Starting Stopped Unreachable ToBeRemoved Removing Removed Adding BeingDecommissioned Decommissioned UpdateCert ToggleTls Resizing SystemdUpgrade Terminating Terminated]
	State string `json:"state,omitempty"`

	// Tablet server HTTP port
	TserverHTTPPort int32 `json:"tserverHttpPort,omitempty"`

	// Tablet server RPC port
	TserverRPCPort int32 `json:"tserverRpcPort,omitempty"`

	// True if this a custom YB AMI
	YbPrebuiltAmi bool `json:"ybPrebuiltAmi,omitempty"`

	// YCQL HTTP port
	YqlServerHTTPPort int32 `json:"yqlServerHttpPort,omitempty"`

	// YCQL RPC port
	YqlServerRPCPort int32 `json:"yqlServerRpcPort,omitempty"`

	// YSQL HTTP port
	YsqlServerHTTPPort int32 `json:"ysqlServerHttpPort,omitempty"`

	// YSQL RPC port
	YsqlServerRPCPort int32 `json:"ysqlServerRpcPort,omitempty"`
}

// Validate validates this node details resp
func (m *NodeDetailsResp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowedActions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzUUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMasterState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeUUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlacementUUID(formats); err != nil {
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

var nodeDetailsRespAllowedActionsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ADD","REMOVE","START","STOP","DELETE","QUERY","RELEASE","START_MASTER","PRECHECK_DETACHED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nodeDetailsRespAllowedActionsItemsEnum = append(nodeDetailsRespAllowedActionsItemsEnum, v)
	}
}

func (m *NodeDetailsResp) validateAllowedActionsItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, nodeDetailsRespAllowedActionsItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NodeDetailsResp) validateAllowedActions(formats strfmt.Registry) error {
	if swag.IsZero(m.AllowedActions) { // not required
		return nil
	}

	if err := validate.UniqueItems("allowedActions", "body", m.AllowedActions); err != nil {
		return err
	}

	for i := 0; i < len(m.AllowedActions); i++ {

		// value enum
		if err := m.validateAllowedActionsItemsEnum("allowedActions"+"."+strconv.Itoa(i), "body", m.AllowedActions[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *NodeDetailsResp) validateAzUUID(formats strfmt.Registry) error {
	if swag.IsZero(m.AzUUID) { // not required
		return nil
	}

	if err := validate.FormatOf("azUuid", "body", "uuid", m.AzUUID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NodeDetailsResp) validateCloudInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.CloudInfo) { // not required
		return nil
	}

	if m.CloudInfo != nil {
		if err := m.CloudInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloudInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cloudInfo")
			}
			return err
		}
	}

	return nil
}

var nodeDetailsRespTypeMasterStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["None","ToStart","Configured","ToStop"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nodeDetailsRespTypeMasterStatePropEnum = append(nodeDetailsRespTypeMasterStatePropEnum, v)
	}
}

const (

	// NodeDetailsRespMasterStateNone captures enum value "None"
	NodeDetailsRespMasterStateNone string = "None"

	// NodeDetailsRespMasterStateToStart captures enum value "ToStart"
	NodeDetailsRespMasterStateToStart string = "ToStart"

	// NodeDetailsRespMasterStateConfigured captures enum value "Configured"
	NodeDetailsRespMasterStateConfigured string = "Configured"

	// NodeDetailsRespMasterStateToStop captures enum value "ToStop"
	NodeDetailsRespMasterStateToStop string = "ToStop"
)

// prop value enum
func (m *NodeDetailsResp) validateMasterStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, nodeDetailsRespTypeMasterStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NodeDetailsResp) validateMasterState(formats strfmt.Registry) error {
	if swag.IsZero(m.MasterState) { // not required
		return nil
	}

	// value enum
	if err := m.validateMasterStateEnum("masterState", "body", m.MasterState); err != nil {
		return err
	}

	return nil
}

func (m *NodeDetailsResp) validateNodeUUID(formats strfmt.Registry) error {
	if swag.IsZero(m.NodeUUID) { // not required
		return nil
	}

	if err := validate.FormatOf("nodeUuid", "body", "uuid", m.NodeUUID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NodeDetailsResp) validatePlacementUUID(formats strfmt.Registry) error {
	if swag.IsZero(m.PlacementUUID) { // not required
		return nil
	}

	if err := validate.FormatOf("placementUuid", "body", "uuid", m.PlacementUUID.String(), formats); err != nil {
		return err
	}

	return nil
}

var nodeDetailsRespTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ToBeAdded","InstanceCreated","ServerSetup","ToJoinCluster","Reprovisioning","Provisioned","SoftwareInstalled","UpgradeSoftware","UpdateGFlags","Live","Stopping","Starting","Stopped","Unreachable","ToBeRemoved","Removing","Removed","Adding","BeingDecommissioned","Decommissioned","UpdateCert","ToggleTls","Resizing","SystemdUpgrade","Terminating","Terminated"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nodeDetailsRespTypeStatePropEnum = append(nodeDetailsRespTypeStatePropEnum, v)
	}
}

const (

	// NodeDetailsRespStateToBeAdded captures enum value "ToBeAdded"
	NodeDetailsRespStateToBeAdded string = "ToBeAdded"

	// NodeDetailsRespStateInstanceCreated captures enum value "InstanceCreated"
	NodeDetailsRespStateInstanceCreated string = "InstanceCreated"

	// NodeDetailsRespStateServerSetup captures enum value "ServerSetup"
	NodeDetailsRespStateServerSetup string = "ServerSetup"

	// NodeDetailsRespStateToJoinCluster captures enum value "ToJoinCluster"
	NodeDetailsRespStateToJoinCluster string = "ToJoinCluster"

	// NodeDetailsRespStateReprovisioning captures enum value "Reprovisioning"
	NodeDetailsRespStateReprovisioning string = "Reprovisioning"

	// NodeDetailsRespStateProvisioned captures enum value "Provisioned"
	NodeDetailsRespStateProvisioned string = "Provisioned"

	// NodeDetailsRespStateSoftwareInstalled captures enum value "SoftwareInstalled"
	NodeDetailsRespStateSoftwareInstalled string = "SoftwareInstalled"

	// NodeDetailsRespStateUpgradeSoftware captures enum value "UpgradeSoftware"
	NodeDetailsRespStateUpgradeSoftware string = "UpgradeSoftware"

	// NodeDetailsRespStateUpdateGFlags captures enum value "UpdateGFlags"
	NodeDetailsRespStateUpdateGFlags string = "UpdateGFlags"

	// NodeDetailsRespStateLive captures enum value "Live"
	NodeDetailsRespStateLive string = "Live"

	// NodeDetailsRespStateStopping captures enum value "Stopping"
	NodeDetailsRespStateStopping string = "Stopping"

	// NodeDetailsRespStateStarting captures enum value "Starting"
	NodeDetailsRespStateStarting string = "Starting"

	// NodeDetailsRespStateStopped captures enum value "Stopped"
	NodeDetailsRespStateStopped string = "Stopped"

	// NodeDetailsRespStateUnreachable captures enum value "Unreachable"
	NodeDetailsRespStateUnreachable string = "Unreachable"

	// NodeDetailsRespStateToBeRemoved captures enum value "ToBeRemoved"
	NodeDetailsRespStateToBeRemoved string = "ToBeRemoved"

	// NodeDetailsRespStateRemoving captures enum value "Removing"
	NodeDetailsRespStateRemoving string = "Removing"

	// NodeDetailsRespStateRemoved captures enum value "Removed"
	NodeDetailsRespStateRemoved string = "Removed"

	// NodeDetailsRespStateAdding captures enum value "Adding"
	NodeDetailsRespStateAdding string = "Adding"

	// NodeDetailsRespStateBeingDecommissioned captures enum value "BeingDecommissioned"
	NodeDetailsRespStateBeingDecommissioned string = "BeingDecommissioned"

	// NodeDetailsRespStateDecommissioned captures enum value "Decommissioned"
	NodeDetailsRespStateDecommissioned string = "Decommissioned"

	// NodeDetailsRespStateUpdateCert captures enum value "UpdateCert"
	NodeDetailsRespStateUpdateCert string = "UpdateCert"

	// NodeDetailsRespStateToggleTLS captures enum value "ToggleTls"
	NodeDetailsRespStateToggleTLS string = "ToggleTls"

	// NodeDetailsRespStateResizing captures enum value "Resizing"
	NodeDetailsRespStateResizing string = "Resizing"

	// NodeDetailsRespStateSystemdUpgrade captures enum value "SystemdUpgrade"
	NodeDetailsRespStateSystemdUpgrade string = "SystemdUpgrade"

	// NodeDetailsRespStateTerminating captures enum value "Terminating"
	NodeDetailsRespStateTerminating string = "Terminating"

	// NodeDetailsRespStateTerminated captures enum value "Terminated"
	NodeDetailsRespStateTerminated string = "Terminated"
)

// prop value enum
func (m *NodeDetailsResp) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, nodeDetailsRespTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NodeDetailsResp) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this node details resp based on the context it is used
func (m *NodeDetailsResp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAllowedActions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCloudInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodeDetailsResp) contextValidateAllowedActions(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "allowedActions", "body", []string(m.AllowedActions)); err != nil {
		return err
	}

	return nil
}

func (m *NodeDetailsResp) contextValidateCloudInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.CloudInfo != nil {
		if err := m.CloudInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloudInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cloudInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NodeDetailsResp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeDetailsResp) UnmarshalBinary(b []byte) error {
	var res NodeDetailsResp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
