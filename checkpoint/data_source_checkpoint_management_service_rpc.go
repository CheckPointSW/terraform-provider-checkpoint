package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementServiceRpc() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementServiceRpcRead,
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
			"keep_connections_open_after_policy_installation": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections.",
			},
			"program_number": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "N/A",
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
			"groups": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of group identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceManagementServiceRpcRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showServiceRpcRes, err := client.ApiCall("show-service-rpc", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showServiceRpcRes.Success {
		if objectNotFound(showServiceRpcRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showServiceRpcRes.ErrorMsg)
	}

	serviceRpc := showServiceRpcRes.GetData()

	log.Println("Read ServiceRpc - Show JSON = ", serviceRpc)

	if v := serviceRpc["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := serviceRpc["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := serviceRpc["keep-connections-open-after-policy-installation"]; v != nil {
		_ = d.Set("keep_connections_open_after_policy_installation", v)
	}

	if v := serviceRpc["program-number"]; v != nil {
		_ = d.Set("program_number", v)
	}

	if serviceRpc["tags"] != nil {
		tagsJson, ok := serviceRpc["tags"].([]interface{})
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

	if v := serviceRpc["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := serviceRpc["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if serviceRpc["groups"] != nil {
		groupsJson, ok := serviceRpc["groups"].([]interface{})
		if ok {
			groupsIds := make([]string, 0)
			if len(groupsJson) > 0 {
				for _, groups := range groupsJson {
					groups := groups.(map[string]interface{})
					groupsIds = append(groupsIds, groups["name"].(string))
				}
			}
			_ = d.Set("groups", groupsIds)
		}
	} else {
		_ = d.Set("groups", nil)
	}

	return nil
}
