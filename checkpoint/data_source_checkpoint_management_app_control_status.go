package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementAppControlStatus() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementAppControlStatusRead,
		Schema: map[string]*schema.Schema{
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object unique identifier.",
			},
			"last_updated": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The last time Application Control & URL Filtering was updated on the management server.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"iso_8601": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Date and time represented in international ISO 8601 format.",
						},
						"posix": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.",
						},
					},
				},
			},
			"installed_version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Installed Application Control & URL Filtering version.",
			},
			"installed_version_creation_time": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Installed Application Control & URL Filtering version creation time.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"iso_8601": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Date and time represented in international ISO 8601 format.",
						},
						"posix": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.",
						},
					},
				},
			},
		},
	}
}

func dataSourceManagementAppControlStatusRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	payload := make(map[string]interface{})

	showAppControlStatusRes, err := client.ApiCallSimple("show-app-control-status", payload)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showAppControlStatusRes.Success {
		return fmt.Errorf(showAppControlStatusRes.ErrorMsg)
	}

	appControlStatus := showAppControlStatusRes.GetData()

	log.Println("Read App Control Status - Show JSON = ", appControlStatus)

	if v := appControlStatus["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := appControlStatus["last-updated"]; v != nil {

		mapToReturn := make(map[string]interface{})

		innerMap := v.(map[string]interface{})

		if v := innerMap["iso-8601"]; v != nil {
			mapToReturn["iso_8601"] = v
		}
		if v := innerMap["posix"]; v != nil {
			mapToReturn["posix"] = v
		}

		_ = d.Set("last_updated", []interface{}{mapToReturn})
	}

	if v := appControlStatus["installed-version"]; v != nil {
		_ = d.Set("installed_version", v)
	}

	if v := appControlStatus["installed-version-creation-time"]; v != nil {

		mapToReturn := make(map[string]interface{})

		innerMap := v.(map[string]interface{})

		if v := innerMap["iso-8601"]; v != nil {
			mapToReturn["iso_8601"] = v
		}
		if v := innerMap["posix"]; v != nil {
			mapToReturn["posix"] = v
		}

		_ = d.Set("installed_version_creation_time", []interface{}{mapToReturn})
	}

	return nil
}
