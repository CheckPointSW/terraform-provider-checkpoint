package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementAccessSection() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementAccessSectionRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object unique identifier.",
			},
			"layer": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Layer that the rule belongs to identified by the name or UID.",
			},
		},
	}
}

func dataSourceManagementAccessSectionRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := map[string]interface{}{
		"layer": d.Get("layer"),
	}

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showAccessSectionRes, err := client.ApiCall("show-access-section", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showAccessSectionRes.Success {
		return fmt.Errorf(showAccessSectionRes.ErrorMsg)
	}

	accessSection := showAccessSectionRes.GetData()

	log.Println("Read AccessSection - Show JSON = ", accessSection)

	if v := accessSection["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := accessSection["name"]; v != nil {
		_ = d.Set("name", v)
	}

	return nil
}
