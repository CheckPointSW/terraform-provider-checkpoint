package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func resourceManagementSetCloudLicensePool() *schema.Resource {
	return &schema.Resource{
		Create: createManagementSetCloudLicensePool,
		Read:   readManagementSetCloudLicensePool,
		Delete: deleteManagementSetCloudLicensePool,
		Schema: map[string]*schema.Schema{
			"pool": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Pool name.",
			},
			"ck": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Contract Key. Required to identify a specific pool when multiple pools share the same name.",
			},
			"default_pool": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Set pool to default. This value can only be changed from false to true. To disable the current default pool, you must set a different pool as the default.",
			},
			"migrate_gateways": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Move gateways from current default pool to the new default pool. Required when default-pool parameter is set to true.",
			},
			"assigned_gateways": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Attach security gateways to the pool. The attached gateways will use licenses from this pool.",
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"gateway": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Gateway name or uid.",
						},
						"domain": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Domain name or uid. Required when running from MDS context.",
						},
					},
				},
			},
			"task_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "set-cloud-license-pool task UID. Use show-task command to check the progress of the task.",
			},
		},
	}
}

func createManagementSetCloudLicensePool(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("pool"); ok {
		payload["pool"] = v.(string)
	}

	if v, ok := d.GetOk("ck"); ok {
		payload["ck"] = v.(string)
	}

	if v, ok := d.GetOkExists("default_pool"); ok {
		payload["default-pool"] = v.(bool)
	}

	if v, ok := d.GetOkExists("migrate_gateways"); ok {
		payload["migrate-gateways"] = v.(bool)
	}

	if v, ok := d.GetOk("assigned_gateways"); ok {

		assignedGatewaysList := v.([]interface{})

		if len(assignedGatewaysList) > 0 {

			var assignedGatewaysPayload []map[string]interface{}

			for i := range assignedGatewaysList {

				Payload := make(map[string]interface{})

				if v, ok := d.GetOk("assigned_gateways." + strconv.Itoa(i) + ".gateway"); ok {
					Payload["gateway"] = v.(string)
				}
				if v, ok := d.GetOk("assigned_gateways." + strconv.Itoa(i) + ".domain"); ok {
					Payload["domain"] = v.(string)
				}
				assignedGatewaysPayload = append(assignedGatewaysPayload, Payload)
			}
			payload["assigned-gateways"] = assignedGatewaysPayload
		}
	}

	SetCloudLicensePoolRes, err := client.ApiCall("set-cloud-license-pool", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !SetCloudLicensePoolRes.Success {
		return fmt.Errorf(SetCloudLicensePoolRes.ErrorMsg)
	}

	d.SetId("set-cloud-license-pool-" + acctest.RandString(10))
	_ = d.Set("task_id", resolveTaskId(SetCloudLicensePoolRes.GetData()))
	return nil
}

func readManagementSetCloudLicensePool(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementSetCloudLicensePool(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
