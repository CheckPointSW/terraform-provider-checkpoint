package Examples

import (
	api_go_sdk "github.com/Checkpoint/api_go_sdk/APIFiles"
	"fmt"
	"os"
)

func ShowHosts() {

	var apiServer string
	var username string
	var password string



	fmt.Printf("Enter server IP address or hostname: ")
	fmt.Scanln(&apiServer)

	fmt.Printf("Enter username: ")
	fmt.Scanln(&username)

	fmt.Printf("Enter password: ")
	fmt.Scanln(&password)




	args := api_go_sdk.APIClientArgs(443, "", "", apiServer, "", -1, "", false, false, "deb.txt", api_go_sdk.WebContext, api_go_sdk.TimeOut, api_go_sdk.SleepTime)

	client := api_go_sdk.APIClient(args)


	if x, _ := client.CheckFingerprint(); !x {
		print("Could not get the server's fingerprint - Check connectivity with the server.\n")
		os.Exit(1)
	}

	loginRes, err := client.Login(username, password, false, "", false, "")
	if err != nil {
		print("Login error.\n")
		os.Exit(1)
	}

	if !loginRes.Success {
		print("Login failed:\n" + loginRes.ErrorMsg)
		os.Exit(1)
	}



	showHostsRes,err2 := client.ApiQuery("show-hosts", "full", "objects", false, map[string]interface{}{})


	if err2 != nil {
		print("Failed to retrieve the hosts\n")
		return
	}



	//fmt.Println(show_sessions_res.GetData())
	for _,sessionObj := range showHostsRes.GetData(){
		fmt.Println("-------------\n--------------")
		fmt.Println(sessionObj.(map[string]interface{})["name"].(string))
		fmt.Println(sessionObj.(map[string]interface{})["ipv4-address"].(string))
	}

}
