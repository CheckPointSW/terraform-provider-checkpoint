package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementMobileAccessProfileSection() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceManagementMobileAccessProfileSectionRead,

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
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}
func dataSourceManagementMobileAccessProfileSectionRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showMobileAccessProfileSectionRes, err := client.ApiCall("show-mobile-access-profile-section", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showMobileAccessProfileSectionRes.Success {
		if objectNotFound(showMobileAccessProfileSectionRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showMobileAccessProfileSectionRes.ErrorMsg)
	}

	mobileAccessProfileSection := showMobileAccessProfileSectionRes.GetData()

	log.Println("Read MobileAccessProfileSection - Show JSON = ", mobileAccessProfileSection)

	if v := mobileAccessProfileSection["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := mobileAccessProfileSection["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if mobileAccessProfileSection["tags"] != nil {
		tagsJson, ok := mobileAccessProfileSection["tags"].([]interface{})
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

	return nil

}
