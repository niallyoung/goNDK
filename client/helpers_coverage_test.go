package client

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFetchEventByID_ContextCanceled(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	rm := NewRelayManager("wss://relay.damus.io")
	connectCtx := context.Background()
	err := rm.Connect(connectCtx)
	if err != nil {
		t.Skip("Could not connect to relay")
	}
	defer rm.Close()

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	_, err = FetchEventByID(ctx, rm, "test")
	assert.Error(t, err)
}

func TestFetchEventByID_Timeout(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rm := NewRelayManager("wss://relay.damus.io")
	err := rm.Connect(ctx)
	if err != nil {
		t.Skip("Could not connect to relay")
	}
	defer rm.Close()

	// Non-existent event
	_, err = FetchEventByID(ctx, rm, "0000000000000000000000000000000000000000000000000000000000000000")
	assert.Error(t, err)
	assert.Equal(t, ErrEventNotFound, err)
}
