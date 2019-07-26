package servicefabric

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

// MeshServiceClient is the service Fabric REST Client APIs allows management of Service Fabric clusters, applications
// and services.
type MeshServiceClient struct {
    BaseClient
}
// NewMeshServiceClient creates an instance of the MeshServiceClient client.
func NewMeshServiceClient() MeshServiceClient {
    return NewMeshServiceClientWithBaseURI(DefaultBaseURI, )
}

// NewMeshServiceClientWithBaseURI creates an instance of the MeshServiceClient client.
    func NewMeshServiceClientWithBaseURI(baseURI string, ) MeshServiceClient {
        return MeshServiceClient{ NewWithBaseURI(baseURI, )}
    }

// Get gets the information about the Service resource with the given name. The information include the description and
// other properties of the Service.
    // Parameters:
        // applicationResourceName - the identity of the application.
        // serviceResourceName - the identity of the service.
func (client MeshServiceClient) Get(ctx context.Context, applicationResourceName string, serviceResourceName string) (result ServiceResourceDescription, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/MeshServiceClient.Get")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.GetPreparer(ctx, applicationResourceName, serviceResourceName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "servicefabric.MeshServiceClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "servicefabric.MeshServiceClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "servicefabric.MeshServiceClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client MeshServiceClient) GetPreparer(ctx context.Context, applicationResourceName string, serviceResourceName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "applicationResourceName": applicationResourceName,
            "serviceResourceName": serviceResourceName,
            }

                        const APIVersion = "6.4-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/Resources/Applications/{applicationResourceName}/Services/{serviceResourceName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client MeshServiceClient) GetSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client MeshServiceClient) GetResponder(resp *http.Response) (result ServiceResourceDescription, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// List gets the information about all services of an application resource. The information include the description and
// other properties of the Service.
    // Parameters:
        // applicationResourceName - the identity of the application.
func (client MeshServiceClient) List(ctx context.Context, applicationResourceName string) (result PagedServiceResourceDescriptionList, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/MeshServiceClient.List")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.ListPreparer(ctx, applicationResourceName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "servicefabric.MeshServiceClient", "List", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "servicefabric.MeshServiceClient", "List", resp, "Failure sending request")
            return
            }

            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "servicefabric.MeshServiceClient", "List", resp, "Failure responding to request")
            }

    return
    }

    // ListPreparer prepares the List request.
    func (client MeshServiceClient) ListPreparer(ctx context.Context, applicationResourceName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "applicationResourceName": applicationResourceName,
            }

                        const APIVersion = "6.4-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/Resources/Applications/{applicationResourceName}/Services",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client MeshServiceClient) ListSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client MeshServiceClient) ListResponder(resp *http.Response) (result PagedServiceResourceDescriptionList, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

