package consumptionapi

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
    "github.com/Azure/azure-sdk-for-go/services/preview/consumption/mgmt/2017-04-24-preview/consumption"
)

        // UsageDetailsClientAPI contains the set of methods on the UsageDetailsClient type.
        type UsageDetailsClientAPI interface {
            List(ctx context.Context, scope string, expand string, filter string, skiptoken string, top *int32) (result consumption.UsageDetailsListResultPage, err error)
        }

        var _ UsageDetailsClientAPI = (*consumption.UsageDetailsClient)(nil)
        // OperationsClientAPI contains the set of methods on the OperationsClient type.
        type OperationsClientAPI interface {
            List(ctx context.Context) (result consumption.OperationListResultPage, err error)
        }

        var _ OperationsClientAPI = (*consumption.OperationsClient)(nil)
