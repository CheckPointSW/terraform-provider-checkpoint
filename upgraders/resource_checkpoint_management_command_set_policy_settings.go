package upgraders

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceManagementCommandSetPolicySettingsV0 is the V0 schema where security_access_defaults was TypeMap.
func ResourceManagementCommandSetPolicySettingsV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"last_in_cell": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Added object after removing the last object in cell.",
			},
			"none_object_behavior": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "'None' object behavior. Rules with object 'None' will never be matched.",
			},
			"security_access_defaults": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Access Policy default values.",
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"destination": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Destination default value for new rule creation. Any or None.",
						},
						"service": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Service and Applications default value for new rule creation. Any or None.",
						},
						"source": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Source default value for new rule creation. Any or None.",
						},
					},
				},
			},
		},
	}
}

// ResourceManagementCommandSetPolicySettingsStateUpgradeV0 converts security_access_defaults from TypeMap to TypeList.
func ResourceManagementCommandSetPolicySettingsStateUpgradeV0(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
	return UpgradeMapsToLists(rawState, "security_access_defaults"), nil
}
