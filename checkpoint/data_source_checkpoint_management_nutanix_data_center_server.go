package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceManagementNutanixDataCenterServer() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNutanixDataCenterServerRead,
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
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
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
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceNutanixDataCenterServerRead(d *schema.ResourceData, m interface{}) error {
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

	showNutanixDataCenterServerRes, err := client.ApiCall("show-data-center-server", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showNutanixDataCenterServerRes.Success {
		return fmt.Errorf(showNutanixDataCenterServerRes.ErrorMsg)
	}

	nutanixDataCenterServer := showNutanixDataCenterServerRes.GetData()

	if v := nutanixDataCenterServer["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := nutanixDataCenterServer["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if nutanixDataCenterServer["properties"] != nil {
		propertiesList := nutanixDataCenterServer["properties"].([]interface{})

		if len(propertiesList) > 0 {
			var propertiesListToReturn []map[string]interface{}

			for i := range propertiesList {
				propertiesMap := propertiesList[i].(map[string]interface{})

				propertiesMapToAdd := make(map[string]interface{})

				if v, _ := propertiesMap["name"]; v != nil {
					propertiesMapToAdd["name"] = v
				}
				if v, _ := propertiesMap["value"]; v != nil {
					propertiesMapToAdd["value"] = v
				}

				propertiesListToReturn = append(propertiesListToReturn, propertiesMapToAdd)
			}

			_ = d.Set("properties", propertiesListToReturn)
		} else {
			_ = d.Set("properties", propertiesList)
		}

	} else {
		_ = d.Set("properties", nil)
	}

	if nutanixDataCenterServer["tags"] != nil {
		tagsJson, ok := nutanixDataCenterServer["tags"].([]interface{})
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

	if v := nutanixDataCenterServer["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := nutanixDataCenterServer["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := nutanixDataCenterServer["automatic-refresh"]; v != nil {
		_ = d.Set("automatic_refresh", v)
	}

	if v := nutanixDataCenterServer["data-center-type"]; v != nil {
		_ = d.Set("data_center_type", v)
	}

	return nil
}
