# Gift Wrap CLI Examples

## Basic Text Message

```bash
# Generate keys automatically on first use
./giftwrap send \
  --destination npub180cvv07tjdrrgpa0j7j7tmnyl2yr6yr7l8j4s3evf6u64th6gkwsyjh6w6 \
  --message "Hello from NOSTR!"
```

## Send File

```bash
# Send an image
cat photo.jpg | ./giftwrap send \
  --destination npub180cvv07tjdrrgpa0j7j7tmnyl2yr6yr7l8j4s3evf6u64th6gkwsyjh6w6 \
  --filename photo.jpg

# Send a PDF
cat document.pdf | ./giftwrap send \
  --destination npub180cvv07tjdrrgpa0j7j7tmnyl2yr6yr7l8j4s3evf6u64th6gkwsyjh6w6 \
  --filename document.pdf
```

## Receive Messages

```bash
# Fetch specific event
./giftwrap receive \
  --event-id abc123... \
  --relay wss://relay.damus.io

# Watch for incoming messages (real-time)
./giftwrap receive --watch

# Receive from stdin (pipe event JSON)
echo '{"kind":1059,"content":"..."}' | ./giftwrap receive
```

## Save Files

```bash
# Save received file to specific directory
./giftwrap receive \
  --event-id abc123... \
  --output ./downloads
```

## Custom Relay

```bash
# Use different relay
./giftwrap send \
  --destination npub180cvv07tjdrrgpa0j7j7tmnyl2yr6yr7l8j4s3evf6u64th6gkwsyjh6w6 \
  --message "Hello" \
  --relay wss://nos.lol
```

## Custom Key Location

```bash
# Use key from custom location
./giftwrap send \
  --destination npub180cvv07tjdrrgpa0j7j7tmnyl2yr6yr7l8j4s3evf6u64th6gkwsyjh6w6 \
  --message "Hello" \
  --key /path/to/my/key
```

## Complete Workflow

### Sender Side

```bash
# 1. Create a message
echo "Secret plans" > message.txt

# 2. Send it
cat message.txt | ./giftwrap send \
  --destination npub1alice... \
  --filename message.txt

# Output:
# 🔑 Loading key from ~/.nostr/key...
# ✓ Using identity: npub1bob...
# 🔌 Connecting to wss://relay.damus.io...
# ✓ Connected
# 📤 Publishing...
# ✅ Sent to npub1alice...
# Event ID: abc123...
```

### Receiver Side

```bash
# 1. Fetch the message
./giftwrap receive \
  --event-id abc123... \
  --output ./received

# Output:
# 🔌 Connecting to wss://relay.damus.io...
# 🔍 Fetching event abc123...
# 📎 File saved: ./received/message.txt (13 bytes)

# 2. Read the message
cat ./received/message.txt
# Secret plans
```

## Watch Mode Example

```bash
# Terminal 1 (receiver watching)
./giftwrap receive --watch
# 👀 Watching for messages...

# Terminal 2 (sender)
./giftwrap send \
  --destination npub1receiver... \
  --message "Real-time message!"

# Terminal 1 output:
# 📨 Message: Real-time message!
```

## Piping Between Commands

```bash
# Generate event and save to file
./giftwrap send \
  --destination npub1alice... \
  --message "Test" > event.json

# Later, decrypt from file
cat event.json | ./giftwrap receive
```

## Multiple Files

```bash
# Send multiple files in sequence
for file in *.jpg; do
  cat "$file" | ./giftwrap send \
    --destination npub1alice... \
    --filename "$file"
  sleep 1
done
```

## Error Handling

```bash
# Invalid destination
./giftwrap send --destination invalid --message "test"
# Error: invalid destination: decode bech32: ...

# Event not found
./giftwrap receive --event-id nonexistent
# Error: fetch event: event not found

# Connection timeout
./giftwrap send \
  --destination npub1alice... \
  --message "test" \
  --relay wss://invalid.relay
# Error: connect to relay: ...
```

## Tips

1. **Key Security**: Keep your `~/.nostr/key` file secure with `chmod 600`
2. **Relay Selection**: Use reliable relays for better delivery
3. **File Size**: Large files work but may take longer to encrypt/decrypt
4. **Watch Mode**: Use `--watch` for real-time messaging applications
5. **Output Directory**: Always specify `--output` when receiving files to avoid overwriting
