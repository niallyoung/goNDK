package goNDK_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/niallyoung/goNDK"
)

func TestNewEvent(t *testing.T) {
	t.Run("NewEvent() returns an Event", func(t *testing.T) {
		e := goNDK.NewEvent("fake-pubkey", time.Unix(0, 0), 1, goNDK.Tags{}, "fake-content", "fake-sig")
		assert.NotNil(t, e)
	})
}
