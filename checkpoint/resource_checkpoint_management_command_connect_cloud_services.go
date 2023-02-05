package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementConnectCloudServices() *schema.Resource {
	return &schema.Resource{
		Create: createManagementConnectCloudServices,
		Read:   readManagementConnectCloudServices,
		Delete: deleteManagementConnectCloudServices,
		Schema: map[string]*schema.Schema{
			"auth_token": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Copy the authentication token from the Smart-1 cloud service hosted in the Infinity Portal.",
			},
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

func createManagementConnectCloudServices(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}

	if v, ok := d.GetOk("auth_token"); ok {
		payload["auth-token"] = v.(string)
	}

	ConnectCloudServicesRes, err := client.ApiCall("connect-cloud-services", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !ConnectCloudServicesRes.Success {
		return fmt.Errorf(ConnectCloudServicesRes.ErrorMsg)
	}

	connectCloudServicesRes := ConnectCloudServicesRes.GetData()

	log.Println("Connect Cloud Services - JSON = ", connectCloudServicesRes)

	d.SetId("connect-cloud-services" + acctest.RandString(5))

	if v := connectCloudServicesRes["status"]; v != nil {
		_ = d.Set("status", v)
	} else {
		_ = d.Set("status", nil)
	}

	if v := connectCloudServicesRes["connected-at"]; v != nil {
		connectedAtShow := connectCloudServicesRes["connected-at"].(map[string]interface{})
		connectedAtState := make(map[string]interface{})
		if v := connectedAtShow["iso-8601"]; v != nil {
			connectedAtState["iso_8601"] = v
		}
		if v := connectedAtShow["posix"]; v != nil {
			connectedAtState["posix"] = v
		}
		_ = d.Set("connected_at", connectedAtState)
	} else {
		_ = d.Set("connected_at", nil)
	}

	if v := connectCloudServicesRes["management-url"]; v != nil {
		_ = d.Set("management_url", v)
	} else {
		_ = d.Set("management_url", nil)
	}

	return readManagementConnectCloudServices(d, m)
}

func readManagementConnectCloudServices(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementConnectCloudServices(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
