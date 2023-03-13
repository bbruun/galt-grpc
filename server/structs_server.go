package server

import (
	"fmt"
)

type GaltServerConfig struct {
	ConfigFileLocation string `yaml:"config_file_location"`
	ReadyToStart       bool   `yaml:"ready_to_start"`
	Server             struct {
		Listen string
		Port   int
	} `yaml:"server"`
}

func (g *GaltServerConfig) GetListenAddress() string {
	return fmt.Sprintf("%s:%d", g.Server.Listen, g.Server.Port)
}
