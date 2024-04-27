// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	v1alpha1 "github.com/tagesjump/provider-upjet-yc/apis/resourcemanager/v1alpha1"
	v1alpha11 "github.com/tagesjump/provider-upjet-yc/apis/vpc/v1alpha1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Recordset.
func (mg *Recordset) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ZoneIDRef,
		Selector:     mg.Spec.ForProvider.ZoneIDSelector,
		To: reference.To{
			List:    &ZoneList{},
			Managed: &Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ZoneID")
	}
	mg.Spec.ForProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ZoneIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ZoneIDRef,
		Selector:     mg.Spec.InitProvider.ZoneIDSelector,
		To: reference.To{
			List:    &ZoneList{},
			Managed: &Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ZoneID")
	}
	mg.Spec.InitProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ZoneIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Zone.
func (mg *Zone) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FolderID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FolderIDRef,
		Selector:     mg.Spec.ForProvider.FolderIDSelector,
		To: reference.To{
			List:    &v1alpha1.FolderList{},
			Managed: &v1alpha1.Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FolderID")
	}
	mg.Spec.ForProvider.FolderID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FolderIDRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.PrivateNetworks),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.PrivateNetworksRefs,
		Selector:      mg.Spec.ForProvider.PrivateNetworksSelector,
		To: reference.To{
			List:    &v1alpha11.NetworkList{},
			Managed: &v1alpha11.Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PrivateNetworks")
	}
	mg.Spec.ForProvider.PrivateNetworks = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.PrivateNetworksRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.FolderID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.FolderIDRef,
		Selector:     mg.Spec.InitProvider.FolderIDSelector,
		To: reference.To{
			List:    &v1alpha1.FolderList{},
			Managed: &v1alpha1.Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.FolderID")
	}
	mg.Spec.InitProvider.FolderID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.FolderIDRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.PrivateNetworks),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.InitProvider.PrivateNetworksRefs,
		Selector:      mg.Spec.InitProvider.PrivateNetworksSelector,
		To: reference.To{
			List:    &v1alpha11.NetworkList{},
			Managed: &v1alpha11.Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.PrivateNetworks")
	}
	mg.Spec.InitProvider.PrivateNetworks = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.PrivateNetworksRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this ZoneIAMBinding.
func (mg *ZoneIAMBinding) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DNSZoneID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.DNSZoneIDRef,
		Selector:     mg.Spec.ForProvider.DNSZoneIDSelector,
		To: reference.To{
			List:    &ZoneList{},
			Managed: &Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DNSZoneID")
	}
	mg.Spec.ForProvider.DNSZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DNSZoneIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DNSZoneID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.DNSZoneIDRef,
		Selector:     mg.Spec.InitProvider.DNSZoneIDSelector,
		To: reference.To{
			List:    &ZoneList{},
			Managed: &Zone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.DNSZoneID")
	}
	mg.Spec.InitProvider.DNSZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.DNSZoneIDRef = rsp.ResolvedReference

	return nil
}
