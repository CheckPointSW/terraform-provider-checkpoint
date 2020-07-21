package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementPackage() *schema.Resource {
	return &schema.Resource{
		Create: createManagementPackage,
		Read:   readManagementPackage,
		Update: updateManagementPackage,
		Delete: deleteManagementPackage,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name. Should be unique in the domain.",
			},
			"access": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "True - enables, False - disables access & NAT policies, empty - nothing is changed.",
				Default:     true,
			},
			"desktop_security": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "True - enables, False - disables Desktop security policy, empty - nothing is changed.",
				Default:     false,
			},
			"installation_targets": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Which Gateways identified by the name or UID to install the policy on.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"qos": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "True - enables, False - disables QoS policy, empty - nothing is changed.",
				Default:     false,
			},
			"qos_policy_type": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "QoS policy type.",
				Default:     "recommended",
			},
			"threat_prevention": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "True - enables, False - disables Threat policy, empty - nothing is changed.",
				Default:     true,
			},
			"vpn_traditional_mode": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "True - enables, False - disables VPN traditional mode, empty - nothing is changed.",
				Default:     false,
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
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func createManagementPackage(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	_package := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		_package["name"] = v.(string)
	}
	if v, ok := d.GetOkExists("access"); ok {
		_package["access"] = v.(bool)
	}
	if v, ok := d.GetOkExists("desktop_security"); ok {
		_package["desktop-security"] = v.(bool)
	}
	if v, ok := d.GetOk("installation_targets"); ok {
		_package["installation-targets"] = v.(*schema.Set).List()
	}
	if v, ok := d.GetOkExists("qos"); ok {
		_package["qos"] = v.(bool)
	}
	if v, ok := d.GetOk("qos_policy_type"); ok {
		_package["qos-policy-type"] = v.(string)
	}
	if v, ok := d.GetOkExists("threat_prevention"); ok {
		_package["threat-prevention"] = v.(bool)
	}
	if v, ok := d.GetOkExists("vpn_traditional_mode"); ok {
		_package["vpn-traditional-mode"] = v.(bool)
	}
	if val, ok := d.GetOk("comments"); ok {
		_package["comments"] = val.(string)
	}
	if val, ok := d.GetOk("tags"); ok {
		_package["tags"] = val.(*schema.Set).List()
	}

	if val, ok := d.GetOk("color"); ok {
		_package["color"] = val.(string)
	}
	if val, ok := d.GetOkExists("ignore_errors"); ok {
		_package["ignore-errors"] = val.(bool)
	}
	if val, ok := d.GetOkExists("ignore_warnings"); ok {
		_package["ignore-warnings"] = val.(bool)
	}

	log.Println("Create Package - Map = ", _package)

	addPackageRes, err := client.ApiCall("add-package", _package, client.GetSessionID(), true, false)
	if err != nil || !addPackageRes.Success {
		if addPackageRes.ErrorMsg != "" {
			return fmt.Errorf(addPackageRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addPackageRes.GetData()["uid"].(string))

	return readManagementPackage(d, m)
}

func readManagementPackage(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showPackageRes, err := client.ApiCall("show-package", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showPackageRes.Success {
		// Handle delete resource from other clients
		if objectNotFound(showPackageRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showPackageRes.ErrorMsg)
	}

	_package := showPackageRes.GetData()

	log.Println("Read Package - Show JSON = ", _package)

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

func updateManagementPackage(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	_package := make(map[string]interface{})

	if d.HasChange("name") {
		oldName, newName := d.GetChange("name")
		_package["name"] = oldName
		_package["new-name"] = newName
	} else {
		_package["name"] = d.Get("name")
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		_package["ignore-errors"] = v.(bool)
	}
	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		_package["ignore-warnings"] = v.(bool)
	}

	if ok := d.HasChange("comments"); ok {
		_package["comments"] = d.Get("comments")
	}

	if ok := d.HasChange("color"); ok {
		_package["color"] = d.Get("color")
	}

	if ok := d.HasChange("access"); ok {
		_package["access"] = d.Get("access")
	}
	if ok := d.HasChange("desktop_security"); ok {
		_package["desktop-security"] = d.Get("desktop_security")
	}

	if ok := d.HasChange("installation_targets"); ok {
		if v, ok := d.GetOk("installation_targets"); ok {
			_package["installation-targets"] = v.(*schema.Set).List()
		} else {
			oldInstallationTargets, _ := d.GetChange("installation_targets")
			_package["installation-targets"] = map[string]interface{}{"remove": oldInstallationTargets.(*schema.Set).List()}
		}
	}
	if ok := d.HasChange("desktop_security"); ok {
		_package["desktop-security"] = d.Get("desktop_security")
	}
	if ok := d.HasChange("qos"); ok {
		_package["qos"] = d.Get("qos")
	}
	if ok := d.HasChange("qos_policy_type"); ok {
		_package["qos-policy-type"] = d.Get("qos_policy_type")
	}
	if ok := d.HasChange("threat_prevention"); ok {
		_package["threat-prevention"] = d.Get("threat_prevention")
	}
	if ok := d.HasChange("vpn_traditional_mode"); ok {
		_package["vpn_traditional_mode"] = d.Get("vpn_traditional_mode")
	}

	if ok := d.HasChange("tags"); ok {
		if v, ok := d.GetOk("tags"); ok {
			_package["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			_package["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	log.Println("Update Package - Map = ", _package)
	updatePackageRes, err := client.ApiCall("set-package", _package, client.GetSessionID(), true, false)
	if err != nil || !updatePackageRes.Success {
		if updatePackageRes.ErrorMsg != "" {
			return fmt.Errorf(updatePackageRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementPackage(d, m)
}

func deleteManagementPackage(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	packagePayload := map[string]interface{}{
		"uid": d.Id(),
	}

	deletePackageRes, err := client.ApiCall("delete-package", packagePayload, client.GetSessionID(), true, false)
	if err != nil || !deletePackageRes.Success {
		if deletePackageRes.ErrorMsg != "" {
			return fmt.Errorf(deletePackageRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
