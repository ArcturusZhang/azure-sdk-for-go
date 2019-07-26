package containerinstanceapi

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
    "github.com/Azure/azure-sdk-for-go/services/containerinstance/mgmt/2018-10-01/containerinstance"
    "github.com/Azure/go-autorest/autorest"
)

        // BaseClientAPI contains the set of methods on the BaseClient type.
        type BaseClientAPI interface {
            ListCachedImages(ctx context.Context, location string) (result containerinstance.CachedImagesListResult, err error)
            ListCapabilities(ctx context.Context, location string) (result containerinstance.CapabilitiesListResult, err error)
        }

        var _ BaseClientAPI = (*containerinstance.BaseClient)(nil)
        // ContainerGroupsClientAPI contains the set of methods on the ContainerGroupsClient type.
        type ContainerGroupsClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, containerGroupName string, containerGroup containerinstance.ContainerGroup) (result containerinstance.ContainerGroupsCreateOrUpdateFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, containerGroupName string) (result containerinstance.ContainerGroup, err error)
            Get(ctx context.Context, resourceGroupName string, containerGroupName string) (result containerinstance.ContainerGroup, err error)
            List(ctx context.Context) (result containerinstance.ContainerGroupListResultPage, err error)
            ListByResourceGroup(ctx context.Context, resourceGroupName string) (result containerinstance.ContainerGroupListResultPage, err error)
            Restart(ctx context.Context, resourceGroupName string, containerGroupName string) (result containerinstance.ContainerGroupsRestartFuture, err error)
            Start(ctx context.Context, resourceGroupName string, containerGroupName string) (result containerinstance.ContainerGroupsStartFuture, err error)
            Stop(ctx context.Context, resourceGroupName string, containerGroupName string) (result autorest.Response, err error)
            Update(ctx context.Context, resourceGroupName string, containerGroupName string, resource containerinstance.Resource) (result containerinstance.ContainerGroup, err error)
        }

        var _ ContainerGroupsClientAPI = (*containerinstance.ContainerGroupsClient)(nil)
        // OperationsClientAPI contains the set of methods on the OperationsClient type.
        type OperationsClientAPI interface {
            List(ctx context.Context) (result containerinstance.OperationListResult, err error)
        }

        var _ OperationsClientAPI = (*containerinstance.OperationsClient)(nil)
        // ContainerGroupUsageClientAPI contains the set of methods on the ContainerGroupUsageClient type.
        type ContainerGroupUsageClientAPI interface {
            List(ctx context.Context, location string) (result containerinstance.UsageListResult, err error)
        }

        var _ ContainerGroupUsageClientAPI = (*containerinstance.ContainerGroupUsageClient)(nil)
        // ContainerClientAPI contains the set of methods on the ContainerClient type.
        type ContainerClientAPI interface {
            ExecuteCommand(ctx context.Context, resourceGroupName string, containerGroupName string, containerName string, containerExecRequest containerinstance.ContainerExecRequest) (result containerinstance.ContainerExecResponse, err error)
            ListLogs(ctx context.Context, resourceGroupName string, containerGroupName string, containerName string, tail *int32) (result containerinstance.Logs, err error)
        }

        var _ ContainerClientAPI = (*containerinstance.ContainerClient)(nil)
        // ServiceAssociationLinkClientAPI contains the set of methods on the ServiceAssociationLinkClient type.
        type ServiceAssociationLinkClientAPI interface {
            Delete(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string) (result autorest.Response, err error)
        }

        var _ ServiceAssociationLinkClientAPI = (*containerinstance.ServiceAssociationLinkClient)(nil)
