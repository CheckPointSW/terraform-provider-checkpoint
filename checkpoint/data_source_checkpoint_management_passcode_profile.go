package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementPasscodeProfile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementPasscodeProfileRead,
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
			"allow_simple_passcode": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "The passcode length is 4 and only numeric values allowed.",
			},
			"min_passcode_length": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Minimum passcode length - relevant if \"allow-simple-passcode\" is disable.",
			},
			"require_alphanumeric_passcode": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Require alphanumeric characters in the passcode - relevant if \"allow-simple-passcode\" is disable.",
			},
			"min_passcode_complex_characters": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Minimum number of complex characters (if \"require-alphanumeric-passcode\" is enabled). The number of the complex characters cannot be greater than number of the passcode length.",
			},
			"force_passcode_expiration": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable/disable expiration date to the passcode.",
			},
			"passcode_expiration_period": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The period in days after which the passcode will expire.",
			},
			"enable_inactivity_time_lock": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Lock the device if app is inactive.",
			},
			"max_inactivity_time_lock": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Time without user input before passcode must be re-entered (in minutes).",
			},
			"enable_passcode_failed_attempts": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Exit after few failures in passcode verification.",
			},
			"max_passcode_failed_attempts": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Number of failed attempts allowed.",
			},
			"enable_passcode_history": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Check passcode history for reparations.",
			},
			"passcode_history": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Number of passcodes that will be kept in history.",
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

func dataSourceManagementPasscodeProfileRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
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

	if v := passcodeProfile["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

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

	return nil

}
