package event_test

import (
	"encoding/json"
	"testing"

	secp "github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/stretchr/testify/assert"

	"github.com/niallyoung/goNDK/event"
)

var randomPrivateKey, _ = secp.GeneratePrivateKey()

func TestEvent_Sign(t *testing.T) {
	t.Run("given a valid Event", func(t *testing.T) {
		e := ValidEvent()
		t.Run("when we sign with an invalid privatekey", func(t *testing.T) {
			err := e.Sign("invalid-private-key")
			t.Run("then we get an error, and the Event is unchanged", func(t *testing.T) {
				assert.Error(t, err)
				assert.Equal(t, e, ValidEvent())
			})
		})
	})

	t.Run("given a valid Event with a valid privatekey", func(t *testing.T) {
		e := ValidEvent()
		t.Run("when we sign with a valid privatekey", func(t *testing.T) {
			err := e.Sign(randomPrivateKey.Key.String())
			t.Run("then we get no error, and the Event is unchanged", func(t *testing.T) {
				assert.NoError(t, err)
				assert.NotEqual(t, e, ValidEvent())
			})
		})
	})

	t.Run("given a valid Event with no tags", func(t *testing.T) {
		e := ValidEventNoTags()
		t.Run("when we sign with a valid privatekey", func(t *testing.T) {
			err := e.Sign(randomPrivateKey.Key.String())
			t.Run("then we get no error, and the Event has changed", func(t *testing.T) {
				assert.NoError(t, err)
				assert.NotEqual(t, e, ValidEvent())
			})
		})
	})
}

func TestEvent_ValidateSignature_JSON_Unmarshal(t *testing.T) {
	t.Run("given valid JSON, when we Unmarshal into an Event", func(t *testing.T) {
		var e event.Event
		err := json.Unmarshal([]byte(ValidEventJSON), &e)
		assert.NoError(t, err)
		t.Run("when we validate the signature", func(t *testing.T) {
			ok, err := e.ValidateSignature()
			t.Run("then we get no error, and successful validation", func(t *testing.T) {
				assert.NoError(t, err)
				assert.True(t, ok)
			})
		})
	})
}

func TestEvent_ValidateSignature_NewEvent(t *testing.T) {
	t.Run("given a valid Event", func(t *testing.T) {
		e := ValidEvent()
		t.Run("when we validate its signature", func(t *testing.T) {
			ok, err := e.ValidateSignature()
			t.Run("then we get no error, and successful validation", func(t *testing.T) {
				assert.NoError(t, err)
				assert.True(t, ok)
			})
		})
	})

	t.Run("given an valid Event with minimal fields (missing PubKey|Sig|ID)", func(t *testing.T) {
		e := ValidEventMinimal()
		t.Run("when we validate its signature", func(t *testing.T) {
			ok, err := e.ValidateSignature()
			t.Run("then we get an error, and unsuccessful validation", func(t *testing.T) {
				assert.False(t, ok)
				assert.Error(t, err)
				assert.ErrorContains(t, err, "unsigned event")
			})
		})
	})

	t.Run("given an invalid Event (bad PubKey)", func(t *testing.T) {
		e := ValidEvent()
		*e.PubKey = "invalid"
		t.Run("when we validate its signature", func(t *testing.T) {
			ok, err := e.ValidateSignature()
			t.Run("then we get an error, and unsuccessful validation", func(t *testing.T) {
				assert.Error(t, err)
				assert.False(t, ok)
			})
		})
	})

	t.Run("given an invalid Event (bad PubKey)", func(t *testing.T) {
		e := ValidEvent()
		*e.PubKey = "invalid"
		t.Run("when we validate its signature", func(t *testing.T) {
			ok, err := e.ValidateSignature()
			t.Run("then we get an error, and unsuccessful validation", func(t *testing.T) {
				assert.Error(t, err)
				assert.False(t, ok)
			})
		})
	})
}
