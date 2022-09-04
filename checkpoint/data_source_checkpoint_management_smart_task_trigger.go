package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementSmartTaskTrigger() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementSmartTaskTriggerRead,
		Schema: map[string]*schema.Schema{
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object unique identifier.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation of the object.",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object type.",
			},
			"before_operation": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether or not this trigger is fired before an operation.",
			},
		},
	}
}

func dataSourceManagementSmartTaskTriggerRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showSmartTaskTriggerRes, err := client.ApiCall("show-smart-task-trigger", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		fmt.Errorf(err.Error())
	}
	if !showSmartTaskTriggerRes.Success {
		fmt.Errorf(showSmartTaskTriggerRes.ErrorMsg)
	}

	smartTaskTrigger := showSmartTaskTriggerRes.GetData()

	log.Println("Read Smart Task Trigger - Show JSON = ", smartTaskTrigger)

	d.SetId("show-smart-task-trigger-" + acctest.RandString(10))

	if v := smartTaskTrigger["uid"]; v != nil {
		_ = d.Set("uid", v)
	}

	if v := smartTaskTrigger["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := smartTaskTrigger["type"]; v != nil {
		_ = d.Set("type", v)
	}

	if v := smartTaskTrigger["before-operation"]; v != nil {
		_ = d.Set("before_operation", v)
	}

	return nil
}
