package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"reflect"
)

func resourceManagementServiceUdp() *schema.Resource {
	return &schema.Resource{
		Create: createManagementServiceUdp,
		Read:   readManagementServiceUdp,
		Update: updateManagementServiceUdp,
		Delete: deleteManagementServiceUdp,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name. Should be unique in the domain.",
			},
			"accept_replies": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "N/A",
				Default:     true,
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
			"port": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The number of the port used to provide this service. To specify a port range, place a hyphen between the lowest and highest port numbers, for example 44-55.",
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
				Default:     40,
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

func createManagementServiceUdp(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	serviceUdp := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		serviceUdp["name"] = v.(string)
	}
	if val, ok := d.GetOkExists("accept_replies"); ok {
		serviceUdp["accept-replies"] = val.(bool)
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
		serviceUdp["aggressive-aging"] = res
	}

	if val, ok := d.GetOkExists("keep_connections_open_after_policy_installation"); ok {
		serviceUdp["keep-connections-open-after-policy-installation"] = val.(bool)
	}
	if val, ok := d.GetOkExists("match_by_protocol_signature"); ok {
		serviceUdp["match-by-protocol-signature"] = val.(bool)
	}
	if val, ok := d.GetOkExists("match_for_any"); ok {
		serviceUdp["match-for-any"] = val.(bool)
	}
	if val, ok := d.GetOkExists("override_default_settings"); ok {
		serviceUdp["override-default-settings"] = val.(bool)
	}
	if val, ok := d.GetOk("port"); ok {
		serviceUdp["port"] = val.(string)
	}
	if val, ok := d.GetOk("protocol"); ok {
		serviceUdp["protocol"] = val.(string)
	}
	if val, ok := d.GetOk("session_timeout"); ok {
		serviceUdp["session-timeout"] = val.(int)
	}
	if val, ok := d.GetOk("source_port"); ok {
		serviceUdp["source-port"] = val.(string)
	}
	if val, ok := d.GetOkExists("sync_connections_on_cluster"); ok {
		serviceUdp["sync-connections-on-cluster"] = val.(bool)
	}
	if val, ok := d.GetOkExists("use_default_session_timeout"); ok {
		serviceUdp["use-default-session-timeout"] = val.(bool)
	}

	if val, ok := d.GetOk("tags"); ok {
		serviceUdp["tags"] = val.(*schema.Set).List()
	}
	if val, ok := d.GetOk("comments"); ok {
		serviceUdp["comments"] = val.(string)
	}
	if val, ok := d.GetOk("color"); ok {
		serviceUdp["color"] = val.(string)
	}
	if val, ok := d.GetOkExists("ignore_errors"); ok {
		serviceUdp["ignore-errors"] = val.(bool)
	}
	if val, ok := d.GetOkExists("ignore_warnings"); ok {
		serviceUdp["ignore-warnings"] = val.(bool)
	}

	log.Println("Create Service Udp - Map = ", serviceUdp)

	addServiceUdpRes, err := client.ApiCall("add-service-udp", serviceUdp, client.GetSessionID(), true, false)
	if err != nil || !addServiceUdpRes.Success {
		if addServiceUdpRes.ErrorMsg != "" {
			return fmt.Errorf(addServiceUdpRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addServiceUdpRes.GetData()["uid"].(string))

	return readManagementServiceUdp(d, m)
}

func readManagementServiceUdp(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showServiceUdpRes, err := client.ApiCall("show-service-udp", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showServiceUdpRes.Success {
		// Handle delete resource from other clients
		if objectNotFound(showServiceUdpRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showServiceUdpRes.ErrorMsg)
	}

	serviceUdp := showServiceUdpRes.GetData()

	if v := serviceUdp["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := serviceUdp["accept-replies"]; v != nil {
		_ = d.Set("accept_replies", v)
	}

	if v := serviceUdp["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := serviceUdp["keep-connections-open-after-policy-installation"]; v != nil {
		_ = d.Set("keep_connections_open_after_policy_installation", v)
	}

	if v := serviceUdp["match-by-protocol-signature"]; v != nil {
		_ = d.Set("match_by_protocol_signature", v)
	}

	if v := serviceUdp["match-for-any"]; v != nil {
		_ = d.Set("match_for_any", v)
	}

	if v := serviceUdp["override-default-settings"]; v != nil {
		_ = d.Set("override_default_settings", v)
	}

	if v := serviceUdp["port"]; v != nil {
		_ = d.Set("port", v)
	}

	if v := serviceUdp["protocol"]; v != nil {
		_ = d.Set("protocol", v)
	}

	if v := serviceUdp["session-timeout"]; v != nil {
		_ = d.Set("session_timeout", v)
	}

	if v := serviceUdp["source-port"]; v != nil {
		_ = d.Set("source_port", v)
	}

	if v := serviceUdp["sync-connections-on-cluster"]; v != nil {
		_ = d.Set("sync_connections_on_cluster", v)
	}

	if v := serviceUdp["use-default-session-timeout"]; v != nil {
		_ = d.Set("use_default_session_timeout", v)
	}

	if serviceUdp["aggressive-aging"] != nil {

		aggressiveAgingMap := serviceUdp["aggressive-aging"].(map[string]interface{})

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

	if serviceUdp["tags"] != nil {
		tagsJson := serviceUdp["tags"].([]interface{})
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

func updateManagementServiceUdp(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	serviceUdp := make(map[string]interface{})

	if d.HasChange("name") {
		oldName, newName := d.GetChange("name")
		serviceUdp["name"] = oldName.(string)
		serviceUdp["new-name"] = newName.(string)
	} else {
		serviceUdp["name"] = d.Get("name")
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
			serviceUdp["aggressive-aging"] = res
		} else { //argument deleted - go back to defaults
			defaultAggressiveAging := map[string]interface{}{
				"enable":              true,
				"timeout":             600,
				"use-default-timeout": true,
				"default-timeout":     0,
			}
			serviceUdp["aggressive-aging"] = defaultAggressiveAging
		}
	}

	if ok := d.HasChange("tags"); ok {
		if v, ok := d.GetOk("tags"); ok {
			serviceUdp["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			serviceUdp["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("accept_replies"); ok {
		serviceUdp["accept-replies"] = d.Get("accept_replies")
	}
	if ok := d.HasChange("comments"); ok {
		serviceUdp["comments"] = d.Get("comments")
	}
	if ok := d.HasChange("keep_connections_open_after_policy_installation"); ok {
		serviceUdp["keep-connections-open-after-policy-installation"] = d.Get("keep_connections_open_after_policy_installation")
	}
	if ok := d.HasChange("match_by_protocol_signature"); ok {
		serviceUdp["match-by-protocol-signature"] = d.Get("match_by_protocol_signature")
	}
	if ok := d.HasChange("match_for_any"); ok {
		serviceUdp["match-for-any"] = d.Get("match_for_any")
	}
	if ok := d.HasChange("override_default_settings"); ok {
		serviceUdp["override-default-settings"] = d.Get("override_default_settings")
	}
	if ok := d.HasChange("port"); ok {
		serviceUdp["port"] = d.Get("port")
	}
	if ok := d.HasChange("protocol"); ok {
		serviceUdp["protocol"] = d.Get("protocol")
	}
	if ok := d.HasChange("session_timeout"); ok {
		serviceUdp["session-timeout"] = d.Get("session_timeout")
	}
	if ok := d.HasChange("source_port"); ok {
		serviceUdp["source-port"] = d.Get("source_port")
	}
	if ok := d.HasChange("sync_connections_on_cluster"); ok {
		serviceUdp["sync-connections-on-cluster"] = d.Get("sync_connections_on_cluster")
	}
	if ok := d.HasChange("use_default_session_timeout"); ok {
		serviceUdp["use-default-session-timeout"] = d.Get("use_default_session_timeout")
	}
	if ok := d.HasChange("color"); ok {
		serviceUdp["color"] = d.Get("color")
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		serviceUdp["ignore-errors"] = v.(bool)
	}
	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		serviceUdp["ignore-warnings"] = v.(bool)
	}

	log.Println("Update Service Udp - Map = ", serviceUdp)
	setServiceUdpRes, _ := client.ApiCall("set-service-udp", serviceUdp, client.GetSessionID(), true, false)
	if !setServiceUdpRes.Success {
		return fmt.Errorf(setServiceUdpRes.ErrorMsg)
	}

	return readManagementServiceUdp(d, m)
}

func deleteManagementServiceUdp(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	payload := map[string]interface{}{
		"uid": d.Id(),
	}
	deleteServiceUdpRes, _ := client.ApiCall("delete-service-udp", payload, client.GetSessionID(), true, false)
	if !deleteServiceUdpRes.Success {
		return fmt.Errorf(deleteServiceUdpRes.ErrorMsg)
	}
	d.SetId("")

	return nil
}
