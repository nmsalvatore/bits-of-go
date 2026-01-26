# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Small Go snippets for learning HTTP server fundamentals. The goal is to understand what each piece does, not just make it work. Each subdirectory in `web/` is a standalone example that builds on previous concepts.

## Commands

Each example is a standalone Go program. Run from the specific directory:

```bash
# Run an example (from its directory)
go run main.go

# Run from repo root
go run ./web/hello-json/main.go
```

All examples listen on port 8080. Test with curl:
```bash
curl localhost:8080
curl "localhost:8080?name=Gopher"
```

## Code Structure

- `web/hello/` - Basic handler, ServeMux, ListenAndServe
- `web/hello-param/` - Query parameter parsing with `r.URL.Query().Get()`
- `web/hello-json/` - JSON responses with `encoding/json`, proper Content-Type header, error handling

## Learning Notes

### HTTP Handler Patterns
- Handler function signature: `func(http.ResponseWriter, *http.Request)`
- Use explicit `mux := http.NewServeMux()` over `http.DefaultServeMux`
- Variable name `mux` is conventional for ServeMux

### Error Handling in Handlers
- Never `log.Fatal` in handlers (kills server)
- Use `http.Error(w, message, statusCode)` then `return`

### JSON Responses
- Set `Content-Type: application/json` before writing
- `json.Marshal` errors should return 500, not panic

## Next Up
- Testing handlers with `net/http/httptest`
