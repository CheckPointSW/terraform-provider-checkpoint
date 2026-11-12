package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func dataSourceManagementCloudLicensePool() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementCloudLicensePoolRead,
		Schema: map[string]*schema.Schema{
			"pool": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Pool name.",
			},
			"ck": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Certificate Key. Required to identify a specific pool when multiple pools share the same name.",
			},
			"cks": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of the licenses CKs (Certificate Keys) that belong to this license pool.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ck": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The license CK (Certificate Key).",
						},
						"expired": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether this CK is expired.",
						},
					},
				},
			},
			"available_quota": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The difference between the pool's total quota and the total cores quantity of the pool's subscribed Security Gateways.",
			},
			"default_pool": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "All new CloudGuard Gateways are automatically subscribed to the default license pool.",
			},
			"total_quota": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A license pool total quota is the total quantity of all the virtual cores provided by all the Central Licenses in this pool.",
			},
			"subscribed_gateways": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of the subscribed CloudGuard Gateways of this license pool.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Gateway name.",
						},
						"uid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Gateway's unique identifier.",
						},
						"used_quota": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Cores quantity used by the gateway.",
						},
						"domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Domain name.",
						},
					},
				},
			},
		},
	}
}

func dataSourceManagementCloudLicensePoolRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{}

	if v, ok := d.GetOk("pool"); ok {
		payload["pool"] = v.(string)
	}

	if v, ok := d.GetOk("ck"); ok {
		payload["ck"] = v.(string)
	}

	showCloudLicensePoolRes, err := client.ApiCall("show-cloud-license-pool", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showCloudLicensePoolRes.Success {
		if objectNotFound(showCloudLicensePoolRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showCloudLicensePoolRes.ErrorMsg)
	}

	cloudLicensePool := showCloudLicensePoolRes.GetData()

	log.Println("Read CloudLicensePool - Show JSON = ", cloudLicensePool)

	if v := cloudLicensePool["pool"]; v != nil {
		d.SetId(v.(string))
		_ = d.Set("pool", v)
	}

	if cloudLicensePool["cks"] != nil {

		cksList := cloudLicensePool["cks"].([]interface{})

		if len(cksList) > 0 {

			var cksListToReturn []map[string]interface{}

			for i := range cksList {

				cksMap := cksList[i].(map[string]interface{})

				cksMapToAdd := make(map[string]interface{})

				if v := cksMap["ck"]; v != nil {
					cksMapToAdd["ck"] = v
				}
				if v := cksMap["expired"]; v != nil {
					cksMapToAdd["expired"] = v
				}

				cksListToReturn = append(cksListToReturn, cksMapToAdd)
			}

			_ = d.Set("cks", cksListToReturn)
		}
	} else {
		_ = d.Set("cks", nil)
	}

	if v := cloudLicensePool["available-quota"]; v != nil {
		_ = d.Set("available_quota", v)
	}

	if v := cloudLicensePool["default-pool"]; v != nil {
		_ = d.Set("default_pool", v)
	}

	if v := cloudLicensePool["total-quota"]; v != nil {
		_ = d.Set("total_quota", v)
	}

	if cloudLicensePool["subscribed-gateways"] != nil {

		subscribedGatewaysList := cloudLicensePool["subscribed-gateways"].([]interface{})

		if len(subscribedGatewaysList) > 0 {

			var subscribedGatewaysListToReturn []map[string]interface{}

			for i := range subscribedGatewaysList {

				subscribedGatewaysMap := subscribedGatewaysList[i].(map[string]interface{})

				subscribedGatewaysMapToAdd := make(map[string]interface{})

				if v := subscribedGatewaysMap["name"]; v != nil {
					subscribedGatewaysMapToAdd["name"] = v
				}
				if v := subscribedGatewaysMap["domain"]; v != nil {
					subscribedGatewaysMapToAdd["domain"] = v
				}
				if v := subscribedGatewaysMap["used-quota"]; v != nil {
					subscribedGatewaysMapToAdd["used_quota"] = v
				}
				if v := subscribedGatewaysMap["uid"]; v != nil {
					subscribedGatewaysMapToAdd["uid"] = v
				}

				subscribedGatewaysListToReturn = append(subscribedGatewaysListToReturn, subscribedGatewaysMapToAdd)
			}

			_ = d.Set("subscribed_gateways", subscribedGatewaysListToReturn)
		}
	} else {
		_ = d.Set("subscribed_gateways", nil)
	}

	return nil

}
