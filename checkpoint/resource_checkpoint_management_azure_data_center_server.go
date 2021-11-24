package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strings"
)

func resourceManagementAzureDataCenterServer() *schema.Resource {
	return &schema.Resource{
		Create: createManagementAzureDataCenterServer,
		Read:   readManagementAzureDataCenterServer,
		Update: updateManagementAzureDataCenterServer,
		Delete: deleteManagementAzureDataCenterServer,
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
				Description: "user-authentication\nUses the Azure AD User to authenticate.\nservice-principal-authentication\nUses the Service Principal to authenticate.",
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "An Azure Active Directory user Format <username>@<domain>.\nRequired for authentication-method: user-authentication.",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Password of the Azure account.\nRequired for authentication-method: user-authentication.",
			},
			"password_base64": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Password of the Azure account encoded in Base64.\nRequired for authentication-method: user-authentication.",
			},
			"application_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Application ID of the Service Principal, in UUID format.\nRequired for authentication-method: service-principal-authentication.",
			},
			"application_key": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "The key created for the Service Principal.\nRequired for authentication-method: service-principal-authentication.",
				Default:     false,
			},
			"directory_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Directory ID of the Azure AD, in UUID format.\nRequired for authentication-method: service-principal-authentication.",
			},
			"environment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Select the Azure Cloud Environment.",
				Default:     "AzureCloud",
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
				Description: "Apply changes ignoring warnings. By Setting this parameter to 'true' test connection failure will be ignored.",
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

func createManagementAzureDataCenterServer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	azureDataCenterServer := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		azureDataCenterServer["name"] = v.(string)
	}

	azureDataCenterServer["type"] = "azure"

	if v, ok := d.GetOk("authentication_method"); ok {
		azureDataCenterServer["authentication-method"] = v.(string)
	}

	if v, ok := d.GetOk("username"); ok {
		azureDataCenterServer["username"] = v.(string)
	}

	if v, ok := d.GetOk("password"); ok {
		azureDataCenterServer["password"] = v.(string)
	}

	if v, ok := d.GetOk("password_base64"); ok {
		azureDataCenterServer["password-base64"] = v.(string)
	}

	if v, ok := d.GetOk("application_id"); ok {
		azureDataCenterServer["application-id"] = v.(string)
	}

	if v, ok := d.GetOk("application_key"); ok {
		azureDataCenterServer["application-key"] = v.(string)
	}

	if v, ok := d.GetOk("directory_id"); ok {
		azureDataCenterServer["directory-id"] = v.(string)
	}
	if v, ok := d.GetOk("environment"); ok {
		azureDataCenterServer["environment"] = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		azureDataCenterServer["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		azureDataCenterServer["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		azureDataCenterServer["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		azureDataCenterServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		azureDataCenterServer["ignore-errors"] = v.(bool)
	}

	log.Println("Create azureDataCenterServer - Map = ", azureDataCenterServer)

	addAzureDataCenterServerRes, err := client.ApiCall("add-data-center-server", azureDataCenterServer, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !addAzureDataCenterServerRes.Success {
		if addAzureDataCenterServerRes.ErrorMsg != "" {
			return fmt.Errorf(addAzureDataCenterServerRes.ErrorMsg)
		}
		msg := createTaskFailMessage("add-data-center-server", addAzureDataCenterServerRes.GetData())
		return fmt.Errorf(msg)
	}
	payload := map[string]interface{}{
		"name": azureDataCenterServer["name"],
	}
	showAzureDataCenterServerRes, err := client.ApiCall("show-data-center-server", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showAzureDataCenterServerRes.Success {
		return fmt.Errorf(showAzureDataCenterServerRes.ErrorMsg)
	}
	d.SetId(showAzureDataCenterServerRes.GetData()["uid"].(string))
	return readManagementAzureDataCenterServer(d, m)
}

func readManagementAzureDataCenterServer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showAzureDataCenterServerRes, err := client.ApiCall("show-data-center-server", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showAzureDataCenterServerRes.Success {
		if objectNotFound(showAzureDataCenterServerRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showAzureDataCenterServerRes.ErrorMsg)
	}
	azureDataCenterServer := showAzureDataCenterServerRes.GetData()

	if v := azureDataCenterServer["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if azureDataCenterServer["properties"] != nil {
		propsJson, ok := azureDataCenterServer["properties"].([]interface{})
		if ok {
			for _, prop := range propsJson {
				propMap := prop.(map[string]interface{})
				propName := strings.ReplaceAll(propMap["name"].(string), "-", "_")
				propValue := propMap["value"]
				_ = d.Set(propName, propValue)
			}
		}
	}

	if azureDataCenterServer["tags"] != nil {
		tagsJson, ok := azureDataCenterServer["tags"].([]interface{})
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

	if v := azureDataCenterServer["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := azureDataCenterServer["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := azureDataCenterServer["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := azureDataCenterServer["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementAzureDataCenterServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	azureDataCenterServer := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		azureDataCenterServer["name"] = oldName
		azureDataCenterServer["new-name"] = newName
	} else {
		azureDataCenterServer["name"] = d.Get("name")
	}

	if ok := d.HasChange("authentication_method"); ok {
		azureDataCenterServer["authentication-method"] = d.Get("authentication_method")
	}

	if ok := d.HasChange("password"); ok {
		azureDataCenterServer["password"] = d.Get("password")
	}

	if ok := d.HasChange("password_base64"); ok {
		azureDataCenterServer["password-base64"] = d.Get("password_base64")
	}

	if d.HasChange("username") {
		azureDataCenterServer["username"] = d.Get("username")
		if v := d.Get("password"); v != nil && v != "" {
			azureDataCenterServer["password"] = v
		}
		if v := d.Get("password_base64"); v != nil && v != "" {
			azureDataCenterServer["password-base64"] = v
		}
	}

	if ok := d.HasChange("application_id"); ok {
		azureDataCenterServer["application-id"] = d.Get("application_id")
	}

	if ok := d.HasChange("application_key"); ok {
		azureDataCenterServer["application-key"] = d.Get("application_key")
	}

	if ok := d.HasChange("directory_id"); ok {
		azureDataCenterServer["directory-id"] = d.Get("directory_id")
	}

	if ok := d.HasChange("environment"); ok {
		azureDataCenterServer["environment"] = d.Get("environment")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			azureDataCenterServer["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			azureDataCenterServer["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		azureDataCenterServer["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		azureDataCenterServer["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		azureDataCenterServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		azureDataCenterServer["ignore-errors"] = v.(bool)
	}

	log.Println("Update azureDataCenterServer - Map = ", azureDataCenterServer)

	updateAzureDataCenterServerRes, err := client.ApiCall("set-data-center-server", azureDataCenterServer, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !updateAzureDataCenterServerRes.Success {
		if updateAzureDataCenterServerRes.ErrorMsg != "" {
			return fmt.Errorf(updateAzureDataCenterServerRes.ErrorMsg)
		}
		msg := createTaskFailMessage("set-data-center-server", updateAzureDataCenterServerRes.GetData())
		return fmt.Errorf(msg)
	}

	return readManagementAzureDataCenterServer(d, m)
}

func deleteManagementAzureDataCenterServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	azureDataCenterServerPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete azureDataCenterServer")

	deleteAzureDataCenterServerRes, err := client.ApiCall("delete-data-center-server", azureDataCenterServerPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteAzureDataCenterServerRes.Success {
		if deleteAzureDataCenterServerRes.ErrorMsg != "" {
			return fmt.Errorf(deleteAzureDataCenterServerRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
