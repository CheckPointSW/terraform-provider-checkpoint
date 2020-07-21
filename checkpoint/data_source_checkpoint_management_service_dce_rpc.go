package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementServiceDceRpc() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementServiceDceRpcRead,
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
			"interface_uuid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Network interface UUID.",
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

func dataSourceManagementServiceDceRpcRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showServiceDceRpcRes, err := client.ApiCall("show-service-dce-rpc", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showServiceDceRpcRes.Success {
		return fmt.Errorf(showServiceDceRpcRes.ErrorMsg)
	}

	serviceDceRpc := showServiceDceRpcRes.GetData()

	log.Println("Read ServiceDceRpc - Show JSON = ", serviceDceRpc)

	if v := serviceDceRpc["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := serviceDceRpc["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := serviceDceRpc["interface-uuid"]; v != nil {
		_ = d.Set("interface_uuid", v)
	}

	if v := serviceDceRpc["keep-connections-open-after-policy-installation"]; v != nil {
		_ = d.Set("keep_connections_open_after_policy_installation", v)
	}

	if serviceDceRpc["tags"] != nil {
		tagsJson, ok := serviceDceRpc["tags"].([]interface{})
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

	if v := serviceDceRpc["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := serviceDceRpc["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if serviceDceRpc["groups"] != nil {
		groupsJson, ok := serviceDceRpc["groups"].([]interface{})
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
