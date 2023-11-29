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

package iamworkforcepool

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const workforcePoolProviderIdRegexp = `^[a-z0-9-]{4,32}$`

func ValidateWorkforcePoolProviderId(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)

	if strings.HasPrefix(value, "gcp-") {
		errors = append(errors, fmt.Errorf(
			"%q (%q) can not start with \"gcp-\". "+
				"The prefix `gcp-` is reserved for use by Google, and may not be specified.", k, value))
	}

	if !regexp.MustCompile(workforcePoolProviderIdRegexp).MatchString(value) {
		errors = append(errors, fmt.Errorf(
			"%q (%q) must be 4-32 characters, and may contain the characters [a-z0-9-].", k, value))
	}

	return
}

const IAMWorkforcePoolWorkforcePoolProviderAssetType string = "iam.googleapis.com/WorkforcePoolProvider"

func ResourceConverterIAMWorkforcePoolWorkforcePoolProvider() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: IAMWorkforcePoolWorkforcePoolProviderAssetType,
		Convert:   GetIAMWorkforcePoolWorkforcePoolProviderCaiObject,
	}
}

func GetIAMWorkforcePoolWorkforcePoolProviderCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//iam.googleapis.com/locations/{{location}}/workforcePools/{{workforce_pool_id}}/providers/{{provider_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetIAMWorkforcePoolWorkforcePoolProviderApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: IAMWorkforcePoolWorkforcePoolProviderAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/iam/v1/rest",
				DiscoveryName:        "WorkforcePoolProvider",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetIAMWorkforcePoolWorkforcePoolProviderApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandIAMWorkforcePoolWorkforcePoolProviderDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandIAMWorkforcePoolWorkforcePoolProviderDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	disabledProp, err := expandIAMWorkforcePoolWorkforcePoolProviderDisabled(d.Get("disabled"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("disabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(disabledProp)) && (ok || !reflect.DeepEqual(v, disabledProp)) {
		obj["disabled"] = disabledProp
	}
	attributeMappingProp, err := expandIAMWorkforcePoolWorkforcePoolProviderAttributeMapping(d.Get("attribute_mapping"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("attribute_mapping"); !tpgresource.IsEmptyValue(reflect.ValueOf(attributeMappingProp)) && (ok || !reflect.DeepEqual(v, attributeMappingProp)) {
		obj["attributeMapping"] = attributeMappingProp
	}
	attributeConditionProp, err := expandIAMWorkforcePoolWorkforcePoolProviderAttributeCondition(d.Get("attribute_condition"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("attribute_condition"); !tpgresource.IsEmptyValue(reflect.ValueOf(attributeConditionProp)) && (ok || !reflect.DeepEqual(v, attributeConditionProp)) {
		obj["attributeCondition"] = attributeConditionProp
	}
	samlProp, err := expandIAMWorkforcePoolWorkforcePoolProviderSaml(d.Get("saml"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("saml"); !tpgresource.IsEmptyValue(reflect.ValueOf(samlProp)) && (ok || !reflect.DeepEqual(v, samlProp)) {
		obj["saml"] = samlProp
	}
	oidcProp, err := expandIAMWorkforcePoolWorkforcePoolProviderOidc(d.Get("oidc"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("oidc"); !tpgresource.IsEmptyValue(reflect.ValueOf(oidcProp)) && (ok || !reflect.DeepEqual(v, oidcProp)) {
		obj["oidc"] = oidcProp
	}

	return obj, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderDisabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderAttributeMapping(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderAttributeCondition(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderSaml(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedIdpMetadataXml, err := expandIAMWorkforcePoolWorkforcePoolProviderSamlIdpMetadataXml(original["idp_metadata_xml"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIdpMetadataXml); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["idpMetadataXml"] = transformedIdpMetadataXml
	}

	return transformed, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderSamlIdpMetadataXml(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderOidc(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedIssuerUri, err := expandIAMWorkforcePoolWorkforcePoolProviderOidcIssuerUri(original["issuer_uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIssuerUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["issuerUri"] = transformedIssuerUri
	}

	transformedClientId, err := expandIAMWorkforcePoolWorkforcePoolProviderOidcClientId(original["client_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedClientId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["clientId"] = transformedClientId
	}

	transformedClientSecret, err := expandIAMWorkforcePoolWorkforcePoolProviderOidcClientSecret(original["client_secret"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedClientSecret); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["clientSecret"] = transformedClientSecret
	}

	transformedWebSsoConfig, err := expandIAMWorkforcePoolWorkforcePoolProviderOidcWebSsoConfig(original["web_sso_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWebSsoConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["webSsoConfig"] = transformedWebSsoConfig
	}

	transformedJwksJson, err := expandIAMWorkforcePoolWorkforcePoolProviderOidcJwksJson(original["jwks_json"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedJwksJson); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["jwksJson"] = transformedJwksJson
	}

	return transformed, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderOidcIssuerUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderOidcClientId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderOidcClientSecret(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedValue, err := expandIAMWorkforcePoolWorkforcePoolProviderOidcClientSecretValue(original["value"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedValue); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["value"] = transformedValue
	}

	return transformed, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderOidcClientSecretValue(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPlainText, err := expandIAMWorkforcePoolWorkforcePoolProviderOidcClientSecretValuePlainText(original["plain_text"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPlainText); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["plainText"] = transformedPlainText
	}

	transformedThumbprint, err := expandIAMWorkforcePoolWorkforcePoolProviderOidcClientSecretValueThumbprint(original["thumbprint"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedThumbprint); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["thumbprint"] = transformedThumbprint
	}

	return transformed, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderOidcClientSecretValuePlainText(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderOidcClientSecretValueThumbprint(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderOidcWebSsoConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedResponseType, err := expandIAMWorkforcePoolWorkforcePoolProviderOidcWebSsoConfigResponseType(original["response_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResponseType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["responseType"] = transformedResponseType
	}

	transformedAssertionClaimsBehavior, err := expandIAMWorkforcePoolWorkforcePoolProviderOidcWebSsoConfigAssertionClaimsBehavior(original["assertion_claims_behavior"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAssertionClaimsBehavior); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["assertionClaimsBehavior"] = transformedAssertionClaimsBehavior
	}

	transformedAdditionalScopes, err := expandIAMWorkforcePoolWorkforcePoolProviderOidcWebSsoConfigAdditionalScopes(original["additional_scopes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAdditionalScopes); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["additionalScopes"] = transformedAdditionalScopes
	}

	return transformed, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderOidcWebSsoConfigResponseType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderOidcWebSsoConfigAssertionClaimsBehavior(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderOidcWebSsoConfigAdditionalScopes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderOidcJwksJson(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
