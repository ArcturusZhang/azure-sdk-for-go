// Package costmanagement implements the Azure ARM Costmanagement service API version 2018-08-01-preview.
//
//
package costmanagement

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
    "context"
    "github.com/Azure/go-autorest/autorest"
    "github.com/Azure/go-autorest/autorest/azure"
    "github.com/Azure/go-autorest/autorest/validation"
    "github.com/Azure/go-autorest/tracing"
    "net/http"
)

const (
// DefaultBaseURI is the default URI used for the service Costmanagement
DefaultBaseURI = "https://management.azure.com")

// BaseClient is the base client for Costmanagement.
type BaseClient struct {
    autorest.Client
    BaseURI string
            SubscriptionID string
}

// New creates an instance of the BaseClient client.
func New(subscriptionID string)BaseClient {
    return NewWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWithBaseURI creates an instance of the BaseClient client.
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
    return BaseClient{
        Client: autorest.NewClientWithUserAgent(UserAgent()),
        BaseURI: baseURI,
                SubscriptionID: subscriptionID,
    }
}

    // QueryBillingAccount lists the usage data for billing account.
        // Parameters:
            // billingAccountID - billingAccount ID
            // parameters - parameters supplied to the CreateOrUpdate Report operation.
    func (client BaseClient) QueryBillingAccount(ctx context.Context, billingAccountID string, parameters ReportDefinition) (result QueryResult, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/BaseClient.QueryBillingAccount")
            defer func() {
                sc := -1
                if result.Response.Response != nil {
                    sc = result.Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
        }
                if err := validation.Validate([]validation.Validation{
                { TargetValue: parameters,
                 Constraints: []validation.Constraint{	{Target: "parameters.Type", Name: validation.Null, Rule: true, Chain: nil },
                	{Target: "parameters.TimePeriod", Name: validation.Null, Rule: false ,
                Chain: []validation.Constraint{	{Target: "parameters.TimePeriod.From", Name: validation.Null, Rule: true, Chain: nil },
                	{Target: "parameters.TimePeriod.To", Name: validation.Null, Rule: true, Chain: nil },
                }},
                	{Target: "parameters.Dataset", Name: validation.Null, Rule: false ,
                Chain: []validation.Constraint{	{Target: "parameters.Dataset.Grouping", Name: validation.Null, Rule: false ,
                Chain: []validation.Constraint{	{Target: "parameters.Dataset.Grouping", Name: validation.MaxItems, Rule: 2, Chain: nil },
                }},
                	{Target: "parameters.Dataset.Filter", Name: validation.Null, Rule: false ,
                Chain: []validation.Constraint{	{Target: "parameters.Dataset.Filter.And", Name: validation.Null, Rule: false ,
                Chain: []validation.Constraint{	{Target: "parameters.Dataset.Filter.And", Name: validation.MinItems, Rule: 2, Chain: nil },
                }},
                	{Target: "parameters.Dataset.Filter.Or", Name: validation.Null, Rule: false ,
                Chain: []validation.Constraint{	{Target: "parameters.Dataset.Filter.Or", Name: validation.MinItems, Rule: 2, Chain: nil },
                }},
                	{Target: "parameters.Dataset.Filter.Not", Name: validation.Null, Rule: false, Chain: nil },
                	{Target: "parameters.Dataset.Filter.Dimension", Name: validation.Null, Rule: false ,
                Chain: []validation.Constraint{	{Target: "parameters.Dataset.Filter.Dimension.Name", Name: validation.Null, Rule: true, Chain: nil },
                	{Target: "parameters.Dataset.Filter.Dimension.Operator", Name: validation.Null, Rule: true, Chain: nil },
                	{Target: "parameters.Dataset.Filter.Dimension.Values", Name: validation.Null, Rule: true ,
                Chain: []validation.Constraint{	{Target: "parameters.Dataset.Filter.Dimension.Values", Name: validation.MinItems, Rule: 1, Chain: nil },
                }},
                }},
                	{Target: "parameters.Dataset.Filter.Tag", Name: validation.Null, Rule: false ,
                Chain: []validation.Constraint{	{Target: "parameters.Dataset.Filter.Tag.Name", Name: validation.Null, Rule: true, Chain: nil },
                	{Target: "parameters.Dataset.Filter.Tag.Operator", Name: validation.Null, Rule: true, Chain: nil },
                	{Target: "parameters.Dataset.Filter.Tag.Values", Name: validation.Null, Rule: true ,
                Chain: []validation.Constraint{	{Target: "parameters.Dataset.Filter.Tag.Values", Name: validation.MinItems, Rule: 1, Chain: nil },
                }},
                }},
                }},
                }}}}}); err != nil {
                return result, validation.NewError("costmanagement.BaseClient", "QueryBillingAccount", err.Error())
                }

                    req, err := client.QueryBillingAccountPreparer(ctx, billingAccountID, parameters)
        if err != nil {
        err = autorest.NewErrorWithError(err, "costmanagement.BaseClient", "QueryBillingAccount", nil , "Failure preparing request")
        return
        }

                resp, err := client.QueryBillingAccountSender(req)
                if err != nil {
                result.Response = autorest.Response{Response: resp}
                err = autorest.NewErrorWithError(err, "costmanagement.BaseClient", "QueryBillingAccount", resp, "Failure sending request")
                return
                }

                result, err = client.QueryBillingAccountResponder(resp)
                if err != nil {
                err = autorest.NewErrorWithError(err, "costmanagement.BaseClient", "QueryBillingAccount", resp, "Failure responding to request")
                }

        return
        }

        // QueryBillingAccountPreparer prepares the QueryBillingAccount request.
        func (client BaseClient) QueryBillingAccountPreparer(ctx context.Context, billingAccountID string, parameters ReportDefinition) (*http.Request, error) {
                pathParameters := map[string]interface{} {
                "billingAccountId": autorest.Encode("path",billingAccountID),
                }

                            const APIVersion = "2018-08-01-preview"
            queryParameters := map[string]interface{} {
            "api-version": APIVersion,
            }

            preparer := autorest.CreatePreparer(
        autorest.AsContentType("application/json; charset=utf-8"),
        autorest.AsPost(),
        autorest.WithBaseURL(client.BaseURI),
        autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/providers/Microsoft.CostManagement/Query",pathParameters),
        autorest.WithJSON(parameters),
        autorest.WithQueryParameters(queryParameters))
        return preparer.Prepare((&http.Request{}).WithContext(ctx))
        }

        // QueryBillingAccountSender sends the QueryBillingAccount request. The method will close the
        // http.Response Body if it receives an error.
        func (client BaseClient) QueryBillingAccountSender(req *http.Request) (*http.Response, error) {
            sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
                return autorest.SendWithSender(client, req, sd...)
                }

    // QueryBillingAccountResponder handles the response to the QueryBillingAccount request. The method always
    // closes the http.Response Body.
    func (client BaseClient) QueryBillingAccountResponder(resp *http.Response) (result QueryResult, err error) {
        err = autorest.Respond(
        resp,
        client.ByInspecting(),
        azure.WithErrorUnlessStatusCode(http.StatusOK),
        autorest.ByUnmarshallingJSON(&result),
        autorest.ByClosing())
        result.Response = autorest.Response{Response: resp}
            return
        }

    // QueryResourceGroup lists the usage data for subscriptionId and resource group.
        // Parameters:
            // resourceGroupName - azure Resource Group Name.
            // parameters - parameters supplied to the CreateOrUpdate Report operation.
    func (client BaseClient) QueryResourceGroup(ctx context.Context, resourceGroupName string, parameters ReportDefinition) (result QueryResult, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/BaseClient.QueryResourceGroup")
            defer func() {
                sc := -1
                if result.Response.Response != nil {
                    sc = result.Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
        }
                if err := validation.Validate([]validation.Validation{
                { TargetValue: parameters,
                 Constraints: []validation.Constraint{	{Target: "parameters.Type", Name: validation.Null, Rule: true, Chain: nil },
                	{Target: "parameters.TimePeriod", Name: validation.Null, Rule: false ,
                Chain: []validation.Constraint{	{Target: "parameters.TimePeriod.From", Name: validation.Null, Rule: true, Chain: nil },
                	{Target: "parameters.TimePeriod.To", Name: validation.Null, Rule: true, Chain: nil },
                }},
                	{Target: "parameters.Dataset", Name: validation.Null, Rule: false ,
                Chain: []validation.Constraint{	{Target: "parameters.Dataset.Grouping", Name: validation.Null, Rule: false ,
                Chain: []validation.Constraint{	{Target: "parameters.Dataset.Grouping", Name: validation.MaxItems, Rule: 2, Chain: nil },
                }},
                	{Target: "parameters.Dataset.Filter", Name: validation.Null, Rule: false ,
                Chain: []validation.Constraint{	{Target: "parameters.Dataset.Filter.And", Name: validation.Null, Rule: false ,
                Chain: []validation.Constraint{	{Target: "parameters.Dataset.Filter.And", Name: validation.MinItems, Rule: 2, Chain: nil },
                }},
                	{Target: "parameters.Dataset.Filter.Or", Name: validation.Null, Rule: false ,
                Chain: []validation.Constraint{	{Target: "parameters.Dataset.Filter.Or", Name: validation.MinItems, Rule: 2, Chain: nil },
                }},
                	{Target: "parameters.Dataset.Filter.Not", Name: validation.Null, Rule: false, Chain: nil },
                	{Target: "parameters.Dataset.Filter.Dimension", Name: validation.Null, Rule: false ,
                Chain: []validation.Constraint{	{Target: "parameters.Dataset.Filter.Dimension.Name", Name: validation.Null, Rule: true, Chain: nil },
                	{Target: "parameters.Dataset.Filter.Dimension.Operator", Name: validation.Null, Rule: true, Chain: nil },
                	{Target: "parameters.Dataset.Filter.Dimension.Values", Name: validation.Null, Rule: true ,
                Chain: []validation.Constraint{	{Target: "parameters.Dataset.Filter.Dimension.Values", Name: validation.MinItems, Rule: 1, Chain: nil },
                }},
                }},
                	{Target: "parameters.Dataset.Filter.Tag", Name: validation.Null, Rule: false ,
                Chain: []validation.Constraint{	{Target: "parameters.Dataset.Filter.Tag.Name", Name: validation.Null, Rule: true, Chain: nil },
                	{Target: "parameters.Dataset.Filter.Tag.Operator", Name: validation.Null, Rule: true, Chain: nil },
                	{Target: "parameters.Dataset.Filter.Tag.Values", Name: validation.Null, Rule: true ,
                Chain: []validation.Constraint{	{Target: "parameters.Dataset.Filter.Tag.Values", Name: validation.MinItems, Rule: 1, Chain: nil },
                }},
                }},
                }},
                }}}}}); err != nil {
                return result, validation.NewError("costmanagement.BaseClient", "QueryResourceGroup", err.Error())
                }

                    req, err := client.QueryResourceGroupPreparer(ctx, resourceGroupName, parameters)
        if err != nil {
        err = autorest.NewErrorWithError(err, "costmanagement.BaseClient", "QueryResourceGroup", nil , "Failure preparing request")
        return
        }

                resp, err := client.QueryResourceGroupSender(req)
                if err != nil {
                result.Response = autorest.Response{Response: resp}
                err = autorest.NewErrorWithError(err, "costmanagement.BaseClient", "QueryResourceGroup", resp, "Failure sending request")
                return
                }

                result, err = client.QueryResourceGroupResponder(resp)
                if err != nil {
                err = autorest.NewErrorWithError(err, "costmanagement.BaseClient", "QueryResourceGroup", resp, "Failure responding to request")
                }

        return
        }

        // QueryResourceGroupPreparer prepares the QueryResourceGroup request.
        func (client BaseClient) QueryResourceGroupPreparer(ctx context.Context, resourceGroupName string, parameters ReportDefinition) (*http.Request, error) {
                pathParameters := map[string]interface{} {
                "resourceGroupName": autorest.Encode("path",resourceGroupName),
                "subscriptionId": autorest.Encode("path",client.SubscriptionID),
                }

                            const APIVersion = "2018-08-01-preview"
            queryParameters := map[string]interface{} {
            "api-version": APIVersion,
            }

            preparer := autorest.CreatePreparer(
        autorest.AsContentType("application/json; charset=utf-8"),
        autorest.AsPost(),
        autorest.WithBaseURL(client.BaseURI),
        autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.CostManagement/Query",pathParameters),
        autorest.WithJSON(parameters),
        autorest.WithQueryParameters(queryParameters))
        return preparer.Prepare((&http.Request{}).WithContext(ctx))
        }

        // QueryResourceGroupSender sends the QueryResourceGroup request. The method will close the
        // http.Response Body if it receives an error.
        func (client BaseClient) QueryResourceGroupSender(req *http.Request) (*http.Response, error) {
            sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
                return autorest.SendWithSender(client, req, sd...)
                }

    // QueryResourceGroupResponder handles the response to the QueryResourceGroup request. The method always
    // closes the http.Response Body.
    func (client BaseClient) QueryResourceGroupResponder(resp *http.Response) (result QueryResult, err error) {
        err = autorest.Respond(
        resp,
        client.ByInspecting(),
        azure.WithErrorUnlessStatusCode(http.StatusOK),
        autorest.ByUnmarshallingJSON(&result),
        autorest.ByClosing())
        result.Response = autorest.Response{Response: resp}
            return
        }

    // QuerySubscription lists the usage data for subscriptionId.
        // Parameters:
            // parameters - parameters supplied to the CreateOrUpdate Report operation.
    func (client BaseClient) QuerySubscription(ctx context.Context, parameters ReportDefinition) (result QueryResult, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/BaseClient.QuerySubscription")
            defer func() {
                sc := -1
                if result.Response.Response != nil {
                    sc = result.Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
        }
                if err := validation.Validate([]validation.Validation{
                { TargetValue: parameters,
                 Constraints: []validation.Constraint{	{Target: "parameters.Type", Name: validation.Null, Rule: true, Chain: nil },
                	{Target: "parameters.TimePeriod", Name: validation.Null, Rule: false ,
                Chain: []validation.Constraint{	{Target: "parameters.TimePeriod.From", Name: validation.Null, Rule: true, Chain: nil },
                	{Target: "parameters.TimePeriod.To", Name: validation.Null, Rule: true, Chain: nil },
                }},
                	{Target: "parameters.Dataset", Name: validation.Null, Rule: false ,
                Chain: []validation.Constraint{	{Target: "parameters.Dataset.Grouping", Name: validation.Null, Rule: false ,
                Chain: []validation.Constraint{	{Target: "parameters.Dataset.Grouping", Name: validation.MaxItems, Rule: 2, Chain: nil },
                }},
                	{Target: "parameters.Dataset.Filter", Name: validation.Null, Rule: false ,
                Chain: []validation.Constraint{	{Target: "parameters.Dataset.Filter.And", Name: validation.Null, Rule: false ,
                Chain: []validation.Constraint{	{Target: "parameters.Dataset.Filter.And", Name: validation.MinItems, Rule: 2, Chain: nil },
                }},
                	{Target: "parameters.Dataset.Filter.Or", Name: validation.Null, Rule: false ,
                Chain: []validation.Constraint{	{Target: "parameters.Dataset.Filter.Or", Name: validation.MinItems, Rule: 2, Chain: nil },
                }},
                	{Target: "parameters.Dataset.Filter.Not", Name: validation.Null, Rule: false, Chain: nil },
                	{Target: "parameters.Dataset.Filter.Dimension", Name: validation.Null, Rule: false ,
                Chain: []validation.Constraint{	{Target: "parameters.Dataset.Filter.Dimension.Name", Name: validation.Null, Rule: true, Chain: nil },
                	{Target: "parameters.Dataset.Filter.Dimension.Operator", Name: validation.Null, Rule: true, Chain: nil },
                	{Target: "parameters.Dataset.Filter.Dimension.Values", Name: validation.Null, Rule: true ,
                Chain: []validation.Constraint{	{Target: "parameters.Dataset.Filter.Dimension.Values", Name: validation.MinItems, Rule: 1, Chain: nil },
                }},
                }},
                	{Target: "parameters.Dataset.Filter.Tag", Name: validation.Null, Rule: false ,
                Chain: []validation.Constraint{	{Target: "parameters.Dataset.Filter.Tag.Name", Name: validation.Null, Rule: true, Chain: nil },
                	{Target: "parameters.Dataset.Filter.Tag.Operator", Name: validation.Null, Rule: true, Chain: nil },
                	{Target: "parameters.Dataset.Filter.Tag.Values", Name: validation.Null, Rule: true ,
                Chain: []validation.Constraint{	{Target: "parameters.Dataset.Filter.Tag.Values", Name: validation.MinItems, Rule: 1, Chain: nil },
                }},
                }},
                }},
                }}}}}); err != nil {
                return result, validation.NewError("costmanagement.BaseClient", "QuerySubscription", err.Error())
                }

                    req, err := client.QuerySubscriptionPreparer(ctx, parameters)
        if err != nil {
        err = autorest.NewErrorWithError(err, "costmanagement.BaseClient", "QuerySubscription", nil , "Failure preparing request")
        return
        }

                resp, err := client.QuerySubscriptionSender(req)
                if err != nil {
                result.Response = autorest.Response{Response: resp}
                err = autorest.NewErrorWithError(err, "costmanagement.BaseClient", "QuerySubscription", resp, "Failure sending request")
                return
                }

                result, err = client.QuerySubscriptionResponder(resp)
                if err != nil {
                err = autorest.NewErrorWithError(err, "costmanagement.BaseClient", "QuerySubscription", resp, "Failure responding to request")
                }

        return
        }

        // QuerySubscriptionPreparer prepares the QuerySubscription request.
        func (client BaseClient) QuerySubscriptionPreparer(ctx context.Context, parameters ReportDefinition) (*http.Request, error) {
                pathParameters := map[string]interface{} {
                "subscriptionId": autorest.Encode("path",client.SubscriptionID),
                }

                            const APIVersion = "2018-08-01-preview"
            queryParameters := map[string]interface{} {
            "api-version": APIVersion,
            }

            preparer := autorest.CreatePreparer(
        autorest.AsContentType("application/json; charset=utf-8"),
        autorest.AsPost(),
        autorest.WithBaseURL(client.BaseURI),
        autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.CostManagement/Query",pathParameters),
        autorest.WithJSON(parameters),
        autorest.WithQueryParameters(queryParameters))
        return preparer.Prepare((&http.Request{}).WithContext(ctx))
        }

        // QuerySubscriptionSender sends the QuerySubscription request. The method will close the
        // http.Response Body if it receives an error.
        func (client BaseClient) QuerySubscriptionSender(req *http.Request) (*http.Response, error) {
            sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
                return autorest.SendWithSender(client, req, sd...)
                }

    // QuerySubscriptionResponder handles the response to the QuerySubscription request. The method always
    // closes the http.Response Body.
    func (client BaseClient) QuerySubscriptionResponder(resp *http.Response) (result QueryResult, err error) {
        err = autorest.Respond(
        resp,
        client.ByInspecting(),
        azure.WithErrorUnlessStatusCode(http.StatusOK),
        autorest.ByUnmarshallingJSON(&result),
        autorest.ByClosing())
        result.Response = autorest.Response{Response: resp}
            return
        }

