package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementApiSettings() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementApiSettingsRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object name. Must be unique in the domain.",
			},
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object unique identifier.",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object type.",
			},
			"accepted_api_calls_from": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Clients allowed to connect to the API Server.",
			},
			"automatic_start": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "MGMT API will start after server will start.",
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

func dataSourceManagementApiSettingsRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	payload := make(map[string]interface{})

	showApiSettingsRes, err := client.ApiCall("show-api-settings", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		fmt.Errorf(err.Error())
	}
	if !showApiSettingsRes.Success {
		fmt.Errorf(showApiSettingsRes.ErrorMsg)
	}

	apiSettings := showApiSettingsRes.GetData()

	log.Println("Read Api Settings - Show JSON = ", apiSettings)

	d.SetId("show-api-settings-" + acctest.RandString(10))

	if v := apiSettings["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := apiSettings["uid"]; v != nil {
		_ = d.Set("uid", v)
	}

	if v := apiSettings["type"]; v != nil {
		_ = d.Set("type", v)
	}

	if v := apiSettings["accepted-api-calls-from"]; v != nil {
		_ = d.Set("accepted_api_calls_from", v)
	}

	if v := apiSettings["automatic-start"]; v != nil {
		_ = d.Set("automatic_start", v)
	}

	if apiSettings["tags"] != nil {
		tagsJson := apiSettings["tags"].([]interface{})
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

	if v := apiSettings["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := apiSettings["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil
}
