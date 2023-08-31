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

package gkehub2

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const GKEHub2ScopeRBACRoleBindingAssetType string = "gkehub.googleapis.com/ScopeRBACRoleBinding"

func ResourceConverterGKEHub2ScopeRBACRoleBinding() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: GKEHub2ScopeRBACRoleBindingAssetType,
		Convert:   GetGKEHub2ScopeRBACRoleBindingCaiObject,
	}
}

func GetGKEHub2ScopeRBACRoleBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//gkehub.googleapis.com/projects/{{project}}/locations/global/scopes/{{scope_id}}/rbacrolebindings/{{scope_rbac_role_binding_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetGKEHub2ScopeRBACRoleBindingApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: GKEHub2ScopeRBACRoleBindingAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/gkehub/v1beta/rest",
				DiscoveryName:        "ScopeRBACRoleBinding",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetGKEHub2ScopeRBACRoleBindingApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	userProp, err := expandGKEHub2ScopeRBACRoleBindingUser(d.Get("user"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("user"); !tpgresource.IsEmptyValue(reflect.ValueOf(userProp)) && (ok || !reflect.DeepEqual(v, userProp)) {
		obj["user"] = userProp
	}
	groupProp, err := expandGKEHub2ScopeRBACRoleBindingGroup(d.Get("group"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("group"); !tpgresource.IsEmptyValue(reflect.ValueOf(groupProp)) && (ok || !reflect.DeepEqual(v, groupProp)) {
		obj["group"] = groupProp
	}
	roleProp, err := expandGKEHub2ScopeRBACRoleBindingRole(d.Get("role"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("role"); !tpgresource.IsEmptyValue(reflect.ValueOf(roleProp)) && (ok || !reflect.DeepEqual(v, roleProp)) {
		obj["role"] = roleProp
	}

	return obj, nil
}

func expandGKEHub2ScopeRBACRoleBindingUser(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEHub2ScopeRBACRoleBindingGroup(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEHub2ScopeRBACRoleBindingRole(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPredefinedRole, err := expandGKEHub2ScopeRBACRoleBindingRolePredefinedRole(original["predefined_role"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPredefinedRole); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["predefinedRole"] = transformedPredefinedRole
	}

	return transformed, nil
}

func expandGKEHub2ScopeRBACRoleBindingRolePredefinedRole(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}