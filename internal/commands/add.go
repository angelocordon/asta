package commands

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/angelocordon/asta/internal/entries"
	"github.com/angelocordon/asta/internal/models"
)

// AddCommand implements the add command for logging accomplishments
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
	entriesPath, err := entries.GetEntriesPath()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return 1
	}

	if _, err := os.Stat(entriesPath); os.IsNotExist(err) {
		fmt.Fprintln(os.Stderr, "Error: asta not initialized. Run 'asta init' first.")
		return 1
	}

	// Read existing entries
	entriesFile, err := entries.ReadEntries()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return 1
	}

	// Create new entry
	newEntry := models.Entry{
		ID:        entries.GenerateEntryID(),
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Entry:     entryText,
	}

	// Append new entry
	entriesFile.Entries = append(entriesFile.Entries, newEntry)

	// Write back to file
	if err := entries.WriteEntries(entriesFile); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return 1
	}

	fmt.Println("✓ Entry added")
	return 0
}
