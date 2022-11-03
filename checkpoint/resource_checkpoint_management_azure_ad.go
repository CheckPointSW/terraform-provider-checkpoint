package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementAzureAd() *schema.Resource {
	return &schema.Resource{
		Create: createManagementAzureAd,
		Read:   readManagementAzureAd,
		Update: updateManagementAzureAd,
		Delete: deleteManagementAzureAd,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"authentication_method": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "user-authentication uses the Azure AD User to authenticate. service-principal-authentication uses the Service Principal to authenticate.",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "Password of the Azure account. Required for authentication-method: user-authentication.",
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "An Azure Active Directory user Format <username>@<domain>. Required for authentication-method: user-authentication",
			},
			"application_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Application ID of the Service Principal, in UUID format. Required for authentication-method: service-principal-authentication.",
			},
			"application_key": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The key created for the Service Principal. Required for authentication-method: service-principal-authentication.",
			},
			"directory_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Directory ID of the Azure AD, in UUID format. Required for authentication-method: service-principal-authentication.",
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
				Description: "Apply changes ignoring warnings.",
				Default:     false,
			},
			"ignore_errors": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.",
				Default:     false,
			},
			"task_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Azure AD Operation task-id, use show-task command to check the progress of the task.",
			},
		},
	}
}

func createManagementAzureAd(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	azureAd := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		azureAd["name"] = v.(string)
	}

	if v, ok := d.GetOk("authentication_method"); ok {
		azureAd["authentication-method"] = v.(string)
	}

	if v, ok := d.GetOk("password"); ok {
		azureAd["password"] = v.(string)
	}

	if v, ok := d.GetOk("username"); ok {
		azureAd["username"] = v.(string)
	}

	if v, ok := d.GetOk("application_id"); ok {
		azureAd["application-id"] = v.(string)
	}

	if v, ok := d.GetOk("application_key"); ok {
		azureAd["application-key"] = v.(string)
	}

	if v, ok := d.GetOk("directory_id"); ok {
		azureAd["directory-id"] = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		azureAd["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		azureAd["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		azureAd["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		azureAd["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		azureAd["ignore-errors"] = v.(bool)
	}

	log.Println("Create AzureAd - Map = ", azureAd)

	addAzureAdRes, err := client.ApiCall("add-azure-ad", azureAd, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !addAzureAdRes.Success {
		if addAzureAdRes.ErrorMsg != "" {
			return fmt.Errorf(addAzureAdRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addAzureAdRes.GetData()["uid"].(string))
	_ = d.Set("task_id", resolveTaskId(addAzureAdRes.GetData()))

	return readManagementAzureAd(d, m)
}

func readManagementAzureAd(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showAzureAdRes, err := client.ApiCall("show-azure-ad", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showAzureAdRes.Success {
		if objectNotFound(showAzureAdRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showAzureAdRes.ErrorMsg)
	}

	azureAd := showAzureAdRes.GetData()

	log.Println("Read AzureAd - Show JSON = ", azureAd)

	if v := azureAd["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := azureAd["authentication-method"]; v != nil {
		_ = d.Set("authentication_method", v)
	}

	if v := azureAd["password"]; v != nil {
		_ = d.Set("password", v)
	}

	if v := azureAd["username"]; v != nil {
		_ = d.Set("username", v)
	}

	if v := azureAd["application-id"]; v != nil {
		_ = d.Set("application_id", v)
	}

	if v := azureAd["application-key"]; v != nil {
		_ = d.Set("application_key", v)
	}

	if v := azureAd["directory-id"]; v != nil {
		_ = d.Set("directory_id", v)
	}

	if azureAd["tags"] != nil {
		tagsJson, ok := azureAd["tags"].([]interface{})
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

	if v := azureAd["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := azureAd["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if azureAd["properties"] != nil {
		propertiesList := azureAd["properties"].([]interface{})

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

	return nil

}

func updateManagementAzureAd(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	azureAd := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		azureAd["name"] = oldName
		azureAd["new-name"] = newName
	} else {
		azureAd["name"] = d.Get("name")
	}

	if ok := d.HasChange("authentication_method"); ok {
		azureAd["authentication-method"] = d.Get("authentication_method")
	}

	if ok := d.HasChange("password"); ok {
		azureAd["password"] = d.Get("password")
	}

	if ok := d.HasChange("username"); ok {
		azureAd["username"] = d.Get("username")
	}

	if ok := d.HasChange("application_id"); ok {
		azureAd["application-id"] = d.Get("application_id")
	}

	if ok := d.HasChange("application_key"); ok {
		azureAd["application-key"] = d.Get("application_key")
	}

	if ok := d.HasChange("directory_id"); ok {
		azureAd["directory-id"] = d.Get("directory_id")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			azureAd["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			azureAd["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		azureAd["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		azureAd["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		azureAd["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		azureAd["ignore-errors"] = v.(bool)
	}

	log.Println("Update AzureAd - Map = ", azureAd)

	updateAzureAdRes, err := client.ApiCall("set-azure-ad", azureAd, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !updateAzureAdRes.Success {
		if updateAzureAdRes.ErrorMsg != "" {
			return fmt.Errorf(updateAzureAdRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementAzureAd(d, m)
}

func deleteManagementAzureAd(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	azureAdPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete AzureAd")

	deleteAzureAdRes, err := client.ApiCall("delete-azure-ad", azureAdPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteAzureAdRes.Success {
		if deleteAzureAdRes.ErrorMsg != "" {
			return fmt.Errorf(deleteAzureAdRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
