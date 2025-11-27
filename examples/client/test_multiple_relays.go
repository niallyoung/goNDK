package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/niallyoung/goNDK/client"
	"github.com/niallyoung/goNDK/event"
)

func testRelay(relayURL string) (int, bool, error) {
	rm := client.NewRelayManager(relayURL)
	ctx := context.Background()

	if err := rm.Connect(ctx); err != nil {
		return 0, false, err
	}
	defer rm.Close()

	filters := []client.Filter{{Kinds: []int{1}, Limit: 10}}
	sub, err := rm.Subscribe(ctx, filters)
	if err != nil {
		return 0, false, err
	}

	receiveCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	eventCount := 0
	go sub.Receive(receiveCtx, func(_ context.Context, e *event.Event) {
		eventCount++
	})

	eoseReceived := false
	select {
	case <-sub.EOSE():
		eoseReceived = true
	case <-receiveCtx.Done():
	}

	time.Sleep(200 * time.Millisecond)
	return eventCount, eoseReceived, nil
}

func main() {
	relays := []string{
		"wss://relay.damus.io",
		"wss://relay.primal.net",
		"wss://nos.lol",
		"wss://relay.nostr.band",
	}

	fmt.Println("🧪 Testing goNDK Client with Public Relays\n")

	passed := 0
	failed := 0

	for _, relay := range relays {
		fmt.Printf("Testing %s... ", relay)
		
		count, eose, err := testRelay(relay)
		
		if err != nil {
			fmt.Printf("❌ FAILED: %v\n", err)
			failed++
			continue
		}

		if count > 0 {
			fmt.Printf("✅ PASSED (%d events, EOSE: %v)\n", count, eose)
			passed++
		} else {
			fmt.Printf("⚠️  WARNING: Connected but no events received\n")
			failed++
		}
	}

	fmt.Printf("\n📊 Results: %d passed, %d failed\n", passed, failed)
	
	if passed > 0 {
		fmt.Println("✅ Integration tests PASSED")
	} else {
		log.Fatal("❌ All integration tests FAILED")
	}
}
