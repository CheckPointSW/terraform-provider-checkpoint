package checkpoint

import (
	"fmt"
	chkp "github.com/Checkpoint/api_go_sdk/APIFiles"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)


func resourcePutFile() *schema.Resource {
	return &schema.Resource{
		Create: createPutFile,
		Read:   readPutFile,
		Update: updatePutFile,
		Delete: deletePutFile,
		Schema: map[string]*schema.Schema{
			"file_name": {
				Type: schema.TypeString,
				Required: true,
				Description: "Filename include the desired path. The file will be created in the user home directory if the full path wasn't provided",
			},
			"text_content": {
				Type: schema.TypeString,
				Required: true,
				Description: "",
			},
			"override": {
				Type: schema.TypeBool,
				Optional: true,
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
	log.Println("Enter createPutFile...")
	client := m.(*chkp.ApiClient)
	payload := putFileParseSchemaToMap(d)
	log.Println(payload)
	setPIRes, _ := client.ApiCall("put-file",payload,client.GetSessionID(),true,false)
	if !setPIRes.Success {
		return fmt.Errorf(setPIRes.ErrorMsg)
	}

	// Set Schema UID = Object key
	d.SetId(payload["file-name"].(string))

	log.Println("Exit createPutFile...")
	return readPutFile(d, m)
}

func readPutFile(d *schema.ResourceData, m interface{}) error {
	return nil
}

func updatePutFile(d *schema.ResourceData, m interface{}) error {
	log.Println("Enter updatePutFile...")
	client := m.(*chkp.ApiClient)
	payload := putFileParseSchemaToMap(d)
	setNetworkRes, _ := client.ApiCall("put-file",payload,client.GetSessionID(),true,false)
	if !setNetworkRes.Success {
		return fmt.Errorf(setNetworkRes.ErrorMsg)
	}
	log.Println("Exit updatePutFile...")
	return readPutFile(d, m)
}

func deletePutFile(d *schema.ResourceData, m interface{}) error {
	log.Println("Enter deletePutFile...")
	d.SetId("") // Destroy resource
	log.Println("Exit deletePutFile...")
	return nil
}
