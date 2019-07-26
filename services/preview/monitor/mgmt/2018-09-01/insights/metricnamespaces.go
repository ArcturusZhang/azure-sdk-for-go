package insights

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

// MetricNamespacesClient is the monitor Management Client
type MetricNamespacesClient struct {
    BaseClient
}
// NewMetricNamespacesClient creates an instance of the MetricNamespacesClient client.
func NewMetricNamespacesClient(subscriptionID string) MetricNamespacesClient {
    return NewMetricNamespacesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewMetricNamespacesClientWithBaseURI creates an instance of the MetricNamespacesClient client.
    func NewMetricNamespacesClientWithBaseURI(baseURI string, subscriptionID string) MetricNamespacesClient {
        return MetricNamespacesClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// List lists the metric namespaces for the resource.
    // Parameters:
        // resourceURI - the identifier of the resource.
        // startTime - the ISO 8601 conform Date start time from which to query for metric namespaces.
func (client MetricNamespacesClient) List(ctx context.Context, resourceURI string, startTime string) (result MetricNamespaceCollection, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/MetricNamespacesClient.List")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.ListPreparer(ctx, resourceURI, startTime)
    if err != nil {
    err = autorest.NewErrorWithError(err, "insights.MetricNamespacesClient", "List", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "insights.MetricNamespacesClient", "List", resp, "Failure sending request")
            return
            }

            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "insights.MetricNamespacesClient", "List", resp, "Failure responding to request")
            }

    return
    }

    // ListPreparer prepares the List request.
    func (client MetricNamespacesClient) ListPreparer(ctx context.Context, resourceURI string, startTime string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceUri": resourceURI,
            }

                        const APIVersion = "2017-12-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }
            if len(startTime) > 0 {
            queryParameters["startTime"] = autorest.Encode("query",startTime)
            }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/{resourceUri}/providers/microsoft.insights/metricNamespaces",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client MetricNamespacesClient) ListSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client MetricNamespacesClient) ListResponder(resp *http.Response) (result MetricNamespaceCollection, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

