package event_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/niallyoung/goNDK/event"
)

var validEvent = event.NewEvent("fake-pubkey", time.Unix(0, 0), 1, event.Tags{}, "fake-content", "fake-sig")

func TestNewEvent(t *testing.T) {
	t.Run("NewEvent() returns an Event", func(t *testing.T) {
		e := validEvent
		assert.NotNil(t, e)
	})

	t.Run("new event has an ID", func(t *testing.T) {
		e := validEvent
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
