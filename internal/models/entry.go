package models

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
