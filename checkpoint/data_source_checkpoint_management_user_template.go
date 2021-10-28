package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"reflect"
	"strconv"
	"strings"
)

func dataSourceManagementUserTemplate() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementUserTemplateRead,
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
			"expiration_by_global_properties": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Expiration date according to global properties.",
			},
			"expiration_date": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Expiration date in format: yyyy-MM-dd.",
			},
			"authentication_method": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Authentication method.",
			},
			"radius_server": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "RADIUS server object identified by the name or UID. Must be set when \"authentication-method\" was selected to be \"RADIUS\".",
			},
			"tacacs_server": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "TACACS server object identified by the name or UID. Must be set when \"authentication-method\" was selected to be \"TACACS\".",
			},
			"connect_on_days": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Days users allow to connect.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"connect_daily": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Connect every day.",
			},
			"from_hour": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Allow users connect from hour.",
			},
			"to_hour": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Allow users connect until hour.",
			},
			"allowed_locations": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "User allowed locations.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"destinations": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Collection of allowed destination locations name or uid.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"sources": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Collection of allowed source locations name or uid.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"encryption": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "User encryption.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_ike": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable IKE encryption for users.",
						},
						"enable_public_key": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable IKE public key.",
						},
						"enable_shared_secret": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable IKE shared secret.",
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

func dataSourceManagementUserTemplateRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showUserTemplateRes, err := client.ApiCall("show-user-template", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showUserTemplateRes.Success {
		return fmt.Errorf(showUserTemplateRes.ErrorMsg)
	}

	userTemplate := showUserTemplateRes.GetData()

	log.Println("Read UserTemplate - Show JSON = ", userTemplate)

	if v := userTemplate["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := userTemplate["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := userTemplate["expiration-by-global-properties"]; v != nil {
		_ = d.Set("expiration_by_global_properties", v)
	}

	if v := userTemplate["expiration-date"]; v != nil {
		isoDate := v.(map[string]interface{})["iso-8601"].(string)
		date := strings.Split(isoDate, "T")[0]
		_ = d.Set("expiration_date", date)
	}

	if v := userTemplate["authentication-method"]; v != nil {
		_ = d.Set("authentication_method", v)
	}

	if v := userTemplate["radius-server"]; v != nil {
		_ = d.Set("radius_server", v.(map[string]interface{})["name"].(string))
	}

	if v := userTemplate["tacacs-server"]; v != nil {
		_ = d.Set("tacacs_server", v.(map[string]interface{})["name"].(string))
	}

	if userTemplate["connect_on_days"] != nil {
		connectOnDaysJson, ok := userTemplate["connect_on_days"].([]interface{})
		if ok {
			_ = d.Set("connect_on_days", connectOnDaysJson)
		}
	} else {
		_ = d.Set("connect_on_days", nil)
	}

	if v := userTemplate["connect-daily"]; v != nil {
		_ = d.Set("connect_daily", v)
	}

	if v := userTemplate["from-hour"]; v != nil {
		_ = d.Set("from_hour", v)
	}

	if v := userTemplate["to-hour"]; v != nil {
		_ = d.Set("to_hour", v)
	}

	if userTemplate["allowed-locations"] != nil {

		allowedLocationsMap := userTemplate["allowed-locations"].(map[string]interface{})

		allowedLocationsMapToReturn := make(map[string]interface{})

		if v, _ := allowedLocationsMap["destinations"]; v != nil {
			allowedLocationsMapToReturn["destinations"] = v
		}
		if v, _ := allowedLocationsMap["sources"]; v != nil {
			allowedLocationsMapToReturn["sources"] = v
		}

		_, allowedLocationsInConf := d.GetOk("allowed_locations")
		defaultAllowedLocations := map[string]interface{}{"sources": "['97aeb369-9aea-11d5-bd16-0090272ccb30']", "destinations": "['97aeb369-9aea-11d5-bd16-0090272ccb30']"}
		if reflect.DeepEqual(defaultAllowedLocations, allowedLocationsMapToReturn) && !allowedLocationsInConf {
			_ = d.Set("allowed_locations", map[string]interface{}{})
		} else {
			_ = d.Set("allowed_locations", allowedLocationsMapToReturn)
		}

	} else {
		_ = d.Set("allowed_locations", nil)
	}

	if userTemplate["encryption"] != nil {

		encryptionMap := userTemplate["encryption"].(map[string]interface{})

		encryptionMapToReturn := make(map[string]interface{})

		if v, _ := encryptionMap["ike"]; v != nil {
			encryptionMapToReturn["enable_ike"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := encryptionMap["public-key"]; v != nil {
			encryptionMapToReturn["enable_public_key"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := encryptionMap["shared-secret"]; v != nil {
			encryptionMapToReturn["enable_shared_secret"] = strconv.FormatBool(v.(bool))
		}

		_, encryptionInConf := d.GetOk("encryption")
		defaultEncryption := map[string]interface{}{"enable_ike": "false"}
		if reflect.DeepEqual(defaultEncryption, encryptionMapToReturn) && !encryptionInConf {
			_ = d.Set("encryption", map[string]interface{}{})
		} else {
			_ = d.Set("encryption", encryptionMapToReturn)
		}

	} else {
		_ = d.Set("encryption", nil)
	}

	if userTemplate["tags"] != nil {
		tagsJson, ok := userTemplate["tags"].([]interface{})
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

	if v := userTemplate["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := userTemplate["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil
}
