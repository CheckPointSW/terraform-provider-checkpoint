package checkpoint

        import (
            "fmt"
            checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
            "github.com/hashicorp/terraform/helper/schema"
            "strconv"
        )

            func resourceManagementSetThreatProtection() *schema.Resource {   
        return &schema.Resource{
            Create: createManagementSetThreatProtection,
            Read:   readManagementSetThreatProtection,
            Delete: deleteManagementSetThreatProtection,
            Schema: map[string]*schema.Schema{ 
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: "Object name.",
            },
            "comments": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: "Protection comments.",
            },
            "follow_up": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:    true,
                Description: "Tag the protection with pre-defined follow-up flag.",
            },
            "overrides": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: "Overrides per profile for this protection<br> Note: Remove override for Core protections removes only the actions override. Remove override for Threat Cloud protections removes the action, track and packet captures.",
                ForceNew:    true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "action": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: "Protection action.",
                        },
                        "profile": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: "Profile name.",
                        },
                        "capture_packets": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: "Capture packets.",
                        },
                        "track": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: "Tracking method for protection.",
                        },
                    },
                },
            },
        },
    }
}

func createManagementSetThreatProtection(d *schema.ResourceData, m interface{}) error {
    return readManagementSetThreatProtection(d, m)
}

func readManagementSetThreatProtection(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    var payload = map[string]interface{}{}
    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOk("comments"); ok {
        payload["comments"] = v.(string)
    }

    if v, ok := d.GetOkExists("follow_up"); ok {
        payload["follow-up"] = v.(bool)
    }

    if v, ok := d.GetOk("overrides"); ok {

        overridesList := v.([]interface{})

        if len(overridesList) > 0 {

            var overridesPayload []map[string]interface{}

            for i := range overridesList {

                Payload := make(map[string]interface{})

                if v, ok := d.GetOk("overrides." + strconv.Itoa(i) + ".action"); ok {
                    Payload["action"] = v.(string)
                }
                if v, ok := d.GetOk("overrides." + strconv.Itoa(i) + ".profile"); ok {
                    Payload["profile"] = v.(string)
                }
                if v, ok := d.GetOk("overrides." + strconv.Itoa(i) + ".capture_packets"); ok {
                    Payload["capture-packets"] = v.(bool)
                }
                if v, ok := d.GetOk("overrides." + strconv.Itoa(i) + ".track"); ok {
                    Payload["track"] = v.(string)
                }
                overridesPayload = append(overridesPayload, Payload)
            }
            payload["overrides"] = overridesPayload
        }
    }

    SetThreatProtectionRes, _ := client.ApiCall("set-threat-protection", payload, client.GetSessionID(), true, false)
    if !SetThreatProtectionRes.Success {
        return fmt.Errorf(SetThreatProtectionRes.ErrorMsg)
    }

    d.SetId("ff")
    return nil
}

func deleteManagementSetThreatProtection(d *schema.ResourceData, m interface{}) error {

    d.SetId("")
    return nil
}

