package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementSmtpServer() *schema.Resource {
	return &schema.Resource{
		Create: createManagementSmtpServer,
		Read:   readManagementSmtpServer,
		Update: updateManagementSmtpServer,
		Delete: deleteManagementSmtpServer,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The SMTP port to use.",
			},
			"server": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The SMTP server address.",
			},
			"authentication": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Does the mail server requires authentication.",
				Default:     false,
			},
			"encryption": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Encryption type.",
				Default:     "none",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A password for the SMTP server.<br><font color=\"red\">Required only if</font> authentication is set to true.",
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A username for the SMTP server.<br><font color=\"red\">Required only if</font> authentication is set to true.",
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
			"domains_to_process": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Indicates which domains to process the commands on. It cannot be used with the details-level full, must be run from the System Domain only and with ignore-warnings true. Valid values are: CURRENT_DOMAIN, ALL_DOMAINS_ON_THIS_SERVER.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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

func createManagementSmtpServer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	smtpServer := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		smtpServer["name"] = v.(string)
	}

	if v, ok := d.GetOk("port"); ok {
		smtpServer["port"] = v.(int)
	}

	if v, ok := d.GetOk("server"); ok {
		smtpServer["server"] = v.(string)
	}

	if v, ok := d.GetOkExists("authentication"); ok {
		smtpServer["authentication"] = v.(bool)
	}

	if v, ok := d.GetOk("encryption"); ok {
		smtpServer["encryption"] = v.(string)
	}

	if v, ok := d.GetOk("password"); ok {
		smtpServer["password"] = v.(string)
	}

	if v, ok := d.GetOk("username"); ok {
		smtpServer["username"] = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		smtpServer["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		smtpServer["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		smtpServer["comments"] = v.(string)
	}

	if v, ok := d.GetOk("domains_to_process"); ok {
		smtpServer["domains-to-process"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		smtpServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		smtpServer["ignore-errors"] = v.(bool)
	}

	log.Println("Create SmtpServer - Map = ", smtpServer)

	addSmtpServerRes, err := client.ApiCall("add-smtp-server", smtpServer, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !addSmtpServerRes.Success {
		if addSmtpServerRes.ErrorMsg != "" {
			return fmt.Errorf(addSmtpServerRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addSmtpServerRes.GetData()["uid"].(string))

	return readManagementSmtpServer(d, m)
}

func readManagementSmtpServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showSmtpServerRes, err := client.ApiCall("show-smtp-server", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showSmtpServerRes.Success {
		if objectNotFound(showSmtpServerRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showSmtpServerRes.ErrorMsg)
	}

	smtpServer := showSmtpServerRes.GetData()

	log.Println("Read SmtpServer - Show JSON = ", smtpServer)

	if v := smtpServer["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := smtpServer["port"]; v != nil {
		_ = d.Set("port", v)
	}

	if v := smtpServer["server"]; v != nil {
		_ = d.Set("server", v)
	}

	if v := smtpServer["authentication"]; v != nil {
		_ = d.Set("authentication", v)
	}

	if v := smtpServer["encryption"]; v != nil {
		_ = d.Set("encryption", v)
	}

	if v := smtpServer["password"]; v != nil {
		_ = d.Set("password", v)
	}

	if v := smtpServer["username"]; v != nil {
		_ = d.Set("username", v)
	}

	if smtpServer["tags"] != nil {
		tagsJson, ok := smtpServer["tags"].([]interface{})
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

	if v := smtpServer["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := smtpServer["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if smtpServer["domains_to_process"] != nil {
		domainsToProcessJson, ok := smtpServer["domains_to_process"].([]interface{})
		if ok {
			domainsToProcessIds := make([]string, 0)
			if len(domainsToProcessJson) > 0 {
				for _, domains_to_process := range domainsToProcessJson {
					domains_to_process := domains_to_process.(map[string]interface{})
					domainsToProcessIds = append(domainsToProcessIds, domains_to_process["name"].(string))
				}
			}
			_ = d.Set("domains_to_process", domainsToProcessIds)
		}
	} else {
		_ = d.Set("domains_to_process", nil)
	}

	if v := smtpServer["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := smtpServer["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementSmtpServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	smtpServer := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		smtpServer["name"] = oldName
		smtpServer["new-name"] = newName
	} else {
		smtpServer["name"] = d.Get("name")
	}

	if ok := d.HasChange("port"); ok {
		smtpServer["port"] = d.Get("port")
	}

	if ok := d.HasChange("server"); ok {
		smtpServer["server"] = d.Get("server")
	}

	if v, ok := d.GetOkExists("authentication"); ok {
		smtpServer["authentication"] = v.(bool)
	}

	if ok := d.HasChange("encryption"); ok {
		smtpServer["encryption"] = d.Get("encryption")
	}

	if ok := d.HasChange("password"); ok {
		smtpServer["password"] = d.Get("password")
	}

	if ok := d.HasChange("username"); ok {
		smtpServer["username"] = d.Get("username")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			smtpServer["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			smtpServer["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		smtpServer["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		smtpServer["comments"] = d.Get("comments")
	}

	if d.HasChange("domains_to_process") {
		if v, ok := d.GetOk("domains_to_process"); ok {
			smtpServer["domains_to_process"] = v.(*schema.Set).List()
		} else {
			oldDomains_To_Process, _ := d.GetChange("domains_to_process")
			smtpServer["domains_to_process"] = map[string]interface{}{"remove": oldDomains_To_Process.(*schema.Set).List()}
		}
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		smtpServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		smtpServer["ignore-errors"] = v.(bool)
	}

	log.Println("Update SmtpServer - Map = ", smtpServer)

	updateSmtpServerRes, err := client.ApiCall("set-smtp-server", smtpServer, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !updateSmtpServerRes.Success {
		if updateSmtpServerRes.ErrorMsg != "" {
			return fmt.Errorf(updateSmtpServerRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementSmtpServer(d, m)
}

func deleteManagementSmtpServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	smtpServerPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete SmtpServer")

	deleteSmtpServerRes, err := client.ApiCall("delete-smtp-server", smtpServerPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteSmtpServerRes.Success {
		if deleteSmtpServerRes.ErrorMsg != "" {
			return fmt.Errorf(deleteSmtpServerRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
