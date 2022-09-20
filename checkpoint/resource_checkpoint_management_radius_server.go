package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"reflect"
	"strconv"
)

func resourceManagementRadiusServer() *schema.Resource {
	return &schema.Resource{
		Create: createManagementRadiusServer,
		Read:   readManagementRadiusServer,
		Update: updateManagementRadiusServer,
		Delete: deleteManagementRadiusServer,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name. Must be unique in the domain.",
			},
			"server": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The UID or Name of the host that is the RADIUS Server.",
			},
			"shared_secret": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The secret between the RADIUS server and the Security Gateway.",
				Sensitive:   true,
			},
			"service": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "RADIUS",
				Description: "The UID or Name of the Service to which the RADIUS server listens.",
			},
			"version": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "RADIUS Ver. 1.0",
				Description: "The version can be either RADIUS Version 1.0, which is RFC 2138 compliant, and RADIUS Version 2.0 which is RFC 2865 compliant.",
			},
			"protocol": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "PAP",
				Description: "The type of authentication protocol that will be used when authenticating the user to the RADIUS server.",
			},
			"priority": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     1,
				Description: "The priority of the RADIUS Server in case it is a member of a RADIUS Group.",
			},
			"accounting": &schema.Schema{
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Accounting settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_ip_pool_management": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "IP pool management, enables Accounting service.",
						},
						"accounting_service": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The UID or Name of the the accounting interface to notify the server when users login and logout which will then lock and release the IP addresses that the server allocated to those users.",
						},
					},
				},
			},
			"tags": &schema.Schema{
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"color": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "black",
				Description: "Color of the object. Should be one of existing colors.",
			},
			"comments": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comments string.",
			},
			"ignore_warnings": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Apply changes ignoring warnings.",
			},
			"ignore_errors": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.\nApply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.",
			},
		},
	}
}

func createManagementRadiusServer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	radiusServer := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		radiusServer["name"] = v.(string)
	}

	if v, ok := d.GetOk("server"); ok {
		radiusServer["server"] = v.(string)
	}

	if v, ok := d.GetOk("shared_secret"); ok {
		radiusServer["shared-secret"] = v.(string)
	}

	if v, ok := d.GetOk("service"); ok {
		radiusServer["service"] = v.(string)
	}

	if v, ok := d.GetOk("version"); ok {
		radiusServer["version"] = v.(string)
	}

	if v, ok := d.GetOk("protocol"); ok {
		radiusServer["protocol"] = v.(string)
	}

	if v, ok := d.GetOk("priority"); ok {
		radiusServer["priority"] = v.(int)
	}

	if _, ok := d.GetOk("accounting"); ok {
		res := make(map[string]interface{})

		if v, ok := d.GetOk("accounting.enable_ip_pool_management"); ok {
			res["enable-ip-pool-management"] = v
		}

		if v, ok := d.GetOk("accounting.accounting_service"); ok {
			res["accounting-service"] = v.(string)
		}

		radiusServer["accounting"] = res
	}

	if v, ok := d.GetOk("tags"); ok {
		radiusServer["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		radiusServer["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		radiusServer["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		radiusServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		radiusServer["ignore-errors"] = v.(bool)
	}

	log.Println("Create Radius Server - Map = ", radiusServer)

	addRadiusServerRes, err := client.ApiCall("add-radius-server", radiusServer, client.GetSessionID(), true, client.IsProxyUsed())

	if err != nil || !addRadiusServerRes.Success {
		if addRadiusServerRes.ErrorMsg != "" {
			return fmt.Errorf(addRadiusServerRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addRadiusServerRes.GetData()["uid"].(string))

	return readManagementRadiusServer(d, m)
}

func readManagementRadiusServer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showRadiusServerRes, err := client.ApiCall("show-radius-server", payload, client.GetSessionID(), true, client.IsProxyUsed())

	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showRadiusServerRes.Success {
		// Handle delete resource from other clients
		if objectNotFound(showRadiusServerRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showRadiusServerRes.ErrorMsg)
	}

	radiusServer := showRadiusServerRes.GetData()

	log.Println("Read Radius Server - Show JSON = ", radiusServer)

	if v := radiusServer["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := radiusServer["server"]; v != nil {
		_ = d.Set("server", v)
	}

	if v := radiusServer["shared-secret"]; v != nil {
		_ = d.Set("shared_secret", v)
	}

	if v := radiusServer["service"]; v != nil {
		_ = d.Set("service", v)
	}

	if v := radiusServer["version"]; v != nil {
		_ = d.Set("version", v)
	}

	if v := radiusServer["protocol"]; v != nil {
		_ = d.Set("protocol", v)
	}

	if v := radiusServer["priority"]; v != nil {
		_ = d.Set("priority", v)
	}

	if radiusServer["accounting"] != nil {
		accountingMap := radiusServer["accounting"].(map[string]interface{})
		accountingMapToReturn := make(map[string]interface{})

		if v, _ := accountingMap["enable-ip-pool-management"]; v != nil {
			accountingMapToReturn["enable_ip_pool_management"] = strconv.FormatBool(v.(bool))
		}

		if v, _ := accountingMap["accounting-service"]; v != nil && v != "" {
			accountingMapToReturn["accounting_service"] = v
		}
		_, accountingInConf := d.GetOk("accounting")
		defaultAccounting := map[string]interface{}{"enable_ip_pool_management": "false"}
		if reflect.DeepEqual(defaultAccounting, accountingMapToReturn) && !accountingInConf {
			_ = d.Set("accounting", map[string]interface{}{})
		} else {
			_ = d.Set("accounting", accountingMapToReturn)
		}

	} else {
		_ = d.Set("accounting", nil)
	}

	if radiusServer["tags"] != nil {
		tagsJson := radiusServer["tags"].([]interface{})
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

	if v := radiusServer["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := radiusServer["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := radiusServer["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := radiusServer["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}
	return nil
}

func updateManagementRadiusServer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	radiusServer := make(map[string]interface{})

	if d.HasChange("name") {
		oldName, newName := d.GetChange("name")
		radiusServer["name"] = oldName
		radiusServer["new-name"] = newName
	} else {
		radiusServer["name"] = d.Get("name")
	}

	if ok := d.HasChange("server"); ok {
		radiusServer["server"] = d.Get("server")
	}

	if ok := d.HasChange("shared_secret"); ok {
		radiusServer["shared-secret"] = d.Get("shared_secret")
	}

	if ok := d.HasChange("service"); ok {
		radiusServer["service"] = d.Get("service")
	}

	if ok := d.HasChange("version"); ok {
		radiusServer["version"] = d.Get("version")
	}

	if ok := d.HasChange("protocol"); ok {
		radiusServer["protocol"] = d.Get("protocol")
	}

	if ok := d.HasChange("priority"); ok {
		radiusServer["priority"] = d.Get("priority")
	}

	if ok := d.HasChange("accounting"); ok {
		if _, ok := d.GetOk("accounting"); ok {
			res := make(map[string]interface{})

			if v, ok := d.GetOk("accounting.enable_ip_pool_management"); ok {
				res["enable-ip-pool-management"] = v
			}

			if v, ok := d.GetOk("accounting.accounting_service"); ok {
				res["accounting-service"] = v.(string)
			}

			radiusServer["accounting"] = res
		}
	} else {
		radiusServer["accounting"] = map[string]interface{}{"enable-ip-pool-management": "false"}
	}

	if ok := d.HasChange("tags"); ok {
		if v, ok := d.GetOk("tags"); ok {
			radiusServer["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			radiusServer["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		radiusServer["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		radiusServer["comments"] = d.Get("comments")
	}

	if ok := d.HasChange("ignore_warnings"); ok {
		radiusServer["ignore-warnings"] = d.Get("ignore_warnings")
	}

	if ok := d.HasChange("ignore_errors"); ok {
		radiusServer["ignore-errors"] = d.Get("ignore_errors")
	}

	log.Println("Update Radius Server - Map = ", radiusServer)
	updateRadiusServerRes, err := client.ApiCall("set-radius-server", radiusServer, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !updateRadiusServerRes.Success {
		if updateRadiusServerRes.ErrorMsg != "" {
			return fmt.Errorf(updateRadiusServerRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementRadiusServer(d, m)
}

func deleteManagementRadiusServer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	radiusServerPayload := map[string]interface{}{
		"uid":             d.Id(),
		"ignore-warnings": "true",
	}

	deleteRadiusServerRes, err := client.ApiCall("delete-radius-server", radiusServerPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteRadiusServerRes.Success {
		if deleteRadiusServerRes.ErrorMsg != "" {
			return fmt.Errorf(deleteRadiusServerRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
