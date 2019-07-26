package sqlvirtualmachineapi

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
    "github.com/Azure/azure-sdk-for-go/services/preview/sqlvirtualmachine/mgmt/2017-03-01-preview/sqlvirtualmachine"
)

        // AvailabilityGroupListenersClientAPI contains the set of methods on the AvailabilityGroupListenersClient type.
        type AvailabilityGroupListenersClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, SQLVirtualMachineGroupName string, availabilityGroupListenerName string, parameters sqlvirtualmachine.AvailabilityGroupListener) (result sqlvirtualmachine.AvailabilityGroupListenersCreateOrUpdateFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, SQLVirtualMachineGroupName string, availabilityGroupListenerName string) (result sqlvirtualmachine.AvailabilityGroupListenersDeleteFuture, err error)
            Get(ctx context.Context, resourceGroupName string, SQLVirtualMachineGroupName string, availabilityGroupListenerName string) (result sqlvirtualmachine.AvailabilityGroupListener, err error)
            ListByGroup(ctx context.Context, resourceGroupName string, SQLVirtualMachineGroupName string) (result sqlvirtualmachine.AvailabilityGroupListenerListResultPage, err error)
        }

        var _ AvailabilityGroupListenersClientAPI = (*sqlvirtualmachine.AvailabilityGroupListenersClient)(nil)
        // OperationsClientAPI contains the set of methods on the OperationsClient type.
        type OperationsClientAPI interface {
            List(ctx context.Context) (result sqlvirtualmachine.OperationListResultPage, err error)
        }

        var _ OperationsClientAPI = (*sqlvirtualmachine.OperationsClient)(nil)
        // GroupsClientAPI contains the set of methods on the GroupsClient type.
        type GroupsClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, SQLVirtualMachineGroupName string, parameters sqlvirtualmachine.Group) (result sqlvirtualmachine.GroupsCreateOrUpdateFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, SQLVirtualMachineGroupName string) (result sqlvirtualmachine.GroupsDeleteFuture, err error)
            Get(ctx context.Context, resourceGroupName string, SQLVirtualMachineGroupName string) (result sqlvirtualmachine.Group, err error)
            List(ctx context.Context) (result sqlvirtualmachine.GroupListResultPage, err error)
            ListByResourceGroup(ctx context.Context, resourceGroupName string) (result sqlvirtualmachine.GroupListResultPage, err error)
            Update(ctx context.Context, resourceGroupName string, SQLVirtualMachineGroupName string, parameters sqlvirtualmachine.GroupUpdate) (result sqlvirtualmachine.GroupsUpdateFuture, err error)
        }

        var _ GroupsClientAPI = (*sqlvirtualmachine.GroupsClient)(nil)
        // SQLVirtualMachinesClientAPI contains the set of methods on the SQLVirtualMachinesClient type.
        type SQLVirtualMachinesClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, SQLVirtualMachineName string, parameters sqlvirtualmachine.SQLVirtualMachine) (result sqlvirtualmachine.SQLVirtualMachinesCreateOrUpdateFutureType, err error)
            Delete(ctx context.Context, resourceGroupName string, SQLVirtualMachineName string) (result sqlvirtualmachine.SQLVirtualMachinesDeleteFutureType, err error)
            Get(ctx context.Context, resourceGroupName string, SQLVirtualMachineName string, expand string) (result sqlvirtualmachine.SQLVirtualMachine, err error)
            List(ctx context.Context) (result sqlvirtualmachine.ListResultPage, err error)
            ListByResourceGroup(ctx context.Context, resourceGroupName string) (result sqlvirtualmachine.ListResultPage, err error)
            Update(ctx context.Context, resourceGroupName string, SQLVirtualMachineName string, parameters sqlvirtualmachine.Update) (result sqlvirtualmachine.SQLVirtualMachinesUpdateFutureType, err error)
        }

        var _ SQLVirtualMachinesClientAPI = (*sqlvirtualmachine.SQLVirtualMachinesClient)(nil)
