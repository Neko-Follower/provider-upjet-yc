/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1alpha12 "github.com/tagesjump/provider-upjet-yc/apis/iam/v1alpha1"
	v1alpha1 "github.com/tagesjump/provider-upjet-yc/apis/resourcemanager/v1alpha1"
	v1alpha11 "github.com/tagesjump/provider-upjet-yc/apis/vpc/v1alpha1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this ClickhouseCluster.
func (mg *ClickhouseCluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
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

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Host); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Host[i3].SubnetID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.Host[i3].SubnetIDRef,
			Selector:     mg.Spec.ForProvider.Host[i3].SubnetIDSelector,
			To: reference.To{
				List:    &v1alpha11.SubnetList{},
				Managed: &v1alpha11.Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Host[i3].SubnetID")
		}
		mg.Spec.ForProvider.Host[i3].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Host[i3].SubnetIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NetworkIDRef,
		Selector:     mg.Spec.ForProvider.NetworkIDSelector,
		To: reference.To{
			List:    &v1alpha11.NetworkList{},
			Managed: &v1alpha11.Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NetworkID")
	}
	mg.Spec.ForProvider.NetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetworkIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceAccountID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ServiceAccountIDRef,
		Selector:     mg.Spec.ForProvider.ServiceAccountIDSelector,
		To: reference.To{
			List:    &v1alpha12.ServiceAccountList{},
			Managed: &v1alpha12.ServiceAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServiceAccountID")
	}
	mg.Spec.ForProvider.ServiceAccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceAccountIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ElasticsearchCluster.
func (mg *ElasticsearchCluster) ResolveReferences(ctx context.Context, c client.Reader) error {
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

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Host); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Host[i3].SubnetID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.Host[i3].SubnetIDRef,
			Selector:     mg.Spec.ForProvider.Host[i3].SubnetIDSelector,
			To: reference.To{
				List:    &v1alpha11.SubnetList{},
				Managed: &v1alpha11.Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Host[i3].SubnetID")
		}
		mg.Spec.ForProvider.Host[i3].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Host[i3].SubnetIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NetworkIDRef,
		Selector:     mg.Spec.ForProvider.NetworkIDSelector,
		To: reference.To{
			List:    &v1alpha11.NetworkList{},
			Managed: &v1alpha11.Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NetworkID")
	}
	mg.Spec.ForProvider.NetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetworkIDRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SecurityGroupIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SecurityGroupIdsRefs,
		Selector:      mg.Spec.ForProvider.SecurityGroupIdsSelector,
		To: reference.To{
			List:    &v1alpha11.SecurityGroupList{},
			Managed: &v1alpha11.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroupIds")
	}
	mg.Spec.ForProvider.SecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SecurityGroupIdsRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceAccountID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ServiceAccountIDRef,
		Selector:     mg.Spec.ForProvider.ServiceAccountIDSelector,
		To: reference.To{
			List:    &v1alpha12.ServiceAccountList{},
			Managed: &v1alpha12.ServiceAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServiceAccountID")
	}
	mg.Spec.ForProvider.ServiceAccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceAccountIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this GreenplumCluster.
func (mg *GreenplumCluster) ResolveReferences(ctx context.Context, c client.Reader) error {
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

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NetworkIDRef,
		Selector:     mg.Spec.ForProvider.NetworkIDSelector,
		To: reference.To{
			List:    &v1alpha11.NetworkList{},
			Managed: &v1alpha11.Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NetworkID")
	}
	mg.Spec.ForProvider.NetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetworkIDRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SecurityGroupIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SecurityGroupIdsRefs,
		Selector:      mg.Spec.ForProvider.SecurityGroupIdsSelector,
		To: reference.To{
			List:    &v1alpha11.SecurityGroupList{},
			Managed: &v1alpha11.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroupIds")
	}
	mg.Spec.ForProvider.SecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SecurityGroupIdsRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubnetID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SubnetIDRef,
		Selector:     mg.Spec.ForProvider.SubnetIDSelector,
		To: reference.To{
			List:    &v1alpha11.SubnetList{},
			Managed: &v1alpha11.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetID")
	}
	mg.Spec.ForProvider.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubnetIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this KafkaCluster.
func (mg *KafkaCluster) ResolveReferences(ctx context.Context, c client.Reader) error {
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

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NetworkIDRef,
		Selector:     mg.Spec.ForProvider.NetworkIDSelector,
		To: reference.To{
			List:    &v1alpha11.NetworkList{},
			Managed: &v1alpha11.Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NetworkID")
	}
	mg.Spec.ForProvider.NetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetworkIDRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SecurityGroupIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SecurityGroupIdsRefs,
		Selector:      mg.Spec.ForProvider.SecurityGroupIdsSelector,
		To: reference.To{
			List:    &v1alpha11.SecurityGroupList{},
			Managed: &v1alpha11.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroupIds")
	}
	mg.Spec.ForProvider.SecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SecurityGroupIdsRefs = mrsp.ResolvedReferences

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SubnetIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SubnetIdsRefs,
		Selector:      mg.Spec.ForProvider.SubnetIdsSelector,
		To: reference.To{
			List:    &v1alpha11.SubnetList{},
			Managed: &v1alpha11.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetIds")
	}
	mg.Spec.ForProvider.SubnetIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SubnetIdsRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this KafkaConnector.
func (mg *KafkaConnector) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterIDRef,
		Selector:     mg.Spec.ForProvider.ClusterIDSelector,
		To: reference.To{
			List:    &KafkaClusterList{},
			Managed: &KafkaCluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterID")
	}
	mg.Spec.ForProvider.ClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this KafkaTopic.
func (mg *KafkaTopic) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterIDRef,
		Selector:     mg.Spec.ForProvider.ClusterIDSelector,
		To: reference.To{
			List:    &KafkaClusterList{},
			Managed: &KafkaCluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterID")
	}
	mg.Spec.ForProvider.ClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this KafkaUser.
func (mg *KafkaUser) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterIDRef,
		Selector:     mg.Spec.ForProvider.ClusterIDSelector,
		To: reference.To{
			List:    &KafkaClusterList{},
			Managed: &KafkaCluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterID")
	}
	mg.Spec.ForProvider.ClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this MongodbCluster.
func (mg *MongodbCluster) ResolveReferences(ctx context.Context, c client.Reader) error {
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

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Host); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Host[i3].SubnetID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.Host[i3].SubnetIDRef,
			Selector:     mg.Spec.ForProvider.Host[i3].SubnetIDSelector,
			To: reference.To{
				List:    &v1alpha11.SubnetList{},
				Managed: &v1alpha11.Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Host[i3].SubnetID")
		}
		mg.Spec.ForProvider.Host[i3].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Host[i3].SubnetIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NetworkIDRef,
		Selector:     mg.Spec.ForProvider.NetworkIDSelector,
		To: reference.To{
			List:    &v1alpha11.NetworkList{},
			Managed: &v1alpha11.Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NetworkID")
	}
	mg.Spec.ForProvider.NetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetworkIDRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SecurityGroupIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SecurityGroupIdsRefs,
		Selector:      mg.Spec.ForProvider.SecurityGroupIdsSelector,
		To: reference.To{
			List:    &v1alpha11.SecurityGroupList{},
			Managed: &v1alpha11.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroupIds")
	}
	mg.Spec.ForProvider.SecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SecurityGroupIdsRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this MySQLCluster.
func (mg *MySQLCluster) ResolveReferences(ctx context.Context, c client.Reader) error {
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

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Host); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Host[i3].SubnetID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.Host[i3].SubnetIDRef,
			Selector:     mg.Spec.ForProvider.Host[i3].SubnetIDSelector,
			To: reference.To{
				List:    &v1alpha11.SubnetList{},
				Managed: &v1alpha11.Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Host[i3].SubnetID")
		}
		mg.Spec.ForProvider.Host[i3].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Host[i3].SubnetIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NetworkIDRef,
		Selector:     mg.Spec.ForProvider.NetworkIDSelector,
		To: reference.To{
			List:    &v1alpha11.NetworkList{},
			Managed: &v1alpha11.Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NetworkID")
	}
	mg.Spec.ForProvider.NetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetworkIDRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SecurityGroupIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SecurityGroupIdsRefs,
		Selector:      mg.Spec.ForProvider.SecurityGroupIdsSelector,
		To: reference.To{
			List:    &v1alpha11.SecurityGroupList{},
			Managed: &v1alpha11.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroupIds")
	}
	mg.Spec.ForProvider.SecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SecurityGroupIdsRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this MySQLDatabase.
func (mg *MySQLDatabase) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterIDRef,
		Selector:     mg.Spec.ForProvider.ClusterIDSelector,
		To: reference.To{
			List:    &MySQLClusterList{},
			Managed: &MySQLCluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterID")
	}
	mg.Spec.ForProvider.ClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this MySQLUser.
func (mg *MySQLUser) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterIDRef,
		Selector:     mg.Spec.ForProvider.ClusterIDSelector,
		To: reference.To{
			List:    &MySQLClusterList{},
			Managed: &MySQLCluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterID")
	}
	mg.Spec.ForProvider.ClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this PostgresqlCluster.
func (mg *PostgresqlCluster) ResolveReferences(ctx context.Context, c client.Reader) error {
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

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Host); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Host[i3].SubnetID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.Host[i3].SubnetIDRef,
			Selector:     mg.Spec.ForProvider.Host[i3].SubnetIDSelector,
			To: reference.To{
				List:    &v1alpha11.SubnetList{},
				Managed: &v1alpha11.Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Host[i3].SubnetID")
		}
		mg.Spec.ForProvider.Host[i3].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Host[i3].SubnetIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NetworkIDRef,
		Selector:     mg.Spec.ForProvider.NetworkIDSelector,
		To: reference.To{
			List:    &v1alpha11.NetworkList{},
			Managed: &v1alpha11.Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NetworkID")
	}
	mg.Spec.ForProvider.NetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetworkIDRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SecurityGroupIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SecurityGroupIdsRefs,
		Selector:      mg.Spec.ForProvider.SecurityGroupIdsSelector,
		To: reference.To{
			List:    &v1alpha11.SecurityGroupList{},
			Managed: &v1alpha11.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroupIds")
	}
	mg.Spec.ForProvider.SecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SecurityGroupIdsRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this PostgresqlDatabase.
func (mg *PostgresqlDatabase) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterIDRef,
		Selector:     mg.Spec.ForProvider.ClusterIDSelector,
		To: reference.To{
			List:    &PostgresqlClusterList{},
			Managed: &PostgresqlCluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterID")
	}
	mg.Spec.ForProvider.ClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this PostgresqlUser.
func (mg *PostgresqlUser) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterIDRef,
		Selector:     mg.Spec.ForProvider.ClusterIDSelector,
		To: reference.To{
			List:    &PostgresqlClusterList{},
			Managed: &PostgresqlCluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterID")
	}
	mg.Spec.ForProvider.ClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RedisCluster.
func (mg *RedisCluster) ResolveReferences(ctx context.Context, c client.Reader) error {
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

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Host); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Host[i3].SubnetID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.Host[i3].SubnetIDRef,
			Selector:     mg.Spec.ForProvider.Host[i3].SubnetIDSelector,
			To: reference.To{
				List:    &v1alpha11.SubnetList{},
				Managed: &v1alpha11.Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Host[i3].SubnetID")
		}
		mg.Spec.ForProvider.Host[i3].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Host[i3].SubnetIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NetworkIDRef,
		Selector:     mg.Spec.ForProvider.NetworkIDSelector,
		To: reference.To{
			List:    &v1alpha11.NetworkList{},
			Managed: &v1alpha11.Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NetworkID")
	}
	mg.Spec.ForProvider.NetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetworkIDRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SecurityGroupIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SecurityGroupIdsRefs,
		Selector:      mg.Spec.ForProvider.SecurityGroupIdsSelector,
		To: reference.To{
			List:    &v1alpha11.SecurityGroupList{},
			Managed: &v1alpha11.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroupIds")
	}
	mg.Spec.ForProvider.SecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SecurityGroupIdsRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this SqlserverCluster.
func (mg *SqlserverCluster) ResolveReferences(ctx context.Context, c client.Reader) error {
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

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Host); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Host[i3].SubnetID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.Host[i3].SubnetIDRef,
			Selector:     mg.Spec.ForProvider.Host[i3].SubnetIDSelector,
			To: reference.To{
				List:    &v1alpha11.SubnetList{},
				Managed: &v1alpha11.Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Host[i3].SubnetID")
		}
		mg.Spec.ForProvider.Host[i3].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Host[i3].SubnetIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NetworkIDRef,
		Selector:     mg.Spec.ForProvider.NetworkIDSelector,
		To: reference.To{
			List:    &v1alpha11.NetworkList{},
			Managed: &v1alpha11.Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NetworkID")
	}
	mg.Spec.ForProvider.NetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetworkIDRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SecurityGroupIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SecurityGroupIdsRefs,
		Selector:      mg.Spec.ForProvider.SecurityGroupIdsSelector,
		To: reference.To{
			List:    &v1alpha11.SecurityGroupList{},
			Managed: &v1alpha11.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroupIds")
	}
	mg.Spec.ForProvider.SecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SecurityGroupIdsRefs = mrsp.ResolvedReferences

	return nil
}
