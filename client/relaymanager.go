package client

import (
	"context"
	"errors"
	"math"

	"nhooyr.io/websocket"
)

type RelayManager struct {
	URL        string
	conn       *websocket.Conn
	doneChan   chan struct{}
	noticeChan chan string
}

func NewRelayManager(url string) RelayManager {
	return RelayManager{
		URL:        url,
		conn:       nil,
		doneChan:   make(chan struct{}),
		noticeChan: make(chan string, 100),
	}
}

func (rm *RelayManager) Connect(ctx context.Context) error {
	conn, _, err := websocket.Dial(ctx, rm.URL, nil)
	if err != nil {
		return errors.Join(err, errors.New("websocket connection"))
	}

	conn.SetReadLimit(math.MaxInt64 - 1) // disable read limit

	rm.conn = conn

	return nil
}
