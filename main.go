package main

import (
	"fmt"
	"os"

	"log"
	"net"

	"github.com/bbruun/galt/config"
	"github.com/bbruun/galt/initializers"
	"google.golang.org/grpc"
)

var VERSION string = "(development build)"
var AUTHOR string = "Bjarke Bruun - https://github.com/bbruun"
var PORT string = ":4505"

var IS_MASTER bool

func init() {
	fmt.Println("setting up config.Config")
	config.NewConfig("./config/galt.yaml")
	fmt.Printf("config.Config set up: %+v\n", config.Config)
	fmt.Printf("Listen address: %s\n", config.Config.GetListenAddress())
	// Parse command line parameters
	initializers.ParseFlags()

	// Load/override loaded configuration
	initializers.LoadEnvVariables()

	// Read default configuration file or specified configuration file
	config.NewConfig(initializers.ConfigFileLocation)
	config.Config.UpdateConfigFromEnvVars()

	// Connect to database and migrate/update it
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	fmt.Printf("Galt %s\n", VERSION)

	fmt.Printf("Config: %+v\n", config.Config)
	asdf := fmt.Sprintf("0.0.0.0:%s", config.Config.GRPCPort)
	fmt.Printf("asdf: %s\n", asdf)

	// listen, err := net.Listen("tcp", config.Config.GetListenAddress())
	listen, err := net.Listen("tcp", ":4505")
	if err != nil {
		log.Fatalf("failed to listen on port %s, %v\n", fmt.Sprintf("0.0.0.0:%s", config.Config.GRPCPort), err)
	}

	fmt.Printf("net.Listen(\"tcp\", \"%s\")\n", config.Config.GetListenAddress())

	fmt.Println(listen.Addr().String())
	fmt.Printf("listen: %+v\n", listen)

	fmt.Printf("Starting gRPC server on port %s", config.Config.GRPCPort)
	os.Exit(0)
	grpcServer := grpc.NewServer()
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to server gRPC over port %s, %v", config.Config.GRPCPort, err)
	} else {
		fmt.Printf("Server is up and running on port %s\n", config.Config.GRPCPort)
	}

}
