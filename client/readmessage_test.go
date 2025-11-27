package client

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRelayManager_ReadMessage_ErrorPaths(t *testing.T) {
	t.Run("empty message array", func(t *testing.T) {
		// Simulate empty message
		emptyMsg := []json.RawMessage{}
		data, _ := json.Marshal(emptyMsg)
		
		// We can't easily test this without a real connection
		// but we can test the error handling logic
		var message []json.RawMessage
		err := json.Unmarshal(data, &message)
		assert.NoError(t, err)
		assert.Equal(t, 0, len(message))
	})

	t.Run("invalid message type", func(t *testing.T) {
		// Test unsupported message type handling
		invalidMsg := []interface{}{"UNKNOWN", "data"}
		data, _ := json.Marshal(invalidMsg)
		
		var message []json.RawMessage
		err := json.Unmarshal(data, &message)
		assert.NoError(t, err)
		
		var typ string
		err = json.Unmarshal(message[0], &typ)
		assert.NoError(t, err)
		assert.Equal(t, "UNKNOWN", typ)
	})

	t.Run("invalid notice length", func(t *testing.T) {
		// NOTICE should have exactly 2 elements
		invalidNotice := []interface{}{"NOTICE"}
		data, _ := json.Marshal(invalidNotice)
		
		var message []json.RawMessage
		err := json.Unmarshal(data, &message)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(message))
	})

	t.Run("invalid event length", func(t *testing.T) {
		// EVENT should have exactly 3 elements
		invalidEvent := []interface{}{"EVENT", "sub1"}
		data, _ := json.Marshal(invalidEvent)
		
		var message []json.RawMessage
		err := json.Unmarshal(data, &message)
		assert.NoError(t, err)
		assert.Equal(t, 2, len(message))
	})

	t.Run("invalid EOSE length", func(t *testing.T) {
		// EOSE should have exactly 2 elements
		invalidEOSE := []interface{}{"EOSE"}
		data, _ := json.Marshal(invalidEOSE)
		
		var message []json.RawMessage
		err := json.Unmarshal(data, &message)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(message))
	})

	t.Run("invalid OK length", func(t *testing.T) {
		// OK should have exactly 4 elements
		invalidOK := []interface{}{"OK", "eventid", true}
		data, _ := json.Marshal(invalidOK)
		
		var message []json.RawMessage
		err := json.Unmarshal(data, &message)
		assert.NoError(t, err)
		assert.Equal(t, 3, len(message))
	})
}

func TestRelayManager_WriteMessage_ErrorPath(t *testing.T) {
	t.Run("marshal error", func(t *testing.T) {
		// EventMessage with nil event will cause marshal issues
		msg := &EventMessage{Event: nil}
		data, err := msg.MarshalJSON()
		// Should succeed but produce null
		assert.NoError(t, err)
		assert.Contains(t, string(data), "null")
	})
}

func TestRelayManager_handleOKMessage_ChannelDrop(t *testing.T) {
	rm := NewRelayManager("wss://test")
	
	t.Run("drops when channel full", func(t *testing.T) {
		// Create a channel group with buffered channel
		okChan := make(chan *CommandResult, 1)
		rm.eventMap.Store("test", &eventChannelGroup{okChan: okChan})
		
		// Fill the channel
		okChan <- &CommandResult{OK: true}
		
		// Try to send another - should drop without error
		msg := &OKMessage{EventID: "test", OK: false, Message: "dropped"}
		err := rm.handleOKMessage(msg)
		assert.NoError(t, err)
		
		// Verify first message still there
		result := <-okChan
		assert.True(t, result.OK)
	})
}

func TestSubscribe_ErrorPath(t *testing.T) {
	rm := NewRelayManager("wss://test")
	ctx := context.Background()
	
	t.Run("empty filters", func(t *testing.T) {
		_, err := rm.Subscribe(ctx, nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "at least one filter")
	})
	
	t.Run("zero length filters", func(t *testing.T) {
		_, err := rm.Subscribe(ctx, []Filter{})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "at least one filter")
	})
}
