package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"strings"
)

func dataSourceManagementKubernetesDataCenterServer() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceKubernetesDataCenterServerRead,
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
			"hostname": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IP address or hostname of the Kubernetes server.",
			},
			"token_file": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Kubernetes access token encoded in base64.",
			},
			"ca_certificate": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Kubernetes public certificate key encoded in base64.",
			},
			"unsafe_auto_accept": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "When set to false, the current Data Center Server's certificate should be trusted, either by providing the certificate-fingerprint argument or by relying on a previously trusted certificate of this hostname.\n\nWhen set to true, trust the current Data Center Server's certificate as-is.",
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

func dataSourceKubernetesDataCenterServerRead(d *schema.ResourceData, m interface{}) error {
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
	showKubernetesDataCenterServerRes, err := client.ApiCall("show-data-center-server", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showKubernetesDataCenterServerRes.Success {
		return fmt.Errorf(showKubernetesDataCenterServerRes.ErrorMsg)
	}
	kubernetesDataCenterServer := showKubernetesDataCenterServerRes.GetData()

	if v := kubernetesDataCenterServer["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := kubernetesDataCenterServer["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if kubernetesDataCenterServer["properties"] != nil {
		propsJson, ok := kubernetesDataCenterServer["properties"].([]interface{})
		if ok {
			for _, prop := range propsJson {
				propMap := prop.(map[string]interface{})
				if propMap["name"] != nil {
					propName := strings.ReplaceAll(propMap["name"].(string), "-", "_")
					propValue := propMap["value"]
					if propName == "unsafe_auto_accept" {
						propValue, _ = strconv.ParseBool(propValue.(string))
					}
					_ = d.Set(propName, propValue)
				}
			}
		}
	}

	if kubernetesDataCenterServer["tags"] != nil {
		tagsJson, ok := kubernetesDataCenterServer["tags"].([]interface{})
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

	if v := kubernetesDataCenterServer["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := kubernetesDataCenterServer["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := kubernetesDataCenterServer["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := kubernetesDataCenterServer["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}
