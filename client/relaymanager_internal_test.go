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
		_, port := server()
		rm := NewRelayManager("ws://" + "localhost:" + strconv.FormatInt(int64(port), 10))
		err := rm.Connect(context.Background())
		assert.NoError(t, err)
		rm.doneChan <- struct{}{}
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
