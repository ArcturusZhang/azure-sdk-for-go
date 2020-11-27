package adpapi

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
	"github.com/Azure/azure-sdk-for-go/services/preview/adp/mgmt/2019-07-01-preview"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result adp.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result adp.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*adp.OperationsClient)(nil)

// AccountsClientAPI contains the set of methods on the AccountsClient type.
type AccountsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, parameters *adp.Account) (result adp.AccountsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string) (result adp.AccountsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string) (result adp.Account, err error)
	List(ctx context.Context) (result adp.AccountListPage, err error)
	ListComplete(ctx context.Context) (result adp.AccountListIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result adp.AccountListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result adp.AccountListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, accountName string, parameters *adp.AccountPatch) (result adp.AccountsUpdateFuture, err error)
}

var _ AccountsClientAPI = (*adp.AccountsClient)(nil)

// DataPoolsClientAPI contains the set of methods on the DataPoolsClient type.
type DataPoolsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, dataPoolName string, parameters *adp.DataPool) (result adp.DataPoolsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, dataPoolName string) (result adp.DataPoolsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, dataPoolName string) (result adp.DataPool, err error)
	List(ctx context.Context, resourceGroupName string, accountName string) (result adp.DataPoolListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, accountName string) (result adp.DataPoolListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, accountName string, dataPoolName string, parameters *adp.DataPoolPatch) (result adp.DataPoolsUpdateFuture, err error)
}

var _ DataPoolsClientAPI = (*adp.DataPoolsClient)(nil)