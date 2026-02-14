# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Small, focused Go snippets ("bits") for learning Go concepts. Each bit is a self-contained directory with its own `main.go` (and sometimes templates, tests, or a `go.mod`).

## Structure

- `web/` - HTTP/web bits, each in its own directory (e.g., `web/hello/`, `web/cookies/`, `web/auth-session-csrf/`)
- Bits with tests have their own `go.mod` and use `package main` (e.g., `web/hello-json-post/`)
- Bits without tests are standalone `main.go` files with no `go.mod`
- Template-based bits include `.html` files alongside the Go code

## Commands

Run a server bit (must `cd` into the bit directory since templates use relative paths):
```
cd web/hello && go run main.go
```

Run tests for a specific bit:
```
cd web/hello-json-post && go test -v
```

Run all tests across the repo:
```
find . -name '*_test.go' -execdir go test -v \;
```

## Conventions

- Each bit is self-contained: no shared packages or imports between bits
- Server bits listen on port 8080 by default
- Tests use `net/http/httptest` with `httptest.NewRequest` and `httptest.NewRecorder`
- Bits that need tests split handler logic into a separate function (not inside `main()`) so it can be tested
- Templates are parsed with `template.Must(template.ParseFiles(...))` at package level
- Standard library only — no third-party dependencies

## Bit Map

### HTTP/Web
- [x] Basic handler
- [x] Query params
- [x] JSON response
- [x] JSON POST
- [x] Cookies
- [x] Templates (data, loops, layouts)
- [x] Forms
- [x] Session auth
- [ ] CSRF
- [ ] Path parameters (`/users/{id}`)
- [ ] Middleware (logging, timing, auth wrapper)
- [ ] File uploads
- [ ] Static file serving
- [ ] HTTP client (making requests)
- [ ] Websockets

### Cryptography
- [ ] Hashing (SHA256, checksums)
- [ ] Password hashing (bcrypt)
- [ ] HMAC signing
- [ ] Symmetric encryption (AES)
- [ ] Asymmetric encryption (RSA basics)
- [ ] TLS/certificates (conceptual)

### Database/Persistence
- [ ] File read/write
- [ ] JSON file storage
- [ ] SQLite basics
- [ ] SQL queries with `database/sql`
- [ ] Migrations
- [ ] Connection pooling

### Concurrency
- [ ] Goroutines
- [ ] Channels
- [ ] WaitGroups
- [ ] Mutexes
- [ ] Worker pools
- [ ] Context for cancellation

### Testing
- [x] Table-driven tests
- [x] httptest
- [ ] Test fixtures
- [ ] Mocking
- [ ] Benchmarks
- [ ] Fuzzing

### Data & Encoding
- [x] JSON encode/decode
- [ ] XML
- [ ] CSV
- [ ] Base64 vs hex
- [ ] Gob (Go binary format)

### CLI & Tools
- [ ] Flag parsing
- [ ] Environment variables
- [ ] Stdin/stdout
- [ ] Building a CLI tool

### Error Handling
- [x] Error wrapping (`fmt.Errorf` with `%w`)
- [x] Custom error types
- [x] Sentinel errors
- [x] Errors vs panics

### Standard Library Gems
- [ ] `time` (parsing, formatting, timezones)
- [ ] `strings` / `bytes` manipulation
- [ ] `regexp`
- [ ] `io` interfaces
- [ ] `context`

## Notes

- The sessions map in auth-session isn't thread-safe - will need a mutex when covering concurrency
- `html/template` escapes automatically, `text/template` does not
- `crypto/rand.Read` never returns an error (interface compliance only)
