package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementDeleteInfinityIdpObject() *schema.Resource {
	return &schema.Resource{
		Create: createManagementDeleteInfinityIdpObject,
		Read:   readManagementDeleteInfinityIdpObject,
		Delete: deleteManagementDeleteInfinityIdpObject,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Object name.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Object UID.",
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
		},
	}
}

func createManagementDeleteInfinityIdpObject(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("name"); ok {
		payload["name"] = v.(string)
	} else {
		if v, ok := d.GetOk("uid"); ok {
			payload["uid"] = v.(string)
		}
	}
	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		payload["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		payload["ignore-errors"] = v.(bool)
	}

	DeleteInfinityIdpObjectRes, _ := client.ApiCall("delete-infinity-idp-object", payload, client.GetSessionID(), true, false)
	if !DeleteInfinityIdpObjectRes.Success {
		return fmt.Errorf(DeleteInfinityIdpObjectRes.ErrorMsg)
	}

	d.SetId("delete-infinity-idp-object-" + acctest.RandString(10))

	return readManagementDeleteInfinityIdpObject(d, m)
}

func readManagementDeleteInfinityIdpObject(d *schema.ResourceData, m interface{}) error {

	return nil
}

func deleteManagementDeleteInfinityIdpObject(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
