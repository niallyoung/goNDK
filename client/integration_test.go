// +build integration

package client_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/niallyoung/goNDK/client"
	"github.com/niallyoung/goNDK/event"
)

// Run with: go test -tags=integration ./client

func TestIntegration_PublicRelay(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	relays := []string{
		"wss://relay.damus.io",
		"wss://relay.primal.net",
		"wss://nos.lol",
	}

	for _, relayURL := range relays {
		t.Run(relayURL, func(t *testing.T) {
			rm := client.NewRelayManager(relayURL)
			ctx := context.Background()

			err := rm.Connect(ctx)
			require.NoError(t, err, "Failed to connect to %s", relayURL)
			defer rm.Close()

			filters := []client.Filter{{Kinds: []int{1}, Limit: 5}}
			sub, err := rm.Subscribe(ctx, filters)
			require.NoError(t, err)

			receiveCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
			defer cancel()

			eventCount := 0
			go sub.Receive(receiveCtx, func(_ context.Context, e *event.Event) {
				eventCount++
				assert.NotNil(t, e.ID)
				assert.NotNil(t, e.Pubkey)
				assert.Equal(t, 1, e.Kind)
			})

			select {
			case <-sub.EOSE():
				// Success
			case <-time.After(10 * time.Second):
				t.Fatal("Timeout waiting for EOSE")
			}

			time.Sleep(200 * time.Millisecond)
			assert.Greater(t, eventCount, 0, "Should receive at least one event")
		})
	}
}
