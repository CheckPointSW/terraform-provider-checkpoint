package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementTestSicStatus() *schema.Resource {
	return &schema.Resource{
		Create: createManagementTestSicStatus,
		Read:   readManagementTestSicStatus,
		Delete: deleteManagementTestSicStatus,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Gateway, cluster member or Check Point host name.",
			},
			"sic_message": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "SIC message from the gateway.",
			},
			"sic_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "SIC (Secure Internal Communication) name.",
			},
			"sic_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "SIC Status reflects the state of the gateway after it has received the certificate issued by the ICA.",
			},
			"task_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Command asynchronous task unique identifier.",
			},
		},
	}
}

func createManagementTestSicStatus(d *schema.ResourceData, m interface{}) error {
	return readManagementTestSicStatus(d, m)
}

func readManagementTestSicStatus(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("name"); ok {
		payload["name"] = v.(string)
	}

	TestSicStatusRes, _ := client.ApiCall("test-sic-status", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if !TestSicStatusRes.Success {
		return fmt.Errorf(TestSicStatusRes.ErrorMsg)
	}

	d.SetId("test-sic-status" + acctest.RandString(10))
	_ = d.Set("task_id", resolveTaskId(TestSicStatusRes.GetData()))
	testSicStatusProfile := TestSicStatusRes.GetData()

	log.Println("Read TestSicStatus - Show JSON = ", testSicStatusProfile)

	if v := testSicStatusProfile["sic-message"]; v != nil {
		_ = d.Set("sic_message", v)
	}

	if v := testSicStatusProfile["sic-name"]; v != nil {
		_ = d.Set("sic_name", v)
	}

	if v := testSicStatusProfile["sic-status"]; v != nil {
		_ = d.Set("sic_status", v)
	}
	return nil
}

func deleteManagementTestSicStatus(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
