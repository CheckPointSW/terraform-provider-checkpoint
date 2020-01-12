package main

import (
	"github.com/Checkpoint/api_go_sdk/Examples"
	"fmt"
	"os"
)

func main(){

	switch os.Args[1] {
	case "discard":
		Examples.DiscardSessions()
	case "rule":
		Examples.AddAccess()
	case "add_host":
		Examples.AddHost()
	case "show_hosts":
		Examples.ShowHosts()
	case "dup_ip":
		Examples.DupIp()
	default:
		fmt.Println("Operation not supported. Optional operations are rule/discard/add_host/show_hosts/dup_ip")
	}
}
