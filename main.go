package main

import "github.com/cliffano/packer-post-processor-json-updater/json_updater"
import "github.com/hashicorp/packer/packer/plugin"

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterPostProcessor(new(json_updater.PostProcessor))
	server.Serve()
}
