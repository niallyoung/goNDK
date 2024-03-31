package event_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/niallyoung/goNDK/event"
)

func TestNewEvent(t *testing.T) {
	t.Run("NewEvent() returns an Event", func(t *testing.T) {
		e := event.NewEvent("fake-pubkey", time.Unix(0, 0), 1, event.Tags{}, "fake-content", "fake-sig")
		assert.NotNil(t, e)
	})
}
