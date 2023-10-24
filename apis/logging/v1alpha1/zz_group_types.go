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

type GroupInitParameters struct {
	DataStream *string `json:"dataStream,omitempty" tf:"data_stream,omitempty"`

	// A description for the Yandex Cloud Logging group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A set of key/value label pairs to assign to the Yandex Cloud Logging group.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name for the Yandex Cloud Logging group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Log entries retention period for the Yandex Cloud Logging group.
	RetentionPeriod *string `json:"retentionPeriod,omitempty" tf:"retention_period,omitempty"`
}

type GroupObservation struct {

	// ID of the cloud that the Yandex Cloud Logging group belong to.
	CloudID *string `json:"cloudId,omitempty" tf:"cloud_id,omitempty"`

	// The Yandex Cloud Logging group creation timestamp.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	DataStream *string `json:"dataStream,omitempty" tf:"data_stream,omitempty"`

	// A description for the Yandex Cloud Logging group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// ID of the folder that the Yandex Cloud Logging group belongs to.
	// It will be deduced from provider configuration if not set explicitly.
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// The Yandex Cloud Logging group ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A set of key/value label pairs to assign to the Yandex Cloud Logging group.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name for the Yandex Cloud Logging group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Log entries retention period for the Yandex Cloud Logging group.
	RetentionPeriod *string `json:"retentionPeriod,omitempty" tf:"retention_period,omitempty"`

	// The Yandex Cloud Logging group status.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type GroupParameters struct {

	// +kubebuilder:validation:Optional
	DataStream *string `json:"dataStream,omitempty" tf:"data_stream,omitempty"`

	// A description for the Yandex Cloud Logging group.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// ID of the folder that the Yandex Cloud Logging group belongs to.
	// It will be deduced from provider configuration if not set explicitly.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/resourcemanager/v1alpha1.Folder
	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// A set of key/value label pairs to assign to the Yandex Cloud Logging group.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name for the Yandex Cloud Logging group.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Log entries retention period for the Yandex Cloud Logging group.
	// +kubebuilder:validation:Optional
	RetentionPeriod *string `json:"retentionPeriod,omitempty" tf:"retention_period,omitempty"`
}

// GroupSpec defines the desired state of Group
type GroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GroupParameters `json:"forProvider"`
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
	InitProvider GroupInitParameters `json:"initProvider,omitempty"`
}

// GroupStatus defines the observed state of Group.
type GroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Group is the Schema for the Groups API. Manages Yandex Cloud Logging group.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type Group struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GroupSpec   `json:"spec"`
	Status            GroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GroupList contains a list of Groups
type GroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Group `json:"items"`
}

// Repository type metadata.
var (
	Group_Kind             = "Group"
	Group_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Group_Kind}.String()
	Group_KindAPIVersion   = Group_Kind + "." + CRDGroupVersion.String()
	Group_GroupVersionKind = CRDGroupVersion.WithKind(Group_Kind)
)

func init() {
	SchemeBuilder.Register(&Group{}, &GroupList{})
}
