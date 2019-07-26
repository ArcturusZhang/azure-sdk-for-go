package subscriptions

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
    "github.com/Azure/go-autorest/autorest"
    "github.com/Azure/go-autorest/autorest/to"
    "github.com/Azure/go-autorest/tracing"
    "net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2015-11-01/subscriptions"

            // ListResult subscription list operation response.
            type ListResult struct {
            autorest.Response `json:"-"`
            // Value - Gets or sets subscriptions.
            Value *[]Subscription `json:"value,omitempty"`
            // NextLink - Gets or sets the URL to get the next set of results.
            NextLink *string `json:"nextLink,omitempty"`
            }

            // ListResultIterator provides access to a complete listing of Subscription values.
            type ListResultIterator struct {
                i int
                page ListResultPage
            }
        // NextWithContext advances to the next value.  If there was an error making
        // the request the iterator does not advance and the error is returned.
        func (iter * ListResultIterator) NextWithContext(ctx context.Context) (err error) {
        if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ListResultIterator.NextWithContext")
        defer func() {
        sc := -1
        if iter.Response().Response.Response != nil {
        sc = iter.Response().Response.Response.StatusCode
        }
        tracing.EndSpan(ctx, sc, err)
        }()
        }
        iter.i++
        if iter.i < len(iter. page.Values()) {
        return nil
        }
        err = iter.page.NextWithContext(ctx)
        if err != nil {
        iter. i--
        return err
        }
        iter.i = 0
        return nil
        }
        // Next advances to the next value.  If there was an error making
        // the request the iterator does not advance and the error is returned.
        // Deprecated: Use NextWithContext() instead.
        func (iter * ListResultIterator) Next() error {
        return iter.NextWithContext(context.Background())
        }
        // NotDone returns true if the enumeration should be started or is not yet complete.
        func (iter ListResultIterator) NotDone() bool {
        return iter.page.NotDone() && iter.i < len(iter. page.Values())
        }
        // Response returns the raw server response from the last page request.
        func (iter ListResultIterator) Response() ListResult {
        return iter.page.Response()
        }
        // Value returns the current value or a zero-initialized value if the
        // iterator has advanced beyond the end of the collection.
        func (iter ListResultIterator) Value() Subscription {
        if !iter.page.NotDone() {
        return Subscription{}
        }
        return iter.page.Values()[iter.i]
        }
        // Creates a new instance of the ListResultIterator type.
        func NewListResultIterator (page ListResultPage) ListResultIterator {
            return ListResultIterator{page: page}
        }


                // IsEmpty returns true if the ListResult contains no values.
                func (lr ListResult) IsEmpty() bool {
                return lr.Value == nil || len(*lr.Value) == 0
                }

                    // listResultPreparer prepares a request to retrieve the next set of results.
                    // It returns nil if no more results exist.
                    func (lr ListResult) listResultPreparer(ctx context.Context) (*http.Request, error) {
                    if lr.NextLink == nil || len(to.String(lr.NextLink)) < 1 {
                    return nil, nil
                    }
                    return autorest.Prepare((&http.Request{}).WithContext(ctx),
                    autorest.AsJSON(),
                    autorest.AsGet(),
                    autorest.WithBaseURL(to.String( lr.NextLink)));
                    }

            // ListResultPage contains a page of Subscription values.
            type ListResultPage struct {
                fn func(context.Context, ListResult) (ListResult, error)
                lr ListResult
            }

        // NextWithContext advances to the next page of values.  If there was an error making
        // the request the page does not advance and the error is returned.
        func (page * ListResultPage) NextWithContext(ctx context.Context) (err error) {
        if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ListResultPage.NextWithContext")
        defer func() {
        sc := -1
        if page.Response().Response.Response != nil {
        sc = page.Response().Response.Response.StatusCode
        }
        tracing.EndSpan(ctx, sc, err)
        }()
        }
        next, err := page.fn(ctx, page.lr)
        if err != nil {
        return err
        }
        page.lr = next
        return nil
        }

        // Next advances to the next page of values.  If there was an error making
        // the request the page does not advance and the error is returned.
        // Deprecated: Use NextWithContext() instead.
        func (page * ListResultPage) Next() error {
        return page.NextWithContext(context.Background())
        }
        // NotDone returns true if the page enumeration should be started or is not yet complete.
        func (page ListResultPage) NotDone() bool {
        return !page.lr.IsEmpty()
        }
        // Response returns the raw server response from the last page request.
        func (page ListResultPage) Response() ListResult {
        return page.lr
        }
        // Values returns the slice of values for the current page or nil if there are no values.
        func (page ListResultPage) Values() []Subscription {
        if page.lr.IsEmpty() {
        return nil
        }
        return *page.lr.Value
        }
        // Creates a new instance of the ListResultPage type.
        func NewListResultPage (getNextPage func(context.Context, ListResult) (ListResult, error)) ListResultPage {
            return ListResultPage{fn: getNextPage}
        }

            // Location location information.
            type Location struct {
            // ID - Gets or sets the ID of the resource (/subscriptions/SubscriptionId).
            ID *string `json:"id,omitempty"`
            // SubscriptionID - Gets or sets the subscription Id.
            SubscriptionID *string `json:"subscriptionId,omitempty"`
            // Name - Gets or sets the location name
            Name *string `json:"name,omitempty"`
            // DisplayName - Gets or sets the display name of the location
            DisplayName *string `json:"displayName,omitempty"`
            // Latitude - Gets or sets the latitude of the location
            Latitude *string `json:"latitude,omitempty"`
            // Longitude - Gets or sets the longitude of the location
            Longitude *string `json:"longitude,omitempty"`
            }

            // LocationListResult location list operation response.
            type LocationListResult struct {
            autorest.Response `json:"-"`
            // Value - Gets the locations.
            Value *[]Location `json:"value,omitempty"`
            }

            // Policies subscription policies.
            type Policies struct {
            // LocationPlacementID - Gets or sets the subscription location placement Id.
            LocationPlacementID *string `json:"locationPlacementId,omitempty"`
            // QuotaID - Gets or sets the subscription quota Id.
            QuotaID *string `json:"quotaId,omitempty"`
            }

            // Subscription subscription information.
            type Subscription struct {
            autorest.Response `json:"-"`
            // ID - Gets or sets the ID of the resource (/subscriptions/SubscriptionId).
            ID *string `json:"id,omitempty"`
            // SubscriptionID - Gets or sets the subscription Id.
            SubscriptionID *string `json:"subscriptionId,omitempty"`
            // DisplayName - Gets or sets the subscription display name
            DisplayName *string `json:"displayName,omitempty"`
            // State - Gets or sets the subscription state
            State *string `json:"state,omitempty"`
            // SubscriptionPolicies - Gets or sets the subscription policies.
            SubscriptionPolicies *Policies `json:"subscriptionPolicies,omitempty"`
            }

            // TenantIDDescription tenant Id information
            type TenantIDDescription struct {
            // ID - Gets or sets Id
            ID *string `json:"id,omitempty"`
            // TenantID - Gets or sets tenantId
            TenantID *string `json:"tenantId,omitempty"`
            }

            // TenantListResult tenant Ids information.
            type TenantListResult struct {
            autorest.Response `json:"-"`
            // Value - Gets or sets tenant Ids.
            Value *[]TenantIDDescription `json:"value,omitempty"`
            // NextLink - Gets or sets the URL to get the next set of results.
            NextLink *string `json:"nextLink,omitempty"`
            }

            // TenantListResultIterator provides access to a complete listing of TenantIDDescription values.
            type TenantListResultIterator struct {
                i int
                page TenantListResultPage
            }
        // NextWithContext advances to the next value.  If there was an error making
        // the request the iterator does not advance and the error is returned.
        func (iter * TenantListResultIterator) NextWithContext(ctx context.Context) (err error) {
        if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/TenantListResultIterator.NextWithContext")
        defer func() {
        sc := -1
        if iter.Response().Response.Response != nil {
        sc = iter.Response().Response.Response.StatusCode
        }
        tracing.EndSpan(ctx, sc, err)
        }()
        }
        iter.i++
        if iter.i < len(iter. page.Values()) {
        return nil
        }
        err = iter.page.NextWithContext(ctx)
        if err != nil {
        iter. i--
        return err
        }
        iter.i = 0
        return nil
        }
        // Next advances to the next value.  If there was an error making
        // the request the iterator does not advance and the error is returned.
        // Deprecated: Use NextWithContext() instead.
        func (iter * TenantListResultIterator) Next() error {
        return iter.NextWithContext(context.Background())
        }
        // NotDone returns true if the enumeration should be started or is not yet complete.
        func (iter TenantListResultIterator) NotDone() bool {
        return iter.page.NotDone() && iter.i < len(iter. page.Values())
        }
        // Response returns the raw server response from the last page request.
        func (iter TenantListResultIterator) Response() TenantListResult {
        return iter.page.Response()
        }
        // Value returns the current value or a zero-initialized value if the
        // iterator has advanced beyond the end of the collection.
        func (iter TenantListResultIterator) Value() TenantIDDescription {
        if !iter.page.NotDone() {
        return TenantIDDescription{}
        }
        return iter.page.Values()[iter.i]
        }
        // Creates a new instance of the TenantListResultIterator type.
        func NewTenantListResultIterator (page TenantListResultPage) TenantListResultIterator {
            return TenantListResultIterator{page: page}
        }


                // IsEmpty returns true if the ListResult contains no values.
                func (tlr TenantListResult) IsEmpty() bool {
                return tlr.Value == nil || len(*tlr.Value) == 0
                }

                    // tenantListResultPreparer prepares a request to retrieve the next set of results.
                    // It returns nil if no more results exist.
                    func (tlr TenantListResult) tenantListResultPreparer(ctx context.Context) (*http.Request, error) {
                    if tlr.NextLink == nil || len(to.String(tlr.NextLink)) < 1 {
                    return nil, nil
                    }
                    return autorest.Prepare((&http.Request{}).WithContext(ctx),
                    autorest.AsJSON(),
                    autorest.AsGet(),
                    autorest.WithBaseURL(to.String( tlr.NextLink)));
                    }

            // TenantListResultPage contains a page of TenantIDDescription values.
            type TenantListResultPage struct {
                fn func(context.Context, TenantListResult) (TenantListResult, error)
                tlr TenantListResult
            }

        // NextWithContext advances to the next page of values.  If there was an error making
        // the request the page does not advance and the error is returned.
        func (page * TenantListResultPage) NextWithContext(ctx context.Context) (err error) {
        if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/TenantListResultPage.NextWithContext")
        defer func() {
        sc := -1
        if page.Response().Response.Response != nil {
        sc = page.Response().Response.Response.StatusCode
        }
        tracing.EndSpan(ctx, sc, err)
        }()
        }
        next, err := page.fn(ctx, page.tlr)
        if err != nil {
        return err
        }
        page.tlr = next
        return nil
        }

        // Next advances to the next page of values.  If there was an error making
        // the request the page does not advance and the error is returned.
        // Deprecated: Use NextWithContext() instead.
        func (page * TenantListResultPage) Next() error {
        return page.NextWithContext(context.Background())
        }
        // NotDone returns true if the page enumeration should be started or is not yet complete.
        func (page TenantListResultPage) NotDone() bool {
        return !page.tlr.IsEmpty()
        }
        // Response returns the raw server response from the last page request.
        func (page TenantListResultPage) Response() TenantListResult {
        return page.tlr
        }
        // Values returns the slice of values for the current page or nil if there are no values.
        func (page TenantListResultPage) Values() []TenantIDDescription {
        if page.tlr.IsEmpty() {
        return nil
        }
        return *page.tlr.Value
        }
        // Creates a new instance of the TenantListResultPage type.
        func NewTenantListResultPage (getNextPage func(context.Context, TenantListResult) (TenantListResult, error)) TenantListResultPage {
            return TenantListResultPage{fn: getNextPage}
        }

