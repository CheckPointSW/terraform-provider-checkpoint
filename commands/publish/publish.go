package main

import (
	"fmt"
	"github.com/CheckPointSW/terraform-provider-checkpoint/commands"
	"os"
)

func main() {
	apiClient, err := commands.InitClient()
	if err != nil {
		fmt.Println("Publish error: " + err.Error())
		os.Exit(1)
	}

	publishRes, err := apiClient.ApiCall("publish", map[string]interface{}{}, apiClient.GetSessionID(), true, false)
	if err != nil {
		fmt.Println("Publish error: " + err.Error())
		os.Exit(1)
	}

	taskId := commands.ResolveTaskId(publishRes.GetData())

	if !publishRes.Success {
		errMsg := fmt.Sprintf("Publish failed: %s.", publishRes.ErrorMsg)
		if taskId != nil {
			errMsg += fmt.Sprintf(" task-id [%s]", taskId)
		}
		fmt.Println(errMsg)
		os.Exit(1)
	}

	fmt.Println(fmt.Sprintf("Publish finished successfully. task-id [%s]", taskId))
}
