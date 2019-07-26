// Package botservice implements the Azure ARM Botservice service API version 2018-07-12.
//
// Azure Bot Service is a platform for creating smart conversational agents.
package botservice

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

const (
// DefaultBaseURI is the default URI used for the service Botservice
DefaultBaseURI = "https://management.azure.com")

// BaseClient is the base client for Botservice.
type BaseClient struct {
    autorest.Client
    BaseURI string
            SubscriptionID string
}

// New creates an instance of the BaseClient client.
func New(subscriptionID string)BaseClient {
    return NewWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWithBaseURI creates an instance of the BaseClient client.
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
    return BaseClient{
        Client: autorest.NewClientWithUserAgent(UserAgent()),
        BaseURI: baseURI,
                SubscriptionID: subscriptionID,
    }
}

