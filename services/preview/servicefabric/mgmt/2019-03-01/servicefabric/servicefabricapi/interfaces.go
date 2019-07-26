package servicefabricapi

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
    "github.com/Azure/azure-sdk-for-go/services/preview/servicefabric/mgmt/2019-03-01/servicefabric"
    "github.com/Azure/go-autorest/autorest"
)

        // ClustersClientAPI contains the set of methods on the ClustersClient type.
        type ClustersClientAPI interface {
            Create(ctx context.Context, resourceGroupName string, clusterName string, parameters servicefabric.Cluster) (result servicefabric.ClustersCreateFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, clusterName string) (result autorest.Response, err error)
            Get(ctx context.Context, resourceGroupName string, clusterName string) (result servicefabric.Cluster, err error)
            List(ctx context.Context) (result servicefabric.ClusterListResult, err error)
            ListByResourceGroup(ctx context.Context, resourceGroupName string) (result servicefabric.ClusterListResult, err error)
            Update(ctx context.Context, resourceGroupName string, clusterName string, parameters servicefabric.ClusterUpdateParameters) (result servicefabric.ClustersUpdateFuture, err error)
        }

        var _ ClustersClientAPI = (*servicefabric.ClustersClient)(nil)
        // ClusterVersionsClientAPI contains the set of methods on the ClusterVersionsClient type.
        type ClusterVersionsClientAPI interface {
            Get(ctx context.Context, location string, clusterVersion string) (result servicefabric.ClusterCodeVersionsListResult, err error)
            GetByEnvironment(ctx context.Context, location string, environment string, clusterVersion string) (result servicefabric.ClusterCodeVersionsListResult, err error)
            List(ctx context.Context, location string) (result servicefabric.ClusterCodeVersionsListResult, err error)
            ListByEnvironment(ctx context.Context, location string, environment string) (result servicefabric.ClusterCodeVersionsListResult, err error)
        }

        var _ ClusterVersionsClientAPI = (*servicefabric.ClusterVersionsClient)(nil)
        // OperationsClientAPI contains the set of methods on the OperationsClient type.
        type OperationsClientAPI interface {
            List(ctx context.Context) (result servicefabric.OperationListResultPage, err error)
        }

        var _ OperationsClientAPI = (*servicefabric.OperationsClient)(nil)
        // ApplicationTypesClientAPI contains the set of methods on the ApplicationTypesClient type.
        type ApplicationTypesClientAPI interface {
            Create(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, parameters servicefabric.ApplicationTypeResource) (result servicefabric.ApplicationTypeResource, err error)
            Delete(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string) (result servicefabric.ApplicationTypesDeleteFuture, err error)
            Get(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string) (result servicefabric.ApplicationTypeResource, err error)
            List(ctx context.Context, resourceGroupName string, clusterName string) (result servicefabric.ApplicationTypeResourceList, err error)
        }

        var _ ApplicationTypesClientAPI = (*servicefabric.ApplicationTypesClient)(nil)
        // ApplicationTypeVersionsClientAPI contains the set of methods on the ApplicationTypeVersionsClient type.
        type ApplicationTypeVersionsClientAPI interface {
            Create(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, version string, parameters servicefabric.ApplicationTypeVersionResource) (result servicefabric.ApplicationTypeVersionsCreateFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, version string) (result servicefabric.ApplicationTypeVersionsDeleteFuture, err error)
            Get(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, version string) (result servicefabric.ApplicationTypeVersionResource, err error)
            List(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string) (result servicefabric.ApplicationTypeVersionResourceList, err error)
        }

        var _ ApplicationTypeVersionsClientAPI = (*servicefabric.ApplicationTypeVersionsClient)(nil)
        // ApplicationsClientAPI contains the set of methods on the ApplicationsClient type.
        type ApplicationsClientAPI interface {
            Create(ctx context.Context, resourceGroupName string, clusterName string, applicationName string, parameters servicefabric.ApplicationResource) (result servicefabric.ApplicationsCreateFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, clusterName string, applicationName string) (result servicefabric.ApplicationsDeleteFuture, err error)
            Get(ctx context.Context, resourceGroupName string, clusterName string, applicationName string) (result servicefabric.ApplicationResource, err error)
            List(ctx context.Context, resourceGroupName string, clusterName string) (result servicefabric.ApplicationResourceList, err error)
            Update(ctx context.Context, resourceGroupName string, clusterName string, applicationName string, parameters servicefabric.ApplicationResourceUpdate) (result servicefabric.ApplicationsUpdateFuture, err error)
        }

        var _ ApplicationsClientAPI = (*servicefabric.ApplicationsClient)(nil)
        // ServicesClientAPI contains the set of methods on the ServicesClient type.
        type ServicesClientAPI interface {
            Create(ctx context.Context, resourceGroupName string, clusterName string, applicationName string, serviceName string, parameters servicefabric.ServiceResource) (result servicefabric.ServicesCreateFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, clusterName string, applicationName string, serviceName string) (result servicefabric.ServicesDeleteFuture, err error)
            Get(ctx context.Context, resourceGroupName string, clusterName string, applicationName string, serviceName string) (result servicefabric.ServiceResource, err error)
            List(ctx context.Context, resourceGroupName string, clusterName string, applicationName string) (result servicefabric.ServiceResourceList, err error)
            Update(ctx context.Context, resourceGroupName string, clusterName string, applicationName string, serviceName string, parameters servicefabric.ServiceResourceUpdate) (result servicefabric.ServicesUpdateFuture, err error)
        }

        var _ ServicesClientAPI = (*servicefabric.ServicesClient)(nil)
