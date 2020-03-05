package main

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/terraform-providers/terraform-provider-checkpoint/commands"
	"log"
)




func addApplicationSite(client checkpoint.ApiClient) error {

	applicationSite := map[string]interface{}{
		"name" : "New Application Site 1",
		"primary-category" : "Social Networking",
		"description" : "My Application Site",
		"url-list" : []string{"www.cnet.com"},
	}

	addApplicationSiteRes, err := client.ApiCall("add-application-site", applicationSite, client.GetSessionID(), true, false)
	if err != nil || !addApplicationSiteRes.Success {
		if addApplicationSiteRes.ErrorMsg != "" {
			return fmt.Errorf(addApplicationSiteRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return nil
}


func addApplicationSiteCategory(client checkpoint.ApiClient) error {

	applicationSiteCategory := map[string]interface{}{
		"name" : "New Application Site Category 1",
	}

	addApplicationSiteRes, err := client.ApiCall("add-application-site-category", applicationSiteCategory, client.GetSessionID(), true, false)
	if err != nil || !addApplicationSiteRes.Success {
		if addApplicationSiteRes.ErrorMsg != "" {
			return fmt.Errorf(addApplicationSiteRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return nil
}

func addHost(client checkpoint.ApiClient) error {

	host := map[string]interface{}{
		"name" : "My Test Host 3",
		"ipv4-address": "10.0.0.1",
	}

	addHostRes, err := client.ApiCall("add-host", host, client.GetSessionID(), true, false)
	if err != nil || !addHostRes.Success {
		if addHostRes.ErrorMsg != "" {
			return fmt.Errorf(addHostRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	host1 := map[string]interface{}{
		"name" : "somehost",
		"ipv4-address": "10.0.0.2",
	}

	addHost1Res, err := client.ApiCall("add-host", host1, client.GetSessionID(), true, false)
	if err != nil || !addHost1Res.Success {
		if addHost1Res.ErrorMsg != "" {
			return fmt.Errorf(addHost1Res.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	host2 := map[string]interface{}{
		"name" : "New Host 1",
		"ipv4-address": "10.0.0.3",
	}

	addHost2Res, err := client.ApiCall("add-host", host2, client.GetSessionID(), true, false)
	if err != nil || !addHost2Res.Success {
		if addHost2Res.ErrorMsg != "" {
			return fmt.Errorf(addHost2Res.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return nil
}

func addGroup(client checkpoint.ApiClient) error {

	group := map[string]interface{}{
		"name" : "new group 1",
	}

	addGroupRes, err := client.ApiCall("add-group", group, client.GetSessionID(), true, false)
	if err != nil || !addGroupRes.Success {
		if addGroupRes.ErrorMsg != "" {
			return fmt.Errorf(addGroupRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	group1 := map[string]interface{}{
		"name" : "new group 2",
	}

	addGroup1Res, err := client.ApiCall("add-group", group1, client.GetSessionID(), true, false)
	if err != nil || !addGroup1Res.Success {
		if addGroup1Res.ErrorMsg != "" {
			return fmt.Errorf(addGroup1Res.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return nil
}

func addThreatLayer(client checkpoint.ApiClient) error {

	threatLayer := map[string]interface{}{
		"name" : "New Layer 1",
	}

	addThreatLayerRes, err := client.ApiCall("add-threat-layer", threatLayer, client.GetSessionID(), true, false)
	if err != nil || !addThreatLayerRes.Success {
		if addThreatLayerRes.ErrorMsg != "" {
			return fmt.Errorf(addThreatLayerRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return nil
}

func addThreatRule(client checkpoint.ApiClient) error {

	threatRule := map[string]interface{}{
		"layer" : "New Layer 1",
		"position" : "top",
		"name" : "First threat rule",
	}

	addThreatRuleRes, err := client.ApiCall("add-threat-rule", threatRule, client.GetSessionID(), true, false)
	if err != nil || !addThreatRuleRes.Success {
		if addThreatRuleRes.ErrorMsg != "" {
			return fmt.Errorf(addThreatRuleRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return nil
}

func addExceptionGroup(client checkpoint.ApiClient) error {

	exceptionGroup := map[string]interface{}{
		"name" : "exception_group_2",
	}

	addExceptionGroupeRes, err := client.ApiCall("add-exception-group", exceptionGroup, client.GetSessionID(), true, false)
	if err != nil || !addExceptionGroupeRes.Success {
		if addExceptionGroupeRes.ErrorMsg != "" {
			return fmt.Errorf(addExceptionGroupeRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return nil
}

func addHTTPSLayer(client checkpoint.ApiClient) error {

	HTTPSLayer := map[string]interface{}{
		"name" : "New Layer 2",
	}

	addHTTPSLayerRes, err := client.ApiCall("add-https-layer", HTTPSLayer, client.GetSessionID(), true, false)
	if err != nil || !addHTTPSLayerRes.Success {
		if addHTTPSLayerRes.ErrorMsg != "" {
			return fmt.Errorf(addHTTPSLayerRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return nil
}

//func addMds(client checkpoint.ApiClient) error {
//
//	mds := map[string]interface{}{
//		"name" : "MDM_Server",
//		"ipAddress4": "10.0.0.1",
//	}
//
//	addMdsRes, err := client.ApiCall("add-generic-object", mds, client.GetSessionID(), true, false)
//	if err != nil || !addMdsRes.Success {
//		if addMdsRes.ErrorMsg != "" {
//			return fmt.Errorf(addMdsRes.ErrorMsg)
//		}
//		return fmt.Errorf(err.Error())
//	}
//
//	return nil
//}

//func addDomain(client checkpoint.ApiClient) error {
//
//	domain := map[string]interface{}{
//		"name" : "New Application Site Category 1",
//	}
//
//	addApplicationSiteRes, err := client.ApiCall("add-application-site-category", applicationSiteCategory, client.GetSessionID(), true, false)
//	if err != nil || !addApplicationSiteRes.Success {
//		if addApplicationSiteRes.ErrorMsg != "" {
//			return fmt.Errorf(addApplicationSiteRes.ErrorMsg)
//		}
//		return fmt.Errorf(err.Error())
//	}
//
//	return nil
//}


func main() {

	apiClient, err := commands.InitClient()
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	err = addApplicationSite(apiClient)
	if err!= nil {
		log.Fatalf("error: %s", err)
	}

	err = addApplicationSiteCategory(apiClient)
	if err!= nil {
		log.Fatalf("error: %s", err)
	}

	err = addHost(apiClient)
	if err!= nil {
		log.Fatalf("error: %s", err)
	}

	err = addGroup(apiClient)
	if err!= nil {
		log.Fatalf("error: %s", err)
	}

	err = addThreatLayer(apiClient)
	if err!= nil {
		log.Fatalf("error: %s", err)
	}

	err = addThreatRule(apiClient)
	if err!= nil {
		log.Fatalf("error: %s", err)
	}

	err = addExceptionGroup(apiClient)
	if err!= nil {
		log.Fatalf("error: %s", err)
	}

	err = addHTTPSLayer(apiClient)
	if err!= nil {
		log.Fatalf("error: %s", err)
	}



	publishRes, err := apiClient.ApiCall("publish", map[string]interface{}{}, apiClient.GetSessionID(), true, false)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	if !publishRes.Success {
		log.Fatalf("error: %s", publishRes.ErrorMsg)
	}

	log.Printf("published successfully")

}
