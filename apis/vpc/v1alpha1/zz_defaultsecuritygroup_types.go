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

type DefaultSecurityGroupInitParameters struct {

	// Description of the security group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A list of egress rules. The structure is documented below.
	Egress []EgressInitParameters `json:"egress,omitempty" tf:"egress,omitempty"`

	// A list of ingress rules.
	Ingress []IngressInitParameters `json:"ingress,omitempty" tf:"ingress,omitempty"`

	// Labels to assign to this security group.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`
}

type DefaultSecurityGroupObservation struct {

	// Creation timestamp of this security group.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Description of the security group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A list of egress rules. The structure is documented below.
	Egress []EgressObservation `json:"egress,omitempty" tf:"egress,omitempty"`

	// ID of the folder this security group belongs to.
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Id of the security group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A list of ingress rules.
	Ingress []IngressObservation `json:"ingress,omitempty" tf:"ingress,omitempty"`

	// Labels to assign to this security group.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name of this security group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the network this security group belongs to.
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Status of this security group.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type DefaultSecurityGroupParameters struct {

	// Description of the security group.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A list of egress rules. The structure is documented below.
	// +kubebuilder:validation:Optional
	Egress []EgressParameters `json:"egress,omitempty" tf:"egress,omitempty"`

	// ID of the folder this security group belongs to.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/resourcemanager/v1alpha1.Folder
	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// A list of ingress rules.
	// +kubebuilder:validation:Optional
	Ingress []IngressParameters `json:"ingress,omitempty" tf:"ingress,omitempty"`

	// Labels to assign to this security group.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// ID of the network this security group belongs to.
	// +crossplane:generate:reference:type=Network
	// +kubebuilder:validation:Optional
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a Network to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a Network to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`
}

type EgressInitParameters struct {

	// Description of the rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Minimum port number.
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// Labels to assign to this security group.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Port number (if applied to a single port).
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Special-purpose targets. self_security_group refers to this particular security group. loadbalancer_healthchecks represents loadbalancer health check nodes.
	PredefinedTarget *string `json:"predefinedTarget,omitempty" tf:"predefined_target,omitempty"`

	// One of ANY, TCP, UDP, ICMP, IPV6_ICMP.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Target security group ID for this rule.
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// Maximum port number.
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`

	// The blocks of IPv4 addresses for this rule.
	V4CidrBlocks []*string `json:"v4CidrBlocks,omitempty" tf:"v4_cidr_blocks,omitempty"`

	// The blocks of IPv6 addresses for this rule. v6_cidr_blocks argument is currently not supported. It will be available in the future.
	V6CidrBlocks []*string `json:"v6CidrBlocks,omitempty" tf:"v6_cidr_blocks,omitempty"`
}

type EgressObservation struct {

	// Description of the rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Minimum port number.
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// Id of the security group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Labels to assign to this security group.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Port number (if applied to a single port).
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Special-purpose targets. self_security_group refers to this particular security group. loadbalancer_healthchecks represents loadbalancer health check nodes.
	PredefinedTarget *string `json:"predefinedTarget,omitempty" tf:"predefined_target,omitempty"`

	// One of ANY, TCP, UDP, ICMP, IPV6_ICMP.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Target security group ID for this rule.
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// Maximum port number.
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`

	// The blocks of IPv4 addresses for this rule.
	V4CidrBlocks []*string `json:"v4CidrBlocks,omitempty" tf:"v4_cidr_blocks,omitempty"`

	// The blocks of IPv6 addresses for this rule. v6_cidr_blocks argument is currently not supported. It will be available in the future.
	V6CidrBlocks []*string `json:"v6CidrBlocks,omitempty" tf:"v6_cidr_blocks,omitempty"`
}

type EgressParameters struct {

	// Description of the rule.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Minimum port number.
	// +kubebuilder:validation:Optional
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// Labels to assign to this security group.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Port number (if applied to a single port).
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Special-purpose targets. self_security_group refers to this particular security group. loadbalancer_healthchecks represents loadbalancer health check nodes.
	// +kubebuilder:validation:Optional
	PredefinedTarget *string `json:"predefinedTarget,omitempty" tf:"predefined_target,omitempty"`

	// One of ANY, TCP, UDP, ICMP, IPV6_ICMP.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// Target security group ID for this rule.
	// +kubebuilder:validation:Optional
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// Maximum port number.
	// +kubebuilder:validation:Optional
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`

	// The blocks of IPv4 addresses for this rule.
	// +kubebuilder:validation:Optional
	V4CidrBlocks []*string `json:"v4CidrBlocks,omitempty" tf:"v4_cidr_blocks,omitempty"`

	// The blocks of IPv6 addresses for this rule. v6_cidr_blocks argument is currently not supported. It will be available in the future.
	// +kubebuilder:validation:Optional
	V6CidrBlocks []*string `json:"v6CidrBlocks,omitempty" tf:"v6_cidr_blocks,omitempty"`
}

type IngressInitParameters struct {

	// Description of the rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Minimum port number.
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// Labels to assign to this rule.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Port number (if applied to a single port).
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Special-purpose targets. self_security_group refers to this particular security group. loadbalancer_healthchecks represents loadbalancer health check nodes.
	PredefinedTarget *string `json:"predefinedTarget,omitempty" tf:"predefined_target,omitempty"`

	// One of ANY, TCP, UDP, ICMP, IPV6_ICMP.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Target security group ID for this rule.
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// Maximum port number.
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`

	// The blocks of IPv4 addresses for this rule.
	V4CidrBlocks []*string `json:"v4CidrBlocks,omitempty" tf:"v4_cidr_blocks,omitempty"`

	// The blocks of IPv6 addresses for this rule. v6_cidr_blocks argument is currently not supported. It will be available in the future.
	V6CidrBlocks []*string `json:"v6CidrBlocks,omitempty" tf:"v6_cidr_blocks,omitempty"`
}

type IngressObservation struct {

	// Description of the rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Minimum port number.
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// Id of the security group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Labels to assign to this rule.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Port number (if applied to a single port).
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Special-purpose targets. self_security_group refers to this particular security group. loadbalancer_healthchecks represents loadbalancer health check nodes.
	PredefinedTarget *string `json:"predefinedTarget,omitempty" tf:"predefined_target,omitempty"`

	// One of ANY, TCP, UDP, ICMP, IPV6_ICMP.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Target security group ID for this rule.
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// Maximum port number.
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`

	// The blocks of IPv4 addresses for this rule.
	V4CidrBlocks []*string `json:"v4CidrBlocks,omitempty" tf:"v4_cidr_blocks,omitempty"`

	// The blocks of IPv6 addresses for this rule. v6_cidr_blocks argument is currently not supported. It will be available in the future.
	V6CidrBlocks []*string `json:"v6CidrBlocks,omitempty" tf:"v6_cidr_blocks,omitempty"`
}

type IngressParameters struct {

	// Description of the rule.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Minimum port number.
	// +kubebuilder:validation:Optional
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// Labels to assign to this rule.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Port number (if applied to a single port).
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Special-purpose targets. self_security_group refers to this particular security group. loadbalancer_healthchecks represents loadbalancer health check nodes.
	// +kubebuilder:validation:Optional
	PredefinedTarget *string `json:"predefinedTarget,omitempty" tf:"predefined_target,omitempty"`

	// One of ANY, TCP, UDP, ICMP, IPV6_ICMP.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// Target security group ID for this rule.
	// +kubebuilder:validation:Optional
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// Maximum port number.
	// +kubebuilder:validation:Optional
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`

	// The blocks of IPv4 addresses for this rule.
	// +kubebuilder:validation:Optional
	V4CidrBlocks []*string `json:"v4CidrBlocks,omitempty" tf:"v4_cidr_blocks,omitempty"`

	// The blocks of IPv6 addresses for this rule. v6_cidr_blocks argument is currently not supported. It will be available in the future.
	// +kubebuilder:validation:Optional
	V6CidrBlocks []*string `json:"v6CidrBlocks,omitempty" tf:"v6_cidr_blocks,omitempty"`
}

// DefaultSecurityGroupSpec defines the desired state of DefaultSecurityGroup
type DefaultSecurityGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DefaultSecurityGroupParameters `json:"forProvider"`
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
	InitProvider DefaultSecurityGroupInitParameters `json:"initProvider,omitempty"`
}

// DefaultSecurityGroupStatus defines the observed state of DefaultSecurityGroup.
type DefaultSecurityGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DefaultSecurityGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DefaultSecurityGroup is the Schema for the DefaultSecurityGroups API. Yandex VPC Default Security Group.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud.upjet}
type DefaultSecurityGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DefaultSecurityGroupSpec   `json:"spec"`
	Status            DefaultSecurityGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DefaultSecurityGroupList contains a list of DefaultSecurityGroups
type DefaultSecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DefaultSecurityGroup `json:"items"`
}

// Repository type metadata.
var (
	DefaultSecurityGroup_Kind             = "DefaultSecurityGroup"
	DefaultSecurityGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DefaultSecurityGroup_Kind}.String()
	DefaultSecurityGroup_KindAPIVersion   = DefaultSecurityGroup_Kind + "." + CRDGroupVersion.String()
	DefaultSecurityGroup_GroupVersionKind = CRDGroupVersion.WithKind(DefaultSecurityGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&DefaultSecurityGroup{}, &DefaultSecurityGroupList{})
}
