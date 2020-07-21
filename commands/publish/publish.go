package main

import (
	"github.com/terraform-providers/terraform-provider-checkpoint/commands"
	"os"
)

func log(msg string) {
	_ = commands.LogToFile("publish.txt", msg)
}

func main() {

	apiClient, err := commands.InitClient()
	if err != nil {
		log("Publish error: " + err.Error())
		os.Exit(1)
	}

	publishRes, err := apiClient.ApiCall("publish", map[string]interface{}{}, apiClient.GetSessionID(), true, false)
	if err != nil {
		log("Publish error: " + err.Error())
		os.Exit(1)
	}
	if !publishRes.Success {
		log("Publish failed: " + publishRes.ErrorMsg)
		os.Exit(1)
	}

	log("Publish finished successfully")
}
