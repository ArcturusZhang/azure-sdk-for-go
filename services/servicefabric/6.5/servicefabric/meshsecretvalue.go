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
    "github.com/Azure/go-autorest/autorest/validation"
)

// MeshSecretValueClient is the service Fabric REST Client APIs allows management of Service Fabric clusters,
// applications and services.
type MeshSecretValueClient struct {
    BaseClient
}
// NewMeshSecretValueClient creates an instance of the MeshSecretValueClient client.
func NewMeshSecretValueClient() MeshSecretValueClient {
    return NewMeshSecretValueClientWithBaseURI(DefaultBaseURI, )
}

// NewMeshSecretValueClientWithBaseURI creates an instance of the MeshSecretValueClient client.
    func NewMeshSecretValueClientWithBaseURI(baseURI string, ) MeshSecretValueClient {
        return MeshSecretValueClient{ NewWithBaseURI(baseURI, )}
    }

// AddValue creates a new value of the specified secret resource. The name of the value is typically the version
// identifier. Once created the value cannot be changed.
    // Parameters:
        // secretResourceName - the name of the secret resource.
        // secretValueResourceName - the name of the secret resource value which is typically the version identifier
        // for the value.
        // secretValueResourceDescription - description for creating a value of a secret resource.
func (client MeshSecretValueClient) AddValue(ctx context.Context, secretResourceName string, secretValueResourceName string, secretValueResourceDescription SecretValueResourceDescription) (result SecretValueResourceDescription, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/MeshSecretValueClient.AddValue")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: secretValueResourceDescription,
             Constraints: []validation.Constraint{	{Target: "secretValueResourceDescription.Name", Name: validation.Null, Rule: true, Chain: nil },
            	{Target: "secretValueResourceDescription.SecretValueResourceProperties", Name: validation.Null, Rule: true, Chain: nil }}}}); err != nil {
            return result, validation.NewError("servicefabric.MeshSecretValueClient", "AddValue", err.Error())
            }

                req, err := client.AddValuePreparer(ctx, secretResourceName, secretValueResourceName, secretValueResourceDescription)
    if err != nil {
    err = autorest.NewErrorWithError(err, "servicefabric.MeshSecretValueClient", "AddValue", nil , "Failure preparing request")
    return
    }

            resp, err := client.AddValueSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "servicefabric.MeshSecretValueClient", "AddValue", resp, "Failure sending request")
            return
            }

            result, err = client.AddValueResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "servicefabric.MeshSecretValueClient", "AddValue", resp, "Failure responding to request")
            }

    return
    }

    // AddValuePreparer prepares the AddValue request.
    func (client MeshSecretValueClient) AddValuePreparer(ctx context.Context, secretResourceName string, secretValueResourceName string, secretValueResourceDescription SecretValueResourceDescription) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "secretResourceName": secretResourceName,
            "secretValueResourceName": secretValueResourceName,
            }

                        const APIVersion = "6.4-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPut(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/Resources/Secrets/{secretResourceName}/values/{secretValueResourceName}",pathParameters),
    autorest.WithJSON(secretValueResourceDescription),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // AddValueSender sends the AddValue request. The method will close the
    // http.Response Body if it receives an error.
    func (client MeshSecretValueClient) AddValueSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// AddValueResponder handles the response to the AddValue request. The method always
// closes the http.Response Body.
func (client MeshSecretValueClient) AddValueResponder(resp *http.Response) (result SecretValueResourceDescription, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated,http.StatusAccepted),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Delete deletes the secret value resource identified by the name. The name of the resource is typically the version
// associated with that value. Deletion will fail if the specified value is in use.
    // Parameters:
        // secretResourceName - the name of the secret resource.
        // secretValueResourceName - the name of the secret resource value which is typically the version identifier
        // for the value.
func (client MeshSecretValueClient) Delete(ctx context.Context, secretResourceName string, secretValueResourceName string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/MeshSecretValueClient.Delete")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.DeletePreparer(ctx, secretResourceName, secretValueResourceName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "servicefabric.MeshSecretValueClient", "Delete", nil , "Failure preparing request")
    return
    }

            resp, err := client.DeleteSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "servicefabric.MeshSecretValueClient", "Delete", resp, "Failure sending request")
            return
            }

            result, err = client.DeleteResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "servicefabric.MeshSecretValueClient", "Delete", resp, "Failure responding to request")
            }

    return
    }

    // DeletePreparer prepares the Delete request.
    func (client MeshSecretValueClient) DeletePreparer(ctx context.Context, secretResourceName string, secretValueResourceName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "secretResourceName": secretResourceName,
            "secretValueResourceName": secretValueResourceName,
            }

                        const APIVersion = "6.4-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsDelete(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/Resources/Secrets/{secretResourceName}/values/{secretValueResourceName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client MeshSecretValueClient) DeleteSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client MeshSecretValueClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Get get the information about the specified named secret value resources. The information does not include the
// actual value of the secret.
    // Parameters:
        // secretResourceName - the name of the secret resource.
        // secretValueResourceName - the name of the secret resource value which is typically the version identifier
        // for the value.
func (client MeshSecretValueClient) Get(ctx context.Context, secretResourceName string, secretValueResourceName string) (result SecretValueResourceDescription, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/MeshSecretValueClient.Get")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.GetPreparer(ctx, secretResourceName, secretValueResourceName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "servicefabric.MeshSecretValueClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "servicefabric.MeshSecretValueClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "servicefabric.MeshSecretValueClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client MeshSecretValueClient) GetPreparer(ctx context.Context, secretResourceName string, secretValueResourceName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "secretResourceName": secretResourceName,
            "secretValueResourceName": secretValueResourceName,
            }

                        const APIVersion = "6.4-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/Resources/Secrets/{secretResourceName}/values/{secretValueResourceName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client MeshSecretValueClient) GetSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client MeshSecretValueClient) GetResponder(resp *http.Response) (result SecretValueResourceDescription, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// List gets information about all secret value resources of the specified secret resource. The information includes
// the names of the secret value resources, but not the actual values.
    // Parameters:
        // secretResourceName - the name of the secret resource.
func (client MeshSecretValueClient) List(ctx context.Context, secretResourceName string) (result PagedSecretValueResourceDescriptionList, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/MeshSecretValueClient.List")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.ListPreparer(ctx, secretResourceName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "servicefabric.MeshSecretValueClient", "List", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "servicefabric.MeshSecretValueClient", "List", resp, "Failure sending request")
            return
            }

            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "servicefabric.MeshSecretValueClient", "List", resp, "Failure responding to request")
            }

    return
    }

    // ListPreparer prepares the List request.
    func (client MeshSecretValueClient) ListPreparer(ctx context.Context, secretResourceName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "secretResourceName": secretResourceName,
            }

                        const APIVersion = "6.4-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/Resources/Secrets/{secretResourceName}/values",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client MeshSecretValueClient) ListSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client MeshSecretValueClient) ListResponder(resp *http.Response) (result PagedSecretValueResourceDescriptionList, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Show lists the decrypted value of the specified named value of the secret resource. This is a privileged operation.
    // Parameters:
        // secretResourceName - the name of the secret resource.
        // secretValueResourceName - the name of the secret resource value which is typically the version identifier
        // for the value.
func (client MeshSecretValueClient) Show(ctx context.Context, secretResourceName string, secretValueResourceName string) (result SecretValue, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/MeshSecretValueClient.Show")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.ShowPreparer(ctx, secretResourceName, secretValueResourceName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "servicefabric.MeshSecretValueClient", "Show", nil , "Failure preparing request")
    return
    }

            resp, err := client.ShowSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "servicefabric.MeshSecretValueClient", "Show", resp, "Failure sending request")
            return
            }

            result, err = client.ShowResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "servicefabric.MeshSecretValueClient", "Show", resp, "Failure responding to request")
            }

    return
    }

    // ShowPreparer prepares the Show request.
    func (client MeshSecretValueClient) ShowPreparer(ctx context.Context, secretResourceName string, secretValueResourceName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "secretResourceName": secretResourceName,
            "secretValueResourceName": secretValueResourceName,
            }

                        const APIVersion = "6.4-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/Resources/Secrets/{secretResourceName}/values/{secretValueResourceName}/list_value",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ShowSender sends the Show request. The method will close the
    // http.Response Body if it receives an error.
    func (client MeshSecretValueClient) ShowSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// ShowResponder handles the response to the Show request. The method always
// closes the http.Response Body.
func (client MeshSecretValueClient) ShowResponder(resp *http.Response) (result SecretValue, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

