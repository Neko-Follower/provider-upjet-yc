/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CiliumInitParameters struct {
}

type CiliumObservation struct {
}

type CiliumParameters struct {
}

type ClusterInitParameters struct {

	// CIDR block. IP range for allocating pod addresses.
	// It should not overlap with any subnet in the network the Kubernetes cluster located in. Static routes will be
	// set up for this CIDR blocks in node subnets.
	ClusterIPv4Range *string `json:"clusterIpv4Range,omitempty" tf:"cluster_ipv4_range,omitempty"`

	// Identical to cluster_ipv4_range but for IPv6 protocol.
	ClusterIPv6Range *string `json:"clusterIpv6Range,omitempty" tf:"cluster_ipv6_range,omitempty"`

	// A description of the Kubernetes cluster.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// cluster KMS provider parameters.
	KMSProvider []KMSProviderInitParameters `json:"kmsProvider,omitempty" tf:"kms_provider,omitempty"`

	// A set of key/value label pairs to assign to the Kubernetes cluster.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Kubernetes master configuration options. The structure is documented below.
	Master []MasterInitParameters `json:"master,omitempty" tf:"master,omitempty"`

	// Name of a specific Kubernetes cluster.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Network Implementation options. The structure is documented below.
	NetworkImplementation []NetworkImplementationInitParameters `json:"networkImplementation,omitempty" tf:"network_implementation,omitempty"`

	// Network policy provider for the cluster. Possible values: CALICO.
	NetworkPolicyProvider *string `json:"networkPolicyProvider,omitempty" tf:"network_policy_provider,omitempty"`

	// Size of the masks that are assigned to each node in the cluster. Effectively limits maximum number of pods for each node.
	NodeIPv4CidrMaskSize *float64 `json:"nodeIpv4CidrMaskSize,omitempty" tf:"node_ipv4_cidr_mask_size,omitempty"`

	// Cluster release channel.
	ReleaseChannel *string `json:"releaseChannel,omitempty" tf:"release_channel,omitempty"`

	// CIDR block. IP range Kubernetes service Kubernetes cluster
	// IP addresses will be allocated from. It should not overlap with any subnet in the network
	// the Kubernetes cluster located in.
	ServiceIPv4Range *string `json:"serviceIpv4Range,omitempty" tf:"service_ipv4_range,omitempty"`

	// Identical to service_ipv4_range but for IPv6 protocol.
	ServiceIPv6Range *string `json:"serviceIpv6Range,omitempty" tf:"service_ipv6_range,omitempty"`
}

type ClusterObservation struct {

	// CIDR block. IP range for allocating pod addresses.
	// It should not overlap with any subnet in the network the Kubernetes cluster located in. Static routes will be
	// set up for this CIDR blocks in node subnets.
	ClusterIPv4Range *string `json:"clusterIpv4Range,omitempty" tf:"cluster_ipv4_range,omitempty"`

	// Identical to cluster_ipv4_range but for IPv6 protocol.
	ClusterIPv6Range *string `json:"clusterIpv6Range,omitempty" tf:"cluster_ipv6_range,omitempty"`

	// (Computed) The Kubernetes cluster creation timestamp.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// A description of the Kubernetes cluster.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the folder that the Kubernetes cluster belongs to.
	// If it is not provided, the default provider folder is used.
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// (Computed) Health of the Kubernetes cluster.
	Health *string `json:"health,omitempty" tf:"health,omitempty"`

	// (Computed) ID of a new Kubernetes cluster.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// cluster KMS provider parameters.
	KMSProvider []KMSProviderObservation `json:"kmsProvider,omitempty" tf:"kms_provider,omitempty"`

	// A set of key/value label pairs to assign to the Kubernetes cluster.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Log group where cluster stores cluster system logs, like audit, events, or controlplane logs.
	LogGroupID *string `json:"logGroupId,omitempty" tf:"log_group_id,omitempty"`

	// Kubernetes master configuration options. The structure is documented below.
	Master []MasterObservation `json:"master,omitempty" tf:"master,omitempty"`

	// Name of a specific Kubernetes cluster.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the cluster network.
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Network Implementation options. The structure is documented below.
	NetworkImplementation []NetworkImplementationObservation `json:"networkImplementation,omitempty" tf:"network_implementation,omitempty"`

	// Network policy provider for the cluster. Possible values: CALICO.
	NetworkPolicyProvider *string `json:"networkPolicyProvider,omitempty" tf:"network_policy_provider,omitempty"`

	// Size of the masks that are assigned to each node in the cluster. Effectively limits maximum number of pods for each node.
	NodeIPv4CidrMaskSize *float64 `json:"nodeIpv4CidrMaskSize,omitempty" tf:"node_ipv4_cidr_mask_size,omitempty"`

	// Service account to be used by the worker nodes of the Kubernetes cluster
	// to access Container Registry or to push node logs and metrics.
	NodeServiceAccountID *string `json:"nodeServiceAccountId,omitempty" tf:"node_service_account_id,omitempty"`

	// Cluster release channel.
	ReleaseChannel *string `json:"releaseChannel,omitempty" tf:"release_channel,omitempty"`

	// Service account to be used for provisioning Compute Cloud and VPC resources
	// for Kubernetes cluster. Selected service account should have edit role on the folder where the Kubernetes
	// cluster will be located and on the folder where selected network resides.
	ServiceAccountID *string `json:"serviceAccountId,omitempty" tf:"service_account_id,omitempty"`

	// CIDR block. IP range Kubernetes service Kubernetes cluster
	// IP addresses will be allocated from. It should not overlap with any subnet in the network
	// the Kubernetes cluster located in.
	ServiceIPv4Range *string `json:"serviceIpv4Range,omitempty" tf:"service_ipv4_range,omitempty"`

	// Identical to service_ipv4_range but for IPv6 protocol.
	ServiceIPv6Range *string `json:"serviceIpv6Range,omitempty" tf:"service_ipv6_range,omitempty"`

	// (Computed)Status of the Kubernetes cluster.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type ClusterParameters struct {

	// CIDR block. IP range for allocating pod addresses.
	// It should not overlap with any subnet in the network the Kubernetes cluster located in. Static routes will be
	// set up for this CIDR blocks in node subnets.
	// +kubebuilder:validation:Optional
	ClusterIPv4Range *string `json:"clusterIpv4Range,omitempty" tf:"cluster_ipv4_range,omitempty"`

	// Identical to cluster_ipv4_range but for IPv6 protocol.
	// +kubebuilder:validation:Optional
	ClusterIPv6Range *string `json:"clusterIpv6Range,omitempty" tf:"cluster_ipv6_range,omitempty"`

	// A description of the Kubernetes cluster.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the folder that the Kubernetes cluster belongs to.
	// If it is not provided, the default provider folder is used.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/resourcemanager/v1alpha1.Folder
	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// cluster KMS provider parameters.
	// +kubebuilder:validation:Optional
	KMSProvider []KMSProviderParameters `json:"kmsProvider,omitempty" tf:"kms_provider,omitempty"`

	// A set of key/value label pairs to assign to the Kubernetes cluster.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Kubernetes master configuration options. The structure is documented below.
	// +kubebuilder:validation:Optional
	Master []MasterParameters `json:"master,omitempty" tf:"master,omitempty"`

	// Name of a specific Kubernetes cluster.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the cluster network.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/vpc/v1alpha1.Network
	// +kubebuilder:validation:Optional
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a Network in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a Network in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// Network Implementation options. The structure is documented below.
	// +kubebuilder:validation:Optional
	NetworkImplementation []NetworkImplementationParameters `json:"networkImplementation,omitempty" tf:"network_implementation,omitempty"`

	// Network policy provider for the cluster. Possible values: CALICO.
	// +kubebuilder:validation:Optional
	NetworkPolicyProvider *string `json:"networkPolicyProvider,omitempty" tf:"network_policy_provider,omitempty"`

	// Size of the masks that are assigned to each node in the cluster. Effectively limits maximum number of pods for each node.
	// +kubebuilder:validation:Optional
	NodeIPv4CidrMaskSize *float64 `json:"nodeIpv4CidrMaskSize,omitempty" tf:"node_ipv4_cidr_mask_size,omitempty"`

	// Service account to be used by the worker nodes of the Kubernetes cluster
	// to access Container Registry or to push node logs and metrics.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/vpc/v1alpha1.SecurityGroup
	// +kubebuilder:validation:Optional
	NodeServiceAccountID *string `json:"nodeServiceAccountId,omitempty" tf:"node_service_account_id,omitempty"`

	// Reference to a SecurityGroup in vpc to populate nodeServiceAccountId.
	// +kubebuilder:validation:Optional
	NodeServiceAccountIDRef *v1.Reference `json:"nodeServiceAccountIdRef,omitempty" tf:"-"`

	// Selector for a SecurityGroup in vpc to populate nodeServiceAccountId.
	// +kubebuilder:validation:Optional
	NodeServiceAccountIDSelector *v1.Selector `json:"nodeServiceAccountIdSelector,omitempty" tf:"-"`

	// Cluster release channel.
	// +kubebuilder:validation:Optional
	ReleaseChannel *string `json:"releaseChannel,omitempty" tf:"release_channel,omitempty"`

	// Service account to be used for provisioning Compute Cloud and VPC resources
	// for Kubernetes cluster. Selected service account should have edit role on the folder where the Kubernetes
	// cluster will be located and on the folder where selected network resides.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/vpc/v1alpha1.SecurityGroup
	// +kubebuilder:validation:Optional
	ServiceAccountID *string `json:"serviceAccountId,omitempty" tf:"service_account_id,omitempty"`

	// Reference to a SecurityGroup in vpc to populate serviceAccountId.
	// +kubebuilder:validation:Optional
	ServiceAccountIDRef *v1.Reference `json:"serviceAccountIdRef,omitempty" tf:"-"`

	// Selector for a SecurityGroup in vpc to populate serviceAccountId.
	// +kubebuilder:validation:Optional
	ServiceAccountIDSelector *v1.Selector `json:"serviceAccountIdSelector,omitempty" tf:"-"`

	// CIDR block. IP range Kubernetes service Kubernetes cluster
	// IP addresses will be allocated from. It should not overlap with any subnet in the network
	// the Kubernetes cluster located in.
	// +kubebuilder:validation:Optional
	ServiceIPv4Range *string `json:"serviceIpv4Range,omitempty" tf:"service_ipv4_range,omitempty"`

	// Identical to service_ipv4_range but for IPv6 protocol.
	// +kubebuilder:validation:Optional
	ServiceIPv6Range *string `json:"serviceIpv6Range,omitempty" tf:"service_ipv6_range,omitempty"`
}

type KMSProviderInitParameters struct {
}

type KMSProviderObservation struct {

	// KMS key ID.
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`
}

type KMSProviderParameters struct {

	// KMS key ID.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/kms/v1alpha1.SymmetricKey
	// +kubebuilder:validation:Optional
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`

	// Reference to a SymmetricKey in kms to populate keyId.
	// +kubebuilder:validation:Optional
	KeyIDRef *v1.Reference `json:"keyIdRef,omitempty" tf:"-"`

	// Selector for a SymmetricKey in kms to populate keyId.
	// +kubebuilder:validation:Optional
	KeyIDSelector *v1.Selector `json:"keyIdSelector,omitempty" tf:"-"`
}

type LocationInitParameters struct {

	// ID of the availability zone.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type LocationObservation struct {

	// ID of the subnet. If no ID is specified, and there only one subnet in specified zone, an address in this subnet will be allocated.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// ID of the availability zone.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type LocationParameters struct {

	// ID of the subnet. If no ID is specified, and there only one subnet in specified zone, an address in this subnet will be allocated.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/vpc/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// ID of the availability zone.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type MaintenancePolicyInitParameters struct {

	// Boolean flag that specifies if master can be upgraded automatically. When omitted, default value is TRUE.
	AutoUpgrade *bool `json:"autoUpgrade,omitempty" tf:"auto_upgrade,omitempty"`

	// (Computed) This structure specifies maintenance window, when update for master is allowed. When omitted, it defaults to any time.
	// To specify time of day interval, for all days, one element should be provided, with two fields set, start_time and duration.
	// Please see zonal_cluster_resource_name config example.
	MaintenanceWindow []MaintenanceWindowInitParameters `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`
}

type MaintenancePolicyObservation struct {

	// Boolean flag that specifies if master can be upgraded automatically. When omitted, default value is TRUE.
	AutoUpgrade *bool `json:"autoUpgrade,omitempty" tf:"auto_upgrade,omitempty"`

	// (Computed) This structure specifies maintenance window, when update for master is allowed. When omitted, it defaults to any time.
	// To specify time of day interval, for all days, one element should be provided, with two fields set, start_time and duration.
	// Please see zonal_cluster_resource_name config example.
	MaintenanceWindow []MaintenanceWindowObservation `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`
}

type MaintenancePolicyParameters struct {

	// Boolean flag that specifies if master can be upgraded automatically. When omitted, default value is TRUE.
	// +kubebuilder:validation:Optional
	AutoUpgrade *bool `json:"autoUpgrade" tf:"auto_upgrade,omitempty"`

	// (Computed) This structure specifies maintenance window, when update for master is allowed. When omitted, it defaults to any time.
	// To specify time of day interval, for all days, one element should be provided, with two fields set, start_time and duration.
	// Please see zonal_cluster_resource_name config example.
	// +kubebuilder:validation:Optional
	MaintenanceWindow []MaintenanceWindowParameters `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`
}

type MaintenanceWindowInitParameters struct {
	Day *string `json:"day,omitempty" tf:"day,omitempty"`

	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`

	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

type MaintenanceWindowObservation struct {
	Day *string `json:"day,omitempty" tf:"day,omitempty"`

	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`

	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

type MaintenanceWindowParameters struct {

	// +kubebuilder:validation:Optional
	Day *string `json:"day,omitempty" tf:"day,omitempty"`

	// +kubebuilder:validation:Optional
	Duration *string `json:"duration" tf:"duration,omitempty"`

	// +kubebuilder:validation:Optional
	StartTime *string `json:"startTime" tf:"start_time,omitempty"`
}

type MasterInitParameters struct {
	ExternalV6Address *string `json:"externalV6Address,omitempty" tf:"external_v6_address,omitempty"`

	// (Computed) Maintenance policy for Kubernetes master.
	// If policy is omitted, automatic revision upgrades of the kubernetes master are enabled and could happen at any time.
	// Revision upgrades are performed only within the same minor version, e.g. 1.13.
	// Minor version upgrades (e.g. 1.13->1.14) should be performed manually. The structure is documented below.
	MaintenancePolicy []MaintenancePolicyInitParameters `json:"maintenancePolicy,omitempty" tf:"maintenance_policy,omitempty"`

	// Master Logging options. The structure is documented below.
	MasterLogging []MasterLoggingInitParameters `json:"masterLogging,omitempty" tf:"master_logging,omitempty"`

	// (Computed) Boolean flag. When true, Kubernetes master will have visible ipv4 address.
	PublicIP *bool `json:"publicIp,omitempty" tf:"public_ip,omitempty"`

	// Initialize parameters for Regional Master (highly available master). The structure is documented below.
	Regional []RegionalInitParameters `json:"regional,omitempty" tf:"regional,omitempty"`

	// (Computed) Version of Kubernetes that will be used for master.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// Initialize parameters for Zonal Master (single node master). The structure is documented below.
	Zonal []ZonalInitParameters `json:"zonal,omitempty" tf:"zonal,omitempty"`
}

type MasterLoggingInitParameters struct {

	// Boolean flag that specifies if kube-apiserver audit logs should be sent to Yandex Cloud Logging.
	AuditEnabled *bool `json:"auditEnabled,omitempty" tf:"audit_enabled,omitempty"`

	// Boolean flag that specifies if cluster-autoscaler logs should be sent to Yandex Cloud Logging.
	ClusterAutoscalerEnabled *bool `json:"clusterAutoscalerEnabled,omitempty" tf:"cluster_autoscaler_enabled,omitempty"`

	// Boolean flag that specifies if master components logs should be sent to Yandex Cloud Logging. The exact components that will send their logs must be configured via the options described below.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Boolean flag that specifies if kubernetes cluster events should be sent to Yandex Cloud Logging.
	EventsEnabled *bool `json:"eventsEnabled,omitempty" tf:"events_enabled,omitempty"`

	// ID of the folder default Log group of which should be used to collect logs.
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Boolean flag that specifies if kube-apiserver logs should be sent to Yandex Cloud Logging.
	KubeApiserverEnabled *bool `json:"kubeApiserverEnabled,omitempty" tf:"kube_apiserver_enabled,omitempty"`

	// ID of the Yandex Cloud Logging Log group.
	LogGroupID *string `json:"logGroupId,omitempty" tf:"log_group_id,omitempty"`
}

type MasterLoggingObservation struct {

	// Boolean flag that specifies if kube-apiserver audit logs should be sent to Yandex Cloud Logging.
	AuditEnabled *bool `json:"auditEnabled,omitempty" tf:"audit_enabled,omitempty"`

	// Boolean flag that specifies if cluster-autoscaler logs should be sent to Yandex Cloud Logging.
	ClusterAutoscalerEnabled *bool `json:"clusterAutoscalerEnabled,omitempty" tf:"cluster_autoscaler_enabled,omitempty"`

	// Boolean flag that specifies if master components logs should be sent to Yandex Cloud Logging. The exact components that will send their logs must be configured via the options described below.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Boolean flag that specifies if kubernetes cluster events should be sent to Yandex Cloud Logging.
	EventsEnabled *bool `json:"eventsEnabled,omitempty" tf:"events_enabled,omitempty"`

	// ID of the folder default Log group of which should be used to collect logs.
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Boolean flag that specifies if kube-apiserver logs should be sent to Yandex Cloud Logging.
	KubeApiserverEnabled *bool `json:"kubeApiserverEnabled,omitempty" tf:"kube_apiserver_enabled,omitempty"`

	// ID of the Yandex Cloud Logging Log group.
	LogGroupID *string `json:"logGroupId,omitempty" tf:"log_group_id,omitempty"`
}

type MasterLoggingParameters struct {

	// Boolean flag that specifies if kube-apiserver audit logs should be sent to Yandex Cloud Logging.
	// +kubebuilder:validation:Optional
	AuditEnabled *bool `json:"auditEnabled,omitempty" tf:"audit_enabled,omitempty"`

	// Boolean flag that specifies if cluster-autoscaler logs should be sent to Yandex Cloud Logging.
	// +kubebuilder:validation:Optional
	ClusterAutoscalerEnabled *bool `json:"clusterAutoscalerEnabled,omitempty" tf:"cluster_autoscaler_enabled,omitempty"`

	// Boolean flag that specifies if master components logs should be sent to Yandex Cloud Logging. The exact components that will send their logs must be configured via the options described below.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Boolean flag that specifies if kubernetes cluster events should be sent to Yandex Cloud Logging.
	// +kubebuilder:validation:Optional
	EventsEnabled *bool `json:"eventsEnabled,omitempty" tf:"events_enabled,omitempty"`

	// ID of the folder default Log group of which should be used to collect logs.
	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Boolean flag that specifies if kube-apiserver logs should be sent to Yandex Cloud Logging.
	// +kubebuilder:validation:Optional
	KubeApiserverEnabled *bool `json:"kubeApiserverEnabled,omitempty" tf:"kube_apiserver_enabled,omitempty"`

	// ID of the Yandex Cloud Logging Log group.
	// +kubebuilder:validation:Optional
	LogGroupID *string `json:"logGroupId,omitempty" tf:"log_group_id,omitempty"`
}

type MasterObservation struct {

	// (Computed) PEM-encoded public certificate that is the root of trust for the Kubernetes cluster.
	ClusterCACertificate *string `json:"clusterCaCertificate,omitempty" tf:"cluster_ca_certificate,omitempty"`

	// (Computed) An IPv4 external network address that is assigned to the master.
	ExternalV4Address *string `json:"externalV4Address,omitempty" tf:"external_v4_address,omitempty"`

	// (Computed) External endpoint that can be used to access Kubernetes cluster API from the internet (outside of the cloud).
	ExternalV4Endpoint *string `json:"externalV4Endpoint,omitempty" tf:"external_v4_endpoint,omitempty"`

	ExternalV6Address *string `json:"externalV6Address,omitempty" tf:"external_v6_address,omitempty"`

	ExternalV6Endpoint *string `json:"externalV6Endpoint,omitempty" tf:"external_v6_endpoint,omitempty"`

	// (Computed) An IPv4 internal network address that is assigned to the master.
	InternalV4Address *string `json:"internalV4Address,omitempty" tf:"internal_v4_address,omitempty"`

	// (Computed) Internal endpoint that can be used to connect to the master from cloud networks.
	InternalV4Endpoint *string `json:"internalV4Endpoint,omitempty" tf:"internal_v4_endpoint,omitempty"`

	// (Computed) Maintenance policy for Kubernetes master.
	// If policy is omitted, automatic revision upgrades of the kubernetes master are enabled and could happen at any time.
	// Revision upgrades are performed only within the same minor version, e.g. 1.13.
	// Minor version upgrades (e.g. 1.13->1.14) should be performed manually. The structure is documented below.
	MaintenancePolicy []MaintenancePolicyObservation `json:"maintenancePolicy,omitempty" tf:"maintenance_policy,omitempty"`

	// Master Logging options. The structure is documented below.
	MasterLogging []MasterLoggingObservation `json:"masterLogging,omitempty" tf:"master_logging,omitempty"`

	// (Computed) Boolean flag. When true, Kubernetes master will have visible ipv4 address.
	PublicIP *bool `json:"publicIp,omitempty" tf:"public_ip,omitempty"`

	// Initialize parameters for Regional Master (highly available master). The structure is documented below.
	Regional []RegionalObservation `json:"regional,omitempty" tf:"regional,omitempty"`

	// List of security group IDs to which the Kubernetes cluster belongs.
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// (Computed) Version of Kubernetes that will be used for master.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// (Computed) Information about cluster version. The structure is documented below.
	VersionInfo []VersionInfoObservation `json:"versionInfo,omitempty" tf:"version_info,omitempty"`

	// Initialize parameters for Zonal Master (single node master). The structure is documented below.
	Zonal []ZonalObservation `json:"zonal,omitempty" tf:"zonal,omitempty"`
}

type MasterParameters struct {

	// +kubebuilder:validation:Optional
	ExternalV6Address *string `json:"externalV6Address,omitempty" tf:"external_v6_address,omitempty"`

	// (Computed) Maintenance policy for Kubernetes master.
	// If policy is omitted, automatic revision upgrades of the kubernetes master are enabled and could happen at any time.
	// Revision upgrades are performed only within the same minor version, e.g. 1.13.
	// Minor version upgrades (e.g. 1.13->1.14) should be performed manually. The structure is documented below.
	// +kubebuilder:validation:Optional
	MaintenancePolicy []MaintenancePolicyParameters `json:"maintenancePolicy,omitempty" tf:"maintenance_policy,omitempty"`

	// Master Logging options. The structure is documented below.
	// +kubebuilder:validation:Optional
	MasterLogging []MasterLoggingParameters `json:"masterLogging,omitempty" tf:"master_logging,omitempty"`

	// (Computed) Boolean flag. When true, Kubernetes master will have visible ipv4 address.
	// +kubebuilder:validation:Optional
	PublicIP *bool `json:"publicIp,omitempty" tf:"public_ip,omitempty"`

	// Initialize parameters for Regional Master (highly available master). The structure is documented below.
	// +kubebuilder:validation:Optional
	Regional []RegionalParameters `json:"regional,omitempty" tf:"regional,omitempty"`

	// List of security group IDs to which the Kubernetes cluster belongs.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/vpc/v1alpha1.SecurityGroup
	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// References to SecurityGroup in vpc to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIdsRefs []v1.Reference `json:"securityGroupIdsRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in vpc to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIdsSelector *v1.Selector `json:"securityGroupIdsSelector,omitempty" tf:"-"`

	// (Computed) Version of Kubernetes that will be used for master.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// Initialize parameters for Zonal Master (single node master). The structure is documented below.
	// +kubebuilder:validation:Optional
	Zonal []ZonalParameters `json:"zonal,omitempty" tf:"zonal,omitempty"`
}

type NetworkImplementationInitParameters struct {

	// Cilium network implementation configuration. No options exist.
	Cilium []CiliumInitParameters `json:"cilium,omitempty" tf:"cilium,omitempty"`
}

type NetworkImplementationObservation struct {

	// Cilium network implementation configuration. No options exist.
	Cilium []CiliumParameters `json:"cilium,omitempty" tf:"cilium,omitempty"`
}

type NetworkImplementationParameters struct {

	// Cilium network implementation configuration. No options exist.
	// +kubebuilder:validation:Optional
	Cilium []CiliumParameters `json:"cilium,omitempty" tf:"cilium,omitempty"`
}

type RegionalInitParameters struct {

	// Array of locations, where master instances will be allocated. The structure is documented below.
	Location []LocationInitParameters `json:"location,omitempty" tf:"location,omitempty"`

	// Name of availability region (e.g. "ru-central1"), where master instances will be allocated.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type RegionalObservation struct {

	// Array of locations, where master instances will be allocated. The structure is documented below.
	Location []LocationObservation `json:"location,omitempty" tf:"location,omitempty"`

	// Name of availability region (e.g. "ru-central1"), where master instances will be allocated.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type RegionalParameters struct {

	// Array of locations, where master instances will be allocated. The structure is documented below.
	// +kubebuilder:validation:Optional
	Location []LocationParameters `json:"location,omitempty" tf:"location,omitempty"`

	// Name of availability region (e.g. "ru-central1"), where master instances will be allocated.
	// +kubebuilder:validation:Optional
	Region *string `json:"region" tf:"region,omitempty"`
}

type VersionInfoInitParameters struct {
}

type VersionInfoObservation struct {

	// Current Kubernetes version, major.minor (e.g. 1.15).
	CurrentVersion *string `json:"currentVersion,omitempty" tf:"current_version,omitempty"`

	// Boolean flag.
	// Newer revisions may include Kubernetes patches (e.g 1.15.1 -> 1.15.2) as well
	// as some internal component updates - new features or bug fixes in yandex-specific
	// components either on the master or nodes.
	NewRevisionAvailable *bool `json:"newRevisionAvailable,omitempty" tf:"new_revision_available,omitempty"`

	// Human readable description of the changes to be applied
	// when updating to the latest revision. Empty if new_revision_available is false.
	NewRevisionSummary *string `json:"newRevisionSummary,omitempty" tf:"new_revision_summary,omitempty"`

	// Boolean flag. The current version is on the deprecation schedule,
	// component (master or node group) should be upgraded.
	VersionDeprecated *bool `json:"versionDeprecated,omitempty" tf:"version_deprecated,omitempty"`
}

type VersionInfoParameters struct {
}

type ZonalInitParameters struct {

	// ID of the availability zone.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type ZonalObservation struct {

	// ID of the subnet. If no ID is specified, and there only one subnet in specified zone, an address in this subnet will be allocated.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// ID of the availability zone.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type ZonalParameters struct {

	// ID of the subnet. If no ID is specified, and there only one subnet in specified zone, an address in this subnet will be allocated.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/vpc/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// ID of the availability zone.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ClusterInitParameters `json:"initProvider,omitempty"`
}

// ClusterStatus defines the observed state of Cluster.
type ClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Cluster is the Schema for the Clusters API. Allows management of Yandex Kubernetes Cluster. For more information, see
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud.upjet}
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.master) || (has(self.initProvider) && has(self.initProvider.master))",message="spec.forProvider.master is a required parameter"
	Spec   ClusterSpec   `json:"spec"`
	Status ClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterList contains a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

// Repository type metadata.
var (
	Cluster_Kind             = "Cluster"
	Cluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Cluster_Kind}.String()
	Cluster_KindAPIVersion   = Cluster_Kind + "." + CRDGroupVersion.String()
	Cluster_GroupVersionKind = CRDGroupVersion.WithKind(Cluster_Kind)
)

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}