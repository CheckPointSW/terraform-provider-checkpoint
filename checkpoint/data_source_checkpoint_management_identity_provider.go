package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementIdentityProvider() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceManagementIdentityProviderRead,
		Schema: map[string]*schema.Schema{
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object uid.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"usage": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Usage of Identity Provider.",
			},
			"gateway": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Gateway for the SAML Identity Provider usage. Identified by name or UID. <font color=\"red\">Required only when</font> 'usage' is set to 'gateway_policy_and_logs'.",
			},
			"service": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Service for the selected gateway. <font color=\"red\">Required only when</font> 'usage' is set to 'gateway_policy_and_logs'.",
			},
			"required_identifier": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Required identifier (Entity ID) for the SAML Identity Provider.",
			},
			"reply_urls": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "List of URLs for the SAML Identity Provider.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"data_receiving": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Data receiving method from the SAML Identity Provider.",
			},
			"received_identifier": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Received Identifier (Entity ID) based on the provider data. <font color=\"red\">Required only when</font> 'data-receiving' is set to 'manually'.",
			},
			"login_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Login URL based on the provider data. <font color=\"red\">Required only when</font> 'data-receiving' is set to 'manually'.",
			},
			"base64_metadata_file": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Metadata file encoded in base64 based on the provider data. <font color=\"red\">Required only when</font> 'data-receiving' is set to 'metadata_file'.",
			},
			"base64_certificate": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Certificate file encoded in base64 based on provider data. <font color=\"red\">Required only when</font> 'data-receiving' is set to 'manually'.",
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

func dataSourceManagementIdentityProviderRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}
	showIdentityProviderRes, err := client.ApiCallSimple("show-identity-provider", payload)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showIdentityProviderRes.Success {
		return fmt.Errorf(showIdentityProviderRes.ErrorMsg)
	}

	identityProvider := showIdentityProviderRes.GetData()

	log.Println("Read IdentityProvider - Show JSON = ", identityProvider)

	if v := identityProvider["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := identityProvider["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := identityProvider["usage"]; v != nil {
		_ = d.Set("usage", v)
	}

	if v := identityProvider["gateway"]; v != nil {
		_ = d.Set("gateway", v.(map[string]interface{})["name"].(string))
	}

	if v := identityProvider["service"]; v != nil {
		_ = d.Set("service", v)
	}

	if v := identityProvider["data-receiving"]; v != nil {
		_ = d.Set("data_receiving", v)
	}

	if v := identityProvider["received-identifier"]; v != nil {
		_ = d.Set("received_identifier", v)
	}

	if v := identityProvider["required-identifier"]; v != nil {
		_ = d.Set("required_identifier", v)
	}

	if identityProvider["reply-urls"] != nil {
		replyUrlsJson, ok := identityProvider["reply-urls"].([]interface{})
		if ok {
			replyUrls := make([]string, 0)
			if len(replyUrlsJson) > 0 {
				for _, reply_url := range replyUrlsJson {
					replyUrls = append(replyUrls, reply_url.(string))
				}
			}
			_ = d.Set("reply_urls", replyUrls)
		}
	} else {
		_ = d.Set("reply_urls", nil)
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

	if v := identityProvider["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := identityProvider["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil

}
