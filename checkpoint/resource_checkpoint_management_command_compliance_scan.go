package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementComplianceScan() *schema.Resource {
	return &schema.Resource{
		Create: createManagementComplianceScan,
		Read:   readManagementComplianceScan,
		Delete: deleteManagementComplianceScan,
		Schema: map[string]*schema.Schema{
			"task_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Compliance scan task UID.",
			},
		},
	}
}

func createManagementComplianceScan(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	ComplianceScanRes, err := client.ApiCall("compliance-scan", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !ComplianceScanRes.Success {
		return fmt.Errorf(ComplianceScanRes.ErrorMsg)
	}

	d.SetId("compliance-scan-" + acctest.RandString(10))
	_ = d.Set("task_id", resolveTaskId(ComplianceScanRes.GetData()))

	return nil
}

func readManagementComplianceScan(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementComplianceScan(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
