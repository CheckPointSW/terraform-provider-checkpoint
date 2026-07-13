package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func dataSourceManagementGuideline() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementGuidelineRead,
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
			"show_indexing_status": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Control whether to show the indexing status of the guideline.",
			},
			"indexing_status_layer": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Relevant only when show-indexing-status is true. The access-layer to show the indexing status of (identified by unique id or 'any' for all attached access-layers).",
			},
			"dereference_group_members": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates whether to dereference \"members\" field by details level for every object in reply.",
			},
			"show_membership": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates whether to calculate and show \"groups\" field for every object in reply.",
			},
			"access_layers": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The access-layers objects attached to the guideline with their policy-package context.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_layer": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The access-layer object attached to the guideline identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.",
						},
						"policy_package": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The policy-package object context for the access-layer (only for global access-layers) identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.",
						},
					},
				},
			},
			"cell_actions_override": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "All the cells that the user changed the default action in.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"from": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Unique identifier of the segment of the cell in the 'from' axis.",
						},
						"from_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the segment in the 'from' axis.",
						},
						"to": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Unique identifier of the segment of the cell in the 'to' axis.",
						},
						"to_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the segment in the 'to' axis.",
						},
						"action": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The action selected for the cell.",
						},
					},
				},
			},
			"default_action": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The default action for guideline cells with two different groups.",
			},
			"default_self_action": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The default action for guideline cells with the same group in both axis.",
			},
			"guideline_groups": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The segments displayed in the guideline matrix in at least one of the axes (from or to).",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"guideline_group": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "The network-group object identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"members": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "N/A",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"position": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The position of the guideline group in the axis.",
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
			"icon": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object icon.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of tag objects identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceManagementGuidelineRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{}

	if v, ok := d.GetOk("name"); ok {
		payload["name"] = v.(string)
	} else if v, ok := d.GetOk("uid"); ok {
		payload["uid"] = v.(string)
	} else {
		return fmt.Errorf("Either name or uid must be specified")
	}

	if v, ok := d.GetOkExists("show_indexing_status"); ok {
		payload["show-indexing-status"] = v.(bool)
	}

	if v, ok := d.GetOk("indexing_status_layer"); ok {
		payload["indexing-status-layer"] = v.(string)
	}

	if v, ok := d.GetOkExists("dereference_group_members"); ok {
		payload["dereference-group-members"] = v.(bool)
	}

	if v, ok := d.GetOkExists("show_membership"); ok {
		payload["show-membership"] = v.(bool)
	}

	showGuidelineRes, err := client.ApiCall("show-guideline", payload, client.GetSessionID(), true, false)
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

	if v := guideline["uid"]; v != nil {
		d.SetId(v.(string))
	}

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

				if v := accessLayersMap["access-layer"]; v != nil {
					accessLayersMapToAdd["access_layer"] = v.(map[string]interface{})["name"]
				}

				if v := accessLayersMap["policy-package"]; v != nil {
					accessLayersMapToAdd["policy_package"] = v.(map[string]interface{})["name"]
				}

				accessLayersListToReturn = append(accessLayersListToReturn, accessLayersMapToAdd)
			}

			_ = d.Set("access_layers", accessLayersListToReturn)
		}
	} else {
		_ = d.Set("access_layers", nil)
	}

	if guideline["cell-actions-override"] != nil {

		cellActionsOverrideList := guideline["cell-actions-override"].([]interface{})

		if len(cellActionsOverrideList) > 0 {

			var cellActionsOverrideListToReturn []map[string]interface{}

			for i := range cellActionsOverrideList {

				cellActionsOverrideMap := cellActionsOverrideList[i].(map[string]interface{})

				cellActionsOverrideMapToAdd := make(map[string]interface{})

				if v := cellActionsOverrideMap["from"]; v != nil {
					cellActionsOverrideMapToAdd["from"] = v
				}
				if v := cellActionsOverrideMap["from-type"]; v != nil {
					cellActionsOverrideMapToAdd["from_type"] = v
				}
				if v := cellActionsOverrideMap["to"]; v != nil {
					cellActionsOverrideMapToAdd["to"] = v
				}
				if v := cellActionsOverrideMap["to-type"]; v != nil {
					cellActionsOverrideMapToAdd["to_type"] = v
				}
				if v := cellActionsOverrideMap["action"]; v != nil {
					cellActionsOverrideMapToAdd["action"] = v
				}

				cellActionsOverrideListToReturn = append(cellActionsOverrideListToReturn, cellActionsOverrideMapToAdd)
			}

			_ = d.Set("cell_actions_override", cellActionsOverrideListToReturn)
		}
	} else {
		_ = d.Set("cell_actions_override", nil)
	}

	if v := guideline["default-action"]; v != nil {
		_ = d.Set("default_action", v)
	}

	if v := guideline["default-self-action"]; v != nil {
		_ = d.Set("default_self_action", v)
	}

	if guideline["guideline-groups"] != nil {

		guidelineGroupsList := guideline["guideline-groups"].([]interface{})

		if len(guidelineGroupsList) > 0 {

			var guidelineGroupsListToReturn []map[string]interface{}

			for i := range guidelineGroupsList {

				guidelineGroupsMap := guidelineGroupsList[i].(map[string]interface{})

				guidelineGroupsMapToAdd := make(map[string]interface{})

				if v := guidelineGroupsMap["guideline-group"]; v != nil {
					guidelineGroupList, ok := v.([]interface{})
					if ok {
						guidelineGroupIds := make([]string, 0)
						if len(guidelineGroupList) > 0 {
							for _, item := range guidelineGroupList {
								item := item.(map[string]interface{})
								guidelineGroupIds = append(guidelineGroupIds, item["name"].(string))
							}
						}
						guidelineGroupsMapToAdd["guideline_group"] = guidelineGroupIds
					}
				}

				if v := guidelineGroupsMap["members"]; v != nil {
					membersList, ok := v.([]interface{})
					if ok {
						membersIds := make([]string, 0)
						if len(membersList) > 0 {
							for _, item := range membersList {
								item := item.(map[string]interface{})
								membersIds = append(membersIds, item["name"].(string))
							}
						}
						guidelineGroupsMapToAdd["members"] = membersIds
					}
				}

				if v := guidelineGroupsMap["position"]; v != nil {
					guidelineGroupsMapToAdd["position"] = v
				}

				guidelineGroupsListToReturn = append(guidelineGroupsListToReturn, guidelineGroupsMapToAdd)
			}

			_ = d.Set("guideline_groups", guidelineGroupsListToReturn)
		}
	} else {
		_ = d.Set("guideline_groups", nil)
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

	if v := guideline["icon"]; v != nil {
		_ = d.Set("icon", v)
	}

	if guideline["tags"] != nil {
		tagsJson, ok := guideline["tags"].([]interface{})
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

	if v := guideline["uid"]; v != nil {
		_ = d.Set("uid", v)
	}

	return nil

}
