package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"reflect"
)

func resourceManagementServiceTcp() *schema.Resource {
	return &schema.Resource{
		Create: createManagementServiceTcp,
		Read:   readManagementServiceTcp,
		Update: updateManagementServiceTcp,
		Delete: deleteManagementServiceTcp,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name. Should be unique in the domain.",
			},
			"port": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The number of the port used to provide this service. To specify a port range, place a hyphen between the lowest and highest port numbers, for example 44-55.",
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
						},
						"enable": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "N/A",
						},
						"timeout": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Aggressive aging timeout in seconds.",
						},
						"use_default_timeout": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "N/A",
						},
					},
				},
			},
			"keep_connections_open_after_policy_installation": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections.",
			},
			"match_by_protocol_signature": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "A value of true enables matching by the selected protocol's signature - the signature identifies the protocol as genuine. Select this option to limit the port to the specified protocol. If the selected protocol does not support matching by signature, this field cannot be set to true.",
				Default:     false,
			},
			"match_for_any": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates whether this service is used when 'Any' is set as the rule's service and there are several service objects with the same source port and protocol.",
				Default:     true,
			},
			"override_default_settings": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates whether this service is a Data Domain service which has been overridden.",
			},
			"protocol": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Select the protocol type associated with the service, and by implication, the management server (if any) that enforces Content Security and Authentication for the service. Selecting a Protocol Type invokes the specific protocol handlers for each protocol type, thus enabling higher level of security by parsing the protocol, and higher level of connectivity by tracking dynamic actions (such as opening of ports).",
			},
			"session_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Time (in seconds) before the session times out.",
				Default:     3600,
			},
			"source_port": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Port number for the client side service. If specified, only those Source port Numbers will be Accepted, Dropped, or Rejected during packet inspection. Otherwise, the source port is not inspected.",
			},
			"sync_connections_on_cluster": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enables state-synchronized High Availability or Load Sharing on a ClusterXL or OPSEC-certified cluster.",
				Default:     true,
			},
			"use_default_session_timeout": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Use default virtual session timeout.",
				Default:     true,
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
			"comments": &schema.Schema{
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

func createManagementServiceTcp(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	serviceTcp := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		serviceTcp["name"] = v.(string)
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
		serviceTcp["aggressive-aging"] = res
	}

	if val, ok := d.GetOkExists("keep_connections_open_after_policy_installation"); ok {
		serviceTcp["keep-connections-open-after-policy-installation"] = val.(bool)
	}
	if val, ok := d.GetOkExists("match_by_protocol_signature"); ok {
		serviceTcp["match-by-protocol-signature"] = val.(bool)
	}
	if val, ok := d.GetOkExists("match_for_any"); ok {
		serviceTcp["match-for-any"] = val.(bool)
	}
	if val, ok := d.GetOkExists("override_default_settings"); ok {
		serviceTcp["override-default-settings"] = val.(bool)
	}
	if val, ok := d.GetOk("port"); ok {
		serviceTcp["port"] = val.(string)
	}
	if val, ok := d.GetOk("protocol"); ok {
		serviceTcp["protocol"] = val.(string)
	}
	if val, ok := d.GetOk("session_timeout"); ok {
		serviceTcp["session-timeout"] = val.(int)
	}
	if val, ok := d.GetOk("source_port"); ok {
		serviceTcp["source-port"] = val.(string)
	}
	if val, ok := d.GetOkExists("sync_connections_on_cluster"); ok {
		serviceTcp["sync-connections-on-cluster"] = val.(bool)
	}
	if val, ok := d.GetOkExists("use_default_session_timeout"); ok {
		serviceTcp["use-default-session-timeout"] = val.(bool)
	}

	if val, ok := d.GetOk("tags"); ok {
		serviceTcp["tags"] = val.(*schema.Set).List()
	}
	if val, ok := d.GetOk("comments"); ok {
		serviceTcp["comments"] = val.(string)
	}
	if val, ok := d.GetOk("color"); ok {
		serviceTcp["color"] = val.(string)
	}
	if val, ok := d.GetOkExists("ignore_errors"); ok {
		serviceTcp["ignore-errors"] = val.(bool)
	}
	if val, ok := d.GetOkExists("ignore_warnings"); ok {
		serviceTcp["ignore-warnings"] = val.(bool)
	}

	log.Println("Create Service Tcp - Map = ", serviceTcp)

	addServiceTcpRes, err := client.ApiCall("add-service-tcp", serviceTcp, client.GetSessionID(), true, false)
	if err != nil || !addServiceTcpRes.Success {
		if addServiceTcpRes.ErrorMsg != "" {
			return fmt.Errorf(addServiceTcpRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addServiceTcpRes.GetData()["uid"].(string))

	return readManagementServiceTcp(d, m)
}

func readManagementServiceTcp(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showServiceTcpRes, err := client.ApiCall("show-service-tcp", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showServiceTcpRes.Success {
		// Handle delete resource from other clients
		if objectNotFound(showServiceTcpRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showServiceTcpRes.ErrorMsg)
	}

	serviceTcp := showServiceTcpRes.GetData()

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

func updateManagementServiceTcp(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	serviceTcp := make(map[string]interface{})

	if d.HasChange("name") {
		oldName, newName := d.GetChange("name")
		serviceTcp["name"] = oldName.(string)
		serviceTcp["new-name"] = newName.(string)
	} else {
		serviceTcp["name"] = d.Get("name")
	}

	if ok := d.HasChange("aggressive_aging"); ok {

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
			serviceTcp["aggressive-aging"] = res
		} else { //argument deleted - go back to defaults
			defaultAggressiveAging := map[string]interface{}{
				"enable":              true,
				"timeout":             600,
				"use-default-timeout": true,
				"default-timeout":     0,
			}
			serviceTcp["aggressive-aging"] = defaultAggressiveAging
		}
	}

	if ok := d.HasChange("tags"); ok {
		if v, ok := d.GetOk("tags"); ok {
			serviceTcp["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			serviceTcp["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("comments"); ok {
		serviceTcp["comments"] = d.Get("comments")
	}
	if ok := d.HasChange("keep_connections_open_after_policy_installation"); ok {
		serviceTcp["keep-connections-open-after-policy-installation"] = d.Get("keep_connections_open_after_policy_installation")
	}
	if ok := d.HasChange("match_by_protocol_signature"); ok {
		serviceTcp["match-by-protocol-signature"] = d.Get("match_by_protocol_signature")
	}
	if ok := d.HasChange("match_for_any"); ok {
		serviceTcp["match-for-any"] = d.Get("match_for_any")
	}
	if ok := d.HasChange("override_default_settings"); ok {
		serviceTcp["override-default-settings"] = d.Get("override_default_settings")
	}
	if ok := d.HasChange("port"); ok {
		serviceTcp["port"] = d.Get("port")
	}
	if ok := d.HasChange("protocol"); ok {
		serviceTcp["protocol"] = d.Get("protocol")
	}
	if ok := d.HasChange("session_timeout"); ok {
		serviceTcp["session-timeout"] = d.Get("session_timeout")
	}
	if ok := d.HasChange("source_port"); ok {
		serviceTcp["source-port"] = d.Get("source_port")
	}
	if ok := d.HasChange("sync_connections_on_cluster"); ok {
		serviceTcp["sync-connections-on-cluster"] = d.Get("sync_connections_on_cluster")
	}
	if ok := d.HasChange("use_default_session_timeout"); ok {
		serviceTcp["use-default-session-timeout"] = d.Get("use_default_session_timeout")
	}
	if ok := d.HasChange("color"); ok {
		serviceTcp["color"] = d.Get("color")
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		serviceTcp["ignore-errors"] = v.(bool)
	}
	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		serviceTcp["ignore-warnings"] = v.(bool)
	}

	log.Println("Update Service Tcp - Map = ", serviceTcp)
	setServiceTcpRes, _ := client.ApiCall("set-service-tcp", serviceTcp, client.GetSessionID(), true, false)
	if !setServiceTcpRes.Success {
		return fmt.Errorf(setServiceTcpRes.ErrorMsg)
	}

	return readManagementServiceTcp(d, m)
}

func deleteManagementServiceTcp(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	payload := map[string]interface{}{
		"uid": d.Id(),
	}
	deleteServiceTcpRes, _ := client.ApiCall("delete-service-tcp", payload, client.GetSessionID(), true, false)
	if !deleteServiceTcpRes.Success {
		return fmt.Errorf(deleteServiceTcpRes.ErrorMsg)
	}
	d.SetId("")

	return nil
}
