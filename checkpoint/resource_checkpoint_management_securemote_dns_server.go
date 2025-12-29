package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strconv"
)

func resourceManagementSecuremoteDnsServer() *schema.Resource {
	return &schema.Resource{
		Create: createManagementSecuremoteDnsServer,
		Read:   readManagementSecuremoteDnsServer,
		Update: updateManagementSecuremoteDnsServer,
		Delete: deleteManagementSecuremoteDnsServer,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name. Should be unique in the domain.",
			},
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "DNS server for remote clients in the Remote access community.",
			},
			"domains": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The DNS domains that the remote clients can access.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain_suffix": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "DNS Domain suffix.",
						},
						"maximum_prefix_label_count": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Maximum number of matching labels preceding the suffix.",
							Default:     1,
						},
					},
				},
			},
			"ignore_warnings": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Apply changes ignoring warnings.",
			},
			"ignore_errors": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.",
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
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func createManagementSecuremoteDnsServer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	securemoteDnsServer := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		securemoteDnsServer["name"] = v.(string)
	}
	if v, ok := d.GetOk("host"); ok {
		securemoteDnsServer["host"] = v.(string)
	}

	//list of objects
	if v, ok := d.GetOk("domains"); ok {

		domainsList := v.([]interface{})
		if len(domainsList) > 0 {

			var domainsPayload []map[string]interface{}

			for i := range domainsList {

				payload := make(map[string]interface{})

				if v, ok := d.GetOk("domains." + strconv.Itoa(i) + ".domain_suffix"); ok {
					payload["domain-suffix"] = v.(string)
				}
				if v, ok := d.GetOk("domains." + strconv.Itoa(i) + ".maximum_prefix_label_count"); ok {
					payload["maximum-prefix-label-count"] = v.(int)
				}
				domainsPayload = append(domainsPayload, payload)
			}

			securemoteDnsServer["domains"] = domainsPayload
		}
	}

	if val, ok := d.GetOk("comments"); ok {
		securemoteDnsServer["comments"] = val.(string)
	}
	if val, ok := d.GetOk("tags"); ok {
		securemoteDnsServer["tags"] = val.(*schema.Set).List()
	}

	if val, ok := d.GetOk("color"); ok {
		securemoteDnsServer["color"] = val.(string)
	}
	if val, ok := d.GetOkExists("ignore_errors"); ok {
		securemoteDnsServer["ignore-errors"] = val.(bool)
	}
	if val, ok := d.GetOkExists("ignore_warnings"); ok {
		securemoteDnsServer["ignore-warnings"] = val.(bool)
	}

	log.Println("Create SecuremoteDnsServer - Map = ", securemoteDnsServer)

	addSecuremoteDnsServerRes, err := client.ApiCallSimple("add-securemote-dns-server", securemoteDnsServer)
	if err != nil || !addSecuremoteDnsServerRes.Success {
		if addSecuremoteDnsServerRes.ErrorMsg != "" {
			return fmt.Errorf(addSecuremoteDnsServerRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addSecuremoteDnsServerRes.GetData()["uid"].(string))

	return readManagementSecuremoteDnsServer(d, m)
}

func readManagementSecuremoteDnsServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showSecuremoteDnsServerRes, err := client.ApiCallSimple("show-securemote-dns-server", payload)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showSecuremoteDnsServerRes.Success {
		// Handle delete resource from other clients
		if objectNotFound(showSecuremoteDnsServerRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showSecuremoteDnsServerRes.ErrorMsg)
	}

	securemoteDnsServer := showSecuremoteDnsServerRes.GetData()

	log.Println("Read SecuremoteDnsServer - Show JSON = ", securemoteDnsServer)

	if v := securemoteDnsServer["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := securemoteDnsServer["host"]; v != nil {
		_ = d.Set("host", v)
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
					domainMapToAdd["domain-suffix"] = v
				}
				if v, _ := domainMap["maximum-prefix-label-count"]; v != nil {
					domainMapToAdd["maximum-prefix-label-count"] = v
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

func updateManagementSecuremoteDnsServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	securemoteDnsServer := make(map[string]interface{})

	securemoteDnsServer["uid"] = d.Id()

	if d.HasChange("name") {
		if v, ok := d.GetOk("name"); ok {
			securemoteDnsServer["new-name"] = v
		}
	}

	if ok := d.HasChange("host"); ok {
		if v, ok := d.GetOk("host"); ok {
			securemoteDnsServer["host"] = v
		}
	}

	if d.HasChange("domains") {

		if v, ok := d.GetOk("domains"); ok {

			domainsList := v.([]interface{})

			var domainsPayload []map[string]interface{}

			for i := range domainsList {

				payload := make(map[string]interface{})

				//name, subnets, mask lengths are required to request
				if v, ok := d.GetOk("domains." + strconv.Itoa(i) + ".domain_suffix"); ok {
					payload["domain-suffix"] = v.(string)
				}
				if v, ok := d.GetOk("domains." + strconv.Itoa(i) + ".maximum_prefix_label_count"); ok {
					payload["maximum-prefix-label-count"] = v.(string)
				}

				domainsPayload = append(domainsPayload, payload)
			}

			securemoteDnsServer["domains"] = domainsPayload

		}
	}

	if ok := d.HasChange("comments"); ok {
		if v, ok := d.GetOk("comments"); ok {
			securemoteDnsServer["comments"] = v
		}
	}

	if ok := d.HasChange("color"); ok {
		if v, ok := d.GetOk("color"); ok {
			securemoteDnsServer["color"] = v
		}
	}

	if ok := d.HasChange("tags"); ok {
		if v, ok := d.GetOk("tags"); ok {
			securemoteDnsServer["tags"] = v.(*schema.Set).List()
		}
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		securemoteDnsServer["ignore-errors"] = v.(bool)
	}
	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		securemoteDnsServer["ignore-warnings"] = v.(bool)
	}

	log.Println("Update SecuremoteDnsServer - Map = ", securemoteDnsServer)
	if len(securemoteDnsServer) != 3 {
		updateSecuremoteDnsServerRes, err := client.ApiCallSimple("set-securemote-dns-server", securemoteDnsServer)
		if err != nil || !updateSecuremoteDnsServerRes.Success {
			if updateSecuremoteDnsServerRes.ErrorMsg != "" {
				return fmt.Errorf(updateSecuremoteDnsServerRes.ErrorMsg)
			}
			return fmt.Errorf(err.Error())
		}
	} else {
		// Payload contain only required fields: uid, ignore-warnings and ignore-errors
		// We got empty update, skip update API call...
		log.Println("Got empty update. Skip update API call...")
	}

	return readManagementSecuremoteDnsServer(d, m)
}

func deleteManagementSecuremoteDnsServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	securemoteDnsServerPayload := map[string]interface{}{
		"uid": d.Id(),
	}
	if v, ok := d.GetOkExists("ignore_errors"); ok {
		securemoteDnsServerPayload["ignore-errors"] = v.(bool)
	}
	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		securemoteDnsServerPayload["ignore-warnings"] = v.(bool)
	}
	deleteSecuremoteDnsServerRes, err := client.ApiCallSimple("delete-securemote-dns-server", securemoteDnsServerPayload)
	if err != nil || !deleteSecuremoteDnsServerRes.Success {
		if deleteSecuremoteDnsServerRes.ErrorMsg != "" {
			return fmt.Errorf(deleteSecuremoteDnsServerRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
