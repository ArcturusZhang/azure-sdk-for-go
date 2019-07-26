package sql

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

// FirewallRulesClient is the the Azure SQL Database management API provides a RESTful set of web services that
// interact with Azure SQL Database services to manage your databases. The API enables you to create, retrieve, update,
// and delete databases.
type FirewallRulesClient struct {
    BaseClient
}
// NewFirewallRulesClient creates an instance of the FirewallRulesClient client.
func NewFirewallRulesClient(subscriptionID string) FirewallRulesClient {
    return NewFirewallRulesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewFirewallRulesClientWithBaseURI creates an instance of the FirewallRulesClient client.
    func NewFirewallRulesClientWithBaseURI(baseURI string, subscriptionID string) FirewallRulesClient {
        return FirewallRulesClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CreateOrUpdate creates or updates a firewall rule.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serverName - the name of the server.
        // firewallRuleName - the name of the firewall rule.
        // parameters - the required parameters for creating or updating a firewall rule.
func (client FirewallRulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string, parameters FirewallRule) (result FirewallRule, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/FirewallRulesClient.CreateOrUpdate")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: parameters,
             Constraints: []validation.Constraint{	{Target: "parameters.FirewallRuleProperties", Name: validation.Null, Rule: false ,
            Chain: []validation.Constraint{	{Target: "parameters.FirewallRuleProperties.StartIPAddress", Name: validation.Null, Rule: true, Chain: nil },
            	{Target: "parameters.FirewallRuleProperties.EndIPAddress", Name: validation.Null, Rule: true, Chain: nil },
            }}}}}); err != nil {
            return result, validation.NewError("sql.FirewallRulesClient", "CreateOrUpdate", err.Error())
            }

                req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serverName, firewallRuleName, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "sql.FirewallRulesClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

            resp, err := client.CreateOrUpdateSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "sql.FirewallRulesClient", "CreateOrUpdate", resp, "Failure sending request")
            return
            }

            result, err = client.CreateOrUpdateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "sql.FirewallRulesClient", "CreateOrUpdate", resp, "Failure responding to request")
            }

    return
    }

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client FirewallRulesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string, parameters FirewallRule) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "firewallRuleName": autorest.Encode("path",firewallRuleName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "serverName": autorest.Encode("path",serverName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2014-04-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

                    parameters.Kind = nil
                parameters.Location = nil
    preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPut(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/firewallRules/{firewallRuleName}",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client FirewallRulesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client FirewallRulesClient) CreateOrUpdateResponder(resp *http.Response) (result FirewallRule, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Delete deletes a firewall rule.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serverName - the name of the server.
        // firewallRuleName - the name of the firewall rule.
func (client FirewallRulesClient) Delete(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/FirewallRulesClient.Delete")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.DeletePreparer(ctx, resourceGroupName, serverName, firewallRuleName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "sql.FirewallRulesClient", "Delete", nil , "Failure preparing request")
    return
    }

            resp, err := client.DeleteSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "sql.FirewallRulesClient", "Delete", resp, "Failure sending request")
            return
            }

            result, err = client.DeleteResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "sql.FirewallRulesClient", "Delete", resp, "Failure responding to request")
            }

    return
    }

    // DeletePreparer prepares the Delete request.
    func (client FirewallRulesClient) DeletePreparer(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "firewallRuleName": autorest.Encode("path",firewallRuleName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "serverName": autorest.Encode("path",serverName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2014-04-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsDelete(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/firewallRules/{firewallRuleName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client FirewallRulesClient) DeleteSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client FirewallRulesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Get gets a firewall rule.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serverName - the name of the server.
        // firewallRuleName - the name of the firewall rule.
func (client FirewallRulesClient) Get(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string) (result FirewallRule, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/FirewallRulesClient.Get")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.GetPreparer(ctx, resourceGroupName, serverName, firewallRuleName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "sql.FirewallRulesClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "sql.FirewallRulesClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "sql.FirewallRulesClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client FirewallRulesClient) GetPreparer(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "firewallRuleName": autorest.Encode("path",firewallRuleName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "serverName": autorest.Encode("path",serverName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2014-04-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/firewallRules/{firewallRuleName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client FirewallRulesClient) GetSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client FirewallRulesClient) GetResponder(resp *http.Response) (result FirewallRule, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// ListByServer returns a list of firewall rules.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serverName - the name of the server.
func (client FirewallRulesClient) ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result FirewallRuleListResult, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/FirewallRulesClient.ListByServer")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.ListByServerPreparer(ctx, resourceGroupName, serverName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "sql.FirewallRulesClient", "ListByServer", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListByServerSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "sql.FirewallRulesClient", "ListByServer", resp, "Failure sending request")
            return
            }

            result, err = client.ListByServerResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "sql.FirewallRulesClient", "ListByServer", resp, "Failure responding to request")
            }

    return
    }

    // ListByServerPreparer prepares the ListByServer request.
    func (client FirewallRulesClient) ListByServerPreparer(ctx context.Context, resourceGroupName string, serverName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "serverName": autorest.Encode("path",serverName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2014-04-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/firewallRules",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByServerSender sends the ListByServer request. The method will close the
    // http.Response Body if it receives an error.
    func (client FirewallRulesClient) ListByServerSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListByServerResponder handles the response to the ListByServer request. The method always
// closes the http.Response Body.
func (client FirewallRulesClient) ListByServerResponder(resp *http.Response) (result FirewallRuleListResult, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

