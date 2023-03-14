/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmdRun

import (
	"context"
	"log"
	"strings"

	"github.com/bbruun/galt/cli"
	"github.com/bbruun/galt/proto"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"
)

var inherited *pflag.FlagSet

// cmd.runCmd represents the cmd.run command
var CmdRunCmd = &cobra.Command{
	Use:   "cmd.run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("cmdRun: args: %v\n", args)
		log.Printf("cmdRun: minions: %s\n", cli.Cli.Minions[0:])
		SendCmdRunCommand(cli.Cli.Minions, args)
	},
}

func SendCmdRunCommand(minions []string, args []string) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:4505", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewMinionServiceClient(conn)

	response, err := c.CmdRun(context.Background(), &proto.CmdRunSend{Minions: strings.Join(minions, "\n"), Command: strings.Join(args, ";")})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	for x := range response.GetMinionCmdResult() {
		log.Printf("Response from server: %s", x)
	}
}

func init() {
	log.Println("CmdRunCmd()")
	inherited = CmdRunCmd.InheritedFlags()
	// cmdRunCmd.AddCommand(cmdRunCmd)

}
