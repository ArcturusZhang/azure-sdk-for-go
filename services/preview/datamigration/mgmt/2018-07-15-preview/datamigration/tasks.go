package datamigration

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

// TasksClient is the data Migration Client
type TasksClient struct {
    BaseClient
}
// NewTasksClient creates an instance of the TasksClient client.
func NewTasksClient(subscriptionID string) TasksClient {
    return NewTasksClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewTasksClientWithBaseURI creates an instance of the TasksClient client.
    func NewTasksClientWithBaseURI(baseURI string, subscriptionID string) TasksClient {
        return TasksClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// Cancel the tasks resource is a nested, proxy-only resource representing work performed by a DMS instance. This
// method cancels a task if it's currently queued or running.
    // Parameters:
        // groupName - name of the resource group
        // serviceName - name of the service
        // projectName - name of the project
        // taskName - name of the Task
func (client TasksClient) Cancel(ctx context.Context, groupName string, serviceName string, projectName string, taskName string) (result ProjectTask, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/TasksClient.Cancel")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.CancelPreparer(ctx, groupName, serviceName, projectName, taskName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "datamigration.TasksClient", "Cancel", nil , "Failure preparing request")
    return
    }

            resp, err := client.CancelSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "datamigration.TasksClient", "Cancel", resp, "Failure sending request")
            return
            }

            result, err = client.CancelResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "datamigration.TasksClient", "Cancel", resp, "Failure responding to request")
            }

    return
    }

    // CancelPreparer prepares the Cancel request.
    func (client TasksClient) CancelPreparer(ctx context.Context, groupName string, serviceName string, projectName string, taskName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "groupName": autorest.Encode("path",groupName),
            "projectName": autorest.Encode("path",projectName),
            "serviceName": autorest.Encode("path",serviceName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "taskName": autorest.Encode("path",taskName),
            }

                        const APIVersion = "2018-07-15-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}/tasks/{taskName}/cancel",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CancelSender sends the Cancel request. The method will close the
    // http.Response Body if it receives an error.
    func (client TasksClient) CancelSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// CancelResponder handles the response to the Cancel request. The method always
// closes the http.Response Body.
func (client TasksClient) CancelResponder(resp *http.Response) (result ProjectTask, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Command the tasks resource is a nested, proxy-only resource representing work performed by a DMS instance. This
// method executes a command on a running task.
    // Parameters:
        // groupName - name of the resource group
        // serviceName - name of the service
        // projectName - name of the project
        // taskName - name of the Task
        // parameters - command to execute
func (client TasksClient) Command(ctx context.Context, groupName string, serviceName string, projectName string, taskName string, parameters BasicCommandProperties) (result CommandPropertiesModel, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/TasksClient.Command")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.CommandPreparer(ctx, groupName, serviceName, projectName, taskName, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "datamigration.TasksClient", "Command", nil , "Failure preparing request")
    return
    }

            resp, err := client.CommandSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "datamigration.TasksClient", "Command", resp, "Failure sending request")
            return
            }

            result, err = client.CommandResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "datamigration.TasksClient", "Command", resp, "Failure responding to request")
            }

    return
    }

    // CommandPreparer prepares the Command request.
    func (client TasksClient) CommandPreparer(ctx context.Context, groupName string, serviceName string, projectName string, taskName string, parameters BasicCommandProperties) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "groupName": autorest.Encode("path",groupName),
            "projectName": autorest.Encode("path",projectName),
            "serviceName": autorest.Encode("path",serviceName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "taskName": autorest.Encode("path",taskName),
            }

                        const APIVersion = "2018-07-15-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}/tasks/{taskName}/command",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CommandSender sends the Command request. The method will close the
    // http.Response Body if it receives an error.
    func (client TasksClient) CommandSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// CommandResponder handles the response to the Command request. The method always
// closes the http.Response Body.
func (client TasksClient) CommandResponder(resp *http.Response) (result CommandPropertiesModel, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// CreateOrUpdate the tasks resource is a nested, proxy-only resource representing work performed by a DMS instance.
// The PUT method creates a new task or updates an existing one, although since tasks have no mutable custom
// properties, there is little reason to update an existing one.
    // Parameters:
        // parameters - information about the task
        // groupName - name of the resource group
        // serviceName - name of the service
        // projectName - name of the project
        // taskName - name of the Task
func (client TasksClient) CreateOrUpdate(ctx context.Context, parameters ProjectTask, groupName string, serviceName string, projectName string, taskName string) (result ProjectTask, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/TasksClient.CreateOrUpdate")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.CreateOrUpdatePreparer(ctx, parameters, groupName, serviceName, projectName, taskName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "datamigration.TasksClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

            resp, err := client.CreateOrUpdateSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "datamigration.TasksClient", "CreateOrUpdate", resp, "Failure sending request")
            return
            }

            result, err = client.CreateOrUpdateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "datamigration.TasksClient", "CreateOrUpdate", resp, "Failure responding to request")
            }

    return
    }

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client TasksClient) CreateOrUpdatePreparer(ctx context.Context, parameters ProjectTask, groupName string, serviceName string, projectName string, taskName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "groupName": autorest.Encode("path",groupName),
            "projectName": autorest.Encode("path",projectName),
            "serviceName": autorest.Encode("path",serviceName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "taskName": autorest.Encode("path",taskName),
            }

                        const APIVersion = "2018-07-15-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPut(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}/tasks/{taskName}",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client TasksClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client TasksClient) CreateOrUpdateResponder(resp *http.Response) (result ProjectTask, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Delete the tasks resource is a nested, proxy-only resource representing work performed by a DMS instance. The DELETE
// method deletes a task, canceling it first if it's running.
    // Parameters:
        // groupName - name of the resource group
        // serviceName - name of the service
        // projectName - name of the project
        // taskName - name of the Task
        // deleteRunningTasks - delete the resource even if it contains running tasks
func (client TasksClient) Delete(ctx context.Context, groupName string, serviceName string, projectName string, taskName string, deleteRunningTasks *bool) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/TasksClient.Delete")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.DeletePreparer(ctx, groupName, serviceName, projectName, taskName, deleteRunningTasks)
    if err != nil {
    err = autorest.NewErrorWithError(err, "datamigration.TasksClient", "Delete", nil , "Failure preparing request")
    return
    }

            resp, err := client.DeleteSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "datamigration.TasksClient", "Delete", resp, "Failure sending request")
            return
            }

            result, err = client.DeleteResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "datamigration.TasksClient", "Delete", resp, "Failure responding to request")
            }

    return
    }

    // DeletePreparer prepares the Delete request.
    func (client TasksClient) DeletePreparer(ctx context.Context, groupName string, serviceName string, projectName string, taskName string, deleteRunningTasks *bool) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "groupName": autorest.Encode("path",groupName),
            "projectName": autorest.Encode("path",projectName),
            "serviceName": autorest.Encode("path",serviceName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "taskName": autorest.Encode("path",taskName),
            }

                        const APIVersion = "2018-07-15-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }
            if deleteRunningTasks != nil {
            queryParameters["deleteRunningTasks"] = autorest.Encode("query",*deleteRunningTasks)
            }

        preparer := autorest.CreatePreparer(
    autorest.AsDelete(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}/tasks/{taskName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client TasksClient) DeleteSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client TasksClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Get the tasks resource is a nested, proxy-only resource representing work performed by a DMS instance. The GET
// method retrieves information about a task.
    // Parameters:
        // groupName - name of the resource group
        // serviceName - name of the service
        // projectName - name of the project
        // taskName - name of the Task
        // expand - expand the response
func (client TasksClient) Get(ctx context.Context, groupName string, serviceName string, projectName string, taskName string, expand string) (result ProjectTask, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/TasksClient.Get")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.GetPreparer(ctx, groupName, serviceName, projectName, taskName, expand)
    if err != nil {
    err = autorest.NewErrorWithError(err, "datamigration.TasksClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "datamigration.TasksClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "datamigration.TasksClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client TasksClient) GetPreparer(ctx context.Context, groupName string, serviceName string, projectName string, taskName string, expand string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "groupName": autorest.Encode("path",groupName),
            "projectName": autorest.Encode("path",projectName),
            "serviceName": autorest.Encode("path",serviceName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "taskName": autorest.Encode("path",taskName),
            }

                        const APIVersion = "2018-07-15-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }
            if len(expand) > 0 {
            queryParameters["$expand"] = autorest.Encode("query",expand)
            }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}/tasks/{taskName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client TasksClient) GetSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client TasksClient) GetResponder(resp *http.Response) (result ProjectTask, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// List the services resource is the top-level resource that represents the Database Migration Service. This method
// returns a list of tasks owned by a service resource. Some tasks may have a status of Unknown, which indicates that
// an error occurred while querying the status of that task.
    // Parameters:
        // groupName - name of the resource group
        // serviceName - name of the service
        // projectName - name of the project
        // taskType - filter tasks by task type
func (client TasksClient) List(ctx context.Context, groupName string, serviceName string, projectName string, taskType string) (result TaskListPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/TasksClient.List")
        defer func() {
            sc := -1
            if result.tl.Response.Response != nil {
                sc = result.tl.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
                result.fn = client.listNextResults
    req, err := client.ListPreparer(ctx, groupName, serviceName, projectName, taskType)
    if err != nil {
    err = autorest.NewErrorWithError(err, "datamigration.TasksClient", "List", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListSender(req)
            if err != nil {
            result.tl.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "datamigration.TasksClient", "List", resp, "Failure sending request")
            return
            }

            result.tl, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "datamigration.TasksClient", "List", resp, "Failure responding to request")
            }

    return
    }

    // ListPreparer prepares the List request.
    func (client TasksClient) ListPreparer(ctx context.Context, groupName string, serviceName string, projectName string, taskType string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "groupName": autorest.Encode("path",groupName),
            "projectName": autorest.Encode("path",projectName),
            "serviceName": autorest.Encode("path",serviceName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-07-15-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }
            if len(taskType) > 0 {
            queryParameters["taskType"] = autorest.Encode("query",taskType)
            }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}/tasks",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client TasksClient) ListSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client TasksClient) ListResponder(resp *http.Response) (result TaskList, err error) {
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
            func (client TasksClient) listNextResults(ctx context.Context, lastResults TaskList) (result TaskList, err error) {
            req, err := lastResults.taskListPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "datamigration.TasksClient", "listNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "datamigration.TasksClient", "listNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "datamigration.TasksClient", "listNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListComplete enumerates all values, automatically crossing page boundaries as required.
    func (client TasksClient) ListComplete(ctx context.Context, groupName string, serviceName string, projectName string, taskType string) (result TaskListIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/TasksClient.List")
            defer func() {
                sc := -1
                if result.Response().Response.Response != nil {
                    sc = result.page.Response().Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
     }
        result.page, err = client.List(ctx, groupName, serviceName, projectName, taskType)
                return
        }

// Update the tasks resource is a nested, proxy-only resource representing work performed by a DMS instance. The PATCH
// method updates an existing task, but since tasks have no mutable custom properties, there is little reason to do so.
    // Parameters:
        // parameters - information about the task
        // groupName - name of the resource group
        // serviceName - name of the service
        // projectName - name of the project
        // taskName - name of the Task
func (client TasksClient) Update(ctx context.Context, parameters ProjectTask, groupName string, serviceName string, projectName string, taskName string) (result ProjectTask, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/TasksClient.Update")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.UpdatePreparer(ctx, parameters, groupName, serviceName, projectName, taskName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "datamigration.TasksClient", "Update", nil , "Failure preparing request")
    return
    }

            resp, err := client.UpdateSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "datamigration.TasksClient", "Update", resp, "Failure sending request")
            return
            }

            result, err = client.UpdateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "datamigration.TasksClient", "Update", resp, "Failure responding to request")
            }

    return
    }

    // UpdatePreparer prepares the Update request.
    func (client TasksClient) UpdatePreparer(ctx context.Context, parameters ProjectTask, groupName string, serviceName string, projectName string, taskName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "groupName": autorest.Encode("path",groupName),
            "projectName": autorest.Encode("path",projectName),
            "serviceName": autorest.Encode("path",serviceName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "taskName": autorest.Encode("path",taskName),
            }

                        const APIVersion = "2018-07-15-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPatch(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}/tasks/{taskName}",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // UpdateSender sends the Update request. The method will close the
    // http.Response Body if it receives an error.
    func (client TasksClient) UpdateSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client TasksClient) UpdateResponder(resp *http.Response) (result ProjectTask, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

