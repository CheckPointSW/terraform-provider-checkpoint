package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/terraform-providers/terraform-provider-checkpoint/checkpoint"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: checkpoint.Provider,
	})
}
