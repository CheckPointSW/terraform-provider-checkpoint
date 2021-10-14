package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"

	"strconv"
)

func resourceManagementAccessRole() *schema.Resource {
	return &schema.Resource{
		Create: createManagementAccessRole,
		Read:   readManagementAccessRole,
		Update: updateManagementAccessRole,
		Delete: deleteManagementAccessRole,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"machines": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Machines that can access the system.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"source": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Active Directory name or UID or Identity Tag.",
						},
						"selection": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Name or UID of an object selected from source.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"base_dn": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "When source is \"Active Directory\" use \"base-dn\" to refine the query in AD database.",
						},
					},
				},
			},
			"networks": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of Network objects identified by the name or UID that can access the system.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"remote_access_clients": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Remote access clients identified by name or UID.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"users": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Users that can access the system.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"source": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Active Directory name or UID or Identity Tag  or Internal User Groups or LDAP groups or Guests.",
						},
						"selection": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Name or UID of an object selected from source.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"base_dn": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "When source is \"Active Directory\" use \"base-dn\" to refine the query in AD database.",
						},
					},
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

func createManagementAccessRole(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	accessRole := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		accessRole["name"] = v.(string)
	}

	if v, ok := d.GetOk("machines"); ok {

		machinesList := v.([]interface{})

		if len(machinesList) > 0 {

			var machinesPayload []map[string]interface{}

			for i := range machinesList {

				Payload := make(map[string]interface{})

				if v, ok := d.GetOk("machines." + strconv.Itoa(i) + ".source"); ok {
					Payload["source"] = v.(string)
				}
				if v, ok := d.GetOk("machines." + strconv.Itoa(i) + ".selection"); ok {
					Payload["selection"] = v.(*schema.Set).List()
				}
				if v, ok := d.GetOk("machines." + strconv.Itoa(i) + ".base_dn"); ok {
					Payload["base-dn"] = v.(string)
				}
				machinesPayload = append(machinesPayload, Payload)
			}
			accessRole["machines"] = machinesPayload
		}
	}

	if v, ok := d.GetOk("networks"); ok {
		accessRole["networks"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("remote_access_clients"); ok {
		accessRole["remote-access-clients"] = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		accessRole["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("users"); ok {

		usersList := v.([]interface{})

		if len(usersList) > 0 {

			var usersPayload []map[string]interface{}

			for i := range usersList {

				Payload := make(map[string]interface{})

				if v, ok := d.GetOk("users." + strconv.Itoa(i) + ".source"); ok {
					Payload["source"] = v.(string)
				}
				if v, ok := d.GetOk("users." + strconv.Itoa(i) + ".selection"); ok {
					Payload["selection"] = v.(*schema.Set).List()
				}
				if v, ok := d.GetOk("users." + strconv.Itoa(i) + ".base_dn"); ok {
					Payload["base-dn"] = v.(string)
				}
				usersPayload = append(usersPayload, Payload)
			}
			accessRole["users"] = usersPayload
		}
	}

	if v, ok := d.GetOk("color"); ok {
		accessRole["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		accessRole["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		accessRole["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		accessRole["ignore-errors"] = v.(bool)
	}

	log.Println("Create AccessRole - Map = ", accessRole)

	addAccessRoleRes, err := client.ApiCall("add-access-role", accessRole, client.GetSessionID(), true, false)
	if err != nil || !addAccessRoleRes.Success {
		if addAccessRoleRes.ErrorMsg != "" {
			return fmt.Errorf(addAccessRoleRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addAccessRoleRes.GetData()["uid"].(string))

	return readManagementAccessRole(d, m)
}

func readManagementAccessRole(d *schema.ResourceData, m interface{}) error {
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

func updateManagementAccessRole(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	accessRole := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		accessRole["name"] = oldName
		accessRole["new-name"] = newName
	} else {
		accessRole["name"] = d.Get("name")
	}

	if d.HasChange("machines") {
		if v, ok := d.GetOk("machines"); ok {

			machinesList := v.([]interface{})
			var machinesPayload []map[string]interface{}

			for i := range machinesList {
				Payload := make(map[string]interface{})
				Payload["source"] = d.Get("machines." + strconv.Itoa(i) + ".source")
				Payload["selection"] = d.Get("machines." + strconv.Itoa(i) + ".selection").(*schema.Set).List()
				Payload["base-dn"] = d.Get("machines." + strconv.Itoa(i) + ".base-dn")
				machinesPayload = append(machinesPayload, Payload)
			}
			accessRole["machines"] = machinesPayload
		} else {
			accessRole["machines"] = "any"
		}
	}
	if d.HasChange("networks") {
		if v, ok := d.GetOk("networks"); ok {
			accessRole["networks"] = v.(*schema.Set).List()
		} else {
			accessRole["networks"] = "any"
		}
	}

	if ok := d.HasChange("remote_access_clients"); ok {
		if v, ok := d.GetOk("remote_access_clients"); ok {
			accessRole["remote-access-clients"] = v.(string)
		} else {
			accessRole["remote-access-clients"] = "any"
		}
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			accessRole["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			accessRole["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if d.HasChange("users") {
		if v, ok := d.GetOk("users"); ok {

			usersList := v.([]interface{})
			var usersPayload []map[string]interface{}

			for i := range usersList {
				Payload := make(map[string]interface{})
				Payload["source"] = d.Get("users." + strconv.Itoa(i) + ".source")
				Payload["selection"] = d.Get("users." + strconv.Itoa(i) + ".selection").(*schema.Set).List()
				Payload["base-dn"] = d.Get("users." + strconv.Itoa(i) + ".base-dn")
				usersPayload = append(usersPayload, Payload)
			}
			accessRole["users"] = usersPayload
		} else {
			accessRole["users"] = "any"
		}
	}

	if ok := d.HasChange("color"); ok {
		accessRole["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		accessRole["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		accessRole["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		accessRole["ignore-errors"] = v.(bool)
	}

	log.Println("Update AccessRole - Map = ", accessRole)

	updateAccessRoleRes, err := client.ApiCall("set-access-role", accessRole, client.GetSessionID(), true, false)
	if err != nil || !updateAccessRoleRes.Success {
		if updateAccessRoleRes.ErrorMsg != "" {
			return fmt.Errorf(updateAccessRoleRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementAccessRole(d, m)
}

func deleteManagementAccessRole(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	accessRolePayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete AccessRole")

	deleteAccessRoleRes, err := client.ApiCall("delete-access-role", accessRolePayload, client.GetSessionID(), true, false)
	if err != nil || !deleteAccessRoleRes.Success {
		if deleteAccessRoleRes.ErrorMsg != "" {
			return fmt.Errorf(deleteAccessRoleRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
