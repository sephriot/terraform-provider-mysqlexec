package main

import (
"github.com/hashicorp/terraform/plugin"
"github.com/hashicorp/terraform/terraform"
"github.com/sephriot/terraform-provider-mysqlexec/mysqlexec"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return mysqlexec.Provider()
		},
	})
}