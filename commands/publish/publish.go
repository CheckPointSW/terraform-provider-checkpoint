package main

import (
	"github.com/terraform-providers/terraform-provider-checkpoint/commands"
	"log"
)

func main() {

	apiClient, err := commands.InitClient()
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	publishRes, err := apiClient.ApiCall("publish", map[string]interface{}{}, apiClient.GetSessionID(),true,false)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	if !publishRes.Success {
		log.Fatalf("error: %s", publishRes.ErrorMsg)
	}

	log.Printf("published successfully")

}
