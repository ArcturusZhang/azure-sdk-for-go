package customerinsightsapi

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
    "github.com/Azure/azure-sdk-for-go/services/customerinsights/mgmt/2017-01-01/customerinsights"
    "github.com/Azure/go-autorest/autorest"
)

        // OperationsClientAPI contains the set of methods on the OperationsClient type.
        type OperationsClientAPI interface {
            List(ctx context.Context) (result customerinsights.OperationListResultPage, err error)
        }

        var _ OperationsClientAPI = (*customerinsights.OperationsClient)(nil)
        // HubsClientAPI contains the set of methods on the HubsClient type.
        type HubsClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, hubName string, parameters customerinsights.Hub) (result customerinsights.Hub, err error)
            Delete(ctx context.Context, resourceGroupName string, hubName string) (result customerinsights.HubsDeleteFuture, err error)
            Get(ctx context.Context, resourceGroupName string, hubName string) (result customerinsights.Hub, err error)
            List(ctx context.Context) (result customerinsights.HubListResultPage, err error)
            ListByResourceGroup(ctx context.Context, resourceGroupName string) (result customerinsights.HubListResultPage, err error)
            Update(ctx context.Context, resourceGroupName string, hubName string, parameters customerinsights.Hub) (result customerinsights.Hub, err error)
        }

        var _ HubsClientAPI = (*customerinsights.HubsClient)(nil)
        // ProfilesClientAPI contains the set of methods on the ProfilesClient type.
        type ProfilesClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, hubName string, profileName string, parameters customerinsights.ProfileResourceFormat) (result customerinsights.ProfilesCreateOrUpdateFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, hubName string, profileName string, localeCode string) (result customerinsights.ProfilesDeleteFuture, err error)
            Get(ctx context.Context, resourceGroupName string, hubName string, profileName string, localeCode string) (result customerinsights.ProfileResourceFormat, err error)
            GetEnrichingKpis(ctx context.Context, resourceGroupName string, hubName string, profileName string) (result customerinsights.ListKpiDefinition, err error)
            ListByHub(ctx context.Context, resourceGroupName string, hubName string, localeCode string) (result customerinsights.ProfileListResultPage, err error)
        }

        var _ ProfilesClientAPI = (*customerinsights.ProfilesClient)(nil)
        // InteractionsClientAPI contains the set of methods on the InteractionsClient type.
        type InteractionsClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, hubName string, interactionName string, parameters customerinsights.InteractionResourceFormat) (result customerinsights.InteractionsCreateOrUpdateFuture, err error)
            Get(ctx context.Context, resourceGroupName string, hubName string, interactionName string, localeCode string) (result customerinsights.InteractionResourceFormat, err error)
            ListByHub(ctx context.Context, resourceGroupName string, hubName string, localeCode string) (result customerinsights.InteractionListResultPage, err error)
            SuggestRelationshipLinks(ctx context.Context, resourceGroupName string, hubName string, interactionName string) (result customerinsights.SuggestRelationshipLinksResponse, err error)
        }

        var _ InteractionsClientAPI = (*customerinsights.InteractionsClient)(nil)
        // RelationshipsClientAPI contains the set of methods on the RelationshipsClient type.
        type RelationshipsClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, hubName string, relationshipName string, parameters customerinsights.RelationshipResourceFormat) (result customerinsights.RelationshipsCreateOrUpdateFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, hubName string, relationshipName string) (result customerinsights.RelationshipsDeleteFuture, err error)
            Get(ctx context.Context, resourceGroupName string, hubName string, relationshipName string) (result customerinsights.RelationshipResourceFormat, err error)
            ListByHub(ctx context.Context, resourceGroupName string, hubName string) (result customerinsights.RelationshipListResultPage, err error)
        }

        var _ RelationshipsClientAPI = (*customerinsights.RelationshipsClient)(nil)
        // RelationshipLinksClientAPI contains the set of methods on the RelationshipLinksClient type.
        type RelationshipLinksClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, hubName string, relationshipLinkName string, parameters customerinsights.RelationshipLinkResourceFormat) (result customerinsights.RelationshipLinksCreateOrUpdateFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, hubName string, relationshipLinkName string) (result customerinsights.RelationshipLinksDeleteFuture, err error)
            Get(ctx context.Context, resourceGroupName string, hubName string, relationshipLinkName string) (result customerinsights.RelationshipLinkResourceFormat, err error)
            ListByHub(ctx context.Context, resourceGroupName string, hubName string) (result customerinsights.RelationshipLinkListResultPage, err error)
        }

        var _ RelationshipLinksClientAPI = (*customerinsights.RelationshipLinksClient)(nil)
        // AuthorizationPoliciesClientAPI contains the set of methods on the AuthorizationPoliciesClient type.
        type AuthorizationPoliciesClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, hubName string, authorizationPolicyName string, parameters customerinsights.AuthorizationPolicyResourceFormat) (result customerinsights.AuthorizationPolicyResourceFormat, err error)
            Get(ctx context.Context, resourceGroupName string, hubName string, authorizationPolicyName string) (result customerinsights.AuthorizationPolicyResourceFormat, err error)
            ListByHub(ctx context.Context, resourceGroupName string, hubName string) (result customerinsights.AuthorizationPolicyListResultPage, err error)
            RegeneratePrimaryKey(ctx context.Context, resourceGroupName string, hubName string, authorizationPolicyName string) (result customerinsights.AuthorizationPolicy, err error)
            RegenerateSecondaryKey(ctx context.Context, resourceGroupName string, hubName string, authorizationPolicyName string) (result customerinsights.AuthorizationPolicy, err error)
        }

        var _ AuthorizationPoliciesClientAPI = (*customerinsights.AuthorizationPoliciesClient)(nil)
        // ConnectorsClientAPI contains the set of methods on the ConnectorsClient type.
        type ConnectorsClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, hubName string, connectorName string, parameters customerinsights.ConnectorResourceFormat) (result customerinsights.ConnectorsCreateOrUpdateFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, hubName string, connectorName string) (result customerinsights.ConnectorsDeleteFuture, err error)
            Get(ctx context.Context, resourceGroupName string, hubName string, connectorName string) (result customerinsights.ConnectorResourceFormat, err error)
            ListByHub(ctx context.Context, resourceGroupName string, hubName string) (result customerinsights.ConnectorListResultPage, err error)
        }

        var _ ConnectorsClientAPI = (*customerinsights.ConnectorsClient)(nil)
        // ConnectorMappingsClientAPI contains the set of methods on the ConnectorMappingsClient type.
        type ConnectorMappingsClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, hubName string, connectorName string, mappingName string, parameters customerinsights.ConnectorMappingResourceFormat) (result customerinsights.ConnectorMappingResourceFormat, err error)
            Delete(ctx context.Context, resourceGroupName string, hubName string, connectorName string, mappingName string) (result autorest.Response, err error)
            Get(ctx context.Context, resourceGroupName string, hubName string, connectorName string, mappingName string) (result customerinsights.ConnectorMappingResourceFormat, err error)
            ListByConnector(ctx context.Context, resourceGroupName string, hubName string, connectorName string) (result customerinsights.ConnectorMappingListResultPage, err error)
        }

        var _ ConnectorMappingsClientAPI = (*customerinsights.ConnectorMappingsClient)(nil)
        // KpiClientAPI contains the set of methods on the KpiClient type.
        type KpiClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, hubName string, kpiName string, parameters customerinsights.KpiResourceFormat) (result customerinsights.KpiCreateOrUpdateFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, hubName string, kpiName string) (result customerinsights.KpiDeleteFuture, err error)
            Get(ctx context.Context, resourceGroupName string, hubName string, kpiName string) (result customerinsights.KpiResourceFormat, err error)
            ListByHub(ctx context.Context, resourceGroupName string, hubName string) (result customerinsights.KpiListResultPage, err error)
            Reprocess(ctx context.Context, resourceGroupName string, hubName string, kpiName string) (result autorest.Response, err error)
        }

        var _ KpiClientAPI = (*customerinsights.KpiClient)(nil)
        // WidgetTypesClientAPI contains the set of methods on the WidgetTypesClient type.
        type WidgetTypesClientAPI interface {
            Get(ctx context.Context, resourceGroupName string, hubName string, widgetTypeName string) (result customerinsights.WidgetTypeResourceFormat, err error)
            ListByHub(ctx context.Context, resourceGroupName string, hubName string) (result customerinsights.WidgetTypeListResultPage, err error)
        }

        var _ WidgetTypesClientAPI = (*customerinsights.WidgetTypesClient)(nil)
        // ViewsClientAPI contains the set of methods on the ViewsClient type.
        type ViewsClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, hubName string, viewName string, parameters customerinsights.ViewResourceFormat) (result customerinsights.ViewResourceFormat, err error)
            Delete(ctx context.Context, resourceGroupName string, hubName string, viewName string, userID string) (result autorest.Response, err error)
            Get(ctx context.Context, resourceGroupName string, hubName string, viewName string, userID string) (result customerinsights.ViewResourceFormat, err error)
            ListByHub(ctx context.Context, resourceGroupName string, hubName string, userID string) (result customerinsights.ViewListResultPage, err error)
        }

        var _ ViewsClientAPI = (*customerinsights.ViewsClient)(nil)
        // LinksClientAPI contains the set of methods on the LinksClient type.
        type LinksClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, hubName string, linkName string, parameters customerinsights.LinkResourceFormat) (result customerinsights.LinksCreateOrUpdateFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, hubName string, linkName string) (result autorest.Response, err error)
            Get(ctx context.Context, resourceGroupName string, hubName string, linkName string) (result customerinsights.LinkResourceFormat, err error)
            ListByHub(ctx context.Context, resourceGroupName string, hubName string) (result customerinsights.LinkListResultPage, err error)
        }

        var _ LinksClientAPI = (*customerinsights.LinksClient)(nil)
        // RolesClientAPI contains the set of methods on the RolesClient type.
        type RolesClientAPI interface {
            ListByHub(ctx context.Context, resourceGroupName string, hubName string) (result customerinsights.RoleListResultPage, err error)
        }

        var _ RolesClientAPI = (*customerinsights.RolesClient)(nil)
        // RoleAssignmentsClientAPI contains the set of methods on the RoleAssignmentsClient type.
        type RoleAssignmentsClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, hubName string, assignmentName string, parameters customerinsights.RoleAssignmentResourceFormat) (result customerinsights.RoleAssignmentsCreateOrUpdateFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, hubName string, assignmentName string) (result autorest.Response, err error)
            Get(ctx context.Context, resourceGroupName string, hubName string, assignmentName string) (result customerinsights.RoleAssignmentResourceFormat, err error)
            ListByHub(ctx context.Context, resourceGroupName string, hubName string) (result customerinsights.RoleAssignmentListResultPage, err error)
        }

        var _ RoleAssignmentsClientAPI = (*customerinsights.RoleAssignmentsClient)(nil)
        // ImagesClientAPI contains the set of methods on the ImagesClient type.
        type ImagesClientAPI interface {
            GetUploadURLForData(ctx context.Context, resourceGroupName string, hubName string, parameters customerinsights.GetImageUploadURLInput) (result customerinsights.ImageDefinition, err error)
            GetUploadURLForEntityType(ctx context.Context, resourceGroupName string, hubName string, parameters customerinsights.GetImageUploadURLInput) (result customerinsights.ImageDefinition, err error)
        }

        var _ ImagesClientAPI = (*customerinsights.ImagesClient)(nil)
