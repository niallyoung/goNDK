# Gift Wrap CLI Tool Plan

## Goal
Command-line tool to send/receive encrypted messages and files via NOSTR gift wrap.

## Commands

### `giftwrap send`
```bash
# Text message
giftwrap send --destination npub1xyz... --message "Hello"

# File
cat image.png | giftwrap send --destination npub1xyz... --filename image.png

# With relay
giftwrap send --destination npub1xyz... --message "Hello" --relay wss://relay.damus.io
```

### `giftwrap receive`
```bash
# By event ID
giftwrap receive --event-id abc123... --relay wss://relay.damus.io

# Listen for incoming
giftwrap receive --relay wss://relay.damus.io --watch

# Pipe event JSON
echo '{"kind":1059,...}' | giftwrap receive
```

## What We Have

### From goNDK
- ✅ `nips/nip44` - Encryption/decryption
- ✅ `nips/nip59` - Wrap/Unwrap
- ✅ `client/RelayManager` - Connect, publish, subscribe
- ✅ `event/Event` - Event structure
- ✅ `identity/Identity` - npub/nsec handling (partial)

## What We Need

### 1. NIP-19 Support (npub/nsec conversion)
**Status:** Missing
**Location:** `identity/nip19.go`
**Dependency:** `github.com/btcsuite/btcd/btcutil/bech32`

### 2. Identity Generation
**Status:** Missing
**Location:** `identity/generate.go`

### 3. File Metadata Encoding
**Status:** Missing
**Location:** `examples/client-giftwrap/internal/file.go`

### 4. Event Fetching Helper
**Status:** Partial - need single event fetch
**Location:** `client/helpers.go`

### 5. Key Storage
**Status:** Missing
**Location:** `examples/client-giftwrap/internal/keys.go`

## Implementation Phases

### Phase 1: Core (Minimal)
1. ✅ NIP-44 encryption
2. ✅ NIP-59 gift wrap
3. ⏳ npub/nsec conversion
4. ⏳ Basic send (text only)
5. ⏳ Basic receive (text only)

### Phase 2: File Support
1. ⏳ File metadata encoding
2. ⏳ Base64 for binary
3. ⏳ File output to disk

### Phase 3: Key Management
1. ⏳ Key generation
2. ⏳ Key storage
3. ⏳ Environment variables



## New Code Needed

### identity/nip19.go
```go
func NpubToHex(npub string) (string, error)
func NsecToHex(nsec string) (string, error)
func HexToNpub(hex string) (string, error)
func HexToNsec(hex string) (string, error)
```

### identity/generate.go
```go
func Generate() (*Identity, error)
func FromNsec(nsec string) (*Identity, error)
```

### client/helpers.go
```go
func FetchEventByID(ctx, rm, eventID) (*Event, error)
```

### examples/client-giftwrap/internal/file.go
```go
type FileMessage struct {
    Type     string
    Content  string
    Metadata *FileMetadata
}
```

## Dependencies

### New
- github.com/btcsuite/btcd/btcutil/bech32
- github.com/spf13/cobra

## Success Criteria

- ✅ Send text messages
- ✅ Send files
- ✅ Receive and decrypt
- ✅ Save files to disk
- ✅ Clean, generic implementation
