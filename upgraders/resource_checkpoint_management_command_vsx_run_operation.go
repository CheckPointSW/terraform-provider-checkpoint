package upgraders

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceManagementCommandVsxRunOperationV0 is the V0 schema where the operation parameter fields were TypeMap.
func ResourceManagementCommandVsxRunOperationV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"operation": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the operation to run. Each operation has its specific parameters.<br>The available operations are:<ul><li><i>upgrade</i> - Upgrades the VSX Gateway or VSX Cluster object to a higher version</li><li><i>downgrade</i> - Downgrades the VSX Gateway or VSX Cluster object to a lower version</li><li><i>add-member</i> - Adds a new VSX Cluster member object</li><li><i>remove-member</i> - Removes a VSX Cluster member object</li><li><i>reconf-gw</i> - Reconfigures a VSX Gateway after a clean install</li><li><i>reconf-member</i> - Reconfigures a VSX Cluster member after a clean install</li></ul>.",
			},
			"add_member_params": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Parameters for the operation to add a VSX Cluster member.",
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv4_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The IPv4 address of the management interface of the VSX Cluster member.",
						},
						"ipv4_sync_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The IPv4 address of the sync interface of the VSX Cluster member.",
						},
						"member_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Name of the new VSX Cluster member object.",
						},
						"vsx_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Name of the VSX Cluster object.",
						},
						"vsx_uid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "UID of the VSX Cluster object.",
						},
					},
				},
			},
			"downgrade_params": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Parameters for the operation to downgrade a VSX Gateway or VSX Cluster object to a lower version.<br>In case the current version is already the target version, or is lower than the target version, no change is done.",
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"target_version": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The target version.",
						},
						"vsx_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Name of the VSX Gateway or VSX Cluster object.",
						},
						"vsx_uid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "UID of the VSX Gateway or VSX Cluster object.",
						},
					},
				},
			},
			"reconf_gw_params": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Parameters for the operation to reconfigure a VSX Gateway after a clean install.",
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv4_corexl_number": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Number of IPv4 CoreXL Firewall instances on the target VSX Gateway.<br>Valid values:<br><ul><li>To configure CoreXL Firewall instances, enter an integer greater or equal to 2.</li><li>To disable CoreXL, enter 1.</li></ul>.",
							Default:     2,
						},
						"one_time_password": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A password required for establishing a Secure Internal Communication (SIC). Enter the same password you used during the First Time Configuration Wizard on the target VSX Gateway.",
						},
						"vsx_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Name of the VSX Gateway object.",
						},
						"vsx_uid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "UID of the VSX Gateway object.",
						},
					},
				},
			},
			"reconf_member_params": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Parameters for the operation to reconfigure a VSX Cluster member after a clean install.",
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv4_corexl_number": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Number of IPv4 CoreXL Firewall instances on the target VSX Cluster member.<br>Valid values:<br><ul><li>To configure CoreXL Firewall instances, enter an integer greater or equal to 2.</li><li>To disable CoreXL, enter 1.</li></ul>Important - The CoreXL configuration must be the same on all the cluster members.",
							Default:     2,
						},
						"member_uid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "UID of the VSX Cluster member object.",
						},
						"member_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Name of the VSX Cluster member object.",
						},
						"one_time_password": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A password required for establishing a Secure Internal Communication (SIC). Enter the same password you used during the First Time Configuration Wizard on the target VSX Cluster member.",
						},
					},
				},
			},
			"remove_member_params": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Parameters for the operation to remove a VSX Cluster member object.",
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"member_uid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "UID of the VSX Cluster member object.",
						},
						"member_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Name of the VSX Cluster member object.",
						},
					},
				},
			},
			"upgrade_params": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Parameters for the operation to upgrade a VSX Gateway or VSX Cluster object to a higher version.<br>In case the current version is already the target version, or is higher than the target version, no change is done.",
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"target_version": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The target version.",
						},
						"vsx_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Name of the VSX Gateway or VSX Cluster object.",
						},
						"vsx_uid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "UID of the VSX Gateway or VSX Cluster object.",
						},
					},
				},
			},
			"task_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Operation task UID.",
			},
		},
	}
}

// ResourceManagementCommandVsxRunOperationStateUpgradeV0 converts the TypeMap parameter fields to TypeList.
func ResourceManagementCommandVsxRunOperationStateUpgradeV0(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
	return UpgradeMapsToLists(rawState,
		"add_member_params", "downgrade_params", "reconf_gw_params",
		"reconf_member_params", "remove_member_params", "upgrade_params",
	), nil
}
