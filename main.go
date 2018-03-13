package main

import (
	"github.com/dot1q/terraform-provider-xenserver/xenserver"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return xenserver.Provider()
		},
	})
}
