package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"reflect"
)

func resourceManagementVpnCommunityRemoteAccess() *schema.Resource {
	return &schema.Resource{
		Create: createManagementVpnCommunityRemoteAccess,
		Read:   readManagementVpnCommunityRemoteAccess,
		Update: updateManagementVpnCommunityRemoteAccess,
		Delete: deleteManagementVpnCommunityRemoteAccess,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"gateways": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of Gateway objects identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"user_groups": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of User group objects identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
			"comments": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comments string.",
			},
			"ignore_warnings": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Apply changes ignoring warnings.",
			},
			"ignore_errors": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.",
			},
		},
	}
}

func createManagementVpnCommunityRemoteAccess(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}

	if v, ok := d.GetOk("name"); ok {
		payload["name"] = v.(string)
	}

	if v, ok := d.GetOk("gateways"); ok {
		payload["gateways"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("user_groups"); ok {
		payload["user-groups"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("tags"); ok {
		payload["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		payload["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		payload["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		payload["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		payload["ignore-errors"] = v.(bool)
	}

	SetVpnCommunityRemoteAccessRes, _ := client.ApiCall("set-vpn-community-remote-access", payload, client.GetSessionID(), true, false)
	if !SetVpnCommunityRemoteAccessRes.Success {
		return fmt.Errorf(SetVpnCommunityRemoteAccessRes.ErrorMsg)
	}

	d.SetId(SetVpnCommunityRemoteAccessRes.GetData()["uid"].(string))

	return readManagementVpnCommunityRemoteAccess(d, m)
}

func updateManagementVpnCommunityRemoteAccess(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		payload["name"] = oldName
		payload["new-name"] = newName
	} else {
		payload["name"] = d.Get("name")
	}

	if ok := d.HasChange("gateways"); ok {
		if v, ok := d.GetOk("gateways"); ok {
			payload["gateways"] = v.(*schema.Set).List()
		} else {
			oldGateways, _ := d.GetChange("gateways")
			payload["gateways"] = map[string]interface{}{"remove": oldGateways.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("user_groups"); ok {
		if v, ok := d.GetOk("user_groups"); ok {
			payload["user-groups"] = v.(*schema.Set).List()
		} else {
			oldUserGroups, _ := d.GetChange("gateways")
			payload["user-groups"] = map[string]interface{}{"remove": oldUserGroups.(*schema.Set).List()}
		}
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			payload["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			payload["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		payload["color"] = d.Get("color").(string)
	}

	if ok := d.HasChange("comments"); ok {
		payload["comments"] = d.Get("comments").(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		payload["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		payload["ignore-errors"] = v.(bool)
	}

	SetVpnCommunityRemoteAccessRes, _ := client.ApiCall("set-vpn-community-remote-access", payload, client.GetSessionID(), true, false)
	if !SetVpnCommunityRemoteAccessRes.Success {
		return fmt.Errorf(SetVpnCommunityRemoteAccessRes.ErrorMsg)
	}

	return readManagementVpnCommunityRemoteAccess(d, m)
}

func readManagementVpnCommunityRemoteAccess(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showVpnCommunityRemoteAccessRes, err := client.ApiCall("show-vpn-community-remote-access", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showVpnCommunityRemoteAccessRes.Success {
		if objectNotFound(showVpnCommunityRemoteAccessRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showVpnCommunityRemoteAccessRes.ErrorMsg)
	}

	vpnCommunityRemoteAccess := showVpnCommunityRemoteAccessRes.GetData()

	log.Println("Read VpnCommunityRemoteAccess - Show JSON = ", vpnCommunityRemoteAccess)

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

func deleteManagementVpnCommunityRemoteAccess(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
