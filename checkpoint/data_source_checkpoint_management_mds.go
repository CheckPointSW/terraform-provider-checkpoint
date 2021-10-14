package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementMds() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementMdsRead,
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
			"hardware": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Hardware name. For example: Open server, Smart-1, Other.",
			},
			"os": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Operating system name. For example: Gaia, Linux, SecurePlatform.",
			},
			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "System version.",
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
				Computed:    true,
				Description: "First IP address in the range.",
			},
			"ip_pool_last": {
				Type:        schema.TypeString,
				Computed:    true,
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
			"server_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of the management server.",
			},
		},
	}
}

func dataSourceManagementMdsRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showMdsRes, err := client.ApiCall("show-mds", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showMdsRes.Success {
		return fmt.Errorf(showMdsRes.ErrorMsg)
	}

	mds := showMdsRes.GetData()

	log.Println("Read Mds - Show JSON = ", mds)

	if v := mds["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

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
