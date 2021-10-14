package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementNatSection() *schema.Resource {
	return &schema.Resource{
		Create: createManagementNatSection,
		Read:   readManagementNatSection,
		Update: updateManagementNatSection,
		Delete: deleteManagementNatSection,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
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
		},
	}
}

func createManagementNatSection(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	natSection := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		natSection["name"] = v.(string)
	}

	if v, ok := d.GetOk("package"); ok {
		natSection["package"] = v.(string)
	}

	if _, ok := d.GetOk("position"); ok {

		if v, ok := d.GetOk("position.top"); ok {
			if v.(string) == "top" {
				natSection["position"] = "top" // entire rule-base
			} else {
				natSection["position"] = map[string]interface{}{"top": v.(string)} // section-name
			}
		}

		if v, ok := d.GetOk("position.above"); ok {
			natSection["position"] = map[string]interface{}{"above": v.(string)}
		}

		if v, ok := d.GetOk("position.below"); ok {
			natSection["position"] = map[string]interface{}{"below": v.(string)}
		}

		if v, ok := d.GetOk("position.bottom"); ok {
			if v.(string) == "bottom" {
				natSection["position"] = "bottom" // entire rule-base
			} else {
				natSection["position"] = map[string]interface{}{"bottom": v.(string)} // section-name
			}
		}
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		natSection["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		natSection["ignore-errors"] = v.(bool)
	}

	log.Println("Create NAT section - Map = ", natSection)

	addNatSectionRes, err := client.ApiCall("add-nat-section", natSection, client.GetSessionID(), true, false)
	if err != nil || !addNatSectionRes.Success {
		if addNatSectionRes.ErrorMsg != "" {
			return fmt.Errorf(addNatSectionRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addNatSectionRes.GetData()["uid"].(string))

	return readManagementNatSection(d, m)
}

func readManagementNatSection(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid":     d.Id(),
		"package": d.Get("package"),
	}

	showNatSectionRes, err := client.ApiCall("show-nat-section", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showNatSectionRes.Success {
		if objectNotFound(showNatSectionRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showNatSectionRes.ErrorMsg)
	}

	natSection := showNatSectionRes.GetData()

	log.Println("Read NatSection - Show JSON = ", natSection)

	if v := natSection["name"]; v != nil {
		_ = d.Set("name", v)
	}

	return nil
}

func updateManagementNatSection(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	natSection := make(map[string]interface{})

	natSection["uid"] = d.Id()
	natSection["package"] = d.Get("package")

	if ok := d.HasChange("name"); ok {
		natSection["new-name"] = d.Get("name")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		natSection["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		natSection["ignore-errors"] = v.(bool)
	}

	log.Println("Update NAT section - Map = ", natSection)

	updateNatSectionRes, err := client.ApiCall("set-nat-section", natSection, client.GetSessionID(), true, false)
	if err != nil || !updateNatSectionRes.Success {
		if updateNatSectionRes.ErrorMsg != "" {
			return fmt.Errorf(updateNatSectionRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementNatSection(d, m)
}

func deleteManagementNatSection(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	natSectionPayload := map[string]interface{}{
		"uid":     d.Id(),
		"package": d.Get("package"),
	}

	log.Println("Delete NAT section")

	deleteNatSectionRes, err := client.ApiCall("delete-nat-section", natSectionPayload, client.GetSessionID(), true, false)
	if err != nil || !deleteNatSectionRes.Success {
		if deleteNatSectionRes.ErrorMsg != "" {
			return fmt.Errorf(deleteNatSectionRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
