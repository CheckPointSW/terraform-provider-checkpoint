package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementGlobalAssignment() *schema.Resource {
	return &schema.Resource{
		Create: createManagementGlobalAssignment,
		Read:   readManagementGlobalAssignment,
		Update: updateManagementGlobalAssignment,
		Delete: deleteManagementGlobalAssignment,
		Schema: map[string]*schema.Schema{
			"dependent_domain": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "N/A",
			},
			"global_access_policy": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Global domain access policy that is assigned to a dependent domain.",
			},
			"global_domain": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "N/A",
			},
			"global_threat_prevention_policy": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Global domain threat prevention policy that is assigned to a dependent domain.",
			},
			"manage_protection_actions": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "N/A",
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
			"assignment_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"assignment_up_to_date": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "The time when the assignment was assigned.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"iso_8601": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Date and time represented in international ISO 8601 format.",
						},
						"posix": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.",
						},
					},
				},
			},
		},
	}
}

func createManagementGlobalAssignment(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	globalAssignment := make(map[string]interface{})

	if v, ok := d.GetOk("dependent_domain"); ok {
		globalAssignment["dependent-domain"] = v.(string)
	}

	if v, ok := d.GetOk("global_access_policy"); ok {
		globalAssignment["global-access-policy"] = v.(string)
	}

	if v, ok := d.GetOk("global_domain"); ok {
		globalAssignment["global-domain"] = v.(string)
	}

	if v, ok := d.GetOk("global_threat_prevention_policy"); ok {
		globalAssignment["global-threat-prevention-policy"] = v.(string)
	}

	if v, ok := d.GetOkExists("manage_protection_actions"); ok {
		globalAssignment["manage-protection-actions"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		globalAssignment["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		globalAssignment["ignore-errors"] = v.(bool)
	}

	log.Println("Create GlobalAssignment - Map = ", globalAssignment)

	addGlobalAssignmentRes, err := client.ApiCall("add-global-assignment", globalAssignment, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !addGlobalAssignmentRes.Success {
		if addGlobalAssignmentRes.ErrorMsg != "" {
			return fmt.Errorf(addGlobalAssignmentRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addGlobalAssignmentRes.GetData()["uid"].(string))

	return readManagementGlobalAssignment(d, m)
}

func readManagementGlobalAssignment(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showGlobalAssignmentRes, err := client.ApiCall("show-global-assignment", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showGlobalAssignmentRes.Success {
		if objectNotFound(showGlobalAssignmentRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showGlobalAssignmentRes.ErrorMsg)
	}

	globalAssignment := showGlobalAssignmentRes.GetData()

	log.Println("Read GlobalAssignment - Show JSON = ", globalAssignment)

	if v := globalAssignment["dependent-domain"]; v != nil {
		_ = d.Set("dependent_domain", v)
	}

	if v := globalAssignment["global-access-policy"]; v != nil {
		_ = d.Set("global_access_policy", v)
	}

	if v := globalAssignment["global-domain"]; v != nil {
		_ = d.Set("global_domain", v)
	}

	if v := globalAssignment["global-threat-prevention-policy"]; v != nil {
		_ = d.Set("global_threat_prevention_policy", v)
	}

	if v := globalAssignment["manage-protection-actions"]; v != nil {
		_ = d.Set("manage_protection_actions", v)
	}

	if v := globalAssignment["assignment-status"]; v != nil {
		_ = d.Set("assignment_status", v)
	}

	if globalAssignment["assignment-up-to-date"] != nil {
		assignmentUpToDateMap := globalAssignment["assignment-up-to-date"].(map[string]interface{})
		assignmentUpToDateMapToReturn := make(map[string]interface{})

		if v, _ := assignmentUpToDateMap["iso-8601"]; v != nil {
			assignmentUpToDateMapToReturn["iso_8601"] = v
		}
		if v, _ := assignmentUpToDateMap["posix"]; v != nil {
			assignmentUpToDateMapToReturn["posix"] = v
		}

		_ = d.Set("assignment_up_to_date", assignmentUpToDateMapToReturn)
	} else {
		_ = d.Set("assignment_up_to_date", nil)
	}

	return nil

}

func updateManagementGlobalAssignment(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	globalAssignment := make(map[string]interface{})

	if ok := d.HasChange("dependent_domain"); ok {
		globalAssignment["dependent-domain"] = d.Get("dependent_domain")
	}

	if ok := d.HasChange("global_access_policy"); ok {
		globalAssignment["global-access-policy"] = d.Get("global_access_policy")
	}

	if ok := d.HasChange("global_domain"); ok {
		globalAssignment["global-domain"] = d.Get("global_domain")
	}

	if ok := d.HasChange("global_threat_prevention_policy"); ok {
		globalAssignment["global-threat-prevention-policy"] = d.Get("global_threat_prevention_policy")
	}

	if v, ok := d.GetOkExists("manage_protection_actions"); ok {
		globalAssignment["manage-protection-actions"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		globalAssignment["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		globalAssignment["ignore-errors"] = v.(bool)
	}

	log.Println("Update GlobalAssignment - Map = ", globalAssignment)

	updateGlobalAssignmentRes, err := client.ApiCall("set-global-assignment", globalAssignment, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !updateGlobalAssignmentRes.Success {
		if updateGlobalAssignmentRes.ErrorMsg != "" {
			return fmt.Errorf(updateGlobalAssignmentRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementGlobalAssignment(d, m)
}

func deleteManagementGlobalAssignment(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	globalAssignmentPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete GlobalAssignment")

	deleteGlobalAssignmentRes, err := client.ApiCall("delete-global-assignment", globalAssignmentPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteGlobalAssignmentRes.Success {
		if deleteGlobalAssignmentRes.ErrorMsg != "" {
			return fmt.Errorf(deleteGlobalAssignmentRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
