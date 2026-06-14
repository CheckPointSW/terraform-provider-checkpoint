package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaScheduledBackup() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaScheduledBackup,
        Read:   readGaiaScheduledBackup,
        Update: updateGaiaScheduledBackup,
        Delete: deleteGaiaScheduledBackup,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `backup schedule name`,
            },
            "host": {
                Type:        schema.TypeList,
                Required:    true,
                Description: `scheduled backup host`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "target": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `backup host type`,
                        },
                        "ip_address": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `backup host IPv4 address`,
                        },
                        "upload_path": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `backup host upload path`,
                        },
                        "username": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `backup host username`,
                        },
                        "password": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Sensitive:   true,
                            Description: `backup host password`,
                        },
                    },
                },
            },
            "recurrence": {
                Type:        schema.TypeList,
                Required:    true,
                Description: `scheduled backup recurrence`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "pattern": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `backup recurrence pattern`,
                        },
                        "days": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Description: `backup recurrence days`,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "months": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Description: `backup recurrence months`,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "weekdays": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Description: `backup recurrence weekdays`,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                    },
                },
            },
            "time": {
                Type:        schema.TypeList,
                Required:    true,
                Description: `scheduled backup time`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "hour": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `backup time hour`,
                        },
                        "minute": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `backup time minute`,
                        },
                    },
                },
            },
            "retention_policy": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Retention-policy for the backup scheduler, supported from R81.10 and above`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "max_disk_space": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `Maximum diskspace to keep on the local machine (MB)`,
                        },
                        "min_num_of_backups": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `Minimum backups to keep`,
                        },
                        "max_num_of_backups": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `Maximum backups to keep`,
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

func createGaiaScheduledBackup(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
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
        if len(recurrenceMap) > 0 {
            payload["recurrence"] = recurrenceMap
        }
    }

    if v := d.Get("time"); len(v.([]interface{})) > 0 {
        timeMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("time.0.hour"); ok {
            timeMap["hour"] = v.(int)
        }
        if v, ok := d.GetOkExists("time.0.minute"); ok {
            timeMap["minute"] = v.(int)
        }
        if len(timeMap) > 0 {
            payload["time"] = timeMap
        }
    }

    if v := d.Get("retention_policy"); len(v.([]interface{})) > 0 {
        _ = v
        retentionpolicyMap := make(map[string]interface{})
        if v, ok := d.GetOk("retention_policy.0.max_disk_space"); ok {
            retentionpolicyMap["max-disk-space"] = v.(int)
        }
        if v, ok := d.GetOk("retention_policy.0.min_num_of_backups"); ok {
            retentionpolicyMap["min-num-of-backups"] = v.(int)
        }
        if v, ok := d.GetOk("retention_policy.0.max_num_of_backups"); ok {
            retentionpolicyMap["max-num-of-backups"] = v.(int)
        }
        if len(retentionpolicyMap) > 0 {
            payload["retention-policy"] = retentionpolicyMap
        }
    }

    log.Println("Create ScheduledBackup - Map = ", payload)

    addScheduledBackupRes, err := client.ApiCallSimple("add-scheduled-backup", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addScheduledBackupRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addScheduledBackupRes.Success {
            errMsg = addScheduledBackupRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addScheduledBackupRes.GetData()
        }

        debugLogOperation(
            "scheduled-backup",        // resource type
            "create",                       // operation
            "add-scheduled-backup",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add scheduled-backup: %v", err)
    }
    if !addScheduledBackupRes.Success {
        if addScheduledBackupRes.ErrorMsg != "" {
            return fmt.Errorf(addScheduledBackupRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("scheduled-backup-" + acctest.RandString(10)))
    return readGaiaScheduledBackup(d, m)
}

func readGaiaScheduledBackup(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showScheduledBackupRes, err := client.ApiCallSimple("show-scheduled-backup", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showScheduledBackupRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showScheduledBackupRes.Success {
            errMsg = showScheduledBackupRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showScheduledBackupRes.GetData()
        }

        debugLogOperation(
            "scheduled-backup",        // resource type
            "read",                       // operation
            "show-scheduled-backup",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show scheduled-backup: %v", err)
    }
    if !showScheduledBackupRes.Success {
        if data := showScheduledBackupRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showScheduledBackupRes.ErrorMsg)
    }

    scheduledBackup := showScheduledBackupRes.GetData()

    log.Println("Read ScheduledBackup - Show JSON = ", scheduledBackup)

    if v, exists := scheduledBackup["name"]; exists {
        d.Set("name", fmt.Sprintf("%v", v))
    }
    if v, exists := scheduledBackup["host"]; exists {
        d.Set("host", v)
    }
    if v, exists := scheduledBackup["recurrence"]; exists {
        d.Set("recurrence", v)
    }
    if v, exists := scheduledBackup["time"]; exists {
        d.Set("time", v)
    }
    if v, exists := scheduledBackup["retention-policy"]; exists {
        d.Set("retention_policy", v)
    }
    if v, exists := scheduledBackup["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaScheduledBackup(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
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
        if len(recurrenceMap) > 0 {
            payload["recurrence"] = recurrenceMap
        }
    }

    if v := d.Get("time"); len(v.([]interface{})) > 0 {
        timeMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("time.0.hour"); ok {
            timeMap["hour"] = v.(int)
        }
        if v, ok := d.GetOkExists("time.0.minute"); ok {
            timeMap["minute"] = v.(int)
        }
        if len(timeMap) > 0 {
            payload["time"] = timeMap
        }
    }

    if v := d.Get("retention_policy"); len(v.([]interface{})) > 0 {
        _ = v
        retentionpolicyMap := make(map[string]interface{})
        if v, ok := d.GetOk("retention_policy.0.max_disk_space"); ok {
            retentionpolicyMap["max-disk-space"] = v.(int)
        }
        if v, ok := d.GetOk("retention_policy.0.min_num_of_backups"); ok {
            retentionpolicyMap["min-num-of-backups"] = v.(int)
        }
        if v, ok := d.GetOk("retention_policy.0.max_num_of_backups"); ok {
            retentionpolicyMap["max-num-of-backups"] = v.(int)
        }
        if len(retentionpolicyMap) > 0 {
            payload["retention-policy"] = retentionpolicyMap
        }
    }

    setScheduledBackupRes, err := client.ApiCallSimple("set-scheduled-backup", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setScheduledBackupRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setScheduledBackupRes.Success {
            errMsg = setScheduledBackupRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setScheduledBackupRes.GetData()
        }

        debugLogOperation(
            "scheduled-backup",        // resource type
            "update",                       // operation
            "set-scheduled-backup",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set scheduled-backup: %v", err)
    }
    if !setScheduledBackupRes.Success {
        return fmt.Errorf(setScheduledBackupRes.ErrorMsg)
    }

    return readGaiaScheduledBackup(d, m)
}

func deleteGaiaScheduledBackup(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    deleteScheduledBackupRes, err := client.ApiCallSimple("delete-scheduled-backup", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && deleteScheduledBackupRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !deleteScheduledBackupRes.Success {
            errMsg = deleteScheduledBackupRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = deleteScheduledBackupRes.GetData()
        }

        debugLogOperation(
            "scheduled-backup",        // resource type
            "delete",                       // operation
            "delete-scheduled-backup",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to delete scheduled-backup: %v", err)
    }
    if !deleteScheduledBackupRes.Success {
        return fmt.Errorf(deleteScheduledBackupRes.ErrorMsg)
    }

    d.SetId("")
    return nil
}

