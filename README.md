# goNDK

[![Run Tests](https://github.com/niallyoung/goNDK/actions/workflows/main.yaml/badge.svg)](https://github.com/niallyoung/goNDK/actions/workflows/main.yaml)

A NOSTR Development Kit in Go - well-engineered, production-ready, and easy to use.

## Features

- ✅ **Event** - Create, sign, and validate NOSTR events (NIP-01)
- ✅ **Identity** - Manage NOSTR identities (npub/nsec, NIP-19)
- ✅ **Client** - Connect to relays, publish events, subscribe to filters
- ✅ **NIP-44** - Encryption/decryption
- ✅ **NIP-59** - Gift wrap envelopes
- ✅ **CLI Tool** - Send/receive encrypted messages and files
- ✅ **81.95% test coverage** - Thoroughly tested and reliable

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

See `examples/client/` and `examples/client-giftwrap/` for complete working examples.

## Development

```shell
make test     # Run tests
make cover    # Coverage report (81.95%)
make lint     # Lint code
```

## Status

**v0.0.9** - Relay client with subscriptions, publishing, and comprehensive test coverage.

See [TODO.md](TODO.md) for roadmap and [CHANGELOG.md](CHANGELOG.md) for release history.

## Thanks

Built upon and inspired by:
- [khatru](https://github.com/fiatjaf/khatru)
- [go-nostr](https://github.com/nbd-wtf/go-nostr)
- [shota3506/go-nostr](https://github.com/shota3506/go-nostr)
- [nostr-domain](https://github.com/dextryz/nostr-domain)

## License

MIT License - Copyright (c) 2024 Niall Young
