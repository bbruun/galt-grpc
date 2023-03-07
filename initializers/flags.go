package initializers

import (
	"flag"
	"fmt"
)

var DEFAULTCONFIGFILE string = "./config/galt.yaml"

var ConfigFileLocation string
var ExitBeforeStart bool
var DEBUG bool

func ParseFlags() {
	fmt.Println("reading parameters")
	location := flag.String("config", DEFAULTCONFIGFILE, "Configuration file (default ./config/galt.yaml)")
	debug := flag.Bool("debug", false, "Enable debug")
	exitBeforeStart := flag.Bool("devexit", false, "Exit before starting actual gRPC server")
	flag.Parse()

	ConfigFileLocation = *location
	DEBUG = *debug
	ExitBeforeStart = *exitBeforeStart

}
