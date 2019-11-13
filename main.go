package main

import (
	"./checkpoint"
	"github.com/hashicorp/terraform/plugin"
	//checkpoint "github.com/terraform-providers/terraform-provider-checkpoint"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: checkpoint.Provider,
	})
}
