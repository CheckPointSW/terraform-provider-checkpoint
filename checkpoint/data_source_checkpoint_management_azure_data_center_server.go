package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"strings"
)

func dataSourceManagementAzureDataCenterServer() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAzureDataCenterServerRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name. Must be unique in the domain.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object unique identifier.",
			},
			"authentication_method": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "user-authentication\nUses the Azure AD User to authenticate.\nservice-principal-authentication\nUses the Service Principal to authenticate.",
			},
			"username": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "An Azure Active Directory user Format <username>@<domain>.\nRequired for authentication-method: user-authentication.",
			},
			"application_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Application ID of the Service Principal, in UUID format.\nRequired for authentication-method: service-principal-authentication.",
			},
			"directory_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Directory ID of the Azure AD, in UUID format.\nRequired for authentication-method: service-principal-authentication.",
			},
			"environment": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Select the Azure Cloud Environment.",
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
			},
		},
	}
}

func dataSourceAzureDataCenterServerRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	var name string
	var uid string

	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}
	if v, ok := d.GetOk("uid"); ok {
		uid = v.(string)
	}
	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}
	showAzureDataCenterServerRes, err := client.ApiCall("show-data-center-server", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showAzureDataCenterServerRes.Success {
		return fmt.Errorf(showAzureDataCenterServerRes.ErrorMsg)
	}
	azureDataCenterServer := showAzureDataCenterServerRes.GetData()

	if v := azureDataCenterServer["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

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
				if propName == "enable_sts_assume_role" {
					propValue, _ = strconv.ParseBool(propValue.(string))
				}
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
