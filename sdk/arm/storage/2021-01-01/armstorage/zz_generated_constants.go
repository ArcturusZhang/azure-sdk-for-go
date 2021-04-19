// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorage

// AccessTier - Required for storage accounts where kind = BlobStorage. The access tier used for billing.
type AccessTier string

const (
	AccessTierHot  AccessTier = "Hot"
	AccessTierCool AccessTier = "Cool"
)

// PossibleAccessTierValues returns the possible values for the AccessTier const type.
func PossibleAccessTierValues() []AccessTier {
	return []AccessTier{
		AccessTierHot,
		AccessTierCool,
	}
}

// ToPtr() returns a *AccessTier pointing to the current value.
func (c AccessTier) ToPtr() *AccessTier {
	return &c
}

// AccountStatus - Gets the status indicating whether the primary location of the storage account is available or unavailable.
type AccountStatus string

const (
	AccountStatusAvailable   AccountStatus = "available"
	AccountStatusUnavailable AccountStatus = "unavailable"
)

// PossibleAccountStatusValues returns the possible values for the AccountStatus const type.
func PossibleAccountStatusValues() []AccountStatus {
	return []AccountStatus{
		AccountStatusAvailable,
		AccountStatusUnavailable,
	}
}

// ToPtr() returns a *AccountStatus pointing to the current value.
func (c AccountStatus) ToPtr() *AccountStatus {
	return &c
}

type BlobInventoryPolicyName string

const (
	BlobInventoryPolicyNameDefault BlobInventoryPolicyName = "default"
)

// PossibleBlobInventoryPolicyNameValues returns the possible values for the BlobInventoryPolicyName const type.
func PossibleBlobInventoryPolicyNameValues() []BlobInventoryPolicyName {
	return []BlobInventoryPolicyName{
		BlobInventoryPolicyNameDefault,
	}
}

// ToPtr() returns a *BlobInventoryPolicyName pointing to the current value.
func (c BlobInventoryPolicyName) ToPtr() *BlobInventoryPolicyName {
	return &c
}

// BlobRestoreProgressStatus - The status of blob restore progress. Possible values are: - InProgress: Indicates that blob restore is ongoing. - Complete:
// Indicates that blob restore has been completed successfully. - Failed:
// Indicates that blob restore is failed.
type BlobRestoreProgressStatus string

const (
	BlobRestoreProgressStatusComplete   BlobRestoreProgressStatus = "Complete"
	BlobRestoreProgressStatusFailed     BlobRestoreProgressStatus = "Failed"
	BlobRestoreProgressStatusInProgress BlobRestoreProgressStatus = "InProgress"
)

// PossibleBlobRestoreProgressStatusValues returns the possible values for the BlobRestoreProgressStatus const type.
func PossibleBlobRestoreProgressStatusValues() []BlobRestoreProgressStatus {
	return []BlobRestoreProgressStatus{
		BlobRestoreProgressStatusComplete,
		BlobRestoreProgressStatusFailed,
		BlobRestoreProgressStatusInProgress,
	}
}

// ToPtr() returns a *BlobRestoreProgressStatus pointing to the current value.
func (c BlobRestoreProgressStatus) ToPtr() *BlobRestoreProgressStatus {
	return &c
}

// Bypass - Specifies whether traffic is bypassed for Logging/Metrics/AzureServices. Possible values are any combination of Logging|Metrics|AzureServices
// (For example, "Logging, Metrics"), or None to bypass none
// of those traffics.
type Bypass string

const (
	BypassAzureServices Bypass = "AzureServices"
	BypassLogging       Bypass = "Logging"
	BypassMetrics       Bypass = "Metrics"
	BypassNone          Bypass = "None"
)

// PossibleBypassValues returns the possible values for the Bypass const type.
func PossibleBypassValues() []Bypass {
	return []Bypass{
		BypassAzureServices,
		BypassLogging,
		BypassMetrics,
		BypassNone,
	}
}

// ToPtr() returns a *Bypass pointing to the current value.
func (c Bypass) ToPtr() *Bypass {
	return &c
}

type CorsRuleAllowedMethodsItem string

const (
	CorsRuleAllowedMethodsItemDELETE  CorsRuleAllowedMethodsItem = "DELETE"
	CorsRuleAllowedMethodsItemGET     CorsRuleAllowedMethodsItem = "GET"
	CorsRuleAllowedMethodsItemHEAD    CorsRuleAllowedMethodsItem = "HEAD"
	CorsRuleAllowedMethodsItemMERGE   CorsRuleAllowedMethodsItem = "MERGE"
	CorsRuleAllowedMethodsItemOPTIONS CorsRuleAllowedMethodsItem = "OPTIONS"
	CorsRuleAllowedMethodsItemPOST    CorsRuleAllowedMethodsItem = "POST"
	CorsRuleAllowedMethodsItemPUT     CorsRuleAllowedMethodsItem = "PUT"
)

// PossibleCorsRuleAllowedMethodsItemValues returns the possible values for the CorsRuleAllowedMethodsItem const type.
func PossibleCorsRuleAllowedMethodsItemValues() []CorsRuleAllowedMethodsItem {
	return []CorsRuleAllowedMethodsItem{
		CorsRuleAllowedMethodsItemDELETE,
		CorsRuleAllowedMethodsItemGET,
		CorsRuleAllowedMethodsItemHEAD,
		CorsRuleAllowedMethodsItemMERGE,
		CorsRuleAllowedMethodsItemOPTIONS,
		CorsRuleAllowedMethodsItemPOST,
		CorsRuleAllowedMethodsItemPUT,
	}
}

// ToPtr() returns a *CorsRuleAllowedMethodsItem pointing to the current value.
func (c CorsRuleAllowedMethodsItem) ToPtr() *CorsRuleAllowedMethodsItem {
	return &c
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// ToPtr() returns a *CreatedByType pointing to the current value.
func (c CreatedByType) ToPtr() *CreatedByType {
	return &c
}

// DefaultAction - Specifies the default action of allow or deny when no other rules match.
type DefaultAction string

const (
	DefaultActionAllow DefaultAction = "Allow"
	DefaultActionDeny  DefaultAction = "Deny"
)

// PossibleDefaultActionValues returns the possible values for the DefaultAction const type.
func PossibleDefaultActionValues() []DefaultAction {
	return []DefaultAction{
		DefaultActionAllow,
		DefaultActionDeny,
	}
}

// ToPtr() returns a *DefaultAction pointing to the current value.
func (c DefaultAction) ToPtr() *DefaultAction {
	return &c
}

// DirectoryServiceOptions - Indicates the directory service used.
type DirectoryServiceOptions string

const (
	DirectoryServiceOptionsAADDS DirectoryServiceOptions = "AADDS"
	DirectoryServiceOptionsAD    DirectoryServiceOptions = "AD"
	DirectoryServiceOptionsNone  DirectoryServiceOptions = "None"
)

// PossibleDirectoryServiceOptionsValues returns the possible values for the DirectoryServiceOptions const type.
func PossibleDirectoryServiceOptionsValues() []DirectoryServiceOptions {
	return []DirectoryServiceOptions{
		DirectoryServiceOptionsAADDS,
		DirectoryServiceOptionsAD,
		DirectoryServiceOptionsNone,
	}
}

// ToPtr() returns a *DirectoryServiceOptions pointing to the current value.
func (c DirectoryServiceOptions) ToPtr() *DirectoryServiceOptions {
	return &c
}

// EnabledProtocols - The authentication protocol that is used for the file share. Can only be specified when creating a share.
type EnabledProtocols string

const (
	EnabledProtocolsNFS EnabledProtocols = "NFS"
	EnabledProtocolsSMB EnabledProtocols = "SMB"
)

// PossibleEnabledProtocolsValues returns the possible values for the EnabledProtocols const type.
func PossibleEnabledProtocolsValues() []EnabledProtocols {
	return []EnabledProtocols{
		EnabledProtocolsNFS,
		EnabledProtocolsSMB,
	}
}

// ToPtr() returns a *EnabledProtocols pointing to the current value.
func (c EnabledProtocols) ToPtr() *EnabledProtocols {
	return &c
}

// EncryptionScopeSource - The provider for the encryption scope. Possible values (case-insensitive): Microsoft.Storage, Microsoft.KeyVault.
type EncryptionScopeSource string

const (
	EncryptionScopeSourceMicrosoftKeyVault EncryptionScopeSource = "Microsoft.KeyVault"
	EncryptionScopeSourceMicrosoftStorage  EncryptionScopeSource = "Microsoft.Storage"
)

// PossibleEncryptionScopeSourceValues returns the possible values for the EncryptionScopeSource const type.
func PossibleEncryptionScopeSourceValues() []EncryptionScopeSource {
	return []EncryptionScopeSource{
		EncryptionScopeSourceMicrosoftKeyVault,
		EncryptionScopeSourceMicrosoftStorage,
	}
}

// ToPtr() returns a *EncryptionScopeSource pointing to the current value.
func (c EncryptionScopeSource) ToPtr() *EncryptionScopeSource {
	return &c
}

// EncryptionScopeState - The state of the encryption scope. Possible values (case-insensitive): Enabled, Disabled.
type EncryptionScopeState string

const (
	EncryptionScopeStateDisabled EncryptionScopeState = "Disabled"
	EncryptionScopeStateEnabled  EncryptionScopeState = "Enabled"
)

// PossibleEncryptionScopeStateValues returns the possible values for the EncryptionScopeState const type.
func PossibleEncryptionScopeStateValues() []EncryptionScopeState {
	return []EncryptionScopeState{
		EncryptionScopeStateDisabled,
		EncryptionScopeStateEnabled,
	}
}

// ToPtr() returns a *EncryptionScopeState pointing to the current value.
func (c EncryptionScopeState) ToPtr() *EncryptionScopeState {
	return &c
}

// ExtendedLocationTypes - The type of extendedLocation.
type ExtendedLocationTypes string

const (
	ExtendedLocationTypesEdgeZone ExtendedLocationTypes = "EdgeZone"
)

// PossibleExtendedLocationTypesValues returns the possible values for the ExtendedLocationTypes const type.
func PossibleExtendedLocationTypesValues() []ExtendedLocationTypes {
	return []ExtendedLocationTypes{
		ExtendedLocationTypesEdgeZone,
	}
}

// ToPtr() returns a *ExtendedLocationTypes pointing to the current value.
func (c ExtendedLocationTypes) ToPtr() *ExtendedLocationTypes {
	return &c
}

// GeoReplicationStatus - The status of the secondary location. Possible values are: - Live: Indicates that the secondary location is active and operational.
// - Bootstrap: Indicates initial synchronization from the primary
// location to the secondary location is in progress.This typically occurs when replication is first enabled. - Unavailable: Indicates that the secondary
// location is temporarily unavailable.
type GeoReplicationStatus string

const (
	GeoReplicationStatusBootstrap   GeoReplicationStatus = "Bootstrap"
	GeoReplicationStatusLive        GeoReplicationStatus = "Live"
	GeoReplicationStatusUnavailable GeoReplicationStatus = "Unavailable"
)

// PossibleGeoReplicationStatusValues returns the possible values for the GeoReplicationStatus const type.
func PossibleGeoReplicationStatusValues() []GeoReplicationStatus {
	return []GeoReplicationStatus{
		GeoReplicationStatusBootstrap,
		GeoReplicationStatusLive,
		GeoReplicationStatusUnavailable,
	}
}

// ToPtr() returns a *GeoReplicationStatus pointing to the current value.
func (c GeoReplicationStatus) ToPtr() *GeoReplicationStatus {
	return &c
}

// HTTPProtocol - The protocol permitted for a request made with the account SAS.
type HTTPProtocol string

const (
	HTTPProtocolHTTPSHTTP HTTPProtocol = "https,http"
	HTTPProtocolHTTPS     HTTPProtocol = "https"
)

// PossibleHTTPProtocolValues returns the possible values for the HTTPProtocol const type.
func PossibleHTTPProtocolValues() []HTTPProtocol {
	return []HTTPProtocol{
		HTTPProtocolHTTPSHTTP,
		HTTPProtocolHTTPS,
	}
}

// ToPtr() returns a *HTTPProtocol pointing to the current value.
func (c HTTPProtocol) ToPtr() *HTTPProtocol {
	return &c
}

// IdentityType - The identity type.
type IdentityType string

const (
	IdentityTypeNone                       IdentityType = "None"
	IdentityTypeSystemAssigned             IdentityType = "SystemAssigned"
	IdentityTypeSystemAssignedUserAssigned IdentityType = "SystemAssigned,UserAssigned"
	IdentityTypeUserAssigned               IdentityType = "UserAssigned"
)

// PossibleIdentityTypeValues returns the possible values for the IdentityType const type.
func PossibleIdentityTypeValues() []IdentityType {
	return []IdentityType{
		IdentityTypeNone,
		IdentityTypeSystemAssigned,
		IdentityTypeSystemAssignedUserAssigned,
		IdentityTypeUserAssigned,
	}
}

// ToPtr() returns a *IdentityType pointing to the current value.
func (c IdentityType) ToPtr() *IdentityType {
	return &c
}

// ImmutabilityPolicyState - The ImmutabilityPolicy state of a blob container, possible values include: Locked and Unlocked.
type ImmutabilityPolicyState string

const (
	ImmutabilityPolicyStateLocked   ImmutabilityPolicyState = "Locked"
	ImmutabilityPolicyStateUnlocked ImmutabilityPolicyState = "Unlocked"
)

// PossibleImmutabilityPolicyStateValues returns the possible values for the ImmutabilityPolicyState const type.
func PossibleImmutabilityPolicyStateValues() []ImmutabilityPolicyState {
	return []ImmutabilityPolicyState{
		ImmutabilityPolicyStateLocked,
		ImmutabilityPolicyStateUnlocked,
	}
}

// ToPtr() returns a *ImmutabilityPolicyState pointing to the current value.
func (c ImmutabilityPolicyState) ToPtr() *ImmutabilityPolicyState {
	return &c
}

// ImmutabilityPolicyUpdateType - The ImmutabilityPolicy update type of a blob container, possible values include: put, lock and extend.
type ImmutabilityPolicyUpdateType string

const (
	ImmutabilityPolicyUpdateTypeExtend ImmutabilityPolicyUpdateType = "extend"
	ImmutabilityPolicyUpdateTypeLock   ImmutabilityPolicyUpdateType = "lock"
	ImmutabilityPolicyUpdateTypePut    ImmutabilityPolicyUpdateType = "put"
)

// PossibleImmutabilityPolicyUpdateTypeValues returns the possible values for the ImmutabilityPolicyUpdateType const type.
func PossibleImmutabilityPolicyUpdateTypeValues() []ImmutabilityPolicyUpdateType {
	return []ImmutabilityPolicyUpdateType{
		ImmutabilityPolicyUpdateTypeExtend,
		ImmutabilityPolicyUpdateTypeLock,
		ImmutabilityPolicyUpdateTypePut,
	}
}

// ToPtr() returns a *ImmutabilityPolicyUpdateType pointing to the current value.
func (c ImmutabilityPolicyUpdateType) ToPtr() *ImmutabilityPolicyUpdateType {
	return &c
}

// InventoryRuleType - The valid value is Inventory
type InventoryRuleType string

const (
	InventoryRuleTypeInventory InventoryRuleType = "Inventory"
)

// PossibleInventoryRuleTypeValues returns the possible values for the InventoryRuleType const type.
func PossibleInventoryRuleTypeValues() []InventoryRuleType {
	return []InventoryRuleType{
		InventoryRuleTypeInventory,
	}
}

// ToPtr() returns a *InventoryRuleType pointing to the current value.
func (c InventoryRuleType) ToPtr() *InventoryRuleType {
	return &c
}

// KeyPermission - Permissions for the key -- read-only or full permissions.
type KeyPermission string

const (
	KeyPermissionRead KeyPermission = "Read"
	KeyPermissionFull KeyPermission = "Full"
)

// PossibleKeyPermissionValues returns the possible values for the KeyPermission const type.
func PossibleKeyPermissionValues() []KeyPermission {
	return []KeyPermission{
		KeyPermissionRead,
		KeyPermissionFull,
	}
}

// ToPtr() returns a *KeyPermission pointing to the current value.
func (c KeyPermission) ToPtr() *KeyPermission {
	return &c
}

// KeySource - The encryption keySource (provider). Possible values (case-insensitive): Microsoft.Storage, Microsoft.Keyvault
type KeySource string

const (
	KeySourceMicrosoftKeyvault KeySource = "Microsoft.Keyvault"
	KeySourceMicrosoftStorage  KeySource = "Microsoft.Storage"
)

// PossibleKeySourceValues returns the possible values for the KeySource const type.
func PossibleKeySourceValues() []KeySource {
	return []KeySource{
		KeySourceMicrosoftKeyvault,
		KeySourceMicrosoftStorage,
	}
}

// ToPtr() returns a *KeySource pointing to the current value.
func (c KeySource) ToPtr() *KeySource {
	return &c
}

// KeyType - Encryption key type to be used for the encryption service. 'Account' key type implies that an account-scoped encryption key will be used. 'Service'
// key type implies that a default service key is used.
type KeyType string

const (
	KeyTypeAccount KeyType = "Account"
	KeyTypeService KeyType = "Service"
)

// PossibleKeyTypeValues returns the possible values for the KeyType const type.
func PossibleKeyTypeValues() []KeyType {
	return []KeyType{
		KeyTypeAccount,
		KeyTypeService,
	}
}

// ToPtr() returns a *KeyType pointing to the current value.
func (c KeyType) ToPtr() *KeyType {
	return &c
}

// Kind - Indicates the type of storage account.
type Kind string

const (
	KindBlobStorage      Kind = "BlobStorage"
	KindBlockBlobStorage Kind = "BlockBlobStorage"
	KindFileStorage      Kind = "FileStorage"
	KindStorage          Kind = "Storage"
	KindStorageV2        Kind = "StorageV2"
)

// PossibleKindValues returns the possible values for the Kind const type.
func PossibleKindValues() []Kind {
	return []Kind{
		KindBlobStorage,
		KindBlockBlobStorage,
		KindFileStorage,
		KindStorage,
		KindStorageV2,
	}
}

// ToPtr() returns a *Kind pointing to the current value.
func (c Kind) ToPtr() *Kind {
	return &c
}

// LargeFileSharesState - Allow large file shares if sets to Enabled. It cannot be disabled once it is enabled.
type LargeFileSharesState string

const (
	LargeFileSharesStateDisabled LargeFileSharesState = "Disabled"
	LargeFileSharesStateEnabled  LargeFileSharesState = "Enabled"
)

// PossibleLargeFileSharesStateValues returns the possible values for the LargeFileSharesState const type.
func PossibleLargeFileSharesStateValues() []LargeFileSharesState {
	return []LargeFileSharesState{
		LargeFileSharesStateDisabled,
		LargeFileSharesStateEnabled,
	}
}

// ToPtr() returns a *LargeFileSharesState pointing to the current value.
func (c LargeFileSharesState) ToPtr() *LargeFileSharesState {
	return &c
}

// LeaseContainerRequestAction - Specifies the lease action. Can be one of the available actions.
type LeaseContainerRequestAction string

const (
	LeaseContainerRequestActionAcquire LeaseContainerRequestAction = "Acquire"
	LeaseContainerRequestActionBreak   LeaseContainerRequestAction = "Break"
	LeaseContainerRequestActionChange  LeaseContainerRequestAction = "Change"
	LeaseContainerRequestActionRelease LeaseContainerRequestAction = "Release"
	LeaseContainerRequestActionRenew   LeaseContainerRequestAction = "Renew"
)

// PossibleLeaseContainerRequestActionValues returns the possible values for the LeaseContainerRequestAction const type.
func PossibleLeaseContainerRequestActionValues() []LeaseContainerRequestAction {
	return []LeaseContainerRequestAction{
		LeaseContainerRequestActionAcquire,
		LeaseContainerRequestActionBreak,
		LeaseContainerRequestActionChange,
		LeaseContainerRequestActionRelease,
		LeaseContainerRequestActionRenew,
	}
}

// ToPtr() returns a *LeaseContainerRequestAction pointing to the current value.
func (c LeaseContainerRequestAction) ToPtr() *LeaseContainerRequestAction {
	return &c
}

// LeaseDuration - Specifies whether the lease on a container is of infinite or fixed duration, only when the container is leased.
type LeaseDuration string

const (
	LeaseDurationFixed    LeaseDuration = "Fixed"
	LeaseDurationInfinite LeaseDuration = "Infinite"
)

// PossibleLeaseDurationValues returns the possible values for the LeaseDuration const type.
func PossibleLeaseDurationValues() []LeaseDuration {
	return []LeaseDuration{
		LeaseDurationFixed,
		LeaseDurationInfinite,
	}
}

// ToPtr() returns a *LeaseDuration pointing to the current value.
func (c LeaseDuration) ToPtr() *LeaseDuration {
	return &c
}

// LeaseState - Lease state of the container.
type LeaseState string

const (
	LeaseStateAvailable LeaseState = "Available"
	LeaseStateBreaking  LeaseState = "Breaking"
	LeaseStateBroken    LeaseState = "Broken"
	LeaseStateExpired   LeaseState = "Expired"
	LeaseStateLeased    LeaseState = "Leased"
)

// PossibleLeaseStateValues returns the possible values for the LeaseState const type.
func PossibleLeaseStateValues() []LeaseState {
	return []LeaseState{
		LeaseStateAvailable,
		LeaseStateBreaking,
		LeaseStateBroken,
		LeaseStateExpired,
		LeaseStateLeased,
	}
}

// ToPtr() returns a *LeaseState pointing to the current value.
func (c LeaseState) ToPtr() *LeaseState {
	return &c
}

// LeaseStatus - The lease status of the container.
type LeaseStatus string

const (
	LeaseStatusLocked   LeaseStatus = "Locked"
	LeaseStatusUnlocked LeaseStatus = "Unlocked"
)

// PossibleLeaseStatusValues returns the possible values for the LeaseStatus const type.
func PossibleLeaseStatusValues() []LeaseStatus {
	return []LeaseStatus{
		LeaseStatusLocked,
		LeaseStatusUnlocked,
	}
}

// ToPtr() returns a *LeaseStatus pointing to the current value.
func (c LeaseStatus) ToPtr() *LeaseStatus {
	return &c
}

type ListContainersInclude string

const (
	ListContainersIncludeDeleted ListContainersInclude = "deleted"
)

// PossibleListContainersIncludeValues returns the possible values for the ListContainersInclude const type.
func PossibleListContainersIncludeValues() []ListContainersInclude {
	return []ListContainersInclude{
		ListContainersIncludeDeleted,
	}
}

// ToPtr() returns a *ListContainersInclude pointing to the current value.
func (c ListContainersInclude) ToPtr() *ListContainersInclude {
	return &c
}

type ListSharesExpand string

const (
	ListSharesExpandDeleted   ListSharesExpand = "deleted"
	ListSharesExpandSnapshots ListSharesExpand = "snapshots"
)

// PossibleListSharesExpandValues returns the possible values for the ListSharesExpand const type.
func PossibleListSharesExpandValues() []ListSharesExpand {
	return []ListSharesExpand{
		ListSharesExpandDeleted,
		ListSharesExpandSnapshots,
	}
}

// ToPtr() returns a *ListSharesExpand pointing to the current value.
func (c ListSharesExpand) ToPtr() *ListSharesExpand {
	return &c
}

type ManagementPolicyName string

const (
	ManagementPolicyNameDefault ManagementPolicyName = "default"
)

// PossibleManagementPolicyNameValues returns the possible values for the ManagementPolicyName const type.
func PossibleManagementPolicyNameValues() []ManagementPolicyName {
	return []ManagementPolicyName{
		ManagementPolicyNameDefault,
	}
}

// ToPtr() returns a *ManagementPolicyName pointing to the current value.
func (c ManagementPolicyName) ToPtr() *ManagementPolicyName {
	return &c
}

// MinimumTLSVersion - Set the minimum TLS version to be permitted on requests to storage. The default interpretation is TLS 1.0 for this property.
type MinimumTLSVersion string

const (
	MinimumTLSVersionTLS10 MinimumTLSVersion = "TLS1_0"
	MinimumTLSVersionTLS11 MinimumTLSVersion = "TLS1_1"
	MinimumTLSVersionTLS12 MinimumTLSVersion = "TLS1_2"
)

// PossibleMinimumTLSVersionValues returns the possible values for the MinimumTLSVersion const type.
func PossibleMinimumTLSVersionValues() []MinimumTLSVersion {
	return []MinimumTLSVersion{
		MinimumTLSVersionTLS10,
		MinimumTLSVersionTLS11,
		MinimumTLSVersionTLS12,
	}
}

// ToPtr() returns a *MinimumTLSVersion pointing to the current value.
func (c MinimumTLSVersion) ToPtr() *MinimumTLSVersion {
	return &c
}

// Name - Name of the policy. The valid value is AccessTimeTracking. This field is currently read only
type Name string

const (
	NameAccessTimeTracking Name = "AccessTimeTracking"
)

// PossibleNameValues returns the possible values for the Name const type.
func PossibleNameValues() []Name {
	return []Name{
		NameAccessTimeTracking,
	}
}

// ToPtr() returns a *Name pointing to the current value.
func (c Name) ToPtr() *Name {
	return &c
}

// Permissions - The signed permissions for the account SAS. Possible values include: Read (r), Write (w), Delete (d), List (l), Add (a), Create (c), Update
// (u) and Process (p).
type Permissions string

const (
	PermissionsA Permissions = "a"
	PermissionsC Permissions = "c"
	PermissionsD Permissions = "d"
	PermissionsL Permissions = "l"
	PermissionsP Permissions = "p"
	PermissionsR Permissions = "r"
	PermissionsU Permissions = "u"
	PermissionsW Permissions = "w"
)

// PossiblePermissionsValues returns the possible values for the Permissions const type.
func PossiblePermissionsValues() []Permissions {
	return []Permissions{
		PermissionsA,
		PermissionsC,
		PermissionsD,
		PermissionsL,
		PermissionsP,
		PermissionsR,
		PermissionsU,
		PermissionsW,
	}
}

// ToPtr() returns a *Permissions pointing to the current value.
func (c Permissions) ToPtr() *Permissions {
	return &c
}

// PrivateEndpointConnectionProvisioningState - The current provisioning state.
type PrivateEndpointConnectionProvisioningState string

const (
	PrivateEndpointConnectionProvisioningStateCreating  PrivateEndpointConnectionProvisioningState = "Creating"
	PrivateEndpointConnectionProvisioningStateDeleting  PrivateEndpointConnectionProvisioningState = "Deleting"
	PrivateEndpointConnectionProvisioningStateFailed    PrivateEndpointConnectionProvisioningState = "Failed"
	PrivateEndpointConnectionProvisioningStateSucceeded PrivateEndpointConnectionProvisioningState = "Succeeded"
)

// PossiblePrivateEndpointConnectionProvisioningStateValues returns the possible values for the PrivateEndpointConnectionProvisioningState const type.
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return []PrivateEndpointConnectionProvisioningState{
		PrivateEndpointConnectionProvisioningStateCreating,
		PrivateEndpointConnectionProvisioningStateDeleting,
		PrivateEndpointConnectionProvisioningStateFailed,
		PrivateEndpointConnectionProvisioningStateSucceeded,
	}
}

// ToPtr() returns a *PrivateEndpointConnectionProvisioningState pointing to the current value.
func (c PrivateEndpointConnectionProvisioningState) ToPtr() *PrivateEndpointConnectionProvisioningState {
	return &c
}

// PrivateEndpointServiceConnectionStatus - The private endpoint connection status.
type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusApproved PrivateEndpointServiceConnectionStatus = "Approved"
	PrivateEndpointServiceConnectionStatusPending  PrivateEndpointServiceConnectionStatus = "Pending"
	PrivateEndpointServiceConnectionStatusRejected PrivateEndpointServiceConnectionStatus = "Rejected"
)

// PossiblePrivateEndpointServiceConnectionStatusValues returns the possible values for the PrivateEndpointServiceConnectionStatus const type.
func PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus {
	return []PrivateEndpointServiceConnectionStatus{
		PrivateEndpointServiceConnectionStatusApproved,
		PrivateEndpointServiceConnectionStatusPending,
		PrivateEndpointServiceConnectionStatusRejected,
	}
}

// ToPtr() returns a *PrivateEndpointServiceConnectionStatus pointing to the current value.
func (c PrivateEndpointServiceConnectionStatus) ToPtr() *PrivateEndpointServiceConnectionStatus {
	return &c
}

// ProvisioningState - Gets the status of the storage account at the time the operation was called.
type ProvisioningState string

const (
	ProvisioningStateCreating     ProvisioningState = "Creating"
	ProvisioningStateResolvingDNS ProvisioningState = "ResolvingDNS"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCreating,
		ProvisioningStateResolvingDNS,
		ProvisioningStateSucceeded,
	}
}

// ToPtr() returns a *ProvisioningState pointing to the current value.
func (c ProvisioningState) ToPtr() *ProvisioningState {
	return &c
}

// PublicAccess - Specifies whether data in the container may be accessed publicly and the level of access.
type PublicAccess string

const (
	PublicAccessContainer PublicAccess = "Container"
	PublicAccessBlob      PublicAccess = "Blob"
	PublicAccessNone      PublicAccess = "None"
)

// PossiblePublicAccessValues returns the possible values for the PublicAccess const type.
func PossiblePublicAccessValues() []PublicAccess {
	return []PublicAccess{
		PublicAccessContainer,
		PublicAccessBlob,
		PublicAccessNone,
	}
}

// ToPtr() returns a *PublicAccess pointing to the current value.
func (c PublicAccess) ToPtr() *PublicAccess {
	return &c
}

type PutSharesExpand string

const (
	PutSharesExpandSnapshots PutSharesExpand = "snapshots"
)

// PossiblePutSharesExpandValues returns the possible values for the PutSharesExpand const type.
func PossiblePutSharesExpandValues() []PutSharesExpand {
	return []PutSharesExpand{
		PutSharesExpandSnapshots,
	}
}

// ToPtr() returns a *PutSharesExpand pointing to the current value.
func (c PutSharesExpand) ToPtr() *PutSharesExpand {
	return &c
}

// Reason - Gets the reason that a storage account name could not be used. The Reason element is only returned if NameAvailable is false.
type Reason string

const (
	ReasonAccountNameInvalid Reason = "AccountNameInvalid"
	ReasonAlreadyExists      Reason = "AlreadyExists"
)

// PossibleReasonValues returns the possible values for the Reason const type.
func PossibleReasonValues() []Reason {
	return []Reason{
		ReasonAccountNameInvalid,
		ReasonAlreadyExists,
	}
}

// ToPtr() returns a *Reason pointing to the current value.
func (c Reason) ToPtr() *Reason {
	return &c
}

// ReasonCode - The reason for the restriction. As of now this can be "QuotaId" or "NotAvailableForSubscription". Quota Id is set when the SKU has requiredQuotas
// parameter as the subscription does not belong to that
// quota. The "NotAvailableForSubscription" is related to capacity at DC.
type ReasonCode string

const (
	ReasonCodeNotAvailableForSubscription ReasonCode = "NotAvailableForSubscription"
	ReasonCodeQuotaID                     ReasonCode = "QuotaId"
)

// PossibleReasonCodeValues returns the possible values for the ReasonCode const type.
func PossibleReasonCodeValues() []ReasonCode {
	return []ReasonCode{
		ReasonCodeNotAvailableForSubscription,
		ReasonCodeQuotaID,
	}
}

// ToPtr() returns a *ReasonCode pointing to the current value.
func (c ReasonCode) ToPtr() *ReasonCode {
	return &c
}

// RootSquashType - The property is for NFS share only. The default is NoRootSquash.
type RootSquashType string

const (
	RootSquashTypeAllSquash    RootSquashType = "AllSquash"
	RootSquashTypeNoRootSquash RootSquashType = "NoRootSquash"
	RootSquashTypeRootSquash   RootSquashType = "RootSquash"
)

// PossibleRootSquashTypeValues returns the possible values for the RootSquashType const type.
func PossibleRootSquashTypeValues() []RootSquashType {
	return []RootSquashType{
		RootSquashTypeAllSquash,
		RootSquashTypeNoRootSquash,
		RootSquashTypeRootSquash,
	}
}

// ToPtr() returns a *RootSquashType pointing to the current value.
func (c RootSquashType) ToPtr() *RootSquashType {
	return &c
}

// RoutingChoice - Routing Choice defines the kind of network routing opted by the user.
type RoutingChoice string

const (
	RoutingChoiceInternetRouting  RoutingChoice = "InternetRouting"
	RoutingChoiceMicrosoftRouting RoutingChoice = "MicrosoftRouting"
)

// PossibleRoutingChoiceValues returns the possible values for the RoutingChoice const type.
func PossibleRoutingChoiceValues() []RoutingChoice {
	return []RoutingChoice{
		RoutingChoiceInternetRouting,
		RoutingChoiceMicrosoftRouting,
	}
}

// ToPtr() returns a *RoutingChoice pointing to the current value.
func (c RoutingChoice) ToPtr() *RoutingChoice {
	return &c
}

// RuleType - The valid value is Lifecycle
type RuleType string

const (
	RuleTypeLifecycle RuleType = "Lifecycle"
)

// PossibleRuleTypeValues returns the possible values for the RuleType const type.
func PossibleRuleTypeValues() []RuleType {
	return []RuleType{
		RuleTypeLifecycle,
	}
}

// ToPtr() returns a *RuleType pointing to the current value.
func (c RuleType) ToPtr() *RuleType {
	return &c
}

// SKUName - The SKU name. Required for account creation; optional for update. Note that in older versions, SKU name was called accountType.
type SKUName string

const (
	SKUNamePremiumLRS     SKUName = "Premium_LRS"
	SKUNamePremiumZRS     SKUName = "Premium_ZRS"
	SKUNameStandardGRS    SKUName = "Standard_GRS"
	SKUNameStandardGZRS   SKUName = "Standard_GZRS"
	SKUNameStandardLRS    SKUName = "Standard_LRS"
	SKUNameStandardRAGRS  SKUName = "Standard_RAGRS"
	SKUNameStandardRAGZRS SKUName = "Standard_RAGZRS"
	SKUNameStandardZRS    SKUName = "Standard_ZRS"
)

// PossibleSKUNameValues returns the possible values for the SKUName const type.
func PossibleSKUNameValues() []SKUName {
	return []SKUName{
		SKUNamePremiumLRS,
		SKUNamePremiumZRS,
		SKUNameStandardGRS,
		SKUNameStandardGZRS,
		SKUNameStandardLRS,
		SKUNameStandardRAGRS,
		SKUNameStandardRAGZRS,
		SKUNameStandardZRS,
	}
}

// ToPtr() returns a *SKUName pointing to the current value.
func (c SKUName) ToPtr() *SKUName {
	return &c
}

// SKUTier - The SKU tier. This is based on the SKU name.
type SKUTier string

const (
	SKUTierStandard SKUTier = "Standard"
	SKUTierPremium  SKUTier = "Premium"
)

// PossibleSKUTierValues returns the possible values for the SKUTier const type.
func PossibleSKUTierValues() []SKUTier {
	return []SKUTier{
		SKUTierStandard,
		SKUTierPremium,
	}
}

// ToPtr() returns a *SKUTier pointing to the current value.
func (c SKUTier) ToPtr() *SKUTier {
	return &c
}

// Services - The signed services accessible with the account SAS. Possible values include: Blob (b), Queue (q), Table (t), File (f).
type Services string

const (
	ServicesB Services = "b"
	ServicesF Services = "f"
	ServicesQ Services = "q"
	ServicesT Services = "t"
)

// PossibleServicesValues returns the possible values for the Services const type.
func PossibleServicesValues() []Services {
	return []Services{
		ServicesB,
		ServicesF,
		ServicesQ,
		ServicesT,
	}
}

// ToPtr() returns a *Services pointing to the current value.
func (c Services) ToPtr() *Services {
	return &c
}

// ShareAccessTier - Access tier for specific share. GpV2 account can choose between TransactionOptimized (default), Hot, and Cool. FileStorage account
// can choose Premium.
type ShareAccessTier string

const (
	ShareAccessTierCool                 ShareAccessTier = "Cool"
	ShareAccessTierHot                  ShareAccessTier = "Hot"
	ShareAccessTierPremium              ShareAccessTier = "Premium"
	ShareAccessTierTransactionOptimized ShareAccessTier = "TransactionOptimized"
)

// PossibleShareAccessTierValues returns the possible values for the ShareAccessTier const type.
func PossibleShareAccessTierValues() []ShareAccessTier {
	return []ShareAccessTier{
		ShareAccessTierCool,
		ShareAccessTierHot,
		ShareAccessTierPremium,
		ShareAccessTierTransactionOptimized,
	}
}

// ToPtr() returns a *ShareAccessTier pointing to the current value.
func (c ShareAccessTier) ToPtr() *ShareAccessTier {
	return &c
}

// SignedResource - The signed services accessible with the service SAS. Possible values include: Blob (b), Container (c), File (f), Share (s).
type SignedResource string

const (
	SignedResourceB SignedResource = "b"
	SignedResourceC SignedResource = "c"
	SignedResourceF SignedResource = "f"
	SignedResourceS SignedResource = "s"
)

// PossibleSignedResourceValues returns the possible values for the SignedResource const type.
func PossibleSignedResourceValues() []SignedResource {
	return []SignedResource{
		SignedResourceB,
		SignedResourceC,
		SignedResourceF,
		SignedResourceS,
	}
}

// ToPtr() returns a *SignedResource pointing to the current value.
func (c SignedResource) ToPtr() *SignedResource {
	return &c
}

// SignedResourceTypes - The signed resource types that are accessible with the account SAS. Service (s): Access to service-level APIs; Container (c): Access
// to container-level APIs; Object (o): Access to object-level APIs
// for blobs, queue messages, table entities, and files.
type SignedResourceTypes string

const (
	SignedResourceTypesC SignedResourceTypes = "c"
	SignedResourceTypesO SignedResourceTypes = "o"
	SignedResourceTypesS SignedResourceTypes = "s"
)

// PossibleSignedResourceTypesValues returns the possible values for the SignedResourceTypes const type.
func PossibleSignedResourceTypesValues() []SignedResourceTypes {
	return []SignedResourceTypes{
		SignedResourceTypesC,
		SignedResourceTypesO,
		SignedResourceTypesS,
	}
}

// ToPtr() returns a *SignedResourceTypes pointing to the current value.
func (c SignedResourceTypes) ToPtr() *SignedResourceTypes {
	return &c
}

// State - Gets the state of virtual network rule.
type State string

const (
	StateDeprovisioning       State = "deprovisioning"
	StateFailed               State = "failed"
	StateNetworkSourceDeleted State = "networkSourceDeleted"
	StateProvisioning         State = "provisioning"
	StateSucceeded            State = "succeeded"
)

// PossibleStateValues returns the possible values for the State const type.
func PossibleStateValues() []State {
	return []State{
		StateDeprovisioning,
		StateFailed,
		StateNetworkSourceDeleted,
		StateProvisioning,
		StateSucceeded,
	}
}

// ToPtr() returns a *State pointing to the current value.
func (c State) ToPtr() *State {
	return &c
}

type StorageAccountExpand string

const (
	StorageAccountExpandGeoReplicationStats StorageAccountExpand = "geoReplicationStats"
	StorageAccountExpandBlobRestoreStatus   StorageAccountExpand = "blobRestoreStatus"
)

// PossibleStorageAccountExpandValues returns the possible values for the StorageAccountExpand const type.
func PossibleStorageAccountExpandValues() []StorageAccountExpand {
	return []StorageAccountExpand{
		StorageAccountExpandGeoReplicationStats,
		StorageAccountExpandBlobRestoreStatus,
	}
}

// ToPtr() returns a *StorageAccountExpand pointing to the current value.
func (c StorageAccountExpand) ToPtr() *StorageAccountExpand {
	return &c
}

// UsageUnit - Gets the unit of measurement.
type UsageUnit string

const (
	UsageUnitCount           UsageUnit = "Count"
	UsageUnitBytes           UsageUnit = "Bytes"
	UsageUnitSeconds         UsageUnit = "Seconds"
	UsageUnitPercent         UsageUnit = "Percent"
	UsageUnitCountsPerSecond UsageUnit = "CountsPerSecond"
	UsageUnitBytesPerSecond  UsageUnit = "BytesPerSecond"
)

// PossibleUsageUnitValues returns the possible values for the UsageUnit const type.
func PossibleUsageUnitValues() []UsageUnit {
	return []UsageUnit{
		UsageUnitCount,
		UsageUnitBytes,
		UsageUnitSeconds,
		UsageUnitPercent,
		UsageUnitCountsPerSecond,
		UsageUnitBytesPerSecond,
	}
}

// ToPtr() returns a *UsageUnit pointing to the current value.
func (c UsageUnit) ToPtr() *UsageUnit {
	return &c
}
