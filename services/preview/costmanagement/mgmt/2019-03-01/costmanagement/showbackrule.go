package costmanagement

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

// ShowbackRuleClient is the client for the ShowbackRule methods of the Costmanagement service.
type ShowbackRuleClient struct {
    BaseClient
}
// NewShowbackRuleClient creates an instance of the ShowbackRuleClient client.
func NewShowbackRuleClient(subscriptionID string) ShowbackRuleClient {
    return NewShowbackRuleClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewShowbackRuleClientWithBaseURI creates an instance of the ShowbackRuleClient client.
    func NewShowbackRuleClientWithBaseURI(baseURI string, subscriptionID string) ShowbackRuleClient {
        return ShowbackRuleClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CreateUpdateRule create/Update showback rule for billing account.
    // Parameters:
        // billingAccountID - billingAccount ID
        // ruleName - showback rule name
        // showbackRule - showback rule to create or update.
func (client ShowbackRuleClient) CreateUpdateRule(ctx context.Context, billingAccountID string, ruleName string, showbackRule ShowbackRule) (result ShowbackRule, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ShowbackRuleClient.CreateUpdateRule")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.CreateUpdateRulePreparer(ctx, billingAccountID, ruleName, showbackRule)
    if err != nil {
    err = autorest.NewErrorWithError(err, "costmanagement.ShowbackRuleClient", "CreateUpdateRule", nil , "Failure preparing request")
    return
    }

            resp, err := client.CreateUpdateRuleSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "costmanagement.ShowbackRuleClient", "CreateUpdateRule", resp, "Failure sending request")
            return
            }

            result, err = client.CreateUpdateRuleResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "costmanagement.ShowbackRuleClient", "CreateUpdateRule", resp, "Failure responding to request")
            }

    return
    }

    // CreateUpdateRulePreparer prepares the CreateUpdateRule request.
    func (client ShowbackRuleClient) CreateUpdateRulePreparer(ctx context.Context, billingAccountID string, ruleName string, showbackRule ShowbackRule) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "billingAccountId": autorest.Encode("path",billingAccountID),
            "ruleName": autorest.Encode("path",ruleName),
            }

                        const APIVersion = "2019-03-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

                    showbackRule.ID = nil
                showbackRule.Name = nil
                showbackRule.Type = nil
    preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPut(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/providers/Microsoft.CostManagement/showbackRules/{ruleName}",pathParameters),
    autorest.WithJSON(showbackRule),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateUpdateRuleSender sends the CreateUpdateRule request. The method will close the
    // http.Response Body if it receives an error.
    func (client ShowbackRuleClient) CreateUpdateRuleSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// CreateUpdateRuleResponder handles the response to the CreateUpdateRule request. The method always
// closes the http.Response Body.
func (client ShowbackRuleClient) CreateUpdateRuleResponder(resp *http.Response) (result ShowbackRule, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// GetBillingAccountID gets the showback rule by rule name.
    // Parameters:
        // billingAccountID - billingAccount ID
        // ruleName - showback rule name
func (client ShowbackRuleClient) GetBillingAccountID(ctx context.Context, billingAccountID string, ruleName string) (result ShowbackRule, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ShowbackRuleClient.GetBillingAccountID")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.GetBillingAccountIDPreparer(ctx, billingAccountID, ruleName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "costmanagement.ShowbackRuleClient", "GetBillingAccountID", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetBillingAccountIDSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "costmanagement.ShowbackRuleClient", "GetBillingAccountID", resp, "Failure sending request")
            return
            }

            result, err = client.GetBillingAccountIDResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "costmanagement.ShowbackRuleClient", "GetBillingAccountID", resp, "Failure responding to request")
            }

    return
    }

    // GetBillingAccountIDPreparer prepares the GetBillingAccountID request.
    func (client ShowbackRuleClient) GetBillingAccountIDPreparer(ctx context.Context, billingAccountID string, ruleName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "billingAccountId": autorest.Encode("path",billingAccountID),
            "ruleName": autorest.Encode("path",ruleName),
            }

                        const APIVersion = "2019-03-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/providers/Microsoft.CostManagement/showbackRules/{ruleName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetBillingAccountIDSender sends the GetBillingAccountID request. The method will close the
    // http.Response Body if it receives an error.
    func (client ShowbackRuleClient) GetBillingAccountIDSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetBillingAccountIDResponder handles the response to the GetBillingAccountID request. The method always
// closes the http.Response Body.
func (client ShowbackRuleClient) GetBillingAccountIDResponder(resp *http.Response) (result ShowbackRule, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

