package migrate

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

// LocationClient is the move your workloads to Azure.
type LocationClient struct {
    BaseClient
}
// NewLocationClient creates an instance of the LocationClient client.
func NewLocationClient(subscriptionID string, acceptLanguage string) LocationClient {
    return NewLocationClientWithBaseURI(DefaultBaseURI, subscriptionID, acceptLanguage)
}

// NewLocationClientWithBaseURI creates an instance of the LocationClient client.
    func NewLocationClientWithBaseURI(baseURI string, subscriptionID string, acceptLanguage string) LocationClient {
        return LocationClient{ NewWithBaseURI(baseURI, subscriptionID, acceptLanguage)}
    }

// CheckNameAvailability checks whether the project name is available in the specified region.
    // Parameters:
        // locationName - the desired region for the name check.
        // parameters - properties needed to check the availability of a name.
func (client LocationClient) CheckNameAvailability(ctx context.Context, locationName string, parameters CheckNameAvailabilityParameters) (result CheckNameAvailabilityResult, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/LocationClient.CheckNameAvailability")
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
             Constraints: []validation.Constraint{	{Target: "parameters.Name", Name: validation.Null, Rule: true, Chain: nil },
            	{Target: "parameters.Type", Name: validation.Null, Rule: true, Chain: nil }}}}); err != nil {
            return result, validation.NewError("migrate.LocationClient", "CheckNameAvailability", err.Error())
            }

                req, err := client.CheckNameAvailabilityPreparer(ctx, locationName, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "migrate.LocationClient", "CheckNameAvailability", nil , "Failure preparing request")
    return
    }

            resp, err := client.CheckNameAvailabilitySender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "migrate.LocationClient", "CheckNameAvailability", resp, "Failure sending request")
            return
            }

            result, err = client.CheckNameAvailabilityResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "migrate.LocationClient", "CheckNameAvailability", resp, "Failure responding to request")
            }

    return
    }

    // CheckNameAvailabilityPreparer prepares the CheckNameAvailability request.
    func (client LocationClient) CheckNameAvailabilityPreparer(ctx context.Context, locationName string, parameters CheckNameAvailabilityParameters) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "locationName": autorest.Encode("path",locationName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-02-02"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Migrate/locations/{locationName}/checkNameAvailability",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CheckNameAvailabilitySender sends the CheckNameAvailability request. The method will close the
    // http.Response Body if it receives an error.
    func (client LocationClient) CheckNameAvailabilitySender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// CheckNameAvailabilityResponder handles the response to the CheckNameAvailability request. The method always
// closes the http.Response Body.
func (client LocationClient) CheckNameAvailabilityResponder(resp *http.Response) (result CheckNameAvailabilityResult, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

