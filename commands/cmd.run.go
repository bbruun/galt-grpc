package commands

import (
	"encoding/json"
	"fmt"
)

type CmdRun struct {
	Commandline string
}

func (c *CmdRun) ToString() string {
	res, err := json.Marshal(c)
	if err != nil {
		fmt.Printf("Failed to marshal command: %s\n", c.Commandline)
		return ""
	}
	return string(res)
}
