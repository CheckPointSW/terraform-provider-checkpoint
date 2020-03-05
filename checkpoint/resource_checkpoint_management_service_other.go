package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"reflect"
	"strconv"
)

func resourceManagementServiceOther() *schema.Resource {   
    return &schema.Resource{
        Create: createManagementServiceOther,
        Read:   readManagementServiceOther,
        Update: updateManagementServiceOther,
        Delete: deleteManagementServiceOther,
        Schema: map[string]*schema.Schema{ 
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "Object name.",
            },
            "accept_replies": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Specifies whether Other Service replies are to be accepted.",
                Default:     false,
            },
            "action": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: "Contains an INSPECT expression that defines the action to take if a rule containing this service is matched. Example: set r_mhandler &open_ssl_handler sets a handler on the connection.",
            },
            "aggressive_aging": {
                Type:        schema.TypeMap,
                Optional:    true,
                Description: "Sets short (aggressive) timeouts for idle connections.",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "default_timeout": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: "Default aggressive aging timeout in seconds.",
                            Default:     0,
                        },
                        "enable": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: "N/A",
                            Default:     true,
                        },
                        "timeout": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: "Aggressive aging timeout in seconds.",
                            Default:     15,
                        },
                        "use_default_timeout": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: "N/A",
                            Default:     true,
                        },
                    },
                },
            },
            "ip_protocol": {
                Type:        schema.TypeInt,
                Optional:    true,
                Description: "IP protocol number.",
            },
            "keep_connections_open_after_policy_installation": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections.",
            },
            "match": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: "Contains an INSPECT expression that defines the matching criteria. The connection is examined against the expression during the first packet. Example: tcp, dport = 21, direction = 0 matches incoming FTP control connections.",
            },
            "match_for_any": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Indicates whether this service is used when 'Any' is set as the rule's service and there are several service objects with the same source port and protocol.",
            },
            "override_default_settings": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Indicates whether this service is a Data Domain service which has been overridden.",
                Default:     false,
            },
            "session_timeout": {
                Type:        schema.TypeInt,
                Optional:    true,
                Description: "Time (in seconds) before the session times out.",
            },
            "sync_connections_on_cluster": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enables state-synchronized High Availability or Load Sharing on a ClusterXL or OPSEC-certified cluster.",
            },
            "tags": {
                Type:        schema.TypeSet,
                Optional:    true,
                Description: "Collection of tag identifiers.",
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
            "use_default_session_timeout": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Use default virtual session timeout.",
                Default:     true,
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
        },
    }
}

func createManagementServiceOther(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)

    serviceOther := make(map[string]interface{})

    if v, ok := d.GetOk("name"); ok {
        serviceOther["name"] = v.(string)
    }

    if v, ok := d.GetOkExists("accept_replies"); ok {
        serviceOther["accept-replies"] = v.(bool)
    }

    if v, ok := d.GetOk("action"); ok {
        serviceOther["action"] = v.(string)
    }

    if _, ok := d.GetOk("aggressive_aging"); ok {

        res := make(map[string]interface{})

        if v, ok := d.GetOk("aggressive_aging.default_timeout"); ok {
            res["default-timeout"] = v
        }
        if v, ok := d.GetOk("aggressive_aging.enable"); ok {
            res["enable"] = v
        }
        if v, ok := d.GetOk("aggressive_aging.timeout"); ok {
            res["timeout"] = v
        }
        if v, ok := d.GetOk("aggressive_aging.use_default_timeout"); ok {
            res["use-default-timeout"] = v
        }
        serviceOther["aggressive-aging"] = res
    }

    if v, ok := d.GetOk("ip_protocol"); ok {
        serviceOther["ip-protocol"] = v.(int)
    }

    if v, ok := d.GetOkExists("keep_connections_open_after_policy_installation"); ok {
        serviceOther["keep-connections-open-after-policy-installation"] = v.(bool)
    }

    if v, ok := d.GetOk("match"); ok {
        serviceOther["match"] = v.(string)
    }

    if v, ok := d.GetOkExists("match_for_any"); ok {
        serviceOther["match-for-any"] = v.(bool)
    }

    if v, ok := d.GetOkExists("override_default_settings"); ok {
        serviceOther["override-default-settings"] = v.(bool)
    }

    if v, ok := d.GetOk("session_timeout"); ok {
        serviceOther["session-timeout"] = v.(int)
    }

    if v, ok := d.GetOkExists("sync_connections_on_cluster"); ok {
        serviceOther["sync-connections-on-cluster"] = v.(bool)
    }

    if v, ok := d.GetOk("tags"); ok {
        serviceOther["tags"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOkExists("use_default_session_timeout"); ok {
        serviceOther["use-default-session-timeout"] = v.(bool)
    }

    if v, ok := d.GetOk("color"); ok {
        serviceOther["color"] = v.(string)
    }

    if v, ok := d.GetOk("comments"); ok {
        serviceOther["comments"] = v.(string)
    }

    if v, ok := d.GetOk("groups"); ok {
        serviceOther["groups"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
        serviceOther["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
        serviceOther["ignore-errors"] = v.(bool)
    }

    log.Println("Create ServiceOther - Map = ", serviceOther)

    addServiceOtherRes, err := client.ApiCall("add-service-other", serviceOther, client.GetSessionID(), true, false)
    if err != nil || !addServiceOtherRes.Success {
        if addServiceOtherRes.ErrorMsg != "" {
            return fmt.Errorf(addServiceOtherRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }

    d.SetId(addServiceOtherRes.GetData()["uid"].(string))

    return readManagementServiceOther(d, m)
}

func readManagementServiceOther(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    payload := map[string]interface{}{
        "uid": d.Id(),
    }

    showServiceOtherRes, err := client.ApiCall("show-service-other", payload, client.GetSessionID(), true, false)
    if err != nil {
		return fmt.Errorf(err.Error())
	}
    if !showServiceOtherRes.Success {
		if objectNotFound(showServiceOtherRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
        return fmt.Errorf(showServiceOtherRes.ErrorMsg)
    }

    serviceOther := showServiceOtherRes.GetData()

    log.Println("Read ServiceOther - Show JSON = ", serviceOther)

	if v := serviceOther["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := serviceOther["accept-replies"]; v != nil {
		_ = d.Set("accept_replies", v)
	}

	if v := serviceOther["action"]; v != nil {
		_ = d.Set("action", v)
	}

    if serviceOther["aggressive-aging"] != nil {

        aggressiveAgingMap := serviceOther["aggressive-aging"].(map[string]interface{})

        aggressiveAgingMapToReturn := make(map[string]interface{})

        if v, _ := aggressiveAgingMap["default-timeout"]; v != nil{
            aggressiveAgingMapToReturn["default_timeout"] = v
        }
        if v, _ := aggressiveAgingMap["enable"]; v != nil{
            aggressiveAgingMapToReturn["enable"] = strconv.FormatBool(v.(bool))
        }
        if v, _ := aggressiveAgingMap["timeout"]; v != nil{
            aggressiveAgingMapToReturn["timeout"] = v
        }
        if v, _ := aggressiveAgingMap["use-default-timeout"]; v != nil{
            aggressiveAgingMapToReturn["use_default_timeout"] = strconv.FormatBool(v.(bool))
        }

        _, aggressiveAgingInConf := d.GetOk("aggressive_aging")
        defaultAggressiveAging := map[string]interface{}{"enable": "true", "timeout": "15", "use_default_timeout": "true", "default_timeout": "0"}
        if reflect.DeepEqual(defaultAggressiveAging, aggressiveAgingMapToReturn) && !aggressiveAgingInConf {
            _ = d.Set("aggressive_aging", map[string]interface{}{})
        } else {
            _ = d.Set("aggressive_aging", aggressiveAgingMapToReturn)
        }

    } else {
        _ = d.Set("aggressive_aging", nil)
    }

	if v := serviceOther["ip-protocol"]; v != nil {
		_ = d.Set("ip_protocol", v)
	}

	if v := serviceOther["keep-connections-open-after-policy-installation"]; v != nil {
		_ = d.Set("keep_connections_open_after_policy_installation", v)
	}

	if v := serviceOther["match"]; v != nil {
		_ = d.Set("match", v)
	}

	if v := serviceOther["match-for-any"]; v != nil {
		_ = d.Set("match_for_any", v)
	}

	if v := serviceOther["override-default-settings"]; v != nil {
		_ = d.Set("override_default_settings", v)
	}

	if v := serviceOther["session-timeout"]; v != nil {
		_ = d.Set("session_timeout", v)
	}

	if v := serviceOther["sync-connections-on-cluster"]; v != nil {
		_ = d.Set("sync_connections_on_cluster", v)
	}

    if serviceOther["tags"] != nil {
        tagsJson, ok := serviceOther["tags"].([]interface{})
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

	if v := serviceOther["use-default-session-timeout"]; v != nil {
		_ = d.Set("use_default_session_timeout", v)
	}

	if v := serviceOther["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := serviceOther["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

    if serviceOther["groups"] != nil {
        groupsJson, ok := serviceOther["groups"].([]interface{})
        if ok {
            groupsIds := make([]string, 0)
            if len(groupsJson) > 0 {
                for _, groups := range groupsJson {
                    groups := groups.(map[string]interface{})
                    groupsIds = append(groupsIds, groups["name"].(string))
                }
            }
        _ = d.Set("groups", groupsIds)
        }
    } else {
        _ = d.Set("groups", nil)
    }

	if v := serviceOther["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := serviceOther["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementServiceOther(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
    serviceOther := make(map[string]interface{})

    if ok := d.HasChange("name"); ok {
        oldName, newName := d.GetChange("name")
        serviceOther["name"] = oldName
        serviceOther["new-name"] = newName
    } else {
        serviceOther["name"] = d.Get("name")
    }

    if v, ok := d.GetOkExists("accept_replies"); ok {
	       serviceOther["accept-replies"] = v.(bool)
    }

    if ok := d.HasChange("action"); ok {
	       serviceOther["action"] = d.Get("action")
    }

    if d.HasChange("aggressive_aging") {

        if _, ok := d.GetOk("aggressive_aging"); ok {

            res := make(map[string]interface{})

            if d.HasChange("aggressive_aging.default_timeout") {
                res["default-timeout"] = d.Get("aggressive_aging.default_timeout")
            }
            if d.HasChange("aggressive_aging.enable") {
                res["enable"] = d.Get("aggressive_aging.enable")
            }
            if d.HasChange("aggressive_aging.timeout") {
                res["timeout"] = d.Get("aggressive_aging.timeout")
            }
            if d.HasChange("aggressive_aging.use_default_timeout") {
                res["use-default-timeout"] = d.Get("aggressive_aging.use_default_timeout")
            }
            serviceOther["aggressive-aging"] = res
        } else {
            serviceOther["aggressive-aging"] = map[string]interface{}{"enable": true, "timeout": "15", "use-default-timeout": true, "default-timeout": "0"}
        }
    }

    if ok := d.HasChange("ip_protocol"); ok {
	       serviceOther["ip-protocol"] = d.Get("ip_protocol")
    }

    if v, ok := d.GetOkExists("keep_connections_open_after_policy_installation"); ok {
	       serviceOther["keep-connections-open-after-policy-installation"] = v.(bool)
    }

    if ok := d.HasChange("match"); ok {
	       serviceOther["match"] = d.Get("match")
    }

    if v, ok := d.GetOkExists("match_for_any"); ok {
	       serviceOther["match-for-any"] = v.(bool)
    }

    if v, ok := d.GetOkExists("override_default_settings"); ok {
	       serviceOther["override-default-settings"] = v.(bool)
    }

    if ok := d.HasChange("session_timeout"); ok {
	       serviceOther["session-timeout"] = d.Get("session_timeout")
    }

    if v, ok := d.GetOkExists("sync_connections_on_cluster"); ok {
	       serviceOther["sync-connections-on-cluster"] = v.(bool)
    }

    if d.HasChange("tags") {
        if v, ok := d.GetOk("tags"); ok {
            serviceOther["tags"] = v.(*schema.Set).List()
        } else {
            oldTags, _ := d.GetChange("tags")
	           serviceOther["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
        }
    }

    if v, ok := d.GetOkExists("use_default_session_timeout"); ok {
	       serviceOther["use-default-session-timeout"] = v.(bool)
    }

    if ok := d.HasChange("color"); ok {
	       serviceOther["color"] = d.Get("color")
    }

    if ok := d.HasChange("comments"); ok {
	       serviceOther["comments"] = d.Get("comments")
    }

    if d.HasChange("groups") {
        if v, ok := d.GetOk("groups"); ok {
            serviceOther["groups"] = v.(*schema.Set).List()
        } else {
            oldGroups, _ := d.GetChange("groups")
	           serviceOther["groups"] = map[string]interface{}{"remove": oldGroups.(*schema.Set).List()}
        }
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
	       serviceOther["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
	       serviceOther["ignore-errors"] = v.(bool)
    }

    log.Println("Update ServiceOther - Map = ", serviceOther)

    updateServiceOtherRes, err := client.ApiCall("set-service-other", serviceOther, client.GetSessionID(), true, false)
    if err != nil || !updateServiceOtherRes.Success {
        if updateServiceOtherRes.ErrorMsg != "" {
            return fmt.Errorf(updateServiceOtherRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }

    return readManagementServiceOther(d, m)
}

func deleteManagementServiceOther(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    serviceOtherPayload := map[string]interface{}{
        "uid": d.Id(),
    }

    log.Println("Delete ServiceOther")

    deleteServiceOtherRes, err := client.ApiCall("delete-service-other", serviceOtherPayload , client.GetSessionID(), true, false)
    if err != nil || !deleteServiceOtherRes.Success {
        if deleteServiceOtherRes.ErrorMsg != "" {
            return fmt.Errorf(deleteServiceOtherRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }
    d.SetId("")

    return nil
}

