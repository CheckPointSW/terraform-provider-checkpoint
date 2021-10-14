package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementNatSection() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementNatSectionRead,
		Schema: map[string]*schema.Schema{
			"package": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the package.",
			},
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
		},
	}
}

func dataSourceManagementNatSectionRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := map[string]interface{}{
		"package": d.Get("package"),
	}

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showNatSectionRes, err := client.ApiCall("show-nat-section", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showNatSectionRes.Success {
		return fmt.Errorf(showNatSectionRes.ErrorMsg)
	}

	natSection := showNatSectionRes.GetData()

	log.Println("Read NatSection - Show JSON = ", natSection)

	if v := natSection["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := natSection["name"]; v != nil {
		_ = d.Set("name", v)
	}

	return nil
}
