package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"strings"
)

func dataSourceManagementAwsDataCenterServer() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAwsDataCenterServerRead,
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
				Description: "user-authentication\nUses the Access keys to authenticate.\nrole-authentication\nUses the AWS IAM role to authenticate.\nThis option requires the Security Management Server be deployed in AWS and has an IAM Role.",
			},
			"access_key_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Access key ID for the AWS account.\nRequired for authentication-method: user-authentication.",
			},
			"region": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Select the AWS region.",
			},
			"enable_sts_assume_role": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enables the STS Assume Role option. After it is enabled, the sts-role field is mandatory, whereas the sts-external-id is optional.",
			},
			"sts_role": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The STS RoleARN of the role to be assumed.\nRequired for enable-sts-assume-role: true.",
			},
			"sts_external_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "An optional STS External-Id to use when assuming the role.",
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

func dataSourceAwsDataCenterServerRead(d *schema.ResourceData, m interface{}) error {
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
	showAwsDataCenterServerRes, err := client.ApiCall("show-data-center-server", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showAwsDataCenterServerRes.Success {
		return fmt.Errorf(showAwsDataCenterServerRes.ErrorMsg)
	}
	awsDataCenterServer := showAwsDataCenterServerRes.GetData()

	if v := awsDataCenterServer["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := awsDataCenterServer["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if awsDataCenterServer["properties"] != nil {
		propsJson, ok := awsDataCenterServer["properties"].([]interface{})
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

	if awsDataCenterServer["tags"] != nil {
		tagsJson, ok := awsDataCenterServer["tags"].([]interface{})
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

	if v := awsDataCenterServer["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := awsDataCenterServer["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := awsDataCenterServer["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := awsDataCenterServer["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}
