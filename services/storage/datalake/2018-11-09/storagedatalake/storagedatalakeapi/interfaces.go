package storagedatalakeapi

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
    "github.com/Azure/azure-sdk-for-go/services/storage/datalake/2018-11-09/storagedatalake"
    "github.com/Azure/go-autorest/autorest"
    "io"
)

        // FilesystemClientAPI contains the set of methods on the FilesystemClient type.
        type FilesystemClientAPI interface {
            Create(ctx context.Context, filesystem string, xMsProperties string, xMsClientRequestID string, timeout *int32, xMsDate string) (result autorest.Response, err error)
            Delete(ctx context.Context, filesystem string, ifModifiedSince string, ifUnmodifiedSince string, xMsClientRequestID string, timeout *int32, xMsDate string) (result autorest.Response, err error)
            GetProperties(ctx context.Context, filesystem string, xMsClientRequestID string, timeout *int32, xMsDate string) (result autorest.Response, err error)
            List(ctx context.Context, prefix string, continuation string, maxResults *int32, xMsClientRequestID string, timeout *int32, xMsDate string) (result storagedatalake.FilesystemList, err error)
            SetProperties(ctx context.Context, filesystem string, xMsProperties string, ifModifiedSince string, ifUnmodifiedSince string, xMsClientRequestID string, timeout *int32, xMsDate string) (result autorest.Response, err error)
        }

        var _ FilesystemClientAPI = (*storagedatalake.FilesystemClient)(nil)
        // PathClientAPI contains the set of methods on the PathClient type.
        type PathClientAPI interface {
            Create(ctx context.Context, filesystem string, pathParameter string, resource storagedatalake.PathResourceType, continuation string, mode storagedatalake.PathRenameMode, cacheControl string, contentEncoding string, contentLanguage string, contentDisposition string, xMsCacheControl string, xMsContentType string, xMsContentEncoding string, xMsContentLanguage string, xMsContentDisposition string, xMsRenameSource string, xMsLeaseID string, xMsSourceLeaseID string, xMsProperties string, xMsPermissions string, xMsUmask string, ifMatch string, ifNoneMatch string, ifModifiedSince string, ifUnmodifiedSince string, xMsSourceIfMatch string, xMsSourceIfNoneMatch string, xMsSourceIfModifiedSince string, xMsSourceIfUnmodifiedSince string, xMsClientRequestID string, timeout *int32, xMsDate string) (result autorest.Response, err error)
            Delete(ctx context.Context, filesystem string, pathParameter string, recursive *bool, continuation string, xMsLeaseID string, ifMatch string, ifNoneMatch string, ifModifiedSince string, ifUnmodifiedSince string, xMsClientRequestID string, timeout *int32, xMsDate string) (result autorest.Response, err error)
            GetProperties(ctx context.Context, filesystem string, pathParameter string, action storagedatalake.PathGetPropertiesAction, upn *bool, xMsLeaseID string, ifMatch string, ifNoneMatch string, ifModifiedSince string, ifUnmodifiedSince string, xMsClientRequestID string, timeout *int32, xMsDate string) (result autorest.Response, err error)
            Lease(ctx context.Context, xMsLeaseAction storagedatalake.PathLeaseAction, filesystem string, pathParameter string, xMsLeaseDuration *int32, xMsLeaseBreakPeriod *int32, xMsLeaseID string, xMsProposedLeaseID string, ifMatch string, ifNoneMatch string, ifModifiedSince string, ifUnmodifiedSince string, xMsClientRequestID string, timeout *int32, xMsDate string) (result autorest.Response, err error)
            List(ctx context.Context, recursive bool, filesystem string, directory string, continuation string, maxResults *int32, upn *bool, xMsClientRequestID string, timeout *int32, xMsDate string) (result storagedatalake.PathList, err error)
            Read(ctx context.Context, filesystem string, pathParameter string, rangeParameter string, xMsLeaseID string, xMsRangeGetContentMd5 *bool, ifMatch string, ifNoneMatch string, ifModifiedSince string, ifUnmodifiedSince string, xMsClientRequestID string, timeout *int32, xMsDate string) (result storagedatalake.ReadCloser, err error)
            Update(ctx context.Context, action storagedatalake.PathUpdateAction, filesystem string, pathParameter string, position *int64, retainUncommittedData *bool, closeParameter *bool, contentLength *int64, contentMD5 string, xMsLeaseID string, xMsCacheControl string, xMsContentType string, xMsContentDisposition string, xMsContentEncoding string, xMsContentLanguage string, xMsContentMd5 string, xMsProperties string, xMsOwner string, xMsGroup string, xMsPermissions string, xMsACL string, ifMatch string, ifNoneMatch string, ifModifiedSince string, ifUnmodifiedSince string, requestBody io.ReadCloser, xMsClientRequestID string, timeout *int32, xMsDate string) (result autorest.Response, err error)
        }

        var _ PathClientAPI = (*storagedatalake.PathClient)(nil)
