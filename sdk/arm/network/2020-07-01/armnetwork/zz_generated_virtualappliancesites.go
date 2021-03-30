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

// VirtualApplianceSitesClient contains the methods for the VirtualApplianceSites group.
// Don't use this type directly, use NewVirtualApplianceSitesClient() instead.
type VirtualApplianceSitesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewVirtualApplianceSitesClient creates a new instance of VirtualApplianceSitesClient with the specified values.
func NewVirtualApplianceSitesClient(con *armcore.Connection, subscriptionID string) *VirtualApplianceSitesClient {
	return &VirtualApplianceSitesClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates the specified Network Virtual Appliance Site.
func (client *VirtualApplianceSitesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, siteName string, parameters VirtualApplianceSite, options *VirtualApplianceSitesBeginCreateOrUpdateOptions) (VirtualApplianceSitePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, networkVirtualApplianceName, siteName, parameters, options)
	if err != nil {
		return VirtualApplianceSitePollerResponse{}, err
	}
	result := VirtualApplianceSitePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualApplianceSitesClient.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return VirtualApplianceSitePollerResponse{}, err
	}
	poller := &virtualApplianceSitePoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (VirtualApplianceSiteResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new VirtualApplianceSitePoller from the specified resume token.
// token - The value must come from a previous call to VirtualApplianceSitePoller.ResumeToken().
func (client *VirtualApplianceSitesClient) ResumeCreateOrUpdate(token string) (VirtualApplianceSitePoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualApplianceSitesClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &virtualApplianceSitePoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates or updates the specified Network Virtual Appliance Site.
func (client *VirtualApplianceSitesClient) createOrUpdate(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, siteName string, parameters VirtualApplianceSite, options *VirtualApplianceSitesBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, networkVirtualApplianceName, siteName, parameters, options)
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
func (client *VirtualApplianceSitesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, siteName string, parameters VirtualApplianceSite, options *VirtualApplianceSitesBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances/{networkVirtualApplianceName}/virtualApplianceSites/{siteName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkVirtualApplianceName == "" {
		return nil, errors.New("parameter networkVirtualApplianceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkVirtualApplianceName}", url.PathEscape(networkVirtualApplianceName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
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
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *VirtualApplianceSitesClient) createOrUpdateHandleResponse(resp *azcore.Response) (VirtualApplianceSiteResponse, error) {
	var val *VirtualApplianceSite
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return VirtualApplianceSiteResponse{}, err
	}
	return VirtualApplianceSiteResponse{RawResponse: resp.Response, VirtualApplianceSite: val}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *VirtualApplianceSitesClient) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDelete - Deletes the specified site from a Virtual Appliance.
func (client *VirtualApplianceSitesClient) BeginDelete(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, siteName string, options *VirtualApplianceSitesBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.delete(ctx, resourceGroupName, networkVirtualApplianceName, siteName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualApplianceSitesClient.Delete", "location", resp, client.deleteHandleError)
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
func (client *VirtualApplianceSitesClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualApplianceSitesClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified site from a Virtual Appliance.
func (client *VirtualApplianceSitesClient) delete(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, siteName string, options *VirtualApplianceSitesBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, networkVirtualApplianceName, siteName, options)
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
func (client *VirtualApplianceSitesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, siteName string, options *VirtualApplianceSitesBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances/{networkVirtualApplianceName}/virtualApplianceSites/{siteName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkVirtualApplianceName == "" {
		return nil, errors.New("parameter networkVirtualApplianceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkVirtualApplianceName}", url.PathEscape(networkVirtualApplianceName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
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
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *VirtualApplianceSitesClient) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets the specified Virtual Appliance Site.
func (client *VirtualApplianceSitesClient) Get(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, siteName string, options *VirtualApplianceSitesGetOptions) (VirtualApplianceSiteResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkVirtualApplianceName, siteName, options)
	if err != nil {
		return VirtualApplianceSiteResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return VirtualApplianceSiteResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return VirtualApplianceSiteResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualApplianceSitesClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, siteName string, options *VirtualApplianceSitesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances/{networkVirtualApplianceName}/virtualApplianceSites/{siteName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkVirtualApplianceName == "" {
		return nil, errors.New("parameter networkVirtualApplianceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkVirtualApplianceName}", url.PathEscape(networkVirtualApplianceName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
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
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualApplianceSitesClient) getHandleResponse(resp *azcore.Response) (VirtualApplianceSiteResponse, error) {
	var val *VirtualApplianceSite
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return VirtualApplianceSiteResponse{}, err
	}
	return VirtualApplianceSiteResponse{RawResponse: resp.Response, VirtualApplianceSite: val}, nil
}

// getHandleError handles the Get error response.
func (client *VirtualApplianceSitesClient) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Lists all Network Virtual Appliance Sites in a Network Virtual Appliance resource.
func (client *VirtualApplianceSitesClient) List(resourceGroupName string, networkVirtualApplianceName string, options *VirtualApplianceSitesListOptions) NetworkVirtualApplianceSiteListResultPager {
	return &networkVirtualApplianceSiteListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, networkVirtualApplianceName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp NetworkVirtualApplianceSiteListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.NetworkVirtualApplianceSiteListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *VirtualApplianceSitesClient) listCreateRequest(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, options *VirtualApplianceSitesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances/{networkVirtualApplianceName}/virtualApplianceSites"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkVirtualApplianceName == "" {
		return nil, errors.New("parameter networkVirtualApplianceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkVirtualApplianceName}", url.PathEscape(networkVirtualApplianceName))
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
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualApplianceSitesClient) listHandleResponse(resp *azcore.Response) (NetworkVirtualApplianceSiteListResultResponse, error) {
	var val *NetworkVirtualApplianceSiteListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return NetworkVirtualApplianceSiteListResultResponse{}, err
	}
	return NetworkVirtualApplianceSiteListResultResponse{RawResponse: resp.Response, NetworkVirtualApplianceSiteListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *VirtualApplianceSitesClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
