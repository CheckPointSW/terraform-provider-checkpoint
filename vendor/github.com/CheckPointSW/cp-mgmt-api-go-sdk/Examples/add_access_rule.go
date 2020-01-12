package Examples

import (
	api_go_sdk "github.com/Checkpoint/api_go_sdk/APIFiles"
	"fmt"
	"os"
)

func AddAccess() {

	var api_server string
	var username string
	var password string

	fmt.Printf("Enter server IP address or hostname: ")
	fmt.Scanln(&api_server)

	fmt.Printf("Enter username: ")
	fmt.Scanln(&username)

	fmt.Printf("Enter password: ")
	fmt.Scanln(&password)

	args := api_go_sdk.APIClientArgs(443, "", "", api_server, "194.29.36.43", 8080, "", false, false, "deb.txt", api_go_sdk.WebContext,api_go_sdk.TimeOut,api_go_sdk.SleepTime)
	client := api_go_sdk.APIClient(args)

	fmt.Printf("Enter the name of the access rule: ")
	var rule_name string
	fmt.Scanln(&rule_name)

	if isFingerPrintTrusted, err := client.CheckFingerprint(); isFingerPrintTrusted == false || err != nil {

		if err != nil {
			fmt.Println(err.Error())
		} else {
			print("Could not get the server's fingerprint - Check connectivity with the server.\n")
		}
		os.Exit(1)
	}

	login_res, err := client.Login(username, password, false, "", false, "")
	if err != nil {
		print("Login error.\n")
		os.Exit(1)
	}



	if login_res.Success == false {
		print("Login failed:\n" + login_res.ErrorMsg)
		os.Exit(1)
	}

	// add a rule to the top of the "Network" layer
	payload := map[string]interface{}{
		"name":     rule_name,
		"layer":    "Network",
		"position": "top",
	}
	add_rule_response, err := client.ApiCall("add-access-rule", payload, client.GetSessionID(), false, true)

	if err != nil {
		print("error" + err.Error() + "\n")
	}

	if add_rule_response.Success {
		print("The rule: " + rule_name + " has been added successfully\n")

		// publish the result
		payload = map[string]interface{}{}

		publish_res, err := client.ApiCall("publish", payload, client.GetSessionID(), false, true)
		if publish_res.Success {
			print("The changes were published successfully.\n")
		} else {
			print("Failed to publish the changes. \n" + err.Error())
		}
	} else {
		print("Failed to add the access-rule: '" + rule_name + "', Error:\n" + add_rule_response.ErrorMsg)
	}

}
