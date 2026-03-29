package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func dataSourceManagementShowUpdatableObject() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementShowUpdatableObjectRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name. Should be unique in the domain.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object unique identifier.",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object type.",
			},
			"name_in_updatable_objects_repository": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object name in the Updatable Objects Repository.",
			},
			"uid_in_updatable_objects_repository": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Unique identifier of the object in the Updatable Objects Repository.",
			},
			"additional_properties": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Additional properties on the object.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Sorts results by the given field in ascending order.",
						},
						"info_text": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Information about the Updatable Object IP ranges source.",
						},
						"info_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URL of the Updatable Object IP ranges source.",
						},
						"uri": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URI of the Updatable Object under the Updatable Objects Repository.",
						},
					},
				},
			},
			"updatable_object_meta_info": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"updated_on_updatable_objects_repository": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Last update time from the Updatable Objects Repository",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"iso_8601": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "",
									},
									"posix": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "",
									},
								},
							},
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
func dataSourceManagementShowUpdatableObjectRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := make(map[string]interface{})

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}
	showUpdatableObjectRes, err := client.ApiCall("show-updatable-object", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	if !showUpdatableObjectRes.Success {
		return fmt.Errorf("%s", showUpdatableObjectRes.ErrorMsg)
	}

	updatableObjectJson := showUpdatableObjectRes.GetData()

	log.Println("Read updatable-object - Show JSON = ", updatableObjectJson)

	if v := updatableObjectJson["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := updatableObjectJson["name"]; v != nil {
		_ = d.Set("name", v)
	}
	if v := updatableObjectJson["type"]; v != nil {
		_ = d.Set("type", v)
	}
	if v := updatableObjectJson["name-in-updatable-objects-repository"]; v != nil {

		_ = d.Set("name_in_updatable_objects_repository", v.(string))

	}

	if v := updatableObjectJson["uid-in-updatable-objects-repository"]; v != nil {
		_ = d.Set("uid_in_updatable_objects_repository", v.(string))
	}
	if updatableObjectJson["additional-properties"] != nil {

		additionalPropertiesMap := updatableObjectJson["additional-properties"].(map[string]interface{})

		additionalPropertiesMapToReturn := make(map[string]interface{})

		if v := additionalPropertiesMap["description"]; v != nil {
			additionalPropertiesMapToReturn["description"] = v
		}
		if v := additionalPropertiesMap["info-text"]; v != nil {
			additionalPropertiesMapToReturn["info_text"] = v
		}
		if v := additionalPropertiesMap["info-url"]; v != nil {
			additionalPropertiesMapToReturn["info_url"] = v
		}
		if v := additionalPropertiesMap["uri"]; v != nil {
			additionalPropertiesMapToReturn["uri"] = v
		}

		_ = d.Set("additional_properties", []interface{}{additionalPropertiesMapToReturn})

	} else {
		_ = d.Set("additional_properties", nil)
	}
	if v := updatableObjectJson["updatable-object-meta-info"]; v != nil {
		metaInfoMap := v.(map[string]interface{})
		metaInfoMapToReturn := make(map[string]interface{})
		if v := metaInfoMap["updated-on-updatable-objects-repository"]; v != nil {
			innerMap := v.(map[string]interface{})
			innerMapToReturn := make(map[string]interface{})
			if v := innerMap["iso-8601"]; v != nil {
				innerMapToReturn["iso_8601"] = v
			}
			if v := innerMap["posix"]; v != nil {
				innerMapToReturn["posix"] = v
			}
			metaInfoMapToReturn["updated_on_updatable_objects_repository"] = []interface{}{innerMapToReturn}
		}
		_ = d.Set("updatable_object_meta_info", []interface{}{metaInfoMapToReturn})
	}
	if updatableObjectJson["tags"] != nil {
		tagsJson, ok := updatableObjectJson["tags"].([]interface{})
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

	if v := updatableObjectJson["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := updatableObjectJson["comments"]; v != nil {
		_ = d.Set("comments", v)
	}
	return nil
}
