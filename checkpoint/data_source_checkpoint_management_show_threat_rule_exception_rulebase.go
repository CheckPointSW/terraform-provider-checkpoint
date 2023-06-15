package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementShowThreatRuleExceptionRuleBase() *schema.Resource {

	return &schema.Resource{

		Read: dataSourceManagementShowThreatRuleExceptionRuleBaseRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object uid",
			},
			"rule_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the parent rule.",
			},
			"rule_uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The UID of the parent rule.",
			},
			"rule_number": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The position in the rulebase of the parent rule.",
			},
			"filter": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Search expression to filter the rulebase.",
			},
			"filter_settings": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Sets filter preferences.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"search_mode": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "search mode",
						},
						"packet_search_settings": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "packet search settings",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"expand_group_members": {

										Type:        schema.TypeBool,
										Optional:    true,
										Description: "When true, if the search expression contains a UID or a name of a group object, results will include rules that match on at least one member of the group.",
									},
									"expand_group_with_exclusion_members": {

										Type:        schema.TypeBool,
										Optional:    true,
										Description: "When true, if the search expression contains a UID or a name of a group-with-exclusion object, results will include rules that match at least one member of the \"include\" part and is not a member of the \"except\" part.",
									},
									"match_on_any": {

										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether to match on 'Any' object",
									},
									"match_on_group_with_exclusion": {

										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether to match on a group-with-exclusion.",
									},
									"match_on_negate": {

										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether to match on a negated cell.",
									},
								},
							},
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
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Sorts the results by search criteria",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ASC": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Sorts results by the given field in ascending order.",
						},
						"DESC": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Sorts results by the given field in descending order.",
						},
					},
				},
			},
			"package": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of the package.",
			},
			"use_object_dictionary": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "",
				Default:     true,
			},
			"from": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "From which element number the query was done.",
			},
			"rulebase": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "group  name",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "rulebase type.",
						},
						"uid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "rulebase uid.",
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
						"rulebase": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "list of rulebases.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "rule name",
									},
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "rule type",
									},
									"uid": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "rule uid",
									},
									"install_on": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "Which Gateways identified by the name or UID to install the policy on.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"source": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "Collection of Network objects identified by the name or UID.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"source_negate": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "True if negate is set for source.",
									},
									"destination": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "Collection of Network objects identified by the name or UID.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"destination_negate": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "True if negate is set for destination.",
									},
									"service": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "Collection of Network objects identified by the name or UID.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"service_negate": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "True if negate is set for service.",
									},
									"protected_scope": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "Collection of objects defining Protected Scope identified by the name or UID.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"protected_scope_negate": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "True if negate is set for Protected Scope.",
									},
									"protection_or_site": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "Name of the protection or site.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"track": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Packet tracking.",
									},
									"action": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Action-the enforced profile.",
									},
									"exception_number": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The UID of the exception-group.",
									},
								},
							},
						},
					},
				},
			},
			"objects_dictionary": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "This table shows the level of details in the Standard level.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "rule name",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "rule type",
						},
						"uid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "rule uid",
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

func dataSourceManagementShowThreatRuleExceptionRuleBaseRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := map[string]interface{}{}
	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}
	ruleuid := d.Get("rule_uid").(string)

	rulename := d.Get("rule_name").(string)

	rulenumber := d.Get("rule_number").(string)

	if rulename != "" {
		payload["rule-name"] = rulename
	} else if ruleuid != "" {
		payload["rule-uid"] = ruleuid
	} else if rulenumber != "" {
		payload["rule-number"] = rulenumber
	}

	if v, ok := d.GetOk("filter"); ok {
		payload["filter"] = v.(string)
	}

	if v, ok := d.GetOk("filter_settings"); ok {
		filters, ok := v.(map[string]interface{})
		if ok {

			filtersMapToReturn := make(map[string]interface{})

			packetSearchMap := make(map[string]interface{})

			if val, ok := filters["search_mode"]; ok {
				filtersMapToReturn["search-mode"] = val
			} else {
				filtersMapToReturn["search-mode"] = "general"
			}

			if val, ok := filters["expand_group_members"]; ok {
				packetSearchMap["expand-group-members"] = val
			} else {
				packetSearchMap["expand-group-members"] = false
			}

			if val, ok := filters["expand_group_with_exclusion_members"]; ok {
				packetSearchMap["expand-group-with-exclusion-members"] = val
			} else {
				packetSearchMap["expand-group-with-exclusion-members"] = false
			}

			if val, ok := filters["match_on_any"]; ok {
				packetSearchMap["match-on-any"] = val
			} else {
				packetSearchMap["match-on-any"] = true
			}

			if val, ok := filters["match_on_group_with_exclusion"]; ok {
				packetSearchMap["match-on-group-with-exclusion"] = val
			} else {
				packetSearchMap["match-on-group-with-exclusion"] = true
			}

			if val, ok := filters["match_on_negate"]; ok {
				packetSearchMap["match-on-negate"] = val
			} else {
				packetSearchMap["match-on-negate"] = true
			}

			filtersMapToReturn["packet-search-settings"] = packetSearchMap
			payload["filter-settings"] = filtersMapToReturn
		}
	}
	if v, ok := d.GetOk("limit"); ok {
		payload["limit"] = v.(int)
	}

	if v, ok := d.GetOk("offset"); ok {
		payload["offset"] = v.(int)
	}

	useObjDict := true
	if v, ok := d.GetOkExists("use_object_dictionary"); ok {
		payload["use-object-dictionary"] = v.(bool)
		useObjDict = v.(bool)
	}

	showThreatRuleExceptionRuleBaseRes, err := client.ApiCall("show-threat-rule-exception-rulebase", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showThreatRuleExceptionRuleBaseRes.Success {
		return fmt.Errorf(showThreatRuleExceptionRuleBaseRes.ErrorMsg)
	}
	threatRuleExceptionRuleBase := showThreatRuleExceptionRuleBaseRes.GetData()

	log.Println("Read ruleBaseJson - Show JSON = ", threatRuleExceptionRuleBase)

	if v := threatRuleExceptionRuleBase["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}
	if v := threatRuleExceptionRuleBase["name"]; v != nil {
		_ = d.Set("name", v)
	}
	if v := threatRuleExceptionRuleBase["from"]; v != nil {

		val := v.(float64)

		_ = d.Set("from", int(val))
	}

	if threatRuleExceptionRuleBase["rulebase"] != nil {

		threatRuleExceptionRuleBaseList := threatRuleExceptionRuleBase["rulebase"].([]interface{})

		var ruleBaseArrToReturn []map[string]interface{}

		if len(threatRuleExceptionRuleBaseList) > 0 {

			for i := range threatRuleExceptionRuleBaseList {

				threatRuleExceptionRuleBaseMap := threatRuleExceptionRuleBaseList[i].(map[string]interface{})

				payload := make(map[string]interface{})

				if v := threatRuleExceptionRuleBaseMap["name"]; v != nil {
					payload["name"] = v.(string)
				}
				if v := threatRuleExceptionRuleBaseMap["type"]; v != nil {
					payload["type"] = v.(string)
				}
				if v := threatRuleExceptionRuleBaseMap["uid"]; v != nil {
					payload["uid"] = v.(string)
				}
				if v := threatRuleExceptionRuleBaseMap["from"]; v != nil {
					payload["from"] = int(v.(float64))
				}
				if v := threatRuleExceptionRuleBaseMap["to"]; v != nil {
					payload["to"] = int(v.(float64))
				}
				if v := threatRuleExceptionRuleBaseMap["rulebase"]; v != nil {
					ruleBaseList := v.([]interface{})
					var ruleBaseListToReturn []map[string]interface{}
					if len(ruleBaseList) > 0 {
						for i := range ruleBaseList {
							ruleBaseObj := ruleBaseList[i].(map[string]interface{})
							innerPayload := make(map[string]interface{})
							if v := ruleBaseObj["name"]; v != nil {
								innerPayload["name"] = v
							}
							if v := ruleBaseObj["type"]; v != nil {
								innerPayload["type"] = v
							}
							if v := ruleBaseObj["uid"]; v != nil {
								innerPayload["uid"] = v
							}
							if v := ruleBaseObj["source-negate"]; v != nil {
								innerPayload["source_negate"] = v.(bool)
							}
							if v := ruleBaseObj["service-negate"]; v != nil {
								innerPayload["service_negate"] = v.(bool)
							}
							if v := ruleBaseObj["destination-negate"]; v != nil {
								innerPayload["destination_negate"] = v.(bool)
							}
							if v := ruleBaseObj["protected-scope-negate"]; v != nil {
								innerPayload["protected_scope_negate"] = v.(bool)
							}
							if useObjDict {
								if v := ruleBaseObj["source"]; v != nil {
									innerPayload["source"] = v
								}
								if v := ruleBaseObj["action"]; v != nil {
									innerPayload["action"] = v.(string)
								}
								if v := ruleBaseObj["track"]; v != nil {
									innerPayload["track"] = v.(string)
								}
								if v := ruleBaseObj["service"]; v != nil {
									innerPayload["service"] = v
								}
								if v := ruleBaseObj["destination"]; v != nil {
									innerPayload["destination"] = v
								}
								if v := ruleBaseObj["protected-scope"]; v != nil {
									innerPayload["protected_scope"] = v
								}
								if v := ruleBaseObj["install-on"]; v != nil {
									innerPayload["install_on"] = v
								}
								if v := ruleBaseObj["protection-or-site"]; v != nil {
									innerPayload["protection_or_site"] = v
								}
							} else {
								if v := ruleBaseObj["action"]; v != nil {
									innerPayload["action"] = v.(map[string]interface{})["name"]
								}
								if v := ruleBaseObj["track"]; v != nil {
									innerPayload["track"] = v.(map[string]interface{})["name"]
								}
								if v := ruleBaseObj["source"]; v != nil {
									sourceJson := v.([]interface{})
									sourceNames := make([]string, 0)
									if len(sourceJson) > 0 {
										for _, source := range sourceJson {
											source := source.(map[string]interface{})
											sourceNames = append(sourceNames, source["name"].(string))
										}
									}
									_, sourceInConf := d.GetOk("source")
									if sourceNames[0] == "Any" && !sourceInConf {
										innerPayload["source"] = []interface{}{}
									} else {
										innerPayload["source"] = sourceNames
									}
								}
								if v := ruleBaseObj["service"]; v != nil {

									serviceJson := v.([]interface{})
									serviceJsonNames := make([]string, 0)
									if len(serviceJson) > 0 {
										for _, service := range serviceJson {
											service := service.(map[string]interface{})
											serviceJsonNames = append(serviceJsonNames, service["name"].(string))
										}
									}
									_, serviceInConf := d.GetOk("service")
									if serviceJsonNames[0] == "Any" && !serviceInConf {
										innerPayload["service"] = []interface{}{}
									} else {
										innerPayload["service"] = serviceJsonNames
									}
								}
								if v := ruleBaseObj["destination"]; v != nil {
									destinationJson := v.([]interface{})
									destinationNames := make([]string, 0)
									if len(destinationJson) > 0 {
										for _, destination := range destinationJson {
											destination := destination.(map[string]interface{})
											destinationNames = append(destinationNames, destination["name"].(string))
										}
									}
									_, destinationInConf := d.GetOk("destination")
									if destinationNames[0] == "Any" && !destinationInConf {
										innerPayload["destination"] = []interface{}{}
									} else {
										innerPayload["destination"] = destinationNames
									}
								}
								if v := ruleBaseObj["protected-scope"]; v != nil {
									protectedScopeJson := v.([]interface{})
									protectedScopeNames := make([]string, 0)
									if len(protectedScopeJson) > 0 {
										for _, protectedScope := range protectedScopeJson {
											protectedScope := protectedScope.(map[string]interface{})
											protectedScopeNames = append(protectedScopeNames, protectedScope["name"].(string))
										}
									}
									_, protectedScopeInConf := d.GetOk("protected_scope")
									if protectedScopeNames[0] == "Any" && !protectedScopeInConf {
										innerPayload["protected_scope"] = []interface{}{}
									} else {
										innerPayload["protected_scope"] = protectedScopeNames
									}
								}
								if v := ruleBaseObj["install-on"]; v != nil {
									installOnJson := v.([]interface{})
									installOnJsonNames := make([]string, 0)
									if len(installOnJson) > 0 {
										for _, installOn := range installOnJson {
											installOn := installOn.(map[string]interface{})
											installOnJsonNames = append(installOnJsonNames, installOn["name"].(string))
										}
									}
									_, installOnInConf := d.GetOk("install_on")
									if installOnJsonNames[0] == "Policy Targets" && !installOnInConf {
										innerPayload["install_on"] = []interface{}{}
									} else {
										innerPayload["install_on"] = installOnJsonNames
									}
								}
								if v := ruleBaseObj["protection-or-site"]; v != nil {
									protectionOrSiteJson := v.([]interface{})
									protectionOrSiteIds := make([]string, 0)
									if len(protectionOrSiteJson) > 0 {
										for _, protectionOrSite := range protectionOrSiteJson {
											protectionOrSite := protectionOrSite.(map[string]interface{})
											protectionOrSiteIds = append(protectionOrSiteIds, protectionOrSite["name"].(string))
										}
									}
									_, protectionOrSiteInConf := d.GetOk("protection_or_site")
									if protectionOrSiteIds[0] == "Any" && !protectionOrSiteInConf {
										innerPayload["protection_or_site"] = []interface{}{}
									} else {
										innerPayload["protection_or_site"] = protectionOrSiteIds
									}
								}
							}
							ruleBaseListToReturn = append(ruleBaseListToReturn, innerPayload)
						}
					}
					payload["rulebase"] = ruleBaseListToReturn
				}
				ruleBaseArrToReturn = append(ruleBaseArrToReturn, payload)
			}
		}
		d.Set("rulebase", ruleBaseArrToReturn)
	}
	if useObjDict {
		if v := threatRuleExceptionRuleBase["objects-dictionary"]; v != nil {
			var listOfObjectToReturn []map[string]interface{}
			objectDictionaryList := v.([]interface{})
			if len(objectDictionaryList) > 0 {
				for i := range objectDictionaryList {
					objDict := objectDictionaryList[i].(map[string]interface{})
					payload := make(map[string]interface{})
					if v := objDict["name"]; v != nil {
						payload["name"] = v.(string)
					}
					if v := objDict["type"]; v != nil {
						payload["type"] = v.(string)
					}
					if v := objDict["uid"]; v != nil {
						payload["uid"] = v.(string)
					}
					listOfObjectToReturn = append(listOfObjectToReturn, payload)
				}
			}
			d.Set("objects_dictionary", listOfObjectToReturn)
		}
	}
	if v := threatRuleExceptionRuleBase["to"]; v != nil {
		val := v.(float64)
		_ = d.Set("to", int(val))
	}
	if v := threatRuleExceptionRuleBase["total"]; v != nil {
		val := v.(float64)
		_ = d.Set("total", int(val))
	}
	return nil
}
