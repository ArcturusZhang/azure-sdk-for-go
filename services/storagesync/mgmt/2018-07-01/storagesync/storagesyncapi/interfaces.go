package storagesyncapi

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
    "github.com/Azure/azure-sdk-for-go/services/storagesync/mgmt/2018-07-01/storagesync"
    "github.com/Azure/go-autorest/autorest"
)

        // OperationsClientAPI contains the set of methods on the OperationsClient type.
        type OperationsClientAPI interface {
            List(ctx context.Context) (result storagesync.OperationEntityListResultPage, err error)
        }

        var _ OperationsClientAPI = (*storagesync.OperationsClient)(nil)
        // ServicesClientAPI contains the set of methods on the ServicesClient type.
        type ServicesClientAPI interface {
            CheckNameAvailability(ctx context.Context, locationName string, parameters storagesync.CheckNameAvailabilityParameters) (result storagesync.CheckNameAvailabilityResult, err error)
            Create(ctx context.Context, resourceGroupName string, storageSyncServiceName string, parameters storagesync.ServiceCreateParameters) (result storagesync.Service, err error)
            Delete(ctx context.Context, resourceGroupName string, storageSyncServiceName string) (result autorest.Response, err error)
            Get(ctx context.Context, resourceGroupName string, storageSyncServiceName string) (result storagesync.Service, err error)
            ListByResourceGroup(ctx context.Context, resourceGroupName string) (result storagesync.ServiceArray, err error)
            ListBySubscription(ctx context.Context) (result storagesync.ServiceArray, err error)
            Update(ctx context.Context, resourceGroupName string, storageSyncServiceName string, parameters *storagesync.ServiceUpdateParameters) (result storagesync.Service, err error)
        }

        var _ ServicesClientAPI = (*storagesync.ServicesClient)(nil)
        // SyncGroupsClientAPI contains the set of methods on the SyncGroupsClient type.
        type SyncGroupsClientAPI interface {
            Create(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, parameters storagesync.SyncGroupCreateParameters) (result storagesync.SyncGroup, err error)
            Delete(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string) (result autorest.Response, err error)
            Get(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string) (result storagesync.SyncGroup, err error)
            ListByStorageSyncService(ctx context.Context, resourceGroupName string, storageSyncServiceName string) (result storagesync.SyncGroupArray, err error)
        }

        var _ SyncGroupsClientAPI = (*storagesync.SyncGroupsClient)(nil)
        // CloudEndpointsClientAPI contains the set of methods on the CloudEndpointsClient type.
        type CloudEndpointsClientAPI interface {
            Create(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, cloudEndpointName string, parameters storagesync.CloudEndpointCreateParameters) (result storagesync.CloudEndpointsCreateFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, cloudEndpointName string) (result storagesync.CloudEndpointsDeleteFuture, err error)
            Get(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, cloudEndpointName string) (result storagesync.CloudEndpoint, err error)
            ListBySyncGroup(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string) (result storagesync.CloudEndpointArray, err error)
            PostBackup(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, cloudEndpointName string, parameters storagesync.BackupRequest) (result storagesync.CloudEndpointsPostBackupFuture, err error)
            PostRestore(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, cloudEndpointName string, parameters storagesync.PostRestoreRequest) (result storagesync.CloudEndpointsPostRestoreFuture, err error)
            PreBackup(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, cloudEndpointName string, parameters storagesync.BackupRequest) (result storagesync.CloudEndpointsPreBackupFuture, err error)
            PreRestore(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, cloudEndpointName string, parameters storagesync.PreRestoreRequest) (result storagesync.CloudEndpointsPreRestoreFuture, err error)
            Restoreheartbeat(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, cloudEndpointName string) (result autorest.Response, err error)
        }

        var _ CloudEndpointsClientAPI = (*storagesync.CloudEndpointsClient)(nil)
        // ServerEndpointsClientAPI contains the set of methods on the ServerEndpointsClient type.
        type ServerEndpointsClientAPI interface {
            Create(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, serverEndpointName string, parameters storagesync.ServerEndpointCreateParameters) (result storagesync.ServerEndpointsCreateFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, serverEndpointName string) (result storagesync.ServerEndpointsDeleteFuture, err error)
            Get(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, serverEndpointName string) (result storagesync.ServerEndpoint, err error)
            ListBySyncGroup(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string) (result storagesync.ServerEndpointArray, err error)
            RecallAction(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, serverEndpointName string, parameters storagesync.RecallActionParameters) (result storagesync.ServerEndpointsRecallActionFuture, err error)
            Update(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, serverEndpointName string, parameters *storagesync.ServerEndpointUpdateParameters) (result storagesync.ServerEndpointsUpdateFuture, err error)
        }

        var _ ServerEndpointsClientAPI = (*storagesync.ServerEndpointsClient)(nil)
        // RegisteredServersClientAPI contains the set of methods on the RegisteredServersClient type.
        type RegisteredServersClientAPI interface {
            Create(ctx context.Context, resourceGroupName string, storageSyncServiceName string, serverID string, parameters storagesync.RegisteredServerCreateParameters) (result storagesync.RegisteredServersCreateFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, storageSyncServiceName string, serverID string) (result storagesync.RegisteredServersDeleteFuture, err error)
            Get(ctx context.Context, resourceGroupName string, storageSyncServiceName string, serverID string) (result storagesync.RegisteredServer, err error)
            ListByStorageSyncService(ctx context.Context, resourceGroupName string, storageSyncServiceName string) (result storagesync.RegisteredServerArray, err error)
            TriggerRollover(ctx context.Context, resourceGroupName string, storageSyncServiceName string, serverID string, parameters storagesync.TriggerRolloverRequest) (result storagesync.RegisteredServersTriggerRolloverFuture, err error)
        }

        var _ RegisteredServersClientAPI = (*storagesync.RegisteredServersClient)(nil)
        // WorkflowsClientAPI contains the set of methods on the WorkflowsClient type.
        type WorkflowsClientAPI interface {
            Abort(ctx context.Context, resourceGroupName string, storageSyncServiceName string, workflowID string) (result autorest.Response, err error)
            Get(ctx context.Context, resourceGroupName string, storageSyncServiceName string, workflowID string) (result storagesync.Workflow, err error)
            ListByStorageSyncService(ctx context.Context, resourceGroupName string, storageSyncServiceName string) (result storagesync.WorkflowArray, err error)
        }

        var _ WorkflowsClientAPI = (*storagesync.WorkflowsClient)(nil)
