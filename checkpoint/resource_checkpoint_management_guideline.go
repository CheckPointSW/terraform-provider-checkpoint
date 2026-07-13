package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"

	"strconv"
)

func resourceManagementGuideline() *schema.Resource {
	return &schema.Resource{
		Create: createManagementGuideline,
		Read:   readManagementGuideline,
		Update: updateManagementGuideline,
		Delete: deleteManagementGuideline,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"access_layers": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "The access-layers (one or more) that will be attached to the guideline, identified by name or UID.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_layer": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Access-layer attached to guideline identified by the name or UID.if Access-Layer is in the global domain due to Global Assignment Local domain Package is required.",
						},
						"policy_package": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Policy package context for the access-layer attached to guideline identified by the name or UID.Package will be ignored if the access-layer is local.",
						},
					},
				},
			},
			"guideline_groups": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "The groups that will be part of the guideline (guideline should have between 2-12 segments, including internet-segment and other-segment). It is recommended to select groups that best represent segments of the network.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Network group name.",
						},
						"uid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Network group unique identifier.",
						},
						"position": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Required:    true,
							Description: "Position in the rulebase.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"top": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "N/A",
									},
									"above": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "N/A",
									},
									"below": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "N/A",
									},
									"bottom": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "N/A",
									},
								},
							},
						},
						"members": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Network group members identified by name.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"cell_actions_override": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Cells that their action will override the default actions of the guideline.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"from": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The segment identifier (name or UID) of the cell in the 'from' axis. The field is mandatory only if \"from-type\" is \"network group\".",
						},
						"from_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The type of the segment in the 'from' axis.",
							Default:     "Network Group",
						},
						"to": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The segment identifier (name or UID) of the cell in the 'to' axis. The field is mandatory only if \"to-type\" is \"network group\".",
						},
						"to_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The type of the segment in the 'to' axis.",
							Default:     "Network Group",
						},
						"action": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The action to be applied to the cell. The field is mandatory at add command.",
						},
						"allowed_services": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Services (identified by name or UID) that are allowed in the cell. Relevant only if the action in the cell is 'All traffic is not allowed'. To remove allowed-services call update with the same \"All traffic is not allowed\" action, or remove the cell-action-override.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"indexing_status": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Task-id map for the indexing tasks of the guideline.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_layer_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the access-layer that is being indexed.",
						},
						"indexing_message": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Message which offers more details on The indexing task.",
						},
						"indexing_task": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the task that is indexing the access-layer. Relevant only if the task is in progress.",
						},
						"last_update_time": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Last time the indexing status was updated.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"iso_8601": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Date and time represented in international ISO 8601 format.",
									},
									"posix": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.",
									},
								},
							},
						},
						"policy_package_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the policy Package that is being indexed.(only used if the layer is global).",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the indexing task.",
						},
					},
				},
			},
			"dereference_group_members": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates whether to dereference \"members\" field by details level for every object in reply.",
				Default:     false,
			},
			"show_membership": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates whether to calculate and show \"groups\" field for every object in reply.",
				Default:     false,
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
		},
	}
}

func createManagementGuideline(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	guideline := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		guideline["name"] = v.(string)
	}

	if v, ok := d.GetOk("access_layers"); ok {

		accessLayersList := v.([]interface{})

		if len(accessLayersList) > 0 {

			var accessLayersPayload []map[string]interface{}

			for i := range accessLayersList {

				Payload := make(map[string]interface{})

				if v, ok := d.GetOk("access_layers." + strconv.Itoa(i) + ".access_layer"); ok {
					Payload["access-layer"] = v.(string)
				}
				if v, ok := d.GetOk("access_layers." + strconv.Itoa(i) + ".policy_package"); ok {
					Payload["policy-package"] = v.(string)
				}
				accessLayersPayload = append(accessLayersPayload, Payload)
			}
			guideline["access-layers"] = accessLayersPayload
		}
	}

	if v, ok := d.GetOk("guideline_groups"); ok {

		guidelineGroupsList := v.([]interface{})

		if len(guidelineGroupsList) > 0 {

			var guidelineGroupsPayload []map[string]interface{}

			for i := range guidelineGroupsList {

				Payload := make(map[string]interface{})

				if v, ok := d.GetOk("guideline_groups." + strconv.Itoa(i) + ".name"); ok {
					Payload["name"] = v.(string)
				}
				if v, ok := d.GetOk("guideline_groups." + strconv.Itoa(i) + ".uid"); ok {
					Payload["uid"] = v.(string)
				}
				if _, ok := d.GetOk("guideline_groups." + strconv.Itoa(i) + ".position"); ok {
					positionPrefix := "guideline_groups." + strconv.Itoa(i) + ".position.0."
					if v, ok := d.GetOk(positionPrefix + "top"); ok {
						if v.(string) == "top" {
							Payload["position"] = "top"
						} else {
							Payload["position"] = map[string]interface{}{"top": v.(string)}
						}
					}
					if v, ok := d.GetOk(positionPrefix + "above"); ok {
						Payload["position"] = map[string]interface{}{"above": v.(string)}
					}
					if v, ok := d.GetOk(positionPrefix + "below"); ok {
						Payload["position"] = map[string]interface{}{"below": v.(string)}
					}
					if v, ok := d.GetOk(positionPrefix + "bottom"); ok {
						if v.(string) == "bottom" {
							Payload["position"] = "bottom"
						} else {
							Payload["position"] = map[string]interface{}{"bottom": v.(string)}
						}
					}
				}
				guidelineGroupsPayload = append(guidelineGroupsPayload, Payload)
			}
			guideline["guideline-groups"] = guidelineGroupsPayload
		}
	}

	if v, ok := d.GetOk("cell_actions_override"); ok {

		cellActionsOverrideList := v.([]interface{})

		if len(cellActionsOverrideList) > 0 {

			var cellActionsOverridePayload []map[string]interface{}

			for i := range cellActionsOverrideList {

				Payload := make(map[string]interface{})

				if v, ok := d.GetOk("cell_actions_override." + strconv.Itoa(i) + ".from"); ok {
					Payload["from"] = v.(string)
				}
				if v, ok := d.GetOk("cell_actions_override." + strconv.Itoa(i) + ".from_type"); ok {
					Payload["from-type"] = v.(string)
				}
				if v, ok := d.GetOk("cell_actions_override." + strconv.Itoa(i) + ".to"); ok {
					Payload["to"] = v.(string)
				}
				if v, ok := d.GetOk("cell_actions_override." + strconv.Itoa(i) + ".to_type"); ok {
					Payload["to-type"] = v.(string)
				}
				if v, ok := d.GetOk("cell_actions_override." + strconv.Itoa(i) + ".action"); ok {
					Payload["action"] = v.(string)
				}
				if v, ok := d.GetOk("cell_actions_override." + strconv.Itoa(i) + ".allowed_services"); ok {
					Payload["allowed-services"] = v
				}
				cellActionsOverridePayload = append(cellActionsOverridePayload, Payload)
			}
			guideline["cell-actions-override"] = cellActionsOverridePayload
		}
	}

	if v, ok := d.GetOkExists("dereference_group_members"); ok {
		guideline["dereference-group-members"] = v.(bool)
	}

	if v, ok := d.GetOkExists("show_membership"); ok {
		guideline["show-membership"] = v.(bool)
	}

	if v, ok := d.GetOk("color"); ok {
		guideline["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		guideline["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		guideline["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		guideline["ignore-errors"] = v.(bool)
	}

	log.Println("Create Guideline - Map = ", guideline)

	addGuidelineRes, err := client.ApiCall("add-guideline", guideline, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !addGuidelineRes.Success {
		if addGuidelineRes.ErrorMsg != "" {
			return fmt.Errorf(addGuidelineRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addGuidelineRes.GetData()["uid"].(string))

	return readManagementGuideline(d, m)
}

func readManagementGuideline(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showGuidelineRes, err := client.ApiCall("show-guideline", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showGuidelineRes.Success {
		if objectNotFound(showGuidelineRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showGuidelineRes.ErrorMsg)
	}

	guideline := showGuidelineRes.GetData()

	log.Println("Read Guideline - Show JSON = ", guideline)

	if v := guideline["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if guideline["access-layers"] != nil {

		accessLayersList := guideline["access-layers"].([]interface{})

		if len(accessLayersList) > 0 {

			var accessLayersListToReturn []map[string]interface{}

			for i := range accessLayersList {

				accessLayersMap := accessLayersList[i].(map[string]interface{})

				accessLayersMapToAdd := make(map[string]interface{})

				if v, _ := accessLayersMap["access-layer"]; v != nil {
					accessLayersMapToAdd["access_layer"] = v.(map[string]interface{})["name"]
				}
				if v, _ := accessLayersMap["policy-package"]; v != nil {
					accessLayersMapToAdd["policy_package"] = v.(map[string]interface{})["name"]
				}
				accessLayersListToReturn = append(accessLayersListToReturn, accessLayersMapToAdd)
			}

			_ = d.Set("access_layers", accessLayersListToReturn)
		} else {
			_ = d.Set("access_layers", accessLayersList)
		}
	} else {
		_ = d.Set("access_layers", nil)
	}

	if guideline["guideline-groups"] != nil {

		guidelineGroupsList := guideline["guideline-groups"].([]interface{})

		if len(guidelineGroupsList) > 0 {

			var guidelineGroupsListToReturn []map[string]interface{}

			for i := range guidelineGroupsList {

				guidelineGroupsMap := guidelineGroupsList[i].(map[string]interface{})

				guidelineGroupsMapToAdd := make(map[string]interface{})

				if v, _ := guidelineGroupsMap["guideline-group"]; v != nil {
					if guidelineGroupList, ok := v.([]interface{}); ok && len(guidelineGroupList) > 0 {
						guidelineGroupsMapToAdd["name"] = guidelineGroupList[0].(map[string]interface{})["name"]
					}
				}
				if v, _ := guidelineGroupsMap["position"]; v != nil {
					guidelineGroupsMapToAdd["position"] = v
				}
				if v, _ := guidelineGroupsMap["members"]; v != nil {
					membersJson, ok := v.([]interface{})
					if ok {
						membersIds := make([]string, 0)
						if len(membersJson) > 0 {
							for _, member := range membersJson {
								member := member.(map[string]interface{})
								membersIds = append(membersIds, member["name"].(string))
							}
						}
						guidelineGroupsMapToAdd["members"] = membersIds
					}
				}
				guidelineGroupsListToReturn = append(guidelineGroupsListToReturn, guidelineGroupsMapToAdd)
			}

			_ = d.Set("guideline_groups", guidelineGroupsListToReturn)
		} else {
			_ = d.Set("guideline_groups", guidelineGroupsList)
		}
	} else {
		_ = d.Set("guideline_groups", nil)
	}

	if guideline["cell-actions-override"] != nil {

		cellActionsOverrideList := guideline["cell-actions-override"].([]interface{})

		if len(cellActionsOverrideList) > 0 {

			var cellActionsOverrideListToReturn []map[string]interface{}

			for i := range cellActionsOverrideList {

				cellActionsOverrideMap := cellActionsOverrideList[i].(map[string]interface{})

				cellActionsOverrideMapToAdd := make(map[string]interface{})

				if v, _ := cellActionsOverrideMap["from"]; v != nil {
					cellActionsOverrideMapToAdd["from"] = v
				}
				if v, _ := cellActionsOverrideMap["from-type"]; v != nil {
					cellActionsOverrideMapToAdd["from_type"] = v
				}
				if v, _ := cellActionsOverrideMap["to"]; v != nil {
					cellActionsOverrideMapToAdd["to"] = v
				}
				if v, _ := cellActionsOverrideMap["to-type"]; v != nil {
					cellActionsOverrideMapToAdd["to_type"] = v
				}
				if v, _ := cellActionsOverrideMap["action"]; v != nil {
					cellActionsOverrideMapToAdd["action"] = v
				}
				if v, _ := cellActionsOverrideMap["allowed-services"]; v != nil {
					cellActionsOverrideMapToAdd["allowed_services"] = v
				}
				cellActionsOverrideListToReturn = append(cellActionsOverrideListToReturn, cellActionsOverrideMapToAdd)
			}

			_ = d.Set("cell_actions_override", cellActionsOverrideListToReturn)
		} else {
			_ = d.Set("cell_actions_override", cellActionsOverrideList)
		}
	} else {
		_ = d.Set("cell_actions_override", nil)
	}

	if v := guideline["dereference-group-members"]; v != nil {
		_ = d.Set("dereference_group_members", v)
	}

	if v := guideline["show-membership"]; v != nil {
		_ = d.Set("show_membership", v)
	}

	if guideline["indexing-status"] != nil {

		indexingStatusList := guideline["indexing-status"].([]interface{})

		if len(indexingStatusList) > 0 {

			var indexingStatusListToReturn []map[string]interface{}

			for i := range indexingStatusList {

				indexingStatusMap := indexingStatusList[i].(map[string]interface{})

				indexingStatusMapToAdd := make(map[string]interface{})

				if v := indexingStatusMap["access-layer-id"]; v != nil {
					indexingStatusMapToAdd["access_layer_id"] = v
				}
				if v := indexingStatusMap["indexing-message"]; v != nil {
					indexingStatusMapToAdd["indexing_message"] = v
				}
				if v := indexingStatusMap["indexing-task"]; v != nil {
					indexingStatusMapToAdd["indexing_task"] = v
				}
				if v := indexingStatusMap["last-update-time"]; v != nil {

					lastUpdateTimeMap := v.(map[string]interface{})

					lastUpdateTimeMapToReturn := make(map[string]interface{})

					if v := lastUpdateTimeMap["iso-8601"]; v != nil {
						lastUpdateTimeMapToReturn["iso_8601"] = v
					}
					if v := lastUpdateTimeMap["posix"]; v != nil {
						lastUpdateTimeMapToReturn["posix"] = v
					}

					indexingStatusMapToAdd["last_update_time"] = []interface{}{lastUpdateTimeMapToReturn}
				}

				if v := indexingStatusMap["policy-package-id"]; v != nil {
					indexingStatusMapToAdd["policy_package_id"] = v
				}
				if v := indexingStatusMap["status"]; v != nil {
					indexingStatusMapToAdd["status"] = v
				}

				indexingStatusListToReturn = append(indexingStatusListToReturn, indexingStatusMapToAdd)
			}

			_ = d.Set("indexing_status", indexingStatusListToReturn)
		}
	} else {
		_ = d.Set("indexing_status", nil)
	}

	if v := guideline["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := guideline["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := guideline["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := guideline["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementGuideline(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	guideline := make(map[string]interface{})

	guideline["uid"] = d.Id()

	if ok := d.HasChange("name"); ok {
		if v, ok := d.GetOk("name"); ok {
			guideline["new-name"] = v.(string)
		}
	}

	if d.HasChange("access_layers") {

		if v, ok := d.GetOk("access_layers"); ok {

			accessLayersList := v.([]interface{})

			var accessLayersPayload []map[string]interface{}

			for i := range accessLayersList {

				Payload := make(map[string]interface{})

				if v, ok := d.GetOk("access_layers." + strconv.Itoa(i) + ".access_layer"); ok {
					Payload["access-layer"] = v
				}
				if v, ok := d.GetOk("access_layers." + strconv.Itoa(i) + ".policy_package"); ok {
					Payload["policy-package"] = v
				}
				if v, ok := d.GetOk("access_layers." + strconv.Itoa(i) + ".domains_to_process"); ok {
					Payload["domains-to-process"] = v
				}
				accessLayersPayload = append(accessLayersPayload, Payload)
			}
			guideline["access-layers"] = accessLayersPayload
		}
	}

	if d.HasChange("guideline_groups") {

		if v, ok := d.GetOk("guideline_groups"); ok {

			guidelineGroupsList := v.([]interface{})

			var guidelineGroupsPayload []map[string]interface{}

			for i := range guidelineGroupsList {

				Payload := make(map[string]interface{})

				if v, ok := d.GetOk("guideline_groups." + strconv.Itoa(i) + ".name"); ok {
					Payload["name"] = v
				}
				if _, ok := d.GetOk("guideline_groups." + strconv.Itoa(i) + ".position"); ok {
					positionPrefix := "guideline_groups." + strconv.Itoa(i) + ".position.0."
					if v, ok := d.GetOk(positionPrefix + "top"); ok {
						if v.(string) == "top" {
							Payload["position"] = "top"
						} else {
							Payload["position"] = map[string]interface{}{"top": v.(string)}
						}
					}
					if v, ok := d.GetOk(positionPrefix + "above"); ok {
						Payload["position"] = map[string]interface{}{"above": v.(string)}
					}
					if v, ok := d.GetOk(positionPrefix + "below"); ok {
						Payload["position"] = map[string]interface{}{"below": v.(string)}
					}
					if v, ok := d.GetOk(positionPrefix + "bottom"); ok {
						if v.(string) == "bottom" {
							Payload["position"] = "bottom"
						} else {
							Payload["position"] = map[string]interface{}{"bottom": v.(string)}
						}
					}
				}
				guidelineGroupsPayload = append(guidelineGroupsPayload, Payload)
			}
			guideline["guideline-groups"] = guidelineGroupsPayload
		}
	}

	if d.HasChange("cell_actions_override") {

		if v, ok := d.GetOk("cell_actions_override"); ok {

			cellActionsOverrideList := v.([]interface{})

			var cellActionsOverridePayload []map[string]interface{}

			for i := range cellActionsOverrideList {

				Payload := make(map[string]interface{})

				if v, ok := d.GetOk("cell_actions_override." + strconv.Itoa(i) + ".from"); ok {
					Payload["from"] = v
				}
				if v, ok := d.GetOk("cell_actions_override." + strconv.Itoa(i) + ".from_type"); ok {
					Payload["from-type"] = v
				}
				if v, ok := d.GetOk("cell_actions_override." + strconv.Itoa(i) + ".to"); ok {
					Payload["to"] = v
				}
				if v, ok := d.GetOk("cell_actions_override." + strconv.Itoa(i) + ".to_type"); ok {
					Payload["to-type"] = v
				}
				if v, ok := d.GetOk("cell_actions_override." + strconv.Itoa(i) + ".action"); ok {
					Payload["action"] = v
				}
				if v, ok := d.GetOk("cell_actions_override." + strconv.Itoa(i) + ".allowed_services"); ok {
					Payload["allowed-services"] = v
				}
				cellActionsOverridePayload = append(cellActionsOverridePayload, Payload)
			}
			guideline["cell-actions-override"] = cellActionsOverridePayload
		}
	}

	if v, ok := d.GetOkExists("dereference_group_members"); ok {
		guideline["dereference-group-members"] = v.(bool)
	}

	if v, ok := d.GetOkExists("show_membership"); ok {
		guideline["show-membership"] = v.(bool)
	}

	if ok := d.HasChange("color"); ok {
		guideline["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		guideline["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		guideline["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		guideline["ignore-errors"] = v.(bool)
	}

	log.Println("Update Guideline - Map = ", guideline)

	updateGuidelineRes, err := client.ApiCall("set-guideline", guideline, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !updateGuidelineRes.Success {
		if updateGuidelineRes.ErrorMsg != "" {
			return fmt.Errorf(updateGuidelineRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementGuideline(d, m)
}

func deleteManagementGuideline(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	guidelinePayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete Guideline")

	deleteGuidelineRes, err := client.ApiCall("delete-guideline", guidelinePayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteGuidelineRes.Success {
		if deleteGuidelineRes.ErrorMsg != "" {
			return fmt.Errorf(deleteGuidelineRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
