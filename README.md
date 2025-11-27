# goNDK

[![Run Tests](https://github.com/niallyoung/goNDK/actions/workflows/main.yaml/badge.svg)](https://github.com/niallyoung/goNDK/actions/workflows/main.yaml)

goNDK is a NOSTR Development Kit in Golang

## Goals

- well-engineered framework of NOSTR types, funcs and interfaces
- 95%+ unit & integration test coverage
- useful and valuable to a wide audience, from rapid prototyping to production-ready solutions at scale
- facilitate NIP experimentation, help define and maintain interoperability, build for the long-term
- ensure the NOSTR development experience in Go is high-quality, productive, flexible, easy and fun

## Status

* `initial_client` branch: Client + RelayManager complete
* total unit test coverage = 68%+ (see `./.meta/cover.sh`)
* ready for merge to main
* v0.1.0 release target: functional relay client

- [x] `Event{}` - 83.3% coverage
- [x] `Identity{}` - 100% coverage
- [x] `Client{}`, `RelayManager{}` - 68.2% coverage
  - [x] `Subscription{}`, `Filter{}`
  - [x] WebSocket connection management
  - [x] Event publishing
  - [x] Subscription with filters
  - [x] Message handling (EVENT, REQ, CLOSE, EOSE, OK, NOTICE)
  - [ ] Retry/reconnection logic
  - [ ] Connection pooling
- [ ] `Relay{}` - Future
  - [ ] `LocalRelay{}`, `ProxyRelay{}`, etc.

~v0.1.0 ETA: Q1 2025: functional `Client{}` + `RelayManager{}`
  * ✅ Event creation and signing
  * ✅ WebSocket relay connections
  * ✅ Event publishing
  * ✅ Subscriptions with filters
  * ✅ 68%+ test coverage
  * ⏳ Merge initial_client → main

~v0.2.0 ETA: Q2 2025: Enhanced client features
  * Retry/reconnection logic
  * Connection pooling
  * Rate limiting
  * 80%+ test coverage

## Development

```shell
make lint     # golangci-lint
make test     # unit tests
make cover    # 95%+ coverage
make generate # code generation (event/event_easyjson.go)

make docker.build
make docker.lint
make docker.test
make docker.cover
make docker.shell
```

## Usage

### Installation

```shell
go get github.com/niallyoung/goNDK
```

### Event{}

```go
import (
    "github.com/niallyoung/goNDK/event"
)

// unmarshal an incoming JSON serialised Event
var event event.Event
err := json.Unmarshal([]byte(`{"kind": 1, "content": "...", ... }`), &event)

// create and sign a new event
e := event.NewEvent(1, "hello world!", event.Tags(nil), nil, nil, nil, nil)
err := e.Sign(privateKey.Key.String()) // Sign an Event

// serialization
text := e.String()
bytes := e.Serialize()

// validation
err := e.Validate()
ok, err := e.ValidateSignature()
```

### Client{}

```go
import (
	"github.com/niallyoung/goNDK/client"
)

c := client.NewClient()
err := c.Validate()
```

### RelayManager{}

```go
import (
	"github.com/niallyoung/goNDK/client"
)

c := client.NewRelayManager(url)
err := c.Connect()
sub, err := c.Subscribe(ctx, filters)

res, err := c.Publish(ctx, event)

err := c.WriteMessage(ctx, message)
err := c.ReadMessage(ctx)
```

### Identity{}

```go
import (
	"github.com/niallyoung/goNDK/identity"
)

i := identity.NewIdentity(pubkey, npub)
err := i.Validate()
```



## Thanks

Built upon, around and inspired by:

* [khatru](https://github.com/fiatjaf/khatru)
* [go-nostr](https://github.com/nbd-wtf/go-nostr)
* [shota3506/go-nostr (client)](https://github.com/shota3506/go-nostr)
* [nostr-domain](https://github.com/dextryz/nostr-domain)

## License

MIT License
Copyright (c) 2024 Niall Young <5465765+niallyoung@users.noreply.github.com>