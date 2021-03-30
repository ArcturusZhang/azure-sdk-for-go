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
	"time"
)

// VirtualNetworkPeeringsClient contains the methods for the VirtualNetworkPeerings group.
// Don't use this type directly, use NewVirtualNetworkPeeringsClient() instead.
type VirtualNetworkPeeringsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewVirtualNetworkPeeringsClient creates a new instance of VirtualNetworkPeeringsClient with the specified values.
func NewVirtualNetworkPeeringsClient(con *armcore.Connection, subscriptionID string) *VirtualNetworkPeeringsClient {
	return &VirtualNetworkPeeringsClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates a peering in the specified virtual network.
func (client *VirtualNetworkPeeringsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkName string, virtualNetworkPeeringName string, virtualNetworkPeeringParameters VirtualNetworkPeering, options *VirtualNetworkPeeringsBeginCreateOrUpdateOptions) (VirtualNetworkPeeringPollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, virtualNetworkName, virtualNetworkPeeringName, virtualNetworkPeeringParameters, options)
	if err != nil {
		return VirtualNetworkPeeringPollerResponse{}, err
	}
	result := VirtualNetworkPeeringPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualNetworkPeeringsClient.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return VirtualNetworkPeeringPollerResponse{}, err
	}
	poller := &virtualNetworkPeeringPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (VirtualNetworkPeeringResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new VirtualNetworkPeeringPoller from the specified resume token.
// token - The value must come from a previous call to VirtualNetworkPeeringPoller.ResumeToken().
func (client *VirtualNetworkPeeringsClient) ResumeCreateOrUpdate(token string) (VirtualNetworkPeeringPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualNetworkPeeringsClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &virtualNetworkPeeringPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates or updates a peering in the specified virtual network.
func (client *VirtualNetworkPeeringsClient) createOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkName string, virtualNetworkPeeringName string, virtualNetworkPeeringParameters VirtualNetworkPeering, options *VirtualNetworkPeeringsBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, virtualNetworkName, virtualNetworkPeeringName, virtualNetworkPeeringParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VirtualNetworkPeeringsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, virtualNetworkPeeringName string, virtualNetworkPeeringParameters VirtualNetworkPeering, options *VirtualNetworkPeeringsBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/virtualNetworkPeerings/{virtualNetworkPeeringName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	if virtualNetworkPeeringName == "" {
		return nil, errors.New("parameter virtualNetworkPeeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkPeeringName}", url.PathEscape(virtualNetworkPeeringName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-05-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(virtualNetworkPeeringParameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *VirtualNetworkPeeringsClient) createOrUpdateHandleResponse(resp *azcore.Response) (VirtualNetworkPeeringResponse, error) {
	var val *VirtualNetworkPeering
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return VirtualNetworkPeeringResponse{}, err
	}
	return VirtualNetworkPeeringResponse{RawResponse: resp.Response, VirtualNetworkPeering: val}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *VirtualNetworkPeeringsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDelete - Deletes the specified virtual network peering.
func (client *VirtualNetworkPeeringsClient) BeginDelete(ctx context.Context, resourceGroupName string, virtualNetworkName string, virtualNetworkPeeringName string, options *VirtualNetworkPeeringsBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.delete(ctx, resourceGroupName, virtualNetworkName, virtualNetworkPeeringName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualNetworkPeeringsClient.Delete", "location", resp, client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *VirtualNetworkPeeringsClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualNetworkPeeringsClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified virtual network peering.
func (client *VirtualNetworkPeeringsClient) delete(ctx context.Context, resourceGroupName string, virtualNetworkName string, virtualNetworkPeeringName string, options *VirtualNetworkPeeringsBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, virtualNetworkName, virtualNetworkPeeringName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *VirtualNetworkPeeringsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, virtualNetworkPeeringName string, options *VirtualNetworkPeeringsBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/virtualNetworkPeerings/{virtualNetworkPeeringName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	if virtualNetworkPeeringName == "" {
		return nil, errors.New("parameter virtualNetworkPeeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkPeeringName}", url.PathEscape(virtualNetworkPeeringName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
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

// deleteHandleError handles the Delete error response.
func (client *VirtualNetworkPeeringsClient) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets the specified virtual network peering.
func (client *VirtualNetworkPeeringsClient) Get(ctx context.Context, resourceGroupName string, virtualNetworkName string, virtualNetworkPeeringName string, options *VirtualNetworkPeeringsGetOptions) (VirtualNetworkPeeringResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, virtualNetworkName, virtualNetworkPeeringName, options)
	if err != nil {
		return VirtualNetworkPeeringResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return VirtualNetworkPeeringResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return VirtualNetworkPeeringResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualNetworkPeeringsClient) getCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, virtualNetworkPeeringName string, options *VirtualNetworkPeeringsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/virtualNetworkPeerings/{virtualNetworkPeeringName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	if virtualNetworkPeeringName == "" {
		return nil, errors.New("parameter virtualNetworkPeeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkPeeringName}", url.PathEscape(virtualNetworkPeeringName))
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
func (client *VirtualNetworkPeeringsClient) getHandleResponse(resp *azcore.Response) (VirtualNetworkPeeringResponse, error) {
	var val *VirtualNetworkPeering
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return VirtualNetworkPeeringResponse{}, err
	}
	return VirtualNetworkPeeringResponse{RawResponse: resp.Response, VirtualNetworkPeering: val}, nil
}

// getHandleError handles the Get error response.
func (client *VirtualNetworkPeeringsClient) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Gets all virtual network peerings in a virtual network.
func (client *VirtualNetworkPeeringsClient) List(resourceGroupName string, virtualNetworkName string, options *VirtualNetworkPeeringsListOptions) VirtualNetworkPeeringListResultPager {
	return &virtualNetworkPeeringListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, virtualNetworkName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp VirtualNetworkPeeringListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.VirtualNetworkPeeringListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *VirtualNetworkPeeringsClient) listCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, options *VirtualNetworkPeeringsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/virtualNetworkPeerings"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
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
func (client *VirtualNetworkPeeringsClient) listHandleResponse(resp *azcore.Response) (VirtualNetworkPeeringListResultResponse, error) {
	var val *VirtualNetworkPeeringListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return VirtualNetworkPeeringListResultResponse{}, err
	}
	return VirtualNetworkPeeringListResultResponse{RawResponse: resp.Response, VirtualNetworkPeeringListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *VirtualNetworkPeeringsClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
