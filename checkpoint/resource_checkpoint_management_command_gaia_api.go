package checkpoint

import (
	"encoding/json"
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementGaiaApi() *schema.Resource {
	return &schema.Resource{
		Create: createManagementGaiaApi,
		Read:   readManagementGaiaApi,
		Delete: deleteManagementGaiaApi,
		Schema: map[string]*schema.Schema{
			"target": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Gateway-object-name or gateway-ip-address or gateway-UID.",
			},
			"other_parameter": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Other input parameters that gateway needs it.",
			},
			"command_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Target's api command.",
			},
			"response_message": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Response's object from the target in json format.\n",
			},
		},
	}
}

func createManagementGaiaApi(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("target"); ok {
		payload["target"] = v.(string)
	}

	if v, ok := d.GetOk("other_parameter"); ok {
		payload["other-parameter"] = v.(string)
	}

	commandName := "gaia-api/" + d.Get("command_name").(string)

	GaiaApiRes, err := client.ApiCall(commandName, payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !GaiaApiRes.Success {
		return fmt.Errorf(GaiaApiRes.ErrorMsg)
	}

	gaiaApi := GaiaApiRes.GetData()

	if v := gaiaApi["command-name"]; v != nil {
		_ = d.Set("command_name", v)
	}

	if v := gaiaApi["response-message"]; v != nil {
		valToReturn, err := json.Marshal(v)

		if err != nil {
			log.Println(err.Error())
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
