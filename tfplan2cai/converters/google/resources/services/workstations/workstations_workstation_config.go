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
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const WorkstationsWorkstationConfigAssetType string = "workstations.googleapis.com/WorkstationConfig"

func ResourceConverterWorkstationsWorkstationConfig() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType: WorkstationsWorkstationConfigAssetType,
		Convert:   GetWorkstationsWorkstationConfigCaiObject,
	}
}

func GetWorkstationsWorkstationConfigCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	name, err := tpgresource.AssetName(d, config, "//workstations.googleapis.com/projects/{{project}}/locations/{{location}}/workstationClusters/{{workstation_cluster_id}}/workstationConfigs/{{workstation_config_id}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}
	if obj, err := GetWorkstationsWorkstationConfigApiObject(d, config); err == nil {
		return []tpgresource.Asset{{
			Name: name,
			Type: WorkstationsWorkstationConfigAssetType,
			Resource: &tpgresource.AssetResource{
				Version:              "v1beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/workstations/v1beta/rest",
				DiscoveryName:        "WorkstationConfig",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []tpgresource.Asset{}, err
	}
}

func GetWorkstationsWorkstationConfigApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandWorkstationsWorkstationConfigDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	labelsProp, err := expandWorkstationsWorkstationConfigLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	annotationsProp, err := expandWorkstationsWorkstationConfigAnnotations(d.Get("annotations"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("annotations"); !tpgresource.IsEmptyValue(reflect.ValueOf(annotationsProp)) && (ok || !reflect.DeepEqual(v, annotationsProp)) {
		obj["annotations"] = annotationsProp
	}
	etagProp, err := expandWorkstationsWorkstationConfigEtag(d.Get("etag"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("etag"); !tpgresource.IsEmptyValue(reflect.ValueOf(etagProp)) && (ok || !reflect.DeepEqual(v, etagProp)) {
		obj["etag"] = etagProp
	}
	idleTimeoutProp, err := expandWorkstationsWorkstationConfigIdleTimeout(d.Get("idle_timeout"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("idle_timeout"); !tpgresource.IsEmptyValue(reflect.ValueOf(idleTimeoutProp)) && (ok || !reflect.DeepEqual(v, idleTimeoutProp)) {
		obj["idleTimeout"] = idleTimeoutProp
	}
	runningTimeoutProp, err := expandWorkstationsWorkstationConfigRunningTimeout(d.Get("running_timeout"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("running_timeout"); !tpgresource.IsEmptyValue(reflect.ValueOf(runningTimeoutProp)) && (ok || !reflect.DeepEqual(v, runningTimeoutProp)) {
		obj["runningTimeout"] = runningTimeoutProp
	}
	hostProp, err := expandWorkstationsWorkstationConfigHost(d.Get("host"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("host"); !tpgresource.IsEmptyValue(reflect.ValueOf(hostProp)) && (ok || !reflect.DeepEqual(v, hostProp)) {
		obj["host"] = hostProp
	}
	persistentDirectoriesProp, err := expandWorkstationsWorkstationConfigPersistentDirectories(d.Get("persistent_directories"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("persistent_directories"); !tpgresource.IsEmptyValue(reflect.ValueOf(persistentDirectoriesProp)) && (ok || !reflect.DeepEqual(v, persistentDirectoriesProp)) {
		obj["persistentDirectories"] = persistentDirectoriesProp
	}
	containerProp, err := expandWorkstationsWorkstationConfigContainer(d.Get("container"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("container"); !tpgresource.IsEmptyValue(reflect.ValueOf(containerProp)) && (ok || !reflect.DeepEqual(v, containerProp)) {
		obj["container"] = containerProp
	}
	encryptionKeyProp, err := expandWorkstationsWorkstationConfigEncryptionKey(d.Get("encryption_key"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("encryption_key"); !tpgresource.IsEmptyValue(reflect.ValueOf(encryptionKeyProp)) && (ok || !reflect.DeepEqual(v, encryptionKeyProp)) {
		obj["encryptionKey"] = encryptionKeyProp
	}

	return obj, nil
}

func expandWorkstationsWorkstationConfigDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkstationsWorkstationConfigLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandWorkstationsWorkstationConfigAnnotations(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandWorkstationsWorkstationConfigEtag(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkstationsWorkstationConfigIdleTimeout(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkstationsWorkstationConfigRunningTimeout(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkstationsWorkstationConfigHost(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedGceInstance, err := expandWorkstationsWorkstationConfigHostGceInstance(original["gce_instance"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGceInstance); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["gceInstance"] = transformedGceInstance
	}

	return transformed, nil
}

func expandWorkstationsWorkstationConfigHostGceInstance(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMachineType, err := expandWorkstationsWorkstationConfigHostGceInstanceMachineType(original["machine_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMachineType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["machineType"] = transformedMachineType
	}

	transformedServiceAccount, err := expandWorkstationsWorkstationConfigHostGceInstanceServiceAccount(original["service_account"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedServiceAccount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["serviceAccount"] = transformedServiceAccount
	}

	transformedPoolSize, err := expandWorkstationsWorkstationConfigHostGceInstancePoolSize(original["pool_size"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPoolSize); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["poolSize"] = transformedPoolSize
	}

	transformedBootDiskSizeGb, err := expandWorkstationsWorkstationConfigHostGceInstanceBootDiskSizeGb(original["boot_disk_size_gb"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBootDiskSizeGb); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["bootDiskSizeGb"] = transformedBootDiskSizeGb
	}

	transformedTags, err := expandWorkstationsWorkstationConfigHostGceInstanceTags(original["tags"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTags); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["tags"] = transformedTags
	}

	transformedDisablePublicIpAddresses, err := expandWorkstationsWorkstationConfigHostGceInstanceDisablePublicIpAddresses(original["disable_public_ip_addresses"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDisablePublicIpAddresses); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["disablePublicIpAddresses"] = transformedDisablePublicIpAddresses
	}

	transformedShieldedInstanceConfig, err := expandWorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfig(original["shielded_instance_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedShieldedInstanceConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["shieldedInstanceConfig"] = transformedShieldedInstanceConfig
	}

	transformedConfidentialInstanceConfig, err := expandWorkstationsWorkstationConfigHostGceInstanceConfidentialInstanceConfig(original["confidential_instance_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedConfidentialInstanceConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["confidentialInstanceConfig"] = transformedConfidentialInstanceConfig
	}

	transformedAccelerators, err := expandWorkstationsWorkstationConfigHostGceInstanceAccelerators(original["accelerators"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAccelerators); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["accelerators"] = transformedAccelerators
	}

	return transformed, nil
}

func expandWorkstationsWorkstationConfigHostGceInstanceMachineType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkstationsWorkstationConfigHostGceInstanceServiceAccount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkstationsWorkstationConfigHostGceInstancePoolSize(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkstationsWorkstationConfigHostGceInstanceBootDiskSizeGb(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkstationsWorkstationConfigHostGceInstanceTags(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkstationsWorkstationConfigHostGceInstanceDisablePublicIpAddresses(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnableSecureBoot, err := expandWorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigEnableSecureBoot(original["enable_secure_boot"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["enableSecureBoot"] = transformedEnableSecureBoot
	}

	transformedEnableVtpm, err := expandWorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigEnableVtpm(original["enable_vtpm"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["enableVtpm"] = transformedEnableVtpm
	}

	transformedEnableIntegrityMonitoring, err := expandWorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigEnableIntegrityMonitoring(original["enable_integrity_monitoring"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["enableIntegrityMonitoring"] = transformedEnableIntegrityMonitoring
	}

	return transformed, nil
}

func expandWorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigEnableSecureBoot(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigEnableVtpm(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigEnableIntegrityMonitoring(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkstationsWorkstationConfigHostGceInstanceConfidentialInstanceConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnableConfidentialCompute, err := expandWorkstationsWorkstationConfigHostGceInstanceConfidentialInstanceConfigEnableConfidentialCompute(original["enable_confidential_compute"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["enableConfidentialCompute"] = transformedEnableConfidentialCompute
	}

	return transformed, nil
}

func expandWorkstationsWorkstationConfigHostGceInstanceConfidentialInstanceConfigEnableConfidentialCompute(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkstationsWorkstationConfigHostGceInstanceAccelerators(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedType, err := expandWorkstationsWorkstationConfigHostGceInstanceAcceleratorsType(original["type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["type"] = transformedType
		}

		transformedCount, err := expandWorkstationsWorkstationConfigHostGceInstanceAcceleratorsCount(original["count"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["count"] = transformedCount
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandWorkstationsWorkstationConfigHostGceInstanceAcceleratorsType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkstationsWorkstationConfigHostGceInstanceAcceleratorsCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkstationsWorkstationConfigPersistentDirectories(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedMountPath, err := expandWorkstationsWorkstationConfigPersistentDirectoriesMountPath(original["mount_path"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedMountPath); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["mountPath"] = transformedMountPath
		}

		transformedGcePd, err := expandWorkstationsWorkstationConfigPersistentDirectoriesGcePd(original["gce_pd"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedGcePd); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["gcePd"] = transformedGcePd
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandWorkstationsWorkstationConfigPersistentDirectoriesMountPath(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkstationsWorkstationConfigPersistentDirectoriesGcePd(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedFsType, err := expandWorkstationsWorkstationConfigPersistentDirectoriesGcePdFsType(original["fs_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFsType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["fsType"] = transformedFsType
	}

	transformedDiskType, err := expandWorkstationsWorkstationConfigPersistentDirectoriesGcePdDiskType(original["disk_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDiskType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["diskType"] = transformedDiskType
	}

	transformedSizeGb, err := expandWorkstationsWorkstationConfigPersistentDirectoriesGcePdSizeGb(original["size_gb"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSizeGb); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["sizeGb"] = transformedSizeGb
	}

	transformedReclaimPolicy, err := expandWorkstationsWorkstationConfigPersistentDirectoriesGcePdReclaimPolicy(original["reclaim_policy"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedReclaimPolicy); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["reclaimPolicy"] = transformedReclaimPolicy
	}

	transformedSourceSnapshot, err := expandWorkstationsWorkstationConfigPersistentDirectoriesGcePdSourceSnapshot(original["source_snapshot"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSourceSnapshot); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["sourceSnapshot"] = transformedSourceSnapshot
	}

	return transformed, nil
}

func expandWorkstationsWorkstationConfigPersistentDirectoriesGcePdFsType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkstationsWorkstationConfigPersistentDirectoriesGcePdDiskType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkstationsWorkstationConfigPersistentDirectoriesGcePdSizeGb(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkstationsWorkstationConfigPersistentDirectoriesGcePdReclaimPolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkstationsWorkstationConfigPersistentDirectoriesGcePdSourceSnapshot(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkstationsWorkstationConfigContainer(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedImage, err := expandWorkstationsWorkstationConfigContainerImage(original["image"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedImage); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["image"] = transformedImage
	}

	transformedCommand, err := expandWorkstationsWorkstationConfigContainerCommand(original["command"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCommand); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["command"] = transformedCommand
	}

	transformedArgs, err := expandWorkstationsWorkstationConfigContainerArgs(original["args"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedArgs); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["args"] = transformedArgs
	}

	transformedWorkingDir, err := expandWorkstationsWorkstationConfigContainerWorkingDir(original["working_dir"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["workingDir"] = transformedWorkingDir
	}

	transformedEnv, err := expandWorkstationsWorkstationConfigContainerEnv(original["env"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnv); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["env"] = transformedEnv
	}

	transformedRunAsUser, err := expandWorkstationsWorkstationConfigContainerRunAsUser(original["run_as_user"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRunAsUser); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["runAsUser"] = transformedRunAsUser
	}

	return transformed, nil
}

func expandWorkstationsWorkstationConfigContainerImage(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkstationsWorkstationConfigContainerCommand(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkstationsWorkstationConfigContainerArgs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkstationsWorkstationConfigContainerWorkingDir(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkstationsWorkstationConfigContainerEnv(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandWorkstationsWorkstationConfigContainerRunAsUser(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkstationsWorkstationConfigEncryptionKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedKmsKey, err := expandWorkstationsWorkstationConfigEncryptionKeyKmsKey(original["kms_key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKey); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKey"] = transformedKmsKey
	}

	transformedKmsKeyServiceAccount, err := expandWorkstationsWorkstationConfigEncryptionKeyKmsKeyServiceAccount(original["kms_key_service_account"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeyServiceAccount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyServiceAccount"] = transformedKmsKeyServiceAccount
	}

	return transformed, nil
}

func expandWorkstationsWorkstationConfigEncryptionKeyKmsKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkstationsWorkstationConfigEncryptionKeyKmsKeyServiceAccount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}