package server

import (
	"log"
	"os"
	"sync"

	"gopkg.in/yaml.v2"
)

var Configuration *GaltServerConfig

var (
	ConfigFile          string   = "./etc/galt/server.yaml"
	ConfigFileLocations []string = []string{"/etc/galt/server.yaml", "./etc/galt/server.yaml"}
)

func init() {

	log.Println("Loading Galt Server Configuration")
}

func NewGaltServerConfig() *GaltServerConfig {
	c := GaltServerConfig{}
	for x := range ConfigFileLocations {
		if _, err := os.Stat(ConfigFileLocations[x]); err != nil {
			continue
		}
		log.Println("server.yaml loaded, parsing")
		lf, err := os.ReadFile(ConfigFileLocations[x])
		if err != nil {
			log.Printf("Faild to load configuration file \"%s\": %s\n", ConfigFileLocations[x], err)
			os.Exit(1)
		}
		yaml.Unmarshal(lf, &c)
	}
	return &c
}

func StartServer() {
	log.Printf("Starting Salt sever")
	var wg sync.WaitGroup
	wg.Add(1)
	go gRPCServer(wg)
	log.Printf("... server started and is ready to server")
	wg.Wait()

}
