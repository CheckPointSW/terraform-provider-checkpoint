package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementNutanixDataCenterServer() *schema.Resource {
	return &schema.Resource{
		Create: createManagementNutanixDataCenterServer,
		Read:   readManagementNutanixDataCenterServer,
		Update: updateManagementNutanixDataCenterServer,
		Delete: deleteManagementNutanixDataCenterServer,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name. Must be unique in the domain.",
			},
			"hostname": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "IP Address or hostname of the Nutanix Prism server.",
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Username of the Nutanix Prism server.",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "Password of the Nutanix Prism server.",
			},
			"certificate_fingerprint": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specify the SHA-1 or SHA-256 fingerprint of the Data Center Server's certificate.",
			},
			"unsafe_auto_accept": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "When set to false, the current Data Center Server's certificate should be trusted, either by providing the certificate-fingerprint argument or by relying on a previously trusted certificate of this hostname. When set to true, trust the current Data Center Server's certificate as-is.",
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
				Default:     "black",
				Description: "Color of the object. Should be one of existing colors.",
			},
			"comments": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comments string",
			},
			"ignore_warnings": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Apply changes ignoring warnings. By Setting this parameter to 'true' test connection failure will be ignored.",
			},
			"ignore_errors": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.",
			},
			"automatic_refresh": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether the data center server's content is automatically updated.",
			},
			"data_center_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Data Center type.",
			},
		},
	}
}

func createManagementNutanixDataCenterServer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	nutanixDataCenterServer := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		nutanixDataCenterServer["name"] = v.(string)
	}

	nutanixDataCenterServer["type"] = "nutanix"

	if v, ok := d.GetOk("hostname"); ok {
		nutanixDataCenterServer["hostname"] = v.(string)
	}

	if v, ok := d.GetOk("username"); ok {
		nutanixDataCenterServer["username"] = v.(string)
	}

	if v, ok := d.GetOk("password"); ok {
		nutanixDataCenterServer["password"] = v.(string)
	}

	if v, ok := d.GetOk("certificate_fingerprint"); ok {
		nutanixDataCenterServer["certificate-fingerprint"] = v.(string)
	}

	if v, ok := d.GetOkExists("unsafe_auto_accept"); ok {
		nutanixDataCenterServer["unsafe-auto-accept"] = v.(bool)
	}

	if v, ok := d.GetOk("tags"); ok {
		nutanixDataCenterServer["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		nutanixDataCenterServer["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		nutanixDataCenterServer["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		nutanixDataCenterServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		nutanixDataCenterServer["ignore-errors"] = v.(bool)
	}

	log.Println("Create Nutanix Data Center Server - Map = ", nutanixDataCenterServer)

	addNutanixDataCenterServerRes, err := client.ApiCall("add-data-center-server", nutanixDataCenterServer, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !addNutanixDataCenterServerRes.Success {
		if addNutanixDataCenterServerRes.ErrorMsg != "" {
			return fmt.Errorf(addNutanixDataCenterServerRes.ErrorMsg)
		}
		msg := createTaskFailMessage("add-data-center-server", addNutanixDataCenterServerRes.GetData())
		return fmt.Errorf(msg)
	}
	payload := map[string]interface{}{
		"name": nutanixDataCenterServer["name"],
	}
	showNutanixDataCenterServerRes, err := client.ApiCall("show-data-center-server", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showNutanixDataCenterServerRes.Success {
		return fmt.Errorf(showNutanixDataCenterServerRes.ErrorMsg)
	}
	d.SetId(showNutanixDataCenterServerRes.GetData()["uid"].(string))
	return readManagementNutanixDataCenterServer(d, m)
}

func readManagementNutanixDataCenterServer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showNutanixDataCenterServerRes, err := client.ApiCall("show-data-center-server", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showNutanixDataCenterServerRes.Success {
		if objectNotFound(showNutanixDataCenterServerRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showNutanixDataCenterServerRes.ErrorMsg)
	}

	nutanixDataCenterServer := showNutanixDataCenterServerRes.GetData()

	log.Println("Read Nutanix Data Center - Show JSON = ", nutanixDataCenterServer)

	if v := nutanixDataCenterServer["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := nutanixDataCenterServer["data-center-type"]; v != nil {
		_ = d.Set("data_center_type", v)
	}

	if nutanixDataCenterServer["properties"] != nil {
		propertiesList := nutanixDataCenterServer["properties"].([]interface{})

		if len(propertiesList) > 0 {
			var propertiesListToReturn []map[string]interface{}

			for i := range propertiesList {
				propertiesMap := propertiesList[i].(map[string]interface{})

				propertiesMapToAdd := make(map[string]interface{})

				if v, _ := propertiesMap["name"]; v != nil {
					propertiesMapToAdd["name"] = v
				}
				if v, _ := propertiesMap["value"]; v != nil {
					propertiesMapToAdd["value"] = v
				}

				propertiesListToReturn = append(propertiesListToReturn, propertiesMapToAdd)
			}

			_ = d.Set("properties", propertiesListToReturn)

		} else {
			_ = d.Set("properties", propertiesList)
		}
	} else {
		_ = d.Set("properties", nil)
	}

	if v := nutanixDataCenterServer["automatic-refresh"]; v != nil {
		_ = d.Set("automatic_refresh", v)
	}

	if nutanixDataCenterServer["tags"] != nil {
		tagsJson, ok := nutanixDataCenterServer["tags"].([]interface{})
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

	return nil
}

func updateManagementNutanixDataCenterServer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	nutanixDataCenterServer := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		nutanixDataCenterServer["name"] = oldName
		nutanixDataCenterServer["new-name"] = newName
	} else {
		nutanixDataCenterServer["name"] = d.Get("name")
	}

	if ok := d.HasChange("hostname"); ok {
		nutanixDataCenterServer["hostname"] = d.Get("hostname")
	}

	if ok := d.HasChange("username"); ok {
		nutanixDataCenterServer["username"] = d.Get("username")
	}

	if ok := d.HasChange("password"); ok {
		nutanixDataCenterServer["password"] = d.Get("password")
	}

	if ok := d.HasChange("certificate_fingerprint"); ok {
		nutanixDataCenterServer["certificate-fingerprint"] = d.Get("certificate_fingerprint")
	}

	if ok := d.HasChange("unsafe_auto_accept"); ok {
		nutanixDataCenterServer["unsafe-auto-accept"] = d.Get("unsafe_auto_accept")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			nutanixDataCenterServer["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			nutanixDataCenterServer["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		nutanixDataCenterServer["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		nutanixDataCenterServer["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		nutanixDataCenterServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		nutanixDataCenterServer["ignore-errors"] = v.(bool)
	}

	log.Println("Update nutanixDataCenterServer - Map = ", nutanixDataCenterServer)

	updateNutanixDataCenterServerRes, err := client.ApiCall("set-data-center-server", nutanixDataCenterServer, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !updateNutanixDataCenterServerRes.Success {
		if updateNutanixDataCenterServerRes.ErrorMsg != "" {
			return fmt.Errorf(updateNutanixDataCenterServerRes.ErrorMsg)
		}
		msg := createTaskFailMessage("set-data-center-server", updateNutanixDataCenterServerRes.GetData())
		return fmt.Errorf(msg)
	}

	return readManagementNutanixDataCenterServer(d, m)
}

func deleteManagementNutanixDataCenterServer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	nutanixDataCenterServerPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		nutanixDataCenterServerPayload["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		nutanixDataCenterServerPayload["ignore-errors"] = v.(bool)
	}

	log.Println("Delete nutanixDataCenterServer")

	deleteNutanixDataCenterServerRes, err := client.ApiCall("delete-data-center-server", nutanixDataCenterServerPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteNutanixDataCenterServerRes.Success {
		if deleteNutanixDataCenterServerRes.ErrorMsg != "" {
			return fmt.Errorf(deleteNutanixDataCenterServerRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
