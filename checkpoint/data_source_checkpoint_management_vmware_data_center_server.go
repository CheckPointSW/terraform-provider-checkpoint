package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"strings"
)

func dataSourceManagementVMwareDataCenterServer() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVMwareDataCenterServerRead,
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
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object type. nsx, nsxt or vcenter or globalnsxt.",
			},
			"hostname": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IP Address or hostname of the vmware server.",
			},
			"username": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Username of the vmware server",
			},
			"certificate_fingerprint": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Specify the SHA-1 or SHA-256 fingerprint of the Data Center Server's certificate.",
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

func dataSourceVMwareDataCenterServerRead(d *schema.ResourceData, m interface{}) error {
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
	showVMwareDataCenterServerRes, err := client.ApiCall("show-data-center-server", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showVMwareDataCenterServerRes.Success {
		return fmt.Errorf(showVMwareDataCenterServerRes.ErrorMsg)
	}
	vmwareDataCenterServer := showVMwareDataCenterServerRes.GetData()

	if v := vmwareDataCenterServer["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := vmwareDataCenterServer["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if vmwareDataCenterServer["properties"] != nil {
		propsJson, ok := vmwareDataCenterServer["properties"].([]interface{})
		if ok {
			for _, prop := range propsJson {
				propMap := prop.(map[string]interface{})
				propName := strings.ReplaceAll(propMap["name"].(string), "-", "_")
				propValue := propMap["value"]
				if propName == "unsafe_auto_accept" || propName == "policy_mode" || propName == "import_vms" {
					propValue, _ = strconv.ParseBool(propValue.(string))
				}
				_ = d.Set(propName, propValue)
			}
		}
	}

	if vmwareDataCenterServer["tags"] != nil {
		tagsJson, ok := vmwareDataCenterServer["tags"].([]interface{})
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

	if v := vmwareDataCenterServer["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := vmwareDataCenterServer["data-center-type"]; v != nil {
		if v == "vcenter" || v == "nsx" || v == "nsxt" || v == "globalnsxt" {
			_ = d.Set("type", v)
		}
	}

	if v := vmwareDataCenterServer["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := vmwareDataCenterServer["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := vmwareDataCenterServer["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}
