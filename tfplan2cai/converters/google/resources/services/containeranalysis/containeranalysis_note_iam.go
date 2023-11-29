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

package containeranalysis

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const ContainerAnalysisNoteIAMAssetType string = "containeranalysis.googleapis.com/Note"

func ResourceConverterContainerAnalysisNoteIamPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         ContainerAnalysisNoteIAMAssetType,
		Convert:           GetContainerAnalysisNoteIamPolicyCaiObject,
		MergeCreateUpdate: MergeContainerAnalysisNoteIamPolicy,
	}
}

func ResourceConverterContainerAnalysisNoteIamBinding() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         ContainerAnalysisNoteIAMAssetType,
		Convert:           GetContainerAnalysisNoteIamBindingCaiObject,
		FetchFullResource: FetchContainerAnalysisNoteIamPolicy,
		MergeCreateUpdate: MergeContainerAnalysisNoteIamBinding,
		MergeDelete:       MergeContainerAnalysisNoteIamBindingDelete,
	}
}

func ResourceConverterContainerAnalysisNoteIamMember() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         ContainerAnalysisNoteIAMAssetType,
		Convert:           GetContainerAnalysisNoteIamMemberCaiObject,
		FetchFullResource: FetchContainerAnalysisNoteIamPolicy,
		MergeCreateUpdate: MergeContainerAnalysisNoteIamMember,
		MergeDelete:       MergeContainerAnalysisNoteIamMemberDelete,
	}
}

func GetContainerAnalysisNoteIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newContainerAnalysisNoteIamAsset(d, config, cai.ExpandIamPolicyBindings)
}

func GetContainerAnalysisNoteIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newContainerAnalysisNoteIamAsset(d, config, cai.ExpandIamRoleBindings)
}

func GetContainerAnalysisNoteIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newContainerAnalysisNoteIamAsset(d, config, cai.ExpandIamMemberBindings)
}

func MergeContainerAnalysisNoteIamPolicy(existing, incoming cai.Asset) cai.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeContainerAnalysisNoteIamBinding(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAuthoritativeBindings)
}

func MergeContainerAnalysisNoteIamBindingDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAuthoritativeBindings)
}

func MergeContainerAnalysisNoteIamMember(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAdditiveBindings)
}

func MergeContainerAnalysisNoteIamMemberDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAdditiveBindings)
}

func newContainerAnalysisNoteIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]cai.IAMBinding, error),
) ([]cai.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []cai.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := cai.AssetName(d, config, "//containeranalysis.googleapis.com/projects/{{project}}/notes/{{note}}")
	if err != nil {
		return []cai.Asset{}, err
	}

	return []cai.Asset{{
		Name: name,
		Type: ContainerAnalysisNoteIAMAssetType,
		IAMPolicy: &cai.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchContainerAnalysisNoteIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (cai.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("note"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}

	return cai.FetchIamPolicy(
		ContainerAnalysisNoteIamUpdaterProducer,
		d,
		config,
		"//containeranalysis.googleapis.com/projects/{{project}}/notes/{{note}}",
		ContainerAnalysisNoteIAMAssetType,
	)
}
