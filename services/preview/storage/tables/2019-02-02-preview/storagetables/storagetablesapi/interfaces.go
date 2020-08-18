package storagetablesapi

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
	"github.com/Azure/azure-sdk-for-go/services/preview/storage/tables/2019-02-02-preview/storagetables"
	"github.com/Azure/go-autorest/autorest"
)

// TableClientAPI contains the set of methods on the TableClient type.
type TableClientAPI interface {
	Create(ctx context.Context, tableProperties storagetables.TableProperties, requestID string, formatParameter storagetables.OdataMetadataFormat, responsePreference storagetables.ResponseFormat) (result storagetables.TableResponse, err error)
	Delete(ctx context.Context, table string, requestID string) (result autorest.Response, err error)
	DeleteEntity(ctx context.Context, table string, partitionKey string, rowKey string, ifMatch string, timeout *int32, requestID string, formatParameter storagetables.OdataMetadataFormat) (result autorest.Response, err error)
	GetAccessPolicy(ctx context.Context, table string, timeout *int32, requestID string) (result storagetables.ListSignedIdentifier, err error)
	InsertEntity(ctx context.Context, table string, timeout *int32, requestID string, formatParameter storagetables.OdataMetadataFormat, tableEntityProperties map[string]interface{}, responsePreference storagetables.ResponseFormat) (result storagetables.SetSetObject, err error)
	MergeEntity(ctx context.Context, table string, partitionKey string, rowKey string, timeout *int32, requestID string, formatParameter storagetables.OdataMetadataFormat, tableEntityProperties map[string]interface{}, ifMatch string) (result autorest.Response, err error)
	Query(ctx context.Context, requestID string, formatParameter storagetables.OdataMetadataFormat, top *int32, selectParameter string, filter string, nextTableName string) (result storagetables.TableQueryResponse, err error)
	QueryEntities(ctx context.Context, table string, timeout *int32, requestID string, formatParameter storagetables.OdataMetadataFormat, top *int32, selectParameter string, filter string, nextPartitionKey string, nextRowKey string) (result storagetables.TableEntityQueryResponse, err error)
	QueryEntitiesWithPartitionAndRowKey(ctx context.Context, table string, partitionKey string, rowKey string, timeout *int32, requestID string, formatParameter storagetables.OdataMetadataFormat, selectParameter string, filter string) (result storagetables.TableEntityQueryResponse, err error)
	SetAccessPolicy(ctx context.Context, table string, tableACL []storagetables.SignedIdentifier, timeout *int32, requestID string) (result autorest.Response, err error)
	UpdateEntity(ctx context.Context, table string, partitionKey string, rowKey string, timeout *int32, requestID string, formatParameter storagetables.OdataMetadataFormat, tableEntityProperties map[string]interface{}, ifMatch string) (result autorest.Response, err error)
}

var _ TableClientAPI = (*storagetables.TableClient)(nil)

// ServiceClientAPI contains the set of methods on the ServiceClient type.
type ServiceClientAPI interface {
	GetProperties(ctx context.Context, timeout *int32, requestID string) (result storagetables.TableServiceProperties, err error)
	GetStatistics(ctx context.Context, timeout *int32, requestID string) (result storagetables.TableServiceStats, err error)
	SetProperties(ctx context.Context, tableServiceProperties storagetables.TableServiceProperties, timeout *int32, requestID string) (result autorest.Response, err error)
}

var _ ServiceClientAPI = (*storagetables.ServiceClient)(nil)