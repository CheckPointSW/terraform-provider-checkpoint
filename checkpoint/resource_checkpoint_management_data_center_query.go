package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"

	"strconv"
)

func resourceManagementDataCenterQuery() *schema.Resource {
	return &schema.Resource{
		Create: createManagementDataCenterQuery,
		Read:   readManagementDataCenterQuery,
		Update: updateManagementDataCenterQuery,
		Delete: deleteManagementDataCenterQuery,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"data_centers": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Collection of Data Center servers identified by the name or UID. Use \"All\" to select all data centers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"query_rules": {
				Type:        schema.TypeList,
				Optional:    true,
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

func createManagementDataCenterQuery(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	dataCenterQuery := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		dataCenterQuery["name"] = v.(string)
	}

	if v, ok := d.GetOk("data_centers"); ok {
		dataCentersList := v.([]interface{})
		if len(dataCentersList) == 1 && dataCentersList[0] == "All" {
			dataCenterQuery["data-centers"] = "All"
		} else {
			dataCenterQuery["data-centers"] = v
		}
	}

	if v, ok := d.GetOk("query_rules"); ok {

		queryRulesList := v.([]interface{})

		if len(queryRulesList) > 0 {

			var queryRulesPayload []map[string]interface{}

			for i := range queryRulesList {

				Payload := make(map[string]interface{})

				if v, ok := d.GetOk("query_rules." + strconv.Itoa(i) + ".key_type"); ok {
					Payload["key-type"] = v.(string)
				}
				if v, ok := d.GetOk("query_rules." + strconv.Itoa(i) + ".key"); ok {
					Payload["key"] = v.(string)
				}
				if v, ok := d.GetOk("query_rules." + strconv.Itoa(i) + ".values"); ok {
					Payload["values"] = v
				}
				queryRulesPayload = append(queryRulesPayload, Payload)
			}
			dataCenterQuery["query-rules"] = queryRulesPayload
		}
	}

	if v, ok := d.GetOk("tags"); ok {
		dataCenterQuery["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		dataCenterQuery["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		dataCenterQuery["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		dataCenterQuery["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		dataCenterQuery["ignore-errors"] = v.(bool)
	}

	log.Println("Create DataCenterQuery - Map = ", dataCenterQuery)

	addDataCenterQueryRes, err := client.ApiCall("add-data-center-query", dataCenterQuery, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !addDataCenterQueryRes.Success {
		if addDataCenterQueryRes.ErrorMsg != "" {
			return fmt.Errorf(addDataCenterQueryRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addDataCenterQueryRes.GetData()["uid"].(string))

	return readManagementDataCenterQuery(d, m)
}

func readManagementDataCenterQuery(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showDataCenterQueryRes, err := client.ApiCall("show-data-center-query", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showDataCenterQueryRes.Success {
		if objectNotFound(showDataCenterQueryRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showDataCenterQueryRes.ErrorMsg)
	}

	dataCenterQuery := showDataCenterQueryRes.GetData()

	KeysToFixedKeys := getKeysToFixedKeys()

	log.Println("Read DataCenterQuery - Show JSON = ", dataCenterQuery)

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

func updateManagementDataCenterQuery(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	dataCenterQuery := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		dataCenterQuery["name"] = oldName
		dataCenterQuery["new-name"] = newName
	} else {
		dataCenterQuery["name"] = d.Get("name")
	}

	if d.HasChange("data_centers") {
		if v, ok := d.GetOk("data_centers"); ok {
			dataCentersList := v.([]interface{})
			if len(dataCentersList) == 1 && dataCentersList[0] == "All" {
				dataCenterQuery["data-centers"] = "All"
			} else {
				dataCenterQuery["data-centers"] = v
			}
		}
	}

	if d.HasChange("query_rules") {

		if v, ok := d.GetOk("query_rules"); ok {

			queryRulesList := v.([]interface{})

			var queryRulesPayload []map[string]interface{}

			for i := range queryRulesList {

				Payload := make(map[string]interface{})
				Payload["key-type"] = d.Get("query_rules." + strconv.Itoa(i) + ".key_type")
				Payload["key"] = d.Get("query_rules." + strconv.Itoa(i) + ".key")
				Payload["values"] = d.Get("query_rules." + strconv.Itoa(i) + ".values")
				queryRulesPayload = append(queryRulesPayload, Payload)
			}
			dataCenterQuery["query-rules"] = queryRulesPayload
		} else {
			dataCenterQuery["query-rules"] = nil
		}
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			dataCenterQuery["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			dataCenterQuery["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		dataCenterQuery["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		dataCenterQuery["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		dataCenterQuery["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		dataCenterQuery["ignore-errors"] = v.(bool)
	}

	log.Println("Update DataCenterQuery - Map = ", dataCenterQuery)

	updateDataCenterQueryRes, err := client.ApiCall("set-data-center-query", dataCenterQuery, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !updateDataCenterQueryRes.Success {
		if updateDataCenterQueryRes.ErrorMsg != "" {
			return fmt.Errorf(updateDataCenterQueryRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementDataCenterQuery(d, m)
}

func deleteManagementDataCenterQuery(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	dataCenterQueryPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		dataCenterQueryPayload["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		dataCenterQueryPayload["ignore-errors"] = v.(bool)
	}

	log.Println("Delete DataCenterQuery")

	deleteDataCenterQueryRes, err := client.ApiCall("delete-data-center-query", dataCenterQueryPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteDataCenterQueryRes.Success {
		if deleteDataCenterQueryRes.ErrorMsg != "" {
			return fmt.Errorf(deleteDataCenterQueryRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
