package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementServiceCitrixTcp() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementServiceCitrixTcpRead,
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
			"application": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Citrix application name.",
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

func dataSourceManagementServiceCitrixTcpRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showServiceCitrixTcpRes, err := client.ApiCall("show-service-citrix-tcp", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showServiceCitrixTcpRes.Success {
		return fmt.Errorf(showServiceCitrixTcpRes.ErrorMsg)
	}

	serviceCitrixTcp := showServiceCitrixTcpRes.GetData()

	log.Println("Read ServiceCitrixTcp - Show JSON = ", serviceCitrixTcp)

	if v := serviceCitrixTcp["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := serviceCitrixTcp["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := serviceCitrixTcp["application"]; v != nil {
		_ = d.Set("application", v)
	}

	if serviceCitrixTcp["tags"] != nil {
		tagsJson, ok := serviceCitrixTcp["tags"].([]interface{})
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

	if v := serviceCitrixTcp["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := serviceCitrixTcp["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil
}
