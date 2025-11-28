package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/niallyoung/goNDK/event"
)

func TestEventMessageMarshalWithoutSubID(t *testing.T) {
	ev := event.NewEvent(1, "test", nil, nil, nil, nil, nil)
	msg := &EventMessage{Event: ev}
	
	data, err := msg.MarshalJSON()
	require.NoError(t, err)
	assert.Contains(t, string(data), "EVENT")
}

func TestEventMessageMarshalWithSubID(t *testing.T) {
	ev := event.NewEvent(1, "test", nil, nil, nil, nil, nil)
	msg := &EventMessage{
		SubscriptionID: "sub123",
		Event:          ev,
	}
	
	data, err := msg.MarshalJSON()
	require.NoError(t, err)
	assert.Contains(t, string(data), "EVENT")
	assert.Contains(t, string(data), "sub123")
}

func TestReqMessageMarshal(t *testing.T) {
	msg := &ReqMessage{
		SubscriptionID: "sub123",
		Filters: []Filter{
			{Kinds: []int{1}},
			{Authors: []string{"pub1"}},
		},
	}
	
	data, err := msg.MarshalJSON()
	require.NoError(t, err)
	assert.Contains(t, string(data), "REQ")
	assert.Contains(t, string(data), "sub123")
}

func TestCloseMessageMarshal(t *testing.T) {
	msg := &CloseMessage{SubscriptionID: "sub123"}
	
	data, err := msg.MarshalJSON()
	require.NoError(t, err)
	assert.Contains(t, string(data), "CLOSE")
	assert.Contains(t, string(data), "sub123")
}
