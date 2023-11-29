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

package storage

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const StorageBucketAccessControlAssetType string = "storage.googleapis.com/BucketAccessControl"

func ResourceConverterStorageBucketAccessControl() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: StorageBucketAccessControlAssetType,
		Convert:   GetStorageBucketAccessControlCaiObject,
	}
}

func GetStorageBucketAccessControlCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//storage.googleapis.com/b/{{bucket}}/acl/{{entity}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetStorageBucketAccessControlApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: StorageBucketAccessControlAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/storage/v1/rest",
				DiscoveryName:        "BucketAccessControl",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetStorageBucketAccessControlApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	bucketProp, err := expandStorageBucketAccessControlBucket(d.Get("bucket"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("bucket"); !tpgresource.IsEmptyValue(reflect.ValueOf(bucketProp)) && (ok || !reflect.DeepEqual(v, bucketProp)) {
		obj["bucket"] = bucketProp
	}
	entityProp, err := expandStorageBucketAccessControlEntity(d.Get("entity"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("entity"); !tpgresource.IsEmptyValue(reflect.ValueOf(entityProp)) && (ok || !reflect.DeepEqual(v, entityProp)) {
		obj["entity"] = entityProp
	}
	roleProp, err := expandStorageBucketAccessControlRole(d.Get("role"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("role"); !tpgresource.IsEmptyValue(reflect.ValueOf(roleProp)) && (ok || !reflect.DeepEqual(v, roleProp)) {
		obj["role"] = roleProp
	}

	return obj, nil
}

func expandStorageBucketAccessControlBucket(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandStorageBucketAccessControlEntity(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandStorageBucketAccessControlRole(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
