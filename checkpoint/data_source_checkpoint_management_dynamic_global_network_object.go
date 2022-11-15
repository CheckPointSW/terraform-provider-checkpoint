package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementDynamicGlobalNetworkObject() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementDynamicGlobalNetworkObjectRead,
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
			"tags": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of tag objects identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.",
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

func dataSourceManagementDynamicGlobalNetworkObjectRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showDynamicGlobalNetworkObjectRes, err := client.ApiCall("show-dynamic-global-network-object", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showDynamicGlobalNetworkObjectRes.Success {
		return fmt.Errorf(showDynamicGlobalNetworkObjectRes.ErrorMsg)
	}

	dynamicGlobalNetworkObject := showDynamicGlobalNetworkObjectRes.GetData()

	log.Println("Read Dynamic Global Network Object - Show JSON = ", dynamicGlobalNetworkObject)

	if v := dynamicGlobalNetworkObject["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := dynamicGlobalNetworkObject["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if dynamicGlobalNetworkObject["tags"] != nil {
		tagsJson, ok := dynamicGlobalNetworkObject["tags"].([]interface{})
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

	if v := dynamicGlobalNetworkObject["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := dynamicGlobalNetworkObject["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil
}
