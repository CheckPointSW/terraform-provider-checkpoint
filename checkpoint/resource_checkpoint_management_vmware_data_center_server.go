package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strconv"
	"strings"
)

func resourceManagementVMwareDataCenterServer() *schema.Resource {
	return &schema.Resource{
		Create: createManagementVMwareDataCenterServer,
		Read:   readManagementVMwareDataCenterServer,
		Update: updateManagementVMwareDataCenterServer,
		Delete: deleteManagementVMwareDataCenterServer,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "VMWare object type. nsx or nsxt or globalnsxt or vcenter.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name. Must be unique in the domain.",
			},
			"hostname": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "IP Address or hostname of the vCenter server.",
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "Username of the vCenter server",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "Password of the vCenter server.",
			},
			"password_base64": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "Password of the vCenter server encoded in Base64.",
			},
			"certificate_fingerprint": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specify the SHA-1 or SHA-256 fingerprint of the Data Center Server's certificate.",
			},
			"unsafe_auto_accept": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "When set to false, the current Data Center Server's certificate should be trusted, either by providing the certificate-fingerprint argument or by relying on a previously trusted certificate of this hostname.\n\nWhen set to true, trust the current Data Center Server's certificate as-is.",
				Default:     false,
			},
			"policy_mode": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "For nsxt type only.\nWhen set to false, the Data Center Server will use Manager Mode APIs.\n\nWhen set to true, the Data Center Server will use Policy Mode APIs.",
			},
			"import_vms": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "For nsxt type only. When set to true, the Data Center Server will import Virtual Machines as well.\nThis feature will create additional API requests toward NSX-T manager\n\nNote: importing Virtual Machines can only be enabled while using Policy Mode APIs.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
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
				Description: "Apply changes ignoring warnings. By Setting this parameter to 'true' test connection failure will be ignored.",
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

func createManagementVMwareDataCenterServer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	vmwareDataCenterServer := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		vmwareDataCenterServer["name"] = v.(string)
	}

	if v, ok := d.GetOk("type"); ok {
		t := v.(string)
		if t != "nsx" && t != "nsxt" && t != "globalnsxt" && t != "vcenter" {
			return fmt.Errorf("invalid type: '%s'. Only 'nsx', 'nsxt', 'globalnsxt' or 'vcenter' are allowed", t)
		}
		vmwareDataCenterServer["type"] = t
		if t == "nsxt" {
			if pm, ok := d.GetOkExists("policy_mode"); ok {
				vmwareDataCenterServer["policy-mode"] = pm.(bool)
			}
			if pm, ok := d.GetOkExists("import_vms"); ok {
				vmwareDataCenterServer["import-vms"] = pm.(bool)
			}
		}
	}

	if v, ok := d.GetOk("hostname"); ok {
		vmwareDataCenterServer["hostname"] = v.(string)
	}

	if v, ok := d.GetOk("username"); ok {
		vmwareDataCenterServer["username"] = v.(string)
	}

	if v, ok := d.GetOk("password"); ok {
		vmwareDataCenterServer["password"] = v.(string)
	}

	if v, ok := d.GetOk("password_base64"); ok {
		vmwareDataCenterServer["password-base64"] = v.(string)
	}

	if v, ok := d.GetOk("certificate_fingerprint"); ok {
		vmwareDataCenterServer["certificate-fingerprint"] = v.(string)
	}

	if v, ok := d.GetOk("unsafe_auto_accept"); ok {
		vmwareDataCenterServer["unsafe-auto-accept"] = v.(bool)
	}

	if v, ok := d.GetOk("tags"); ok {
		vmwareDataCenterServer["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		vmwareDataCenterServer["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		vmwareDataCenterServer["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		vmwareDataCenterServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		vmwareDataCenterServer["ignore-errors"] = v.(bool)
	}

	log.Println("Create vmwareDataCenterServer - Map = ", vmwareDataCenterServer)

	addVMwareDataCenterServerRes, err := client.ApiCall("add-data-center-server", vmwareDataCenterServer, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !addVMwareDataCenterServerRes.Success {
		if addVMwareDataCenterServerRes.ErrorMsg != "" {
			return fmt.Errorf(addVMwareDataCenterServerRes.ErrorMsg)
		}
		msg := createTaskFailMessage("add-data-center-server", addVMwareDataCenterServerRes.GetData())
		return fmt.Errorf(msg)
	}
	payload := map[string]interface{}{
		"name": vmwareDataCenterServer["name"],
	}
	showVMwareDataCenterServerRes, err := client.ApiCall("show-data-center-server", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showVMwareDataCenterServerRes.Success {
		return fmt.Errorf(showVMwareDataCenterServerRes.ErrorMsg)
	}
	d.SetId(showVMwareDataCenterServerRes.GetData()["uid"].(string))
	return readManagementVMwareDataCenterServer(d, m)
}

func readManagementVMwareDataCenterServer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showVMwareDataCenterServerRes, err := client.ApiCall("show-data-center-server", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showVMwareDataCenterServerRes.Success {
		if objectNotFound(showVMwareDataCenterServerRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showVMwareDataCenterServerRes.ErrorMsg)
	}
	vmwareDataCenterServer := showVMwareDataCenterServerRes.GetData()

	if v := vmwareDataCenterServer["name"]; v != nil {
		_ = d.Set("name", v)
	}

	// "policy-mode" and "import-vms" in "properties"
	if vmwareDataCenterServer["properties"] != nil {
		propsJson, ok := vmwareDataCenterServer["properties"].([]interface{})
		if ok {
			for _, prop := range propsJson {
				propMap := prop.(map[string]interface{})
				propName := strings.ReplaceAll(propMap["name"].(string), "-", "_")
				propValue := propMap["value"]
				if propName == "unsafe_auto_accept" || propName == "policy_mode" || propName == "import_vms" {
					// all of them are boolean type so need to convert properly
					propValue, _ = strconv.ParseBool(propValue.(string))
				}
				_ = d.Set(propName, propValue)
			}
		}
	}

	if vmwareDataCenterServer["tags"] != nil {
		tagsJson, ok := vmwareDataCenterServer["tags"].([]interface{})
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

	if v := vmwareDataCenterServer["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := vmwareDataCenterServer["data-center-type"]; v != nil {
		_ = d.Set("type", v)
	}

	if v := vmwareDataCenterServer["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := vmwareDataCenterServer["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := vmwareDataCenterServer["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementVMwareDataCenterServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	vmwareDataCenterServer := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		vmwareDataCenterServer["name"] = oldName
		vmwareDataCenterServer["new-name"] = newName
	} else {
		vmwareDataCenterServer["name"] = d.Get("name")
	}

	if d.HasChange("hostname") {
		vmwareDataCenterServer["hostname"] = d.Get("hostname")
	}

	if d.HasChange("password") {
		vmwareDataCenterServer["password"] = d.Get("password")
	}

	if d.HasChange("password_base64") {
		vmwareDataCenterServer["password-base64"] = d.Get("password_base64")
	}

	if d.HasChange("username") {
		vmwareDataCenterServer["username"] = d.Get("username")
		if v := d.Get("password"); v != nil && v != "" {
			vmwareDataCenterServer["password"] = v
		}
		if v := d.Get("password_base64"); v != nil && v != "" {
			vmwareDataCenterServer["password-base64"] = v
		}
	}

	if d.HasChange("certificate_fingerprint") {
		vmwareDataCenterServer["certificate-fingerprint"] = d.Get("certificate_fingerprint")
	}

	if d.HasChange("unsafe_auto_accept") {
		vmwareDataCenterServer["unsafe-auto-accept"] = d.Get("unsafe_auto_accept")
	}

	if d.HasChange("policy_mode") {
		vmwareDataCenterServer["policy-mode"] = d.Get("policy_mode")
	}

	if d.HasChange("import_vms") {
		vmwareDataCenterServer["import-vms"] = d.Get("import_vms")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			vmwareDataCenterServer["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			vmwareDataCenterServer["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		vmwareDataCenterServer["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		vmwareDataCenterServer["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		vmwareDataCenterServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		vmwareDataCenterServer["ignore-errors"] = v.(bool)
	}

	log.Println("Update vmwareDataCenterServer - Map = ", vmwareDataCenterServer)

	updateVMwareDataCenterServerRes, err := client.ApiCall("set-data-center-server", vmwareDataCenterServer, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !updateVMwareDataCenterServerRes.Success {
		if updateVMwareDataCenterServerRes.ErrorMsg != "" {
			return fmt.Errorf(updateVMwareDataCenterServerRes.ErrorMsg)
		}
		msg := createTaskFailMessage("set-data-center-server", updateVMwareDataCenterServerRes.GetData())
		return fmt.Errorf(msg)
	}

	return readManagementVMwareDataCenterServer(d, m)
}

func deleteManagementVMwareDataCenterServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	vmwareDataCenterServerPayload := map[string]interface{}{
		"uid": d.Id(),
	}
	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		vmwareDataCenterServerPayload["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		vmwareDataCenterServerPayload["ignore-errors"] = v.(bool)
	}
	log.Println("Delete vmwareDataCenterServer")

	deleteVMwareDataCenterServerRes, err := client.ApiCall("delete-data-center-server", vmwareDataCenterServerPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteVMwareDataCenterServerRes.Success {
		if deleteVMwareDataCenterServerRes.ErrorMsg != "" {
			return fmt.Errorf(deleteVMwareDataCenterServerRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
