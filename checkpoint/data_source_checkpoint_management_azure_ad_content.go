package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strconv"
)

func dataSourceManagementAzureAdContent() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAzureAdContentRead,
		Schema: map[string]*schema.Schema{
			"azure_ad_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of the Azure AD Server where to search for objects.",
			},
			"azure_ad_uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Unique identifier of the Azure AD Server where to search for objects.",
			},
			"limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     50,
				Description: "The maximal number of returned results.",
			},
			"offset": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
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
			"uid_in_azure_ad": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Return result matching the unique identifier of the object on the Azure AD Server.",
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
							Description: "Return results under the specified Data Center Object (identified by URI).",
						},
						"parent_uid_in_data_center": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Return results under the specified Data Center Object (identified by UID).",
						},
					},
				},
			},
			"from": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "From which element number the query was done.",
			},
			"objects": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Remote objects views.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name_in_azure_ad": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object name in the Azure AD.",
						},
						"uid_in_azure_ad": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Unique identifier of the object in the Azure AD.",
						},
						"azure_ad_object": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The imported management object (if exists). Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object management name.",
						},
						"type_in_azure_ad": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object type in Azure AD.",
						},
						"additional_properties": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Additional properties on the object.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"value": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
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
		},
	}
}

func dataSourceAzureAdContentRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	payload := make(map[string]interface{})

	if v, ok := d.GetOk("azure_ad_name"); ok {
		payload["azure-ad-name"] = v.(string)
	}

	if v, ok := d.GetOk("azure_ad_uid"); ok {
		payload["azure-ad-uid"] = v.(string)
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
					payload["asc"] = v.(string)
				}
				if v, ok := d.GetOk("order." + strconv.Itoa(i) + ".desc"); ok {
					payload["desc"] = v.(string)
				}

				orderPayload = append(orderPayload, payload)
			}

			payload["order"] = orderPayload
		}
	}

	if v, ok := d.GetOk("uid_in_azure_ad"); ok {
		payload["uid-in-azure-ad"] = v.(string)
	}

	if _, ok := d.GetOk("filter"); ok {
		res := make(map[string]interface{})

		if v, ok := d.GetOk("filter.text"); ok {
			res["text"] = v.(string)
		}
		if v, ok := d.GetOk("filter.uri"); ok {
			res["uri"] = v.(string)
		}
		if v, ok := d.GetOk("filter.parent_uid_in_data_center"); ok {
			res["parent-uid-in-data-center"] = v.(string)
		}
		payload["filter"] = res
	}

	showAzureAdContentRes, err := client.ApiCall("show-azure-ad-content", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showAzureAdContentRes.Success {
		return fmt.Errorf(showAzureAdContentRes.ErrorMsg)
	}

	azureAdContent := showAzureAdContentRes.GetData()

	log.Println("Read Azure Ad Content - Show JSON = ", azureAdContent)

	if v := azureAdContent["from"]; v != nil {
		_ = d.Set("from", v)
	}

	if azureAdContent["objects"] != nil {
		objectsList := azureAdContent["objects"].([]interface{})

		if len(objectsList) > 0 {
			var objectsListToReturn []map[string]interface{}

			for i := range objectsList {
				objectsMap := objectsList[i].(map[string]interface{})

				objectsMapToAdd := make(map[string]interface{})

				if v, _ := objectsMap["name-in-azure-ad"]; v != nil {
					objectsMapToAdd["name_in_azure_ad"] = v
				}
				if v, _ := objectsMap["uid-in-azure-ad"]; v != nil {
					objectsMapToAdd["uid_in_azure_ad"] = v
				}

				if objectsMap["azure-ad-object"] != nil {
					azureAdObjectMap := objectsMap["azure-ad-object"].(map[string]interface{})

					objectsMapToAdd["azure_ad_object"] = azureAdObjectMap["name"]
				}

				if v, _ := objectsMap["name"]; v != nil {
					objectsMapToAdd["name"] = v
				}
				if v, _ := objectsMap["type-in-azure-ad"]; v != nil {
					objectsMapToAdd["type_in_azure_ad"] = v
				}
				if objectsMap["additional-properties"] != nil {
					additionalPropertiesList := objectsMap["additional-properties"].([]interface{})

					if len(additionalPropertiesList) > 0 {
						var additionalPropertiesListToReturn []map[string]interface{}

						for i := range additionalPropertiesList {
							additionalPropertiesMap := additionalPropertiesList[i].(map[string]interface{})

							additionalPropertiesMapToAdd := make(map[string]interface{})

							if v, _ := additionalPropertiesMap["name"]; v != nil {
								additionalPropertiesMapToAdd["name"] = v
							}
							if v, _ := additionalPropertiesMap["value"]; v != nil {
								additionalPropertiesMapToAdd["value"] = v
							}

							additionalPropertiesListToReturn = append(additionalPropertiesListToReturn, additionalPropertiesMapToAdd)
						}
						objectsMapToAdd["additional_properties"] = additionalPropertiesListToReturn
					}
				} else {
					objectsMapToAdd["additional_properties"] = nil
				}

				objectsListToReturn = append(objectsListToReturn, objectsMapToAdd)
			}

			_ = d.Set("objects", objectsListToReturn)
		} else {
			_ = d.Set("objects", objectsList)
		}

	} else {
		_ = d.Set("objects", nil)
	}

	if v := azureAdContent["to"]; v != nil {
		_ = d.Set("to", v)
	}

	if v := azureAdContent["total"]; v != nil {
		_ = d.Set("total", v)
	}

	return nil
}
