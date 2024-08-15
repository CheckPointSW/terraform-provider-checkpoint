package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementMultipleKeyExchanges() *schema.Resource {
	return &schema.Resource{
		Create: createManagementMultipleKeyExchanges,
		Read:   readManagementMultipleKeyExchanges,
		Update: updateManagementMultipleKeyExchanges,
		Delete: deleteManagementMultipleKeyExchanges,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"key_exchange_methods": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "Key-Exchange methods to use. Can contain only Diffie-Hellman groups.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"additional_key_exchange_1_methods": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Additional Key-Exchange 1 methods to use.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"additional_key_exchange_2_methods": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Additional Key-Exchange 2 methods to use.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"additional_key_exchange_3_methods": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Additional Key-Exchange 3 methods to use.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"additional_key_exchange_4_methods": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Additional Key-Exchange 4 methods to use.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"additional_key_exchange_5_methods": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Additional Key-Exchange 5 methods to use.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"additional_key_exchange_6_methods": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Additional Key-Exchange 6 methods to use.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"additional_key_exchange_7_methods": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Additional Key-Exchange 7 methods to use.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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

func createManagementMultipleKeyExchanges(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	multipleKeyExchanges := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		multipleKeyExchanges["name"] = v.(string)
	}

	if v, ok := d.GetOk("key_exchange_methods"); ok {
		multipleKeyExchanges["key-exchange-methods"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("additional_key_exchange_1_methods"); ok {
		multipleKeyExchanges["additional-key-exchange-1-methods"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("additional_key_exchange_2_methods"); ok {
		multipleKeyExchanges["additional-key-exchange-2-methods"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("additional_key_exchange_3_methods"); ok {
		multipleKeyExchanges["additional-key-exchange-3-methods"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("additional_key_exchange_4_methods"); ok {
		multipleKeyExchanges["additional-key-exchange-4-methods"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("additional_key_exchange_5_methods"); ok {
		multipleKeyExchanges["additional-key-exchange-5-methods"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("additional_key_exchange_6_methods"); ok {
		multipleKeyExchanges["additional-key-exchange-6-methods"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("additional_key_exchange_7_methods"); ok {
		multipleKeyExchanges["additional-key-exchange-7-methods"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("tags"); ok {
		multipleKeyExchanges["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		multipleKeyExchanges["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		multipleKeyExchanges["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		multipleKeyExchanges["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		multipleKeyExchanges["ignore-errors"] = v.(bool)
	}

	log.Println("Create MultipleKeyExchanges - Map = ", multipleKeyExchanges)

	addMultipleKeyExchangesRes, err := client.ApiCall("add-multiple-key-exchanges", multipleKeyExchanges, client.GetSessionID(), true, false)
	if err != nil || !addMultipleKeyExchangesRes.Success {
		if addMultipleKeyExchangesRes.ErrorMsg != "" {
			return fmt.Errorf(addMultipleKeyExchangesRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addMultipleKeyExchangesRes.GetData()["uid"].(string))

	return readManagementMultipleKeyExchanges(d, m)
}

func readManagementMultipleKeyExchanges(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showMultipleKeyExchangesRes, err := client.ApiCall("show-multiple-key-exchanges", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showMultipleKeyExchangesRes.Success {
		if objectNotFound(showMultipleKeyExchangesRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showMultipleKeyExchangesRes.ErrorMsg)
	}

	multipleKeyExchanges := showMultipleKeyExchangesRes.GetData()

	log.Println("Read MultipleKeyExchanges - Show JSON = ", multipleKeyExchanges)

	if v := multipleKeyExchanges["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := multipleKeyExchanges["key-exchange-methods"]; v != nil {

		_ = d.Set("key_exchange_methods", v.([]interface{}))

	} else {
		_ = d.Set("key_exchange_methods", nil)
	}

	if v := multipleKeyExchanges["additional-key-exchange-1-methods"]; v != nil {
		_ = d.Set("additional_key_exchange_1_methods", v.([]interface{}))

	} else {
		_ = d.Set("additional_key_exchange_1_methods", nil)
	}

	if v := multipleKeyExchanges["additional-key-exchange-2-methods"]; v != nil {
		_ = d.Set("additional_key_exchange_2_methods", v.([]interface{}))
	} else {
		_ = d.Set("additional_key_exchange_2_methods", nil)
	}

	if v := multipleKeyExchanges["additional-key-exchange-3-methods"]; v != nil {
		_ = d.Set("additional_key_exchange_3_methods", v.([]interface{}))
	} else {
		_ = d.Set("additional_key_exchange_3_methods", nil)
	}

	if v := multipleKeyExchanges["additional-key-exchange-4-methods"]; v != nil {
		_ = d.Set("additional_key_exchange_4_methods", v.([]interface{}))
	} else {
		_ = d.Set("additional_key_exchange_4_methods", nil)
	}

	if v := multipleKeyExchanges["additional-key-exchange-5-methods"]; v != nil {
		_ = d.Set("additional_key_exchange_5_methods", v.([]interface{}))
	} else {
		_ = d.Set("additional_key_exchange_5_methods", nil)
	}

	if v := multipleKeyExchanges["additional-key-exchange-6-methods"]; v != nil {
		_ = d.Set("additional_key_exchange_6_methods", v.([]interface{}))
	} else {
		_ = d.Set("additional_key_exchange_6_methods", nil)
	}

	if v := multipleKeyExchanges["additional-key-exchange-7-methods"]; v != nil {
		_ = d.Set("additional_key_exchange_7_methods", v.([]interface{}))
	} else {
		_ = d.Set("additional_key_exchange_7_methods", nil)
	}

	if multipleKeyExchanges["tags"] != nil {
		tagsJson, ok := multipleKeyExchanges["tags"].([]interface{})
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

	if v := multipleKeyExchanges["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := multipleKeyExchanges["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := multipleKeyExchanges["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := multipleKeyExchanges["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementMultipleKeyExchanges(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	multipleKeyExchanges := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		multipleKeyExchanges["name"] = oldName
		multipleKeyExchanges["new-name"] = newName
	} else {
		multipleKeyExchanges["name"] = d.Get("name")
	}

	if d.HasChange("key_exchange_methods") {
		if v, ok := d.GetOk("key_exchange_methods"); ok {
			multipleKeyExchanges["key-exchange-methods"] = v.(*schema.Set).List()
		} else {
			oldKey_Exchange_Methods, _ := d.GetChange("key_exchange_methods")
			multipleKeyExchanges["key-exchange-methods"] = map[string]interface{}{"remove": oldKey_Exchange_Methods.(*schema.Set).List()}
		}
	}

	if d.HasChange("additional_key_exchange_1_methods") {
		if v, ok := d.GetOk("additional_key_exchange_1_methods"); ok {
			multipleKeyExchanges["additional-key-exchange-1-methods"] = v.(*schema.Set).List()
		} else {
			oldAdditional_Key_Exchange_1_Methods, _ := d.GetChange("additional_key_exchange_1_methods")
			multipleKeyExchanges["additional-key-exchange-1-methods"] = map[string]interface{}{"remove": oldAdditional_Key_Exchange_1_Methods.(*schema.Set).List()}
		}
	}

	if d.HasChange("additional_key_exchange_2_methods") {
		if v, ok := d.GetOk("additional_key_exchange_2_methods"); ok {
			multipleKeyExchanges["additional-key-exchange-2-methods"] = v.(*schema.Set).List()
		} else {
			oldAdditional_Key_Exchange_2_Methods, _ := d.GetChange("additional_key_exchange_2_methods")
			multipleKeyExchanges["additional-key-exchange-2-methods"] = map[string]interface{}{"remove": oldAdditional_Key_Exchange_2_Methods.(*schema.Set).List()}
		}
	}

	if d.HasChange("additional_key_exchange_3_methods") {
		if v, ok := d.GetOk("additional_key_exchange_3_methods"); ok {
			multipleKeyExchanges["additional-key-exchange-3-methods"] = v.(*schema.Set).List()
		} else {
			oldAdditional_Key_Exchange_3_Methods, _ := d.GetChange("additional_key_exchange_3_methods")
			multipleKeyExchanges["additional-key-exchange-3-methods"] = map[string]interface{}{"remove": oldAdditional_Key_Exchange_3_Methods.(*schema.Set).List()}
		}
	}

	if d.HasChange("additional_key_exchange_4_methods") {
		if v, ok := d.GetOk("additional_key_exchange_4_methods"); ok {
			multipleKeyExchanges["additional-key-exchange-4-methods"] = v.(*schema.Set).List()
		} else {
			oldAdditional_Key_Exchange_4_Methods, _ := d.GetChange("additional_key_exchange_4_methods")
			multipleKeyExchanges["additional-key-exchange-4-methods"] = map[string]interface{}{"remove": oldAdditional_Key_Exchange_4_Methods.(*schema.Set).List()}
		}
	}

	if d.HasChange("additional_key_exchange_5_methods") {
		if v, ok := d.GetOk("additional_key_exchange_5_methods"); ok {
			multipleKeyExchanges["additional-key-exchange-5-methods"] = v.(*schema.Set).List()
		} else {
			oldAdditional_Key_Exchange_5_Methods, _ := d.GetChange("additional_key_exchange_5_methods")
			multipleKeyExchanges["additional-key-exchange-5-methods"] = map[string]interface{}{"remove": oldAdditional_Key_Exchange_5_Methods.(*schema.Set).List()}
		}
	}

	if d.HasChange("additional_key_exchange_6_methods") {
		if v, ok := d.GetOk("additional_key_exchange_6_methods"); ok {
			multipleKeyExchanges["additional-key-exchange-6-methods"] = v.(*schema.Set).List()
		} else {
			oldAdditional_Key_Exchange_6_Methods, _ := d.GetChange("additional_key_exchange_6_methods")
			multipleKeyExchanges["additional-key-exchange-6-methods"] = map[string]interface{}{"remove": oldAdditional_Key_Exchange_6_Methods.(*schema.Set).List()}
		}
	}

	if d.HasChange("additional_key_exchange_7_methods") {
		if v, ok := d.GetOk("additional_key_exchange_7_methods"); ok {
			multipleKeyExchanges["additional-key-exchange-7-methods"] = v.(*schema.Set).List()
		} else {
			oldAdditional_Key_Exchange_7_Methods, _ := d.GetChange("additional_key_exchange_7_methods")
			multipleKeyExchanges["additional-key-exchange-7-methods"] = map[string]interface{}{"remove": oldAdditional_Key_Exchange_7_Methods.(*schema.Set).List()}
		}
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			multipleKeyExchanges["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			multipleKeyExchanges["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		multipleKeyExchanges["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		multipleKeyExchanges["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		multipleKeyExchanges["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		multipleKeyExchanges["ignore-errors"] = v.(bool)
	}

	log.Println("Update MultipleKeyExchanges - Map = ", multipleKeyExchanges)

	updateMultipleKeyExchangesRes, err := client.ApiCall("set-multiple-key-exchanges", multipleKeyExchanges, client.GetSessionID(), true, false)
	if err != nil || !updateMultipleKeyExchangesRes.Success {
		if updateMultipleKeyExchangesRes.ErrorMsg != "" {
			return fmt.Errorf(updateMultipleKeyExchangesRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementMultipleKeyExchanges(d, m)
}

func deleteManagementMultipleKeyExchanges(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	multipleKeyExchangesPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete MultipleKeyExchanges")

	deleteMultipleKeyExchangesRes, err := client.ApiCall("delete-multiple-key-exchanges", multipleKeyExchangesPayload, client.GetSessionID(), true, false)
	if err != nil || !deleteMultipleKeyExchangesRes.Success {
		if deleteMultipleKeyExchangesRes.ErrorMsg != "" {
			return fmt.Errorf(deleteMultipleKeyExchangesRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
