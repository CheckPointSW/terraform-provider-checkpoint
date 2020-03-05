package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform/helper/schema"
)

    func resourceManagementUninstallSoftwarePackage() *schema.Resource {   
        return &schema.Resource{
            Create: createManagementUninstallSoftwarePackage,
            Read:   readManagementUninstallSoftwarePackage,
            Delete: deleteManagementUninstallSoftwarePackage,
            Schema: map[string]*schema.Schema{ 
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: "The name of the software package.",
            },
            "targets": {
                Type:        schema.TypeSet,
                Required:    true,
                ForceNew:    true,
                Description: "On what targets to execute this command. Targets may be identified by their name, or object unique identifier.",
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
            "cluster_installation_settings": {
                Type:        schema.TypeMap,
                Optional:    true,
                Description: "Installation settings for cluster.",
                ForceNew:    true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "cluster_delay": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: "The delay between end of installation on one cluster members and start of installation on the next cluster member.",
                            Default:     0,
                        },
                        "cluster_strategy": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: "The cluster installation strategy.",
                            Default:     "legacy",
                        },
                    },
                },
            },
            "concurrency_limit": {
                Type:        schema.TypeInt,
                Optional:    true,
                ForceNew:    true,
                Description: "The number of targets, on which the same package is installed at the same time.",
            },
        },
    }
}

func createManagementUninstallSoftwarePackage(d *schema.ResourceData, m interface{}) error {
    return readManagementUninstallSoftwarePackage(d, m)
}

func readManagementUninstallSoftwarePackage(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    var payload = map[string]interface{}{}
    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOk("targets"); ok {
        payload["targets"] = v.(*schema.Set).List()
    }

    if _, ok := d.GetOk("cluster_installation_settings"); ok {

        res := make(map[string]interface{})

        if v, ok := d.GetOk("cluster_installation_settings.cluster_delay"); ok {
            res["cluster-delay"] = v
        }
        if v, ok := d.GetOk("cluster_installation_settings.cluster_strategy"); ok {
            res["cluster-strategy"] = v.(string)
        }
        payload["cluster-installation-settings"] = res
    }

    if v, ok := d.GetOk("concurrency_limit"); ok {
        payload["concurrency-limit"] = v.(int)
    }

    UninstallSoftwarePackageRes, _ := client.ApiCall("uninstall-software-package", payload, client.GetSessionID(), true, false)
    if !UninstallSoftwarePackageRes.Success {
        return fmt.Errorf(UninstallSoftwarePackageRes.ErrorMsg)
    }

    d.SetId("ff")
    return nil
}

func deleteManagementUninstallSoftwarePackage(d *schema.ResourceData, m interface{}) error {

    d.SetId("")
    return nil
}

