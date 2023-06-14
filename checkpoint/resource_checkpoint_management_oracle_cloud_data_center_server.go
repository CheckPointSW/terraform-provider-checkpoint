package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementOracleCloudDataCenterServer() *schema.Resource {
	return &schema.Resource{
		Create: createManagementOracleCloudDataCenterServer,
		Read:   readManagementOracleCloudDataCenterServer,
		Update: updateManagementOracleCloudDataCenterServer,
		Delete: deleteManagementOracleCloudDataCenterServer,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name. Must be unique in the domain.",
			},
			"authentication_method": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "key-authentication Uses the Service Account private key file to authenticate. vm-instance-authentication Uses VM Instance to authenticate. This option requires the Security Management Server deployed in Oracle Cloud, and running in a dynamic group with the required permissions",
			},
			"private_key": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: " An Oracle Cloud API key PEM file, encoded in base64. Required for authentication-method: key-authentication.",
			},
			"key_user": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "An Oracle Cloud user id associated with key. Required for authentication-method: key-authentication.",
			},
			"key_tenant": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "An Oracle Cloud tenancy id where the key was created. Required for authentication-method: key-authentication.",
			},
			"key_region": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "An Oracle Cloud region for where to create scanner. Required for authentication-method: key-authentication.",
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
				Description: "Color of the object. Should be one of existing colors.",
			},
			"comments": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comments string.",
			},
			"ignore_warnings": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Apply changes ignoring warnings. By Setting this parameter to 'true' test connection failure will be ignored.",
			},
			"ignore_errors": {
				Type:        schema.TypeBool,
				Optional:    true,
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

func createManagementOracleCloudDataCenterServer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	oracleCloudDataCenterServer := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		oracleCloudDataCenterServer["name"] = v.(string)
	}

	oracleCloudDataCenterServer["type"] = "oci"

	if v, ok := d.GetOk("authentication_method"); ok {
		oracleCloudDataCenterServer["authentication-method"] = v.(string)
	}

	if v, ok := d.GetOk("private_key"); ok {
		oracleCloudDataCenterServer["private-key"] = v.(string)
	}

	if v, ok := d.GetOk("key_user"); ok {
		oracleCloudDataCenterServer["key-user"] = v.(string)
	}

	if v, ok := d.GetOk("key_tenant"); ok {
		oracleCloudDataCenterServer["key-tenant"] = v.(string)
	}

	if v, ok := d.GetOk("key_region"); ok {
		oracleCloudDataCenterServer["key-region"] = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		oracleCloudDataCenterServer["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		oracleCloudDataCenterServer["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		oracleCloudDataCenterServer["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		oracleCloudDataCenterServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		oracleCloudDataCenterServer["ignore-errors"] = v.(bool)
	}

	log.Println("Create oracleCloudDataCenterServer - Map = ", oracleCloudDataCenterServer)

	addOracleCloudDataCenterServerRes, err := client.ApiCall("add-data-center-server", oracleCloudDataCenterServer, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !addOracleCloudDataCenterServerRes.Success {
		if addOracleCloudDataCenterServerRes.ErrorMsg != "" {
			return fmt.Errorf(addOracleCloudDataCenterServerRes.ErrorMsg)
		}
		msg := createTaskFailMessage("add-data-center-server", addOracleCloudDataCenterServerRes.GetData())
		return fmt.Errorf(msg)
	}
	payload := map[string]interface{}{
		"name": oracleCloudDataCenterServer["name"],
	}

	showOracleCloudDataCenterServerRes, err := client.ApiCall("show-data-center-server", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showOracleCloudDataCenterServerRes.Success {
		return fmt.Errorf(showOracleCloudDataCenterServerRes.ErrorMsg)
	}
	d.SetId(showOracleCloudDataCenterServerRes.GetData()["uid"].(string))

	return readManagementOracleCloudDataCenterServer(d, m)
}

func readManagementOracleCloudDataCenterServer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showOracleCloudDataCenterServerRes, err := client.ApiCall("show-data-center-server", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showOracleCloudDataCenterServerRes.Success {
		if objectNotFound(showOracleCloudDataCenterServerRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showOracleCloudDataCenterServerRes.ErrorMsg)
	}

	oracleCloudDataCenterServer := showOracleCloudDataCenterServerRes.GetData()

	if v := oracleCloudDataCenterServer["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if oracleCloudDataCenterServer["properties"] != nil {
		propertiesList := oracleCloudDataCenterServer["properties"].([]interface{})

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

	if oracleCloudDataCenterServer["tags"] != nil {
		tagsJson, ok := oracleCloudDataCenterServer["tags"].([]interface{})
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

	if v := oracleCloudDataCenterServer["automatic-refresh"]; v != nil {
		_ = d.Set("automatic_refresh", v)
	}

	if v := oracleCloudDataCenterServer["data-center-type"]; v != nil {
		_ = d.Set("data_center_type", v)
	}

	return nil
}

func updateManagementOracleCloudDataCenterServer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	oracleCloudDataCenterServer := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		oracleCloudDataCenterServer["name"] = oldName
		oracleCloudDataCenterServer["new-name"] = newName
	} else {
		oracleCloudDataCenterServer["name"] = d.Get("name")
	}

	if ok := d.HasChange("authentication_method"); ok {
		oracleCloudDataCenterServer["authentication-method"] = d.Get("authentication_method")
	}

	if ok := d.HasChange("private_key"); ok {
		oracleCloudDataCenterServer["private-key"] = d.Get("private_key")
	}

	if ok := d.HasChange("key_user"); ok {
		oracleCloudDataCenterServer["key-user"] = d.Get("key_user")
	}

	if ok := d.HasChange("key_tenant"); ok {
		oracleCloudDataCenterServer["key-tenant"] = d.Get("key_tenant")
	}

	if ok := d.HasChange("key_region"); ok {
		oracleCloudDataCenterServer["key-region"] = d.Get("key_region")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			oracleCloudDataCenterServer["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			oracleCloudDataCenterServer["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		oracleCloudDataCenterServer["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		oracleCloudDataCenterServer["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		oracleCloudDataCenterServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		oracleCloudDataCenterServer["ignore-errors"] = v.(bool)
	}

	log.Println("Update oracleCloudDataCenterServer - Map = ", oracleCloudDataCenterServer)

	updateOracleCloudDataCenterServerRes, err := client.ApiCall("set-data-center-server", oracleCloudDataCenterServer, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !updateOracleCloudDataCenterServerRes.Success {
		if updateOracleCloudDataCenterServerRes.ErrorMsg != "" {
			return fmt.Errorf(updateOracleCloudDataCenterServerRes.ErrorMsg)
		}
		msg := createTaskFailMessage("set-data-center-server", updateOracleCloudDataCenterServerRes.GetData())
		return fmt.Errorf(msg)
	}

	return readManagementOracleCloudDataCenterServer(d, m)
}

func deleteManagementOracleCloudDataCenterServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	oracleCloudDataCenterServerPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		oracleCloudDataCenterServerPayload["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		oracleCloudDataCenterServerPayload["ignore-errors"] = v.(bool)
	}

	deleteOracleCloudDataCenterServerRes, err := client.ApiCall("delete-data-center-server", oracleCloudDataCenterServerPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteOracleCloudDataCenterServerRes.Success {
		if deleteOracleCloudDataCenterServerRes.ErrorMsg != "" {
			return fmt.Errorf(deleteOracleCloudDataCenterServerRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
