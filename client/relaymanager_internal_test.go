package client

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRelayManager_Connect_doneChan(t *testing.T) {
	t.Run("doneChan receive breaks out", func(t *testing.T) {
		_, port := FakeRelay(Echo)

		rm := NewRelayManager("ws://localhost:" + port)
		err := rm.Connect(context.Background())
		assert.NoError(t, err)

		rm.doneChan <- struct{}{}
		err = rm.Connect(context.Background())
		assert.NoError(t, err)
	})
}

func TestRelayManager_Publish(t *testing.T) {
	t.Run("publish", func(t *testing.T) {
		ctx := context.Background()
		_, port := FakeRelay(Echo)

		rm := NewRelayManager("ws://localhost:" + port)
		err := rm.Connect(ctx)
		assert.NoError(t, err)

		// TODO flesh out FakeRelay() for all future tests
		//res, err := rm.Publish(ctx, &event.Event{})
		//assert.NoError(t, err)
		//assert.NotNil(t, res)
	})
}
