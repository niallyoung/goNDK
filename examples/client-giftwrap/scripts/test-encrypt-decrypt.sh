#!/bin/bash
set -e

echo "=== Gift Wrap Encrypt/Decrypt Test ==="
echo ""

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
GIFTWRAP_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"
GONDK_DIR="$(cd "$GIFTWRAP_DIR/../.." && pwd)"

cd "$GIFTWRAP_DIR"

# Build tools
echo "Building tools..."
make build > /dev/null 2>&1
cd "$GONDK_DIR/examples/identity-generator" && make build > /dev/null 2>&1
cd "$GIFTWRAP_DIR"

# Create test message
TEST_MESSAGE="Hello NOSTR! This is a test of NIP-59 gift wrap encryption."
echo "1. Test message: \"$TEST_MESSAGE\""
echo ""

# Generate fresh receiver keypair
echo "2. Generating receiver keypair..."
RECEIVER_OUTPUT=$("$GONDK_DIR/examples/identity-generator/identity-generator")
RECEIVER_NPUB=$(echo "$RECEIVER_OUTPUT" | grep npub | awk '{print $2}')
RECEIVER_NSEC=$(echo "$RECEIVER_OUTPUT" | grep nsec | awk '{print $2}')

echo "   Receiver: $RECEIVER_NPUB"
echo "   (nsec: $RECEIVER_NSEC)"
echo ""

# Save receiver key temporarily
TMPDIR="/tmp/giftwrap-test-$$"
mkdir -p "$TMPDIR"
trap "rm -rf $TMPDIR" EXIT
echo "$RECEIVER_NSEC" > "$TMPDIR/receiver.key"

# Create encrypted event (generates sender automatically)
echo "3. Creating encrypted gift wrap event..."
./giftwrap create --destination "$RECEIVER_NPUB" --message "$TEST_MESSAGE" > "$TMPDIR/event.json" 2>&1

# Extract just the JSON (skip the "Generated sender" line)
sed -i.bak '/Generated sender/d' "$TMPDIR/event.json"

echo "   ✓ Event created"
echo ""

# Decrypt
echo "4. Decrypting gift wrap event..."
DECRYPTED=$(cat "$TMPDIR/event.json" | ./giftwrap receive --key "$TMPDIR/receiver.key" 2>&1 | grep "Message:" | sed 's/📨 Message: //')

echo "   Decrypted: \"$DECRYPTED\""
echo ""

# Verify
if [ "$DECRYPTED" = "$TEST_MESSAGE" ]; then
    echo "✅ SUCCESS: Message encrypted and decrypted correctly!"
    exit 0
else
    echo "❌ FAILED: Decrypted message doesn't match"
    echo "   Expected: $TEST_MESSAGE"
    echo "   Got:      $DECRYPTED"
    exit 1
fi
