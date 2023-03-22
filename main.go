/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"

	"github.com/bbruun/galt/cli"
	"github.com/bbruun/galt/cmd"
	"github.com/bbruun/galt/messaging"
	"github.com/bbruun/galt/server"
)

var mc *messaging.Minions

func init() {
	log.Println("Galt - the fast way to manage your servers")

	messaging.MinionStateCollector = messaging.NewMinions()
	mc = messaging.MinionStateCollector
	server.Configuration = server.NewGaltServerConfig()

}

func main() {
	cli.Cli = cli.NewCli()

	cmd.Execute()
}
