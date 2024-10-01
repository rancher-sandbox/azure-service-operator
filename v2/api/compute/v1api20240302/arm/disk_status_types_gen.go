// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

// Disk resource.
type Disk_STATUS struct {
	// ExtendedLocation: The extended location where the disk will be created. Extended location cannot be changed.
	ExtendedLocation *ExtendedLocation_STATUS `json:"extendedLocation,omitempty"`

	// Id: Resource Id
	Id *string `json:"id,omitempty"`

	// Location: Resource location
	Location *string `json:"location,omitempty"`

	// ManagedBy: A relative URI containing the ID of the VM that has the disk attached.
	ManagedBy *string `json:"managedBy,omitempty"`

	// ManagedByExtended: List of relative URIs containing the IDs of the VMs that have the disk attached. maxShares should be
	// set to a value greater than one for disks to allow attaching them to multiple VMs.
	ManagedByExtended []string `json:"managedByExtended,omitempty"`

	// Name: Resource name
	Name *string `json:"name,omitempty"`

	// Properties: Disk resource properties.
	Properties *DiskProperties_STATUS `json:"properties,omitempty"`

	// Sku: The disks sku name. Can be Standard_LRS, Premium_LRS, StandardSSD_LRS, UltraSSD_LRS, Premium_ZRS, StandardSSD_ZRS,
	// or  PremiumV2_LRS.
	Sku *DiskSku_STATUS `json:"sku,omitempty"`

	// Tags: Resource tags
	Tags map[string]string `json:"tags,omitempty"`

	// Type: Resource type
	Type *string `json:"type,omitempty"`

	// Zones: The Logical zone list for Disk.
	Zones []string `json:"zones,omitempty"`
}

// Disk resource properties.
type DiskProperties_STATUS struct {
	// BurstingEnabled: Set to true to enable bursting beyond the provisioned performance target of the disk. Bursting is
	// disabled by default. Does not apply to Ultra disks.
	BurstingEnabled *bool `json:"burstingEnabled,omitempty"`

	// BurstingEnabledTime: Latest time when bursting was last enabled on a disk.
	BurstingEnabledTime *string `json:"burstingEnabledTime,omitempty"`

	// CompletionPercent: Percentage complete for the background copy when a resource is created via the CopyStart operation.
	CompletionPercent *float64 `json:"completionPercent,omitempty"`

	// CreationData: Disk source information. CreationData information cannot be changed after the disk has been created.
	CreationData *CreationData_STATUS `json:"creationData,omitempty"`

	// DataAccessAuthMode: Additional authentication requirements when exporting or uploading to a disk or snapshot.
	DataAccessAuthMode *DataAccessAuthMode_STATUS `json:"dataAccessAuthMode,omitempty"`

	// DiskAccessId: ARM id of the DiskAccess resource for using private endpoints on disks.
	DiskAccessId *string `json:"diskAccessId,omitempty"`

	// DiskIOPSReadOnly: The total number of IOPS that will be allowed across all VMs mounting the shared disk as ReadOnly. One
	// operation can transfer between 4k and 256k bytes.
	DiskIOPSReadOnly *int `json:"diskIOPSReadOnly,omitempty"`

	// DiskIOPSReadWrite: The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can
	// transfer between 4k and 256k bytes.
	DiskIOPSReadWrite *int `json:"diskIOPSReadWrite,omitempty"`

	// DiskMBpsReadOnly: The total throughput (MBps) that will be allowed across all VMs mounting the shared disk as ReadOnly.
	// MBps means millions of bytes per second - MB here uses the ISO notation, of powers of 10.
	DiskMBpsReadOnly *int `json:"diskMBpsReadOnly,omitempty"`

	// DiskMBpsReadWrite: The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes
	// per second - MB here uses the ISO notation, of powers of 10.
	DiskMBpsReadWrite *int `json:"diskMBpsReadWrite,omitempty"`

	// DiskSizeBytes: The size of the disk in bytes. This field is read only.
	DiskSizeBytes *int `json:"diskSizeBytes,omitempty"`

	// DiskSizeGB: If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to
	// create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only
	// allowed if the disk is not attached to a running VM, and can only increase the disk's size.
	DiskSizeGB *int `json:"diskSizeGB,omitempty"`

	// DiskState: The state of the disk.
	DiskState *DiskState_STATUS `json:"diskState,omitempty"`

	// Encryption: Encryption property can be used to encrypt data at rest with customer managed keys or platform managed keys.
	Encryption *Encryption_STATUS `json:"encryption,omitempty"`

	// EncryptionSettingsCollection: Encryption settings collection used for Azure Disk Encryption, can contain multiple
	// encryption settings per disk or snapshot.
	EncryptionSettingsCollection *EncryptionSettingsCollection_STATUS `json:"encryptionSettingsCollection,omitempty"`

	// HyperVGeneration: The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
	HyperVGeneration *DiskProperties_HyperVGeneration_STATUS `json:"hyperVGeneration,omitempty"`

	// LastOwnershipUpdateTime: The UTC time when the ownership state of the disk was last changed i.e., the time the disk was
	// last attached or detached from a VM or the time when the VM to which the disk was attached was deallocated or started.
	LastOwnershipUpdateTime *string `json:"LastOwnershipUpdateTime,omitempty"`

	// MaxShares: The maximum number of VMs that can attach to the disk at the same time. Value greater than one indicates a
	// disk that can be mounted on multiple VMs at the same time.
	MaxShares *int `json:"maxShares,omitempty"`

	// NetworkAccessPolicy: Policy for accessing the disk via network.
	NetworkAccessPolicy *NetworkAccessPolicy_STATUS `json:"networkAccessPolicy,omitempty"`

	// OptimizedForFrequentAttach: Setting this property to true improves reliability and performance of data disks that are
	// frequently (more than 5 times a day) by detached from one virtual machine and attached to another. This property should
	// not be set for disks that are not detached and attached frequently as it causes the disks to not align with the fault
	// domain of the virtual machine.
	OptimizedForFrequentAttach *bool `json:"optimizedForFrequentAttach,omitempty"`

	// OsType: The Operating System type.
	OsType *DiskProperties_OsType_STATUS `json:"osType,omitempty"`

	// PropertyUpdatesInProgress: Properties of the disk for which update is pending.
	PropertyUpdatesInProgress *PropertyUpdatesInProgress_STATUS `json:"propertyUpdatesInProgress,omitempty"`

	// ProvisioningState: The disk provisioning state.
	ProvisioningState *string `json:"provisioningState,omitempty"`

	// PublicNetworkAccess: Policy for controlling export on the disk.
	PublicNetworkAccess *PublicNetworkAccess_STATUS `json:"publicNetworkAccess,omitempty"`

	// PurchasePlan: Purchase plan information for the the image from which the OS disk was created. E.g. - {name:
	// 2019-Datacenter, publisher: MicrosoftWindowsServer, product: WindowsServer}
	PurchasePlan *PurchasePlan_STATUS `json:"purchasePlan,omitempty"`

	// SecurityProfile: Contains the security related information for the resource.
	SecurityProfile *DiskSecurityProfile_STATUS `json:"securityProfile,omitempty"`

	// ShareInfo: Details of the list of all VMs that have the disk attached. maxShares should be set to a value greater than
	// one for disks to allow attaching them to multiple VMs.
	ShareInfo []ShareInfoElement_STATUS `json:"shareInfo,omitempty"`

	// SupportedCapabilities: List of supported capabilities for the image from which the OS disk was created.
	SupportedCapabilities *SupportedCapabilities_STATUS `json:"supportedCapabilities,omitempty"`

	// SupportsHibernation: Indicates the OS on a disk supports hibernation.
	SupportsHibernation *bool `json:"supportsHibernation,omitempty"`

	// Tier: Performance tier of the disk (e.g, P4, S10) as described here:
	// https://azure.microsoft.com/en-us/pricing/details/managed-disks/. Does not apply to Ultra disks.
	Tier *string `json:"tier,omitempty"`

	// TimeCreated: The time when the disk was created.
	TimeCreated *string `json:"timeCreated,omitempty"`

	// UniqueId: Unique Guid identifying the resource.
	UniqueId *string `json:"uniqueId,omitempty"`
}

// The disks sku name. Can be Standard_LRS, Premium_LRS, StandardSSD_LRS, UltraSSD_LRS, Premium_ZRS, StandardSSD_ZRS, or
// PremiumV2_LRS.
type DiskSku_STATUS struct {
	// Name: The sku name.
	Name *DiskSku_Name_STATUS `json:"name,omitempty"`

	// Tier: The sku tier.
	Tier *string `json:"tier,omitempty"`
}

// Data used when creating a disk.
type CreationData_STATUS struct {
	// CreateOption: This enumerates the possible sources of a disk's creation.
	CreateOption *CreationData_CreateOption_STATUS `json:"createOption,omitempty"`

	// ElasticSanResourceId: Required if createOption is CopyFromSanSnapshot. This is the ARM id of the source elastic san
	// volume snapshot.
	ElasticSanResourceId *string `json:"elasticSanResourceId,omitempty"`

	// GalleryImageReference: Required if creating from a Gallery Image. The id/sharedGalleryImageId/communityGalleryImageId of
	// the ImageDiskReference will be the ARM id of the shared galley image version from which to create a disk.
	GalleryImageReference *ImageDiskReference_STATUS `json:"galleryImageReference,omitempty"`

	// ImageReference: Disk source information for PIR or user images.
	ImageReference *ImageDiskReference_STATUS `json:"imageReference,omitempty"`

	// LogicalSectorSize: Logical sector size in bytes for Ultra disks. Supported values are 512 ad 4096. 4096 is the default.
	LogicalSectorSize *int `json:"logicalSectorSize,omitempty"`

	// PerformancePlus: Set this flag to true to get a boost on the performance target of the disk deployed, see here on the
	// respective performance target. This flag can only be set on disk creation time and cannot be disabled after enabled.
	PerformancePlus *bool `json:"performancePlus,omitempty"`

	// ProvisionedBandwidthCopySpeed: If this field is set on a snapshot and createOption is CopyStart, the snapshot will be
	// copied at a quicker speed.
	ProvisionedBandwidthCopySpeed *CreationData_ProvisionedBandwidthCopySpeed_STATUS `json:"provisionedBandwidthCopySpeed,omitempty"`

	// SecurityDataUri: If createOption is ImportSecure, this is the URI of a blob to be imported into VM guest state.
	SecurityDataUri *string `json:"securityDataUri,omitempty"`

	// SourceResourceId: If createOption is Copy, this is the ARM id of the source snapshot or disk.
	SourceResourceId *string `json:"sourceResourceId,omitempty"`

	// SourceUniqueId: If this field is set, this is the unique id identifying the source of this resource.
	SourceUniqueId *string `json:"sourceUniqueId,omitempty"`

	// SourceUri: If createOption is Import, this is the URI of a blob to be imported into a managed disk.
	SourceUri *string `json:"sourceUri,omitempty"`

	// StorageAccountId: Required if createOption is Import. The Azure Resource Manager identifier of the storage account
	// containing the blob to import as a disk.
	StorageAccountId *string `json:"storageAccountId,omitempty"`

	// UploadSizeBytes: If createOption is Upload, this is the size of the contents of the upload including the VHD footer.
	// This value should be between 20972032 (20 MiB + 512 bytes for the VHD footer) and 35183298347520 bytes (32 TiB + 512
	// bytes for the VHD footer).
	UploadSizeBytes *int `json:"uploadSizeBytes,omitempty"`
}

// Additional authentication requirements when exporting or uploading to a disk or snapshot.
type DataAccessAuthMode_STATUS string

const (
	DataAccessAuthMode_STATUS_AzureActiveDirectory = DataAccessAuthMode_STATUS("AzureActiveDirectory")
	DataAccessAuthMode_STATUS_None                 = DataAccessAuthMode_STATUS("None")
)

// Mapping from string to DataAccessAuthMode_STATUS
var dataAccessAuthMode_STATUS_Values = map[string]DataAccessAuthMode_STATUS{
	"azureactivedirectory": DataAccessAuthMode_STATUS_AzureActiveDirectory,
	"none":                 DataAccessAuthMode_STATUS_None,
}

type DiskProperties_HyperVGeneration_STATUS string

const (
	DiskProperties_HyperVGeneration_STATUS_V1 = DiskProperties_HyperVGeneration_STATUS("V1")
	DiskProperties_HyperVGeneration_STATUS_V2 = DiskProperties_HyperVGeneration_STATUS("V2")
)

// Mapping from string to DiskProperties_HyperVGeneration_STATUS
var diskProperties_HyperVGeneration_STATUS_Values = map[string]DiskProperties_HyperVGeneration_STATUS{
	"v1": DiskProperties_HyperVGeneration_STATUS_V1,
	"v2": DiskProperties_HyperVGeneration_STATUS_V2,
}

type DiskProperties_OsType_STATUS string

const (
	DiskProperties_OsType_STATUS_Linux   = DiskProperties_OsType_STATUS("Linux")
	DiskProperties_OsType_STATUS_Windows = DiskProperties_OsType_STATUS("Windows")
)

// Mapping from string to DiskProperties_OsType_STATUS
var diskProperties_OsType_STATUS_Values = map[string]DiskProperties_OsType_STATUS{
	"linux":   DiskProperties_OsType_STATUS_Linux,
	"windows": DiskProperties_OsType_STATUS_Windows,
}

// Contains the security related information for the resource.
type DiskSecurityProfile_STATUS struct {
	// SecureVMDiskEncryptionSetId: ResourceId of the disk encryption set associated to Confidential VM supported disk
	// encrypted with customer managed key
	SecureVMDiskEncryptionSetId *string `json:"secureVMDiskEncryptionSetId,omitempty"`

	// SecurityType: Specifies the SecurityType of the VM. Applicable for OS disks only.
	SecurityType *DiskSecurityType_STATUS `json:"securityType,omitempty"`
}

type DiskSku_Name_STATUS string

const (
	DiskSku_Name_STATUS_PremiumV2_LRS   = DiskSku_Name_STATUS("PremiumV2_LRS")
	DiskSku_Name_STATUS_Premium_LRS     = DiskSku_Name_STATUS("Premium_LRS")
	DiskSku_Name_STATUS_Premium_ZRS     = DiskSku_Name_STATUS("Premium_ZRS")
	DiskSku_Name_STATUS_StandardSSD_LRS = DiskSku_Name_STATUS("StandardSSD_LRS")
	DiskSku_Name_STATUS_StandardSSD_ZRS = DiskSku_Name_STATUS("StandardSSD_ZRS")
	DiskSku_Name_STATUS_Standard_LRS    = DiskSku_Name_STATUS("Standard_LRS")
	DiskSku_Name_STATUS_UltraSSD_LRS    = DiskSku_Name_STATUS("UltraSSD_LRS")
)

// Mapping from string to DiskSku_Name_STATUS
var diskSku_Name_STATUS_Values = map[string]DiskSku_Name_STATUS{
	"premiumv2_lrs":   DiskSku_Name_STATUS_PremiumV2_LRS,
	"premium_lrs":     DiskSku_Name_STATUS_Premium_LRS,
	"premium_zrs":     DiskSku_Name_STATUS_Premium_ZRS,
	"standardssd_lrs": DiskSku_Name_STATUS_StandardSSD_LRS,
	"standardssd_zrs": DiskSku_Name_STATUS_StandardSSD_ZRS,
	"standard_lrs":    DiskSku_Name_STATUS_Standard_LRS,
	"ultrassd_lrs":    DiskSku_Name_STATUS_UltraSSD_LRS,
}

// This enumerates the possible state of the disk.
type DiskState_STATUS string

const (
	DiskState_STATUS_ActiveSAS       = DiskState_STATUS("ActiveSAS")
	DiskState_STATUS_ActiveSASFrozen = DiskState_STATUS("ActiveSASFrozen")
	DiskState_STATUS_ActiveUpload    = DiskState_STATUS("ActiveUpload")
	DiskState_STATUS_Attached        = DiskState_STATUS("Attached")
	DiskState_STATUS_Frozen          = DiskState_STATUS("Frozen")
	DiskState_STATUS_ReadyToUpload   = DiskState_STATUS("ReadyToUpload")
	DiskState_STATUS_Reserved        = DiskState_STATUS("Reserved")
	DiskState_STATUS_Unattached      = DiskState_STATUS("Unattached")
)

// Mapping from string to DiskState_STATUS
var diskState_STATUS_Values = map[string]DiskState_STATUS{
	"activesas":       DiskState_STATUS_ActiveSAS,
	"activesasfrozen": DiskState_STATUS_ActiveSASFrozen,
	"activeupload":    DiskState_STATUS_ActiveUpload,
	"attached":        DiskState_STATUS_Attached,
	"frozen":          DiskState_STATUS_Frozen,
	"readytoupload":   DiskState_STATUS_ReadyToUpload,
	"reserved":        DiskState_STATUS_Reserved,
	"unattached":      DiskState_STATUS_Unattached,
}

// Encryption at rest settings for disk or snapshot
type Encryption_STATUS struct {
	// DiskEncryptionSetId: ResourceId of the disk encryption set to use for enabling encryption at rest.
	DiskEncryptionSetId *string `json:"diskEncryptionSetId,omitempty"`

	// Type: The type of key used to encrypt the data of the disk.
	Type *EncryptionType_STATUS `json:"type,omitempty"`
}

// Encryption settings for disk or snapshot
type EncryptionSettingsCollection_STATUS struct {
	// Enabled: Set this flag to true and provide DiskEncryptionKey and optional KeyEncryptionKey to enable encryption. Set
	// this flag to false and remove DiskEncryptionKey and KeyEncryptionKey to disable encryption. If EncryptionSettings is
	// null in the request object, the existing settings remain unchanged.
	Enabled *bool `json:"enabled,omitempty"`

	// EncryptionSettings: A collection of encryption settings, one for each disk volume.
	EncryptionSettings []EncryptionSettingsElement_STATUS `json:"encryptionSettings,omitempty"`

	// EncryptionSettingsVersion: Describes what type of encryption is used for the disks. Once this field is set, it cannot be
	// overwritten. '1.0' corresponds to Azure Disk Encryption with AAD app.'1.1' corresponds to Azure Disk Encryption.
	EncryptionSettingsVersion *string `json:"encryptionSettingsVersion,omitempty"`
}

// Policy for accessing the disk via network.
type NetworkAccessPolicy_STATUS string

const (
	NetworkAccessPolicy_STATUS_AllowAll     = NetworkAccessPolicy_STATUS("AllowAll")
	NetworkAccessPolicy_STATUS_AllowPrivate = NetworkAccessPolicy_STATUS("AllowPrivate")
	NetworkAccessPolicy_STATUS_DenyAll      = NetworkAccessPolicy_STATUS("DenyAll")
)

// Mapping from string to NetworkAccessPolicy_STATUS
var networkAccessPolicy_STATUS_Values = map[string]NetworkAccessPolicy_STATUS{
	"allowall":     NetworkAccessPolicy_STATUS_AllowAll,
	"allowprivate": NetworkAccessPolicy_STATUS_AllowPrivate,
	"denyall":      NetworkAccessPolicy_STATUS_DenyAll,
}

// Properties of the disk for which update is pending.
type PropertyUpdatesInProgress_STATUS struct {
	// TargetTier: The target performance tier of the disk if a tier change operation is in progress.
	TargetTier *string `json:"targetTier,omitempty"`
}

// Policy for controlling export on the disk.
type PublicNetworkAccess_STATUS string

const (
	PublicNetworkAccess_STATUS_Disabled = PublicNetworkAccess_STATUS("Disabled")
	PublicNetworkAccess_STATUS_Enabled  = PublicNetworkAccess_STATUS("Enabled")
)

// Mapping from string to PublicNetworkAccess_STATUS
var publicNetworkAccess_STATUS_Values = map[string]PublicNetworkAccess_STATUS{
	"disabled": PublicNetworkAccess_STATUS_Disabled,
	"enabled":  PublicNetworkAccess_STATUS_Enabled,
}

// Used for establishing the purchase context of any 3rd Party artifact through MarketPlace.
type PurchasePlan_STATUS struct {
	// Name: The plan ID.
	Name *string `json:"name,omitempty"`

	// Product: Specifies the product of the image from the marketplace. This is the same value as Offer under the
	// imageReference element.
	Product *string `json:"product,omitempty"`

	// PromotionCode: The Offer Promotion Code.
	PromotionCode *string `json:"promotionCode,omitempty"`

	// Publisher: The publisher ID.
	Publisher *string `json:"publisher,omitempty"`
}

type ShareInfoElement_STATUS struct {
	// VmUri: A relative URI containing the ID of the VM that has the disk attached.
	VmUri *string `json:"vmUri,omitempty"`
}

// List of supported capabilities persisted on the disk resource for VM use.
type SupportedCapabilities_STATUS struct {
	// AcceleratedNetwork: True if the image from which the OS disk is created supports accelerated networking.
	AcceleratedNetwork *bool `json:"acceleratedNetwork,omitempty"`

	// Architecture: CPU architecture supported by an OS disk.
	Architecture *SupportedCapabilities_Architecture_STATUS `json:"architecture,omitempty"`

	// DiskControllerTypes: The disk controllers that an OS disk supports. If set it can be SCSI or SCSI, NVME or NVME, SCSI.
	DiskControllerTypes *string `json:"diskControllerTypes,omitempty"`
}

type CreationData_CreateOption_STATUS string

const (
	CreationData_CreateOption_STATUS_Attach               = CreationData_CreateOption_STATUS("Attach")
	CreationData_CreateOption_STATUS_Copy                 = CreationData_CreateOption_STATUS("Copy")
	CreationData_CreateOption_STATUS_CopyFromSanSnapshot  = CreationData_CreateOption_STATUS("CopyFromSanSnapshot")
	CreationData_CreateOption_STATUS_CopyStart            = CreationData_CreateOption_STATUS("CopyStart")
	CreationData_CreateOption_STATUS_Empty                = CreationData_CreateOption_STATUS("Empty")
	CreationData_CreateOption_STATUS_FromImage            = CreationData_CreateOption_STATUS("FromImage")
	CreationData_CreateOption_STATUS_Import               = CreationData_CreateOption_STATUS("Import")
	CreationData_CreateOption_STATUS_ImportSecure         = CreationData_CreateOption_STATUS("ImportSecure")
	CreationData_CreateOption_STATUS_Restore              = CreationData_CreateOption_STATUS("Restore")
	CreationData_CreateOption_STATUS_Upload               = CreationData_CreateOption_STATUS("Upload")
	CreationData_CreateOption_STATUS_UploadPreparedSecure = CreationData_CreateOption_STATUS("UploadPreparedSecure")
)

// Mapping from string to CreationData_CreateOption_STATUS
var creationData_CreateOption_STATUS_Values = map[string]CreationData_CreateOption_STATUS{
	"attach":               CreationData_CreateOption_STATUS_Attach,
	"copy":                 CreationData_CreateOption_STATUS_Copy,
	"copyfromsansnapshot":  CreationData_CreateOption_STATUS_CopyFromSanSnapshot,
	"copystart":            CreationData_CreateOption_STATUS_CopyStart,
	"empty":                CreationData_CreateOption_STATUS_Empty,
	"fromimage":            CreationData_CreateOption_STATUS_FromImage,
	"import":               CreationData_CreateOption_STATUS_Import,
	"importsecure":         CreationData_CreateOption_STATUS_ImportSecure,
	"restore":              CreationData_CreateOption_STATUS_Restore,
	"upload":               CreationData_CreateOption_STATUS_Upload,
	"uploadpreparedsecure": CreationData_CreateOption_STATUS_UploadPreparedSecure,
}

type CreationData_ProvisionedBandwidthCopySpeed_STATUS string

const (
	CreationData_ProvisionedBandwidthCopySpeed_STATUS_Enhanced = CreationData_ProvisionedBandwidthCopySpeed_STATUS("Enhanced")
	CreationData_ProvisionedBandwidthCopySpeed_STATUS_None     = CreationData_ProvisionedBandwidthCopySpeed_STATUS("None")
)

// Mapping from string to CreationData_ProvisionedBandwidthCopySpeed_STATUS
var creationData_ProvisionedBandwidthCopySpeed_STATUS_Values = map[string]CreationData_ProvisionedBandwidthCopySpeed_STATUS{
	"enhanced": CreationData_ProvisionedBandwidthCopySpeed_STATUS_Enhanced,
	"none":     CreationData_ProvisionedBandwidthCopySpeed_STATUS_None,
}

// Specifies the SecurityType of the VM. Applicable for OS disks only.
type DiskSecurityType_STATUS string

const (
	DiskSecurityType_STATUS_ConfidentialVM_DiskEncryptedWithCustomerKey             = DiskSecurityType_STATUS("ConfidentialVM_DiskEncryptedWithCustomerKey")
	DiskSecurityType_STATUS_ConfidentialVM_DiskEncryptedWithPlatformKey             = DiskSecurityType_STATUS("ConfidentialVM_DiskEncryptedWithPlatformKey")
	DiskSecurityType_STATUS_ConfidentialVM_NonPersistedTPM                          = DiskSecurityType_STATUS("ConfidentialVM_NonPersistedTPM")
	DiskSecurityType_STATUS_ConfidentialVM_VMGuestStateOnlyEncryptedWithPlatformKey = DiskSecurityType_STATUS("ConfidentialVM_VMGuestStateOnlyEncryptedWithPlatformKey")
	DiskSecurityType_STATUS_TrustedLaunch                                           = DiskSecurityType_STATUS("TrustedLaunch")
)

// Mapping from string to DiskSecurityType_STATUS
var diskSecurityType_STATUS_Values = map[string]DiskSecurityType_STATUS{
	"confidentialvm_diskencryptedwithcustomerkey":             DiskSecurityType_STATUS_ConfidentialVM_DiskEncryptedWithCustomerKey,
	"confidentialvm_diskencryptedwithplatformkey":             DiskSecurityType_STATUS_ConfidentialVM_DiskEncryptedWithPlatformKey,
	"confidentialvm_nonpersistedtpm":                          DiskSecurityType_STATUS_ConfidentialVM_NonPersistedTPM,
	"confidentialvm_vmgueststateonlyencryptedwithplatformkey": DiskSecurityType_STATUS_ConfidentialVM_VMGuestStateOnlyEncryptedWithPlatformKey,
	"trustedlaunch": DiskSecurityType_STATUS_TrustedLaunch,
}

// Encryption settings for one disk volume.
type EncryptionSettingsElement_STATUS struct {
	// DiskEncryptionKey: Key Vault Secret Url and vault id of the disk encryption key
	DiskEncryptionKey *KeyVaultAndSecretReference_STATUS `json:"diskEncryptionKey,omitempty"`

	// KeyEncryptionKey: Key Vault Key Url and vault id of the key encryption key. KeyEncryptionKey is optional and when
	// provided is used to unwrap the disk encryption key.
	KeyEncryptionKey *KeyVaultAndKeyReference_STATUS `json:"keyEncryptionKey,omitempty"`
}

// The type of key used to encrypt the data of the disk.
type EncryptionType_STATUS string

const (
	EncryptionType_STATUS_EncryptionAtRestWithCustomerKey             = EncryptionType_STATUS("EncryptionAtRestWithCustomerKey")
	EncryptionType_STATUS_EncryptionAtRestWithPlatformAndCustomerKeys = EncryptionType_STATUS("EncryptionAtRestWithPlatformAndCustomerKeys")
	EncryptionType_STATUS_EncryptionAtRestWithPlatformKey             = EncryptionType_STATUS("EncryptionAtRestWithPlatformKey")
)

// Mapping from string to EncryptionType_STATUS
var encryptionType_STATUS_Values = map[string]EncryptionType_STATUS{
	"encryptionatrestwithcustomerkey":             EncryptionType_STATUS_EncryptionAtRestWithCustomerKey,
	"encryptionatrestwithplatformandcustomerkeys": EncryptionType_STATUS_EncryptionAtRestWithPlatformAndCustomerKeys,
	"encryptionatrestwithplatformkey":             EncryptionType_STATUS_EncryptionAtRestWithPlatformKey,
}

// The source image used for creating the disk.
type ImageDiskReference_STATUS struct {
	// CommunityGalleryImageId: A relative uri containing a community Azure Compute Gallery image reference.
	CommunityGalleryImageId *string `json:"communityGalleryImageId,omitempty"`

	// Id: A relative uri containing either a Platform Image Repository, user image, or Azure Compute Gallery image reference.
	Id *string `json:"id,omitempty"`

	// Lun: If the disk is created from an image's data disk, this is an index that indicates which of the data disks in the
	// image to use. For OS disks, this field is null.
	Lun *int `json:"lun,omitempty"`

	// SharedGalleryImageId: A relative uri containing a direct shared Azure Compute Gallery image reference.
	SharedGalleryImageId *string `json:"sharedGalleryImageId,omitempty"`
}

type SupportedCapabilities_Architecture_STATUS string

const (
	SupportedCapabilities_Architecture_STATUS_Arm64 = SupportedCapabilities_Architecture_STATUS("Arm64")
	SupportedCapabilities_Architecture_STATUS_X64   = SupportedCapabilities_Architecture_STATUS("x64")
)

// Mapping from string to SupportedCapabilities_Architecture_STATUS
var supportedCapabilities_Architecture_STATUS_Values = map[string]SupportedCapabilities_Architecture_STATUS{
	"arm64": SupportedCapabilities_Architecture_STATUS_Arm64,
	"x64":   SupportedCapabilities_Architecture_STATUS_X64,
}

// Key Vault Key Url and vault id of KeK, KeK is optional and when provided is used to unwrap the encryptionKey
type KeyVaultAndKeyReference_STATUS struct {
	// KeyUrl: Url pointing to a key or secret in KeyVault
	KeyUrl *string `json:"keyUrl,omitempty"`

	// SourceVault: Resource id of the KeyVault containing the key or secret
	SourceVault *SourceVault_STATUS `json:"sourceVault,omitempty"`
}

// Key Vault Secret Url and vault id of the encryption key
type KeyVaultAndSecretReference_STATUS struct {
	// SecretUrl: Url pointing to a key or secret in KeyVault
	SecretUrl *string `json:"secretUrl,omitempty"`

	// SourceVault: Resource id of the KeyVault containing the key or secret
	SourceVault *SourceVault_STATUS `json:"sourceVault,omitempty"`
}