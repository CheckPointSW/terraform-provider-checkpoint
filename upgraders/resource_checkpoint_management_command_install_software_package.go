package upgraders

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceManagementCommandInstallSoftwarePackageV0 is the V0 schema where cluster_installation_settings was TypeMap.
func ResourceManagementCommandInstallSoftwarePackageV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the software package.",
			},
			"targets": {
				Type:        schema.TypeSet,
				Required:    true,
				ForceNew:    true,
				Description: "On what targets to execute this command. Targets may be identified by their name, or object unique identifier.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"cluster_installation_settings": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Installation settings for cluster.",
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cluster_delay": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The delay between end of installation on one cluster members and start of installation on the next cluster member.",
							Default:     0,
						},
						"cluster_strategy": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The cluster installation strategy.",
							Default:     "legacy",
						},
					},
				},
			},
			"concurrency_limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "The number of targets, on which the same package is installed at the same time.",
			},
			"task_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Command asynchronous task unique identifier.",
			},
		},
	}
}

// ResourceManagementCommandInstallSoftwarePackageStateUpgradeV0 converts cluster_installation_settings from TypeMap to TypeList.
func ResourceManagementCommandInstallSoftwarePackageStateUpgradeV0(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
	return UpgradeMapsToLists(rawState, "cluster_installation_settings"), nil
}
