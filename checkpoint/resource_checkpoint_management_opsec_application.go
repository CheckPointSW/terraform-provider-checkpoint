package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	
	"strconv"
)

func resourceManagementOpsecApplication() *schema.Resource {   
    return &schema.Resource{
        Create: createManagementOpsecApplication,
        Read:   readManagementOpsecApplication,
        Update: updateManagementOpsecApplication,
        Delete: deleteManagementOpsecApplication,
        Schema: map[string]*schema.Schema{ 
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "Object name.",
            },
            "cpmi": {
                Type:        schema.TypeMap,
                Optional:    true,
                Description: "Used to setup the CPMI client entity.",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "administrator_profile": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: "A profile to set the log reading permissions by for the client entity.",
                        },
                        "enabled": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: "Whether to enable this client entity on the Opsec Application.",
                        },
                        "use_administrator_credentials": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: "Whether to use the Admin's credentials to login to the security management server.",
                        },
                    },
                },
            },
            "host": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: "The host where the server is running. Pre-define the host as a network object.",
            },
            "lea": {
                Type:        schema.TypeMap,
                Optional:    true,
                Description: "Used to setup the LEA client entity.",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "access_permissions": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: "Log reading permissions for the LEA client entity.",
                            Default:     "show all",
                        },
                        "administrator_profile": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: "A profile to set the log reading permissions by for the client entity.",
                        },
                        "enabled": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: "Whether to enable this client entity on the Opsec Application.",
                        },
                    },
                },
            },
            "one_time_password": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: "A password required for establishing a Secure Internal Communication (SIC).",
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

func createManagementOpsecApplication(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)

    opsecApplication := make(map[string]interface{})

    if v, ok := d.GetOk("name"); ok {
        opsecApplication["name"] = v.(string)
    }

    if _, ok := d.GetOk("cpmi"); ok {

        res := make(map[string]interface{})

        if v, ok := d.GetOk("cpmi.administrator_profile"); ok {
            res["administrator-profile"] = v.(string)
        }
        if v, ok := d.GetOk("cpmi.enabled"); ok {
            res["enabled"] = v
        }
        if v, ok := d.GetOk("cpmi.use_administrator_credentials"); ok {
            res["use-administrator-credentials"] = v
        }
        opsecApplication["cpmi"] = res
    }

    if v, ok := d.GetOk("host"); ok {
        opsecApplication["host"] = v.(string)
    }

    if _, ok := d.GetOk("lea"); ok {

        res := make(map[string]interface{})

        if v, ok := d.GetOk("lea.access_permissions"); ok {
            res["access-permissions"] = v.(string)
        }
        if v, ok := d.GetOk("lea.administrator_profile"); ok {
            res["administrator-profile"] = v.(string)
        }
        if v, ok := d.GetOk("lea.enabled"); ok {
            res["enabled"] = v
        }
        opsecApplication["lea"] = res
    }

    if v, ok := d.GetOk("one_time_password"); ok {
        opsecApplication["one-time-password"] = v.(string)
    }

    if v, ok := d.GetOk("tags"); ok {
        opsecApplication["tags"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOk("color"); ok {
        opsecApplication["color"] = v.(string)
    }

    if v, ok := d.GetOk("comments"); ok {
        opsecApplication["comments"] = v.(string)
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
        opsecApplication["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
        opsecApplication["ignore-errors"] = v.(bool)
    }

    log.Println("Create OpsecApplication - Map = ", opsecApplication)

    addOpsecApplicationRes, err := client.ApiCall("add-opsec-application", opsecApplication, client.GetSessionID(), true, false)
    if err != nil || !addOpsecApplicationRes.Success {
        if addOpsecApplicationRes.ErrorMsg != "" {
            return fmt.Errorf(addOpsecApplicationRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }

    d.SetId(addOpsecApplicationRes.GetData()["uid"].(string))

    return readManagementOpsecApplication(d, m)
}

func readManagementOpsecApplication(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    payload := map[string]interface{}{
        "uid": d.Id(),
    }

    showOpsecApplicationRes, err := client.ApiCall("show-opsec-application", payload, client.GetSessionID(), true, false)
    if err != nil {
		return fmt.Errorf(err.Error())
	}
    if !showOpsecApplicationRes.Success {
		if objectNotFound(showOpsecApplicationRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
        return fmt.Errorf(showOpsecApplicationRes.ErrorMsg)
    }

    opsecApplication := showOpsecApplicationRes.GetData()

    log.Println("Read OpsecApplication - Show JSON = ", opsecApplication)

	if v := opsecApplication["name"]; v != nil {
		_ = d.Set("name", v)
	}

    if opsecApplication["cpmi"] != nil {

        cpmiMap := opsecApplication["cpmi"].(map[string]interface{})

        cpmiMapToReturn := make(map[string]interface{})

        if v, _ := cpmiMap["administrator-profile"]; v != nil{
            cpmiMapToReturn["administrator_profile"] = v
        }
        if v, _ := cpmiMap["enabled"]; v != nil{
            cpmiMapToReturn["enabled"] = strconv.FormatBool(v.(bool))
        }
        if v, _ := cpmiMap["use-administrator-credentials"]; v != nil{
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

        if v, _ := leaMap["access-permissions"]; v != nil{
            leaMapToReturn["access_permissions"] = v
        }
        if v, _ := leaMap["administrator-profile"]; v != nil{
            leaMapToReturn["administrator_profile"] = v
        }
        if v, _ := leaMap["enabled"]; v != nil{
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

	if v := opsecApplication["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := opsecApplication["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementOpsecApplication(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
    opsecApplication := make(map[string]interface{})

    if ok := d.HasChange("name"); ok {
        oldName, newName := d.GetChange("name")
        opsecApplication["name"] = oldName
        opsecApplication["new-name"] = newName
    } else {
        opsecApplication["name"] = d.Get("name")
    }

    if d.HasChange("cpmi") {

        if _, ok := d.GetOk("cpmi"); ok {

            res := make(map[string]interface{})

            if d.HasChange("cpmi.administrator_profile") {
                res["administrator-profile"] = d.Get("cpmi.administrator_profile")
            }
            if d.HasChange("cpmi.enabled") {
                res["enabled"] = d.Get("cpmi.enabled")
            }
            if d.HasChange("cpmi.use_administrator_credentials") {
                res["use-administrator-credentials"] = d.Get("cpmi.use_administrator_credentials")
            }
            opsecApplication["cpmi"] = res
        }
    }

    if ok := d.HasChange("host"); ok {
	       opsecApplication["host"] = d.Get("host")
    }

    if d.HasChange("lea") {

        if _, ok := d.GetOk("lea"); ok {

            res := make(map[string]interface{})

            if d.HasChange("lea.access_permissions") {
                res["access-permissions"] = d.Get("lea.access_permissions")
            }
            if d.HasChange("lea.administrator_profile") {
                res["administrator-profile"] = d.Get("lea.administrator_profile")
            }
            if d.HasChange("lea.enabled") {
                res["enabled"] = d.Get("lea.enabled")
            }
            opsecApplication["lea"] = res
        }
    }

    if ok := d.HasChange("one_time_password"); ok {
	       opsecApplication["one-time-password"] = d.Get("one_time_password")
    }

    if d.HasChange("tags") {
        if v, ok := d.GetOk("tags"); ok {
            opsecApplication["tags"] = v.(*schema.Set).List()
        } else {
            oldTags, _ := d.GetChange("tags")
	           opsecApplication["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
        }
    }

    if ok := d.HasChange("color"); ok {
	       opsecApplication["color"] = d.Get("color")
    }

    if ok := d.HasChange("comments"); ok {
	       opsecApplication["comments"] = d.Get("comments")
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
	       opsecApplication["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
	       opsecApplication["ignore-errors"] = v.(bool)
    }

    log.Println("Update OpsecApplication - Map = ", opsecApplication)

    updateOpsecApplicationRes, err := client.ApiCall("set-opsec-application", opsecApplication, client.GetSessionID(), true, false)
    if err != nil || !updateOpsecApplicationRes.Success {
        if updateOpsecApplicationRes.ErrorMsg != "" {
            return fmt.Errorf(updateOpsecApplicationRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }

    return readManagementOpsecApplication(d, m)
}

func deleteManagementOpsecApplication(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    opsecApplicationPayload := map[string]interface{}{
        "uid": d.Id(),
    }

    log.Println("Delete OpsecApplication")

    deleteOpsecApplicationRes, err := client.ApiCall("delete-opsec-application", opsecApplicationPayload , client.GetSessionID(), true, false)
    if err != nil || !deleteOpsecApplicationRes.Success {
        if deleteOpsecApplicationRes.ErrorMsg != "" {
            return fmt.Errorf(deleteOpsecApplicationRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }
    d.SetId("")

    return nil
}

