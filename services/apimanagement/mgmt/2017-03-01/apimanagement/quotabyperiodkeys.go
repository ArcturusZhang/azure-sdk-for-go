package apimanagement

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
    "github.com/Azure/go-autorest/autorest"
    "github.com/Azure/go-autorest/autorest/azure"
    "net/http"
    "context"
    "github.com/Azure/go-autorest/tracing"
    "github.com/Azure/go-autorest/autorest/validation"
)

// QuotaByPeriodKeysClient is the apiManagement Client
type QuotaByPeriodKeysClient struct {
    BaseClient
}
// NewQuotaByPeriodKeysClient creates an instance of the QuotaByPeriodKeysClient client.
func NewQuotaByPeriodKeysClient(subscriptionID string) QuotaByPeriodKeysClient {
    return NewQuotaByPeriodKeysClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewQuotaByPeriodKeysClientWithBaseURI creates an instance of the QuotaByPeriodKeysClient client.
    func NewQuotaByPeriodKeysClientWithBaseURI(baseURI string, subscriptionID string) QuotaByPeriodKeysClient {
        return QuotaByPeriodKeysClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// Get gets the value of the quota counter associated with the counter-key in the policy for the specific period in
// service instance.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // serviceName - the name of the API Management service.
        // quotaCounterKey - quota counter key identifier.This is the result of expression defined in counter-key
        // attribute of the quota-by-key policy.For Example, if you specify counter-key="boo" in the policy, then it’s
        // accessible by "boo" counter key. But if it’s defined as counter-key="@("b"+"a")" then it will be accessible
        // by "ba" key
        // quotaPeriodKey - quota period key identifier.
func (client QuotaByPeriodKeysClient) Get(ctx context.Context, resourceGroupName string, serviceName string, quotaCounterKey string, quotaPeriodKey string) (result QuotaCounterContract, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/QuotaByPeriodKeysClient.Get")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: serviceName,
             Constraints: []validation.Constraint{	{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil },
            	{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("apimanagement.QuotaByPeriodKeysClient", "Get", err.Error())
            }

                req, err := client.GetPreparer(ctx, resourceGroupName, serviceName, quotaCounterKey, quotaPeriodKey)
    if err != nil {
    err = autorest.NewErrorWithError(err, "apimanagement.QuotaByPeriodKeysClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "apimanagement.QuotaByPeriodKeysClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "apimanagement.QuotaByPeriodKeysClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client QuotaByPeriodKeysClient) GetPreparer(ctx context.Context, resourceGroupName string, serviceName string, quotaCounterKey string, quotaPeriodKey string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "quotaCounterKey": autorest.Encode("path",quotaCounterKey),
            "quotaPeriodKey": autorest.Encode("path",quotaPeriodKey),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "serviceName": autorest.Encode("path",serviceName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2017-03-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/quotas/{quotaCounterKey}/{quotaPeriodKey}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client QuotaByPeriodKeysClient) GetSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client QuotaByPeriodKeysClient) GetResponder(resp *http.Response) (result QuotaCounterContract, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Update updates an existing quota counter value in the specified service instance.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // serviceName - the name of the API Management service.
        // quotaCounterKey - quota counter key identifier.This is the result of expression defined in counter-key
        // attribute of the quota-by-key policy.For Example, if you specify counter-key="boo" in the policy, then it’s
        // accessible by "boo" counter key. But if it’s defined as counter-key="@("b"+"a")" then it will be accessible
        // by "ba" key
        // quotaPeriodKey - quota period key identifier.
        // parameters - the value of the Quota counter to be applied on the specified period.
func (client QuotaByPeriodKeysClient) Update(ctx context.Context, resourceGroupName string, serviceName string, quotaCounterKey string, quotaPeriodKey string, parameters QuotaCounterValueContractProperties) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/QuotaByPeriodKeysClient.Update")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: serviceName,
             Constraints: []validation.Constraint{	{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil },
            	{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("apimanagement.QuotaByPeriodKeysClient", "Update", err.Error())
            }

                req, err := client.UpdatePreparer(ctx, resourceGroupName, serviceName, quotaCounterKey, quotaPeriodKey, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "apimanagement.QuotaByPeriodKeysClient", "Update", nil , "Failure preparing request")
    return
    }

            resp, err := client.UpdateSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "apimanagement.QuotaByPeriodKeysClient", "Update", resp, "Failure sending request")
            return
            }

            result, err = client.UpdateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "apimanagement.QuotaByPeriodKeysClient", "Update", resp, "Failure responding to request")
            }

    return
    }

    // UpdatePreparer prepares the Update request.
    func (client QuotaByPeriodKeysClient) UpdatePreparer(ctx context.Context, resourceGroupName string, serviceName string, quotaCounterKey string, quotaPeriodKey string, parameters QuotaCounterValueContractProperties) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "quotaCounterKey": autorest.Encode("path",quotaCounterKey),
            "quotaPeriodKey": autorest.Encode("path",quotaPeriodKey),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "serviceName": autorest.Encode("path",serviceName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2017-03-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPatch(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/quotas/{quotaCounterKey}/{quotaPeriodKey}",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // UpdateSender sends the Update request. The method will close the
    // http.Response Body if it receives an error.
    func (client QuotaByPeriodKeysClient) UpdateSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client QuotaByPeriodKeysClient) UpdateResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

