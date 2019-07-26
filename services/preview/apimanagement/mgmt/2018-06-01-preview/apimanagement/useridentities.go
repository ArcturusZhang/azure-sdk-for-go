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

// UserIdentitiesClient is the apiManagement Client
type UserIdentitiesClient struct {
    BaseClient
}
// NewUserIdentitiesClient creates an instance of the UserIdentitiesClient client.
func NewUserIdentitiesClient(subscriptionID string) UserIdentitiesClient {
    return NewUserIdentitiesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewUserIdentitiesClientWithBaseURI creates an instance of the UserIdentitiesClient client.
    func NewUserIdentitiesClientWithBaseURI(baseURI string, subscriptionID string) UserIdentitiesClient {
        return UserIdentitiesClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// List list of all user identities.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // serviceName - the name of the API Management service.
        // userID - user identifier. Must be unique in the current API Management service instance.
func (client UserIdentitiesClient) List(ctx context.Context, resourceGroupName string, serviceName string, userID string) (result UserIdentityCollectionPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/UserIdentitiesClient.List")
        defer func() {
            sc := -1
            if result.uic.Response.Response != nil {
                sc = result.uic.Response.Response.StatusCode
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
            	{Target: "userID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("apimanagement.UserIdentitiesClient", "List", err.Error())
            }

                        result.fn = client.listNextResults
    req, err := client.ListPreparer(ctx, resourceGroupName, serviceName, userID)
    if err != nil {
    err = autorest.NewErrorWithError(err, "apimanagement.UserIdentitiesClient", "List", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListSender(req)
            if err != nil {
            result.uic.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "apimanagement.UserIdentitiesClient", "List", resp, "Failure sending request")
            return
            }

            result.uic, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "apimanagement.UserIdentitiesClient", "List", resp, "Failure responding to request")
            }

    return
    }

    // ListPreparer prepares the List request.
    func (client UserIdentitiesClient) ListPreparer(ctx context.Context, resourceGroupName string, serviceName string, userID string) (*http.Request, error) {
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

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{userId}/identities",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client UserIdentitiesClient) ListSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client UserIdentitiesClient) ListResponder(resp *http.Response) (result UserIdentityCollection, err error) {
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
            func (client UserIdentitiesClient) listNextResults(ctx context.Context, lastResults UserIdentityCollection) (result UserIdentityCollection, err error) {
            req, err := lastResults.userIdentityCollectionPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "apimanagement.UserIdentitiesClient", "listNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "apimanagement.UserIdentitiesClient", "listNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "apimanagement.UserIdentitiesClient", "listNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListComplete enumerates all values, automatically crossing page boundaries as required.
    func (client UserIdentitiesClient) ListComplete(ctx context.Context, resourceGroupName string, serviceName string, userID string) (result UserIdentityCollectionIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/UserIdentitiesClient.List")
            defer func() {
                sc := -1
                if result.Response().Response.Response != nil {
                    sc = result.page.Response().Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
     }
        result.page, err = client.List(ctx, resourceGroupName, serviceName, userID)
                return
        }

