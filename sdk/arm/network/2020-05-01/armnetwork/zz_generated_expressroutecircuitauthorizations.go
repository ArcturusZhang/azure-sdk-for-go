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

// ExpressRouteCircuitAuthorizationsClient contains the methods for the ExpressRouteCircuitAuthorizations group.
// Don't use this type directly, use NewExpressRouteCircuitAuthorizationsClient() instead.
type ExpressRouteCircuitAuthorizationsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewExpressRouteCircuitAuthorizationsClient creates a new instance of ExpressRouteCircuitAuthorizationsClient with the specified values.
func NewExpressRouteCircuitAuthorizationsClient(con *armcore.Connection, subscriptionID string) *ExpressRouteCircuitAuthorizationsClient {
	return &ExpressRouteCircuitAuthorizationsClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates an authorization in the specified express route circuit.
func (client *ExpressRouteCircuitAuthorizationsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, circuitName string, authorizationName string, authorizationParameters ExpressRouteCircuitAuthorization, options *ExpressRouteCircuitAuthorizationsBeginCreateOrUpdateOptions) (ExpressRouteCircuitAuthorizationPollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, circuitName, authorizationName, authorizationParameters, options)
	if err != nil {
		return ExpressRouteCircuitAuthorizationPollerResponse{}, err
	}
	result := ExpressRouteCircuitAuthorizationPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ExpressRouteCircuitAuthorizationsClient.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return ExpressRouteCircuitAuthorizationPollerResponse{}, err
	}
	poller := &expressRouteCircuitAuthorizationPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ExpressRouteCircuitAuthorizationResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new ExpressRouteCircuitAuthorizationPoller from the specified resume token.
// token - The value must come from a previous call to ExpressRouteCircuitAuthorizationPoller.ResumeToken().
func (client *ExpressRouteCircuitAuthorizationsClient) ResumeCreateOrUpdate(token string) (ExpressRouteCircuitAuthorizationPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ExpressRouteCircuitAuthorizationsClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &expressRouteCircuitAuthorizationPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates or updates an authorization in the specified express route circuit.
func (client *ExpressRouteCircuitAuthorizationsClient) createOrUpdate(ctx context.Context, resourceGroupName string, circuitName string, authorizationName string, authorizationParameters ExpressRouteCircuitAuthorization, options *ExpressRouteCircuitAuthorizationsBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, circuitName, authorizationName, authorizationParameters, options)
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
func (client *ExpressRouteCircuitAuthorizationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, authorizationName string, authorizationParameters ExpressRouteCircuitAuthorization, options *ExpressRouteCircuitAuthorizationsBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/authorizations/{authorizationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if circuitName == "" {
		return nil, errors.New("parameter circuitName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	if authorizationName == "" {
		return nil, errors.New("parameter authorizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationName}", url.PathEscape(authorizationName))
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
	return req, req.MarshalAsJSON(authorizationParameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ExpressRouteCircuitAuthorizationsClient) createOrUpdateHandleResponse(resp *azcore.Response) (ExpressRouteCircuitAuthorizationResponse, error) {
	var val *ExpressRouteCircuitAuthorization
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ExpressRouteCircuitAuthorizationResponse{}, err
	}
	return ExpressRouteCircuitAuthorizationResponse{RawResponse: resp.Response, ExpressRouteCircuitAuthorization: val}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ExpressRouteCircuitAuthorizationsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDelete - Deletes the specified authorization from the specified express route circuit.
func (client *ExpressRouteCircuitAuthorizationsClient) BeginDelete(ctx context.Context, resourceGroupName string, circuitName string, authorizationName string, options *ExpressRouteCircuitAuthorizationsBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.delete(ctx, resourceGroupName, circuitName, authorizationName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ExpressRouteCircuitAuthorizationsClient.Delete", "location", resp, client.deleteHandleError)
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
func (client *ExpressRouteCircuitAuthorizationsClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ExpressRouteCircuitAuthorizationsClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified authorization from the specified express route circuit.
func (client *ExpressRouteCircuitAuthorizationsClient) delete(ctx context.Context, resourceGroupName string, circuitName string, authorizationName string, options *ExpressRouteCircuitAuthorizationsBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, circuitName, authorizationName, options)
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
func (client *ExpressRouteCircuitAuthorizationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, authorizationName string, options *ExpressRouteCircuitAuthorizationsBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/authorizations/{authorizationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if circuitName == "" {
		return nil, errors.New("parameter circuitName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	if authorizationName == "" {
		return nil, errors.New("parameter authorizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationName}", url.PathEscape(authorizationName))
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
func (client *ExpressRouteCircuitAuthorizationsClient) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets the specified authorization from the specified express route circuit.
func (client *ExpressRouteCircuitAuthorizationsClient) Get(ctx context.Context, resourceGroupName string, circuitName string, authorizationName string, options *ExpressRouteCircuitAuthorizationsGetOptions) (ExpressRouteCircuitAuthorizationResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, circuitName, authorizationName, options)
	if err != nil {
		return ExpressRouteCircuitAuthorizationResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ExpressRouteCircuitAuthorizationResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ExpressRouteCircuitAuthorizationResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ExpressRouteCircuitAuthorizationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, authorizationName string, options *ExpressRouteCircuitAuthorizationsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/authorizations/{authorizationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if circuitName == "" {
		return nil, errors.New("parameter circuitName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	if authorizationName == "" {
		return nil, errors.New("parameter authorizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationName}", url.PathEscape(authorizationName))
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
func (client *ExpressRouteCircuitAuthorizationsClient) getHandleResponse(resp *azcore.Response) (ExpressRouteCircuitAuthorizationResponse, error) {
	var val *ExpressRouteCircuitAuthorization
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ExpressRouteCircuitAuthorizationResponse{}, err
	}
	return ExpressRouteCircuitAuthorizationResponse{RawResponse: resp.Response, ExpressRouteCircuitAuthorization: val}, nil
}

// getHandleError handles the Get error response.
func (client *ExpressRouteCircuitAuthorizationsClient) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Gets all authorizations in an express route circuit.
func (client *ExpressRouteCircuitAuthorizationsClient) List(resourceGroupName string, circuitName string, options *ExpressRouteCircuitAuthorizationsListOptions) AuthorizationListResultPager {
	return &authorizationListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, circuitName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp AuthorizationListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.AuthorizationListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *ExpressRouteCircuitAuthorizationsClient) listCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, options *ExpressRouteCircuitAuthorizationsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/authorizations"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if circuitName == "" {
		return nil, errors.New("parameter circuitName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
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
func (client *ExpressRouteCircuitAuthorizationsClient) listHandleResponse(resp *azcore.Response) (AuthorizationListResultResponse, error) {
	var val *AuthorizationListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return AuthorizationListResultResponse{}, err
	}
	return AuthorizationListResultResponse{RawResponse: resp.Response, AuthorizationListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *ExpressRouteCircuitAuthorizationsClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
