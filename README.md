## Shell Emulator (Go)

- Minimal interactive shell implemented in Go with built-ins, PATH lookups, and history support.

## Supported Features

- Interactive prompt powered by `liner` with persistent in-memory history.
- Built-in commands: `exit`, `echo`, `type`, `pwd`, `cd`, `history` (with optional limit).
- External command execution by scanning `$PATH` and delegating to binaries.
- Shell-style argument parsing supporting quotes, escaping, and whitespace preservation.
- Optional manual readline implementation with arrow-key history navigation (`readline.go`).

## Project Structure

```
.
├── app/
│   ├── main.go
│   ├── readline.go
│   ├── commands/
│   │   ├── cd.go
│   │   ├── history.go
│   │   ├── path.go
│   │   └── type.go
│   └── utils/
│       └── utils.go
├── go.mod
├── README.md
└── terminal.sh
```

## File-Level Responsibilities

- `app/main.go` — entrypoint; initializes the prompt, dispatches built-ins, and proxies external commands via `PathCommand`.
- `app/readline.go` — alternative stdin reader supporting arrow-key history; useful if `liner` is unavailable.
- `app/commands/cd.go` — implements directory changes with `~` expansion and error handling.
- `app/commands/history.go` — prints recorded commands, capped by an optional numeric limit.
- `app/commands/path.go` — spawns external binaries resolved from `$PATH` and streams their combined output.
- `app/commands/type.go` — reports whether a token is a builtin or resolves to an executable path.
- `app/utils/utils.go` — stores history, builtin metadata, argument parsing logic, and dynamic `$PATH` scanning.
- `terminal.sh` — helper script that builds (`go build`) and runs the shell binary from the repository root.
- `go.mod` — Go module definition pinning dependencies (e.g., `github.com/peterh/liner`).

## Getting Started

- Requirements: Go 1.20+ (module-enabled environment).
- Build & run via helper script:
  ```sh
  ./terminal.sh
  ```
- Or run directly with `go run`:
  ```sh
  go run app/*.go
  ```

## Usage Notes

- Prompt: `$ `; `Ctrl+C` aborts the current line, `Ctrl+D` exits.
- `history` supports `history 5` to print only the last five entries.
- External commands inherit arguments parsed by `utils.ParseArguments`, so quoting behaves similarly to POSIX shells.
- `TypeCommand` prefers builtins over PATH executables to match typical shell behavior.

## Implementation Highlights

- Uses `liner` for robust line editing; `utils.ScanPathCommands` caches discoverable executables on each prompt iteration.
- Argument parsing normalizes whitespace, supports nested quote contexts, and respects escape sequences like `\n` inside double quotes.
- External command execution uses `exec.Command` with combined stdout/stderr output to emulate shell output semantics.
