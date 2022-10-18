package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementTacacsServer() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementTacacsServerRead,
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
			"encryption": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Is there a secret key defined on the server. Must be set true when \"server-type\" was selected to be \"TACACS+\".",
			},
			"groups": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"priority": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The priority of the TACACS Server in case it is a member of a TACACS Group.",
			},
			"server": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "The UID or Name of the host that is the TACACS Server.",
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
					},
				},
			},
			"server_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Server type, TACACS or TACACS+.",
			},
			"service": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Server service, only relevant when \"server-type\" is TACACS.",
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
					},
				},
			},
		},
	}
}

func dataSourceManagementTacacsServerRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showTacacsServerRes, err := client.ApiCall("show-tacacs-server", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showTacacsServerRes.Success {
		return fmt.Errorf(showTacacsServerRes.ErrorMsg)
	}

	tacacsServer := showTacacsServerRes.GetData()

	log.Println("Read Tacacs Server - Show JSON = ", tacacsServer)

	if v := tacacsServer["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := tacacsServer["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := tacacsServer["encryption"]; v != nil {
		_ = d.Set("encryption", v)
	}

	if tacacsServer["groups"] != nil {
		groupsJson := tacacsServer["groups"].([]interface{})
		groupsIds := make([]string, 0)
		if len(groupsJson) > 0 {
			// Create slice of group names
			for _, group := range groupsJson {
				group := group.(map[string]interface{})
				groupsIds = append(groupsIds, group["name"].(string))
			}
		}
		_ = d.Set("groups", groupsIds)
	} else {
		_ = d.Set("groups", nil)
	}

	if v := tacacsServer["priority"]; v != nil {
		_ = d.Set("priority", v)
	}

	if tacacsServer["server"] != nil {
		serverMap := tacacsServer["server"].(map[string]interface{})

		serverMapToReturn := make(map[string]interface{})

		if v, _ := serverMap["name"]; v != nil {
			serverMapToReturn["name"] = v
		}
		if v, _ := serverMap["uid"]; v != nil {
			serverMapToReturn["uid"] = v
		}

		_ = d.Set("server", serverMapToReturn)
	} else {
		_ = d.Set("server", nil)
	}

	if v := tacacsServer["server-type"]; v != nil {
		_ = d.Set("server_type", v)
	}

	if tacacsServer["service"] != nil {
		serviceMap := tacacsServer["service"].(map[string]interface{})
		log.Println("service detected!!!")
		serviceMapToReturn := make(map[string]interface{})

		if v, _ := serviceMap["name"]; v != nil {
			serviceMapToReturn["name"] = v
		}
		if v, _ := serviceMap["uid"]; v != nil {
			serviceMapToReturn["uid"] = v
		}

		_ = d.Set("service", serviceMapToReturn)

	} else {
		_ = d.Set("service", nil)
	}

	return nil
}
