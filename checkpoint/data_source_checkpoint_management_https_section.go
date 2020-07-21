package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementHttpsSection() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementHttpsSectionRead,
		Schema: map[string]*schema.Schema{
			"layer": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Layer that holds the Object. Identified by the Name or UID.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object unique identifier.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
		},
	}
}

func dataSourceManagementHttpsSectionRead(d *schema.ResourceData, m interface{}) error {

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

	showHttpsSectionRes, err := client.ApiCall("show-https-section", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showHttpsSectionRes.Success {
		return fmt.Errorf(showHttpsSectionRes.ErrorMsg)
	}

	httpsSection := showHttpsSectionRes.GetData()

	log.Println("Read HttpsSection - Show JSON = ", httpsSection)

	if v := httpsSection["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := httpsSection["name"]; v != nil {
		_ = d.Set("name", v)
	}

	return nil
}
