package apimanagement

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

// UserGroupClient is the apiManagement Client
type UserGroupClient struct {
    BaseClient
}
// NewUserGroupClient creates an instance of the UserGroupClient client.
func NewUserGroupClient(subscriptionID string) UserGroupClient {
    return NewUserGroupClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewUserGroupClientWithBaseURI creates an instance of the UserGroupClient client.
    func NewUserGroupClientWithBaseURI(baseURI string, subscriptionID string) UserGroupClient {
        return UserGroupClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// List lists all user groups.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // serviceName - the name of the API Management service.
        // userID - user identifier. Must be unique in the current API Management service instance.
        // filter - | Field       | Supported operators    | Supported functions               |
        // |-------------|------------------------|-----------------------------------|
        //
        // |name | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith|
        // |displayName | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith|
        // |description | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith|
        // top - number of records to return.
        // skip - number of records to skip.
func (client UserGroupClient) List(ctx context.Context, resourceGroupName string, serviceName string, userID string, filter string, top *int32, skip *int32) (result GroupCollectionPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/UserGroupClient.List")
        defer func() {
            sc := -1
            if result.gc.Response.Response != nil {
                sc = result.gc.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: serviceName,
             Constraints: []validation.Constraint{	{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil },
            	{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil }}},
            { TargetValue: userID,
             Constraints: []validation.Constraint{	{Target: "userID", Name: validation.MaxLength, Rule: 80, Chain: nil },
            	{Target: "userID", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "userID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil }}},
            { TargetValue: top,
             Constraints: []validation.Constraint{	{Target: "top", Name: validation.Null, Rule: false ,
            Chain: []validation.Constraint{	{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil },
            }}}},
            { TargetValue: skip,
             Constraints: []validation.Constraint{	{Target: "skip", Name: validation.Null, Rule: false ,
            Chain: []validation.Constraint{	{Target: "skip", Name: validation.InclusiveMinimum, Rule: 0, Chain: nil },
            }}}}}); err != nil {
            return result, validation.NewError("apimanagement.UserGroupClient", "List", err.Error())
            }

                        result.fn = client.listNextResults
    req, err := client.ListPreparer(ctx, resourceGroupName, serviceName, userID, filter, top, skip)
    if err != nil {
    err = autorest.NewErrorWithError(err, "apimanagement.UserGroupClient", "List", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListSender(req)
            if err != nil {
            result.gc.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "apimanagement.UserGroupClient", "List", resp, "Failure sending request")
            return
            }

            result.gc, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "apimanagement.UserGroupClient", "List", resp, "Failure responding to request")
            }

    return
    }

    // ListPreparer prepares the List request.
    func (client UserGroupClient) ListPreparer(ctx context.Context, resourceGroupName string, serviceName string, userID string, filter string, top *int32, skip *int32) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "serviceName": autorest.Encode("path",serviceName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "userId": autorest.Encode("path",userID),
            }

                        const APIVersion = "2018-06-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }
            if len(filter) > 0 {
            queryParameters["$filter"] = autorest.Encode("query",filter)
            }
            if top != nil {
            queryParameters["$top"] = autorest.Encode("query",*top)
            }
            if skip != nil {
            queryParameters["$skip"] = autorest.Encode("query",*skip)
            }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{userId}/groups",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client UserGroupClient) ListSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client UserGroupClient) ListResponder(resp *http.Response) (result GroupCollection, err error) {
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
            func (client UserGroupClient) listNextResults(ctx context.Context, lastResults GroupCollection) (result GroupCollection, err error) {
            req, err := lastResults.groupCollectionPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "apimanagement.UserGroupClient", "listNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "apimanagement.UserGroupClient", "listNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "apimanagement.UserGroupClient", "listNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListComplete enumerates all values, automatically crossing page boundaries as required.
    func (client UserGroupClient) ListComplete(ctx context.Context, resourceGroupName string, serviceName string, userID string, filter string, top *int32, skip *int32) (result GroupCollectionIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/UserGroupClient.List")
            defer func() {
                sc := -1
                if result.Response().Response.Response != nil {
                    sc = result.page.Response().Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
     }
        result.page, err = client.List(ctx, resourceGroupName, serviceName, userID, filter, top, skip)
                return
        }

