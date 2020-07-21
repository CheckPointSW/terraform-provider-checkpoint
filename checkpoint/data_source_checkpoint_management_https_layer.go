package checkpoint

import (
	"fmt"
	"log"

	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceManagementHttpsLayer() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementHttpsLayerRead,
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
			"shared": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Define the Layer as Shared (TRUE/FALSE).",
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

func dataSourceManagementHttpsLayerRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showHttpsLayerRes, err := client.ApiCall("show-https-layer", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showHttpsLayerRes.Success {
		return fmt.Errorf(showHttpsLayerRes.ErrorMsg)
	}

	httpsLayer := showHttpsLayerRes.GetData()

	log.Println("Read HttpsLayer - Show JSON = ", httpsLayer)

	if v := httpsLayer["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := httpsLayer["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := httpsLayer["shared"]; v != nil {
		_ = d.Set("shared", v)
	}

	if httpsLayer["tags"] != nil {
		tagsJson, ok := httpsLayer["tags"].([]interface{})
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

	if v := httpsLayer["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := httpsLayer["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil
}
