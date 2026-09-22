# English–German translation experiments (Go)

Historical learning experiments from 2023 exploring Go, API requests, and text translation. The original project used AI-generated example code.

## Status

These are two separate, incomplete experiments, not a maintained application. They both declare `main`, there is no `go.mod`, and the examples use different legacy integrations. Running `go run .` is not supported. Do not use browser session cookies as a substitute for a supported API.

| File | Purpose | Known limitation |
| --- | --- | --- |
| `ENG-GER GPT3 Translator.go` | Early completions-API example | Missing client import and module setup; legacy interface |
| `main.go` | Historical unofficial browser-session client | Unsupported integration; not a recommended setup |

Credential literals have been removed from the current files. If any published value was real, it must be revoked by its owner. Earlier Git history can still contain those values; this cleanup does not claim to erase history or revoke credentials.

Before reviving this project, consolidate the entry point, create a Go module, replace the legacy integrations with a supported API, and add tests using mocked responses. Never commit keys, session tokens, or clearance cookies.
