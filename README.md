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

* initial `v0.0.x` Type build-out underway, commenced Easter 2024
* total unit test coverage = 96.97%+ (see `./.meta/cover.sh`)
* integration tests pending stable interfaces, after types established
* v0.1.0 release should actually work

- [x] `Event{}`
- [x] `Identity{}` **WIP**
  - [ ] `IdentityProvider{}`? NWC / Keystore integration
- [x] `Client{}`, `RelayManager{}` **WIP**
  - [ ] `Subscription{}`, `SubscriptionFilter{}`
  - [ ] `CommandlineClient{}`, ...
- [ ] `*Config{}`
- [ ] `Relay{}`, `Transport{}`, `ClientManager{}`,
  - [ ] `Outbox{}`, `Inbox{}`
  - [ ] `LocalRelay{}`, `ProxyRelay{}`, `CommandlineRelay{}`, ...

~v0.1.0 ETA: ~Jul/Aug 2024: functional `Client{}` publishing `Event{}`s
  * cohesive types, moderately de-coupled
  * interfaces established, all messaging via interfaces
  * publishing to a public relay with successful downstream propagation

~v0.2.0 ETA: ~Sep/Oct 2024: functional `Relay{}`
  * all dependencies injected, optional adapters to shim to interfaces
  * mocks generated from all interfaces, refactor unit tests with injection
  * `Client{}` publishing to `Relay{}` with successful downstream propagation

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