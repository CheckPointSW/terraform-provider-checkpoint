package Examples

import (
	api_go_sdk "github.com/Checkpoint/api_go_sdk/APIFiles"
	"fmt"
	"os"
)

func DupIp() {

	fmt.Printf("Enter server IP address or hostname: ")
	var apiServer string
	fmt.Scanln(&apiServer)

	fmt.Printf("Enter username: ")
	var username string
	fmt.Scanln(&username)

	fmt.Printf("Enter password: ")
	var pass string
	fmt.Scanln(&pass)

	args := api_go_sdk.APIClientArgs(443, "", "", apiServer, "", -1, "", false, false, "deb.txt", api_go_sdk.WebContext, api_go_sdk.TimeOut, api_go_sdk.SleepTime)

	client := api_go_sdk.APIClient(args)

	if x, _ := client.CheckFingerprint(); x == false {
		print("Could not get the server's fingerprint - Check connectivity with the server.\n")
		os.Exit(1)
	}

	loginRes, err := client.Login(username, pass, false, "", false, "")
	if err != nil {
		print("Login error.\n")
		os.Exit(1)
	}

	if loginRes.Success == false {
		print("Login failed:\n" + loginRes.ErrorMsg)
		os.Exit(1)
	}

	showHostsRes, err := client.ApiQuery("show-hosts", "full", "objects", false, map[string]interface{}{})

	if err != nil {
		print("Failed to retrieve the hosts\n")
		return
	}

	// objDictionary - for a given IP address, get an array of hosts (name, unique-ID) that use this IP address.
	var objDictionary = map[string][]map[string]string{}

	//dupIpSlice - a collection of the duplicate IP addresses in all the host objects.
	var dupIpSlice []string

	for _, host := range showHostsRes.GetData(){
		ipaddr := host.(map[string]interface{})["ipv4-address"].(string)
		if ipaddr == "" {
			print(host.(map[string]interface{})["name"].(string) + " has no IPv4 address. Skipping...")
			continue
		}
		hostData := map[string]string{"name": host.(map[string]interface{})["name"].(string), "uid": host.(map[string]interface{})["uid"].(string)}

		//ok will be set to true if ipaddr is actually in objDictionary
		if _, ok := objDictionary[ipaddr]; ok {
			ipExists := false
			for _, ip := range dupIpSlice {
				if ip == ipaddr{
					ipExists = true
					break
				}
			}
			if !ipExists {
				dupIpSlice = append(dupIpSlice, ipaddr)
			}
			objDictionary[ipaddr] = append(objDictionary[ipaddr], hostData)
		} else {

			objDictionary[ipaddr] = []map[string]string{}
			objDictionary[ipaddr] = append(objDictionary[ipaddr], hostData)
		}

	}

	//print list of duplicate IP addresses to the console
	fmt.Println("List of Duplicate IP addresses: ")
	fmt.Println("-------------------------------")

	if len(dupIpSlice) == 0 {
		fmt.Println("No hosts with duplicate IP addresses")
	}

	//for every duplicate ip - print hosts with that ip:
	for _, dup_ip :=  range dupIpSlice {

		fmt.Println("\nIP Address: " + dup_ip + "")
		fmt.Println("-------------------------------")

		for _, hostData := range objDictionary[dup_ip]{

			fmt.Println("host name: " + hostData["name"] + " host uid: " + hostData["uid"])
			//fmt.Println(hostData[1])

		}

	}


}
