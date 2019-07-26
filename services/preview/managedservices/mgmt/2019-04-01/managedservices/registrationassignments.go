package managedservices

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

// RegistrationAssignmentsClient is the specification for ManagedServices.
type RegistrationAssignmentsClient struct {
    BaseClient
}
// NewRegistrationAssignmentsClient creates an instance of the RegistrationAssignmentsClient client.
func NewRegistrationAssignmentsClient() RegistrationAssignmentsClient {
    return NewRegistrationAssignmentsClientWithBaseURI(DefaultBaseURI, )
}

// NewRegistrationAssignmentsClientWithBaseURI creates an instance of the RegistrationAssignmentsClient client.
    func NewRegistrationAssignmentsClientWithBaseURI(baseURI string, ) RegistrationAssignmentsClient {
        return RegistrationAssignmentsClient{ NewWithBaseURI(baseURI, )}
    }

// CreateOrUpdate creates or updates a registration assignment.
    // Parameters:
        // scope - scope of the resource.
        // registrationAssignmentID - guid of the registration assignment.
        // requestBody - the parameters required to create new registration assignment.
func (client RegistrationAssignmentsClient) CreateOrUpdate(ctx context.Context, scope string, registrationAssignmentID string, requestBody RegistrationAssignment) (result RegistrationAssignmentsCreateOrUpdateFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/RegistrationAssignmentsClient.CreateOrUpdate")
        defer func() {
            sc := -1
            if result.Response() != nil {
                sc = result.Response().StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: requestBody,
             Constraints: []validation.Constraint{	{Target: "requestBody.Properties", Name: validation.Null, Rule: false ,
            Chain: []validation.Constraint{	{Target: "requestBody.Properties.RegistrationDefinitionID", Name: validation.Null, Rule: true, Chain: nil },
            	{Target: "requestBody.Properties.RegistrationDefinition", Name: validation.Null, Rule: false ,
            Chain: []validation.Constraint{	{Target: "requestBody.Properties.RegistrationDefinition.Plan", Name: validation.Null, Rule: false ,
            Chain: []validation.Constraint{	{Target: "requestBody.Properties.RegistrationDefinition.Plan.Name", Name: validation.Null, Rule: true, Chain: nil },
            	{Target: "requestBody.Properties.RegistrationDefinition.Plan.Publisher", Name: validation.Null, Rule: true, Chain: nil },
            	{Target: "requestBody.Properties.RegistrationDefinition.Plan.Product", Name: validation.Null, Rule: true, Chain: nil },
            	{Target: "requestBody.Properties.RegistrationDefinition.Plan.Version", Name: validation.Null, Rule: true, Chain: nil },
            }},
            }},
            }}}}}); err != nil {
            return result, validation.NewError("managedservices.RegistrationAssignmentsClient", "CreateOrUpdate", err.Error())
            }

                req, err := client.CreateOrUpdatePreparer(ctx, scope, registrationAssignmentID, requestBody)
    if err != nil {
    err = autorest.NewErrorWithError(err, "managedservices.RegistrationAssignmentsClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

            result, err = client.CreateOrUpdateSender(req)
            if err != nil {
            err = autorest.NewErrorWithError(err, "managedservices.RegistrationAssignmentsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
            return
            }

    return
    }

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client RegistrationAssignmentsClient) CreateOrUpdatePreparer(ctx context.Context, scope string, registrationAssignmentID string, requestBody RegistrationAssignment) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "registrationAssignmentId": autorest.Encode("path",registrationAssignmentID),
            "scope": scope,
            }

                        const APIVersion = "2019-04-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

                    requestBody.ID = nil
                requestBody.Type = nil
                requestBody.Name = nil
    preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPut(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/{scope}/providers/Microsoft.ManagedServices/registrationAssignments/{registrationAssignmentId}",pathParameters),
    autorest.WithJSON(requestBody),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client RegistrationAssignmentsClient) CreateOrUpdateSender(req *http.Request) (future RegistrationAssignmentsCreateOrUpdateFuture, err error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
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
func (client RegistrationAssignmentsClient) CreateOrUpdateResponder(resp *http.Response) (result RegistrationAssignment, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Delete deletes the specified registration assignment.
    // Parameters:
        // scope - scope of the resource.
        // registrationAssignmentID - guid of the registration assignment.
func (client RegistrationAssignmentsClient) Delete(ctx context.Context, scope string, registrationAssignmentID string) (result RegistrationAssignmentsDeleteFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/RegistrationAssignmentsClient.Delete")
        defer func() {
            sc := -1
            if result.Response() != nil {
                sc = result.Response().StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.DeletePreparer(ctx, scope, registrationAssignmentID)
    if err != nil {
    err = autorest.NewErrorWithError(err, "managedservices.RegistrationAssignmentsClient", "Delete", nil , "Failure preparing request")
    return
    }

            result, err = client.DeleteSender(req)
            if err != nil {
            err = autorest.NewErrorWithError(err, "managedservices.RegistrationAssignmentsClient", "Delete", result.Response(), "Failure sending request")
            return
            }

    return
    }

    // DeletePreparer prepares the Delete request.
    func (client RegistrationAssignmentsClient) DeletePreparer(ctx context.Context, scope string, registrationAssignmentID string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "registrationAssignmentId": autorest.Encode("path",registrationAssignmentID),
            "scope": scope,
            }

                        const APIVersion = "2019-04-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsDelete(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/{scope}/providers/Microsoft.ManagedServices/registrationAssignments/{registrationAssignmentId}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client RegistrationAssignmentsClient) DeleteSender(req *http.Request) (future RegistrationAssignmentsDeleteFuture, err error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
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
func (client RegistrationAssignmentsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Get gets the details of specified registration assignment.
    // Parameters:
        // scope - scope of the resource.
        // registrationAssignmentID - guid of the registration assignment.
        // expandRegistrationDefinition - tells whether to return registration definition details also along with
        // registration assignment details.
func (client RegistrationAssignmentsClient) Get(ctx context.Context, scope string, registrationAssignmentID string, expandRegistrationDefinition *bool) (result RegistrationAssignment, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/RegistrationAssignmentsClient.Get")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.GetPreparer(ctx, scope, registrationAssignmentID, expandRegistrationDefinition)
    if err != nil {
    err = autorest.NewErrorWithError(err, "managedservices.RegistrationAssignmentsClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "managedservices.RegistrationAssignmentsClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "managedservices.RegistrationAssignmentsClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client RegistrationAssignmentsClient) GetPreparer(ctx context.Context, scope string, registrationAssignmentID string, expandRegistrationDefinition *bool) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "registrationAssignmentId": autorest.Encode("path",registrationAssignmentID),
            "scope": scope,
            }

                        const APIVersion = "2019-04-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }
            if expandRegistrationDefinition != nil {
            queryParameters["$expandRegistrationDefinition"] = autorest.Encode("query",*expandRegistrationDefinition)
            }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/{scope}/providers/Microsoft.ManagedServices/registrationAssignments/{registrationAssignmentId}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client RegistrationAssignmentsClient) GetSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RegistrationAssignmentsClient) GetResponder(resp *http.Response) (result RegistrationAssignment, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// List gets a list of the registration assignments.
    // Parameters:
        // scope - scope of the resource.
        // expandRegistrationDefinition - tells whether to return registration definition details also along with
        // registration assignment details.
func (client RegistrationAssignmentsClient) List(ctx context.Context, scope string, expandRegistrationDefinition *bool) (result RegistrationAssignmentListPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/RegistrationAssignmentsClient.List")
        defer func() {
            sc := -1
            if result.ral.Response.Response != nil {
                sc = result.ral.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
                result.fn = client.listNextResults
    req, err := client.ListPreparer(ctx, scope, expandRegistrationDefinition)
    if err != nil {
    err = autorest.NewErrorWithError(err, "managedservices.RegistrationAssignmentsClient", "List", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListSender(req)
            if err != nil {
            result.ral.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "managedservices.RegistrationAssignmentsClient", "List", resp, "Failure sending request")
            return
            }

            result.ral, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "managedservices.RegistrationAssignmentsClient", "List", resp, "Failure responding to request")
            }

    return
    }

    // ListPreparer prepares the List request.
    func (client RegistrationAssignmentsClient) ListPreparer(ctx context.Context, scope string, expandRegistrationDefinition *bool) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "scope": scope,
            }

                        const APIVersion = "2019-04-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }
            if expandRegistrationDefinition != nil {
            queryParameters["$expandRegistrationDefinition"] = autorest.Encode("query",*expandRegistrationDefinition)
            }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/{scope}/providers/Microsoft.ManagedServices/registrationAssignments",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client RegistrationAssignmentsClient) ListSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client RegistrationAssignmentsClient) ListResponder(resp *http.Response) (result RegistrationAssignmentList, err error) {
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
            func (client RegistrationAssignmentsClient) listNextResults(ctx context.Context, lastResults RegistrationAssignmentList) (result RegistrationAssignmentList, err error) {
            req, err := lastResults.registrationAssignmentListPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "managedservices.RegistrationAssignmentsClient", "listNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "managedservices.RegistrationAssignmentsClient", "listNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "managedservices.RegistrationAssignmentsClient", "listNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListComplete enumerates all values, automatically crossing page boundaries as required.
    func (client RegistrationAssignmentsClient) ListComplete(ctx context.Context, scope string, expandRegistrationDefinition *bool) (result RegistrationAssignmentListIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/RegistrationAssignmentsClient.List")
            defer func() {
                sc := -1
                if result.Response().Response.Response != nil {
                    sc = result.page.Response().Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
     }
        result.page, err = client.List(ctx, scope, expandRegistrationDefinition)
                return
        }

