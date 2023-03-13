/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"

	"github.com/bbruun/galt/cmd"
	"github.com/bbruun/galt/server"
)

func init() {

	server.Configuration = server.NewGaltServerConfig()

}

func main() {

	log.Println("Starting server")
	cmd.Execute()
}
