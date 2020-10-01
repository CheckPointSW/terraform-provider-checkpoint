package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementInstallPolicy() *schema.Resource {
	return &schema.Resource{
		Create: createManagementInstallPolicy,
		Read:   readManagementInstallPolicy,
		Delete: deleteManagementInstallPolicy,
		Schema: map[string]*schema.Schema{
			"policy_package": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the Policy Package to be installed.",
			},
			"targets": {
				Type:        schema.TypeSet,
				Required:    true,
				ForceNew:    true,
				Description: "On what targets to execute this command. Targets may be identified by their name, or object unique identifier.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"access": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Set to be true in order to install the Access Control policy. By default, the value is true if Access Control policy is enabled on the input policy package, otherwise false.",
			},
			"desktop_security": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Set to be true in order to install the Desktop Security policy. By default, the value is true if desktop security policy is enabled on the input policy package, otherwise false.",
			},
			"qos": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Set to be true in order to install the QoS policy. By default, the value is true if Quality-of-Service policy is enabled on the input policy package, otherwise false.",
			},
			"threat_prevention": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Set to be true in order to install the Threat Prevention policy. By default, the value is true if Threat Prevention policy is enabled on the input policy package, otherwise false.",
			},
			"install_on_all_cluster_members_or_fail": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Relevant for the gateway clusters. If true, the policy is installed on all the cluster members. If the installation on a cluster member fails, don't install on that cluster.",
				Default:     true,
			},
			"prepare_only": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "If true, prepares the policy for the installation, but doesn't install it on an installation target.",
				Default:     false,
			},
			"revision": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The UID of the revision of the policy to install.",
			},
			"task_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Command asynchronous task unique identifier.",
			},
		},
	}
}

func createManagementInstallPolicy(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = make(map[string]interface{})

	if v, ok := d.GetOk("policy_package"); ok {
		payload["policy-package"] = v.(string)
	}
	if v, ok := d.GetOk("targets"); ok {
		payload["targets"] = v.(*schema.Set).List()
	}
	if v, ok := d.GetOk("access"); ok {
		payload["access"] = v.(bool)
	}
	if v, ok := d.GetOk("desktop_security"); ok {
		payload["desktop-security"] = v.(bool)
	}
	if v, ok := d.GetOk("qos"); ok {
		payload["qos"] = v.(bool)
	}
	if v, ok := d.GetOk("threat_prevention"); ok {
		payload["threat-prevention"] = v.(bool)
	}
	if v, ok := d.GetOk("install_on_all_cluster_members_or_fail"); ok {
		payload["install-on-all-cluster-members-or-fail"] = v.(bool)
	}
	if v, ok := d.GetOk("revision"); ok {
		payload["revision"] = v.(bool)
	}

	installPolicyRes, _ := client.ApiCall("install-policy", payload, client.GetSessionID(), true, false)
	if !installPolicyRes.Success {
		return fmt.Errorf(installPolicyRes.ErrorMsg)
	}

	d.SetId("install-policy-" + acctest.RandString(10))
	_ = d.Set("task_id", resolveTaskId(installPolicyRes.GetData()))
	return readManagementInstallPolicy(d, m)
}

func readManagementInstallPolicy(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementInstallPolicy(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
