package logic

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

// SessionsClient is the REST API for Azure Logic Apps.
type SessionsClient struct {
    BaseClient
}
// NewSessionsClient creates an instance of the SessionsClient client.
func NewSessionsClient(subscriptionID string) SessionsClient {
    return NewSessionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSessionsClientWithBaseURI creates an instance of the SessionsClient client.
    func NewSessionsClientWithBaseURI(baseURI string, subscriptionID string) SessionsClient {
        return SessionsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CreateOrUpdate creates or updates an integration account session.
    // Parameters:
        // resourceGroupName - the resource group name.
        // integrationAccountName - the integration account name.
        // sessionName - the integration account session name.
        // session - the integration account session.
func (client SessionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, integrationAccountName string, sessionName string, session IntegrationAccountSession) (result IntegrationAccountSession, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SessionsClient.CreateOrUpdate")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: session,
             Constraints: []validation.Constraint{	{Target: "session.IntegrationAccountSessionProperties", Name: validation.Null, Rule: true, Chain: nil }}}}); err != nil {
            return result, validation.NewError("logic.SessionsClient", "CreateOrUpdate", err.Error())
            }

                req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, integrationAccountName, sessionName, session)
    if err != nil {
    err = autorest.NewErrorWithError(err, "logic.SessionsClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

            resp, err := client.CreateOrUpdateSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "logic.SessionsClient", "CreateOrUpdate", resp, "Failure sending request")
            return
            }

            result, err = client.CreateOrUpdateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "logic.SessionsClient", "CreateOrUpdate", resp, "Failure responding to request")
            }

    return
    }

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client SessionsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, integrationAccountName string, sessionName string, session IntegrationAccountSession) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "integrationAccountName": autorest.Encode("path",integrationAccountName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "sessionName": autorest.Encode("path",sessionName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2016-06-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPut(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/sessions/{sessionName}",pathParameters),
    autorest.WithJSON(session),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client SessionsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client SessionsClient) CreateOrUpdateResponder(resp *http.Response) (result IntegrationAccountSession, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Delete deletes an integration account session.
    // Parameters:
        // resourceGroupName - the resource group name.
        // integrationAccountName - the integration account name.
        // sessionName - the integration account session name.
func (client SessionsClient) Delete(ctx context.Context, resourceGroupName string, integrationAccountName string, sessionName string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SessionsClient.Delete")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.DeletePreparer(ctx, resourceGroupName, integrationAccountName, sessionName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "logic.SessionsClient", "Delete", nil , "Failure preparing request")
    return
    }

            resp, err := client.DeleteSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "logic.SessionsClient", "Delete", resp, "Failure sending request")
            return
            }

            result, err = client.DeleteResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "logic.SessionsClient", "Delete", resp, "Failure responding to request")
            }

    return
    }

    // DeletePreparer prepares the Delete request.
    func (client SessionsClient) DeletePreparer(ctx context.Context, resourceGroupName string, integrationAccountName string, sessionName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "integrationAccountName": autorest.Encode("path",integrationAccountName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "sessionName": autorest.Encode("path",sessionName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2016-06-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsDelete(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/sessions/{sessionName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client SessionsClient) DeleteSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SessionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Get gets an integration account session.
    // Parameters:
        // resourceGroupName - the resource group name.
        // integrationAccountName - the integration account name.
        // sessionName - the integration account session name.
func (client SessionsClient) Get(ctx context.Context, resourceGroupName string, integrationAccountName string, sessionName string) (result IntegrationAccountSession, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SessionsClient.Get")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.GetPreparer(ctx, resourceGroupName, integrationAccountName, sessionName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "logic.SessionsClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "logic.SessionsClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "logic.SessionsClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client SessionsClient) GetPreparer(ctx context.Context, resourceGroupName string, integrationAccountName string, sessionName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "integrationAccountName": autorest.Encode("path",integrationAccountName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "sessionName": autorest.Encode("path",sessionName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2016-06-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/sessions/{sessionName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client SessionsClient) GetSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SessionsClient) GetResponder(resp *http.Response) (result IntegrationAccountSession, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// ListByIntegrationAccounts gets a list of integration account sessions.
    // Parameters:
        // resourceGroupName - the resource group name.
        // integrationAccountName - the integration account name.
        // top - the number of items to be included in the result.
        // filter - the filter to apply on the operation. Options for filters include: ChangedTime.
func (client SessionsClient) ListByIntegrationAccounts(ctx context.Context, resourceGroupName string, integrationAccountName string, top *int32, filter string) (result IntegrationAccountSessionListResultPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SessionsClient.ListByIntegrationAccounts")
        defer func() {
            sc := -1
            if result.iaslr.Response.Response != nil {
                sc = result.iaslr.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
                result.fn = client.listByIntegrationAccountsNextResults
    req, err := client.ListByIntegrationAccountsPreparer(ctx, resourceGroupName, integrationAccountName, top, filter)
    if err != nil {
    err = autorest.NewErrorWithError(err, "logic.SessionsClient", "ListByIntegrationAccounts", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListByIntegrationAccountsSender(req)
            if err != nil {
            result.iaslr.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "logic.SessionsClient", "ListByIntegrationAccounts", resp, "Failure sending request")
            return
            }

            result.iaslr, err = client.ListByIntegrationAccountsResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "logic.SessionsClient", "ListByIntegrationAccounts", resp, "Failure responding to request")
            }

    return
    }

    // ListByIntegrationAccountsPreparer prepares the ListByIntegrationAccounts request.
    func (client SessionsClient) ListByIntegrationAccountsPreparer(ctx context.Context, resourceGroupName string, integrationAccountName string, top *int32, filter string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "integrationAccountName": autorest.Encode("path",integrationAccountName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2016-06-01"
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
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/sessions",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByIntegrationAccountsSender sends the ListByIntegrationAccounts request. The method will close the
    // http.Response Body if it receives an error.
    func (client SessionsClient) ListByIntegrationAccountsSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListByIntegrationAccountsResponder handles the response to the ListByIntegrationAccounts request. The method always
// closes the http.Response Body.
func (client SessionsClient) ListByIntegrationAccountsResponder(resp *http.Response) (result IntegrationAccountSessionListResult, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

            // listByIntegrationAccountsNextResults retrieves the next set of results, if any.
            func (client SessionsClient) listByIntegrationAccountsNextResults(ctx context.Context, lastResults IntegrationAccountSessionListResult) (result IntegrationAccountSessionListResult, err error) {
            req, err := lastResults.integrationAccountSessionListResultPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "logic.SessionsClient", "listByIntegrationAccountsNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListByIntegrationAccountsSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "logic.SessionsClient", "listByIntegrationAccountsNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListByIntegrationAccountsResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "logic.SessionsClient", "listByIntegrationAccountsNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListByIntegrationAccountsComplete enumerates all values, automatically crossing page boundaries as required.
    func (client SessionsClient) ListByIntegrationAccountsComplete(ctx context.Context, resourceGroupName string, integrationAccountName string, top *int32, filter string) (result IntegrationAccountSessionListResultIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/SessionsClient.ListByIntegrationAccounts")
            defer func() {
                sc := -1
                if result.Response().Response.Response != nil {
                    sc = result.page.Response().Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
     }
        result.page, err = client.ListByIntegrationAccounts(ctx, resourceGroupName, integrationAccountName, top, filter)
                return
        }

