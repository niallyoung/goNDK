package event_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvent_Serialize(t *testing.T) {
	t.Run("given a valid Event, Serialize() returns the expected JSON", func(t *testing.T) {
		e := ValidEvent()
		bytes := e.Serialize()
		assert.Equal(t, ValidEventSerialize, string(bytes))
	})
}

func TestEvent_String(t *testing.T) {
	t.Run("given a valid Event, String() returns the expected string", func(t *testing.T) {
		e := ValidEvent()
		assert.Equal(t, ValidEventString, e.String())
	})
}
