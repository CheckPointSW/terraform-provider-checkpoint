package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strconv"
)

func resourceManagementDomain() *schema.Resource {
	return &schema.Resource{
		Create: createManagementDomain,
		Read:   readManagementDomain,
		Update: updateManagementDomain,
		Delete: deleteManagementDomain,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"servers": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "Domain servers.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Object name. Must be unique in the domain.",
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
						"multi_domain_server": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Multi Domain server name or UID.",
						},
						"active": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Activate domain server. Only one domain server is allowed to be active.",
						},
						"skip_start_domain_server": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Set this value to be true to prevent starting the new created domain.",
						},
						"type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Domain server type.",
						},
					},
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

func createManagementDomain(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	domain := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		domain["name"] = v.(string)
	}

	if v, ok := d.GetOk("servers"); ok {

		serversList := v.([]interface{})

		if len(serversList) > 0 {
			var serversPayload []map[string]interface{}

			for i := range serversList {

				Payload := make(map[string]interface{})

				if v, ok := d.GetOk("servers." + strconv.Itoa(i) + ".name"); ok {
					Payload["name"] = v.(string)
				}
				if v, ok := d.GetOk("servers." + strconv.Itoa(i) + ".ipv4_address"); ok {
					Payload["ipv4-address"] = v.(string)
				}
				if v, ok := d.GetOk("servers." + strconv.Itoa(i) + ".ipv6_address"); ok {
					Payload["ipv6-address"] = v.(string)
				}
				if v, ok := d.GetOk("servers." + strconv.Itoa(i) + ".multi_domain_server"); ok {
					Payload["multi-domain-server"] = v.(string)
				}
				if v, ok := d.GetOk("servers." + strconv.Itoa(i) + ".active"); ok {
					Payload["active"] = v.(bool)
				}
				if v, ok := d.GetOk("servers." + strconv.Itoa(i) + ".skip_start_domain_server"); ok {
					Payload["skip-start-domain-server"] = v.(bool)
				}
				if v, ok := d.GetOk("servers." + strconv.Itoa(i) + ".type"); ok {
					Payload["type"] = v.(string)
				}
				serversPayload = append(serversPayload, Payload)
			}
			domain["servers"] = serversPayload
		}
	}

	if v, ok := d.GetOk("color"); ok {
		domain["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		domain["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		domain["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		domain["ignore-errors"] = v.(bool)
	}

	log.Println("Create Domain - Map = ", domain)

	addDomainRes, err := client.ApiCall("add-domain", domain, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !addDomainRes.Success {
		if addDomainRes.ErrorMsg != "" {
			return fmt.Errorf(addDomainRes.ErrorMsg)
		}
		msg := createTaskFailMessage("add-domain", addDomainRes.GetData())
		return fmt.Errorf(msg)
	}

	// add-simple-cluster returns task-id. Call show-simple-cluster for object uid.
	showDomainRes, err := client.ApiCall("show-domain", map[string]interface{}{"name": d.Get("name")}, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showDomainRes.Success {
		return fmt.Errorf(showDomainRes.ErrorMsg)
	}

	d.SetId(showDomainRes.GetData()["uid"].(string))

	return readManagementDomain(d, m)
}

func readManagementDomain(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showDomainRes, err := client.ApiCall("show-domain", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showDomainRes.Success {
		if objectNotFound(showDomainRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showDomainRes.ErrorMsg)
	}

	domain := showDomainRes.GetData()

	log.Println("Read Domain - Show JSON = ", domain)

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

func updateManagementDomain(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	domain := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		domain["name"] = oldName
		domain["new-name"] = newName
	} else {
		domain["name"] = d.Get("name")
	}

	if d.HasChange("servers") {
		if v, ok := d.GetOk("servers"); ok {
			serversList := v.([]interface{})
			var serversPayload []map[string]interface{}

			for i := range serversList {
				serverPayload := make(map[string]interface{})
				if v, ok := d.GetOk("servers." + strconv.Itoa(i) + ".name"); ok {
					serverPayload["name"] = v
				}
				if v, ok := d.GetOk("servers." + strconv.Itoa(i) + ".ipv4_address"); ok {
					serverPayload["ipv4-address"] = v
				}
				if v, ok := d.GetOk("servers." + strconv.Itoa(i) + ".ipv6_address"); ok {
					serverPayload["ipv6-address"] = v
				}
				if v, ok := d.GetOk("servers." + strconv.Itoa(i) + ".multi_domain_server"); ok {
					serverPayload["multi-domain-server"] = v
				}
				if v, ok := d.GetOk("servers." + strconv.Itoa(i) + ".active"); ok {
					serverPayload["active"] = v.(bool)
				}
				if v, ok := d.GetOk("servers." + strconv.Itoa(i) + ".skip_start_domain_server"); ok {
					serverPayload["skip-start-domain-server"] = v.(bool)
				}
				if v, ok := d.GetOk("servers." + strconv.Itoa(i) + ".type"); ok {
					serverPayload["type"] = v
				}
				serversPayload = append(serversPayload, serverPayload)
			}
			domain["servers"] = serversPayload
		}
	}

	if ok := d.HasChange("color"); ok {
		domain["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		domain["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		domain["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		domain["ignore-errors"] = v.(bool)
	}

	log.Println("Update Domain - Map = ", domain)

	updateDomainRes, err := client.ApiCall("set-domain", domain, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !updateDomainRes.Success {
		if updateDomainRes.ErrorMsg != "" {
			return fmt.Errorf(updateDomainRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementDomain(d, m)
}

func deleteManagementDomain(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	domainPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		domainPayload["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		domainPayload["ignore-errors"] = v.(bool)
	}

	log.Println("Delete Domain")

	deleteDomainRes, err := client.ApiCall("delete-domain", domainPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteDomainRes.Success {
		if deleteDomainRes.ErrorMsg != "" {
			return fmt.Errorf(deleteDomainRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
