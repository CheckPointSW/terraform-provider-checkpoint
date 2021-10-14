package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementAccessPointName() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementAccessPointNameRead,
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
			"apn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "APN name.",
			},
			"enforce_end_user_domain": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable enforce end user domain.",
			},
			"block_traffic_other_end_user_domains": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Block MS to MS traffic between this and other APN end user domains.",
			},
			"block_traffic_this_end_user_domain": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Block MS to MS traffic within this end user domain.",
			},
			"end_user_domain": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "End user domain name or UID.",
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
func dataSourceManagementAccessPointNameRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showAccessPointNameRes, err := client.ApiCall("show-access-point-name", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showAccessPointNameRes.Success {
		return fmt.Errorf(showAccessPointNameRes.ErrorMsg)
	}

	accessPointName := showAccessPointNameRes.GetData()

	log.Println("Read AccessPointName - Show JSON = ", accessPointName)

	if v := accessPointName["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := accessPointName["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := accessPointName["apn"]; v != nil {
		_ = d.Set("apn", v)
	}

	if v := accessPointName["enforce-end-user-domain"]; v != nil {
		_ = d.Set("enforce_end_user_domain", v)
	}

	if v := accessPointName["block-traffic-other-end-user-domains"]; v != nil {
		_ = d.Set("block_traffic_other_end_user_domains", v)
	}

	if v := accessPointName["block-traffic-this-end-user-domain"]; v != nil {
		_ = d.Set("block_traffic_this_end_user_domain", v)
	}

	if v := accessPointName["end-user-domain"]; v != nil {
		_ = d.Set("end_user_domain", v)
	}

	if accessPointName["tags"] != nil {
		tagsJson, ok := accessPointName["tags"].([]interface{})
		if ok {
			tagsIds := make([]string, 0)
			if len(tagsJson) > 0 {
				for _, tags := range tagsJson {
					tags := tags.(map[string]interface{})
					tagsIds = append(tagsIds, tags["name"].(string))
				}
			}
			_ = d.Set("tags", tagsIds)
		}
	} else {
		_ = d.Set("tags", nil)
	}

	if v := accessPointName["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := accessPointName["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil
}
