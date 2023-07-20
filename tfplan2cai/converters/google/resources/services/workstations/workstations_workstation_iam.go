// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package workstations

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgiamresource"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const WorkstationsWorkstationIAMAssetType string = "workstations.googleapis.com/Workstation"

func ResourceConverterWorkstationsWorkstationIamPolicy() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         WorkstationsWorkstationIAMAssetType,
		Convert:           GetWorkstationsWorkstationIamPolicyCaiObject,
		MergeCreateUpdate: MergeWorkstationsWorkstationIamPolicy,
	}
}

func ResourceConverterWorkstationsWorkstationIamBinding() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         WorkstationsWorkstationIAMAssetType,
		Convert:           GetWorkstationsWorkstationIamBindingCaiObject,
		FetchFullResource: FetchWorkstationsWorkstationIamPolicy,
		MergeCreateUpdate: MergeWorkstationsWorkstationIamBinding,
		MergeDelete:       MergeWorkstationsWorkstationIamBindingDelete,
	}
}

func ResourceConverterWorkstationsWorkstationIamMember() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         WorkstationsWorkstationIAMAssetType,
		Convert:           GetWorkstationsWorkstationIamMemberCaiObject,
		FetchFullResource: FetchWorkstationsWorkstationIamPolicy,
		MergeCreateUpdate: MergeWorkstationsWorkstationIamMember,
		MergeDelete:       MergeWorkstationsWorkstationIamMemberDelete,
	}
}

func GetWorkstationsWorkstationIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newWorkstationsWorkstationIamAsset(d, config, tpgiamresource.ExpandIamPolicyBindings)
}

func GetWorkstationsWorkstationIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newWorkstationsWorkstationIamAsset(d, config, tpgiamresource.ExpandIamRoleBindings)
}

func GetWorkstationsWorkstationIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newWorkstationsWorkstationIamAsset(d, config, tpgiamresource.ExpandIamMemberBindings)
}

func MergeWorkstationsWorkstationIamPolicy(existing, incoming tpgresource.Asset) tpgresource.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeWorkstationsWorkstationIamBinding(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeIamAssets(existing, incoming, tpgiamresource.MergeAuthoritativeBindings)
}

func MergeWorkstationsWorkstationIamBindingDelete(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeDeleteIamAssets(existing, incoming, tpgiamresource.MergeDeleteAuthoritativeBindings)
}

func MergeWorkstationsWorkstationIamMember(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeIamAssets(existing, incoming, tpgiamresource.MergeAdditiveBindings)
}

func MergeWorkstationsWorkstationIamMemberDelete(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeDeleteIamAssets(existing, incoming, tpgiamresource.MergeDeleteAdditiveBindings)
}

func newWorkstationsWorkstationIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]tpgresource.IAMBinding, error),
) ([]tpgresource.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []tpgresource.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := tpgresource.AssetName(d, config, "//workstations.googleapis.com/projects/{{project}}/locations/{{location}}/workstationClusters/{{workstation_cluster_id}}/workstationConfigs/{{workstation_config_id}}/workstations/{{workstation_id}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}

	return []tpgresource.Asset{{
		Name: name,
		Type: WorkstationsWorkstationIAMAssetType,
		IAMPolicy: &tpgresource.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchWorkstationsWorkstationIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgresource.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("location"); !ok {
		return tpgresource.Asset{}, tpgresource.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("workstation_cluster_id"); !ok {
		return tpgresource.Asset{}, tpgresource.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("workstation_config_id"); !ok {
		return tpgresource.Asset{}, tpgresource.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("workstation_id"); !ok {
		return tpgresource.Asset{}, tpgresource.ErrEmptyIdentityField
	}

	return tpgiamresource.FetchIamPolicy(
		WorkstationsWorkstationIamUpdaterProducer,
		d,
		config,
		"//workstations.googleapis.com/projects/{{project}}/locations/{{location}}/workstationClusters/{{workstation_cluster_id}}/workstationConfigs/{{workstation_config_id}}/workstations/{{workstation_id}}",
		WorkstationsWorkstationIAMAssetType,
	)
}
