package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaPasswordPolicy() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaPasswordPolicy,
        Read:   readGaiaPasswordPolicy,
        Update: updateGaiaPasswordPolicy,
        Delete: deleteGaiaPasswordPolicy,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "lock_settings": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `password change configuration`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "inactivity_settings": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `inactivity configuration`,
                            MaxItems:    1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "lock_unused_accounts_enabled": {
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        Description: `Password lock unused accounts, default: false`,
                                    },
                                    "inactivity_threshold_days": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Description: `Inactivity days to password expiration lockout, default value is 365 days`,
                                    },
                                },
                            },
                        },
                        "failed_attempts_settings": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `failed attempts configuration`,
                            MaxItems:    1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "failed_lock_duration_seconds": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Description: `Password failed logging lockout duration, default value is 1200`,
                                    },
                                    "failed_lock_enforced_on_admin": {
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        Description: `Enforce failed lockout on admin user, default value is false`,
                                    },
                                    "failed_lock_enabled": {
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        Description: `Lock user after exceeded maximum allowed login attempts, default value is false`,
                                    },
                                    "failed_attempts_allowed": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Description: `Amount of login attempts allowed before lockout, default value is 10 attempts`,
                                    },
                                },
                            },
                        },
                        "password_expiration_days": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Sensitive:   true,
                            Description: `Password expiration lifetime, default value is 'never'`,
                        },
                        "password_expiration_warning_days": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Sensitive:   true,
                            Description: `Number of days before a password expires that the user gets warned, default value is 7 days`,
                        },
                        "password_expiration_maximum_days_before_lock": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Sensitive:   true,
                            Description: `Password expiration lockout in days, default value is 'never'`,
                        },
                        "must_one_time_password_enabled": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Sensitive:   true,
                            Description: `Forces a user to change their password after                it has been set via \"User Management\" (but not via \"Self Password Change\" or forced change at login).Use this command to set the value. Default value is false`,
                        },
                    },
                },
            },
            "password_history": {
                Type:        schema.TypeList,
                Optional:    true,
                Sensitive:   true,
                Description: `password history configuration`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "check_history_enabled": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Password history check, default value is false`,
                        },
                        "repeated_history_length": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `Password history length, default value is 10 entries`,
                        },
                    },
                },
            },
            "password_strength": {
                Type:        schema.TypeList,
                Optional:    true,
                Sensitive:   true,
                Description: `password strength configuration`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "minimum_length": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `default length is 6`,
                        },
                        "complexity": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `default value is 2`,
                        },
                        "palindrome_check_enabled": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Password palindrome check, default value is true`,
                        },
                    },
                },
            },
            "all_users_require_two_factor_authentication": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `Force Two-Factor Authentication for all users. Upon their next login, if Two-Factor Authentication is not already set up, the users will be required to generate the authentication keys.`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
        },
    }
}

func createGaiaPasswordPolicy(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v := d.Get("lock_settings"); len(v.([]interface{})) > 0 {
        if lsMap, ok := v.([]interface{})[0].(map[string]interface{}); ok {
            lockSettingsMap := map[string]interface{}{}
            if val, ok2 := lsMap["password_expiration_days"]; ok2 {
                lockSettingsMap["password-expiration-days"] = val
            }
            if val, ok2 := lsMap["password_expiration_warning_days"]; ok2 {
                lockSettingsMap["password-expiration-warning-days"] = val
            }
            if val, ok2 := lsMap["password_expiration_maximum_days_before_lock"]; ok2 {
                lockSettingsMap["password-expiration-maximum-days-before-lock"] = val
            }
            if v2, ok2 := d.GetOkExists("lock_settings.0.must_one_time_password_enabled"); ok2 {
                lockSettingsMap["must-one-time-password-enabled"] = v2.(bool)
            }
            if iaList := d.Get("lock_settings.0.inactivity_settings"); len(iaList.([]interface{})) > 0 {
                if iaMap, ok2 := iaList.([]interface{})[0].(map[string]interface{}); ok2 {
                    inactivityMap := map[string]interface{}{}
                    if v2, ok3 := d.GetOkExists("lock_settings.0.inactivity_settings.0.lock_unused_accounts_enabled"); ok3 {
                        inactivityMap["lock-unused-accounts-enabled"] = v2.(bool)
                    }
                    if val, ok3 := iaMap["inactivity_threshold_days"]; ok3 {
                        inactivityMap["inactivity-threshold-days"] = val
                    }
                    lockSettingsMap["inactivity-settings"] = inactivityMap
                }
            }
            if faList := d.Get("lock_settings.0.failed_attempts_settings"); len(faList.([]interface{})) > 0 {
                if faMap, ok2 := faList.([]interface{})[0].(map[string]interface{}); ok2 {
                    failedMap := map[string]interface{}{}
                    if val, ok3 := faMap["failed_lock_duration_seconds"]; ok3 {
                        failedMap["failed-lock-duration-seconds"] = val
                    }
                    if v2, ok3 := d.GetOkExists("lock_settings.0.failed_attempts_settings.0.failed_lock_enforced_on_admin"); ok3 {
                        failedMap["failed-lock-enforced-on-admin"] = v2.(bool)
                    }
                    if v2, ok3 := d.GetOkExists("lock_settings.0.failed_attempts_settings.0.failed_lock_enabled"); ok3 {
                        failedMap["failed-lock-enabled"] = v2.(bool)
                    }
                    if val, ok3 := faMap["failed_attempts_allowed"]; ok3 {
                        failedMap["failed-attempts-allowed"] = val
                    }
                    lockSettingsMap["failed-attempts-settings"] = failedMap
                }
            }
            payload["lock-settings"] = lockSettingsMap
        }
    }

    if v := d.Get("password_history"); len(v.([]interface{})) > 0 {
        phMap := map[string]interface{}{}
        if v2, ok := d.GetOkExists("password_history.0.check_history_enabled"); ok {
            phMap["check-history-enabled"] = v2.(bool)
        }
        if val, ok := d.GetOk("password_history.0.repeated_history_length"); ok {
            phMap["repeated-history-length"] = val
        }
        payload["password-history"] = phMap
    }

    if v := d.Get("password_strength"); len(v.([]interface{})) > 0 {
        psMap := map[string]interface{}{}
        if val, ok := d.GetOk("password_strength.0.minimum_length"); ok {
            psMap["minimum-length"] = val
        }
        if val, ok := d.GetOk("password_strength.0.complexity"); ok {
            psMap["complexity"] = val
        }
        if v2, ok := d.GetOkExists("password_strength.0.palindrome_check_enabled"); ok {
            psMap["palindrome-check-enabled"] = v2.(bool)
        }
        payload["password-strength"] = psMap
    }

    if v, ok := d.GetOkExists("all_users_require_two_factor_authentication"); ok {
        payload["all-users-require-two-factor-authentication"] = v.(bool)
    }

    log.Println("Create PasswordPolicy - Map = ", payload)

    addPasswordPolicyRes, err := client.ApiCallSimple("set-password-policy", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addPasswordPolicyRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addPasswordPolicyRes.Success {
            errMsg = addPasswordPolicyRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addPasswordPolicyRes.GetData()
        }

        debugLogOperation(
            "password-policy",        // resource type
            "create",                       // operation
            "set-password-policy",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add password-policy: %v", err)
    }
    if !addPasswordPolicyRes.Success {
        if addPasswordPolicyRes.ErrorMsg != "" {
            return fmt.Errorf(addPasswordPolicyRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("password-policy-" + acctest.RandString(10)))
    return readGaiaPasswordPolicy(d, m)
}

func readGaiaPasswordPolicy(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showPasswordPolicyRes, err := client.ApiCallSimple("show-password-policy", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showPasswordPolicyRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showPasswordPolicyRes.Success {
            errMsg = showPasswordPolicyRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showPasswordPolicyRes.GetData()
        }

        debugLogOperation(
            "password-policy",        // resource type
            "read",                       // operation
            "show-password-policy",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show password-policy: %v", err)
    }
    if !showPasswordPolicyRes.Success {
        if data := showPasswordPolicyRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showPasswordPolicyRes.ErrorMsg)
    }

    passwordPolicy := showPasswordPolicyRes.GetData()

    log.Println("Read PasswordPolicy - Show JSON = ", passwordPolicy)

    if v, exists := passwordPolicy["lock-settings"]; exists {
        if m, ok := v.(map[string]interface{}); ok {
            lsMap := map[string]interface{}{
                "password_expiration_days":                     m["password-expiration-days"],
                "password_expiration_warning_days":             m["password-expiration-warning-days"],
                "password_expiration_maximum_days_before_lock": m["password-expiration-maximum-days-before-lock"],
                "must_one_time_password_enabled":               m["must-one-time-password-enabled"],
            }
            if ia, ok := m["inactivity-settings"].(map[string]interface{}); ok {
                lsMap["inactivity_settings"] = []interface{}{map[string]interface{}{
                    "lock_unused_accounts_enabled": ia["lock-unused-accounts-enabled"],
                    "inactivity_threshold_days":    ia["inactivity-threshold-days"],
                }}
            }
            if fa, ok := m["failed-attempts-settings"].(map[string]interface{}); ok {
                lsMap["failed_attempts_settings"] = []interface{}{map[string]interface{}{
                    "failed_lock_duration_seconds":  fa["failed-lock-duration-seconds"],
                    "failed_lock_enforced_on_admin": fa["failed-lock-enforced-on-admin"],
                    "failed_lock_enabled":           fa["failed-lock-enabled"],
                    "failed_attempts_allowed":       fa["failed-attempts-allowed"],
                }}
            }
            d.Set("lock_settings", []interface{}{lsMap})
        }
    }
    if v, exists := passwordPolicy["password-history"]; exists {
        if m, ok := v.(map[string]interface{}); ok {
            d.Set("password_history", []interface{}{map[string]interface{}{
                "check_history_enabled":   m["check-history-enabled"],
                "repeated_history_length": m["repeated-history-length"],
            }})
        }
    }
    if v, exists := passwordPolicy["password-strength"]; exists {
        if m, ok := v.(map[string]interface{}); ok {
            d.Set("password_strength", []interface{}{map[string]interface{}{
                "minimum_length":           m["minimum-length"],
                "complexity":               m["complexity"],
                "palindrome_check_enabled": m["palindrome-check-enabled"],
            }})
        }
    }
    if v, exists := passwordPolicy["all-users-require-two-factor-authentication"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("all_users_require_two_factor_authentication", b)
        } else if s, ok := v.(string); ok {
            d.Set("all_users_require_two_factor_authentication", s == "true")
        }
    }
    if v, exists := passwordPolicy["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaPasswordPolicy(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v := d.Get("lock_settings"); len(v.([]interface{})) > 0 {
        if lsMap, ok := v.([]interface{})[0].(map[string]interface{}); ok {
            lockSettingsMap := map[string]interface{}{}
            if val, ok2 := lsMap["password_expiration_days"]; ok2 {
                lockSettingsMap["password-expiration-days"] = val
            }
            if val, ok2 := lsMap["password_expiration_warning_days"]; ok2 {
                lockSettingsMap["password-expiration-warning-days"] = val
            }
            if val, ok2 := lsMap["password_expiration_maximum_days_before_lock"]; ok2 {
                lockSettingsMap["password-expiration-maximum-days-before-lock"] = val
            }
            if v2, ok2 := d.GetOkExists("lock_settings.0.must_one_time_password_enabled"); ok2 {
                lockSettingsMap["must-one-time-password-enabled"] = v2.(bool)
            }
            if iaList := d.Get("lock_settings.0.inactivity_settings"); len(iaList.([]interface{})) > 0 {
                if iaMap, ok2 := iaList.([]interface{})[0].(map[string]interface{}); ok2 {
                    inactivityMap := map[string]interface{}{}
                    if v2, ok3 := d.GetOkExists("lock_settings.0.inactivity_settings.0.lock_unused_accounts_enabled"); ok3 {
                        inactivityMap["lock-unused-accounts-enabled"] = v2.(bool)
                    }
                    if val, ok3 := iaMap["inactivity_threshold_days"]; ok3 {
                        inactivityMap["inactivity-threshold-days"] = val
                    }
                    lockSettingsMap["inactivity-settings"] = inactivityMap
                }
            }
            if faList := d.Get("lock_settings.0.failed_attempts_settings"); len(faList.([]interface{})) > 0 {
                if faMap, ok2 := faList.([]interface{})[0].(map[string]interface{}); ok2 {
                    failedMap := map[string]interface{}{}
                    if val, ok3 := faMap["failed_lock_duration_seconds"]; ok3 {
                        failedMap["failed-lock-duration-seconds"] = val
                    }
                    if v2, ok3 := d.GetOkExists("lock_settings.0.failed_attempts_settings.0.failed_lock_enforced_on_admin"); ok3 {
                        failedMap["failed-lock-enforced-on-admin"] = v2.(bool)
                    }
                    if v2, ok3 := d.GetOkExists("lock_settings.0.failed_attempts_settings.0.failed_lock_enabled"); ok3 {
                        failedMap["failed-lock-enabled"] = v2.(bool)
                    }
                    if val, ok3 := faMap["failed_attempts_allowed"]; ok3 {
                        failedMap["failed-attempts-allowed"] = val
                    }
                    lockSettingsMap["failed-attempts-settings"] = failedMap
                }
            }
            payload["lock-settings"] = lockSettingsMap
        }
    }

    if v := d.Get("password_history"); len(v.([]interface{})) > 0 {
        phMap := map[string]interface{}{}
        if v2, ok := d.GetOkExists("password_history.0.check_history_enabled"); ok {
            phMap["check-history-enabled"] = v2.(bool)
        }
        if val, ok := d.GetOk("password_history.0.repeated_history_length"); ok {
            phMap["repeated-history-length"] = val
        }
        payload["password-history"] = phMap
    }

    if v := d.Get("password_strength"); len(v.([]interface{})) > 0 {
        psMap := map[string]interface{}{}
        if val, ok := d.GetOk("password_strength.0.minimum_length"); ok {
            psMap["minimum-length"] = val
        }
        if val, ok := d.GetOk("password_strength.0.complexity"); ok {
            psMap["complexity"] = val
        }
        if v2, ok := d.GetOkExists("password_strength.0.palindrome_check_enabled"); ok {
            psMap["palindrome-check-enabled"] = v2.(bool)
        }
        payload["password-strength"] = psMap
    }

    if v, ok := d.GetOkExists("all_users_require_two_factor_authentication"); ok {
        payload["all-users-require-two-factor-authentication"] = v.(bool)
    }

    setPasswordPolicyRes, err := client.ApiCallSimple("set-password-policy", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setPasswordPolicyRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setPasswordPolicyRes.Success {
            errMsg = setPasswordPolicyRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setPasswordPolicyRes.GetData()
        }

        debugLogOperation(
            "password-policy",        // resource type
            "update",                       // operation
            "set-password-policy",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set password-policy: %v", err)
    }
    if !setPasswordPolicyRes.Success {
        return fmt.Errorf(setPasswordPolicyRes.ErrorMsg)
    }

    return readGaiaPasswordPolicy(d, m)
}

func deleteGaiaPasswordPolicy(d *schema.ResourceData, m interface{}) error {


        // No API call - just remove the ID to indicate resource deletion
        d.SetId("")
        return nil
    }

    