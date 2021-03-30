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

// RoutesClient contains the methods for the Routes group.
// Don't use this type directly, use NewRoutesClient() instead.
type RoutesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewRoutesClient creates a new instance of RoutesClient with the specified values.
func NewRoutesClient(con *armcore.Connection, subscriptionID string) *RoutesClient {
	return &RoutesClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates a route in the specified route table.
func (client *RoutesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, routeTableName string, routeName string, routeParameters Route, options *RoutesBeginCreateOrUpdateOptions) (RoutePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, routeTableName, routeName, routeParameters, options)
	if err != nil {
		return RoutePollerResponse{}, err
	}
	result := RoutePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("RoutesClient.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return RoutePollerResponse{}, err
	}
	poller := &routePoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (RouteResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new RoutePoller from the specified resume token.
// token - The value must come from a previous call to RoutePoller.ResumeToken().
func (client *RoutesClient) ResumeCreateOrUpdate(token string) (RoutePoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("RoutesClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &routePoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates or updates a route in the specified route table.
func (client *RoutesClient) createOrUpdate(ctx context.Context, resourceGroupName string, routeTableName string, routeName string, routeParameters Route, options *RoutesBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, routeTableName, routeName, routeParameters, options)
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
func (client *RoutesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, routeTableName string, routeName string, routeParameters Route, options *RoutesBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}/routes/{routeName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if routeTableName == "" {
		return nil, errors.New("parameter routeTableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeTableName}", url.PathEscape(routeTableName))
	if routeName == "" {
		return nil, errors.New("parameter routeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeName}", url.PathEscape(routeName))
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
	return req, req.MarshalAsJSON(routeParameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *RoutesClient) createOrUpdateHandleResponse(resp *azcore.Response) (RouteResponse, error) {
	var val *Route
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return RouteResponse{}, err
	}
	return RouteResponse{RawResponse: resp.Response, Route: val}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *RoutesClient) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDelete - Deletes the specified route from a route table.
func (client *RoutesClient) BeginDelete(ctx context.Context, resourceGroupName string, routeTableName string, routeName string, options *RoutesBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.delete(ctx, resourceGroupName, routeTableName, routeName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("RoutesClient.Delete", "location", resp, client.deleteHandleError)
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
func (client *RoutesClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("RoutesClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified route from a route table.
func (client *RoutesClient) delete(ctx context.Context, resourceGroupName string, routeTableName string, routeName string, options *RoutesBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, routeTableName, routeName, options)
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
func (client *RoutesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, routeTableName string, routeName string, options *RoutesBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}/routes/{routeName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if routeTableName == "" {
		return nil, errors.New("parameter routeTableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeTableName}", url.PathEscape(routeTableName))
	if routeName == "" {
		return nil, errors.New("parameter routeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeName}", url.PathEscape(routeName))
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
func (client *RoutesClient) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets the specified route from a route table.
func (client *RoutesClient) Get(ctx context.Context, resourceGroupName string, routeTableName string, routeName string, options *RoutesGetOptions) (RouteResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, routeTableName, routeName, options)
	if err != nil {
		return RouteResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return RouteResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return RouteResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RoutesClient) getCreateRequest(ctx context.Context, resourceGroupName string, routeTableName string, routeName string, options *RoutesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}/routes/{routeName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if routeTableName == "" {
		return nil, errors.New("parameter routeTableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeTableName}", url.PathEscape(routeTableName))
	if routeName == "" {
		return nil, errors.New("parameter routeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeName}", url.PathEscape(routeName))
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
func (client *RoutesClient) getHandleResponse(resp *azcore.Response) (RouteResponse, error) {
	var val *Route
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return RouteResponse{}, err
	}
	return RouteResponse{RawResponse: resp.Response, Route: val}, nil
}

// getHandleError handles the Get error response.
func (client *RoutesClient) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Gets all routes in a route table.
func (client *RoutesClient) List(resourceGroupName string, routeTableName string, options *RoutesListOptions) RouteListResultPager {
	return &routeListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, routeTableName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp RouteListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.RouteListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *RoutesClient) listCreateRequest(ctx context.Context, resourceGroupName string, routeTableName string, options *RoutesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}/routes"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if routeTableName == "" {
		return nil, errors.New("parameter routeTableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeTableName}", url.PathEscape(routeTableName))
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
func (client *RoutesClient) listHandleResponse(resp *azcore.Response) (RouteListResultResponse, error) {
	var val *RouteListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return RouteListResultResponse{}, err
	}
	return RouteListResultResponse{RawResponse: resp.Response, RouteListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *RoutesClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
