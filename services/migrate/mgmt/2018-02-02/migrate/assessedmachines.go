package migrate

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

// AssessedMachinesClient is the move your workloads to Azure.
type AssessedMachinesClient struct {
    BaseClient
}
// NewAssessedMachinesClient creates an instance of the AssessedMachinesClient client.
func NewAssessedMachinesClient(subscriptionID string, acceptLanguage string) AssessedMachinesClient {
    return NewAssessedMachinesClientWithBaseURI(DefaultBaseURI, subscriptionID, acceptLanguage)
}

// NewAssessedMachinesClientWithBaseURI creates an instance of the AssessedMachinesClient client.
    func NewAssessedMachinesClientWithBaseURI(baseURI string, subscriptionID string, acceptLanguage string) AssessedMachinesClient {
        return AssessedMachinesClient{ NewWithBaseURI(baseURI, subscriptionID, acceptLanguage)}
    }

// Get get an assessed machine with its size & cost estimate that was evaluated in the specified assessment.
    // Parameters:
        // resourceGroupName - name of the Azure Resource Group that project is part of.
        // projectName - name of the Azure Migrate project.
        // groupName - unique name of a group within a project.
        // assessmentName - unique name of an assessment within a project.
        // assessedMachineName - unique name of an assessed machine evaluated as part of an assessment.
func (client AssessedMachinesClient) Get(ctx context.Context, resourceGroupName string, projectName string, groupName string, assessmentName string, assessedMachineName string) (result AssessedMachine, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/AssessedMachinesClient.Get")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.GetPreparer(ctx, resourceGroupName, projectName, groupName, assessmentName, assessedMachineName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "migrate.AssessedMachinesClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "migrate.AssessedMachinesClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "migrate.AssessedMachinesClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client AssessedMachinesClient) GetPreparer(ctx context.Context, resourceGroupName string, projectName string, groupName string, assessmentName string, assessedMachineName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "assessedMachineName": autorest.Encode("path",assessedMachineName),
            "assessmentName": autorest.Encode("path",assessmentName),
            "groupName": autorest.Encode("path",groupName),
            "projectName": autorest.Encode("path",projectName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-02-02"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}/groups/{groupName}/assessments/{assessmentName}/assessedMachines/{assessedMachineName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
            if len(client.AcceptLanguage) > 0 {
            preparer = autorest.DecoratePreparer(preparer,
            autorest.WithHeader("Accept-Language",autorest.String(client.AcceptLanguage)))
            }
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client AssessedMachinesClient) GetSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AssessedMachinesClient) GetResponder(resp *http.Response) (result AssessedMachine, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// ListByAssessment get list of machines that assessed as part of the specified assessment. Returns a json array of
// objects of type 'assessedMachine' as specified in the Models section.
//
// Whenever an assessment is created or updated, it goes under computation. During this phase, the 'status' field of
// Assessment object reports 'Computing'.
// During the period when the assessment is under computation, the list of assessed machines is empty and no assessed
// machines are returned by this call.
    // Parameters:
        // resourceGroupName - name of the Azure Resource Group that project is part of.
        // projectName - name of the Azure Migrate project.
        // groupName - unique name of a group within a project.
        // assessmentName - unique name of an assessment within a project.
func (client AssessedMachinesClient) ListByAssessment(ctx context.Context, resourceGroupName string, projectName string, groupName string, assessmentName string) (result AssessedMachineResultList, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/AssessedMachinesClient.ListByAssessment")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.ListByAssessmentPreparer(ctx, resourceGroupName, projectName, groupName, assessmentName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "migrate.AssessedMachinesClient", "ListByAssessment", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListByAssessmentSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "migrate.AssessedMachinesClient", "ListByAssessment", resp, "Failure sending request")
            return
            }

            result, err = client.ListByAssessmentResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "migrate.AssessedMachinesClient", "ListByAssessment", resp, "Failure responding to request")
            }

    return
    }

    // ListByAssessmentPreparer prepares the ListByAssessment request.
    func (client AssessedMachinesClient) ListByAssessmentPreparer(ctx context.Context, resourceGroupName string, projectName string, groupName string, assessmentName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "assessmentName": autorest.Encode("path",assessmentName),
            "groupName": autorest.Encode("path",groupName),
            "projectName": autorest.Encode("path",projectName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-02-02"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}/groups/{groupName}/assessments/{assessmentName}/assessedMachines",pathParameters),
    autorest.WithQueryParameters(queryParameters))
            if len(client.AcceptLanguage) > 0 {
            preparer = autorest.DecoratePreparer(preparer,
            autorest.WithHeader("Accept-Language",autorest.String(client.AcceptLanguage)))
            }
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByAssessmentSender sends the ListByAssessment request. The method will close the
    // http.Response Body if it receives an error.
    func (client AssessedMachinesClient) ListByAssessmentSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListByAssessmentResponder handles the response to the ListByAssessment request. The method always
// closes the http.Response Body.
func (client AssessedMachinesClient) ListByAssessmentResponder(resp *http.Response) (result AssessedMachineResultList, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

