/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"strings"

	"github.com/bbruun/galt/server"
	"github.com/spf13/cobra"
)

var (
	configFile string
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the Galt server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if configFile != "" {
			server.ConfigFileLocations = append(server.ConfigFileLocations, configFile)
		}

		server.StartServer()
		log.Printf("starting server")
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().StringVarP(&configFile, "config", "c", strings.Join(server.ConfigFileLocations, " "), "Galt server config file")
}
