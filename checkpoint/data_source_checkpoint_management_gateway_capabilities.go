package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementSetGatewayCapabilities() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceManagementSetGatewayCapabilitiesRead,

		Schema: map[string]*schema.Schema{
			"hardware": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Check Point hardware.",
			},
			"platform": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Check Point gateway platform.",
			},
			"version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Gateway platform version.",
			},
			"restrictions": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Set of restrictions.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hardware": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the regulation.",
						},
						"platform": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Check Point gateway platform.",
						},
						"version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Gateway platform version.",
						},
					},
				},
			},
			"supported_platforms": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Set of restrictions.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"default": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the regulation.",
						},
						"platforms": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "List of Check Point gateway platforms.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"supported_blades": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Set of restrictions.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"management": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Management blades.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"default": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "N/A",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "N/A",
									},
									"readonly": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "N/A",
									},
								},
							},
						},
						"network_security": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Network Security blades.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"default": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "N/A",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "N/A",
									},
									"readonly": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "N/A",
									},
								},
							},
						},
						"threat_prevention": {
							Type:        schema.TypeList,
							Computed:    true,
							MaxItems:    1,
							Description: "Threat Prevention blades.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"autonomous": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "N/A",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"default": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "N/A",
												},
												"name": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "N/A",
												},
												"readonly": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "N/A",
												},
											},
										},
									},
									"custom": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "N/A",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"default": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "N/A",
												},
												"name": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "N/A",
												},
												"readonly": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "N/A",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"supported_firmware_platforms": {
				Type:        schema.TypeList,
				Computed:    true,
				MaxItems:    1,
				Description: "Supported firmware platforms according to restrictions.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"default": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Default gateway firmware platform.",
						},
						"firmware_platforms": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "List of gateway firmware platforms.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"supported_hardware": {
				Type:        schema.TypeList,
				Computed:    true,
				MaxItems:    1,
				Description: "Supported firmware platforms according to restrictions.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"default": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Default gateway firmware platform.",
						},
						"hardware": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "List of gateway firmware platforms.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"supported_versions": {
				Type:        schema.TypeList,
				Computed:    true,
				MaxItems:    1,
				Description: "Supported firmware platforms according to restrictions.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"default": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Default gateway firmware platform.",
						},
						"versions": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "List of gateway firmware platforms.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceManagementSetGatewayCapabilitiesRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{}

	if v, ok := d.GetOk("hardware"); ok {
		payload["hardware"] = v
	}
	if v, ok := d.GetOk("platform"); ok {
		payload["platform"] = v
	}
	if v, ok := d.GetOk("version"); ok {
		payload["version"] = v
	}

	showGatewayCapabilitiesRes, err := client.ApiCall("show-gateway-capabilities", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showGatewayCapabilitiesRes.Success {
		return fmt.Errorf(showGatewayCapabilitiesRes.ErrorMsg)
	}

	gatewayCapabilities := showGatewayCapabilitiesRes.GetData()

	log.Println("Read Gateway Capabilities - Show JSON = ", gatewayCapabilities)

	d.SetId("show-global-capabilities-" + acctest.RandString(10))

	if v := gatewayCapabilities["restrictions"]; v != nil {

		objMap := v.(map[string]interface{})

		restrictionsMapToAdd := make(map[string]interface{})

		if v := objMap["hardware"]; v != nil {
			restrictionsMapToAdd["hardware"] = v
		}
		if v := objMap["platform"]; v != nil {
			restrictionsMapToAdd["platform"] = v
		}
		if v := objMap["version"]; v != nil {
			restrictionsMapToAdd["version"] = v
		}
		_ = d.Set("restrictions", restrictionsMapToAdd)
	}

	if v := gatewayCapabilities["supported-platforms"]; v != nil {

		objMap := v.(map[string]interface{})

		mapToReturn := make(map[string]interface{})

		if v := objMap["default"]; v != nil {
			mapToReturn["default"] = v
		}
		if v := objMap["platforms"]; v != nil {

			mapToReturn["platforms"] = v
		}

		_ = d.Set("supported_platforms", []interface{}{mapToReturn})
	}

	if v := gatewayCapabilities["supported-blades"]; v != nil {

		innerMap := v.(map[string]interface{})

		supportedBladesMap := make(map[string]interface{})

		if v := innerMap["management"]; v != nil {

			managementList := v.([]interface{})

			if len(managementList) > 0 {

				var managementObjectsPayload []map[string]interface{}

				for i := range managementList {

					objMap := managementList[i].(map[string]interface{})

					mapToReturn := make(map[string]interface{})

					if v := objMap["default"]; v != nil {
						mapToReturn["default"] = v
					}
					if v := objMap["name"]; v != nil {
						mapToReturn["name"] = v
					}
					if v := objMap["readonly"]; v != nil {
						mapToReturn["readonly"] = v
					}
					managementObjectsPayload = append(managementObjectsPayload, mapToReturn)
				}
				supportedBladesMap["management"] = managementObjectsPayload
			}
		}

		if v := innerMap["network-security"]; v != nil {

			networkSecurityList := v.([]interface{})

			if len(networkSecurityList) > 0 {

				var networkSecurityObjectsPayload []map[string]interface{}

				for i := range networkSecurityList {

					objMap := networkSecurityList[i].(map[string]interface{})

					mapToReturn := make(map[string]interface{})

					if v := objMap["default"]; v != nil {
						mapToReturn["default"] = v
					}
					if v := objMap["name"]; v != nil {
						mapToReturn["name"] = v
					}
					if v := objMap["readonly"]; v != nil {
						mapToReturn["readonly"] = v
					}
					networkSecurityObjectsPayload = append(networkSecurityObjectsPayload, mapToReturn)
				}
				supportedBladesMap["network_security"] = networkSecurityObjectsPayload
			}
		}

		if v := innerMap["threat-prevention"]; v != nil {

			threatPreventionMapToReturn := make(map[string]interface{})

			threatPreventionMap := v.(map[string]interface{})

			if v := threatPreventionMap["autonomous"]; v != nil {

				autonomousList := v.([]interface{})

				if len(autonomousList) > 0 {

					var autonomousObjectsPayload []map[string]interface{}

					for i := range autonomousList {

						objMap := autonomousList[i].(map[string]interface{})

						mapToReturn := make(map[string]interface{})

						if v := objMap["default"]; v != nil {
							mapToReturn["default"] = v
						}
						if v := objMap["name"]; v != nil {
							mapToReturn["name"] = v
						}
						if v := objMap["readonly"]; v != nil {
							mapToReturn["readonly"] = v
						}
						autonomousObjectsPayload = append(autonomousObjectsPayload, mapToReturn)
					}
					threatPreventionMapToReturn["autonomous"] = autonomousObjectsPayload
				}

			}

			if v := threatPreventionMap["custom"]; v != nil {

				customList := v.([]interface{})

				if len(customList) > 0 {

					var customObjectsPayload []map[string]interface{}

					for i := range customList {

						objMap := customList[i].(map[string]interface{})

						mapToReturn := make(map[string]interface{})

						if v := objMap["default"]; v != nil {
							mapToReturn["default"] = v
						}
						if v := objMap["name"]; v != nil {
							mapToReturn["name"] = v
						}
						if v := objMap["readonly"]; v != nil {
							mapToReturn["readonly"] = v
						}
						customObjectsPayload = append(customObjectsPayload, mapToReturn)
					}
					threatPreventionMapToReturn["custom"] = customObjectsPayload
				}

			}

			supportedBladesMap["threat_prevention"] = []interface{}{threatPreventionMapToReturn}

		}

		_ = d.Set("supported_blades", []interface{}{supportedBladesMap})
	}

	if v := gatewayCapabilities["supported-firmware-platforms"]; v != nil {

		objMap := v.(map[string]interface{})

		mapToReturn := make(map[string]interface{})

		if v := objMap["default"]; v != nil {
			mapToReturn["default"] = v
		}
		if v := objMap["firmwarePlatforms"]; v != nil {
			mapToReturn["firmware_platforms"] = v
		}

		_ = d.Set("supported_firmware_platforms", []interface{}{mapToReturn})
	}

	if v := gatewayCapabilities["supported-hardware"]; v != nil {

		objMap := v.(map[string]interface{})

		mapToReturn := make(map[string]interface{})

		if v := objMap["default"]; v != nil {
			mapToReturn["default"] = v
		}
		if v := objMap["hardware"]; v != nil {
			mapToReturn["hardware"] = v
		}

		_ = d.Set("supported_hardware", []interface{}{mapToReturn})
	}

	if v := gatewayCapabilities["supported-versions"]; v != nil {

		objMap := v.(map[string]interface{})

		mapToReturn := make(map[string]interface{})

		if v := objMap["default"]; v != nil {
			mapToReturn["default"] = v
		}
		if v := objMap["versions"]; v != nil {
			mapToReturn["versions"] = v
		}

		_ = d.Set("supported_versions", []interface{}{mapToReturn})
	}

	return nil
}
