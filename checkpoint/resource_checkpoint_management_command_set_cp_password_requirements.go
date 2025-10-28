package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementSetCpPasswordRequirements() *schema.Resource {
	return &schema.Resource{
		Create: createManagementSetCpPasswordRequirements,
		Read:   readManagementSetCpPasswordRequirements,
		Delete: deleteManagementSetCpPasswordRequirements,
		Schema: map[string]*schema.Schema{
			"min_password_length": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "Minimum Check Point password length.",
			},
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				ForceNew:    true,
				Description: "Object unique identifier.",
			},
		},
	}
}

func createManagementSetCpPasswordRequirements(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("min_password_length"); ok {
		payload["min-password-length"] = v.(int)
	}

	SetCpPasswordRequirementsRes, err := client.ApiCall("set-cp-password-requirements", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !SetCpPasswordRequirementsRes.Success {
		return fmt.Errorf(SetCpPasswordRequirementsRes.ErrorMsg)
	}

	cpPasswordRequirements := SetCpPasswordRequirementsRes.GetData()

	_ = d.Set("uid", cpPasswordRequirements["uid"])
	d.SetId(cpPasswordRequirements["uid"].(string))

	return nil
}

func readManagementSetCpPasswordRequirements(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementSetCpPasswordRequirements(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
