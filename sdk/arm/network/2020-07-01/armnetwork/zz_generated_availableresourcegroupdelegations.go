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

// AvailableResourceGroupDelegationsClient contains the methods for the AvailableResourceGroupDelegations group.
// Don't use this type directly, use NewAvailableResourceGroupDelegationsClient() instead.
type AvailableResourceGroupDelegationsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewAvailableResourceGroupDelegationsClient creates a new instance of AvailableResourceGroupDelegationsClient with the specified values.
func NewAvailableResourceGroupDelegationsClient(con *armcore.Connection, subscriptionID string) *AvailableResourceGroupDelegationsClient {
	return &AvailableResourceGroupDelegationsClient{con: con, subscriptionID: subscriptionID}
}

// List - Gets all of the available subnet delegations for this resource group in this region.
func (client *AvailableResourceGroupDelegationsClient) List(location string, resourceGroupName string, options *AvailableResourceGroupDelegationsListOptions) AvailableDelegationsResultPager {
	return &availableDelegationsResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, location, resourceGroupName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp AvailableDelegationsResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.AvailableDelegationsResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *AvailableResourceGroupDelegationsClient) listCreateRequest(ctx context.Context, location string, resourceGroupName string, options *AvailableResourceGroupDelegationsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/locations/{location}/availableDelegations"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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
func (client *AvailableResourceGroupDelegationsClient) listHandleResponse(resp *azcore.Response) (AvailableDelegationsResultResponse, error) {
	var val *AvailableDelegationsResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return AvailableDelegationsResultResponse{}, err
	}
	return AvailableDelegationsResultResponse{RawResponse: resp.Response, AvailableDelegationsResult: val}, nil
}

// listHandleError handles the List error response.
func (client *AvailableResourceGroupDelegationsClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
