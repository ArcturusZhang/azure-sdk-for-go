package customproviders

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

// CustomResourceProviderClient is the allows extension of ARM control plane with custom resource providers.
type CustomResourceProviderClient struct {
    BaseClient
}
// NewCustomResourceProviderClient creates an instance of the CustomResourceProviderClient client.
func NewCustomResourceProviderClient(subscriptionID string) CustomResourceProviderClient {
    return NewCustomResourceProviderClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewCustomResourceProviderClientWithBaseURI creates an instance of the CustomResourceProviderClient client.
    func NewCustomResourceProviderClientWithBaseURI(baseURI string, subscriptionID string) CustomResourceProviderClient {
        return CustomResourceProviderClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CreateOrUpdate creates or updates the custom resource provider.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // resourceProviderName - the name of the resource provider.
        // resourceProvider - the parameters required to create or update a custom resource provider definition.
func (client CustomResourceProviderClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceProviderName string, resourceProvider CustomRPManifest) (result CustomResourceProviderCreateOrUpdateFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/CustomResourceProviderClient.CreateOrUpdate")
        defer func() {
            sc := -1
            if result.Response() != nil {
                sc = result.Response().StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: resourceProviderName,
             Constraints: []validation.Constraint{	{Target: "resourceProviderName", Name: validation.MaxLength, Rule: 64, Chain: nil },
            	{Target: "resourceProviderName", Name: validation.MinLength, Rule: 3, Chain: nil }}}}); err != nil {
            return result, validation.NewError("customproviders.CustomResourceProviderClient", "CreateOrUpdate", err.Error())
            }

                req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, resourceProviderName, resourceProvider)
    if err != nil {
    err = autorest.NewErrorWithError(err, "customproviders.CustomResourceProviderClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

            result, err = client.CreateOrUpdateSender(req)
            if err != nil {
            err = autorest.NewErrorWithError(err, "customproviders.CustomResourceProviderClient", "CreateOrUpdate", result.Response(), "Failure sending request")
            return
            }

    return
    }

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client CustomResourceProviderClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, resourceProviderName string, resourceProvider CustomRPManifest) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "resourceProviderName": autorest.Encode("path",resourceProviderName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-09-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPut(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders/{resourceProviderName}",pathParameters),
    autorest.WithJSON(resourceProvider),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client CustomResourceProviderClient) CreateOrUpdateSender(req *http.Request) (future CustomResourceProviderCreateOrUpdateFuture, err error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            var resp *http.Response
            resp, err = autorest.SendWithSender(client, req, sd...)
            if err != nil {
            return
            }
            future.Future, err = azure.NewFutureFromResponse(resp)
            return
            }

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client CustomResourceProviderClient) CreateOrUpdateResponder(resp *http.Response) (result CustomRPManifest, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Delete deletes the custom resource provider.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // resourceProviderName - the name of the resource provider.
func (client CustomResourceProviderClient) Delete(ctx context.Context, resourceGroupName string, resourceProviderName string) (result CustomResourceProviderDeleteFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/CustomResourceProviderClient.Delete")
        defer func() {
            sc := -1
            if result.Response() != nil {
                sc = result.Response().StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: resourceProviderName,
             Constraints: []validation.Constraint{	{Target: "resourceProviderName", Name: validation.MaxLength, Rule: 64, Chain: nil },
            	{Target: "resourceProviderName", Name: validation.MinLength, Rule: 3, Chain: nil }}}}); err != nil {
            return result, validation.NewError("customproviders.CustomResourceProviderClient", "Delete", err.Error())
            }

                req, err := client.DeletePreparer(ctx, resourceGroupName, resourceProviderName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "customproviders.CustomResourceProviderClient", "Delete", nil , "Failure preparing request")
    return
    }

            result, err = client.DeleteSender(req)
            if err != nil {
            err = autorest.NewErrorWithError(err, "customproviders.CustomResourceProviderClient", "Delete", result.Response(), "Failure sending request")
            return
            }

    return
    }

    // DeletePreparer prepares the Delete request.
    func (client CustomResourceProviderClient) DeletePreparer(ctx context.Context, resourceGroupName string, resourceProviderName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "resourceProviderName": autorest.Encode("path",resourceProviderName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-09-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsDelete(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders/{resourceProviderName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client CustomResourceProviderClient) DeleteSender(req *http.Request) (future CustomResourceProviderDeleteFuture, err error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            var resp *http.Response
            resp, err = autorest.SendWithSender(client, req, sd...)
            if err != nil {
            return
            }
            future.Future, err = azure.NewFutureFromResponse(resp)
            return
            }

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client CustomResourceProviderClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Get gets the custom resource provider manifest.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // resourceProviderName - the name of the resource provider.
func (client CustomResourceProviderClient) Get(ctx context.Context, resourceGroupName string, resourceProviderName string) (result CustomRPManifest, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/CustomResourceProviderClient.Get")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: resourceProviderName,
             Constraints: []validation.Constraint{	{Target: "resourceProviderName", Name: validation.MaxLength, Rule: 64, Chain: nil },
            	{Target: "resourceProviderName", Name: validation.MinLength, Rule: 3, Chain: nil }}}}); err != nil {
            return result, validation.NewError("customproviders.CustomResourceProviderClient", "Get", err.Error())
            }

                req, err := client.GetPreparer(ctx, resourceGroupName, resourceProviderName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "customproviders.CustomResourceProviderClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "customproviders.CustomResourceProviderClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "customproviders.CustomResourceProviderClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client CustomResourceProviderClient) GetPreparer(ctx context.Context, resourceGroupName string, resourceProviderName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "resourceProviderName": autorest.Encode("path",resourceProviderName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-09-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders/{resourceProviderName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client CustomResourceProviderClient) GetSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client CustomResourceProviderClient) GetResponder(resp *http.Response) (result CustomRPManifest, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// ListByResourceGroup gets all the custom resource providers within a resource group.
    // Parameters:
        // resourceGroupName - the name of the resource group.
func (client CustomResourceProviderClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result ListByCustomRPManifestPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/CustomResourceProviderClient.ListByResourceGroup")
        defer func() {
            sc := -1
            if result.lbcrm.Response.Response != nil {
                sc = result.lbcrm.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
                result.fn = client.listByResourceGroupNextResults
    req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "customproviders.CustomResourceProviderClient", "ListByResourceGroup", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListByResourceGroupSender(req)
            if err != nil {
            result.lbcrm.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "customproviders.CustomResourceProviderClient", "ListByResourceGroup", resp, "Failure sending request")
            return
            }

            result.lbcrm, err = client.ListByResourceGroupResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "customproviders.CustomResourceProviderClient", "ListByResourceGroup", resp, "Failure responding to request")
            }

    return
    }

    // ListByResourceGroupPreparer prepares the ListByResourceGroup request.
    func (client CustomResourceProviderClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-09-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
    // http.Response Body if it receives an error.
    func (client CustomResourceProviderClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client CustomResourceProviderClient) ListByResourceGroupResponder(resp *http.Response) (result ListByCustomRPManifest, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

            // listByResourceGroupNextResults retrieves the next set of results, if any.
            func (client CustomResourceProviderClient) listByResourceGroupNextResults(ctx context.Context, lastResults ListByCustomRPManifest) (result ListByCustomRPManifest, err error) {
            req, err := lastResults.listByCustomRPManifestPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "customproviders.CustomResourceProviderClient", "listByResourceGroupNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListByResourceGroupSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "customproviders.CustomResourceProviderClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListByResourceGroupResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "customproviders.CustomResourceProviderClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
    func (client CustomResourceProviderClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result ListByCustomRPManifestIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/CustomResourceProviderClient.ListByResourceGroup")
            defer func() {
                sc := -1
                if result.Response().Response.Response != nil {
                    sc = result.page.Response().Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
     }
        result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
                return
        }

// ListBySubscription gets all the custom resource providers within a subscription.
func (client CustomResourceProviderClient) ListBySubscription(ctx context.Context) (result ListByCustomRPManifestPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/CustomResourceProviderClient.ListBySubscription")
        defer func() {
            sc := -1
            if result.lbcrm.Response.Response != nil {
                sc = result.lbcrm.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
                result.fn = client.listBySubscriptionNextResults
    req, err := client.ListBySubscriptionPreparer(ctx)
    if err != nil {
    err = autorest.NewErrorWithError(err, "customproviders.CustomResourceProviderClient", "ListBySubscription", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListBySubscriptionSender(req)
            if err != nil {
            result.lbcrm.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "customproviders.CustomResourceProviderClient", "ListBySubscription", resp, "Failure sending request")
            return
            }

            result.lbcrm, err = client.ListBySubscriptionResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "customproviders.CustomResourceProviderClient", "ListBySubscription", resp, "Failure responding to request")
            }

    return
    }

    // ListBySubscriptionPreparer prepares the ListBySubscription request.
    func (client CustomResourceProviderClient) ListBySubscriptionPreparer(ctx context.Context) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-09-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.CustomProviders/resourceProviders",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListBySubscriptionSender sends the ListBySubscription request. The method will close the
    // http.Response Body if it receives an error.
    func (client CustomResourceProviderClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client CustomResourceProviderClient) ListBySubscriptionResponder(resp *http.Response) (result ListByCustomRPManifest, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

            // listBySubscriptionNextResults retrieves the next set of results, if any.
            func (client CustomResourceProviderClient) listBySubscriptionNextResults(ctx context.Context, lastResults ListByCustomRPManifest) (result ListByCustomRPManifest, err error) {
            req, err := lastResults.listByCustomRPManifestPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "customproviders.CustomResourceProviderClient", "listBySubscriptionNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListBySubscriptionSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "customproviders.CustomResourceProviderClient", "listBySubscriptionNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListBySubscriptionResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "customproviders.CustomResourceProviderClient", "listBySubscriptionNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListBySubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
    func (client CustomResourceProviderClient) ListBySubscriptionComplete(ctx context.Context) (result ListByCustomRPManifestIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/CustomResourceProviderClient.ListBySubscription")
            defer func() {
                sc := -1
                if result.Response().Response.Response != nil {
                    sc = result.page.Response().Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
     }
        result.page, err = client.ListBySubscription(ctx)
                return
        }

// Update updates an existing custom resource provider. The only value that can be updated via PATCH currently is the
// tags.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // resourceProviderName - the name of the resource provider.
        // patchableResource - the updatable fields of a custom resource provider.
func (client CustomResourceProviderClient) Update(ctx context.Context, resourceGroupName string, resourceProviderName string, patchableResource ResourceProvidersUpdate) (result CustomRPManifest, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/CustomResourceProviderClient.Update")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: resourceProviderName,
             Constraints: []validation.Constraint{	{Target: "resourceProviderName", Name: validation.MaxLength, Rule: 64, Chain: nil },
            	{Target: "resourceProviderName", Name: validation.MinLength, Rule: 3, Chain: nil }}}}); err != nil {
            return result, validation.NewError("customproviders.CustomResourceProviderClient", "Update", err.Error())
            }

                req, err := client.UpdatePreparer(ctx, resourceGroupName, resourceProviderName, patchableResource)
    if err != nil {
    err = autorest.NewErrorWithError(err, "customproviders.CustomResourceProviderClient", "Update", nil , "Failure preparing request")
    return
    }

            resp, err := client.UpdateSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "customproviders.CustomResourceProviderClient", "Update", resp, "Failure sending request")
            return
            }

            result, err = client.UpdateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "customproviders.CustomResourceProviderClient", "Update", resp, "Failure responding to request")
            }

    return
    }

    // UpdatePreparer prepares the Update request.
    func (client CustomResourceProviderClient) UpdatePreparer(ctx context.Context, resourceGroupName string, resourceProviderName string, patchableResource ResourceProvidersUpdate) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "resourceProviderName": autorest.Encode("path",resourceProviderName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-09-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPatch(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders/{resourceProviderName}",pathParameters),
    autorest.WithJSON(patchableResource),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // UpdateSender sends the Update request. The method will close the
    // http.Response Body if it receives an error.
    func (client CustomResourceProviderClient) UpdateSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client CustomResourceProviderClient) UpdateResponder(resp *http.Response) (result CustomRPManifest, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

