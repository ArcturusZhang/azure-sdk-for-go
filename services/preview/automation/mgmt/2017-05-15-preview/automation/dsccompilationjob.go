package automation

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
    "github.com/satori/go.uuid"
)

// DscCompilationJobClient is the automation Client
type DscCompilationJobClient struct {
    BaseClient
}
// NewDscCompilationJobClient creates an instance of the DscCompilationJobClient client.
func NewDscCompilationJobClient(subscriptionID string) DscCompilationJobClient {
    return NewDscCompilationJobClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDscCompilationJobClientWithBaseURI creates an instance of the DscCompilationJobClient client.
    func NewDscCompilationJobClientWithBaseURI(baseURI string, subscriptionID string) DscCompilationJobClient {
        return DscCompilationJobClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// Create creates the Dsc compilation job of the configuration.
    // Parameters:
        // resourceGroupName - name of an Azure Resource group.
        // automationAccountName - the name of the automation account.
        // compilationJobID - the the DSC configuration Id.
        // parameters - the parameters supplied to the create compilation job operation.
func (client DscCompilationJobClient) Create(ctx context.Context, resourceGroupName string, automationAccountName string, compilationJobID uuid.UUID, parameters DscCompilationJobCreateParameters) (result DscCompilationJob, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/DscCompilationJobClient.Create")
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
            	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil }}},
            { TargetValue: parameters,
             Constraints: []validation.Constraint{	{Target: "parameters.DscCompilationJobCreateProperties", Name: validation.Null, Rule: true ,
            Chain: []validation.Constraint{	{Target: "parameters.DscCompilationJobCreateProperties.Configuration", Name: validation.Null, Rule: true, Chain: nil },
            }}}}}); err != nil {
            return result, validation.NewError("automation.DscCompilationJobClient", "Create", err.Error())
            }

                req, err := client.CreatePreparer(ctx, resourceGroupName, automationAccountName, compilationJobID, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "automation.DscCompilationJobClient", "Create", nil , "Failure preparing request")
    return
    }

            resp, err := client.CreateSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "automation.DscCompilationJobClient", "Create", resp, "Failure sending request")
            return
            }

            result, err = client.CreateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "automation.DscCompilationJobClient", "Create", resp, "Failure responding to request")
            }

    return
    }

    // CreatePreparer prepares the Create request.
    func (client DscCompilationJobClient) CreatePreparer(ctx context.Context, resourceGroupName string, automationAccountName string, compilationJobID uuid.UUID, parameters DscCompilationJobCreateParameters) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "automationAccountName": autorest.Encode("path",automationAccountName),
            "compilationJobId": autorest.Encode("path",compilationJobID),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2015-10-31"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPut(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/compilationjobs/{compilationJobId}",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateSender sends the Create request. The method will close the
    // http.Response Body if it receives an error.
    func (client DscCompilationJobClient) CreateSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client DscCompilationJobClient) CreateResponder(resp *http.Response) (result DscCompilationJob, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Get retrieve the Dsc configuration compilation job identified by job id.
    // Parameters:
        // resourceGroupName - name of an Azure Resource group.
        // automationAccountName - the name of the automation account.
        // compilationJobID - the Dsc configuration compilation job id.
func (client DscCompilationJobClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, compilationJobID uuid.UUID) (result DscCompilationJob, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/DscCompilationJobClient.Get")
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
            	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("automation.DscCompilationJobClient", "Get", err.Error())
            }

                req, err := client.GetPreparer(ctx, resourceGroupName, automationAccountName, compilationJobID)
    if err != nil {
    err = autorest.NewErrorWithError(err, "automation.DscCompilationJobClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "automation.DscCompilationJobClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "automation.DscCompilationJobClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client DscCompilationJobClient) GetPreparer(ctx context.Context, resourceGroupName string, automationAccountName string, compilationJobID uuid.UUID) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "automationAccountName": autorest.Encode("path",automationAccountName),
            "compilationJobId": autorest.Encode("path",compilationJobID),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2015-10-31"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/compilationjobs/{compilationJobId}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client DscCompilationJobClient) GetSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DscCompilationJobClient) GetResponder(resp *http.Response) (result DscCompilationJob, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// GetStream retrieve the job stream identified by job stream id.
    // Parameters:
        // resourceGroupName - name of an Azure Resource group.
        // automationAccountName - the name of the automation account.
        // jobID - the job id.
        // jobStreamID - the job stream id.
func (client DscCompilationJobClient) GetStream(ctx context.Context, resourceGroupName string, automationAccountName string, jobID uuid.UUID, jobStreamID string) (result JobStream, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/DscCompilationJobClient.GetStream")
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
            	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("automation.DscCompilationJobClient", "GetStream", err.Error())
            }

                req, err := client.GetStreamPreparer(ctx, resourceGroupName, automationAccountName, jobID, jobStreamID)
    if err != nil {
    err = autorest.NewErrorWithError(err, "automation.DscCompilationJobClient", "GetStream", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetStreamSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "automation.DscCompilationJobClient", "GetStream", resp, "Failure sending request")
            return
            }

            result, err = client.GetStreamResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "automation.DscCompilationJobClient", "GetStream", resp, "Failure responding to request")
            }

    return
    }

    // GetStreamPreparer prepares the GetStream request.
    func (client DscCompilationJobClient) GetStreamPreparer(ctx context.Context, resourceGroupName string, automationAccountName string, jobID uuid.UUID, jobStreamID string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "automationAccountName": autorest.Encode("path",automationAccountName),
            "jobId": autorest.Encode("path",jobID),
            "jobStreamId": autorest.Encode("path",jobStreamID),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2015-10-31"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/compilationjobs/{jobId}/streams/{jobStreamId}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetStreamSender sends the GetStream request. The method will close the
    // http.Response Body if it receives an error.
    func (client DscCompilationJobClient) GetStreamSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetStreamResponder handles the response to the GetStream request. The method always
// closes the http.Response Body.
func (client DscCompilationJobClient) GetStreamResponder(resp *http.Response) (result JobStream, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// ListByAutomationAccount retrieve a list of dsc compilation jobs.
    // Parameters:
        // resourceGroupName - name of an Azure Resource group.
        // automationAccountName - the name of the automation account.
        // filter - the filter to apply on the operation.
func (client DscCompilationJobClient) ListByAutomationAccount(ctx context.Context, resourceGroupName string, automationAccountName string, filter string) (result DscCompilationJobListResultPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/DscCompilationJobClient.ListByAutomationAccount")
        defer func() {
            sc := -1
            if result.dcjlr.Response.Response != nil {
                sc = result.dcjlr.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: resourceGroupName,
             Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("automation.DscCompilationJobClient", "ListByAutomationAccount", err.Error())
            }

                        result.fn = client.listByAutomationAccountNextResults
    req, err := client.ListByAutomationAccountPreparer(ctx, resourceGroupName, automationAccountName, filter)
    if err != nil {
    err = autorest.NewErrorWithError(err, "automation.DscCompilationJobClient", "ListByAutomationAccount", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListByAutomationAccountSender(req)
            if err != nil {
            result.dcjlr.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "automation.DscCompilationJobClient", "ListByAutomationAccount", resp, "Failure sending request")
            return
            }

            result.dcjlr, err = client.ListByAutomationAccountResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "automation.DscCompilationJobClient", "ListByAutomationAccount", resp, "Failure responding to request")
            }

    return
    }

    // ListByAutomationAccountPreparer prepares the ListByAutomationAccount request.
    func (client DscCompilationJobClient) ListByAutomationAccountPreparer(ctx context.Context, resourceGroupName string, automationAccountName string, filter string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "automationAccountName": autorest.Encode("path",automationAccountName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2015-10-31"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }
            if len(filter) > 0 {
            queryParameters["$filter"] = autorest.Encode("query",filter)
            }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/compilationjobs",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByAutomationAccountSender sends the ListByAutomationAccount request. The method will close the
    // http.Response Body if it receives an error.
    func (client DscCompilationJobClient) ListByAutomationAccountSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListByAutomationAccountResponder handles the response to the ListByAutomationAccount request. The method always
// closes the http.Response Body.
func (client DscCompilationJobClient) ListByAutomationAccountResponder(resp *http.Response) (result DscCompilationJobListResult, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

            // listByAutomationAccountNextResults retrieves the next set of results, if any.
            func (client DscCompilationJobClient) listByAutomationAccountNextResults(ctx context.Context, lastResults DscCompilationJobListResult) (result DscCompilationJobListResult, err error) {
            req, err := lastResults.dscCompilationJobListResultPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "automation.DscCompilationJobClient", "listByAutomationAccountNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListByAutomationAccountSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "automation.DscCompilationJobClient", "listByAutomationAccountNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListByAutomationAccountResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "automation.DscCompilationJobClient", "listByAutomationAccountNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListByAutomationAccountComplete enumerates all values, automatically crossing page boundaries as required.
    func (client DscCompilationJobClient) ListByAutomationAccountComplete(ctx context.Context, resourceGroupName string, automationAccountName string, filter string) (result DscCompilationJobListResultIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/DscCompilationJobClient.ListByAutomationAccount")
            defer func() {
                sc := -1
                if result.Response().Response.Response != nil {
                    sc = result.page.Response().Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
     }
        result.page, err = client.ListByAutomationAccount(ctx, resourceGroupName, automationAccountName, filter)
                return
        }

