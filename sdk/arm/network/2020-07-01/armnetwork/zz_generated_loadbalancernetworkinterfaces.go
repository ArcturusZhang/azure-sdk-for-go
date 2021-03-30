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

// LoadBalancerNetworkInterfacesClient contains the methods for the LoadBalancerNetworkInterfaces group.
// Don't use this type directly, use NewLoadBalancerNetworkInterfacesClient() instead.
type LoadBalancerNetworkInterfacesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewLoadBalancerNetworkInterfacesClient creates a new instance of LoadBalancerNetworkInterfacesClient with the specified values.
func NewLoadBalancerNetworkInterfacesClient(con *armcore.Connection, subscriptionID string) *LoadBalancerNetworkInterfacesClient {
	return &LoadBalancerNetworkInterfacesClient{con: con, subscriptionID: subscriptionID}
}

// List - Gets associated load balancer network interfaces.
func (client *LoadBalancerNetworkInterfacesClient) List(resourceGroupName string, loadBalancerName string, options *LoadBalancerNetworkInterfacesListOptions) NetworkInterfaceListResultPager {
	return &networkInterfaceListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, loadBalancerName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp NetworkInterfaceListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.NetworkInterfaceListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *LoadBalancerNetworkInterfacesClient) listCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, options *LoadBalancerNetworkInterfacesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/networkInterfaces"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
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
func (client *LoadBalancerNetworkInterfacesClient) listHandleResponse(resp *azcore.Response) (NetworkInterfaceListResultResponse, error) {
	var val *NetworkInterfaceListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return NetworkInterfaceListResultResponse{}, err
	}
	return NetworkInterfaceListResultResponse{RawResponse: resp.Response, NetworkInterfaceListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *LoadBalancerNetworkInterfacesClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
