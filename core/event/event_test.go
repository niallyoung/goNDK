package event_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/niallyoung/gondo/core/event"
)

func TestNewEvent(t *testing.T) {
	t.Run("NewEvent() returns an Event", func(t *testing.T) {
		e := event.NewEvent()
		assert.NotNil(t, e)
	})
}
