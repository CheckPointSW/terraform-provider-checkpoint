package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementLockObject() *schema.Resource {
	return &schema.Resource{
		Create: createManagementLockObject,
		Read:   readManagementLockObject,
		Delete: deleteManagementLockObject,
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
				Description: "Object unique identifier.",
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

func createManagementLockObject(d *schema.ResourceData, m interface{}) error {
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

	LockObjectRes, err := client.ApiCall("lock-object", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !LockObjectRes.Success {
		return fmt.Errorf(LockObjectRes.ErrorMsg)
	}

	lockObject := LockObjectRes.GetData()

	if lockObject["object"] != nil {
		objectMap := lockObject["object"].(map[string]interface{})

		if v, _ := objectMap["uid"]; v != nil {
			_ = d.Set("uid", v)
			d.SetId(v.(string))
		}
		if v, _ := objectMap["name"]; v != nil {
			_ = d.Set("name", v)
		}
	}

	return readManagementLockObject(d, m)
}

func readManagementLockObject(d *schema.ResourceData, m interface{}) error {

	return nil
}

func deleteManagementLockObject(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
