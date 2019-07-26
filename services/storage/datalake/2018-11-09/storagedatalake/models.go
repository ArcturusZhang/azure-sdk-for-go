package storagedatalake

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
    "io"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/storage/datalake/2018-11-09/storagedatalake"

        // PathGetPropertiesAction enumerates the values for path get properties action.
    type PathGetPropertiesAction string

    const (
                // GetAccessControl ...
        GetAccessControl PathGetPropertiesAction = "getAccessControl"
                // GetStatus ...
        GetStatus PathGetPropertiesAction = "getStatus"
            )
    // PossiblePathGetPropertiesActionValues returns an array of possible values for the PathGetPropertiesAction const type.
    func PossiblePathGetPropertiesActionValues() []PathGetPropertiesAction {
        return []PathGetPropertiesAction{GetAccessControl,GetStatus}
    }

        // PathLeaseAction enumerates the values for path lease action.
    type PathLeaseAction string

    const (
                // Acquire ...
        Acquire PathLeaseAction = "acquire"
                // Break ...
        Break PathLeaseAction = "break"
                // Change ...
        Change PathLeaseAction = "change"
                // Release ...
        Release PathLeaseAction = "release"
                // Renew ...
        Renew PathLeaseAction = "renew"
            )
    // PossiblePathLeaseActionValues returns an array of possible values for the PathLeaseAction const type.
    func PossiblePathLeaseActionValues() []PathLeaseAction {
        return []PathLeaseAction{Acquire,Break,Change,Release,Renew}
    }

        // PathRenameMode enumerates the values for path rename mode.
    type PathRenameMode string

    const (
                // Legacy ...
        Legacy PathRenameMode = "legacy"
                // Posix ...
        Posix PathRenameMode = "posix"
            )
    // PossiblePathRenameModeValues returns an array of possible values for the PathRenameMode const type.
    func PossiblePathRenameModeValues() []PathRenameMode {
        return []PathRenameMode{Legacy,Posix}
    }

        // PathResourceType enumerates the values for path resource type.
    type PathResourceType string

    const (
                // Directory ...
        Directory PathResourceType = "directory"
                // File ...
        File PathResourceType = "file"
            )
    // PossiblePathResourceTypeValues returns an array of possible values for the PathResourceType const type.
    func PossiblePathResourceTypeValues() []PathResourceType {
        return []PathResourceType{Directory,File}
    }

        // PathUpdateAction enumerates the values for path update action.
    type PathUpdateAction string

    const (
                // Append ...
        Append PathUpdateAction = "append"
                // Flush ...
        Flush PathUpdateAction = "flush"
                // SetAccessControl ...
        SetAccessControl PathUpdateAction = "setAccessControl"
                // SetProperties ...
        SetProperties PathUpdateAction = "setProperties"
            )
    // PossiblePathUpdateActionValues returns an array of possible values for the PathUpdateAction const type.
    func PossiblePathUpdateActionValues() []PathUpdateAction {
        return []PathUpdateAction{Append,Flush,SetAccessControl,SetProperties}
    }

            // DataLakeStorageError ...
            type DataLakeStorageError struct {
            // Error - The service error response object.
            Error *DataLakeStorageErrorError `json:"error,omitempty"`
            }

            // DataLakeStorageErrorError the service error response object.
            type DataLakeStorageErrorError struct {
            // Code - The service error code.
            Code *string `json:"code,omitempty"`
            // Message - The service error message.
            Message *string `json:"message,omitempty"`
            }

            // Filesystem ...
            type Filesystem struct {
            Name *string `json:"name,omitempty"`
            LastModified *string `json:"lastModified,omitempty"`
            ETag *string `json:"eTag,omitempty"`
            }

            // FilesystemList ...
            type FilesystemList struct {
            autorest.Response `json:"-"`
            Filesystems *[]Filesystem `json:"filesystems,omitempty"`
            }

            // Path ...
            type Path struct {
            Name *string `json:"name,omitempty"`
            IsDirectory *bool `json:"isDirectory,omitempty"`
            LastModified *string `json:"lastModified,omitempty"`
            ETag *string `json:"eTag,omitempty"`
            ContentLength *int64 `json:"contentLength,omitempty"`
            Owner *string `json:"owner,omitempty"`
            Group *string `json:"group,omitempty"`
            Permissions *string `json:"permissions,omitempty"`
            }

            // PathList ...
            type PathList struct {
            autorest.Response `json:"-"`
            Paths *[]Path `json:"paths,omitempty"`
            }

            // ReadCloser ...
            type ReadCloser struct {
            autorest.Response `json:"-"`
            Value *io.ReadCloser `json:"value,omitempty"`
            }

