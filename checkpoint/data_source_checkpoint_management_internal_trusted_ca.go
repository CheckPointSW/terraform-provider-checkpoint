package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceManagementSetInternalTrustedCa() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceManagementSetInternalTrustedCaRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object unique identifier.",
			},
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object unique identifier.",
			},
			"base64_certificate": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object unique identifier.",
			},
			"retrieve_crl_from_http_servers": {
				Type:     schema.TypeBool,
				Computed: true,

				Description: "Whether to retrieve Certificate Revocation List from http servers.",
			},
			"retrieve_crl_from_ldap_servers": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether to retrieve Certificate Revocation List from ldap servers.",
			},
			"cache_crl": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Cache Certificate Revocation List on the Security Gateway.",
			},
			"crl_cache_method": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Weather to retrieve new Certificate Revocation List after the certificate expires or after a fixed period.",
			},
			"crl_cache_timeout": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "When to fetch new Certificate Revocation List (in minutes).",
			},
			"allow_certificates_from_branches": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Allow only certificates from listed branches.",
			},
			"branches": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Branches to allow certificates from. Required only if \"allow-certificates-from-branches\" set to \"true\".",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"tags": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"color": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Color of the object. Should be one of existing colors.",
			},
			"comments": {
				Type:     schema.TypeString,
				Computed: true,

				Description: "Comments string.",
			},
		},
	}
}

func dataSourceManagementSetInternalTrustedCaRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}

	internalTrustedCaRes, _ := client.ApiCall("show-internal-trusted-ca", payload, client.GetSessionID(), true, false)
	if !internalTrustedCaRes.Success {
		return fmt.Errorf(internalTrustedCaRes.ErrorMsg)
	}
	internalTrustedCaData := internalTrustedCaRes.GetData()

	if v := internalTrustedCaData["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := internalTrustedCaData["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := internalTrustedCaData["base64-certificate"]; v != nil {
		_ = d.Set("base64_certificate", v)
	}

	if v := internalTrustedCaData["retrieve-crl-from-http-servers"]; v != nil {
		_ = d.Set("retrieve_crl_from_http_servers", v)
	}

	if v := internalTrustedCaData["retrieve-crl-from-ldap-servers"]; v != nil {
		_ = d.Set("retrieve_crl_from_ldap_servers", v)
	}

	if v := internalTrustedCaData["cache-crl"]; v != nil {
		_ = d.Set("cache_crl", v)
	}

	if v := internalTrustedCaData["crl-cache-method"]; v != nil {
		_ = d.Set("crl_cache_method", v)
	}

	if v := internalTrustedCaData["crl-cache-timeout"]; v != nil {
		_ = d.Set("crl_cache_timeout", v)
	}

	if v := internalTrustedCaData["allow-certificates-from-branches"]; v != nil {
		_ = d.Set("allow_certificates_from_branches", v)
	}

	if v := internalTrustedCaData["branches"]; v != nil {
		_ = d.Set("branches", v)
	}

	if internalTrustedCaData["tags"] != nil {
		tagsJson, ok := internalTrustedCaData["tags"].([]interface{})
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

	if v := internalTrustedCaData["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := internalTrustedCaData["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil

}
