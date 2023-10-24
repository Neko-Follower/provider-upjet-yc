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

type RecordsetInitParameters struct {

	// The string data for the records in this record set.
	Data []*string `json:"data,omitempty" tf:"data,omitempty"`

	// The DNS name this record set will apply to.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The time-to-live of this record set (seconds).
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// The DNS record set type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RecordsetObservation struct {

	// The string data for the records in this record set.
	Data []*string `json:"data,omitempty" tf:"data,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The DNS name this record set will apply to.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The time-to-live of this record set (seconds).
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// The DNS record set type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The id of the zone in which this record set will reside.
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type RecordsetParameters struct {

	// The string data for the records in this record set.
	// +kubebuilder:validation:Optional
	Data []*string `json:"data,omitempty" tf:"data,omitempty"`

	// The DNS name this record set will apply to.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The time-to-live of this record set (seconds).
	// +kubebuilder:validation:Optional
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// The DNS record set type.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The id of the zone in which this record set will reside.
	// +crossplane:generate:reference:type=Zone
	// +kubebuilder:validation:Optional
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`

	// Reference to a Zone to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDRef *v1.Reference `json:"zoneIdRef,omitempty" tf:"-"`

	// Selector for a Zone to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDSelector *v1.Selector `json:"zoneIdSelector,omitempty" tf:"-"`
}

// RecordsetSpec defines the desired state of Recordset
type RecordsetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RecordsetParameters `json:"forProvider"`
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
	InitProvider RecordsetInitParameters `json:"initProvider,omitempty"`
}

// RecordsetStatus defines the observed state of Recordset.
type RecordsetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RecordsetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Recordset is the Schema for the Recordsets API. Manages a DNS Recordset within Yandex.Cloud.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type Recordset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.data) || (has(self.initProvider) && has(self.initProvider.data))",message="spec.forProvider.data is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ttl) || (has(self.initProvider) && has(self.initProvider.ttl))",message="spec.forProvider.ttl is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   RecordsetSpec   `json:"spec"`
	Status RecordsetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RecordsetList contains a list of Recordsets
type RecordsetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Recordset `json:"items"`
}

// Repository type metadata.
var (
	Recordset_Kind             = "Recordset"
	Recordset_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Recordset_Kind}.String()
	Recordset_KindAPIVersion   = Recordset_Kind + "." + CRDGroupVersion.String()
	Recordset_GroupVersionKind = CRDGroupVersion.WithKind(Recordset_Kind)
)

func init() {
	SchemeBuilder.Register(&Recordset{}, &RecordsetList{})
}
