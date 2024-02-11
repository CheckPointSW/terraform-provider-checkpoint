package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementCMEAccountsAzure() *schema.Resource {
	return &schema.Resource{
		Create: createManagementCMEAccountsAzure,
		Update: updateManagementCMEAccountsAzure,
		Read:   readManagementCMEAccountsAzure,
		Delete: deleteManagementCMEAccountsAzure,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique account name for identification.",
			},
			"subscription": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Azure subscription ID.",
			},
			"directory_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Azure Active Directory tenant ID.",
			},
			"application_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The application ID with which the service principal is associated.",
			},
			"client_secret": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The service principal's client secret.",
				Sensitive:   true,
			},
			"deletion_tolerance": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     3,
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

func deleteManagementCMEAccountsAzure(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var name string
	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}

	log.Println("Delete cme Azure account - name = ", name)
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

func readManagementCMEAccountsAzure(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var name string

	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}

	log.Println("Read cme Azure account - name = ", name)
	url := CmeApiPath + "/accounts/" + name

	AzureAccountRes, err := client.ApiCall(url, nil, client.GetSessionID(), true, client.IsProxyUsed(), "GET")

	if err != nil {
		return fmt.Errorf(err.Error())
	}
	account := AzureAccountRes.GetData()
	if checkIfRequestFailed(account) {
		if cmeObjectNotFound(account) {
			d.SetId("")
			return nil
		}
		errMessage := buildErrorMessage(account)
		return fmt.Errorf(errMessage)
	}

	AzureAccount := account["result"].(map[string]interface{})

	_ = d.Set("name", AzureAccount["name"])

	_ = d.Set("subscription", AzureAccount["subscription"])

	_ = d.Set("directory_id", AzureAccount["directory_id"])

	_ = d.Set("application_id", AzureAccount["application_id"])

	_ = d.Set("deletion_tolerance", AzureAccount["deletion_tolerance"])

	_ = d.Set("domain", AzureAccount["domain"])

	_ = d.Set("platform", AzureAccount["platform"])

	_ = d.Set("gw_configurations", AzureAccount["gw_configurations"])

	return nil
}

func createManagementCMEAccountsAzure(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	payload := make(map[string]interface{})

	if v, ok := d.GetOk("subscription"); ok {
		payload["subscription"] = v.(string)
	}
	if v, ok := d.GetOk("directory_id"); ok {
		payload["directory_id"] = v.(string)
	}
	if v, ok := d.GetOk("application_id"); ok {
		payload["application_id"] = v.(string)
	}
	if v, ok := d.GetOk("client_secret"); ok {
		payload["client_secret"] = v.(string)
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
	log.Println("Create cme Azure account - name = ", payload["name"])

	url := CmeApiPath + "/accounts/azure"

	cmeAccountsRes, err := client.ApiCall(url, payload, client.GetSessionID(), true, client.IsProxyUsed())

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	data := cmeAccountsRes.GetData()
	if checkIfRequestFailed(data) {
		errMessage := buildErrorMessage(data)
		return fmt.Errorf(errMessage)
	}
	d.SetId("cme-azure-account-" + d.Get("name").(string) + "-" + acctest.RandString(10))

	return readManagementCMEAccountsAzure(d, m)
}

func updateManagementCMEAccountsAzure(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	payload := make(map[string]interface{})

	if d.HasChange("subscription") {
		payload["subscription"] = d.Get("subscription")
	}
	if d.HasChange("directory_id") {
		payload["directory_id"] = d.Get("directory_id")
	}
	if d.HasChange("application_id") {
		payload["application_id"] = d.Get("application_id")
	}
	if d.HasChange("client_secret") {
		payload["client_secret"] = d.Get("client_secret")
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
	log.Println("Set cme Azure account - name = ", name)

	url := CmeApiPath + "/accounts/azure/" + name
	cmeAccountsRes, err := client.ApiCall(url, payload, client.GetSessionID(), true, client.IsProxyUsed(), "PUT")

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	data := cmeAccountsRes.GetData()
	if checkIfRequestFailed(data) {
		errMessage := buildErrorMessage(data)
		return fmt.Errorf(errMessage)
	}

	return readManagementCMEAccountsAzure(d, m)
}
