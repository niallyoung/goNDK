package client_test

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

	"github.com/niallyoung/goNDK/client"
	"github.com/niallyoung/goNDK/event"
)

func TestNewRelayManager(t *testing.T) {
	t.Run("constructor", func(t *testing.T) {
		rm := client.NewRelayManager("wss://localhost")
		assert.NotNil(t, rm)
	})
}

func TestRelayManager_Connect(t *testing.T) {
	t.Run("connect to invalid url", func(t *testing.T) {
		rm := client.NewRelayManager("wss://localhost:0")
		err := rm.Connect(context.Background())
		assert.ErrorContains(t, err, "failed to WebSocket dial: failed to send handshake request")
	})

	t.Run("connect to valid url", func(t *testing.T) {
		_, port := server()
		rm := client.NewRelayManager("ws://" + "localhost:" + strconv.FormatInt(int64(port), 10))
		err := rm.Connect(context.Background())
		assert.NoError(t, err)
	})

	t.Run("ReadMessage() error during Connect", func(t *testing.T) {
		_, port := server()
		rm := client.NewRelayManager("ws://" + "localhost:" + strconv.FormatInt(int64(port), 10))
		err := rm.Connect(context.Background())
		assert.NoError(t, err)
		err = rm.WriteMessage(context.Background(), event.Event{})
		err = rm.Connect(context.Background())
		assert.NoError(t, err)
	})
}

func server() (*httptest.Server, int) {
	// custom listener to start up on a random port
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal(err)
	}

	s := httptest.NewUnstartedServer(http.HandlerFunc(echo))
	_ = s.Listener.Close()
	s.Listener = l
	s.Start()

	return s, l.Addr().(*net.TCPAddr).Port
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
