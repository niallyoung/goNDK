# goNDK

[![Run Tests](https://github.com/niallyoung/goNDK/actions/workflows/main.yaml/badge.svg)](https://github.com/niallyoung/goNDK/actions/workflows/main.yaml)

goNDK is a NOSTR Development Kit, for Golang.

## Goals

* a solid NOSTR foundation in Go, useful for rapid prototyping and prod-ready solutions at scale
* cohesive, well tested, injectable, idiomatic low-level funcs, types and interfaces
* future?: build up from there - a variety of simple/example clients, relays, go funcs, helpers
* help NOSTR grow, make it easy to run, build, publish and read NOSTR

## Status

WIP, not even at v0.0.1 yet but very soon... ETA: mid-April 2024

NOTE: I'm focused on the fundamantals here, not any particular architecture. I want this to be useful, generic, fun and
easy-to-use, for myself **and others** to build upon. I'm building from the ground up, deliberately slow and methodical.
For compatibility I've started with the same ~Event serialization shared across `khatru`, `go-noster` and `nostr-domain`
but the rest is a clean-room implementation.


## TODO
- WIP: test coverage Event{}, EventRequest{} stub
- v0.0.1 tag/release
- flesh out Relay{}, funcs, types, interfaces, tests
- flesh out relay/*: ConnectionHandler, Outbox, Inbox

## Thanks

Built upon, around and inspired by:
* [go-nostr](https://github.com/nbd-wtf/go-nostr)
* [khatru](https://github.com/fiatjaf/khatru)
* [nostr-domain](https://github.com/dextryz/nostr-domain)

MIT License
Copyright (c) 2024 Niall Young <5465765+niallyoung@users.noreply.github.com>