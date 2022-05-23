package main

import (
	"fmt"
	"github.com/CheckPointSW/terraform-provider-checkpoint/commands"
	"os"
)

func main() {
	apiClient, err := commands.InitClient()
	if err != nil || len(os.Args) < 3 {
		fmt.Println("Approve Session error: " + err.Error())
		os.Exit(1)
	}
	if len(os.Args) < 2 {
		fmt.Println("no session uid or comments provided")
		os.Exit(1)
	}
	payload := make(map[string]interface{})
	payload["uid"] = os.Args[1]
	payload["comments"] = os.Args[2]
	rejectSessionRes, err := apiClient.ApiCall("reject-session", payload, apiClient.GetSessionID(), true, apiClient.IsProxyUsed())
	if err != nil {
		fmt.Println("Reject Session error: " + err.Error())
		os.Exit(1)
	}

	if !rejectSessionRes.Success {
		errMsg := fmt.Sprintf("Reject Session failed: %s.", rejectSessionRes.ErrorMsg)
		fmt.Println(errMsg)
		os.Exit(1)
	}

	fmt.Println(fmt.Sprintf("Reject Session finished successfully."))
}
