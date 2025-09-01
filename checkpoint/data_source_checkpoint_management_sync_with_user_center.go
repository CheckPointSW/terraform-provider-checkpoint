package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceManagementSyncWIthUserCenter() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementSyncWIthUserCenterRead,
		Schema: map[string]*schema.Schema{
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object unique identifier.",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "This indicates whether the information is being synchronized with the user center once a day.",
			},
		},
	}
}

func dataSourceManagementSyncWIthUserCenterRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}

	syncWIthUserCenterRes, err := client.ApiCallSimple("show-sync-with-user-center", payload)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !syncWIthUserCenterRes.Success {
		return fmt.Errorf(syncWIthUserCenterRes.ErrorMsg)
	}
	syncWIthUserCenterData := syncWIthUserCenterRes.GetData()

	if v := syncWIthUserCenterData["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := syncWIthUserCenterData["enabled"]; v != nil {
		_ = d.Set("enabled", v)
	}

	return nil
}
