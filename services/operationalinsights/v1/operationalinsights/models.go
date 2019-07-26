package operationalinsights

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
    "github.com/Azure/go-autorest/autorest"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/operationalinsights/v1/operationalinsights"

            // Column a column in a table.
            type Column struct {
            // Name - The name of this column.
            Name *string `json:"name,omitempty"`
            // Type - The data type of this column.
            Type *string `json:"type,omitempty"`
            }

            // ErrorDetail ...
            type ErrorDetail struct {
            // Code - The error's code.
            Code *string `json:"code,omitempty"`
            // Message - A human readable error message.
            Message *string `json:"message,omitempty"`
            // Target - Indicates which property in the request is responsible for the error.
            Target *string `json:"target,omitempty"`
            // Value - Indicates which value in 'target' is responsible for the error.
            Value *string `json:"value,omitempty"`
            // Resources - Indicates resources which were responsible for the error.
            Resources *[]string `json:"resources,omitempty"`
            AdditionalProperties interface{} `json:"additionalProperties,omitempty"`
            }

            // ErrorInfo ...
            type ErrorInfo struct {
            // Code - A machine readable error code.
            Code *string `json:"code,omitempty"`
            // Message - A human readable error message.
            Message *string `json:"message,omitempty"`
            // Details - error details.
            Details *[]ErrorDetail `json:"details,omitempty"`
            // Innererror - Inner error details if they exist.
            Innererror *ErrorInfo `json:"innererror,omitempty"`
            AdditionalProperties interface{} `json:"additionalProperties,omitempty"`
            }

            // ErrorResponse contains details when the response code indicates an error.
            type ErrorResponse struct {
            // Error - The error details.
            Error *ErrorInfo `json:"error,omitempty"`
            }

            // QueryBody the Analytics query. Learn more about the [Analytics query
            // syntax](https://azure.microsoft.com/documentation/articles/app-insights-analytics-reference/)
            type QueryBody struct {
            // Query - The query to execute.
            Query *string `json:"query,omitempty"`
            // Timespan - Optional. The timespan over which to query data. This is an ISO8601 time period value.  This timespan is applied in addition to any that are specified in the query expression.
            Timespan *string `json:"timespan,omitempty"`
            // Workspaces - A list of workspaces that are included in the query.
            Workspaces *[]string `json:"workspaces,omitempty"`
            }

            // QueryResults contains the tables, columns & rows resulting from a query.
            type QueryResults struct {
            autorest.Response `json:"-"`
            // Tables - The list of tables, columns and rows.
            Tables *[]Table `json:"tables,omitempty"`
            }

            // Table contains the columns and rows for one table in a query response.
            type Table struct {
            // Name - The name of the table.
            Name *string `json:"name,omitempty"`
            // Columns - The list of columns in this table.
            Columns *[]Column `json:"columns,omitempty"`
            // Rows - The resulting rows from this query.
            Rows *[][]interface{} `json:"rows,omitempty"`
            }

