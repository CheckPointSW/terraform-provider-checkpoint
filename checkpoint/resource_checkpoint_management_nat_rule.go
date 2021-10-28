package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strings"
)

func resourceManagementNatRule() *schema.Resource {
	return &schema.Resource{
		Create: createManagementNatRule,
		Read:   readManagementNatRule,
		Update: updateManagementNatRule,
		Delete: deleteManagementNatRule,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				arr := strings.Split(d.Id(), ";")
				if len(arr) != 2 {
					return nil, fmt.Errorf("invalid unique identifier format. UID format: <PACKAGE_NAME>;<RULE_UID>")
				}
				_ = d.Set("package", arr[0])
				d.SetId(arr[1])
				return []*schema.ResourceData{d}, nil
			},
		},
		Schema: map[string]*schema.Schema{
			"package": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the package.",
			},
			"position": {
				Type:        schema.TypeMap,
				Required:    true,
				Description: "Position in the rulebase.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"top": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Add rule on top of specific section identified by uid or name. Select value 'top' for entire rule base.",
						},
						"above": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Add rule above specific section/rule identified by uid or name.",
						},
						"below": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Add rule below specific section/rule identified by uid or name.",
						},
						"bottom": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Add rule in the bottom of specific section identified by uid or name. Select value 'bottom' for entire rule base.",
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Rule name.",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable/Disable the rule.",
				Default:     true,
			},
			"install_on": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Which Gateways identified by the name or UID to install the policy on.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"method": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Nat method.",
				Default:     "static",
			},
			"original_destination": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Original destination.",
				Default:     "Any",
			},
			"original_service": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Original service.",
				Default:     "Any",
			},
			"original_source": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Original source.",
				Default:     "Any",
			},
			"translated_destination": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Translated destination.",
				Default:     "Original",
			},
			"translated_service": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Translated service.",
				Default:     "Original",
			},
			"translated_source": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Translated source.",
				Default:     "Original",
			},
			"auto_generated": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Auto generated.",
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
			"comments": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comments string.",
			},
		},
	}
}

func createManagementNatRule(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	natRule := make(map[string]interface{})

	if v, ok := d.GetOk("package"); ok {
		natRule["package"] = v.(string)
	}
	if _, ok := d.GetOk("position"); ok {

		if v, ok := d.GetOk("position.top"); ok {
			if v.(string) == "top" {
				natRule["position"] = "top" // entire rule-base
			} else {
				natRule["position"] = map[string]interface{}{"top": v.(string)} // section-name
			}
		}

		if v, ok := d.GetOk("position.above"); ok {
			natRule["position"] = map[string]interface{}{"above": v.(string)}
		}

		if v, ok := d.GetOk("position.below"); ok {
			natRule["position"] = map[string]interface{}{"below": v.(string)}
		}

		if v, ok := d.GetOk("position.bottom"); ok {
			if v.(string) == "bottom" {
				natRule["position"] = "bottom" // entire rule-base
			} else {
				natRule["position"] = map[string]interface{}{"bottom": v.(string)} // section-name
			}
		}
	}

	if v, ok := d.GetOk("name"); ok {
		natRule["name"] = v.(string)
	}

	if v, ok := d.GetOkExists("enabled"); ok {
		natRule["enabled"] = v.(bool)
	}

	if val, ok := d.GetOk("install_on"); ok {
		natRule["install-on"] = val.(*schema.Set).List()
	}

	if v, ok := d.GetOk("original_destination"); ok {
		natRule["original-destination"] = v.(string)
	}

	if v, ok := d.GetOk("original_service"); ok {
		natRule["original-service"] = v.(string)
	}

	if v, ok := d.GetOk("original_source"); ok {
		natRule["original-source"] = v.(string)
	}

	if v, ok := d.GetOk("translated_destination"); ok {
		natRule["translated-destination"] = v.(string)
	}

	if v, ok := d.GetOk("translated_service"); ok {
		natRule["translated-service"] = v.(string)
	}

	if v, ok := d.GetOk("translated_source"); ok {
		natRule["translated-source"] = v.(string)
	}

	if val, ok := d.GetOk("comments"); ok {
		natRule["comments"] = val.(string)
	}

	if val, ok := d.GetOkExists("ignore_errors"); ok {
		natRule["ignore-errors"] = val.(bool)
	}
	if val, ok := d.GetOkExists("ignore_warnings"); ok {
		natRule["ignore-warnings"] = val.(bool)
	}

	log.Println("Create NAT Rule - Map = ", natRule)

	addNatRuleRes, err := client.ApiCall("add-nat-rule", natRule, client.GetSessionID(), true, false)
	if err != nil || !addNatRuleRes.Success {
		if addNatRuleRes.ErrorMsg != "" {
			return fmt.Errorf(addNatRuleRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addNatRuleRes.GetData()["uid"].(string))

	return readManagementNatRule(d, m)
}

func readManagementNatRule(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid":     d.Id(),
		"package": d.Get("package"),
	}

	showNatRuleRes, err := client.ApiCall("show-nat-rule", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showNatRuleRes.Success {
		// Handle delete resource from other clients
		if objectNotFound(showNatRuleRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showNatRuleRes.ErrorMsg)
	}

	natRule := showNatRuleRes.GetData()

	log.Println("Read NAT Rule - Show JSON = ", natRule)

	if v := natRule["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := natRule["auto-generated"]; v != nil {
		_ = d.Set("auto_generated", v)
	}

	if v := natRule["enabled"]; v != nil {
		_ = d.Set("enabled", v)
	}

	if v := natRule["method"]; v != nil {
		_ = d.Set("method", v)
	}

	if natRule["install-on"] != nil {
		installOnJson := natRule["install-on"].([]interface{})
		installOnJsonIds := make([]string, 0)
		if len(installOnJson) > 0 {
			for _, installOn := range installOnJson {
				installOn := installOn.(map[string]interface{})
				installOnJsonIds = append(installOnJsonIds, installOn["name"].(string))
			}
		}
		_, installOnInConf := d.GetOk("install_on")
		if installOnJsonIds[0] == "Policy Targets" && !installOnInConf {
			_ = d.Set("install_on", []interface{}{})
		} else {
			_ = d.Set("install_on", installOnJsonIds)
		}
	}

	if v := natRule["original-destination"]; v != nil {
		_ = d.Set("original_destination", v.(map[string]interface{})["name"])
	}

	if v := natRule["original-service"]; v != nil {
		_ = d.Set("original_service", v.(map[string]interface{})["name"])
	}

	if v := natRule["original-source"]; v != nil {
		_ = d.Set("original_source", v.(map[string]interface{})["name"])
	}

	if v := natRule["translated-destination"]; v != nil {
		_ = d.Set("translated_destination", v.(map[string]interface{})["name"])
	}

	if v := natRule["translated-service"]; v != nil {
		_ = d.Set("translated_service", v.(map[string]interface{})["name"])
	}

	if v := natRule["translated-source"]; v != nil {
		_ = d.Set("translated_source", v.(map[string]interface{})["name"])
	}

	if v := natRule["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil
}

func updateManagementNatRule(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	natRule := make(map[string]interface{})

	natRule["uid"] = d.Id()
	natRule["package"] = d.Get("package")

	if d.HasChange("name") {
		natRule["new-name"] = d.Get("name")
	}

	if d.HasChange("position") {
		if _, ok := d.GetOk("position"); ok {

			if v, ok := d.GetOk("position.top"); ok {
				if v.(string) == "top" {
					natRule["new-position"] = "top" // entire rule-base
				} else {
					natRule["new-position"] = map[string]interface{}{"top": v.(string)} // specific section-name
				}
			}

			if v, ok := d.GetOk("position.above"); ok {
				natRule["new-position"] = map[string]interface{}{"above": v.(string)}
			}

			if v, ok := d.GetOk("position.below"); ok {
				natRule["new-position"] = map[string]interface{}{"below": v.(string)}
			}

			if v, ok := d.GetOk("position.bottom"); ok {
				if v.(string) == "bottom" {
					natRule["new-position"] = "bottom" // entire rule-base
				} else {
					natRule["new-position"] = map[string]interface{}{"bottom": v.(string)} // specific section-name
				}
			}
		}
	}

	if d.HasChange("install_on") {
		if v, ok := d.GetOk("install_on"); ok {
			natRule["install-on"] = v.(*schema.Set).List()
		} else {
			oldInstallOn, _ := d.GetChange("install_on")
			natRule["install-on"] = map[string]interface{}{"remove": oldInstallOn.(*schema.Set).List()}
		}
	}

	if d.HasChange("enabled") {
		natRule["enabled"] = d.Get("enabled")
	}

	if d.HasChange("method") {
		natRule["method"] = d.Get("method")
	}

	if d.HasChange("original_destination") {
		natRule["original-destination"] = d.Get("original_destination")
	}

	if d.HasChange("original_service") {
		natRule["original-service"] = d.Get("original_service")
	}

	if d.HasChange("original_source") {
		natRule["original-source"] = d.Get("original_source")
	}

	if d.HasChange("translated_destination") {
		natRule["translated-destination"] = d.Get("translated_destination")
	}

	if d.HasChange("translated_service") {
		natRule["translated-service"] = d.Get("translated_service")
	}

	if d.HasChange("translated_source") {
		natRule["translated-source"] = d.Get("translated_source")
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		natRule["ignore-errors"] = v.(bool)
	}
	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		natRule["ignore-warnings"] = v.(bool)
	}

	if d.HasChange("comments") {
		natRule["comments"] = d.Get("comments")
	}

	log.Println("Update NAT Rule - Map = ", natRule)

	updateNatRuleRes, err := client.ApiCall("set-nat-rule", natRule, client.GetSessionID(), true, false)
	if err != nil || !updateNatRuleRes.Success {
		if updateNatRuleRes.ErrorMsg != "" {
			return fmt.Errorf(updateNatRuleRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	return readManagementNatRule(d, m)
}

func deleteManagementNatRule(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	natRulePayload := map[string]interface{}{
		"uid":     d.Id(),
		"package": d.Get("package"),
	}

	deleteAccessRuleRes, err := client.ApiCall("delete-nat-rule", natRulePayload, client.GetSessionID(), true, false)
	if err != nil || !deleteAccessRuleRes.Success {
		if deleteAccessRuleRes.ErrorMsg != "" {
			return fmt.Errorf(deleteAccessRuleRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId("")
	return nil
}
