package commands

import (
	"fmt"
)

// VersionCommand displays the version
type VersionCommand struct {
	Version string
}

func (c *VersionCommand) Help() string {
	return `Usage: asta version

  Print the version number of asta.`
}

func (c *VersionCommand) Synopsis() string {
	return "Print the version number"
}

func (c *VersionCommand) Run(args []string) int {
	fmt.Printf("asta version %s\n", c.Version)
	return 0
}
