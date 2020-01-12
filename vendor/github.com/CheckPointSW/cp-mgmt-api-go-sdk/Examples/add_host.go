package Examples

import (
	api_go_sdk "github.com/Checkpoint/api_go_sdk/APIFiles"
	"fmt"
	"os"
)

func AddHost() {
	var api_server string
	var username string
	var password string

	fmt.Printf("Enter server IP address or hostname: ")
	fmt.Scanln(&api_server)

	fmt.Printf("Enter username: ")
	fmt.Scanln(&username)

	fmt.Printf("Enter password: ")
	fmt.Scanln(&password)

	args := api_go_sdk.APIClientArgs(443, "", "", api_server, "194.29.36.43", 8080, "", false, false, "deb.txt", api_go_sdk.WebContext, api_go_sdk.TimeOut, api_go_sdk.SleepTime)

	client := api_go_sdk.APIClient(args)

	fmt.Printf("Enter the name of the host: ")
	var host_name string
	fmt.Scanln(&host_name)

	fmt.Printf("Enter the ip of the host: ")
	var host_ip string
	fmt.Scanln(&host_ip)

	if x, _ := client.CheckFingerprint(); x == false {
		print("Could not get the server's fingerprint - Check connectivity with the server.\n")
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
		"name":       host_name,
		"ip-address": host_ip,
	}
	add_host_response, err := client.ApiCall("add-host", payload, client.GetSessionID(), false, true)

	if err != nil {
		print("error" + err.Error() + "\n")
	}

	if add_host_response.Success {
		print("The host: " + host_name + " has been added successfully\n")

		// publish the result
		payload = map[string]interface{}{}

		publish_res, err := client.ApiCall("publish", payload, client.GetSessionID(), true, true)
		if publish_res.Success {
			print("The changes were published successfully.\n")
		} else {
			print("Failed to publish the changes. \n" + err.Error())
		}
	} else {
		print("Failed to add the host: '" + host_name + "', Error:\n" + add_host_response.ErrorMsg)
	}

}
