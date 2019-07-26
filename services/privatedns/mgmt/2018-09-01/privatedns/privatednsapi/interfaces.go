package privatednsapi

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
    "github.com/Azure/azure-sdk-for-go/services/privatedns/mgmt/2018-09-01/privatedns"
    "github.com/Azure/go-autorest/autorest"
)

        // PrivateZonesClientAPI contains the set of methods on the PrivateZonesClient type.
        type PrivateZonesClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, privateZoneName string, parameters privatedns.PrivateZone, ifMatch string, ifNoneMatch string) (result privatedns.PrivateZonesCreateOrUpdateFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, privateZoneName string, ifMatch string) (result privatedns.PrivateZonesDeleteFuture, err error)
            Get(ctx context.Context, resourceGroupName string, privateZoneName string) (result privatedns.PrivateZone, err error)
            List(ctx context.Context, top *int32) (result privatedns.PrivateZoneListResultPage, err error)
            ListByResourceGroup(ctx context.Context, resourceGroupName string, top *int32) (result privatedns.PrivateZoneListResultPage, err error)
            Update(ctx context.Context, resourceGroupName string, privateZoneName string, parameters privatedns.PrivateZone, ifMatch string) (result privatedns.PrivateZonesUpdateFuture, err error)
        }

        var _ PrivateZonesClientAPI = (*privatedns.PrivateZonesClient)(nil)
        // VirtualNetworkLinksClientAPI contains the set of methods on the VirtualNetworkLinksClient type.
        type VirtualNetworkLinksClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, privateZoneName string, virtualNetworkLinkName string, parameters privatedns.VirtualNetworkLink, ifMatch string, ifNoneMatch string) (result privatedns.VirtualNetworkLinksCreateOrUpdateFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, privateZoneName string, virtualNetworkLinkName string, ifMatch string) (result privatedns.VirtualNetworkLinksDeleteFuture, err error)
            Get(ctx context.Context, resourceGroupName string, privateZoneName string, virtualNetworkLinkName string) (result privatedns.VirtualNetworkLink, err error)
            List(ctx context.Context, resourceGroupName string, privateZoneName string, top *int32) (result privatedns.VirtualNetworkLinkListResultPage, err error)
            Update(ctx context.Context, resourceGroupName string, privateZoneName string, virtualNetworkLinkName string, parameters privatedns.VirtualNetworkLink, ifMatch string) (result privatedns.VirtualNetworkLinksUpdateFuture, err error)
        }

        var _ VirtualNetworkLinksClientAPI = (*privatedns.VirtualNetworkLinksClient)(nil)
        // RecordSetsClientAPI contains the set of methods on the RecordSetsClient type.
        type RecordSetsClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, privateZoneName string, recordType privatedns.RecordType, relativeRecordSetName string, parameters privatedns.RecordSet, ifMatch string, ifNoneMatch string) (result privatedns.RecordSet, err error)
            Delete(ctx context.Context, resourceGroupName string, privateZoneName string, recordType privatedns.RecordType, relativeRecordSetName string, ifMatch string) (result autorest.Response, err error)
            Get(ctx context.Context, resourceGroupName string, privateZoneName string, recordType privatedns.RecordType, relativeRecordSetName string) (result privatedns.RecordSet, err error)
            List(ctx context.Context, resourceGroupName string, privateZoneName string, top *int32, recordsetnamesuffix string) (result privatedns.RecordSetListResultPage, err error)
            ListByType(ctx context.Context, resourceGroupName string, privateZoneName string, recordType privatedns.RecordType, top *int32, recordsetnamesuffix string) (result privatedns.RecordSetListResultPage, err error)
            Update(ctx context.Context, resourceGroupName string, privateZoneName string, recordType privatedns.RecordType, relativeRecordSetName string, parameters privatedns.RecordSet, ifMatch string) (result privatedns.RecordSet, err error)
        }

        var _ RecordSetsClientAPI = (*privatedns.RecordSetsClient)(nil)
