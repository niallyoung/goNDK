package client_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/niallyoung/goNDK/client"
	"github.com/niallyoung/goNDK/event"
)

func TestEventMessage_MarshalJSON_ErrorPaths(t *testing.T) {
	t.Run("with subscription ID", func(t *testing.T) {
		e := event.NewEvent(1, "test", nil, nil, nil, nil, nil)
		msg := &client.EventMessage{
			SubscriptionID: "sub123",
			Event:          e,
		}
		data, err := json.Marshal(msg)
		require.NoError(t, err)
		assert.Contains(t, string(data), "sub123")
	})

	t.Run("without subscription ID", func(t *testing.T) {
		e := event.NewEvent(1, "test", nil, nil, nil, nil, nil)
		msg := &client.EventMessage{
			Event: e,
		}
		data, err := json.Marshal(msg)
		require.NoError(t, err)
		assert.NotContains(t, string(data), "sub")
	})
}

func TestReqMessage_MarshalJSON_Coverage(t *testing.T) {
	t.Run("with single filter", func(t *testing.T) {
		msg := &client.ReqMessage{
			SubscriptionID: "sub1",
			Filters:        []client.Filter{{Kinds: []int{1}}},
		}
		data, err := json.Marshal(msg)
		require.NoError(t, err)
		assert.Contains(t, string(data), "REQ")
		assert.Contains(t, string(data), "sub1")
	})

	t.Run("with multiple filters", func(t *testing.T) {
		msg := &client.ReqMessage{
			SubscriptionID: "sub2",
			Filters: []client.Filter{
				{Kinds: []int{1}},
				{Kinds: []int{2}},
			},
		}
		data, err := json.Marshal(msg)
		require.NoError(t, err)
		assert.Contains(t, string(data), "REQ")
	})

	t.Run("with empty filters", func(t *testing.T) {
		msg := &client.ReqMessage{
			SubscriptionID: "sub3",
			Filters:        []client.Filter{},
		}
		data, err := json.Marshal(msg)
		require.NoError(t, err)
		assert.Contains(t, string(data), "REQ")
	})
}

func TestCloseMessage_MarshalJSON_Coverage(t *testing.T) {
	t.Run("marshal close message", func(t *testing.T) {
		msg := &client.CloseMessage{
			SubscriptionID: "sub123",
		}
		data, err := json.Marshal(msg)
		require.NoError(t, err)
		assert.Contains(t, string(data), "CLOSE")
		assert.Contains(t, string(data), "sub123")
	})
}
