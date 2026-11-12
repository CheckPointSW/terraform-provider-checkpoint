package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func dataSourceManagementPackage() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementPackageRead,
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
			"show_installation_targets": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates whether to calculate and show \"installation-targets\" field in reply.",
			},
			"access": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "True - enables, False - disables access & NAT policies, empty - nothing is changed.",
			},
			"desktop_security": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "True - enables, False - disables Desktop security policy, empty - nothing is changed.",
			},
			"installation_targets": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Which Gateways identified by the name or UID to install the policy on.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"access_layers": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Access policy layers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"threat_layers": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Threat policy layers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"https_inspection_layers": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "HTTPS inspection policy layers.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"inbound_https_layer": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "HTTPS inspection policy inbound layer identified by name or UID.",
						},
						"outbound_https_layer": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "HTTPS inspection policy outbound layer identified by name or UID.",
						},
					},
				},
			},
			"qos": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "True - enables, False - disables QoS policy, empty - nothing is changed.",
			},
			"qos_policy_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "QoS policy type.",
			},
			"threat_prevention": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "True - enables, False - disables Threat policy, empty - nothing is changed.",
			},
			"vpn_traditional_mode": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "True - enables, False - disables VPN traditional mode, empty - nothing is changed.",
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
			"tags": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"autonomous_threat_policy": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "UID of the Autonomous Threat Prevention policy.",
			},
			"https_inspection_policy": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "True - enables, False - disables HTTPS Inspection policy, empty - nothing is changed.",
			},
			"nat_layer": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "UID of the NAT policy.",
			},
			"nat_policy": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "True - enables, False - disables NAT policy, empty - nothing is changed.",
			},
			"sd_wan": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "True - enables, False - disables SD-WAN policy, empty - nothing is changed.",
			},
			"installation_targets_revision": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of installation targets and revisions on which this policy package was installed.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cluster_members_revision": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "If this target is a cluster, this list shows a revision which was installed on each cluster member.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"revision": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The revision installed on this target. Level of details in the output corresponds to the number of details for search. This table shows the level of d...",
									},
									"target_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the installation target.",
									},
									"target_uid": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Installation target unique identifier.",
									},
								},
							},
						},
						"revision": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The revision installed on this target. Level of details in the output corresponds to the number of details for search. This table shows the level of d...",
						},
						"target_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the installation target.",
						},
						"target_uid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Installation target unique identifier.",
						},
					},
				},
			},
			"sd_wan_layer": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "SD-WAN policy layer. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Stan...",
			},
		},
	}
}

func dataSourceManagementPackageRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	if v, ok := d.GetOkExists("show_installation_targets"); ok {
		payload["show-installation-targets"] = v.(bool)
	}

	showPackageRes, err := client.ApiCall("show-package", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	if !showPackageRes.Success {
		return fmt.Errorf("%s", showPackageRes.ErrorMsg)
	}

	_package := showPackageRes.GetData()

	log.Println("Read Package - Show JSON = ", _package)

	if v := _package["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := _package["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := _package["access"]; v != nil {
		_ = d.Set("access", v)
	}

	if v := _package["desktop-security"]; v != nil {
		_ = d.Set("desktop_security", v)
	}

	if v := _package["installation-targets"]; v != nil {

		installationTargetsIds := make([]string, 0)
		if v == "all" {
			installationTargetsIds = append(installationTargetsIds, v.(string))
		} else {
			installationTargetsJson := _package["installation-targets"].([]interface{})
			if len(installationTargetsJson) > 0 {
				for _, installationTarget := range installationTargetsJson {
					installationTarget := installationTarget.(map[string]interface{})
					installationTargetsIds = append(installationTargetsIds, installationTarget["name"].(string))
				}
			}
		}
		_, installationTargetsInConf := d.GetOk("installation_targets")
		if len(installationTargetsIds) == 1 && installationTargetsIds[0] == "all" && !installationTargetsInConf {
			_ = d.Set("installation_targets", []interface{}{})
		} else {
			_ = d.Set("installation_targets", installationTargetsIds)
		}

	} else {
		_ = d.Set("installation_targets", nil)
	}

	if v := _package["access-layers"]; v != nil {
		layersJson := v.([]interface{})
		names := make([]string, 0, len(layersJson))
		for _, layer := range layersJson {
			names = append(names, layer.(map[string]interface{})["name"].(string))
		}
		_ = d.Set("access_layers", names)
	} else {
		_ = d.Set("access_layers", nil)
	}

	if v := _package["threat-layers"]; v != nil {
		layersJson := v.([]interface{})
		names := make([]string, 0, len(layersJson))
		for _, layer := range layersJson {
			names = append(names, layer.(map[string]interface{})["name"].(string))
		}
		_ = d.Set("threat_layers", names)
	} else {
		_ = d.Set("threat_layers", nil)
	}

	if v := _package["https-inspection-layers"]; v != nil {
		raw := v.(map[string]interface{})
		entry := make(map[string]interface{})
		if inb := raw["inbound-https-layer"]; inb != nil {
			entry["inbound_https_layer"] = inb.(map[string]interface{})["name"].(string)
		}
		if out := raw["outbound-https-layer"]; out != nil {
			entry["outbound_https_layer"] = out.(map[string]interface{})["name"].(string)
		}
		_ = d.Set("https_inspection_layers", []map[string]interface{}{entry})
	} else {
		_ = d.Set("https_inspection_layers", nil)
	}

	if v := _package["qos"]; v != nil {
		_ = d.Set("qos", v)
	}

	if v := _package["qos-policy-type"]; v != nil {
		_ = d.Set("qos_policy_type", v)
	}

	if v := _package["threat-prevention"]; v != nil {
		_ = d.Set("threat_prevention", v)
	}

	if v := _package["vpn-traditional-mode"]; v != nil {
		_ = d.Set("vpn_traditional_mode", v)
	}

	if v := _package["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := _package["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if _package["tags"] != nil {
		tagsJson := _package["tags"].([]interface{})
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

	if v := _package["autonomous-threat-policy"]; v != nil {
		_ = d.Set("autonomous_threat_policy", v)
	}

	if v := _package["https-inspection-policy"]; v != nil {
		_ = d.Set("https_inspection_policy", v)
	}

	if v := _package["nat-layer"]; v != nil {
		_ = d.Set("nat_layer", v)
	}

	if v := _package["nat-policy"]; v != nil {
		_ = d.Set("nat_policy", v)
	}

	if v := _package["sd-wan"]; v != nil {
		_ = d.Set("sd_wan", v)
	}

	if v := _package["installation-targets-revision"]; v != nil {
		installationTargetsRevisionList := v.([]interface{})
		var installationTargetsRevisionListState []map[string]interface{}
		for i := range installationTargetsRevisionList {
			installationTargetsRevisionShow := installationTargetsRevisionList[i].(map[string]interface{})
			installationTargetsRevisionState := make(map[string]interface{})
			if v := installationTargetsRevisionShow["cluster-members-revision"]; v != nil {
				clusterMembersRevisionList := v.([]interface{})
				var clusterMembersRevisionListState []map[string]interface{}
				for i := range clusterMembersRevisionList {
					clusterMembersRevisionShow := clusterMembersRevisionList[i].(map[string]interface{})
					clusterMembersRevisionState := make(map[string]interface{})
					if v := clusterMembersRevisionShow["revision"]; v != nil {
						clusterMembersRevisionState["revision"] = v.(map[string]interface{})["name"]
					}
					if v := clusterMembersRevisionShow["target-name"]; v != nil {
						clusterMembersRevisionState["target_name"] = v
					}
					if v := clusterMembersRevisionShow["target-uid"]; v != nil {
						clusterMembersRevisionState["target_uid"] = v
					}
					clusterMembersRevisionListState = append(clusterMembersRevisionListState, clusterMembersRevisionState)
				}
				installationTargetsRevisionState["cluster_members_revision"] = clusterMembersRevisionListState
			}
			if v := installationTargetsRevisionShow["revision"]; v != nil {
				installationTargetsRevisionState["revision"] = v.(map[string]interface{})["name"]
			}
			if v := installationTargetsRevisionShow["target-name"]; v != nil {
				installationTargetsRevisionState["target_name"] = v
			}
			if v := installationTargetsRevisionShow["target-uid"]; v != nil {
				installationTargetsRevisionState["target_uid"] = v
			}
			installationTargetsRevisionListState = append(installationTargetsRevisionListState, installationTargetsRevisionState)
		}
		_ = d.Set("installation_targets_revision", installationTargetsRevisionListState)
	}

	if v := _package["sd-wan-layer"]; v != nil {
		_ = d.Set("sd_wan_layer", v.(map[string]interface{})["name"])
	}

	return nil
}
