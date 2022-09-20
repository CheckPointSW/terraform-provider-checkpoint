package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementVsxRunOperation() *schema.Resource {
	return &schema.Resource{
		Create: createManagementVsxRunOperation,
		Read:   readManagementVsxRunOperation,
		Delete: deleteManagementVsxRunOperation,
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

func createManagementVsxRunOperation(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("operation"); ok {
		payload["operation"] = v.(string)
	}

	if _, ok := d.GetOk("add_member_params"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("add_member_params.ipv4_address"); ok {
			res["ipv4-address"] = v.(string)
		}
		if v, ok := d.GetOk("add_member_params.ipv4_sync_address"); ok {
			res["ipv4-sync-address"] = v.(string)
		}
		if v, ok := d.GetOk("add_member_params.member_name"); ok {
			res["member-name"] = v.(string)
		}
		if v, ok := d.GetOk("add_member_params.vsx_name"); ok {
			res["vsx-name"] = v.(string)
		}
		if v, ok := d.GetOk("add_member_params.vsx_uid"); ok {
			res["vsx-uid"] = v.(string)
		}
		payload["add-member-params"] = res
	}

	if _, ok := d.GetOk("downgrade_params"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("downgrade_params.target_version"); ok {
			res["target-version"] = v.(string)
		}
		if v, ok := d.GetOk("downgrade_params.vsx_name"); ok {
			res["vsx-name"] = v.(string)
		}
		if v, ok := d.GetOk("downgrade_params.vsx_uid"); ok {
			res["vsx-uid"] = v.(string)
		}
		payload["downgrade-params"] = res
	}

	if _, ok := d.GetOk("reconf_gw_params"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("reconf_gw_params.ipv4_corexl_number"); ok {
			res["ipv4-corexl-number"] = v
		}
		if v, ok := d.GetOk("reconf_gw_params.one_time_password"); ok {
			res["one-time-password"] = v.(string)
		}
		if v, ok := d.GetOk("reconf_gw_params.vsx_name"); ok {
			res["vsx-name"] = v.(string)
		}
		if v, ok := d.GetOk("reconf_gw_params.vsx_uid"); ok {
			res["vsx-uid"] = v.(string)
		}
		payload["reconf-gw-params"] = res
	}

	if _, ok := d.GetOk("reconf_member_params"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("reconf_member_params.ipv4_corexl_number"); ok {
			res["ipv4-corexl-number"] = v
		}
		if v, ok := d.GetOk("reconf_member_params.member_uid"); ok {
			res["member-uid"] = v.(string)
		}
		if v, ok := d.GetOk("reconf_member_params.member_name"); ok {
			res["member-name"] = v.(string)
		}
		if v, ok := d.GetOk("reconf_member_params.one_time_password"); ok {
			res["one-time-password"] = v.(string)
		}
		payload["reconf-member-params"] = res
	}

	if _, ok := d.GetOk("remove_member_params"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("remove_member_params.member_uid"); ok {
			res["member-uid"] = v.(string)
		}
		if v, ok := d.GetOk("remove_member_params.member_name"); ok {
			res["member-name"] = v.(string)
		}
		payload["remove-member-params"] = res
	}

	if _, ok := d.GetOk("upgrade_params"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("upgrade_params.target_version"); ok {
			res["target-version"] = v.(string)
		}
		if v, ok := d.GetOk("upgrade_params.vsx_name"); ok {
			res["vsx-name"] = v.(string)
		}
		if v, ok := d.GetOk("upgrade_params.vsx_uid"); ok {
			res["vsx-uid"] = v.(string)
		}
		payload["upgrade-params"] = res
	}

	VsxRunOperationRes, _ := client.ApiCall("vsx-run-operation", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if !VsxRunOperationRes.Success {
		return fmt.Errorf(VsxRunOperationRes.ErrorMsg)
	}

	d.SetId("vsx-run-operation-" + acctest.RandString(10))
	_ = d.Set("task_id", resolveTaskId(VsxRunOperationRes.GetData()))
	return readManagementVsxRunOperation(d, m)
}

func readManagementVsxRunOperation(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementVsxRunOperation(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
