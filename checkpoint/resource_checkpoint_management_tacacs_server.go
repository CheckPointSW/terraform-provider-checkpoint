package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
    "log"
)

func resourceManagementTacacsServer() *schema.Resource {
    return &schema.Resource{
        Create: createManagementTacacsServer,
        Read:   readManagementTacacsServer,
        Update: updateManagementTacacsServer,
        Delete: deleteManagementTacacsServer,
        Schema: map[string]*schema.Schema{
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "Object name. Must be unique in the domain.",
            },
            "secret_key": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: "The server's secret key. Required only when \"server-type\" was selected to be \"TACACS+\".",
            },
            "server": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "The UID or Name of the host that is the TACACS Server.",
            },
            "encryption": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Is there a secret key defined on the server. Must be set true when \"server-type\" was selected to be \"TACACS+\".",
                Default:     false,
            },
            "priority": {
                Type:        schema.TypeInt,
                Optional:    true,
                Description: "The priority of the TACACS Server in case it is a member of a TACACS Group.",
                Default:     1,
            },
            "server_type": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: "Server type, TACACS or TACACS+.",
                Default:     "TACACS",
            },
            "service": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: "Server service, only relevant when \"server-type\" is TACACS.",
                Default:     "TACACS",
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

func createManagementTacacsServer(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)

    tacacsServer := make(map[string]interface{})

    if v, ok := d.GetOk("name"); ok {
        tacacsServer["name"] = v.(string)
    }

    if v, ok := d.GetOk("secret_key"); ok {
        tacacsServer["secret-key"] = v.(string)
    }

    if v, ok := d.GetOk("server"); ok {
        tacacsServer["server"] = v.(string)
    }

    if v, ok := d.GetOkExists("encryption"); ok {
        tacacsServer["encryption"] = v.(bool)
    }

    if v, ok := d.GetOk("priority"); ok {
        tacacsServer["priority"] = v.(int)
    }

    if v, ok := d.GetOk("server_type"); ok {
        tacacsServer["server-type"] = v.(string)
    }

    if v, ok := d.GetOk("service"); ok {
        tacacsServer["service"] = v.(string)
    }

    if v, ok := d.GetOk("tags"); ok {
        tacacsServer["tags"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOk("color"); ok {
        tacacsServer["color"] = v.(string)
    }

    if v, ok := d.GetOk("comments"); ok {
        tacacsServer["comments"] = v.(string)
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
        tacacsServer["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
        tacacsServer["ignore-errors"] = v.(bool)
    }

    log.Println("Create TacacsServer - Map = ", tacacsServer)

    addTacacsServerRes, err := client.ApiCall("add-tacacs-server", tacacsServer, client.GetSessionID(), true, client.IsProxyUsed())
    if err != nil || !addTacacsServerRes.Success {
        if addTacacsServerRes.ErrorMsg != "" {
            return fmt.Errorf(addTacacsServerRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }

    d.SetId(addTacacsServerRes.GetData()["uid"].(string))

    return readManagementTacacsServer(d, m)
}

func readManagementTacacsServer(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    payload := map[string]interface{}{
        "uid": d.Id(),
    }

    showTacacsServerRes, err := client.ApiCall("show-tacacs-server", payload, client.GetSessionID(), true, client.IsProxyUsed())
    if err != nil {
        return fmt.Errorf(err.Error())
    }
    if !showTacacsServerRes.Success {
        if objectNotFound(showTacacsServerRes.GetData()["code"].(string)) {
            d.SetId("")
            return nil
        }
        return fmt.Errorf(showTacacsServerRes.ErrorMsg)
    }

    tacacsServer := showTacacsServerRes.GetData()

    log.Println("Read TacacsServer - Show JSON = ", tacacsServer)

    if v := tacacsServer["name"]; v != nil {
        _ = d.Set("name", v)
    }

    if v := tacacsServer["secret-key"]; v != nil {
        _ = d.Set("secret_key", v)
    }

    if v := tacacsServer["server"]; v != nil {
        _ = d.Set("server", v)
    }

    if v := tacacsServer["encryption"]; v != nil {
        _ = d.Set("encryption", v)
    }

    if v := tacacsServer["priority"]; v != nil {
        _ = d.Set("priority", v)
    }

    if v := tacacsServer["server-type"]; v != nil {
        _ = d.Set("server_type", v)
    }

    if v := tacacsServer["service"]; v != nil {
        _ = d.Set("service", v)
    }

    if tacacsServer["tags"] != nil {
        tagsJson, ok := tacacsServer["tags"].([]interface{})
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

    if v := tacacsServer["color"]; v != nil {
        _ = d.Set("color", v)
    }

    if v := tacacsServer["comments"]; v != nil {
        _ = d.Set("comments", v)
    }

    if v := tacacsServer["ignore-warnings"]; v != nil {
        _ = d.Set("ignore_warnings", v)
    }

    if v := tacacsServer["ignore-errors"]; v != nil {
        _ = d.Set("ignore_errors", v)
    }

    return nil

}

func updateManagementTacacsServer(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    tacacsServer := make(map[string]interface{})

    if ok := d.HasChange("name"); ok {
        oldName, newName := d.GetChange("name")
        tacacsServer["name"] = oldName
        tacacsServer["new-name"] = newName
    } else {
        tacacsServer["name"] = d.Get("name")
    }

    if ok := d.HasChange("secret_key"); ok {
        tacacsServer["secret-key"] = d.Get("secret_key")
    }

    if ok := d.HasChange("server"); ok {
        tacacsServer["server"] = d.Get("server")
    }

    if v, ok := d.GetOkExists("encryption"); ok {
        tacacsServer["encryption"] = v.(bool)
    }

    if ok := d.HasChange("priority"); ok {
        tacacsServer["priority"] = d.Get("priority")
    }

    if ok := d.HasChange("server_type"); ok {
        tacacsServer["server-type"] = d.Get("server_type")
    }

    if ok := d.HasChange("service"); ok {
        tacacsServer["service"] = d.Get("service")
    }

    if d.HasChange("tags") {
        if v, ok := d.GetOk("tags"); ok {
            tacacsServer["tags"] = v.(*schema.Set).List()
        } else {
            oldTags, _ := d.GetChange("tags")
            tacacsServer["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
        }
    }

    if ok := d.HasChange("color"); ok {
        tacacsServer["color"] = d.Get("color")
    }

    if ok := d.HasChange("comments"); ok {
        tacacsServer["comments"] = d.Get("comments")
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
        tacacsServer["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
        tacacsServer["ignore-errors"] = v.(bool)
    }

    log.Println("Update TacacsServer - Map = ", tacacsServer)

    updateTacacsServerRes, err := client.ApiCall("set-tacacs-server", tacacsServer, client.GetSessionID(), true, client.IsProxyUsed())
    if err != nil || !updateTacacsServerRes.Success {
        if updateTacacsServerRes.ErrorMsg != "" {
            return fmt.Errorf(updateTacacsServerRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }

    return readManagementTacacsServer(d, m)
}

func deleteManagementTacacsServer(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    tacacsServerPayload := map[string]interface{}{
        "uid": d.Id(),
    }

    log.Println("Delete TacacsServer")

    deleteTacacsServerRes, err := client.ApiCall("delete-tacacs-server", tacacsServerPayload, client.GetSessionID(), true, false)
    if err != nil || !deleteTacacsServerRes.Success {
        if deleteTacacsServerRes.ErrorMsg != "" {
            return fmt.Errorf(deleteTacacsServerRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }
    d.SetId("")

    return nil
}
