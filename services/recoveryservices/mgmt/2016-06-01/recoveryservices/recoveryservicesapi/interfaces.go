package recoveryservicesapi

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
    "github.com/Azure/azure-sdk-for-go/services/recoveryservices/mgmt/2016-06-01/recoveryservices"
    "github.com/Azure/go-autorest/autorest"
)

        // VaultCertificatesClientAPI contains the set of methods on the VaultCertificatesClient type.
        type VaultCertificatesClientAPI interface {
            Create(ctx context.Context, resourceGroupName string, vaultName string, certificateName string, certificateRequest recoveryservices.CertificateRequest) (result recoveryservices.VaultCertificateResponse, err error)
        }

        var _ VaultCertificatesClientAPI = (*recoveryservices.VaultCertificatesClient)(nil)
        // RegisteredIdentitiesClientAPI contains the set of methods on the RegisteredIdentitiesClient type.
        type RegisteredIdentitiesClientAPI interface {
            Delete(ctx context.Context, resourceGroupName string, vaultName string, identityName string) (result autorest.Response, err error)
        }

        var _ RegisteredIdentitiesClientAPI = (*recoveryservices.RegisteredIdentitiesClient)(nil)
        // ReplicationUsagesClientAPI contains the set of methods on the ReplicationUsagesClient type.
        type ReplicationUsagesClientAPI interface {
            List(ctx context.Context, resourceGroupName string, vaultName string) (result recoveryservices.ReplicationUsageList, err error)
        }

        var _ ReplicationUsagesClientAPI = (*recoveryservices.ReplicationUsagesClient)(nil)
        // ClientAPI contains the set of methods on the Client type.
        type ClientAPI interface {
            CheckNameAvailability(ctx context.Context, resourceGroupName string, location string, input recoveryservices.CheckNameAvailabilityParameters) (result recoveryservices.CheckNameAvailabilityResultResource, err error)
        }

        var _ ClientAPI = (*recoveryservices.Client)(nil)
        // VaultsClientAPI contains the set of methods on the VaultsClient type.
        type VaultsClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, vaultName string, vault recoveryservices.Vault) (result recoveryservices.Vault, err error)
            Delete(ctx context.Context, resourceGroupName string, vaultName string) (result autorest.Response, err error)
            Get(ctx context.Context, resourceGroupName string, vaultName string) (result recoveryservices.Vault, err error)
            ListByResourceGroup(ctx context.Context, resourceGroupName string) (result recoveryservices.VaultListPage, err error)
            ListBySubscriptionID(ctx context.Context) (result recoveryservices.VaultListPage, err error)
            Update(ctx context.Context, resourceGroupName string, vaultName string, vault recoveryservices.PatchVault) (result recoveryservices.Vault, err error)
        }

        var _ VaultsClientAPI = (*recoveryservices.VaultsClient)(nil)
        // OperationsClientAPI contains the set of methods on the OperationsClient type.
        type OperationsClientAPI interface {
            List(ctx context.Context) (result recoveryservices.ClientDiscoveryResponsePage, err error)
        }

        var _ OperationsClientAPI = (*recoveryservices.OperationsClient)(nil)
        // VaultExtendedInfoClientAPI contains the set of methods on the VaultExtendedInfoClient type.
        type VaultExtendedInfoClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, vaultName string, resourceResourceExtendedInfoDetails recoveryservices.VaultExtendedInfoResource) (result recoveryservices.VaultExtendedInfoResource, err error)
            Get(ctx context.Context, resourceGroupName string, vaultName string) (result recoveryservices.VaultExtendedInfoResource, err error)
            Update(ctx context.Context, resourceGroupName string, vaultName string, resourceResourceExtendedInfoDetails recoveryservices.VaultExtendedInfoResource) (result recoveryservices.VaultExtendedInfoResource, err error)
        }

        var _ VaultExtendedInfoClientAPI = (*recoveryservices.VaultExtendedInfoClient)(nil)
        // UsagesClientAPI contains the set of methods on the UsagesClient type.
        type UsagesClientAPI interface {
            ListByVaults(ctx context.Context, resourceGroupName string, vaultName string) (result recoveryservices.VaultUsageList, err error)
        }

        var _ UsagesClientAPI = (*recoveryservices.UsagesClient)(nil)
