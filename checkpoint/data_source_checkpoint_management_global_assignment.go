package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementGlobalAssignment() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementGlobalAssignmentRead,
		Schema: map[string]*schema.Schema{
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object unique identifier.",
			},
			"dependent_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"global_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object name. Must be unique in the domain.",
			},
			"assignment_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"assignment_up_to_date": {
				Type:     schema.TypeMap,
				Computed: true,
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
			"global_access_policy": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Global domain access policy that is assigned to a dependent domain.",
			},
			"global_threat_prevention_policy": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Global domain threat prevention policy that is assigned to a dependent domain.",
			},
			"manage_protection_actions": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"tags": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of tag objects identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
		},
	}
}

func dataSourceManagementGlobalAssignmentRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	uid := d.Get("uid").(string)
	dependentDomain := d.Get("dependent_domain")
	globalDomain := d.Get("global_domain")

	payload := make(map[string]interface{})

	if uid != "" {
		payload["uid"] = uid
	} else if dependentDomain != "" {
		payload["dependent-domain"] = dependentDomain
	}
	if globalDomain != "" {
		payload["global-domain"] = globalDomain
	}

	showGlobalAssignmentRes, err := client.ApiCall("show-global-assignment", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showGlobalAssignmentRes.Success {
		return fmt.Errorf(showGlobalAssignmentRes.ErrorMsg)
	}

	globalAssignment := showGlobalAssignmentRes.GetData()

	log.Println("Read Global Assignment - Show JSON = ", globalAssignment)

	if v := globalAssignment["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := globalAssignment["assignment-status"]; v != nil {
		_ = d.Set("assignment_status", v)
	}

	if globalAssignment["assigment-up-to-date"] != nil {
		assignmentUpToDateMap := globalAssignment["assigment-up-to-date"].(map[string]interface{})
		assignmentUpToDateMapToReturn := make(map[string]interface{})

		if v, _ := assignmentUpToDateMap["iso-8601"]; v != nil {
			assignmentUpToDateMapToReturn["iso_8601"] = v
		}
		if v, _ := assignmentUpToDateMap["posix"]; v != nil {
			assignmentUpToDateMapToReturn["posix"] = v
		}

		_ = d.Set("assignment_up_to_date", assignmentUpToDateMapToReturn)
	}

	if globalAssignment["dependent-domain"] != nil {
		dependentDomainMap := globalAssignment["dependent-domain"].(map[string]interface{})
		_ = d.Set("dependent_domain", dependentDomainMap["name"])
	}

	if v := globalAssignment["global-access-policy"]; v != nil {
		_ = d.Set("global_access_policy", v)
	}

	if globalAssignment["global-domain"] != nil {
		globalDomainMap := globalAssignment["global-domain"].(map[string]interface{})
		_ = d.Set("global_domain", globalDomainMap["name"])
	}

	if v := globalAssignment["global-threat-prevention-policy"]; v != nil {
		_ = d.Set("global_threat_prevention_policy", v)
	}

	if v := globalAssignment["manage-protection-actions"]; v != nil {
		_ = d.Set("manage_protection_actions", v)
	}

	if globalAssignment["tags"] != nil {
		tagsJson := globalAssignment["tags"].([]interface{})
		var tagsIds = make([]string, 0)
		if len(tagsJson) > 0 {
			// Create slice of tag names
			for _, tag := range tagsJson {
				tag := tag.(map[string]interface{})
				tagsIds = append(tagsIds, tag["name"].(string))
			}
		}
		_ = d.Set("tags", tagsIds)
	} else {
		_ = d.Set("tags", nil)
	}

	if v := globalAssignment["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := globalAssignment["color"]; v != nil {
		_ = d.Set("color", v)
	}

	return nil
}
