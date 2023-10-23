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

type SnapshotInitParameters struct {

	// Description of the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A set of key/value label pairs to assign to the snapshot.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// A name for the resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type SnapshotObservation struct {

	// Creation timestamp of the snapshot.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Description of the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Size of the disk when the snapshot was created, specified in GB.
	DiskSize *float64 `json:"diskSize,omitempty" tf:"disk_size,omitempty"`

	// The ID of the folder that the resource belongs to. If it
	// is not provided, the default provider folder is used.
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A set of key/value label pairs to assign to the snapshot.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// A name for the resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the disk to create a snapshot from.
	SourceDiskID *string `json:"sourceDiskId,omitempty" tf:"source_disk_id,omitempty"`

	// Size of the snapshot, specified in GB.
	StorageSize *float64 `json:"storageSize,omitempty" tf:"storage_size,omitempty"`
}

type SnapshotParameters struct {

	// Description of the resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the folder that the resource belongs to. If it
	// is not provided, the default provider folder is used.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/resourcemanager/v1alpha1.Folder
	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// A set of key/value label pairs to assign to the snapshot.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// A name for the resource.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the disk to create a snapshot from.
	// +crossplane:generate:reference:type=Disk
	// +kubebuilder:validation:Optional
	SourceDiskID *string `json:"sourceDiskId,omitempty" tf:"source_disk_id,omitempty"`

	// Reference to a Disk to populate sourceDiskId.
	// +kubebuilder:validation:Optional
	SourceDiskIDRef *v1.Reference `json:"sourceDiskIdRef,omitempty" tf:"-"`

	// Selector for a Disk to populate sourceDiskId.
	// +kubebuilder:validation:Optional
	SourceDiskIDSelector *v1.Selector `json:"sourceDiskIdSelector,omitempty" tf:"-"`
}

// SnapshotSpec defines the desired state of Snapshot
type SnapshotSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SnapshotParameters `json:"forProvider"`
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
	InitProvider SnapshotInitParameters `json:"initProvider,omitempty"`
}

// SnapshotStatus defines the observed state of Snapshot.
type SnapshotStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SnapshotObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Snapshot is the Schema for the Snapshots API. Creates a new snapshot of a disk.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud.upjet}
type Snapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SnapshotSpec   `json:"spec"`
	Status            SnapshotStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SnapshotList contains a list of Snapshots
type SnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Snapshot `json:"items"`
}

// Repository type metadata.
var (
	Snapshot_Kind             = "Snapshot"
	Snapshot_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Snapshot_Kind}.String()
	Snapshot_KindAPIVersion   = Snapshot_Kind + "." + CRDGroupVersion.String()
	Snapshot_GroupVersionKind = CRDGroupVersion.WithKind(Snapshot_Kind)
)

func init() {
	SchemeBuilder.Register(&Snapshot{}, &SnapshotList{})
}
