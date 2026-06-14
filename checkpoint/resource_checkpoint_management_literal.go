package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"net"
	"strconv"
	"strings"
)

func resourceManagementLiteral() *schema.Resource {
	return &schema.Resource{
		Create: createManagementLiteral,
		Read:   readManagementLiteral,
		Delete: deleteManagementLiteral,
		Schema: map[string]*schema.Schema{
			"literal": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Literal represents IPv4 or IPv6 or Network CIDR or Service (tcp/<port number> or udp/<port number>) or DNS Domain (a name starting with '.')",
				ForceNew:    true,
			},
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object unique identifier",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object name",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object type",
			},
		},
	}
}

func createManagementLiteral(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	literal := d.Get("literal").(string)

	// Check if it's an IP address
	if ip := net.ParseIP(literal); ip != nil {
		response, err := handleHostLiteral(literal, client)
		if err != nil {
			return err
		}
		if !response.Success {
			return fmt.Errorf(response.ErrorMsg)
		}
		initLiteralData(response.GetData(), d)
		return readManagementLiteral(d, m)
	}

	// Check if it's a CIDR network
	if _, ipNet, err := net.ParseCIDR(literal); err == nil {
		// Get the subnet (network address)
		subnet := ipNet.IP.String()
		// Get the mask length
		maskLength, _ := ipNet.Mask.Size()
		response, err := handleNetworkLiteral(subnet, maskLength, client)
		if err != nil {
			return err
		}
		if !response.Success {
			return fmt.Errorf(response.ErrorMsg)
		}
		initLiteralData(response.GetData(), d)
		return readManagementLiteral(d, m)
	}

	// Check for tcp/<port> or udp/<port>
	if strings.HasPrefix(literal, "tcp/") || strings.HasPrefix(literal, "udp/") {
		parts := strings.SplitN(literal, "/", 2)
		if len(parts) == 2 {
			protocol := parts[0]
			portStr := parts[1]
			if port, err := strconv.Atoi(portStr); err == nil {
				response, err := handleServiceLiteral(protocol, port, client)
				if err != nil {
					return err
				}
				if !response.Success {
					return fmt.Errorf(response.ErrorMsg)
				}
				initLiteralData(response.GetData(), d)
				return readManagementLiteral(d, m)
			} else {
				return fmt.Errorf("invalid service port number")
			}
		} else {
			return fmt.Errorf("invalid service format. Expected tcp/<port> or udp/<port>")
		}
	}

	// Check for DNS Domain (a name starting with '.')
	if strings.HasPrefix(literal, ".") {
		response, err := handleDnsDomainLiteral(literal, client)
		if err != nil {
			return err
		}
		if !response.Success {
			return fmt.Errorf(response.ErrorMsg)
		}
		initLiteralData(response.GetData(), d)
		return readManagementLiteral(d, m)
	}

	return fmt.Errorf("invalid literal. The following types are supported: IPv4 or IPv6 or Network CIDR or Service (tcp/<port number> or udp/<port number>) or DNS Domain (a name starting with '.')")
}

func initLiteralData(object map[string]interface{}, d *schema.ResourceData) {
	if v := object["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := object["type"]; v != nil {
		_ = d.Set("type", v)
	}

	if v := object["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}
}

func handleServiceLiteral(protocol string, port int, client *checkpoint.ApiClient) (checkpoint.APIResponse, error) {
	showServicePayload := make(map[string]interface{})
	name := "service_" + protocol + "_" + strconv.Itoa(port)
	showServicePayload["name"] = name
	showCommand := "show-service-" + protocol

	addServicePayload := make(map[string]interface{})
	addServicePayload["name"] = name
	addServicePayload["ignore-warnings"] = true
	addServicePayload["port"] = strconv.Itoa(port)
	addCommand := "add-service-" + protocol

	response, err := handleObjectLiteral(showCommand, showServicePayload, addCommand, addServicePayload, client)
	if err != nil {
		return response, err
	}
	if !response.Success {
		return response, fmt.Errorf(response.ErrorMsg)
	}

	return response, nil
}

func handleNetworkLiteral(subnet string, maskLength int, client *checkpoint.ApiClient) (checkpoint.APIResponse, error) {
	showNetworkPayload := make(map[string]interface{})
	name := "network_" + subnet + "/" + strconv.Itoa(maskLength)
	showNetworkPayload["name"] = name

	addNetworkPayload := make(map[string]interface{})
	addNetworkPayload["name"] = name
	addNetworkPayload["ignore-warnings"] = true
	addNetworkPayload["subnet"] = subnet
	addNetworkPayload["mask-length"] = maskLength

	response, err := handleObjectLiteral("show-network", showNetworkPayload, "add-network", addNetworkPayload, client)
	if err != nil {
		return response, err
	}
	if !response.Success {
		return response, fmt.Errorf(response.ErrorMsg)
	}

	return response, nil
}

func handleHostLiteral(ip string, client *checkpoint.ApiClient) (checkpoint.APIResponse, error) {
	showHostPayload := make(map[string]interface{})
	name := "host_" + ip
	showHostPayload["name"] = name

	addHostPayload := make(map[string]interface{})
	addHostPayload["name"] = name
	addHostPayload["ignore-warnings"] = true
	addHostPayload["ip-address"] = ip

	response, err := handleObjectLiteral("show-host", showHostPayload, "add-host", addHostPayload, client)
	if err != nil {
		return response, err
	}
	if !response.Success {
		return response, fmt.Errorf(response.ErrorMsg)
	}

	return response, nil
}

func handleDnsDomainLiteral(domain string, client *checkpoint.ApiClient) (checkpoint.APIResponse, error) {
	showDnsDomainPayload := make(map[string]interface{})
	showDnsDomainPayload["name"] = domain

	addDnsDomainPayload := make(map[string]interface{})
	addDnsDomainPayload["name"] = domain
	addDnsDomainPayload["ignore-warnings"] = true
	addDnsDomainPayload["is-sub-domain"] = false

	response, err := handleObjectLiteral("show-dns-domain", showDnsDomainPayload, "add-dns-domain", addDnsDomainPayload, client)
	if err != nil {
		return response, err
	}
	if !response.Success {
		return response, fmt.Errorf(response.ErrorMsg)
	}

	return response, nil
}

func readManagementLiteral(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid":           d.Id(),
		"details-level": "full",
	}

	response, err := client.ApiCallSimple("show-object", payload)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !response.Success {
		if objectNotFound(response.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(response.ErrorMsg)
	}

	object := response.GetData()["object"].(map[string]interface{})

	if v := object["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := object["uid"]; v != nil {
		_ = d.Set("uid", v)
	}

	if objType := object["type"]; objType != nil {
		_ = d.Set("type", objType)
		literal := ""
		if objType == "host" {
			if v := object["ipv4-address"]; v != nil {
				literal = v.(string)
			} else if v := object["ipv6-address"]; v != nil {
				literal = v.(string)
			}
		} else if objType == "network" {
			if object["subnet4"] != nil && object["mask-length4"] != nil {
				literal = object["subnet4"].(string) + "/" + strconv.Itoa(int(object["mask-length4"].(float64)))
			} else if object["subnet6"] != nil && object["mask-length6"] != nil {
				literal = object["subnet6"].(string) + "/" + strconv.Itoa(int(object["mask-length6"].(float64)))
			}
		} else if objType == "service-tcp" {
			if v := object["port"]; v != nil {
				literal = "tcp/" + v.(string)
			}
		} else if objType == "service-udp" {
			if v := object["port"]; v != nil {
				literal = "udp/" + v.(string)
			}
		} else if objType == "dns-domain" {
			if v := object["name"]; v != nil {
				literal = v.(string)
			}
		} else {
			return fmt.Errorf("invalid object type. Expected to host, network, service-tcp, service-udp or dns-domain")
		}
		_ = d.Set("literal", literal)
	}

	return nil
}

func deleteManagementLiteral(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
