package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"reflect"
	"strconv"
)

func dataSourceManagementRadiusServer() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementRadiusServerRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object unique identifier.",
			},
			"server": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The UID or Name of the host that is the RADIUS Server.",
			},
			"shared_secret": {
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "The secret between the RADIUS server and the Security Gateway.",
			},
			"service": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The UID or Name of the Service to which the RADIUS server listens.",
			},
			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The version can be either RADIUS Version 1.0, which is RFC 2138 compliant, and RADIUS Version 2.0 which is RFC 2865 compliant.",
			},
			"protocol": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of authentication protocol that will be used when authenticating the user to the RADIUS server.",
			},
			"priority": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The priority of the RADIUS Server in case it is a member of a RADIUS Group.",
			},
			"accounting": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Accounting settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_ip_pool_management": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "IP pool management, enables Accounting service.",
						},
						"accounting_service": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The UID or Name of the the accounting interface to notify the server when users login and logout which will then lock and release the IP addresses that the server allocated to those users.",
						},
					},
				},
			},
			"tags": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"color": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Color of the object. Should be one of existing colors.",
			},
			"comments": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Comments string.",
			},
		},
	}
}

func dataSourceManagementRadiusServerRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showRadiusServerRes, err := client.ApiCall("show-radius-server", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showRadiusServerRes.Success {
		return fmt.Errorf(showRadiusServerRes.ErrorMsg)
	}

	radiusServer := showRadiusServerRes.GetData()
	log.Println("Read Radius Server - Show JSON = ", radiusServer)

	if v := radiusServer["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := radiusServer["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := radiusServer["server"]; v != nil {
		_ = d.Set("server", v)
	}

	if v := radiusServer["shared-secret"]; v != nil {
		_ = d.Set("shared_secret", v)
	}

	if v := radiusServer["service"]; v != nil {
		_ = d.Set("service", v)
	}

	if v := radiusServer["version"]; v != nil {
		_ = d.Set("version", v)
	}

	if v := radiusServer["protocol"]; v != nil {
		_ = d.Set("protocol", v)
	}

	if v := radiusServer["priority"]; v != nil {
		_ = d.Set("priority", v)
	}

	if radiusServer["accounting"] != nil {
		accountingMap := radiusServer["accounting"].(map[string]interface{})

		accountingMapToReturn := make(map[string]interface{})

		if v, _ := accountingMap["enable-ip-pool-management"]; v != nil {
			accountingMapToReturn["enable_ip_pool_management"] = strconv.FormatBool(v.(bool))
		}

		if v, _ := accountingMap["accounting-service"]; v != "" && v != nil {
			accountingMapToReturn["accounting_service"] = v
		}

		_, accountingInConf := d.GetOk("accounting")
		defaultAccounting := map[string]interface{}{"enable_ip_pool_management": "false"}
		if reflect.DeepEqual(defaultAccounting, accountingMapToReturn) && !accountingInConf {
			_ = d.Set("accounting", map[string]interface{}{})
		} else {
			_ = d.Set("accounting", accountingMapToReturn)
		}

	} else {
		_ = d.Set("accounting", nil)
	}

	if radiusServer["tags"] != nil {
		tagsJson := radiusServer["tags"].([]interface{})
		var tagsIds = make([]string, 0)
		if len(tagsJson) > 0 {
			// Create slice of tag names
			for _, tag := range tagsJson {
				tag := tag.(map[string]interface{})
				tagsIds = append(tagsIds, tag["name"].(string))
			}
		}
		_ = d.Set("tags", tagsIds)
	} else {
		_ = d.Set("tags", nil)
	}

	if v := radiusServer["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := radiusServer["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil
}
