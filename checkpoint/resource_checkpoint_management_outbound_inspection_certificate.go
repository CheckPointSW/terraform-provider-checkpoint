package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementOutboundInspectionCertificate() *schema.Resource {
	return &schema.Resource{
		Create: createManagementOutboundInspectionCertificate,
		Read:   readManagementOutboundInspectionCertificate,
		Update: updateManagementOutboundInspectionCertificate,
		Delete: deleteManagementOutboundInspectionCertificate,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"issued_by": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The DN (Distinguished Name) of the certificate.",
			},
			"base64_password": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Password (encoded in Base64 with padding) for the certificate file.",
			},
			"base64_certificate": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Password (encoded in Base64 with padding) for the certificate file.",
			},
			"base64_public_certificate": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Password (encoded in Base64 with padding) for the certificate file.",
			},
			"valid_from": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The date, from which the certificate is valid. Format: YYYY-MM-DD.",
			},
			"valid_to": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The certificate expiration date. Format: YYYY-MM-DD.",
			},
			"is_default": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Is the certificate the default certificate.",
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

func createManagementOutboundInspectionCertificate(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	outboundInspectionCertificate := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		outboundInspectionCertificate["name"] = v.(string)
	}

	if v, ok := d.GetOk("issued_by"); ok {
		outboundInspectionCertificate["issued-by"] = v.(string)
	}

	if v, ok := d.GetOk("base64_password"); ok {
		outboundInspectionCertificate["base64-password"] = v.(string)
	}

	if v, ok := d.GetOk("valid_from"); ok {
		outboundInspectionCertificate["valid-from"] = v.(string)
	}

	if v, ok := d.GetOk("valid_to"); ok {
		outboundInspectionCertificate["valid-to"] = v.(string)
	}

	if v, ok := d.GetOkExists("is_default"); ok {
		outboundInspectionCertificate["is-default"] = v.(bool)
	}

	if v, ok := d.GetOk("tags"); ok {
		outboundInspectionCertificate["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		outboundInspectionCertificate["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		outboundInspectionCertificate["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		outboundInspectionCertificate["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		outboundInspectionCertificate["ignore-errors"] = v.(bool)
	}

	log.Println("Create OutboundInspectionCertificate - Map = ", outboundInspectionCertificate)

	addOutboundInspectionCertificateRes, err := client.ApiCall("add-outbound-inspection-certificate", outboundInspectionCertificate, client.GetSessionID(), true, false)
	if err != nil || !addOutboundInspectionCertificateRes.Success {
		if addOutboundInspectionCertificateRes.ErrorMsg != "" {
			return fmt.Errorf(addOutboundInspectionCertificateRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addOutboundInspectionCertificateRes.GetData()["uid"].(string))

	return readManagementOutboundInspectionCertificate(d, m)
}

func readManagementOutboundInspectionCertificate(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showOutboundInspectionCertificateRes, err := client.ApiCall("show-outbound-inspection-certificate", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showOutboundInspectionCertificateRes.Success {
		if objectNotFound(showOutboundInspectionCertificateRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showOutboundInspectionCertificateRes.ErrorMsg)
	}

	outboundInspectionCertificate := showOutboundInspectionCertificateRes.GetData()

	log.Println("Read OutboundInspectionCertificate - Show JSON = ", outboundInspectionCertificate)

	if v := outboundInspectionCertificate["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := outboundInspectionCertificate["issued-by"]; v != nil {
		_ = d.Set("issued_by", removeCnPrefix(v.(string)))
	}

	if v := outboundInspectionCertificate["base64-certificate"]; v != nil {
		_ = d.Set("base64_certificate", v)
	}

	if v := outboundInspectionCertificate["base64-public-certificate"]; v != nil {
		_ = d.Set("base64_public_certificate", v)
	}

	if v := outboundInspectionCertificate["valid-from"]; v != nil {
		dateStr, err := convertDateFormat(v.(string))
		if err != nil {
			return fmt.Errorf("failed to convert the value %s from field valid-from to format yyyy-mm-dd ", v)
		}
		_ = d.Set("valid_from", dateStr)
	}

	if v := outboundInspectionCertificate["valid-to"]; v != nil {
		dateStr, err := convertDateFormat(v.(string))
		if err != nil {
			return fmt.Errorf("failed to convert the value %s from field valid-to to format yyyy-mm-dd ", v)
		}
		_ = d.Set("valid_to", dateStr)
	}

	if v := outboundInspectionCertificate["is-default"]; v != nil {
		_ = d.Set("is_default", v)
	}

	if outboundInspectionCertificate["tags"] != nil {
		tagsJson, ok := outboundInspectionCertificate["tags"].([]interface{})
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

	if v := outboundInspectionCertificate["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := outboundInspectionCertificate["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := outboundInspectionCertificate["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := outboundInspectionCertificate["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementOutboundInspectionCertificate(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	outboundInspectionCertificate := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		outboundInspectionCertificate["name"] = oldName
		outboundInspectionCertificate["new-name"] = newName
	} else {
		outboundInspectionCertificate["name"] = d.Get("name")
	}

	if ok := d.HasChange("issued_by"); ok {
		outboundInspectionCertificate["issued-by"] = d.Get("issued_by")
	}

	if ok := d.HasChange("base64_password"); ok {
		outboundInspectionCertificate["base64-password"] = d.Get("base64_password")
	}

	if ok := d.HasChange("valid_from"); ok {
		outboundInspectionCertificate["valid-from"] = d.Get("valid_from")
	}

	if ok := d.HasChange("valid_to"); ok {
		outboundInspectionCertificate["valid-to"] = d.Get("valid_to")
	}

	if v, ok := d.GetOkExists("is_default"); ok {
		outboundInspectionCertificate["is-default"] = v.(bool)
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			outboundInspectionCertificate["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			outboundInspectionCertificate["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		outboundInspectionCertificate["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		outboundInspectionCertificate["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		outboundInspectionCertificate["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		outboundInspectionCertificate["ignore-errors"] = v.(bool)
	}

	log.Println("Update OutboundInspectionCertificate - Map = ", outboundInspectionCertificate)

	updateOutboundInspectionCertificateRes, err := client.ApiCall("set-outbound-inspection-certificate", outboundInspectionCertificate, client.GetSessionID(), true, false)
	if err != nil || !updateOutboundInspectionCertificateRes.Success {
		if updateOutboundInspectionCertificateRes.ErrorMsg != "" {
			return fmt.Errorf(updateOutboundInspectionCertificateRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementOutboundInspectionCertificate(d, m)
}

func deleteManagementOutboundInspectionCertificate(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	outboundInspectionCertificatePayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete OutboundInspectionCertificate")

	deleteOutboundInspectionCertificateRes, err := client.ApiCall("delete-outbound-inspection-certificate", outboundInspectionCertificatePayload, client.GetSessionID(), true, false)
	if err != nil || !deleteOutboundInspectionCertificateRes.Success {
		if deleteOutboundInspectionCertificateRes.ErrorMsg != "" {
			return fmt.Errorf(deleteOutboundInspectionCertificateRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
