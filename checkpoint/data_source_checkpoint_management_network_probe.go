package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementNetworkProbe() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceManagementNetworkProbeRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"http_options": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Additional options when [protocol] is set to \"http\".",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"destination": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The destination URL.",
						},
					},
				},
			},
			"icmp_options": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Additional options when [protocol] is set to \"icmp\".",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"destination": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "One of these:<br>- Name or UID of an existing object with a unicast IPv4 address (Host, Security Gateway, and so on).<br>- A unicast IPv4 address string (if you do not want to create such an object).",
						},
						"source": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "One of these:<br>- The string \"main-ip\" (the probe uses the main IPv4 address of the Security Gateway objects you specified in the parameter [install-on]).<br>- Name or UID of an existing object of type 'Host' with a unicast IPv4 address.<br>- A unicast IPv4 address string (if you do not want to create such an object).",
						},
					},
				},
			},
			"install_on": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of Check Point Security Gateways that generate the probe, identified by name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"protocol": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The probing protocol to use.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"interval": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The time interval (in seconds) between each probe request.<br>Best Practice - The interval value should be lower than the timeout value.",
			},
			"timeout": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The probe expiration timeout (in seconds). If there is not a single reply within this time, the status of the probe changes to \"Down\".",
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

func dataSourceManagementNetworkProbeRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showNetworkProbeRes, err := client.ApiCall("show-network-probe", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showNetworkProbeRes.Success {
		if objectNotFound(showNetworkProbeRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showNetworkProbeRes.ErrorMsg)
	}

	networkProbe := showNetworkProbeRes.GetData()

	log.Println("Read NetworkProbe - Show JSON = ", networkProbe)

	if v := networkProbe["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := networkProbe["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if networkProbe["http-options"] != nil {

		httpOptionsMap := networkProbe["http-options"].(map[string]interface{})

		httpOptionsMapToReturn := make(map[string]interface{})

		if v, _ := httpOptionsMap["destination"]; v != nil {
			httpOptionsMapToReturn["destination"] = v
		}
		_ = d.Set("http_options", httpOptionsMapToReturn)
	} else {
		_ = d.Set("http_options", nil)
	}

	if networkProbe["icmp-options"] != nil {

		icmpOptionsMap := networkProbe["icmp-options"].(map[string]interface{})

		icmpOptionsMapToReturn := make(map[string]interface{})

		if v, _ := icmpOptionsMap["destination"]; v != nil {
			icmpOptionsMapToReturn["destination"] = v
		}
		if v, _ := icmpOptionsMap["source"]; v != nil {
			icmpOptionsMapToReturn["source"] = v
		}
		_ = d.Set("icmp_options", icmpOptionsMapToReturn)
	} else {
		_ = d.Set("icmp_options", nil)
	}

	if networkProbe["install-on"] != nil {
		installOnJson, ok := networkProbe["install-on"].([]interface{})
		if ok {
			installOnIds := make([]string, 0)
			if len(installOnJson) > 0 {
				for _, install_on := range installOnJson {
					install_on := install_on.(map[string]interface{})
					installOnIds = append(installOnIds, install_on["name"].(string))
				}
			}
			_ = d.Set("install_on", installOnIds)
		}
	} else {
		_ = d.Set("install_on", nil)
	}

	if v := networkProbe["protocol"]; v != nil {
		_ = d.Set("protocol", v)
	}

	if networkProbe["tags"] != nil {
		tagsJson, ok := networkProbe["tags"].([]interface{})
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

	if v := networkProbe["interval"]; v != nil {
		_ = d.Set("interval", v)
	}

	if v := networkProbe["timeout"]; v != nil {
		_ = d.Set("timeout", v)
	}

	if v := networkProbe["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := networkProbe["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := networkProbe["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := networkProbe["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}
