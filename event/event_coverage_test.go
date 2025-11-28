package event

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEventAllParams(t *testing.T) {
	pubkey := "test"
	sig := "sig"
	id := "id"
	ts := int64(123)
	
	ev := NewEvent(1, "content", Tags{{"e", "123"}}, &ts, &id, &pubkey, &sig)
	
	assert.Equal(t, 1, ev.Kind)
	assert.Equal(t, "content", ev.Content)
	assert.Len(t, ev.Tags, 1)
	assert.NotZero(t, ev.CreatedAt)
	assert.NotNil(t, ev.ID)
	assert.NotNil(t, ev.Pubkey)
	assert.NotNil(t, ev.Sig)
}

func TestNewEventNilParams(t *testing.T) {
	ev := NewEvent(1, "content", nil, nil, nil, nil, nil)
	
	assert.Equal(t, 1, ev.Kind)
	assert.Equal(t, "content", ev.Content)
	assert.NotZero(t, ev.CreatedAt)
}

func TestTagsIteration(t *testing.T) {
	tags := Tags{
		{"e", "event1"},
		{"p", "pubkey1"},
		{"e", "event2"},
	}
	
	count := 0
	for _, tag := range tags {
		if len(tag) > 0 {
			count++
		}
	}
	assert.Equal(t, 3, count)
}
