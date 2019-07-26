package face

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

// PersonGroupClient is the an API for face detection, verification, and identification.
type PersonGroupClient struct {
    BaseClient
}
// NewPersonGroupClient creates an instance of the PersonGroupClient client.
func NewPersonGroupClient(endpoint string) PersonGroupClient {
    return PersonGroupClient{ New(endpoint)}
}

// Create create a new person group with specified personGroupId, name, user-provided userData and recognitionModel.
// <br /> A person group is the container of the uploaded person data, including face recognition features.
// <br /> After creation, use [PersonGroup Person -
// Create](/docs/services/563879b61984550e40cbbe8d/operations/563879b61984550f3039523c) to add persons into the group,
// and then call [PersonGroup - Train](/docs/services/563879b61984550e40cbbe8d/operations/563879b61984550f30395249) to
// get this group ready for [Face -
// Identify](/docs/services/563879b61984550e40cbbe8d/operations/563879b61984550f30395239).
// <br /> No image will be stored. Only the person's extracted face features and userData will be stored on server
// until [PersonGroup Person - Delete](/docs/services/563879b61984550e40cbbe8d/operations/563879b61984550f3039523d) or
// [PersonGroup - Delete](/docs/services/563879b61984550e40cbbe8d/operations/563879b61984550f30395245) is called.
// <br/>'recognitionModel' should be specified to associate with this person group. The default value for
// 'recognitionModel' is 'recognition_01', if the latest model needed, please explicitly specify the model you need in
// this parameter. New faces that are added to an existing person group will use the recognition model that's already
// associated with the collection. Existing face features in a person group can't be updated to features extracted by
// another version of recognition model.
// * 'recognition_01': The default recognition model for [PersonGroup -
// Create](/docs/services/563879b61984550e40cbbe8d/operations/563879b61984550f30395244). All those person groups
// created before 2019 March are bonded with this recognition model.
// * 'recognition_02': Recognition model released in 2019 March. 'recognition_02' is recommended since its overall
// accuracy is improved compared with 'recognition_01'.
//
// Person group quota:
// * Free-tier subscription quota: 1,000 person groups. Each holds up to 1,000 persons.
// * S0-tier subscription quota: 1,000,000 person groups. Each holds up to 10,000 persons.
// * to handle larger scale face identification problem, please consider using
// [LargePersonGroup](/docs/services/563879b61984550e40cbbe8d/operations/599acdee6ac60f11b48b5a9d).
    // Parameters:
        // personGroupID - id referencing a particular person group.
        // body - request body for creating new person group.
func (client PersonGroupClient) Create(ctx context.Context, personGroupID string, body MetaDataContract) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/PersonGroupClient.Create")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: personGroupID,
             Constraints: []validation.Constraint{	{Target: "personGroupID", Name: validation.MaxLength, Rule: 64, Chain: nil },
            	{Target: "personGroupID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("face.PersonGroupClient", "Create", err.Error())
            }

                req, err := client.CreatePreparer(ctx, personGroupID, body)
    if err != nil {
    err = autorest.NewErrorWithError(err, "face.PersonGroupClient", "Create", nil , "Failure preparing request")
    return
    }

            resp, err := client.CreateSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "face.PersonGroupClient", "Create", resp, "Failure sending request")
            return
            }

            result, err = client.CreateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "face.PersonGroupClient", "Create", resp, "Failure responding to request")
            }

    return
    }

    // CreatePreparer prepares the Create request.
    func (client PersonGroupClient) CreatePreparer(ctx context.Context, personGroupID string, body MetaDataContract) (*http.Request, error) {
            urlParameters := map[string]interface{} {
            "Endpoint": client.Endpoint,
            }

            pathParameters := map[string]interface{} {
            "personGroupId": autorest.Encode("path",personGroupID),
            }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPut(),
    autorest.WithCustomBaseURL("{Endpoint}/face/v1.0", urlParameters),
    autorest.WithPathParameters("/persongroups/{personGroupId}",pathParameters),
    autorest.WithJSON(body))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateSender sends the Create request. The method will close the
    // http.Response Body if it receives an error.
    func (client PersonGroupClient) CreateSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client PersonGroupClient) CreateResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Delete delete an existing person group. Persisted face features of all people in the person group will also be
// deleted.
    // Parameters:
        // personGroupID - id referencing a particular person group.
func (client PersonGroupClient) Delete(ctx context.Context, personGroupID string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/PersonGroupClient.Delete")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: personGroupID,
             Constraints: []validation.Constraint{	{Target: "personGroupID", Name: validation.MaxLength, Rule: 64, Chain: nil },
            	{Target: "personGroupID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("face.PersonGroupClient", "Delete", err.Error())
            }

                req, err := client.DeletePreparer(ctx, personGroupID)
    if err != nil {
    err = autorest.NewErrorWithError(err, "face.PersonGroupClient", "Delete", nil , "Failure preparing request")
    return
    }

            resp, err := client.DeleteSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "face.PersonGroupClient", "Delete", resp, "Failure sending request")
            return
            }

            result, err = client.DeleteResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "face.PersonGroupClient", "Delete", resp, "Failure responding to request")
            }

    return
    }

    // DeletePreparer prepares the Delete request.
    func (client PersonGroupClient) DeletePreparer(ctx context.Context, personGroupID string) (*http.Request, error) {
            urlParameters := map[string]interface{} {
            "Endpoint": client.Endpoint,
            }

            pathParameters := map[string]interface{} {
            "personGroupId": autorest.Encode("path",personGroupID),
            }

        preparer := autorest.CreatePreparer(
    autorest.AsDelete(),
    autorest.WithCustomBaseURL("{Endpoint}/face/v1.0", urlParameters),
    autorest.WithPathParameters("/persongroups/{personGroupId}",pathParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client PersonGroupClient) DeleteSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client PersonGroupClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Get retrieve person group name, userData and recognitionModel. To get person information under this personGroup, use
// [PersonGroup Person - List](/docs/services/563879b61984550e40cbbe8d/operations/563879b61984550f30395241).
    // Parameters:
        // personGroupID - id referencing a particular person group.
        // returnRecognitionModel - a value indicating whether the operation should return 'recognitionModel' in
        // response.
func (client PersonGroupClient) Get(ctx context.Context, personGroupID string, returnRecognitionModel *bool) (result PersonGroup, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/PersonGroupClient.Get")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: personGroupID,
             Constraints: []validation.Constraint{	{Target: "personGroupID", Name: validation.MaxLength, Rule: 64, Chain: nil },
            	{Target: "personGroupID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("face.PersonGroupClient", "Get", err.Error())
            }

                req, err := client.GetPreparer(ctx, personGroupID, returnRecognitionModel)
    if err != nil {
    err = autorest.NewErrorWithError(err, "face.PersonGroupClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "face.PersonGroupClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "face.PersonGroupClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client PersonGroupClient) GetPreparer(ctx context.Context, personGroupID string, returnRecognitionModel *bool) (*http.Request, error) {
            urlParameters := map[string]interface{} {
            "Endpoint": client.Endpoint,
            }

            pathParameters := map[string]interface{} {
            "personGroupId": autorest.Encode("path",personGroupID),
            }

                    queryParameters := map[string]interface{} {
        }
            if returnRecognitionModel != nil {
            queryParameters["returnRecognitionModel"] = autorest.Encode("query",*returnRecognitionModel)
                } else {
                queryParameters["returnRecognitionModel"] = autorest.Encode("query",false)
            }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithCustomBaseURL("{Endpoint}/face/v1.0", urlParameters),
    autorest.WithPathParameters("/persongroups/{personGroupId}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client PersonGroupClient) GetSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PersonGroupClient) GetResponder(resp *http.Response) (result PersonGroup, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// GetTrainingStatus retrieve the training status of a person group (completed or ongoing).
    // Parameters:
        // personGroupID - id referencing a particular person group.
func (client PersonGroupClient) GetTrainingStatus(ctx context.Context, personGroupID string) (result TrainingStatus, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/PersonGroupClient.GetTrainingStatus")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: personGroupID,
             Constraints: []validation.Constraint{	{Target: "personGroupID", Name: validation.MaxLength, Rule: 64, Chain: nil },
            	{Target: "personGroupID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("face.PersonGroupClient", "GetTrainingStatus", err.Error())
            }

                req, err := client.GetTrainingStatusPreparer(ctx, personGroupID)
    if err != nil {
    err = autorest.NewErrorWithError(err, "face.PersonGroupClient", "GetTrainingStatus", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetTrainingStatusSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "face.PersonGroupClient", "GetTrainingStatus", resp, "Failure sending request")
            return
            }

            result, err = client.GetTrainingStatusResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "face.PersonGroupClient", "GetTrainingStatus", resp, "Failure responding to request")
            }

    return
    }

    // GetTrainingStatusPreparer prepares the GetTrainingStatus request.
    func (client PersonGroupClient) GetTrainingStatusPreparer(ctx context.Context, personGroupID string) (*http.Request, error) {
            urlParameters := map[string]interface{} {
            "Endpoint": client.Endpoint,
            }

            pathParameters := map[string]interface{} {
            "personGroupId": autorest.Encode("path",personGroupID),
            }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithCustomBaseURL("{Endpoint}/face/v1.0", urlParameters),
    autorest.WithPathParameters("/persongroups/{personGroupId}/training",pathParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetTrainingStatusSender sends the GetTrainingStatus request. The method will close the
    // http.Response Body if it receives an error.
    func (client PersonGroupClient) GetTrainingStatusSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetTrainingStatusResponder handles the response to the GetTrainingStatus request. The method always
// closes the http.Response Body.
func (client PersonGroupClient) GetTrainingStatusResponder(resp *http.Response) (result TrainingStatus, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// List list person groups’ personGroupId, name, userData and recognitionModel.<br />
// * Person groups are stored in alphabetical order of personGroupId.
// * "start" parameter (string, optional) is a user-provided personGroupId value that returned entries have larger ids
// by string comparison. "start" set to empty to indicate return from the first item.
// * "top" parameter (int, optional) specifies the number of entries to return. A maximal of 1000 entries can be
// returned in one call. To fetch more, you can specify "start" with the last returned entry’s Id of the current call.
// <br />
// For example, total 5 person groups: "group1", ..., "group5".
// <br /> "start=&top=" will return all 5 groups.
// <br /> "start=&top=2" will return "group1", "group2".
// <br /> "start=group2&top=3" will return "group3", "group4", "group5".
    // Parameters:
        // start - list person groups from the least personGroupId greater than the "start".
        // top - the number of person groups to list.
        // returnRecognitionModel - a value indicating whether the operation should return 'recognitionModel' in
        // response.
func (client PersonGroupClient) List(ctx context.Context, start string, top *int32, returnRecognitionModel *bool) (result ListPersonGroup, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/PersonGroupClient.List")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: start,
             Constraints: []validation.Constraint{	{Target: "start", Name: validation.MaxLength, Rule: 64, Chain: nil }}},
            { TargetValue: top,
             Constraints: []validation.Constraint{	{Target: "top", Name: validation.Null, Rule: false ,
            Chain: []validation.Constraint{	{Target: "top", Name: validation.InclusiveMaximum, Rule: int64(1000), Chain: nil },
            	{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil },
            }}}}}); err != nil {
            return result, validation.NewError("face.PersonGroupClient", "List", err.Error())
            }

                req, err := client.ListPreparer(ctx, start, top, returnRecognitionModel)
    if err != nil {
    err = autorest.NewErrorWithError(err, "face.PersonGroupClient", "List", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "face.PersonGroupClient", "List", resp, "Failure sending request")
            return
            }

            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "face.PersonGroupClient", "List", resp, "Failure responding to request")
            }

    return
    }

    // ListPreparer prepares the List request.
    func (client PersonGroupClient) ListPreparer(ctx context.Context, start string, top *int32, returnRecognitionModel *bool) (*http.Request, error) {
            urlParameters := map[string]interface{} {
            "Endpoint": client.Endpoint,
            }

                        queryParameters := map[string]interface{} {
        }
            if len(start) > 0 {
            queryParameters["start"] = autorest.Encode("query",start)
            }
            if top != nil {
            queryParameters["top"] = autorest.Encode("query",*top)
                } else {
                queryParameters["top"] = autorest.Encode("query",1000)
            }
            if returnRecognitionModel != nil {
            queryParameters["returnRecognitionModel"] = autorest.Encode("query",*returnRecognitionModel)
                } else {
                queryParameters["returnRecognitionModel"] = autorest.Encode("query",false)
            }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithCustomBaseURL("{Endpoint}/face/v1.0", urlParameters),
    autorest.WithPath("/persongroups"),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client PersonGroupClient) ListSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client PersonGroupClient) ListResponder(resp *http.Response) (result ListPersonGroup, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result.Value),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Train queue a person group training task, the training task may not be started immediately.
    // Parameters:
        // personGroupID - id referencing a particular person group.
func (client PersonGroupClient) Train(ctx context.Context, personGroupID string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/PersonGroupClient.Train")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: personGroupID,
             Constraints: []validation.Constraint{	{Target: "personGroupID", Name: validation.MaxLength, Rule: 64, Chain: nil },
            	{Target: "personGroupID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("face.PersonGroupClient", "Train", err.Error())
            }

                req, err := client.TrainPreparer(ctx, personGroupID)
    if err != nil {
    err = autorest.NewErrorWithError(err, "face.PersonGroupClient", "Train", nil , "Failure preparing request")
    return
    }

            resp, err := client.TrainSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "face.PersonGroupClient", "Train", resp, "Failure sending request")
            return
            }

            result, err = client.TrainResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "face.PersonGroupClient", "Train", resp, "Failure responding to request")
            }

    return
    }

    // TrainPreparer prepares the Train request.
    func (client PersonGroupClient) TrainPreparer(ctx context.Context, personGroupID string) (*http.Request, error) {
            urlParameters := map[string]interface{} {
            "Endpoint": client.Endpoint,
            }

            pathParameters := map[string]interface{} {
            "personGroupId": autorest.Encode("path",personGroupID),
            }

        preparer := autorest.CreatePreparer(
    autorest.AsPost(),
    autorest.WithCustomBaseURL("{Endpoint}/face/v1.0", urlParameters),
    autorest.WithPathParameters("/persongroups/{personGroupId}/train",pathParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // TrainSender sends the Train request. The method will close the
    // http.Response Body if it receives an error.
    func (client PersonGroupClient) TrainSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// TrainResponder handles the response to the Train request. The method always
// closes the http.Response Body.
func (client PersonGroupClient) TrainResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Update update an existing person group's display name and userData. The properties which does not appear in request
// body will not be updated.
    // Parameters:
        // personGroupID - id referencing a particular person group.
        // body - request body for updating person group.
func (client PersonGroupClient) Update(ctx context.Context, personGroupID string, body NameAndUserDataContract) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/PersonGroupClient.Update")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: personGroupID,
             Constraints: []validation.Constraint{	{Target: "personGroupID", Name: validation.MaxLength, Rule: 64, Chain: nil },
            	{Target: "personGroupID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("face.PersonGroupClient", "Update", err.Error())
            }

                req, err := client.UpdatePreparer(ctx, personGroupID, body)
    if err != nil {
    err = autorest.NewErrorWithError(err, "face.PersonGroupClient", "Update", nil , "Failure preparing request")
    return
    }

            resp, err := client.UpdateSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "face.PersonGroupClient", "Update", resp, "Failure sending request")
            return
            }

            result, err = client.UpdateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "face.PersonGroupClient", "Update", resp, "Failure responding to request")
            }

    return
    }

    // UpdatePreparer prepares the Update request.
    func (client PersonGroupClient) UpdatePreparer(ctx context.Context, personGroupID string, body NameAndUserDataContract) (*http.Request, error) {
            urlParameters := map[string]interface{} {
            "Endpoint": client.Endpoint,
            }

            pathParameters := map[string]interface{} {
            "personGroupId": autorest.Encode("path",personGroupID),
            }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPatch(),
    autorest.WithCustomBaseURL("{Endpoint}/face/v1.0", urlParameters),
    autorest.WithPathParameters("/persongroups/{personGroupId}",pathParameters),
    autorest.WithJSON(body))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // UpdateSender sends the Update request. The method will close the
    // http.Response Body if it receives an error.
    func (client PersonGroupClient) UpdateSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client PersonGroupClient) UpdateResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByClosing())
    result.Response = resp
        return
    }

