# goNDK

[![Run Tests](https://github.com/niallyoung/goNDK/actions/workflows/main.yaml/badge.svg)](https://github.com/niallyoung/goNDK/actions/workflows/main.yaml)

goNDK is a NOSTR Development Kit in Golang

### GOALS

* cohesive, well tested, injectable, idiomatic low-level funcs, types and interfaces
* intended for rapid prototyping and production-ready solutions at scale
* (TODO) a reference library of NOSTR Clients, Relays, go funcs, helpers
* help NOSTR to grow: make it easy to build and run NOSTR relays and clients in Go

### STATUS

goNDK was established Easter 2024, and is a **Work In Progress**

### USAGE

```shell
make lint     // golangci-lint
make test     // unit tests
make cover    // total coverage %, calls ./.meta/cover.sh
make generate // code generation (easyjson)

make docker.build
make docker.lint
make docker.test
make docker.cover
make docker.shell
```

### EXAMPLE

```shell
go get github.com/niallyoung/goNDK@latest
```

```go
import (
    "github.com/niallyoung/goNDK/event"
)

// unmarshal an incoming JSON serialised Event
var event event.Event
err := json.Unmarshal([]byte(`{"kind": 1, "content": "hello world!", ... }`), &event)

// create a new, unsigned event
var e := event.NewEvent(1, "hello world!", event.Tags(nil), nil, nil, nil, nil)

err := e.Validate() // Event value validation

err := e.Sign(privateKey)        // Sign an Event
ok, err := e.ValidateSignature() // Validate its Signature

text := e.String()     // Event JSON serialization
bytes := e.Serialize() // custom JSON Array serialization, for identity / authentication
```

### TODO

- `Client{}`, `Identity{}`, stream events from relays, publish to relays
- `Relay{}`, `Transport{}`, `ConnectionHandler{}`, `Outbox{}`, `Inbox{}`

### THANKS

Built upon, around and inspired by:

* [go-nostr](https://github.com/nbd-wtf/go-nostr)
* [khatru](https://github.com/fiatjaf/khatru)
* [nostr-domain](https://github.com/dextryz/nostr-domain)

MIT License
Copyright (c) 2024 Niall Young <5465765+niallyoung@users.noreply.github.com>