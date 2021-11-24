package checkpoint

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementAwsDataCenterServer() *schema.Resource {
	return &schema.Resource{
		Create: createManagementAwsDataCenterServer,
		Read:   readManagementAwsDataCenterServer,
		Update: updateManagementAwsDataCenterServer,
		Delete: deleteManagementAwsDataCenterServer,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name. Must be unique in the domain.",
			},
			"authentication_method": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "user-authentication\nUses the Access keys to authenticate.\nrole-authentication\nUses the AWS IAM role to authenticate.\nThis option requires the Security Management Server be deployed in AWS and has an IAM Role.",
			},
			"access_key_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Access key ID for the AWS account.\nRequired for authentication-method: user-authentication.",
			},
			"secret_access_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Secret access key for the AWS account.\nRequired for authentication-method: user-authentication.",
			},
			"region": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Select the AWS region.",
			},
			"enable_sts_assume_role": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enables the STS Assume Role option. After it is enabled, the sts-role field is mandatory, whereas the sts-external-id is optional.",
				Default:     false,
			},
			"sts_role": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The STS RoleARN of the role to be assumed.\nRequired for enable-sts-assume-role: true.",
			},
			"sts_external_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "An optional STS External-Id to use when assuming the role.",
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

func createManagementAwsDataCenterServer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	awsDataCenterServer := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		awsDataCenterServer["name"] = v.(string)
	}

	awsDataCenterServer["type"] = "aws"

	if v, ok := d.GetOk("authentication_method"); ok {
		awsDataCenterServer["authentication-method"] = v.(string)
	}

	if v, ok := d.GetOk("access_key_id"); ok {
		awsDataCenterServer["access-key-id"] = v.(string)
	}

	if v, ok := d.GetOk("secret_access_key"); ok {
		awsDataCenterServer["secret-access-key"] = v.(string)
	}

	if v, ok := d.GetOk("region"); ok {
		awsDataCenterServer["region"] = v.(string)
	}

	if v, ok := d.GetOk("enable_sts_assume_role"); ok {
		awsDataCenterServer["enable-sts-assume-role"] = v.(string)
	}

	if v, ok := d.GetOk("sts_role"); ok {
		awsDataCenterServer["sts-role"] = v.(string)
	}

	if v, ok := d.GetOk("sts_external_id"); ok {
		awsDataCenterServer["custom-value"] = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		awsDataCenterServer["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		awsDataCenterServer["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		awsDataCenterServer["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		awsDataCenterServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		awsDataCenterServer["ignore-errors"] = v.(bool)
	}

	log.Println("Create awsDataCenterServer - Map = ", awsDataCenterServer)

	addAwsDataCenterServerRes, err := client.ApiCall("add-data-center-server", awsDataCenterServer, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !addAwsDataCenterServerRes.Success {
		if addAwsDataCenterServerRes.ErrorMsg != "" {
			return fmt.Errorf(addAwsDataCenterServerRes.ErrorMsg)
		}
		msg := createTaskFailMessage("add-data-center-server", addAwsDataCenterServerRes.GetData())
		return fmt.Errorf(msg)
	}
	payload := map[string]interface{}{
		"name": awsDataCenterServer["name"],
	}
	showAwsDataCenterServerRes, err := client.ApiCall("show-data-center-server", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showAwsDataCenterServerRes.Success {
		return fmt.Errorf(showAwsDataCenterServerRes.ErrorMsg)
	}
	d.SetId(showAwsDataCenterServerRes.GetData()["uid"].(string))
	return readManagementAwsDataCenterServer(d, m)
}

func readManagementAwsDataCenterServer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showAwsDataCenterServerRes, err := client.ApiCall("show-data-center-server", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showAwsDataCenterServerRes.Success {
		if objectNotFound(showAwsDataCenterServerRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showAwsDataCenterServerRes.ErrorMsg)
	}
	awsDataCenterServer := showAwsDataCenterServerRes.GetData()

	if v := awsDataCenterServer["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if awsDataCenterServer["properties"] != nil {
		propsJson, ok := awsDataCenterServer["properties"].([]interface{})
		if ok {
			for _, prop := range propsJson {
				propMap := prop.(map[string]interface{})
				propName := strings.ReplaceAll(propMap["name"].(string), "-", "_")
				propValue := propMap["value"]
				if propName == "enable_sts_assume_role" {
					propValue, _ = strconv.ParseBool(propValue.(string))
				}
				_ = d.Set(propName, propValue)
			}
		}
	}

	if awsDataCenterServer["tags"] != nil {
		tagsJson, ok := awsDataCenterServer["tags"].([]interface{})
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

	if v := awsDataCenterServer["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := awsDataCenterServer["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := awsDataCenterServer["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := awsDataCenterServer["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementAwsDataCenterServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	awsDataCenterServer := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		awsDataCenterServer["name"] = oldName
		awsDataCenterServer["new-name"] = newName
	} else {
		awsDataCenterServer["name"] = d.Get("name")
	}

	if ok := d.HasChange("secret_access_key"); ok {
		awsDataCenterServer["secret-access-key"] = d.Get("secret_access_key")
	}

	if ok := d.HasChange("authentication_method"); ok {
		awsDataCenterServer["authentication-method"] = d.Get("authentication_method")
		if awsDataCenterServer["authentication-method"] == "user-authentication" {
			awsDataCenterServer["secret-access-key"] = d.Get("secret_access_key")
		}
	}

	if ok := d.HasChange("access_key_id"); ok {
		awsDataCenterServer["access-key-id"] = d.Get("access_key_id")
	}

	if ok := d.HasChange("region"); ok {
		awsDataCenterServer["region"] = d.Get("region")
	}

	if ok := d.HasChange("enable_sts_assume_role"); ok {
		awsDataCenterServer["enable-sts-assume-role"] = d.Get("enable_sts_assume_role")
	}

	if ok := d.HasChange("sts_role"); ok {
		awsDataCenterServer["sts-role"] = d.Get("sts_role")
	}

	if ok := d.HasChange("sts_external_id"); ok {
		awsDataCenterServer["sts-external-id"] = d.Get("sts_external_id")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			awsDataCenterServer["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			awsDataCenterServer["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		awsDataCenterServer["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		awsDataCenterServer["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		awsDataCenterServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		awsDataCenterServer["ignore-errors"] = v.(bool)
	}

	log.Println("Update awsDataCenterServer - Map = ", awsDataCenterServer)

	updateAwsDataCenterServerRes, err := client.ApiCall("set-data-center-server", awsDataCenterServer, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !updateAwsDataCenterServerRes.Success {
		if updateAwsDataCenterServerRes.ErrorMsg != "" {
			return fmt.Errorf(updateAwsDataCenterServerRes.ErrorMsg)
		}
		msg := createTaskFailMessage("set-data-center-server", updateAwsDataCenterServerRes.GetData())
		return fmt.Errorf(msg)
	}

	return readManagementAwsDataCenterServer(d, m)
}

func deleteManagementAwsDataCenterServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	awsDataCenterServerPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete awsDataCenterServer")

	deleteAwsDataCenterServerRes, err := client.ApiCall("delete-data-center-server", awsDataCenterServerPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteAwsDataCenterServerRes.Success {
		if deleteAwsDataCenterServerRes.ErrorMsg != "" {
			return fmt.Errorf(deleteAwsDataCenterServerRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
