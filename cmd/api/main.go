package main

import (
	"github.com/DanialIdham/benz/pkg/cmd"
	"github.com/DanialIdham/benz/pkg/cmd/server"
)

func main() {
	err := server.NewCommand().Execute()
	cmd.CheckError(err)
}
