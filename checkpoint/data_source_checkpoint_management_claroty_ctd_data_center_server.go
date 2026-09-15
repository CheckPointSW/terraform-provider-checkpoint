package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	"strings"
)

func dataSourceManagementClarotyCtdDataCenterServer() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceClarotyCtdDataCenterServerRead,
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
				Description: "IP address or hostname of the Claroty CTD server.",
			},
			"username": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "User name for Claroty CTD.",
			},
			"certificate_fingerprint": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Specify the SHA-1 or SHA-256 fingerprint of the Data Center Server's certificate.",
			},
			"unsafe_auto_accept": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "When set to false, the current Data Center Server's certificate should be trusted, either by providing the certificate-fingerprint argument or by relying on a previously trusted certificate of this hostname. When set to true, trust the current Data Center Server's certificate as-is.",
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
			"automatic_refresh": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether the data center server's content is automatically updated.",
			},
			"data_center_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Data Center type.",
			},
			"properties": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Data Center properties.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "N/A",
						},
						"value": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "N/A",
						},
					},
				},
			},
		},
	}
}

func dataSourceClarotyCtdDataCenterServerRead(d *schema.ResourceData, m interface{}) error {
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
	showRes, err := client.ApiCall("show-data-center-server", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	if !showRes.Success {
		return fmt.Errorf("%s", showRes.ErrorMsg)
	}
	clarotyCtdDataCenterServer := showRes.GetData()

	if v := clarotyCtdDataCenterServer["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := clarotyCtdDataCenterServer["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if clarotyCtdDataCenterServer["properties"] != nil {
		propsJson, ok := clarotyCtdDataCenterServer["properties"].([]interface{})
		if ok {
			for _, prop := range propsJson {
				propMap := prop.(map[string]interface{})
				propName := strings.ReplaceAll(propMap["name"].(string), "-", "_")
				propValue := propMap["value"]
				if propName == "unsafe_auto_accept" {
					propValue, _ = strconv.ParseBool(propValue.(string))
				}
				_ = d.Set(propName, propValue)
			}
		}
	}

	if clarotyCtdDataCenterServer["tags"] != nil {
		tagsJson, ok := clarotyCtdDataCenterServer["tags"].([]interface{})
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

	if v := clarotyCtdDataCenterServer["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := clarotyCtdDataCenterServer["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := clarotyCtdDataCenterServer["automatic-refresh"]; v != nil {
		_ = d.Set("automatic_refresh", v)
	}

	if v := clarotyCtdDataCenterServer["data-center-type"]; v != nil {
		_ = d.Set("data_center_type", v)
	}

	if v := clarotyCtdDataCenterServer["properties"]; v != nil {
		propertiesList := v.([]interface{})
		var propertiesListState []map[string]interface{}
		for i := range propertiesList {
			propertiesShow := propertiesList[i].(map[string]interface{})
			propertiesState := make(map[string]interface{})
			if v := propertiesShow["name"]; v != nil {
				propertiesState["name"] = v
			}
			if v := propertiesShow["value"]; v != nil {
				propertiesState["value"] = v
			}
			propertiesListState = append(propertiesListState, propertiesState)
		}
		_ = d.Set("properties", propertiesListState)
	}

	return nil

}
