package checkpoint

import (
	"github.com/CheckPointSW/terraform-provider-checkpoint/v3/upgraders"
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceManagementServiceSctp() *schema.Resource {
	return &schema.Resource{
		Create: createManagementServiceSctp,
		Read:   readManagementServiceSctp,
		Update: updateManagementServiceSctp,
		Delete: deleteManagementServiceSctp,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    upgraders.ResourceManagementServiceSctpV0().CoreConfigSchema().ImpliedType(),
				Upgrade: upgraders.ResourceManagementServiceSctpStateUpgradeV0,
				Version: 0,
			},
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"aggressive_aging": {
				Type:        schema.TypeList,
				MaxItems:    1,
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
							Default:     600,
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
			"keep_connections_open_after_policy_installation": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections.",
			},
			"match_for_any": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates whether this service is used when 'Any' is set as the rule's service and there are several service objects with the same source port and protocol.",
			},
			"port": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Port number. To specify a port range add a hyphen between the lowest and the highest port numbers, for example 44-45.",
			},
			"session_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Time (in seconds) before the session times out.",
			},
			"source_port": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Source port number. To specify a port range add a hyphen between the lowest and the highest port numbers, for example 44-45.",
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

func createManagementServiceSctp(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	serviceSctp := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		serviceSctp["name"] = v.(string)
	}

	if v, ok := d.GetOk("aggressive_aging"); ok {

		aggressiveAgingList := v.([]interface{})

		if len(aggressiveAgingList) > 0 {

			aggressiveAgingPayload := make(map[string]interface{})

			if v, ok := d.GetOk("aggressive_aging.0.default_timeout"); ok {
				aggressiveAgingPayload["default-timeout"] = v.(int)
			}
			if v, ok := d.GetOkExists("aggressive_aging.0.enable"); ok {
				aggressiveAgingPayload["enable"] = v.(bool)
			}
			if v, ok := d.GetOk("aggressive_aging.0.timeout"); ok {
				aggressiveAgingPayload["timeout"] = v.(int)
			}
			if v, ok := d.GetOkExists("aggressive_aging.0.use_default_timeout"); ok {
				aggressiveAgingPayload["use-default-timeout"] = v.(bool)
			}
			serviceSctp["aggressive-aging"] = aggressiveAgingPayload
		}
	}

	if v, ok := d.GetOkExists("keep_connections_open_after_policy_installation"); ok {
		serviceSctp["keep-connections-open-after-policy-installation"] = v.(bool)
	}

	if v, ok := d.GetOkExists("match_for_any"); ok {
		serviceSctp["match-for-any"] = v.(bool)
	}

	if v, ok := d.GetOk("port"); ok {
		serviceSctp["port"] = v.(string)
	}

	if v, ok := d.GetOk("session_timeout"); ok {
		serviceSctp["session-timeout"] = v.(int)
	}

	if v, ok := d.GetOk("source_port"); ok {
		serviceSctp["source-port"] = v.(string)
	}

	if v, ok := d.GetOkExists("sync_connections_on_cluster"); ok {
		serviceSctp["sync-connections-on-cluster"] = v.(bool)
	}

	if v, ok := d.GetOk("tags"); ok {
		serviceSctp["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOkExists("use_default_session_timeout"); ok {
		serviceSctp["use-default-session-timeout"] = v.(bool)
	}

	if v, ok := d.GetOk("color"); ok {
		serviceSctp["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		serviceSctp["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		serviceSctp["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		serviceSctp["ignore-errors"] = v.(bool)
	}

	log.Println("Create ServiceSctp - Map = ", serviceSctp)

	addServiceSctpRes, err := client.ApiCall("add-service-sctp", serviceSctp, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !addServiceSctpRes.Success {
		if addServiceSctpRes.ErrorMsg != "" {
			return fmt.Errorf("%s", addServiceSctpRes.ErrorMsg)
		}
		return fmt.Errorf("%s", err.Error())
	}

	d.SetId(addServiceSctpRes.GetData()["uid"].(string))

	return readManagementServiceSctp(d, m)
}

func readManagementServiceSctp(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showServiceSctpRes, err := client.ApiCall("show-service-sctp", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	if !showServiceSctpRes.Success {
		if objectNotFound(showServiceSctpRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("%s", showServiceSctpRes.ErrorMsg)
	}

	serviceSctp := showServiceSctpRes.GetData()

	log.Println("Read ServiceSctp - Show JSON = ", serviceSctp)

	if v := serviceSctp["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if serviceSctp["aggressive-aging"] != nil {

		aggressiveAgingMap := serviceSctp["aggressive-aging"].(map[string]interface{})

		aggressiveAgingMapToReturn := make(map[string]interface{})

		if v := aggressiveAgingMap["default-timeout"]; v != nil {
			aggressiveAgingMapToReturn["default_timeout"] = v
		}
		if v := aggressiveAgingMap["enable"]; v != nil {
			aggressiveAgingMapToReturn["enable"] = v
		}
		if v := aggressiveAgingMap["timeout"]; v != nil {
			aggressiveAgingMapToReturn["timeout"] = v
		}
		if v := aggressiveAgingMap["use-default-timeout"]; v != nil {
			aggressiveAgingMapToReturn["use_default_timeout"] = v
		}

		_ = d.Set("aggressive_aging", []interface{}{aggressiveAgingMapToReturn})

	} else {
		_ = d.Set("aggressive_aging", nil)
	}

	if v := serviceSctp["keep-connections-open-after-policy-installation"]; v != nil {
		_ = d.Set("keep_connections_open_after_policy_installation", v)
	}

	if v := serviceSctp["match-for-any"]; v != nil {
		_ = d.Set("match_for_any", v)
	}

	if v := serviceSctp["port"]; v != nil {
		_ = d.Set("port", v)
	}

	if v := serviceSctp["session-timeout"]; v != nil {
		_ = d.Set("session_timeout", v)
	}

	if v := serviceSctp["source-port"]; v != nil {
		_ = d.Set("source_port", v)
	}

	if v := serviceSctp["sync-connections-on-cluster"]; v != nil {
		_ = d.Set("sync_connections_on_cluster", v)
	}

	if serviceSctp["tags"] != nil {
		tagsJson, ok := serviceSctp["tags"].([]interface{})
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

	if v := serviceSctp["use-default-session-timeout"]; v != nil {
		_ = d.Set("use_default_session_timeout", v)
	}

	if v := serviceSctp["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := serviceSctp["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := serviceSctp["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := serviceSctp["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementServiceSctp(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	serviceSctp := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		serviceSctp["name"] = oldName
		serviceSctp["new-name"] = newName
	} else {
		serviceSctp["name"] = d.Get("name")
	}

	if d.HasChange("aggressive_aging") {

		if v, ok := d.GetOk("aggressive_aging"); ok {

			aggressiveAgingList := v.([]interface{})

			if len(aggressiveAgingList) > 0 {

				aggressiveAgingPayload := make(map[string]interface{})

				if v, ok := d.GetOk("aggressive_aging.0.default_timeout"); ok {
					aggressiveAgingPayload["default-timeout"] = v.(int)
				}
				if v, ok := d.GetOkExists("aggressive_aging.0.enable"); ok {
					aggressiveAgingPayload["enable"] = v.(bool)
				}
				if v, ok := d.GetOk("aggressive_aging.0.timeout"); ok {
					aggressiveAgingPayload["timeout"] = v.(int)
				}
				if v, ok := d.GetOkExists("aggressive_aging.0.use_default_timeout"); ok {
					aggressiveAgingPayload["use-default-timeout"] = v.(bool)
				}
				serviceSctp["aggressive-aging"] = aggressiveAgingPayload
			}
		}
	}

	if v, ok := d.GetOkExists("keep_connections_open_after_policy_installation"); ok {
		serviceSctp["keep-connections-open-after-policy-installation"] = v.(bool)
	}

	if v, ok := d.GetOkExists("match_for_any"); ok {
		serviceSctp["match-for-any"] = v.(bool)
	}

	if ok := d.HasChange("port"); ok {
		serviceSctp["port"] = d.Get("port")
	}

	if ok := d.HasChange("session_timeout"); ok {
		serviceSctp["session-timeout"] = d.Get("session_timeout")
	}

	if ok := d.HasChange("source_port"); ok {
		serviceSctp["source-port"] = d.Get("source_port")
	}

	if v, ok := d.GetOkExists("sync_connections_on_cluster"); ok {
		serviceSctp["sync-connections-on-cluster"] = v.(bool)
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			serviceSctp["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			serviceSctp["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if v, ok := d.GetOkExists("use_default_session_timeout"); ok {
		serviceSctp["use-default-session-timeout"] = v.(bool)
	}

	if ok := d.HasChange("color"); ok {
		serviceSctp["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		serviceSctp["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		serviceSctp["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		serviceSctp["ignore-errors"] = v.(bool)
	}

	log.Println("Update ServiceSctp - Map = ", serviceSctp)

	updateServiceSctpRes, err := client.ApiCall("set-service-sctp", serviceSctp, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !updateServiceSctpRes.Success {
		if updateServiceSctpRes.ErrorMsg != "" {
			return fmt.Errorf("%s", updateServiceSctpRes.ErrorMsg)
		}
		return fmt.Errorf("%s", err.Error())
	}

	return readManagementServiceSctp(d, m)
}

func deleteManagementServiceSctp(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	serviceSctpPayload := map[string]interface{}{
		"uid": d.Id(),
	}
	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		serviceSctpPayload["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		serviceSctpPayload["ignore-errors"] = v.(bool)
	}
	log.Println("Delete ServiceSctp")

	deleteServiceSctpRes, err := client.ApiCall("delete-service-sctp", serviceSctpPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteServiceSctpRes.Success {
		if deleteServiceSctpRes.ErrorMsg != "" {
			return fmt.Errorf("%s", deleteServiceSctpRes.ErrorMsg)
		}
		return fmt.Errorf("%s", err.Error())
	}
	d.SetId("")

	return nil
}
