package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strconv"
)

func dataSourceManagementShowUpdatableObjectsRepositoryContent() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementShowUpdatableObjectsRepositoryContentRead,
		Schema: map[string]*schema.Schema{
			"uid_in_updatable_objects_repository": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The object's unique identifier in the Updatable Objects repository.",
			},
			"filter": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Return results matching the specified filter.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"text": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Return results containing the specified text value.",
						},
						"uri": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Return results under the specified uri value.",
						},
						"parent_uid_in_updatable_objects_repository": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Return results under the specified Updatable Object.",
						},
					},
				},
			},
			"limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The maximal number of returned results.",
			},
			"offset": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Number of the results to initially skip.",
			},
			"order": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Sorts the results by search criteria. Automatically sorts the results by Name, in the ascending order.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"asc": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Sorts results by the given field in ascending order.",
						},
						"desc": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Sorts results by the given field in descending order.",
						},
					},
				},
			},
			"from": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "From which element number the query was done.",
			},
			"to": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "To which element number the query was done.",
			},
			"total": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Total number of elements returned by the query.",
			},
			"objects": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Objects list",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
							Type:        schema.TypeMap,
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
						"updatable_object": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "The imported management object (if exists).",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Object name.",
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
									"domain": {
										Type:        schema.TypeMap,
										Computed:    true,
										Description: "Information about the domain that holds the Object.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Object name.",
												},
												"uid": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Object unique identifier.",
												},
												"domain_type": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Domain type.",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceManagementShowUpdatableObjectsRepositoryContentRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := make(map[string]interface{})

	if v, ok := d.GetOk("uid_in_updatable_objects_repository"); ok {
		payload["uid-in-updatable-objects-repository"] = v.(string)
	}

	if _, ok := d.GetOk("filter"); ok {
		filter := make(map[string]interface{})

		if v, ok := d.GetOk("filter.text"); ok {
			filter["text"] = v.(string)
		}

		if v, ok := d.GetOk("filter.uri"); ok {
			filter["uri"] = v.(string)
		}

		if v, ok := d.GetOk("filter.parent_uid_in_updatable_objects_repository"); ok {
			filter["parent-uid-in-updatable-objects-repository"] = v.(string)
		}

		payload["filter"] = filter
	}

	if v, ok := d.GetOk("limit"); ok {
		payload["limit"] = v.(int)
	}

	if v, ok := d.GetOk("offset"); ok {
		payload["offset"] = v.(int)
	}

	if v, ok := d.GetOk("order"); ok {

		orderList := v.([]interface{})

		if len(orderList) > 0 {
			var orderPayload []map[string]interface{}

			for i := range orderList {
				payload := make(map[string]interface{})

				if v, ok := d.GetOk("order." + strconv.Itoa(i) + ".asc"); ok {
					payload["ASC"] = v.(string)
				}

				if v, ok := d.GetOk("order." + strconv.Itoa(i) + ".desc"); ok {
					payload["DESC"] = v.(string)
				}

				orderPayload = append(orderPayload, payload)
			}

			payload["order"] = orderPayload
		}
	}

	showUpdatableObjectsRepositoryContentRes, err := client.ApiCall("show-updatable-objects-repository-content", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showUpdatableObjectsRepositoryContentRes.Success {
		return fmt.Errorf(showUpdatableObjectsRepositoryContentRes.ErrorMsg)
	}

	updatableObjectsRepositoryContentResData := showUpdatableObjectsRepositoryContentRes.GetData()

	log.Println("show-updatable-objects-repository-content JSON = ", updatableObjectsRepositoryContentResData)

	d.SetId("show-updatable-objects-repository-content-" + acctest.RandString(10))

	if v := updatableObjectsRepositoryContentResData["from"]; v != nil {
		_ = d.Set("from", v)
	}

	if v := updatableObjectsRepositoryContentResData["to"]; v != nil {
		_ = d.Set("to", v)
	}

	if v := updatableObjectsRepositoryContentResData["total"]; v != nil {
		_ = d.Set("total", v)
	}

	if v := updatableObjectsRepositoryContentResData["objects"]; v != nil {
		objectsList := v.([]interface{})
		if len(objectsList) > 0 {
			var objectsListState []map[string]interface{}
			for i := range objectsList {
				objectJson := objectsList[i].(map[string]interface{})
				objectState := make(map[string]interface{})

				if v := objectJson["name-in-updatable-objects-repository"]; v != nil {
					objectState["name_in_updatable_objects_repository"] = v
				}

				if v := objectJson["uid-in-updatable-objects-repository"]; v != nil {
					objectState["uid_in_updatable_objects_repository"] = v
				}

				if v := objectJson["additional-properties"]; v != nil {
					additionalPropertiesJson := v.(map[string]interface{})
					additionalPropertiesState := make(map[string]interface{})
					if v := additionalPropertiesJson["description"]; v != nil {
						additionalPropertiesState["description"] = v
					}
					if v := additionalPropertiesJson["info-text"]; v != nil {
						additionalPropertiesState["info_text"] = v
					}
					if v := additionalPropertiesJson["info-url"]; v != nil {
						additionalPropertiesState["info_url"] = v
					}
					if v := additionalPropertiesJson["uri"]; v != nil {
						additionalPropertiesState["uri"] = v
					}
					objectState["additional_properties"] = additionalPropertiesState
				}

				if v := objectJson["updatable-object"]; v != nil {
					updatableObjectJson := v.(map[string]interface{})
					updatableObjectState := make(map[string]interface{})

					if v := updatableObjectJson["name"]; v != nil {
						updatableObjectState["name"] = v
					}

					if v := updatableObjectJson["uid"]; v != nil {
						updatableObjectState["uid"] = v
					}

					if v := updatableObjectJson["type"]; v != nil {
						updatableObjectState["type"] = v
					}

					if v := updatableObjectJson["domain"]; v != nil {
						domainJson := v.(map[string]interface{})
						domainState := make(map[string]interface{})

						if v := domainJson["name"]; v != nil {
							domainState["name"] = v
						}

						if v := domainJson["uid"]; v != nil {
							domainState["uid"] = v
						}

						if v := domainJson["domain-type"]; v != nil {
							domainState["domain_type"] = v
						}
						updatableObjectState["domain"] = domainState
					}
					objectState["updatable_object"] = updatableObjectState
				}

				objectsListState = append(objectsListState, objectState)
			}
			_ = d.Set("objects", objectsListState)
		} else {
			_ = d.Set("objects", objectsList)
		}
	} else {
		_ = d.Set("objects", nil)
	}

	return nil
}
