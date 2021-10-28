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

func resourceManagementUserTemplate() *schema.Resource {
	return &schema.Resource{
		Create: createManagementUserTemplate,
		Read:   readManagementUserTemplate,
		Update: updateManagementUserTemplate,
		Delete: deleteManagementUserTemplate,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"expiration_by_global_properties": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Expiration date according to global properties.",
			},
			"expiration_date": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Expiration date in format: yyyy-MM-dd.",
			},
			"authentication_method": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Authentication method.",
				Default:     "undefined",
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
				Default:     true,
			},
			"from_hour": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Allow users connect from hour.",
				Default:     "00:00",
			},
			"to_hour": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Allow users connect until hour.",
				Default:     "23:59",
			},
			"allowed_locations": {
				Type:        schema.TypeMap,
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
				Type:        schema.TypeMap,
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
		},
	}
}

func createManagementUserTemplate(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	userTemplate := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		userTemplate["name"] = v.(string)
	}

	if v, ok := d.GetOkExists("expiration_by_global_properties"); ok {
		userTemplate["expiration-by-global-properties"] = v.(bool)
	}

	if v, ok := d.GetOk("expiration_date"); ok {
		userTemplate["expiration-date"] = v.(string)
	}

	if v, ok := d.GetOk("authentication_method"); ok {
		userTemplate["authentication-method"] = v.(string)
	}

	if v, ok := d.GetOk("radius_server"); ok {
		userTemplate["radius-server"] = v.(string)
	}

	if v, ok := d.GetOk("tacacs_server"); ok {
		userTemplate["tacacs-server"] = v.(string)
	}

	if v, ok := d.GetOk("connect_on_days"); ok {
		userTemplate["connect-on-days"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOkExists("connect_daily"); ok {
		userTemplate["connect-daily"] = v.(bool)
	}

	if v, ok := d.GetOk("from_hour"); ok {
		userTemplate["from-hour"] = v.(string)
	}

	if v, ok := d.GetOk("to_hour"); ok {
		userTemplate["to-hour"] = v.(string)
	}

	if _, ok := d.GetOk("allowed_locations"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("allowed_locations.destinations"); ok {
			res["destinations"] = v.(*schema.Set).List()
		}
		if v, ok := d.GetOk("allowed_locations.sources"); ok {
			res["sources"] = v.(*schema.Set).List()
		}
		userTemplate["allowed-locations"] = res
	}

	if _, ok := d.GetOk("encryption"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("encryption.enable_ike"); ok {
			res["enable-ike"] = v.(bool)
		}
		if v, ok := d.GetOk("encryption.enable_public_key"); ok {
			res["enable-public-key"] = v.(bool)
		}
		if v, ok := d.GetOk("encryption.enable_shared_secret"); ok {
			res["enable-shared-secret"] = v.(bool)
		}
		userTemplate["encryption"] = res
	}

	if v, ok := d.GetOk("tags"); ok {
		userTemplate["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		userTemplate["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		userTemplate["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		userTemplate["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		userTemplate["ignore-errors"] = v.(bool)
	}

	log.Println("Create UserTemplate - Map = ", userTemplate)

	addUserTemplateRes, err := client.ApiCall("add-user-template", userTemplate, client.GetSessionID(), true, false)
	if err != nil || !addUserTemplateRes.Success {
		if addUserTemplateRes.ErrorMsg != "" {
			return fmt.Errorf(addUserTemplateRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addUserTemplateRes.GetData()["uid"].(string))

	return readManagementUserTemplate(d, m)
}

func readManagementUserTemplate(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showUserTemplateRes, err := client.ApiCall("show-user-template", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showUserTemplateRes.Success {
		if objectNotFound(showUserTemplateRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showUserTemplateRes.ErrorMsg)
	}

	userTemplate := showUserTemplateRes.GetData()

	log.Println("Read UserTemplate - Show JSON = ", userTemplate)

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

	if v := userTemplate["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := userTemplate["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementUserTemplate(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	userTemplate := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		userTemplate["name"] = oldName
		userTemplate["new-name"] = newName
	} else {
		userTemplate["name"] = d.Get("name")
	}

	if v, ok := d.GetOkExists("expiration_by_global_properties"); ok {
		userTemplate["expiration-by-global-properties"] = v.(bool)
	}

	if ok := d.HasChange("expiration_date"); ok {
		userTemplate["expiration-date"] = d.Get("expiration_date")
	}

	if ok := d.HasChange("authentication_method"); ok {
		userTemplate["authentication-method"] = d.Get("authentication_method")
	}

	if ok := d.HasChange("radius_server"); ok {
		userTemplate["radius-server"] = d.Get("radius_server")
	}

	if ok := d.HasChange("tacacs_server"); ok {
		userTemplate["tacacs-server"] = d.Get("tacacs_server")
	}

	if d.HasChange("connect_on_days") {
		if v, ok := d.GetOk("connect_on_days"); ok {
			userTemplate["connect_on_days"] = v.(*schema.Set).List()
		}
	}

	if v, ok := d.GetOkExists("connect_daily"); ok {
		userTemplate["connect-daily"] = v.(bool)
	}

	if ok := d.HasChange("from_hour"); ok {
		userTemplate["from-hour"] = d.Get("from_hour")
	}

	if ok := d.HasChange("to_hour"); ok {
		userTemplate["to-hour"] = d.Get("to_hour")
	}

	if d.HasChange("allowed_locations") {
		defaultLocationUid := "97aeb369-9aea-11d5-bd16-0090272ccb30"

		if _, ok := d.GetOk("allowed_locations"); ok {

			res := make(map[string]interface{})

			if d.HasChange("allowed_locations.destinations") {
				if v, ok := d.GetOk("allowed_locations.destinations"); ok {
					res["destinations"] = v.(*schema.Set).List()
				} else {
					res["destinations"] = defaultLocationUid
				}
			}

			if d.HasChange("allowed_locations.sources") {
				if v, ok := d.GetOk("allowed_locations.destinations"); ok {
					res["sources"] = v.(*schema.Set).List()
				} else {
					res["sources"] = defaultLocationUid
				}
			}

			userTemplate["allowed-locations"] = res
		} else {
			userTemplate["allowed-locations"] = map[string]interface{}{"sources": defaultLocationUid, "destinations": defaultLocationUid}
		}
	}

	if d.HasChange("encryption") {

		if _, ok := d.GetOk("encryption"); ok {

			res := make(map[string]interface{})

			if d.HasChange("encryption.enable_ike") {
				res["enable-ike"] = d.Get("encryption.enable_ike")
			}
			if d.HasChange("encryption.enable_public_key") {
				res["enable-public-key"] = d.Get("encryption.enable_public_key")
			}
			if d.HasChange("encryption.enable_shared_secret") {
				res["enable-shared-secret"] = d.Get("encryption.enable_shared_secret")
			}
			userTemplate["encryption"] = res
		} else {
			userTemplate["encryption"] = map[string]interface{}{"enable-ike": false}
		}
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			userTemplate["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			userTemplate["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		userTemplate["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		userTemplate["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		userTemplate["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		userTemplate["ignore-errors"] = v.(bool)
	}

	log.Println("Update UserTemplate - Map = ", userTemplate)

	updateUserTemplateRes, err := client.ApiCall("set-user-template", userTemplate, client.GetSessionID(), true, false)
	if err != nil || !updateUserTemplateRes.Success {
		if updateUserTemplateRes.ErrorMsg != "" {
			return fmt.Errorf(updateUserTemplateRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementUserTemplate(d, m)
}

func deleteManagementUserTemplate(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	userTemplatePayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete UserTemplate")

	deleteUserTemplateRes, err := client.ApiCall("delete-user-template", userTemplatePayload, client.GetSessionID(), true, false)
	if err != nil || !deleteUserTemplateRes.Success {
		if deleteUserTemplateRes.ErrorMsg != "" {
			return fmt.Errorf(deleteUserTemplateRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
