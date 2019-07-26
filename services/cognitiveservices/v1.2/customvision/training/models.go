package training

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
    "github.com/Azure/go-autorest/autorest/date"
    "github.com/satori/go.uuid"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v1.2/customvision/training"

        // DomainType enumerates the values for domain type.
    type DomainType string

    const (
                // Classification ...
        Classification DomainType = "Classification"
                // ObjectDetection ...
        ObjectDetection DomainType = "ObjectDetection"
            )
    // PossibleDomainTypeValues returns an array of possible values for the DomainType const type.
    func PossibleDomainTypeValues() []DomainType {
        return []DomainType{Classification,ObjectDetection}
    }

        // ExportFlavor enumerates the values for export flavor.
    type ExportFlavor string

    const (
                // Linux ...
        Linux ExportFlavor = "Linux"
                // Windows ...
        Windows ExportFlavor = "Windows"
            )
    // PossibleExportFlavorValues returns an array of possible values for the ExportFlavor const type.
    func PossibleExportFlavorValues() []ExportFlavor {
        return []ExportFlavor{Linux,Windows}
    }

        // ExportPlatform enumerates the values for export platform.
    type ExportPlatform string

    const (
                // CoreML ...
        CoreML ExportPlatform = "CoreML"
                // DockerFile ...
        DockerFile ExportPlatform = "DockerFile"
                // ONNX ...
        ONNX ExportPlatform = "ONNX"
                // TensorFlow ...
        TensorFlow ExportPlatform = "TensorFlow"
            )
    // PossibleExportPlatformValues returns an array of possible values for the ExportPlatform const type.
    func PossibleExportPlatformValues() []ExportPlatform {
        return []ExportPlatform{CoreML,DockerFile,ONNX,TensorFlow}
    }

        // ExportStatusModel enumerates the values for export status model.
    type ExportStatusModel string

    const (
                // Done ...
        Done ExportStatusModel = "Done"
                // Exporting ...
        Exporting ExportStatusModel = "Exporting"
                // Failed ...
        Failed ExportStatusModel = "Failed"
            )
    // PossibleExportStatusModelValues returns an array of possible values for the ExportStatusModel const type.
    func PossibleExportStatusModelValues() []ExportStatusModel {
        return []ExportStatusModel{Done,Exporting,Failed}
    }

        // ImageUploadStatus enumerates the values for image upload status.
    type ImageUploadStatus string

    const (
                // ErrorImageFormat ...
        ErrorImageFormat ImageUploadStatus = "ErrorImageFormat"
                // ErrorImageSize ...
        ErrorImageSize ImageUploadStatus = "ErrorImageSize"
                // ErrorLimitExceed ...
        ErrorLimitExceed ImageUploadStatus = "ErrorLimitExceed"
                // ErrorRegionLimitExceed ...
        ErrorRegionLimitExceed ImageUploadStatus = "ErrorRegionLimitExceed"
                // ErrorSource ...
        ErrorSource ImageUploadStatus = "ErrorSource"
                // ErrorStorage ...
        ErrorStorage ImageUploadStatus = "ErrorStorage"
                // ErrorTagLimitExceed ...
        ErrorTagLimitExceed ImageUploadStatus = "ErrorTagLimitExceed"
                // ErrorUnknown ...
        ErrorUnknown ImageUploadStatus = "ErrorUnknown"
                // OK ...
        OK ImageUploadStatus = "OK"
                // OKDuplicate ...
        OKDuplicate ImageUploadStatus = "OKDuplicate"
            )
    // PossibleImageUploadStatusValues returns an array of possible values for the ImageUploadStatus const type.
    func PossibleImageUploadStatusValues() []ImageUploadStatus {
        return []ImageUploadStatus{ErrorImageFormat,ErrorImageSize,ErrorLimitExceed,ErrorRegionLimitExceed,ErrorSource,ErrorStorage,ErrorTagLimitExceed,ErrorUnknown,OK,OKDuplicate}
    }

        // OrderBy enumerates the values for order by.
    type OrderBy string

    const (
                // Newest ...
        Newest OrderBy = "Newest"
                // Oldest ...
        Oldest OrderBy = "Oldest"
                // Suggested ...
        Suggested OrderBy = "Suggested"
            )
    // PossibleOrderByValues returns an array of possible values for the OrderBy const type.
    func PossibleOrderByValues() []OrderBy {
        return []OrderBy{Newest,Oldest,Suggested}
    }

            // BoundingBox ...
            type BoundingBox struct {
            Left *float64 `json:"left,omitempty"`
            Top *float64 `json:"top,omitempty"`
            Width *float64 `json:"width,omitempty"`
            Height *float64 `json:"height,omitempty"`
            }

            // Domain ...
            type Domain struct {
            autorest.Response `json:"-"`
            // ID - READ-ONLY
            ID *uuid.UUID `json:"id,omitempty"`
            // Name - READ-ONLY
            Name *string `json:"name,omitempty"`
            // Type - READ-ONLY; Possible values include: 'Classification', 'ObjectDetection'
            Type DomainType `json:"type,omitempty"`
            // Exportable - READ-ONLY
            Exportable *bool `json:"exportable,omitempty"`
            // Enabled - READ-ONLY
            Enabled *bool `json:"enabled,omitempty"`
            }

            // Export ...
            type Export struct {
            autorest.Response `json:"-"`
            // Platform - READ-ONLY; Possible values include: 'CoreML', 'TensorFlow', 'DockerFile', 'ONNX'
            Platform ExportPlatform `json:"platform,omitempty"`
            // Status - READ-ONLY; Possible values include: 'Exporting', 'Failed', 'Done'
            Status ExportStatusModel `json:"status,omitempty"`
            // DownloadURI - READ-ONLY
            DownloadURI *string `json:"downloadUri,omitempty"`
            // Flavor - READ-ONLY; Possible values include: 'Linux', 'Windows'
            Flavor ExportFlavor `json:"flavor,omitempty"`
            }

            // Image image model to be sent as JSON
            type Image struct {
            // ID - READ-ONLY
            ID *uuid.UUID `json:"id,omitempty"`
            // Created - READ-ONLY
            Created *date.Time `json:"created,omitempty"`
            // Width - READ-ONLY
            Width *int32 `json:"width,omitempty"`
            // Height - READ-ONLY
            Height *int32 `json:"height,omitempty"`
            // ImageURI - READ-ONLY
            ImageURI *string `json:"imageUri,omitempty"`
            // ThumbnailURI - READ-ONLY
            ThumbnailURI *string `json:"thumbnailUri,omitempty"`
            // Tags - READ-ONLY
            Tags *[]ImageTag `json:"tags,omitempty"`
            // Regions - READ-ONLY
            Regions *[]ImageRegion `json:"regions,omitempty"`
            }

            // ImageCreateResult ...
            type ImageCreateResult struct {
            // SourceURL - READ-ONLY
            SourceURL *string `json:"sourceUrl,omitempty"`
            // Status - READ-ONLY; Possible values include: 'OK', 'OKDuplicate', 'ErrorSource', 'ErrorImageFormat', 'ErrorImageSize', 'ErrorStorage', 'ErrorLimitExceed', 'ErrorTagLimitExceed', 'ErrorRegionLimitExceed', 'ErrorUnknown'
            Status ImageUploadStatus `json:"status,omitempty"`
            // Image - READ-ONLY
            Image *Image `json:"image,omitempty"`
            }

            // ImageCreateSummary ...
            type ImageCreateSummary struct {
            autorest.Response `json:"-"`
            // IsBatchSuccessful - READ-ONLY
            IsBatchSuccessful *bool `json:"isBatchSuccessful,omitempty"`
            // Images - READ-ONLY
            Images *[]ImageCreateResult `json:"images,omitempty"`
            }

            // ImageFileCreateBatch ...
            type ImageFileCreateBatch struct {
            Images *[]ImageFileCreateEntry `json:"images,omitempty"`
            TagIds *[]uuid.UUID `json:"tagIds,omitempty"`
            }

            // ImageFileCreateEntry ...
            type ImageFileCreateEntry struct {
            Name *string `json:"name,omitempty"`
            Contents *[]byte `json:"contents,omitempty"`
            TagIds *[]uuid.UUID `json:"tagIds,omitempty"`
            Regions *[]Region `json:"regions,omitempty"`
            }

            // ImageIDCreateBatch ...
            type ImageIDCreateBatch struct {
            Images *[]ImageIDCreateEntry `json:"images,omitempty"`
            TagIds *[]uuid.UUID `json:"tagIds,omitempty"`
            }

            // ImageIDCreateEntry ...
            type ImageIDCreateEntry struct {
            ID *uuid.UUID `json:"id,omitempty"`
            TagIds *[]uuid.UUID `json:"tagIds,omitempty"`
            Regions *[]Region `json:"regions,omitempty"`
            }

            // ImagePerformance image performance model
            type ImagePerformance struct {
            // Predictions - READ-ONLY
            Predictions *[]Prediction `json:"predictions,omitempty"`
            // ID - READ-ONLY
            ID *uuid.UUID `json:"id,omitempty"`
            // Created - READ-ONLY
            Created *date.Time `json:"created,omitempty"`
            // Width - READ-ONLY
            Width *int32 `json:"width,omitempty"`
            // Height - READ-ONLY
            Height *int32 `json:"height,omitempty"`
            // ImageURI - READ-ONLY
            ImageURI *string `json:"imageUri,omitempty"`
            // ThumbnailURI - READ-ONLY
            ThumbnailURI *string `json:"thumbnailUri,omitempty"`
            // Tags - READ-ONLY
            Tags *[]ImageTag `json:"tags,omitempty"`
            // Regions - READ-ONLY
            Regions *[]ImageRegion `json:"regions,omitempty"`
            }

            // ImagePrediction ...
            type ImagePrediction struct {
            autorest.Response `json:"-"`
            // ID - READ-ONLY
            ID *uuid.UUID `json:"id,omitempty"`
            // Project - READ-ONLY
            Project *uuid.UUID `json:"project,omitempty"`
            // Iteration - READ-ONLY
            Iteration *uuid.UUID `json:"iteration,omitempty"`
            // Created - READ-ONLY
            Created *date.Time `json:"created,omitempty"`
            // Predictions - READ-ONLY
            Predictions *[]Prediction `json:"predictions,omitempty"`
            }

            // ImageRegion ...
            type ImageRegion struct {
            // RegionID - READ-ONLY
            RegionID *uuid.UUID `json:"regionId,omitempty"`
            // TagName - READ-ONLY
            TagName *string `json:"tagName,omitempty"`
            // Created - READ-ONLY
            Created *date.Time `json:"created,omitempty"`
            TagID *uuid.UUID `json:"tagId,omitempty"`
            Left *float64 `json:"left,omitempty"`
            Top *float64 `json:"top,omitempty"`
            Width *float64 `json:"width,omitempty"`
            Height *float64 `json:"height,omitempty"`
            }

            // ImageRegionCreateBatch batch of image region information to create.
            type ImageRegionCreateBatch struct {
            Regions *[]ImageRegionCreateEntry `json:"regions,omitempty"`
            }

            // ImageRegionCreateEntry ...
            type ImageRegionCreateEntry struct {
            ImageID *uuid.UUID `json:"imageId,omitempty"`
            TagID *uuid.UUID `json:"tagId,omitempty"`
            Left *float64 `json:"left,omitempty"`
            Top *float64 `json:"top,omitempty"`
            Width *float64 `json:"width,omitempty"`
            Height *float64 `json:"height,omitempty"`
            }

            // ImageRegionCreateResult ...
            type ImageRegionCreateResult struct {
            // ImageID - READ-ONLY
            ImageID *uuid.UUID `json:"imageId,omitempty"`
            // RegionID - READ-ONLY
            RegionID *uuid.UUID `json:"regionId,omitempty"`
            // TagName - READ-ONLY
            TagName *string `json:"tagName,omitempty"`
            // Created - READ-ONLY
            Created *date.Time `json:"created,omitempty"`
            TagID *uuid.UUID `json:"tagId,omitempty"`
            Left *float64 `json:"left,omitempty"`
            Top *float64 `json:"top,omitempty"`
            Width *float64 `json:"width,omitempty"`
            Height *float64 `json:"height,omitempty"`
            }

            // ImageRegionCreateSummary ...
            type ImageRegionCreateSummary struct {
            autorest.Response `json:"-"`
            Created *[]ImageRegionCreateResult `json:"created,omitempty"`
            Duplicated *[]ImageRegionCreateEntry `json:"duplicated,omitempty"`
            Exceeded *[]ImageRegionCreateEntry `json:"exceeded,omitempty"`
            }

            // ImageRegionProposal ...
            type ImageRegionProposal struct {
            autorest.Response `json:"-"`
            // ProjectID - READ-ONLY
            ProjectID *uuid.UUID `json:"projectId,omitempty"`
            // ImageID - READ-ONLY
            ImageID *uuid.UUID `json:"imageId,omitempty"`
            // Proposals - READ-ONLY
            Proposals *[]RegionProposal `json:"proposals,omitempty"`
            }

            // ImageTag ...
            type ImageTag struct {
            // TagID - READ-ONLY
            TagID *uuid.UUID `json:"tagId,omitempty"`
            // TagName - READ-ONLY
            TagName *string `json:"tagName,omitempty"`
            // Created - READ-ONLY
            Created *date.Time `json:"created,omitempty"`
            }

            // ImageTagCreateBatch ...
            type ImageTagCreateBatch struct {
            Tags *[]ImageTagCreateEntry `json:"tags,omitempty"`
            }

            // ImageTagCreateEntry ...
            type ImageTagCreateEntry struct {
            ImageID *uuid.UUID `json:"imageId,omitempty"`
            TagID *uuid.UUID `json:"tagId,omitempty"`
            }

            // ImageTagCreateSummary ...
            type ImageTagCreateSummary struct {
            autorest.Response `json:"-"`
            Created *[]ImageTagCreateEntry `json:"created,omitempty"`
            Duplicated *[]ImageTagCreateEntry `json:"duplicated,omitempty"`
            Exceeded *[]ImageTagCreateEntry `json:"exceeded,omitempty"`
            }

            // ImageURL ...
            type ImageURL struct {
            URL *string `json:"url,omitempty"`
            }

            // ImageURLCreateBatch ...
            type ImageURLCreateBatch struct {
            Images *[]ImageURLCreateEntry `json:"images,omitempty"`
            TagIds *[]uuid.UUID `json:"tagIds,omitempty"`
            }

            // ImageURLCreateEntry ...
            type ImageURLCreateEntry struct {
            URL *string `json:"url,omitempty"`
            TagIds *[]uuid.UUID `json:"tagIds,omitempty"`
            Regions *[]Region `json:"regions,omitempty"`
            }

            // Int32 ...
            type Int32 struct {
            autorest.Response `json:"-"`
            Value *int32 `json:"value,omitempty"`
            }

            // Iteration iteration model to be sent over JSON
            type Iteration struct {
            autorest.Response `json:"-"`
            // ID - READ-ONLY; Gets the id of the iteration
            ID *uuid.UUID `json:"id,omitempty"`
            // Name - Gets or sets the name of the iteration
            Name *string `json:"name,omitempty"`
            // IsDefault - Gets or sets a value indicating whether the iteration is the default iteration for the project
            IsDefault *bool `json:"isDefault,omitempty"`
            // Status - READ-ONLY; Gets the current iteration status
            Status *string `json:"status,omitempty"`
            // Created - READ-ONLY; Gets the time this iteration was completed
            Created *date.Time `json:"created,omitempty"`
            // LastModified - READ-ONLY; Gets the time this iteration was last modified
            LastModified *date.Time `json:"lastModified,omitempty"`
            // TrainedAt - READ-ONLY; Gets the time this iteration was last modified
            TrainedAt *date.Time `json:"trainedAt,omitempty"`
            // ProjectID - READ-ONLY; Gets the project id of the iteration
            ProjectID *uuid.UUID `json:"projectId,omitempty"`
            // Exportable - READ-ONLY; Whether the iteration can be exported to another format for download
            Exportable *bool `json:"exportable,omitempty"`
            // DomainID - READ-ONLY; Get or sets a guid of the domain the iteration has been trained on
            DomainID *uuid.UUID `json:"domainId,omitempty"`
            }

            // IterationPerformance represents the detailed performance data for a trained iteration
            type IterationPerformance struct {
            autorest.Response `json:"-"`
            // PerTagPerformance - READ-ONLY; Gets the per-tag performance details for this iteration
            PerTagPerformance *[]TagPerformance `json:"perTagPerformance,omitempty"`
            // Precision - READ-ONLY; Gets the precision
            Precision *float64 `json:"precision,omitempty"`
            // PrecisionStdDeviation - READ-ONLY; Gets the standard deviation for the precision
            PrecisionStdDeviation *float64 `json:"precisionStdDeviation,omitempty"`
            // Recall - READ-ONLY; Gets the recall
            Recall *float64 `json:"recall,omitempty"`
            // RecallStdDeviation - READ-ONLY; Gets the standard deviation for the recall
            RecallStdDeviation *float64 `json:"recallStdDeviation,omitempty"`
            // AveragePrecision - READ-ONLY; Gets the average precision when applicable
            AveragePrecision *float64 `json:"averagePrecision,omitempty"`
            }

            // ListDomain ...
            type ListDomain struct {
            autorest.Response `json:"-"`
            Value *[]Domain `json:"value,omitempty"`
            }

            // ListExport ...
            type ListExport struct {
            autorest.Response `json:"-"`
            Value *[]Export `json:"value,omitempty"`
            }

            // ListImage ...
            type ListImage struct {
            autorest.Response `json:"-"`
            Value *[]Image `json:"value,omitempty"`
            }

            // ListImagePerformance ...
            type ListImagePerformance struct {
            autorest.Response `json:"-"`
            Value *[]ImagePerformance `json:"value,omitempty"`
            }

            // ListIteration ...
            type ListIteration struct {
            autorest.Response `json:"-"`
            Value *[]Iteration `json:"value,omitempty"`
            }

            // ListProject ...
            type ListProject struct {
            autorest.Response `json:"-"`
            Value *[]Project `json:"value,omitempty"`
            }

            // ListTag ...
            type ListTag struct {
            autorest.Response `json:"-"`
            Value *[]Tag `json:"value,omitempty"`
            }

            // Prediction ...
            type Prediction struct {
            // Probability - READ-ONLY
            Probability *float64 `json:"probability,omitempty"`
            // TagID - READ-ONLY
            TagID *uuid.UUID `json:"tagId,omitempty"`
            // TagName - READ-ONLY
            TagName *string `json:"tagName,omitempty"`
            // BoundingBox - READ-ONLY
            BoundingBox *BoundingBox `json:"boundingBox,omitempty"`
            }

            // PredictionQueryResult ...
            type PredictionQueryResult struct {
            autorest.Response `json:"-"`
            // Token - READ-ONLY
            Token *PredictionQueryToken `json:"token,omitempty"`
            // Results - READ-ONLY
            Results *[]StoredImagePrediction `json:"results,omitempty"`
            }

            // PredictionQueryTag ...
            type PredictionQueryTag struct {
            // ID - READ-ONLY
            ID *uuid.UUID `json:"id,omitempty"`
            // MinThreshold - READ-ONLY
            MinThreshold *float64 `json:"minThreshold,omitempty"`
            // MaxThreshold - READ-ONLY
            MaxThreshold *float64 `json:"maxThreshold,omitempty"`
            }

            // PredictionQueryToken ...
            type PredictionQueryToken struct {
            Session *string `json:"session,omitempty"`
            Continuation *string `json:"continuation,omitempty"`
            MaxCount *int32 `json:"maxCount,omitempty"`
            // OrderBy - Possible values include: 'Newest', 'Oldest', 'Suggested'
            OrderBy OrderBy `json:"orderBy,omitempty"`
            Tags *[]PredictionQueryTag `json:"tags,omitempty"`
            IterationID *uuid.UUID `json:"iterationId,omitempty"`
            StartTime *date.Time `json:"startTime,omitempty"`
            EndTime *date.Time `json:"endTime,omitempty"`
            Application *string `json:"application,omitempty"`
            }

            // Project represents a project
            type Project struct {
            autorest.Response `json:"-"`
            // ID - READ-ONLY; Gets the project id
            ID *uuid.UUID `json:"id,omitempty"`
            // Name - Gets or sets the name of the project
            Name *string `json:"name,omitempty"`
            // Description - Gets or sets the description of the project
            Description *string `json:"description,omitempty"`
            // Settings - Gets or sets the project settings
            Settings *ProjectSettings `json:"settings,omitempty"`
            // Created - READ-ONLY; Gets the date this project was created
            Created *date.Time `json:"created,omitempty"`
            // LastModified - READ-ONLY; Gets the date this project was last modified
            LastModified *date.Time `json:"lastModified,omitempty"`
            // ThumbnailURI - READ-ONLY; Gets the thumbnail url representing the project
            ThumbnailURI *string `json:"thumbnailUri,omitempty"`
            }

            // ProjectSettings represents settings associated with a project
            type ProjectSettings struct {
            // DomainID - Gets or sets the id of the Domain to use with this project
            DomainID *uuid.UUID `json:"domainId,omitempty"`
            }

            // Region ...
            type Region struct {
            TagID *uuid.UUID `json:"tagId,omitempty"`
            Left *float64 `json:"left,omitempty"`
            Top *float64 `json:"top,omitempty"`
            Width *float64 `json:"width,omitempty"`
            Height *float64 `json:"height,omitempty"`
            }

            // RegionProposal ...
            type RegionProposal struct {
            // Confidence - READ-ONLY
            Confidence *float64 `json:"confidence,omitempty"`
            // BoundingBox - READ-ONLY
            BoundingBox *BoundingBox `json:"boundingBox,omitempty"`
            }

            // StoredImagePrediction result of an image classification request
            type StoredImagePrediction struct {
            // ImageURI - READ-ONLY
            ImageURI *string `json:"imageUri,omitempty"`
            // ThumbnailURI - READ-ONLY
            ThumbnailURI *string `json:"thumbnailUri,omitempty"`
            // Domain - READ-ONLY
            Domain *uuid.UUID `json:"domain,omitempty"`
            // ID - READ-ONLY
            ID *uuid.UUID `json:"id,omitempty"`
            // Project - READ-ONLY
            Project *uuid.UUID `json:"project,omitempty"`
            // Iteration - READ-ONLY
            Iteration *uuid.UUID `json:"iteration,omitempty"`
            // Created - READ-ONLY
            Created *date.Time `json:"created,omitempty"`
            // Predictions - READ-ONLY
            Predictions *[]Prediction `json:"predictions,omitempty"`
            }

            // Tag represents a Tag
            type Tag struct {
            autorest.Response `json:"-"`
            // ID - READ-ONLY; Gets the Tag ID
            ID *uuid.UUID `json:"id,omitempty"`
            // Name - Gets or sets the name of the tag
            Name *string `json:"name,omitempty"`
            // Description - Gets or sets the description of the tag
            Description *string `json:"description,omitempty"`
            // ImageCount - READ-ONLY; Gets the number of images with this tag
            ImageCount *int32 `json:"imageCount,omitempty"`
            }

            // TagPerformance represents performance data for a particular tag in a trained iteration
            type TagPerformance struct {
            // ID - READ-ONLY
            ID *uuid.UUID `json:"id,omitempty"`
            // Name - READ-ONLY
            Name *string `json:"name,omitempty"`
            // Precision - READ-ONLY; Gets the precision
            Precision *float64 `json:"precision,omitempty"`
            // PrecisionStdDeviation - READ-ONLY; Gets the standard deviation for the precision
            PrecisionStdDeviation *float64 `json:"precisionStdDeviation,omitempty"`
            // Recall - READ-ONLY; Gets the recall
            Recall *float64 `json:"recall,omitempty"`
            // RecallStdDeviation - READ-ONLY; Gets the standard deviation for the recall
            RecallStdDeviation *float64 `json:"recallStdDeviation,omitempty"`
            // AveragePrecision - READ-ONLY; Gets the average precision when applicable
            AveragePrecision *float64 `json:"averagePrecision,omitempty"`
            }

