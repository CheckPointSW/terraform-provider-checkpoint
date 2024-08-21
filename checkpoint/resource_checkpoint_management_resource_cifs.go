package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"

	"strconv"
)

func resourceManagementResourceCifs() *schema.Resource {
	return &schema.Resource{
		Create: createManagementResourceCifs,
		Read:   readManagementResourceCifs,
		Update: updateManagementResourceCifs,
		Delete: deleteManagementResourceCifs,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"allowed_disk_and_print_shares": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "The list of Allowed Disk and Print Shares. Must be added in pairs.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"server_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Blocks the ability to remotely manipulate a the window's registry.",
						},
						"share_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Disk shares.",
						},
					},
				},
			},
			"log_mapped_shares": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Logs each share map attempt.",
				Default:     false,
			},
			"log_access_violation": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Logs any attempt to violate the restrictions imposed by the Resource.",
				Default:     false,
			},
			"block_remote_registry_access": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Blocks the ability to remotely manipulate a the window's registry.",
				Default:     true,
			},
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"color": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Color of the object. Should be one of existing colors.",
				Default:     "black",
			},
			"comments": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comments string.",
			},
			"ignore_warnings": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Apply changes ignoring warnings.",
				Default:     false,
			},
			"ignore_errors": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.",
				Default:     false,
			},
		},
	}
}

func createManagementResourceCifs(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	resourceCifs := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		resourceCifs["name"] = v.(string)
	}

	if v, ok := d.GetOk("allowed_disk_and_print_shares"); ok {

		allowedDiskAndPrintSharesList := v.([]interface{})

		if len(allowedDiskAndPrintSharesList) > 0 {

			var allowedDiskAndPrintSharesPayload []map[string]interface{}

			for i := range allowedDiskAndPrintSharesList {

				Payload := make(map[string]interface{})

				if v, ok := d.GetOk("allowed_disk_and_print_shares." + strconv.Itoa(i) + ".server_name"); ok {
					Payload["server-name"] = v.(string)
				}
				if v, ok := d.GetOk("allowed_disk_and_print_shares." + strconv.Itoa(i) + ".share_name"); ok {
					Payload["share-name"] = v.(string)
				}
				allowedDiskAndPrintSharesPayload = append(allowedDiskAndPrintSharesPayload, Payload)
			}
			resourceCifs["allowed-disk-and-print-shares"] = allowedDiskAndPrintSharesPayload
		}
	}

	if v, ok := d.GetOkExists("log_mapped_shares"); ok {
		resourceCifs["log-mapped-shares"] = v.(bool)
	}

	if v, ok := d.GetOkExists("log_access_violation"); ok {
		resourceCifs["log-access-violation"] = v.(bool)
	}

	if v, ok := d.GetOkExists("block_remote_registry_access"); ok {
		resourceCifs["block-remote-registry-access"] = v.(bool)
	}

	if v, ok := d.GetOk("tags"); ok {
		resourceCifs["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		resourceCifs["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		resourceCifs["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		resourceCifs["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		resourceCifs["ignore-errors"] = v.(bool)
	}

	log.Println("Create ResourceCifs - Map = ", resourceCifs)

	addResourceCifsRes, err := client.ApiCall("add-resource-cifs", resourceCifs, client.GetSessionID(), true, false)
	if err != nil || !addResourceCifsRes.Success {
		if addResourceCifsRes.ErrorMsg != "" {
			return fmt.Errorf(addResourceCifsRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addResourceCifsRes.GetData()["uid"].(string))

	return readManagementResourceCifs(d, m)
}

func readManagementResourceCifs(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
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

func updateManagementResourceCifs(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	resourceCifs := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		resourceCifs["name"] = oldName
		resourceCifs["new-name"] = newName
	} else {
		resourceCifs["name"] = d.Get("name")
	}

	if d.HasChange("allowed_disk_and_print_shares") {

		if v, ok := d.GetOk("allowed_disk_and_print_shares"); ok {

			allowedDiskAndPrintSharesList := v.([]interface{})

			var allowedDiskAndPrintSharesPayload []map[string]interface{}

			for i := range allowedDiskAndPrintSharesList {

				Payload := make(map[string]interface{})

				localMap := allowedDiskAndPrintSharesList[i].(map[string]interface{})

				if v := localMap["server_name"]; v != nil {
					Payload["server-name"] = v
				}

				if v := localMap["share_name"]; v != nil {
					Payload["share-name"] = v
				}

				allowedDiskAndPrintSharesPayload = append(allowedDiskAndPrintSharesPayload, Payload)
			}
			resourceCifs["allowed-disk-and-print-shares"] = allowedDiskAndPrintSharesPayload
		} else {
			oldallowedDiskAndPrintShares, _ := d.GetChange("allowed_disk_and_print_shares")
			var allowedDiskAndPrintSharesToDelete []interface{}
			for _, i := range oldallowedDiskAndPrintShares.([]interface{}) {
				allowedDiskAndPrintSharesToDelete = append(allowedDiskAndPrintSharesToDelete, i.(map[string]interface{})["name"].(string))
			}
			resourceCifs["allowed-disk-and-print-shares"] = map[string]interface{}{"remove": allowedDiskAndPrintSharesToDelete}
		}
	}

	if v, ok := d.GetOkExists("log_mapped_shares"); ok {
		resourceCifs["log-mapped-shares"] = v.(bool)
	}

	if v, ok := d.GetOkExists("log_access_violation"); ok {
		resourceCifs["log-access-violation"] = v.(bool)
	}

	if v, ok := d.GetOkExists("block_remote_registry_access"); ok {
		resourceCifs["block-remote-registry-access"] = v.(bool)
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			resourceCifs["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			resourceCifs["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		resourceCifs["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		resourceCifs["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		resourceCifs["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		resourceCifs["ignore-errors"] = v.(bool)
	}

	log.Println("Update ResourceCifs - Map = ", resourceCifs)

	updateResourceCifsRes, err := client.ApiCall("set-resource-cifs", resourceCifs, client.GetSessionID(), true, false)
	if err != nil || !updateResourceCifsRes.Success {
		if updateResourceCifsRes.ErrorMsg != "" {
			return fmt.Errorf(updateResourceCifsRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementResourceCifs(d, m)
}

func deleteManagementResourceCifs(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	resourceCifsPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete ResourceCifs")

	deleteResourceCifsRes, err := client.ApiCall("delete-resource-cifs", resourceCifsPayload, client.GetSessionID(), true, false)
	if err != nil || !deleteResourceCifsRes.Success {
		if deleteResourceCifsRes.ErrorMsg != "" {
			return fmt.Errorf(deleteResourceCifsRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
