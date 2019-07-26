package security

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

// WorkspaceSettingsClient is the API spec for Microsoft.Security (Azure Security Center) resource provider
type WorkspaceSettingsClient struct {
    BaseClient
}
// NewWorkspaceSettingsClient creates an instance of the WorkspaceSettingsClient client.
func NewWorkspaceSettingsClient(subscriptionID string, ascLocation string) WorkspaceSettingsClient {
    return NewWorkspaceSettingsClientWithBaseURI(DefaultBaseURI, subscriptionID, ascLocation)
}

// NewWorkspaceSettingsClientWithBaseURI creates an instance of the WorkspaceSettingsClient client.
    func NewWorkspaceSettingsClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) WorkspaceSettingsClient {
        return WorkspaceSettingsClient{ NewWithBaseURI(baseURI, subscriptionID, ascLocation)}
    }

// Create creating settings about where we should store your security data and logs
    // Parameters:
        // workspaceSettingName - name of the security setting
        // workspaceSetting - security data setting object
func (client WorkspaceSettingsClient) Create(ctx context.Context, workspaceSettingName string, workspaceSetting WorkspaceSetting) (result WorkspaceSetting, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/WorkspaceSettingsClient.Create")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: client.SubscriptionID,
             Constraints: []validation.Constraint{	{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil }}},
            { TargetValue: workspaceSetting,
             Constraints: []validation.Constraint{	{Target: "workspaceSetting.WorkspaceSettingProperties", Name: validation.Null, Rule: false ,
            Chain: []validation.Constraint{	{Target: "workspaceSetting.WorkspaceSettingProperties.WorkspaceID", Name: validation.Null, Rule: true, Chain: nil },
            	{Target: "workspaceSetting.WorkspaceSettingProperties.Scope", Name: validation.Null, Rule: true, Chain: nil },
            }}}}}); err != nil {
            return result, validation.NewError("security.WorkspaceSettingsClient", "Create", err.Error())
            }

                req, err := client.CreatePreparer(ctx, workspaceSettingName, workspaceSetting)
    if err != nil {
    err = autorest.NewErrorWithError(err, "security.WorkspaceSettingsClient", "Create", nil , "Failure preparing request")
    return
    }

            resp, err := client.CreateSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "security.WorkspaceSettingsClient", "Create", resp, "Failure sending request")
            return
            }

            result, err = client.CreateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "security.WorkspaceSettingsClient", "Create", resp, "Failure responding to request")
            }

    return
    }

    // CreatePreparer prepares the Create request.
    func (client WorkspaceSettingsClient) CreatePreparer(ctx context.Context, workspaceSettingName string, workspaceSetting WorkspaceSetting) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "workspaceSettingName": autorest.Encode("path",workspaceSettingName),
            }

                        const APIVersion = "2017-08-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPut(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/workspaceSettings/{workspaceSettingName}",pathParameters),
    autorest.WithJSON(workspaceSetting),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateSender sends the Create request. The method will close the
    // http.Response Body if it receives an error.
    func (client WorkspaceSettingsClient) CreateSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client WorkspaceSettingsClient) CreateResponder(resp *http.Response) (result WorkspaceSetting, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Delete deletes the custom workspace settings for this subscription. new VMs will report to the default workspace
    // Parameters:
        // workspaceSettingName - name of the security setting
func (client WorkspaceSettingsClient) Delete(ctx context.Context, workspaceSettingName string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/WorkspaceSettingsClient.Delete")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: client.SubscriptionID,
             Constraints: []validation.Constraint{	{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("security.WorkspaceSettingsClient", "Delete", err.Error())
            }

                req, err := client.DeletePreparer(ctx, workspaceSettingName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "security.WorkspaceSettingsClient", "Delete", nil , "Failure preparing request")
    return
    }

            resp, err := client.DeleteSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "security.WorkspaceSettingsClient", "Delete", resp, "Failure sending request")
            return
            }

            result, err = client.DeleteResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "security.WorkspaceSettingsClient", "Delete", resp, "Failure responding to request")
            }

    return
    }

    // DeletePreparer prepares the Delete request.
    func (client WorkspaceSettingsClient) DeletePreparer(ctx context.Context, workspaceSettingName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "workspaceSettingName": autorest.Encode("path",workspaceSettingName),
            }

                        const APIVersion = "2017-08-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsDelete(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/workspaceSettings/{workspaceSettingName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client WorkspaceSettingsClient) DeleteSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client WorkspaceSettingsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Get settings about where we should store your security data and logs. If the result is empty, it means that no
// custom-workspace configuration was set
    // Parameters:
        // workspaceSettingName - name of the security setting
func (client WorkspaceSettingsClient) Get(ctx context.Context, workspaceSettingName string) (result WorkspaceSetting, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/WorkspaceSettingsClient.Get")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: client.SubscriptionID,
             Constraints: []validation.Constraint{	{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("security.WorkspaceSettingsClient", "Get", err.Error())
            }

                req, err := client.GetPreparer(ctx, workspaceSettingName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "security.WorkspaceSettingsClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "security.WorkspaceSettingsClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "security.WorkspaceSettingsClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client WorkspaceSettingsClient) GetPreparer(ctx context.Context, workspaceSettingName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "workspaceSettingName": autorest.Encode("path",workspaceSettingName),
            }

                        const APIVersion = "2017-08-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/workspaceSettings/{workspaceSettingName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client WorkspaceSettingsClient) GetSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client WorkspaceSettingsClient) GetResponder(resp *http.Response) (result WorkspaceSetting, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// List settings about where we should store your security data and logs. If the result is empty, it means that no
// custom-workspace configuration was set
func (client WorkspaceSettingsClient) List(ctx context.Context) (result WorkspaceSettingListPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/WorkspaceSettingsClient.List")
        defer func() {
            sc := -1
            if result.wsl.Response.Response != nil {
                sc = result.wsl.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: client.SubscriptionID,
             Constraints: []validation.Constraint{	{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("security.WorkspaceSettingsClient", "List", err.Error())
            }

                        result.fn = client.listNextResults
    req, err := client.ListPreparer(ctx)
    if err != nil {
    err = autorest.NewErrorWithError(err, "security.WorkspaceSettingsClient", "List", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListSender(req)
            if err != nil {
            result.wsl.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "security.WorkspaceSettingsClient", "List", resp, "Failure sending request")
            return
            }

            result.wsl, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "security.WorkspaceSettingsClient", "List", resp, "Failure responding to request")
            }

    return
    }

    // ListPreparer prepares the List request.
    func (client WorkspaceSettingsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2017-08-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/workspaceSettings",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client WorkspaceSettingsClient) ListSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client WorkspaceSettingsClient) ListResponder(resp *http.Response) (result WorkspaceSettingList, err error) {
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
            func (client WorkspaceSettingsClient) listNextResults(ctx context.Context, lastResults WorkspaceSettingList) (result WorkspaceSettingList, err error) {
            req, err := lastResults.workspaceSettingListPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "security.WorkspaceSettingsClient", "listNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "security.WorkspaceSettingsClient", "listNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "security.WorkspaceSettingsClient", "listNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListComplete enumerates all values, automatically crossing page boundaries as required.
    func (client WorkspaceSettingsClient) ListComplete(ctx context.Context) (result WorkspaceSettingListIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/WorkspaceSettingsClient.List")
            defer func() {
                sc := -1
                if result.Response().Response.Response != nil {
                    sc = result.page.Response().Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
     }
        result.page, err = client.List(ctx)
                return
        }

// Update settings about where we should store your security data and logs
    // Parameters:
        // workspaceSettingName - name of the security setting
        // workspaceSetting - security data setting object
func (client WorkspaceSettingsClient) Update(ctx context.Context, workspaceSettingName string, workspaceSetting WorkspaceSetting) (result WorkspaceSetting, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/WorkspaceSettingsClient.Update")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: client.SubscriptionID,
             Constraints: []validation.Constraint{	{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("security.WorkspaceSettingsClient", "Update", err.Error())
            }

                req, err := client.UpdatePreparer(ctx, workspaceSettingName, workspaceSetting)
    if err != nil {
    err = autorest.NewErrorWithError(err, "security.WorkspaceSettingsClient", "Update", nil , "Failure preparing request")
    return
    }

            resp, err := client.UpdateSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "security.WorkspaceSettingsClient", "Update", resp, "Failure sending request")
            return
            }

            result, err = client.UpdateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "security.WorkspaceSettingsClient", "Update", resp, "Failure responding to request")
            }

    return
    }

    // UpdatePreparer prepares the Update request.
    func (client WorkspaceSettingsClient) UpdatePreparer(ctx context.Context, workspaceSettingName string, workspaceSetting WorkspaceSetting) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "workspaceSettingName": autorest.Encode("path",workspaceSettingName),
            }

                        const APIVersion = "2017-08-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPatch(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/workspaceSettings/{workspaceSettingName}",pathParameters),
    autorest.WithJSON(workspaceSetting),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // UpdateSender sends the Update request. The method will close the
    // http.Response Body if it receives an error.
    func (client WorkspaceSettingsClient) UpdateSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client WorkspaceSettingsClient) UpdateResponder(resp *http.Response) (result WorkspaceSetting, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

