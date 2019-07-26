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
)

// MediaservicesClient is the client for the Mediaservices methods of the Media service.
type MediaservicesClient struct {
    BaseClient
}
// NewMediaservicesClient creates an instance of the MediaservicesClient client.
func NewMediaservicesClient(subscriptionID string) MediaservicesClient {
    return NewMediaservicesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewMediaservicesClientWithBaseURI creates an instance of the MediaservicesClient client.
    func NewMediaservicesClientWithBaseURI(baseURI string, subscriptionID string) MediaservicesClient {
        return MediaservicesClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CreateOrUpdate creates or updates a Media Services account
    // Parameters:
        // resourceGroupName - the name of the resource group within the Azure subscription.
        // accountName - the Media Services account name.
        // parameters - the request parameters
func (client MediaservicesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, parameters Service) (result Service, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/MediaservicesClient.CreateOrUpdate")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, accountName, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "media.MediaservicesClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

            resp, err := client.CreateOrUpdateSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "media.MediaservicesClient", "CreateOrUpdate", resp, "Failure sending request")
            return
            }

            result, err = client.CreateOrUpdateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "media.MediaservicesClient", "CreateOrUpdate", resp, "Failure responding to request")
            }

    return
    }

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client MediaservicesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, accountName string, parameters Service) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "accountName": autorest.Encode("path",accountName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-06-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPut(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client MediaservicesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client MediaservicesClient) CreateOrUpdateResponder(resp *http.Response) (result Service, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Delete deletes a Media Services account
    // Parameters:
        // resourceGroupName - the name of the resource group within the Azure subscription.
        // accountName - the Media Services account name.
func (client MediaservicesClient) Delete(ctx context.Context, resourceGroupName string, accountName string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/MediaservicesClient.Delete")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.DeletePreparer(ctx, resourceGroupName, accountName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "media.MediaservicesClient", "Delete", nil , "Failure preparing request")
    return
    }

            resp, err := client.DeleteSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "media.MediaservicesClient", "Delete", resp, "Failure sending request")
            return
            }

            result, err = client.DeleteResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "media.MediaservicesClient", "Delete", resp, "Failure responding to request")
            }

    return
    }

    // DeletePreparer prepares the Delete request.
    func (client MediaservicesClient) DeletePreparer(ctx context.Context, resourceGroupName string, accountName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "accountName": autorest.Encode("path",accountName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-06-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsDelete(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client MediaservicesClient) DeleteSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client MediaservicesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Get get the details of a Media Services account
    // Parameters:
        // resourceGroupName - the name of the resource group within the Azure subscription.
        // accountName - the Media Services account name.
func (client MediaservicesClient) Get(ctx context.Context, resourceGroupName string, accountName string) (result Service, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/MediaservicesClient.Get")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.GetPreparer(ctx, resourceGroupName, accountName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "media.MediaservicesClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "media.MediaservicesClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "media.MediaservicesClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client MediaservicesClient) GetPreparer(ctx context.Context, resourceGroupName string, accountName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "accountName": autorest.Encode("path",accountName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-06-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client MediaservicesClient) GetSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client MediaservicesClient) GetResponder(resp *http.Response) (result Service, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// GetBySubscription get the details of a Media Services account
    // Parameters:
        // accountName - the Media Services account name.
func (client MediaservicesClient) GetBySubscription(ctx context.Context, accountName string) (result SubscriptionMediaService, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/MediaservicesClient.GetBySubscription")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.GetBySubscriptionPreparer(ctx, accountName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "media.MediaservicesClient", "GetBySubscription", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetBySubscriptionSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "media.MediaservicesClient", "GetBySubscription", resp, "Failure sending request")
            return
            }

            result, err = client.GetBySubscriptionResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "media.MediaservicesClient", "GetBySubscription", resp, "Failure responding to request")
            }

    return
    }

    // GetBySubscriptionPreparer prepares the GetBySubscription request.
    func (client MediaservicesClient) GetBySubscriptionPreparer(ctx context.Context, accountName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "accountName": autorest.Encode("path",accountName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-06-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Media/mediaservices/{accountName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetBySubscriptionSender sends the GetBySubscription request. The method will close the
    // http.Response Body if it receives an error.
    func (client MediaservicesClient) GetBySubscriptionSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetBySubscriptionResponder handles the response to the GetBySubscription request. The method always
// closes the http.Response Body.
func (client MediaservicesClient) GetBySubscriptionResponder(resp *http.Response) (result SubscriptionMediaService, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// List list Media Services accounts in the resource group
    // Parameters:
        // resourceGroupName - the name of the resource group within the Azure subscription.
func (client MediaservicesClient) List(ctx context.Context, resourceGroupName string) (result ServiceCollectionPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/MediaservicesClient.List")
        defer func() {
            sc := -1
            if result.sc.Response.Response != nil {
                sc = result.sc.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
                result.fn = client.listNextResults
    req, err := client.ListPreparer(ctx, resourceGroupName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "media.MediaservicesClient", "List", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListSender(req)
            if err != nil {
            result.sc.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "media.MediaservicesClient", "List", resp, "Failure sending request")
            return
            }

            result.sc, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "media.MediaservicesClient", "List", resp, "Failure responding to request")
            }

    return
    }

    // ListPreparer prepares the List request.
    func (client MediaservicesClient) ListPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-06-01-preview"
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

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client MediaservicesClient) ListSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client MediaservicesClient) ListResponder(resp *http.Response) (result ServiceCollection, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

            // listNextResults retrieves the next set of results, if any.
            func (client MediaservicesClient) listNextResults(ctx context.Context, lastResults ServiceCollection) (result ServiceCollection, err error) {
            req, err := lastResults.serviceCollectionPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "media.MediaservicesClient", "listNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "media.MediaservicesClient", "listNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "media.MediaservicesClient", "listNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListComplete enumerates all values, automatically crossing page boundaries as required.
    func (client MediaservicesClient) ListComplete(ctx context.Context, resourceGroupName string) (result ServiceCollectionIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/MediaservicesClient.List")
            defer func() {
                sc := -1
                if result.Response().Response.Response != nil {
                    sc = result.page.Response().Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
     }
        result.page, err = client.List(ctx, resourceGroupName)
                return
        }

// ListBySubscription list Media Services accounts in the subscription.
func (client MediaservicesClient) ListBySubscription(ctx context.Context) (result SubscriptionMediaServiceCollectionPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/MediaservicesClient.ListBySubscription")
        defer func() {
            sc := -1
            if result.smsc.Response.Response != nil {
                sc = result.smsc.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
                result.fn = client.listBySubscriptionNextResults
    req, err := client.ListBySubscriptionPreparer(ctx)
    if err != nil {
    err = autorest.NewErrorWithError(err, "media.MediaservicesClient", "ListBySubscription", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListBySubscriptionSender(req)
            if err != nil {
            result.smsc.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "media.MediaservicesClient", "ListBySubscription", resp, "Failure sending request")
            return
            }

            result.smsc, err = client.ListBySubscriptionResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "media.MediaservicesClient", "ListBySubscription", resp, "Failure responding to request")
            }

    return
    }

    // ListBySubscriptionPreparer prepares the ListBySubscription request.
    func (client MediaservicesClient) ListBySubscriptionPreparer(ctx context.Context) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-06-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Media/mediaservices",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListBySubscriptionSender sends the ListBySubscription request. The method will close the
    // http.Response Body if it receives an error.
    func (client MediaservicesClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client MediaservicesClient) ListBySubscriptionResponder(resp *http.Response) (result SubscriptionMediaServiceCollection, err error) {
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
            func (client MediaservicesClient) listBySubscriptionNextResults(ctx context.Context, lastResults SubscriptionMediaServiceCollection) (result SubscriptionMediaServiceCollection, err error) {
            req, err := lastResults.subscriptionMediaServiceCollectionPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "media.MediaservicesClient", "listBySubscriptionNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListBySubscriptionSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "media.MediaservicesClient", "listBySubscriptionNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListBySubscriptionResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "media.MediaservicesClient", "listBySubscriptionNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListBySubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
    func (client MediaservicesClient) ListBySubscriptionComplete(ctx context.Context) (result SubscriptionMediaServiceCollectionIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/MediaservicesClient.ListBySubscription")
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

// SyncStorageKeys synchronizes storage account keys for a storage account associated with the Media Service account.
    // Parameters:
        // resourceGroupName - the name of the resource group within the Azure subscription.
        // accountName - the Media Services account name.
        // parameters - the request parameters
func (client MediaservicesClient) SyncStorageKeys(ctx context.Context, resourceGroupName string, accountName string, parameters SyncStorageKeysInput) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/MediaservicesClient.SyncStorageKeys")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.SyncStorageKeysPreparer(ctx, resourceGroupName, accountName, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "media.MediaservicesClient", "SyncStorageKeys", nil , "Failure preparing request")
    return
    }

            resp, err := client.SyncStorageKeysSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "media.MediaservicesClient", "SyncStorageKeys", resp, "Failure sending request")
            return
            }

            result, err = client.SyncStorageKeysResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "media.MediaservicesClient", "SyncStorageKeys", resp, "Failure responding to request")
            }

    return
    }

    // SyncStorageKeysPreparer prepares the SyncStorageKeys request.
    func (client MediaservicesClient) SyncStorageKeysPreparer(ctx context.Context, resourceGroupName string, accountName string, parameters SyncStorageKeysInput) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "accountName": autorest.Encode("path",accountName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-06-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/syncStorageKeys",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // SyncStorageKeysSender sends the SyncStorageKeys request. The method will close the
    // http.Response Body if it receives an error.
    func (client MediaservicesClient) SyncStorageKeysSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// SyncStorageKeysResponder handles the response to the SyncStorageKeys request. The method always
// closes the http.Response Body.
func (client MediaservicesClient) SyncStorageKeysResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Update updates an existing Media Services account
    // Parameters:
        // resourceGroupName - the name of the resource group within the Azure subscription.
        // accountName - the Media Services account name.
        // parameters - the request parameters
func (client MediaservicesClient) Update(ctx context.Context, resourceGroupName string, accountName string, parameters Service) (result Service, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/MediaservicesClient.Update")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.UpdatePreparer(ctx, resourceGroupName, accountName, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "media.MediaservicesClient", "Update", nil , "Failure preparing request")
    return
    }

            resp, err := client.UpdateSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "media.MediaservicesClient", "Update", resp, "Failure sending request")
            return
            }

            result, err = client.UpdateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "media.MediaservicesClient", "Update", resp, "Failure responding to request")
            }

    return
    }

    // UpdatePreparer prepares the Update request.
    func (client MediaservicesClient) UpdatePreparer(ctx context.Context, resourceGroupName string, accountName string, parameters Service) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "accountName": autorest.Encode("path",accountName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-06-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPatch(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // UpdateSender sends the Update request. The method will close the
    // http.Response Body if it receives an error.
    func (client MediaservicesClient) UpdateSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client MediaservicesClient) UpdateResponder(resp *http.Response) (result Service, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

