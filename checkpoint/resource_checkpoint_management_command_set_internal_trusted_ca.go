package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementSetInternalTrustedCa() *schema.Resource {
	return &schema.Resource{
		Create: createManagementSetInternalTrustedCa,
		Read:   readManagementSetInternalTrustedCa,
		Delete: deleteManagementSetInternalTrustedCa,
		Schema: map[string]*schema.Schema{
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object unique identifier.",
			},
			"retrieve_crl_from_http_servers": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Whether to retrieve Certificate Revocation List from http servers.",
			},
			"cache_crl": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Cache Certificate Revocation List on the Security Gateway.",
			},
			"crl_cache_method": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Weather to retrieve new Certificate Revocation List after the certificate expires or after a fixed period.",
			},
			"crl_cache_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "When to fetch new Certificate Revocation List (in minutes).",
			},
			"allow_certificates_from_branches": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Allow only certificates from listed branches.",
			},
			"branches": {
				Type:        schema.TypeSet,
				Optional:    true,
				ForceNew:    true,
				Description: "Branches to allow certificates from. Required only if \"allow-certificates-from-branches\" set to \"true\".",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				ForceNew:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"color": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Color of the object. Should be one of existing colors.",
			},
			"comments": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Comments string.",
			},
			"domains_to_process": {
				Type:        schema.TypeSet,
				Optional:    true,
				ForceNew:    true,
				Description: "Indicates which domains to process the commands on. It cannot be used with the details-level full, must be run from the System Domain only and with ignore-warnings true. Valid values are: CURRENT_DOMAIN, ALL_DOMAINS_ON_THIS_SERVER.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ignore_warnings": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Apply changes ignoring warnings.",
			},
			"ignore_errors": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.",
			},
		},
	}
}

func createManagementSetInternalTrustedCa(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOkExists("retrieve_crl_from_http_servers"); ok {
		payload["retrieve-crl-from-http-servers"] = v.(bool)
	}

	if v, ok := d.GetOkExists("cache_crl"); ok {
		payload["cache-crl"] = v.(bool)
	}

	if v, ok := d.GetOk("crl_cache_method"); ok {
		payload["crl-cache-method"] = v.(string)
	}

	if v, ok := d.GetOk("crl_cache_timeout"); ok {
		payload["crl-cache-timeout"] = v.(int)
	}

	if v, ok := d.GetOkExists("allow_certificates_from_branches"); ok {
		payload["allow-certificates-from-branches"] = v.(bool)
	}

	if v, ok := d.GetOk("branches"); ok {
		payload["branches"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("tags"); ok {
		payload["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		payload["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		payload["comments"] = v.(string)
	}

	if v, ok := d.GetOk("domains_to_process"); ok {
		payload["domains-to-process"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		payload["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		payload["ignore-errors"] = v.(bool)
	}

	SetInternalTrustedCaRes, _ := client.ApiCall("set-internal-trusted-ca", payload, client.GetSessionID(), true, false)
	if !SetInternalTrustedCaRes.Success {
		return fmt.Errorf(SetInternalTrustedCaRes.ErrorMsg)
	}

	res := SetInternalTrustedCaRes.GetData()

	_ = d.Set("uid", res["uid"])
	d.SetId(res["uid"].(string))

	return readManagementSetInternalTrustedCa(d, m)
}

func readManagementSetInternalTrustedCa(d *schema.ResourceData, m interface{}) error {

	return nil
}

func deleteManagementSetInternalTrustedCa(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
