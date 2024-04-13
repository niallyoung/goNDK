package event_test

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/niallyoung/goNDK/event"
)

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
var randomPrivateKey = fmt.Sprintf("%x", rnd.Uint64())

func TestEvent_Sign(t *testing.T) {
	t.Run("sign valid Event with an invalid privatekey", func(t *testing.T) {
		e := validEvent()
		err := e.Sign("invalid-private-key")
		assert.Error(t, err)
	})

	t.Run("sign valid Event with a valid privatekey", func(t *testing.T) {
		e := validEvent()
		err := e.Sign(randomPrivateKey)
		assert.NoError(t, err)
	})

	t.Run("sign valid Event with no tags, with a valid privatekey", func(t *testing.T) {
		e := validEventNoTags()
		err := e.Sign(randomPrivateKey)
		assert.NoError(t, err)
	})
}

func TestEvent_ValidateSignature_JSON_Unmarshal(t *testing.T) {
	t.Run("sign json.Unmarshal with a valid privatekey", func(t *testing.T) {
		var e event.Event
		err := json.Unmarshal([]byte(validEventJSON), &e)
		ok, err := e.ValidateSignature()
		assert.True(t, ok)
		assert.NoError(t, err)
	})
}

func TestEvent_ValidateSignature(t *testing.T) {
	t.Run("sign NewEvent() with a valid privatekey", func(t *testing.T) {
		e := validEvent()
		ok, err := e.ValidateSignature()
		assert.True(t, ok) // FIXME failing Step 9 of schnorrVerify()
		assert.NoError(t, err)
	})

	t.Run("sign an invalid pubkey, with a valid privatekey", func(t *testing.T) {
		e := validEvent()
		e.PubKey = "invalid"
		ok, err := e.ValidateSignature()
		assert.False(t, ok)
		assert.Error(t, err)
	})

	t.Run("sign an invalid pubkey, with a valid privatekey", func(t *testing.T) {
		e := validEvent()
		e.PubKey = "invalid"
		ok, err := e.ValidateSignature()
		assert.False(t, ok)
		assert.Error(t, err)
	})
}