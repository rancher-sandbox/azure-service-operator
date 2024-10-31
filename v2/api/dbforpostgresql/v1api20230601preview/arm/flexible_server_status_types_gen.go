// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

type FlexibleServer_STATUS struct {
	// Id: Fully qualified resource ID for the resource. E.g.
	// "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id *string `json:"id,omitempty"`

	// Identity: Describes the identity of the application.
	Identity *UserAssignedIdentity_STATUS `json:"identity,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the server.
	Properties *ServerProperties_STATUS `json:"properties,omitempty"`

	// Sku: The SKU (pricing tier) of the server.
	Sku *Sku_STATUS `json:"sku,omitempty"`

	// SystemData: Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData_STATUS `json:"systemData,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// The properties of a server.
type ServerProperties_STATUS struct {
	// AdministratorLogin: The administrator's login name of a server. Can only be specified when the server is being created
	// (and is required for creation).
	AdministratorLogin *string `json:"administratorLogin,omitempty"`

	// AuthConfig: AuthConfig properties of a server.
	AuthConfig *AuthConfig_STATUS `json:"authConfig,omitempty"`

	// AvailabilityZone: availability zone information of the server.
	AvailabilityZone *string `json:"availabilityZone,omitempty"`

	// Backup: Backup properties of a server.
	Backup *Backup_STATUS `json:"backup,omitempty"`

	// CreateMode: The mode to create a new PostgreSQL server.
	CreateMode *ServerProperties_CreateMode_STATUS `json:"createMode,omitempty"`

	// DataEncryption: Data encryption properties of a server.
	DataEncryption *DataEncryption_STATUS `json:"dataEncryption,omitempty"`

	// FullyQualifiedDomainName: The fully qualified domain name of a server.
	FullyQualifiedDomainName *string `json:"fullyQualifiedDomainName,omitempty"`

	// HighAvailability: High availability properties of a server.
	HighAvailability *HighAvailability_STATUS `json:"highAvailability,omitempty"`

	// MaintenanceWindow: Maintenance window properties of a server.
	MaintenanceWindow *MaintenanceWindow_STATUS `json:"maintenanceWindow,omitempty"`

	// MinorVersion: The minor version of the server.
	MinorVersion *string `json:"minorVersion,omitempty"`

	// Network: Network properties of a server. This Network property is required to be passed only in case you want the server
	// to be Private access server.
	Network *Network_STATUS `json:"network,omitempty"`

	// PointInTimeUTC: Restore point creation time (ISO8601 format), specifying the time to restore from. It's required when
	// 'createMode' is 'PointInTimeRestore' or 'GeoRestore' or 'ReviveDropped'.
	PointInTimeUTC *string `json:"pointInTimeUTC,omitempty"`

	// PrivateEndpointConnections: List of private endpoint connections associated with the specified resource.
	PrivateEndpointConnections []PrivateEndpointConnection_STATUS `json:"privateEndpointConnections,omitempty"`

	// Replica: Replica properties of a server. These Replica properties are required to be passed only in case you want to
	// Promote a server.
	Replica *Replica_STATUS `json:"replica,omitempty"`

	// ReplicaCapacity: Replicas allowed for a server.
	ReplicaCapacity *int `json:"replicaCapacity,omitempty"`

	// ReplicationRole: Replication role of the server
	ReplicationRole *ReplicationRole_STATUS `json:"replicationRole,omitempty"`

	// SourceServerResourceId: The source server resource ID to restore from. It's required when 'createMode' is
	// 'PointInTimeRestore' or 'GeoRestore' or 'Replica' or 'ReviveDropped'. This property is returned only for Replica server
	SourceServerResourceId *string `json:"sourceServerResourceId,omitempty"`

	// State: A state of a server that is visible to user.
	State *ServerProperties_State_STATUS `json:"state,omitempty"`

	// Storage: Storage properties of a server.
	Storage *Storage_STATUS `json:"storage,omitempty"`

	// Version: PostgreSQL Server version.
	Version *ServerVersion_STATUS `json:"version,omitempty"`
}

// Sku information related properties of a server.
type Sku_STATUS struct {
	// Name: The name of the sku, typically, tier + family + cores, e.g. Standard_D4s_v3.
	Name *string `json:"name,omitempty"`

	// Tier: The tier of the particular SKU, e.g. Burstable.
	Tier *Sku_Tier_STATUS `json:"tier,omitempty"`
}

// Metadata pertaining to creation and last modification of the resource.
type SystemData_STATUS struct {
	// CreatedAt: The timestamp of resource creation (UTC).
	CreatedAt *string `json:"createdAt,omitempty"`

	// CreatedBy: The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// CreatedByType: The type of identity that created the resource.
	CreatedByType *SystemData_CreatedByType_STATUS `json:"createdByType,omitempty"`

	// LastModifiedAt: The timestamp of resource last modification (UTC)
	LastModifiedAt *string `json:"lastModifiedAt,omitempty"`

	// LastModifiedBy: The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// LastModifiedByType: The type of identity that last modified the resource.
	LastModifiedByType *SystemData_LastModifiedByType_STATUS `json:"lastModifiedByType,omitempty"`
}

// Information describing the identities associated with this application.
type UserAssignedIdentity_STATUS struct {
	// TenantId: Tenant id of the server.
	TenantId *string `json:"tenantId,omitempty"`

	// Type: the types of identities associated with this resource; currently restricted to 'None and UserAssigned'
	Type *UserAssignedIdentity_Type_STATUS `json:"type,omitempty"`

	// UserAssignedIdentities: represents user assigned identities map.
	UserAssignedIdentities map[string]UserIdentity_STATUS `json:"userAssignedIdentities,omitempty"`
}

// Authentication configuration properties of a server
type AuthConfig_STATUS struct {
	// ActiveDirectoryAuth: If Enabled, Azure Active Directory authentication is enabled.
	ActiveDirectoryAuth *AuthConfig_ActiveDirectoryAuth_STATUS `json:"activeDirectoryAuth,omitempty"`

	// PasswordAuth: If Enabled, Password authentication is enabled.
	PasswordAuth *AuthConfig_PasswordAuth_STATUS `json:"passwordAuth,omitempty"`

	// TenantId: Tenant id of the server.
	TenantId *string `json:"tenantId,omitempty"`
}

// Backup properties of a server
type Backup_STATUS struct {
	// BackupRetentionDays: Backup retention days for the server.
	BackupRetentionDays *int `json:"backupRetentionDays,omitempty"`

	// EarliestRestoreDate: The earliest restore point time (ISO8601 format) for server.
	EarliestRestoreDate *string `json:"earliestRestoreDate,omitempty"`

	// GeoRedundantBackup: A value indicating whether Geo-Redundant backup is enabled on the server.
	GeoRedundantBackup *Backup_GeoRedundantBackup_STATUS `json:"geoRedundantBackup,omitempty"`
}

// Data encryption properties of a server
type DataEncryption_STATUS struct {
	// GeoBackupEncryptionKeyStatus: Geo-backup encryption key status for Data encryption enabled server.
	GeoBackupEncryptionKeyStatus *DataEncryption_GeoBackupEncryptionKeyStatus_STATUS `json:"geoBackupEncryptionKeyStatus,omitempty"`

	// GeoBackupKeyURI: URI for the key in keyvault for data encryption for geo-backup of server.
	GeoBackupKeyURI *string `json:"geoBackupKeyURI,omitempty"`

	// GeoBackupUserAssignedIdentityId: Resource Id for the User assigned identity to be used for data encryption for
	// geo-backup of server.
	GeoBackupUserAssignedIdentityId *string `json:"geoBackupUserAssignedIdentityId,omitempty"`

	// PrimaryEncryptionKeyStatus: Primary encryption key status for Data encryption enabled server.
	PrimaryEncryptionKeyStatus *DataEncryption_PrimaryEncryptionKeyStatus_STATUS `json:"primaryEncryptionKeyStatus,omitempty"`

	// PrimaryKeyURI: URI for the key in keyvault for data encryption of the primary server.
	PrimaryKeyURI *string `json:"primaryKeyURI,omitempty"`

	// PrimaryUserAssignedIdentityId: Resource Id for the User assigned identity to be used for data encryption of the primary
	// server.
	PrimaryUserAssignedIdentityId *string `json:"primaryUserAssignedIdentityId,omitempty"`

	// Type: Data encryption type to depict if it is System Managed vs Azure Key vault.
	Type *DataEncryption_Type_STATUS `json:"type,omitempty"`
}

// High availability properties of a server
type HighAvailability_STATUS struct {
	// Mode: The HA mode for the server.
	Mode *HighAvailability_Mode_STATUS `json:"mode,omitempty"`

	// StandbyAvailabilityZone: availability zone information of the standby.
	StandbyAvailabilityZone *string `json:"standbyAvailabilityZone,omitempty"`

	// State: A state of a HA server that is visible to user.
	State *HighAvailability_State_STATUS `json:"state,omitempty"`
}

// Maintenance window properties of a server.
type MaintenanceWindow_STATUS struct {
	// CustomWindow: indicates whether custom window is enabled or disabled
	CustomWindow *string `json:"customWindow,omitempty"`

	// DayOfWeek: day of week for maintenance window
	DayOfWeek *int `json:"dayOfWeek,omitempty"`

	// StartHour: start hour for maintenance window
	StartHour *int `json:"startHour,omitempty"`

	// StartMinute: start minute for maintenance window
	StartMinute *int `json:"startMinute,omitempty"`
}

// Network properties of a server.
type Network_STATUS struct {
	// DelegatedSubnetResourceId: Delegated subnet arm resource id. This is required to be passed during create, in case we
	// want the server to be VNET injected, i.e. Private access server. During update, pass this only if we want to update the
	// value for Private DNS zone.
	DelegatedSubnetResourceId *string `json:"delegatedSubnetResourceId,omitempty"`

	// PrivateDnsZoneArmResourceId: Private dns zone arm resource id. This is required to be passed during create, in case we
	// want the server to be VNET injected, i.e. Private access server. During update, pass this only if we want to update the
	// value for Private DNS zone.
	PrivateDnsZoneArmResourceId *string `json:"privateDnsZoneArmResourceId,omitempty"`

	// PublicNetworkAccess: public network access is enabled or not
	PublicNetworkAccess *Network_PublicNetworkAccess_STATUS `json:"publicNetworkAccess,omitempty"`
}

// The private endpoint connection resource.
type PrivateEndpointConnection_STATUS struct {
	// Id: Fully qualified resource ID for the resource. E.g.
	// "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id *string `json:"id,omitempty"`
}

// Replica properties of a server
type Replica_STATUS struct {
	// Capacity: Replicas allowed for a server.
	Capacity *int `json:"capacity,omitempty"`

	// PromoteMode: Sets the promote mode for a replica server. This is a write only property.
	PromoteMode *Replica_PromoteMode_STATUS `json:"promoteMode,omitempty"`

	// PromoteOption: Sets the promote options for a replica server. This is a write only property.
	PromoteOption *Replica_PromoteOption_STATUS `json:"promoteOption,omitempty"`

	// ReplicationState: Gets the replication state of a replica server. This property is returned only for replicas api call.
	// Supported values are Active, Catchup, Provisioning, Updating, Broken, Reconfiguring
	ReplicationState *Replica_ReplicationState_STATUS `json:"replicationState,omitempty"`

	// Role: Used to indicate role of the server in replication set.
	Role *ReplicationRole_STATUS `json:"role,omitempty"`
}

// Used to indicate role of the server in replication set.
type ReplicationRole_STATUS string

const (
	ReplicationRole_STATUS_AsyncReplica    = ReplicationRole_STATUS("AsyncReplica")
	ReplicationRole_STATUS_GeoAsyncReplica = ReplicationRole_STATUS("GeoAsyncReplica")
	ReplicationRole_STATUS_None            = ReplicationRole_STATUS("None")
	ReplicationRole_STATUS_Primary         = ReplicationRole_STATUS("Primary")
)

// Mapping from string to ReplicationRole_STATUS
var replicationRole_STATUS_Values = map[string]ReplicationRole_STATUS{
	"asyncreplica":    ReplicationRole_STATUS_AsyncReplica,
	"geoasyncreplica": ReplicationRole_STATUS_GeoAsyncReplica,
	"none":            ReplicationRole_STATUS_None,
	"primary":         ReplicationRole_STATUS_Primary,
}

type ServerProperties_CreateMode_STATUS string

const (
	ServerProperties_CreateMode_STATUS_Create             = ServerProperties_CreateMode_STATUS("Create")
	ServerProperties_CreateMode_STATUS_Default            = ServerProperties_CreateMode_STATUS("Default")
	ServerProperties_CreateMode_STATUS_GeoRestore         = ServerProperties_CreateMode_STATUS("GeoRestore")
	ServerProperties_CreateMode_STATUS_PointInTimeRestore = ServerProperties_CreateMode_STATUS("PointInTimeRestore")
	ServerProperties_CreateMode_STATUS_Replica            = ServerProperties_CreateMode_STATUS("Replica")
	ServerProperties_CreateMode_STATUS_ReviveDropped      = ServerProperties_CreateMode_STATUS("ReviveDropped")
	ServerProperties_CreateMode_STATUS_Update             = ServerProperties_CreateMode_STATUS("Update")
)

// Mapping from string to ServerProperties_CreateMode_STATUS
var serverProperties_CreateMode_STATUS_Values = map[string]ServerProperties_CreateMode_STATUS{
	"create":             ServerProperties_CreateMode_STATUS_Create,
	"default":            ServerProperties_CreateMode_STATUS_Default,
	"georestore":         ServerProperties_CreateMode_STATUS_GeoRestore,
	"pointintimerestore": ServerProperties_CreateMode_STATUS_PointInTimeRestore,
	"replica":            ServerProperties_CreateMode_STATUS_Replica,
	"revivedropped":      ServerProperties_CreateMode_STATUS_ReviveDropped,
	"update":             ServerProperties_CreateMode_STATUS_Update,
}

type ServerProperties_State_STATUS string

const (
	ServerProperties_State_STATUS_Disabled = ServerProperties_State_STATUS("Disabled")
	ServerProperties_State_STATUS_Dropping = ServerProperties_State_STATUS("Dropping")
	ServerProperties_State_STATUS_Ready    = ServerProperties_State_STATUS("Ready")
	ServerProperties_State_STATUS_Starting = ServerProperties_State_STATUS("Starting")
	ServerProperties_State_STATUS_Stopped  = ServerProperties_State_STATUS("Stopped")
	ServerProperties_State_STATUS_Stopping = ServerProperties_State_STATUS("Stopping")
	ServerProperties_State_STATUS_Updating = ServerProperties_State_STATUS("Updating")
)

// Mapping from string to ServerProperties_State_STATUS
var serverProperties_State_STATUS_Values = map[string]ServerProperties_State_STATUS{
	"disabled": ServerProperties_State_STATUS_Disabled,
	"dropping": ServerProperties_State_STATUS_Dropping,
	"ready":    ServerProperties_State_STATUS_Ready,
	"starting": ServerProperties_State_STATUS_Starting,
	"stopped":  ServerProperties_State_STATUS_Stopped,
	"stopping": ServerProperties_State_STATUS_Stopping,
	"updating": ServerProperties_State_STATUS_Updating,
}

// The version of a server.
type ServerVersion_STATUS string

const (
	ServerVersion_STATUS_11 = ServerVersion_STATUS("11")
	ServerVersion_STATUS_12 = ServerVersion_STATUS("12")
	ServerVersion_STATUS_13 = ServerVersion_STATUS("13")
	ServerVersion_STATUS_14 = ServerVersion_STATUS("14")
	ServerVersion_STATUS_15 = ServerVersion_STATUS("15")
	ServerVersion_STATUS_16 = ServerVersion_STATUS("16")
)

// Mapping from string to ServerVersion_STATUS
var serverVersion_STATUS_Values = map[string]ServerVersion_STATUS{
	"11": ServerVersion_STATUS_11,
	"12": ServerVersion_STATUS_12,
	"13": ServerVersion_STATUS_13,
	"14": ServerVersion_STATUS_14,
	"15": ServerVersion_STATUS_15,
	"16": ServerVersion_STATUS_16,
}

type Sku_Tier_STATUS string

const (
	Sku_Tier_STATUS_Burstable       = Sku_Tier_STATUS("Burstable")
	Sku_Tier_STATUS_GeneralPurpose  = Sku_Tier_STATUS("GeneralPurpose")
	Sku_Tier_STATUS_MemoryOptimized = Sku_Tier_STATUS("MemoryOptimized")
)

// Mapping from string to Sku_Tier_STATUS
var sku_Tier_STATUS_Values = map[string]Sku_Tier_STATUS{
	"burstable":       Sku_Tier_STATUS_Burstable,
	"generalpurpose":  Sku_Tier_STATUS_GeneralPurpose,
	"memoryoptimized": Sku_Tier_STATUS_MemoryOptimized,
}

// Storage properties of a server
type Storage_STATUS struct {
	// AutoGrow: Flag to enable / disable Storage Auto grow for flexible server.
	AutoGrow *Storage_AutoGrow_STATUS `json:"autoGrow,omitempty"`

	// Iops: Storage tier IOPS quantity. This property is required to be set for storage Type PremiumV2_LRS
	Iops *int `json:"iops,omitempty"`

	// StorageSizeGB: Max storage allowed for a server.
	StorageSizeGB *int `json:"storageSizeGB,omitempty"`

	// Throughput: Storage throughput for the server. This is required to be set for storage Type PremiumV2_LRS
	Throughput *int `json:"throughput,omitempty"`

	// Tier: Name of storage tier for IOPS.
	Tier *Storage_Tier_STATUS `json:"tier,omitempty"`

	// Type: Storage type for the server. Allowed values are Premium_LRS and PremiumV2_LRS, and default is Premium_LRS if not
	// specified
	Type *Storage_Type_STATUS `json:"type,omitempty"`
}

type SystemData_CreatedByType_STATUS string

const (
	SystemData_CreatedByType_STATUS_Application     = SystemData_CreatedByType_STATUS("Application")
	SystemData_CreatedByType_STATUS_Key             = SystemData_CreatedByType_STATUS("Key")
	SystemData_CreatedByType_STATUS_ManagedIdentity = SystemData_CreatedByType_STATUS("ManagedIdentity")
	SystemData_CreatedByType_STATUS_User            = SystemData_CreatedByType_STATUS("User")
)

// Mapping from string to SystemData_CreatedByType_STATUS
var systemData_CreatedByType_STATUS_Values = map[string]SystemData_CreatedByType_STATUS{
	"application":     SystemData_CreatedByType_STATUS_Application,
	"key":             SystemData_CreatedByType_STATUS_Key,
	"managedidentity": SystemData_CreatedByType_STATUS_ManagedIdentity,
	"user":            SystemData_CreatedByType_STATUS_User,
}

type SystemData_LastModifiedByType_STATUS string

const (
	SystemData_LastModifiedByType_STATUS_Application     = SystemData_LastModifiedByType_STATUS("Application")
	SystemData_LastModifiedByType_STATUS_Key             = SystemData_LastModifiedByType_STATUS("Key")
	SystemData_LastModifiedByType_STATUS_ManagedIdentity = SystemData_LastModifiedByType_STATUS("ManagedIdentity")
	SystemData_LastModifiedByType_STATUS_User            = SystemData_LastModifiedByType_STATUS("User")
)

// Mapping from string to SystemData_LastModifiedByType_STATUS
var systemData_LastModifiedByType_STATUS_Values = map[string]SystemData_LastModifiedByType_STATUS{
	"application":     SystemData_LastModifiedByType_STATUS_Application,
	"key":             SystemData_LastModifiedByType_STATUS_Key,
	"managedidentity": SystemData_LastModifiedByType_STATUS_ManagedIdentity,
	"user":            SystemData_LastModifiedByType_STATUS_User,
}

type UserAssignedIdentity_Type_STATUS string

const (
	UserAssignedIdentity_Type_STATUS_None         = UserAssignedIdentity_Type_STATUS("None")
	UserAssignedIdentity_Type_STATUS_UserAssigned = UserAssignedIdentity_Type_STATUS("UserAssigned")
)

// Mapping from string to UserAssignedIdentity_Type_STATUS
var userAssignedIdentity_Type_STATUS_Values = map[string]UserAssignedIdentity_Type_STATUS{
	"none":         UserAssignedIdentity_Type_STATUS_None,
	"userassigned": UserAssignedIdentity_Type_STATUS_UserAssigned,
}

// Describes a single user-assigned identity associated with the application.
type UserIdentity_STATUS struct {
	// ClientId: the client identifier of the Service Principal which this identity represents.
	ClientId *string `json:"clientId,omitempty"`

	// PrincipalId: the object identifier of the Service Principal which this identity represents.
	PrincipalId *string `json:"principalId,omitempty"`
}

type AuthConfig_ActiveDirectoryAuth_STATUS string

const (
	AuthConfig_ActiveDirectoryAuth_STATUS_Disabled = AuthConfig_ActiveDirectoryAuth_STATUS("Disabled")
	AuthConfig_ActiveDirectoryAuth_STATUS_Enabled  = AuthConfig_ActiveDirectoryAuth_STATUS("Enabled")
)

// Mapping from string to AuthConfig_ActiveDirectoryAuth_STATUS
var authConfig_ActiveDirectoryAuth_STATUS_Values = map[string]AuthConfig_ActiveDirectoryAuth_STATUS{
	"disabled": AuthConfig_ActiveDirectoryAuth_STATUS_Disabled,
	"enabled":  AuthConfig_ActiveDirectoryAuth_STATUS_Enabled,
}

type AuthConfig_PasswordAuth_STATUS string

const (
	AuthConfig_PasswordAuth_STATUS_Disabled = AuthConfig_PasswordAuth_STATUS("Disabled")
	AuthConfig_PasswordAuth_STATUS_Enabled  = AuthConfig_PasswordAuth_STATUS("Enabled")
)

// Mapping from string to AuthConfig_PasswordAuth_STATUS
var authConfig_PasswordAuth_STATUS_Values = map[string]AuthConfig_PasswordAuth_STATUS{
	"disabled": AuthConfig_PasswordAuth_STATUS_Disabled,
	"enabled":  AuthConfig_PasswordAuth_STATUS_Enabled,
}

type Backup_GeoRedundantBackup_STATUS string

const (
	Backup_GeoRedundantBackup_STATUS_Disabled = Backup_GeoRedundantBackup_STATUS("Disabled")
	Backup_GeoRedundantBackup_STATUS_Enabled  = Backup_GeoRedundantBackup_STATUS("Enabled")
)

// Mapping from string to Backup_GeoRedundantBackup_STATUS
var backup_GeoRedundantBackup_STATUS_Values = map[string]Backup_GeoRedundantBackup_STATUS{
	"disabled": Backup_GeoRedundantBackup_STATUS_Disabled,
	"enabled":  Backup_GeoRedundantBackup_STATUS_Enabled,
}

type DataEncryption_GeoBackupEncryptionKeyStatus_STATUS string

const (
	DataEncryption_GeoBackupEncryptionKeyStatus_STATUS_Invalid = DataEncryption_GeoBackupEncryptionKeyStatus_STATUS("Invalid")
	DataEncryption_GeoBackupEncryptionKeyStatus_STATUS_Valid   = DataEncryption_GeoBackupEncryptionKeyStatus_STATUS("Valid")
)

// Mapping from string to DataEncryption_GeoBackupEncryptionKeyStatus_STATUS
var dataEncryption_GeoBackupEncryptionKeyStatus_STATUS_Values = map[string]DataEncryption_GeoBackupEncryptionKeyStatus_STATUS{
	"invalid": DataEncryption_GeoBackupEncryptionKeyStatus_STATUS_Invalid,
	"valid":   DataEncryption_GeoBackupEncryptionKeyStatus_STATUS_Valid,
}

type DataEncryption_PrimaryEncryptionKeyStatus_STATUS string

const (
	DataEncryption_PrimaryEncryptionKeyStatus_STATUS_Invalid = DataEncryption_PrimaryEncryptionKeyStatus_STATUS("Invalid")
	DataEncryption_PrimaryEncryptionKeyStatus_STATUS_Valid   = DataEncryption_PrimaryEncryptionKeyStatus_STATUS("Valid")
)

// Mapping from string to DataEncryption_PrimaryEncryptionKeyStatus_STATUS
var dataEncryption_PrimaryEncryptionKeyStatus_STATUS_Values = map[string]DataEncryption_PrimaryEncryptionKeyStatus_STATUS{
	"invalid": DataEncryption_PrimaryEncryptionKeyStatus_STATUS_Invalid,
	"valid":   DataEncryption_PrimaryEncryptionKeyStatus_STATUS_Valid,
}

type DataEncryption_Type_STATUS string

const (
	DataEncryption_Type_STATUS_AzureKeyVault = DataEncryption_Type_STATUS("AzureKeyVault")
	DataEncryption_Type_STATUS_SystemManaged = DataEncryption_Type_STATUS("SystemManaged")
)

// Mapping from string to DataEncryption_Type_STATUS
var dataEncryption_Type_STATUS_Values = map[string]DataEncryption_Type_STATUS{
	"azurekeyvault": DataEncryption_Type_STATUS_AzureKeyVault,
	"systemmanaged": DataEncryption_Type_STATUS_SystemManaged,
}

type HighAvailability_Mode_STATUS string

const (
	HighAvailability_Mode_STATUS_Disabled      = HighAvailability_Mode_STATUS("Disabled")
	HighAvailability_Mode_STATUS_SameZone      = HighAvailability_Mode_STATUS("SameZone")
	HighAvailability_Mode_STATUS_ZoneRedundant = HighAvailability_Mode_STATUS("ZoneRedundant")
)

// Mapping from string to HighAvailability_Mode_STATUS
var highAvailability_Mode_STATUS_Values = map[string]HighAvailability_Mode_STATUS{
	"disabled":      HighAvailability_Mode_STATUS_Disabled,
	"samezone":      HighAvailability_Mode_STATUS_SameZone,
	"zoneredundant": HighAvailability_Mode_STATUS_ZoneRedundant,
}

type HighAvailability_State_STATUS string

const (
	HighAvailability_State_STATUS_CreatingStandby = HighAvailability_State_STATUS("CreatingStandby")
	HighAvailability_State_STATUS_FailingOver     = HighAvailability_State_STATUS("FailingOver")
	HighAvailability_State_STATUS_Healthy         = HighAvailability_State_STATUS("Healthy")
	HighAvailability_State_STATUS_NotEnabled      = HighAvailability_State_STATUS("NotEnabled")
	HighAvailability_State_STATUS_RemovingStandby = HighAvailability_State_STATUS("RemovingStandby")
	HighAvailability_State_STATUS_ReplicatingData = HighAvailability_State_STATUS("ReplicatingData")
)

// Mapping from string to HighAvailability_State_STATUS
var highAvailability_State_STATUS_Values = map[string]HighAvailability_State_STATUS{
	"creatingstandby": HighAvailability_State_STATUS_CreatingStandby,
	"failingover":     HighAvailability_State_STATUS_FailingOver,
	"healthy":         HighAvailability_State_STATUS_Healthy,
	"notenabled":      HighAvailability_State_STATUS_NotEnabled,
	"removingstandby": HighAvailability_State_STATUS_RemovingStandby,
	"replicatingdata": HighAvailability_State_STATUS_ReplicatingData,
}

type Network_PublicNetworkAccess_STATUS string

const (
	Network_PublicNetworkAccess_STATUS_Disabled = Network_PublicNetworkAccess_STATUS("Disabled")
	Network_PublicNetworkAccess_STATUS_Enabled  = Network_PublicNetworkAccess_STATUS("Enabled")
)

// Mapping from string to Network_PublicNetworkAccess_STATUS
var network_PublicNetworkAccess_STATUS_Values = map[string]Network_PublicNetworkAccess_STATUS{
	"disabled": Network_PublicNetworkAccess_STATUS_Disabled,
	"enabled":  Network_PublicNetworkAccess_STATUS_Enabled,
}

type Replica_PromoteMode_STATUS string

const (
	Replica_PromoteMode_STATUS_Standalone = Replica_PromoteMode_STATUS("standalone")
	Replica_PromoteMode_STATUS_Switchover = Replica_PromoteMode_STATUS("switchover")
)

// Mapping from string to Replica_PromoteMode_STATUS
var replica_PromoteMode_STATUS_Values = map[string]Replica_PromoteMode_STATUS{
	"standalone": Replica_PromoteMode_STATUS_Standalone,
	"switchover": Replica_PromoteMode_STATUS_Switchover,
}

type Replica_PromoteOption_STATUS string

const (
	Replica_PromoteOption_STATUS_Forced  = Replica_PromoteOption_STATUS("forced")
	Replica_PromoteOption_STATUS_Planned = Replica_PromoteOption_STATUS("planned")
)

// Mapping from string to Replica_PromoteOption_STATUS
var replica_PromoteOption_STATUS_Values = map[string]Replica_PromoteOption_STATUS{
	"forced":  Replica_PromoteOption_STATUS_Forced,
	"planned": Replica_PromoteOption_STATUS_Planned,
}

type Replica_ReplicationState_STATUS string

const (
	Replica_ReplicationState_STATUS_Active        = Replica_ReplicationState_STATUS("Active")
	Replica_ReplicationState_STATUS_Broken        = Replica_ReplicationState_STATUS("Broken")
	Replica_ReplicationState_STATUS_Catchup       = Replica_ReplicationState_STATUS("Catchup")
	Replica_ReplicationState_STATUS_Provisioning  = Replica_ReplicationState_STATUS("Provisioning")
	Replica_ReplicationState_STATUS_Reconfiguring = Replica_ReplicationState_STATUS("Reconfiguring")
	Replica_ReplicationState_STATUS_Updating      = Replica_ReplicationState_STATUS("Updating")
)

// Mapping from string to Replica_ReplicationState_STATUS
var replica_ReplicationState_STATUS_Values = map[string]Replica_ReplicationState_STATUS{
	"active":        Replica_ReplicationState_STATUS_Active,
	"broken":        Replica_ReplicationState_STATUS_Broken,
	"catchup":       Replica_ReplicationState_STATUS_Catchup,
	"provisioning":  Replica_ReplicationState_STATUS_Provisioning,
	"reconfiguring": Replica_ReplicationState_STATUS_Reconfiguring,
	"updating":      Replica_ReplicationState_STATUS_Updating,
}

type Storage_AutoGrow_STATUS string

const (
	Storage_AutoGrow_STATUS_Disabled = Storage_AutoGrow_STATUS("Disabled")
	Storage_AutoGrow_STATUS_Enabled  = Storage_AutoGrow_STATUS("Enabled")
)

// Mapping from string to Storage_AutoGrow_STATUS
var storage_AutoGrow_STATUS_Values = map[string]Storage_AutoGrow_STATUS{
	"disabled": Storage_AutoGrow_STATUS_Disabled,
	"enabled":  Storage_AutoGrow_STATUS_Enabled,
}

type Storage_Tier_STATUS string

const (
	Storage_Tier_STATUS_P1  = Storage_Tier_STATUS("P1")
	Storage_Tier_STATUS_P10 = Storage_Tier_STATUS("P10")
	Storage_Tier_STATUS_P15 = Storage_Tier_STATUS("P15")
	Storage_Tier_STATUS_P2  = Storage_Tier_STATUS("P2")
	Storage_Tier_STATUS_P20 = Storage_Tier_STATUS("P20")
	Storage_Tier_STATUS_P3  = Storage_Tier_STATUS("P3")
	Storage_Tier_STATUS_P30 = Storage_Tier_STATUS("P30")
	Storage_Tier_STATUS_P4  = Storage_Tier_STATUS("P4")
	Storage_Tier_STATUS_P40 = Storage_Tier_STATUS("P40")
	Storage_Tier_STATUS_P50 = Storage_Tier_STATUS("P50")
	Storage_Tier_STATUS_P6  = Storage_Tier_STATUS("P6")
	Storage_Tier_STATUS_P60 = Storage_Tier_STATUS("P60")
	Storage_Tier_STATUS_P70 = Storage_Tier_STATUS("P70")
	Storage_Tier_STATUS_P80 = Storage_Tier_STATUS("P80")
)

// Mapping from string to Storage_Tier_STATUS
var storage_Tier_STATUS_Values = map[string]Storage_Tier_STATUS{
	"p1":  Storage_Tier_STATUS_P1,
	"p10": Storage_Tier_STATUS_P10,
	"p15": Storage_Tier_STATUS_P15,
	"p2":  Storage_Tier_STATUS_P2,
	"p20": Storage_Tier_STATUS_P20,
	"p3":  Storage_Tier_STATUS_P3,
	"p30": Storage_Tier_STATUS_P30,
	"p4":  Storage_Tier_STATUS_P4,
	"p40": Storage_Tier_STATUS_P40,
	"p50": Storage_Tier_STATUS_P50,
	"p6":  Storage_Tier_STATUS_P6,
	"p60": Storage_Tier_STATUS_P60,
	"p70": Storage_Tier_STATUS_P70,
	"p80": Storage_Tier_STATUS_P80,
}

type Storage_Type_STATUS string

const (
	Storage_Type_STATUS_PremiumV2_LRS = Storage_Type_STATUS("PremiumV2_LRS")
	Storage_Type_STATUS_Premium_LRS   = Storage_Type_STATUS("Premium_LRS")
)

// Mapping from string to Storage_Type_STATUS
var storage_Type_STATUS_Values = map[string]Storage_Type_STATUS{
	"premiumv2_lrs": Storage_Type_STATUS_PremiumV2_LRS,
	"premium_lrs":   Storage_Type_STATUS_Premium_LRS,
}