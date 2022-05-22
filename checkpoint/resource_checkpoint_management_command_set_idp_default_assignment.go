package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementSetIdpDefaultAssignment() *schema.Resource {
	return &schema.Resource{
		Create: createManagementSetIdpDefaultAssignment,
		Read:   readManagementSetIdpDefaultAssignment,
		Delete: deleteManagementSetIdpDefaultAssignment,
		Schema: map[string]*schema.Schema{
			"identity_provider": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Represents the Identity Provider to be used for Login by this assignment identified by the name or UID, to cancel existing assignment should set to 'none'.",
			},
			"ignore_warnings": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Apply changes ignoring warnings.",
			},
			"ignore_errors": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.",
			},
			"task_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Command asynchronous task unique identifier.",
			},
		},
	}
}

func createManagementSetIdpDefaultAssignment(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("identity_provider"); ok {
		payload["identity-provider"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		payload["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		payload["ignore-errors"] = v.(bool)
	}

	SetIdpDefaultAssignmentRes, _ := client.ApiCall("set-idp-default-assignment", payload, client.GetSessionID(), true, false)
	if !SetIdpDefaultAssignmentRes.Success {
		return fmt.Errorf(SetIdpDefaultAssignmentRes.ErrorMsg)
	}

	d.SetId("set-idp-default-assignment" + acctest.RandString(10))
	_ = d.Set("task_id", resolveTaskId(SetIdpDefaultAssignmentRes.GetData()))
	return readManagementSetIdpDefaultAssignment(d, m)
}

func readManagementSetIdpDefaultAssignment(d *schema.ResourceData, m interface{}) error {

	return nil
}

func deleteManagementSetIdpDefaultAssignment(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
