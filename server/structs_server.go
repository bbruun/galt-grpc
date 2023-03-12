package server

type GaltServerConfig struct {
	ConfigFileLocation string `yaml:"config_file_location"`
	ReadyToStart       bool   `yaml:"ready_to_start"`
	Server             struct {
		Port int
	} `yaml:"server"`
	Logger struct {
		LogLevel  string `yaml:"logLevel"`
		Directory string `yaml:"directory"`
		Filename  string `yaml:"filename"`
		Format    string `yaml:"format"`
	} `yaml:"logging"`
}
