package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementThreatLayer() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementThreatLayerRead,
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
			"ips_layer": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "N/A",
			},
			"parent_layer": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "N/A",
			},
		},
	}
}

func dataSourceManagementThreatLayerRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showThreatLayerRes, err := client.ApiCall("show-threat-layer", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showThreatLayerRes.Success {
		return fmt.Errorf(showThreatLayerRes.ErrorMsg)
	}

	threatLayer := showThreatLayerRes.GetData()

	log.Println("Read Threat Layer - Show JSON = ", threatLayer)

	if v := threatLayer["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := threatLayer["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if threatLayer["tags"] != nil {
		tagsJson := threatLayer["tags"].([]interface{})
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

	if v := threatLayer["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := threatLayer["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := threatLayer["ips-layer"]; v != nil {
		_ = d.Set("ips_layer", v)
	}

	if v := threatLayer["parent-layer"]; v != nil {
		_ = d.Set("parent_layer", v)
	}

	return nil
}
