package media

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

// ServiceClient is the media Services resource management APIs.
type ServiceClient struct {
    BaseClient
}
// NewServiceClient creates an instance of the ServiceClient client.
func NewServiceClient(subscriptionID string) ServiceClient {
    return NewServiceClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewServiceClientWithBaseURI creates an instance of the ServiceClient client.
    func NewServiceClientWithBaseURI(baseURI string, subscriptionID string) ServiceClient {
        return ServiceClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CheckNameAvailability checks whether the Media Service resource name is available. The name must be globally unique.
    // Parameters:
        // parameters - properties needed to check the availability of a name.
func (client ServiceClient) CheckNameAvailability(ctx context.Context, parameters CheckNameAvailabilityInput) (result CheckNameAvailabilityOutput, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ServiceClient.CheckNameAvailability")
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
             Constraints: []validation.Constraint{	{Target: "parameters.Name", Name: validation.Null, Rule: true ,
            Chain: []validation.Constraint{	{Target: "parameters.Name", Name: validation.MaxLength, Rule: 24, Chain: nil },
            	{Target: "parameters.Name", Name: validation.MinLength, Rule: 3, Chain: nil },
            	{Target: "parameters.Name", Name: validation.Pattern, Rule: `^[a-z0-9]{3,24}$`, Chain: nil },
            }},
            	{Target: "parameters.Type", Name: validation.Null, Rule: true, Chain: nil }}}}); err != nil {
            return result, validation.NewError("media.ServiceClient", "CheckNameAvailability", err.Error())
            }

                req, err := client.CheckNameAvailabilityPreparer(ctx, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "media.ServiceClient", "CheckNameAvailability", nil , "Failure preparing request")
    return
    }

            resp, err := client.CheckNameAvailabilitySender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "media.ServiceClient", "CheckNameAvailability", resp, "Failure sending request")
            return
            }

            result, err = client.CheckNameAvailabilityResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "media.ServiceClient", "CheckNameAvailability", resp, "Failure responding to request")
            }

    return
    }

    // CheckNameAvailabilityPreparer prepares the CheckNameAvailability request.
    func (client ServiceClient) CheckNameAvailabilityPreparer(ctx context.Context, parameters CheckNameAvailabilityInput) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2015-10-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Media/CheckNameAvailability",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CheckNameAvailabilitySender sends the CheckNameAvailability request. The method will close the
    // http.Response Body if it receives an error.
    func (client ServiceClient) CheckNameAvailabilitySender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// CheckNameAvailabilityResponder handles the response to the CheckNameAvailability request. The method always
// closes the http.Response Body.
func (client ServiceClient) CheckNameAvailabilityResponder(resp *http.Response) (result CheckNameAvailabilityOutput, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Create creates a Media Service.
    // Parameters:
        // resourceGroupName - name of the resource group within the Azure subscription.
        // mediaServiceName - name of the Media Service.
        // parameters - media Service properties needed for creation.
func (client ServiceClient) Create(ctx context.Context, resourceGroupName string, mediaServiceName string, parameters Service) (result Service, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ServiceClient.Create")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: mediaServiceName,
             Constraints: []validation.Constraint{	{Target: "mediaServiceName", Name: validation.MaxLength, Rule: 24, Chain: nil },
            	{Target: "mediaServiceName", Name: validation.MinLength, Rule: 3, Chain: nil },
            	{Target: "mediaServiceName", Name: validation.Pattern, Rule: `^[a-z0-9]{3,24}$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("media.ServiceClient", "Create", err.Error())
            }

                req, err := client.CreatePreparer(ctx, resourceGroupName, mediaServiceName, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "media.ServiceClient", "Create", nil , "Failure preparing request")
    return
    }

            resp, err := client.CreateSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "media.ServiceClient", "Create", resp, "Failure sending request")
            return
            }

            result, err = client.CreateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "media.ServiceClient", "Create", resp, "Failure responding to request")
            }

    return
    }

    // CreatePreparer prepares the Create request.
    func (client ServiceClient) CreatePreparer(ctx context.Context, resourceGroupName string, mediaServiceName string, parameters Service) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "mediaServiceName": autorest.Encode("path",mediaServiceName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2015-10-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPut(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{mediaServiceName}",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateSender sends the Create request. The method will close the
    // http.Response Body if it receives an error.
    func (client ServiceClient) CreateSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ServiceClient) CreateResponder(resp *http.Response) (result Service, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Delete deletes a Media Service.
    // Parameters:
        // resourceGroupName - name of the resource group within the Azure subscription.
        // mediaServiceName - name of the Media Service.
func (client ServiceClient) Delete(ctx context.Context, resourceGroupName string, mediaServiceName string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ServiceClient.Delete")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: mediaServiceName,
             Constraints: []validation.Constraint{	{Target: "mediaServiceName", Name: validation.MaxLength, Rule: 24, Chain: nil },
            	{Target: "mediaServiceName", Name: validation.MinLength, Rule: 3, Chain: nil },
            	{Target: "mediaServiceName", Name: validation.Pattern, Rule: `^[a-z0-9]{3,24}$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("media.ServiceClient", "Delete", err.Error())
            }

                req, err := client.DeletePreparer(ctx, resourceGroupName, mediaServiceName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "media.ServiceClient", "Delete", nil , "Failure preparing request")
    return
    }

            resp, err := client.DeleteSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "media.ServiceClient", "Delete", resp, "Failure sending request")
            return
            }

            result, err = client.DeleteResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "media.ServiceClient", "Delete", resp, "Failure responding to request")
            }

    return
    }

    // DeletePreparer prepares the Delete request.
    func (client ServiceClient) DeletePreparer(ctx context.Context, resourceGroupName string, mediaServiceName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "mediaServiceName": autorest.Encode("path",mediaServiceName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2015-10-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsDelete(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{mediaServiceName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client ServiceClient) DeleteSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ServiceClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Get gets a Media Service.
    // Parameters:
        // resourceGroupName - name of the resource group within the Azure subscription.
        // mediaServiceName - name of the Media Service.
func (client ServiceClient) Get(ctx context.Context, resourceGroupName string, mediaServiceName string) (result Service, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ServiceClient.Get")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: mediaServiceName,
             Constraints: []validation.Constraint{	{Target: "mediaServiceName", Name: validation.MaxLength, Rule: 24, Chain: nil },
            	{Target: "mediaServiceName", Name: validation.MinLength, Rule: 3, Chain: nil },
            	{Target: "mediaServiceName", Name: validation.Pattern, Rule: `^[a-z0-9]{3,24}$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("media.ServiceClient", "Get", err.Error())
            }

                req, err := client.GetPreparer(ctx, resourceGroupName, mediaServiceName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "media.ServiceClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "media.ServiceClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "media.ServiceClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client ServiceClient) GetPreparer(ctx context.Context, resourceGroupName string, mediaServiceName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "mediaServiceName": autorest.Encode("path",mediaServiceName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2015-10-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{mediaServiceName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client ServiceClient) GetSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ServiceClient) GetResponder(resp *http.Response) (result Service, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// ListByResourceGroup lists all of the Media Services in a resource group.
    // Parameters:
        // resourceGroupName - name of the resource group within the Azure subscription.
func (client ServiceClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result ServiceCollection, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ServiceClient.ListByResourceGroup")
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
    err = autorest.NewErrorWithError(err, "media.ServiceClient", "ListByResourceGroup", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListByResourceGroupSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "media.ServiceClient", "ListByResourceGroup", resp, "Failure sending request")
            return
            }

            result, err = client.ListByResourceGroupResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "media.ServiceClient", "ListByResourceGroup", resp, "Failure responding to request")
            }

    return
    }

    // ListByResourceGroupPreparer prepares the ListByResourceGroup request.
    func (client ServiceClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2015-10-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
    // http.Response Body if it receives an error.
    func (client ServiceClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client ServiceClient) ListByResourceGroupResponder(resp *http.Response) (result ServiceCollection, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// ListKeys lists the keys for a Media Service.
    // Parameters:
        // resourceGroupName - name of the resource group within the Azure subscription.
        // mediaServiceName - name of the Media Service.
func (client ServiceClient) ListKeys(ctx context.Context, resourceGroupName string, mediaServiceName string) (result ServiceKeys, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ServiceClient.ListKeys")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: mediaServiceName,
             Constraints: []validation.Constraint{	{Target: "mediaServiceName", Name: validation.MaxLength, Rule: 24, Chain: nil },
            	{Target: "mediaServiceName", Name: validation.MinLength, Rule: 3, Chain: nil },
            	{Target: "mediaServiceName", Name: validation.Pattern, Rule: `^[a-z0-9]{3,24}$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("media.ServiceClient", "ListKeys", err.Error())
            }

                req, err := client.ListKeysPreparer(ctx, resourceGroupName, mediaServiceName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "media.ServiceClient", "ListKeys", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListKeysSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "media.ServiceClient", "ListKeys", resp, "Failure sending request")
            return
            }

            result, err = client.ListKeysResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "media.ServiceClient", "ListKeys", resp, "Failure responding to request")
            }

    return
    }

    // ListKeysPreparer prepares the ListKeys request.
    func (client ServiceClient) ListKeysPreparer(ctx context.Context, resourceGroupName string, mediaServiceName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "mediaServiceName": autorest.Encode("path",mediaServiceName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2015-10-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{mediaServiceName}/listKeys",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListKeysSender sends the ListKeys request. The method will close the
    // http.Response Body if it receives an error.
    func (client ServiceClient) ListKeysSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListKeysResponder handles the response to the ListKeys request. The method always
// closes the http.Response Body.
func (client ServiceClient) ListKeysResponder(resp *http.Response) (result ServiceKeys, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// RegenerateKey regenerates a primary or secondary key for a Media Service.
    // Parameters:
        // resourceGroupName - name of the resource group within the Azure subscription.
        // mediaServiceName - name of the Media Service.
        // parameters - properties needed to regenerate the Media Service key.
func (client ServiceClient) RegenerateKey(ctx context.Context, resourceGroupName string, mediaServiceName string, parameters RegenerateKeyInput) (result RegenerateKeyOutput, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ServiceClient.RegenerateKey")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: mediaServiceName,
             Constraints: []validation.Constraint{	{Target: "mediaServiceName", Name: validation.MaxLength, Rule: 24, Chain: nil },
            	{Target: "mediaServiceName", Name: validation.MinLength, Rule: 3, Chain: nil },
            	{Target: "mediaServiceName", Name: validation.Pattern, Rule: `^[a-z0-9]{3,24}$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("media.ServiceClient", "RegenerateKey", err.Error())
            }

                req, err := client.RegenerateKeyPreparer(ctx, resourceGroupName, mediaServiceName, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "media.ServiceClient", "RegenerateKey", nil , "Failure preparing request")
    return
    }

            resp, err := client.RegenerateKeySender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "media.ServiceClient", "RegenerateKey", resp, "Failure sending request")
            return
            }

            result, err = client.RegenerateKeyResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "media.ServiceClient", "RegenerateKey", resp, "Failure responding to request")
            }

    return
    }

    // RegenerateKeyPreparer prepares the RegenerateKey request.
    func (client ServiceClient) RegenerateKeyPreparer(ctx context.Context, resourceGroupName string, mediaServiceName string, parameters RegenerateKeyInput) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "mediaServiceName": autorest.Encode("path",mediaServiceName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2015-10-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{mediaServiceName}/regenerateKey",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // RegenerateKeySender sends the RegenerateKey request. The method will close the
    // http.Response Body if it receives an error.
    func (client ServiceClient) RegenerateKeySender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// RegenerateKeyResponder handles the response to the RegenerateKey request. The method always
// closes the http.Response Body.
func (client ServiceClient) RegenerateKeyResponder(resp *http.Response) (result RegenerateKeyOutput, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// SyncStorageKeys synchronizes storage account keys for a storage account associated with the Media Service account.
    // Parameters:
        // resourceGroupName - name of the resource group within the Azure subscription.
        // mediaServiceName - name of the Media Service.
        // parameters - properties needed to synchronize the keys for a storage account to the Media Service.
func (client ServiceClient) SyncStorageKeys(ctx context.Context, resourceGroupName string, mediaServiceName string, parameters SyncStorageKeysInput) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ServiceClient.SyncStorageKeys")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: mediaServiceName,
             Constraints: []validation.Constraint{	{Target: "mediaServiceName", Name: validation.MaxLength, Rule: 24, Chain: nil },
            	{Target: "mediaServiceName", Name: validation.MinLength, Rule: 3, Chain: nil },
            	{Target: "mediaServiceName", Name: validation.Pattern, Rule: `^[a-z0-9]{3,24}$`, Chain: nil }}},
            { TargetValue: parameters,
             Constraints: []validation.Constraint{	{Target: "parameters.ID", Name: validation.Null, Rule: true, Chain: nil }}}}); err != nil {
            return result, validation.NewError("media.ServiceClient", "SyncStorageKeys", err.Error())
            }

                req, err := client.SyncStorageKeysPreparer(ctx, resourceGroupName, mediaServiceName, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "media.ServiceClient", "SyncStorageKeys", nil , "Failure preparing request")
    return
    }

            resp, err := client.SyncStorageKeysSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "media.ServiceClient", "SyncStorageKeys", resp, "Failure sending request")
            return
            }

            result, err = client.SyncStorageKeysResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "media.ServiceClient", "SyncStorageKeys", resp, "Failure responding to request")
            }

    return
    }

    // SyncStorageKeysPreparer prepares the SyncStorageKeys request.
    func (client ServiceClient) SyncStorageKeysPreparer(ctx context.Context, resourceGroupName string, mediaServiceName string, parameters SyncStorageKeysInput) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "mediaServiceName": autorest.Encode("path",mediaServiceName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2015-10-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{mediaServiceName}/syncStorageKeys",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // SyncStorageKeysSender sends the SyncStorageKeys request. The method will close the
    // http.Response Body if it receives an error.
    func (client ServiceClient) SyncStorageKeysSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// SyncStorageKeysResponder handles the response to the SyncStorageKeys request. The method always
// closes the http.Response Body.
func (client ServiceClient) SyncStorageKeysResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Update updates a Media Service.
    // Parameters:
        // resourceGroupName - name of the resource group within the Azure subscription.
        // mediaServiceName - name of the Media Service.
        // parameters - media Service properties needed for update.
func (client ServiceClient) Update(ctx context.Context, resourceGroupName string, mediaServiceName string, parameters Service) (result Service, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ServiceClient.Update")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: mediaServiceName,
             Constraints: []validation.Constraint{	{Target: "mediaServiceName", Name: validation.MaxLength, Rule: 24, Chain: nil },
            	{Target: "mediaServiceName", Name: validation.MinLength, Rule: 3, Chain: nil },
            	{Target: "mediaServiceName", Name: validation.Pattern, Rule: `^[a-z0-9]{3,24}$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("media.ServiceClient", "Update", err.Error())
            }

                req, err := client.UpdatePreparer(ctx, resourceGroupName, mediaServiceName, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "media.ServiceClient", "Update", nil , "Failure preparing request")
    return
    }

            resp, err := client.UpdateSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "media.ServiceClient", "Update", resp, "Failure sending request")
            return
            }

            result, err = client.UpdateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "media.ServiceClient", "Update", resp, "Failure responding to request")
            }

    return
    }

    // UpdatePreparer prepares the Update request.
    func (client ServiceClient) UpdatePreparer(ctx context.Context, resourceGroupName string, mediaServiceName string, parameters Service) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "mediaServiceName": autorest.Encode("path",mediaServiceName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2015-10-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPatch(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{mediaServiceName}",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // UpdateSender sends the Update request. The method will close the
    // http.Response Body if it receives an error.
    func (client ServiceClient) UpdateSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ServiceClient) UpdateResponder(resp *http.Response) (result Service, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

