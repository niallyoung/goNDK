package event_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/niallyoung/gondo/core/event"
)

func TestEvent(t *testing.T) {
	t.Run("stub", func(t *testing.T) {
		e := event.Event{}
		assert.NotNil(t, e)
	})
}
