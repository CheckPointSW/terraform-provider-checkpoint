package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementLogicalServer() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementLogicalServerRead,
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
			"ipv4_address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IPv4 address.",
			},
			"ipv6_address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IPv6 address.",
			},
			"server_group": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Server group associated with the logical server.  Identified by name or UID.",
			},
			"server_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of server for the logical server.",
			},
			"persistence_mode": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates if persistence mode is enabled for the logical server.",
			},
			"persistency_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Persistency type for the logical server.",
			},
			"balance_method": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Load balancing method for the logical server.",
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
			"tags": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"icon": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object icon.",
			},
		},
	}
}

func dataSourceManagementLogicalServerRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{}

	if v, ok := d.GetOk("name"); ok {
		payload["name"] = v.(string)
	} else if v, ok := d.GetOk("uid"); ok {
		payload["uid"] = v.(string)
	} else {
		return fmt.Errorf("Either name or uid must be specified")
	}

	showLogicalServerRes, err := client.ApiCall("show-logical-server", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showLogicalServerRes.Success {
		if objectNotFound(showLogicalServerRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showLogicalServerRes.ErrorMsg)
	}

	logicalServer := showLogicalServerRes.GetData()

	log.Println("Read LogicalServer - Show JSON = ", logicalServer)

	if v := logicalServer["uid"]; v != nil {
		d.SetId(v.(string))
	}

	if v := logicalServer["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := logicalServer["ipv4-address"]; v != nil {
		_ = d.Set("ipv4_address", v)
	}

	if v := logicalServer["ipv6-address"]; v != nil {
		_ = d.Set("ipv6_address", v)
	}

	if v := logicalServer["server-group"]; v != nil {
		_ = d.Set("server_group", v.(map[string]interface{})["name"].(string))
	}

	if v := logicalServer["server-type"]; v != nil {
		_ = d.Set("server_type", v)
	}

	if v := logicalServer["persistence-mode"]; v != nil {
		_ = d.Set("persistence_mode", v)
	}

	if v := logicalServer["persistency-type"]; v != nil {
		_ = d.Set("persistency_type", v)
	}

	if v := logicalServer["balance-method"]; v != nil {
		_ = d.Set("balance_method", v)
	}

	if v := logicalServer["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := logicalServer["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := logicalServer["icon"]; v != nil {
		_ = d.Set("icon", v)
	}

	if logicalServer["tags"] != nil {
		tagsJson, ok := logicalServer["tags"].([]interface{})
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
