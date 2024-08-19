package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementExternalTrustedCa() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceManagementExternalTrustedCaRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object unique identifier.",
			},
			"base64_certificate": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Certificate file encoded in base64.",
			},
			"retrieve_crl_from_http_servers": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether to retrieve Certificate Revocation List from http servers.",
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
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Comments string.",
			},
		},
	}
}

func dataSourceManagementExternalTrustedCaRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
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

	if v := externalTrustedCa["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

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

	if externalTrustedCa["domains_to_process"] != nil {
		domainsToProcessJson, ok := externalTrustedCa["domains_to_process"].([]interface{})
		if ok {
			domainsToProcessIds := make([]string, 0)
			if len(domainsToProcessJson) > 0 {
				for _, domains_to_process := range domainsToProcessJson {
					domains_to_process := domains_to_process.(map[string]interface{})
					domainsToProcessIds = append(domainsToProcessIds, domains_to_process["name"].(string))
				}
			}
			_ = d.Set("domains_to_process", domainsToProcessIds)
		}
	} else {
		_ = d.Set("domains_to_process", nil)
	}

	if v := externalTrustedCa["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := externalTrustedCa["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}
