package checkpoint

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceManagementUser() *schema.Resource {
	return &schema.Resource{
		Create: createManagementUser,
		Read:   readManagementUser,
		Update: updateManagementUser,
		Delete: deleteManagementUser,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"email": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User email.",
			},
			"expiration_date": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Expiration date in format: yyyy-MM-dd.",
			},
			"phone_number": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User phone number.",
			},
			"authentication_method": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Authentication method.",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "Checkpoint password authentication method identified by the name or UID. Must be set when \"authentication-method\" was selected to be \"Check Point Password\".",
			},
			"radius_server": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "RADIUS server object identified by the name or UID. Must be set when \"authentication-method\" was selected to be \"RADIUS\".",
			},
			"tacacs_server": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "TACACS server object identified by the name or UID. Must be set when \"authentication-method\" was selected to be \"TACACS\".",
			},
			"connect_on_days": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Days users allow to connect.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"connect_daily": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Connect every day.",
			},
			"from_hour": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Allow users connect from hour.",
			},
			"to_hour": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Allow users connect until hour.",
			},
			"allowed_locations": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "User allowed locations.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"destinations": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Collection of allowed destination locations name or uid.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"sources": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Collection of allowed source locations name or uid.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"encryption": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "User encryption.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_ike": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable IKE encryption for users.",
						},
						"enable_public_key": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable IKE public key.",
						},
						"enable_shared_secret": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable IKE shared secret.",
						},
						"shared_secret": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IKE shared secret.",
						},
					},
				},
			},
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"color": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Color of the object. Should be one of existing colors.",
				Default:     "black",
			},
			"comments": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comments string.",
			},
			"ignore_warnings": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Apply changes ignoring warnings.",
				Default:     false,
			},
			"ignore_errors": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.",
				Default:     false,
			},
			"template": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User template name or UID.",
				Default:     "Default",
			},
		},
	}
}

func createManagementUser(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	user := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		user["name"] = v.(string)
	}

	if v, ok := d.GetOk("email"); ok {
		user["email"] = v.(string)
	}

	if v, ok := d.GetOk("expiration_date"); ok {
		user["expiration-date"] = v.(string)
	}

	if v, ok := d.GetOk("phone_number"); ok {
		user["phone-number"] = v.(string)
	}

	if v, ok := d.GetOk("authentication_method"); ok {
		user["authentication-method"] = v.(string)
	}

	if v, ok := d.GetOk("password"); ok {
		user["password"] = v.(string)
	}

	if v, ok := d.GetOk("radius_server"); ok {
		user["radius-server"] = v.(string)
	}

	if v, ok := d.GetOk("tacacs_server"); ok {
		user["tacacs-server"] = v.(string)
	}

	if v, ok := d.GetOk("connect_on_days"); ok {
		user["connect-on-days"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOkExists("connect_daily"); ok {
		user["connect-daily"] = v.(bool)
	}

	if v, ok := d.GetOk("from_hour"); ok {
		user["from-hour"] = v.(string)
	}

	if v, ok := d.GetOk("to_hour"); ok {
		user["to-hour"] = v.(string)
	}

	if v, ok := d.GetOk("allowed_locations"); ok {

		allowedLocationsList := v.([]interface{})

		if len(allowedLocationsList) > 0 {

			allowedLocationsPayload := make(map[string]interface{})

			if v, ok := d.GetOk("allowed_locations.0.destinations"); ok {
				allowedLocationsPayload["destinations"] = v.(*schema.Set).List()
			}
			if v, ok := d.GetOk("allowed_locations.0.sources"); ok {
				allowedLocationsPayload["sources"] = v.(*schema.Set).List()
			}
			user["allowed-locations"] = allowedLocationsPayload
		}
	}

	if v, ok := d.GetOk("encryption"); ok {

		encryptionList := v.([]interface{})

		if len(encryptionList) > 0 {

			encryptionPayload := make(map[string]interface{})

			if v, ok := d.GetOk("encryption.0.enable_ike"); ok {
				encryptionPayload["enable-ike"] = v.(bool)
			}
			if v, ok := d.GetOk("encryption.0.enable_public_key"); ok {
				encryptionPayload["enable-public-key"] = v.(bool)
			}
			if v, ok := d.GetOk("encryption.0.enable_shared_secret"); ok {
				encryptionPayload["enable-shared-secret"] = v.(bool)
			}
			if v, ok := d.GetOk("encryption.0.shared_secret"); ok {
				encryptionPayload["shared-secret"] = v.(string)
			}
			user["encryption"] = encryptionPayload
		}
	}

	if v, ok := d.GetOk("tags"); ok {
		user["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		user["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		user["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		user["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		user["ignore-errors"] = v.(bool)
	}

	if v, ok := d.GetOk("template"); ok {
		user["template"] = v.(string)
	}

	log.Println("Create User - Map = ", user)

	addUserRes, err := client.ApiCall("add-user", user, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !addUserRes.Success {
		if addUserRes.ErrorMsg != "" {
			return fmt.Errorf(addUserRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addUserRes.GetData()["uid"].(string))

	return readManagementUser(d, m)
}

func readManagementUser(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showUserRes, err := client.ApiCall("show-user", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showUserRes.Success {
		if objectNotFound(showUserRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showUserRes.ErrorMsg)
	}

	user := showUserRes.GetData()

	log.Println("Read User - Show JSON = ", user)

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

		if v := allowedLocationsMap["destinations"]; v != nil {
			allowedLocationsMapToReturn["destinations"] = v
		}
		if v := allowedLocationsMap["sources"]; v != nil {
			allowedLocationsMapToReturn["sources"] = v
		}
		_, allowedLocationsInConf := d.GetOk("allowed_locations")
		defaultAllowedLocations := map[string]interface{}{"sources": "['97aeb369-9aea-11d5-bd16-0090272ccb30']", "destinations": "['97aeb369-9aea-11d5-bd16-0090272ccb30']"}
		if reflect.DeepEqual(defaultAllowedLocations, allowedLocationsMapToReturn) && !allowedLocationsInConf {
			_ = d.Set("allowed_locations", []interface{}{})
		} else {
			_ = d.Set("allowed_locations", []interface{}{allowedLocationsMapToReturn})
		}

	} else {
		_ = d.Set("allowed_locations", nil)
	}

	if user["encryption"] != nil {

		encryptionMap := user["encryption"].(map[string]interface{})

		encryptionMapToReturn := make(map[string]interface{})

		if v := encryptionMap["enable-ike"]; v != nil {
			encryptionMapToReturn["enable_ike"] = v
		}
		if v := encryptionMap["enable-public-key"]; v != nil {
			encryptionMapToReturn["enable_public_key"] = v
		}
		if v := encryptionMap["enable-shared-secret"]; v != nil {
			encryptionMapToReturn["enable_shared_secret"] = v
		}
		if v := encryptionMap["shared-secret"]; v != nil {
			encryptionMapToReturn["shared_secret"] = v
		}
		_, encryptionInConf := d.GetOk("encryption")
		defaultEncryption := map[string]interface{}{"enable_ike": "false"}
		if reflect.DeepEqual(defaultEncryption, encryptionMapToReturn) && !encryptionInConf {
			_ = d.Set("encryption", []interface{}{})
		} else {
			_ = d.Set("encryption", []interface{}{encryptionMapToReturn})
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

func updateManagementUser(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	user := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		user["name"] = oldName
		user["new-name"] = newName
	} else {
		user["name"] = d.Get("name")
	}

	if ok := d.HasChange("email"); ok {
		user["email"] = d.Get("email")
	}

	if ok := d.HasChange("expiration_date"); ok {
		user["expiration-date"] = d.Get("expiration_date")
	}

	if ok := d.HasChange("phone_number"); ok {
		user["phone-number"] = d.Get("phone_number")
	}

	if ok := d.HasChange("authentication_method"); ok {
		user["authentication-method"] = d.Get("authentication_method")
	}

	if ok := d.HasChange("password"); ok {
		user["password"] = d.Get("password")
	}

	if ok := d.HasChange("radius_server"); ok {
		user["radius-server"] = d.Get("radius_server")
	}

	if ok := d.HasChange("tacacs_server"); ok {
		user["tacacs-server"] = d.Get("tacacs_server")
	}

	if d.HasChange("connect_on_days") {
		if v, ok := d.GetOk("connect_on_days"); ok {
			user["connect_on_days"] = v.(*schema.Set).List()
		}
	}

	if v, ok := d.GetOkExists("connect_daily"); ok {
		user["connect-daily"] = v.(bool)
	}

	if ok := d.HasChange("from_hour"); ok {
		user["from-hour"] = d.Get("from_hour")
	}

	if ok := d.HasChange("to_hour"); ok {
		user["to-hour"] = d.Get("to_hour")
	}

	if d.HasChange("allowed_locations") {

		if v, ok := d.GetOk("allowed_locations"); ok {

			allowedLocationsList := v.([]interface{})

			if len(allowedLocationsList) > 0 {

				allowedLocationsPayload := make(map[string]interface{})

				if v, ok := d.GetOk("allowed_locations.0.destinations"); ok {
					allowedLocationsPayload["destinations"] = v.(*schema.Set).List()
				}
				if v, ok := d.GetOk("allowed_locations.0.sources"); ok {
					allowedLocationsPayload["sources"] = v.(*schema.Set).List()
				}
				user["allowed-locations"] = allowedLocationsPayload
			}
		}
	}

	if d.HasChange("encryption") {

		if v, ok := d.GetOk("encryption"); ok {

			encryptionList := v.([]interface{})

			if len(encryptionList) > 0 {

				encryptionPayload := make(map[string]interface{})

				if v, ok := d.GetOkExists("encryption.0.enable_ike"); ok {
					encryptionPayload["enable-ike"] = v.(bool)
				}
				if v, ok := d.GetOkExists("encryption.0.enable_public_key"); ok {
					encryptionPayload["enable-public-key"] = v.(bool)
				}
				if v, ok := d.GetOkExists("encryption.0.enable_shared_secret"); ok {
					encryptionPayload["enable-shared-secret"] = v.(bool)
				}
				if v, ok := d.GetOk("encryption.0.shared_secret"); ok {
					encryptionPayload["shared-secret"] = v.(string)
				}
				user["encryption"] = encryptionPayload
			}
		}
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			user["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			user["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		user["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		user["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		user["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		user["ignore-errors"] = v.(bool)
	}

	log.Println("Update User - Map = ", user)

	updateUserRes, err := client.ApiCall("set-user", user, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !updateUserRes.Success {
		if updateUserRes.ErrorMsg != "" {
			return fmt.Errorf(updateUserRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementUser(d, m)
}

func deleteManagementUser(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	userPayload := map[string]interface{}{
		"uid": d.Id(),
	}
	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		userPayload["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		userPayload["ignore-errors"] = v.(bool)
	}
	log.Println("Delete User")

	deleteUserRes, err := client.ApiCall("delete-user", userPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteUserRes.Success {
		if deleteUserRes.ErrorMsg != "" {
			return fmt.Errorf(deleteUserRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
