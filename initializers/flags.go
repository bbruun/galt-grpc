package initializers

import (
	"flag"
	"fmt"
)

var DEFAULTCONFIGFILE string = "./config/galt.yaml"

var ConfigFileLocation string

func ParseFlags() {
	fmt.Println("reading parameters")
	locationTmp := flag.String("config", DEFAULTCONFIGFILE, "Configuration file (default ./config/galt.yaml)")
	location := string(*locationTmp)
	fmt.Printf("configFileLocation = %s\n", location)
	flag.Parse()
	fmt.Println("- parameters parsed")
	// map flags to global configurations
	// GlobalFlags.SetConfigFileLocation(location)
	// config.Config = config.NewConfig(location)
	ConfigFileLocation = *locationTmp
	fmt.Println("- ParseFlags() done")
}
