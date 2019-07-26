package virtualmachineimagebuilderapi

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
    "github.com/Azure/azure-sdk-for-go/services/preview/virtualmachineimagebuilder/mgmt/2018-02-01-preview/virtualmachineimagebuilder"
)

        // VirtualMachineImageTemplateClientAPI contains the set of methods on the VirtualMachineImageTemplateClient type.
        type VirtualMachineImageTemplateClientAPI interface {
            CreateOrUpdate(ctx context.Context, parameters virtualmachineimagebuilder.ImageTemplate, resourceGroupName string, imageTemplateName string) (result virtualmachineimagebuilder.VirtualMachineImageTemplateCreateOrUpdateFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, imageTemplateName string) (result virtualmachineimagebuilder.VirtualMachineImageTemplateDeleteFuture, err error)
            Get(ctx context.Context, resourceGroupName string, imageTemplateName string) (result virtualmachineimagebuilder.ImageTemplate, err error)
            GetRunOutput(ctx context.Context, resourceGroupName string, imageTemplateName string, runOutputName string) (result virtualmachineimagebuilder.RunOutput, err error)
            List(ctx context.Context) (result virtualmachineimagebuilder.ImageTemplateListResultPage, err error)
            ListByResourceGroup(ctx context.Context, resourceGroupName string) (result virtualmachineimagebuilder.ImageTemplateListResultPage, err error)
            ListRunOutputs(ctx context.Context, resourceGroupName string, imageTemplateName string) (result virtualmachineimagebuilder.RunOutputCollectionPage, err error)
            Run(ctx context.Context, resourceGroupName string, imageTemplateName string) (result virtualmachineimagebuilder.VirtualMachineImageTemplateRunFuture, err error)
            Update(ctx context.Context, parameters virtualmachineimagebuilder.ImageTemplateUpdateParameters, resourceGroupName string, imageTemplateName string) (result virtualmachineimagebuilder.ImageTemplate, err error)
        }

        var _ VirtualMachineImageTemplateClientAPI = (*virtualmachineimagebuilder.VirtualMachineImageTemplateClient)(nil)
        // OperationsClientAPI contains the set of methods on the OperationsClient type.
        type OperationsClientAPI interface {
            List(ctx context.Context) (result virtualmachineimagebuilder.OperationListResultPage, err error)
        }

        var _ OperationsClientAPI = (*virtualmachineimagebuilder.OperationsClient)(nil)
