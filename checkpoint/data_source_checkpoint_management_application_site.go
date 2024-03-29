package checkpoint

import (
	"fmt"
	"log"

	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceManagementApplicationSite() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementApplicationSiteRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object unique identifier.",
			},
			"additional_categories": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Used to configure or edit the additional categories of a custom application / site used in the Application and URL Filtering or Threat Prevention.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A description for the application.",
			},
			"primary_category": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Each application is assigned to one primary category based on its most defining aspect.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"url_list": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "URLs that determine this particular application.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"application_signature": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Application signature generated by <a href=\"https://supportcenter.checkpoint.com/supportcenter/portal?eventSubmit_doGoviewsolutiondetails=&solutionid=sk103051\">Signature Tool</a>.",
			},
			"urls_defined_as_regular_expression": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "States whether the URL is defined as a Regular Expression or not.",
			},
			"color": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Color of the object. Should be one of existing colors.",
			},
			"comments": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Comments string.",
			},
			"groups": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of group identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceManagementApplicationSiteRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showApplicationSiteRes, err := client.ApiCall("show-application-site", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showApplicationSiteRes.Success {
		return fmt.Errorf(showApplicationSiteRes.ErrorMsg)
	}

	applicationSite := showApplicationSiteRes.GetData()

	log.Println("Read ApplicationSite - Show JSON = ", applicationSite)

	if v := applicationSite["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := applicationSite["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if applicationSite["additional-categories"] != nil {
		additionalCategoriesJson, ok := applicationSite["additional-categories"].([]interface{})
		if ok {
			additionalCategoriesIds := make([]string, 0, len(additionalCategoriesJson))
			if len(additionalCategoriesJson) > 0 {
				for _, additional_categories := range additionalCategoriesJson {
					additional_categories := additional_categories.(string)
					additionalCategoriesIds = append(additionalCategoriesIds, additional_categories)
				}
			}
			_ = d.Set("additional_categories", additionalCategoriesIds)
		}
	} else {
		_ = d.Set("additional_categories", nil)
	}

	if v := applicationSite["description"]; v != nil {
		_ = d.Set("description", v)
	}

	if v := applicationSite["primary-category"]; v != nil {
		_ = d.Set("primary_category", v)
	}

	if applicationSite["tags"] != nil {
		tagsJson, ok := applicationSite["tags"].([]interface{})
		if ok {
			tagsIds := make([]string, 0)
			if len(tagsJson) > 0 {
				for _, tags := range tagsJson {
					tags := tags.(map[string]interface{})
					tagsIds = append(tagsIds, tags["name"].(string))
				}
			}
			_ = d.Set("tags", tagsIds)
		}
	} else {
		_ = d.Set("tags", nil)
	}

	if applicationSite["url-list"] != nil {
		urlListJson, ok := applicationSite["url-list"].([]interface{})
		if ok {
			urlListIds := make([]string, 0)
			if len(urlListJson) > 0 {
				for _, url_list := range urlListJson {
					url_list := url_list.(string)
					urlListIds = append(urlListIds, url_list)
				}
			}
			_ = d.Set("url_list", urlListIds)
		}
	} else {
		_ = d.Set("url_list", nil)
	}

	if v := applicationSite["application-signature"]; v != nil {
		_ = d.Set("application_signature", v)
	}

	if v := applicationSite["urls-defined-as-regular-expression"]; v != nil {
		_ = d.Set("urls_defined_as_regular_expression", v)
	}

	if v := applicationSite["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := applicationSite["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if applicationSite["groups"] != nil {
		groupsJson, ok := applicationSite["groups"].([]interface{})
		if ok {
			groupsIds := make([]string, 0)
			if len(groupsJson) > 0 {
				for _, groups := range groupsJson {
					groups := groups.(map[string]interface{})
					groupsIds = append(groupsIds, groups["name"].(string))
				}
			}
			_ = d.Set("groups", groupsIds)
		}
	} else {
		_ = d.Set("groups", nil)
	}

	return nil
}
