package upgraders

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceManagementDataCenterObjectV0 is the V0 schema where updated_on_data_center was TypeMap.
func ResourceManagementDataCenterObjectV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"data_center_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of the Data Center Server the object is in.",
			},
			"data_center_uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Unique identifier of the Data Center Server the object is in.",
			},
			"uri": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "URI of the object in the Data Center Server.",
			},
			"uid_in_data_center": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Unique identifier of the object in the Data Center Server.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Override default name on data-center.",
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
			"groups": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of group identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
			"name_in_data_center": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object Name in Data Center",
			},
			"data_center": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Data Center Object",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object name. Must be unique in the domain.",
						},
						"uid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object unique identifier.",
						},
						"automatic_refresh": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "UID.",
						},
						"data_center_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Data Center Type.",
						},
						"properties": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Data Center properties",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Object name. Must be unique in the domain.",
									},
									"value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Object name. Must be unique in the domain.",
									},
								},
							},
						},
					},
				},
			},
			"updated_on_data_center": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Last update time of data center",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"iso_8601": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"posix": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"deleted": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates if the object is inaccessible or deleted on Data Center Server.",
			},
			"type_in_data_center": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object type in Data Center.",
			},
			"additional_properties": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Additional properties on the object.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "",
						},
						"value": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "",
						},
					},
				},
			},
			"wait_for_object_sync": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "When set to true, the provider will wait for object sync with the management server",
			},
		},
	}
}

// ResourceManagementDataCenterObjectStateUpgradeV0 converts updated_on_data_center from TypeMap to TypeList.
func ResourceManagementDataCenterObjectStateUpgradeV0(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
	return UpgradeMapsToLists(rawState, "updated_on_data_center"), nil
}
