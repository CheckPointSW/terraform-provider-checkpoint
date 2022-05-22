package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementRepositoryScript() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementRepositoryScriptRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"script_body": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The entire content of the script.",
			},
			"script_body_base64": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The entire content of the script encoded in Base64.",
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

func dataSourceManagementRepositoryScriptRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showRepositoryScriptRes, err := client.ApiCall("show-repository-script", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showRepositoryScriptRes.Success {
		if objectNotFound(showRepositoryScriptRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showRepositoryScriptRes.ErrorMsg)
	}

	repositoryScript := showRepositoryScriptRes.GetData()

	log.Println("Read RepositoryScript - Show JSON = ", repositoryScript)

	if v := repositoryScript["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := repositoryScript["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := repositoryScript["script-body"]; v != nil {
		_ = d.Set("script_body", v)
	}

	if v := repositoryScript["script-body-base64"]; v != nil {
		_ = d.Set("script_body_base64", v)
	}

	if repositoryScript["tags"] != nil {
		tagsJson, ok := repositoryScript["tags"].([]interface{})
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

	if v := repositoryScript["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := repositoryScript["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil

}
