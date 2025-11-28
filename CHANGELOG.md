# CHANGELOG

## [v0.1.0](https://github.com/niallyoung/goNDK/tree/v0.1.0) (2025-01-27)

[Full Changelog](https://github.com/niallyoung/goNDK/compare/v0.0.9...v0.1.0)

**Major Features:**
- NIP-19 identity encoding/decoding (npub/nsec ↔ hex)
- Identity generation and key management
- NIP-44 encryption/decryption
- NIP-59 gift wrap envelopes
- Complete CLI tool for encrypted messaging
- File transfer support with MIME detection
- FetchEventByID helper for single event retrieval

**Gift Wrap CLI:**
- Send text messages and files
- Receive and decrypt messages
- Watch mode for real-time messaging
- Automatic key generation and storage
- Progress indicators and error handling
- Comprehensive examples and documentation

**New Packages:**
- `identity/nip19.go` - Bech32 encoding/decoding
- `identity/generate.go` - Key generation
- `client/helpers.go` - Event fetching utilities
- `examples/client-giftwrap/` - Complete CLI application

**Test Coverage:**
- All new features fully tested
- Integration tests for CLI
- MIME type detection tests
- Key storage tests

## [v0.0.9](https://github.com/niallyoung/goNDK/tree/v0.0.9) (2025-01-27)

[Full Changelog](https://github.com/niallyoung/goNDK/compare/v0.0.8...v0.0.9)

**Major Features:**
- Full relay client implementation (Client, RelayManager, Subscription)
- WebSocket connection management
- Event publishing and subscription with filters
- All NIP-01 message types (EVENT, REQ, CLOSE, EOSE, OK, NOTICE)
- Integration tests with 4 public relays
- Working examples in `examples/client/`

**Test Coverage:**
- Overall: 81.95% (was ~80%)
- client: 74.9%
- event: 83.3%
- identity: 100%

**Changes:**
- Add RelayManager for relay connections
- Add Subscription with filter support
- Add message types (EventMessage, ReqMessage, CloseMessage, etc.)
- Add integration tests against public relays
- Add examples (fetch_events.go, test_multiple_relays.go)
- Fix Makefile to exclude examples from tests
- Comprehensive unit tests for error paths

## [v0.0.8](https://github.com/niallyoung/goNDK/tree/v0.0.8) (2024-04-21)

[Full Changelog](https://github.com/niallyoung/goNDK/compare/v0.0.7...v0.0.8)

- initial Identity{}
- 96.55% test coverage

## [v0.0.7](https://github.com/niallyoung/goNDK/tree/v0.0.7) (2024-04-20)

[Full Changelog](https://github.com/niallyoung/goNDK/compare/v0.0.6...v0.0.7)

- 96.47% test coverage
- fixed serialization tests
- removed unnecessary deps
- simplify and polish

## [v0.0.6](https://github.com/niallyoung/goNDK/tree/v0.0.6) (2024-04-20)

[Full Changelog](https://github.com/niallyoung/goNDK/compare/v0.0.5...v0.0.6)

- event.Sign() bugfix

## [v0.0.5](https://github.com/niallyoung/goNDK/tree/v0.0.5) (2024-04-19)

[Full Changelog](https://github.com/niallyoung/goNDK/compare/v0.0.4...v0.0.5)

- remove eventrequest for now = 85.71% coverage

## [v0.0.4](https://github.com/niallyoung/goNDK/tree/v0.0.4) (2024-04-19)

[Full Changelog](https://github.com/niallyoung/goNDK/compare/v0.0.3...v0.0.4)

- additional test coverage = 84.88% total

## [v0.0.3](https://github.com/niallyoung/goNDK/tree/v0.0.3) (2024-04-18)

[Full Changelog](https://github.com/niallyoung/goNDK/compare/v0.0.2...v0.0.3)

- update README usage
- fix Event.ValidateSignature() when unsigned

## [v0.0.2](https://github.com/niallyoung/goNDK/tree/v0.0.2) (2024-04-14)

[Full Changelog](https://github.com/niallyoung/goNDK/compare/v0.0.1...v0.0.2)

- update README with usage and example code
- CHANGELOG.md

## [v0.0.1](https://github.com/niallyoung/goNDK/tree/v0.0.1) (2024-04-14)

[Full Changelog](https://github.com/niallyoung/goNDK/compare/aa6aa22...v0.0.1)

**Initial Release**

goNDK is a NOSTR Development Kit, for Golang.

- event.Event{}
- eventrequest.EventRequest{}
