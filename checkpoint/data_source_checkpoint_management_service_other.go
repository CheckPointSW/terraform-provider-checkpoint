package checkpoint

import (
	"fmt"
	"log"
	"reflect"

	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceManagementServiceOther() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementServiceOtherRead,
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
			"accept_replies": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Specifies whether Other Service replies are to be accepted.",
			},
			"action": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Contains an INSPECT expression that defines the action to take if a rule containing this service is matched. Example: set r_mhandler &open_ssl_handler sets a handler on the connection.",
			},
			"aggressive_aging": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Sets short (aggressive) timeouts for idle connections.",
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
			"ip_protocol": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "IP protocol number.",
			},
			"keep_connections_open_after_policy_installation": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections.",
			},
			"match": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Contains an INSPECT expression that defines the matching criteria. The connection is examined against the expression during the first packet. Example: tcp, dport = 21, direction = 0 matches incoming FTP control connections.",
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
			"session_timeout": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Time (in seconds) before the session times out.",
			},
			"sync_connections_on_cluster": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enables state-synchronized High Availability or Load Sharing on a ClusterXL or OPSEC-certified cluster.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"use_default_session_timeout": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Use default virtual session timeout.",
			},
			"color": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Color of the object. Should be one of existing colors.",
			},
			"comments": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comments string.",
			},
			"groups": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of group identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceManagementServiceOtherRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showServiceOtherRes, err := client.ApiCall("show-service-other", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showServiceOtherRes.Success {
		return fmt.Errorf(showServiceOtherRes.ErrorMsg)
	}

	serviceOther := showServiceOtherRes.GetData()

	log.Println("Read ServiceOther - Show JSON = ", serviceOther)

	if v := serviceOther["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

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

		if v, _ := aggressiveAgingMap["default-timeout"]; v != nil {
			aggressiveAgingMapToReturn["default_timeout"] = v
		}
		if v, _ := aggressiveAgingMap["enable"]; v != nil {
			aggressiveAgingMapToReturn["enable"] = v
		}
		if v, _ := aggressiveAgingMap["timeout"]; v != nil {
			aggressiveAgingMapToReturn["timeout"] = v
		}
		if v, _ := aggressiveAgingMap["use-default-timeout"]; v != nil {
			aggressiveAgingMapToReturn["use_default_timeout"] = v
		}

		_, aggressiveAgingInConf := d.GetOk("aggressive_aging")
		defaultAggressiveAging := map[string]interface{}{
			"enable":              true,
			"timeout":             600,
			"use_default_timeout": true,
			"default_timeout":     600,
		}
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

	return nil
}
