package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementIfMapServer() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceManagementIfMapServerRead,
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
			"port": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "IF-MAP server port number.",
			},
			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IF-MAP version.",
			},
			"host": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Host that is IF-MAP server. Identified by name or UID.",
			},
			"path": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "N/A",
			},
			"monitored_ips": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "IP ranges to be monitored by the IF-MAP client.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"first_ip": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "First IPv4 address in the range to be monitored.",
						},
						"last_ip": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Last IPv4 address in the range to be monitored.",
						},
					},
				},
			},
			"query_whole_ranges": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicate whether to query whole ranges instead of single IP.",
			},
			"authentication": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Authentication configuration for the IF-MAP server.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication_method": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Authentication method for the IF-MAP server.",
						},
						"username": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Username for the IF-MAP server authentication. <font color=\"red\">Required only when</font> 'authentication-method' is set to 'basic'.",
						},
						"password": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Username for the IF-MAP server authentication. <font color=\"red\">Required only when</font> 'authentication-method' is set to 'basic'.",
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
			"tags": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceManagementIfMapServerRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}
	showIfMapServerRes, err := client.ApiCallSimple("show-if-map-server", payload)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showIfMapServerRes.Success {
		return fmt.Errorf(showIfMapServerRes.ErrorMsg)
	}

	ifMapServer := showIfMapServerRes.GetData()

	log.Println("Read IfMapServer - Show JSON = ", ifMapServer)

	if v := ifMapServer["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := ifMapServer["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := ifMapServer["port"]; v != nil {
		_ = d.Set("port", v)
	}

	if v := ifMapServer["version"]; v != nil {
		_ = d.Set("version", v)
	}

	if v := ifMapServer["host"]; v != nil {
		_ = d.Set("host", v.(map[string]interface{})["name"].(string))
	}

	if v := ifMapServer["path"]; v != nil {
		_ = d.Set("path", v)
	}

	if ifMapServer["monitored-ips"] != nil {

		monitoredIpsList := ifMapServer["monitored-ips"].([]interface{})

		if len(monitoredIpsList) > 0 {

			var monitoredIpsListToReturn []map[string]interface{}

			for i := range monitoredIpsList {

				monitoredIpMap := monitoredIpsList[i].(map[string]interface{})

				monitoredIpMapToAdd := make(map[string]interface{})

				if v, _ := monitoredIpMap["first-ip"]; v != nil {
					monitoredIpMapToAdd["first_ip"] = v
				}
				if v, _ := monitoredIpMap["last-ip"]; v != nil {
					monitoredIpMapToAdd["last_ip"] = v
				}

				monitoredIpsListToReturn = append(monitoredIpsListToReturn, monitoredIpMapToAdd)
			}

			_ = d.Set("monitored_ips", monitoredIpsListToReturn)
		} else {
			_ = d.Set("monitored_ips", monitoredIpsList)
		}
	} else {
		_ = d.Set("monitored_ips", nil)
	}

	if v := ifMapServer["query-whole-ranges"]; v != nil {
		_ = d.Set("query_whole_ranges", v)
	}

	if ifMapServer["authentication"] != nil {

		authenticationMap := ifMapServer["authentication"].(map[string]interface{})

		authenticationMapToReturn := make(map[string]interface{})

		if v, _ := authenticationMap["authentication-method"]; v != "" && v != nil {
			authenticationMapToReturn["authentication_method"] = v
		}

		if v, _ := authenticationMap["username"]; v != nil {
			authenticationMapToReturn["username"] = v
		}

		_ = d.Set("authentication", []interface{}{authenticationMapToReturn})

	} else {
		_ = d.Set("authentication", nil)
	}

	if ifMapServer["tags"] != nil {
		tagsJson, ok := ifMapServer["tags"].([]interface{})
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

	if v := ifMapServer["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := ifMapServer["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil
}
