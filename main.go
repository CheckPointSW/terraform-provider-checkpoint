package main

import (
	"github.com/hashicorp/terraform/plugin"
	"terraform-provider-checkpoint/checkpoint"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: checkpoint.Provider,
	})
}
