package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/angelocordon/asta/internal/models"
)

// InitCommand implements the init command for initializing the asta repository
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
	} else if !os.IsNotExist(err) {
		// Some other error occurred (permissions, etc.)
		fmt.Fprintf(os.Stderr, "Error: Could not check file %s: %v\n", entriesFile, err)
		return 1
	}

	// Create entries.json with empty entries array
	initialData := models.EntriesFile{
		Entries: []models.Entry{},
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
