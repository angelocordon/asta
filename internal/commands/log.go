package commands

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/angelocordon/asta/internal/entries"
)

// LogCommand implements the log command for viewing entries
type LogCommand struct{}

func (c *LogCommand) Help() string {
	return `Usage: asta log

  View all your accomplishments, grouped by day.`
}

func (c *LogCommand) Synopsis() string {
	return "View all your accomplishments"
}

func (c *LogCommand) Run(args []string) int {
	// Check if asta is initialized
	entriesPath, err := entries.GetEntriesPath()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return 1
	}

	// Check if file exists
	if _, err := os.Stat(entriesPath); os.IsNotExist(err) {
		// File doesn't exist - show empty state
		fmt.Println("Total: 0 entries")
		return 0
	}

	// Read entries
	entriesFile, err := entries.ReadEntries()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return 1
	}

	// Handle empty entries
	if len(entriesFile.Entries) == 0 {
		fmt.Println("Total: 0 entries")
		return 0
	}

	// Display each entry in git log style
	for i, entry := range entriesFile.Entries {
		if i > 0 {
			fmt.Println() // Add blank line between entries
		}

		// Parse timestamp
		t, err := time.Parse(time.RFC3339, entry.Timestamp)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: Could not parse timestamp for entry %s: %v\n", entry.ID, err)
			continue
		}

		// Display entry header
		fmt.Printf("entry %s\n", entry.ID)
		fmt.Printf("Date: %s\n", t.Format("January 2, 2006 15:04:05"))
		fmt.Println()

		// Display entry text with 4-space indentation
		lines := strings.Split(entry.Entry, "\n")
		for _, line := range lines {
			fmt.Printf("    %s\n", line)
		}
	}

	// Display total count
	fmt.Printf("\nTotal: %d entries\n", len(entriesFile.Entries))
	return 0
}
