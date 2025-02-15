// Code generated by go-swagger; DO NOT EDIT.

package db_clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new db clusters API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for db clusters API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteDBCluster(params *DeleteDBClusterParams, opts ...ClientOption) (*DeleteDBClusterOK, error)

	GetDBCluster(params *GetDBClusterParams, opts ...ClientOption) (*GetDBClusterOK, error)

	ListDBClusters(params *ListDBClustersParams, opts ...ClientOption) (*ListDBClustersOK, error)

	RestartDBCluster(params *RestartDBClusterParams, opts ...ClientOption) (*RestartDBClusterOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteDBCluster deletes DB cluster deletes DB cluster
*/
func (a *Client) DeleteDBCluster(params *DeleteDBClusterParams, opts ...ClientOption) (*DeleteDBClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDBClusterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteDBCluster",
		Method:             "POST",
		PathPattern:        "/v1/management/DBaaS/DBClusters/Delete",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteDBClusterReader{formats: a.formats},
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
	success, ok := result.(*DeleteDBClusterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteDBClusterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetDBCluster gets DB cluster returns parameters used to create a database cluster
*/
func (a *Client) GetDBCluster(params *GetDBClusterParams, opts ...ClientOption) (*GetDBClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDBClusterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetDBCluster",
		Method:             "POST",
		PathPattern:        "/v1/management/DBaaS/DBClusters/Get",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDBClusterReader{formats: a.formats},
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
	success, ok := result.(*GetDBClusterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetDBClusterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListDBClusters lists DB clusters returns a list of DB clusters
*/
func (a *Client) ListDBClusters(params *ListDBClustersParams, opts ...ClientOption) (*ListDBClustersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListDBClustersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListDBClusters",
		Method:             "POST",
		PathPattern:        "/v1/management/DBaaS/DBClusters/List",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListDBClustersReader{formats: a.formats},
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
	success, ok := result.(*ListDBClustersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListDBClustersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
RestartDBCluster restarts DB cluster restarts DB cluster
*/
func (a *Client) RestartDBCluster(params *RestartDBClusterParams, opts ...ClientOption) (*RestartDBClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRestartDBClusterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RestartDBCluster",
		Method:             "POST",
		PathPattern:        "/v1/management/DBaaS/DBClusters/Restart",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RestartDBClusterReader{formats: a.formats},
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
	success, ok := result.(*RestartDBClusterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RestartDBClusterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
