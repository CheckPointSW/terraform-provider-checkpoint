package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementExceptionGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementExceptionGroupRead,
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
			"applied_profile": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The threat profile to apply this group to in the case of apply-on threat-rules-with-specific-profile.",
			},
			"apply_on": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "An exception group can be set to apply on all threat rules, all threat rules which have a specific profile, or those rules manually chosen by the user.",
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

func dataSourceManagementExceptionGroupRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showExceptionGroupRes, err := client.ApiCall("show-exception-group", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showExceptionGroupRes.Success {
		return fmt.Errorf(showExceptionGroupRes.ErrorMsg)
	}

	exceptionGroup := showExceptionGroupRes.GetData()

	log.Println("Read ExceptionGroup - Show JSON = ", exceptionGroup)

	if v := exceptionGroup["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := exceptionGroup["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := exceptionGroup["applied-profile"]; v != nil {
		_ = d.Set("applied_profile", v)
	}

	if v := exceptionGroup["apply-on"]; v != nil {
		_ = d.Set("apply_on", v)
	}

	if exceptionGroup["tags"] != nil {
		tagsJson, ok := exceptionGroup["tags"].([]interface{})
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

	if v := exceptionGroup["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := exceptionGroup["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil
}
