package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementServiceCompoundTcp() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementServiceCompoundTcpRead,
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
			"compound_service": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Compound service type.",
			},
			"keep_connections_open_after_policy_installation": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections.",
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

func dataSourceManagementServiceCompoundTcpRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showServiceCompoundTcpRes, err := client.ApiCall("show-service-compound-tcp", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showServiceCompoundTcpRes.Success {
		return fmt.Errorf(showServiceCompoundTcpRes.ErrorMsg)
	}

	serviceCompoundTcp := showServiceCompoundTcpRes.GetData()

	log.Println("Read ServiceCompoundTcp - Show JSON = ", serviceCompoundTcp)

	if v := serviceCompoundTcp["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := serviceCompoundTcp["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := serviceCompoundTcp["compound-service"]; v != nil {
		_ = d.Set("compound_service", v)
	}

	if v := serviceCompoundTcp["keep-connections-open-after-policy-installation"]; v != nil {
		_ = d.Set("keep_connections_open_after_policy_installation", v)
	}

	if serviceCompoundTcp["tags"] != nil {
		tagsJson, ok := serviceCompoundTcp["tags"].([]interface{})
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

	if v := serviceCompoundTcp["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := serviceCompoundTcp["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil
}
