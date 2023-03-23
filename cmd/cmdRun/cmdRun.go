/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmdRun

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
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

func SendCmdRunCommand(minions []string, args []string) error {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:4505", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewMinionServiceClient(conn)

	hostname, _ := os.Hostname()
	response, err := c.CmdRun(context.Background(), &proto.CmdRunFromClient{
		Name:              hostname,
		MessageFromClient: "",
		MessageFromServer: "",
		TargetMinions:     strings.Join(minions, ", "),
		Command:           strings.Join(args, ";"),
	})
	if err != nil {
		log.Fatalf("Error when creating message to Server: %s", err)
	}
	fmt.Println("command sent to Galt master")

	for {
		// resp, err := response.Recv()
		msg := proto.SendCommandResultToMinion{}
		err := response.RecvMsg(&msg)
		if err == io.EOF {
			return nil
		} else if err == nil {
			valStr := fmt.Sprintf("Response val: %s", msg.MessageFromServer)
			log.Println(valStr)
		}

		if err != nil {
			panic(err) // dont use panic in your real project
		}

	}
}

func init() {
	log.Println("CmdRunCmd()")
	inherited = CmdRunCmd.InheritedFlags()
	// cmdRunCmd.AddCommand(cmdRunCmd)

}
