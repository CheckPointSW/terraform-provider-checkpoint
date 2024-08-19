package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementMultipleKeyExchanges() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementMultipleKeyExchangesRead,
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
			"key_exchange_methods": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Key-Exchange methods to use. Can contain only Diffie-Hellman groups.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"additional_key_exchange_1_methods": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Additional Key-Exchange 1 methods to use.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"additional_key_exchange_2_methods": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Additional Key-Exchange 2 methods to use.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"additional_key_exchange_3_methods": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Additional Key-Exchange 3 methods to use.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"additional_key_exchange_4_methods": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Additional Key-Exchange 4 methods to use.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"additional_key_exchange_5_methods": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Additional Key-Exchange 5 methods to use.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"additional_key_exchange_6_methods": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Additional Key-Exchange 6 methods to use.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"additional_key_exchange_7_methods": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Additional Key-Exchange 7 methods to use.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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

func dataSourceManagementMultipleKeyExchangesRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
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

	if v := multipleKeyExchanges["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

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

	return nil

}
