package event_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
var randomPrivateKey = fmt.Sprintf("%x", rnd.Uint64())

func TestEvent_Sign(t *testing.T) {
	t.Run("sign", func(t *testing.T) {
		e := validEvent()
		err := e.Sign(randomPrivateKey)
		assert.NoError(t, err)

		ok, err := e.ValidateSignature()
		assert.True(t, ok)
		assert.NoError(t, err)
	})
}
