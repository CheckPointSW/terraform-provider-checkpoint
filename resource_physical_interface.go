package main

import (
	chkp "api_go_sdk/APIFiles"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)


func resourcePhysicalInterface() *schema.Resource {
	return &schema.Resource{
		Create: createPhysicalInterface,
		Read:   readPhysicalInterface,
		Update: updatePhysicalInterface,
		Delete: deletePhysicalInterface,
		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString,
				Required: true,
				Description: "interface name",
			},
			"auto_negotiation": {
				Type: schema.TypeBool,
				Optional: true,
				Description: "Activating auto_negotiation will skip the speed and duplex configuration",
			},
			"comments": {
				Type: schema.TypeString,
				Optional: true,
				Description: "comments",
			},
			"duplex": {
				Type: schema.TypeString,
				Optional: true,
				Description: "Duplex is not relevant when 'auto_negotiation' is enabled.",
			},
			"enabled": {
				Type: schema.TypeBool,
				Optional: true,
				Description: "",
			},
			"ipv4_address": {
				Type: schema.TypeString,
				Optional: true,
				Description: "",
			},
			"monitor_mode": {
				Type: schema.TypeBool,
				Optional: true,
				Description: "",
			},
			"ipv6_autoconfig": {
				Type: schema.TypeBool,
				Optional: true,
				Description: "",
			},
			"mac_addr": {
				Type: schema.TypeString,
				Optional: true,
				Description: "",
			},
			"mtu": {
				Type: schema.TypeInt,
				Optional: true,
				Description: "",
			},
			"rx_ringsize": {
				Type:         schema.TypeInt,
				Optional:     true,
				Description:  "",
			},
			"ipv6_mask_length": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ipv6_address": {
				Type: schema.TypeString,
				Optional: true,
				Description: "",
			},
			"ipv4_mask_length": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"speed": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Speed is not relevant when 'auto_negotiation' is enabled",
			},
			"tx_ringsize": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func physicalInterfaceParseSchemaToMap(d *schema.ResourceData, createResource bool) map[string]interface{} {
	networkMap := make(map[string]interface{})

	if val, ok := d.GetOk("name"); ok {
		networkMap["name"] = val.(string)
	}
	if val, ok := d.GetOk("monitor_mode"); ok {
		networkMap["monitor-mode"] = val.(bool)
	}
	if val, ok := d.GetOk("duplex"); ok {
		networkMap["duplex"] = val.(string)
	}
	if val, ok := d.GetOk("ipv6_autoconfig"); ok {
		networkMap["ipv6-autoconfig"] = val.(bool)
	}

	if val, ok := d.GetOk("mac_addr"); ok {
		networkMap["mac-addr"] = val.(int)
	}
	if val, ok := d.GetOk("enabled"); ok {
		networkMap["enabled"] = val.(bool)
	}
	if val, ok := d.GetOk("comments"); ok {
		networkMap["comments"] = val.(string)
	}
	if val, ok := d.GetOk("mtu"); ok {
		networkMap["mtu"] = val.(int)
	}

	if val, ok := d.GetOk("rx_ringsize"); ok {
		networkMap["rx-ringsize"] = val.(int)
	}
	if val, ok := d.GetOk("ipv6_mask_length"); ok {
		networkMap["ipv6-mask-length"] = val.(int)
	}
	if val, ok := d.GetOk("ipv6_address"); ok {
		networkMap["ipv6-address"] = val.(string)
	}
	if val, ok := d.GetOk("ipv4_address"); ok {
		networkMap["ipv4-address"] = val.(string)
	}
	if val, ok := d.GetOk("ipv4_mask_length"); ok {
		networkMap["ipv4-mask-length"] = val.(int)
	}
	if val, ok := d.GetOk("speed"); ok {
		networkMap["speed"] = val.(string)
	}
	if val, ok := d.GetOk("tx_ringsize"); ok {
		networkMap["tx-ringsize"] = val.(int)
	}
	if val, ok := d.GetOk("auto_negotiation"); ok {
		networkMap["auto-negotiation"] = val.(bool)
	}

	if !createResource {
		// Call from updatePhysicalInterface
		// Remove attributes that cannot be set from map - Schema contain ADD + SET attr.
		delete(networkMap,"set-if-exists")
	}
	return networkMap
}

func createPhysicalInterface(d *schema.ResourceData, m interface{}) error {
	log.Println("Enter createPhysicalInterface...")
	client := m.(*chkp.ApiClient)
	payload := physicalInterfaceParseSchemaToMap(d, true)
	log.Println(payload)
	setPIRes, _ := client.ApiCall("set-physical-interface",payload,client.GetSessionID(),true,false)
	if !setPIRes.Success {
		return fmt.Errorf(setPIRes.ErrorMsg)
	}

	// Set Schema UID = Object key
	d.SetId(setPIRes.GetData()["name"].(string))

	log.Println("Exit createPhysicalInterface...")
	return readPhysicalInterface(d, m)
}

func readPhysicalInterface(d *schema.ResourceData, m interface{}) error {
	log.Println("Enter readPhysicalInterface...")
	client := m.(*chkp.ApiClient)
	payload := map[string]interface{}{
		"name": d.Get("name"),
	}
	showPIRes, _ := client.ApiCall("show-physical-interface",payload,client.GetSessionID(),true,false)
	if !showPIRes.Success {
		// Handle deletion of an object from other clients - Object not found
		if objectNotFound(showPIRes.GetData()["code"].(string)) {
			d.SetId("") // Destroy resource
			return nil
		}
		return fmt.Errorf(showPIRes.ErrorMsg)
	}
	PIJson := showPIRes.GetData()
	log.Println(PIJson)

	if _, ok := d.GetOk("name"); ok {
		_ = d.Set("name", PIJson["name"].(string))
	}

	log.Println("Exit readPhysicalInterface...")
	return nil
}

func updatePhysicalInterface(d *schema.ResourceData, m interface{}) error {
	log.Println("Enter updatePhysicalInterface...")
	client := m.(*chkp.ApiClient)
	payload := physicalInterfaceParseSchemaToMap(d, false)
	setNetworkRes, _ := client.ApiCall("set-physical-interface",payload,client.GetSessionID(),true,false)
	if !setNetworkRes.Success {
		return fmt.Errorf(setNetworkRes.ErrorMsg)
	}
	log.Println("Exit updatePhysicalInterface...")
	return readPhysicalInterface(d, m)
}

func deletePhysicalInterface(d *schema.ResourceData, m interface{}) error {
	log.Println("Enter deletePhysicalInterface...")
	d.SetId("") // Destroy resource
	log.Println("Exit deletePhysicalInterface...")
	return nil
}
