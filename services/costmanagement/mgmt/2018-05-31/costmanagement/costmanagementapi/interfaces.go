package costmanagementapi

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
    "github.com/Azure/azure-sdk-for-go/services/costmanagement/mgmt/2018-05-31/costmanagement"
    "github.com/Azure/go-autorest/autorest"
)

        // BaseClientAPI contains the set of methods on the BaseClient type.
        type BaseClientAPI interface {
            QueryBillingAccount(ctx context.Context, billingAccountID string, parameters costmanagement.ReportConfigDefinition) (result costmanagement.QueryResult, err error)
            QueryResourceGroup(ctx context.Context, resourceGroupName string, parameters costmanagement.ReportConfigDefinition) (result costmanagement.QueryResult, err error)
            QuerySubscription(ctx context.Context, parameters costmanagement.ReportConfigDefinition) (result costmanagement.QueryResult, err error)
        }

        var _ BaseClientAPI = (*costmanagement.BaseClient)(nil)
        // ReportConfigClientAPI contains the set of methods on the ReportConfigClient type.
        type ReportConfigClientAPI interface {
            CreateOrUpdate(ctx context.Context, reportConfigName string, parameters costmanagement.ReportConfig) (result costmanagement.ReportConfig, err error)
            CreateOrUpdateByResourceGroupName(ctx context.Context, resourceGroupName string, reportConfigName string, parameters costmanagement.ReportConfig) (result costmanagement.ReportConfig, err error)
            Delete(ctx context.Context, reportConfigName string) (result autorest.Response, err error)
            DeleteByResourceGroupName(ctx context.Context, resourceGroupName string, reportConfigName string) (result autorest.Response, err error)
            Get(ctx context.Context, reportConfigName string) (result costmanagement.ReportConfig, err error)
            GetByResourceGroupName(ctx context.Context, resourceGroupName string, reportConfigName string) (result costmanagement.ReportConfig, err error)
            List(ctx context.Context) (result costmanagement.ReportConfigListResult, err error)
            ListByResourceGroupName(ctx context.Context, resourceGroupName string) (result costmanagement.ReportConfigListResult, err error)
        }

        var _ ReportConfigClientAPI = (*costmanagement.ReportConfigClient)(nil)
        // BillingAccountDimensionsClientAPI contains the set of methods on the BillingAccountDimensionsClient type.
        type BillingAccountDimensionsClientAPI interface {
            List(ctx context.Context, billingAccountID string, filter string, expand string, skiptoken string, top *int32) (result costmanagement.DimensionsListResult, err error)
        }

        var _ BillingAccountDimensionsClientAPI = (*costmanagement.BillingAccountDimensionsClient)(nil)
        // SubscriptionDimensionsClientAPI contains the set of methods on the SubscriptionDimensionsClient type.
        type SubscriptionDimensionsClientAPI interface {
            List(ctx context.Context, filter string, expand string, skiptoken string, top *int32) (result costmanagement.DimensionsListResult, err error)
        }

        var _ SubscriptionDimensionsClientAPI = (*costmanagement.SubscriptionDimensionsClient)(nil)
        // ResourceGroupDimensionsClientAPI contains the set of methods on the ResourceGroupDimensionsClient type.
        type ResourceGroupDimensionsClientAPI interface {
            List(ctx context.Context, resourceGroupName string, filter string, expand string, skiptoken string, top *int32) (result costmanagement.DimensionsListResult, err error)
        }

        var _ ResourceGroupDimensionsClientAPI = (*costmanagement.ResourceGroupDimensionsClient)(nil)
        // OperationsClientAPI contains the set of methods on the OperationsClient type.
        type OperationsClientAPI interface {
            List(ctx context.Context) (result costmanagement.OperationListResultPage, err error)
        }

        var _ OperationsClientAPI = (*costmanagement.OperationsClient)(nil)
