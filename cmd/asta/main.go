package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hashicorp/cli"
)

const version = "0.1.0"

func main() {
	c := cli.NewCLI("asta", version)
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"init": func() (cli.Command, error) {
			return &InitCommand{}, nil
		},
		"add": func() (cli.Command, error) {
			return &AddCommand{}, nil
		},
		"log": func() (cli.Command, error) {
			return &LogCommand{}, nil
		},
		"version": func() (cli.Command, error) {
			return &VersionCommand{
				Version: version,
			}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	os.Exit(exitStatus)
}

// InitCommand is a placeholder for the init command
type InitCommand struct{}

func (c *InitCommand) Help() string {
	return `Usage: asta init

  Initialize your asta repository. Creates ~/.asta/entries.json.`
}

func (c *InitCommand) Synopsis() string {
	return "Initialize your asta repository"
}

func (c *InitCommand) Run(args []string) int {
	// Get home directory
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Could not determine home directory: %v\n", err)
		return 1
	}

	// Define paths
	astaDir := filepath.Join(home, ".asta")
	entriesFile := filepath.Join(astaDir, "entries.json")

	// Create .asta directory if it doesn't exist
	if err := os.MkdirAll(astaDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error: Could not create directory %s: %v\n", astaDir, err)
		return 1
	}

	// Check if entries.json already exists
	_, err = os.Stat(entriesFile)
	if err == nil {
		// File already exists
		fmt.Println("✓ Initialized asta at ~/.asta")
		fmt.Println("  (Repository already initialized)")
		return 0
	}

	// Create entries.json with empty entries array
	initialData := map[string]interface{}{
		"entries": []interface{}{},
	}

	jsonData, err := json.MarshalIndent(initialData, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Could not create initial data: %v\n", err)
		return 1
	}

	if err := os.WriteFile(entriesFile, jsonData, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error: Could not write to %s: %v\n", entriesFile, err)
		return 1
	}

	fmt.Println("✓ Initialized asta at ~/.asta")
	return 0
}

// AddCommand is a placeholder for the add command
type AddCommand struct{}

func (c *AddCommand) Help() string {
	return `Usage: asta add <entry>

  Log an accomplishment. Records your achievement with a timestamp.

Example:
  asta add "Fixed critical auth bug, reducing login failures by 95%"`
}

func (c *AddCommand) Synopsis() string {
	return "Log an accomplishment"
}

func (c *AddCommand) Run(args []string) int {
	if len(args) < 1 {
		fmt.Println("Error: entry text required")
		fmt.Println(c.Help())
		return 1
	}
	fmt.Println("✓ Entry added")
	fmt.Println("(Placeholder - functionality not yet implemented)")
	return 0
}

// LogCommand is a placeholder for the log command
type LogCommand struct{}

func (c *LogCommand) Help() string {
	return `Usage: asta log

  View all your accomplishments, grouped by day.`
}

func (c *LogCommand) Synopsis() string {
	return "View all your accomplishments"
}

func (c *LogCommand) Run(args []string) int {
	fmt.Println("Total: 0 entries")
	fmt.Println("(Placeholder - functionality not yet implemented)")
	return 0
}

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
