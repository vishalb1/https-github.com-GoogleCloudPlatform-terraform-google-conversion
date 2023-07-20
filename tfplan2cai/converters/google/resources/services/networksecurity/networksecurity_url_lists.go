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

package networksecurity

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const NetworkSecurityUrlListsAssetType string = "networksecurity.googleapis.com/UrlLists"

func ResourceConverterNetworkSecurityUrlLists() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType: NetworkSecurityUrlListsAssetType,
		Convert:   GetNetworkSecurityUrlListsCaiObject,
	}
}

func GetNetworkSecurityUrlListsCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	name, err := tpgresource.AssetName(d, config, "//networksecurity.googleapis.com/projects/{{project}}/locations/{{location}}/urlLists/{{name}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}
	if obj, err := GetNetworkSecurityUrlListsApiObject(d, config); err == nil {
		return []tpgresource.Asset{{
			Name: name,
			Type: NetworkSecurityUrlListsAssetType,
			Resource: &tpgresource.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/networksecurity/v1beta1/rest",
				DiscoveryName:        "UrlLists",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []tpgresource.Asset{}, err
	}
}

func GetNetworkSecurityUrlListsApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkSecurityUrlListsDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	valuesProp, err := expandNetworkSecurityUrlListsValues(d.Get("values"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("values"); !tpgresource.IsEmptyValue(reflect.ValueOf(valuesProp)) && (ok || !reflect.DeepEqual(v, valuesProp)) {
		obj["values"] = valuesProp
	}

	return obj, nil
}

func expandNetworkSecurityUrlListsDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityUrlListsValues(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
