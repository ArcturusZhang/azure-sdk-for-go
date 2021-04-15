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

// ConnectionMonitorsClient contains the methods for the ConnectionMonitors group.
// Don't use this type directly, use NewConnectionMonitorsClient() instead.
type ConnectionMonitorsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewConnectionMonitorsClient creates a new instance of ConnectionMonitorsClient with the specified values.
func NewConnectionMonitorsClient(con *armcore.Connection, subscriptionID string) *ConnectionMonitorsClient {
	return &ConnectionMonitorsClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Create or update a connection monitor.
func (client *ConnectionMonitorsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, parameters ConnectionMonitor, options *ConnectionMonitorsBeginCreateOrUpdateOptions) (ConnectionMonitorResultPollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, networkWatcherName, connectionMonitorName, parameters, options)
	if err != nil {
		return ConnectionMonitorResultPollerResponse{}, err
	}
	result := ConnectionMonitorResultPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ConnectionMonitorsClient.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return ConnectionMonitorResultPollerResponse{}, err
	}
	poller := &connectionMonitorResultPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ConnectionMonitorResultResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new ConnectionMonitorResultPoller from the specified resume token.
// token - The value must come from a previous call to ConnectionMonitorResultPoller.ResumeToken().
func (client *ConnectionMonitorsClient) ResumeCreateOrUpdate(token string) (ConnectionMonitorResultPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ConnectionMonitorsClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &connectionMonitorResultPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Create or update a connection monitor.
func (client *ConnectionMonitorsClient) createOrUpdate(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, parameters ConnectionMonitor, options *ConnectionMonitorsBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, networkWatcherName, connectionMonitorName, parameters, options)
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
func (client *ConnectionMonitorsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, parameters ConnectionMonitor, options *ConnectionMonitorsBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/connectionMonitors/{connectionMonitorName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if connectionMonitorName == "" {
		return nil, errors.New("parameter connectionMonitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionMonitorName}", url.PathEscape(connectionMonitorName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-07-01")
	if options != nil && options.Migrate != nil {
		reqQP.Set("migrate", *options.Migrate)
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ConnectionMonitorsClient) createOrUpdateHandleResponse(resp *azcore.Response) (ConnectionMonitorResultResponse, error) {
	var val *ConnectionMonitorResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ConnectionMonitorResultResponse{}, err
	}
	return ConnectionMonitorResultResponse{RawResponse: resp.Response, ConnectionMonitorResult: val}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ConnectionMonitorsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDelete - Deletes the specified connection monitor.
func (client *ConnectionMonitorsClient) BeginDelete(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, networkWatcherName, connectionMonitorName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ConnectionMonitorsClient.Delete", "location", resp, client.deleteHandleError)
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
func (client *ConnectionMonitorsClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ConnectionMonitorsClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified connection monitor.
func (client *ConnectionMonitorsClient) deleteOperation(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, networkWatcherName, connectionMonitorName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ConnectionMonitorsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/connectionMonitors/{connectionMonitorName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if connectionMonitorName == "" {
		return nil, errors.New("parameter connectionMonitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionMonitorName}", url.PathEscape(connectionMonitorName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-07-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ConnectionMonitorsClient) deleteHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets a connection monitor by name.
func (client *ConnectionMonitorsClient) Get(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsGetOptions) (ConnectionMonitorResultResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkWatcherName, connectionMonitorName, options)
	if err != nil {
		return ConnectionMonitorResultResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ConnectionMonitorResultResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ConnectionMonitorResultResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ConnectionMonitorsClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/connectionMonitors/{connectionMonitorName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if connectionMonitorName == "" {
		return nil, errors.New("parameter connectionMonitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionMonitorName}", url.PathEscape(connectionMonitorName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-07-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ConnectionMonitorsClient) getHandleResponse(resp *azcore.Response) (ConnectionMonitorResultResponse, error) {
	var val *ConnectionMonitorResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ConnectionMonitorResultResponse{}, err
	}
	return ConnectionMonitorResultResponse{RawResponse: resp.Response, ConnectionMonitorResult: val}, nil
}

// getHandleError handles the Get error response.
func (client *ConnectionMonitorsClient) getHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Lists all connection monitors for the specified Network Watcher.
func (client *ConnectionMonitorsClient) List(ctx context.Context, resourceGroupName string, networkWatcherName string, options *ConnectionMonitorsListOptions) (ConnectionMonitorListResultResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, networkWatcherName, options)
	if err != nil {
		return ConnectionMonitorListResultResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ConnectionMonitorListResultResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ConnectionMonitorListResultResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *ConnectionMonitorsClient) listCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, options *ConnectionMonitorsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/connectionMonitors"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-07-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ConnectionMonitorsClient) listHandleResponse(resp *azcore.Response) (ConnectionMonitorListResultResponse, error) {
	var val *ConnectionMonitorListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ConnectionMonitorListResultResponse{}, err
	}
	return ConnectionMonitorListResultResponse{RawResponse: resp.Response, ConnectionMonitorListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *ConnectionMonitorsClient) listHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginQuery - Query a snapshot of the most recent connection states.
func (client *ConnectionMonitorsClient) BeginQuery(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsBeginQueryOptions) (ConnectionMonitorQueryResultPollerResponse, error) {
	resp, err := client.query(ctx, resourceGroupName, networkWatcherName, connectionMonitorName, options)
	if err != nil {
		return ConnectionMonitorQueryResultPollerResponse{}, err
	}
	result := ConnectionMonitorQueryResultPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ConnectionMonitorsClient.Query", "location", resp, client.queryHandleError)
	if err != nil {
		return ConnectionMonitorQueryResultPollerResponse{}, err
	}
	poller := &connectionMonitorQueryResultPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ConnectionMonitorQueryResultResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeQuery creates a new ConnectionMonitorQueryResultPoller from the specified resume token.
// token - The value must come from a previous call to ConnectionMonitorQueryResultPoller.ResumeToken().
func (client *ConnectionMonitorsClient) ResumeQuery(token string) (ConnectionMonitorQueryResultPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ConnectionMonitorsClient.Query", token, client.queryHandleError)
	if err != nil {
		return nil, err
	}
	return &connectionMonitorQueryResultPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Query - Query a snapshot of the most recent connection states.
func (client *ConnectionMonitorsClient) query(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsBeginQueryOptions) (*azcore.Response, error) {
	req, err := client.queryCreateRequest(ctx, resourceGroupName, networkWatcherName, connectionMonitorName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.queryHandleError(resp)
	}
	return resp, nil
}

// queryCreateRequest creates the Query request.
func (client *ConnectionMonitorsClient) queryCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsBeginQueryOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/connectionMonitors/{connectionMonitorName}/query"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if connectionMonitorName == "" {
		return nil, errors.New("parameter connectionMonitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionMonitorName}", url.PathEscape(connectionMonitorName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-07-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// queryHandleResponse handles the Query response.
func (client *ConnectionMonitorsClient) queryHandleResponse(resp *azcore.Response) (ConnectionMonitorQueryResultResponse, error) {
	var val *ConnectionMonitorQueryResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ConnectionMonitorQueryResultResponse{}, err
	}
	return ConnectionMonitorQueryResultResponse{RawResponse: resp.Response, ConnectionMonitorQueryResult: val}, nil
}

// queryHandleError handles the Query error response.
func (client *ConnectionMonitorsClient) queryHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginStart - Starts the specified connection monitor.
func (client *ConnectionMonitorsClient) BeginStart(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsBeginStartOptions) (HTTPPollerResponse, error) {
	resp, err := client.start(ctx, resourceGroupName, networkWatcherName, connectionMonitorName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ConnectionMonitorsClient.Start", "location", resp, client.startHandleError)
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

// ResumeStart creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *ConnectionMonitorsClient) ResumeStart(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ConnectionMonitorsClient.Start", token, client.startHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Start - Starts the specified connection monitor.
func (client *ConnectionMonitorsClient) start(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsBeginStartOptions) (*azcore.Response, error) {
	req, err := client.startCreateRequest(ctx, resourceGroupName, networkWatcherName, connectionMonitorName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.startHandleError(resp)
	}
	return resp, nil
}

// startCreateRequest creates the Start request.
func (client *ConnectionMonitorsClient) startCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsBeginStartOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/connectionMonitors/{connectionMonitorName}/start"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if connectionMonitorName == "" {
		return nil, errors.New("parameter connectionMonitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionMonitorName}", url.PathEscape(connectionMonitorName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-07-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// startHandleError handles the Start error response.
func (client *ConnectionMonitorsClient) startHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginStop - Stops the specified connection monitor.
func (client *ConnectionMonitorsClient) BeginStop(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsBeginStopOptions) (HTTPPollerResponse, error) {
	resp, err := client.stop(ctx, resourceGroupName, networkWatcherName, connectionMonitorName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ConnectionMonitorsClient.Stop", "location", resp, client.stopHandleError)
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

// ResumeStop creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *ConnectionMonitorsClient) ResumeStop(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ConnectionMonitorsClient.Stop", token, client.stopHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Stop - Stops the specified connection monitor.
func (client *ConnectionMonitorsClient) stop(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsBeginStopOptions) (*azcore.Response, error) {
	req, err := client.stopCreateRequest(ctx, resourceGroupName, networkWatcherName, connectionMonitorName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.stopHandleError(resp)
	}
	return resp, nil
}

// stopCreateRequest creates the Stop request.
func (client *ConnectionMonitorsClient) stopCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsBeginStopOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/connectionMonitors/{connectionMonitorName}/stop"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if connectionMonitorName == "" {
		return nil, errors.New("parameter connectionMonitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionMonitorName}", url.PathEscape(connectionMonitorName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-07-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// stopHandleError handles the Stop error response.
func (client *ConnectionMonitorsClient) stopHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// UpdateTags - Update tags of the specified connection monitor.
func (client *ConnectionMonitorsClient) UpdateTags(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, parameters TagsObject, options *ConnectionMonitorsUpdateTagsOptions) (ConnectionMonitorResultResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, networkWatcherName, connectionMonitorName, parameters, options)
	if err != nil {
		return ConnectionMonitorResultResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ConnectionMonitorResultResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ConnectionMonitorResultResponse{}, client.updateTagsHandleError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *ConnectionMonitorsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, parameters TagsObject, options *ConnectionMonitorsUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/connectionMonitors/{connectionMonitorName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if connectionMonitorName == "" {
		return nil, errors.New("parameter connectionMonitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionMonitorName}", url.PathEscape(connectionMonitorName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-07-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *ConnectionMonitorsClient) updateTagsHandleResponse(resp *azcore.Response) (ConnectionMonitorResultResponse, error) {
	var val *ConnectionMonitorResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ConnectionMonitorResultResponse{}, err
	}
	return ConnectionMonitorResultResponse{RawResponse: resp.Response, ConnectionMonitorResult: val}, nil
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *ConnectionMonitorsClient) updateTagsHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
