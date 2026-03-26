package upgraders

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceManagementUserV0 is the V0 schema where allowed_locations and encryption were TypeMap.
func ResourceManagementUserV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"email": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User email.",
			},
			"expiration_date": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Expiration date in format: yyyy-MM-dd.",
			},
			"phone_number": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User phone number.",
			},
			"authentication_method": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Authentication method.",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "Checkpoint password authentication method identified by the name or UID. Must be set when \"authentication-method\" was selected to be \"Check Point Password\".",
			},
			"radius_server": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "RADIUS server object identified by the name or UID. Must be set when \"authentication-method\" was selected to be \"RADIUS\".",
			},
			"tacacs_server": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "TACACS server object identified by the name or UID. Must be set when \"authentication-method\" was selected to be \"TACACS\".",
			},
			"connect_on_days": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Days users allow to connect.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"connect_daily": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Connect every day.",
			},
			"from_hour": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Allow users connect from hour.",
			},
			"to_hour": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Allow users connect until hour.",
			},
			"allowed_locations": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "User allowed locations.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"destinations": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Collection of allowed destination locations name or uid.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"sources": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Collection of allowed source locations name or uid.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"encryption": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "User encryption.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_ike": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable IKE encryption for users.",
						},
						"enable_public_key": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable IKE public key.",
						},
						"enable_shared_secret": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable IKE shared secret.",
						},
						"shared_secret": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IKE shared secret.",
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
			"template": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User template name or UID.",
				Default:     "Default",
			},
		},
	}
}

// ResourceManagementUserStateUpgradeV0 converts allowed_locations and encryption from TypeMap to TypeList.
func ResourceManagementUserStateUpgradeV0(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
	return UpgradeMapsToLists(rawState, "allowed_locations", "encryption"), nil
}
