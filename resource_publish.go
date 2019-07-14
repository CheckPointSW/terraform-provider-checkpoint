package main
import (
	chkp "api_go_sdk/APIFiles"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)


func resourcePublish() *schema.Resource {
	return &schema.Resource{
		Create: createPublish,
		Read:   readPublish,
		Delete: deletePublish,
		Schema: map[string]*schema.Schema{
			"uid": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: "Session unique identifier. Specify it to publish a different session than the one you currently use.",
			},
		},
	}
}

func createPublish(d *schema.ResourceData, m interface{}) error {
	client := m.(*chkp.ApiClient)
	var uid string
	var payload = make(map[string]interface{})

	if v, ok := d.GetOk("uid"); ok {
		uid = v.(string)
		payload["uid"] = uid
		log.Println("publish other session uid - ", uid)
	} else {
		// Publish current session
		s, err := GetSid()
		if err != nil {
			return err
		}
		uid = s.Uid
		log.Println("publish current session uid - ", uid)
	}
	publishRes, _ := client.ApiCall("publish",payload,client.GetSessionID(),true,false)
	if !publishRes.Success {
		return fmt.Errorf(publishRes.ErrorMsg)
	}
	// Set Schema UID = Session UID
	d.SetId(uid)
	return readPublish(d, m)
}

func readPublish(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deletePublish(d *schema.ResourceData, m interface{}) error {
	d.SetId("") // Destroy resource
	return nil
}
