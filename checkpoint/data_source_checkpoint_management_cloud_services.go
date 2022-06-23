package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementCloudServices() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementCloudServicesRead,
		Schema: map[string]*schema.Schema{
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status of the connection to the Infinity Portal.",
			},
			"connected_at": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "The time of the connection between the Management Server and the Infinity Portal.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"iso_8601": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Date and time represented in international ISO 8601 format.",
						},
						"posix": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.",
						},
					},
				},
			},
			"management_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Management Server's public URL.",
			},
		},
	}
}

func dataSourceManagementCloudServicesRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	showCloudServices, err := client.ApiCall("show-cloud-services", make(map[string]interface{}), client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showCloudServices.Success {
		return fmt.Errorf(showCloudServices.ErrorMsg)
	}

	showCloudServicesRes := showCloudServices.GetData()

	log.Println("Show Cloud Services - JSON = ", showCloudServicesRes)

	if v := showCloudServicesRes["status"]; v != nil {
		_ = d.Set("status", v)
	}else{
		_ = d.Set("status", nil)
	}

	if v := showCloudServicesRes["connected-at"]; v != nil {
		if connectedAtShow, ok := showCloudServicesRes["connected-at"].(map[string]interface{}); ok {
			connectedAtState := make(map[string]interface{})
			if v := connectedAtShow["iso-8601"]; v != nil {
				connectedAtState["iso_8601"] = v
			}
			if v := connectedAtShow["posix"]; v != nil {
				connectedAtState["posix"] = v
			}
			_ = d.Set("connected_at", connectedAtState)
		}
	}else{
		_ = d.Set("connected_at", nil)
	}

	if v := showCloudServicesRes["management-url"]; v != nil {
		_ = d.Set("management_url", v)
	}else{
		_ = d.Set("management_url", nil)
	}

	d.SetId("show-cloud-services-" + acctest.RandString(5))

	return nil
}
