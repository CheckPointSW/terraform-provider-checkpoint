package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementDataCenterQuery() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDataCenterQueryRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name. Must be unique in the domain.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object unique identifier.",
			},
			"data_centers": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Collection of Data Center servers identified by the name or UID. Use \"All\" to select all data centers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"query_rules": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Data Center Query Rules.<br>There is an 'AND' operation between multiple Query Rules.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The type of the \"key\" parameter.<br>Use \"predefined\" for these keys: type-in-data-center, name-in-data-center, and ip-address.<br>Use \"tag\" to query the Data Center tag�s property.",
						},
						"key": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Defines in which Data Center property to query.<br>For key-type \"predefined\", use these keys: type-in-data-center, name-in-data-center, and ip-address.<br>For key-type \"tag\", use the Data Center tag key to query.<br>Keys are case-insensitive.",
						},
						"values": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "The value(s) of the Data Center property to match the Query Rule.<br>Values are case-insensitive.<br>There is an 'OR' operation between multiple values.<br>For key-type \"predefined\" and key 'ip-address', the values must be an IPv4 or IPv6 address.<br>For key-type \"tag\", the values must be the Data Center tag values.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
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

func dataSourceDataCenterQueryRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	var name string
	var uid string

	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}
	if v, ok := d.GetOk("uid"); ok {
		uid = v.(string)
	}
	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}
	showDataCenterQueryRes, err := client.ApiCall("show-data-center-query", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showDataCenterQueryRes.Success {
		return fmt.Errorf(showDataCenterQueryRes.ErrorMsg)
	}

	dataCenterQuery := showDataCenterQueryRes.GetData()

	KeysToFixedKeys := getKeysToFixedKeys()

	log.Println("Read DataCenterQuery - Show JSON = ", dataCenterQuery)

	if v := dataCenterQuery["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := dataCenterQuery["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if dataCenterQuery["data-centers"] != nil {
		dataCentersJson, ok := dataCenterQuery["data-centers"].([]interface{})
		if ok {
			dataCentersIds := make([]string, 0)
			if len(dataCentersJson) > 0 {
				for _, data_centers := range dataCentersJson {
					data_centers := data_centers.(map[string]interface{})
					dataCentersIds = append(dataCentersIds, data_centers["name"].(string))
				}
				_ = d.Set("data_centers", dataCentersIds)
			} else {
				_ = d.Set("data_centers", []string{"All"})
			}
		}
	}

	if dataCenterQuery["query-rules"] != nil {

		queryRulesList, ok := dataCenterQuery["query-rules"].([]interface{})

		if ok {

			if len(queryRulesList) > 0 {

				var queryRulesListToReturn []map[string]interface{}

				for i := range queryRulesList {

					queryRulesMap := queryRulesList[i].(map[string]interface{})

					queryRulesMapToAdd := make(map[string]interface{})

					if v, _ := queryRulesMap["key-type"]; v != nil {
						keyType := v.(string)
						if newType, ok := KeysToFixedKeys[keyType]; ok {
							queryRulesMapToAdd["key_type"] = newType
						} else {
							queryRulesMapToAdd["key_type"] = v
						}
					}
					if v, _ := queryRulesMap["key"]; v != nil {
						key := v.(string)
						if newType, ok := KeysToFixedKeys[key]; ok {
							queryRulesMapToAdd["key"] = newType
						} else {
							queryRulesMapToAdd["key"] = v
						}
					}
					if v, _ := queryRulesMap["values"]; v != nil {
						queryRulesMapToAdd["values"] = v
					}
					queryRulesListToReturn = append(queryRulesListToReturn, queryRulesMapToAdd)
				}
				_ = d.Set("query_rules", queryRulesListToReturn)
			}
		}
	}

	if dataCenterQuery["tags"] != nil {
		tagsJson, ok := dataCenterQuery["tags"].([]interface{})
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

	if v := dataCenterQuery["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := dataCenterQuery["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := dataCenterQuery["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := dataCenterQuery["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}
