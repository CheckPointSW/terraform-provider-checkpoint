package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"

	"strconv"
)

func dataSourceManagementOpsecApplication() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementOpsecApplicationRead,
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
			"cpmi": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Used to setup the CPMI client entity.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"administrator_profile": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A profile to set the log reading permissions by for the client entity.",
						},
						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to enable this client entity on the Opsec Application.",
						},
						"use_administrator_credentials": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to use the Admin's credentials to login to the security management server.",
						},
					},
				},
			},
			"lea": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Used to setup the LEA client entity.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_permissions": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Log reading permissions for the LEA client entity.",
						},
						"administrator_profile": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A profile to set the log reading permissions by for the client entity.",
						},
						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to enable this client entity on the Opsec Application.",
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
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Comments string.",
			},
		},
	}
}

func dataSourceManagementOpsecApplicationRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showOpsecApplicationRes, err := client.ApiCall("show-opsec-application", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showOpsecApplicationRes.Success {
		return fmt.Errorf(showOpsecApplicationRes.ErrorMsg)
	}

	opsecApplication := showOpsecApplicationRes.GetData()

	log.Println("Read OpsecApplication - Show JSON = ", opsecApplication)

	if v := opsecApplication["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := opsecApplication["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if opsecApplication["cpmi"] != nil {

		cpmiMap := opsecApplication["cpmi"].(map[string]interface{})

		cpmiMapToReturn := make(map[string]interface{})

		if v, _ := cpmiMap["administrator-profile"]; v != nil {
			cpmiMapToReturn["administrator_profile"] = v
		}
		if v, _ := cpmiMap["enabled"]; v != nil {
			cpmiMapToReturn["enabled"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := cpmiMap["use-administrator-credentials"]; v != nil {
			cpmiMapToReturn["use_administrator_credentials"] = strconv.FormatBool(v.(bool))
		}
		_ = d.Set("cpmi", cpmiMapToReturn)
	} else {
		_ = d.Set("cpmi", nil)
	}

	if v := opsecApplication["host"]; v != nil {
		_ = d.Set("host", v)
	}

	if opsecApplication["lea"] != nil {

		leaMap := opsecApplication["lea"].(map[string]interface{})

		leaMapToReturn := make(map[string]interface{})

		if v, _ := leaMap["access-permissions"]; v != nil {
			leaMapToReturn["access_permissions"] = v
		}
		if v, _ := leaMap["administrator-profile"]; v != nil {
			leaMapToReturn["administrator_profile"] = v
		}
		if v, _ := leaMap["enabled"]; v != nil {
			leaMapToReturn["enabled"] = strconv.FormatBool(v.(bool))
		}
		_ = d.Set("lea", leaMapToReturn)
	} else {
		_ = d.Set("lea", nil)
	}

	if v := opsecApplication["one-time-password"]; v != nil {
		_ = d.Set("one_time_password", v)
	}

	if opsecApplication["tags"] != nil {
		tagsJson, ok := opsecApplication["tags"].([]interface{})
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

	if v := opsecApplication["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := opsecApplication["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil
}
