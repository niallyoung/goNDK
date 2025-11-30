# keygen - NOSTR Keypair Generator

Generate NOSTR keypairs (nsec/npub) for testing and development.

## Installation

```bash
go install github.com/niallyoung/goNDK/examples/keygen@latest
```

Or build locally:

```bash
cd examples/keygen
go build
```

## Usage

### Default (bech32 format)
```bash
$ keygen
nsec1...|npub1...
```

### JSON format
```bash
$ keygen -format json
{
  "nsec": "nsec1...",
  "npub": "npub1...",
  "privkey_hex": "...",
  "pubkey_hex": "..."
}
```

### Hex format
```bash
$ keygen -format hex
<64-char-privkey>|<64-char-pubkey>
```

## Examples

### Generate and store in variables
```bash
KEY=$(keygen)
NSEC=$(echo $KEY | cut -d'|' -f1)
NPUB=$(echo $KEY | cut -d'|' -f2)
echo "Private: $NSEC"
echo "Public: $NPUB"
```

### Generate and save to files
```bash
keygen -format json > keypair.json
```

### Use in scripts
```bash
#!/bin/bash
KEY=$(keygen)
NSEC=$(echo $KEY | cut -d'|' -f1)

# Store in AWS SSM
aws ssm put-parameter \
  --name "/my-service/PrivateKey" \
  --value "$NSEC" \
  --type "SecureString"
```

## Output Formats

### bech32 (default)
- Format: `nsec1...|npub1...`
- Use: Human-readable, NOSTR standard
- Length: 63 characters each

### hex
- Format: `<privkey>|<pubkey>`
- Use: Programming, cryptographic operations
- Length: 64 hex characters each

### json
- Format: JSON object with all representations
- Use: Structured data, APIs
- Fields: nsec, npub, privkey_hex, pubkey_hex

## Testing

```bash
go test -v
```

## Security Notes

- **Private keys (nsec)**: Keep secret, never commit to git
- **Public keys (npub)**: Safe to share
- **Production**: Use hardware security modules or key management services
- **Storage**: Encrypt private keys at rest

## See Also

- [NIP-19: bech32-encoded entities](https://github.com/nostr-protocol/nips/blob/master/19.md)
- [goNDK identity package](../../identity/)
