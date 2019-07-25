package main

import (
	chkp "api_go_sdk/APIFiles"
	"fmt"
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

func putFileParseSchemaToMap(d *schema.ResourceData, createResource bool) map[string]interface{} {
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

	if !createResource {
		// Call from updatePutFile
		// Remove attributes that cannot be set from map - Schema contain ADD + SET attr.
		delete(putFileMap,"set-if-exists")
	}
	return putFileMap
}

func createPutFile(d *schema.ResourceData, m interface{}) error {
	log.Println("Enter createPutFile...")
	client := m.(*chkp.ApiClient)
	payload := putFileParseSchemaToMap(d, true)
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
	log.Println("Enter readPutFile...")
	//client := m.(*chkp.ApiClient)
	//payload := map[string]interface{}{
	//	"name": d.Get("name"),
	//}
	//showPIRes, _ := client.ApiCall("show-physical-interface",payload,client.GetSessionID(),true,false)
	//if !showPIRes.Success {
	//	// Handle deletion of an object from other clients - Object not found
	//	if objectNotFound(showPIRes.GetData()["code"].(string)) {
	//		d.SetId("") // Destroy resource
	//		return nil
	//	}
	//	return fmt.Errorf(showPIRes.ErrorMsg)
	//}
	//PIJson := showPIRes.GetData()
	//log.Println(PIJson)
	//
	//if _, ok := d.GetOk("name"); ok {
	//	_ = d.Set("name", PIJson["name"].(string))
	//}

	log.Println("Exit readPutFile...")
	return nil
}

func updatePutFile(d *schema.ResourceData, m interface{}) error {
	log.Println("Enter updatePutFile...")
	client := m.(*chkp.ApiClient)
	payload := putFileParseSchemaToMap(d, false)
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
