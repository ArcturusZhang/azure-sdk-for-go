package cognitiveservices

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

// AccountsClient is the cognitive Services Management Client
type AccountsClient struct {
    BaseClient
}
// NewAccountsClient creates an instance of the AccountsClient client.
func NewAccountsClient(subscriptionID string) AccountsClient {
    return NewAccountsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAccountsClientWithBaseURI creates an instance of the AccountsClient client.
    func NewAccountsClientWithBaseURI(baseURI string, subscriptionID string) AccountsClient {
        return AccountsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// Create create Cognitive Services Account. Accounts is a resource group wide resource type. It holds the keys for
// developer to access intelligent APIs. It's also the resource type for billing.
    // Parameters:
        // resourceGroupName - the name of the resource group within the user's subscription.
        // accountName - the name of the cognitive services account within the specified resource group. Cognitive
        // Services account names must be between 3 and 24 characters in length and use numbers and lower-case letters
        // only.
        // parameters - the parameters to provide for the created account.
func (client AccountsClient) Create(ctx context.Context, resourceGroupName string, accountName string, parameters AccountCreateParameters) (result Account, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/AccountsClient.Create")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: accountName,
             Constraints: []validation.Constraint{	{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil },
            	{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil },
            	{Target: "accountName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil }}},
            { TargetValue: parameters,
             Constraints: []validation.Constraint{	{Target: "parameters.Sku", Name: validation.Null, Rule: true, Chain: nil },
            	{Target: "parameters.Location", Name: validation.Null, Rule: true, Chain: nil },
            	{Target: "parameters.Properties", Name: validation.Null, Rule: true, Chain: nil }}}}); err != nil {
            return result, validation.NewError("cognitiveservices.AccountsClient", "Create", err.Error())
            }

                req, err := client.CreatePreparer(ctx, resourceGroupName, accountName, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "Create", nil , "Failure preparing request")
    return
    }

            resp, err := client.CreateSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "Create", resp, "Failure sending request")
            return
            }

            result, err = client.CreateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "Create", resp, "Failure responding to request")
            }

    return
    }

    // CreatePreparer prepares the Create request.
    func (client AccountsClient) CreatePreparer(ctx context.Context, resourceGroupName string, accountName string, parameters AccountCreateParameters) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "accountName": autorest.Encode("path",accountName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2016-02-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPut(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateSender sends the Create request. The method will close the
    // http.Response Body if it receives an error.
    func (client AccountsClient) CreateSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client AccountsClient) CreateResponder(resp *http.Response) (result Account, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Delete deletes a Cognitive Services account from the resource group.
    // Parameters:
        // resourceGroupName - the name of the resource group within the user's subscription.
        // accountName - the name of the cognitive services account within the specified resource group. Cognitive
        // Services account names must be between 3 and 24 characters in length and use numbers and lower-case letters
        // only.
func (client AccountsClient) Delete(ctx context.Context, resourceGroupName string, accountName string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/AccountsClient.Delete")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: accountName,
             Constraints: []validation.Constraint{	{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil },
            	{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil },
            	{Target: "accountName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("cognitiveservices.AccountsClient", "Delete", err.Error())
            }

                req, err := client.DeletePreparer(ctx, resourceGroupName, accountName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "Delete", nil , "Failure preparing request")
    return
    }

            resp, err := client.DeleteSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "Delete", resp, "Failure sending request")
            return
            }

            result, err = client.DeleteResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "Delete", resp, "Failure responding to request")
            }

    return
    }

    // DeletePreparer prepares the Delete request.
    func (client AccountsClient) DeletePreparer(ctx context.Context, resourceGroupName string, accountName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "accountName": autorest.Encode("path",accountName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2016-02-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsDelete(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client AccountsClient) DeleteSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client AccountsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// GetProperties returns a Cognitive Services account specified by the parameters.
    // Parameters:
        // resourceGroupName - the name of the resource group within the user's subscription.
        // accountName - the name of the cognitive services account within the specified resource group. Cognitive
        // Services account names must be between 3 and 24 characters in length and use numbers and lower-case letters
        // only.
func (client AccountsClient) GetProperties(ctx context.Context, resourceGroupName string, accountName string) (result Account, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/AccountsClient.GetProperties")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: accountName,
             Constraints: []validation.Constraint{	{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil },
            	{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil },
            	{Target: "accountName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("cognitiveservices.AccountsClient", "GetProperties", err.Error())
            }

                req, err := client.GetPropertiesPreparer(ctx, resourceGroupName, accountName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "GetProperties", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetPropertiesSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "GetProperties", resp, "Failure sending request")
            return
            }

            result, err = client.GetPropertiesResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "GetProperties", resp, "Failure responding to request")
            }

    return
    }

    // GetPropertiesPreparer prepares the GetProperties request.
    func (client AccountsClient) GetPropertiesPreparer(ctx context.Context, resourceGroupName string, accountName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "accountName": autorest.Encode("path",accountName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2016-02-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetPropertiesSender sends the GetProperties request. The method will close the
    // http.Response Body if it receives an error.
    func (client AccountsClient) GetPropertiesSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetPropertiesResponder handles the response to the GetProperties request. The method always
// closes the http.Response Body.
func (client AccountsClient) GetPropertiesResponder(resp *http.Response) (result Account, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// List returns all the resources of a particular type belonging to a subscription.
func (client AccountsClient) List(ctx context.Context) (result AccountListResult, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/AccountsClient.List")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.ListPreparer(ctx)
    if err != nil {
    err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "List", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "List", resp, "Failure sending request")
            return
            }

            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "List", resp, "Failure responding to request")
            }

    return
    }

    // ListPreparer prepares the List request.
    func (client AccountsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2016-02-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.CognitiveServices/accounts",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client AccountsClient) ListSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client AccountsClient) ListResponder(resp *http.Response) (result AccountListResult, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// ListByResourceGroup returns all the resources of a particular type belonging to a resource group
    // Parameters:
        // resourceGroupName - the name of the resource group within the user's subscription.
func (client AccountsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result AccountListResult, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/AccountsClient.ListByResourceGroup")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "ListByResourceGroup", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListByResourceGroupSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "ListByResourceGroup", resp, "Failure sending request")
            return
            }

            result, err = client.ListByResourceGroupResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "ListByResourceGroup", resp, "Failure responding to request")
            }

    return
    }

    // ListByResourceGroupPreparer prepares the ListByResourceGroup request.
    func (client AccountsClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2016-02-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
    // http.Response Body if it receives an error.
    func (client AccountsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client AccountsClient) ListByResourceGroupResponder(resp *http.Response) (result AccountListResult, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// ListKeys lists the account keys for the specified Cognitive Services account.
    // Parameters:
        // resourceGroupName - the name of the resource group within the user's subscription.
        // accountName - the name of the cognitive services account within the specified resource group. Cognitive
        // Services account names must be between 3 and 24 characters in length and use numbers and lower-case letters
        // only.
func (client AccountsClient) ListKeys(ctx context.Context, resourceGroupName string, accountName string) (result AccountKeys, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/AccountsClient.ListKeys")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: accountName,
             Constraints: []validation.Constraint{	{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil },
            	{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil },
            	{Target: "accountName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("cognitiveservices.AccountsClient", "ListKeys", err.Error())
            }

                req, err := client.ListKeysPreparer(ctx, resourceGroupName, accountName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "ListKeys", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListKeysSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "ListKeys", resp, "Failure sending request")
            return
            }

            result, err = client.ListKeysResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "ListKeys", resp, "Failure responding to request")
            }

    return
    }

    // ListKeysPreparer prepares the ListKeys request.
    func (client AccountsClient) ListKeysPreparer(ctx context.Context, resourceGroupName string, accountName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "accountName": autorest.Encode("path",accountName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2016-02-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}/listKeys",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListKeysSender sends the ListKeys request. The method will close the
    // http.Response Body if it receives an error.
    func (client AccountsClient) ListKeysSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListKeysResponder handles the response to the ListKeys request. The method always
// closes the http.Response Body.
func (client AccountsClient) ListKeysResponder(resp *http.Response) (result AccountKeys, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// ListSkus list available SKUs for the requested Cognitive Services account
    // Parameters:
        // resourceGroupName - the name of the resource group within the user's subscription.
        // accountName - the name of the cognitive services account within the specified resource group. Cognitive
        // Services account names must be between 3 and 24 characters in length and use numbers and lower-case letters
        // only.
func (client AccountsClient) ListSkus(ctx context.Context, resourceGroupName string, accountName string) (result AccountEnumerateSkusResult, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/AccountsClient.ListSkus")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: accountName,
             Constraints: []validation.Constraint{	{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil },
            	{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil },
            	{Target: "accountName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("cognitiveservices.AccountsClient", "ListSkus", err.Error())
            }

                req, err := client.ListSkusPreparer(ctx, resourceGroupName, accountName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "ListSkus", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListSkusSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "ListSkus", resp, "Failure sending request")
            return
            }

            result, err = client.ListSkusResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "ListSkus", resp, "Failure responding to request")
            }

    return
    }

    // ListSkusPreparer prepares the ListSkus request.
    func (client AccountsClient) ListSkusPreparer(ctx context.Context, resourceGroupName string, accountName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "accountName": autorest.Encode("path",accountName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2016-02-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}/skus",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSkusSender sends the ListSkus request. The method will close the
    // http.Response Body if it receives an error.
    func (client AccountsClient) ListSkusSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListSkusResponder handles the response to the ListSkus request. The method always
// closes the http.Response Body.
func (client AccountsClient) ListSkusResponder(resp *http.Response) (result AccountEnumerateSkusResult, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// RegenerateKey regenerates the specified account key for the specified Cognitive Services account.
    // Parameters:
        // resourceGroupName - the name of the resource group within the user's subscription.
        // accountName - the name of the cognitive services account within the specified resource group. Cognitive
        // Services account names must be between 3 and 24 characters in length and use numbers and lower-case letters
        // only.
        // body - regenerate key parameters.
func (client AccountsClient) RegenerateKey(ctx context.Context, resourceGroupName string, accountName string, body RegenerateKeyParameters) (result AccountKeys, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/AccountsClient.RegenerateKey")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: accountName,
             Constraints: []validation.Constraint{	{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil },
            	{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil },
            	{Target: "accountName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("cognitiveservices.AccountsClient", "RegenerateKey", err.Error())
            }

                req, err := client.RegenerateKeyPreparer(ctx, resourceGroupName, accountName, body)
    if err != nil {
    err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "RegenerateKey", nil , "Failure preparing request")
    return
    }

            resp, err := client.RegenerateKeySender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "RegenerateKey", resp, "Failure sending request")
            return
            }

            result, err = client.RegenerateKeyResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "RegenerateKey", resp, "Failure responding to request")
            }

    return
    }

    // RegenerateKeyPreparer prepares the RegenerateKey request.
    func (client AccountsClient) RegenerateKeyPreparer(ctx context.Context, resourceGroupName string, accountName string, body RegenerateKeyParameters) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "accountName": autorest.Encode("path",accountName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2016-02-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}/regenerateKey",pathParameters),
    autorest.WithJSON(body),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // RegenerateKeySender sends the RegenerateKey request. The method will close the
    // http.Response Body if it receives an error.
    func (client AccountsClient) RegenerateKeySender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// RegenerateKeyResponder handles the response to the RegenerateKey request. The method always
// closes the http.Response Body.
func (client AccountsClient) RegenerateKeyResponder(resp *http.Response) (result AccountKeys, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Update updates a Cognitive Services account
    // Parameters:
        // resourceGroupName - the name of the resource group within the user's subscription.
        // accountName - the name of the cognitive services account within the specified resource group. Cognitive
        // Services account names must be between 3 and 24 characters in length and use numbers and lower-case letters
        // only.
        // body - the parameters to provide for the created account.
func (client AccountsClient) Update(ctx context.Context, resourceGroupName string, accountName string, body AccountUpdateParameters) (result Account, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/AccountsClient.Update")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: accountName,
             Constraints: []validation.Constraint{	{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil },
            	{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil },
            	{Target: "accountName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("cognitiveservices.AccountsClient", "Update", err.Error())
            }

                req, err := client.UpdatePreparer(ctx, resourceGroupName, accountName, body)
    if err != nil {
    err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "Update", nil , "Failure preparing request")
    return
    }

            resp, err := client.UpdateSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "Update", resp, "Failure sending request")
            return
            }

            result, err = client.UpdateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "Update", resp, "Failure responding to request")
            }

    return
    }

    // UpdatePreparer prepares the Update request.
    func (client AccountsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, accountName string, body AccountUpdateParameters) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "accountName": autorest.Encode("path",accountName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2016-02-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPatch(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}",pathParameters),
    autorest.WithJSON(body),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // UpdateSender sends the Update request. The method will close the
    // http.Response Body if it receives an error.
    func (client AccountsClient) UpdateSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client AccountsClient) UpdateResponder(resp *http.Response) (result Account, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

