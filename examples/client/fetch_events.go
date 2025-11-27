package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/niallyoung/goNDK/client"
	"github.com/niallyoung/goNDK/event"
)

func main() {
	// Connect to Damus relay
	relayURL := "wss://relay.damus.io"
	fmt.Printf("Connecting to %s...\n", relayURL)

	rm := client.NewRelayManager(relayURL)
	ctx := context.Background()

	if err := rm.Connect(ctx); err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer rm.Close()

	fmt.Println("✅ Connected successfully")

	// Subscribe to recent text notes (kind 1)
	filters := []client.Filter{
		{
			Kinds: []int{1}, // Text notes
			Limit: 12,       // Fetch 12 events
		},
	}

	sub, err := rm.Subscribe(ctx, filters)
	if err != nil {
		log.Fatalf("Failed to subscribe: %v", err)
	}

	fmt.Println("📡 Subscribed, waiting for events...")

	// Receive events with timeout
	receiveCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	eventCount := 0
	eoseReceived := false

	go func() {
		err := sub.Receive(receiveCtx, func(_ context.Context, e *event.Event) {
			eventCount++
			fmt.Printf("\n📨 Event #%d:\n", eventCount)
			fmt.Printf("  ID: %s\n", *e.ID)
			fmt.Printf("  Kind: %d\n", e.Kind)
			fmt.Printf("  Author: %s\n", *e.Pubkey)
			fmt.Printf("  Content: %.80s...\n", e.Content)
			fmt.Printf("  Created: %s\n", time.Unix(int64(e.CreatedAt), 0).Format(time.RFC3339))
		})
		if err != nil && err != context.DeadlineExceeded {
			log.Printf("Receive error: %v", err)
		}
	}()

	// Wait for EOSE or timeout
	select {
	case <-sub.EOSE():
		eoseReceived = true
		fmt.Println("\n✅ EOSE received (all stored events sent)")
	case <-receiveCtx.Done():
		fmt.Println("\n⏱️  Timeout reached")
	}

	// Give a moment for final events to process
	time.Sleep(500 * time.Millisecond)

	fmt.Printf("\n📊 Summary:\n")
	fmt.Printf("  Events received: %d\n", eventCount)
	fmt.Printf("  EOSE received: %v\n", eoseReceived)
	fmt.Printf("  Relay: %s\n", relayURL)

	if eventCount > 0 {
		fmt.Println("\n✅ Integration test PASSED")
	} else {
		fmt.Println("\n❌ Integration test FAILED - no events received")
	}
}
