package client_test

import (
	"context"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"

	"github.com/niallyoung/goNDK/client"
)

const serverAddress = "localhost:8080"

func TestNewRelayManager(t *testing.T) {
	t.Run("constructor", func(t *testing.T) {
		rm := client.NewRelayManager("wss://" + serverAddress)
		assert.NotNil(t, rm)
	})
}

func TestRelayManager_Connect(t *testing.T) {
	t.Run("connect to invalid url", func(t *testing.T) {
		rm := client.NewRelayManager("wss://" + serverAddress)
		err := rm.Connect(context.Background())
		assert.ErrorContains(t, err, "failed to WebSocket dial: failed to send handshake request")
	})

	t.Run("connect to valid url (local test server)", func(t *testing.T) {
		_ = server()
		rm := client.NewRelayManager("ws://" + serverAddress)
		err := rm.Connect(context.Background())
		assert.NoError(t, err)
	})
}

func server() *httptest.Server {
	// custom listener to start up with a specific port
	l, err := net.Listen("tcp", serverAddress)
	if err != nil {
		log.Fatal(err)
	}

	s := httptest.NewUnstartedServer(http.HandlerFunc(echo))
	s.Listener.Close()
	s.Listener = l
	s.Start()

	return s
}

func echo(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{}
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	defer c.Close()

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
