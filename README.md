# asta

A simple, local-first CLI tool for tracking your daily professional accomplishments.  Like git, but for your career.

## Why asta?

Performance review season comes around, and you're staring at a blank page trying to remember what you did for the past 6 months. Sound familiar?

**asta** helps you build a "brag document" by making it easy to log your daily accomplishments. When review time comes, you'll have concrete evidence of your impact.

## Philosophy

- **Local-first**: Your data lives in `~/.asta/entries.json` - no accounts, no cloud, just files
- **Simple**: Three commands.  One file. That's it.
- **Fast**: Log accomplishments in seconds, not minutes
- **Evidence-based**:  Capture impact, not just activity

## Installation

```bash
# Clone and build (for now)
git clone https://github.com/angelocordon/asta.git
cd asta
go build -o asta cmd/asta/main.go
mv asta /usr/local/bin/  # or anywhere in your PATH
```

## Quick Start

```bash
# Initialize your accomplishment log
asta init

# Log accomplishments as you complete them
asta add "Fixed critical auth bug, reducing login failures by 95% for 50k users"
asta add "Led architecture review with 4 teams, aligned on API migration strategy"
asta add "Mentored junior engineer on testing practices, helped them increase coverage from 40% to 85%"

# Review your accomplishments anytime
asta log
```

## Commands

### `asta init`

Initialize your asta repository.  Creates `~/.asta/entries.json`.

```bash
$ asta init
✓ Initialized asta at ~/.asta
```

### `asta add <entry>`

Log an accomplishment. 

```bash
$ asta add "Shipped payment reconciliation dashboard, saving finance team 4 hours daily"
✓ Entry added
```

### `asta log`

View all your accomplishments, grouped by day.

```bash
$ asta log

December 30, 2025
  • Fixed auth bug, reducing login failures by 95% for EU users (4:45 PM)
  • Reviewed 5 PRs for payments team (5:30 PM)

December 29, 2025
  • Led incident response for payment outage (2:15 PM)
  • Mentored junior engineer on testing practices (10:00 AM)

Total: 4 entries
```

## Data Storage

All data is stored locally in `~/.asta/entries.json`:

```json
{
  "entries": [
    {
      "id": "e-1735574400-a1b2",
      "timestamp": "2025-12-30T16:45:00Z",
      "entry":  "Fixed auth bug, reducing login failures by 95% for EU users"
    }
  ]
}
```

This means:
- ✅ **You own your data** - it's just a JSON file on your machine
- ✅ **Easy to backup** - copy the file anywhere
- ✅ **Portable** - works offline, no dependencies
- ✅ **Private** - nothing leaves your machine

## MVP Scope

This is the minimal viable product.  It does one thing well:  **log and review accomplishments**. 

**Current features:**
- ✅ Initialize repository
- ✅ Add entries
- ✅ View all entries

**Future features** (Phase 2+):
- [ ] Filter by date (`--since`, `--today`, `--week`, `--month`)
- [ ] Export to markdown/PDF
- [ ] Edit/delete entries
- [ ] Search entries
- [ ] Tags and categories
- [ ] Interactive entry mode with prompts
- [ ] Entry templates
- [ ] Quality coaching
- [ ] Stats and insights
- [ ] Cloud sync (optional)
- [ ] Web interface
