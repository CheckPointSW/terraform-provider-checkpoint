package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"math"
	"strconv"
	"strings"
)

func dataSourceManagementDataCenterContent() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementDataCenterContentRead,

		Schema: map[string]*schema.Schema{
			"data_center_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of the Data Center Server where to search for objects.",
			},
			"data_center_uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Unique identifier of the Data Center Server where to search for objects.",
			},
			"limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The maximal number of returned results.",
				Default:     50,
			},
			"offset": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Number of the results to initially skip.",
				Default:     0,
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
			"uid_in_data_center": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Return result matching the unique identifier of the object on the Data Center Server.",
			},
			"filter": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Return results matching the specified filter.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"text": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Return results containing the specified text value.",
						},
						"uri": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Return results under the specified Data Center Object (identified by URI).",
						},
						"parent_uid_in_data_center": {
							Type:        schema.TypeString,
							Computed:    true,
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
						"name_in_data_center": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object name in the Data Center.",
						},
						"uid_in_data_center": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Unique identifier of the object in the Data Center.",
						},
						"data_center_object": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "The imported management object (if exists). Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.",
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
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Object type.",
									},
								},
							},
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object management name.",
						},
						"type_in_data_center": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object type in Data Center.",
						},
						"additional_properties": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Additional properties on the object.\nRemote objects views.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "N/A",
									},
									"value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "N/A",
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

func dataSourceManagementDataCenterContentRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("data_center_name").(string)
	uid := d.Get("data_center_uid").(string)

	payload := map[string]interface{}{}

	if name != "" {
		payload["data-center-name"] = name
	} else if uid != "" {
		payload["data-center-uid"] = uid
	}

	if v, ok := d.GetOk("limit"); ok {
		payload["limit"] = v.(int)
	}
	if v, ok := d.GetOk("offset"); ok {
		payload["offset"] = v.(int)
	}
	if v, ok := d.GetOk("order"); ok {

		ordersList, ok := v.([]interface{})
		var ordersDictToReturn []map[string]interface{}

		if ok {
			for i := range ordersList {

				objectsMap := ordersList[i].(map[string]interface{})

				tempOrder := make(map[string]interface{})

				if v, _ := objectsMap["asc"]; v != nil && v != "" {
					tempOrder["ASC"] = v
				}

				if v, _ := objectsMap["desc"]; v != nil && v != "" {
					tempOrder["DESC"] = v
				}

				ordersDictToReturn = append(ordersDictToReturn, tempOrder)
			}
			payload["order"] = ordersDictToReturn
		}
	}
	if v, ok := d.GetOk("uid_in_data_center"); ok {
		payload["uid-in-data-center"] = v.(string)
	}
	if v, ok := d.GetOk("filter"); ok {
		dataCenterContentFilter := v.(map[string]interface{})
		dataCenterContentFilterMap := make(map[string]interface{})
		if v, ok := dataCenterContentFilter["text"]; ok {
			dataCenterContentFilterMap["text"] = v.(string)
		}
		if v, ok := dataCenterContentFilter["uri"]; ok {
			dataCenterContentFilterMap["uri"] = v.(string)
		}
		if v, ok := dataCenterContentFilter["parent_uid_in_data_center"]; ok {
			dataCenterContentFilterMap["parent-uid-in-data-center"] = v.(string)
		}
		payload["filter"] = dataCenterContentFilterMap
	}
	showDataCenterContentRes, err := client.ApiCall("show-data-center-content", payload, client.GetSessionID(), true, client.IsProxyUsed())

	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showDataCenterContentRes.Success {
		return fmt.Errorf(showDataCenterContentRes.ErrorMsg)
	}
	DataCenterContent := showDataCenterContentRes.GetData()

	log.Println("Read DataCenterContent - Show JSON = ", DataCenterContent)

	d.SetId("show-data-center-content-" + acctest.RandString(10))

	if v := DataCenterContent["data-center-name"]; v != nil {
		_ = d.Set("data_center_name", v)
	}

	if v := DataCenterContent["from"]; v != nil {
		_ = d.Set("from", int(math.Round(v.(float64))))
	}

	if DataCenterContent["objects"] != nil {
		dataCenterContentObjects := DataCenterContent["objects"].([]interface{})
		var dataCenterContentObjectToReturn []map[string]interface{}
		for i, _ := range dataCenterContentObjects {
			dataCenterContentObject := dataCenterContentObjects[i].(map[string]interface{})
			dataCenterContentObjectsMap := make(map[string]interface{})
			if v, ok := dataCenterContentObject["name-in-data-center"]; ok {
				dataCenterContentObjectsMap["name_in_data_center"] = v.(string)
			}
			if v, ok := dataCenterContentObject["uid-in-data-center"]; ok {
				dataCenterContentObjectsMap["uid_in_data_center"] = v.(string)
			}
			if v, ok := dataCenterContentObject["data-center-object"]; ok {
				dataCenterContentObjectsMap["data_center_object"] = v
			}
			if v, ok := dataCenterContentObject["name"]; ok {
				dataCenterContentObjectsMap["name"] = v.(string)
			}
			if v, ok := dataCenterContentObject["type-in-data-center"]; ok {
				dataCenterContentObjectsMap["type_in_data_center"] = v.(string)
			}
			if v, ok := dataCenterContentObject["additional-properties"]; ok {
				propsJson, ok := v.([]interface{})
				propsMapToReturn := make(map[string]interface{})
				if ok {
					for _, prop := range propsJson {
						propMap := prop.(map[string]interface{})
						propName := strings.ReplaceAll(propMap["name"].(string), "-", "_")
						propValue := propMap["value"]
						if propName == "unsafe_auto_accept" || propName == "enable_sts_assume_role" {
							propValue, _ = strconv.ParseBool(propValue.(string))
						}
						if propName == "urls" || propName == "hostnames" {
							propValue = strings.Split(propValue.(string), ";")
						}
						propsMapToReturn[propName] = propValue
					}
				}
				dataCenterContentObjectsMap["additional_properties"] = propsMapToReturn
			}
			dataCenterContentObjectToReturn = append(dataCenterContentObjectToReturn, dataCenterContentObjectsMap)
		}
		_ = d.Set("objects", dataCenterContentObjectToReturn)
	}

	if v := DataCenterContent["to"]; v != nil {
		_ = d.Set("to", int(math.Round(v.(float64))))
	}
	if v := DataCenterContent["total"]; v != nil {
		_ = d.Set("total", int(math.Round(v.(float64))))
	}
	return nil
}
