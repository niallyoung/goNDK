package event_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/niallyoung/goNDK/event"
)

var jsonEvent = `{"pubkey": "fake-pubkey", "created_at:" 0, "kind": 1, "content": "fake-content", "sig": "fake-sig"}`
var newEvent = event.NewEvent("fake-pubkey", time.Unix(0, 0), 1, event.Tags{}, "fake-content", "fake-sig")
var validEvent = &event.Event{PubKey: "fake-pubkey", CreatedAt: time.Unix(0, 0), Kind: 1, Tags: event.Tags{}, Content: "fake-content", Sig: "fake-sig"}

func TestNewEvent(t *testing.T) {
	t.Run("NewEvent() returns an Event", func(t *testing.T) {
		e := newEvent
		assert.NotNil(t, e)
	})

	t.Run("new Event has an ID", func(t *testing.T) {
		e := newEvent
		assert.Equal(t, "9d004c6d691bb765165f30dfa8854355033025b12f564707331d8de5f0a77a72", e.ID())
	})
}

func TestEvent_Validate(t *testing.T) {
	t.Run("validate", func(t *testing.T) {
		e := validEvent
		err := e.Validate()
		assert.NoError(t, err)
	})
}
