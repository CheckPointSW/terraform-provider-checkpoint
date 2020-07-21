package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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

	showPackageRes, err := client.ApiCall("show-package", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showPackageRes.Success {
		return fmt.Errorf(showPackageRes.ErrorMsg)
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
		if installationTargetsIds[0] == "all" && !installationTargetsInConf {
			_ = d.Set("installation_targets", []interface{}{})
		} else {
			_ = d.Set("installation_targets", installationTargetsIds)
		}

	} else {
		_ = d.Set("installation_targets", nil)
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

	return nil
}
