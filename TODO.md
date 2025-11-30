# TODO

# Core Types

## Relay Client

- [x] `Event{}` - 82.1% coverage
  - [x] Relaxed validation (empty content, special chars in tags)
  - [x] Separate signature validation (Validate vs ValidateComplete)
  - [x] Robust nil tag handling
- [x] `Identity{}` - 83.8% coverage (NIP-19 encoding/decoding, key generation)
- [x] `Client{}`, `RelayManager{}` - 80.2% coverage
  - [x] `Subscription{}`, `Filter{}`
  - [x] WebSocket connection management
  - [x] Event publishing
  - [x] Subscription with filters
  - [x] Message handling (EVENT, REQ, CLOSE, EOSE, OK, NOTICE)
  - [x] Integration tests with public relays
  - [x] FetchEventByID helper
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

## Identity and Key Management
Assuming a future NFC/QR hardware signer:
- [ ] Isolate signing behind an Interface and centralised service
- [ ] Key storage and retrieval mechanisms
- [ ] Integration with secure enclaves or OS key stores
- [ ] Key rotation and revocation R&D
- [ ] Support for external signers (NFC, QR code, hardware wallets)

# Examples

## Gift Wrap CLI

- [x] NIP-19 support (npub/nsec conversion)
- [x] Identity generation
- [x] File metadata encoding
- [x] Key storage
- [x] Send command (text and files)
- [x] Receive command (fetch, watch, stdin)
- [x] Working CLI tool

## Identity CLI
- [x] NIP-19 support (npub/nsec conversion)
- [x] Identity generation
- [x] Key storage
- [x] Display identity info
- [x] Working CLI tool

# Next Targets

## Documentation
- [ ] Comprehensive README
- [ ] Code comments and docstrings
- [ ] Usage examples and tutorials
- [ ] NIP references and explanations
- [ ] NOSTR protocol specification, artefacts, diagrams, visualisations

## Interfaces
- [ ] Define clear interfaces for core components and layers of responsibility
- [ ] Modular and pluggable implementation #1 refactor behind well-exercised interfaces
- [ ] Modular and pluggable implementation #2 to exercise and prove out architecture
- [ ] consider, rationalise, exercise, compare, iterate 🤔

## NOSTR Features
- [ ] multi-homed transports
- [ ] backend storage and query options, local and remote
- [ ] low-latency high-throughput event streaming
- [ ] relay multiplexing and distribution
- [ ] client and relay version/config negotiation
- [ ] clients in relays (relay has a key-pair)
- [ ] relays in clients (embedded local relay)
- [ ] privacy and anonymity improvements at a relay-network architecture level (+ZKP?)
- [ ] relay specialization
- [ ] client experiments
- [ ] goNDK suitable for long-term protocol R&D, experimentation and extension

## Reproducible Builds
- [ ] deterministic builds
- [ ] reproducible binaries
- [ ] verifiable dependencies
- [ ] build provenance tracking
- [ ] release checksums published on NOSTR
- [ ] review latest git/NOSTR/Blossom options
- [ ] bootstrap shenanigans

## Repo Structure
- [ ] consider splitting into 2..N repos
  - [ ] examples/* -> standalone polished examples
  - [ ] docs/* protocol/* -> standalone repos
  - [ ] nips/* -> reference or just 1 impl? -> external lib?

## Performance Optimization
- [ ] Benchmarking
- [ ] Profiling and bottleneck identification
- [ ] Load testing
- [ ] Scalability improvements
- [ ] local/dev, cicd, mobile, cloud, onprem, mcu?

## Security Audits
- [ ] basic govuln etc. in Makefile, CICD
- [ ] Review and Implementation of security best practices
- [ ] focused Code Reviews for security vulnerabilities
- [ ] Third-party security audit
- [ ] Long-term strategy around external dependencies

## Fuzzing
- [ ] Fuzz testing for core types
- [ ] Fuzz testing *
