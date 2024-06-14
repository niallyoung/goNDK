package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"

	"github.com/niallyoung/goNDK/client/test"
)

func TestRelayManager_Connect_doneChan(t *testing.T) {
	t.Run("doneChan receive breaks out", func(t *testing.T) {
		_, port := test.FakeRelay(echo)

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
		_, port := test.FakeRelay(echo)

		rm := NewRelayManager("ws://localhost:" + port)
		err := rm.Connect(ctx)
		assert.NoError(t, err)

		// TODO flesh out FakeRelay() for all future tests
		//res, err := rm.Publish(ctx, &event.Event{})
		//assert.NoError(t, err)
		//assert.NotNil(t, res)
	})
}

func echo(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{}
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	defer func() { _ = c.Close() }()

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			break
		}
		err = c.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}
}
