package client

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/niallyoung/goNDK/event"
)

func TestSubscribeNoFilters(t *testing.T) {
	rm := NewRelayManager("wss://relay.damus.io")
	_, err := rm.Subscribe(context.Background(), []Filter{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "at least one filter")
}

func TestSubscriptionID(t *testing.T) {
	rm := NewRelayManager("wss://relay.damus.io")
	sub, err := rm.Subscribe(context.Background(), []Filter{{Kinds: []int{1}}})
	require.NoError(t, err)
	assert.NotEmpty(t, sub.ID())
}

func TestSubscriptionEOSE(t *testing.T) {
	rm := NewRelayManager("wss://relay.damus.io")
	sub, err := rm.Subscribe(context.Background(), []Filter{{Kinds: []int{1}}})
	require.NoError(t, err)
	assert.NotNil(t, sub.EOSE())
}

func TestNoticeChannel(t *testing.T) {
	rm := NewRelayManager("wss://relay.damus.io")
	noticeChan := rm.Notice()
	assert.NotNil(t, noticeChan)
}

func TestCloseIdempotent(t *testing.T) {
	rm := NewRelayManager("wss://relay.damus.io")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	err := rm.Connect(ctx)
	require.NoError(t, err)
	
	err1 := rm.Close()
	err2 := rm.Close()
	
	// Second close should not error
	assert.NoError(t, err1)
	assert.NoError(t, err2)
}

func TestPublishTimeout(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	
	rm := NewRelayManager("wss://relay.damus.io")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	err := rm.Connect(ctx)
	require.NoError(t, err)
	defer rm.Close()
	
	ev := event.NewEvent(1, "test", nil, nil, nil, nil, nil)
	ev.Sign("0000000000000000000000000000000000000000000000000000000000000001")
	
	// Use very short timeout to force timeout
	pubCtx, pubCancel := context.WithTimeout(ctx, 100*time.Millisecond)
	defer pubCancel()
	
	_, err = rm.Publish(pubCtx, ev)
	// May timeout or succeed depending on relay speed
	if err != nil {
		assert.Contains(t, err.Error(), "context deadline exceeded")
	}
}
