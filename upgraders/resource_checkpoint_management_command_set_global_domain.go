package upgraders

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceManagementCommandSetGlobalDomainV0 is the V0 schema where servers was TypeMap.
func ResourceManagementCommandSetGlobalDomainV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Object name.",
			},
			"servers": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Multi Domain Servers. When the field is provided, 'set-global-domain' command is executed asynchronously.",
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"add": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Adds to collection of values",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				ForceNew:    true,
				Description: "Collection of tag identifiers. Note: The list of tags can not be modified in a singlecommand together with the domain servers. To modify tags, please use the separate 'set-global-domain' command, without providing the list of domain servers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"color": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Color of the object. Should be one of existing colors.",
			},
			"comments": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Comments string.",
			},
			"ignore_warnings": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Apply changes ignoring warnings.",
			},
			"ignore_errors": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.",
			},
			"tasks": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Command asynchronous task unique identifiers",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

// ResourceManagementCommandSetGlobalDomainStateUpgradeV0 converts servers from TypeMap to TypeList.
func ResourceManagementCommandSetGlobalDomainStateUpgradeV0(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
	return UpgradeMapsToLists(rawState, "servers"), nil
}
