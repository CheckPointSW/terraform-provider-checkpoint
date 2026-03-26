package checkpoint

import (
	"github.com/CheckPointSW/terraform-provider-checkpoint/upgraders"
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceManagementInstallSoftwarePackage() *schema.Resource {
	return &schema.Resource{
		Create: createManagementInstallSoftwarePackage,
		Read:   readManagementInstallSoftwarePackage,
		Delete: deleteManagementInstallSoftwarePackage,
		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    upgraders.ResourceManagementCommandInstallSoftwarePackageV0().CoreConfigSchema().ImpliedType(),
				Upgrade: upgraders.ResourceManagementCommandInstallSoftwarePackageStateUpgradeV0,
				Version: 0,
			},
		},
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
				Type:        schema.TypeList,
				MaxItems:    1,
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
			"task_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Command asynchronous task unique identifier.",
			},
		},
	}
}

func createManagementInstallSoftwarePackage(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("name"); ok {
		payload["name"] = v.(string)
	}

	if v, ok := d.GetOk("targets"); ok {
		payload["targets"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("cluster_installation_settings"); ok {

		clusterInstallationSettingsList := v.([]interface{})

		if len(clusterInstallationSettingsList) > 0 {

			clusterInstallationSettingsPayload := make(map[string]interface{})

			if v, ok := d.GetOk("cluster_installation_settings.0.cluster_delay"); ok {
				clusterInstallationSettingsPayload["cluster-delay"] = v.(int)
			}
			if v, ok := d.GetOk("cluster_installation_settings.0.cluster_strategy"); ok {
				clusterInstallationSettingsPayload["cluster-strategy"] = v.(string)
			}
			payload["cluster-installation-settings"] = clusterInstallationSettingsPayload
		}
	}

	if v, ok := d.GetOk("concurrency_limit"); ok {
		payload["concurrency-limit"] = v.(int)
	}

	InstallSoftwarePackageRes, _ := client.ApiCall("install-software-package", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if !InstallSoftwarePackageRes.Success {
		return fmt.Errorf(InstallSoftwarePackageRes.ErrorMsg)
	}

	d.SetId("install-software-package-" + acctest.RandString(10))
	_ = d.Set("task_id", resolveTaskId(InstallSoftwarePackageRes.GetData()))

	return readManagementInstallSoftwarePackage(d, m)
}

func readManagementInstallSoftwarePackage(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementInstallSoftwarePackage(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
