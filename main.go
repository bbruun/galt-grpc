/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"

	"github.com/bbruun/galt/cli"
	"github.com/bbruun/galt/cmd"
	"github.com/bbruun/galt/server"
)

func init() {
	log.Println("Galt - the fast way to manage your servers")

	server.Configuration = server.NewGaltServerConfig()

}

func main() {
	cli.Cli = cli.NewCli()

	cmd.Execute()
}
