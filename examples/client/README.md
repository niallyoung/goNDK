# Client Examples

## Running Examples

### Fetch Events from Public Relay
```bash
cd examples/client
go run fetch_events.go
```

Connects to Damus relay, subscribes to recent text notes, and displays them.

### Test Multiple Relays
```bash
cd examples/client
go run test_multiple_relays.go
```

Tests connectivity and event fetching from multiple public relays.

## Integration Tests

Run integration tests against real public relays:

```bash
# From repository root
go test -tags=integration ./client -v

# Skip in CI/short mode
go test -short ./client
```

## Public Relays Used

- `wss://relay.damus.io` - Damus relay
- `wss://relay.primal.net` - Primal relay
- `wss://nos.lol` - Nos relay
- `wss://relay.nostr.band` - Nostr Band relay

All relays tested and working as of 2025-01-27.
