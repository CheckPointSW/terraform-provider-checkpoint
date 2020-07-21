package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"reflect"
	"strconv"
)

func resourcePhysicalInterface() *schema.Resource {
	return &schema.Resource{
		Create: createPhysicalInterface,
		Read:   readPhysicalInterface,
		Update: updatePhysicalInterface,
		Delete: deletePhysicalInterface,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "interface name",
			},
			"auto_negotiation": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Activating auto_negotiation will skip the speed and duplex configuration",
				Default:     "Not-Configured",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// if the parameter does not exist in the tf configuration file it gets the default value. we
					// don't want it to be considered as a change on the 'plan' stage,
					// we don't want to see change from (for example) "true" to ""
					if new == "Not-Configured" || new == old {
						return true
					}
					return false
				},
			},
			"comments": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "comments",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if new == "" || new == old {
						return true
					}
					return false
				},
			},
			"duplex": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Duplex is not relevant when 'auto_negotiation' is enabled.",
				Default:     "Not-Configured",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if new == "Not-Configured" || new == old {
						return true
					}
					return false
				},
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "",
			},
			"ipv4_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
				Default:     "Not-Configured",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if new == "Not-Configured" || new == old {
						return true
					}
					return false
				},
			},
			"monitor_mode": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
				Default:     "Not-Configured",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if new == "Not-Configured" || new == old {
						return true
					}
					return false
				},
			},
			"ipv6_autoconfig": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
				Default:     "Not-Configured",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if new == "Not-Configured" || new == old {
						return true
					}
					return false
				},
			},
			"mac_addr": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
				Default:     "Not-Configured",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if new == "Not-Configured" || new == old {
						return true
					}
					return false
				},
			},
			"mtu": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if new == "0" || new == old {
						return true
					}
					return false
				},
			},
			"rx_ringsize": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
				Default:     "Not-Configured",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if new == "Not-Configured" || new == old {
						return true
					}
					return false
				},
			},
			"ipv6_mask_length": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if new == "0" || new == old {
						return true
					}
					return false
				},
			},
			"ipv6_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
				Default:     "Not-Configured",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if new == "Not-Configured" || new == old {
						return true
					}
					return false
				},
			},
			"ipv4_mask_length": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if new == "0" || new == old {
						return true
					}
					return false
				},
			},

			"speed": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Speed is not relevant when 'auto_negotiation' is enabled",
				Default:     "Not-Configured",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if new == "Not-Configured" || new == old {
						return true
					}
					return false
				},
			},
			"tx_ringsize": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
				Default:     "Not-Configured",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if new == "Not-Configured" || new == old {
						return true
					}
					return false
				},
			},
		},
	}
}

func physicalInterfaceParseSchemaToMap(d *schema.ResourceData) map[string]interface{} {
	physicalInterfaceMap := make(map[string]interface{})

	if val, ok := d.GetOk("name"); ok {
		physicalInterfaceMap["name"] = val.(string)
	}

	if d.HasChange("monitor_mode") {
		physicalInterfaceMap["monitor-mode"] = d.Get("monitor_mode")
	}

	if d.HasChange("duplex") {
		physicalInterfaceMap["duplex"] = d.Get("duplex")
	}

	if d.HasChange("ipv6_autoconfig") {
		physicalInterfaceMap["ipv6-autoconfig"] = d.Get("ipv6_autoconfig")
	}

	if d.HasChange("mac_addr") {
		physicalInterfaceMap["mac-addr"] = d.Get("mac_addr")
	}

	if val, ok := d.GetOkExists("enabled"); ok {
		physicalInterfaceMap["enabled"] = val.(bool)
	}

	if val, ok := d.GetOk("comments"); ok {
		physicalInterfaceMap["comments"] = val.(string)
	}

	if d.HasChange("mtu") {
		physicalInterfaceMap["mtu"] = d.Get("mtu")
	}

	if d.HasChange("rx_ringsize") {
		physicalInterfaceMap["rx-ringsize"] = d.Get("rx_ringsize")
	}

	if d.HasChange("rx_ringsize") {
		physicalInterfaceMap["rx-ringsize"] = d.Get("rx_ringsize")
	}

	if d.HasChange("speed") {
		physicalInterfaceMap["speed"] = d.Get("speed")
	}

	if d.HasChange("rx_ringsize") {
		physicalInterfaceMap["rx-ringsize"] = d.Get("rx_ringsize")
	}

	if d.HasChange("tx_ringsize") {
		physicalInterfaceMap["tx-ringsize"] = d.Get("tx_ringsize")
	}

	if d.HasChange("auto_negotiation") {
		physicalInterfaceMap["auto-negotiation"] = d.Get("auto_negotiation")
	}

	return physicalInterfaceMap
}

func createPhysicalInterface(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	payload := physicalInterfaceParseSchemaToMap(d)

	if v, ok := d.GetOk("ipv6_address"); ok {
		payload["ipv6-address"] = v.(string)
	}
	if v, ok := d.GetOk("ipv6_mask_length"); ok {
		payload["ipv6-mask-length"] = v.(int)
	}

	if v, ok := d.GetOk("ipv4_address"); ok {
		payload["ipv4-address"] = v.(string)
	}
	if v, ok := d.GetOk("ipv4_mask_length"); ok {
		payload["ipv4-mask-length"] = v.(int)
	}

	setPIRes, _ := client.ApiCall("set-physical-interface", payload, client.GetSessionID(), true, false)
	if !setPIRes.Success {
		return fmt.Errorf(setPIRes.ErrorMsg)
	}

	// Set Schema UID = Object key
	d.SetId(setPIRes.GetData()["name"].(string))

	return readPhysicalInterface(d, m)
}

func readPhysicalInterface(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	payload := map[string]interface{}{
		"name": d.Get("name"),
	}
	showPIRes, _ := client.ApiCall("show-physical-interface", payload, client.GetSessionID(), true, false)
	if !showPIRes.Success {
		// Handle deletion of an object from other clients - Object not found
		if objectNotFound(showPIRes.GetData()["code"].(string)) {
			d.SetId("") // Destroy resource
			return nil
		}
		return fmt.Errorf(showPIRes.ErrorMsg)
	}
	PIJson := showPIRes.GetData()

	_ = d.Set("name", PIJson["name"].(string))

	if reflect.TypeOf(PIJson["monitor-mode"]).Kind() == reflect.Bool {
		_ = d.Set("monitor_mode", strconv.FormatBool(PIJson["monitor-mode"].(bool)))
	} else {
		_ = d.Set("monitor_mode", PIJson["monitor-mode"].(string))
	}

	_ = d.Set("duplex", PIJson["duplex"].(string))
	_ = d.Set("ipv6_autoconfig", PIJson["ipv6-autoconfig"].(string))
	_ = d.Set("mac_addr", PIJson["mac-addr"].(string))
	_ = d.Set("enabled", PIJson["enabled"].(bool))
	_ = d.Set("mtu", PIJson["mtu"].(string))
	_ = d.Set("ipv6_mask_length", PIJson["ipv6-mask-length"].(string))
	_ = d.Set("rx_ringsize", PIJson["rx-ringsize"].(string))
	_ = d.Set("ipv6_address", PIJson["ipv6-address"].(string))
	_ = d.Set("ipv4_address", PIJson["ipv4-address"].(string))
	_ = d.Set("ipv4_mask_length", PIJson["ipv4-mask-length"].(string))

	_ = d.Set("speed", PIJson["speed"].(string))
	_ = d.Set("comments", PIJson["comments"].(string))

	_ = d.Set("tx_ringsize", PIJson["tx-ringsize"].(string))
	if reflect.TypeOf(PIJson["auto-negotiation"]).Kind() == reflect.Bool {
		_ = d.Set("auto_negotiation", strconv.FormatBool(PIJson["auto-negotiation"].(bool)))
	} else {
		_ = d.Set("auto_negotiation", PIJson["auto-negotiation"].(string))
	}

	return nil
}

func updatePhysicalInterface(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	payload := physicalInterfaceParseSchemaToMap(d)

	if d.HasChange("ipv6_address") || d.HasChange("ipv6_mask_length") {
		payload["ipv6-address"] = d.Get("ipv6_address")
		payload["ipv6-mask-length"] = d.Get("ipv6_mask_length")
	}

	if d.HasChange("ipv4_address") || d.HasChange("ipv4_mask_length") {
		payload["ipv4-address"] = d.Get("ipv4_address")
		payload["ipv4-mask-length"] = d.Get("ipv4_mask_length")
	}

	setNetworkRes, _ := client.ApiCall("set-physical-interface", payload, client.GetSessionID(), true, false)
	if !setNetworkRes.Success {
		return fmt.Errorf(setNetworkRes.ErrorMsg)
	}
	return readPhysicalInterface(d, m)
}

func deletePhysicalInterface(d *schema.ResourceData, m interface{}) error {
	d.SetId("") // Destroy resource
	return nil
}
