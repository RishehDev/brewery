package main

import (
	"brewery/infrastructure/cmd"
	"brewery/registry"
)

func main() {
	r := registry.NewRegistry()

	cmd.Execute(r.NewAppController())
}
