package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strconv"
	"time"
)

func resourceDataCenterObject() *schema.Resource {
	return &schema.Resource{
		Create: createManagementDataCenterObject,
		Read:   readManagementDataCenterObject,
		Update: updateManagementDataCenterObject,
		Delete: deleteManagementDataCenterObject,
		Schema: map[string]*schema.Schema{
			"data_center_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of the Data Center Server the object is in.",
			},
			"data_center_uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Unique identifier of the Data Center Server the object is in.",
			},
			"uri": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "URI of the object in the Data Center Server.",
			},
			"uid_in_data_center": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Unique identifier of the object in the Data Center Server.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Override default name on data-center.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"color": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Color of the object. Should be one of existing colors.",
				Default:     "black",
			},
			"comments": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comments string.",
			},
			"groups": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of group identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ignore_warnings": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Apply changes ignoring warnings.",
				Default:     false,
			},
			"ignore_errors": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.",
				Default:     false,
			},
			"name_in_data_center": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object Name in Data Center",
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
			"wait_until_sync_object": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "When set to true, the provider will wait until object is synced with the management server",
			},
		},
	}
}

func createManagementDataCenterObject(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("data_center_name"); ok {
		payload["data-center-name"] = v.(string)
	} else {
		if v, ok := d.GetOk("data_center_uid"); ok {
			payload["data-center-uid"] = v.(string)
		}
	}

	if v, ok := d.GetOk("uri"); ok {
		payload["uri"] = v.(string)
	} else {
		if v, ok := d.GetOk("uid_in_data_center"); ok {
			payload["uid-in-data-center"] = v.(string)
		}
	}

	if v, ok := d.GetOk("name"); ok {
		payload["name"] = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		payload["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		payload["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		payload["comments"] = v.(string)
	}

	if v, ok := d.GetOk("groups"); ok {
		payload["groups"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		payload["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		payload["ignore-errors"] = v.(bool)
	}

	AddDataCenterObjectRes, err := client.ApiCall("add-data-center-object", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !AddDataCenterObjectRes.Success {
		return fmt.Errorf(AddDataCenterObjectRes.ErrorMsg)
	}

	d.SetId(AddDataCenterObjectRes.GetData()["uid"].(string))

	if doWait, ok := d.GetOkExists("wait_until_sync_object"); ok {
		if doWait.(bool) {
			objUidInDataCenter := AddDataCenterObjectRes.GetData()["uid-in-data-center"].(string)
			dataCenterUid := AddDataCenterObjectRes.GetData()["data-center"].(map[string]interface{})["uid"].(string)
			showDataCenterContentPayload := make(map[string]interface{})
			showDataCenterContentPayload["data-center-uid"] = dataCenterUid
			showDataCenterContentPayload["uid-in-data-center"] = objUidInDataCenter
			maxRetry := 20
			retry := 1
			for retry < maxRetry {
				showDataCenterContentRes, err := client.ApiCallSimple("show-data-center-content", showDataCenterContentPayload)
				if err != nil || !showDataCenterContentRes.Success {
					// failed to check if data center object exist. stop check...
					break
				}
				showDataCenterContentData := showDataCenterContentRes.GetData()
				if len(showDataCenterContentData["objects"].([]interface{})) > 0 {
					obj := showDataCenterContentData["objects"].([]interface{})[0].(map[string]interface{})
					if obj["uid-in-data-center"] == objUidInDataCenter {
						break
					}
				}
				time.Sleep(time.Second * 10)
				retry++
			}
		}
	}

	return readManagementDataCenterObject(d, m)
}

func readManagementDataCenterObject(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid":           d.Id(),
		"details-level": "full",
	}

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

	if v := dataCenterObj["name-in-data-center"]; v != nil {
		_ = d.Set("name_in_data_center", v)
	}
	if v := dataCenterObj["uid-in-data-center"]; v != nil {
		_ = d.Set("uid_in_data_center", v)
	}
	if v := dataCenterObj["data-center"]; v != nil {

		dataCenterMap, ok := dataCenterObj["data-center"].(map[string]interface{})

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
func updateManagementDataCenterObject(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	dataCenter := make(map[string]interface{})

	if name, ok := d.GetOk("name"); ok {
		dataCenter["name"] = name
	} else {
		dataCenter["uid"] = d.Id()
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			dataCenter["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			dataCenter["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		dataCenter["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		dataCenter["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		dataCenter["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		dataCenter["ignore-errors"] = v.(bool)
	}

	log.Println("Update Data Center - Map = ", dataCenter)

	updateDataCenterRes, err := client.ApiCall("set-data-center-object", dataCenter, client.GetSessionID(), true, false)
	if err != nil || !updateDataCenterRes.Success {
		if updateDataCenterRes.ErrorMsg != "" {
			return fmt.Errorf(updateDataCenterRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementDataCenterObject(d, m)
}

func deleteManagementDataCenterObject(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	dataCenterPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete Data Center")
	deleteLsmClusterRes, err := client.ApiCall("delete-data-center-object", dataCenterPayload, client.GetSessionID(), true, false)
	if err != nil || !deleteLsmClusterRes.Success {
		if deleteLsmClusterRes.ErrorMsg != "" {
			return fmt.Errorf(deleteLsmClusterRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
