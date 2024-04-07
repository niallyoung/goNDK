package event_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/niallyoung/goNDK/event"
)

//var validEventJSON = `{
//	"id": "b52cc46fc9e38e51e8774cc13c00523c013d371d1dd5f42113f06e43ed870a76",
//	"pubkey": "234dd2c21135830a960a462defdb410ac26f178cbf8e13fbe43890f8656ee983",
//	"created_at": 1712350548,
//	"kind": 1,
//	"tags": [],
//	"content": "GM nostr welcome to Saturday!",
//	"sig": "46d7935c4f26f7c20da1f5cdd919f397dc1f63339fadf0b8145eb1fa6a92fae05ef12b5faa8b45794c2700c268ffe0fc389e1894b5fd09195a65e72df7d9e7c1"
//}`

var validEvent = func() (event.Eventer, error) {
	return event.NewEvent(
		"b52cc46fc9e38e51e8774cc13c00523c013d371d1dd5f42113f06e43ed870a76",
		"234dd2c21135830a960a462defdb410ac26f178cbf8e13fbe43890f8656ee983",
		int64(1712350548),
		1,
		event.Tags{nil},
		"GM nostr welcome to Saturday!",
		"46d7935c4f26f7c20da1f5cdd919f397dc1f63339fadf0b8145eb1fa6a92fae05ef12b5faa8b45794c2700c268ffe0fc389e1894b5fd09195a65e72df7d9e7c1",
	)
}

func TestNewEvent(t *testing.T) {
	t.Run("NewEvent() returns an Event", func(t *testing.T) {
		e, err := validEvent()
		assert.NoError(t, err)
		assert.NotNil(t, e)
	})
}

func TestEvent_Validate(t *testing.T) {
	t.Run("validate", func(t *testing.T) {
		e, err := validEvent()
		assert.NoError(t, err)
		err = e.Validate()
		assert.NoError(t, err)
	})
}
