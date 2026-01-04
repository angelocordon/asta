# Copilot Instructions for asta

## About This Project

asta is a simple, local-first CLI tool for tracking daily professional accomplishments. Built with Go, it follows a minimalist philosophy: three commands, one file, that's it.

## Git Commit Guidelines

When creating commits for this repository, follow these rules:

- **Do not commit an empty commit**. Do not make a commit without any changes.
- **Short (72 chars or less) summary** in imperative mood
- **Blank line** separating summary from body
- **Body paragraphs** wrapped at 72 characters
- **Write in imperative mood**: "Fix bug" not "Fixed bug" or "Fixes bug"
- **Avoid listing file changes** - git shows that; focus on what and why
- **Body sentences start with imperative verbs** without subjects (e.g., "Improves" not "This improves")
- Separate paragraphs with blank lines when needed

A properly formed git commit subject line should complete:

> If applied, this commit will _[your subject line here]_

### Examples

```
Add filter by date functionality

Implement --since, --today, --week, and --month flags for the log
command to allow filtering entries by date range.

Allows users to quickly view accomplishments from specific time
periods without manually searching through all entries.
```

```
Refactor into internal package structure

Extract commands, models, and utilities from monolithic main.go
into a clean internal package structure following Go best
practices. Improves code organization, maintainability, and
separation of concerns.

The new structure makes it easier to test individual components,
add new commands, and maintain clear boundaries between different
parts of the application. Package naming uses domain-specific
terms (entries, models, commands) for better code readability.
```

## Code Style

- Follow standard Go conventions and use `gofmt` for formatting
- Keep functions focused and single-purpose
- Prefer simplicity over cleverness
- Write clear, descriptive variable names
- Add comments for complex logic, but let the code speak for itself when possible

## Testing

- Write tests for new functionality
- Ensure existing tests pass before committing
- Use table-driven tests where appropriate (Go best practice)

## Project Philosophy

- **Local-first**: Data stays on the user's machine
- **Simple**: Minimal features, maximum impact
- **Fast**: Operations should complete in milliseconds
- **Evidence-based**: Focus on capturing impact, not just activity
