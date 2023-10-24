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

type QueueInitParameters struct {

	// Enables content-based deduplication. Can be used only if queue is FIFO.
	ContentBasedDeduplication *bool `json:"contentBasedDeduplication,omitempty" tf:"content_based_deduplication,omitempty"`

	// Number of seconds to delay the message from being available for processing. Valid values: from 0 to 900 seconds (15 minutes). Default: 0.
	DelaySeconds *float64 `json:"delaySeconds,omitempty" tf:"delay_seconds,omitempty"`

	// Is this queue FIFO. If this parameter is not used, a standard queue is created. You cannot change the parameter value for a created queue.
	FifoQueue *bool `json:"fifoQueue,omitempty" tf:"fifo_queue,omitempty"`

	// Maximum message size in bytes. Valid values: from 1024 bytes (1 KB) to 262144 bytes (256 KB). Default: 262144 (256 KB). For more information see documentation.
	MaxMessageSize *float64 `json:"maxMessageSize,omitempty" tf:"max_message_size,omitempty"`

	// The length of time in seconds to retain a message. Valid values: from 60 seconds (1 minute) to 1209600 seconds (14 days). Default: 345600 (4 days). For more information see documentation.
	MessageRetentionSeconds *float64 `json:"messageRetentionSeconds,omitempty" tf:"message_retention_seconds,omitempty"`

	// Queue name. The maximum length is 80 characters. You can use numbers, letters, underscores, and hyphens in the name. The name of a FIFO queue must end with the .fifo suffix. If not specified, random name will be generated. Conflicts with name_prefix. For more information see documentation.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Generates random name with the specified prefix. Conflicts with name.
	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`

	// Wait time for the ReceiveMessage method (for long polling), in seconds. Valid values: from 0 to 20 seconds. Default: 0. For more information about long polling see documentation.
	ReceiveWaitTimeSeconds *float64 `json:"receiveWaitTimeSeconds,omitempty" tf:"receive_wait_time_seconds,omitempty"`

	// Message redrive policy in Dead Letter Queue. The source queue and DLQ must be the same type: for FIFO queues, the DLQ must also be a FIFO queue. For more information about redrive policy see documentation. Also you can use example in this page.
	RedrivePolicy *string `json:"redrivePolicy,omitempty" tf:"redrive_policy,omitempty"`

	// ID of the region where the message queue is located at.
	// The default is 'ru-central1'.
	RegionID *string `json:"regionId,omitempty" tf:"region_id,omitempty"`

	// Visibility timeout for messages in a queue, specified in seconds. Valid values: from 0 to 43200 seconds (12 hours). Default: 30.
	VisibilityTimeoutSeconds *float64 `json:"visibilityTimeoutSeconds,omitempty" tf:"visibility_timeout_seconds,omitempty"`
}

type QueueObservation struct {

	// The access key to use when applying changes. If omitted, ymq_access_key specified in provider config is used. For more information see documentation.
	AccessKey *string `json:"accessKey,omitempty" tf:"access_key,omitempty"`

	// ARN of the Yandex Message Queue. It is used for setting up a redrive policy. See documentation.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Enables content-based deduplication. Can be used only if queue is FIFO.
	ContentBasedDeduplication *bool `json:"contentBasedDeduplication,omitempty" tf:"content_based_deduplication,omitempty"`

	// Number of seconds to delay the message from being available for processing. Valid values: from 0 to 900 seconds (15 minutes). Default: 0.
	DelaySeconds *float64 `json:"delaySeconds,omitempty" tf:"delay_seconds,omitempty"`

	// Is this queue FIFO. If this parameter is not used, a standard queue is created. You cannot change the parameter value for a created queue.
	FifoQueue *bool `json:"fifoQueue,omitempty" tf:"fifo_queue,omitempty"`

	// URL of the Yandex Message Queue.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Maximum message size in bytes. Valid values: from 1024 bytes (1 KB) to 262144 bytes (256 KB). Default: 262144 (256 KB). For more information see documentation.
	MaxMessageSize *float64 `json:"maxMessageSize,omitempty" tf:"max_message_size,omitempty"`

	// The length of time in seconds to retain a message. Valid values: from 60 seconds (1 minute) to 1209600 seconds (14 days). Default: 345600 (4 days). For more information see documentation.
	MessageRetentionSeconds *float64 `json:"messageRetentionSeconds,omitempty" tf:"message_retention_seconds,omitempty"`

	// Queue name. The maximum length is 80 characters. You can use numbers, letters, underscores, and hyphens in the name. The name of a FIFO queue must end with the .fifo suffix. If not specified, random name will be generated. Conflicts with name_prefix. For more information see documentation.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Generates random name with the specified prefix. Conflicts with name.
	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`

	// Wait time for the ReceiveMessage method (for long polling), in seconds. Valid values: from 0 to 20 seconds. Default: 0. For more information about long polling see documentation.
	ReceiveWaitTimeSeconds *float64 `json:"receiveWaitTimeSeconds,omitempty" tf:"receive_wait_time_seconds,omitempty"`

	// Message redrive policy in Dead Letter Queue. The source queue and DLQ must be the same type: for FIFO queues, the DLQ must also be a FIFO queue. For more information about redrive policy see documentation. Also you can use example in this page.
	RedrivePolicy *string `json:"redrivePolicy,omitempty" tf:"redrive_policy,omitempty"`

	// ID of the region where the message queue is located at.
	// The default is 'ru-central1'.
	RegionID *string `json:"regionId,omitempty" tf:"region_id,omitempty"`

	// Visibility timeout for messages in a queue, specified in seconds. Valid values: from 0 to 43200 seconds (12 hours). Default: 30.
	VisibilityTimeoutSeconds *float64 `json:"visibilityTimeoutSeconds,omitempty" tf:"visibility_timeout_seconds,omitempty"`
}

type QueueParameters struct {

	// The access key to use when applying changes. If omitted, ymq_access_key specified in provider config is used. For more information see documentation.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/iam/v1alpha1.ServiceAccountStaticAccessKey
	// +crossplane:generate:reference:extractor=github.com/tagesjump/provider-upjet-yc/config/common.ExtractAccessKey()
	// +kubebuilder:validation:Optional
	AccessKey *string `json:"accessKey,omitempty" tf:"access_key,omitempty"`

	// Reference to a ServiceAccountStaticAccessKey in iam to populate accessKey.
	// +kubebuilder:validation:Optional
	AccessKeyRef *v1.Reference `json:"accessKeyRef,omitempty" tf:"-"`

	// Selector for a ServiceAccountStaticAccessKey in iam to populate accessKey.
	// +kubebuilder:validation:Optional
	AccessKeySelector *v1.Selector `json:"accessKeySelector,omitempty" tf:"-"`

	// Enables content-based deduplication. Can be used only if queue is FIFO.
	// +kubebuilder:validation:Optional
	ContentBasedDeduplication *bool `json:"contentBasedDeduplication,omitempty" tf:"content_based_deduplication,omitempty"`

	// Number of seconds to delay the message from being available for processing. Valid values: from 0 to 900 seconds (15 minutes). Default: 0.
	// +kubebuilder:validation:Optional
	DelaySeconds *float64 `json:"delaySeconds,omitempty" tf:"delay_seconds,omitempty"`

	// Is this queue FIFO. If this parameter is not used, a standard queue is created. You cannot change the parameter value for a created queue.
	// +kubebuilder:validation:Optional
	FifoQueue *bool `json:"fifoQueue,omitempty" tf:"fifo_queue,omitempty"`

	// Maximum message size in bytes. Valid values: from 1024 bytes (1 KB) to 262144 bytes (256 KB). Default: 262144 (256 KB). For more information see documentation.
	// +kubebuilder:validation:Optional
	MaxMessageSize *float64 `json:"maxMessageSize,omitempty" tf:"max_message_size,omitempty"`

	// The length of time in seconds to retain a message. Valid values: from 60 seconds (1 minute) to 1209600 seconds (14 days). Default: 345600 (4 days). For more information see documentation.
	// +kubebuilder:validation:Optional
	MessageRetentionSeconds *float64 `json:"messageRetentionSeconds,omitempty" tf:"message_retention_seconds,omitempty"`

	// Queue name. The maximum length is 80 characters. You can use numbers, letters, underscores, and hyphens in the name. The name of a FIFO queue must end with the .fifo suffix. If not specified, random name will be generated. Conflicts with name_prefix. For more information see documentation.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Generates random name with the specified prefix. Conflicts with name.
	// +kubebuilder:validation:Optional
	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`

	// Wait time for the ReceiveMessage method (for long polling), in seconds. Valid values: from 0 to 20 seconds. Default: 0. For more information about long polling see documentation.
	// +kubebuilder:validation:Optional
	ReceiveWaitTimeSeconds *float64 `json:"receiveWaitTimeSeconds,omitempty" tf:"receive_wait_time_seconds,omitempty"`

	// Message redrive policy in Dead Letter Queue. The source queue and DLQ must be the same type: for FIFO queues, the DLQ must also be a FIFO queue. For more information about redrive policy see documentation. Also you can use example in this page.
	// +kubebuilder:validation:Optional
	RedrivePolicy *string `json:"redrivePolicy,omitempty" tf:"redrive_policy,omitempty"`

	// ID of the region where the message queue is located at.
	// The default is 'ru-central1'.
	// +kubebuilder:validation:Optional
	RegionID *string `json:"regionId,omitempty" tf:"region_id,omitempty"`

	// The secret key to use when applying changes. If omitted, ymq_secret_key specified in provider config is used. For more information see documentation.
	// +kubebuilder:validation:Optional
	SecretKeySecretRef *v1.SecretKeySelector `json:"secretKeySecretRef,omitempty" tf:"-"`

	// Visibility timeout for messages in a queue, specified in seconds. Valid values: from 0 to 43200 seconds (12 hours). Default: 30.
	// +kubebuilder:validation:Optional
	VisibilityTimeoutSeconds *float64 `json:"visibilityTimeoutSeconds,omitempty" tf:"visibility_timeout_seconds,omitempty"`
}

// QueueSpec defines the desired state of Queue
type QueueSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     QueueParameters `json:"forProvider"`
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
	InitProvider QueueInitParameters `json:"initProvider,omitempty"`
}

// QueueStatus defines the observed state of Queue.
type QueueStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        QueueObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Queue is the Schema for the Queues API. Allows management of a Yandex.Cloud Message Queue.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type Queue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              QueueSpec   `json:"spec"`
	Status            QueueStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// QueueList contains a list of Queues
type QueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Queue `json:"items"`
}

// Repository type metadata.
var (
	Queue_Kind             = "Queue"
	Queue_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Queue_Kind}.String()
	Queue_KindAPIVersion   = Queue_Kind + "." + CRDGroupVersion.String()
	Queue_GroupVersionKind = CRDGroupVersion.WithKind(Queue_Kind)
)

func init() {
	SchemeBuilder.Register(&Queue{}, &QueueList{})
}
