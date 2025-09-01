package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strconv"
)

func resourceManagementIfMapServer() *schema.Resource {
	return &schema.Resource{
		Create: createManagementIfMapServer,
		Read:   readManagementIfMapServer,
		Update: updateManagementIfMapServer,
		Delete: deleteManagementIfMapServer,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "IF-MAP server port number.",
				Default:     443,
			},
			"version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IF-MAP version.",
				Default:     "2.0",
			},
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Host that is IF-MAP server.  Identified by name or UID.",
			},
			"path": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "N/A",
			},
			"monitored_ips": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "IP ranges to be monitored by the IF-MAP client.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"first_ip": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "First IPv4 address in the range to be monitored.",
							Default:     "0.0.0.0",
						},
						"last_ip": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Last IPv4 address in the range to be monitored.",
							Default:     "0.0.0.0",
						},
					},
				},
			},
			"query_whole_ranges": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicate whether to query whole ranges instead of single IP.",
				Default:     true,
			},
			"authentication": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Authentication configuration for the IF-MAP server.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication_method": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Authentication method for the IF-MAP server.",
						},
						"username": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Username for the IF-MAP server authentication. <font color=\"red\">Required only when</font> 'authentication-method' is set to 'basic'.",
						},
						"password": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Username for the IF-MAP server authentication. <font color=\"red\">Required only when</font> 'authentication-method' is set to 'basic'.",
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
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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

func createManagementIfMapServer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	ifMapServer := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		ifMapServer["name"] = v.(string)
	}

	if v, ok := d.GetOk("port"); ok {
		ifMapServer["port"] = v.(int)
	}

	if v, ok := d.GetOk("version"); ok {
		ifMapServer["version"] = v.(string)
	}

	if v, ok := d.GetOk("host"); ok {
		ifMapServer["host"] = v.(string)
	}

	if v, ok := d.GetOk("path"); ok {
		ifMapServer["path"] = v.(string)
	}

	if v, ok := d.GetOk("monitored_ips"); ok {

		monitoredIpsList := v.([]interface{})

		if len(monitoredIpsList) > 0 {

			var monitoredIpsPayload []map[string]interface{}

			for i := range monitoredIpsList {

				Payload := make(map[string]interface{})

				if v, ok := d.GetOk("monitored_ips." + strconv.Itoa(i) + ".first_ip"); ok {
					Payload["first-ip"] = v.(string)
				}
				if v, ok := d.GetOk("monitored_ips." + strconv.Itoa(i) + ".last_ip"); ok {
					Payload["last-ip"] = v.(string)
				}
				monitoredIpsPayload = append(monitoredIpsPayload, Payload)
			}
			ifMapServer["monitored-ips"] = monitoredIpsPayload
		}
	}

	if v, ok := d.GetOkExists("query_whole_ranges"); ok {
		ifMapServer["query-whole-ranges"] = v.(bool)
	}

	if _, ok := d.GetOk("authentication"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("authentication.0.authentication_method"); ok {
			res["authentication-method"] = v.(string)
		}
		if v, ok := d.GetOk("authentication.0.username"); ok {
			res["username"] = v.(string)
		}
		if v, ok := d.GetOk("authentication.0.password"); ok {
			res["password"] = v.(string)
		}
		ifMapServer["authentication"] = res
	}

	if v, ok := d.GetOk("color"); ok {
		ifMapServer["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		ifMapServer["comments"] = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		ifMapServer["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		ifMapServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		ifMapServer["ignore-errors"] = v.(bool)
	}

	log.Println("Create IfMapServer - Map = ", ifMapServer)

	addIfMapServerRes, err := client.ApiCallSimple("add-if-map-server", ifMapServer)
	if err != nil || !addIfMapServerRes.Success {
		if addIfMapServerRes.ErrorMsg != "" {
			return fmt.Errorf(addIfMapServerRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addIfMapServerRes.GetData()["uid"].(string))

	return readManagementIfMapServer(d, m)
}

func readManagementIfMapServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showIfMapServerRes, err := client.ApiCallSimple("show-if-map-server", payload)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showIfMapServerRes.Success {
		if objectNotFound(showIfMapServerRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showIfMapServerRes.ErrorMsg)
	}

	ifMapServer := showIfMapServerRes.GetData()

	log.Println("Read IfMapServer - Show JSON = ", ifMapServer)

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
		_ = d.Set("host", v)
	}

	if v := ifMapServer["path"]; v != nil {
		_ = d.Set("path", v)
	}

	if ifMapServer["monitored-ips"] != nil {

		monitoredIpsList, ok := ifMapServer["monitored-ips"].([]interface{})

		if ok {

			if len(monitoredIpsList) > 0 {

				var monitoredIpsListToReturn []map[string]interface{}

				for i := range monitoredIpsList {

					monitoredIpsMap := monitoredIpsList[i].(map[string]interface{})

					monitoredIpsMapToAdd := make(map[string]interface{})

					if v, _ := monitoredIpsMap["first-ip"]; v != nil {
						monitoredIpsMapToAdd["first_ip"] = v
					}
					if v, _ := monitoredIpsMap["last-ip"]; v != nil {
						monitoredIpsMapToAdd["last_ip"] = v
					}
					monitoredIpsListToReturn = append(monitoredIpsListToReturn, monitoredIpsMapToAdd)
				}
			}
		}
	}

	if v := ifMapServer["query-whole-ranges"]; v != nil {
		_ = d.Set("query_whole_ranges", v)
	}

	if ifMapServer["authentication"] != nil {

		authenticationMap := ifMapServer["authentication"].(map[string]interface{})

		authenticationMapToReturn := make(map[string]interface{})

		if v, _ := authenticationMap["authentication-method"]; v != nil {
			authenticationMapToReturn["authentication_method"] = v
		}
		if v, _ := authenticationMap["username"]; v != nil {
			authenticationMapToReturn["username"] = v
		}
		if v, _ := authenticationMap["password"]; v != nil {
			authenticationMapToReturn["password"] = v
		}

		if len(authenticationMapToReturn) > 0 {
			_ = d.Set("authentication", []interface{}{authenticationMapToReturn})
		} else {
			_ = d.Set("authentication", nil)
		}
	} else {
		_ = d.Set("authentication", nil)
	}

	if v := ifMapServer["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := ifMapServer["comments"]; v != nil {
		_ = d.Set("comments", v)
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

	if v := ifMapServer["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := ifMapServer["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementIfMapServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	ifMapServer := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		ifMapServer["name"] = oldName
		ifMapServer["new-name"] = newName
	} else {
		ifMapServer["name"] = d.Get("name")
	}

	if ok := d.HasChange("port"); ok {
		ifMapServer["port"] = d.Get("port")
	}

	if ok := d.HasChange("version"); ok {
		ifMapServer["version"] = d.Get("version")
	}

	if ok := d.HasChange("host"); ok {
		ifMapServer["host"] = d.Get("host")
	}

	if ok := d.HasChange("path"); ok {
		ifMapServer["path"] = d.Get("path")
	}

	if d.HasChange("monitored_ips") {

		if v, ok := d.GetOk("monitored_ips"); ok {

			monitoredIpsList := v.([]interface{})

			var monitoredIpsPayload []map[string]interface{}

			for i := range monitoredIpsList {

				Payload := make(map[string]interface{})

				if v, ok := d.GetOk("monitored_ips." + strconv.Itoa(i) + ".first_ip"); ok {
					Payload["first-ip"] = v.(string)
				}

				if v, ok := d.GetOk("monitored_ips." + strconv.Itoa(i) + ".last_ip"); ok {
					Payload["last-ip"] = v.(string)
				}

				monitoredIpsPayload = append(monitoredIpsPayload, Payload)
			}
			ifMapServer["monitored-ips"] = monitoredIpsPayload
		}
	}

	if v, ok := d.GetOkExists("query_whole_ranges"); ok {
		ifMapServer["query-whole-ranges"] = v.(bool)
	}

	if d.HasChange("authentication") {

		if _, ok := d.GetOk("authentication"); ok {

			res := make(map[string]interface{})

			if v, ok := d.GetOk("authentication.0.authentication_method"); ok {
				res["authentication-method"] = v.(string)
			}

			if v, ok := d.GetOk("authentication.0.username"); ok {
				res["username"] = v.(string)
			}

			if v, ok := d.GetOk("authentication.0.password"); ok {
				res["password"] = v.(string)
			}

			ifMapServer["authentication"] = res
		}
	}

	if ok := d.HasChange("color"); ok {
		ifMapServer["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		ifMapServer["comments"] = d.Get("comments")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			ifMapServer["tags"] = v.(*schema.Set).List()
		}
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		ifMapServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		ifMapServer["ignore-errors"] = v.(bool)
	}

	log.Println("Update IfMapServer - Map = ", ifMapServer)

	updateIfMapServerRes, err := client.ApiCallSimple("set-if-map-server", ifMapServer)
	if err != nil || !updateIfMapServerRes.Success {
		if updateIfMapServerRes.ErrorMsg != "" {
			return fmt.Errorf(updateIfMapServerRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementIfMapServer(d, m)
}

func deleteManagementIfMapServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	ifMapServerPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete IfMapServer")

	deleteIfMapServerRes, err := client.ApiCallSimple("delete-if-map-server", ifMapServerPayload)
	if err != nil || !deleteIfMapServerRes.Success {
		if deleteIfMapServerRes.ErrorMsg != "" {
			return fmt.Errorf(deleteIfMapServerRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
