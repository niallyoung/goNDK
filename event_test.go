package gondo_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/niallyoung/gondo"
)

func TestNewEvent(t *testing.T) {
	t.Run("NewEvent() returns an Event", func(t *testing.T) {
		e := gondo.NewEvent("", 0, 1, nil, "fake-content", "fake-sig")
		assert.NotNil(t, e)
	})
}
