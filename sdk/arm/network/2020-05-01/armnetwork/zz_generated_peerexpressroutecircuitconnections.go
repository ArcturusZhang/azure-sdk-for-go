// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// PeerExpressRouteCircuitConnectionsClient contains the methods for the PeerExpressRouteCircuitConnections group.
// Don't use this type directly, use NewPeerExpressRouteCircuitConnectionsClient() instead.
type PeerExpressRouteCircuitConnectionsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewPeerExpressRouteCircuitConnectionsClient creates a new instance of PeerExpressRouteCircuitConnectionsClient with the specified values.
func NewPeerExpressRouteCircuitConnectionsClient(con *armcore.Connection, subscriptionID string) *PeerExpressRouteCircuitConnectionsClient {
	return &PeerExpressRouteCircuitConnectionsClient{con: con, subscriptionID: subscriptionID}
}

// Get - Gets the specified Peer Express Route Circuit Connection from the specified express route circuit.
func (client *PeerExpressRouteCircuitConnectionsClient) Get(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, connectionName string, options *PeerExpressRouteCircuitConnectionsGetOptions) (PeerExpressRouteCircuitConnectionResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, circuitName, peeringName, connectionName, options)
	if err != nil {
		return PeerExpressRouteCircuitConnectionResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return PeerExpressRouteCircuitConnectionResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return PeerExpressRouteCircuitConnectionResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PeerExpressRouteCircuitConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, connectionName string, options *PeerExpressRouteCircuitConnectionsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/peerConnections/{connectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if circuitName == "" {
		return nil, errors.New("parameter circuitName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-05-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PeerExpressRouteCircuitConnectionsClient) getHandleResponse(resp *azcore.Response) (PeerExpressRouteCircuitConnectionResponse, error) {
	var val *PeerExpressRouteCircuitConnection
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return PeerExpressRouteCircuitConnectionResponse{}, err
	}
	return PeerExpressRouteCircuitConnectionResponse{RawResponse: resp.Response, PeerExpressRouteCircuitConnection: val}, nil
}

// getHandleError handles the Get error response.
func (client *PeerExpressRouteCircuitConnectionsClient) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Gets all global reach peer connections associated with a private peering in an express route circuit.
func (client *PeerExpressRouteCircuitConnectionsClient) List(resourceGroupName string, circuitName string, peeringName string, options *PeerExpressRouteCircuitConnectionsListOptions) PeerExpressRouteCircuitConnectionListResultPager {
	return &peerExpressRouteCircuitConnectionListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, circuitName, peeringName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp PeerExpressRouteCircuitConnectionListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.PeerExpressRouteCircuitConnectionListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *PeerExpressRouteCircuitConnectionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, options *PeerExpressRouteCircuitConnectionsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/peerConnections"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if circuitName == "" {
		return nil, errors.New("parameter circuitName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-05-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PeerExpressRouteCircuitConnectionsClient) listHandleResponse(resp *azcore.Response) (PeerExpressRouteCircuitConnectionListResultResponse, error) {
	var val *PeerExpressRouteCircuitConnectionListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return PeerExpressRouteCircuitConnectionListResultResponse{}, err
	}
	return PeerExpressRouteCircuitConnectionListResultResponse{RawResponse: resp.Response, PeerExpressRouteCircuitConnectionListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *PeerExpressRouteCircuitConnectionsClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
