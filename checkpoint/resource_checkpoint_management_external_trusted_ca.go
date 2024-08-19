package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementExternalTrustedCa() *schema.Resource {
	return &schema.Resource{
		Create: createManagementExternalTrustedCa,
		Read:   readManagementExternalTrustedCa,
		Update: updateManagementExternalTrustedCa,
		Delete: deleteManagementExternalTrustedCa,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"base64_certificate": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Certificate file encoded in base64.",
			},
			"retrieve_crl_from_http_servers": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether to retrieve Certificate Revocation List from http servers.",
				Default:     true,
			},
			"crl_cache_method": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Weather to retrieve new Certificate Revocation List after the certificate expires or after a fixed period.",
				Default:     "timeout",
			},
			"crl_cache_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "When to fetch new Certificate Revocation List (in minutes).",
				Default:     1440,
			},
			"allow_certificates_from_branches": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Allow only certificates from listed branches.",
				Default:     false,
			},
			"branches": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Branches to allow certificates from. Required only if \"allow-certificates-from-branches\" set to \"true\".",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
		},
	}
}

func createManagementExternalTrustedCa(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	externalTrustedCa := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		externalTrustedCa["name"] = v.(string)
	}

	if v, ok := d.GetOk("base64_certificate"); ok {
		externalTrustedCa["base64-certificate"] = v.(string)
	}

	if v, ok := d.GetOkExists("retrieve_crl_from_http_servers"); ok {
		externalTrustedCa["retrieve-crl-from-http-servers"] = v.(bool)
	}

	if v, ok := d.GetOk("crl_cache_method"); ok {
		externalTrustedCa["crl-cache-method"] = v.(string)
	}

	if v, ok := d.GetOk("crl_cache_timeout"); ok {
		externalTrustedCa["crl-cache-timeout"] = v.(int)
	}

	if v, ok := d.GetOkExists("allow_certificates_from_branches"); ok {
		externalTrustedCa["allow-certificates-from-branches"] = v.(bool)
	}

	if v, ok := d.GetOk("branches"); ok {
		externalTrustedCa["branches"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("tags"); ok {
		externalTrustedCa["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		externalTrustedCa["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		externalTrustedCa["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		externalTrustedCa["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		externalTrustedCa["ignore-errors"] = v.(bool)
	}

	log.Println("Create ExternalTrustedCa - Map = ", externalTrustedCa)

	addExternalTrustedCaRes, err := client.ApiCall("add-external-trusted-ca", externalTrustedCa, client.GetSessionID(), true, false)
	if err != nil || !addExternalTrustedCaRes.Success {
		if addExternalTrustedCaRes.ErrorMsg != "" {
			return fmt.Errorf(addExternalTrustedCaRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addExternalTrustedCaRes.GetData()["uid"].(string))

	return readManagementExternalTrustedCa(d, m)
}

func readManagementExternalTrustedCa(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showExternalTrustedCaRes, err := client.ApiCall("show-external-trusted-ca", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showExternalTrustedCaRes.Success {
		if objectNotFound(showExternalTrustedCaRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showExternalTrustedCaRes.ErrorMsg)
	}

	externalTrustedCa := showExternalTrustedCaRes.GetData()

	log.Println("Read ExternalTrustedCa - Show JSON = ", externalTrustedCa)

	if v := externalTrustedCa["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := externalTrustedCa["base64-certificate"]; v != nil {
		_ = d.Set("base64_certificate", cleanseCertificate(v.(string)))
	}

	if v := externalTrustedCa["retrieve-crl-from-http-servers"]; v != nil {
		_ = d.Set("retrieve_crl_from_http_servers", v)
	}

	if v := externalTrustedCa["crl-cache-method"]; v != nil {
		_ = d.Set("crl_cache_method", v)
	}

	if v := externalTrustedCa["crl-cache-timeout"]; v != nil {
		_ = d.Set("crl_cache_timeout", v)
	}

	if v := externalTrustedCa["allow-certificates-from-branches"]; v != nil {
		_ = d.Set("allow_certificates_from_branches", v)
	}

	if externalTrustedCa["branches"] != nil {
		branchesJson, ok := externalTrustedCa["branches"].([]interface{})
		if ok {
			branchesIds := make([]string, 0)
			if len(branchesJson) > 0 {
				for _, branches := range branchesJson {
					branches := branches.(map[string]interface{})
					branchesIds = append(branchesIds, branches["name"].(string))
				}
			}
			_ = d.Set("branches", branchesIds)
		}
	} else {
		_ = d.Set("branches", nil)
	}

	if externalTrustedCa["tags"] != nil {
		tagsJson, ok := externalTrustedCa["tags"].([]interface{})
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

	if v := externalTrustedCa["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := externalTrustedCa["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := externalTrustedCa["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := externalTrustedCa["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementExternalTrustedCa(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	externalTrustedCa := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		externalTrustedCa["name"] = oldName
		externalTrustedCa["new-name"] = newName
	} else {
		externalTrustedCa["name"] = d.Get("name")
	}

	if ok := d.HasChange("base64_certificate"); ok {
		externalTrustedCa["base64-certificate"] = d.Get("base64_certificate")
	}

	if v, ok := d.GetOkExists("retrieve_crl_from_http_servers"); ok {
		externalTrustedCa["retrieve-crl-from-http-servers"] = v.(bool)
	}

	if ok := d.HasChange("crl_cache_method"); ok {
		externalTrustedCa["crl-cache-method"] = d.Get("crl_cache_method")
	}

	if ok := d.HasChange("crl_cache_timeout"); ok {
		externalTrustedCa["crl-cache-timeout"] = d.Get("crl_cache_timeout")
	}

	if v, ok := d.GetOkExists("allow_certificates_from_branches"); ok {
		externalTrustedCa["allow-certificates-from-branches"] = v.(bool)
	}

	if d.HasChange("branches") {
		if v, ok := d.GetOk("branches"); ok {
			externalTrustedCa["branches"] = v.(*schema.Set).List()
		} else {
			oldBranches, _ := d.GetChange("branches")
			externalTrustedCa["branches"] = map[string]interface{}{"remove": oldBranches.(*schema.Set).List()}
		}
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			externalTrustedCa["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			externalTrustedCa["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		externalTrustedCa["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		externalTrustedCa["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		externalTrustedCa["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		externalTrustedCa["ignore-errors"] = v.(bool)
	}

	log.Println("Update ExternalTrustedCa - Map = ", externalTrustedCa)

	updateExternalTrustedCaRes, err := client.ApiCall("set-external-trusted-ca", externalTrustedCa, client.GetSessionID(), true, false)
	if err != nil || !updateExternalTrustedCaRes.Success {
		if updateExternalTrustedCaRes.ErrorMsg != "" {
			return fmt.Errorf(updateExternalTrustedCaRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementExternalTrustedCa(d, m)
}

func deleteManagementExternalTrustedCa(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	externalTrustedCaPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete ExternalTrustedCa")

	deleteExternalTrustedCaRes, err := client.ApiCall("delete-external-trusted-ca", externalTrustedCaPayload, client.GetSessionID(), true, false)
	if err != nil || !deleteExternalTrustedCaRes.Success {
		if deleteExternalTrustedCaRes.ErrorMsg != "" {
			return fmt.Errorf(deleteExternalTrustedCaRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
