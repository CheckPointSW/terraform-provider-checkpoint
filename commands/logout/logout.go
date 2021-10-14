package main

import (
	"fmt"
	"github.com/CheckPointSW/terraform-provider-checkpoint/commands"
	"os"
)

func main() {
	apiClient, err := commands.InitClient()
	if err != nil {
		fmt.Println("logout error: " + err.Error())
		os.Exit(1)
	}

	logoutRes, err := apiClient.ApiCall("logout", map[string]interface{}{}, apiClient.GetSessionID(), true, false)
	if err != nil {
		fmt.Println("logout error: " + err.Error())
		os.Exit(1)
	}

	taskId := commands.ResolveTaskId(logoutRes.GetData())

	if !logoutRes.Success {
		errMsg := fmt.Sprintf("logout failed: %s.", logoutRes.ErrorMsg)
		if taskId != nil {
			errMsg += fmt.Sprintf(" task-id [%s]", taskId)
		}
		fmt.Println(errMsg)
		os.Exit(1)
	}

	fmt.Println(fmt.Sprintf("logout finished successfully. task-id [%s]", taskId))
}