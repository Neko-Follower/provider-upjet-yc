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

type SecretCiphertextInitParameters struct {

	// Additional authenticated data (AAD context), optional. If specified, this data will be required for decryption with the SymmetricDecryptRequest
	AadContext *string `json:"aadContext,omitempty" tf:"aad_context,omitempty"`
}

type SecretCiphertextObservation struct {

	// Additional authenticated data (AAD context), optional. If specified, this data will be required for decryption with the SymmetricDecryptRequest
	AadContext *string `json:"aadContext,omitempty" tf:"aad_context,omitempty"`

	// Resulting ciphertext, encoded with "standard" base64 alphabet as defined in RFC 4648 section 4
	Ciphertext *string `json:"ciphertext,omitempty" tf:"ciphertext,omitempty"`

	// an identifier for the resource with format {{key_id}}/{{ciphertext}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ID of the symmetric KMS key to use for encryption.
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`
}

type SecretCiphertextParameters struct {

	// Additional authenticated data (AAD context), optional. If specified, this data will be required for decryption with the SymmetricDecryptRequest
	// +kubebuilder:validation:Optional
	AadContext *string `json:"aadContext,omitempty" tf:"aad_context,omitempty"`

	// ID of the symmetric KMS key to use for encryption.
	// +crossplane:generate:reference:type=SymmetricKey
	// +kubebuilder:validation:Optional
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`

	// Reference to a SymmetricKey to populate keyId.
	// +kubebuilder:validation:Optional
	KeyIDRef *v1.Reference `json:"keyIdRef,omitempty" tf:"-"`

	// Selector for a SymmetricKey to populate keyId.
	// +kubebuilder:validation:Optional
	KeyIDSelector *v1.Selector `json:"keyIdSelector,omitempty" tf:"-"`

	// Plaintext to be encrypted.
	// +kubebuilder:validation:Optional
	PlaintextSecretRef v1.SecretKeySelector `json:"plaintextSecretRef" tf:"-"`
}

// SecretCiphertextSpec defines the desired state of SecretCiphertext
type SecretCiphertextSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecretCiphertextParameters `json:"forProvider"`
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
	InitProvider SecretCiphertextInitParameters `json:"initProvider,omitempty"`
}

// SecretCiphertextStatus defines the observed state of SecretCiphertext.
type SecretCiphertextStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecretCiphertextObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecretCiphertext is the Schema for the SecretCiphertexts API. Encrypts given plaintext with the specified Yandex KMS key and provides access to the ciphertext.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud.upjet}
type SecretCiphertext struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.plaintextSecretRef)",message="spec.forProvider.plaintextSecretRef is a required parameter"
	Spec   SecretCiphertextSpec   `json:"spec"`
	Status SecretCiphertextStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretCiphertextList contains a list of SecretCiphertexts
type SecretCiphertextList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretCiphertext `json:"items"`
}

// Repository type metadata.
var (
	SecretCiphertext_Kind             = "SecretCiphertext"
	SecretCiphertext_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecretCiphertext_Kind}.String()
	SecretCiphertext_KindAPIVersion   = SecretCiphertext_Kind + "." + CRDGroupVersion.String()
	SecretCiphertext_GroupVersionKind = CRDGroupVersion.WithKind(SecretCiphertext_Kind)
)

func init() {
	SchemeBuilder.Register(&SecretCiphertext{}, &SecretCiphertextList{})
}
