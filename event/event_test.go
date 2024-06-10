package event_test

import (
	"encoding/json"
	"strings"
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
	t.Run("marshal zero-object Event{} literal, matches String()", func(t *testing.T) {
		var e event.Event
		jsonBytes, err := json.Marshal(e)
		e2 := event.Event{}
		expectedString := e2.String()
		assert.NoError(t, err)
		assert.Equal(t, expectedString, string(jsonBytes))
	})

	t.Run("marshal NewEvent(), matches String()", func(t *testing.T) {
		e := event.NewEvent(0, "", event.Tags{}, nil, nil, nil, nil)
		jsonBytes, err := json.Marshal(e)
		e2 := event.NewEvent(0, "", event.Tags{}, nil, nil, nil, nil)
		expectedString := e2.String()
		assert.NoError(t, err)
		assert.Equal(t, expectedString, string(jsonBytes))
	})
}

func TestEvent_UnmarshalJSON(t *testing.T) {
	t.Run("unmarshal ValidEventJSON, matches ValidEvent(), metadata matches", func(t *testing.T) {
		var e event.Event
		err := json.Unmarshal([]byte(ValidEventJSON), &e)
		assert.NoError(t, err)
		assert.Equal(t, *ValidEvent(), e)
		assert.True(t, strings.Contains(e.Content, "GM nostr"))
	})

	t.Run("unmarshal ValidEvent2JSON, metadata matches", func(t *testing.T) {
		var e event.Event
		err := json.Unmarshal([]byte(ValidEvent2JSON), &e)
		assert.NoError(t, err)
		assert.Equal(t, event.Timestamp(1673311423), e.CreatedAt)
		assert.True(t, strings.Contains(e.Content, "thousands of smaller ultilities coming together"))
		assert.True(t, strings.Contains(e.Content, "will be the magic"))
	})

	t.Run("unmarshal ValidEvent3JSON, metadata matches", func(t *testing.T) {
		var e event.Event
		err := json.Unmarshal([]byte(ValidEvent3JSON), &e)
		assert.NoError(t, err)
		assert.Equal(t, event.Timestamp(1717137637), e.CreatedAt)
		assert.True(t, strings.Contains(e.Content, "ecash improves these users' privacy"))
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
