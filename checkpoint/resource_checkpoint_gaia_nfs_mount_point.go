package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaNfsMountPoint() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaNfsMountPoint,
        Read:   readGaiaNfsMountPoint,
        Delete: deleteGaiaNfsMountPoint,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:    true,
                Description: "Enable debug logging for this resource.",
            },
            "mount_point": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: `The directory on your root file system from which it will be possible to access the content of the device. Mount points should not have spaces in the names.`,
            },
            "device_path": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: `The device that contains a file system.`,
            },
            "options": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: `Mount options of access to the device. For the list of the supported options, see the Gaia Administration Guide, or the built-in help in the corresponding Gaia Clish command. For explanations about these options, see the Linux man pages 'mount(8)' and 'nfs(5)'.`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
        },
    }
}

func createGaiaNfsMountPoint(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("mount_point"); ok {
        payload["mount-point"] = v.(string)
    }

    if v, ok := d.GetOk("device_path"); ok {
        payload["device-path"] = v.(string)
    }

    if v, ok := d.GetOk("options"); ok {
        payload["options"] = v.(string)
    }

    log.Println("Create NfsMountPoint - Map = ", payload)

    addNfsMountPointRes, err := client.ApiCallSimple("add-nfs-mount-point", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addNfsMountPointRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addNfsMountPointRes.Success {
            errMsg = addNfsMountPointRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addNfsMountPointRes.GetData()
        }

        debugLogOperation(
            "nfs-mount-point",        // resource type
            "create",                       // operation
            "add-nfs-mount-point",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add nfs-mount-point: %v", err)
    }
    if !addNfsMountPointRes.Success {
        if addNfsMountPointRes.ErrorMsg != "" {
            return fmt.Errorf(addNfsMountPointRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("nfs-mount-point-" + acctest.RandString(10)))
    return readGaiaNfsMountPoint(d, m)
}

func readGaiaNfsMountPoint(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("mount_point"); ok {
        payload["mount-point"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showNfsMountPointRes, err := client.ApiCallSimple("show-nfs-mount-point", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showNfsMountPointRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showNfsMountPointRes.Success {
            errMsg = showNfsMountPointRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showNfsMountPointRes.GetData()
        }

        debugLogOperation(
            "nfs-mount-point",        // resource type
            "read",                       // operation
            "show-nfs-mount-point",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show nfs-mount-point: %v", err)
    }
    if !showNfsMountPointRes.Success {
        if data := showNfsMountPointRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showNfsMountPointRes.ErrorMsg)
    }

    nfsMountPoint := showNfsMountPointRes.GetData()

    log.Println("Read NfsMountPoint - Show JSON = ", nfsMountPoint)

    if v, exists := nfsMountPoint["mount-point"]; exists {
        d.Set("mount_point", fmt.Sprintf("%v", v))
    }
    if v, exists := nfsMountPoint["device-path"]; exists {
        d.Set("device_path", fmt.Sprintf("%v", v))
    }
    if v, exists := nfsMountPoint["options"]; exists {
        d.Set("options", fmt.Sprintf("%v", v))
    }
    if v, exists := nfsMountPoint["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func deleteGaiaNfsMountPoint(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("mount_point"); ok {
        payload["mount-point"] = v.(string)
    }

    deleteNfsMountPointRes, err := client.ApiCallSimple("delete-nfs-mount-point", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && deleteNfsMountPointRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !deleteNfsMountPointRes.Success {
            errMsg = deleteNfsMountPointRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = deleteNfsMountPointRes.GetData()
        }

        debugLogOperation(
            "nfs-mount-point",        // resource type
            "delete",                       // operation
            "delete-nfs-mount-point",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to delete nfs-mount-point: %v", err)
    }
    if !deleteNfsMountPointRes.Success {
        return fmt.Errorf(deleteNfsMountPointRes.ErrorMsg)
    }

    d.SetId("")
    return nil
}

