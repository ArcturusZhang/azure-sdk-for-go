package armcompute_test

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/compute/armcompute"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/mock"
)

func TestAvailabilitySetsClient_CreateOrUpdate(t *testing.T) {
	const response = `{
        "sku": {
          "name": "Classic"
        },
        "name": "myAvailabilitySet",
        "properties": {
          "platformFaultDomainCount": 2,
          "platformUpdateDomainCount": 20
        },
        "location": "westus",
        "type": "Microsoft.Compute/availabilitySets",
        "id": "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/availabilitySets/myAvailabilitySet"
      }`
	const request = `{
      "location": "westus",
      "properties": {
        "platformFaultDomainCount": 2,
        "platformUpdateDomainCount": 20
      }
    }`
	var expectedReq, expectedResp armcompute.AvailabilitySet
	if err := json.Unmarshal([]byte(request), &expectedReq); err != nil {
		t.Fatalf("cannot unmarshal request: %+v", err)
	}
	if err := json.Unmarshal([]byte(response), &expectedResp); err != nil {
		t.Fatalf("cannot unmarshal response: %+v", err)
	}
	srv, close := mock.NewServer()
	defer close()
	srv.AppendResponse(mock.WithPredicate(func(req *http.Request) bool {
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			t.Fatalf("cannot read body: %+v", err)
		}
		var reqBody armcompute.AvailabilitySet
		_ = json.Unmarshal(b, &reqBody)
		r := reflect.DeepEqual(reqBody, expectedReq)
		return r
	}),
		mock.WithBody([]byte(response)), // this is for the predicate returns true
	)
	srv.AppendResponse(
		mock.WithStatusCode(http.StatusBadRequest), // this is for the predicate returns false
	)
	conn := armcore.NewConnection(srv.URL(), mockTokenCred{}, nil)
	client := armcompute.NewAvailabilitySetsClient(conn, "mockSubs")
	resp, err := client.CreateOrUpdate(context.Background(), "mock", "mock", expectedReq, nil)
	if err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}
	if resp.AvailabilitySet == nil {
		t.Fatalf("resp is empty")
	}
	if !reflect.DeepEqual(expectedResp, *resp.AvailabilitySet) {
		t.Fatalf("expected %+v, but got %+v", expectedResp, *resp.AvailabilitySet)
	}
}

type mockTokenCred struct{}

func (mockTokenCred) AuthenticationPolicy(azcore.AuthenticationPolicyOptions) azcore.Policy {
	return azcore.PolicyFunc(func(req *azcore.Request) (*azcore.Response, error) {
		return req.Next()
	})
}

func (mockTokenCred) GetToken(context.Context, azcore.TokenRequestOptions) (*azcore.AccessToken, error) {
	return &azcore.AccessToken{
		Token:     "abc123",
		ExpiresOn: time.Now().Add(1 * time.Hour),
	}, nil
}
