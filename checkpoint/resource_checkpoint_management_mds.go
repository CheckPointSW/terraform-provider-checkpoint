package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementMds() *schema.Resource {
	return &schema.Resource{
		Create: createManagementMds,
		Read:   readManagementMds,
		Update: updateManagementMds,
		Delete: deleteManagementMds,
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
			"hardware": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Hardware name. For example: Open server, Smart-1, Other.",
			},
			"os": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Operating system name. For example: Gaia, Linux, SecurePlatform.",
			},
			"version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "System version.",
				Default:     "R81",
			},
			"one_time_password": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Secure internal connection one time password.",
			},
			"sic_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the Secure Internal Connection Trust.",
			},
			"sic_state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "State the Secure Internal Connection Trust.",
			},
			"ip_pool_first": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "First IP address in the range.",
			},
			"ip_pool_last": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Last IP address in the range.",
			},
			"domains": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of Domain objects identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"global_domains": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of Global domain objects identified by the name or UID.",
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
			"server_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Type of the management server.",
				Default:     "multi-domain server",
			},
		},
	}
}

func createManagementMds(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	mds := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		mds["name"] = v.(string)
	}

	if v, ok := d.GetOk("ipv4_address"); ok {
		mds["ipv4-address"] = v.(string)
	}

	if v, ok := d.GetOk("ipv6_address"); ok {
		mds["ipv6-address"] = v.(string)
	}

	if v, ok := d.GetOk("hardware"); ok {
		mds["hardware"] = v.(string)
	}

	if v, ok := d.GetOk("os"); ok {
		mds["os"] = v.(string)
	}

	if v, ok := d.GetOk("version"); ok {
		mds["version"] = v.(string)
	}

	if v, ok := d.GetOk("one_time_password"); ok {
		mds["one-time-password"] = v.(string)
	}

	if v, ok := d.GetOk("ip_pool_first"); ok {
		mds["ip-pool-first"] = v.(string)
	}

	if v, ok := d.GetOk("ip_pool_last"); ok {
		mds["ip-pool-last"] = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		mds["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		mds["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		mds["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		mds["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		mds["ignore-errors"] = v.(bool)
	}

	if v, ok := d.GetOk("server_type"); ok {
		mds["server-type"] = v.(string)
	}

	log.Println("Create Mds - Map = ", mds)

	addMdsRes, err := client.ApiCall("add-mds", mds, client.GetSessionID(), true, false)
	if err != nil || !addMdsRes.Success {
		if addMdsRes.ErrorMsg != "" {
			return fmt.Errorf(addMdsRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addMdsRes.GetData()["uid"].(string))

	return readManagementMds(d, m)
}

func readManagementMds(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showMdsRes, err := client.ApiCall("show-mds", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showMdsRes.Success {
		if objectNotFound(showMdsRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showMdsRes.ErrorMsg)
	}

	mds := showMdsRes.GetData()

	log.Println("Read Mds - Show JSON = ", mds)

	if v := mds["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := mds["ipv4-address"]; v != nil {
		_ = d.Set("ipv4_address", v)
	}

	if v := mds["ipv6-address"]; v != nil {
		_ = d.Set("ipv6_address", v)
	}

	if v := mds["hardware"]; v != nil {
		_ = d.Set("hardware", v.(map[string]interface{})["name"].(string))
	}

	if v := mds["os"]; v != nil {
		_ = d.Set("os", v.(map[string]interface{})["name"].(string))
	}

	if v := mds["version"]; v != nil {
		_ = d.Set("version", v.(map[string]interface{})["name"].(string))
	}

	if v := mds["sic_name"]; v != nil {
		_ = d.Set("sic_name", v)
	}

	if v := mds["sic_state"]; v != nil {
		_ = d.Set("sic_state", v)
	}

	if v := mds["ip-pool-first"]; v != nil {
		_ = d.Set("ip_pool_first", v)
	}

	if v := mds["ip-pool-last"]; v != nil {
		_ = d.Set("ip_pool_last", v)
	}

	if mds["domains"] != nil {
		domainsJson, ok := mds["domains"].([]interface{})
		if ok {
			domainsIds := make([]string, 0)
			if len(domainsJson) > 0 {
				for _, domain := range domainsJson {
					domainsIds = append(domainsIds, domain.(map[string]interface{})["name"].(string))
				}
			}
			_ = d.Set("domains", domainsIds)
		}
	} else {
		_ = d.Set("domains", nil)
	}

	if mds["global-domains"] != nil {
		globalDomainsJson, ok := mds["global-domains"].([]interface{})
		if ok {
			globalDomainsIds := make([]string, 0)
			if len(globalDomainsJson) > 0 {
				for _, globalDomain := range globalDomainsJson {
					globalDomainsIds = append(globalDomainsIds, globalDomain.(map[string]interface{})["name"].(string))
				}
			}
			_ = d.Set("global_domains", globalDomainsIds)
		}
	} else {
		_ = d.Set("global_domains", nil)
	}

	if mds["tags"] != nil {
		tagsJson, ok := mds["tags"].([]interface{})
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

	if v := mds["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := mds["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := mds["server-type"]; v != nil {
		_ = d.Set("server_type", v)
	}

	return nil

}

func updateManagementMds(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	mds := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		mds["name"] = oldName
		mds["new-name"] = newName
	} else {
		mds["name"] = d.Get("name")
	}

	if ok := d.HasChange("ipv4_address"); ok {
		mds["ipv4-address"] = d.Get("ipv4_address")
	}

	if ok := d.HasChange("ipv6_address"); ok {
		mds["ipv6-address"] = d.Get("ipv6_address")
	}

	if ok := d.HasChange("hardware"); ok {
		mds["hardware"] = d.Get("hardware")
	}

	if ok := d.HasChange("os"); ok {
		mds["os"] = d.Get("os")
	}

	if ok := d.HasChange("version"); ok {
		mds["version"] = d.Get("version")
	}

	if ok := d.HasChange("one_time_password"); ok {
		mds["one-time-password"] = d.Get("one_time_password")
	}

	if ok := d.HasChange("ip_pool_first"); ok {
		mds["ip-pool-first"] = d.Get("ip_pool_first")
	}

	if ok := d.HasChange("ip_pool_last"); ok {
		mds["ip-pool-last"] = d.Get("ipv4_pool_last")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			mds["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			mds["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		mds["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		mds["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		mds["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		mds["ignore-errors"] = v.(bool)
	}

	if ok := d.HasChange("server_type"); ok {
		mds["server-type"] = d.Get("server_type")
	}

	log.Println("Update Mds - Map = ", mds)

	updateMdsRes, err := client.ApiCall("set-mds", mds, client.GetSessionID(), true, false)
	if err != nil || !updateMdsRes.Success {
		if updateMdsRes.ErrorMsg != "" {
			return fmt.Errorf(updateMdsRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementMds(d, m)
}

func deleteManagementMds(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	mdsPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete Mds")

	deleteMdsRes, err := client.ApiCall("delete-mds", mdsPayload, client.GetSessionID(), true, false)
	if err != nil || !deleteMdsRes.Success {
		if deleteMdsRes.ErrorMsg != "" {
			return fmt.Errorf(deleteMdsRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
