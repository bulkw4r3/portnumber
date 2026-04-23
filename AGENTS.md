# portnumber

## Commands
- Run: `go run .` (then enter port number when prompted)
- Build: `go build`
- Test: `go test` (no tests currently)

## Behavior
- Reads a port number (1-48) from stdin
- Outputs IP: 172.16.76.(48+port)
- Example: `echo 5 | go run .` outputs "Portnumber 5 -> 172.16.76.53"

## Notes
- The built binary is named `portnumber` (same as directory)
- Go version: 1.26.2 (from go.mod)