// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type ManagedClustersAgentPool_Spec struct {
	Name string `json:"name,omitempty"`

	// Properties: Properties of an agent pool.
	Properties *ManagedClusterAgentPoolProfileProperties `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &ManagedClustersAgentPool_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2024-09-01"
func (pool ManagedClustersAgentPool_Spec) GetAPIVersion() string {
	return "2024-09-01"
}

// GetName returns the Name of the resource
func (pool *ManagedClustersAgentPool_Spec) GetName() string {
	return pool.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ContainerService/managedClusters/agentPools"
func (pool *ManagedClustersAgentPool_Spec) GetType() string {
	return "Microsoft.ContainerService/managedClusters/agentPools"
}

// Properties for the container service agent pool profile.
type ManagedClusterAgentPoolProfileProperties struct {
	// AvailabilityZones: The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType
	// property is 'VirtualMachineScaleSets'.
	AvailabilityZones          []string `json:"availabilityZones"`
	CapacityReservationGroupID *string  `json:"capacityReservationGroupID,omitempty"`

	// Count: Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 1000 (inclusive)
	// for user pools and in the range of 1 to 1000 (inclusive) for system pools. The default value is 1.
	Count *int `json:"count,omitempty"`

	// CreationData: CreationData to be used to specify the source Snapshot ID if the node pool will be created/upgraded using
	// a snapshot.
	CreationData *CreationData `json:"creationData,omitempty"`

	// EnableAutoScaling: Whether to enable auto-scaler
	EnableAutoScaling *bool `json:"enableAutoScaling,omitempty"`

	// EnableEncryptionAtHost: This is only supported on certain VM sizes and in certain Azure regions. For more information,
	// see: https://docs.microsoft.com/azure/aks/enable-host-encryption
	EnableEncryptionAtHost *bool `json:"enableEncryptionAtHost,omitempty"`

	// EnableFIPS: See [Add a FIPS-enabled node
	// pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#add-a-fips-enabled-node-pool-preview) for more
	// details.
	EnableFIPS *bool `json:"enableFIPS,omitempty"`

	// EnableNodePublicIP: Some scenarios may require nodes in a node pool to receive their own dedicated public IP addresses.
	// A common scenario is for gaming workloads, where a console needs to make a direct connection to a cloud virtual machine
	// to minimize hops. For more information see [assigning a public IP per
	// node](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#assign-a-public-ip-per-node-for-your-node-pools). The
	// default is false.
	EnableNodePublicIP *bool `json:"enableNodePublicIP,omitempty"`

	// EnableUltraSSD: Whether to enable UltraSSD
	EnableUltraSSD *bool `json:"enableUltraSSD,omitempty"`

	// GpuInstanceProfile: GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.
	GpuInstanceProfile *GPUInstanceProfile `json:"gpuInstanceProfile,omitempty"`
	HostGroupID        *string             `json:"hostGroupID,omitempty"`

	// KubeletConfig: The Kubelet configuration on the agent pool nodes.
	KubeletConfig *KubeletConfig `json:"kubeletConfig,omitempty"`

	// KubeletDiskType: Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral
	// storage.
	KubeletDiskType *KubeletDiskType `json:"kubeletDiskType,omitempty"`

	// LinuxOSConfig: The OS configuration of Linux agent nodes.
	LinuxOSConfig *LinuxOSConfig `json:"linuxOSConfig,omitempty"`

	// MaxCount: The maximum number of nodes for auto-scaling
	MaxCount *int `json:"maxCount,omitempty"`

	// MaxPods: The maximum number of pods that can run on a node.
	MaxPods *int `json:"maxPods,omitempty"`

	// MinCount: The minimum number of nodes for auto-scaling
	MinCount *int `json:"minCount,omitempty"`

	// Mode: A cluster must have at least one 'System' Agent Pool at all times. For additional information on agent pool
	// restrictions  and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools
	Mode *AgentPoolMode `json:"mode,omitempty"`

	// NetworkProfile: Network-related settings of an agent pool.
	NetworkProfile *AgentPoolNetworkProfile `json:"networkProfile,omitempty"`

	// NodeLabels: The node labels to be persisted across all nodes in agent pool.
	NodeLabels           map[string]string `json:"nodeLabels" serializationType:"explicitEmptyCollection"`
	NodePublicIPPrefixID *string           `json:"nodePublicIPPrefixID,omitempty"`

	// NodeTaints: The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.
	NodeTaints []string `json:"nodeTaints" serializationType:"explicitEmptyCollection"`

	// OrchestratorVersion: Both patch version <major.minor.patch> (e.g. 1.20.13) and <major.minor> (e.g. 1.20) are supported.
	// When <major.minor> is specified, the latest supported GA patch version is chosen automatically. Updating the cluster
	// with the same <major.minor> once it has been created (e.g. 1.14.x -> 1.14) will not trigger an upgrade, even if a newer
	// patch version is available. As a best practice, you should upgrade all node pools in an AKS cluster to the same
	// Kubernetes version. The node pool version must have the same major version as the control plane. The node pool minor
	// version must be within two minor versions of the control plane version. The node pool version cannot be greater than the
	// control plane version. For more information see [upgrading a node
	// pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool).
	OrchestratorVersion *string `json:"orchestratorVersion,omitempty"`
	OsDiskSizeGB        *int    `json:"osDiskSizeGB,omitempty"`

	// OsDiskType: The default is 'Ephemeral' if the VM supports it and has a cache disk larger than the requested
	// OSDiskSizeGB. Otherwise,  defaults to 'Managed'. May not be changed after creation. For more information see [Ephemeral
	// OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).
	OsDiskType *OSDiskType `json:"osDiskType,omitempty"`

	// OsSKU: Specifies the OS SKU used by the agent pool. The default is Ubuntu if OSType is Linux. The default is Windows2019
	// when  Kubernetes <= 1.24 or Windows2022 when Kubernetes >= 1.25 if OSType is Windows.
	OsSKU *OSSKU `json:"osSKU,omitempty"`

	// OsType: The operating system type. The default is Linux.
	OsType      *OSType `json:"osType,omitempty"`
	PodSubnetID *string `json:"podSubnetID,omitempty"`

	// PowerState: When an Agent Pool is first created it is initially Running. The Agent Pool can be stopped by setting this
	// field to Stopped. A stopped Agent Pool stops all of its VMs and does not accrue billing charges. An Agent Pool can only
	// be stopped if it is Running and provisioning state is Succeeded
	PowerState                *PowerState `json:"powerState,omitempty"`
	ProximityPlacementGroupID *string     `json:"proximityPlacementGroupID,omitempty"`

	// ScaleDownMode: This also effects the cluster autoscaler behavior. If not specified, it defaults to Delete.
	ScaleDownMode *ScaleDownMode `json:"scaleDownMode,omitempty"`

	// ScaleSetEvictionPolicy: This cannot be specified unless the scaleSetPriority is 'Spot'. If not specified, the default is
	// 'Delete'.
	ScaleSetEvictionPolicy *ScaleSetEvictionPolicy `json:"scaleSetEvictionPolicy,omitempty"`

	// ScaleSetPriority: The Virtual Machine Scale Set priority. If not specified, the default is 'Regular'.
	ScaleSetPriority *ScaleSetPriority `json:"scaleSetPriority,omitempty"`

	// SecurityProfile: The security settings of an agent pool.
	SecurityProfile *AgentPoolSecurityProfile `json:"securityProfile,omitempty"`

	// SpotMaxPrice: Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any
	// on-demand price. For more details on spot pricing, see [spot VMs
	// pricing](https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing)
	SpotMaxPrice *float64 `json:"spotMaxPrice,omitempty"`

	// Tags: The tags to be persisted on the agent pool virtual machine scale set.
	Tags map[string]string `json:"tags" serializationType:"explicitEmptyCollection"`

	// Type: The type of Agent Pool.
	Type *AgentPoolType `json:"type,omitempty"`

	// UpgradeSettings: Settings for upgrading the agentpool
	UpgradeSettings *AgentPoolUpgradeSettings `json:"upgradeSettings,omitempty"`

	// VmSize: VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods
	// might fail to run correctly. For more details on restricted VM sizes, see:
	// https://docs.microsoft.com/azure/aks/quotas-skus-regions
	VmSize       *string `json:"vmSize,omitempty"`
	VnetSubnetID *string `json:"vnetSubnetID,omitempty"`

	// WindowsProfile: The Windows agent pool's specific profile.
	WindowsProfile *AgentPoolWindowsProfile `json:"windowsProfile,omitempty"`

	// WorkloadRuntime: Determines the type of workload a node can run.
	WorkloadRuntime *WorkloadRuntime `json:"workloadRuntime,omitempty"`
}

// A cluster must have at least one 'System' Agent Pool at all times. For additional information on agent pool restrictions
// and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools
// +kubebuilder:validation:Enum={"System","User"}
type AgentPoolMode string

const (
	AgentPoolMode_System = AgentPoolMode("System")
	AgentPoolMode_User   = AgentPoolMode("User")
)

// Mapping from string to AgentPoolMode
var agentPoolMode_Values = map[string]AgentPoolMode{
	"system": AgentPoolMode_System,
	"user":   AgentPoolMode_User,
}

// Network settings of an agent pool.
type AgentPoolNetworkProfile struct {
	// AllowedHostPorts: The port ranges that are allowed to access. The specified ranges are allowed to overlap.
	AllowedHostPorts          []PortRange `json:"allowedHostPorts"`
	ApplicationSecurityGroups []string    `json:"applicationSecurityGroups,omitempty"`

	// NodePublicIPTags: IPTags of instance-level public IPs.
	NodePublicIPTags []IPTag `json:"nodePublicIPTags"`
}

// The security settings of an agent pool.
type AgentPoolSecurityProfile struct {
	// EnableSecureBoot: Secure Boot is a feature of Trusted Launch which ensures that only signed operating systems and
	// drivers can boot. For more details, see aka.ms/aks/trustedlaunch.  If not specified, the default is false.
	EnableSecureBoot *bool `json:"enableSecureBoot,omitempty"`

	// EnableVTPM: vTPM is a Trusted Launch feature for configuring a dedicated secure vault for keys and measurements held
	// locally on the node. For more details, see aka.ms/aks/trustedlaunch. If not specified, the default is false.
	EnableVTPM *bool `json:"enableVTPM,omitempty"`
}

// The type of Agent Pool.
// +kubebuilder:validation:Enum={"AvailabilitySet","VirtualMachineScaleSets"}
type AgentPoolType string

const (
	AgentPoolType_AvailabilitySet         = AgentPoolType("AvailabilitySet")
	AgentPoolType_VirtualMachineScaleSets = AgentPoolType("VirtualMachineScaleSets")
)

// Mapping from string to AgentPoolType
var agentPoolType_Values = map[string]AgentPoolType{
	"availabilityset":         AgentPoolType_AvailabilitySet,
	"virtualmachinescalesets": AgentPoolType_VirtualMachineScaleSets,
}

// Settings for upgrading an agentpool
type AgentPoolUpgradeSettings struct {
	// DrainTimeoutInMinutes: The amount of time (in minutes) to wait on eviction of pods and graceful termination per node.
	// This eviction wait time honors waiting on pod disruption budgets. If this time is exceeded, the upgrade fails. If not
	// specified, the default is 30 minutes.
	DrainTimeoutInMinutes *int `json:"drainTimeoutInMinutes,omitempty"`

	// MaxSurge: This can either be set to an integer (e.g. '5') or a percentage (e.g. '50%'). If a percentage is specified, it
	// is the percentage of the total agent pool size at the time of the upgrade. For percentages, fractional nodes are rounded
	// up. If not specified, the default is 1. For more information, including best practices, see:
	// https://docs.microsoft.com/azure/aks/upgrade-cluster#customize-node-surge-upgrade
	MaxSurge *string `json:"maxSurge,omitempty"`

	// NodeSoakDurationInMinutes: The amount of time (in minutes) to wait after draining a node and before reimaging it and
	// moving on to next node. If not specified, the default is 0 minutes.
	NodeSoakDurationInMinutes *int `json:"nodeSoakDurationInMinutes,omitempty"`
}

// The Windows agent pool's specific profile.
type AgentPoolWindowsProfile struct {
	// DisableOutboundNat: The default value is false. Outbound NAT can only be disabled if the cluster outboundType is NAT
	// Gateway and the Windows agent pool does not have node public IP enabled.
	DisableOutboundNat *bool `json:"disableOutboundNat,omitempty"`
}

// Data used when creating a target resource from a source resource.
type CreationData struct {
	SourceResourceId *string `json:"sourceResourceId,omitempty"`
}

// GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.
// +kubebuilder:validation:Enum={"MIG1g","MIG2g","MIG3g","MIG4g","MIG7g"}
type GPUInstanceProfile string

const (
	GPUInstanceProfile_MIG1G = GPUInstanceProfile("MIG1g")
	GPUInstanceProfile_MIG2G = GPUInstanceProfile("MIG2g")
	GPUInstanceProfile_MIG3G = GPUInstanceProfile("MIG3g")
	GPUInstanceProfile_MIG4G = GPUInstanceProfile("MIG4g")
	GPUInstanceProfile_MIG7G = GPUInstanceProfile("MIG7g")
)

// Mapping from string to GPUInstanceProfile
var gPUInstanceProfile_Values = map[string]GPUInstanceProfile{
	"mig1g": GPUInstanceProfile_MIG1G,
	"mig2g": GPUInstanceProfile_MIG2G,
	"mig3g": GPUInstanceProfile_MIG3G,
	"mig4g": GPUInstanceProfile_MIG4G,
	"mig7g": GPUInstanceProfile_MIG7G,
}

// See [AKS custom node configuration](https://docs.microsoft.com/azure/aks/custom-node-configuration) for more details.
type KubeletConfig struct {
	// AllowedUnsafeSysctls: Allowed list of unsafe sysctls or unsafe sysctl patterns (ending in `*`).
	AllowedUnsafeSysctls []string `json:"allowedUnsafeSysctls"`

	// ContainerLogMaxFiles: The maximum number of container log files that can be present for a container. The number must be
	// ≥ 2.
	ContainerLogMaxFiles *int `json:"containerLogMaxFiles,omitempty"`

	// ContainerLogMaxSizeMB: The maximum size (e.g. 10Mi) of container log file before it is rotated.
	ContainerLogMaxSizeMB *int `json:"containerLogMaxSizeMB,omitempty"`

	// CpuCfsQuota: The default is true.
	CpuCfsQuota *bool `json:"cpuCfsQuota,omitempty"`

	// CpuCfsQuotaPeriod: The default is '100ms.' Valid values are a sequence of decimal numbers with an optional fraction and
	// a unit suffix. For example: '300ms', '2h45m'. Supported units are 'ns', 'us', 'ms', 's', 'm', and 'h'.
	CpuCfsQuotaPeriod *string `json:"cpuCfsQuotaPeriod,omitempty"`

	// CpuManagerPolicy: The default is 'none'. See [Kubernetes CPU management
	// policies](https://kubernetes.io/docs/tasks/administer-cluster/cpu-management-policies/#cpu-management-policies) for more
	// information. Allowed values are 'none' and 'static'.
	CpuManagerPolicy *string `json:"cpuManagerPolicy,omitempty"`

	// FailSwapOn: If set to true it will make the Kubelet fail to start if swap is enabled on the node.
	FailSwapOn *bool `json:"failSwapOn,omitempty"`

	// ImageGcHighThreshold: To disable image garbage collection, set to 100. The default is 85%
	ImageGcHighThreshold *int `json:"imageGcHighThreshold,omitempty"`

	// ImageGcLowThreshold: This cannot be set higher than imageGcHighThreshold. The default is 80%
	ImageGcLowThreshold *int `json:"imageGcLowThreshold,omitempty"`

	// PodMaxPids: The maximum number of processes per pod.
	PodMaxPids *int `json:"podMaxPids,omitempty"`

	// TopologyManagerPolicy: For more information see [Kubernetes Topology
	// Manager](https://kubernetes.io/docs/tasks/administer-cluster/topology-manager). The default is 'none'. Allowed values
	// are 'none', 'best-effort', 'restricted', and 'single-numa-node'.
	TopologyManagerPolicy *string `json:"topologyManagerPolicy,omitempty"`
}

// Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.
// +kubebuilder:validation:Enum={"OS","Temporary"}
type KubeletDiskType string

const (
	KubeletDiskType_OS        = KubeletDiskType("OS")
	KubeletDiskType_Temporary = KubeletDiskType("Temporary")
)

// Mapping from string to KubeletDiskType
var kubeletDiskType_Values = map[string]KubeletDiskType{
	"os":        KubeletDiskType_OS,
	"temporary": KubeletDiskType_Temporary,
}

// See [AKS custom node configuration](https://docs.microsoft.com/azure/aks/custom-node-configuration) for more details.
type LinuxOSConfig struct {
	// SwapFileSizeMB: The size in MB of a swap file that will be created on each node.
	SwapFileSizeMB *int `json:"swapFileSizeMB,omitempty"`

	// Sysctls: Sysctl settings for Linux agent nodes.
	Sysctls *SysctlConfig `json:"sysctls,omitempty"`

	// TransparentHugePageDefrag: Valid values are 'always', 'defer', 'defer+madvise', 'madvise' and 'never'. The default is
	// 'madvise'. For more information see [Transparent
	// Hugepages](https://www.kernel.org/doc/html/latest/admin-guide/mm/transhuge.html#admin-guide-transhuge).
	TransparentHugePageDefrag *string `json:"transparentHugePageDefrag,omitempty"`

	// TransparentHugePageEnabled: Valid values are 'always', 'madvise', and 'never'. The default is 'always'. For more
	// information see [Transparent
	// Hugepages](https://www.kernel.org/doc/html/latest/admin-guide/mm/transhuge.html#admin-guide-transhuge).
	TransparentHugePageEnabled *string `json:"transparentHugePageEnabled,omitempty"`
}

// The default is 'Ephemeral' if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB. Otherwise,
// defaults to 'Managed'. May not be changed after creation. For more information see [Ephemeral
// OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).
// +kubebuilder:validation:Enum={"Ephemeral","Managed"}
type OSDiskType string

const (
	OSDiskType_Ephemeral = OSDiskType("Ephemeral")
	OSDiskType_Managed   = OSDiskType("Managed")
)

// Mapping from string to OSDiskType
var oSDiskType_Values = map[string]OSDiskType{
	"ephemeral": OSDiskType_Ephemeral,
	"managed":   OSDiskType_Managed,
}

// Specifies the OS SKU used by the agent pool. The default is Ubuntu if OSType is Linux. The default is Windows2019 when
// Kubernetes <= 1.24 or Windows2022 when Kubernetes >= 1.25 if OSType is Windows.
// +kubebuilder:validation:Enum={"AzureLinux","CBLMariner","Ubuntu","Windows2019","Windows2022"}
type OSSKU string

const (
	OSSKU_AzureLinux  = OSSKU("AzureLinux")
	OSSKU_CBLMariner  = OSSKU("CBLMariner")
	OSSKU_Ubuntu      = OSSKU("Ubuntu")
	OSSKU_Windows2019 = OSSKU("Windows2019")
	OSSKU_Windows2022 = OSSKU("Windows2022")
)

// Mapping from string to OSSKU
var oSSKU_Values = map[string]OSSKU{
	"azurelinux":  OSSKU_AzureLinux,
	"cblmariner":  OSSKU_CBLMariner,
	"ubuntu":      OSSKU_Ubuntu,
	"windows2019": OSSKU_Windows2019,
	"windows2022": OSSKU_Windows2022,
}

// The operating system type. The default is Linux.
// +kubebuilder:validation:Enum={"Linux","Windows"}
type OSType string

const (
	OSType_Linux   = OSType("Linux")
	OSType_Windows = OSType("Windows")
)

// Mapping from string to OSType
var oSType_Values = map[string]OSType{
	"linux":   OSType_Linux,
	"windows": OSType_Windows,
}

// Describes the Power State of the cluster
type PowerState struct {
	// Code: Tells whether the cluster is Running or Stopped
	Code *PowerState_Code `json:"code,omitempty"`
}

// Describes how VMs are added to or removed from Agent Pools. See [billing
// states](https://docs.microsoft.com/azure/virtual-machines/states-billing).
// +kubebuilder:validation:Enum={"Deallocate","Delete"}
type ScaleDownMode string

const (
	ScaleDownMode_Deallocate = ScaleDownMode("Deallocate")
	ScaleDownMode_Delete     = ScaleDownMode("Delete")
)

// Mapping from string to ScaleDownMode
var scaleDownMode_Values = map[string]ScaleDownMode{
	"deallocate": ScaleDownMode_Deallocate,
	"delete":     ScaleDownMode_Delete,
}

// The eviction policy specifies what to do with the VM when it is evicted. The default is Delete. For more information
// about eviction see [spot VMs](https://docs.microsoft.com/azure/virtual-machines/spot-vms)
// +kubebuilder:validation:Enum={"Deallocate","Delete"}
type ScaleSetEvictionPolicy string

const (
	ScaleSetEvictionPolicy_Deallocate = ScaleSetEvictionPolicy("Deallocate")
	ScaleSetEvictionPolicy_Delete     = ScaleSetEvictionPolicy("Delete")
)

// Mapping from string to ScaleSetEvictionPolicy
var scaleSetEvictionPolicy_Values = map[string]ScaleSetEvictionPolicy{
	"deallocate": ScaleSetEvictionPolicy_Deallocate,
	"delete":     ScaleSetEvictionPolicy_Delete,
}

// The Virtual Machine Scale Set priority.
// +kubebuilder:validation:Enum={"Regular","Spot"}
type ScaleSetPriority string

const (
	ScaleSetPriority_Regular = ScaleSetPriority("Regular")
	ScaleSetPriority_Spot    = ScaleSetPriority("Spot")
)

// Mapping from string to ScaleSetPriority
var scaleSetPriority_Values = map[string]ScaleSetPriority{
	"regular": ScaleSetPriority_Regular,
	"spot":    ScaleSetPriority_Spot,
}

// Determines the type of workload a node can run.
// +kubebuilder:validation:Enum={"OCIContainer","WasmWasi"}
type WorkloadRuntime string

const (
	WorkloadRuntime_OCIContainer = WorkloadRuntime("OCIContainer")
	WorkloadRuntime_WasmWasi     = WorkloadRuntime("WasmWasi")
)

// Mapping from string to WorkloadRuntime
var workloadRuntime_Values = map[string]WorkloadRuntime{
	"ocicontainer": WorkloadRuntime_OCIContainer,
	"wasmwasi":     WorkloadRuntime_WasmWasi,
}

// Contains the IPTag associated with the object.
type IPTag struct {
	// IpTagType: The IP tag type. Example: RoutingPreference.
	IpTagType *string `json:"ipTagType,omitempty"`

	// Tag: The value of the IP tag associated with the public IP. Example: Internet.
	Tag *string `json:"tag,omitempty"`
}

// The port range.
type PortRange struct {
	// PortEnd: The maximum port that is included in the range. It should be ranged from 1 to 65535, and be greater than or
	// equal to portStart.
	PortEnd *int `json:"portEnd,omitempty"`

	// PortStart: The minimum port that is included in the range. It should be ranged from 1 to 65535, and be less than or
	// equal to portEnd.
	PortStart *int `json:"portStart,omitempty"`

	// Protocol: The network protocol of the port.
	Protocol *PortRange_Protocol `json:"protocol,omitempty"`
}

// +kubebuilder:validation:Enum={"Running","Stopped"}
type PowerState_Code string

const (
	PowerState_Code_Running = PowerState_Code("Running")
	PowerState_Code_Stopped = PowerState_Code("Stopped")
)

// Mapping from string to PowerState_Code
var powerState_Code_Values = map[string]PowerState_Code{
	"running": PowerState_Code_Running,
	"stopped": PowerState_Code_Stopped,
}

// Sysctl settings for Linux agent nodes.
type SysctlConfig struct {
	// FsAioMaxNr: Sysctl setting fs.aio-max-nr.
	FsAioMaxNr *int `json:"fsAioMaxNr,omitempty"`

	// FsFileMax: Sysctl setting fs.file-max.
	FsFileMax *int `json:"fsFileMax,omitempty"`

	// FsInotifyMaxUserWatches: Sysctl setting fs.inotify.max_user_watches.
	FsInotifyMaxUserWatches *int `json:"fsInotifyMaxUserWatches,omitempty"`

	// FsNrOpen: Sysctl setting fs.nr_open.
	FsNrOpen *int `json:"fsNrOpen,omitempty"`

	// KernelThreadsMax: Sysctl setting kernel.threads-max.
	KernelThreadsMax *int `json:"kernelThreadsMax,omitempty"`

	// NetCoreNetdevMaxBacklog: Sysctl setting net.core.netdev_max_backlog.
	NetCoreNetdevMaxBacklog *int `json:"netCoreNetdevMaxBacklog,omitempty"`

	// NetCoreOptmemMax: Sysctl setting net.core.optmem_max.
	NetCoreOptmemMax *int `json:"netCoreOptmemMax,omitempty"`

	// NetCoreRmemDefault: Sysctl setting net.core.rmem_default.
	NetCoreRmemDefault *int `json:"netCoreRmemDefault,omitempty"`

	// NetCoreRmemMax: Sysctl setting net.core.rmem_max.
	NetCoreRmemMax *int `json:"netCoreRmemMax,omitempty"`

	// NetCoreSomaxconn: Sysctl setting net.core.somaxconn.
	NetCoreSomaxconn *int `json:"netCoreSomaxconn,omitempty"`

	// NetCoreWmemDefault: Sysctl setting net.core.wmem_default.
	NetCoreWmemDefault *int `json:"netCoreWmemDefault,omitempty"`

	// NetCoreWmemMax: Sysctl setting net.core.wmem_max.
	NetCoreWmemMax *int `json:"netCoreWmemMax,omitempty"`

	// NetIpv4IpLocalPortRange: Sysctl setting net.ipv4.ip_local_port_range.
	NetIpv4IpLocalPortRange *string `json:"netIpv4IpLocalPortRange,omitempty"`

	// NetIpv4NeighDefaultGcThresh1: Sysctl setting net.ipv4.neigh.default.gc_thresh1.
	NetIpv4NeighDefaultGcThresh1 *int `json:"netIpv4NeighDefaultGcThresh1,omitempty"`

	// NetIpv4NeighDefaultGcThresh2: Sysctl setting net.ipv4.neigh.default.gc_thresh2.
	NetIpv4NeighDefaultGcThresh2 *int `json:"netIpv4NeighDefaultGcThresh2,omitempty"`

	// NetIpv4NeighDefaultGcThresh3: Sysctl setting net.ipv4.neigh.default.gc_thresh3.
	NetIpv4NeighDefaultGcThresh3 *int `json:"netIpv4NeighDefaultGcThresh3,omitempty"`

	// NetIpv4TcpFinTimeout: Sysctl setting net.ipv4.tcp_fin_timeout.
	NetIpv4TcpFinTimeout *int `json:"netIpv4TcpFinTimeout,omitempty"`

	// NetIpv4TcpKeepaliveProbes: Sysctl setting net.ipv4.tcp_keepalive_probes.
	NetIpv4TcpKeepaliveProbes *int `json:"netIpv4TcpKeepaliveProbes,omitempty"`

	// NetIpv4TcpKeepaliveTime: Sysctl setting net.ipv4.tcp_keepalive_time.
	NetIpv4TcpKeepaliveTime *int `json:"netIpv4TcpKeepaliveTime,omitempty"`

	// NetIpv4TcpMaxSynBacklog: Sysctl setting net.ipv4.tcp_max_syn_backlog.
	NetIpv4TcpMaxSynBacklog *int `json:"netIpv4TcpMaxSynBacklog,omitempty"`

	// NetIpv4TcpMaxTwBuckets: Sysctl setting net.ipv4.tcp_max_tw_buckets.
	NetIpv4TcpMaxTwBuckets *int `json:"netIpv4TcpMaxTwBuckets,omitempty"`

	// NetIpv4TcpTwReuse: Sysctl setting net.ipv4.tcp_tw_reuse.
	NetIpv4TcpTwReuse *bool `json:"netIpv4TcpTwReuse,omitempty"`

	// NetIpv4TcpkeepaliveIntvl: Sysctl setting net.ipv4.tcp_keepalive_intvl.
	NetIpv4TcpkeepaliveIntvl *int `json:"netIpv4TcpkeepaliveIntvl,omitempty"`

	// NetNetfilterNfConntrackBuckets: Sysctl setting net.netfilter.nf_conntrack_buckets.
	NetNetfilterNfConntrackBuckets *int `json:"netNetfilterNfConntrackBuckets,omitempty"`

	// NetNetfilterNfConntrackMax: Sysctl setting net.netfilter.nf_conntrack_max.
	NetNetfilterNfConntrackMax *int `json:"netNetfilterNfConntrackMax,omitempty"`

	// VmMaxMapCount: Sysctl setting vm.max_map_count.
	VmMaxMapCount *int `json:"vmMaxMapCount,omitempty"`

	// VmSwappiness: Sysctl setting vm.swappiness.
	VmSwappiness *int `json:"vmSwappiness,omitempty"`

	// VmVfsCachePressure: Sysctl setting vm.vfs_cache_pressure.
	VmVfsCachePressure *int `json:"vmVfsCachePressure,omitempty"`
}

// +kubebuilder:validation:Enum={"TCP","UDP"}
type PortRange_Protocol string

const (
	PortRange_Protocol_TCP = PortRange_Protocol("TCP")
	PortRange_Protocol_UDP = PortRange_Protocol("UDP")
)

// Mapping from string to PortRange_Protocol
var portRange_Protocol_Values = map[string]PortRange_Protocol{
	"tcp": PortRange_Protocol_TCP,
	"udp": PortRange_Protocol_UDP,
}
