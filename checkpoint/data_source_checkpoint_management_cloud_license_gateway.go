package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func dataSourceManagementCloudLicenseGateway() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementCloudLicenseGatewayRead,
		Schema: map[string]*schema.Schema{
			"gateway": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Security gateway name or UID.",
			},
			"domain": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Domain name or UID of security gateway. Required when running from MDS context.",
			},
			"used_quota": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of licenses used by the gateway.",
			},
			"enable_auto_distribution": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether automatic distribution is enabled for this gateway.",
			},
			"cks": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "List of CKs assigned to the gateway.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"pool": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Pool information including name, total quota, and available quota.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pool": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A group of CloudGuard Central Licenses with the same valid contract blades.",
						},
						"available_quota": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The difference between the pool's total quota and the total cores quantity of the pool's subscribed Security Gateways.",
						},
						"total_quota": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A license pool total quota is the total quantity of all the virtual cores provided by all the Central Licenses in this pool.",
						},
					},
				},
			},
			"scheduled_auto_distribution": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Time until the next scheduled automatic distribution.",
			},
		},
	}
}

func dataSourceManagementCloudLicenseGatewayRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{}

	if v, ok := d.GetOk("name"); ok {
		payload["name"] = v.(string)
	} else if v, ok := d.GetOk("uid"); ok {
		payload["uid"] = v.(string)
	} else {
		return fmt.Errorf("Either name or uid must be specified")
	}

	if v, ok := d.GetOk("gateway"); ok {
		payload["gateway"] = v.(string)
	}

	if v, ok := d.GetOk("domain"); ok {
		payload["domain"] = v.(string)
	}

	showCloudLicenseGatewayRes, err := client.ApiCall("show-cloud-license-gateway", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showCloudLicenseGatewayRes.Success {
		if objectNotFound(showCloudLicenseGatewayRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showCloudLicenseGatewayRes.ErrorMsg)
	}

	cloudLicenseGateway := showCloudLicenseGatewayRes.GetData()

	log.Println("Read CloudLicenseGateway - Show JSON = ", cloudLicenseGateway)

	if cloudLicenseGateway["gateway"] != nil {

		gatewayMap := cloudLicenseGateway["gateway"].(map[string]interface{})

		if v := gatewayMap["uid"]; v != nil {
			d.SetId(v.(string))
		}
		if v := gatewayMap["name"]; v != nil {
			_ = d.Set("gateway", v)
		}
		if v := gatewayMap["domain"]; v != nil {
			_ = d.Set("domain", v)
		}
		if v := gatewayMap["used-quota"]; v != nil {
			_ = d.Set("used_quota", v)
		}
	}

	if v := cloudLicenseGateway["enable-auto-distribution"]; v != nil {
		_ = d.Set("enable_auto_distribution", v)
	}

	if cloudLicenseGateway["cks"] != nil {
		cksJson, ok := cloudLicenseGateway["cks"].([]interface{})
		if ok {
			cksIds := make([]string, 0)
			if len(cksJson) > 0 {
				for _, cks := range cksJson {
					cksIds = append(cksIds, cks.(string))
				}
			}
			_ = d.Set("cks", cksIds)
		}
	} else {
		_ = d.Set("cks", nil)
	}

	if cloudLicenseGateway["pool"] != nil {

		poolMap := cloudLicenseGateway["pool"].(map[string]interface{})

		poolMapToReturn := make(map[string]interface{})

		if v := poolMap["pool"]; v != nil {
			poolMapToReturn["pool"] = v
		}
		if v := poolMap["available-quota"]; v != nil {
			poolMapToReturn["available_quota"] = v
		}
		if v := poolMap["total-quota"]; v != nil {
			poolMapToReturn["total_quota"] = v
		}

		_ = d.Set("pool", []interface{}{poolMapToReturn})

	} else {
		_ = d.Set("pool", nil)
	}

	if v := cloudLicenseGateway["scheduled-auto-distribution"]; v != nil {
		_ = d.Set("scheduled_auto_distribution", v)
	}

	return nil

}
