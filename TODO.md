# TODO

## Relay Client

- [x] `Event{}` - 83.3% coverage
- [x] `Identity{}` - 100% coverage
- [x] `Client{}`, `RelayManager{}` - 74.9% coverage
  - [x] `Subscription{}`, `Filter{}`
  - [x] WebSocket connection management
  - [x] Event publishing
  - [x] Subscription with filters
  - [x] Message handling (EVENT, REQ, CLOSE, EOSE, OK, NOTICE)
  - [x] Integration tests with public relays
  - [ ] Retry/reconnection logic
  - [ ] Connection pooling

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
