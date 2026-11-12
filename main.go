package main

import (
	"github.com/CheckPointSW/terraform-provider-checkpoint/v3/checkpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: checkpoint.Provider,
	})
}
