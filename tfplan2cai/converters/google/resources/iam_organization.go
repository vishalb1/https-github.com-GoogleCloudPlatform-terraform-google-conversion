package google

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/resourcemanager"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgiamresource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

var IamOrganizationSchema = resourcemanager.IamOrganizationSchema

func NewOrganizationIamUpdater(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgiamresource.ResourceIamUpdater, error) {
	return resourcemanager.NewOrganizationIamUpdater(d, config)
}

func OrgIdParseFunc(d *schema.ResourceData, _ *transport_tpg.Config) error {
	return resourcemanager.OrgIdParseFunc(d, nil)
}
