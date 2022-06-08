package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementSmtpServer() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementSmtpServerRead,
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
			"port": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The SMTP port to use.",
			},
			"server": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SMTP server address.",
			},
			"authentication": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Does the mail server requires authentication.",
			},
			"encryption": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Encryption type.",
			},
			"password": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A password for the SMTP server.<br><font color=\"red\">Required only if</font> authentication is set to true.",
			},
			"username": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A username for the SMTP server.<br><font color=\"red\">Required only if</font> authentication is set to true.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
			"domains_to_process": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Indicates which domains to process the commands on. It cannot be used with the details-level full, must be run from the System Domain only and with ignore-warnings true. Valid values are: CURRENT_DOMAIN, ALL_DOMAINS_ON_THIS_SERVER.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceManagementSmtpServerRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
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

	if v := smtpServer["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

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

	return nil

}
