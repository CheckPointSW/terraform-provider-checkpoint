package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementDeleteDataCenterObject() *schema.Resource {
	return &schema.Resource{
		Create:             createManagementDeleteDataCenterObject,
		Read:               readManagementDeleteDataCenterObject,
		Delete:             deleteManagementDeleteDataCenterObject,
		DeprecationMessage: "This resource is deprecated. please use the `checkpoint_management_data_center_object` resource.",
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Object name.",
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

func createManagementDeleteDataCenterObject(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("name"); ok {
		payload["name"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		payload["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		payload["ignore-errors"] = v.(bool)
	}

	DeleteDataCenterObjectRes, _ := client.ApiCall("delete-data-center-object", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if !DeleteDataCenterObjectRes.Success {
		return fmt.Errorf(DeleteDataCenterObjectRes.ErrorMsg)
	}

	d.SetId("delete-data-center-object-" + acctest.RandString(10))
	return readManagementDeleteDataCenterObject(d, m)
}

func readManagementDeleteDataCenterObject(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementDeleteDataCenterObject(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
