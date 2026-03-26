package upgraders

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceManagementHostV0 is the full schema at version 0, where nat_settings was TypeMap.
// ALL fields must be present so the cty type decoder can preserve them during upgrade.
func ResourceManagementHostV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ipv4_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv6_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"interfaces": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name":            {Type: schema.TypeString, Optional: true},
						"subnet4":         {Type: schema.TypeString, Optional: true},
						"subnet6":         {Type: schema.TypeString, Optional: true},
						"mask_length4":    {Type: schema.TypeInt, Optional: true},
						"mask_length6":    {Type: schema.TypeInt, Optional: true},
						"ignore_warnings": {Type: schema.TypeBool, Optional: true, Default: false},
						"ignore_errors":   {Type: schema.TypeBool, Optional: true, Default: false},
						"color":           {Type: schema.TypeString, Optional: true, Default: "black"},
						"comments":        {Type: schema.TypeString, Optional: true},
					},
				},
			},
			// nat_settings was TypeMap in V0 — all values stored as strings.
			"nat_settings": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"host_servers": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dns_server":  {Type: schema.TypeBool, Optional: true, Default: false},
						"mail_server": {Type: schema.TypeBool, Optional: true, Default: false},
						"web_server":  {Type: schema.TypeBool, Optional: true, Default: false},
						"web_server_config": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"additional_ports":     {Type: schema.TypeSet, Optional: true, Elem: &schema.Schema{Type: schema.TypeString}},
									"application_engines":  {Type: schema.TypeSet, Optional: true, Elem: &schema.Schema{Type: schema.TypeString}},
									"listen_standard_port": {Type: schema.TypeBool, Optional: true, Default: true},
									"operating_system":     {Type: schema.TypeString, Optional: true, Default: "other"},
									"protected_by":         {Type: schema.TypeString, Optional: true, Default: "97aeb368-9aea-11d5-bd16-0090272ccb30"},
								},
							},
						},
					},
				},
			},
			"ignore_warnings": {Type: schema.TypeBool, Optional: true, Default: false},
			"ignore_errors":   {Type: schema.TypeBool, Optional: true, Default: false},
			"color":           {Type: schema.TypeString, Optional: true, Default: "black"},
			"comments":        {Type: schema.TypeString, Optional: true},
			"tags": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

// ResourceManagementHostStateUpgradeV0 converts nat_settings from TypeMap to TypeList with MaxItems:1.
func ResourceManagementHostStateUpgradeV0(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
	return UpgradeMapsToLists(rawState, "nat_settings"), nil
}
