package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementIdpDefaultAssignment() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementIdpDefaultAssignmentRead,
		Schema: map[string]*schema.Schema{
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"identity_provider": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Represents the Identity Provider to be used for Login by this assignment identified by the name or UID, to cancel existing assignment should set to 'none'.",
			},
			"identity_provider_set": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "True if 'identity-provider' value is set.",
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

func dataSourceManagementIdpDefaultAssignmentRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := make(map[string]interface{})

	showIdpDefaultAssignmentRes, err := client.ApiCall("show-idp-default-assignment", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showIdpDefaultAssignmentRes.Success {
		if objectNotFound(showIdpDefaultAssignmentRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showIdpDefaultAssignmentRes.ErrorMsg)
	}

	idpDefaultAssignment := showIdpDefaultAssignmentRes.GetData()

	log.Println("Read IdpDefaultAssignment - Show JSON = ", idpDefaultAssignment)

	if v := idpDefaultAssignment["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := idpDefaultAssignment["identity-provider"]; v != nil {
		_ = d.Set("identity_provider", v)
	}

	if v := idpDefaultAssignment["identity-provider-set"]; v != nil {
		_ = d.Set("identity_provider_set", v)
	}

	if idpDefaultAssignment["tags"] != nil {
		tagsJson, ok := idpDefaultAssignment["tags"].([]interface{})
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

	return nil

}
