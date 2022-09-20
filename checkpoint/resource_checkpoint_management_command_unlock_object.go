package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementUnlockObject() *schema.Resource {
	return &schema.Resource{
		Create: createManagementUnlockObject,
		Read:   readManagementUnlockObject,
		Delete: deleteManagementUnlockObject,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Object name. Must be unique in the domain.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Object unique identifier. When using uid, there is no need to use name/type parameters",
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Object type.",
			},
			"layer": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Object layer, need to specify the layer if the object is rule/section and uid is not supplied.",
			},
		},
	}
}

func createManagementUnlockObject(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("name"); ok {
		payload["name"] = v.(string)
	} else if v, ok := d.GetOk("uid"); ok {
		payload["uid"] = v.(string)
	}

	if v, ok := d.GetOk("type"); ok {
		payload["type"] = v.(string)
	}

	if v, ok := d.GetOk("layer"); ok {
		payload["layer"] = v.(string)
	}

	UnlockObjectRes, err := client.ApiCall("unlock-object", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !UnlockObjectRes.Success {
		return fmt.Errorf(UnlockObjectRes.ErrorMsg)
	}

	unlockObject := UnlockObjectRes.GetData()

	if unlockObject["object"] != nil {
		objectMap := unlockObject["object"].(map[string]interface{})

		if v, _ := objectMap["uid"]; v != nil {
			_ = d.Set("uid", v)
			d.SetId(v.(string))
		}
		if v, _ := objectMap["name"]; v != nil {
			_ = d.Set("name", v)
		}
	}

	return readManagementUnlockObject(d, m)
}

func readManagementUnlockObject(d *schema.ResourceData, m interface{}) error {

	return nil
}

func deleteManagementUnlockObject(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}