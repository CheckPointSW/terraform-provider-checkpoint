package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementTrustedClient() *schema.Resource {
	return &schema.Resource{
		Create: createManagementTrustedClient,
		Read:   readManagementTrustedClient,
		Update: updateManagementTrustedClient,
		Delete: deleteManagementTrustedClient,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"ipv4_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IPv4 address.",
			},
			"ipv6_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IPv6 address.",
			},
			"domains_assignment": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Domains to be added to this profile. Use domain name only. See example below: \"add-trusted-client (with domain)\".",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ipv4_address_first": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "First IPv4 address in the range.",
			},
			"ipv6_address_first": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "First IPv6 address in the range.",
			},
			"ipv4_address_last": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Last IPv4 address in the range.",
			},
			"ipv6_address_last": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Last IPv6 address in the range.",
			},
			"mask_length4": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "IPv4 mask length.",
			},
			"mask_length6": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "IPv6 mask length.",
			},
			"multi_domain_server_trusted_client": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Let this trusted client connect to all Multi-Domain Servers in the deployment.",
				Default:     true,
			},
			"wild_card": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IP wild card (e.g. 192.0.2.*).",
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Trusted client type.",
				Default:     "ipv4 address",
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

func createManagementTrustedClient(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	trustedClient := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		trustedClient["name"] = v.(string)
	}

	if v, ok := d.GetOk("ipv4_address"); ok {
		trustedClient["ipv4-address"] = v.(string)
	}

	if v, ok := d.GetOk("ipv6_address"); ok {
		trustedClient["ipv6-address"] = v.(string)
	}

	if v, ok := d.GetOk("domains_assignment"); ok {
		trustedClient["domains-assignment"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("ipv4_address_first"); ok {
		trustedClient["ipv4-address-first"] = v.(string)
	}

	if v, ok := d.GetOk("ipv6_address_first"); ok {
		trustedClient["ipv6-address-first"] = v.(string)
	}

	if v, ok := d.GetOk("ipv4_address_last"); ok {
		trustedClient["ipv4-address-last"] = v.(string)
	}

	if v, ok := d.GetOk("ipv6_address_last"); ok {
		trustedClient["ipv6-address-last"] = v.(string)
	}

	if v, ok := d.GetOk("mask_length4"); ok {
		trustedClient["mask-length4"] = v.(int)
	}

	if v, ok := d.GetOk("mask_length6"); ok {
		trustedClient["mask-length6"] = v.(int)
	}

	if v, ok := d.GetOk("multi_domain_server_trusted_client"); ok {
		trustedClient["multi-domain-server-trusted-client"] = v.(bool)
	}

	if v, ok := d.GetOk("tags"); ok {
		trustedClient["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("wild_card"); ok {
		trustedClient["wild-card"] = v.(string)
	}

	if v, ok := d.GetOk("type"); ok {
		trustedClient["type"] = v.(string)
	}

	if v, ok := d.GetOk("color"); ok {
		trustedClient["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		trustedClient["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		trustedClient["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		trustedClient["ignore-errors"] = v.(bool)
	}

	log.Println("Create TrustedClient - Map = ", trustedClient)

	addTrustedClientRes, err := client.ApiCall("add-trusted-client", trustedClient, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !addTrustedClientRes.Success {
		if addTrustedClientRes.ErrorMsg != "" {
			return fmt.Errorf(addTrustedClientRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addTrustedClientRes.GetData()["uid"].(string))

	return readManagementTrustedClient(d, m)
}

func readManagementTrustedClient(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showTrustedClientRes, err := client.ApiCall("show-trusted-client", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showTrustedClientRes.Success {
		if objectNotFound(showTrustedClientRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showTrustedClientRes.ErrorMsg)
	}

	trustedClient := showTrustedClientRes.GetData()

	log.Println("Read TrustedClient - Show JSON = ", trustedClient)

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

func updateManagementTrustedClient(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	trustedClient := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		trustedClient["name"] = oldName
		trustedClient["new-name"] = newName
	} else {
		trustedClient["name"] = d.Get("name")
	}

	if ok := d.HasChange("ipv4_address"); ok {
		trustedClient["ipv4-address"] = d.Get("ipv4_address")
	}

	if ok := d.HasChange("ipv6_address"); ok {
		trustedClient["ipv6-address"] = d.Get("ipv6_address")
	}

	if d.HasChange("domains_assignment") {
		if v, ok := d.GetOk("domains_assignment"); ok {
			trustedClient["domains-assignment"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("domains_assignment")
			trustedClient["domains-assignment"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("ipv4_address_first"); ok {
		trustedClient["ipv4-address-first"] = d.Get("ipv4_address_first")
	}

	if ok := d.HasChange("ipv6_address_first"); ok {
		trustedClient["ipv6-address-first"] = d.Get("ipv6_address_first")
	}

	if ok := d.HasChange("ipv4_address_last"); ok {
		trustedClient["ipv4-address-last"] = d.Get("ipv4_address_last")
	}

	if ok := d.HasChange("ipv6_address_last"); ok {
		trustedClient["ipv6-address-last"] = d.Get("ipv6_address_last")
	}

	if ok := d.HasChange("multi_domain_server_trusted_client"); ok {
		trustedClient["multi-domain-server-trusted-client"] = d.Get("multi_domain_server_trusted_client").(bool)
	}

	if ok := d.HasChange("mask_length4"); ok {
		trustedClient["mask-length4"] = d.Get("mask_length4")
	}

	if ok := d.HasChange("mask_length6"); ok {
		trustedClient["mask-length6"] = d.Get("mask_length6")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			trustedClient["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			trustedClient["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("wild_card"); ok {
		trustedClient["wild-card"] = d.Get("wild_card")
	}

	if ok := d.HasChange("type"); ok {
		trustedClient["type"] = d.Get("type")
	}

	if ok := d.HasChange("color"); ok {
		trustedClient["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		trustedClient["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		trustedClient["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		trustedClient["ignore-errors"] = v.(bool)
	}

	log.Println("Update TrustedClient - Map = ", trustedClient)

	updateTrustedClientRes, err := client.ApiCall("set-trusted-client", trustedClient, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !updateTrustedClientRes.Success {
		if updateTrustedClientRes.ErrorMsg != "" {
			return fmt.Errorf(updateTrustedClientRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementTrustedClient(d, m)
}

func deleteManagementTrustedClient(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	trustedClientPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete TrustedClient")

	deleteTrustedClientRes, err := client.ApiCall("delete-trusted-client", trustedClientPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteTrustedClientRes.Success {
		if deleteTrustedClientRes.ErrorMsg != "" {
			return fmt.Errorf(deleteTrustedClientRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
