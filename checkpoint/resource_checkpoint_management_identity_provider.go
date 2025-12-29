package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementIdentityProvider() *schema.Resource {
	return &schema.Resource{
		Create: createManagementIdentityProvider,
		Read:   readManagementIdentityProvider,
		Update: updateManagementIdentityProvider,
		Delete: deleteManagementIdentityProvider,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"usage": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Usage of Identity Provider.",
				Default:     "gateway_policy_and_logs",
			},
			"gateway": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Gateway for the SAML Identity Provider usage. Identified by name or UID. <font color=\"red\">Required only when</font> 'usage' is set to 'gateway_policy_and_logs'.",
			},
			"service": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Service for the selected gateway. <font color=\"red\">Required only when</font> 'usage' is set to 'gateway_policy_and_logs'.",
			},
			"data_receiving": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Data receiving method from the SAML Identity Provider.",
				Default:     "manually",
			},
			"received_identifier": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Received Identifier (Entity ID) based on the provider data. <font color=\"red\">Required only when</font> 'data-receiving' is set to 'manually'.",
			},
			"login_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Login URL based on the provider data. <font color=\"red\">Required only when</font> 'data-receiving' is set to 'manually'.",
			},
			"base64_metadata_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Metadata file encoded in base64 based on the provider data. <font color=\"red\">Required only when</font> 'data-receiving' is set to 'metadata_file'.",
			},
			"base64_certificate": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Certificate file encoded in base64 based on provider data. <font color=\"red\">Required only when</font> 'data-receiving' is set to 'manually'.",
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
			},
			"ignore_errors": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.",
			},
		},
	}
}

func createManagementIdentityProvider(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	identityProvider := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		identityProvider["name"] = v.(string)
	}

	if v, ok := d.GetOk("usage"); ok {
		identityProvider["usage"] = v.(string)
	}

	if v, ok := d.GetOk("gateway"); ok {
		identityProvider["gateway"] = v.(string)
	}

	if v, ok := d.GetOk("service"); ok {
		identityProvider["service"] = v.(string)
	}

	if v, ok := d.GetOk("data_receiving"); ok {
		identityProvider["data-receiving"] = v.(string)
	}

	if v, ok := d.GetOk("received_identifier"); ok {
		identityProvider["received-identifier"] = v.(string)
	}

	if v, ok := d.GetOk("login_url"); ok {
		identityProvider["login-url"] = v.(string)
	}

	if v, ok := d.GetOk("base64_metadata_file"); ok {
		identityProvider["base64-metadata-file"] = v.(string)
	}

	if v, ok := d.GetOk("base64_certificate"); ok {
		identityProvider["base64-certificate"] = v.(string)
	}

	if v, ok := d.GetOk("color"); ok {
		identityProvider["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		identityProvider["comments"] = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		identityProvider["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		identityProvider["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		identityProvider["ignore-errors"] = v.(bool)
	}

	log.Println("Create IdentityProvider - Map = ", identityProvider)

	addIdentityProviderRes, err := client.ApiCallSimple("add-identity-provider", identityProvider)
	if err != nil || !addIdentityProviderRes.Success {
		if addIdentityProviderRes.ErrorMsg != "" {
			return fmt.Errorf(addIdentityProviderRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addIdentityProviderRes.GetData()["uid"].(string))

	return readManagementIdentityProvider(d, m)
}

func readManagementIdentityProvider(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showIdentityProviderRes, err := client.ApiCallSimple("show-identity-provider", payload)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showIdentityProviderRes.Success {
		if objectNotFound(showIdentityProviderRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showIdentityProviderRes.ErrorMsg)
	}

	identityProvider := showIdentityProviderRes.GetData()

	log.Println("Read IdentityProvider - Show JSON = ", identityProvider)

	if v := identityProvider["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := identityProvider["usage"]; v != nil {
		_ = d.Set("usage", v)
	}

	if v := identityProvider["gateway"]; v != nil {
		_ = d.Set("gateway", v)
	}

	if d.Get("service").(string) != "" {
		if v := identityProvider["service"]; v != nil {
			_ = d.Set("service", v)
		}
	}

	if v := identityProvider["data-receiving"]; v != nil {
		_ = d.Set("data_receiving", v)
	}

	if v := identityProvider["received-identifier"]; v != nil {
		_ = d.Set("received_identifier", v)
	}

	if v := identityProvider["login-url"]; v != nil {
		_ = d.Set("login_url", v)
	}

	if v := identityProvider["base64-metadata-file"]; v != nil {
		_ = d.Set("base64_metadata_file", v)
	}

	if v := identityProvider["base64-certificate"]; v != nil {
		_ = d.Set("base64_certificate", v)
	}

	if v := identityProvider["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := identityProvider["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if identityProvider["tags"] != nil {
		tagsJson, ok := identityProvider["tags"].([]interface{})
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

	if v := identityProvider["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := identityProvider["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementIdentityProvider(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	identityProvider := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		identityProvider["name"] = oldName
		identityProvider["new-name"] = newName
	} else {
		identityProvider["name"] = d.Get("name")
	}

	if ok := d.HasChange("usage"); ok {
		identityProvider["usage"] = d.Get("usage")
	}

	if ok := d.HasChange("gateway"); ok {
		identityProvider["gateway"] = d.Get("gateway")
	}

	if ok := d.HasChange("service"); ok {
		identityProvider["service"] = d.Get("service")
	}

	if ok := d.HasChange("data_receiving"); ok {
		identityProvider["data-receiving"] = d.Get("data_receiving")
	}

	if ok := d.HasChange("received_identifier"); ok {
		identityProvider["received-identifier"] = d.Get("received_identifier")
	}

	if ok := d.HasChange("login_url"); ok {
		identityProvider["login-url"] = d.Get("login_url")
	}

	if ok := d.HasChange("base64_metadata_file"); ok {
		identityProvider["base64-metadata-file"] = d.Get("base64_metadata_file")
	}

	if ok := d.HasChange("base64_certificate"); ok {
		identityProvider["base64-certificate"] = d.Get("base64_certificate")
	}

	if ok := d.HasChange("color"); ok {
		identityProvider["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		identityProvider["comments"] = d.Get("comments")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			identityProvider["tags"] = v.(*schema.Set).List()
		}
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		identityProvider["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		identityProvider["ignore-errors"] = v.(bool)
	}

	log.Println("Update IdentityProvider - Map = ", identityProvider)

	updateIdentityProviderRes, err := client.ApiCallSimple("set-identity-provider", identityProvider)
	if err != nil || !updateIdentityProviderRes.Success {
		if updateIdentityProviderRes.ErrorMsg != "" {
			return fmt.Errorf(updateIdentityProviderRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementIdentityProvider(d, m)
}

func deleteManagementIdentityProvider(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	identityProviderPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete IdentityProvider")

	deleteIdentityProviderRes, err := client.ApiCallSimple("delete-identity-provider", identityProviderPayload)
	if err != nil || !deleteIdentityProviderRes.Success {
		if deleteIdentityProviderRes.ErrorMsg != "" {
			return fmt.Errorf(deleteIdentityProviderRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
