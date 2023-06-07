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

package resourcemanager

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const ResourceManagerLienAssetType string = "cloudresourcemanager.googleapis.com/Lien"

func ResourceConverterResourceManagerLien() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType: ResourceManagerLienAssetType,
		Convert:   GetResourceManagerLienCaiObject,
	}
}

func GetResourceManagerLienCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	name, err := tpgresource.AssetName(d, config, "//cloudresourcemanager.googleapis.com/liens?parent={{parent}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}
	if obj, err := GetResourceManagerLienApiObject(d, config); err == nil {
		return []tpgresource.Asset{{
			Name: name,
			Type: ResourceManagerLienAssetType,
			Resource: &tpgresource.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/cloudresourcemanager/v1/rest",
				DiscoveryName:        "Lien",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []tpgresource.Asset{}, err
	}
}

func GetResourceManagerLienApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	reasonProp, err := expandResourceManagerLienReason(d.Get("reason"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("reason"); !tpgresource.IsEmptyValue(reflect.ValueOf(reasonProp)) && (ok || !reflect.DeepEqual(v, reasonProp)) {
		obj["reason"] = reasonProp
	}
	originProp, err := expandResourceManagerLienOrigin(d.Get("origin"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("origin"); !tpgresource.IsEmptyValue(reflect.ValueOf(originProp)) && (ok || !reflect.DeepEqual(v, originProp)) {
		obj["origin"] = originProp
	}
	parentProp, err := expandResourceManagerLienParent(d.Get("parent"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("parent"); !tpgresource.IsEmptyValue(reflect.ValueOf(parentProp)) && (ok || !reflect.DeepEqual(v, parentProp)) {
		obj["parent"] = parentProp
	}
	restrictionsProp, err := expandResourceManagerLienRestrictions(d.Get("restrictions"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("restrictions"); !tpgresource.IsEmptyValue(reflect.ValueOf(restrictionsProp)) && (ok || !reflect.DeepEqual(v, restrictionsProp)) {
		obj["restrictions"] = restrictionsProp
	}

	return obj, nil
}

func expandResourceManagerLienReason(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandResourceManagerLienOrigin(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandResourceManagerLienParent(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandResourceManagerLienRestrictions(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}