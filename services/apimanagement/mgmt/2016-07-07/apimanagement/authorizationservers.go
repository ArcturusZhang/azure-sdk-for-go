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

// AuthorizationServersClient is the apiManagement Client
type AuthorizationServersClient struct {
    BaseClient
}
// NewAuthorizationServersClient creates an instance of the AuthorizationServersClient client.
func NewAuthorizationServersClient(subscriptionID string) AuthorizationServersClient {
    return NewAuthorizationServersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAuthorizationServersClientWithBaseURI creates an instance of the AuthorizationServersClient client.
    func NewAuthorizationServersClientWithBaseURI(baseURI string, subscriptionID string) AuthorizationServersClient {
        return AuthorizationServersClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CreateOrUpdate creates new authorization server or updates an existing authorization server.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // serviceName - the name of the API Management service.
        // authsid - identifier of the authorization server.
        // parameters - create or update parameters.
func (client AuthorizationServersClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, authsid string, parameters OAuth2AuthorizationServerContract) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/AuthorizationServersClient.CreateOrUpdate")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: serviceName,
             Constraints: []validation.Constraint{	{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil },
            	{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil }}},
            { TargetValue: authsid,
             Constraints: []validation.Constraint{	{Target: "authsid", Name: validation.MaxLength, Rule: 256, Chain: nil },
            	{Target: "authsid", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil }}},
            { TargetValue: parameters,
             Constraints: []validation.Constraint{	{Target: "parameters.Name", Name: validation.Null, Rule: true ,
            Chain: []validation.Constraint{	{Target: "parameters.Name", Name: validation.MaxLength, Rule: 50, Chain: nil },
            	{Target: "parameters.Name", Name: validation.MinLength, Rule: 1, Chain: nil },
            }},
            	{Target: "parameters.ClientRegistrationEndpoint", Name: validation.Null, Rule: true, Chain: nil },
            	{Target: "parameters.AuthorizationEndpoint", Name: validation.Null, Rule: true, Chain: nil },
            	{Target: "parameters.GrantTypes", Name: validation.Null, Rule: true, Chain: nil },
            	{Target: "parameters.ClientID", Name: validation.Null, Rule: true, Chain: nil }}}}); err != nil {
            return result, validation.NewError("apimanagement.AuthorizationServersClient", "CreateOrUpdate", err.Error())
            }

                req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serviceName, authsid, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "apimanagement.AuthorizationServersClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

            resp, err := client.CreateOrUpdateSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "apimanagement.AuthorizationServersClient", "CreateOrUpdate", resp, "Failure sending request")
            return
            }

            result, err = client.CreateOrUpdateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "apimanagement.AuthorizationServersClient", "CreateOrUpdate", resp, "Failure responding to request")
            }

    return
    }

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client AuthorizationServersClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serviceName string, authsid string, parameters OAuth2AuthorizationServerContract) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "authsid": autorest.Encode("path",authsid),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "serviceName": autorest.Encode("path",serviceName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2016-07-07"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

                    parameters.ID = nil
    preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPut(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/authorizationServers/{authsid}",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client AuthorizationServersClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client AuthorizationServersClient) CreateOrUpdateResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Delete deletes specific authorization server instance.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // serviceName - the name of the API Management service.
        // authsid - identifier of the authorization server.
        // ifMatch - the entity state (Etag) version of the authentication server to delete. A value of "*" can be used
        // for If-Match to unconditionally apply the operation.
func (client AuthorizationServersClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, authsid string, ifMatch string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/AuthorizationServersClient.Delete")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: serviceName,
             Constraints: []validation.Constraint{	{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil },
            	{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("apimanagement.AuthorizationServersClient", "Delete", err.Error())
            }

                req, err := client.DeletePreparer(ctx, resourceGroupName, serviceName, authsid, ifMatch)
    if err != nil {
    err = autorest.NewErrorWithError(err, "apimanagement.AuthorizationServersClient", "Delete", nil , "Failure preparing request")
    return
    }

            resp, err := client.DeleteSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "apimanagement.AuthorizationServersClient", "Delete", resp, "Failure sending request")
            return
            }

            result, err = client.DeleteResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "apimanagement.AuthorizationServersClient", "Delete", resp, "Failure responding to request")
            }

    return
    }

    // DeletePreparer prepares the Delete request.
    func (client AuthorizationServersClient) DeletePreparer(ctx context.Context, resourceGroupName string, serviceName string, authsid string, ifMatch string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "authsid": autorest.Encode("path",authsid),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "serviceName": autorest.Encode("path",serviceName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2016-07-07"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsDelete(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/authorizationServers/{authsid}",pathParameters),
    autorest.WithQueryParameters(queryParameters),
    autorest.WithHeader("If-Match", autorest.String(ifMatch)))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client AuthorizationServersClient) DeleteSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client AuthorizationServersClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Get gets the details of the authorization server specified by its identifier.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // serviceName - the name of the API Management service.
        // authsid - identifier of the authorization server.
func (client AuthorizationServersClient) Get(ctx context.Context, resourceGroupName string, serviceName string, authsid string) (result OAuth2AuthorizationServerContract, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/AuthorizationServersClient.Get")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: serviceName,
             Constraints: []validation.Constraint{	{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil },
            	{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil }}},
            { TargetValue: authsid,
             Constraints: []validation.Constraint{	{Target: "authsid", Name: validation.MaxLength, Rule: 256, Chain: nil },
            	{Target: "authsid", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("apimanagement.AuthorizationServersClient", "Get", err.Error())
            }

                req, err := client.GetPreparer(ctx, resourceGroupName, serviceName, authsid)
    if err != nil {
    err = autorest.NewErrorWithError(err, "apimanagement.AuthorizationServersClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "apimanagement.AuthorizationServersClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "apimanagement.AuthorizationServersClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client AuthorizationServersClient) GetPreparer(ctx context.Context, resourceGroupName string, serviceName string, authsid string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "authsid": autorest.Encode("path",authsid),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "serviceName": autorest.Encode("path",serviceName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2016-07-07"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/authorizationServers/{authsid}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client AuthorizationServersClient) GetSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AuthorizationServersClient) GetResponder(resp *http.Response) (result OAuth2AuthorizationServerContract, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// ListByService lists a collection of authorization servers defined within a service instance.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // serviceName - the name of the API Management service.
        // filter - | Field | Supported operators    | Supported functions                         |
        // |-------|------------------------|---------------------------------------------|
        // | id    | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
        // | name  | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
        // top - number of records to return.
        // skip - number of records to skip.
func (client AuthorizationServersClient) ListByService(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result AuthorizationServerCollectionPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/AuthorizationServersClient.ListByService")
        defer func() {
            sc := -1
            if result.asc.Response.Response != nil {
                sc = result.asc.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: serviceName,
             Constraints: []validation.Constraint{	{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil },
            	{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil }}},
            { TargetValue: top,
             Constraints: []validation.Constraint{	{Target: "top", Name: validation.Null, Rule: false ,
            Chain: []validation.Constraint{	{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil },
            }}}},
            { TargetValue: skip,
             Constraints: []validation.Constraint{	{Target: "skip", Name: validation.Null, Rule: false ,
            Chain: []validation.Constraint{	{Target: "skip", Name: validation.InclusiveMinimum, Rule: 0, Chain: nil },
            }}}}}); err != nil {
            return result, validation.NewError("apimanagement.AuthorizationServersClient", "ListByService", err.Error())
            }

                        result.fn = client.listByServiceNextResults
    req, err := client.ListByServicePreparer(ctx, resourceGroupName, serviceName, filter, top, skip)
    if err != nil {
    err = autorest.NewErrorWithError(err, "apimanagement.AuthorizationServersClient", "ListByService", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListByServiceSender(req)
            if err != nil {
            result.asc.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "apimanagement.AuthorizationServersClient", "ListByService", resp, "Failure sending request")
            return
            }

            result.asc, err = client.ListByServiceResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "apimanagement.AuthorizationServersClient", "ListByService", resp, "Failure responding to request")
            }

    return
    }

    // ListByServicePreparer prepares the ListByService request.
    func (client AuthorizationServersClient) ListByServicePreparer(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "serviceName": autorest.Encode("path",serviceName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2016-07-07"
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
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/authorizationServers",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByServiceSender sends the ListByService request. The method will close the
    // http.Response Body if it receives an error.
    func (client AuthorizationServersClient) ListByServiceSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListByServiceResponder handles the response to the ListByService request. The method always
// closes the http.Response Body.
func (client AuthorizationServersClient) ListByServiceResponder(resp *http.Response) (result AuthorizationServerCollection, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

            // listByServiceNextResults retrieves the next set of results, if any.
            func (client AuthorizationServersClient) listByServiceNextResults(ctx context.Context, lastResults AuthorizationServerCollection) (result AuthorizationServerCollection, err error) {
            req, err := lastResults.authorizationServerCollectionPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "apimanagement.AuthorizationServersClient", "listByServiceNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListByServiceSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "apimanagement.AuthorizationServersClient", "listByServiceNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListByServiceResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "apimanagement.AuthorizationServersClient", "listByServiceNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListByServiceComplete enumerates all values, automatically crossing page boundaries as required.
    func (client AuthorizationServersClient) ListByServiceComplete(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result AuthorizationServerCollectionIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/AuthorizationServersClient.ListByService")
            defer func() {
                sc := -1
                if result.Response().Response.Response != nil {
                    sc = result.page.Response().Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
     }
        result.page, err = client.ListByService(ctx, resourceGroupName, serviceName, filter, top, skip)
                return
        }

// Update updates the details of the authorization server specified by its identifier.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // serviceName - the name of the API Management service.
        // authsid - identifier of the authorization server.
        // parameters - oAuth2 Server settings Update parameters.
        // ifMatch - the entity state (Etag) version of the authorization server to update. A value of "*" can be used
        // for If-Match to unconditionally apply the operation.
func (client AuthorizationServersClient) Update(ctx context.Context, resourceGroupName string, serviceName string, authsid string, parameters OAuth2AuthorizationServerUpdateContract, ifMatch string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/AuthorizationServersClient.Update")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: serviceName,
             Constraints: []validation.Constraint{	{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil },
            	{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("apimanagement.AuthorizationServersClient", "Update", err.Error())
            }

                req, err := client.UpdatePreparer(ctx, resourceGroupName, serviceName, authsid, parameters, ifMatch)
    if err != nil {
    err = autorest.NewErrorWithError(err, "apimanagement.AuthorizationServersClient", "Update", nil , "Failure preparing request")
    return
    }

            resp, err := client.UpdateSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "apimanagement.AuthorizationServersClient", "Update", resp, "Failure sending request")
            return
            }

            result, err = client.UpdateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "apimanagement.AuthorizationServersClient", "Update", resp, "Failure responding to request")
            }

    return
    }

    // UpdatePreparer prepares the Update request.
    func (client AuthorizationServersClient) UpdatePreparer(ctx context.Context, resourceGroupName string, serviceName string, authsid string, parameters OAuth2AuthorizationServerUpdateContract, ifMatch string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "authsid": autorest.Encode("path",authsid),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "serviceName": autorest.Encode("path",serviceName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2016-07-07"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPatch(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/authorizationServers/{authsid}",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters),
    autorest.WithHeader("If-Match", autorest.String(ifMatch)))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // UpdateSender sends the Update request. The method will close the
    // http.Response Body if it receives an error.
    func (client AuthorizationServersClient) UpdateSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client AuthorizationServersClient) UpdateResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

