package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strconv"
)

func dataSourceManagementServicesUdp() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementServicesUdpRead,
		Schema: map[string]*schema.Schema{
			"filter": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Search expression to filter objects by.",
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
			"fetch_all": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If true, fetches all results.",
				Default:     false,
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
						"accept_replies": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "N/A",
						},
						"aggressive_aging": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Sets short (aggressive) timeouts for idle connections.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"default_timeout": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Default aggressive aging timeout in seconds.",
									},
									"enable": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "N/A",
									},
									"timeout": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Aggressive aging timeout in seconds.",
									},
									"use_default_timeout": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "N/A",
									},
								},
							},
						},
						"keep_connections_open_after_policy_installation": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections.",
						},
						"match_by_protocol_signature": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "A value of true enables matching by the selected protocol's signature - the signature identifies the protocol as genuine. Select this option to limit the port to the specified protocol. If the selected protocol does not support matching by signature, this field cannot be set to true.",
						},
						"match_for_any": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether this service is used when 'Any' is set as the rule's service and there are several service objects with the same source port and protocol.",
						},
						"override_default_settings": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether this service is a Data Domain service which has been overridden.",
						},
						"port": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The number of the port used to provide this service. To specify a port range, place a hyphen between the lowest and highest port numbers, for example 44-55.",
						},
						"protocol": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Select the protocol type associated with the service, and by implication, the management server (if any) that enforces Content Security and Authentication for the service. Selecting a Protocol Type invokes the specific protocol handlers for each protocol type, thus enabling higher level of security by parsing the protocol, and higher level of connectivity by tracking dynamic actions (such as opening of ports).",
						},
						"session_timeout": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Time (in seconds) before the session times out.",
						},
						"source_port": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Port number for the client side service. If specified, only those Source port Numbers will be Accepted, Dropped, or Rejected during packet inspection. Otherwise, the source port is not inspected.",
						},
						"sync_connections_on_cluster": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enables state-synchronized High Availability or Load Sharing on a ClusterXL or OPSEC-certified cluster.",
						},
						"use_default_session_timeout": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Use default virtual session timeout.",
						},
						"groups": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Collection of group name.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
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
						"icon": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object icon.",
						},
					},
				},
			},
		},
	}
}

func dataSourceManagementServicesUdpRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	showServicesUdpRes := checkpoint.APIResponse{}
	var err error
	fetchAll, _ := d.GetOkExists("fetch_all")

	if fetchAll.(bool) {
		showServicesUdpRes, err = client.ApiQuery("show-services-udp", "full", "objects", true, map[string]interface{}{})
	} else {
		payload := make(map[string]interface{})

		if val, ok := d.GetOk("filter"); ok {
			payload["filter"] = val.(string)
		}

		if val, ok := d.GetOk("limit"); ok {
			payload["limit"] = val.(int)
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

		if val, ok := d.GetOk("offset"); ok {
			payload["offset"] = val.(int)
		}

		payload["details-level"] = "full"
		showServicesUdpRes, err = client.ApiCallSimple("show-services-udp", payload)
	}

	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showServicesUdpRes.Success {
		return fmt.Errorf(showServicesUdpRes.ErrorMsg)
	}

	servicesUdp := showServicesUdpRes.GetData()

	log.Println("Read ServicesUdp - Show JSON = ", servicesUdp)

	d.SetId("show-services-udp-" + acctest.RandString(10))

	if v := servicesUdp["from"]; v != nil {
		_ = d.Set("from", v)
	}

	if v := servicesUdp["to"]; v != nil {
		_ = d.Set("to", v)
	}

	if v := servicesUdp["total"]; v != nil {
		_ = d.Set("total", v)
	}

	if v := servicesUdp["objects"]; v != nil {
		objectsList := v.([]interface{})
		if len(objectsList) > 0 {
			var objectsListState []map[string]interface{}
			for i := range objectsList {
				objectMap := objectsList[i].(map[string]interface{})
				objectMapToAdd := make(map[string]interface{})

				if v := objectMap["name"]; v != nil {
					objectMapToAdd["name"] = v
				}

				if v := objectMap["uid"]; v != nil {
					objectMapToAdd["uid"] = v
				}

				if v := objectMap["comments"]; v != nil {
					objectMapToAdd["comments"] = v
				}

				if v := objectMap["accept-replies"]; v != nil {
					objectMapToAdd["accept_replies"] = v
				}

				if v := objectMap["keep-connections-open-after-policy-installation"]; v != nil {
					objectMapToAdd["keep_connections_open_after_policy_installation"] = v
				}

				if v := objectMap["match-by-protocol-signature"]; v != nil {
					objectMapToAdd["match_by_protocol_signature"] = v
				}

				if v := objectMap["match-for-any"]; v != nil {
					objectMapToAdd["match_for_any"] = v
				}

				if v := objectMap["override-default-settings"]; v != nil {
					objectMapToAdd["override_default_settings"] = v
				}

				if v := objectMap["port"]; v != nil {
					objectMapToAdd["port"] = v
				}

				if v := objectMap["protocol"]; v != nil {
					objectMapToAdd["protocol"] = v
				}

				if v := objectMap["session-timeout"]; v != nil {
					objectMapToAdd["session_timeout"] = v
				}

				if v := objectMap["source-port"]; v != nil {
					objectMapToAdd["source_port"] = v
				}

				if v := objectMap["sync-connections-on-cluster"]; v != nil {
					objectMapToAdd["sync_connections_on_cluster"] = v
				}

				if v := objectMap["use-default-session-timeout"]; v != nil {
					objectMapToAdd["use_default_session_timeout"] = v
				}

				if v := objectMap["color"]; v != nil {
					objectMapToAdd["color"] = v
				}

				if objectMap["aggressive-aging"] != nil {

					aggressiveAgingMap := objectMap["aggressive-aging"].(map[string]interface{})

					aggressiveAgingMapToReturn := make(map[string]interface{})

					if v, _ := aggressiveAgingMap["default-timeout"]; v != nil {
						aggressiveAgingMapToReturn["default_timeout"] = int(v.(float64))
					}
					if v, _ := aggressiveAgingMap["enable"]; v != nil {
						aggressiveAgingMapToReturn["enable"] = v
					}
					if v, _ := aggressiveAgingMap["timeout"]; v != nil {
						aggressiveAgingMapToReturn["timeout"] = int(v.(float64))
					}
					if v, _ := aggressiveAgingMap["use-default-timeout"]; v != nil {
						aggressiveAgingMapToReturn["use_default_timeout"] = v
					}

					objectMapToAdd["aggressive_aging"] = []interface{}{aggressiveAgingMapToReturn}

				}

				if objectMap["groups"] != nil {
					groupsJson := objectMap["groups"].([]interface{})
					groupsIds := make([]string, 0)
					if len(groupsJson) > 0 {
						// Create slice of group names
						for _, group_ := range groupsJson {
							group_ := group_.(map[string]interface{})
							groupsIds = append(groupsIds, group_["name"].(string))
						}
					}
					objectMapToAdd["groups"] = groupsIds
				} else {
					objectMapToAdd["groups"] = nil
				}

				if objectMap["tags"] != nil {
					tagsJson := objectMap["tags"].([]interface{})
					var tagsIds = make([]string, 0)
					if len(tagsJson) > 0 {
						// Create slice of tag names
						for _, tag := range tagsJson {
							tag := tag.(map[string]interface{})
							tagsIds = append(tagsIds, tag["name"].(string))
						}
					}
					objectMapToAdd["tags"] = tagsIds
				} else {
					objectMapToAdd["tags"] = nil
				}

				if v := objectMap["domain"]; v != nil {
					domainMap := v.(map[string]interface{})
					domainMapToAdd := make(map[string]interface{})

					if v := domainMap["name"]; v != nil {
						domainMapToAdd["name"] = v
					}

					if v := domainMap["uid"]; v != nil {
						domainMapToAdd["uid"] = v
					}

					if v := domainMap["domain-type"]; v != nil {
						domainMapToAdd["domain_type"] = v
					}
					objectMapToAdd["domain"] = domainMapToAdd
				}

				if v := objectMap["icon"]; v != nil {
					objectMapToAdd["icon"] = v
				}

				objectsListState = append(objectsListState, objectMapToAdd)
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
