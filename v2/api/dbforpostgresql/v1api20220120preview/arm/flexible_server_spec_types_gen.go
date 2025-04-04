// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type FlexibleServer_Spec struct {
	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Properties of the server.
	Properties *ServerProperties `json:"properties,omitempty"`

	// Sku: The SKU (pricing tier) of the server.
	Sku *Sku `json:"sku,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &FlexibleServer_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-01-20-preview"
func (server FlexibleServer_Spec) GetAPIVersion() string {
	return "2022-01-20-preview"
}

// GetName returns the Name of the resource
func (server *FlexibleServer_Spec) GetName() string {
	return server.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforPostgreSQL/flexibleServers"
func (server *FlexibleServer_Spec) GetType() string {
	return "Microsoft.DBforPostgreSQL/flexibleServers"
}

// The properties of a server.
type ServerProperties struct {
	// AdministratorLogin: The administrator's login name of a server. Can only be specified when the server is being created
	// (and is required for creation).
	AdministratorLogin *string `json:"administratorLogin,omitempty"`

	// AdministratorLoginPassword: The administrator login password (required for server creation).
	AdministratorLoginPassword *string `json:"administratorLoginPassword,omitempty"`

	// AvailabilityZone: availability zone information of the server.
	AvailabilityZone *string `json:"availabilityZone,omitempty"`

	// Backup: Backup properties of a server.
	Backup *Backup `json:"backup,omitempty"`

	// CreateMode: The mode to create a new PostgreSQL server.
	CreateMode *ServerProperties_CreateMode `json:"createMode,omitempty"`

	// HighAvailability: High availability properties of a server.
	HighAvailability *HighAvailability `json:"highAvailability,omitempty"`

	// MaintenanceWindow: Maintenance window properties of a server.
	MaintenanceWindow *MaintenanceWindow `json:"maintenanceWindow,omitempty"`

	// Network: Network properties of a server.
	Network *Network `json:"network,omitempty"`

	// PointInTimeUTC: Restore point creation time (ISO8601 format), specifying the time to restore from. It's required when
	// 'createMode' is 'PointInTimeRestore'.
	PointInTimeUTC         *string `json:"pointInTimeUTC,omitempty"`
	SourceServerResourceId *string `json:"sourceServerResourceId,omitempty"`

	// Storage: Storage properties of a server.
	Storage *Storage `json:"storage,omitempty"`

	// Version: PostgreSQL Server version.
	Version *ServerVersion `json:"version,omitempty"`
}

// Sku information related properties of a server.
type Sku struct {
	// Name: The name of the sku, typically, tier + family + cores, e.g. Standard_D4s_v3.
	Name *string `json:"name,omitempty"`

	// Tier: The tier of the particular SKU, e.g. Burstable.
	Tier *Sku_Tier `json:"tier,omitempty"`
}

// Backup properties of a server
type Backup struct {
	// BackupRetentionDays: Backup retention days for the server.
	BackupRetentionDays *int `json:"backupRetentionDays,omitempty"`

	// GeoRedundantBackup: A value indicating whether Geo-Redundant backup is enabled on the server.
	GeoRedundantBackup *Backup_GeoRedundantBackup `json:"geoRedundantBackup,omitempty"`
}

// High availability properties of a server
type HighAvailability struct {
	// Mode: The HA mode for the server.
	Mode *HighAvailability_Mode `json:"mode,omitempty"`

	// StandbyAvailabilityZone: availability zone information of the standby.
	StandbyAvailabilityZone *string `json:"standbyAvailabilityZone,omitempty"`
}

// Maintenance window properties of a server.
type MaintenanceWindow struct {
	// CustomWindow: indicates whether custom window is enabled or disabled
	CustomWindow *string `json:"customWindow,omitempty"`

	// DayOfWeek: day of week for maintenance window
	DayOfWeek *int `json:"dayOfWeek,omitempty"`

	// StartHour: start hour for maintenance window
	StartHour *int `json:"startHour,omitempty"`

	// StartMinute: start minute for maintenance window
	StartMinute *int `json:"startMinute,omitempty"`
}

// Network properties of a server
type Network struct {
	DelegatedSubnetResourceId   *string `json:"delegatedSubnetResourceId,omitempty"`
	PrivateDnsZoneArmResourceId *string `json:"privateDnsZoneArmResourceId,omitempty"`
}

// +kubebuilder:validation:Enum={"Create","Default","PointInTimeRestore","Update"}
type ServerProperties_CreateMode string

const (
	ServerProperties_CreateMode_Create             = ServerProperties_CreateMode("Create")
	ServerProperties_CreateMode_Default            = ServerProperties_CreateMode("Default")
	ServerProperties_CreateMode_PointInTimeRestore = ServerProperties_CreateMode("PointInTimeRestore")
	ServerProperties_CreateMode_Update             = ServerProperties_CreateMode("Update")
)

// Mapping from string to ServerProperties_CreateMode
var serverProperties_CreateMode_Values = map[string]ServerProperties_CreateMode{
	"create":             ServerProperties_CreateMode_Create,
	"default":            ServerProperties_CreateMode_Default,
	"pointintimerestore": ServerProperties_CreateMode_PointInTimeRestore,
	"update":             ServerProperties_CreateMode_Update,
}

// The version of a server.
// +kubebuilder:validation:Enum={"11","12","13","14"}
type ServerVersion string

const (
	ServerVersion_11 = ServerVersion("11")
	ServerVersion_12 = ServerVersion("12")
	ServerVersion_13 = ServerVersion("13")
	ServerVersion_14 = ServerVersion("14")
)

// Mapping from string to ServerVersion
var serverVersion_Values = map[string]ServerVersion{
	"11": ServerVersion_11,
	"12": ServerVersion_12,
	"13": ServerVersion_13,
	"14": ServerVersion_14,
}

// +kubebuilder:validation:Enum={"Burstable","GeneralPurpose","MemoryOptimized"}
type Sku_Tier string

const (
	Sku_Tier_Burstable       = Sku_Tier("Burstable")
	Sku_Tier_GeneralPurpose  = Sku_Tier("GeneralPurpose")
	Sku_Tier_MemoryOptimized = Sku_Tier("MemoryOptimized")
)

// Mapping from string to Sku_Tier
var sku_Tier_Values = map[string]Sku_Tier{
	"burstable":       Sku_Tier_Burstable,
	"generalpurpose":  Sku_Tier_GeneralPurpose,
	"memoryoptimized": Sku_Tier_MemoryOptimized,
}

// Storage properties of a server
type Storage struct {
	// StorageSizeGB: Max storage allowed for a server.
	StorageSizeGB *int `json:"storageSizeGB,omitempty"`
}

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type Backup_GeoRedundantBackup string

const (
	Backup_GeoRedundantBackup_Disabled = Backup_GeoRedundantBackup("Disabled")
	Backup_GeoRedundantBackup_Enabled  = Backup_GeoRedundantBackup("Enabled")
)

// Mapping from string to Backup_GeoRedundantBackup
var backup_GeoRedundantBackup_Values = map[string]Backup_GeoRedundantBackup{
	"disabled": Backup_GeoRedundantBackup_Disabled,
	"enabled":  Backup_GeoRedundantBackup_Enabled,
}

// +kubebuilder:validation:Enum={"Disabled","SameZone","ZoneRedundant"}
type HighAvailability_Mode string

const (
	HighAvailability_Mode_Disabled      = HighAvailability_Mode("Disabled")
	HighAvailability_Mode_SameZone      = HighAvailability_Mode("SameZone")
	HighAvailability_Mode_ZoneRedundant = HighAvailability_Mode("ZoneRedundant")
)

// Mapping from string to HighAvailability_Mode
var highAvailability_Mode_Values = map[string]HighAvailability_Mode{
	"disabled":      HighAvailability_Mode_Disabled,
	"samezone":      HighAvailability_Mode_SameZone,
	"zoneredundant": HighAvailability_Mode_ZoneRedundant,
}
