package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/hashicorp/cli"
)

const version = "0.1.0"

// Entry represents a single accomplishment entry
type Entry struct {
	ID        string `json:"id"`
	Timestamp string `json:"timestamp"`
	Entry     string `json:"entry"`
}

// EntriesFile represents the structure of entries.json
type EntriesFile struct {
	Entries []Entry `json:"entries"`
}

// getEntriesPath returns the path to the entries.json file
func getEntriesPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not determine home directory: %w", err)
	}
	return filepath.Join(home, ".asta", "entries.json"), nil
}

// readEntries reads and parses entries.json
func readEntries() (*EntriesFile, error) {
	entriesPath, err := getEntriesPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(entriesPath)
	if err != nil {
		return nil, fmt.Errorf("could not read entries file: %w", err)
	}

	var entriesFile EntriesFile
	if err := json.Unmarshal(data, &entriesFile); err != nil {
		return nil, fmt.Errorf("could not parse entries file: %w", err)
	}

	return &entriesFile, nil
}

// writeEntries writes entries to entries.json
func writeEntries(entriesFile *EntriesFile) error {
	entriesPath, err := getEntriesPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(entriesFile, "", "  ")
	if err != nil {
		return fmt.Errorf("could not marshal entries: %w", err)
	}

	if err := os.WriteFile(entriesPath, data, 0644); err != nil {
		return fmt.Errorf("could not write entries file: %w", err)
	}

	return nil
}

// generateEntryID generates a unique ID for an entry in the format e-{timestamp}-{random}
func generateEntryID() string {
	timestamp := time.Now().Unix()
	// Generate a random 4-character alphanumeric suffix
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	suffix := make([]byte, 4)
	for i := range suffix {
		suffix[i] = chars[rand.Intn(len(chars))]
	}
	return fmt.Sprintf("e-%d-%s", timestamp, string(suffix))
}

func main() {
	// Initialize random number generator for entry IDs
	rand.Seed(time.Now().UnixNano())

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
	initialData := EntriesFile{
		Entries: []Entry{},
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
		fmt.Fprintln(os.Stderr, "Error: entry text required")
		fmt.Fprintln(os.Stderr, c.Help())
		return 1
	}

	// Join all args as the entry text
	entryText := strings.Join(args, " ")

	// Check if asta is initialized
	entriesPath, err := getEntriesPath()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return 1
	}

	if _, err := os.Stat(entriesPath); os.IsNotExist(err) {
		fmt.Fprintln(os.Stderr, "Error: asta not initialized. Run 'asta init' first.")
		return 1
	}

	// Read existing entries
	entriesFile, err := readEntries()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return 1
	}

	// Create new entry
	newEntry := Entry{
		ID:        generateEntryID(),
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Entry:     entryText,
	}

	// Append new entry
	entriesFile.Entries = append(entriesFile.Entries, newEntry)

	// Write back to file
	if err := writeEntries(entriesFile); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return 1
	}

	fmt.Println("✓ Entry added")
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
