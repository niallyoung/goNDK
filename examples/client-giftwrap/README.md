# Gift Wrap CLI

Command-line tool for sending and receiving encrypted NOSTR messages and files using NIP-59 gift wrap.

## Installation

```bash
go build -o giftwrap .
```

## Usage

### Send Text Message

```bash
./giftwrap send --destination npub1xyz... --message "Hello, World!"
```

### Send File

```bash
cat image.png | ./giftwrap send --destination npub1xyz... --filename image.png
```

### Receive by Event ID

```bash
./giftwrap receive --event-id abc123...
```

### Watch for Incoming Messages

```bash
./giftwrap receive --watch
```

### Receive from Stdin

```bash
echo '{"kind":1059,...}' | ./giftwrap receive
```

## Options

### Send
- `-d, --destination` - Recipient npub (required)
- `-m, --message` - Text message to send
- `-f, --filename` - Filename for piped data
- `-r, --relay` - Relay URL (default: wss://relay.damus.io)
- `-k, --key` - Path to private key (default: ~/.nostr/key)

### Receive
- `-e, --event-id` - Event ID to fetch
- `-w, --watch` - Watch for incoming messages
- `-o, --output` - Output directory for files
- `-r, --relay` - Relay URL (default: wss://relay.damus.io)
- `-k, --key` - Path to private key (default: ~/.nostr/key)

## Key Management

Keys are automatically generated on first use and stored at `~/.nostr/key` (or path specified with `-k`).

To use an existing key, save your nsec to the key file:

```bash
echo "nsec1..." > ~/.nostr/key
chmod 600 ~/.nostr/key
```

## Features

- ✅ NIP-44 encryption
- ✅ NIP-59 gift wrap envelopes
- ✅ Text messages
- ✅ File transfers (base64 encoded)
- ✅ Automatic key generation
- ✅ Watch mode for real-time messages
- ✅ Stdin/stdout support for piping

## Architecture

```
Message → JSON → Inner Event → Sign → NIP-44 Encrypt → Gift Wrap → Relay
```

## Examples

### Send secret message
```bash
./giftwrap send -d npub1alice... -m "Meet at noon"
```

### Send encrypted file
```bash
cat secret.pdf | ./giftwrap send -d npub1bob... -f secret.pdf
```

### Watch for messages
```bash
./giftwrap receive --watch
```

### Fetch specific event
```bash
./giftwrap receive -e event123... -o ./downloads
```
