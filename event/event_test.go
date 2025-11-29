package event_test

import (
	"testing"

	"github.com/niallyoung/goNDK/event"
	"github.com/stretchr/testify/assert"
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
		err := e.ValidateComplete()
		assert.Error(t, err)
		assert.ErrorContains(t, err, "signature not valid")
	})

	t.Run("valid Event without signature", func(t *testing.T) {
		e := ValidEventMinimal()
		err := e.Validate()
		assert.NoError(t, err)
	})

	t.Run("empty content is valid", func(t *testing.T) {
		e := ValidEventMinimal()
		e.Content = ""
		err := e.Validate()
		assert.NoError(t, err)
	})

	t.Run("tags with special characters are valid", func(t *testing.T) {
		e := ValidEventMinimal()
		e.Tags = event.Tags{
			event.Tag{"p", "npub1abc123..."},
			event.Tag{"e", "event-id-hex"},
			event.Tag{"r", "https://example.com/image.jpg"},
			event.Tag{"t", "#bitcoin"},
		}
		err := e.Validate()
		assert.NoError(t, err)
	})

	t.Run("nil tags are handled gracefully", func(t *testing.T) {
		e := ValidEventMinimal()
		e.Tags = event.Tags{nil, event.Tag{"p", "test"}}
		err := e.Validate()
		assert.NoError(t, err)
	})
}
