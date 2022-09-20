package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementGetInterfaces() *schema.Resource {
	return &schema.Resource{
		Create: createManagementGetInterfaces,
		Read:   readManagementGetInterfaces,
		Delete: deleteManagementGetInterfaces,
		Schema: map[string]*schema.Schema{
			"target_uid": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Target unique identifier.",
			},
			"target_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Target name.",
			},
			"group_interfaces_by_subnet": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Specify whether to group the cluster interfaces by a subnet. Otherwise, group the cluster interfaces by their names.",
			},
			"use_defined_by_routes": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Specify whether to configure the topology \"Defined by Routes\" where applicable. Otherwise, configure the topology to \"This Network\" as default for internal interfaces.",
			},
			"with_topology": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Specify whether to fetch the interfaces with their topology. Otherwise, the Management Server fetches the interfaces without their topology.",
			},
			"task_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Command asynchronous task unique identifier.",
			},
		},
	}
}

func createManagementGetInterfaces(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("target_uid"); ok {
		payload["target-uid"] = v.(string)
	}

	if v, ok := d.GetOk("target_name"); ok {
		payload["target-name"] = v.(string)
	}

	if v, ok := d.GetOkExists("group_interfaces_by_subnet"); ok {
		payload["group-interfaces-by-subnet"] = v.(bool)
	}

	if v, ok := d.GetOkExists("use_defined_by_routes"); ok {
		payload["use-defined-by-routes"] = v.(bool)
	}

	if v, ok := d.GetOkExists("with_topology"); ok {
		payload["with-topology"] = v.(bool)
	}

	GetInterfacesRes, _ := client.ApiCall("get-interfaces", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if !GetInterfacesRes.Success {
		return fmt.Errorf(GetInterfacesRes.ErrorMsg)
	}

	d.SetId("get-interfaces-" + acctest.RandString(10))
	_ = d.Set("task_id", resolveTaskId(GetInterfacesRes.GetData()))
	return readManagementGetInterfaces(d, m)
}

func readManagementGetInterfaces(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementGetInterfaces(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
