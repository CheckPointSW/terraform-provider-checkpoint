package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"reflect"
)

func dataSourceManagementServiceTcp() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementServiceTcpRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name. Should be unique in the domain.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object unique identifier.",
			},
			"port": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The number of the port used to provide this service. To specify a port range, place a hyphen between the lowest and highest port numbers, for example 44-55.",
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
		},
	}
}

func dataSourceManagementServiceTcpRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showServiceTcpRes, err := client.ApiCall("show-service-tcp", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showServiceTcpRes.Success {
		return fmt.Errorf(showServiceTcpRes.ErrorMsg)
	}

	serviceTcp := showServiceTcpRes.GetData()

	if v := serviceTcp["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := serviceTcp["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := serviceTcp["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := serviceTcp["keep-connections-open-after-policy-installation"]; v != nil {
		_ = d.Set("keep_connections_open_after_policy_installation", v)
	}

	if v := serviceTcp["match-by-protocol-signature"]; v != nil {
		_ = d.Set("match_by_protocol_signature", v)
	}

	if v := serviceTcp["match-for-any"]; v != nil {
		_ = d.Set("match_for_any", v)
	}

	if v := serviceTcp["override-default-settings"]; v != nil {
		_ = d.Set("override_default_settings", v)
	}

	if v := serviceTcp["port"]; v != nil {
		_ = d.Set("port", v)
	}

	if v := serviceTcp["protocol"]; v != nil {
		_ = d.Set("protocol", v)
	}

	if v := serviceTcp["session-timeout"]; v != nil {
		_ = d.Set("session_timeout", v)
	}

	if v := serviceTcp["source-port"]; v != nil {
		_ = d.Set("source_port", v)
	}

	if v := serviceTcp["sync-connections-on-cluster"]; v != nil {
		_ = d.Set("sync_connections_on_cluster", v)
	}

	if v := serviceTcp["use-default-session-timeout"]; v != nil {
		_ = d.Set("use_default_session_timeout", v)
	}

	if v := serviceTcp["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if serviceTcp["aggressive-aging"] != nil {

		aggressiveAgingMap := serviceTcp["aggressive-aging"].(map[string]interface{})

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
			"default_timeout":     0,
		}
		if reflect.DeepEqual(defaultAggressiveAging, aggressiveAgingMapToReturn) && !aggressiveAgingInConf {
			_ = d.Set("aggressive_aging", map[string]interface{}{})
		} else {
			_ = d.Set("aggressive_aging", aggressiveAgingMapToReturn)
		}
	} else {
		_ = d.Set("aggressive_aging", nil)
	}

	if serviceTcp["groups"] != nil {
		groupsJson := serviceTcp["groups"].([]interface{})
		groupsIds := make([]string, 0)
		if len(groupsJson) > 0 {
			// Create slice of group names
			for _, group_ := range groupsJson {
				group_ := group_.(map[string]interface{})
				groupsIds = append(groupsIds, group_["name"].(string))
			}
		}
		_ = d.Set("groups", groupsIds)
	} else {
		_ = d.Set("groups", nil)
	}

	if serviceTcp["tags"] != nil {
		tagsJson := serviceTcp["tags"].([]interface{})
		var tagsIds = make([]string, 0)
		if len(tagsJson) > 0 {
			// Create slice of tag names
			for _, tag := range tagsJson {
				tag := tag.(map[string]interface{})
				tagsIds = append(tagsIds, tag["name"].(string))
			}
		}
		_ = d.Set("tags", tagsIds)
	} else {
		_ = d.Set("tags", nil)
	}

	return nil
}
