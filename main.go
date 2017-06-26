package main

import "github.com/hashicorp/terraform/plugin"
import "github.com/xiaopal/terraform-provider-npc/npc"

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: npc.Provider})
}
