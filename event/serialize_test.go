package event_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/niallyoung/goNDK/event"
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
		assert.Equal(t, ValidEventJSON, e.String())
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
