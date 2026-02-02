# Bits of Go

Small, focused snippets for learning Go concepts.

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
- [ ] Error wrapping (`fmt.Errorf` with `%w`)
- [ ] Custom error types
- [ ] Sentinel errors
- [ ] Errors vs panics

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
