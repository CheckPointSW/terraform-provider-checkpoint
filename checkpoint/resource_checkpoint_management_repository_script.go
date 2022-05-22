package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementRepositoryScript() *schema.Resource {
	return &schema.Resource{
		Create: createManagementRepositoryScript,
		Read:   readManagementRepositoryScript,
		Update: updateManagementRepositoryScript,
		Delete: deleteManagementRepositoryScript,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"script_body": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The entire content of the script.",
			},
			"script_body_base64": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The entire content of the script encoded in Base64.",
			},
			"script_body_base64_return": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The entire content of the script encoded in Base64.",
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

func createManagementRepositoryScript(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	repositoryScript := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		repositoryScript["name"] = v.(string)
	}

	if v, ok := d.GetOk("script_body"); ok {
		repositoryScript["script-body"] = v.(string)
	}

	if v, ok := d.GetOk("script_body_base64"); ok {
		repositoryScript["script-body-base64"] = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		repositoryScript["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		repositoryScript["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		repositoryScript["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		repositoryScript["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		repositoryScript["ignore-errors"] = v.(bool)
	}

	log.Println("Create RepositoryScript - Map = ", repositoryScript)

	addRepositoryScriptRes, err := client.ApiCall("add-repository-script", repositoryScript, client.GetSessionID(), true, false)
	if err != nil || !addRepositoryScriptRes.Success {
		if addRepositoryScriptRes.ErrorMsg != "" {
			return fmt.Errorf(addRepositoryScriptRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addRepositoryScriptRes.GetData()["uid"].(string))

	return readManagementRepositoryScript(d, m)
}

func readManagementRepositoryScript(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showRepositoryScriptRes, err := client.ApiCall("show-repository-script", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showRepositoryScriptRes.Success {
		if objectNotFound(showRepositoryScriptRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showRepositoryScriptRes.ErrorMsg)
	}

	repositoryScript := showRepositoryScriptRes.GetData()

	log.Println("Read RepositoryScript - Show JSON = ", repositoryScript)

	if v := repositoryScript["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := repositoryScript["script-body"]; v != nil {
		_ = d.Set("script_body_base64_return", v)
	}

	if v := repositoryScript["script-body-base64"]; v != nil {
		_ = d.Set("script_body_base64", v)
	}

	if repositoryScript["tags"] != nil {
		tagsJson, ok := repositoryScript["tags"].([]interface{})
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

	if v := repositoryScript["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := repositoryScript["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := repositoryScript["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := repositoryScript["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementRepositoryScript(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	repositoryScript := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		repositoryScript["name"] = oldName
		repositoryScript["new-name"] = newName
	} else {
		repositoryScript["name"] = d.Get("name")
	}

	if ok := d.HasChange("script_body"); ok {
		repositoryScript["script-body"] = d.Get("script_body")
	}

	if ok := d.HasChange("script_body_base64"); ok {
		repositoryScript["script-body-base64"] = d.Get("script_body_base64")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			repositoryScript["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			repositoryScript["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		repositoryScript["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		repositoryScript["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		repositoryScript["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		repositoryScript["ignore-errors"] = v.(bool)
	}

	log.Println("Update RepositoryScript - Map = ", repositoryScript)

	updateRepositoryScriptRes, err := client.ApiCall("set-repository-script", repositoryScript, client.GetSessionID(), true, false)
	if err != nil || !updateRepositoryScriptRes.Success {
		if updateRepositoryScriptRes.ErrorMsg != "" {
			return fmt.Errorf(updateRepositoryScriptRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementRepositoryScript(d, m)
}

func deleteManagementRepositoryScript(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	repositoryScriptPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete RepositoryScript")

	deleteRepositoryScriptRes, err := client.ApiCall("delete-repository-script", repositoryScriptPayload, client.GetSessionID(), true, false)
	if err != nil || !deleteRepositoryScriptRes.Success {
		if deleteRepositoryScriptRes.ErrorMsg != "" {
			return fmt.Errorf(deleteRepositoryScriptRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
