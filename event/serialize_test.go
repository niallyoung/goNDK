package event_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvent_Serialize(t *testing.T) {
	t.Run("NewEvent() returns an Event", func(t *testing.T) {
		e, _ := validEvent()
		bytes := e.Serialize()
		assert.Equal(t, validEventSerialize, string(bytes))
	})
}

func TestEvent_String(t *testing.T) {
	t.Run("validEvent.String() returns expected string", func(t *testing.T) {
		e, _ := validEvent()
		assert.Equal(t, validEventString, e.String())
	})
}
