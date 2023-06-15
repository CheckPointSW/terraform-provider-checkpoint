package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementGaiaBestPractice() *schema.Resource {
	return &schema.Resource{
		Create: createManagementGaiaBestPractice,
		Read:   readManagementGaiaBestPractice,
		Update: updateManagementGaiaBestPractice,
		Delete: deleteManagementGaiaBestPractice,
		Schema: map[string]*schema.Schema{
			"best_practice_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Best Practice ID.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Best Practice Name.",
			},
			"action_item": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "To comply with Best Practice, do this action item.",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Description of the Best Practice.",
			},
			"expected_output_text": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The expected output of the script as plain text.",
			},
			"expected_output_base64": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The expected output of the script as Base64.",
			},
			"practice_script_path": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The absolute path of the script on the Management Server to run on Gaia Security Gateways during the Compliance scans.",
			},
			"practice_script_base64": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The entire content of the script encoded in Base64 to run on Gaia Security Gateways during the Compliance scans.",
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
			"regulations": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The applicable regulations of the Gaia Best Practice. Appear only when the value of the 'details-level' parameter is set to 'full'.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"regulation_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the regulation.",
						},
						"requirement_description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the requirement.",
						},
						"requirement_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the requirement.",
						},
						"requirement_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the requirement.",
						},
					},
				},
			},
			"relevant_objects": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The applicable objects of the Gaia Best Practice. Appear only when the value of the 'details-level' parameter is set to 'full'.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Determines if the relevant object is enabled or not.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the relevant object.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the relevant object.",
						},
						"uid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The uid of the relevant object.",
						},
					},
				},
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The current status of the Best Practice.",
			},
			"user_defined": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Determines if the Gaia Best Practice is a user-defined best practice.",
			},
		},
	}
}

func createManagementGaiaBestPractice(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	gaiaBestPractice := make(map[string]interface{})

	if v, ok := d.GetOk("best_practice_id"); ok {
		gaiaBestPractice["best-practice-id"] = v.(string)
	}

	if v, ok := d.GetOk("name"); ok {
		gaiaBestPractice["name"] = v.(string)
	}

	if v, ok := d.GetOk("action_item"); ok {
		gaiaBestPractice["action-item"] = v.(string)
	}

	if v, ok := d.GetOk("description"); ok {
		gaiaBestPractice["description"] = v.(string)
	}

	if v, ok := d.GetOk("expected_output_text"); ok {
		gaiaBestPractice["expected-output-text"] = v.(string)
	}

	if v, ok := d.GetOk("expected_output_base64"); ok {
		gaiaBestPractice["expected-output-base64"] = v.(string)
	}

	if v, ok := d.GetOk("practice_script_path"); ok {
		gaiaBestPractice["practice-script-path"] = v.(string)
	}

	if v, ok := d.GetOk("practice_script_base64"); ok {
		gaiaBestPractice["practice-script-base64"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		gaiaBestPractice["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		gaiaBestPractice["ignore-errors"] = v.(bool)
	}

	log.Println("Create GaiaBestPractice - Map = ", gaiaBestPractice)

	addGaiaBestPracticeRes, err := client.ApiCall("add-gaia-best-practice", gaiaBestPractice, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !addGaiaBestPracticeRes.Success {
		if addGaiaBestPracticeRes.ErrorMsg != "" {
			return fmt.Errorf(addGaiaBestPracticeRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addGaiaBestPracticeRes.GetData()["uid"].(string))

	return readManagementGaiaBestPractice(d, m)
}

func readManagementGaiaBestPractice(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showGaiaBestPracticeRes, err := client.ApiCall("show-gaia-best-practice", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showGaiaBestPracticeRes.Success {
		if objectNotFound(showGaiaBestPracticeRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showGaiaBestPracticeRes.ErrorMsg)
	}

	gaiaBestPractice := showGaiaBestPracticeRes.GetData()

	log.Println("Read GaiaBestPractice - Show JSON = ", gaiaBestPractice)

	if v := gaiaBestPractice["best-practice-id"]; v != nil {
		_ = d.Set("best_practice_id", v)
	}

	if v := gaiaBestPractice["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := gaiaBestPractice["action-item"]; v != nil {
		_ = d.Set("action_item", v)
	}

	if v := gaiaBestPractice["description"]; v != nil {
		_ = d.Set("description", v)
	}

	if v := gaiaBestPractice["expected-output-base64"]; v != nil {
		_ = d.Set("expected_output_base64", v)
	}

	if v := gaiaBestPractice["practice-script-path"]; v != nil {
		_ = d.Set("practice_script_path", v)
	}

	if v := gaiaBestPractice["practice-script-base64"]; v != nil {
		_ = d.Set("practice_script_base64", v)
	}

	if gaiaBestPractice["regulations"] != nil {
		regulationsList := gaiaBestPractice["regulations"].([]interface{})

		if len(regulationsList) > 0 {
			var regulationsListToReturn []map[string]interface{}

			for i := range regulationsList {
				regulationsMap := regulationsList[i].(map[string]interface{})

				regulationsMapToAdd := make(map[string]interface{})

				if v, _ := regulationsMap["regulation-name"]; v != nil {
					regulationsMapToAdd["regulation_name"] = v
				}
				if v, _ := regulationsMap["regulation-description"]; v != nil {
					regulationsMapToAdd["regulation_description"] = v
				}
				if v, _ := regulationsMap["requirement-id"]; v != nil {
					regulationsMapToAdd["requirement_id"] = v
				}
				if v, _ := regulationsMap["requirement-status"]; v != nil {
					regulationsMapToAdd["requirement_status"] = v
				}

				regulationsListToReturn = append(regulationsListToReturn, regulationsMapToAdd)
			}

			_ = d.Set("regulations", regulationsListToReturn)
		} else {
			_ = d.Set("regulations", regulationsList)
		}
	} else {
		_ = d.Set("regulations", nil)
	}

	if gaiaBestPractice["relevant-objects"] != nil {
		relevantObjectsList := gaiaBestPractice["relevant-objects"].([]interface{})

		if len(relevantObjectsList) > 0 {
			var relevantObjectsListToReturn []map[string]interface{}

			for i := range relevantObjectsList {
				relevantObjectsMap := relevantObjectsList[i].(map[string]interface{})

				relevantObjectsMapToAdd := make(map[string]interface{})

				if v, _ := relevantObjectsMap["enabled"]; v != nil {
					relevantObjectsMapToAdd["enabled"] = v
				}
				if v, _ := relevantObjectsMap["name"]; v != nil {
					relevantObjectsMapToAdd["name"] = v
				}
				if v, _ := relevantObjectsMap["status"]; v != nil {
					relevantObjectsMapToAdd["status"] = v
				}
				if v, _ := relevantObjectsMap["uid"]; v != nil {
					relevantObjectsMapToAdd["uid"] = v
				}

				relevantObjectsListToReturn = append(relevantObjectsListToReturn, relevantObjectsMapToAdd)
			}

			_ = d.Set("relevant_objects", relevantObjectsListToReturn)
		} else {
			_ = d.Set("relevant_objects", relevantObjectsList)
		}
	} else {
		_ = d.Set("relevant_objects", nil)
	}

	if v := gaiaBestPractice["status"]; v != nil {
		_ = d.Set("status", v)
	}

	if v := gaiaBestPractice["user-defined"]; v != nil {
		_ = d.Set("user_defined", v)
	}

	return nil

}

func updateManagementGaiaBestPractice(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	gaiaBestPractice := make(map[string]interface{})

	if ok := d.HasChange("best_practice_id"); ok {
		gaiaBestPractice["best-practice-id"] = d.Get("best_practice_id")
	}

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		gaiaBestPractice["name"] = oldName
		gaiaBestPractice["new-name"] = newName
	} else {
		gaiaBestPractice["name"] = d.Get("name")
	}

	if ok := d.HasChange("action_item"); ok {
		gaiaBestPractice["action-item"] = d.Get("action_item")
	}

	if ok := d.HasChange("description"); ok {
		gaiaBestPractice["description"] = d.Get("description")
	}

	if ok := d.HasChange("expected_output_text"); ok {
		gaiaBestPractice["expected-output-text"] = d.Get("expected_output_text")
	}

	if ok := d.HasChange("expected_output_base64"); ok {
		gaiaBestPractice["expected-output-base64"] = d.Get("expected_output_base64")
	}

	if ok := d.HasChange("practice_script_path"); ok {
		gaiaBestPractice["practice-script-path"] = d.Get("practice_script_path")
	}

	if ok := d.HasChange("practice_script_base64"); ok {
		gaiaBestPractice["practice-script-base64"] = d.Get("practice_script_base64")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		gaiaBestPractice["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		gaiaBestPractice["ignore-errors"] = v.(bool)
	}

	log.Println("Update GaiaBestPractice - Map = ", gaiaBestPractice)

	updateGaiaBestPracticeRes, err := client.ApiCall("set-gaia-best-practice", gaiaBestPractice, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !updateGaiaBestPracticeRes.Success {
		if updateGaiaBestPracticeRes.ErrorMsg != "" {
			return fmt.Errorf(updateGaiaBestPracticeRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementGaiaBestPractice(d, m)
}

func deleteManagementGaiaBestPractice(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	gaiaBestPracticePayload := map[string]interface{}{
		"uid": d.Id(),
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		gaiaBestPracticePayload["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		gaiaBestPracticePayload["ignore-errors"] = v.(bool)
	}

	log.Println("Delete GaiaBestPractice")

	deleteGaiaBestPracticeRes, err := client.ApiCall("delete-gaia-best-practice", gaiaBestPracticePayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteGaiaBestPracticeRes.Success {
		if deleteGaiaBestPracticeRes.ErrorMsg != "" {
			return fmt.Errorf(deleteGaiaBestPracticeRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
