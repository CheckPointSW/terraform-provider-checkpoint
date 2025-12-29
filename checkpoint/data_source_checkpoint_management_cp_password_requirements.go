package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceManagementCpPasswordRequirements() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementCpPasswordRequirementsRead,
		Schema: map[string]*schema.Schema{
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object unique identifier.",
			},
			"min_password_length": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Minimum Check Point password length.",
			},
		},
	}
}

func dataSourceManagementCpPasswordRequirementsRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}

	cpPasswordRequirementsRes, err := client.ApiCallSimple("show-cp-password-requirements", payload)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !cpPasswordRequirementsRes.Success {
		return fmt.Errorf(cpPasswordRequirementsRes.ErrorMsg)
	}
	cpPasswordRequirementsData := cpPasswordRequirementsRes.GetData()

	if v := cpPasswordRequirementsData["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := cpPasswordRequirementsData["min-password-length"]; v != nil {
		_ = d.Set("min_password_length", v)
	}

	return nil
}
