/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"

	"github.com/bbruun/galt/cmd"
	"github.com/bbruun/galt/logging"
	"github.com/bbruun/galt/server"
)

var log *logging.LoggerStruct

func main() {

	server.Configuration = server.NewGaltServerConfig()
	fmt.Printf("serverconfig: %+v\n", server.Configuration)
	logging.Logger = logging.NewLogger()
	log = logging.Logger
	log.Debug("Starting server")
	cmd.Execute()
}
