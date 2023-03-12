package server

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

var Configuration *GaltServerConfig

var (
	ConfigFile          string   = "./etc/galt/server.yaml"
	ConfigFileLocations []string = []string{"/etc/galt/server.yaml", "./etc/galt/server.yaml"}
)

func NewGaltServerConfig() *GaltServerConfig {
	c := GaltServerConfig{}
	for x := range ConfigFileLocations {
		if _, err := os.Stat(ConfigFileLocations[x]); err != nil {
			continue
		}

		lf, err := os.ReadFile(ConfigFileLocations[x])
		if err != nil {
			fmt.Printf("Faild to load configuration file \"%s\": %s\n", ConfigFileLocations[x], err)
			os.Exit(1)
		}

		yaml.Unmarshal(lf, &c)
		fmt.Printf("Server configuration loaded: %+v\n", c)
	}

	return &c
}
func StartServer() {
	fmt.Printf("Starting Salt sever")
}
