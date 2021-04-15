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

// InboundSecurityRuleClient contains the methods for the InboundSecurityRule group.
// Don't use this type directly, use NewInboundSecurityRuleClient() instead.
type InboundSecurityRuleClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewInboundSecurityRuleClient creates a new instance of InboundSecurityRuleClient with the specified values.
func NewInboundSecurityRuleClient(con *armcore.Connection, subscriptionID string) *InboundSecurityRuleClient {
	return &InboundSecurityRuleClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates the specified Network Virtual Appliance Inbound Security Rules.
func (client *InboundSecurityRuleClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, ruleCollectionName string, parameters InboundSecurityRule, options *InboundSecurityRuleBeginCreateOrUpdateOptions) (InboundSecurityRulePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, networkVirtualApplianceName, ruleCollectionName, parameters, options)
	if err != nil {
		return InboundSecurityRulePollerResponse{}, err
	}
	result := InboundSecurityRulePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("InboundSecurityRuleClient.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return InboundSecurityRulePollerResponse{}, err
	}
	poller := &inboundSecurityRulePoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (InboundSecurityRuleResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new InboundSecurityRulePoller from the specified resume token.
// token - The value must come from a previous call to InboundSecurityRulePoller.ResumeToken().
func (client *InboundSecurityRuleClient) ResumeCreateOrUpdate(token string) (InboundSecurityRulePoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("InboundSecurityRuleClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &inboundSecurityRulePoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates or updates the specified Network Virtual Appliance Inbound Security Rules.
func (client *InboundSecurityRuleClient) createOrUpdate(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, ruleCollectionName string, parameters InboundSecurityRule, options *InboundSecurityRuleBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, networkVirtualApplianceName, ruleCollectionName, parameters, options)
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
func (client *InboundSecurityRuleClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, ruleCollectionName string, parameters InboundSecurityRule, options *InboundSecurityRuleBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances/{networkVirtualApplianceName}/inboundSecurityRules/{ruleCollectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkVirtualApplianceName == "" {
		return nil, errors.New("parameter networkVirtualApplianceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkVirtualApplianceName}", url.PathEscape(networkVirtualApplianceName))
	if ruleCollectionName == "" {
		return nil, errors.New("parameter ruleCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleCollectionName}", url.PathEscape(ruleCollectionName))
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
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *InboundSecurityRuleClient) createOrUpdateHandleResponse(resp *azcore.Response) (InboundSecurityRuleResponse, error) {
	var val *InboundSecurityRule
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return InboundSecurityRuleResponse{}, err
	}
	return InboundSecurityRuleResponse{RawResponse: resp.Response, InboundSecurityRule: val}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *InboundSecurityRuleClient) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
