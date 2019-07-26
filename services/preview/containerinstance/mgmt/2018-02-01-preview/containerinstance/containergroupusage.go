package containerinstance

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

// ContainerGroupUsageClient is the client for the ContainerGroupUsage methods of the Containerinstance service.
type ContainerGroupUsageClient struct {
    BaseClient
}
// NewContainerGroupUsageClient creates an instance of the ContainerGroupUsageClient client.
func NewContainerGroupUsageClient(subscriptionID string) ContainerGroupUsageClient {
    return NewContainerGroupUsageClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewContainerGroupUsageClientWithBaseURI creates an instance of the ContainerGroupUsageClient client.
    func NewContainerGroupUsageClientWithBaseURI(baseURI string, subscriptionID string) ContainerGroupUsageClient {
        return ContainerGroupUsageClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// List get the usage for a subscription
    // Parameters:
        // location - the identifier for the physical azure location.
func (client ContainerGroupUsageClient) List(ctx context.Context, location string) (result UsageListResult, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ContainerGroupUsageClient.List")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.ListPreparer(ctx, location)
    if err != nil {
    err = autorest.NewErrorWithError(err, "containerinstance.ContainerGroupUsageClient", "List", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "containerinstance.ContainerGroupUsageClient", "List", resp, "Failure sending request")
            return
            }

            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "containerinstance.ContainerGroupUsageClient", "List", resp, "Failure responding to request")
            }

    return
    }

    // ListPreparer prepares the List request.
    func (client ContainerGroupUsageClient) ListPreparer(ctx context.Context, location string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "location": autorest.Encode("path",location),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-02-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.ContainerInstance/locations/{location}/usages",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client ContainerGroupUsageClient) ListSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ContainerGroupUsageClient) ListResponder(resp *http.Response) (result UsageListResult, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

