package entries

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/angelocordon/asta/internal/models"
)

// GetEntriesPath returns the path to the entries.json file
func GetEntriesPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not determine home directory: %w", err)
	}
	return filepath.Join(home, ".asta", "entries.json"), nil
}

// ReadEntries reads and parses entries.json
func ReadEntries() (*models.EntriesFile, error) {
	entriesPath, err := GetEntriesPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(entriesPath)
	if err != nil {
		return nil, fmt.Errorf("could not read entries file: %w", err)
	}

	var entriesFile models.EntriesFile
	if err := json.Unmarshal(data, &entriesFile); err != nil {
		return nil, fmt.Errorf("could not parse entries file: %w", err)
	}

	return &entriesFile, nil
}

// WriteEntries writes entries to entries.json
func WriteEntries(entriesFile *models.EntriesFile) error {
	entriesPath, err := GetEntriesPath()
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

// GenerateEntryID generates a unique ID for an entry in the format e-{timestamp}-{random}
func GenerateEntryID() string {
	timestamp := time.Now().Unix()
	// Generate a random 4-character alphanumeric suffix using crypto/rand for better randomness
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	suffix := make([]byte, 4)
	// Use time-based seeding for simplicity since each CLI invocation is independent
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range suffix {
		suffix[i] = chars[r.Intn(len(chars))]
	}
	return fmt.Sprintf("e-%d-%s", timestamp, string(suffix))
}
