package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"

	"strconv"
)

func resourceManagementGuidelineCellApprovals() *schema.Resource {
	return &schema.Resource{
		Create: createManagementGuidelineCellApprovals,
		Read:   readManagementGuidelineCellApprovals,
		Update: updateManagementGuidelineCellApprovals,
		Delete: deleteManagementGuidelineCellApprovals,
		Schema: map[string]*schema.Schema{
			"guideline": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The guideline (identified by UID or name) in which we approve the violation.",
			},
			"approvals": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "List of approved rules.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rules": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "The full paths (pairs of layer and rule) of the approved rules.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"layer": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The Layer identifier (name or UID).",
									},
									"rule": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The rule identifier (name if unique, rule position number in rule-base or UID).",
									},
								},
							},
						},
					},
				},
			},
			"from": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "\"from\" segment (identified by UID or name), or 'any' to approved the rule across all cells (possible only if \"to\" is also 'any'). This field is mandatory if \"from-type\" is 'Network Group'.",
			},
			"to": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "\"to\" segment (identified by UID or name), or 'any' to approved the rule across all cells (possible only if \"from\" is also 'any'). This field is mandatory if \"to-type\" is 'Network Group'.",
			},
			"comment": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Comment on the approvals. The same comment to all the requested approvals.",
			},
			"from_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The type of the segment in the 'from' axis.",
				Default:     "Network Group",
			},
			"to_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The type of the segment in the 'to' axis.",
				Default:     "Network Group",
			},
			"policy_package": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The policy package (identified by UID or name) in which we approve the violation. This field is mandatory only if the ordered-access-layer (first layer in path) is from a global domain with AGP.",
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
			"delete_scope": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Indicates whether to delete all the approval scope, or only remove the requested cell from the scope. Relevant only for guideline approvals.",
				Default:     "Single Cell",
			},
		},
	}
}

func createManagementGuidelineCellApprovals(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	guidelineCellApprovals := make(map[string]interface{})

	if v, ok := d.GetOk("guideline"); ok {
		guidelineCellApprovals["guideline"] = v.(string)
	}

	if v, ok := d.GetOk("from"); ok {
		guidelineCellApprovals["from"] = v.(string)
	}

	if v, ok := d.GetOk("to"); ok {
		guidelineCellApprovals["to"] = v.(string)
	}

	if v, ok := d.GetOk("comment"); ok {
		guidelineCellApprovals["comment"] = v.(string)
	}

	if v, ok := d.GetOk("approvals"); ok {

		approvalsList := v.([]interface{})

		if len(approvalsList) > 0 {

			var approvalsPayload []map[string]interface{}

			for i := range approvalsList {

				Payload := make(map[string]interface{})

				if v, ok := d.GetOk("approvals." + strconv.Itoa(i) + ".rules"); ok {

					rulesList := v.([]interface{})

					if len(rulesList) > 0 {

						var rulesPayload []map[string]interface{}

						for j := range rulesList {

							rulesMapToAdd := make(map[string]interface{})

							if v, ok := d.GetOk("approvals." + strconv.Itoa(i) + ".rules." + strconv.Itoa(j) + ".layer"); ok {
								rulesMapToAdd["layer"] = v.(string)
							}
							if v, ok := d.GetOk("approvals." + strconv.Itoa(i) + ".rules." + strconv.Itoa(j) + ".rule"); ok {
								rulesMapToAdd["rule"] = v.(string)
							}
							rulesPayload = append(rulesPayload, rulesMapToAdd)
						}
						Payload["rules"] = rulesPayload
					}
				}
				approvalsPayload = append(approvalsPayload, Payload)
			}
			guidelineCellApprovals["approvals"] = approvalsPayload
		}
	}

	if v, ok := d.GetOk("from_type"); ok {
		guidelineCellApprovals["from-type"] = v.(string)
	}

	if v, ok := d.GetOk("to_type"); ok {
		guidelineCellApprovals["to-type"] = v.(string)
	}

	if v, ok := d.GetOk("policy_package"); ok {
		guidelineCellApprovals["policy-package"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		guidelineCellApprovals["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		guidelineCellApprovals["ignore-errors"] = v.(bool)
	}

	log.Println("Create GuidelineCellApprovals - Map = ", guidelineCellApprovals)

	addGuidelineCellApprovalsRes, err := client.ApiCall("add-guideline-cell-approvals", guidelineCellApprovals, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !addGuidelineCellApprovalsRes.Success {
		if addGuidelineCellApprovalsRes.ErrorMsg != "" {
			return fmt.Errorf(addGuidelineCellApprovalsRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addGuidelineCellApprovalsRes.GetData()["uid"].(string))

	return readManagementGuidelineCellApprovals(d, m)
}

func readManagementGuidelineCellApprovals(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showGuidelineCellApprovalsRes, err := client.ApiCall("show-guideline-cell-approvals", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showGuidelineCellApprovalsRes.Success {
		if objectNotFound(showGuidelineCellApprovalsRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showGuidelineCellApprovalsRes.ErrorMsg)
	}

	guidelineCellApprovals := showGuidelineCellApprovalsRes.GetData()

	log.Println("Read GuidelineCellApprovals - Show JSON = ", guidelineCellApprovals)

	if v := guidelineCellApprovals["guideline"]; v != nil {
		_ = d.Set("guideline", v)
	}

	if v := guidelineCellApprovals["from"]; v != nil {
		if fromMap, ok := v.(map[string]interface{}); ok {
			_ = d.Set("from", fromMap["name"])
		} else {
			_ = d.Set("from", v)
		}
	}

	if v := guidelineCellApprovals["to"]; v != nil {
		if toMap, ok := v.(map[string]interface{}); ok {
			_ = d.Set("to", toMap["name"])
		} else {
			_ = d.Set("to", v)
		}
	}

	if v := guidelineCellApprovals["comment"]; v != nil {
		_ = d.Set("comment", v)
	}

	if guidelineCellApprovals["approvals"] != nil {

		approvalsList := guidelineCellApprovals["approvals"].([]interface{})

		if len(approvalsList) > 0 {

			var approvalsListToReturn []map[string]interface{}

			for i := range approvalsList {

				approvalsMap := approvalsList[i].(map[string]interface{})

				approvalsMapToAdd := make(map[string]interface{})

				if approvalsMap["rules"] != nil {

					rulesList := approvalsMap["rules"].([]interface{})

					if len(rulesList) > 0 {

						var rulesListToReturn []map[string]interface{}

						for j := range rulesList {

							rulesMap := rulesList[j].(map[string]interface{})

							rulesMapToAdd := make(map[string]interface{})

							if v, _ := rulesMap["layer"]; v != nil {
								rulesMapToAdd["layer"] = v
							}
							if v, _ := rulesMap["rule"]; v != nil {
								rulesMapToAdd["rule"] = v
							}
							rulesListToReturn = append(rulesListToReturn, rulesMapToAdd)
						}
						approvalsMapToAdd["rules"] = rulesListToReturn
					} else {
						approvalsMapToAdd["rules"] = rulesList
					}
				} else {
					approvalsMapToAdd["rules"] = nil
				}

				approvalsListToReturn = append(approvalsListToReturn, approvalsMapToAdd)
			}

			_ = d.Set("approvals", approvalsListToReturn)
		} else {
			_ = d.Set("approvals", approvalsList)
		}
	} else {
		_ = d.Set("approvals", nil)
	}

	if v := guidelineCellApprovals["from-type"]; v != nil {
		_ = d.Set("from_type", v)
	}

	if v := guidelineCellApprovals["to-type"]; v != nil {
		_ = d.Set("to_type", v)
	}

	if v := guidelineCellApprovals["policy-package"]; v != nil {
		_ = d.Set("policy_package", v)
	}

	if v := guidelineCellApprovals["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := guidelineCellApprovals["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementGuidelineCellApprovals(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	guidelineCellApprovals := make(map[string]interface{})

	guidelineCellApprovals["uid"] = d.Id()

	if ok := d.HasChange("guideline"); ok {
		guidelineCellApprovals["guideline"] = d.Get("guideline")
	}

	if ok := d.HasChange("from"); ok {
		guidelineCellApprovals["from"] = d.Get("from")
	}

	if ok := d.HasChange("to"); ok {
		guidelineCellApprovals["to"] = d.Get("to")
	}

	if ok := d.HasChange("comment"); ok {
		guidelineCellApprovals["comment"] = d.Get("comment")
	}

	if d.HasChange("approvals") {

		if v, ok := d.GetOk("approvals"); ok {

			approvalsList := v.([]interface{})

			var approvalsPayload []map[string]interface{}

			for i := range approvalsList {

				Payload := make(map[string]interface{})

				if v, ok := d.GetOk("approvals." + strconv.Itoa(i) + ".rules"); ok {

					rulesList := v.([]interface{})

					if len(rulesList) > 0 {

						var rulesPayload []map[string]interface{}

						for j := range rulesList {

							rulesMapToAdd := make(map[string]interface{})

							if v, ok := d.GetOk("approvals." + strconv.Itoa(i) + ".rules." + strconv.Itoa(j) + ".layer"); ok {
								rulesMapToAdd["layer"] = v.(string)
							}
							if v, ok := d.GetOk("approvals." + strconv.Itoa(i) + ".rules." + strconv.Itoa(j) + ".rule"); ok {
								rulesMapToAdd["rule"] = v.(string)
							}
							rulesPayload = append(rulesPayload, rulesMapToAdd)
						}
						Payload["rules"] = rulesPayload
					}
				}
				approvalsPayload = append(approvalsPayload, Payload)
			}
			guidelineCellApprovals["approvals"] = approvalsPayload
		} else {
			oldapprovals, _ := d.GetChange("approvals")
			var approvalsToDelete []interface{}
			for _, i := range oldapprovals.([]interface{}) {
				approvalsToDelete = append(approvalsToDelete, i.(map[string]interface{})["name"].(string))
			}
			guidelineCellApprovals["approvals"] = map[string]interface{}{"remove": approvalsToDelete}
		}
	}

	if ok := d.HasChange("from_type"); ok {
		guidelineCellApprovals["from-type"] = d.Get("from_type")
	}

	if ok := d.HasChange("to_type"); ok {
		guidelineCellApprovals["to-type"] = d.Get("to_type")
	}

	if ok := d.HasChange("policy_package"); ok {
		guidelineCellApprovals["policy-package"] = d.Get("policy_package")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		guidelineCellApprovals["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		guidelineCellApprovals["ignore-errors"] = v.(bool)
	}

	log.Println("Update GuidelineCellApprovals - Map = ", guidelineCellApprovals)

	updateGuidelineCellApprovalsRes, err := client.ApiCall("set-guideline-cell-approvals", guidelineCellApprovals, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !updateGuidelineCellApprovalsRes.Success {
		if updateGuidelineCellApprovalsRes.ErrorMsg != "" {
			return fmt.Errorf(updateGuidelineCellApprovalsRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementGuidelineCellApprovals(d, m)
}

func deleteManagementGuidelineCellApprovals(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	guidelineCellApprovalsPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	if v, ok := d.GetOk("delete_scope"); ok {
		guidelineCellApprovalsPayload["delete-scope"] = v.(string)
	}

	log.Println("Delete GuidelineCellApprovals")

	deleteGuidelineCellApprovalsRes, err := client.ApiCall("delete-guideline-cell-approvals", guidelineCellApprovalsPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteGuidelineCellApprovalsRes.Success {
		if deleteGuidelineCellApprovalsRes.ErrorMsg != "" {
			return fmt.Errorf(deleteGuidelineCellApprovalsRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
