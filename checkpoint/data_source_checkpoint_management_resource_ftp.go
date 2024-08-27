package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementResourceFtp() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementResourceFtpRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object uid.",
			},
			"resource_matching_method": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "GET allows Downloads from the server to the client. PUT allows Uploads from the client to the server.",
			},
			"exception_track": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The UID or Name of the exception track to be used to log actions taken as a result of a match on the resource.",
			},
			"resources_path": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Refers to a location on the FTP server.",
			},

			"cvp": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Configure CVP inspection on mail messages.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_cvp": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Select to enable the Content Vectoring Protocol.",
						},
						"server": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The UID or Name of the CVP server, make sure the CVP server is already be defined as an OPSEC Application.",
						},
						"allowed_to_modify_content": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Configures the CVP server to inspect but not modify content.",
						},
						"reply_order": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Designates when the CVP server returns data to the Security Gateway security server.",
						},
					},
				},
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
func dataSourceManagementResourceFtpRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showResourceFtpRes, err := client.ApiCall("show-resource-ftp", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showResourceFtpRes.Success {
		if objectNotFound(showResourceFtpRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showResourceFtpRes.ErrorMsg)
	}

	resourceFtp := showResourceFtpRes.GetData()

	log.Println("Read ResourceFtp - Show JSON = ", resourceFtp)

	if v := resourceFtp["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := resourceFtp["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := resourceFtp["exception-track"]; v != nil {

		objMap := v.(map[string]interface{})
		if v := objMap["name"]; v != nil {
			_ = d.Set("exception_track", v)
		}
	}

	if v := resourceFtp["resources-path"]; v != nil {
		_ = d.Set("resources_path", v)
	}

	if v := resourceFtp["resource-matching-method"]; v != nil {
		_ = d.Set("resource_matching_method", v)
	}

	if resourceFtp["cvp"] != nil {

		cvpMap := resourceFtp["cvp"].(map[string]interface{})

		cvpMapToReturn := make(map[string]interface{})

		if v, _ := cvpMap["enable-cvp"]; v != nil {
			cvpMapToReturn["enable_cvp"] = v
		}
		if v, _ := cvpMap["server"]; v != nil {

			objMap := v.(map[string]interface{})

			if v := objMap["name"]; v != nil {
				cvpMapToReturn["server"] = v
			}
		}
		if v, _ := cvpMap["cvp-server-is-allowed-to-modify-content"]; v != nil {
			cvpMapToReturn["allowed_to_modify_content"] = v
		}
		if v, _ := cvpMap["reply-order"]; v != nil {
			cvpMapToReturn["reply_order"] = v
		}
		_ = d.Set("cvp", []interface{}{cvpMapToReturn})
	} else {
		_ = d.Set("cvp", nil)
	}

	if resourceFtp["tags"] != nil {
		tagsJson, ok := resourceFtp["tags"].([]interface{})
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

	if v := resourceFtp["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := resourceFtp["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := resourceFtp["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := resourceFtp["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}
