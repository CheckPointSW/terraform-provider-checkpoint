package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strconv"
)

func dataSourceManagementDomain() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementDomainRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object unique identifier.",
			},
			"servers": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Domain servers.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object name. Must be unique in the domain.",
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
						"multi_domain_server": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Multi Domain server name or UID.",
						},
						"active": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Activate domain server. Only one domain server is allowed to be active.",
						},
						"skip_start_domain_server": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Set this value to be true to prevent starting the new created domain.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Domain server type.",
						},
					},
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

func dataSourceManagementDomainRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showDomainRes, err := client.ApiCall("show-domain", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showDomainRes.Success {
		return fmt.Errorf(showDomainRes.ErrorMsg)
	}

	domain := showDomainRes.GetData()

	log.Println("Read Domain - Show JSON = ", domain)

	if v := domain["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := domain["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if domain["servers"] != nil {

		serversList, ok := domain["servers"].([]interface{})

		if ok {

			if len(serversList) > 0 {

				var serversListToReturn []map[string]interface{}

				for i := range serversList {

					serversMap := serversList[i].(map[string]interface{})

					serversMapToAdd := make(map[string]interface{})

					if v, _ := serversMap["name"]; v != nil {
						serversMapToAdd["name"] = v
					}
					if v, _ := serversMap["ipv4-address"]; v != nil {
						serversMapToAdd["ipv4_address"] = v
					}
					if v, _ := serversMap["ipv6-address"]; v != nil {
						serversMapToAdd["ipv4_address"] = v
					}
					if v, _ := serversMap["multi-domain-server"]; v != nil {
						serversMapToAdd["multi_domain_server"] = v
					}
					if v, _ := serversMap["active"]; v != nil {
						serversMapToAdd["active"] = strconv.FormatBool(v.(bool))
					} else {
						serversMapToAdd["active"] = true
					}
					if v, _ := serversMap["skip-start-domain-server"]; v != nil {
						serversMapToAdd["skip_start_domain_server"] = strconv.FormatBool(v.(bool))
					} else {
						serversMapToAdd["skip_start_domain_server"] = false
					}
					if v, _ := serversMap["type"]; v != nil {
						serversMapToAdd["type"] = v
					}

					serversListToReturn = append(serversListToReturn, serversMapToAdd)
				}
				_ = d.Set("servers", serversListToReturn)
			}
		}
	}

	if v := domain["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := domain["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil
}
