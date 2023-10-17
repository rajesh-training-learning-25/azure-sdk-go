//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerservice

const (
	moduleName    = "armcontainerservice"
	moduleVersion = "v4.4.0-beta.2"
)

// AgentPoolMode - A cluster must have at least one 'System' Agent Pool at all times. For additional information on agent
// pool restrictions and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools
type AgentPoolMode string

const (
	// AgentPoolModeSystem - System agent pools are primarily for hosting critical system pods such as CoreDNS and metrics-server.
	// System agent pools osType must be Linux. System agent pools VM SKU must have at least 2vCPUs and 4GB of memory.
	AgentPoolModeSystem AgentPoolMode = "System"
	// AgentPoolModeUser - User agent pools are primarily for hosting your application pods.
	AgentPoolModeUser AgentPoolMode = "User"
)

// PossibleAgentPoolModeValues returns the possible values for the AgentPoolMode const type.
func PossibleAgentPoolModeValues() []AgentPoolMode {
	return []AgentPoolMode{
		AgentPoolModeSystem,
		AgentPoolModeUser,
	}
}

// AgentPoolType - The type of Agent Pool.
type AgentPoolType string

const (
	// AgentPoolTypeAvailabilitySet - Use of this is strongly discouraged.
	AgentPoolTypeAvailabilitySet AgentPoolType = "AvailabilitySet"
	// AgentPoolTypeVirtualMachineScaleSets - Create an Agent Pool backed by a Virtual Machine Scale Set.
	AgentPoolTypeVirtualMachineScaleSets AgentPoolType = "VirtualMachineScaleSets"
)

// PossibleAgentPoolTypeValues returns the possible values for the AgentPoolType const type.
func PossibleAgentPoolTypeValues() []AgentPoolType {
	return []AgentPoolType{
		AgentPoolTypeAvailabilitySet,
		AgentPoolTypeVirtualMachineScaleSets,
	}
}

// Code - Tells whether the cluster is Running or Stopped
type Code string

const (
	// CodeRunning - The cluster is running.
	CodeRunning Code = "Running"
	// CodeStopped - The cluster is stopped.
	CodeStopped Code = "Stopped"
)

// PossibleCodeValues returns the possible values for the Code const type.
func PossibleCodeValues() []Code {
	return []Code{
		CodeRunning,
		CodeStopped,
	}
}

// ConnectionStatus - The private link service connection status.
type ConnectionStatus string

const (
	ConnectionStatusApproved     ConnectionStatus = "Approved"
	ConnectionStatusDisconnected ConnectionStatus = "Disconnected"
	ConnectionStatusPending      ConnectionStatus = "Pending"
	ConnectionStatusRejected     ConnectionStatus = "Rejected"
)

// PossibleConnectionStatusValues returns the possible values for the ConnectionStatus const type.
func PossibleConnectionStatusValues() []ConnectionStatus {
	return []ConnectionStatus{
		ConnectionStatusApproved,
		ConnectionStatusDisconnected,
		ConnectionStatusPending,
		ConnectionStatusRejected,
	}
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

// Expander - If not specified, the default is 'random'. See expanders [https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#what-are-expanders]
// for more information.
type Expander string

const (
	// ExpanderLeastWaste - Selects the node group that will have the least idle CPU (if tied, unused memory) after scale-up.
	// This is useful when you have different classes of nodes, for example, high CPU or high memory nodes, and only want to expand
	// those when there are pending pods that need a lot of those resources.
	ExpanderLeastWaste Expander = "least-waste"
	// ExpanderMostPods - Selects the node group that would be able to schedule the most pods when scaling up. This is useful
	// when you are using nodeSelector to make sure certain pods land on certain nodes. Note that this won't cause the autoscaler
	// to select bigger nodes vs. smaller, as it can add multiple smaller nodes at once.
	ExpanderMostPods Expander = "most-pods"
	// ExpanderPriority - Selects the node group that has the highest priority assigned by the user. It's configuration is described
	// in more details [here](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/expander/priority/readme.md).
	ExpanderPriority Expander = "priority"
	// ExpanderRandom - Used when you don't have a particular need for the node groups to scale differently.
	ExpanderRandom Expander = "random"
)

// PossibleExpanderValues returns the possible values for the Expander const type.
func PossibleExpanderValues() []Expander {
	return []Expander{
		ExpanderLeastWaste,
		ExpanderMostPods,
		ExpanderPriority,
		ExpanderRandom,
	}
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

type Format string

const (
	// FormatAzure - Return azure auth-provider kubeconfig. This format is deprecated in v1.22 and will be fully removed in v1.26.
	// See: https://aka.ms/k8s/changes-1-26.
	FormatAzure Format = "azure"
	// FormatExec - Return exec format kubeconfig. This format requires kubelogin binary in the path.
	FormatExec Format = "exec"
)

// PossibleFormatValues returns the possible values for the Format const type.
func PossibleFormatValues() []Format {
	return []Format{
		FormatAzure,
		FormatExec,
	}
}

// GPUInstanceProfile - GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.
type GPUInstanceProfile string

const (
	GPUInstanceProfileMIG1G GPUInstanceProfile = "MIG1g"
	GPUInstanceProfileMIG2G GPUInstanceProfile = "MIG2g"
	GPUInstanceProfileMIG3G GPUInstanceProfile = "MIG3g"
	GPUInstanceProfileMIG4G GPUInstanceProfile = "MIG4g"
	GPUInstanceProfileMIG7G GPUInstanceProfile = "MIG7g"
)

// PossibleGPUInstanceProfileValues returns the possible values for the GPUInstanceProfile const type.
func PossibleGPUInstanceProfileValues() []GPUInstanceProfile {
	return []GPUInstanceProfile{
		GPUInstanceProfileMIG1G,
		GPUInstanceProfileMIG2G,
		GPUInstanceProfileMIG3G,
		GPUInstanceProfileMIG4G,
		GPUInstanceProfileMIG7G,
	}
}

// IPFamily - The IP version to use for cluster networking and IP assignment.
type IPFamily string

const (
	IPFamilyIPv4 IPFamily = "IPv4"
	IPFamilyIPv6 IPFamily = "IPv6"
)

// PossibleIPFamilyValues returns the possible values for the IPFamily const type.
func PossibleIPFamilyValues() []IPFamily {
	return []IPFamily{
		IPFamilyIPv4,
		IPFamilyIPv6,
	}
}

// IstioIngressGatewayMode - Mode of an ingress gateway.
type IstioIngressGatewayMode string

const (
	// IstioIngressGatewayModeExternal - The ingress gateway is assigned a public IP address and is publicly accessible.
	IstioIngressGatewayModeExternal IstioIngressGatewayMode = "External"
	// IstioIngressGatewayModeInternal - The ingress gateway is assigned an internal IP address and cannot is accessed publicly.
	IstioIngressGatewayModeInternal IstioIngressGatewayMode = "Internal"
)

// PossibleIstioIngressGatewayModeValues returns the possible values for the IstioIngressGatewayMode const type.
func PossibleIstioIngressGatewayModeValues() []IstioIngressGatewayMode {
	return []IstioIngressGatewayMode{
		IstioIngressGatewayModeExternal,
		IstioIngressGatewayModeInternal,
	}
}

// KeyVaultNetworkAccessTypes - Network access of key vault. The possible values are Public and Private. Public means the
// key vault allows public access from all networks. Private means the key vault disables public access and
// enables private link. The default value is Public.
type KeyVaultNetworkAccessTypes string

const (
	KeyVaultNetworkAccessTypesPrivate KeyVaultNetworkAccessTypes = "Private"
	KeyVaultNetworkAccessTypesPublic  KeyVaultNetworkAccessTypes = "Public"
)

// PossibleKeyVaultNetworkAccessTypesValues returns the possible values for the KeyVaultNetworkAccessTypes const type.
func PossibleKeyVaultNetworkAccessTypesValues() []KeyVaultNetworkAccessTypes {
	return []KeyVaultNetworkAccessTypes{
		KeyVaultNetworkAccessTypesPrivate,
		KeyVaultNetworkAccessTypesPublic,
	}
}

// KubeletDiskType - Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.
type KubeletDiskType string

const (
	// KubeletDiskTypeOS - Kubelet will use the OS disk for its data.
	KubeletDiskTypeOS KubeletDiskType = "OS"
	// KubeletDiskTypeTemporary - Kubelet will use the temporary disk for its data.
	KubeletDiskTypeTemporary KubeletDiskType = "Temporary"
)

// PossibleKubeletDiskTypeValues returns the possible values for the KubeletDiskType const type.
func PossibleKubeletDiskTypeValues() []KubeletDiskType {
	return []KubeletDiskType{
		KubeletDiskTypeOS,
		KubeletDiskTypeTemporary,
	}
}

// KubernetesSupportPlan - Different support tiers for AKS managed clusters
type KubernetesSupportPlan string

const (
	// KubernetesSupportPlanAKSLongTermSupport - Support for the version extended past the KubernetesOfficial support of 1 year.
	// AKS continues to patch CVEs for another 1 year, for a total of 2 years of support.
	KubernetesSupportPlanAKSLongTermSupport KubernetesSupportPlan = "AKSLongTermSupport"
	// KubernetesSupportPlanKubernetesOfficial - Support for the version is the same as for the open source Kubernetes offering.
	// Official Kubernetes open source community support versions for 1 year after release.
	KubernetesSupportPlanKubernetesOfficial KubernetesSupportPlan = "KubernetesOfficial"
)

// PossibleKubernetesSupportPlanValues returns the possible values for the KubernetesSupportPlan const type.
func PossibleKubernetesSupportPlanValues() []KubernetesSupportPlan {
	return []KubernetesSupportPlan{
		KubernetesSupportPlanAKSLongTermSupport,
		KubernetesSupportPlanKubernetesOfficial,
	}
}

// LicenseType - The license type to use for Windows VMs. See Azure Hybrid User Benefits [https://azure.microsoft.com/pricing/hybrid-benefit/faq/]
// for more details.
type LicenseType string

const (
	// LicenseTypeNone - No additional licensing is applied.
	LicenseTypeNone LicenseType = "None"
	// LicenseTypeWindowsServer - Enables Azure Hybrid User Benefits for Windows VMs.
	LicenseTypeWindowsServer LicenseType = "Windows_Server"
)

// PossibleLicenseTypeValues returns the possible values for the LicenseType const type.
func PossibleLicenseTypeValues() []LicenseType {
	return []LicenseType{
		LicenseTypeNone,
		LicenseTypeWindowsServer,
	}
}

// LoadBalancerSKU - The default is 'standard'. See Azure Load Balancer SKUs [https://docs.microsoft.com/azure/load-balancer/skus]
// for more information about the differences between load balancer SKUs.
type LoadBalancerSKU string

const (
	// LoadBalancerSKUBasic - Use a basic Load Balancer with limited functionality.
	LoadBalancerSKUBasic LoadBalancerSKU = "basic"
	// LoadBalancerSKUStandard - Use a a standard Load Balancer. This is the recommended Load Balancer SKU. For more information
	// about on working with the load balancer in the managed cluster, see the [standard Load Balancer](https://docs.microsoft.com/azure/aks/load-balancer-standard)
	// article.
	LoadBalancerSKUStandard LoadBalancerSKU = "standard"
)

// PossibleLoadBalancerSKUValues returns the possible values for the LoadBalancerSKU const type.
func PossibleLoadBalancerSKUValues() []LoadBalancerSKU {
	return []LoadBalancerSKU{
		LoadBalancerSKUBasic,
		LoadBalancerSKUStandard,
	}
}

// ManagedClusterPodIdentityProvisioningState - The current provisioning state of the pod identity.
type ManagedClusterPodIdentityProvisioningState string

const (
	ManagedClusterPodIdentityProvisioningStateAssigned  ManagedClusterPodIdentityProvisioningState = "Assigned"
	ManagedClusterPodIdentityProvisioningStateCanceled  ManagedClusterPodIdentityProvisioningState = "Canceled"
	ManagedClusterPodIdentityProvisioningStateDeleting  ManagedClusterPodIdentityProvisioningState = "Deleting"
	ManagedClusterPodIdentityProvisioningStateFailed    ManagedClusterPodIdentityProvisioningState = "Failed"
	ManagedClusterPodIdentityProvisioningStateSucceeded ManagedClusterPodIdentityProvisioningState = "Succeeded"
	ManagedClusterPodIdentityProvisioningStateUpdating  ManagedClusterPodIdentityProvisioningState = "Updating"
)

// PossibleManagedClusterPodIdentityProvisioningStateValues returns the possible values for the ManagedClusterPodIdentityProvisioningState const type.
func PossibleManagedClusterPodIdentityProvisioningStateValues() []ManagedClusterPodIdentityProvisioningState {
	return []ManagedClusterPodIdentityProvisioningState{
		ManagedClusterPodIdentityProvisioningStateAssigned,
		ManagedClusterPodIdentityProvisioningStateCanceled,
		ManagedClusterPodIdentityProvisioningStateDeleting,
		ManagedClusterPodIdentityProvisioningStateFailed,
		ManagedClusterPodIdentityProvisioningStateSucceeded,
		ManagedClusterPodIdentityProvisioningStateUpdating,
	}
}

// ManagedClusterSKUName - The name of a managed cluster SKU.
type ManagedClusterSKUName string

const (
	// ManagedClusterSKUNameBase - Base option for the AKS control plane.
	ManagedClusterSKUNameBase ManagedClusterSKUName = "Base"
)

// PossibleManagedClusterSKUNameValues returns the possible values for the ManagedClusterSKUName const type.
func PossibleManagedClusterSKUNameValues() []ManagedClusterSKUName {
	return []ManagedClusterSKUName{
		ManagedClusterSKUNameBase,
	}
}

// ManagedClusterSKUTier - If not specified, the default is 'Free'. See AKS Pricing Tier [https://learn.microsoft.com/azure/aks/free-standard-pricing-tiers]
// for more details.
type ManagedClusterSKUTier string

const (
	// ManagedClusterSKUTierFree - The cluster management is free, but charged for VM, storage, and networking usage. Best for
	// experimenting, learning, simple testing, or workloads with fewer than 10 nodes. Not recommended for production use cases.
	ManagedClusterSKUTierFree ManagedClusterSKUTier = "Free"
	// ManagedClusterSKUTierPremium - Cluster has premium capabilities in addition to all of the capabilities included in 'Standard'.
	// Premium enables selection of LongTermSupport (aka.ms/aks/lts) for certain Kubernetes versions.
	ManagedClusterSKUTierPremium ManagedClusterSKUTier = "Premium"
	// ManagedClusterSKUTierStandard - Recommended for mission-critical and production workloads. Includes Kubernetes control
	// plane autoscaling, workload-intensive testing, and up to 5,000 nodes per cluster. Guarantees 99.95% availability of the
	// Kubernetes API server endpoint for clusters that use Availability Zones and 99.9% of availability for clusters that don't
	// use Availability Zones.
	ManagedClusterSKUTierStandard ManagedClusterSKUTier = "Standard"
)

// PossibleManagedClusterSKUTierValues returns the possible values for the ManagedClusterSKUTier const type.
func PossibleManagedClusterSKUTierValues() []ManagedClusterSKUTier {
	return []ManagedClusterSKUTier{
		ManagedClusterSKUTierFree,
		ManagedClusterSKUTierPremium,
		ManagedClusterSKUTierStandard,
	}
}

// NetworkDataplane - Network dataplane used in the Kubernetes cluster.
type NetworkDataplane string

const (
	// NetworkDataplaneAzure - Use Azure network dataplane.
	NetworkDataplaneAzure NetworkDataplane = "azure"
	// NetworkDataplaneCilium - Use Cilium network dataplane. See [Azure CNI Powered by Cilium](https://learn.microsoft.com/azure/aks/azure-cni-powered-by-cilium)
	// for more information.
	NetworkDataplaneCilium NetworkDataplane = "cilium"
)

// PossibleNetworkDataplaneValues returns the possible values for the NetworkDataplane const type.
func PossibleNetworkDataplaneValues() []NetworkDataplane {
	return []NetworkDataplane{
		NetworkDataplaneAzure,
		NetworkDataplaneCilium,
	}
}

// NetworkMode - This cannot be specified if networkPlugin is anything other than 'azure'.
type NetworkMode string

const (
	// NetworkModeBridge - This is no longer supported
	NetworkModeBridge NetworkMode = "bridge"
	// NetworkModeTransparent - No bridge is created. Intra-VM Pod to Pod communication is through IP routes created by Azure
	// CNI. See [Transparent Mode](https://docs.microsoft.com/azure/aks/faq#transparent-mode) for more information.
	NetworkModeTransparent NetworkMode = "transparent"
)

// PossibleNetworkModeValues returns the possible values for the NetworkMode const type.
func PossibleNetworkModeValues() []NetworkMode {
	return []NetworkMode{
		NetworkModeBridge,
		NetworkModeTransparent,
	}
}

// NetworkPlugin - Network plugin used for building the Kubernetes network.
type NetworkPlugin string

const (
	// NetworkPluginAzure - Use the Azure CNI network plugin. See [Azure CNI (advanced) networking](https://docs.microsoft.com/azure/aks/concepts-network#azure-cni-advanced-networking)
	// for more information.
	NetworkPluginAzure NetworkPlugin = "azure"
	// NetworkPluginKubenet - Use the Kubenet network plugin. See [Kubenet (basic) networking](https://docs.microsoft.com/azure/aks/concepts-network#kubenet-basic-networking)
	// for more information.
	NetworkPluginKubenet NetworkPlugin = "kubenet"
	// NetworkPluginNone - No CNI plugin is pre-installed. See [BYO CNI](https://docs.microsoft.com/en-us/azure/aks/use-byo-cni)
	// for more information.
	NetworkPluginNone NetworkPlugin = "none"
)

// PossibleNetworkPluginValues returns the possible values for the NetworkPlugin const type.
func PossibleNetworkPluginValues() []NetworkPlugin {
	return []NetworkPlugin{
		NetworkPluginAzure,
		NetworkPluginKubenet,
		NetworkPluginNone,
	}
}

// NetworkPluginMode - The mode the network plugin should use.
type NetworkPluginMode string

const (
	// NetworkPluginModeOverlay - Used with networkPlugin=azure, pods are given IPs from the PodCIDR address space but use Azure
	// Routing Domains rather than Kubenet's method of route tables. For more information visit https://aka.ms/aks/azure-cni-overlay.
	NetworkPluginModeOverlay NetworkPluginMode = "overlay"
)

// PossibleNetworkPluginModeValues returns the possible values for the NetworkPluginMode const type.
func PossibleNetworkPluginModeValues() []NetworkPluginMode {
	return []NetworkPluginMode{
		NetworkPluginModeOverlay,
	}
}

// NetworkPolicy - Network policy used for building the Kubernetes network.
type NetworkPolicy string

const (
	// NetworkPolicyAzure - Use Azure network policies. See [differences between Azure and Calico policies](https://docs.microsoft.com/azure/aks/use-network-policies#differences-between-azure-and-calico-policies-and-their-capabilities)
	// for more information.
	NetworkPolicyAzure NetworkPolicy = "azure"
	// NetworkPolicyCalico - Use Calico network policies. See [differences between Azure and Calico policies](https://docs.microsoft.com/azure/aks/use-network-policies#differences-between-azure-and-calico-policies-and-their-capabilities)
	// for more information.
	NetworkPolicyCalico NetworkPolicy = "calico"
	// NetworkPolicyCilium - Use Cilium to enforce network policies. This requires networkDataplane to be 'cilium'.
	NetworkPolicyCilium NetworkPolicy = "cilium"
)

// PossibleNetworkPolicyValues returns the possible values for the NetworkPolicy const type.
func PossibleNetworkPolicyValues() []NetworkPolicy {
	return []NetworkPolicy{
		NetworkPolicyAzure,
		NetworkPolicyCalico,
		NetworkPolicyCilium,
	}
}

// NodeOSUpgradeChannel - Manner in which the OS on your nodes is updated. The default is NodeImage.
type NodeOSUpgradeChannel string

const (
	// NodeOSUpgradeChannelNodeImage - AKS will update the nodes with a newly patched VHD containing security fixes and bugfixes
	// on a weekly cadence. With the VHD update machines will be rolling reimaged to that VHD following maintenance windows and
	// surge settings. No extra VHD cost is incurred when choosing this option as AKS hosts the images.
	NodeOSUpgradeChannelNodeImage NodeOSUpgradeChannel = "NodeImage"
	// NodeOSUpgradeChannelNone - No attempt to update your machines OS will be made either by OS or by rolling VHDs. This means
	// you are responsible for your security updates
	NodeOSUpgradeChannelNone NodeOSUpgradeChannel = "None"
	// NodeOSUpgradeChannelUnmanaged - OS updates will be applied automatically through the OS built-in patching infrastructure.
	// Newly scaled in machines will be unpatched initially and will be patched at some point by the OS's infrastructure. Behavior
	// of this option depends on the OS in question. Ubuntu and Mariner apply security patches through unattended upgrade roughly
	// once a day around 06:00 UTC. Windows does not apply security patches automatically and so for them this option is equivalent
	// to None till further notice
	NodeOSUpgradeChannelUnmanaged NodeOSUpgradeChannel = "Unmanaged"
)

// PossibleNodeOSUpgradeChannelValues returns the possible values for the NodeOSUpgradeChannel const type.
func PossibleNodeOSUpgradeChannelValues() []NodeOSUpgradeChannel {
	return []NodeOSUpgradeChannel{
		NodeOSUpgradeChannelNodeImage,
		NodeOSUpgradeChannelNone,
		NodeOSUpgradeChannelUnmanaged,
	}
}

// OSDiskType - The default is 'Ephemeral' if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB.
// Otherwise, defaults to 'Managed'. May not be changed after creation. For more information
// see Ephemeral OS [https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os].
type OSDiskType string

const (
	// OSDiskTypeEphemeral - Ephemeral OS disks are stored only on the host machine, just like a temporary disk. This provides
	// lower read/write latency, along with faster node scaling and cluster upgrades.
	OSDiskTypeEphemeral OSDiskType = "Ephemeral"
	// OSDiskTypeManaged - Azure replicates the operating system disk for a virtual machine to Azure storage to avoid data loss
	// should the VM need to be relocated to another host. Since containers aren't designed to have local state persisted, this
	// behavior offers limited value while providing some drawbacks, including slower node provisioning and higher read/write
	// latency.
	OSDiskTypeManaged OSDiskType = "Managed"
)

// PossibleOSDiskTypeValues returns the possible values for the OSDiskType const type.
func PossibleOSDiskTypeValues() []OSDiskType {
	return []OSDiskType{
		OSDiskTypeEphemeral,
		OSDiskTypeManaged,
	}
}

// OSSKU - Specifies the OS SKU used by the agent pool. The default is Ubuntu if OSType is Linux. The default is Windows2019
// when Kubernetes = 1.25 if OSType is Windows.
type OSSKU string

const (
	// OSSKUAzureLinux - Use AzureLinux as the OS for node images. Azure Linux is a container-optimized Linux distro built by
	// Microsoft, visit https://aka.ms/azurelinux for more information.
	OSSKUAzureLinux OSSKU = "AzureLinux"
	// OSSKUCBLMariner - Deprecated OSSKU. Microsoft recommends that new deployments choose 'AzureLinux' instead.
	OSSKUCBLMariner OSSKU = "CBLMariner"
	// OSSKUUbuntu - Use Ubuntu as the OS for node images.
	OSSKUUbuntu OSSKU = "Ubuntu"
	// OSSKUWindows2019 - Use Windows2019 as the OS for node images. Unsupported for system node pools. Windows2019 only supports
	// Windows2019 containers; it cannot run Windows2022 containers and vice versa.
	OSSKUWindows2019 OSSKU = "Windows2019"
	// OSSKUWindows2022 - Use Windows2022 as the OS for node images. Unsupported for system node pools. Windows2022 only supports
	// Windows2022 containers; it cannot run Windows2019 containers and vice versa.
	OSSKUWindows2022 OSSKU = "Windows2022"
)

// PossibleOSSKUValues returns the possible values for the OSSKU const type.
func PossibleOSSKUValues() []OSSKU {
	return []OSSKU{
		OSSKUAzureLinux,
		OSSKUCBLMariner,
		OSSKUUbuntu,
		OSSKUWindows2019,
		OSSKUWindows2022,
	}
}

// OSType - The operating system type. The default is Linux.
type OSType string

const (
	// OSTypeLinux - Use Linux.
	OSTypeLinux OSType = "Linux"
	// OSTypeWindows - Use Windows.
	OSTypeWindows OSType = "Windows"
)

// PossibleOSTypeValues returns the possible values for the OSType const type.
func PossibleOSTypeValues() []OSType {
	return []OSType{
		OSTypeLinux,
		OSTypeWindows,
	}
}

// OutboundType - This can only be set at cluster creation time and cannot be changed later. For more information see egress
// outbound type [https://docs.microsoft.com/azure/aks/egress-outboundtype].
type OutboundType string

const (
	// OutboundTypeLoadBalancer - The load balancer is used for egress through an AKS assigned public IP. This supports Kubernetes
	// services of type 'loadBalancer'. For more information see [outbound type loadbalancer](https://docs.microsoft.com/azure/aks/egress-outboundtype#outbound-type-of-loadbalancer).
	OutboundTypeLoadBalancer OutboundType = "loadBalancer"
	// OutboundTypeManagedNATGateway - The AKS-managed NAT gateway is used for egress.
	OutboundTypeManagedNATGateway OutboundType = "managedNATGateway"
	// OutboundTypeUserAssignedNATGateway - The user-assigned NAT gateway associated to the cluster subnet is used for egress.
	// This is an advanced scenario and requires proper network configuration.
	OutboundTypeUserAssignedNATGateway OutboundType = "userAssignedNATGateway"
	// OutboundTypeUserDefinedRouting - Egress paths must be defined by the user. This is an advanced scenario and requires proper
	// network configuration. For more information see [outbound type userDefinedRouting](https://docs.microsoft.com/azure/aks/egress-outboundtype#outbound-type-of-userdefinedrouting).
	OutboundTypeUserDefinedRouting OutboundType = "userDefinedRouting"
)

// PossibleOutboundTypeValues returns the possible values for the OutboundType const type.
func PossibleOutboundTypeValues() []OutboundType {
	return []OutboundType{
		OutboundTypeLoadBalancer,
		OutboundTypeManagedNATGateway,
		OutboundTypeUserAssignedNATGateway,
		OutboundTypeUserDefinedRouting,
	}
}

// PrivateEndpointConnectionProvisioningState - The current provisioning state.
type PrivateEndpointConnectionProvisioningState string

const (
	PrivateEndpointConnectionProvisioningStateCanceled  PrivateEndpointConnectionProvisioningState = "Canceled"
	PrivateEndpointConnectionProvisioningStateCreating  PrivateEndpointConnectionProvisioningState = "Creating"
	PrivateEndpointConnectionProvisioningStateDeleting  PrivateEndpointConnectionProvisioningState = "Deleting"
	PrivateEndpointConnectionProvisioningStateFailed    PrivateEndpointConnectionProvisioningState = "Failed"
	PrivateEndpointConnectionProvisioningStateSucceeded PrivateEndpointConnectionProvisioningState = "Succeeded"
)

// PossiblePrivateEndpointConnectionProvisioningStateValues returns the possible values for the PrivateEndpointConnectionProvisioningState const type.
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return []PrivateEndpointConnectionProvisioningState{
		PrivateEndpointConnectionProvisioningStateCanceled,
		PrivateEndpointConnectionProvisioningStateCreating,
		PrivateEndpointConnectionProvisioningStateDeleting,
		PrivateEndpointConnectionProvisioningStateFailed,
		PrivateEndpointConnectionProvisioningStateSucceeded,
	}
}

// PublicNetworkAccess - Allow or deny public network access for AKS
type PublicNetworkAccess string

const (
	PublicNetworkAccessDisabled PublicNetworkAccess = "Disabled"
	PublicNetworkAccessEnabled  PublicNetworkAccess = "Enabled"
)

// PossiblePublicNetworkAccessValues returns the possible values for the PublicNetworkAccess const type.
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return []PublicNetworkAccess{
		PublicNetworkAccessDisabled,
		PublicNetworkAccessEnabled,
	}
}

// ResourceIdentityType - For more information see use managed identities in AKS [https://docs.microsoft.com/azure/aks/use-managed-identity].
type ResourceIdentityType string

const (
	// ResourceIdentityTypeNone - Do not use a managed identity for the Managed Cluster, service principal will be used instead.
	ResourceIdentityTypeNone ResourceIdentityType = "None"
	// ResourceIdentityTypeSystemAssigned - Use an implicitly created system assigned managed identity to manage cluster resources.
	// Master components in the control plane such as kube-controller-manager will use the system assigned managed identity to
	// manipulate Azure resources.
	ResourceIdentityTypeSystemAssigned ResourceIdentityType = "SystemAssigned"
	// ResourceIdentityTypeUserAssigned - Use a user-specified identity to manage cluster resources. Master components in the
	// control plane such as kube-controller-manager will use the specified user assigned managed identity to manipulate Azure
	// resources.
	ResourceIdentityTypeUserAssigned ResourceIdentityType = "UserAssigned"
)

// PossibleResourceIdentityTypeValues returns the possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{
		ResourceIdentityTypeNone,
		ResourceIdentityTypeSystemAssigned,
		ResourceIdentityTypeUserAssigned,
	}
}

// ScaleDownMode - Describes how VMs are added to or removed from Agent Pools. See billing states [https://docs.microsoft.com/azure/virtual-machines/states-billing].
type ScaleDownMode string

const (
	// ScaleDownModeDeallocate - Attempt to start deallocated instances (if they exist) during scale up and deallocate instances
	// during scale down.
	ScaleDownModeDeallocate ScaleDownMode = "Deallocate"
	// ScaleDownModeDelete - Create new instances during scale up and remove instances during scale down.
	ScaleDownModeDelete ScaleDownMode = "Delete"
)

// PossibleScaleDownModeValues returns the possible values for the ScaleDownMode const type.
func PossibleScaleDownModeValues() []ScaleDownMode {
	return []ScaleDownMode{
		ScaleDownModeDeallocate,
		ScaleDownModeDelete,
	}
}

// ScaleSetEvictionPolicy - The eviction policy specifies what to do with the VM when it is evicted. The default is Delete.
// For more information about eviction see spot VMs
// [https://docs.microsoft.com/azure/virtual-machines/spot-vms]
type ScaleSetEvictionPolicy string

const (
	// ScaleSetEvictionPolicyDeallocate - Nodes in the underlying Scale Set of the node pool are set to the stopped-deallocated
	// state upon eviction. Nodes in the stopped-deallocated state count against your compute quota and can cause issues with
	// cluster scaling or upgrading.
	ScaleSetEvictionPolicyDeallocate ScaleSetEvictionPolicy = "Deallocate"
	// ScaleSetEvictionPolicyDelete - Nodes in the underlying Scale Set of the node pool are deleted when they're evicted.
	ScaleSetEvictionPolicyDelete ScaleSetEvictionPolicy = "Delete"
)

// PossibleScaleSetEvictionPolicyValues returns the possible values for the ScaleSetEvictionPolicy const type.
func PossibleScaleSetEvictionPolicyValues() []ScaleSetEvictionPolicy {
	return []ScaleSetEvictionPolicy{
		ScaleSetEvictionPolicyDeallocate,
		ScaleSetEvictionPolicyDelete,
	}
}

// ScaleSetPriority - The Virtual Machine Scale Set priority.
type ScaleSetPriority string

const (
	// ScaleSetPriorityRegular - Regular VMs will be used.
	ScaleSetPriorityRegular ScaleSetPriority = "Regular"
	// ScaleSetPrioritySpot - Spot priority VMs will be used. There is no SLA for spot nodes. See [spot on AKS](https://docs.microsoft.com/azure/aks/spot-node-pool)
	// for more information.
	ScaleSetPrioritySpot ScaleSetPriority = "Spot"
)

// PossibleScaleSetPriorityValues returns the possible values for the ScaleSetPriority const type.
func PossibleScaleSetPriorityValues() []ScaleSetPriority {
	return []ScaleSetPriority{
		ScaleSetPriorityRegular,
		ScaleSetPrioritySpot,
	}
}

// ServiceMeshMode - Mode of the service mesh.
type ServiceMeshMode string

const (
	// ServiceMeshModeDisabled - Mesh is disabled.
	ServiceMeshModeDisabled ServiceMeshMode = "Disabled"
	// ServiceMeshModeIstio - Istio deployed as an AKS addon.
	ServiceMeshModeIstio ServiceMeshMode = "Istio"
)

// PossibleServiceMeshModeValues returns the possible values for the ServiceMeshMode const type.
func PossibleServiceMeshModeValues() []ServiceMeshMode {
	return []ServiceMeshMode{
		ServiceMeshModeDisabled,
		ServiceMeshModeIstio,
	}
}

// SnapshotType - The type of a snapshot. The default is NodePool.
type SnapshotType string

const (
	// SnapshotTypeNodePool - The snapshot is a snapshot of a node pool.
	SnapshotTypeNodePool SnapshotType = "NodePool"
)

// PossibleSnapshotTypeValues returns the possible values for the SnapshotType const type.
func PossibleSnapshotTypeValues() []SnapshotType {
	return []SnapshotType{
		SnapshotTypeNodePool,
	}
}

// Type - Specifies on which week of the month the dayOfWeek applies.
type Type string

const (
	// TypeFirst - First week of the month.
	TypeFirst Type = "First"
	// TypeFourth - Fourth week of the month.
	TypeFourth Type = "Fourth"
	// TypeLast - Last week of the month.
	TypeLast Type = "Last"
	// TypeSecond - Second week of the month.
	TypeSecond Type = "Second"
	// TypeThird - Third week of the month.
	TypeThird Type = "Third"
)

// PossibleTypeValues returns the possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{
		TypeFirst,
		TypeFourth,
		TypeLast,
		TypeSecond,
		TypeThird,
	}
}

// UpgradeChannel - For more information see setting the AKS cluster auto-upgrade channel [https://docs.microsoft.com/azure/aks/upgrade-cluster#set-auto-upgrade-channel].
type UpgradeChannel string

const (
	// UpgradeChannelNodeImage - Automatically upgrade the node image to the latest version available. Consider using nodeOSUpgradeChannel
	// instead as that allows you to configure node OS patching separate from Kubernetes version patching
	UpgradeChannelNodeImage UpgradeChannel = "node-image"
	// UpgradeChannelNone - Disables auto-upgrades and keeps the cluster at its current version of Kubernetes.
	UpgradeChannelNone UpgradeChannel = "none"
	// UpgradeChannelPatch - Automatically upgrade the cluster to the latest supported patch version when it becomes available
	// while keeping the minor version the same. For example, if a cluster is running version 1.17.7 and versions 1.17.9, 1.18.4,
	// 1.18.6, and 1.19.1 are available, your cluster is upgraded to 1.17.9.
	UpgradeChannelPatch UpgradeChannel = "patch"
	// UpgradeChannelRapid - Automatically upgrade the cluster to the latest supported patch release on the latest supported minor
	// version. In cases where the cluster is at a version of Kubernetes that is at an N-2 minor version where N is the latest
	// supported minor version, the cluster first upgrades to the latest supported patch version on N-1 minor version. For example,
	// if a cluster is running version 1.17.7 and versions 1.17.9, 1.18.4, 1.18.6, and 1.19.1 are available, your cluster first
	// is upgraded to 1.18.6, then is upgraded to 1.19.1.
	UpgradeChannelRapid UpgradeChannel = "rapid"
	// UpgradeChannelStable - Automatically upgrade the cluster to the latest supported patch release on minor version N-1, where
	// N is the latest supported minor version. For example, if a cluster is running version 1.17.7 and versions 1.17.9, 1.18.4,
	// 1.18.6, and 1.19.1 are available, your cluster is upgraded to 1.18.6.
	UpgradeChannelStable UpgradeChannel = "stable"
)

// PossibleUpgradeChannelValues returns the possible values for the UpgradeChannel const type.
func PossibleUpgradeChannelValues() []UpgradeChannel {
	return []UpgradeChannel{
		UpgradeChannelNodeImage,
		UpgradeChannelNone,
		UpgradeChannelPatch,
		UpgradeChannelRapid,
		UpgradeChannelStable,
	}
}

// WeekDay - The weekday enum.
type WeekDay string

const (
	WeekDayFriday    WeekDay = "Friday"
	WeekDayMonday    WeekDay = "Monday"
	WeekDaySaturday  WeekDay = "Saturday"
	WeekDaySunday    WeekDay = "Sunday"
	WeekDayThursday  WeekDay = "Thursday"
	WeekDayTuesday   WeekDay = "Tuesday"
	WeekDayWednesday WeekDay = "Wednesday"
)

// PossibleWeekDayValues returns the possible values for the WeekDay const type.
func PossibleWeekDayValues() []WeekDay {
	return []WeekDay{
		WeekDayFriday,
		WeekDayMonday,
		WeekDaySaturday,
		WeekDaySunday,
		WeekDayThursday,
		WeekDayTuesday,
		WeekDayWednesday,
	}
}

// WorkloadRuntime - Determines the type of workload a node can run.
type WorkloadRuntime string

const (
	// WorkloadRuntimeOCIContainer - Nodes will use Kubelet to run standard OCI container workloads.
	WorkloadRuntimeOCIContainer WorkloadRuntime = "OCIContainer"
	// WorkloadRuntimeWasmWasi - Nodes will use Krustlet to run WASM workloads using the WASI provider (Preview).
	WorkloadRuntimeWasmWasi WorkloadRuntime = "WasmWasi"
)

// PossibleWorkloadRuntimeValues returns the possible values for the WorkloadRuntime const type.
func PossibleWorkloadRuntimeValues() []WorkloadRuntime {
	return []WorkloadRuntime{
		WorkloadRuntimeOCIContainer,
		WorkloadRuntimeWasmWasi,
	}
}
