package main

import (
	"github.com/danial.idham/benz/pkg/cmd"
	"github.com/danial.idham/benz/pkg/cmd/server"
)

func main() {
	err := server.NewCommand().Execute()
	cmd.CheckError(err)
}
