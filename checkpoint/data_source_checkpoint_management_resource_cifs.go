package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementResourceCifs() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementResourceCifsRead,
		Schema: map[string]*schema.Schema{
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object uid.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"allowed_disk_and_print_shares": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of Allowed Disk and Print Shares. Must be added in pairs.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"server_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Blocks the ability to remotely manipulate a the window's registry.",
						},
						"share_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Disk shares.",
						},
					},
				},
			},
			"log_mapped_shares": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Logs each share map attempt.",
			},
			"log_access_violation": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Logs any attempt to violate the restrictions imposed by the Resource.",
			},
			"block_remote_registry_access": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Blocks the ability to remotely manipulate a the window's registry.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"color": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Color of the object. Should be one of existing colors.",
			},
			"comments": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Comments string.",
			},
		},
	}
}
func dataSourceManagementResourceCifsRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showResourceCifsRes, err := client.ApiCall("show-resource-cifs", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showResourceCifsRes.Success {
		if objectNotFound(showResourceCifsRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showResourceCifsRes.ErrorMsg)
	}

	resourceCifs := showResourceCifsRes.GetData()

	log.Println("Read ResourceCifs - Show JSON = ", resourceCifs)

	if v := resourceCifs["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := resourceCifs["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if resourceCifs["allowed-disk-and-print-shares"] != nil {

		allowedDiskAndPrintSharesList, ok := resourceCifs["allowed-disk-and-print-shares"].([]interface{})

		if ok {

			if len(allowedDiskAndPrintSharesList) > 0 {

				var allowedDiskAndPrintSharesListToReturn []map[string]interface{}

				for i := range allowedDiskAndPrintSharesList {

					allowedDiskAndPrintSharesMap := allowedDiskAndPrintSharesList[i].(map[string]interface{})

					allowedDiskAndPrintSharesMapToAdd := make(map[string]interface{})

					if v, _ := allowedDiskAndPrintSharesMap["server-name"]; v != nil {
						allowedDiskAndPrintSharesMapToAdd["server_name"] = v
					}
					if v, _ := allowedDiskAndPrintSharesMap["share-name"]; v != nil {
						allowedDiskAndPrintSharesMapToAdd["share_name"] = v
					}
					allowedDiskAndPrintSharesListToReturn = append(allowedDiskAndPrintSharesListToReturn, allowedDiskAndPrintSharesMapToAdd)
				}
				_ = d.Set("allowed_disk_and_print_shares", allowedDiskAndPrintSharesListToReturn)
			}
		}
	}

	if v := resourceCifs["log-mapped-shares"]; v != nil {
		_ = d.Set("log_mapped_shares", v)
	}

	if v := resourceCifs["log-access-violation"]; v != nil {
		_ = d.Set("log_access_violation", v)
	}

	if v := resourceCifs["block-remote-registry-access"]; v != nil {
		_ = d.Set("block_remote_registry_access", v)
	}

	if resourceCifs["tags"] != nil {
		tagsJson, ok := resourceCifs["tags"].([]interface{})
		if ok {
			tagsIds := make([]string, 0)
			if len(tagsJson) > 0 {
				for _, tags := range tagsJson {
					tags := tags.(map[string]interface{})
					tagsIds = append(tagsIds, tags["name"].(string))
				}
			}
			_ = d.Set("tags", tagsIds)
		}
	} else {
		_ = d.Set("tags", nil)
	}

	if v := resourceCifs["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := resourceCifs["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := resourceCifs["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := resourceCifs["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}
