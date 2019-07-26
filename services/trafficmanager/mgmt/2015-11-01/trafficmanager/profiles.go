package trafficmanager

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

// ProfilesClient is the client for the Profiles methods of the Trafficmanager service.
type ProfilesClient struct {
    BaseClient
}
// NewProfilesClient creates an instance of the ProfilesClient client.
func NewProfilesClient(subscriptionID string) ProfilesClient {
    return NewProfilesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProfilesClientWithBaseURI creates an instance of the ProfilesClient client.
    func NewProfilesClientWithBaseURI(baseURI string, subscriptionID string) ProfilesClient {
        return ProfilesClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CheckTrafficManagerRelativeDNSNameAvailability checks the availability of a Traffic Manager Relative DNS name.
    // Parameters:
        // parameters - the Traffic Manager name parameters supplied to the CheckTrafficManagerNameAvailability
        // operation.
func (client ProfilesClient) CheckTrafficManagerRelativeDNSNameAvailability(ctx context.Context, parameters CheckTrafficManagerRelativeDNSNameAvailabilityParameters) (result NameAvailability, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ProfilesClient.CheckTrafficManagerRelativeDNSNameAvailability")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.CheckTrafficManagerRelativeDNSNameAvailabilityPreparer(ctx, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "trafficmanager.ProfilesClient", "CheckTrafficManagerRelativeDNSNameAvailability", nil , "Failure preparing request")
    return
    }

            resp, err := client.CheckTrafficManagerRelativeDNSNameAvailabilitySender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "trafficmanager.ProfilesClient", "CheckTrafficManagerRelativeDNSNameAvailability", resp, "Failure sending request")
            return
            }

            result, err = client.CheckTrafficManagerRelativeDNSNameAvailabilityResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "trafficmanager.ProfilesClient", "CheckTrafficManagerRelativeDNSNameAvailability", resp, "Failure responding to request")
            }

    return
    }

    // CheckTrafficManagerRelativeDNSNameAvailabilityPreparer prepares the CheckTrafficManagerRelativeDNSNameAvailability request.
    func (client ProfilesClient) CheckTrafficManagerRelativeDNSNameAvailabilityPreparer(ctx context.Context, parameters CheckTrafficManagerRelativeDNSNameAvailabilityParameters) (*http.Request, error) {
                    const APIVersion = "2015-11-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPath("/providers/Microsoft.Network/checkTrafficManagerNameAvailability"),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CheckTrafficManagerRelativeDNSNameAvailabilitySender sends the CheckTrafficManagerRelativeDNSNameAvailability request. The method will close the
    // http.Response Body if it receives an error.
    func (client ProfilesClient) CheckTrafficManagerRelativeDNSNameAvailabilitySender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// CheckTrafficManagerRelativeDNSNameAvailabilityResponder handles the response to the CheckTrafficManagerRelativeDNSNameAvailability request. The method always
// closes the http.Response Body.
func (client ProfilesClient) CheckTrafficManagerRelativeDNSNameAvailabilityResponder(resp *http.Response) (result NameAvailability, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// CreateOrUpdate create or update a Traffic Manager profile.
    // Parameters:
        // resourceGroupName - the name of the resource group containing the Traffic Manager profile.
        // profileName - the name of the Traffic Manager profile.
        // parameters - the Traffic Manager profile parameters supplied to the CreateOrUpdate operation.
func (client ProfilesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, profileName string, parameters Profile) (result Profile, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ProfilesClient.CreateOrUpdate")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, profileName, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "trafficmanager.ProfilesClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

            resp, err := client.CreateOrUpdateSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "trafficmanager.ProfilesClient", "CreateOrUpdate", resp, "Failure sending request")
            return
            }

            result, err = client.CreateOrUpdateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "trafficmanager.ProfilesClient", "CreateOrUpdate", resp, "Failure responding to request")
            }

    return
    }

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client ProfilesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, profileName string, parameters Profile) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "profileName": autorest.Encode("path",profileName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2015-11-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPut(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client ProfilesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ProfilesClient) CreateOrUpdateResponder(resp *http.Response) (result Profile, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Delete deletes a Traffic Manager profile.
    // Parameters:
        // resourceGroupName - the name of the resource group containing the Traffic Manager profile to be deleted.
        // profileName - the name of the Traffic Manager profile to be deleted.
func (client ProfilesClient) Delete(ctx context.Context, resourceGroupName string, profileName string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ProfilesClient.Delete")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.DeletePreparer(ctx, resourceGroupName, profileName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "trafficmanager.ProfilesClient", "Delete", nil , "Failure preparing request")
    return
    }

            resp, err := client.DeleteSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "trafficmanager.ProfilesClient", "Delete", resp, "Failure sending request")
            return
            }

            result, err = client.DeleteResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "trafficmanager.ProfilesClient", "Delete", resp, "Failure responding to request")
            }

    return
    }

    // DeletePreparer prepares the Delete request.
    func (client ProfilesClient) DeletePreparer(ctx context.Context, resourceGroupName string, profileName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "profileName": autorest.Encode("path",profileName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2015-11-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsDelete(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client ProfilesClient) DeleteSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ProfilesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Get gets a Traffic Manager profile.
    // Parameters:
        // resourceGroupName - the name of the resource group containing the Traffic Manager profile.
        // profileName - the name of the Traffic Manager profile.
func (client ProfilesClient) Get(ctx context.Context, resourceGroupName string, profileName string) (result Profile, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ProfilesClient.Get")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.GetPreparer(ctx, resourceGroupName, profileName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "trafficmanager.ProfilesClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "trafficmanager.ProfilesClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "trafficmanager.ProfilesClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client ProfilesClient) GetPreparer(ctx context.Context, resourceGroupName string, profileName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "profileName": autorest.Encode("path",profileName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2015-11-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client ProfilesClient) GetSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ProfilesClient) GetResponder(resp *http.Response) (result Profile, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// ListAll lists all Traffic Manager profiles within a subscription.
func (client ProfilesClient) ListAll(ctx context.Context) (result ProfileListResult, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ProfilesClient.ListAll")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.ListAllPreparer(ctx)
    if err != nil {
    err = autorest.NewErrorWithError(err, "trafficmanager.ProfilesClient", "ListAll", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListAllSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "trafficmanager.ProfilesClient", "ListAll", resp, "Failure sending request")
            return
            }

            result, err = client.ListAllResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "trafficmanager.ProfilesClient", "ListAll", resp, "Failure responding to request")
            }

    return
    }

    // ListAllPreparer prepares the ListAll request.
    func (client ProfilesClient) ListAllPreparer(ctx context.Context) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2015-11-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/trafficmanagerprofiles",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListAllSender sends the ListAll request. The method will close the
    // http.Response Body if it receives an error.
    func (client ProfilesClient) ListAllSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListAllResponder handles the response to the ListAll request. The method always
// closes the http.Response Body.
func (client ProfilesClient) ListAllResponder(resp *http.Response) (result ProfileListResult, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// ListAllInResourceGroup lists all Traffic Manager profiles within a resource group.
    // Parameters:
        // resourceGroupName - the name of the resource group containing the Traffic Manager profiles to be listed.
func (client ProfilesClient) ListAllInResourceGroup(ctx context.Context, resourceGroupName string) (result ProfileListResult, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ProfilesClient.ListAllInResourceGroup")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.ListAllInResourceGroupPreparer(ctx, resourceGroupName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "trafficmanager.ProfilesClient", "ListAllInResourceGroup", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListAllInResourceGroupSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "trafficmanager.ProfilesClient", "ListAllInResourceGroup", resp, "Failure sending request")
            return
            }

            result, err = client.ListAllInResourceGroupResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "trafficmanager.ProfilesClient", "ListAllInResourceGroup", resp, "Failure responding to request")
            }

    return
    }

    // ListAllInResourceGroupPreparer prepares the ListAllInResourceGroup request.
    func (client ProfilesClient) ListAllInResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2015-11-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListAllInResourceGroupSender sends the ListAllInResourceGroup request. The method will close the
    // http.Response Body if it receives an error.
    func (client ProfilesClient) ListAllInResourceGroupSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListAllInResourceGroupResponder handles the response to the ListAllInResourceGroup request. The method always
// closes the http.Response Body.
func (client ProfilesClient) ListAllInResourceGroupResponder(resp *http.Response) (result ProfileListResult, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Update update a Traffic Manager profile.
    // Parameters:
        // resourceGroupName - the name of the resource group containing the Traffic Manager profile.
        // profileName - the name of the Traffic Manager profile.
        // parameters - the Traffic Manager profile parameters supplied to the Update operation.
func (client ProfilesClient) Update(ctx context.Context, resourceGroupName string, profileName string, parameters Profile) (result Profile, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ProfilesClient.Update")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.UpdatePreparer(ctx, resourceGroupName, profileName, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "trafficmanager.ProfilesClient", "Update", nil , "Failure preparing request")
    return
    }

            resp, err := client.UpdateSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "trafficmanager.ProfilesClient", "Update", resp, "Failure sending request")
            return
            }

            result, err = client.UpdateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "trafficmanager.ProfilesClient", "Update", resp, "Failure responding to request")
            }

    return
    }

    // UpdatePreparer prepares the Update request.
    func (client ProfilesClient) UpdatePreparer(ctx context.Context, resourceGroupName string, profileName string, parameters Profile) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "profileName": autorest.Encode("path",profileName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2015-11-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPatch(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // UpdateSender sends the Update request. The method will close the
    // http.Response Body if it receives an error.
    func (client ProfilesClient) UpdateSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ProfilesClient) UpdateResponder(resp *http.Response) (result Profile, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

