package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func dataSourceManagementNatRule() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementNatRuleRead,
		Schema: map[string]*schema.Schema{
			"package": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the package.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Rule name.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Rule UID.",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable/Disable the rule.",
			},
			"install_on": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Which Gateways identified by the name or UID to install the policy on.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"method": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Nat method.",
			},
			"original_destination": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Original destination.",
			},
			"original_service": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Original service.",
			},
			"original_source": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Original source.",
			},
			"translated_destination": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Translated destination.",
			},
			"translated_service": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Translated service.",
			},
			"translated_source": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Translated source.",
			},
			"auto_generated": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Auto generated.",
			},
			"comments": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Comments string.",
			},
			"hits": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Hits count object.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"first_date": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "N/A",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"iso_8601": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Date and time represented in international ISO 8601 format.",
									},
									"posix": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.",
									},
								},
							},
						},
						"last_date": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "N/A",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"iso_8601": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Date and time represented in international ISO 8601 format.",
									},
									"posix": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.",
									},
								},
							},
						},
						"level": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "N/A",
						},
						"percentage": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "N/A",
						},
						"value": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "N/A",
						},
					},
				},
			},
		},
	}
}

func dataSourceManagementNatRuleRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := map[string]interface{}{
		"package": d.Get("package"),
	}

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showNatRuleRes, err := client.ApiCall("show-nat-rule", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	if !showNatRuleRes.Success {
		return fmt.Errorf("%s", showNatRuleRes.ErrorMsg)
	}

	natRule := showNatRuleRes.GetData()

	log.Println("Read NAT Rule - Show JSON = ", natRule)

	if v := natRule["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := natRule["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := natRule["auto-generated"]; v != nil {
		_ = d.Set("auto_generated", v)
	}

	if v := natRule["enabled"]; v != nil {
		_ = d.Set("enabled", v)
	}

	if v := natRule["method"]; v != nil {
		_ = d.Set("method", v)
	}

	if natRule["install-on"] != nil {
		installOnJson := natRule["install-on"].([]interface{})
		installOnJsonIds := make([]string, 0)
		if len(installOnJson) > 0 {
			for _, installOn := range installOnJson {
				installOn := installOn.(map[string]interface{})
				installOnJsonIds = append(installOnJsonIds, installOn["name"].(string))
			}
		}
		_, installOnInConf := d.GetOk("install_on")
		if installOnJsonIds[0] == "Policy Targets" && !installOnInConf {
			_ = d.Set("install_on", []interface{}{})
		} else {
			_ = d.Set("install_on", installOnJsonIds)
		}
	}

	if v := natRule["original-destination"]; v != nil {
		_ = d.Set("original_destination", v.(map[string]interface{})["name"])
	}

	if v := natRule["original-service"]; v != nil {
		_ = d.Set("original_service", v.(map[string]interface{})["name"])
	}

	if v := natRule["original-source"]; v != nil {
		_ = d.Set("original_source", v.(map[string]interface{})["name"])
	}

	if v := natRule["translated-destination"]; v != nil {
		_ = d.Set("translated_destination", v.(map[string]interface{})["name"])
	}

	if v := natRule["translated-service"]; v != nil {
		_ = d.Set("translated_service", v.(map[string]interface{})["name"])
	}

	if v := natRule["translated-source"]; v != nil {
		_ = d.Set("translated_source", v.(map[string]interface{})["name"])
	}

	if v := natRule["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := natRule["hits"]; v != nil {
		hitsShow := v.(map[string]interface{})
		hitsState := make(map[string]interface{})
		if v := hitsShow["first-date"]; v != nil {
			firstDateShow := v.(map[string]interface{})
			firstDateState := make(map[string]interface{})
			if v := firstDateShow["iso-8601"]; v != nil {
				firstDateState["iso_8601"] = v
			}
			if v := firstDateShow["posix"]; v != nil {
				firstDateState["posix"] = v
			}
			hitsState["first_date"] = []interface{}{firstDateState}
		}
		if v := hitsShow["last-date"]; v != nil {
			lastDateShow := v.(map[string]interface{})
			lastDateState := make(map[string]interface{})
			if v := lastDateShow["iso-8601"]; v != nil {
				lastDateState["iso_8601"] = v
			}
			if v := lastDateShow["posix"]; v != nil {
				lastDateState["posix"] = v
			}
			hitsState["last_date"] = []interface{}{lastDateState}
		}
		if v := hitsShow["level"]; v != nil {
			hitsState["level"] = v
		}
		if v := hitsShow["percentage"]; v != nil {
			hitsState["percentage"] = v
		}
		if v := hitsShow["value"]; v != nil {
			hitsState["value"] = v
		}
		_ = d.Set("hits", []interface{}{hitsState})
	}

	return nil
}
