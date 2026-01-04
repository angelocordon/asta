package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/cli"

	"github.com/angelocordon/asta/internal/commands"
	"github.com/angelocordon/asta/internal/version"
)

func main() {
	c := cli.NewCLI("asta", version.Version)
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"init": func() (cli.Command, error) {
			return &commands.InitCommand{}, nil
		},
		"add": func() (cli.Command, error) {
			return &commands.AddCommand{}, nil
		},
		"log": func() (cli.Command, error) {
			return &commands.LogCommand{}, nil
		},
		"version": func() (cli.Command, error) {
			return &commands.VersionCommand{
				Version: version.Version,
			}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	os.Exit(exitStatus)
}
