# goNDK

[![Run Tests](https://github.com/niallyoung/goNDK/actions/workflows/main.yaml/badge.svg)](https://github.com/niallyoung/goNDK/actions/workflows/main.yaml)

NOSTR Development Kit (Go) - well-engineered, comprehensively tested, easy to use, and intended (soon) for production solutions at scale.

NOTE: this is VERY EARLY R&D level exploratory work. Interfaces will be established ASAP, then a modular and consistent pluggable architecture throughout. Your patience is appreciated 🙇‍♂

## Features

- ✅ **Event System** - Create, sign, and validate NOSTR events (NIP-01)
  - Flexible validation (structure vs complete with signatures)
  - Support for empty content and complex tags
  - JSON serialization with proper escaping
- ✅ **Identity Management** - Complete NOSTR identity handling (NIP-19)
  - Generate new identities or load from nsec/hex
  - npub/nsec bech32 encoding/decoding
  - Schnorr signature creation and verification
- ✅ **Relay Client** - Full-featured relay communication
  - WebSocket connections with proper message handling
  - Event publishing and subscription with filters
  - Integration with multiple public relays
- ✅ **Encryption** - Privacy-preserving communication
  - NIP-44 encryption/decryption
  - NIP-59 gift wrap envelopes for metadata privacy
- ✅ **CLI Tools** - Ready-to-use command line utilities
  - Encrypted messaging with file transfer support
  - Identity generation and management
- ✅ **84.5% test coverage** - Thoroughly tested and reliable

## Installation

### Library

```shell
go get github.com/niallyoung/goNDK
```

### CLI Tool

```shell
cd examples/client-giftwrap
make build
make install  # Copies to ~/bin/giftwrap
```

## Quick Start

### Connect to a Relay

```go
import "github.com/niallyoung/goNDK/client"

rm := client.NewRelayManager("wss://relay.damus.io")
err := rm.Connect(ctx)
defer rm.Close()
```

### Subscribe to Events

```go
filters := []client.Filter{{Kinds: []int{1}, Limit: 10}}
sub, err := rm.Subscribe(ctx, filters)

sub.Receive(ctx, func(ctx context.Context, e *event.Event) {
    fmt.Printf("Received: %s\n", e.Content)
})
```

### Create and Sign Events

```go
import "github.com/niallyoung/goNDK/event"

e := event.NewEvent(1, "Hello NOSTR!", nil, nil, nil, nil, nil)
err := e.Sign(privateKeyHex)
```

### Send Encrypted Messages (Gift Wrap)

```go
import "github.com/niallyoung/goNDK/nips/nip59"

// Create inner event
innerEvent := event.NewEvent(1, "Secret message", nil, nil, nil, &senderPubKey, nil)
innerEvent.Sign(senderPrivKey)

// Wrap and encrypt
wrapped, _ := nip59.Wrap(innerEvent, senderPrivKey, recipientPubKey)

// Publish
rm.Publish(ctx, wrapped)
```

### Real-World Example

Successfully sent encrypted message to jb55 using the CLI:

```bash
cd examples/client-giftwrap
make build

./giftwrap send \
  --destination npub1xtscya34g58tk0z605fvr788k263gsu6cy9x0mhnm87echrgufzsevkk5s \
  --message "successful vibed gift-wrap for you bro, see more at https://github.com/niallyoung/goNDK" \
  --relay wss://relay.damus.io
```

This NIP-59 gift-wrapped event was published to relay.damus.io and can only be decrypted by jb55.

## Examples

### Basic Examples
- **[examples/client/](examples/client/)** - Relay connection and event fetching
- **[examples/identity-generator/](examples/identity-generator/)** - Identity creation and management

### Advanced Examples  
- **[examples/client-giftwrap/](examples/client-giftwrap/)** - Complete encrypted messaging CLI
  - Send/receive encrypted messages and files
  - Real-world NIP-59 gift wrap implementation
  - Integration with public relays

See individual example directories for detailed READMEs and usage instructions.

## Development

```shell
make test     # Run tests
make cover    # Coverage report (81.95%)
make lint     # Lint code
```

## Status

**v0.x.x** - Early R&D Work
- High test coverage (84.5%) with integration tests
- CLI tools for practical NOSTR usage
- Relay client with real-world testing on public relays
- Complete NIP-01, NIP-19, NIP-44, NIP-59 implementations

See [TODO.md](TODO.md) for roadmap and [CHANGELOG.md](CHANGELOG.md) for release history.

## Thanks

Built upon and inspired by:
- [khatru](https://github.com/fiatjaf/khatru)
- [go-nostr](https://github.com/nbd-wtf/go-nostr)
- [shota3506/go-nostr](https://github.com/shota3506/go-nostr)
- [nostr-domain](https://github.com/dextryz/nostr-domain)

## License

MIT License - Copyright (c) 2024-2025 Niall Young
