package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementGsnHandoverGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementGsnHandoverGroupRead,
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
			"enforce_gtp": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable enforce GTP signal packet rate limit from this group.",
			},
			"gtp_rate": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Limit of the GTP rate in PDU/sec.",
			},
			"members": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of GSN handover group members identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
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
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Comments string.",
			},
		},
	}
}

func dataSourceManagementGsnHandoverGroupRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showGsnHandoverGroupRes, err := client.ApiCall("show-gsn-handover-group", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showGsnHandoverGroupRes.Success {
		return fmt.Errorf(showGsnHandoverGroupRes.ErrorMsg)
	}

	gsnHandoverGroup := showGsnHandoverGroupRes.GetData()

	log.Println("Read GsnHandoverGroup - Show JSON = ", gsnHandoverGroup)

	if v := gsnHandoverGroup["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := gsnHandoverGroup["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := gsnHandoverGroup["enforce-gtp"]; v != nil {
		_ = d.Set("enforce_gtp", v)
	}

	if v := gsnHandoverGroup["gtp-rate"]; v != nil {
		_ = d.Set("gtp_rate", v)
	}

	if gsnHandoverGroup["members"] != nil {
		membersJson, ok := gsnHandoverGroup["members"].([]interface{})
		if ok {
			membersIds := make([]string, 0)
			if len(membersJson) > 0 {
				for _, members := range membersJson {
					members := members.(map[string]interface{})
					membersIds = append(membersIds, members["name"].(string))
				}
			}
			_ = d.Set("members", membersIds)
		}
	} else {
		_ = d.Set("members", nil)
	}

	if gsnHandoverGroup["tags"] != nil {
		tagsJson, ok := gsnHandoverGroup["tags"].([]interface{})
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

	if v := gsnHandoverGroup["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := gsnHandoverGroup["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil
}
