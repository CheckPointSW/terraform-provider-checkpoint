package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceManagementRegulation() *schema.Resource {
	return &schema.Resource{
		Create: createManagementRegulation,
		Read:   readManagementRegulation,
		Update: updateManagementRegulation,
		Delete: deleteManagementRegulation,
		Schema: map[string]*schema.Schema{
			"full_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Regulation full name.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Regulation name.",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Determines if the regulation is enabled.",
				Default:     true,
			},
			"show_requirements": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Show the requirements of the regulation.",
				Default:     false,
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
				Description: "Comments about this regulation.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
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
			"requirements": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The requirements of the regulation, identified by name.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func createManagementRegulation(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	regulation := make(map[string]interface{})

	if v, ok := d.GetOk("full_name"); ok {
		regulation["full-name"] = v.(string)
	}

	if v, ok := d.GetOk("name"); ok {
		regulation["name"] = v.(string)
	}

	if v, ok := d.GetOkExists("enabled"); ok {
		regulation["enabled"] = v.(bool)
	}

	if v, ok := d.GetOkExists("show_requirements"); ok {
		regulation["show-requirements"] = v.(bool)
	}

	if v, ok := d.GetOk("color"); ok {
		regulation["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		regulation["comments"] = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		regulation["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		regulation["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		regulation["ignore-errors"] = v.(bool)
	}

	log.Println("Create Regulation - Map = ", regulation)

	addRegulationRes, err := client.ApiCall("add-regulation", regulation, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !addRegulationRes.Success {
		if addRegulationRes.ErrorMsg != "" {
			return fmt.Errorf(addRegulationRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addRegulationRes.GetData()["uid"].(string))

	return readManagementRegulation(d, m)
}

func readManagementRegulation(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showRegulationRes, err := client.ApiCall("show-regulation", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showRegulationRes.Success {
		if objectNotFound(showRegulationRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showRegulationRes.ErrorMsg)
	}

	regulation := showRegulationRes.GetData()

	log.Println("Read Regulation - Show JSON = ", regulation)

	if v := regulation["full-name"]; v != nil {
		_ = d.Set("full_name", v)
	}

	if v := regulation["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := regulation["enabled"]; v != nil {
		_ = d.Set("enabled", v)
	}

	if v := regulation["show-requirements"]; v != nil {
		_ = d.Set("show_requirements", v)
	}

	if regulation["requirements"] != nil {
		requirementsJson, ok := regulation["requirements"].([]interface{})
		if ok {
			requirementsIds := make([]string, 0)
			if len(requirementsJson) > 0 {
				for _, requirement := range requirementsJson {
					if requirementMap, ok := requirement.(map[string]interface{}); ok {
						requirementsIds = append(requirementsIds, requirementMap["name"].(string))
					} else {
						requirementsIds = append(requirementsIds, requirement.(string))
					}
				}
			}
			_ = d.Set("requirements", requirementsIds)
		}
	} else {
		_ = d.Set("requirements", nil)
	}

	if v := regulation["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := regulation["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if regulation["tags"] != nil {
		tagsJson, ok := regulation["tags"].([]interface{})
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

	if v := regulation["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := regulation["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementRegulation(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	regulation := make(map[string]interface{})

	regulation["uid"] = d.Id()

	if ok := d.HasChange("full_name"); ok {
		regulation["full-name"] = d.Get("full_name")
	}

	if ok := d.HasChange("name"); ok {
		if v, ok := d.GetOk("name"); ok {
			regulation["new-name"] = v.(string)
		}
	}

	if v, ok := d.GetOkExists("enabled"); ok {
		regulation["enabled"] = v.(bool)
	}

	if v, ok := d.GetOkExists("show_requirements"); ok {
		regulation["show-requirements"] = v.(bool)
	}

	if ok := d.HasChange("color"); ok {
		regulation["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		regulation["comments"] = d.Get("comments")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			regulation["tags"] = v.(*schema.Set).List()
		}
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		regulation["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		regulation["ignore-errors"] = v.(bool)
	}

	log.Println("Update Regulation - Map = ", regulation)

	updateRegulationRes, err := client.ApiCall("set-regulation", regulation, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !updateRegulationRes.Success {
		if updateRegulationRes.ErrorMsg != "" {
			return fmt.Errorf(updateRegulationRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementRegulation(d, m)
}

func deleteManagementRegulation(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	regulationPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete Regulation")

	deleteRegulationRes, err := client.ApiCall("delete-regulation", regulationPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteRegulationRes.Success {
		if deleteRegulationRes.ErrorMsg != "" {
			return fmt.Errorf(deleteRegulationRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
