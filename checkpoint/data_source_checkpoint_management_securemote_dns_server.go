package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementSecuremoteDnsServer() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementSecuremoteDnsServerRead,
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
			"host": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "DNS server for remote clients in the Remote access community.",
			},
			"domains": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "SecuremoteDnsServer domains.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain_suffix": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "DNS Domain suffix.",
						},
						"maximum_prefix_label_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Maximum number of matching labels preceding the suffix.",
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

func dataSourceManagementSecuremoteDnsServerRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showSecuremoteDnsServerRes, err := client.ApiCallSimple("show-securemote-dns-server", payload)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showSecuremoteDnsServerRes.Success {
		return fmt.Errorf(showSecuremoteDnsServerRes.ErrorMsg)
	}

	securemoteDnsServer := showSecuremoteDnsServerRes.GetData()

	log.Println("Read SecuremoteDnsServer - Show JSON = ", securemoteDnsServer)

	if v := securemoteDnsServer["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := securemoteDnsServer["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := securemoteDnsServer["host"]; v != nil {
		_ = d.Set("host", v.(map[string]interface{})["name"].(string))
	}

	if v := securemoteDnsServer["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := securemoteDnsServer["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if securemoteDnsServer["domains"] != nil {

		domainsList := securemoteDnsServer["domains"].([]interface{})

		if len(domainsList) > 0 {

			var domainsListToReturn []map[string]interface{}

			for i := range domainsList {

				domainMap := domainsList[i].(map[string]interface{})

				domainMapToAdd := make(map[string]interface{})

				if v, _ := domainMap["domain-suffix"]; v != nil {
					domainMapToAdd["domain_suffix"] = v
				}
				if v, _ := domainMap["maximum-prefix-label-count"]; v != nil {
					domainMapToAdd["maximum_prefix_label_count"] = v
				}

				domainsListToReturn = append(domainsListToReturn, domainMapToAdd)
			}

			_ = d.Set("domains", domainsListToReturn)
		} else {
			_ = d.Set("domains", domainsList)
		}
	} else {
		_ = d.Set("domains", nil)
	}

	if securemoteDnsServer["tags"] != nil {
		tagsJson := securemoteDnsServer["tags"].([]interface{})
		var tagsIds = make([]string, 0)
		if len(tagsJson) > 0 {
			// Create slice of tag names
			for _, tag := range tagsJson {
				tag := tag.(map[string]interface{})
				tagsIds = append(tagsIds, tag["name"].(string))
			}
		}
		_ = d.Set("tags", tagsIds)
	} else {
		_ = d.Set("tags", nil)
	}

	return nil
}
