package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strings"
)

func dataSourceManagementCheckpointSaseDataCenterServer() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCheckpointSaseDataCenterServerRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name. Must be unique in the domain.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object unique identifier.",
			},
			"connect_to": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "connected-portal: Connect to the connected Check Point Portal Account. other-portal: Connect to a different Check Point Portal Account.",
			},
			"hostname": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL from Check Point Portal. Required for connect-to: other-portal.",
			},
			"client_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Client ID for Check Point SASE account. Required for connect-to: other-portal.",
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
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceCheckpointSaseDataCenterServerRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	var name string
	var uid string

	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}
	if v, ok := d.GetOk("uid"); ok {
		uid = v.(string)
	}
	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}
	showCheckpointSaseDataCenterServerRes, err := client.ApiCall("show-data-center-server", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	if !showCheckpointSaseDataCenterServerRes.Success {
		return fmt.Errorf("%s", showCheckpointSaseDataCenterServerRes.ErrorMsg)
	}
	checkpointSaseDataCenterServer := showCheckpointSaseDataCenterServerRes.GetData()

	if v := checkpointSaseDataCenterServer["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := checkpointSaseDataCenterServer["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if checkpointSaseDataCenterServer["properties"] != nil {
		propsJson, ok := checkpointSaseDataCenterServer["properties"].([]interface{})
		if ok {
			for _, prop := range propsJson {
				propMap := prop.(map[string]interface{})
				propName := strings.ReplaceAll(propMap["name"].(string), "-", "_")
				propValue := propMap["value"]
				_ = d.Set(propName, propValue)
			}
		}
	}

	if checkpointSaseDataCenterServer["tags"] != nil {
		tagsJson, ok := checkpointSaseDataCenterServer["tags"].([]interface{})
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

	if v := checkpointSaseDataCenterServer["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := checkpointSaseDataCenterServer["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil

}
