package upgraders

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceManagementOpsecApplicationV0 is the V0 schema where cpmi and lea were TypeMap.
func ResourceManagementOpsecApplicationV0() *schema.Resource {
	return &schema.Resource{
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
				Sensitive:   true,
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

// ResourceManagementOpsecApplicationStateUpgradeV0 converts cpmi and lea from TypeMap to TypeList.
func ResourceManagementOpsecApplicationStateUpgradeV0(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
	return UpgradeMapsToLists(rawState, "cpmi", "lea"), nil
}
