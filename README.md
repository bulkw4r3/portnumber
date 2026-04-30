# portnumber

## Commands
- Run: `go run .` (launches Bubble Tea TUI)
- Build: `go build`
- Test: `go test` (no tests currently)

## Behavior
- Interactive TUI using Bubble Tea framework
- Enter port number (1-48) to see computed IP: 172.16.76.(48+port)
- Spinner animates next to "Port: " label while waiting for input
- Press Enter to confirm, Esc/Ctrl+C to quit
- Example output: "Portnumber 5 -> 172.16.76.53"
- All confirmed results are persisted to a local SQLite database (`portnumber.db`)

## Notes
- The built binary is named `portnumber` (same as directory)
- Go version: 1.26.2 (from go.mod)
- Uses charmbracelet/bubbletea, bubbles (spinner, textinput), and lipgloss
- Uses `modernc.org/sqlite` (pure Go SQLite, no CGO required)
- Database file `portnumber.db` is auto-created on startup
- Table: `results (id, port, ip, created_at)`
