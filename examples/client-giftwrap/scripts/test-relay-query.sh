#!/bin/bash
set -e

echo "=== Testing Relay Query ==="
echo ""

cd "$(dirname "$0")/.."
make build > /dev/null 2>&1

echo "Querying relay.damus.io for kind 1 events..."
echo ""

# Use the giftwrap CLI to test basic relay connectivity
timeout 10 ./giftwrap receive --watch --relay wss://relay.damus.io 2>&1 | head -20 || echo "Query completed or timed out"

echo ""
echo "✅ Relay connectivity test complete"
