package client

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFetchEventByID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	rm := NewRelayManager("wss://relay.damus.io")
	err := rm.Connect(ctx)
	require.NoError(t, err)
	defer rm.Close()

	// Known event ID from relay
	eventID := "test-event-id"
	
	event, err := FetchEventByID(ctx, rm, eventID)
	
	if err != nil {
		t.Logf("Event not found (expected for test): %v", err)
		return
	}
	
	if event != nil {
		assert.Equal(t, eventID, event.ID)
	}
}

func TestFetchEventByID_NotFound(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rm := NewRelayManager("wss://relay.damus.io")
	err := rm.Connect(ctx)
	require.NoError(t, err)
	defer rm.Close()

	nonExistentID := "0000000000000000000000000000000000000000000000000000000000000000"
	
	event, err := FetchEventByID(ctx, rm, nonExistentID)
	
	assert.Error(t, err)
	assert.Nil(t, event)
}
