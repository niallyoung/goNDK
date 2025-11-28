# TODO

## Relay Client

- [x] `Event{}` - 83.3% coverage
- [x] `Identity{}` - NIP-19 encoding/decoding, key generation
- [x] `Client{}`, `RelayManager{}` - 74.9% coverage
  - [x] `Subscription{}`, `Filter{}`
  - [x] WebSocket connection management
  - [x] Event publishing
  - [x] Subscription with filters
  - [x] Message handling (EVENT, REQ, CLOSE, EOSE, OK, NOTICE)
  - [x] Integration tests with public relays
  - [x] FetchEventByID helper
  - [ ] Retry/reconnection logic
  - [ ] Connection pooling

## Gift Wrap CLI

- [x] NIP-19 support (npub/nsec conversion)
- [x] Identity generation
- [x] File metadata encoding
- [x] Key storage
- [x] Send command (text and files)
- [x] Receive command (fetch, watch, stdin)
- [x] Working CLI tool

## Enhanced Client

- [ ] Retry/reconnection logic
- [ ] Connection pooling
- [ ] Rate limiting
- [ ] Improved filter implementation
- [ ] 85%+ test coverage

## Relay Implementation

- [ ] `Relay{}` server implementation
  - [ ] `LocalRelay{}`, `ProxyRelay{}`, etc.
- [ ] NIP implementations
- [ ] Advanced subscription patterns
