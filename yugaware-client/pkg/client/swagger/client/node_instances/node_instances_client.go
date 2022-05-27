// Code generated by go-swagger; DO NOT EDIT.

package node_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new node instances API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for node instances API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateNodeInstance(params *CreateNodeInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateNodeInstanceOK, error)

	DeleteInstance(params *DeleteInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteInstanceOK, error)

	DetachedNodeAction(params *DetachedNodeActionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DetachedNodeActionOK, error)

	GetNodeInstance(params *GetNodeInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetNodeInstanceOK, error)

	ListByProvider(params *ListByProviderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListByProviderOK, error)

	ListByZone(params *ListByZoneParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListByZoneOK, error)

	NodeAction(params *NodeActionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*NodeActionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateNodeInstance creates a node instance
*/
func (a *Client) CreateNodeInstance(params *CreateNodeInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateNodeInstanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateNodeInstanceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createNodeInstance",
		Method:             "POST",
		PathPattern:        "/api/v1/customers/{cUUID}/zones/{azUUID}/nodes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateNodeInstanceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateNodeInstanceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createNodeInstance: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteInstance deletes a node instance
*/
func (a *Client) DeleteInstance(params *DeleteInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteInstanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteInstanceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteInstance",
		Method:             "DELETE",
		PathPattern:        "/api/v1/customers/{cUUID}/providers/{pUUID}/instances/{instanceIP}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteInstanceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteInstanceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteInstance: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DetachedNodeAction detacheds node action
*/
func (a *Client) DetachedNodeAction(params *DetachedNodeActionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DetachedNodeActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDetachedNodeActionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "detachedNodeAction",
		Method:             "POST",
		PathPattern:        "/api/v1/customers/{cUUID}/providers/{pUUID}/instances/{instanceIP}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DetachedNodeActionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DetachedNodeActionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for detachedNodeAction: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetNodeInstance gets a node instance
*/
func (a *Client) GetNodeInstance(params *GetNodeInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetNodeInstanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNodeInstanceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getNodeInstance",
		Method:             "GET",
		PathPattern:        "/api/v1/customers/{cUUID}/nodes/{nodeUUID}/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNodeInstanceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNodeInstanceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNodeInstance: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListByProvider lists all of a provider s node instances
*/
func (a *Client) ListByProvider(params *ListByProviderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListByProviderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListByProviderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listByProvider",
		Method:             "GET",
		PathPattern:        "/api/v1/customers/{cUUID}/providers/{pUUID}/nodes/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListByProviderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListByProviderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listByProvider: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListByZone lists all of a zone s node instances
*/
func (a *Client) ListByZone(params *ListByZoneParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListByZoneOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListByZoneParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listByZone",
		Method:             "GET",
		PathPattern:        "/api/v1/customers/{cUUID}/zones/{azUUID}/nodes/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListByZoneReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListByZoneOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listByZone: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  NodeAction updates a node
*/
func (a *Client) NodeAction(params *NodeActionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*NodeActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNodeActionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "nodeAction",
		Method:             "PUT",
		PathPattern:        "/api/v1/customers/{cUUID}/universes/{universeUUID}/nodes/{nodeName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &NodeActionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*NodeActionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for nodeAction: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
