package linksapi

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
    "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2016-09-01/links"
    "github.com/Azure/go-autorest/autorest"
)

        // OperationsClientAPI contains the set of methods on the OperationsClient type.
        type OperationsClientAPI interface {
            List(ctx context.Context) (result links.OperationListResultPage, err error)
        }

        var _ OperationsClientAPI = (*links.OperationsClient)(nil)
        // ResourceLinksClientAPI contains the set of methods on the ResourceLinksClient type.
        type ResourceLinksClientAPI interface {
            CreateOrUpdate(ctx context.Context, linkID string, parameters links.ResourceLink) (result links.ResourceLink, err error)
            Delete(ctx context.Context, linkID string) (result autorest.Response, err error)
            Get(ctx context.Context, linkID string) (result links.ResourceLink, err error)
            ListAtSourceScope(ctx context.Context, scope string, filter links.Filter) (result links.ResourceLinkResultPage, err error)
            ListAtSubscription(ctx context.Context, filter string) (result links.ResourceLinkResultPage, err error)
        }

        var _ ResourceLinksClientAPI = (*links.ResourceLinksClient)(nil)
