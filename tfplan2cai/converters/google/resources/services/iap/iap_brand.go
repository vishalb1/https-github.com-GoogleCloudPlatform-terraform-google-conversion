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

package iap

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const IapBrandAssetType string = "iap.googleapis.com/Brand"

func ResourceConverterIapBrand() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType: IapBrandAssetType,
		Convert:   GetIapBrandCaiObject,
	}
}

func GetIapBrandCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	name, err := tpgresource.AssetName(d, config, "//iap.googleapis.com/{{name}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}
	if obj, err := GetIapBrandApiObject(d, config); err == nil {
		return []tpgresource.Asset{{
			Name: name,
			Type: IapBrandAssetType,
			Resource: &tpgresource.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/iap/v1/rest",
				DiscoveryName:        "Brand",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []tpgresource.Asset{}, err
	}
}

func GetIapBrandApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	supportEmailProp, err := expandIapBrandSupportEmail(d.Get("support_email"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("support_email"); !tpgresource.IsEmptyValue(reflect.ValueOf(supportEmailProp)) && (ok || !reflect.DeepEqual(v, supportEmailProp)) {
		obj["supportEmail"] = supportEmailProp
	}
	applicationTitleProp, err := expandIapBrandApplicationTitle(d.Get("application_title"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("application_title"); !tpgresource.IsEmptyValue(reflect.ValueOf(applicationTitleProp)) && (ok || !reflect.DeepEqual(v, applicationTitleProp)) {
		obj["applicationTitle"] = applicationTitleProp
	}

	return obj, nil
}

func expandIapBrandSupportEmail(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIapBrandApplicationTitle(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}