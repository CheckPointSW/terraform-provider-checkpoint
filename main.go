package main

import (
	"github.com/hashicorp/terraform/plugin"
	chkp "terraform-provider-checkpoint/checkpoint"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: chkp.Provider,
	})
}
