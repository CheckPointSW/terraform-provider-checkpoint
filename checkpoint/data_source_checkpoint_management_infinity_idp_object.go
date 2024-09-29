package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementInfinityIdpObject() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceManagementDeleteInfinityIdpObjectRead,
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
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description string.",
			},
			"display_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Entity name in the Management Server.",
			},
			"ext_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Entity unique identifier in the Identity Provider.",
			},
			"idp_display_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Identity Provider name in Management Server.",
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
			"object_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Entity type - can be user/group/machine.",
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

func dataSourceManagementDeleteInfinityIdpObjectRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showInfinityIdpRes, err := client.ApiCall("show-infinity-idp-object", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showInfinityIdpRes.Success {
		return fmt.Errorf(showInfinityIdpRes.ErrorMsg)
	}

	infinityIdp := showInfinityIdpRes.GetData()

	log.Println("Read Infinity-Idp-Object - Show JSON = ", infinityIdp)

	if v := infinityIdp["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := infinityIdp["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := infinityIdp["description"]; v != nil {
		_ = d.Set("description", v)
	}
	if v := infinityIdp["display-name"]; v != nil {
		_ = d.Set("display_name", v)
	}
	if v := infinityIdp["ext-id"]; v != nil {
		_ = d.Set("ext_id", v)
	}
	if v := infinityIdp["idp-display-name"]; v != nil {
		_ = d.Set("idp_display_name", v)
	}
	if v := infinityIdp["idp-id"]; v != nil {
		_ = d.Set("idp_id", v)
	}
	if v := infinityIdp["idp-name"]; v != nil {
		_ = d.Set("idp_name", v)
	}
	if v := infinityIdp["object-type"]; v != nil {
		_ = d.Set("object_type", v)
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
