package checkpoint

import (
	"encoding/json"
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementGenericApi() *schema.Resource {
	return &schema.Resource{
		Create: createManagementGenericApi,
		Read:   readManagementGenericApi,
		Delete: deleteManagementGenericApi,
		Schema: map[string]*schema.Schema{
			"api_command": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "API command name or path",
			},
			"payload": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Request payload in JSON format",
			},
			"method": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "HTTP request method",
				Default:     "POST",
			},
			"response": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Response message in JSON format",
			},
		},
	}
}

func createManagementGenericApi(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	apiCommand := d.Get("api_command").(string)

	// Convert payload from string to map
	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("payload"); ok {
		err := json.Unmarshal([]byte(v.(string)), &payload)
		if err != nil {
			return fmt.Errorf(err.Error())
		}
	}

	var method string
	if v, ok := d.GetOk("method"); ok {
		method = v.(string)
	}

	genericApiRes, err := client.ApiCall(apiCommand, payload, client.GetSessionID(), true, client.IsProxyUsed(), method)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !genericApiRes.Success {
		return fmt.Errorf(genericApiRes.ErrorMsg)
	}

	// Convert response from map to string
	jsonResponse, err := json.Marshal(genericApiRes.GetData())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if jsonResponse != nil {
		_ = d.Set("response", string(jsonResponse))
	}

	d.SetId("generic-api-" + apiCommand + "-" + acctest.RandString(10))

	return readManagementGaiaApi(d, m)
}

func readManagementGenericApi(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementGenericApi(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
