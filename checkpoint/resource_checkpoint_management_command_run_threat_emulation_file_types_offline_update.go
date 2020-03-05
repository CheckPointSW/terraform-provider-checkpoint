package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform/helper/schema"
)

    func resourceManagementRunThreatEmulationFileTypesOfflineUpdate() *schema.Resource {   
        return &schema.Resource{
            Create: createManagementRunThreatEmulationFileTypesOfflineUpdate,
            Read:   readManagementRunThreatEmulationFileTypesOfflineUpdate,
            Delete: deleteManagementRunThreatEmulationFileTypesOfflineUpdate,
            Schema: map[string]*schema.Schema{ 
            "file_path": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: "File path for offline update of Threat Emulation file types, the file path should be on the management machine.",
            },
            "file_raw_data": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: "The contents of a file containing the Threat Emulation file types.",
            },
        },
    }
}

func createManagementRunThreatEmulationFileTypesOfflineUpdate(d *schema.ResourceData, m interface{}) error {
    return readManagementRunThreatEmulationFileTypesOfflineUpdate(d, m)
}

func readManagementRunThreatEmulationFileTypesOfflineUpdate(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    var payload = map[string]interface{}{}
    if v, ok := d.GetOk("file_path"); ok {
        payload["file-path"] = v.(string)
    }

    if v, ok := d.GetOk("file_raw_data"); ok {
        payload["file-raw-data"] = v.(string)
    }

    RunThreatEmulationFileTypesOfflineUpdateRes, _ := client.ApiCall("run-threat-emulation-file-types-offline-update", payload, client.GetSessionID(), true, false)
    if !RunThreatEmulationFileTypesOfflineUpdateRes.Success {
        return fmt.Errorf(RunThreatEmulationFileTypesOfflineUpdateRes.ErrorMsg)
    }

    d.SetId("ff")
    return nil
}

func deleteManagementRunThreatEmulationFileTypesOfflineUpdate(d *schema.ResourceData, m interface{}) error {

    d.SetId("")
    return nil
}

