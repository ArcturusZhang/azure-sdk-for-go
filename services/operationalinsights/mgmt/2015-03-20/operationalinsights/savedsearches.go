package operationalinsights

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

// SavedSearchesClient is the operational Insights Client
type SavedSearchesClient struct {
    BaseClient
}
// NewSavedSearchesClient creates an instance of the SavedSearchesClient client.
func NewSavedSearchesClient(subscriptionID string, purgeID string) SavedSearchesClient {
    return NewSavedSearchesClientWithBaseURI(DefaultBaseURI, subscriptionID, purgeID)
}

// NewSavedSearchesClientWithBaseURI creates an instance of the SavedSearchesClient client.
    func NewSavedSearchesClientWithBaseURI(baseURI string, subscriptionID string, purgeID string) SavedSearchesClient {
        return SavedSearchesClient{ NewWithBaseURI(baseURI, subscriptionID, purgeID)}
    }

// CreateOrUpdate creates or updates a saved search for a given workspace.
    // Parameters:
        // resourceGroupName - the Resource Group name.
        // workspaceName - the Log Analytics Workspace name.
        // savedSearchID - the id of the saved search.
        // parameters - the parameters required to save a search.
func (client SavedSearchesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, savedSearchID string, parameters SavedSearch) (result SavedSearch, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SavedSearchesClient.CreateOrUpdate")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: resourceGroupName,
             Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}},
            { TargetValue: parameters,
             Constraints: []validation.Constraint{	{Target: "parameters.SavedSearchProperties", Name: validation.Null, Rule: true ,
            Chain: []validation.Constraint{	{Target: "parameters.SavedSearchProperties.Category", Name: validation.Null, Rule: true, Chain: nil },
            	{Target: "parameters.SavedSearchProperties.DisplayName", Name: validation.Null, Rule: true, Chain: nil },
            	{Target: "parameters.SavedSearchProperties.Query", Name: validation.Null, Rule: true, Chain: nil },
            	{Target: "parameters.SavedSearchProperties.Version", Name: validation.Null, Rule: false ,
            Chain: []validation.Constraint{	{Target: "parameters.SavedSearchProperties.Version", Name: validation.InclusiveMaximum, Rule: int64(2), Chain: nil },
            	{Target: "parameters.SavedSearchProperties.Version", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil },
            }},
            }}}}}); err != nil {
            return result, validation.NewError("operationalinsights.SavedSearchesClient", "CreateOrUpdate", err.Error())
            }

                req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, workspaceName, savedSearchID, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "operationalinsights.SavedSearchesClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

            resp, err := client.CreateOrUpdateSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "operationalinsights.SavedSearchesClient", "CreateOrUpdate", resp, "Failure sending request")
            return
            }

            result, err = client.CreateOrUpdateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "operationalinsights.SavedSearchesClient", "CreateOrUpdate", resp, "Failure responding to request")
            }

    return
    }

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client SavedSearchesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, workspaceName string, savedSearchID string, parameters SavedSearch) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "savedSearchId": autorest.Encode("path",savedSearchID),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "workspaceName": autorest.Encode("path",workspaceName),
            }

                        const APIVersion = "2015-03-20"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

                    parameters.ID = nil
                parameters.Name = nil
                parameters.Type = nil
    preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPut(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/savedSearches/{savedSearchId}",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client SavedSearchesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client SavedSearchesClient) CreateOrUpdateResponder(resp *http.Response) (result SavedSearch, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Delete deletes the specified saved search in a given workspace.
    // Parameters:
        // resourceGroupName - the Resource Group name.
        // workspaceName - the Log Analytics Workspace name.
        // savedSearchID - the id of the saved search.
func (client SavedSearchesClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, savedSearchID string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SavedSearchesClient.Delete")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: resourceGroupName,
             Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("operationalinsights.SavedSearchesClient", "Delete", err.Error())
            }

                req, err := client.DeletePreparer(ctx, resourceGroupName, workspaceName, savedSearchID)
    if err != nil {
    err = autorest.NewErrorWithError(err, "operationalinsights.SavedSearchesClient", "Delete", nil , "Failure preparing request")
    return
    }

            resp, err := client.DeleteSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "operationalinsights.SavedSearchesClient", "Delete", resp, "Failure sending request")
            return
            }

            result, err = client.DeleteResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "operationalinsights.SavedSearchesClient", "Delete", resp, "Failure responding to request")
            }

    return
    }

    // DeletePreparer prepares the Delete request.
    func (client SavedSearchesClient) DeletePreparer(ctx context.Context, resourceGroupName string, workspaceName string, savedSearchID string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "savedSearchId": autorest.Encode("path",savedSearchID),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "workspaceName": autorest.Encode("path",workspaceName),
            }

                        const APIVersion = "2015-03-20"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsDelete(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/savedSearches/{savedSearchId}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client SavedSearchesClient) DeleteSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SavedSearchesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Get gets the specified saved search for a given workspace.
    // Parameters:
        // resourceGroupName - the Resource Group name.
        // workspaceName - the Log Analytics Workspace name.
        // savedSearchID - the id of the saved search.
func (client SavedSearchesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, savedSearchID string) (result SavedSearch, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SavedSearchesClient.Get")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: resourceGroupName,
             Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("operationalinsights.SavedSearchesClient", "Get", err.Error())
            }

                req, err := client.GetPreparer(ctx, resourceGroupName, workspaceName, savedSearchID)
    if err != nil {
    err = autorest.NewErrorWithError(err, "operationalinsights.SavedSearchesClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "operationalinsights.SavedSearchesClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "operationalinsights.SavedSearchesClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client SavedSearchesClient) GetPreparer(ctx context.Context, resourceGroupName string, workspaceName string, savedSearchID string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "savedSearchId": autorest.Encode("path",savedSearchID),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "workspaceName": autorest.Encode("path",workspaceName),
            }

                        const APIVersion = "2015-03-20"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/savedSearches/{savedSearchId}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client SavedSearchesClient) GetSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SavedSearchesClient) GetResponder(resp *http.Response) (result SavedSearch, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// GetResults gets the results from a saved search for a given workspace.
    // Parameters:
        // resourceGroupName - the Resource Group name.
        // workspaceName - the Log Analytics Workspace name.
        // savedSearchID - the id of the saved search.
func (client SavedSearchesClient) GetResults(ctx context.Context, resourceGroupName string, workspaceName string, savedSearchID string) (result SearchResultsResponse, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SavedSearchesClient.GetResults")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: resourceGroupName,
             Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("operationalinsights.SavedSearchesClient", "GetResults", err.Error())
            }

                req, err := client.GetResultsPreparer(ctx, resourceGroupName, workspaceName, savedSearchID)
    if err != nil {
    err = autorest.NewErrorWithError(err, "operationalinsights.SavedSearchesClient", "GetResults", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetResultsSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "operationalinsights.SavedSearchesClient", "GetResults", resp, "Failure sending request")
            return
            }

            result, err = client.GetResultsResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "operationalinsights.SavedSearchesClient", "GetResults", resp, "Failure responding to request")
            }

    return
    }

    // GetResultsPreparer prepares the GetResults request.
    func (client SavedSearchesClient) GetResultsPreparer(ctx context.Context, resourceGroupName string, workspaceName string, savedSearchID string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "savedSearchId": autorest.Encode("path",savedSearchID),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "workspaceName": autorest.Encode("path",workspaceName),
            }

                        const APIVersion = "2015-03-20"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/savedSearches/{savedSearchId}/results",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetResultsSender sends the GetResults request. The method will close the
    // http.Response Body if it receives an error.
    func (client SavedSearchesClient) GetResultsSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetResultsResponder handles the response to the GetResults request. The method always
// closes the http.Response Body.
func (client SavedSearchesClient) GetResultsResponder(resp *http.Response) (result SearchResultsResponse, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// ListByWorkspace gets the saved searches for a given Log Analytics Workspace
    // Parameters:
        // resourceGroupName - the Resource Group name.
        // workspaceName - the Log Analytics Workspace name.
func (client SavedSearchesClient) ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string) (result SavedSearchesListResult, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SavedSearchesClient.ListByWorkspace")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: resourceGroupName,
             Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("operationalinsights.SavedSearchesClient", "ListByWorkspace", err.Error())
            }

                req, err := client.ListByWorkspacePreparer(ctx, resourceGroupName, workspaceName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "operationalinsights.SavedSearchesClient", "ListByWorkspace", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListByWorkspaceSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "operationalinsights.SavedSearchesClient", "ListByWorkspace", resp, "Failure sending request")
            return
            }

            result, err = client.ListByWorkspaceResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "operationalinsights.SavedSearchesClient", "ListByWorkspace", resp, "Failure responding to request")
            }

    return
    }

    // ListByWorkspacePreparer prepares the ListByWorkspace request.
    func (client SavedSearchesClient) ListByWorkspacePreparer(ctx context.Context, resourceGroupName string, workspaceName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "workspaceName": autorest.Encode("path",workspaceName),
            }

                        const APIVersion = "2015-03-20"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/savedSearches",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByWorkspaceSender sends the ListByWorkspace request. The method will close the
    // http.Response Body if it receives an error.
    func (client SavedSearchesClient) ListByWorkspaceSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListByWorkspaceResponder handles the response to the ListByWorkspace request. The method always
// closes the http.Response Body.
func (client SavedSearchesClient) ListByWorkspaceResponder(resp *http.Response) (result SavedSearchesListResult, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

