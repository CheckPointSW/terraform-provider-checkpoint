package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceManagementRequirement() *schema.Resource {
	return &schema.Resource{
		Create: createManagementRequirement,
		Read:   readManagementRequirement,
		Update: updateManagementRequirement,
		Delete: deleteManagementRequirement,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Requirement name.",
			},
			"regulation": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The relevant regulation. Identified by name or UID.",
			},
			"best_practices": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "UIDs or IDs of the relevant best practices for the requirement.",
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
				Description: "The requirement comments.",
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
			"new_regulation_uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The new regulation UID for this requirement.",
			},
			"regulation_uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The relevant regulation UID for this requirement.",
			},
		},
	}
}

func createManagementRequirement(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	requirement := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		requirement["name"] = v.(string)
	}

	if v, ok := d.GetOk("regulation"); ok {
		requirement["regulation"] = v.(string)
	}

	if v, ok := d.GetOk("best_practices"); ok {
		requirement["best-practices"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		requirement["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		requirement["comments"] = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		requirement["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		requirement["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		requirement["ignore-errors"] = v.(bool)
	}

	log.Println("Create Requirement - Map = ", requirement)

	if v, ok := d.GetOk("regulation_uid"); ok {
		requirement["regulation-uid"] = v.(string)
	}

	addRequirementRes, err := client.ApiCall("add-requirement", requirement, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !addRequirementRes.Success {
		if addRequirementRes.ErrorMsg != "" {
			return fmt.Errorf(addRequirementRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addRequirementRes.GetData()["uid"].(string))

	return readManagementRequirement(d, m)
}

func readManagementRequirement(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showRequirementRes, err := client.ApiCall("show-requirement", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showRequirementRes.Success {
		if objectNotFound(showRequirementRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showRequirementRes.ErrorMsg)
	}

	requirement := showRequirementRes.GetData()

	log.Println("Read Requirement - Show JSON = ", requirement)

	if v := requirement["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := requirement["regulation"]; v != nil {
		if regulationMap, ok := v.(map[string]interface{}); ok {
			_ = d.Set("regulation", regulationMap["name"])
		} else {
			_ = d.Set("regulation", v)
		}
	}

	if requirement["best-practices"] != nil {
		bestPracticesJson, ok := requirement["best-practices"].([]interface{})
		if ok {
			bestPracticesIds := make([]string, 0)
			if len(bestPracticesJson) > 0 {
				for _, best_practices := range bestPracticesJson {
					best_practices := best_practices.(map[string]interface{})
					bestPracticesIds = append(bestPracticesIds, best_practices["name"].(string))
				}
			}
			_ = d.Set("best_practices", bestPracticesIds)
		}
	} else {
		_ = d.Set("best_practices", nil)
	}

	if v := requirement["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := requirement["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if requirement["tags"] != nil {
		tagsJson, ok := requirement["tags"].([]interface{})
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

	if v := requirement["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := requirement["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementRequirement(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	requirement := make(map[string]interface{})

	requirement["uid"] = d.Id()

	if ok := d.HasChange("name"); ok {
		if v, ok := d.GetOk("name"); ok {
			requirement["new-name"] = v.(string)
		}
	}

	if ok := d.HasChange("regulation"); ok {
		requirement["new-regulation"] = d.Get("regulation")
	}

	if d.HasChange("best_practices") {
		if v, ok := d.GetOk("best_practices"); ok {
			requirement["best-practices"] = v.(*schema.Set).List()
		}
	}

	if ok := d.HasChange("color"); ok {
		requirement["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		requirement["comments"] = d.Get("comments")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			requirement["tags"] = v.(*schema.Set).List()
		}
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		requirement["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		requirement["ignore-errors"] = v.(bool)
	}

	log.Println("Update Requirement - Map = ", requirement)

	if ok := d.HasChange("new_regulation_uid"); ok {
		if v, ok := d.GetOk("new_regulation_uid"); ok {
			requirement["new-regulation-uid"] = v.(string)
		}
	}

	updateRequirementRes, err := client.ApiCall("set-requirement", requirement, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !updateRequirementRes.Success {
		if updateRequirementRes.ErrorMsg != "" {
			return fmt.Errorf(updateRequirementRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementRequirement(d, m)
}

func deleteManagementRequirement(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	requirementPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete Requirement")

	deleteRequirementRes, err := client.ApiCall("delete-requirement", requirementPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteRequirementRes.Success {
		if deleteRequirementRes.ErrorMsg != "" {
			return fmt.Errorf(deleteRequirementRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
