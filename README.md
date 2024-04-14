# goNDK

[![Run Tests](https://github.com/niallyoung/goNDK/actions/workflows/main.yaml/badge.svg)](https://github.com/niallyoung/goNDK/actions/workflows/main.yaml)

goNDK is a NOSTR Development Kit, for Golang.

## Goals

* a solid NOSTR foundation in Go, useful for rapid prototyping and production-ready solutions at scale
* cohesive, well tested, injectable, idiomatic low-level funcs, types and interfaces
* a variety of simple/example clients, relays, go funcs, helpers
* help NOSTR to grow: make it easy to build and run NOSTR relays and clients in Go

## TODO

- `Client{}`, `Identity{}`, stream events from relays, publish to relays
- `Relay{}`, `Transport{}`, `ConnectionHandler{}`, `Outbox{}`, `Inbox{}`

## Thanks

Built upon, around and inspired by:

* [go-nostr](https://github.com/nbd-wtf/go-nostr)
* [khatru](https://github.com/fiatjaf/khatru)
* [nostr-domain](https://github.com/dextryz/nostr-domain)

MIT License
Copyright (c) 2024 Niall Young <5465765+niallyoung@users.noreply.github.com>