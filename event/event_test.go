package event_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/niallyoung/goNDK/event"
)

func TestNewEvent(t *testing.T) {
	t.Run("given ValidEvent(), NewEvent() returns an Event", func(t *testing.T) {
		e := ValidEvent()
		assert.NotNil(t, e)
	})

	t.Run("given ValidEventMinimal(), NewEvent() returns an Event", func(t *testing.T) {
		e := ValidEventMinimal()
		assert.NotNil(t, e)
	})
}

func TestEvent_MarshalJSON(t *testing.T) {
	t.Run("unmarshal ValidEventJSON to Event{}", func(t *testing.T) {
		var e event.Event
		err := json.Unmarshal([]byte(ValidEventJSON), &e)
		assert.NoError(t, err)
	})

	t.Run("unmarshal ValidEvent2JSON to Event{}", func(t *testing.T) {
		var e event.Event
		err := json.Unmarshal([]byte(ValidEvent2JSON), &e)
		assert.NoError(t, err)
	})
}

func TestEvent_Validate(t *testing.T) {
	t.Run("valid Event", func(t *testing.T) {
		e := ValidEvent()
		err := e.Validate()
		assert.NoError(t, err)
	})

	t.Run("invalid Event.CreatedAt", func(t *testing.T) {
		e := InvalidEventCreatedAt()
		err := e.Validate()
		assert.Error(t, err)
	})

	t.Run("invalid Event.Sig", func(t *testing.T) {
		e := InvalidEventSignature()
		err := e.Validate()
		assert.Error(t, err)
		assert.ErrorContains(t, err, "signature not valid")
	})
}
