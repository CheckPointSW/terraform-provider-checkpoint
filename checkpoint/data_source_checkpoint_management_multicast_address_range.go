package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementMulticastAddressRange() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementMulticastAddressRangeRead,
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
			"ipv4_address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IPv4 address.",
			},
			"ipv6_address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IPv6 address.",
			},
			"ipv4_address_first": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "First IPv4 address in the range.",
			},
			"ipv6_address_first": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "First IPv6 address in the range.",
			},
			"ipv4_address_last": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Last IPv4 address in the range.",
			},
			"ipv6_address_last": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Last IPv6 address in the range.",
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
			"groups": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of group identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceManagementMulticastAddressRangeRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showMulticastAddressRangeRes, err := client.ApiCall("show-multicast-address-range", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showMulticastAddressRangeRes.Success {
		return fmt.Errorf(showMulticastAddressRangeRes.ErrorMsg)
	}

	multicastAddressRange := showMulticastAddressRangeRes.GetData()

	log.Println("Read MulticastAddressRange - Show JSON = ", multicastAddressRange)

	if v := multicastAddressRange["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := multicastAddressRange["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := multicastAddressRange["ipv4-address"]; v != nil {
		_ = d.Set("ipv4_address", v)
	}

	if v := multicastAddressRange["ipv6-address"]; v != nil {
		_ = d.Set("ipv6_address", v)
	}

	if v := multicastAddressRange["ipv4-address-first"]; v != nil {
		_ = d.Set("ipv4_address_first", v)
	}

	if v := multicastAddressRange["ipv6-address-first"]; v != nil {
		_ = d.Set("ipv6_address_first", v)
	}

	if v := multicastAddressRange["ipv4-address-last"]; v != nil {
		_ = d.Set("ipv4_address_last", v)
	}

	if v := multicastAddressRange["ipv6-address-last"]; v != nil {
		_ = d.Set("ipv6_address_last", v)
	}

	if multicastAddressRange["tags"] != nil {
		tagsJson, ok := multicastAddressRange["tags"].([]interface{})
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

	if v := multicastAddressRange["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := multicastAddressRange["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if multicastAddressRange["groups"] != nil {
		groupsJson, ok := multicastAddressRange["groups"].([]interface{})
		if ok {
			groupsIds := make([]string, 0)
			if len(groupsJson) > 0 {
				for _, groups := range groupsJson {
					groups := groups.(map[string]interface{})
					groupsIds = append(groupsIds, groups["name"].(string))
				}
			}
			_ = d.Set("groups", groupsIds)
		}
	} else {
		_ = d.Set("groups", nil)
	}

	if v := multicastAddressRange["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := multicastAddressRange["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}
