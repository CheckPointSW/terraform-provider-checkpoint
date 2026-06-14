package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaScheduledSnapshot() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaScheduledSnapshot,
        Read:   readGaiaScheduledSnapshot,
        Update: updateGaiaScheduledSnapshot,
        Delete: deleteGaiaScheduledSnapshot,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "enabled": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `State of the snapshot scheduler`,
            },
            "name_prefix": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Prefix for the snapshots name created by the scheduler`,
            },
            "description": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Description of the scheduled snapshot`,
            },
            "host": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Target host for the snapshots creation`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "target": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Host target type`,
                        },
                        "ip_address": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `IP-Address of the target`,
                        },
                        "upload_path": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Upload path for scp/ftp targets`,
                        },
                        "username": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Username for scp/ftp targets`,
                        },
                        "password": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Sensitive:   true,
                            Description: `Password for scp/ftp targets`,
                        },
                    },
                },
            },
            "recurrence": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Recurrence of the scheduled snapshot`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "pattern": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Recurrence pattern`,
                        },
                        "days": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Description: `Recurrence days`,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "months": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Description: `Recurrence months`,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "weekdays": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Description: `Recurrence weekdays`,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "time": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `Recurrence time`,
                            MaxItems:    1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "hour": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Description: `Time hour`,
                                    },
                                    "minute": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Description: `Time minute`,
                                    },
                                },
                            },
                        },
                    },
                },
            },
            "retention_policy": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Retention-policy for the snapshot scheduler`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "keep_disk_space_above_in_gb": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `Minimum diskspace to keep on the local machine (GB)`,
                        },
                        "min_snapshots_to_keep": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `Minimum snapshots to keep`,
                        },
                        "max_snapshots_to_keep": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `Maximum snapshots to keep`,
                        },
                    },
                },
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
        },
    }
}

func createGaiaScheduledSnapshot(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOkExists("enabled"); ok {
        payload["enabled"] = v.(bool)
    }

    if v, ok := d.GetOk("name_prefix"); ok {
        payload["name-prefix"] = v.(string)
    }

    if v, ok := d.GetOk("description"); ok {
        payload["description"] = v.(string)
    }

    if v := d.Get("host"); len(v.([]interface{})) > 0 {
        _ = v
        hostMap := make(map[string]interface{})
        if v, ok := d.GetOk("host.0.target"); ok {
            hostMap["target"] = v.(string)
        }
        if v, ok := d.GetOk("host.0.ip_address"); ok {
            hostMap["ip-address"] = v.(string)
        }
        if v, ok := d.GetOk("host.0.upload_path"); ok {
            hostMap["upload-path"] = v.(string)
        }
        if v, ok := d.GetOk("host.0.username"); ok {
            hostMap["username"] = v.(string)
        }
        if v, ok := d.GetOk("host.0.password"); ok {
            hostMap["password"] = v.(string)
        }
        if len(hostMap) > 0 {
            payload["host"] = hostMap
        }
    }

    if v := d.Get("recurrence"); len(v.([]interface{})) > 0 {
        _ = v
        recurrenceMap := make(map[string]interface{})
        if v, ok := d.GetOk("recurrence.0.pattern"); ok {
            recurrenceMap["pattern"] = v.(string)
        }
        if v := d.Get("recurrence.0.days"); len(v.(*schema.Set).List()) > 0 {
            recurrenceMap["days"] = v.(*schema.Set).List()
        }
        if v := d.Get("recurrence.0.months"); len(v.(*schema.Set).List()) > 0 {
            recurrenceMap["months"] = v.(*schema.Set).List()
        }
        if v := d.Get("recurrence.0.weekdays"); len(v.(*schema.Set).List()) > 0 {
            recurrenceMap["weekdays"] = v.(*schema.Set).List()
        }
        if v, ok := d.GetOk("recurrence.0.time"); ok {
            _ = v
            timeMap := make(map[string]interface{})
            if v, ok := d.GetOk("recurrence.0.time.0.hour"); ok {
                timeMap["hour"] = v.(int)
            }
            if v, ok := d.GetOk("recurrence.0.time.0.minute"); ok {
                timeMap["minute"] = v.(int)
            }
            if len(timeMap) > 0 {
                recurrenceMap["time"] = timeMap
            }
        }
        if len(recurrenceMap) > 0 {
            payload["recurrence"] = recurrenceMap
        }
    }

    if v := d.Get("retention_policy"); len(v.([]interface{})) > 0 {
        _ = v
        retentionpolicyMap := make(map[string]interface{})
        if v, ok := d.GetOk("retention_policy.0.keep_disk_space_above_in_gb"); ok {
            retentionpolicyMap["keep-disk-space-above-in-GB"] = v.(int)
        }
        if v, ok := d.GetOk("retention_policy.0.min_snapshots_to_keep"); ok {
            retentionpolicyMap["min-snapshots-to-keep"] = v.(int)
        }
        if v, ok := d.GetOk("retention_policy.0.max_snapshots_to_keep"); ok {
            retentionpolicyMap["max-snapshots-to-keep"] = v.(int)
        }
        if len(retentionpolicyMap) > 0 {
            payload["retention-policy"] = retentionpolicyMap
        }
    }

    log.Println("Create ScheduledSnapshot - Map = ", payload)

    addScheduledSnapshotRes, err := client.ApiCallSimple("set-scheduled-snapshot", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addScheduledSnapshotRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addScheduledSnapshotRes.Success {
            errMsg = addScheduledSnapshotRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addScheduledSnapshotRes.GetData()
        }

        debugLogOperation(
            "scheduled-snapshot",        // resource type
            "create",                       // operation
            "set-scheduled-snapshot",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add scheduled-snapshot: %v", err)
    }
    if !addScheduledSnapshotRes.Success {
        if addScheduledSnapshotRes.ErrorMsg != "" {
            return fmt.Errorf(addScheduledSnapshotRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("scheduled-snapshot-" + acctest.RandString(10)))
    return readGaiaScheduledSnapshot(d, m)
}

func readGaiaScheduledSnapshot(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showScheduledSnapshotRes, err := client.ApiCallSimple("show-scheduled-snapshot", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showScheduledSnapshotRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showScheduledSnapshotRes.Success {
            errMsg = showScheduledSnapshotRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showScheduledSnapshotRes.GetData()
        }

        debugLogOperation(
            "scheduled-snapshot",        // resource type
            "read",                       // operation
            "show-scheduled-snapshot",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show scheduled-snapshot: %v", err)
    }
    if !showScheduledSnapshotRes.Success {
        if data := showScheduledSnapshotRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showScheduledSnapshotRes.ErrorMsg)
    }

    scheduledSnapshot := showScheduledSnapshotRes.GetData()

    log.Println("Read ScheduledSnapshot - Show JSON = ", scheduledSnapshot)

    if v, exists := scheduledSnapshot["enabled"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enabled", b)
        } else if s, ok := v.(string); ok {
            d.Set("enabled", s == "true")
        }
    }
    if v, exists := scheduledSnapshot["name-prefix"]; exists {
        d.Set("name_prefix", fmt.Sprintf("%v", v))
    }
    if v, exists := scheduledSnapshot["description"]; exists {
        d.Set("description", fmt.Sprintf("%v", v))
    }
    if v, exists := scheduledSnapshot["host"]; exists {
        d.Set("host", v)
    }
    if v, exists := scheduledSnapshot["recurrence"]; exists {
        d.Set("recurrence", v)
    }
    if v, exists := scheduledSnapshot["retention-policy"]; exists {
        d.Set("retention_policy", v)
    }
    if v, exists := scheduledSnapshot["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaScheduledSnapshot(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOkExists("enabled"); ok {
        payload["enabled"] = v.(bool)
    }

    if v, ok := d.GetOk("name_prefix"); ok {
        payload["name-prefix"] = v.(string)
    }

    if v, ok := d.GetOk("description"); ok {
        payload["description"] = v.(string)
    }

    if v := d.Get("host"); len(v.([]interface{})) > 0 {
        _ = v
        hostMap := make(map[string]interface{})
        if v, ok := d.GetOk("host.0.target"); ok {
            hostMap["target"] = v.(string)
        }
        if v, ok := d.GetOk("host.0.ip_address"); ok {
            hostMap["ip-address"] = v.(string)
        }
        if v, ok := d.GetOk("host.0.upload_path"); ok {
            hostMap["upload-path"] = v.(string)
        }
        if v, ok := d.GetOk("host.0.username"); ok {
            hostMap["username"] = v.(string)
        }
        if v, ok := d.GetOk("host.0.password"); ok {
            hostMap["password"] = v.(string)
        }
        if len(hostMap) > 0 {
            payload["host"] = hostMap
        }
    }

    if v := d.Get("recurrence"); len(v.([]interface{})) > 0 {
        _ = v
        recurrenceMap := make(map[string]interface{})
        if v, ok := d.GetOk("recurrence.0.pattern"); ok {
            recurrenceMap["pattern"] = v.(string)
        }
        if v := d.Get("recurrence.0.days"); len(v.(*schema.Set).List()) > 0 {
            recurrenceMap["days"] = v.(*schema.Set).List()
        }
        if v := d.Get("recurrence.0.months"); len(v.(*schema.Set).List()) > 0 {
            recurrenceMap["months"] = v.(*schema.Set).List()
        }
        if v := d.Get("recurrence.0.weekdays"); len(v.(*schema.Set).List()) > 0 {
            recurrenceMap["weekdays"] = v.(*schema.Set).List()
        }
        if v, ok := d.GetOk("recurrence.0.time"); ok {
            _ = v
            timeMap := make(map[string]interface{})
            if v, ok := d.GetOk("recurrence.0.time.0.hour"); ok {
                timeMap["hour"] = v.(int)
            }
            if v, ok := d.GetOk("recurrence.0.time.0.minute"); ok {
                timeMap["minute"] = v.(int)
            }
            if len(timeMap) > 0 {
                recurrenceMap["time"] = timeMap
            }
        }
        if len(recurrenceMap) > 0 {
            payload["recurrence"] = recurrenceMap
        }
    }

    if v := d.Get("retention_policy"); len(v.([]interface{})) > 0 {
        _ = v
        retentionpolicyMap := make(map[string]interface{})
        if v, ok := d.GetOk("retention_policy.0.keep_disk_space_above_in_gb"); ok {
            retentionpolicyMap["keep-disk-space-above-in-GB"] = v.(int)
        }
        if v, ok := d.GetOk("retention_policy.0.min_snapshots_to_keep"); ok {
            retentionpolicyMap["min-snapshots-to-keep"] = v.(int)
        }
        if v, ok := d.GetOk("retention_policy.0.max_snapshots_to_keep"); ok {
            retentionpolicyMap["max-snapshots-to-keep"] = v.(int)
        }
        if len(retentionpolicyMap) > 0 {
            payload["retention-policy"] = retentionpolicyMap
        }
    }

    setScheduledSnapshotRes, err := client.ApiCallSimple("set-scheduled-snapshot", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setScheduledSnapshotRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setScheduledSnapshotRes.Success {
            errMsg = setScheduledSnapshotRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setScheduledSnapshotRes.GetData()
        }

        debugLogOperation(
            "scheduled-snapshot",        // resource type
            "update",                       // operation
            "set-scheduled-snapshot",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set scheduled-snapshot: %v", err)
    }
    if !setScheduledSnapshotRes.Success {
        return fmt.Errorf(setScheduledSnapshotRes.ErrorMsg)
    }

    return readGaiaScheduledSnapshot(d, m)
}

func deleteGaiaScheduledSnapshot(d *schema.ResourceData, m interface{}) error {


        // No API call - just remove the ID to indicate resource deletion
        d.SetId("")
        return nil
    }

    