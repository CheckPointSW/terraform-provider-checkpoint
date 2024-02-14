package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strconv"
)

func dataSourceManagementDataCenterObject() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementDataCenterObjectRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Override default name on data-center.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object unique identifier.",
			},
			"data_center_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the Data Center Server the object is in.",
			},
			"name_in_data_center": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object Name in Data Center",
			},
			"uid_in_data_center": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Unique identifier of the object in the Data Center Server.",
			},
			"data_center": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Data Center Object",
				MaxItems:    1,
				Elem: &schema.Resource{
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
						"automatic_refresh": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "UID.",
						},
						"data_center_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Data Center Type.",
						},
						"properties": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Data Center properties",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Object name. Must be unique in the domain.",
									},
									"value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Object name. Must be unique in the domain.",
									},
								},
							},
						},
					},
				},
			},
			"updated_on_data_center": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Last update time of data center",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"iso_8601": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"posix": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"deleted": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates if the object is inaccessible or deleted on Data Center Server.",
			},
			"type_in_data_center": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object type in Data Center.",
			},
			"additional_properties": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Additional properties on the object.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "",
						},
						"value": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "",
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
			"groups": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of group identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ignore_warnings": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Apply changes ignoring warnings.",
			},
			"ignore_errors": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.",
			},
		},
	}
}
func dataSourceManagementDataCenterObjectRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := make(map[string]interface{})

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}
	payload["details-level"] = "full"
	showDataCenterObjRes, err := client.ApiCall("show-data-center-object", payload, client.GetSessionID(), true, client.IsProxyUsed())

	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showDataCenterObjRes.Success {
		// Handle delete resource from other clients
		if objectNotFound(showDataCenterObjRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showDataCenterObjRes.ErrorMsg)
	}

	dataCenterObj := showDataCenterObjRes.GetData()

	log.Println("Read Data Center Object - Show JSON = ", dataCenterObj)
	if v := dataCenterObj["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := dataCenterObj["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := dataCenterObj["name-in-data-center"]; v != nil {
		_ = d.Set("name_in_data_center", v)
	}
	if v := dataCenterObj["uid-in-data-center"]; v != nil {
		_ = d.Set("uid_in_data_center", v)
	}

	if v := dataCenterObj["data-center"]; v != nil {

		dataCenterMap, ok := v.(map[string]interface{})

		if ok {
			dataCenterMapToReturn := make(map[string]interface{})

			if v := dataCenterMap["name"]; v != nil {
				dataCenterMapToReturn["name"] = v
			}
			if v := dataCenterMap["uid"]; v != nil {
				dataCenterMapToReturn["uid"] = v
			}
			if v := dataCenterMap["automatic-refresh"]; v != nil {
				dataCenterMapToReturn["automatic_refresh"] = v
			}
			if v := dataCenterMap["data-center-type"]; v != nil {
				dataCenterMapToReturn["data_center_type"] = v
			}
			if v := dataCenterMap["properties"]; v != nil {

				propertiesList := v.([]interface{})

				var propertiesToReturn []map[string]interface{}

				if len(propertiesList) > 0 {

					for j := range propertiesList {

						innerMapToAdd := make(map[string]interface{})

						property := propertiesList[j].(map[string]interface{})

						if v := property["name"]; v != nil {
							innerMapToAdd["name"] = v
						}
						if v := property["value"]; v != nil {
							innerMapToAdd["value"] = v
						}

						propertiesToReturn = append(propertiesToReturn, innerMapToAdd)
					}

				}
				dataCenterMapToReturn["properties"] = propertiesToReturn
			}
			_ = d.Set("data_center", []interface{}{dataCenterMapToReturn})
		}
	} else {
		_ = d.Set("data_center", nil)
	}
	if v := dataCenterObj["data-center-object-meta-info"]; v != nil {

		dataCenterObjMetaInfoMap := v.(map[string]interface{})
		dataCenterObjMetaInfoState := make(map[string]interface{})

		if v := dataCenterObjMetaInfoMap["updated-on-data-center"]; v != nil {

			updatedOnDataCenterMap := v.(map[string]interface{})
			if v := updatedOnDataCenterMap["iso-8601"]; v != nil {
				dataCenterObjMetaInfoState["iso_8601"] = v
			}
			if v := updatedOnDataCenterMap["posix"]; v != nil {
				dataCenterObjMetaInfoState["posix"] = strconv.Itoa(int(v.(float64)))
			}
		}
		_ = d.Set("updated_on_data_center", dataCenterObjMetaInfoState)

	} else {
		_ = d.Set("updated_on_data_center", nil)
	}
	if v := dataCenterObj["deleted"]; v != nil {
		_ = d.Set("deleted", v)
	}

	if v := dataCenterObj["type-in-data-center"]; v != nil {
		_ = d.Set("type_in_data_center", v)
	}

	if v := dataCenterObj["additional-properties"]; v != nil {

		var listToReturn []map[string]interface{}
		propertiesList := v.([]interface{})

		if len(propertiesList) > 0 {
			for i := range propertiesList {
				innerMap := propertiesList[i].(map[string]interface{})
				mapToReturn := make(map[string]interface{})
				if v := innerMap["name"]; v != nil {
					mapToReturn["name"] = v
				}
				if v := innerMap["value"]; v != nil {
					mapToReturn["value"] = v
				}
				listToReturn = append(listToReturn, mapToReturn)
			}
		}
		_ = d.Set("additional_properties", listToReturn)
	} else {
		_ = d.Set("additional_properties", nil)
	}

	if dataCenterObj["tags"] != nil {
		tagsJson, ok := dataCenterObj["tags"].([]interface{})
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

	if dataCenterObj["groups"] != nil {
		groupsJson, ok := dataCenterObj["groups"].([]interface{})
		if ok {
			groupsIds := make([]string, 0)
			if len(groupsJson) > 0 {
				for _, tags := range groupsJson {
					tags := tags.(map[string]interface{})
					groupsIds = append(groupsIds, tags["name"].(string))
				}
			}
			_ = d.Set("groups", groupsIds)
		}
	} else {
		_ = d.Set("groups", nil)
	}

	if v := dataCenterObj["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := dataCenterObj["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := dataCenterObj["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := dataCenterObj["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil
}
