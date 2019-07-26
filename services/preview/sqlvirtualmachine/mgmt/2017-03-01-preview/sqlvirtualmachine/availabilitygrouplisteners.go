package sqlvirtualmachine

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

// AvailabilityGroupListenersClient is the the SQL virtual machine management API provides a RESTful set of web APIs
// that interact with Azure Compute, Network & Storage services to manage your SQL Server virtual machine. The API
// enables users to create, delete and retrieve a SQL virtual machine, SQL virtual machine group or availability group
// listener.
type AvailabilityGroupListenersClient struct {
    BaseClient
}
// NewAvailabilityGroupListenersClient creates an instance of the AvailabilityGroupListenersClient client.
func NewAvailabilityGroupListenersClient(subscriptionID string) AvailabilityGroupListenersClient {
    return NewAvailabilityGroupListenersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAvailabilityGroupListenersClientWithBaseURI creates an instance of the AvailabilityGroupListenersClient client.
    func NewAvailabilityGroupListenersClientWithBaseURI(baseURI string, subscriptionID string) AvailabilityGroupListenersClient {
        return AvailabilityGroupListenersClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CreateOrUpdate creates or updates an availability group listener.
    // Parameters:
        // resourceGroupName - name of the resource group that contains the resource. You can obtain this value from
        // the Azure Resource Manager API or the portal.
        // SQLVirtualMachineGroupName - name of the SQL virtual machine group.
        // availabilityGroupListenerName - name of the availability group listener.
        // parameters - the availability group listener.
func (client AvailabilityGroupListenersClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, SQLVirtualMachineGroupName string, availabilityGroupListenerName string, parameters AvailabilityGroupListener) (result AvailabilityGroupListenersCreateOrUpdateFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/AvailabilityGroupListenersClient.CreateOrUpdate")
        defer func() {
            sc := -1
            if result.Response() != nil {
                sc = result.Response().StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, SQLVirtualMachineGroupName, availabilityGroupListenerName, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "sqlvirtualmachine.AvailabilityGroupListenersClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

            result, err = client.CreateOrUpdateSender(req)
            if err != nil {
            err = autorest.NewErrorWithError(err, "sqlvirtualmachine.AvailabilityGroupListenersClient", "CreateOrUpdate", result.Response(), "Failure sending request")
            return
            }

    return
    }

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client AvailabilityGroupListenersClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, SQLVirtualMachineGroupName string, availabilityGroupListenerName string, parameters AvailabilityGroupListener) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "availabilityGroupListenerName": autorest.Encode("path",availabilityGroupListenerName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "sqlVirtualMachineGroupName": autorest.Encode("path",SQLVirtualMachineGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2017-03-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPut(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/{sqlVirtualMachineGroupName}/availabilityGroupListeners/{availabilityGroupListenerName}",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client AvailabilityGroupListenersClient) CreateOrUpdateSender(req *http.Request) (future AvailabilityGroupListenersCreateOrUpdateFuture, err error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
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
func (client AvailabilityGroupListenersClient) CreateOrUpdateResponder(resp *http.Response) (result AvailabilityGroupListener, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Delete deletes an availability group listener.
    // Parameters:
        // resourceGroupName - name of the resource group that contains the resource. You can obtain this value from
        // the Azure Resource Manager API or the portal.
        // SQLVirtualMachineGroupName - name of the SQL virtual machine group.
        // availabilityGroupListenerName - name of the availability group listener.
func (client AvailabilityGroupListenersClient) Delete(ctx context.Context, resourceGroupName string, SQLVirtualMachineGroupName string, availabilityGroupListenerName string) (result AvailabilityGroupListenersDeleteFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/AvailabilityGroupListenersClient.Delete")
        defer func() {
            sc := -1
            if result.Response() != nil {
                sc = result.Response().StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.DeletePreparer(ctx, resourceGroupName, SQLVirtualMachineGroupName, availabilityGroupListenerName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "sqlvirtualmachine.AvailabilityGroupListenersClient", "Delete", nil , "Failure preparing request")
    return
    }

            result, err = client.DeleteSender(req)
            if err != nil {
            err = autorest.NewErrorWithError(err, "sqlvirtualmachine.AvailabilityGroupListenersClient", "Delete", result.Response(), "Failure sending request")
            return
            }

    return
    }

    // DeletePreparer prepares the Delete request.
    func (client AvailabilityGroupListenersClient) DeletePreparer(ctx context.Context, resourceGroupName string, SQLVirtualMachineGroupName string, availabilityGroupListenerName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "availabilityGroupListenerName": autorest.Encode("path",availabilityGroupListenerName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "sqlVirtualMachineGroupName": autorest.Encode("path",SQLVirtualMachineGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2017-03-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsDelete(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/{sqlVirtualMachineGroupName}/availabilityGroupListeners/{availabilityGroupListenerName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client AvailabilityGroupListenersClient) DeleteSender(req *http.Request) (future AvailabilityGroupListenersDeleteFuture, err error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
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
func (client AvailabilityGroupListenersClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Get gets an availability group listener.
    // Parameters:
        // resourceGroupName - name of the resource group that contains the resource. You can obtain this value from
        // the Azure Resource Manager API or the portal.
        // SQLVirtualMachineGroupName - name of the SQL virtual machine group.
        // availabilityGroupListenerName - name of the availability group listener.
func (client AvailabilityGroupListenersClient) Get(ctx context.Context, resourceGroupName string, SQLVirtualMachineGroupName string, availabilityGroupListenerName string) (result AvailabilityGroupListener, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/AvailabilityGroupListenersClient.Get")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.GetPreparer(ctx, resourceGroupName, SQLVirtualMachineGroupName, availabilityGroupListenerName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "sqlvirtualmachine.AvailabilityGroupListenersClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "sqlvirtualmachine.AvailabilityGroupListenersClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "sqlvirtualmachine.AvailabilityGroupListenersClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client AvailabilityGroupListenersClient) GetPreparer(ctx context.Context, resourceGroupName string, SQLVirtualMachineGroupName string, availabilityGroupListenerName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "availabilityGroupListenerName": autorest.Encode("path",availabilityGroupListenerName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "sqlVirtualMachineGroupName": autorest.Encode("path",SQLVirtualMachineGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2017-03-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/{sqlVirtualMachineGroupName}/availabilityGroupListeners/{availabilityGroupListenerName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client AvailabilityGroupListenersClient) GetSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AvailabilityGroupListenersClient) GetResponder(resp *http.Response) (result AvailabilityGroupListener, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// ListByGroup lists all availability group listeners in a SQL virtual machine group.
    // Parameters:
        // resourceGroupName - name of the resource group that contains the resource. You can obtain this value from
        // the Azure Resource Manager API or the portal.
        // SQLVirtualMachineGroupName - name of the SQL virtual machine group.
func (client AvailabilityGroupListenersClient) ListByGroup(ctx context.Context, resourceGroupName string, SQLVirtualMachineGroupName string) (result AvailabilityGroupListenerListResultPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/AvailabilityGroupListenersClient.ListByGroup")
        defer func() {
            sc := -1
            if result.agllr.Response.Response != nil {
                sc = result.agllr.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
                result.fn = client.listByGroupNextResults
    req, err := client.ListByGroupPreparer(ctx, resourceGroupName, SQLVirtualMachineGroupName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "sqlvirtualmachine.AvailabilityGroupListenersClient", "ListByGroup", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListByGroupSender(req)
            if err != nil {
            result.agllr.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "sqlvirtualmachine.AvailabilityGroupListenersClient", "ListByGroup", resp, "Failure sending request")
            return
            }

            result.agllr, err = client.ListByGroupResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "sqlvirtualmachine.AvailabilityGroupListenersClient", "ListByGroup", resp, "Failure responding to request")
            }

    return
    }

    // ListByGroupPreparer prepares the ListByGroup request.
    func (client AvailabilityGroupListenersClient) ListByGroupPreparer(ctx context.Context, resourceGroupName string, SQLVirtualMachineGroupName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "sqlVirtualMachineGroupName": autorest.Encode("path",SQLVirtualMachineGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2017-03-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/{sqlVirtualMachineGroupName}/availabilityGroupListeners",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByGroupSender sends the ListByGroup request. The method will close the
    // http.Response Body if it receives an error.
    func (client AvailabilityGroupListenersClient) ListByGroupSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListByGroupResponder handles the response to the ListByGroup request. The method always
// closes the http.Response Body.
func (client AvailabilityGroupListenersClient) ListByGroupResponder(resp *http.Response) (result AvailabilityGroupListenerListResult, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

            // listByGroupNextResults retrieves the next set of results, if any.
            func (client AvailabilityGroupListenersClient) listByGroupNextResults(ctx context.Context, lastResults AvailabilityGroupListenerListResult) (result AvailabilityGroupListenerListResult, err error) {
            req, err := lastResults.availabilityGroupListenerListResultPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "sqlvirtualmachine.AvailabilityGroupListenersClient", "listByGroupNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListByGroupSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "sqlvirtualmachine.AvailabilityGroupListenersClient", "listByGroupNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListByGroupResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "sqlvirtualmachine.AvailabilityGroupListenersClient", "listByGroupNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListByGroupComplete enumerates all values, automatically crossing page boundaries as required.
    func (client AvailabilityGroupListenersClient) ListByGroupComplete(ctx context.Context, resourceGroupName string, SQLVirtualMachineGroupName string) (result AvailabilityGroupListenerListResultIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/AvailabilityGroupListenersClient.ListByGroup")
            defer func() {
                sc := -1
                if result.Response().Response.Response != nil {
                    sc = result.page.Response().Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
     }
        result.page, err = client.ListByGroup(ctx, resourceGroupName, SQLVirtualMachineGroupName)
                return
        }

