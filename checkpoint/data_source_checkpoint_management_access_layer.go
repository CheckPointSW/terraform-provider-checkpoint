package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementAccessLayer() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementAccessLayerRead,
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
			"applications_and_url_filtering": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether to enable Applications & URL Filtering blade on the layer.",
			},
			"content_awareness": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether to enable Content Awareness blade on the layer.",
			},
			"detect_using_x_forward_for": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether to use X-Forward-For HTTP header, which is added by the  proxy server to keep track of the original source IP.",
			},
			"firewall": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether to enable Firewall blade on the layer.",
			},
			"implicit_cleanup_action": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The default \"catch-all\" action for traffic that does not match any explicit or implied rules in the layer.",
			},
			"mobile_access": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether to enable Mobile Access blade on the layer.",
			},
			"shared": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether this layer is shared.",
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

func dataSourceManagementAccessLayerRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showAccessLayerRes, err := client.ApiCall("show-access-layer", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showAccessLayerRes.Success {
		return fmt.Errorf(showAccessLayerRes.ErrorMsg)
	}

	accessLayer := showAccessLayerRes.GetData()

	log.Println("Read AccessLayer - Show JSON = ", accessLayer)

	if v := accessLayer["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := accessLayer["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := accessLayer["applications-and-url-filtering"]; v != nil {
		_ = d.Set("applications_and_url_filtering", v)
	}

	if v := accessLayer["content-awareness"]; v != nil {
		_ = d.Set("content_awareness", v)
	}

	if v := accessLayer["detect-using-x-forward-for"]; v != nil {
		_ = d.Set("detect_using_x_forward_for", v)
	}

	if v := accessLayer["firewall"]; v != nil {
		_ = d.Set("firewall", v)
	}

	if v := accessLayer["implicit-cleanup-action"]; v != nil {
		_ = d.Set("implicit_cleanup_action", v)
	}

	if v := accessLayer["mobile-access"]; v != nil {
		_ = d.Set("mobile_access", v)
	}

	if v := accessLayer["shared"]; v != nil {
		_ = d.Set("shared", v)
	}

	if accessLayer["tags"] != nil {
		tagsJson, ok := accessLayer["tags"].([]interface{})
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

	if v := accessLayer["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := accessLayer["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil
}
