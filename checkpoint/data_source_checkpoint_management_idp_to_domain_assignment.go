package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementIdpToDomainAssignment() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementIdpToDomainAssignmentRead,
		Schema: map[string]*schema.Schema{
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object unique identifier.",
			},
			"assigned_domain": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Represents the Domain assigned by 'idp-to-domain-assignment', need to be domain name or UID.",
			},
			"identity_provider": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Represents the Identity Provider to be used for Login by this assignment. Must be set when \"using-default\" was set to be false.",
			},
			"identity_provider_set": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "True if 'identity-provider' value is set.",
			},
			"using_default": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Is this assignment override by 'idp-default-assignment'.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceManagementIdpToDomainAssignmentRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("assigned_domain").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["assigned_domain"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showIdpToDomainAssignmentRes, err := client.ApiCall("show-idp-to-domain-assignment", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showIdpToDomainAssignmentRes.Success {
		if objectNotFound(showIdpToDomainAssignmentRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showIdpToDomainAssignmentRes.ErrorMsg)
	}

	idpToDomainAssignment := showIdpToDomainAssignmentRes.GetData()

	log.Println("Read IdpToDomainAssignment - Show JSON = ", idpToDomainAssignment)

	if v := idpToDomainAssignment["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := idpToDomainAssignment["assigned-domain"]; v != nil {
		_ = d.Set("assigned_domain", v.(map[string]interface{})[name])
	}

	if v := idpToDomainAssignment["identity-provider"]; v != nil {
		_ = d.Set("identity_provider", v)
	}

	if v := idpToDomainAssignment["identity-provider-set"]; v != nil {
		_ = d.Set("identity_provider_set", v)
	}

	if idpToDomainAssignment["tags"] != nil {
		tagsJson, ok := idpToDomainAssignment["tags"].([]interface{})
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

	if v := idpToDomainAssignment["using-default"]; v != nil {
		_ = d.Set("using_default", v)
	}
	return nil

}
