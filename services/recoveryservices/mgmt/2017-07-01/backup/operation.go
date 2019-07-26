package backup

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

// OperationClient is the open API 2.0 Specs for Azure RecoveryServices Backup service
type OperationClient struct {
    BaseClient
}
// NewOperationClient creates an instance of the OperationClient client.
func NewOperationClient(subscriptionID string) OperationClient {
    return NewOperationClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewOperationClientWithBaseURI creates an instance of the OperationClient client.
    func NewOperationClientWithBaseURI(baseURI string, subscriptionID string) OperationClient {
        return OperationClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// Validate validate operation for specified backed up item. This is a synchronous operation.
    // Parameters:
        // vaultName - the name of the recovery services vault.
        // resourceGroupName - the name of the resource group where the recovery services vault is present.
        // parameters - resource validate operation request
func (client OperationClient) Validate(ctx context.Context, vaultName string, resourceGroupName string, parameters BasicValidateOperationRequest) (result ValidateOperationsResponse, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/OperationClient.Validate")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.ValidatePreparer(ctx, vaultName, resourceGroupName, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "backup.OperationClient", "Validate", nil , "Failure preparing request")
    return
    }

            resp, err := client.ValidateSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "backup.OperationClient", "Validate", resp, "Failure sending request")
            return
            }

            result, err = client.ValidateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "backup.OperationClient", "Validate", resp, "Failure responding to request")
            }

    return
    }

    // ValidatePreparer prepares the Validate request.
    func (client OperationClient) ValidatePreparer(ctx context.Context, vaultName string, resourceGroupName string, parameters BasicValidateOperationRequest) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "vaultName": autorest.Encode("path",vaultName),
            }

                        const APIVersion = "2017-07-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupValidateOperation",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ValidateSender sends the Validate request. The method will close the
    // http.Response Body if it receives an error.
    func (client OperationClient) ValidateSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ValidateResponder handles the response to the Validate request. The method always
// closes the http.Response Body.
func (client OperationClient) ValidateResponder(resp *http.Response) (result ValidateOperationsResponse, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

