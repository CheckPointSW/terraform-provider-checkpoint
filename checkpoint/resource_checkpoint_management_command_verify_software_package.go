package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceManagementVerifySoftwarePackage() *schema.Resource {
	return &schema.Resource{
		Create: createManagementVerifySoftwarePackage,
		Read:   readManagementVerifySoftwarePackage,
		Delete: deleteManagementVerifySoftwarePackage,
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
			"operation_context": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The operation can be: 'install' (default) or 'uninstall'.",
			},
			"download_package_from": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Where is the package located.",
			},
			"download_package": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Should the package be downloaded before verification.",
			},
			"cluster_installation_settings": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				ForceNew:    true,
				Description: "Installation settings for cluster.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cluster_delay": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The delay between end of installation on one cluster members and start of installation on the next cluster member.",
						},
						"cluster_strategy": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The cluster installation strategy. all-members - Install the package on all members in the cluster non-active-members-and-failover - In the High Avail...",
						},
					},
				},
			},
		},
	}
}

func createManagementVerifySoftwarePackage(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("name"); ok {
		payload["name"] = v.(string)
	}

	if v, ok := d.GetOk("targets"); ok {
		payload["targets"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("concurrency_limit"); ok {
		payload["concurrency-limit"] = v.(int)
	}

	if v, ok := d.GetOk("operation_context"); ok {
		payload["operation-context"] = v.(string)
	}

	if v, ok := d.GetOk("download_package_from"); ok {
		payload["download-package-from"] = v.(string)
	}

	if v, ok := d.GetOkExists("download_package"); ok {
		payload["download-package"] = v.(bool)
	}

	if _, ok := d.GetOk("cluster_installation_settings"); ok {
		clusterInstallationSettingsPayload := make(map[string]interface{})
		if v, ok := d.GetOk("cluster_installation_settings.0.cluster_delay"); ok {
			clusterInstallationSettingsPayload["cluster-delay"] = v.(int)
		}
		if v, ok := d.GetOk("cluster_installation_settings.0.cluster_strategy"); ok {
			clusterInstallationSettingsPayload["cluster-strategy"] = v.(string)
		}
		payload["cluster-installation-settings"] = clusterInstallationSettingsPayload
	}

	VerifySoftwarePackageRes, _ := client.ApiCall("verify-software-package", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if !VerifySoftwarePackageRes.Success {
		return fmt.Errorf("%s", VerifySoftwarePackageRes.ErrorMsg)
	}

	d.SetId("verify-software-package-" + acctest.RandString(10))
	_ = d.Set("task_id", resolveTaskId(VerifySoftwarePackageRes.GetData()))

	return readManagementVerifySoftwarePackage(d, m)
}

func readManagementVerifySoftwarePackage(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementVerifySoftwarePackage(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
