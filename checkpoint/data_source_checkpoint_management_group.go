package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceManagementGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementGroupRead,
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
				Description: "Displays the group's matched content as ranges of IP addresses, in case 'show-as-ranges' is set to true.<br />In this case, the 'members' parameter is...",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"excluded_others": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Objects which are not represented as IP addresses and are negated in the given rule - for example if negate is set for the source or destination of th...",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"ipv4": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Range of IPv4 addresses that match in the given rule.",
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
						"ipv6": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Range of IPv6 addresses that match in the given rule.",
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
						"others": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Objects which are not represented as IP addresses and match the given rule. The details-level parameter of the request determines whether they are dis...",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceManagementGroupRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showGroupRes, err := client.ApiCall("show-group", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	if !showGroupRes.Success {
		return fmt.Errorf("%s", showGroupRes.ErrorMsg)
	}

	group := showGroupRes.GetData()

	if v := group["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := group["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := group["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := group["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if group["members"] != nil {
		membersJson := group["members"].([]interface{})
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

	if group["groups"] != nil {
		groupsJson := group["groups"].([]interface{})
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

	if group["tags"] != nil {
		tagsJson := group["tags"].([]interface{})
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

	if v := group["ranges"]; v != nil {
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
		if v := rangesShow["ipv4"]; v != nil {
			ipv4List := v.([]interface{})
			var ipv4ListState []map[string]interface{}
			for i := range ipv4List {
				ipv4Show := ipv4List[i].(map[string]interface{})
				ipv4State := make(map[string]interface{})
				if v := ipv4Show["end"]; v != nil {
					ipv4State["end"] = v
				}
				if v := ipv4Show["start"]; v != nil {
					ipv4State["start"] = v
				}
				ipv4ListState = append(ipv4ListState, ipv4State)
			}
			rangesState["ipv4"] = ipv4ListState
		}
		if v := rangesShow["ipv6"]; v != nil {
			ipv6List := v.([]interface{})
			var ipv6ListState []map[string]interface{}
			for i := range ipv6List {
				ipv6Show := ipv6List[i].(map[string]interface{})
				ipv6State := make(map[string]interface{})
				if v := ipv6Show["end"]; v != nil {
					ipv6State["end"] = v
				}
				if v := ipv6Show["start"]; v != nil {
					ipv6State["start"] = v
				}
				ipv6ListState = append(ipv6ListState, ipv6State)
			}
			rangesState["ipv6"] = ipv6ListState
		}
		if v := rangesShow["others"]; v != nil {
			othersIdsList := v.([]interface{})
			var othersIds = make([]string, 0)
			for _, item := range othersIdsList {
				othersIds = append(othersIds, item.(map[string]interface{})["name"].(string))
			}
			rangesState["others"] = othersIds
		}
		_ = d.Set("ranges", []interface{}{rangesState})
	}

	return nil
}
