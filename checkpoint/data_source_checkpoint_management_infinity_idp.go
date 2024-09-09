package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementInfinityIdp() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceManagementDeleteInfinityIdpRead,
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
			"idp_domains": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"idp_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object unique identifier.",
			},
			"idp_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object unique identifier.",
			},
			"idp_type": {
				Type:        schema.TypeString,
				Computed:    true,
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
		},
	}
}
func dataSourceManagementDeleteInfinityIdpRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showInfinityIdpRes, err := client.ApiCall("show-infinity-idp", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showInfinityIdpRes.Success {
		return fmt.Errorf(showInfinityIdpRes.ErrorMsg)
	}

	infinityIdp := showInfinityIdpRes.GetData()

	log.Println("Read Infinity-Idp - Show JSON = ", infinityIdp)

	if v := infinityIdp["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := infinityIdp["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := infinityIdp["idp-domains"]; v != nil {
		_ = d.Set("idp_domains", v.([]interface{}))
	}
	if v := infinityIdp["idp-id"]; v != nil {
		_ = d.Set("idp_id", v)
	}
	if v := infinityIdp["idp-name"]; v != nil {
		_ = d.Set("idp_name", v)
	}
	if v := infinityIdp["idp-type"]; v != nil {
		_ = d.Set("idp_type", v)
	}

	if infinityIdp["tags"] != nil {
		tagsJson := infinityIdp["tags"].([]interface{})
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

	return nil
}
