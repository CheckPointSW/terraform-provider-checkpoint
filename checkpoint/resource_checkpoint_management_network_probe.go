package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementNetworkProbe() *schema.Resource {
	return &schema.Resource{
		Create: createManagementNetworkProbe,
		Read:   readManagementNetworkProbe,
		Update: updateManagementNetworkProbe,
		Delete: deleteManagementNetworkProbe,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"http_options": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Additional options when [protocol] is set to \"http\".",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"destination": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The destination URL.",
						},
					},
				},
			},
			"icmp_options": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Additional options when [protocol] is set to \"icmp\".",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"destination": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "One of these:<br>- Name or UID of an existing object with a unicast IPv4 address (Host, Security Gateway, and so on).<br>- A unicast IPv4 address string (if you do not want to create such an object).",
						},
						"source": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "One of these:<br>- The string \"main-ip\" (the probe uses the main IPv4 address of the Security Gateway objects you specified in the parameter [install-on]).<br>- Name or UID of an existing object of type 'Host' with a unicast IPv4 address.<br>- A unicast IPv4 address string (if you do not want to create such an object).",
							Default:     "main-ip",
						},
					},
				},
			},
			"install_on": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "Collection of Check Point Security Gateways that generate the probe, identified by name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"protocol": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The probing protocol to use.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The time interval (in seconds) between each probe request.<br>Best Practice - The interval value should be lower than the timeout value.",
				Default:     10,
			},
			"timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The probe expiration timeout (in seconds). If there is not a single reply within this time, the status of the probe changes to \"Down\".",
				Default:     20,
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

func createManagementNetworkProbe(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	networkProbe := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		networkProbe["name"] = v.(string)
	}

	if _, ok := d.GetOk("http_options"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("http_options.destination"); ok {
			res["destination"] = v.(string)
		}
		networkProbe["http-options"] = res
	}

	if _, ok := d.GetOk("icmp_options"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("icmp_options.destination"); ok {
			res["destination"] = v.(string)
		}
		if v, ok := d.GetOk("icmp_options.source"); ok {
			res["source"] = v.(string)
		}
		networkProbe["icmp-options"] = res
	}

	if v, ok := d.GetOk("install_on"); ok {
		networkProbe["install-on"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("protocol"); ok {
		networkProbe["protocol"] = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		networkProbe["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("interval"); ok {
		networkProbe["interval"] = v.(int)
	}

	if v, ok := d.GetOk("timeout"); ok {
		networkProbe["timeout"] = v.(int)
	}

	if v, ok := d.GetOk("color"); ok {
		networkProbe["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		networkProbe["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		networkProbe["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		networkProbe["ignore-errors"] = v.(bool)
	}

	log.Println("Create NetworkProbe - Map = ", networkProbe)

	addNetworkProbeRes, err := client.ApiCall("add-network-probe", networkProbe, client.GetSessionID(), true, false)
	if err != nil || !addNetworkProbeRes.Success {
		if addNetworkProbeRes.ErrorMsg != "" {
			return fmt.Errorf(addNetworkProbeRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addNetworkProbeRes.GetData()["uid"].(string))

	return readManagementNetworkProbe(d, m)
}

func readManagementNetworkProbe(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
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

func updateManagementNetworkProbe(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	networkProbe := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		networkProbe["name"] = oldName
		networkProbe["new-name"] = newName
	} else {
		networkProbe["name"] = d.Get("name")
	}

	if d.HasChange("http_options") {

		if _, ok := d.GetOk("http_options"); ok {

			res := make(map[string]interface{})

			if v, ok := d.GetOk("http_options.destination"); ok {
				res["destination"] = v.(string)
			}
			networkProbe["http-options"] = res
		}
	}

	if d.HasChange("icmp_options") {

		if _, ok := d.GetOk("icmp_options"); ok {

			res := make(map[string]interface{})

			if v, ok := d.GetOk("icmp_options.destination"); ok {
				res["destination"] = v.(string)
			}
			if d.HasChange("icmp_options.source") {
				res["source"] = d.Get("icmp_options.source")
			}
			networkProbe["icmp-options"] = res
		}
	}

	if d.HasChange("install_on") {
		if v, ok := d.GetOk("install_on"); ok {
			networkProbe["install-on"] = v.(*schema.Set).List()
		} else {
			oldInstall_On, _ := d.GetChange("install_on")
			networkProbe["install-on"] = map[string]interface{}{"remove": oldInstall_On.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("protocol"); ok {
		networkProbe["protocol"] = d.Get("protocol")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			networkProbe["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			networkProbe["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("interval"); ok {
		networkProbe["interval"] = d.Get("interval")
	}

	if ok := d.HasChange("timeout"); ok {
		networkProbe["timeout"] = d.Get("timeout")
	}

	if ok := d.HasChange("color"); ok {
		networkProbe["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		networkProbe["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		networkProbe["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		networkProbe["ignore-errors"] = v.(bool)
	}

	log.Println("Update NetworkProbe - Map = ", networkProbe)

	updateNetworkProbeRes, err := client.ApiCall("set-network-probe", networkProbe, client.GetSessionID(), true, false)
	if err != nil || !updateNetworkProbeRes.Success {
		if updateNetworkProbeRes.ErrorMsg != "" {
			return fmt.Errorf(updateNetworkProbeRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementNetworkProbe(d, m)
}

func deleteManagementNetworkProbe(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	networkProbePayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete NetworkProbe")

	deleteNetworkProbeRes, err := client.ApiCall("delete-network-probe", networkProbePayload, client.GetSessionID(), true, false)
	if err != nil || !deleteNetworkProbeRes.Success {
		if deleteNetworkProbeRes.ErrorMsg != "" {
			return fmt.Errorf(deleteNetworkProbeRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
