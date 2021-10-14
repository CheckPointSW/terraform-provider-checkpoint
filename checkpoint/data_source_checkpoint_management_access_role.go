package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementAccessRole() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementAccessRoleRead,
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
			"machines": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Machines that can access the system.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"source": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Active Directory name or UID or Identity Tag.",
						},
						"selection": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Name or UID of an object selected from source.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"base_dn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "When source is \"Active Directory\" use \"base-dn\" to refine the query in AD database.",
						},
					},
				},
			},
			"networks": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of Network objects identified by the name or UID that can access the system.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"remote_access_clients": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Remote access clients identified by name or UID.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"users": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Users that can access the system.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"source": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Active Directory name or UID or Identity Tag  or Internal User Groups or LDAP groups or Guests.",
						},
						"selection": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Name or UID of an object selected from source.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"base_dn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "When source is \"Active Directory\" use \"base-dn\" to refine the query in AD database.",
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

func dataSourceManagementAccessRoleRead(d *schema.ResourceData, m interface{}) error {
	// TODO: Ask Roy
	TYPE_TO_SOURCE := map[string]string{
		"identity-tag":  "Identity Tag",
		"user-group":    "Internal User Groups",
		"CpmiAnyObject": "Guests",
	}
	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showAccessRoleRes, err := client.ApiCall("show-access-role", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showAccessRoleRes.Success {
		if objectNotFound(showAccessRoleRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showAccessRoleRes.ErrorMsg)
	}

	accessRole := showAccessRoleRes.GetData()

	log.Println("Read AccessRole - Show JSON = ", accessRole)

	if v := accessRole["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if accessRole["machines"] != nil {

		machinesList, ok := accessRole["machines"].([]interface{})

		if ok {

			if len(machinesList) > 0 {

				var machinesListToReturn []map[string]interface{}
				machinesByType := make(map[string]map[string]interface{})
				for i := range machinesList {
					machinesMap := machinesList[i].(map[string]interface{})
					machinesMapToAdd := make(map[string]interface{})
					if v, _ := machinesMap["base-dn"]; v != nil {
						machinesMapToAdd["base_dn"] = v
					}
					if v, _ := machinesMap["type"]; v != nil {
						if value, ok := machinesByType[TYPE_TO_SOURCE[v.(string)]]; ok {
							machinesMapToAdd = value
							if val, _ := machinesMap["name"]; val != nil {
								machinesMapToAdd["selection"] = append(machinesMapToAdd["selection"].([]string), machinesMap["name"].(string))
								machinesByType[TYPE_TO_SOURCE[v.(string)]] = machinesMapToAdd
							}
						} else {
							machinesMapToAdd["source"] = TYPE_TO_SOURCE[v.(string)]
							if val, _ := machinesMap["name"]; val != nil {
								machinesMapToAdd["selection"] = []string{machinesMap["name"].(string)}
								machinesByType[TYPE_TO_SOURCE[v.(string)]] = machinesMapToAdd
							}
						}
					}
				}
				for _, v := range machinesByType {
					machinesListToReturn = append(machinesListToReturn, v)
				}
				_ = d.Set("machines", machinesListToReturn)
			}
		} else {
			_ = d.Set("machines", nil)
		}
	}
	if accessRole["networks"] != nil {
		networksJson, ok := accessRole["networks"].([]interface{})
		if ok {
			networksIds := make([]string, 0)
			if len(networksJson) > 0 {
				for _, networks := range networksJson {
					networks := networks.(map[string]interface{})
					networksIds = append(networksIds, networks["name"].(string))
				}
			}
			_ = d.Set("networks", networksIds)
		} else {
			_ = d.Set("networks", nil)
		}
	} else {
		_ = d.Set("networks", nil)
	}

	if v := accessRole["remote-access-client"]; v != nil {
		//TODO: Ask Roy
		_ = d.Set("remote_access_clients", v.(map[string]interface{})["name"])
	} else {
		_ = d.Set("remote_access_clients", v)
	}

	if accessRole["tags"] != nil {
		tagsJson, ok := accessRole["tags"].([]interface{})
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
	log.Println("what is the user?")
	log.Println(accessRole["users"])
	log.Println("what is the user?")
	if accessRole["users"] != nil {

		usersList, ok := accessRole["users"].([]interface{})

		if ok {

			if len(usersList) > 0 {

				var usersListToReturn []map[string]interface{}
				usersByType := make(map[string]map[string]interface{})
				for i := range usersList {
					usersMap := usersList[i].(map[string]interface{})
					usersMapToAdd := make(map[string]interface{})
					if v, _ := usersMap["base-dn"]; v != nil {
						usersMapToAdd["base_dn"] = v
					}
					log.Println("typeeee")
					log.Println(usersMap["type"])
					log.Println("typeeee")
					if v, _ := usersMap["type"]; v != nil {
						if value, ok := usersByType[TYPE_TO_SOURCE[v.(string)]]; ok {
							usersMapToAdd = value
							if val, _ := usersMap["name"]; val != nil {
								usersMapToAdd["selection"] = append(usersMapToAdd["selection"].([]string), usersMap["name"].(string))
								usersByType[TYPE_TO_SOURCE[v.(string)]] = usersMapToAdd
							}
						} else {
							usersMapToAdd["source"] = TYPE_TO_SOURCE[v.(string)]
							if val, _ := usersMap["name"]; val != nil {
								usersMapToAdd["selection"] = []string{usersMap["name"].(string)}
								usersByType[TYPE_TO_SOURCE[v.(string)]] = usersMapToAdd
							}
						}

					}
				}
				for _, v := range usersByType {
					usersListToReturn = append(usersListToReturn, v)
				}
				_ = d.Set("users", usersListToReturn)
			}
		} else {
			_ = d.Set("users", nil)
		}
	}

	if v := accessRole["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := accessRole["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := accessRole["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := accessRole["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}
