package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementSyslogServer() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementSyslogServerRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"host": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Host server object identified by the name or UID.",
			},
			"port": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Port number.",
			},
			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "RFC version.",
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
			"icon": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object icon.",
			},
		},
	}
}

func dataSourceManagementSyslogServerRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showSyslogServerRes, err := client.ApiCallSimple("show-syslog-server", payload)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showSyslogServerRes.Success {
		return fmt.Errorf(showSyslogServerRes.ErrorMsg)
	}

	syslogServer := showSyslogServerRes.GetData()

	log.Println("Read SyslogServer - Show JSON = ", syslogServer)

	if v := syslogServer["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := syslogServer["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := syslogServer["host"]; v != nil {
		_ = d.Set("host", v.(map[string]interface{})["name"].(string))
	}

	if v := syslogServer["port"]; v != nil {
		_ = d.Set("port", v)
	}

	if v := syslogServer["version"]; v != nil {
		_ = d.Set("version", v)
	}

	if syslogServer["tags"] != nil {
		tagsJson, ok := syslogServer["tags"].([]interface{})
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

	if v := syslogServer["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := syslogServer["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := syslogServer["icon"]; v != nil {
		_ = d.Set("icon", v)
	}

	return nil

}
