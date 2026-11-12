package upgraders

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceManagementTimeV0 is the V0 schema where start and end were TypeMap.
func ResourceManagementTimeV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"end": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "End time. Note: Each gateway may interpret this time differently according to its time zone.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"date": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Date in format dd-MMM-yyyy.",
						},
						"time": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Time in format HH:mm.",
						},
					},
				},
			},
			"end_never": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "End never.",
				Default:     true,
			},
			"hours_ranges": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Hours recurrence. Note: Each gateway may interpret this time differently according to its time zone.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Is hour range enabled.",
						},
						"from": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Time in format HH:MM.",
						},
						"index": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Hour range index.",
						},
						"to": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Time in format HH:MM.",
						},
					},
				},
			},
			"start": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Starting time. Note: Each gateway may interpret this time differently according to its time zone.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"date": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Date in format dd-MMM-yyyy.",
						},
						"time": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Time in format HH:mm.",
						},
					},
				},
			},
			"start_now": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Start immediately.",
				Default:     true,
			},
			"recurrence": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Days recurrence.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"days": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Valid on specific days. Multiple options, support range of days in months. Example:[\"1\",\"3\",\"9-20\"].",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"month": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Valid on month. Example: \"1\", \"2\",\"12\",\"Any\".",
						},
						"pattern": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Valid on \"Interval\", \"Daily\", \"Weekly\", \"Monthly\" base.",
						},
						"weekdays": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Valid on weekdays. Example: \"Sun\", \"Mon\"...\"Sat\".",
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

// ResourceManagementTimeStateUpgradeV0 converts start and end from TypeMap to TypeList.
func ResourceManagementTimeStateUpgradeV0(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
	return UpgradeMapsToLists(rawState, "start", "end"), nil
}
