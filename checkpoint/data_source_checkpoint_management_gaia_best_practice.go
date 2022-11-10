package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementGaiaBestPractice() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementGaiaBestPracticeRead,
		Schema: map[string]*schema.Schema{
			"best_practice_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Best Practice ID.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Best Practice Name.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Best Practice UID.",
			},
			"action_item": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Action item to comply with the Best Practice.",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description of the Best Practice.",
			},
			"expected_output_base64": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The expected output of the script in Base64. Available only for user-defined best practices.",
			},
			"practice_script_base64": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The script to run on Gaia Security Gateways during the Compliance scans in Base64. Available only for user-defined best practices.",
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

func dataSourceManagementGaiaBestPracticeRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)
	bestPracticeId := d.Get("best_practice_id").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	} else if bestPracticeId != "" {
		payload["best-practice-id"] = bestPracticeId
	}

	showGaiaBestPractice, err := client.ApiCall("show-gaia-best-practice", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showGaiaBestPractice.Success {
		return fmt.Errorf(showGaiaBestPractice.ErrorMsg)
	}

	gaiaBestPractice := showGaiaBestPractice.GetData()

	log.Println("Read Gaia Best Practice - Show JSON = ", gaiaBestPractice)

	if v := gaiaBestPractice["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := gaiaBestPractice["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := gaiaBestPractice["best-practice-id"]; v != nil {
		_ = d.Set("best_practice_id", v)
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
