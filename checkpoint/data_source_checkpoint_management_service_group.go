package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceManagementServiceGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementServiceGroupRead,
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
			"members": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of Network objects identified by the name or UID.",
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
			"groups": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of group name.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ranges": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Displays the service group's matched content as ranges of port numbers, in case 'show-as-ranges' is set to true.<br />In this case, the 'members' para...",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"excluded_others": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Objects which are not represented as port numbers and are negated in the given rule - for example if negate is set for the service of this rule. The d...",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"others": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Objects which are not represented as port numbers and match the given rule. The details-level parameter of the request determines whether they are dis...",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"tcp": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Range of TCP ports that match in the given rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"end": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "N/A",
									},
									"start": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "N/A",
									},
								},
							},
						},
						"udp": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Range of UDP ports that match in the given rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"end": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "N/A",
									},
									"start": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "N/A",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceManagementServiceGroupRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showServiceGroupRes, err := client.ApiCall("show-service-group", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	if !showServiceGroupRes.Success {
		// Handle delete resource from other clients
		if objectNotFound(showServiceGroupRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("%s", showServiceGroupRes.ErrorMsg)
	}

	serviceGroup := showServiceGroupRes.GetData()

	if v := serviceGroup["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := serviceGroup["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := serviceGroup["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := serviceGroup["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if serviceGroup["members"] != nil {
		membersJson := serviceGroup["members"].([]interface{})
		membersIds := make([]string, 0)
		if len(membersJson) > 0 {
			// Create slice of members names
			for _, member := range membersJson {
				member := member.(map[string]interface{})
				membersIds = append(membersIds, member["name"].(string))
			}
		}
		_ = d.Set("members", membersIds)
	} else {
		_ = d.Set("members", nil)
	}

	if serviceGroup["groups"] != nil {
		groupsJson := serviceGroup["groups"].([]interface{})
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

	if serviceGroup["tags"] != nil {
		tagsJson := serviceGroup["tags"].([]interface{})
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

	if v := serviceGroup["ranges"]; v != nil {
		rangesShow := v.(map[string]interface{})
		rangesState := make(map[string]interface{})
		if v := rangesShow["excluded-others"]; v != nil {
			excludedOthersIdsList := v.([]interface{})
			var excludedOthersIds = make([]string, 0)
			for _, item := range excludedOthersIdsList {
				excludedOthersIds = append(excludedOthersIds, item.(map[string]interface{})["name"].(string))
			}
			rangesState["excluded_others"] = excludedOthersIds
		}
		if v := rangesShow["others"]; v != nil {
			othersIdsList := v.([]interface{})
			var othersIds = make([]string, 0)
			for _, item := range othersIdsList {
				othersIds = append(othersIds, item.(map[string]interface{})["name"].(string))
			}
			rangesState["others"] = othersIds
		}
		if v := rangesShow["tcp"]; v != nil {
			tcpList := v.([]interface{})
			var tcpListState []map[string]interface{}
			for i := range tcpList {
				tcpShow := tcpList[i].(map[string]interface{})
				tcpState := make(map[string]interface{})
				if v := tcpShow["end"]; v != nil {
					tcpState["end"] = v
				}
				if v := tcpShow["start"]; v != nil {
					tcpState["start"] = v
				}
				tcpListState = append(tcpListState, tcpState)
			}
			rangesState["tcp"] = tcpListState
		}
		if v := rangesShow["udp"]; v != nil {
			udpList := v.([]interface{})
			var udpListState []map[string]interface{}
			for i := range udpList {
				udpShow := udpList[i].(map[string]interface{})
				udpState := make(map[string]interface{})
				if v := udpShow["end"]; v != nil {
					udpState["end"] = v
				}
				if v := udpShow["start"]; v != nil {
					udpState["start"] = v
				}
				udpListState = append(udpListState, udpState)
			}
			rangesState["udp"] = udpListState
		}
		_ = d.Set("ranges", []interface{}{rangesState})
	}

	return nil
}
