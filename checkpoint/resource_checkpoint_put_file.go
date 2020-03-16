package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourcePutFile() *schema.Resource {
	return &schema.Resource{
		Create: createPutFile,
		Read:   readPutFile,
		Update: updatePutFile,
		Delete: deletePutFile,
		Schema: map[string]*schema.Schema{
			"file_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Filename include the desired path. The file will be created in the user home directory if the full path wasn't provided",
			},
			"text_content": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "",
			},
			"override": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "comments",
			},
		},
	}
}

func putFileParseSchemaToMap(d *schema.ResourceData) map[string]interface{} {
	putFileMap := make(map[string]interface{})

	if val, ok := d.GetOk("file_name"); ok {
		putFileMap["file-name"] = val.(string)
	}
	if val, ok := d.GetOk("text_content"); ok {
		putFileMap["text-content"] = val.(string)
	}
	if val, ok := d.GetOk("override"); ok {
		putFileMap["override"] = val.(bool)
	}

	return putFileMap
}

func createPutFile(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	payload := putFileParseSchemaToMap(d)
	setPIRes, _ := client.ApiCall("put-file", payload, client.GetSessionID(), true, false)
	if !setPIRes.Success {
		return fmt.Errorf(setPIRes.ErrorMsg)
	}

	// Set Schema UID = Object key
	d.SetId(payload["file-name"].(string))

	return readPutFile(d, m)
}

func readPutFile(d *schema.ResourceData, m interface{}) error {
	return nil
}

func updatePutFile(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	payload := putFileParseSchemaToMap(d)
	setNetworkRes, _ := client.ApiCall("put-file", payload, client.GetSessionID(), true, false)
	if !setNetworkRes.Success {
		return fmt.Errorf(setNetworkRes.ErrorMsg)
	}
	return readPutFile(d, m)
}

func deletePutFile(d *schema.ResourceData, m interface{}) error {
	d.SetId("") // Destroy resource
	return nil
}
