package client

import (
	"context"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/websocket"

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

func FakeRelay(f http.HandlerFunc) (*httptest.Server, string) {
	// custom listener to start up on a random port
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal(err)
	}

	s := httptest.NewUnstartedServer(f)
	_ = s.Listener.Close()
	s.Listener = l
	s.Start()

	port := l.Addr().(*net.TCPAddr).Port
	portString := strconv.FormatInt(int64(port), 10)

	return s, portString
}

func Echo(w http.ResponseWriter, r *http.Request) {
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
