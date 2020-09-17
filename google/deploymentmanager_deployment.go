// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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

package google

import (
	"context"
	"log"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func customDiffDeploymentManagerDeployment(_ context.Context, d *schema.ResourceDiff, meta interface{}) error {
	if preview := d.Get("preview").(bool); preview {
		log.Printf("[WARN] Deployment preview set to true - Terraform will treat Deployment as recreate-only")

		if d.HasChange("preview") {
			d.ForceNew("preview")
		}

		if d.HasChange("target") {
			d.ForceNew("target")
		}

		if d.HasChange("labels") {
			d.ForceNew("labels")
		}
	}
	return nil
}

func GetDeploymentManagerDeploymentCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	name, err := assetName(d, config, "//deploymentmanager.googleapis.com/projects/{{project}}/global/deployments/{{name}}")
	if err != nil {
		return Asset{}, err
	}
	if obj, err := GetDeploymentManagerDeploymentApiObject(d, config); err == nil {
		return Asset{
			Name: name,
			Type: "deploymentmanager.googleapis.com/Deployment",
			Resource: &AssetResource{
				Version:              "v2",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/deploymentmanager/v2/rest",
				DiscoveryName:        "Deployment",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetDeploymentManagerDeploymentApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandDeploymentManagerDeploymentName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandDeploymentManagerDeploymentDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandDeploymentManagerDeploymentLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); ok || !reflect.DeepEqual(v, labelsProp) {
		obj["labels"] = labelsProp
	}
	targetProp, err := expandDeploymentManagerDeploymentTarget(d.Get("target"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("target"); !isEmptyValue(reflect.ValueOf(targetProp)) && (ok || !reflect.DeepEqual(v, targetProp)) {
		obj["target"] = targetProp
	}

	return obj, nil
}

func expandDeploymentManagerDeploymentName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDeploymentManagerDeploymentDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDeploymentManagerDeploymentLabels(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedKey, err := expandDeploymentManagerDeploymentLabelsKey(original["key"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedKey); val.IsValid() && !isEmptyValue(val) {
			transformed["key"] = transformedKey
		}

		transformedValue, err := expandDeploymentManagerDeploymentLabelsValue(original["value"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedValue); val.IsValid() && !isEmptyValue(val) {
			transformed["value"] = transformedValue
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDeploymentManagerDeploymentLabelsKey(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDeploymentManagerDeploymentLabelsValue(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDeploymentManagerDeploymentTarget(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedConfig, err := expandDeploymentManagerDeploymentTargetConfig(original["config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedConfig); val.IsValid() && !isEmptyValue(val) {
		transformed["config"] = transformedConfig
	}

	transformedImports, err := expandDeploymentManagerDeploymentTargetImports(original["imports"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedImports); val.IsValid() && !isEmptyValue(val) {
		transformed["imports"] = transformedImports
	}

	return transformed, nil
}

func expandDeploymentManagerDeploymentTargetConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedContent, err := expandDeploymentManagerDeploymentTargetConfigContent(original["content"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedContent); val.IsValid() && !isEmptyValue(val) {
		transformed["content"] = transformedContent
	}

	return transformed, nil
}

func expandDeploymentManagerDeploymentTargetConfigContent(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDeploymentManagerDeploymentTargetImports(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedContent, err := expandDeploymentManagerDeploymentTargetImportsContent(original["content"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedContent); val.IsValid() && !isEmptyValue(val) {
			transformed["content"] = transformedContent
		}

		transformedName, err := expandDeploymentManagerDeploymentTargetImportsName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
			transformed["name"] = transformedName
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDeploymentManagerDeploymentTargetImportsContent(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDeploymentManagerDeploymentTargetImportsName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
