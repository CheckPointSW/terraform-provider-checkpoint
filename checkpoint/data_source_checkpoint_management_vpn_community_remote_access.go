package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"reflect"
)

func dataSourceManagementVpnCommunityRemoteAccess() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementVpnCommunityRemoteAccessRead,
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
			"gateways": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of Gateway objects identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"user_groups": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of User group objects identified by the name or UID.",
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

func dataSourceManagementVpnCommunityRemoteAccessRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showVpnCommunityRemoteAccessRes, err := client.ApiCall("show-vpn-community-remote-access", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showVpnCommunityRemoteAccessRes.Success {
		return fmt.Errorf(showVpnCommunityRemoteAccessRes.ErrorMsg)
	}

	vpnCommunityRemoteAccess := showVpnCommunityRemoteAccessRes.GetData()

	log.Println("Read VpnCommunityRemoteAccess - Show JSON = ", vpnCommunityRemoteAccess)

	if v := vpnCommunityRemoteAccess["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := vpnCommunityRemoteAccess["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if vpnCommunityRemoteAccess["gateways"] != nil {
		gatewaysJson, ok := vpnCommunityRemoteAccess["gateways"].([]interface{})
		if ok {
			gwIds := make([]string, 0)
			if len(gatewaysJson) > 0 {
				for _, gw := range gatewaysJson {
					gwIds = append(gwIds, gw.(map[string]interface{})["name"].(string))
				}
			}
			_ = d.Set("gateways", gwIds)
		}
	} else {
		_ = d.Set("gateways", nil)
	}

	if vpnCommunityRemoteAccess["user-groups"] != nil {
		userGroupsJson, ok := vpnCommunityRemoteAccess["user-groups"].([]interface{})
		userGroupIds := make([]string, 0)
		if ok {
			if len(userGroupsJson) > 0 {
				for _, userGroup := range userGroupsJson {
					userGroupIds = append(userGroupIds, userGroup.(map[string]interface{})["name"].(string))
				}
			}
		}
		_, userGroupsInConf := d.GetOk("user_groups")
		defaultUserGroups := []string{"All Users"}
		if reflect.DeepEqual(defaultUserGroups, userGroupIds) && !userGroupsInConf {
			_ = d.Set("user_groups", []string{})
		} else {
			_ = d.Set("user_groups", userGroupIds)
		}
	} else {
		_ = d.Set("user_groups", nil)
	}

	if vpnCommunityRemoteAccess["tags"] != nil {
		tagsJson, ok := vpnCommunityRemoteAccess["tags"].([]interface{})
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

	if v := vpnCommunityRemoteAccess["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := vpnCommunityRemoteAccess["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil
}
