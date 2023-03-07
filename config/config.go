package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type GlobalConfiguration struct {
	GRPCAddress string `yaml:"grpcaddress"`
	GRPCPort    string `yaml:"grpcport"`
	Keys        struct {
		PublicKey  string `yaml:"publickey"`
		PrivateKey string `yaml:"privatekey"`
	} `yaml:"keys"`
}

func (c *GlobalConfiguration) GetListenAddress() string {
	fmt.Printf("Returning listening address: %s:%s\n", c.GRPCAddress, c.GRPCPort)
	return fmt.Sprintf("%s:%s", c.GRPCAddress, c.GRPCPort)
}

var Config GlobalConfiguration

func NewConfig(location string) *GlobalConfiguration {
	if _, err := os.Stat(location); err != nil {
		log.Fatalf("configuration file \"%s\" file found: %v\n", location, err)
	}
	yf, err := ioutil.ReadFile(location)
	if err != nil {
		log.Fatalf("%s is not understood: %v\n", location, err)
	}
	c := GlobalConfiguration{}
	yaml.Unmarshal(yf, &c)

	Config = c
	// fmt.Printf("Config created: %+v\n", Config)
	return &c
}

func (c *GlobalConfiguration) UpdateConfigFromEnvVars() {
	asdf := os.Getenv("PORT")
	if asdf != "" {
		Config.GRPCPort = asdf
	}

	asdf = os.Getenv("PUBLICKEY")
	if asdf != "" {
		Config.Keys.PublicKey = asdf
	}
	asdf = os.Getenv("PRIVATEKEY")
	if asdf != "" {
		Config.Keys.PrivateKey = asdf
	}
}
