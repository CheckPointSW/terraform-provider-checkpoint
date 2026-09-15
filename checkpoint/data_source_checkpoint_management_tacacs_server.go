package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func dataSourceManagementTacacsServer() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementTacacsServerRead,
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
			"encryption": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Is there a secret key defined on the server. Must be set true when \"server-type\" was selected to be \"TACACS+\".",
			},
			"groups": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"priority": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The priority of the TACACS Server in case it is a member of a TACACS Group.",
			},
			"server": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The UID or Name of the host that is the TACACS Server.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object name. Must be unique in the domain.",
						},
						"uid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object unique identifier.",
						},
					},
				},
			},
			"server_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Server type, TACACS or TACACS+.",
			},
			"service": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Server service, only relevant when \"server-type\" is TACACS.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"groups": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"aggressive_aging": {
							Type:        schema.TypeList,
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
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object name. Must be unique in the domain.",
						},
						"uid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object unique identifier.",
						},
						"keep_connections_open_after_policy_installation": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection...",
						},
						"match_for_any": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether this service is used when 'Any' is set as the rule's service and there are several service objects with the same source port and pro...",
						},
						"port": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The number of the port used to provide this service.",
						},
						"session_timeout": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Time (in seconds) before the session times out.",
						},
						"source_port": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Port number for the client side service. If specified, only those Source port Numbers will be Accepted, Dropped, or Rejected during packet inspection....",
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
					},
				},
			},
		},
	}
}

func dataSourceManagementTacacsServerRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showTacacsServerRes, err := client.ApiCall("show-tacacs-server", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	if !showTacacsServerRes.Success {
		return fmt.Errorf("%s", showTacacsServerRes.ErrorMsg)
	}

	tacacsServer := showTacacsServerRes.GetData()

	log.Println("Read Tacacs Server - Show JSON = ", tacacsServer)

	if v := tacacsServer["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := tacacsServer["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := tacacsServer["encryption"]; v != nil {
		_ = d.Set("encryption", v)
	}

	if tacacsServer["groups"] != nil {
		groupsJson := tacacsServer["groups"].([]interface{})
		groupsIds := make([]string, 0)
		if len(groupsJson) > 0 {
			// Create slice of group names
			for _, group := range groupsJson {
				group := group.(map[string]interface{})
				groupsIds = append(groupsIds, group["name"].(string))
			}
		}
		_ = d.Set("groups", groupsIds)
	} else {
		_ = d.Set("groups", nil)
	}

	if v := tacacsServer["priority"]; v != nil {
		_ = d.Set("priority", v)
	}

	if tacacsServer["server"] != nil {

		serverMap := tacacsServer["server"].(map[string]interface{})

		serverMapToReturn := make(map[string]interface{})

		if v := serverMap["name"]; v != nil {
			serverMapToReturn["name"] = v
		}
		if v := serverMap["uid"]; v != nil {
			serverMapToReturn["uid"] = v
		}

		_ = d.Set("server", []interface{}{serverMapToReturn})

	} else {
		_ = d.Set("server", nil)
	}

	if v := tacacsServer["server-type"]; v != nil {
		_ = d.Set("server_type", v)
	}

	if tacacsServer["service"] != nil {

		serviceMap := tacacsServer["service"].(map[string]interface{})

		serviceMapToReturn := make(map[string]interface{})

		if v := serviceMap["aggressive-aging"]; v != nil {
			aggressiveAgingShow := v.(map[string]interface{})
			aggressiveAgingState := make(map[string]interface{})
			if v := aggressiveAgingShow["default-timeout"]; v != nil {
				aggressiveAgingState["default_timeout"] = v
			}
			if v := aggressiveAgingShow["enable"]; v != nil {
				aggressiveAgingState["enable"] = v
			}
			if v := aggressiveAgingShow["timeout"]; v != nil {
				aggressiveAgingState["timeout"] = v
			}
			if v := aggressiveAgingShow["use-default-timeout"]; v != nil {
				aggressiveAgingState["use_default_timeout"] = v
			}
			serviceMapToReturn["aggressive_aging"] = []interface{}{aggressiveAgingState}
		}
		if v := serviceMap["name"]; v != nil {
			serviceMapToReturn["name"] = v
		}
		if v := serviceMap["uid"]; v != nil {
			serviceMapToReturn["uid"] = v
		}

		if v := serviceMap["keep-connections-open-after-policy-installation"]; v != nil {
			serviceMapToReturn["keep_connections_open_after_policy_installation"] = v
		}
		if v := serviceMap["match-for-any"]; v != nil {
			serviceMapToReturn["match_for_any"] = v
		}
		if v := serviceMap["port"]; v != nil {
			serviceMapToReturn["port"] = v
		}
		if v := serviceMap["session-timeout"]; v != nil {
			serviceMapToReturn["session_timeout"] = v
		}
		if v := serviceMap["source-port"]; v != nil {
			serviceMapToReturn["source_port"] = v
		}
		if v := serviceMap["sync-connections-on-cluster"]; v != nil {
			serviceMapToReturn["sync_connections_on_cluster"] = v
		}
		if v := serviceMap["use-default-session-timeout"]; v != nil {
			serviceMapToReturn["use_default_session_timeout"] = v
		}
		if v := serviceMap["groups"]; v != nil {
			groupsList := v.([]interface{})
			groupsIds := make([]string, 0)
			for _, item := range groupsList {
				groupsIds = append(groupsIds, item.(map[string]interface{})["name"].(string))
			}
			serviceMapToReturn["groups"] = groupsIds
		}
		_ = d.Set("service", []interface{}{serviceMapToReturn})

	} else {
		_ = d.Set("service", nil)
	}

	return nil
}
