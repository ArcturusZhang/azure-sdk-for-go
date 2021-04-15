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

// ServiceTagsClient contains the methods for the ServiceTags group.
// Don't use this type directly, use NewServiceTagsClient() instead.
type ServiceTagsClient struct {
	con *armcore.Connection
	subscriptionID string
}

// NewServiceTagsClient creates a new instance of ServiceTagsClient with the specified values.
func NewServiceTagsClient(con *armcore.Connection, subscriptionID string) *ServiceTagsClient {
	return &ServiceTagsClient{con: con, subscriptionID: subscriptionID}
}

// List - Gets a list of service tag information resources.
func (client *ServiceTagsClient) List(ctx context.Context, location string, options *ServiceTagsListOptions) (ServiceTagsListResultResponse, error) {
	req, err := client.listCreateRequest(ctx, location, options)
	if err != nil {
		return ServiceTagsListResultResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ServiceTagsListResultResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ServiceTagsListResultResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *ServiceTagsClient) listCreateRequest(ctx context.Context, location string, options *ServiceTagsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/serviceTags"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
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
func (client *ServiceTagsClient) listHandleResponse(resp *azcore.Response) (ServiceTagsListResultResponse, error) {
	var val *ServiceTagsListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ServiceTagsListResultResponse{}, err
	}
return ServiceTagsListResultResponse{RawResponse: resp.Response, ServiceTagsListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *ServiceTagsClient) listHandleError(resp *azcore.Response) error {
var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

