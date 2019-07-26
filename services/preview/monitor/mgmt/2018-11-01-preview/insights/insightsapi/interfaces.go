package insightsapi

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
    "github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2018-11-01-preview/insights"
    "github.com/Azure/go-autorest/autorest"
)

        // AutoscaleSettingsClientAPI contains the set of methods on the AutoscaleSettingsClient type.
        type AutoscaleSettingsClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, autoscaleSettingName string, parameters insights.AutoscaleSettingResource) (result insights.AutoscaleSettingResource, err error)
            Delete(ctx context.Context, resourceGroupName string, autoscaleSettingName string) (result autorest.Response, err error)
            Get(ctx context.Context, resourceGroupName string, autoscaleSettingName string) (result insights.AutoscaleSettingResource, err error)
            ListByResourceGroup(ctx context.Context, resourceGroupName string) (result insights.AutoscaleSettingResourceCollectionPage, err error)
            ListBySubscription(ctx context.Context) (result insights.AutoscaleSettingResourceCollectionPage, err error)
            Update(ctx context.Context, resourceGroupName string, autoscaleSettingName string, autoscaleSettingResource insights.AutoscaleSettingResourcePatch) (result insights.AutoscaleSettingResource, err error)
        }

        var _ AutoscaleSettingsClientAPI = (*insights.AutoscaleSettingsClient)(nil)
        // OperationsClientAPI contains the set of methods on the OperationsClient type.
        type OperationsClientAPI interface {
            List(ctx context.Context) (result insights.OperationListResult, err error)
        }

        var _ OperationsClientAPI = (*insights.OperationsClient)(nil)
        // AlertRuleIncidentsClientAPI contains the set of methods on the AlertRuleIncidentsClient type.
        type AlertRuleIncidentsClientAPI interface {
            Get(ctx context.Context, resourceGroupName string, ruleName string, incidentName string) (result insights.Incident, err error)
            ListByAlertRule(ctx context.Context, resourceGroupName string, ruleName string) (result insights.IncidentListResult, err error)
        }

        var _ AlertRuleIncidentsClientAPI = (*insights.AlertRuleIncidentsClient)(nil)
        // AlertRulesClientAPI contains the set of methods on the AlertRulesClient type.
        type AlertRulesClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, ruleName string, parameters insights.AlertRuleResource) (result insights.AlertRuleResource, err error)
            Delete(ctx context.Context, resourceGroupName string, ruleName string) (result autorest.Response, err error)
            Get(ctx context.Context, resourceGroupName string, ruleName string) (result insights.AlertRuleResource, err error)
            ListByResourceGroup(ctx context.Context, resourceGroupName string) (result insights.AlertRuleResourceCollection, err error)
            ListBySubscription(ctx context.Context) (result insights.AlertRuleResourceCollection, err error)
            Update(ctx context.Context, resourceGroupName string, ruleName string, alertRulesResource insights.AlertRuleResourcePatch) (result insights.AlertRuleResource, err error)
        }

        var _ AlertRulesClientAPI = (*insights.AlertRulesClient)(nil)
        // LogProfilesClientAPI contains the set of methods on the LogProfilesClient type.
        type LogProfilesClientAPI interface {
            CreateOrUpdate(ctx context.Context, logProfileName string, parameters insights.LogProfileResource) (result insights.LogProfileResource, err error)
            Delete(ctx context.Context, logProfileName string) (result autorest.Response, err error)
            Get(ctx context.Context, logProfileName string) (result insights.LogProfileResource, err error)
            List(ctx context.Context) (result insights.LogProfileCollection, err error)
            Update(ctx context.Context, logProfileName string, logProfilesResource insights.LogProfileResourcePatch) (result insights.LogProfileResource, err error)
        }

        var _ LogProfilesClientAPI = (*insights.LogProfilesClient)(nil)
        // DiagnosticSettingsClientAPI contains the set of methods on the DiagnosticSettingsClient type.
        type DiagnosticSettingsClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceURI string, parameters insights.DiagnosticSettingsResource, name string) (result insights.DiagnosticSettingsResource, err error)
            Delete(ctx context.Context, resourceURI string, name string) (result autorest.Response, err error)
            Get(ctx context.Context, resourceURI string, name string) (result insights.DiagnosticSettingsResource, err error)
            List(ctx context.Context, resourceURI string) (result insights.DiagnosticSettingsResourceCollection, err error)
        }

        var _ DiagnosticSettingsClientAPI = (*insights.DiagnosticSettingsClient)(nil)
        // DiagnosticSettingsCategoryClientAPI contains the set of methods on the DiagnosticSettingsCategoryClient type.
        type DiagnosticSettingsCategoryClientAPI interface {
            Get(ctx context.Context, resourceURI string, name string) (result insights.DiagnosticSettingsCategoryResource, err error)
            List(ctx context.Context, resourceURI string) (result insights.DiagnosticSettingsCategoryResourceCollection, err error)
        }

        var _ DiagnosticSettingsCategoryClientAPI = (*insights.DiagnosticSettingsCategoryClient)(nil)
        // ActionGroupsClientAPI contains the set of methods on the ActionGroupsClient type.
        type ActionGroupsClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, actionGroupName string, actionGroup insights.ActionGroupResource) (result insights.ActionGroupResource, err error)
            Delete(ctx context.Context, resourceGroupName string, actionGroupName string) (result autorest.Response, err error)
            EnableReceiver(ctx context.Context, resourceGroupName string, actionGroupName string, enableRequest insights.EnableRequest) (result autorest.Response, err error)
            Get(ctx context.Context, resourceGroupName string, actionGroupName string) (result insights.ActionGroupResource, err error)
            ListByResourceGroup(ctx context.Context, resourceGroupName string) (result insights.ActionGroupList, err error)
            ListBySubscriptionID(ctx context.Context) (result insights.ActionGroupList, err error)
            Update(ctx context.Context, resourceGroupName string, actionGroupName string, actionGroupPatch insights.ActionGroupPatchBody) (result insights.ActionGroupResource, err error)
        }

        var _ ActionGroupsClientAPI = (*insights.ActionGroupsClient)(nil)
        // ActivityLogAlertsClientAPI contains the set of methods on the ActivityLogAlertsClient type.
        type ActivityLogAlertsClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, activityLogAlertName string, activityLogAlert insights.ActivityLogAlertResource) (result insights.ActivityLogAlertResource, err error)
            Delete(ctx context.Context, resourceGroupName string, activityLogAlertName string) (result autorest.Response, err error)
            Get(ctx context.Context, resourceGroupName string, activityLogAlertName string) (result insights.ActivityLogAlertResource, err error)
            ListByResourceGroup(ctx context.Context, resourceGroupName string) (result insights.ActivityLogAlertList, err error)
            ListBySubscriptionID(ctx context.Context) (result insights.ActivityLogAlertList, err error)
            Update(ctx context.Context, resourceGroupName string, activityLogAlertName string, activityLogAlertPatch insights.ActivityLogAlertPatchBody) (result insights.ActivityLogAlertResource, err error)
        }

        var _ ActivityLogAlertsClientAPI = (*insights.ActivityLogAlertsClient)(nil)
        // ActivityLogsClientAPI contains the set of methods on the ActivityLogsClient type.
        type ActivityLogsClientAPI interface {
            List(ctx context.Context, filter string, selectParameter string) (result insights.EventDataCollectionPage, err error)
        }

        var _ ActivityLogsClientAPI = (*insights.ActivityLogsClient)(nil)
        // EventCategoriesClientAPI contains the set of methods on the EventCategoriesClient type.
        type EventCategoriesClientAPI interface {
            List(ctx context.Context) (result insights.EventCategoryCollection, err error)
        }

        var _ EventCategoriesClientAPI = (*insights.EventCategoriesClient)(nil)
        // TenantActivityLogsClientAPI contains the set of methods on the TenantActivityLogsClient type.
        type TenantActivityLogsClientAPI interface {
            List(ctx context.Context, filter string, selectParameter string) (result insights.EventDataCollectionPage, err error)
        }

        var _ TenantActivityLogsClientAPI = (*insights.TenantActivityLogsClient)(nil)
        // MetricDefinitionsClientAPI contains the set of methods on the MetricDefinitionsClient type.
        type MetricDefinitionsClientAPI interface {
            List(ctx context.Context, resourceURI string, metricnamespace string) (result insights.MetricDefinitionCollection, err error)
        }

        var _ MetricDefinitionsClientAPI = (*insights.MetricDefinitionsClient)(nil)
        // MetricsClientAPI contains the set of methods on the MetricsClient type.
        type MetricsClientAPI interface {
            List(ctx context.Context, resourceURI string, timespan string, interval *string, metricnames string, aggregation string, top *int32, orderby string, filter string, resultType insights.ResultType, metricnamespace string) (result insights.Response, err error)
        }

        var _ MetricsClientAPI = (*insights.MetricsClient)(nil)
        // MetricBaselineClientAPI contains the set of methods on the MetricBaselineClient type.
        type MetricBaselineClientAPI interface {
            CalculateBaseline(ctx context.Context, resourceURI string, timeSeriesInformation insights.TimeSeriesInformation) (result insights.CalculateBaselineResponse, err error)
            Get(ctx context.Context, resourceURI string, metricName string, timespan string, interval *string, aggregation string, sensitivities string, resultType insights.ResultType) (result insights.BaselineResponse, err error)
        }

        var _ MetricBaselineClientAPI = (*insights.MetricBaselineClient)(nil)
        // MetricAlertsClientAPI contains the set of methods on the MetricAlertsClient type.
        type MetricAlertsClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, ruleName string, parameters insights.MetricAlertResource) (result insights.MetricAlertResource, err error)
            Delete(ctx context.Context, resourceGroupName string, ruleName string) (result autorest.Response, err error)
            Get(ctx context.Context, resourceGroupName string, ruleName string) (result insights.MetricAlertResource, err error)
            ListByResourceGroup(ctx context.Context, resourceGroupName string) (result insights.MetricAlertResourceCollection, err error)
            ListBySubscription(ctx context.Context) (result insights.MetricAlertResourceCollection, err error)
            Update(ctx context.Context, resourceGroupName string, ruleName string, parameters insights.MetricAlertResourcePatch) (result insights.MetricAlertResource, err error)
        }

        var _ MetricAlertsClientAPI = (*insights.MetricAlertsClient)(nil)
        // MetricAlertsStatusClientAPI contains the set of methods on the MetricAlertsStatusClient type.
        type MetricAlertsStatusClientAPI interface {
            List(ctx context.Context, resourceGroupName string, ruleName string) (result insights.MetricAlertStatusCollection, err error)
            ListByName(ctx context.Context, resourceGroupName string, ruleName string, statusName string) (result insights.MetricAlertStatusCollection, err error)
        }

        var _ MetricAlertsStatusClientAPI = (*insights.MetricAlertsStatusClient)(nil)
        // ScheduledQueryRulesClientAPI contains the set of methods on the ScheduledQueryRulesClient type.
        type ScheduledQueryRulesClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, ruleName string, parameters insights.LogSearchRuleResource) (result insights.LogSearchRuleResource, err error)
            Delete(ctx context.Context, resourceGroupName string, ruleName string) (result autorest.Response, err error)
            Get(ctx context.Context, resourceGroupName string, ruleName string) (result insights.LogSearchRuleResource, err error)
            ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string) (result insights.LogSearchRuleResourceCollection, err error)
            ListBySubscription(ctx context.Context, filter string) (result insights.LogSearchRuleResourceCollection, err error)
            Update(ctx context.Context, resourceGroupName string, ruleName string, parameters insights.LogSearchRuleResourcePatch) (result insights.LogSearchRuleResource, err error)
        }

        var _ ScheduledQueryRulesClientAPI = (*insights.ScheduledQueryRulesClient)(nil)
        // MetricNamespacesClientAPI contains the set of methods on the MetricNamespacesClient type.
        type MetricNamespacesClientAPI interface {
            List(ctx context.Context, resourceURI string, startTime string) (result insights.MetricNamespaceCollection, err error)
        }

        var _ MetricNamespacesClientAPI = (*insights.MetricNamespacesClient)(nil)
        // VMInsightsClientAPI contains the set of methods on the VMInsightsClient type.
        type VMInsightsClientAPI interface {
            GetOnboardingStatus(ctx context.Context, resourceURI string) (result insights.VMInsightsOnboardingStatus, err error)
        }

        var _ VMInsightsClientAPI = (*insights.VMInsightsClient)(nil)
