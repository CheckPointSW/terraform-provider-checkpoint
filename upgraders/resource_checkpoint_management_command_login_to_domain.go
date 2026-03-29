package upgraders

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceManagementCommandLoginToDomainV0 is the V0 schema where last_login_was_at and login_message were TypeMap.
func ResourceManagementCommandLoginToDomainV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"domain": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Domain identified by the name or UID.",
			},
			"continue_last_session": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Default:     false,
				Description: "When 'continue-last-session' is set to 'True', the new session would continue where the last session was stopped. This option is available when the administrator has only one session that can be continued. If there is more than one session, see 'switch-session' API.",
			},
			"read_only": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Default:     false,
				Description: "Login with Read Only permissions. This parameter is not considered in case continue-last-session is true.",
			},
			"sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Session unique identifier. Enter this session unique identifier in the 'X-chkp-sid' header of each request.",
			},
			"api_server_version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "API Server version.",
			},
			"disk_space_message": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Information about the available disk space on the management server.",
			},
			"last_login_was_at": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Timestamp when administrator last accessed the management server.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"iso_8601": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Date and time represented in international ISO 8601 format.",
						},
						"posix": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.",
						},
					},
				},
			},
			"login_message": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Login message.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"header": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Message header.",
						},
						"message": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Message content.",
						},
					},
				},
			},
			"session_timeout": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Session expiration timeout in seconds.",
			},
			"standby": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "True if this management server is in the standby mode.",
			},
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Session object unique identifier. This identifier may be used in the discard API to discard changes that were made in this session, when administrator is working from another session, or in the 'switch-session' API.",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL that was used to reach the API server.",
			},
		},
	}
}

// ResourceManagementCommandLoginToDomainStateUpgradeV0 converts last_login_was_at and login_message from TypeMap to TypeList.
func ResourceManagementCommandLoginToDomainStateUpgradeV0(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
	return UpgradeMapsToLists(rawState, "last_login_was_at", "login_message"), nil
}
