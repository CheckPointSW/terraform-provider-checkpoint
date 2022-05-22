package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strconv"
)

func dataSourceManagementClusterMember() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementClusterMemberRead,
		Schema: map[string]*schema.Schema{
			"uid": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Cluster member unique identifier.",
			},
			"limit_interfaces": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Limit number of cluster member interfaces to show.",
				Default:     50,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "N/A",
			},
			"cluster_uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Cluster object (the owner of this member) uid.",
			},
			"interfaces": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Cluster member network interfaces.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv4_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv4 address.",
						},
						"ipv4_mask_length": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "IPv4 network mask length.",
						},
						"ipv4_network_mask": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv4 network mask.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Cluster member interface name.",
						},
						"ipv6_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv6 address.",
						},
						"ipv6_mask_length": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "IPv6 network mask length.",
						},
						"ipv6_network_mask": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv6 network mask.",
						},
						"uid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Cluster member interface object UID.",
						},
					},
				},
			},
			"ip_address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Cluster member IP address.",
			},
			"ipv6_address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Cluster member IPv6 address.",
			},
			"sic_message": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Secure Internal Communication message.",
			},
			"sic_state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Secure Internal Communication state.",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object type.",
			},
			"priority": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "In a High Availability New mode cluster each machine is given a priority. The highest priority machine serves as the gateway in normal circumstances. If this machine fails, control is passed to the next highest priority machine. If that machine fails, control is passed to the next machine, and so on. In Load Sharing Unicast mode cluster, the highest priority is the pivot machine. The values must be in a range from 1 to N, where N is number of cluster members.",
			},
			"nat_settings": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "NAT settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_rule": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether to add automatic address translation rules.",
						},
						"hide_behind": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Hide behind method. This parameter is forbidden in case \"method\" parameter is \"static\".",
						},
						"install_on": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Which gateway should apply the NAT translation.",
						},
						"ipv4_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IPv4 address.",
						},
						"ipv6_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IPv6 address.",
						},
						"method": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "NAT translation method.",
						},
					},
				},
			},
		},
	}
}

func dataSourceManagementClusterMemberRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	uid := d.Get("uid").(string)
	d.SetId(uid)

	payload := make(map[string]interface{})

	payload["uid"] = uid

	if v, ok := d.GetOk("limit_interfaces"); ok {
		payload["limit-interfaces"] = v
	}

	showClusterMemberRes, err := client.ApiCall("show-cluster-member", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showClusterMemberRes.Success {
		if objectNotFound(showClusterMemberRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showClusterMemberRes.ErrorMsg)
	}

	clusterMember := showClusterMemberRes.GetData()

	log.Println("Read ClusterMember - Show JSON = ", clusterMember)

	if v := clusterMember["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := clusterMember["type"]; v != nil {
		_ = d.Set("type", v)
	}

	if v := clusterMember["cluster-uid"]; v != nil {
		_ = d.Set("cluster_uid", v)
	}

	if v := clusterMember["ip-address"]; v != nil {
		_ = d.Set("ip_address", v)
	}

	if v := clusterMember["ipv6-address"]; v != nil {
		_ = d.Set("ipv6_address", v)
	}

	if v := clusterMember["sic-message"]; v != nil {
		_ = d.Set("sic_message", v)
	}

	if v := clusterMember["sic-state"]; v != nil {
		_ = d.Set("sic_state", v)
	}

	if v := clusterMember["priority"]; v != nil {
		_ = d.Set("priority", v.(int))
	}

	if clusterMember["interfaces"] != nil {

		interfacesList, ok := clusterMember["interfaces"].([]interface{})

		var interfacesListToReturn []map[string]interface{}

		if ok {

			if len(interfacesList) > 0 {

				for i := range interfacesList {

					interfacesMap := interfacesList[i].(map[string]interface{})

					interfacesMapToAdd := make(map[string]interface{})

					if v, _ := interfacesMap["ipv4-address"]; v != nil {
						interfacesMapToAdd["ipv4_address"] = v
					}
					if v, _ := interfacesMap["ipv4-mask-length"]; v != nil {
						interfacesMapToAdd["ipv4_mask_length"] = v
					}
					if v, _ := interfacesMap["ipv4-network-mask"]; v != nil {
						interfacesMapToAdd["ipv4_network_mask"] = v
					}
					if v, _ := interfacesMap["name"]; v != nil {
						interfacesMapToAdd["name"] = v
					}
					if v, _ := interfacesMap["ipv6-address"]; v != nil {
						interfacesMapToAdd["ipv^_address"] = v
					}
					if v, _ := interfacesMap["ipv6-mask-length"]; v != nil {
						interfacesMapToAdd["ipv6_mask_length"] = v
					}
					if v, _ := interfacesMap["ipv6-network-mask"]; v != nil {
						interfacesMapToAdd["ipv6_network_mask"] = v
					}
					if v, _ := interfacesMap["uid"]; v != nil {
						interfacesMapToAdd["uid"] = v
					}
					interfacesListToReturn = append(interfacesListToReturn, interfacesMapToAdd)
				}
			}
		}
		_ = d.Set("interfaces", interfacesListToReturn)
	}

	if clusterMember["nat-settings"] != nil {

		actionSettingsMap := clusterMember["nat-settings"].(map[string]interface{})

		actionSettingsMapToReturn := make(map[string]interface{})

		if v, _ := actionSettingsMap["auto-rule"]; v != nil {
			actionSettingsMapToReturn["auto_rule"] = strconv.FormatBool(v.(bool))
		}

		if v, _ := actionSettingsMap["hide-behind"]; v != nil {
			actionSettingsMapToReturn["hide_behind"] = v
		}

		if v, _ := actionSettingsMap["install-on"]; v != nil {
			actionSettingsMapToReturn["install_on"] = v
		}

		if v, _ := actionSettingsMap["ipv4-address"]; v != nil {
			actionSettingsMapToReturn["ipv4_address"] = v
		}

		if v, _ := actionSettingsMap["ipv6-address"]; v != nil {
			actionSettingsMapToReturn["ipv6_address"] = v
		}

		if v, _ := actionSettingsMap["method"]; v != nil {
			actionSettingsMapToReturn["method"] = v
		}

		_ = d.Set("nat_settings", actionSettingsMapToReturn)
	} else {
		_ = d.Set("nat_settings", nil)
	}

	return nil

}
