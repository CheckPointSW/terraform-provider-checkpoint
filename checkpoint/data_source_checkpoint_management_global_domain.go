package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementGlobalDomain() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementGlobalDomainRead,
		Schema: map[string]*schema.Schema{
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object unique identifier.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object type.",
			},
			"domain_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "N/A",
			},
			"global_domain_assignments": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "N/A",
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
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object type.",
						},
						"assignment_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "N/A",
						},
						"assignment_up_to_date": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "The time when the assignment was assigned.",
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
						"dependent_domain": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Dependent domain. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.",
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
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Object type.",
									},
									"color": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Color of the object. Should be one of existing colors.",
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
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "N/A",
						},
						"tags": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Collection of tag identifiers.",
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
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Object type.",
									},
									"color": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Color of the object. Should be one of existing colors.",
									},
								},
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
				},
			},
			"servers": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Domain servers.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object name. Must be unique in the domain.",
						},
						"active": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Domain server status.",
						},
						"ipv4_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv4 address.",
						},
						"ipv6_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv6 address.",
						},
						"multi_domain_server": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Multi Domain server name or UID.",
						},
						"skip_start_domain_server": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Set this value to be true to prevent starting the new created domain.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Domain server type.",
						},
					},
				},
			},
			"tags": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Collection of tag objects identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.",
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
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object type.",
						},
						"color": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Color of the object. Should be one of existing colors.",
						},
					},
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

func dataSourceManagementGlobalDomainRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showGlobalDomainRes, err := client.ApiCall("show-global-domain", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showGlobalDomainRes.Success {
		return fmt.Errorf(showGlobalDomainRes.ErrorMsg)
	}

	globalDomain := showGlobalDomainRes.GetData()

	log.Println("Read Global Domain - Show JSON = ", globalDomain)

	d.SetId("show-global-domain-" + acctest.RandString(10))

	if v := globalDomain["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := globalDomain["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := globalDomain["type"]; v != nil {
		_ = d.Set("type", v)
	}

	if v := globalDomain["domain-type"]; v != nil {
		_ = d.Set("domain_type", v)
	}

	if globalDomain["global-domain-assignments"] != nil {
		globalDomainAssignmentsList := globalDomain["global-domain-assignments"].([]interface{})

		if len(globalDomainAssignmentsList) > 0 {
			var globalDomainAssignmentsListToReturn []map[string]interface{}

			for i := range globalDomainAssignmentsList {
				globalDomainAssignmentsMap := globalDomainAssignmentsList[i].(map[string]interface{})

				globalDomainAssignmentsMapToAdd := make(map[string]interface{})

				if v, _ := globalDomainAssignmentsMap["name"]; v != nil {
					globalDomainAssignmentsMapToAdd["name"] = v
				}
				if v, _ := globalDomainAssignmentsMap["uid"]; v != nil {
					globalDomainAssignmentsMapToAdd["uid"] = v
				}
				if v, _ := globalDomainAssignmentsMap["type"]; v != nil {
					globalDomainAssignmentsMapToAdd["type"] = v
				}
				if v, _ := globalDomainAssignmentsMap["assignment-status"]; v != nil {
					globalDomainAssignmentsMapToAdd["assignments_status"] = v
				}
				if globalDomainAssignmentsMap["assignment-up-to-date"] != nil {
					assignmentUpToDateMap := globalDomainAssignmentsMap["assignment-up-to-date"].(map[string]interface{})

					assignmentUpToDateMapToReturn := make(map[string]interface{})

					if v, _ := assignmentUpToDateMap["iso-8601"]; v != nil {
						assignmentUpToDateMapToReturn["iso_8601"] = v
					}
					if v, _ := assignmentUpToDateMap["posix"]; v != nil {
						assignmentUpToDateMapToReturn["posix"] = v
					}

					globalDomainAssignmentsMapToAdd["assignment_up_to_date"] = assignmentUpToDateMapToReturn
				}

				if globalDomainAssignmentsMap["dependent-domain"] != nil {
					dependentDomainMap := globalDomainAssignmentsMap["dependent-domain"].(map[string]interface{})

					dependentDomainMapToReturn := make(map[string]interface{})

					if v, _ := dependentDomainMap["name"]; v != nil {
						dependentDomainMapToReturn["name"] = v
					}
					if v, _ := dependentDomainMap["uid"]; v != nil {
						dependentDomainMapToReturn["uid"] = v
					}
					if v, _ := dependentDomainMap["type"]; v != nil {
						dependentDomainMapToReturn["type"] = v
					}
					if v, _ := dependentDomainMap["color"]; v != nil {
						dependentDomainMapToReturn["color"] = v
					}

					globalDomainAssignmentsMapToAdd["dependent_domain"] = dependentDomainMapToReturn
				}

				if v, _ := globalDomainAssignmentsMap["global-access-policy"]; v != nil {
					globalDomainAssignmentsMapToAdd["global_access_policy"] = v
				}
				if v, _ := globalDomainAssignmentsMap["global-threat-prevention-policy"]; v != nil {
					globalDomainAssignmentsMapToAdd["global_threat_prevention_policy"] = v
				}
				if v, _ := globalDomainAssignmentsMap["manage-protection-actions"]; v != nil {
					globalDomainAssignmentsMapToAdd["manage_protection_actions"] = v
				}

				if globalDomainAssignmentsMap["tags"] != nil {
					tagsList := globalDomainAssignmentsMap["tags"].([]interface{})

					if len(tagsList) > 0 {
						var tagsListToReturn []map[string]interface{}

						for i := range tagsList {
							tagsMap := tagsList[i].(map[string]interface{})

							tagsMapToAdd := make(map[string]interface{})

							if v, _ := tagsMap["name"]; v != nil {
								tagsMapToAdd["name"] = v
							}
							if v, _ := tagsMap["uid"]; v != nil {
								tagsMapToAdd["uid"] = v
							}
							if v, _ := tagsMap["color"]; v != nil {
								tagsMapToAdd["color"] = v
							}
							tagsListToReturn = append(tagsListToReturn, tagsMapToAdd)
						}

						globalDomainAssignmentsMapToAdd["tags"] = tagsListToReturn
					} else {
						globalDomainAssignmentsMapToAdd["tags"] = tagsList
					}
				}
				if v, _ := globalDomainAssignmentsMap["color"]; v != nil {
					globalDomainAssignmentsMapToAdd["color"] = v
				}
				if v, _ := globalDomainAssignmentsMap["comments"]; v != nil {
					globalDomainAssignmentsMapToAdd["comments"] = v
				}

				globalDomainAssignmentsListToReturn = append(globalDomainAssignmentsListToReturn, globalDomainAssignmentsMapToAdd)
			}

			_ = d.Set("global_domain_assignments", globalDomainAssignmentsListToReturn)
		} else {
			_ = d.Set("global_domain_assignments", globalDomainAssignmentsList)
		}
	} else {
		_ = d.Set("global_domain_assignments", nil)
	}

	if globalDomain["servers"] != nil {
		serversList := globalDomain["servers"].([]interface{})

		if len(serversList) > 0 {
			var serversListToReturn []map[string]interface{}

			for i := range serversList {
				serversMap := serversList[i].(map[string]interface{})

				serversMapToAdd := make(map[string]interface{})

				if v, _ := serversMap["name"]; v != nil {
					serversMapToAdd["name"] = v
				}
				if v, _ := serversMap["active"]; v != nil {
					serversMapToAdd["active"] = v
				}
				if v, _ := serversMap["ipv4-address"]; v != nil {
					serversMapToAdd["ipv4_address"] = v
				}
				if v, _ := serversMap["ipv6-address"]; v != nil {
					serversMapToAdd["ipv6_address"] = v
				}
				if v, _ := serversMap["multi-domain-server"]; v != nil {
					serversMapToAdd["multi_domain_server"] = v
				}
				if v, _ := serversMap["skip-start-domain-server"]; v != nil {
					serversMapToAdd["skip_start_domain_server"] = v
				}
				if v, _ := serversMap["type"]; v != nil {
					serversMapToAdd["type"] = v
				}

				serversListToReturn = append(serversListToReturn, serversMapToAdd)
			}

			_ = d.Set("servers", serversListToReturn)
		} else {
			_ = d.Set("servers", serversList)
		}
	} else {
		_ = d.Set("servers", nil)
	}

	if globalDomain["tags"] != nil {
		tagsList := globalDomain["tags"].([]interface{})

		if len(tagsList) > 0 {
			var tagsListToReturn []map[string]interface{}

			for i := range tagsList {
				tagsMap := tagsList[i].(map[string]interface{})

				tagsMapToAdd := make(map[string]interface{})

				if v, _ := tagsMap["name"]; v != nil {
					tagsMapToAdd["name"] = v
				}
				if v, _ := tagsMap["uid"]; v != nil {
					tagsMapToAdd["uid"] = v
				}
				if v, _ := tagsMap["type"]; v != nil {
					tagsMapToAdd["type"] = v
				}
				if v, _ := tagsMap["color"]; v != nil {
					tagsMapToAdd["color"] = v
				}

				tagsListToReturn = append(tagsListToReturn, tagsMapToAdd)
			}

			_ = d.Set("tags", tagsListToReturn)
		} else {
			_ = d.Set("tags", tagsList)
		}
	} else {
		_ = d.Set("tags", nil)
	}

	if v := globalDomain["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := globalDomain["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil
}
