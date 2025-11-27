package client

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/niallyoung/goNDK/event"
)

func TestRelayManager_handleEventMessage_Errors(t *testing.T) {
	rm := NewRelayManager("wss://test")

	t.Run("unaddressed subscription", func(t *testing.T) {
		msg := &EventMessage{
			SubscriptionID: "unknown",
			Event:          &event.Event{},
		}
		err := rm.handleEventMessage(msg)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unaddressed event message")
	})

	t.Run("invalid value in map", func(t *testing.T) {
		rm.subMap.Store("test", "invalid")
		msg := &EventMessage{
			SubscriptionID: "test",
			Event:          &event.Event{},
		}
		err := rm.handleEventMessage(msg)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid value")
	})
}

func TestRelayManager_handleEOSEMessage_Errors(t *testing.T) {
	rm := NewRelayManager("wss://test")

	t.Run("unaddressed subscription", func(t *testing.T) {
		msg := &EOSEMessage{SubscriptionID: "unknown"}
		err := rm.handleEOSEMessage(msg)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unaddressed EOSE message")
	})

	t.Run("invalid value in map", func(t *testing.T) {
		rm.subMap.Store("test", "invalid")
		msg := &EOSEMessage{SubscriptionID: "test"}
		err := rm.handleEOSEMessage(msg)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid value")
	})
}

func TestRelayManager_handleOKMessage_Errors(t *testing.T) {
	rm := NewRelayManager("wss://test")

	t.Run("unaddressed event", func(t *testing.T) {
		msg := &OKMessage{EventID: "unknown", OK: true}
		err := rm.handleOKMessage(msg)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unaddressed OK message")
	})

	t.Run("invalid value in map", func(t *testing.T) {
		rm.eventMap.Store("test", "invalid")
		msg := &OKMessage{EventID: "test", OK: true}
		err := rm.handleOKMessage(msg)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid value")
	})
}

func TestRelayManager_handleNoticeMessage(t *testing.T) {
	rm := NewRelayManager("wss://test")

	t.Run("sends to channel", func(t *testing.T) {
		msg := &NoticeMessage{Message: "test"}
		err := rm.handleNoticeMessage(msg)
		assert.NoError(t, err)

		select {
		case notice := <-rm.noticeChan:
			assert.Equal(t, "test", notice)
		default:
			t.Fatal("notice not received")
		}
	})

	t.Run("drops when channel full", func(t *testing.T) {
		// Fill channel
		for i := 0; i < 100; i++ {
			rm.noticeChan <- "fill"
		}

		msg := &NoticeMessage{Message: "dropped"}
		err := rm.handleNoticeMessage(msg)
		assert.NoError(t, err) // Should not error, just drop
	})
}
