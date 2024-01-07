package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strings"
)

func resourceManagementCMEAccountsGCP() *schema.Resource {
	return &schema.Resource{
		Create: createManagementCMEAccountsGCP,
		Update: updateManagementCMEAccountsGCP,
		Read:   readManagementCMEAccountsGCP,
		Delete: deleteManagementCMEAccountsGCP,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The account name.",
			},
			"project_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The project id.",
			},
			"credentials_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The credentials file.",
			},
			"credentials_data": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Base64 encoded string that represents the content of the credentials file.",
			},
			"deletion_tolerance": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The number of CME cycles to wait when the cloud provider does not return a GW until its deletion.",
			},
			"domain": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The account's domain name in MDS environment.",
			},
			"platform": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The platform of the account.",
			},
			"gw_configurations": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of GW configurations attached to the account",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func deleteManagementCMEAccountsGCP(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var name string
	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}

	log.Println("Delete cme GCP account - name = ", name)
	url := CmeApiPath + "/accounts/" + name

	res, err := client.ApiCall(url, nil, client.GetSessionID(), true, client.IsProxyUsed(), "DELETE")

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	data := res.GetData()
	if checkIfRequestFailed(data) {
		errMessage := buildErrorMessage(data)
		return fmt.Errorf(errMessage)
	}

	d.SetId("")
	return nil
}

func readManagementCMEAccountsGCP(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var name string

	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}

	log.Println("Read cme GCP account - name = ", name)
	url := CmeApiPath + "/accounts/" + name

	GCPAccountRes, err := client.ApiCall(url, nil, client.GetSessionID(), true, client.IsProxyUsed(), "GET")

	if err != nil {
		return fmt.Errorf(err.Error())
	}
	account := GCPAccountRes.GetData()
	if checkIfRequestFailed(account) {
		if cmeObjectNotFound(account) {
			d.SetId("")
			return nil
		}
		errMessage := buildErrorMessage(account)
		return fmt.Errorf(errMessage)
	}

	GCPAccount := account["result"].(map[string]interface{})

	_ = d.Set("name", GCPAccount["name"])

	_ = d.Set("project_id", GCPAccount["project_id"])

	credFile := strings.TrimPrefix(GCPAccount["credentials_file"].(string), "$FWDIR/conf/")
	_ = d.Set("credentials_file", credFile)

	_ = d.Set("credentials_data", GCPAccount["credentials_data"])

	_ = d.Set("deletion_tolerance", GCPAccount["deletion_tolerance"])

	_ = d.Set("domain", GCPAccount["domain"])

	_ = d.Set("platform", GCPAccount["platform"])

	_ = d.Set("gw_configurations", GCPAccount["gw_configurations"])

	return nil

}

func createManagementCMEAccountsGCP(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	payload := make(map[string]interface{})

	if v, ok := d.GetOk("project_id"); ok {
		payload["project_id"] = v.(string)
	}
	if v, ok := d.GetOk("credentials_file"); ok {
		payload["credentials_file"] = v.(string)
	}
	if v, ok := d.GetOk("credentials_data"); ok {
		payload["credentials_data"] = v.(string)
	}
	if v, ok := d.GetOk("deletion_tolerance"); ok {
		payload["deletion_tolerance"] = v.(int)
	}
	if v, ok := d.GetOk("domain"); ok {
		payload["domain"] = v.(string)
	}
	if v, ok := d.GetOk("name"); ok {
		payload["name"] = v.(string)
	}
	log.Println("Create cme GCP account - name = ", payload["name"])

	url := CmeApiPath + "/accounts/gcp"

	cmeAccountsRes, err := client.ApiCall(url, payload, client.GetSessionID(), true, client.IsProxyUsed())

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	data := cmeAccountsRes.GetData()
	if checkIfRequestFailed(data) {
		errMessage := buildErrorMessage(data)
		return fmt.Errorf(errMessage)
	}
	d.SetId("cme-gcp-account-" + d.Get("name").(string) + "-" + acctest.RandString(10))

	return readManagementCMEAccountsGCP(d, m)
}

func updateManagementCMEAccountsGCP(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	payload := make(map[string]interface{})

	if d.HasChange("project_id") {
		payload["project_id"] = d.Get("project_id")
	}
	if d.HasChange("credentials_file") {
		payload["credentials_file"] = d.Get("credentials_file")
	}
	if d.HasChange("credentials_data") {
		payload["credentials_data"] = d.Get("credentials_data")
	}
	if d.HasChange("deletion_tolerance") {
		payload["deletion_tolerance"] = d.Get("deletion_tolerance")
	}
	if d.HasChange("domain") {
		payload["domain"] = d.Get("domain")
	}

	var name string

	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}
	log.Println("Set cme GCP account - name = ", name)

	url := CmeApiPath + "/accounts/gcp/" + name
	cmeAccountsRes, err := client.ApiCall(url, payload, client.GetSessionID(), true, client.IsProxyUsed(), "PUT")

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	data := cmeAccountsRes.GetData()
	if checkIfRequestFailed(data) {
		errMessage := buildErrorMessage(data)
		return fmt.Errorf(errMessage)
	}

	return readManagementCMEAccountsGCP(d, m)
}
