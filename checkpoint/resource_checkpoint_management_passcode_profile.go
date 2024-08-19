package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementPasscodeProfile() *schema.Resource {
	return &schema.Resource{
		Create: createManagementPasscodeProfile,
		Read:   readManagementPasscodeProfile,
		Update: updateManagementPasscodeProfile,
		Delete: deleteManagementPasscodeProfile,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"allow_simple_passcode": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "The passcode length is 4 and only numeric values allowed.",
				Default:     true,
			},
			"min_passcode_length": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Minimum passcode length - relevant if \"allow-simple-passcode\" is disable.",
				Default:     4,
			},
			"require_alphanumeric_passcode": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Require alphanumeric characters in the passcode - relevant if \"allow-simple-passcode\" is disable.",
				Default:     false,
			},
			"min_passcode_complex_characters": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Minimum number of complex characters (if \"require-alphanumeric-passcode\" is enabled). The number of the complex characters cannot be greater than number of the passcode length.",
				Default:     0,
			},
			"force_passcode_expiration": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable/disable expiration date to the passcode.",
				Default:     false,
			},
			"passcode_expiration_period": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The period in days after which the passcode will expire.",
				Default:     90,
			},
			"enable_inactivity_time_lock": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Lock the device if app is inactive.",
				Default:     false,
			},
			"max_inactivity_time_lock": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Time without user input before passcode must be re-entered (in minutes).",
				Default:     15,
			},
			"enable_passcode_failed_attempts": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Exit after few failures in passcode verification.",
				Default:     false,
			},
			"max_passcode_failed_attempts": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Number of failed attempts allowed.",
				Default:     4,
			},
			"enable_passcode_history": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Check passcode history for reparations.",
				Default:     false,
			},
			"passcode_history": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Number of passcodes that will be kept in history.",
				Default:     8,
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

func createManagementPasscodeProfile(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	passcodeProfile := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		passcodeProfile["name"] = v.(string)
	}

	if v, ok := d.GetOkExists("allow_simple_passcode"); ok {
		passcodeProfile["allow-simple-passcode"] = v.(bool)
	}

	if v, ok := d.GetOk("min_passcode_length"); ok {
		passcodeProfile["min-passcode-length"] = v.(int)
	}

	if v, ok := d.GetOkExists("require_alphanumeric_passcode"); ok {
		passcodeProfile["require-alphanumeric-passcode"] = v.(bool)
	}

	if v, ok := d.GetOk("min_passcode_complex_characters"); ok {
		passcodeProfile["min-passcode-complex-characters"] = v.(int)
	}

	if v, ok := d.GetOkExists("force_passcode_expiration"); ok {
		passcodeProfile["force-passcode-expiration"] = v.(bool)
	}

	if v, ok := d.GetOk("passcode_expiration_period"); ok {
		passcodeProfile["passcode-expiration-period"] = v.(int)
	}

	if v, ok := d.GetOkExists("enable_inactivity_time_lock"); ok {
		passcodeProfile["enable-inactivity-time-lock"] = v.(bool)
	}

	if v, ok := d.GetOk("max_inactivity_time_lock"); ok {
		passcodeProfile["max-inactivity-time-lock"] = v.(int)
	}

	if v, ok := d.GetOkExists("enable_passcode_failed_attempts"); ok {
		passcodeProfile["enable-passcode-failed-attempts"] = v.(bool)
	}

	if v, ok := d.GetOk("max_passcode_failed_attempts"); ok {
		passcodeProfile["max-passcode-failed-attempts"] = v.(int)
	}

	if v, ok := d.GetOkExists("enable_passcode_history"); ok {
		passcodeProfile["enable-passcode-history"] = v.(bool)
	}

	if v, ok := d.GetOk("passcode_history"); ok {
		passcodeProfile["passcode-history"] = v.(int)
	}

	if v, ok := d.GetOk("tags"); ok {
		passcodeProfile["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		passcodeProfile["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		passcodeProfile["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		passcodeProfile["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		passcodeProfile["ignore-errors"] = v.(bool)
	}

	log.Println("Create PasscodeProfile - Map = ", passcodeProfile)

	addPasscodeProfileRes, err := client.ApiCall("add-passcode-profile", passcodeProfile, client.GetSessionID(), true, false)
	if err != nil || !addPasscodeProfileRes.Success {
		if addPasscodeProfileRes.ErrorMsg != "" {
			return fmt.Errorf(addPasscodeProfileRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addPasscodeProfileRes.GetData()["uid"].(string))

	return readManagementPasscodeProfile(d, m)
}

func readManagementPasscodeProfile(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showPasscodeProfileRes, err := client.ApiCall("show-passcode-profile", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showPasscodeProfileRes.Success {
		if objectNotFound(showPasscodeProfileRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showPasscodeProfileRes.ErrorMsg)
	}

	passcodeProfile := showPasscodeProfileRes.GetData()

	log.Println("Read PasscodeProfile - Show JSON = ", passcodeProfile)

	if v := passcodeProfile["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := passcodeProfile["allow-simple-passcode"]; v != nil {
		_ = d.Set("allow_simple_passcode", v)
	}

	if v := passcodeProfile["min-passcode-length"]; v != nil {
		_ = d.Set("min_passcode_length", v)
	}

	if v := passcodeProfile["require-alphanumeric-passcode"]; v != nil {
		_ = d.Set("require_alphanumeric_passcode", v)
	}

	if v := passcodeProfile["min-passcode-complex-characters"]; v != nil {
		_ = d.Set("min_passcode_complex_characters", v)
	}

	if v := passcodeProfile["force-passcode-expiration"]; v != nil {
		_ = d.Set("force_passcode_expiration", v)
	}

	if v := passcodeProfile["passcode-expiration-period"]; v != nil {
		_ = d.Set("passcode_expiration_period", v)
	}

	if v := passcodeProfile["enable-inactivity-time-lock"]; v != nil {
		_ = d.Set("enable_inactivity_time_lock", v)
	}

	if v := passcodeProfile["max-inactivity-time-lock"]; v != nil {
		_ = d.Set("max_inactivity_time_lock", v)
	}

	if v := passcodeProfile["enable-passcode-failed-attempts"]; v != nil {
		_ = d.Set("enable_passcode_failed_attempts", v)
	}

	if v := passcodeProfile["max-passcode-failed-attempts"]; v != nil {
		_ = d.Set("max_passcode_failed_attempts", v)
	}

	if v := passcodeProfile["enable-passcode-history"]; v != nil {
		_ = d.Set("enable_passcode_history", v)
	}

	if v := passcodeProfile["passcode-history"]; v != nil {
		_ = d.Set("passcode_history", v)
	}

	if passcodeProfile["tags"] != nil {
		tagsJson, ok := passcodeProfile["tags"].([]interface{})
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

	if v := passcodeProfile["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := passcodeProfile["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := passcodeProfile["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := passcodeProfile["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementPasscodeProfile(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	passcodeProfile := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		passcodeProfile["name"] = oldName
		passcodeProfile["new-name"] = newName
	} else {
		passcodeProfile["name"] = d.Get("name")
	}

	if v, ok := d.GetOkExists("allow_simple_passcode"); ok {
		passcodeProfile["allow-simple-passcode"] = v.(bool)
	}

	if ok := d.HasChange("min_passcode_length"); ok {
		passcodeProfile["min-passcode-length"] = d.Get("min_passcode_length")
	}

	if v, ok := d.GetOkExists("require_alphanumeric_passcode"); ok {
		passcodeProfile["require-alphanumeric-passcode"] = v.(bool)
	}

	if ok := d.HasChange("min_passcode_complex_characters"); ok {
		passcodeProfile["min-passcode-complex-characters"] = d.Get("min_passcode_complex_characters")
	}

	if v, ok := d.GetOkExists("force_passcode_expiration"); ok {
		passcodeProfile["force-passcode-expiration"] = v.(bool)
	}

	if ok := d.HasChange("passcode_expiration_period"); ok {
		passcodeProfile["passcode-expiration-period"] = d.Get("passcode_expiration_period")
	}

	if v, ok := d.GetOkExists("enable_inactivity_time_lock"); ok {
		passcodeProfile["enable-inactivity-time-lock"] = v.(bool)
	}

	if ok := d.HasChange("max_inactivity_time_lock"); ok {
		passcodeProfile["max-inactivity-time-lock"] = d.Get("max_inactivity_time_lock")
	}

	if v, ok := d.GetOkExists("enable_passcode_failed_attempts"); ok {
		passcodeProfile["enable-passcode-failed-attempts"] = v.(bool)
	}

	if ok := d.HasChange("max_passcode_failed_attempts"); ok {
		passcodeProfile["max-passcode-failed-attempts"] = d.Get("max_passcode_failed_attempts")
	}

	if v, ok := d.GetOkExists("enable_passcode_history"); ok {
		passcodeProfile["enable-passcode-history"] = v.(bool)
	}

	if ok := d.HasChange("passcode_history"); ok {
		passcodeProfile["passcode-history"] = d.Get("passcode_history")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			passcodeProfile["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			passcodeProfile["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		passcodeProfile["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		passcodeProfile["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		passcodeProfile["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		passcodeProfile["ignore-errors"] = v.(bool)
	}

	log.Println("Update PasscodeProfile - Map = ", passcodeProfile)

	updatePasscodeProfileRes, err := client.ApiCall("set-passcode-profile", passcodeProfile, client.GetSessionID(), true, false)
	if err != nil || !updatePasscodeProfileRes.Success {
		if updatePasscodeProfileRes.ErrorMsg != "" {
			return fmt.Errorf(updatePasscodeProfileRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementPasscodeProfile(d, m)
}

func deleteManagementPasscodeProfile(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	passcodeProfilePayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete PasscodeProfile")

	deletePasscodeProfileRes, err := client.ApiCall("delete-passcode-profile", passcodeProfilePayload, client.GetSessionID(), true, false)
	if err != nil || !deletePasscodeProfileRes.Success {
		if deletePasscodeProfileRes.ErrorMsg != "" {
			return fmt.Errorf(deletePasscodeProfileRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
