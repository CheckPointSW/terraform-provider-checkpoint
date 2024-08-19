---
layout: "checkpoint"
page_title: "checkpoint_management_app_control_advanced_settings"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-app-control-advanced-settings"
description: |-
Use this data source to get information on an existing Check Point Set App Control Advanced Settings.
---

# checkpoint_management_app_control_advanced_settings

Use this data source to get information on an existing Check Point Set App Control Advanced Settings.

## Example Usage


```hcl
resource "checkpoint_management_set_app_control_advanced_settings" "example" {
  url_filtering_settings = {
    categorize_cached_and_translated_pages = "true"
    categorize_https_websites = "false"
    enforce_safe_search ="true"
  }
  custom_categorization_settings = {
    social_network_widgets_mode = "hold"
    url_filtering_mode = "background"
  }
  web_browsing_services = ["https","AH"]
  match_application_on_any_port = "false"
}

data "checkpoint_management_app_control_advanced_settings" "data" {

  depends_on = [checkpoint_management_app_control_advanced_settings.example]

}
```

## Argument Reference

The following arguments are supported:
* `uid` -  Object unique identifier.
* `internal_error_fail_mode` -  In case of internal system error, allow or block all connections.<br>This property is not available in the Global domain of an MDS machine.
* `url_filtering_settings` - In this section user can enable  URL Filtering features.<br>This property is not available in the Global domain of an MDS machine.url_filtering_settings blocks are documented below.
* `web_browsing_services` -  Web browsing services are the services that match a Web-based custom Application/Site.web_browsing_services blocks are documented below.
* `match_application_on_any_port` -  Match Web application on 'Any' port when used in Block rule - By default this is set to true. and so applications are matched on all services when used in a Block rule.
* `enable_web_browsing` -  If you do not enable URL Filtering on the Security Gateway, you can use a generic Web browser application called Web Browsing in the rule.<br>This application includes all HTTP traffic that is not a defined application
  Application and URL Filtering assigns Web Browsing as the default application for all HTTP traffic that does not match an application in the Application and URL Filtering Database.<br>This property is not available in the Global domain of an MDS machine.
* `httpi_non_standard_ports` -  Enable HTTP inspection on non standard ports for application and URL filtering.<br>This property is not available in the Global domain of an MDS machine.
* `block_request_when_web_service_is_unavailable` - Block requests when the web service is unavailable.
  <br>When selected, requests are blocked when there is no connectivity to the Check Point Online Web Service.<br>When cleared, requests are allowed when there is no connectivity.<br>This property is not available in the Global domain of an MDS machine.
* `website_categorization_mode` - Hold - Requests are blocked until categorization is complete.<br>Background - Requests are allowed until categorization is complete.<br>Custom - configure different settings depending on the service -Lets you set different modes for URL Filtering and Social Networking Widgets.<br>This property is not available in the Global domain of an MDS machine.
* `custom_categorization_settings` - Website categorization mode - select the mode that is used for website categorization.<br>This property is not available in the Global domain of an MDS machine.custom_categorization_settings blocks are documented below.
* `categorize_social_network_widgets` - When selected, the Security Gateway connects to the Check Point Online Web Service to identify social networking widgets that it does not recognize.<br>When cleared or there is no connectivity between the Security Gateway and the Check Point Online Web, the unknown widget is treated as Web Browsing traffic.<br>This property is not available in the Global domain of an MDS machine.
* `domain_level_permission` -  Allows the editing of applications, categories, and services. This property is used only in the Global Domain of an MDS machine.


`url_filtering_settings` supports the following:

* `categorize_https_websites` - This option lets Application and URL Filtering assign categories to HTTPS sites without activating HTTPS inspection. It assigns a site category based on its domain name and whether the site has a valid certificate. If the server certificate is:<br> Trusted - Application and URL Filtering gets the domain name from the certificate and uses it to categorize the site.<br>Not Trusted - Application and URL Filtering assigns a category based on the IP address.<br>This property is not available in the Global domain of an MDS machine.
* `enforce_safe_search` - Select this option to require use of the safe search feature in search engines. When activated, the URL Filtering Policy uses the strictest available safe search option for the specified search engine.<br>This option overrides user specified search engine options to block offensive material in search results.<br>This property is not available in the Global domain of an MDS machine.
* `categorize_cached_and_translated_pages` - Select this option to assign categories to cached search engine results and translated pages.<br>When this option is selected, Application and URL Filtering assigns categories based on the original Web site instead of the 'search engine pages' category.<br>This property is not available in the Global domain of an MDS machine.


`custom_categorization_settings` supports the following:

* `url_filtering_mode` - Hold - Requests are blocked until categorization is complete.<br>Background - Requests are allowed until categorization is complete.<br>This property is not available in the Global domain of an MDS machine.
* `social_network_widgets_mode` - Hold - Requests are blocked until categorization is complete.<br>Background - Requests are allowed until categorization is complete.<br>This property is not available in the Global domain of an MDS machine.


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.  

