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

func dataSourceManagementUser() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementUserRead,
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
			"email": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "User email.",
			},
			"expiration_date": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Expiration date in format: yyyy-MM-dd.",
			},
			"phone_number": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "User phone number.",
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

func dataSourceManagementUserRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showUserRes, err := client.ApiCall("show-user", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showUserRes.Success {
		return fmt.Errorf(showUserRes.ErrorMsg)
	}

	user := showUserRes.GetData()

	log.Println("Read User - Show JSON = ", user)

	if v := user["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := user["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := user["email"]; v != nil {
		_ = d.Set("email", v)
	}

	if v := user["expiration-date"]; v != nil {
		isoDate := v.(map[string]interface{})["iso-8601"].(string)
		date := strings.Split(isoDate, "T")[0]
		_ = d.Set("expiration_date", date)
	}

	if v := user["phone-number"]; v != nil {
		_ = d.Set("phone_number", v)
	}

	if v := user["authentication-method"]; v != nil {
		_ = d.Set("authentication_method", v)
	}

	if v := user["radius-server"]; v != nil {
		_ = d.Set("radius_server", v.(map[string]interface{})["name"].(string))
	}

	if v := user["tacacs-server"]; v != nil {
		_ = d.Set("tacacs_server", v.(map[string]interface{})["name"].(string))
	}

	if user["connect_on_days"] != nil {
		if connectOnDaysJson, ok := user["connect_on_days"].([]interface{}); ok {
			_ = d.Set("connect_on_days", connectOnDaysJson)
		}
	} else {
		_ = d.Set("connect_on_days", nil)
	}

	if v := user["connect-daily"]; v != nil {
		_ = d.Set("connect_daily", v)
	}

	if v := user["from-hour"]; v != nil {
		_ = d.Set("from_hour", v)
	}

	if v := user["to-hour"]; v != nil {
		_ = d.Set("to_hour", v)
	}

	if user["allowed-locations"] != nil {

		allowedLocationsMap := user["allowed-locations"].(map[string]interface{})

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

	if user["encryption"] != nil {

		encryptionMap := user["encryption"].(map[string]interface{})

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
		defaultEncryption := map[string]interface{}{"enable_ike": "true", "enable_public_key": "true", "enable_shared_secret": "false"}
		if reflect.DeepEqual(defaultEncryption, encryptionMapToReturn) && !encryptionInConf {
			_ = d.Set("encryption", map[string]interface{}{})
		} else {
			_ = d.Set("encryption", encryptionMapToReturn)
		}
	} else {
		_ = d.Set("encryption", nil)
	}

	if user["tags"] != nil {
		tagsJson, ok := user["tags"].([]interface{})
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

	if v := user["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := user["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil
}
