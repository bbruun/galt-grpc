package cli

var Cli *CliValues

type CliValues struct {
	Minions []string
}

func NewCli() *CliValues {
	return &CliValues{}
}
