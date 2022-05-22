package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementSetIdpToDomainAssignment() *schema.Resource {
	return &schema.Resource{
		Create: createManagementSetIdpToDomainAssignment,
		Read:   readManagementSetIdpToDomainAssignment,
		Delete: deleteManagementSetIdpToDomainAssignment,
		Schema: map[string]*schema.Schema{
			"assigned_domain": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Represents the Domain assigned by 'idp-to-domain-assignment', need to be domain name or UID.",
			},
			"identity_provider": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Represents the Identity Provider to be used for Login by this assignment. Must be set when \"using-default\" was set to be false.",
			},
			"using_default": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Is this assignment override by 'idp-default-assignment'.",
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

func createManagementSetIdpToDomainAssignment(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("assigned_domain"); ok {
		payload["assigned-domain"] = v.(string)
	}

	if v, ok := d.GetOk("identity_provider"); ok {
		payload["identity-provider"] = v.(string)
	}

	if v, ok := d.GetOkExists("using_default"); ok {
		payload["using-default"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		payload["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		payload["ignore-errors"] = v.(bool)
	}

	SetIdpToDomainAssignmentRes, _ := client.ApiCall("set-idp-to-domain-assignment", payload, client.GetSessionID(), true, false)
	if !SetIdpToDomainAssignmentRes.Success {
		return fmt.Errorf(SetIdpToDomainAssignmentRes.ErrorMsg)
	}

	d.SetId("set-idp-to-domain-assignment" + acctest.RandString(10))
	_ = d.Set("task_id", resolveTaskId(SetIdpToDomainAssignmentRes.GetData()))
	return readManagementSetIdpToDomainAssignment(d, m)
}

func readManagementSetIdpToDomainAssignment(d *schema.ResourceData, m interface{}) error {

	return nil
}

func deleteManagementSetIdpToDomainAssignment(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
