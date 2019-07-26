// Package anomalydetector implements the Azure ARM Anomalydetector service API version 1.0.
//
// The Anomaly Detector API detects anomalies automatically in time series data. It supports two functionalities, one
// is for detecting the whole series with model trained by the timeseries, another is detecting last point with model
// trained by points before. By using this service, business customers can discover incidents and establish a logic
// flow for root cause analysis.
package anomalydetector

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
    "context"
"crypto/tls"
    "github.com/Azure/go-autorest/autorest"
    "github.com/Azure/go-autorest/autorest/azure"
    "github.com/Azure/go-autorest/autorest/validation"
    "github.com/Azure/go-autorest/tracing"
    "net/http"
)


// BaseClient is the base client for Anomalydetector.
type BaseClient struct {
    autorest.Client
            Endpoint string
}

// New creates an instance of the BaseClient client.
func New(endpoint string)BaseClient {
    return NewWithoutDefaults(endpoint)
}

// NewWithoutDefaults creates an instance of the BaseClient client.
func NewWithoutDefaults(endpoint string) BaseClient {
    return BaseClient{
        Client: autorest.NewClientWithOptions(autorest.ClientOptions{UserAgent: UserAgent(), Renegotiation: tls.RenegotiateFreelyAsClient}),
                Endpoint: endpoint,
    }
}

    // EntireDetect this operation generates a model using an entire series, each point is detected with the same model.
    // With this method, points before and after a certain point are used to determine whether it is an anomaly. The entire
    // detection can give user an overall status of the time series.
        // Parameters:
            // body - time series points and period if needed. Advanced model parameters can also be set in the request.
    func (client BaseClient) EntireDetect(ctx context.Context, body Request) (result EntireDetectResponse, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/BaseClient.EntireDetect")
            defer func() {
                sc := -1
                if result.Response.Response != nil {
                    sc = result.Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
        }
                if err := validation.Validate([]validation.Validation{
                { TargetValue: body,
                 Constraints: []validation.Constraint{	{Target: "body.Series", Name: validation.Null, Rule: true, Chain: nil }}}}); err != nil {
                return result, validation.NewError("anomalydetector.BaseClient", "EntireDetect", err.Error())
                }

                    req, err := client.EntireDetectPreparer(ctx, body)
        if err != nil {
        err = autorest.NewErrorWithError(err, "anomalydetector.BaseClient", "EntireDetect", nil , "Failure preparing request")
        return
        }

                resp, err := client.EntireDetectSender(req)
                if err != nil {
                result.Response = autorest.Response{Response: resp}
                err = autorest.NewErrorWithError(err, "anomalydetector.BaseClient", "EntireDetect", resp, "Failure sending request")
                return
                }

                result, err = client.EntireDetectResponder(resp)
                if err != nil {
                err = autorest.NewErrorWithError(err, "anomalydetector.BaseClient", "EntireDetect", resp, "Failure responding to request")
                }

        return
        }

        // EntireDetectPreparer prepares the EntireDetect request.
        func (client BaseClient) EntireDetectPreparer(ctx context.Context, body Request) (*http.Request, error) {
                urlParameters := map[string]interface{} {
                "Endpoint": client.Endpoint,
                }

            preparer := autorest.CreatePreparer(
        autorest.AsContentType("application/json; charset=utf-8"),
        autorest.AsPost(),
        autorest.WithCustomBaseURL("{Endpoint}/anomalydetector/v1.0", urlParameters),
        autorest.WithPath("/timeseries/entire/detect"),
        autorest.WithJSON(body))
        return preparer.Prepare((&http.Request{}).WithContext(ctx))
        }

        // EntireDetectSender sends the EntireDetect request. The method will close the
        // http.Response Body if it receives an error.
        func (client BaseClient) EntireDetectSender(req *http.Request) (*http.Response, error) {
            sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
                return autorest.SendWithSender(client, req, sd...)
                }

    // EntireDetectResponder handles the response to the EntireDetect request. The method always
    // closes the http.Response Body.
    func (client BaseClient) EntireDetectResponder(resp *http.Response) (result EntireDetectResponse, err error) {
        err = autorest.Respond(
        resp,
        client.ByInspecting(),
        azure.WithErrorUnlessStatusCode(http.StatusOK),
        autorest.ByUnmarshallingJSON(&result),
        autorest.ByClosing())
        result.Response = autorest.Response{Response: resp}
            return
        }

    // LastDetect this operation generates a model using points before the latest one. With this method, only historical
    // points are used to determine whether the target point is an anomaly. The latest point detecting operation matches
    // the scenario of real-time monitoring of business metrics.
        // Parameters:
            // body - time series points and period if needed. Advanced model parameters can also be set in the request.
    func (client BaseClient) LastDetect(ctx context.Context, body Request) (result LastDetectResponse, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/BaseClient.LastDetect")
            defer func() {
                sc := -1
                if result.Response.Response != nil {
                    sc = result.Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
        }
                if err := validation.Validate([]validation.Validation{
                { TargetValue: body,
                 Constraints: []validation.Constraint{	{Target: "body.Series", Name: validation.Null, Rule: true, Chain: nil }}}}); err != nil {
                return result, validation.NewError("anomalydetector.BaseClient", "LastDetect", err.Error())
                }

                    req, err := client.LastDetectPreparer(ctx, body)
        if err != nil {
        err = autorest.NewErrorWithError(err, "anomalydetector.BaseClient", "LastDetect", nil , "Failure preparing request")
        return
        }

                resp, err := client.LastDetectSender(req)
                if err != nil {
                result.Response = autorest.Response{Response: resp}
                err = autorest.NewErrorWithError(err, "anomalydetector.BaseClient", "LastDetect", resp, "Failure sending request")
                return
                }

                result, err = client.LastDetectResponder(resp)
                if err != nil {
                err = autorest.NewErrorWithError(err, "anomalydetector.BaseClient", "LastDetect", resp, "Failure responding to request")
                }

        return
        }

        // LastDetectPreparer prepares the LastDetect request.
        func (client BaseClient) LastDetectPreparer(ctx context.Context, body Request) (*http.Request, error) {
                urlParameters := map[string]interface{} {
                "Endpoint": client.Endpoint,
                }

            preparer := autorest.CreatePreparer(
        autorest.AsContentType("application/json; charset=utf-8"),
        autorest.AsPost(),
        autorest.WithCustomBaseURL("{Endpoint}/anomalydetector/v1.0", urlParameters),
        autorest.WithPath("/timeseries/last/detect"),
        autorest.WithJSON(body))
        return preparer.Prepare((&http.Request{}).WithContext(ctx))
        }

        // LastDetectSender sends the LastDetect request. The method will close the
        // http.Response Body if it receives an error.
        func (client BaseClient) LastDetectSender(req *http.Request) (*http.Response, error) {
            sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
                return autorest.SendWithSender(client, req, sd...)
                }

    // LastDetectResponder handles the response to the LastDetect request. The method always
    // closes the http.Response Body.
    func (client BaseClient) LastDetectResponder(resp *http.Response) (result LastDetectResponse, err error) {
        err = autorest.Respond(
        resp,
        client.ByInspecting(),
        azure.WithErrorUnlessStatusCode(http.StatusOK),
        autorest.ByUnmarshallingJSON(&result),
        autorest.ByClosing())
        result.Response = autorest.Response{Response: resp}
            return
        }

