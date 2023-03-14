/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/bbruun/galt/cli"
	cmdRunner "github.com/bbruun/galt/cmd/cmdRun"

	"github.com/spf13/cobra"
)

var (
	minionsToTarget string
	cmdRun          bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "galt",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if minionsToTarget != "" {
			// Parse normal wildcards from CLI
			// https://www.replacemagic.com/help/RegularExpressionsMatchwildcards.html
			// targetArray := []string{}
			minionsToTarget = strings.ReplaceAll(minionsToTarget, "*", ".*")
			minionsToTarget = strings.ReplaceAll(minionsToTarget, " and ", "|")

			// Setup parser to find minions
			regex, err := regexp.CompilePOSIX(minionsToTarget)
			if err != nil {
				log.Printf("The target \"%s\" is not valid\n", minionsToTarget)
				os.Exit(1)
			}
			minionsFound := []string{}
			listOfMinonsOnSystem := []string{"iad1bastion", "us4repo01", "iad1repo01", "dkcph-p-salt-001", "dkcph-s-pay-jms-003"}

			for x := range listOfMinonsOnSystem {
				target := listOfMinonsOnSystem[x]
				// log.Printf("\tchecking minion \"%s\" against: \"%v\"\n", target, regex)
				found := regex.FindAllString(target, -1)
				if len(found) != 0 {
					minionsFound = append(minionsFound, found[0])
				}
			}
			if len(minionsFound) == 0 {
				log.Println("No minions found")
				os.Exit(1)
			}
			cli.Cli.Minions = minionsFound
			log.Println("rootCmd: targets: ", strings.Join(minionsFound, ", "))

			// if cmdRun {
			// 	a := []string{}
			// 	a = append(a, args[1:]...)
			// 	cmdRunner.CmdRunCmd.Run(cmd, a)

			// }
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	log.Println("rootCmd()")
	//
	//
	// rootCmd.Run(cmd, args)
	log.Println("rootCmd.Flags() called")
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "Galt server config file")
	rootCmd.PersistentFlags().StringVarP(&minionsToTarget, "minions", "m", "", "Minions to match - regular (wildcard) regex")
	//TODO: Add grains support to "-m """
	// rootCmd.MarkPersistentFlagRequired("minions")

	log.Println("cmdRunner.CmdRunCmd called")
	rootCmd.AddCommand(cmdRunner.CmdRunCmd)

}
