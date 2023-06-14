package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementThreatLayer() *schema.Resource {
	return &schema.Resource{
		Create: createManagementThreatLayer,
		Read:   readManagementThreatLayer,
		Update: updateManagementThreatLayer,
		Delete: deleteManagementThreatLayer,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name. Must be unique in the domain.",
			},
			"add_default_rule": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Indicates whether to include a default rule in the new layer.",
			},
			"tags": &schema.Schema{
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"color": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "black",
				Description: "Color of the object. Should be one of existing colors.",
			},
			"comments": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comments string.",
			},
			"ignore_warnings": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Apply changes ignoring warnings.",
			},
			"ignore_errors": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.",
			},
			"ips_layer": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "N/A",
			},
			"parent_layer": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "N/A",
			},
		},
	}
}

func createManagementThreatLayer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	threatLayer := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		threatLayer["name"] = v.(string)
	}

	if v, ok := d.GetOkExists("add_default_rule"); ok {
		threatLayer["add-default-rule"] = v.(bool)
	}

	if v, ok := d.GetOk("tags"); ok {
		threatLayer["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		threatLayer["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		threatLayer["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		threatLayer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		threatLayer["ignore-errors"] = v.(bool)
	}

	log.Println("Create Threat Layer - Map = ", threatLayer)

	addThreatLayerRes, err := client.ApiCall("add-threat-layer", threatLayer, client.GetSessionID(), true, client.IsProxyUsed())

	if err != nil || !addThreatLayerRes.Success {
		if addThreatLayerRes.ErrorMsg != "" {
			return fmt.Errorf(addThreatLayerRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addThreatLayerRes.GetData()["uid"].(string))

	return readManagementThreatLayer(d, m)
}

func readManagementThreatLayer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showThreatLayerRes, err := client.ApiCall("show-threat-layer", payload, client.GetSessionID(), true, client.IsProxyUsed())

	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showThreatLayerRes.Success {
		// Handle delete resource from other clients
		if objectNotFound(showThreatLayerRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showThreatLayerRes.ErrorMsg)
	}

	threatLayer := showThreatLayerRes.GetData()

	log.Println("Read Threat Layer - Show JSON = ", threatLayer)

	if v := threatLayer["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := threatLayer["add-default-rule"]; v != nil {
		_ = d.Set("add_default_rule", v)
	}

	if threatLayer["tags"] != nil {
		tagsJson := threatLayer["tags"].([]interface{})
		var tagsIds = make([]string, 0)
		if len(tagsJson) > 0 {
			// Create slice of tag names
			for _, tag := range tagsJson {
				tag := tag.(map[string]interface{})
				tagsIds = append(tagsIds, tag["name"].(string))
			}
		}
		_ = d.Set("tags", tagsIds)
	} else {
		_ = d.Set("tags", nil)
	}

	if v := threatLayer["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := threatLayer["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := threatLayer["ips-layer"]; v != nil {
		_ = d.Set("ips_layer", v)
	}

	if v := threatLayer["parent-layer"]; v != nil {
		_ = d.Set("parent_layer", v)
	}

	return nil
}

func updateManagementThreatLayer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	threatLayer := make(map[string]interface{})

	if d.HasChange("name") {
		oldName, newName := d.GetChange("name")
		threatLayer["name"] = oldName
		threatLayer["new-name"] = newName
	} else {
		threatLayer["name"] = d.Get("name")
	}

	if ok := d.HasChange("tags"); ok {
		if v, ok := d.GetOk("tags"); ok {
			threatLayer["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			threatLayer["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if d.HasChange("color") {
		threatLayer["color"] = d.Get("color")
	}

	if d.HasChange("comments") {
		threatLayer["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		threatLayer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		threatLayer["ignore-errors"] = v.(bool)
	}

	log.Println("Update Threat Layer - Map = ", threatLayer)
	updateThreatLayerRes, err := client.ApiCall("set-threat-layer", threatLayer, client.GetSessionID(), true, client.IsProxyUsed())

	if err != nil || !updateThreatLayerRes.Success {
		if updateThreatLayerRes.ErrorMsg != "" {
			return fmt.Errorf(updateThreatLayerRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementThreatLayer(d, m)
}

func deleteManagementThreatLayer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	threatLayerPayload := map[string]interface{}{
		"uid": d.Id(),
	}
	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		threatLayerPayload["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		threatLayerPayload["ignore-errors"] = v.(bool)
	}
	deleteThreatLayerRes, err := client.ApiCall("delete-threat-layer", threatLayerPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteThreatLayerRes.Success {
		if deleteThreatLayerRes.ErrorMsg != "" {
			return fmt.Errorf(deleteThreatLayerRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
