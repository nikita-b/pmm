// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new nodes API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for nodes API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddContainerNode(params *AddContainerNodeParams, opts ...ClientOption) (*AddContainerNodeOK, error)

	AddGenericNode(params *AddGenericNodeParams, opts ...ClientOption) (*AddGenericNodeOK, error)

	AddRemoteAzureDatabaseNode(params *AddRemoteAzureDatabaseNodeParams, opts ...ClientOption) (*AddRemoteAzureDatabaseNodeOK, error)

	AddRemoteNode(params *AddRemoteNodeParams, opts ...ClientOption) (*AddRemoteNodeOK, error)

	AddRemoteRDSNode(params *AddRemoteRDSNodeParams, opts ...ClientOption) (*AddRemoteRDSNodeOK, error)

	GetNode(params *GetNodeParams, opts ...ClientOption) (*GetNodeOK, error)

	ListNodes(params *ListNodesParams, opts ...ClientOption) (*ListNodesOK, error)

	RemoveNode(params *RemoveNodeParams, opts ...ClientOption) (*RemoveNodeOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AddContainerNode adds container node

Adds container Node.
*/
func (a *Client) AddContainerNode(params *AddContainerNodeParams, opts ...ClientOption) (*AddContainerNodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddContainerNodeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AddContainerNode",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Nodes/AddContainer",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddContainerNodeReader{formats: a.formats},
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
	success, ok := result.(*AddContainerNodeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddContainerNodeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
AddGenericNode adds generic node

Adds generic Node.
*/
func (a *Client) AddGenericNode(params *AddGenericNodeParams, opts ...ClientOption) (*AddGenericNodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddGenericNodeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AddGenericNode",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Nodes/AddGeneric",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddGenericNodeReader{formats: a.formats},
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
	success, ok := result.(*AddGenericNodeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddGenericNodeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
AddRemoteAzureDatabaseNode adds remote azure database node

Adds remote Azure database Node.
*/
func (a *Client) AddRemoteAzureDatabaseNode(params *AddRemoteAzureDatabaseNodeParams, opts ...ClientOption) (*AddRemoteAzureDatabaseNodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddRemoteAzureDatabaseNodeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AddRemoteAzureDatabaseNode",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Nodes/AddRemoteAzureDatabase",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddRemoteAzureDatabaseNodeReader{formats: a.formats},
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
	success, ok := result.(*AddRemoteAzureDatabaseNodeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddRemoteAzureDatabaseNodeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
AddRemoteNode adds remote node

Adds remote Node.
*/
func (a *Client) AddRemoteNode(params *AddRemoteNodeParams, opts ...ClientOption) (*AddRemoteNodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddRemoteNodeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AddRemoteNode",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Nodes/AddRemote",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddRemoteNodeReader{formats: a.formats},
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
	success, ok := result.(*AddRemoteNodeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddRemoteNodeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
AddRemoteRDSNode adds remote RDS node

Adds remote RDS Node.
*/
func (a *Client) AddRemoteRDSNode(params *AddRemoteRDSNodeParams, opts ...ClientOption) (*AddRemoteRDSNodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddRemoteRDSNodeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AddRemoteRDSNode",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Nodes/AddRemoteRDS",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddRemoteRDSNodeReader{formats: a.formats},
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
	success, ok := result.(*AddRemoteRDSNodeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddRemoteRDSNodeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetNode gets node

Returns a single Node by ID.
*/
func (a *Client) GetNode(params *GetNodeParams, opts ...ClientOption) (*GetNodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNodeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetNode",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Nodes/Get",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNodeReader{formats: a.formats},
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
	success, ok := result.(*GetNodeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetNodeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListNodes lists nodes

Returns a list of all Nodes.
*/
func (a *Client) ListNodes(params *ListNodesParams, opts ...ClientOption) (*ListNodesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListNodesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListNodes",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Nodes/List",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListNodesReader{formats: a.formats},
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
	success, ok := result.(*ListNodesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListNodesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
RemoveNode removes node

Removes Node.
*/
func (a *Client) RemoveNode(params *RemoveNodeParams, opts ...ClientOption) (*RemoveNodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveNodeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RemoveNode",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Nodes/Remove",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RemoveNodeReader{formats: a.formats},
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
	success, ok := result.(*RemoveNodeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RemoveNodeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
