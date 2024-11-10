package checkpoint

import (
	"encoding/json"
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementGaiaApi() *schema.Resource {
	return &schema.Resource{
		Create: createManagementGaiaApi,
		Read:   readManagementGaiaApi,
		Delete: deleteManagementGaiaApi,
		Schema: map[string]*schema.Schema{
			"command_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "GAIA API command name or path",
			},
			"target": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Gateway object name or Gateway IP address or Gateway UID",
			},
			"other_parameter": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Other input parameters for the request payload in JSON format",
			},
			"response_message": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Response message in JSON format",
			},
		},
	}
}

func createManagementGaiaApi(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}

	if v, ok := d.GetOk("other_parameter"); ok {
		err := json.Unmarshal([]byte(v.(string)), &payload)
		if err != nil {
			return fmt.Errorf(err.Error())
		}
	}

	if v, ok := d.GetOk("target"); ok {
		payload["target"] = v.(string)
	}

	commandName := "gaia-api/" + d.Get("command_name").(string)

	GaiaApiRes, err := client.ApiCall(commandName, payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !GaiaApiRes.Success {
		return fmt.Errorf(GaiaApiRes.ErrorMsg)
	}

	gaiaApiResponse := GaiaApiRes.GetData()

	if v := gaiaApiResponse["response-message"]; v != nil {
		valToReturn, err := json.Marshal(v)
		if err != nil {
			return fmt.Errorf(err.Error())
		}
		_ = d.Set("response_message", string(valToReturn))
	}

	d.SetId("gaia-api-" + acctest.RandString(10))
	return readManagementGaiaApi(d, m)
}

func readManagementGaiaApi(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementGaiaApi(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
