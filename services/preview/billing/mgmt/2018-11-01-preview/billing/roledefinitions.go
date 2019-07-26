package billing

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

// RoleDefinitionsClient is the billing client provides access to billing resources for Azure subscriptions.
type RoleDefinitionsClient struct {
    BaseClient
}
// NewRoleDefinitionsClient creates an instance of the RoleDefinitionsClient client.
func NewRoleDefinitionsClient(subscriptionID string) RoleDefinitionsClient {
    return NewRoleDefinitionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRoleDefinitionsClientWithBaseURI creates an instance of the RoleDefinitionsClient client.
    func NewRoleDefinitionsClientWithBaseURI(baseURI string, subscriptionID string) RoleDefinitionsClient {
        return RoleDefinitionsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// GetByBillingAccountName gets the role definition for a role
    // Parameters:
        // billingAccountName - billing Account Id.
        // billingRoleDefinitionName - role definition id.
func (client RoleDefinitionsClient) GetByBillingAccountName(ctx context.Context, billingAccountName string, billingRoleDefinitionName string) (result RoleDefinition, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/RoleDefinitionsClient.GetByBillingAccountName")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.GetByBillingAccountNamePreparer(ctx, billingAccountName, billingRoleDefinitionName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "billing.RoleDefinitionsClient", "GetByBillingAccountName", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetByBillingAccountNameSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "billing.RoleDefinitionsClient", "GetByBillingAccountName", resp, "Failure sending request")
            return
            }

            result, err = client.GetByBillingAccountNameResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "billing.RoleDefinitionsClient", "GetByBillingAccountName", resp, "Failure responding to request")
            }

    return
    }

    // GetByBillingAccountNamePreparer prepares the GetByBillingAccountName request.
    func (client RoleDefinitionsClient) GetByBillingAccountNamePreparer(ctx context.Context, billingAccountName string, billingRoleDefinitionName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "billingAccountName": autorest.Encode("path",billingAccountName),
            "billingRoleDefinitionName": autorest.Encode("path",billingRoleDefinitionName),
            }

                        const APIVersion = "2018-11-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/providers/Microsoft.Billing/billingRoleDefinitions/{billingRoleDefinitionName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetByBillingAccountNameSender sends the GetByBillingAccountName request. The method will close the
    // http.Response Body if it receives an error.
    func (client RoleDefinitionsClient) GetByBillingAccountNameSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetByBillingAccountNameResponder handles the response to the GetByBillingAccountName request. The method always
// closes the http.Response Body.
func (client RoleDefinitionsClient) GetByBillingAccountNameResponder(resp *http.Response) (result RoleDefinition, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// GetByBillingProfileName gets the role definition for a role
    // Parameters:
        // billingAccountName - billing Account Id.
        // billingProfileName - billing Profile Id.
        // billingRoleDefinitionName - role definition id.
func (client RoleDefinitionsClient) GetByBillingProfileName(ctx context.Context, billingAccountName string, billingProfileName string, billingRoleDefinitionName string) (result RoleDefinition, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/RoleDefinitionsClient.GetByBillingProfileName")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.GetByBillingProfileNamePreparer(ctx, billingAccountName, billingProfileName, billingRoleDefinitionName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "billing.RoleDefinitionsClient", "GetByBillingProfileName", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetByBillingProfileNameSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "billing.RoleDefinitionsClient", "GetByBillingProfileName", resp, "Failure sending request")
            return
            }

            result, err = client.GetByBillingProfileNameResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "billing.RoleDefinitionsClient", "GetByBillingProfileName", resp, "Failure responding to request")
            }

    return
    }

    // GetByBillingProfileNamePreparer prepares the GetByBillingProfileName request.
    func (client RoleDefinitionsClient) GetByBillingProfileNamePreparer(ctx context.Context, billingAccountName string, billingProfileName string, billingRoleDefinitionName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "billingAccountName": autorest.Encode("path",billingAccountName),
            "billingProfileName": autorest.Encode("path",billingProfileName),
            "billingRoleDefinitionName": autorest.Encode("path",billingRoleDefinitionName),
            }

                        const APIVersion = "2018-11-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/providers/Microsoft.Billing/billingRoleDefinitions/{billingRoleDefinitionName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetByBillingProfileNameSender sends the GetByBillingProfileName request. The method will close the
    // http.Response Body if it receives an error.
    func (client RoleDefinitionsClient) GetByBillingProfileNameSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetByBillingProfileNameResponder handles the response to the GetByBillingProfileName request. The method always
// closes the http.Response Body.
func (client RoleDefinitionsClient) GetByBillingProfileNameResponder(resp *http.Response) (result RoleDefinition, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// GetByInvoiceSectionName gets the role definition for a role
    // Parameters:
        // billingAccountName - billing Account Id.
        // invoiceSectionName - invoiceSection Id.
        // billingRoleDefinitionName - role definition id.
func (client RoleDefinitionsClient) GetByInvoiceSectionName(ctx context.Context, billingAccountName string, invoiceSectionName string, billingRoleDefinitionName string) (result RoleDefinition, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/RoleDefinitionsClient.GetByInvoiceSectionName")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.GetByInvoiceSectionNamePreparer(ctx, billingAccountName, invoiceSectionName, billingRoleDefinitionName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "billing.RoleDefinitionsClient", "GetByInvoiceSectionName", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetByInvoiceSectionNameSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "billing.RoleDefinitionsClient", "GetByInvoiceSectionName", resp, "Failure sending request")
            return
            }

            result, err = client.GetByInvoiceSectionNameResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "billing.RoleDefinitionsClient", "GetByInvoiceSectionName", resp, "Failure responding to request")
            }

    return
    }

    // GetByInvoiceSectionNamePreparer prepares the GetByInvoiceSectionName request.
    func (client RoleDefinitionsClient) GetByInvoiceSectionNamePreparer(ctx context.Context, billingAccountName string, invoiceSectionName string, billingRoleDefinitionName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "billingAccountName": autorest.Encode("path",billingAccountName),
            "billingRoleDefinitionName": autorest.Encode("path",billingRoleDefinitionName),
            "invoiceSectionName": autorest.Encode("path",invoiceSectionName),
            }

                        const APIVersion = "2018-11-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/invoiceSections/{invoiceSectionName}/providers/Microsoft.Billing/billingRoleDefinitions/{billingRoleDefinitionName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetByInvoiceSectionNameSender sends the GetByInvoiceSectionName request. The method will close the
    // http.Response Body if it receives an error.
    func (client RoleDefinitionsClient) GetByInvoiceSectionNameSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetByInvoiceSectionNameResponder handles the response to the GetByInvoiceSectionName request. The method always
// closes the http.Response Body.
func (client RoleDefinitionsClient) GetByInvoiceSectionNameResponder(resp *http.Response) (result RoleDefinition, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// ListByBillingAccountName lists the role definition for a billing account
    // Parameters:
        // billingAccountName - billing Account Id.
func (client RoleDefinitionsClient) ListByBillingAccountName(ctx context.Context, billingAccountName string) (result RoleDefinitionListResult, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/RoleDefinitionsClient.ListByBillingAccountName")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.ListByBillingAccountNamePreparer(ctx, billingAccountName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "billing.RoleDefinitionsClient", "ListByBillingAccountName", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListByBillingAccountNameSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "billing.RoleDefinitionsClient", "ListByBillingAccountName", resp, "Failure sending request")
            return
            }

            result, err = client.ListByBillingAccountNameResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "billing.RoleDefinitionsClient", "ListByBillingAccountName", resp, "Failure responding to request")
            }

    return
    }

    // ListByBillingAccountNamePreparer prepares the ListByBillingAccountName request.
    func (client RoleDefinitionsClient) ListByBillingAccountNamePreparer(ctx context.Context, billingAccountName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "billingAccountName": autorest.Encode("path",billingAccountName),
            }

                        const APIVersion = "2018-11-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/providers/Microsoft.Billing/billingRoleDefinitions",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByBillingAccountNameSender sends the ListByBillingAccountName request. The method will close the
    // http.Response Body if it receives an error.
    func (client RoleDefinitionsClient) ListByBillingAccountNameSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListByBillingAccountNameResponder handles the response to the ListByBillingAccountName request. The method always
// closes the http.Response Body.
func (client RoleDefinitionsClient) ListByBillingAccountNameResponder(resp *http.Response) (result RoleDefinitionListResult, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// ListByBillingProfileName lists the role definition for a Billing Profile
    // Parameters:
        // billingAccountName - billing Account Id.
        // billingProfileName - billing Profile Id.
func (client RoleDefinitionsClient) ListByBillingProfileName(ctx context.Context, billingAccountName string, billingProfileName string) (result RoleDefinitionListResult, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/RoleDefinitionsClient.ListByBillingProfileName")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.ListByBillingProfileNamePreparer(ctx, billingAccountName, billingProfileName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "billing.RoleDefinitionsClient", "ListByBillingProfileName", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListByBillingProfileNameSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "billing.RoleDefinitionsClient", "ListByBillingProfileName", resp, "Failure sending request")
            return
            }

            result, err = client.ListByBillingProfileNameResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "billing.RoleDefinitionsClient", "ListByBillingProfileName", resp, "Failure responding to request")
            }

    return
    }

    // ListByBillingProfileNamePreparer prepares the ListByBillingProfileName request.
    func (client RoleDefinitionsClient) ListByBillingProfileNamePreparer(ctx context.Context, billingAccountName string, billingProfileName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "billingAccountName": autorest.Encode("path",billingAccountName),
            "billingProfileName": autorest.Encode("path",billingProfileName),
            }

                        const APIVersion = "2018-11-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/providers/Microsoft.Billing/billingRoleDefinitions",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByBillingProfileNameSender sends the ListByBillingProfileName request. The method will close the
    // http.Response Body if it receives an error.
    func (client RoleDefinitionsClient) ListByBillingProfileNameSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListByBillingProfileNameResponder handles the response to the ListByBillingProfileName request. The method always
// closes the http.Response Body.
func (client RoleDefinitionsClient) ListByBillingProfileNameResponder(resp *http.Response) (result RoleDefinitionListResult, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// ListByInvoiceSectionName lists the role definition for an invoice Section
    // Parameters:
        // billingAccountName - billing Account Id.
        // invoiceSectionName - invoiceSection Id.
func (client RoleDefinitionsClient) ListByInvoiceSectionName(ctx context.Context, billingAccountName string, invoiceSectionName string) (result RoleDefinitionListResult, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/RoleDefinitionsClient.ListByInvoiceSectionName")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.ListByInvoiceSectionNamePreparer(ctx, billingAccountName, invoiceSectionName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "billing.RoleDefinitionsClient", "ListByInvoiceSectionName", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListByInvoiceSectionNameSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "billing.RoleDefinitionsClient", "ListByInvoiceSectionName", resp, "Failure sending request")
            return
            }

            result, err = client.ListByInvoiceSectionNameResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "billing.RoleDefinitionsClient", "ListByInvoiceSectionName", resp, "Failure responding to request")
            }

    return
    }

    // ListByInvoiceSectionNamePreparer prepares the ListByInvoiceSectionName request.
    func (client RoleDefinitionsClient) ListByInvoiceSectionNamePreparer(ctx context.Context, billingAccountName string, invoiceSectionName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "billingAccountName": autorest.Encode("path",billingAccountName),
            "invoiceSectionName": autorest.Encode("path",invoiceSectionName),
            }

                        const APIVersion = "2018-11-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/invoiceSections/{invoiceSectionName}/providers/Microsoft.Billing/billingRoleDefinitions",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByInvoiceSectionNameSender sends the ListByInvoiceSectionName request. The method will close the
    // http.Response Body if it receives an error.
    func (client RoleDefinitionsClient) ListByInvoiceSectionNameSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListByInvoiceSectionNameResponder handles the response to the ListByInvoiceSectionName request. The method always
// closes the http.Response Body.
func (client RoleDefinitionsClient) ListByInvoiceSectionNameResponder(resp *http.Response) (result RoleDefinitionListResult, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

