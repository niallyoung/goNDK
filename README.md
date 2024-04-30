# goNDK

[![Run Tests](https://github.com/niallyoung/goNDK/actions/workflows/main.yaml/badge.svg)](https://github.com/niallyoung/goNDK/actions/workflows/main.yaml)

goNDK is a NOSTR Development Kit in Golang

### GOALS

- well-engineered framework of NOSTR types, funcs and interfaces
- 95%+ test coverage, for rapid prototyping and production-ready solutions at scale
- help NOSTR development in Go to be high-quality, fast, easy and fun

### STATUS

* initial type build underway, commenced Easter 2024

- [x] `Event{}`
- [ ] `Identity{}` WIP
- [ ] `Client{}`, `RelayManager{}` WIP
  - [ ] `Subscription{}`, `SubscriptionFilter{}`
- [ ] `Relay{}`, `Transport{}`, `ClientManager{}`,
- [ ] `Outbox{}`, `Inbox{}`

### DEVELOPMENT

```shell
make lint     # golangci-lint
make test     # unit tests
make cover    # 95%+ coverage
make generate # code generation (easyjson)

make docker.build
make docker.lint
make docker.test
make docker.cover
make docker.shell
```

### USAGE

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

Engineered upon, around and inspired by:

* [go-nostr](https://github.com/nbd-wtf/go-nostr)
* [khatru](https://github.com/fiatjaf/khatru)
* [nostr-domain](https://github.com/dextryz/nostr-domain)
* [go-nostr (client)](https://github.com/shota3506/go-nostr)

MIT License
Copyright (c) 2024 Niall Young <5465765+niallyoung@users.noreply.github.com>