package upgraders

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceManagementCommandSetAppControlAdvancedSettingsV0 is the V0 schema where
// url_filtering_settings and custom_categorization_settings were TypeMap.
func ResourceManagementCommandSetAppControlAdvancedSettingsV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object unique identifier.",
			},
			"internal_error_fail_mode": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "In case of internal system error, allow or block all connections.<br>This property is not available in the Global domain of an MDS machine.",
			},
			"url_filtering_settings": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "In this section user can enable  URL Filtering features.<br>This property is not available in the Global domain of an MDS machine.",
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"categorize_https_websites": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "This option lets Application and URL Filtering assign categories to HTTPS sites without activating HTTPS inspection. It assigns a site category based on its domain name and whether the site has a valid certificate. If the server certificate is:<br> Trusted - Application and URL Filtering gets the domain name from the certificate and uses it to categorize the site.<br>Not Trusted - Application and URL Filtering assigns a category based on the IP address.<br>This property is not available in the Global domain of an MDS machine.",
						},
						"enforce_safe_search": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Select this option to require use of the safe search feature in search engines. When activated, the URL Filtering Policy uses the strictest available safe search option for the specified search engine.<br>This option overrides user specified search engine options to block offensive material in search results.<br>This property is not available in the Global domain of an MDS machine.",
						},
						"categorize_cached_and_translated_pages": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Select this option to assign categories to cached search engine results and translated pages.<br>When this option is selected, Application and URL Filtering assigns categories based on the original Web site instead of the 'search engine pages' category.<br>This property is not available in the Global domain of an MDS machine.",
						},
					},
				},
			},
			"web_browsing_services": {
				Type:        schema.TypeSet,
				Optional:    true,
				ForceNew:    true,
				Description: "Web browsing services are the services that match a Web-based custom Application/Site.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"match_application_on_any_port": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Match Web application on 'Any' port when used in Block rule - By default this is set to true. and so applications are matched on all services when used in a Block rule.",
			},
			"enable_web_browsing": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "If you do not enable URL Filtering on the Security Gateway, you can use a generic Web browser application called Web Browsing in the rule.<br>This application includes all HTTP traffic that is not a defined application Application and URL Filtering assigns Web Browsing as the default application for all HTTP traffic that does not match an application in the Application and URL Filtering Database.<br>This property is not available in the Global domain of an MDS machine.",
			},
			"httpi_non_standard_ports": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Enable HTTP inspection on non standard ports for application and URL filtering.<br>This property is not available in the Global domain of an MDS machine.",
			},
			"block_request_when_web_service_is_unavailable": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Block requests when the web service is unavailable. <br>When selected, requests are blocked when there is no connectivity to the Check Point Online Web Service.<br>When cleared, requests are allowed when there is no connectivity.<br>This property is not available in the Global domain of an MDS machine.",
			},
			"website_categorization_mode": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Hold - Requests are blocked until categorization is complete.<br>Background - Requests are allowed until categorization is complete.<br>Custom - configure different settings depending on the service -Lets you set different modes for URL Filtering and Social Networking Widgets.<br>This property is not available in the Global domain of an MDS machine.",
			},
			"custom_categorization_settings": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Website categorization mode - select the mode that is used for website categorization.<br>This property is not available in the Global domain of an MDS machine.",
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"url_filtering_mode": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Hold - Requests are blocked until categorization is complete.<br>Background - Requests are allowed until categorization is complete.<br>This property is not available in the Global domain of an MDS machine.",
						},
						"social_network_widgets_mode": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Hold - Requests are blocked until categorization is complete.<br>Background - Requests are allowed until categorization is complete.<br>This property is not available in the Global domain of an MDS machine.",
						},
					},
				},
			},
			"categorize_social_network_widgets": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "When selected, the Security Gateway connects to the Check Point Online Web Service to identify social networking widgets that it does not recognize.<br>When cleared or there is no connectivity between the Security Gateway and the Check Point Online Web, the unknown widget is treated as Web Browsing traffic.<br>This property is not available in the Global domain of an MDS machine.",
			},
			"domain_level_permission": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Allows the editing of applications, categories, and services. This property is used only in the Global Domain of an MDS machine.",
			},
		},
	}
}

// ResourceManagementCommandSetAppControlAdvancedSettingsStateUpgradeV0 converts the TypeMap fields to TypeList.
func ResourceManagementCommandSetAppControlAdvancedSettingsStateUpgradeV0(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
	return UpgradeMapsToLists(rawState, "url_filtering_settings", "custom_categorization_settings"), nil
}
