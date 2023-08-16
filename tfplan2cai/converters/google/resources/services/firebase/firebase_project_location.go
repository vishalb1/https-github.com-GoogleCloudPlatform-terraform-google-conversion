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

package firebase

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const FirebaseProjectLocationAssetType string = "firebase.googleapis.com/ProjectLocation"

func ResourceConverterFirebaseProjectLocation() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: FirebaseProjectLocationAssetType,
		Convert:   GetFirebaseProjectLocationCaiObject,
	}
}

func GetFirebaseProjectLocationCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//firebase.googleapis.com/projects/{{project}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetFirebaseProjectLocationApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: FirebaseProjectLocationAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/firebase/v1beta1/rest",
				DiscoveryName:        "ProjectLocation",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetFirebaseProjectLocationApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	locationIdProp, err := expandFirebaseProjectLocationLocationId(d.Get("location_id"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("location_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(locationIdProp)) && (ok || !reflect.DeepEqual(v, locationIdProp)) {
		obj["locationId"] = locationIdProp
	}

	return obj, nil
}

func expandFirebaseProjectLocationLocationId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
