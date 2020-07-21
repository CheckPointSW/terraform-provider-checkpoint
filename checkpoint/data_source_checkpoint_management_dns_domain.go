package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementDnsDomain() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementDnsDomainRead,
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
			"is_sub_domain": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether to match sub-domains in addition to the domain itself.",
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

func dataSourceManagementDnsDomainRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showDnsDomainRes, err := client.ApiCall("show-dns-domain", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showDnsDomainRes.Success {
		return fmt.Errorf(showDnsDomainRes.ErrorMsg)
	}

	dnsDomain := showDnsDomainRes.GetData()

	log.Println("Read DnsDomain - Show JSON = ", dnsDomain)

	if v := dnsDomain["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := dnsDomain["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := dnsDomain["is-sub-domain"]; v != nil {
		_ = d.Set("is_sub_domain", v)
	}

	if dnsDomain["tags"] != nil {
		tagsJson, ok := dnsDomain["tags"].([]interface{})
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

	if v := dnsDomain["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := dnsDomain["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil
}
