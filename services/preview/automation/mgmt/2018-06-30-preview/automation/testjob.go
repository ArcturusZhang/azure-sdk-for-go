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
)

// TestJobClient is the automation Client
type TestJobClient struct {
    BaseClient
}
// NewTestJobClient creates an instance of the TestJobClient client.
func NewTestJobClient(subscriptionID string) TestJobClient {
    return NewTestJobClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewTestJobClientWithBaseURI creates an instance of the TestJobClient client.
    func NewTestJobClientWithBaseURI(baseURI string, subscriptionID string) TestJobClient {
        return TestJobClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// Create create a test job of the runbook.
    // Parameters:
        // resourceGroupName - name of an Azure Resource group.
        // automationAccountName - the name of the automation account.
        // runbookName - the parameters supplied to the create test job operation.
        // parameters - the parameters supplied to the create test job operation.
func (client TestJobClient) Create(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, parameters TestJobCreateParameters) (result TestJob, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/TestJobClient.Create")
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
            return result, validation.NewError("automation.TestJobClient", "Create", err.Error())
            }

                req, err := client.CreatePreparer(ctx, resourceGroupName, automationAccountName, runbookName, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "automation.TestJobClient", "Create", nil , "Failure preparing request")
    return
    }

            resp, err := client.CreateSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "automation.TestJobClient", "Create", resp, "Failure sending request")
            return
            }

            result, err = client.CreateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "automation.TestJobClient", "Create", resp, "Failure responding to request")
            }

    return
    }

    // CreatePreparer prepares the Create request.
    func (client TestJobClient) CreatePreparer(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, parameters TestJobCreateParameters) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "automationAccountName": autorest.Encode("path",automationAccountName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "runbookName": autorest.Encode("path",runbookName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-06-30"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPut(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/testJob",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateSender sends the Create request. The method will close the
    // http.Response Body if it receives an error.
    func (client TestJobClient) CreateSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client TestJobClient) CreateResponder(resp *http.Response) (result TestJob, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Get retrieve the test job for the specified runbook.
    // Parameters:
        // resourceGroupName - name of an Azure Resource group.
        // automationAccountName - the name of the automation account.
        // runbookName - the runbook name.
func (client TestJobClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string) (result TestJob, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/TestJobClient.Get")
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
            return result, validation.NewError("automation.TestJobClient", "Get", err.Error())
            }

                req, err := client.GetPreparer(ctx, resourceGroupName, automationAccountName, runbookName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "automation.TestJobClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "automation.TestJobClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "automation.TestJobClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client TestJobClient) GetPreparer(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "automationAccountName": autorest.Encode("path",automationAccountName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "runbookName": autorest.Encode("path",runbookName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-06-30"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/testJob",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client TestJobClient) GetSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client TestJobClient) GetResponder(resp *http.Response) (result TestJob, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Resume resume the test job.
    // Parameters:
        // resourceGroupName - name of an Azure Resource group.
        // automationAccountName - the name of the automation account.
        // runbookName - the runbook name.
func (client TestJobClient) Resume(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/TestJobClient.Resume")
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
            	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("automation.TestJobClient", "Resume", err.Error())
            }

                req, err := client.ResumePreparer(ctx, resourceGroupName, automationAccountName, runbookName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "automation.TestJobClient", "Resume", nil , "Failure preparing request")
    return
    }

            resp, err := client.ResumeSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "automation.TestJobClient", "Resume", resp, "Failure sending request")
            return
            }

            result, err = client.ResumeResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "automation.TestJobClient", "Resume", resp, "Failure responding to request")
            }

    return
    }

    // ResumePreparer prepares the Resume request.
    func (client TestJobClient) ResumePreparer(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "automationAccountName": autorest.Encode("path",automationAccountName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "runbookName": autorest.Encode("path",runbookName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-06-30"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/testJob/resume",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ResumeSender sends the Resume request. The method will close the
    // http.Response Body if it receives an error.
    func (client TestJobClient) ResumeSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ResumeResponder handles the response to the Resume request. The method always
// closes the http.Response Body.
func (client TestJobClient) ResumeResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Stop stop the test job.
    // Parameters:
        // resourceGroupName - name of an Azure Resource group.
        // automationAccountName - the name of the automation account.
        // runbookName - the runbook name.
func (client TestJobClient) Stop(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/TestJobClient.Stop")
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
            	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("automation.TestJobClient", "Stop", err.Error())
            }

                req, err := client.StopPreparer(ctx, resourceGroupName, automationAccountName, runbookName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "automation.TestJobClient", "Stop", nil , "Failure preparing request")
    return
    }

            resp, err := client.StopSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "automation.TestJobClient", "Stop", resp, "Failure sending request")
            return
            }

            result, err = client.StopResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "automation.TestJobClient", "Stop", resp, "Failure responding to request")
            }

    return
    }

    // StopPreparer prepares the Stop request.
    func (client TestJobClient) StopPreparer(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "automationAccountName": autorest.Encode("path",automationAccountName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "runbookName": autorest.Encode("path",runbookName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-06-30"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/testJob/stop",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // StopSender sends the Stop request. The method will close the
    // http.Response Body if it receives an error.
    func (client TestJobClient) StopSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// StopResponder handles the response to the Stop request. The method always
// closes the http.Response Body.
func (client TestJobClient) StopResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Suspend suspend the test job.
    // Parameters:
        // resourceGroupName - name of an Azure Resource group.
        // automationAccountName - the name of the automation account.
        // runbookName - the runbook name.
func (client TestJobClient) Suspend(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/TestJobClient.Suspend")
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
            	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("automation.TestJobClient", "Suspend", err.Error())
            }

                req, err := client.SuspendPreparer(ctx, resourceGroupName, automationAccountName, runbookName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "automation.TestJobClient", "Suspend", nil , "Failure preparing request")
    return
    }

            resp, err := client.SuspendSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "automation.TestJobClient", "Suspend", resp, "Failure sending request")
            return
            }

            result, err = client.SuspendResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "automation.TestJobClient", "Suspend", resp, "Failure responding to request")
            }

    return
    }

    // SuspendPreparer prepares the Suspend request.
    func (client TestJobClient) SuspendPreparer(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "automationAccountName": autorest.Encode("path",automationAccountName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "runbookName": autorest.Encode("path",runbookName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-06-30"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/testJob/suspend",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // SuspendSender sends the Suspend request. The method will close the
    // http.Response Body if it receives an error.
    func (client TestJobClient) SuspendSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// SuspendResponder handles the response to the Suspend request. The method always
// closes the http.Response Body.
func (client TestJobClient) SuspendResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByClosing())
    result.Response = resp
        return
    }

