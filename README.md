# goNDK

[![Run Tests](https://github.com/niallyoung/goNDK/actions/workflows/main.yaml/badge.svg)](https://github.com/niallyoung/goNDK/actions/workflows/main.yaml)

goNDK is a NOSTR Development Kit in Golang

### Goals

- well-engineered framework of NOSTR types, funcs and interfaces
- 95%+ test coverage, for rapid prototyping and production-ready solutions at scale
- ensure the NOSTR development experience in Go is high-quality, productive, easy and fun

### Status

* initial `v0.0.x` Type build-out underway, commenced Easter 2024
* total test coverage = 96.97%+ (see `./.meta/cover.sh`)

- [x] `Event{}`
- [x] `Identity{}` **WIP**
- [x] `Client{}`, `RelayManager{}` **WIP**
  - [ ] `Subscription{}`, `SubscriptionFilter{}`
- [ ] `Relay{}`, `Transport{}`, `ClientManager{}`,
  - [ ] `Outbox{}`, `Inbox{}`

* ETA: ~Jul/Aug 2024?: functional `Client{}` publishing to `Relay{}` with network propagation
  * all types standalone and de-coupled
  * interfaces established, all types injected, all messages via interfaces

### Development

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

### Usage

```shell
go get github.com/niallyoung/goNDK@latest
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

### Identity{}

```go
import (
	"github.com/niallyoung/goNDK/identity"
)

i := identity.NewIdentity(pubkey, npub)
err := i.Validate()
```

### Client{}

```go
import (
	"github.com/niallyoung/goNDK/client"
)

c := client.NewClient()
err := c.Validate()
```

### THANKS

Built upon, around and inspired by:

* [go-nostr](https://github.com/nbd-wtf/go-nostr)
* [khatru](https://github.com/fiatjaf/khatru)
* [nostr-domain](https://github.com/dextryz/nostr-domain)
* [go-nostr (client)](https://github.com/shota3506/go-nostr)

MIT License
Copyright (c) 2024 Niall Young <5465765+niallyoung@users.noreply.github.com>