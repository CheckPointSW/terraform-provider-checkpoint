package upgraders

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceManagementCommandConnectCloudServicesV0 is the V0 schema where connected_at was TypeMap.
func ResourceManagementCommandConnectCloudServicesV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"auth_token": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Copy the authentication token from the Smart-1 cloud service hosted in the Infinity Portal.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status of the connection to the Infinity Portal.",
			},
			"connected_at": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "The time of the connection between the Management Server and the Infinity Portal.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"iso_8601": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Date and time represented in international ISO 8601 format.",
						},
						"posix": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.",
						},
					},
				},
			},
			"management_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Management Server's public URL.",
			},
		},
	}
}

// ResourceManagementCommandConnectCloudServicesStateUpgradeV0 converts connected_at from TypeMap to TypeList.
func ResourceManagementCommandConnectCloudServicesStateUpgradeV0(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
	return UpgradeMapsToLists(rawState, "connected_at"), nil
}
