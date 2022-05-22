package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementTrustedClient() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementTrustedClientRead,
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
			"domains_assignment": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Domains to be added to this profile. Use domain name only. See example below: \"add-trusted-client (with domain)\".",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
			"mask_length4": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "IPv4 mask length.",
			},
			"mask_length6": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "IPv6 mask length.",
			},
			"multi_domain_server_trusted_client": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Let this trusted client connect to all Multi-Domain Servers in the deployment.",
			},
			"wild_card": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IP wild card (e.g. 192.0.2.*).",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Trusted client type.",
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

func dataSourceManagementTrustedClientRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showTrustedClientRes, err := client.ApiCall("show-trusted-client", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showTrustedClientRes.Success {
		return fmt.Errorf(showTrustedClientRes.ErrorMsg)
	}

	trustedClient := showTrustedClientRes.GetData()

	log.Println("Read TrustedClient - Show JSON = ", trustedClient)

	if v := trustedClient["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := trustedClient["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := trustedClient["ipv4-address"]; v != nil {
		_ = d.Set("ipv4_address", v)
	}

	if v := trustedClient["ipv6-address"]; v != nil {
		_ = d.Set("ipv6_address", v)
	}

	if trustedClient["domains-assignment"] != nil {
		tagsJson, ok := trustedClient["domains-assignment"].([]interface{})
		if ok {
			tagsIds := make([]string, 0)
			if len(tagsJson) > 0 {
				for _, tags := range tagsJson {
					tags := tags.(map[string]interface{})
					tagsIds = append(tagsIds, tags["name"].(string))
				}
			}
			_ = d.Set("domains_assignment", tagsIds)
		}
	} else {
		_ = d.Set("domains_assignment", nil)
	}

	if v := trustedClient["ipv4-address-first"]; v != nil {
		_ = d.Set("ipv4_address_first", v)
	}

	if v := trustedClient["ipv6-address-first"]; v != nil {
		_ = d.Set("ipv6_address_first", v)
	}

	if v := trustedClient["ipv4-address-last"]; v != nil {
		_ = d.Set("ipv4_address_last", v)
	}

	if v := trustedClient["ipv6-address-last"]; v != nil {
		_ = d.Set("ipv6_address_last", v)
	}

	if v := trustedClient["mask-length4"]; v != nil {
		_ = d.Set("mask_length4", v)
	}

	if v := trustedClient["mask-length6"]; v != nil {
		_ = d.Set("mask_length6", v)
	}

	if v := trustedClient["multi-domain-server-trusted-client"]; v != nil {
		_ = d.Set("multi_domain_server_trusted_client", v)
	}

	if v := trustedClient["wild-card"]; v != nil {
		_ = d.Set("wild_card", v)
	}

	if v := trustedClient["type"]; v != nil {
		_ = d.Set("type", v)
	}

	if trustedClient["tags"] != nil {
		tagsJson, ok := trustedClient["tags"].([]interface{})
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

	if v := trustedClient["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := trustedClient["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := trustedClient["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := trustedClient["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}
