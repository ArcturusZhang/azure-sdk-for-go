package web

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

// ConnectionsClient is the webSite Management Client
type ConnectionsClient struct {
    BaseClient
}
// NewConnectionsClient creates an instance of the ConnectionsClient client.
func NewConnectionsClient(subscriptionID string) ConnectionsClient {
    return NewConnectionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewConnectionsClientWithBaseURI creates an instance of the ConnectionsClient client.
    func NewConnectionsClientWithBaseURI(baseURI string, subscriptionID string) ConnectionsClient {
        return ConnectionsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// ConfirmConsentCode confirms consent code of a connection.
    // Parameters:
        // resourceGroupName - the resource group name.
        // connectionName - the connection name.
        // content - the content.
func (client ConnectionsClient) ConfirmConsentCode(ctx context.Context, resourceGroupName string, connectionName string, content ConfirmConsentCodeInput) (result Connection, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ConnectionsClient.ConfirmConsentCode")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.ConfirmConsentCodePreparer(ctx, resourceGroupName, connectionName, content)
    if err != nil {
    err = autorest.NewErrorWithError(err, "web.ConnectionsClient", "ConfirmConsentCode", nil , "Failure preparing request")
    return
    }

            resp, err := client.ConfirmConsentCodeSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "web.ConnectionsClient", "ConfirmConsentCode", resp, "Failure sending request")
            return
            }

            result, err = client.ConfirmConsentCodeResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "web.ConnectionsClient", "ConfirmConsentCode", resp, "Failure responding to request")
            }

    return
    }

    // ConfirmConsentCodePreparer prepares the ConfirmConsentCode request.
    func (client ConnectionsClient) ConfirmConsentCodePreparer(ctx context.Context, resourceGroupName string, connectionName string, content ConfirmConsentCodeInput) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "connectionName": autorest.Encode("path",connectionName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2015-08-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/connections/{connectionName}/confirmConsentCode",pathParameters),
    autorest.WithJSON(content),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ConfirmConsentCodeSender sends the ConfirmConsentCode request. The method will close the
    // http.Response Body if it receives an error.
    func (client ConnectionsClient) ConfirmConsentCodeSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ConfirmConsentCodeResponder handles the response to the ConfirmConsentCode request. The method always
// closes the http.Response Body.
func (client ConnectionsClient) ConfirmConsentCodeResponder(resp *http.Response) (result Connection, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// CreateOrUpdate creates or updates a connection.
    // Parameters:
        // resourceGroupName - the resource group name.
        // connectionName - the connection name.
        // connection - the connection.
func (client ConnectionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, connectionName string, connection Connection) (result Connection, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ConnectionsClient.CreateOrUpdate")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, connectionName, connection)
    if err != nil {
    err = autorest.NewErrorWithError(err, "web.ConnectionsClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

            resp, err := client.CreateOrUpdateSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "web.ConnectionsClient", "CreateOrUpdate", resp, "Failure sending request")
            return
            }

            result, err = client.CreateOrUpdateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "web.ConnectionsClient", "CreateOrUpdate", resp, "Failure responding to request")
            }

    return
    }

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client ConnectionsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, connectionName string, connection Connection) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "connectionName": autorest.Encode("path",connectionName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2015-08-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPut(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/connections/{connectionName}",pathParameters),
    autorest.WithJSON(connection),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client ConnectionsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ConnectionsClient) CreateOrUpdateResponder(resp *http.Response) (result Connection, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Delete deletes a connection.
    // Parameters:
        // resourceGroupName - the resource group name.
        // connectionName - the connection name.
func (client ConnectionsClient) Delete(ctx context.Context, resourceGroupName string, connectionName string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ConnectionsClient.Delete")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.DeletePreparer(ctx, resourceGroupName, connectionName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "web.ConnectionsClient", "Delete", nil , "Failure preparing request")
    return
    }

            resp, err := client.DeleteSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "web.ConnectionsClient", "Delete", resp, "Failure sending request")
            return
            }

            result, err = client.DeleteResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "web.ConnectionsClient", "Delete", resp, "Failure responding to request")
            }

    return
    }

    // DeletePreparer prepares the Delete request.
    func (client ConnectionsClient) DeletePreparer(ctx context.Context, resourceGroupName string, connectionName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "connectionName": autorest.Encode("path",connectionName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2015-08-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsDelete(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/connections/{connectionName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client ConnectionsClient) DeleteSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ConnectionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Get gets a connection.
    // Parameters:
        // resourceGroupName - the resource group name.
        // connectionName - the connection name.
func (client ConnectionsClient) Get(ctx context.Context, resourceGroupName string, connectionName string) (result Connection, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ConnectionsClient.Get")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.GetPreparer(ctx, resourceGroupName, connectionName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "web.ConnectionsClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "web.ConnectionsClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "web.ConnectionsClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client ConnectionsClient) GetPreparer(ctx context.Context, resourceGroupName string, connectionName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "connectionName": autorest.Encode("path",connectionName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2015-08-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/connections/{connectionName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client ConnectionsClient) GetSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ConnectionsClient) GetResponder(resp *http.Response) (result Connection, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// List gets a list of connections.
    // Parameters:
        // resourceGroupName - resource Group Name
        // top - the number of items to be included in the result.
        // filter - the filter to apply on the operation.
func (client ConnectionsClient) List(ctx context.Context, resourceGroupName string, top *int32, filter string) (result ConnectionCollectionPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ConnectionsClient.List")
        defer func() {
            sc := -1
            if result.cc.Response.Response != nil {
                sc = result.cc.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
                result.fn = client.listNextResults
    req, err := client.ListPreparer(ctx, resourceGroupName, top, filter)
    if err != nil {
    err = autorest.NewErrorWithError(err, "web.ConnectionsClient", "List", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListSender(req)
            if err != nil {
            result.cc.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "web.ConnectionsClient", "List", resp, "Failure sending request")
            return
            }

            result.cc, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "web.ConnectionsClient", "List", resp, "Failure responding to request")
            }

    return
    }

    // ListPreparer prepares the List request.
    func (client ConnectionsClient) ListPreparer(ctx context.Context, resourceGroupName string, top *int32, filter string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2015-08-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }
            if top != nil {
            queryParameters["$top"] = autorest.Encode("query",*top)
            }
            if len(filter) > 0 {
            queryParameters["$filter"] = autorest.Encode("query",filter)
            }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/connections",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client ConnectionsClient) ListSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ConnectionsClient) ListResponder(resp *http.Response) (result ConnectionCollection, err error) {
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
            func (client ConnectionsClient) listNextResults(ctx context.Context, lastResults ConnectionCollection) (result ConnectionCollection, err error) {
            req, err := lastResults.connectionCollectionPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "web.ConnectionsClient", "listNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "web.ConnectionsClient", "listNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "web.ConnectionsClient", "listNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListComplete enumerates all values, automatically crossing page boundaries as required.
    func (client ConnectionsClient) ListComplete(ctx context.Context, resourceGroupName string, top *int32, filter string) (result ConnectionCollectionIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/ConnectionsClient.List")
            defer func() {
                sc := -1
                if result.Response().Response.Response != nil {
                    sc = result.page.Response().Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
     }
        result.page, err = client.List(ctx, resourceGroupName, top, filter)
                return
        }

// ListConnectionKeys lists connection keys.
    // Parameters:
        // resourceGroupName - the resource group name.
        // connectionName - the connection name.
        // content - the content.
func (client ConnectionsClient) ListConnectionKeys(ctx context.Context, resourceGroupName string, connectionName string, content ListConnectionKeysInput) (result ConnectionSecrets, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ConnectionsClient.ListConnectionKeys")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.ListConnectionKeysPreparer(ctx, resourceGroupName, connectionName, content)
    if err != nil {
    err = autorest.NewErrorWithError(err, "web.ConnectionsClient", "ListConnectionKeys", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListConnectionKeysSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "web.ConnectionsClient", "ListConnectionKeys", resp, "Failure sending request")
            return
            }

            result, err = client.ListConnectionKeysResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "web.ConnectionsClient", "ListConnectionKeys", resp, "Failure responding to request")
            }

    return
    }

    // ListConnectionKeysPreparer prepares the ListConnectionKeys request.
    func (client ConnectionsClient) ListConnectionKeysPreparer(ctx context.Context, resourceGroupName string, connectionName string, content ListConnectionKeysInput) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "connectionName": autorest.Encode("path",connectionName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2015-08-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/connections/{connectionName}/listConnectionKeys",pathParameters),
    autorest.WithJSON(content),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListConnectionKeysSender sends the ListConnectionKeys request. The method will close the
    // http.Response Body if it receives an error.
    func (client ConnectionsClient) ListConnectionKeysSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListConnectionKeysResponder handles the response to the ListConnectionKeys request. The method always
// closes the http.Response Body.
func (client ConnectionsClient) ListConnectionKeysResponder(resp *http.Response) (result ConnectionSecrets, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// ListConsentLinks lists consent links of a connection.
    // Parameters:
        // resourceGroupName - the resource group name.
        // connectionName - the connection name.
        // content - the content.
func (client ConnectionsClient) ListConsentLinks(ctx context.Context, resourceGroupName string, connectionName string, content ConsentLinkInput) (result ConsentLinkPayload, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ConnectionsClient.ListConsentLinks")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.ListConsentLinksPreparer(ctx, resourceGroupName, connectionName, content)
    if err != nil {
    err = autorest.NewErrorWithError(err, "web.ConnectionsClient", "ListConsentLinks", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListConsentLinksSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "web.ConnectionsClient", "ListConsentLinks", resp, "Failure sending request")
            return
            }

            result, err = client.ListConsentLinksResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "web.ConnectionsClient", "ListConsentLinks", resp, "Failure responding to request")
            }

    return
    }

    // ListConsentLinksPreparer prepares the ListConsentLinks request.
    func (client ConnectionsClient) ListConsentLinksPreparer(ctx context.Context, resourceGroupName string, connectionName string, content ConsentLinkInput) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "connectionName": autorest.Encode("path",connectionName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2015-08-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/connections/{connectionName}/listConsentLinks",pathParameters),
    autorest.WithJSON(content),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListConsentLinksSender sends the ListConsentLinks request. The method will close the
    // http.Response Body if it receives an error.
    func (client ConnectionsClient) ListConsentLinksSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListConsentLinksResponder handles the response to the ListConsentLinks request. The method always
// closes the http.Response Body.
func (client ConnectionsClient) ListConsentLinksResponder(resp *http.Response) (result ConsentLinkPayload, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

