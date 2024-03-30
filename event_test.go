package goNDK

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/niallyoung/goNDK"
)

func TestNewEvent(t *testing.T) {
	t.Run("NewEvent() returns an Event", func(t *testing.T) {
		e := goNDK.NewEvent("", 0, 1, nil, "fake-content", "fake-sig")
		assert.NotNil(t, e)
	})
}
