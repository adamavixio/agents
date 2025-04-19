# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

- Build: `go build -o ai ./cmd/main.go`
- Run: `go run cmd/main.go`
- Test all: `go test ./...`
- Test single package: `go test ./pkg/package_name`
- Test specific test: `go test -run TestName ./pkg/package_name`
- Lint: `golangci-lint run`
- Format: `gofmt -w .`

## Code Guidelines

- Formatting: Follow standard Go style with `gofmt`
- Imports: Group standard library, third-party, and internal imports
- Error handling: Always check errors, never use `_` for discarding errors
- Naming: Use CamelCase for exported names, camelCase for unexported
- Comments: Document all exported functions, types, and constants
- Types: Use strong typing, avoid interface{} when possible
- Error handling: Return errors rather than panicking
- Testing: Write table-driven tests with descriptive names
