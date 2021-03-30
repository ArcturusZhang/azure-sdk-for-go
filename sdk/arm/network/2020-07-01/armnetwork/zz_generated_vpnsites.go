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

// VPNSitesClient contains the methods for the VPNSites group.
// Don't use this type directly, use NewVPNSitesClient() instead.
type VPNSitesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewVPNSitesClient creates a new instance of VPNSitesClient with the specified values.
func NewVPNSitesClient(con *armcore.Connection, subscriptionID string) *VPNSitesClient {
	return &VPNSitesClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates a VpnSite resource if it doesn't exist else updates the existing VpnSite.
func (client *VPNSitesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, vpnSiteName string, vpnSiteParameters VPNSite, options *VPNSitesBeginCreateOrUpdateOptions) (VPNSitePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, vpnSiteName, vpnSiteParameters, options)
	if err != nil {
		return VPNSitePollerResponse{}, err
	}
	result := VPNSitePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VPNSitesClient.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return VPNSitePollerResponse{}, err
	}
	poller := &vpnSitePoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (VPNSiteResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new VPNSitePoller from the specified resume token.
// token - The value must come from a previous call to VPNSitePoller.ResumeToken().
func (client *VPNSitesClient) ResumeCreateOrUpdate(token string) (VPNSitePoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VPNSitesClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &vpnSitePoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates a VpnSite resource if it doesn't exist else updates the existing VpnSite.
func (client *VPNSitesClient) createOrUpdate(ctx context.Context, resourceGroupName string, vpnSiteName string, vpnSiteParameters VPNSite, options *VPNSitesBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, vpnSiteName, vpnSiteParameters, options)
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
func (client *VPNSitesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, vpnSiteName string, vpnSiteParameters VPNSite, options *VPNSitesBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites/{vpnSiteName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vpnSiteName == "" {
		return nil, errors.New("parameter vpnSiteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vpnSiteName}", url.PathEscape(vpnSiteName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(vpnSiteParameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *VPNSitesClient) createOrUpdateHandleResponse(resp *azcore.Response) (VPNSiteResponse, error) {
	var val *VPNSite
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return VPNSiteResponse{}, err
	}
	return VPNSiteResponse{RawResponse: resp.Response, VPNSite: val}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *VPNSitesClient) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDelete - Deletes a VpnSite.
func (client *VPNSitesClient) BeginDelete(ctx context.Context, resourceGroupName string, vpnSiteName string, options *VPNSitesBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.delete(ctx, resourceGroupName, vpnSiteName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VPNSitesClient.Delete", "location", resp, client.deleteHandleError)
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
func (client *VPNSitesClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VPNSitesClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes a VpnSite.
func (client *VPNSitesClient) delete(ctx context.Context, resourceGroupName string, vpnSiteName string, options *VPNSitesBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, vpnSiteName, options)
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
func (client *VPNSitesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, vpnSiteName string, options *VPNSitesBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites/{vpnSiteName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vpnSiteName == "" {
		return nil, errors.New("parameter vpnSiteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vpnSiteName}", url.PathEscape(vpnSiteName))
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
func (client *VPNSitesClient) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Retrieves the details of a VPN site.
func (client *VPNSitesClient) Get(ctx context.Context, resourceGroupName string, vpnSiteName string, options *VPNSitesGetOptions) (VPNSiteResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, vpnSiteName, options)
	if err != nil {
		return VPNSiteResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return VPNSiteResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return VPNSiteResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VPNSitesClient) getCreateRequest(ctx context.Context, resourceGroupName string, vpnSiteName string, options *VPNSitesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites/{vpnSiteName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vpnSiteName == "" {
		return nil, errors.New("parameter vpnSiteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vpnSiteName}", url.PathEscape(vpnSiteName))
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
func (client *VPNSitesClient) getHandleResponse(resp *azcore.Response) (VPNSiteResponse, error) {
	var val *VPNSite
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return VPNSiteResponse{}, err
	}
	return VPNSiteResponse{RawResponse: resp.Response, VPNSite: val}, nil
}

// getHandleError handles the Get error response.
func (client *VPNSitesClient) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Lists all the VpnSites in a subscription.
func (client *VPNSitesClient) List(options *VPNSitesListOptions) ListVPNSitesResultPager {
	return &listVPNSitesResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp ListVPNSitesResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ListVPNSitesResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *VPNSitesClient) listCreateRequest(ctx context.Context, options *VPNSitesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/vpnSites"
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
func (client *VPNSitesClient) listHandleResponse(resp *azcore.Response) (ListVPNSitesResultResponse, error) {
	var val *ListVPNSitesResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ListVPNSitesResultResponse{}, err
	}
	return ListVPNSitesResultResponse{RawResponse: resp.Response, ListVPNSitesResult: val}, nil
}

// listHandleError handles the List error response.
func (client *VPNSitesClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListByResourceGroup - Lists all the vpnSites in a resource group.
func (client *VPNSitesClient) ListByResourceGroup(resourceGroupName string, options *VPNSitesListByResourceGroupOptions) ListVPNSitesResultPager {
	return &listVPNSitesResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.listByResourceGroupHandleResponse,
		errorer:   client.listByResourceGroupHandleError,
		advancer: func(ctx context.Context, resp ListVPNSitesResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ListVPNSitesResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *VPNSitesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *VPNSitesListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *VPNSitesClient) listByResourceGroupHandleResponse(resp *azcore.Response) (ListVPNSitesResultResponse, error) {
	var val *ListVPNSitesResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ListVPNSitesResultResponse{}, err
	}
	return ListVPNSitesResultResponse{RawResponse: resp.Response, ListVPNSitesResult: val}, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *VPNSitesClient) listByResourceGroupHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// UpdateTags - Updates VpnSite tags.
func (client *VPNSitesClient) UpdateTags(ctx context.Context, resourceGroupName string, vpnSiteName string, vpnSiteParameters TagsObject, options *VPNSitesUpdateTagsOptions) (VPNSiteResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, vpnSiteName, vpnSiteParameters, options)
	if err != nil {
		return VPNSiteResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return VPNSiteResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return VPNSiteResponse{}, client.updateTagsHandleError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *VPNSitesClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, vpnSiteName string, vpnSiteParameters TagsObject, options *VPNSitesUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites/{vpnSiteName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vpnSiteName == "" {
		return nil, errors.New("parameter vpnSiteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vpnSiteName}", url.PathEscape(vpnSiteName))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(vpnSiteParameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *VPNSitesClient) updateTagsHandleResponse(resp *azcore.Response) (VPNSiteResponse, error) {
	var val *VPNSite
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return VPNSiteResponse{}, err
	}
	return VPNSiteResponse{RawResponse: resp.Response, VPNSite: val}, nil
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *VPNSitesClient) updateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
