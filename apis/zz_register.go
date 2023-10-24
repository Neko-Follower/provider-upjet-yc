/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha1 "github.com/tagesjump/provider-upjet-yc/apis/alb/v1alpha1"
	v1alpha1compute "github.com/tagesjump/provider-upjet-yc/apis/compute/v1alpha1"
	v1alpha1container "github.com/tagesjump/provider-upjet-yc/apis/container/v1alpha1"
	v1alpha1iam "github.com/tagesjump/provider-upjet-yc/apis/iam/v1alpha1"
	v1alpha1kms "github.com/tagesjump/provider-upjet-yc/apis/kms/v1alpha1"
	v1alpha1kubernetes "github.com/tagesjump/provider-upjet-yc/apis/kubernetes/v1alpha1"
	v1alpha1mdb "github.com/tagesjump/provider-upjet-yc/apis/mdb/v1alpha1"
	v1alpha1organizationmanager "github.com/tagesjump/provider-upjet-yc/apis/organizationmanager/v1alpha1"
	v1alpha1resourcemanager "github.com/tagesjump/provider-upjet-yc/apis/resourcemanager/v1alpha1"
	v1alpha1storage "github.com/tagesjump/provider-upjet-yc/apis/storage/v1alpha1"
	v1alpha1apis "github.com/tagesjump/provider-upjet-yc/apis/v1alpha1"
	v1beta1 "github.com/tagesjump/provider-upjet-yc/apis/v1beta1"
	v1alpha1vpc "github.com/tagesjump/provider-upjet-yc/apis/vpc/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1alpha1compute.SchemeBuilder.AddToScheme,
		v1alpha1container.SchemeBuilder.AddToScheme,
		v1alpha1iam.SchemeBuilder.AddToScheme,
		v1alpha1kms.SchemeBuilder.AddToScheme,
		v1alpha1kubernetes.SchemeBuilder.AddToScheme,
		v1alpha1mdb.SchemeBuilder.AddToScheme,
		v1alpha1organizationmanager.SchemeBuilder.AddToScheme,
		v1alpha1resourcemanager.SchemeBuilder.AddToScheme,
		v1alpha1storage.SchemeBuilder.AddToScheme,
		v1alpha1apis.SchemeBuilder.AddToScheme,
		v1beta1.SchemeBuilder.AddToScheme,
		v1alpha1vpc.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
