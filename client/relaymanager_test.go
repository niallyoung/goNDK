package client_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/niallyoung/goNDK/client"
)

func TestNewRelayManager(t *testing.T) {
	t.Run("constructor", func(t *testing.T) {
		rm := client.NewRelayManager("wss://fake-01")
		assert.NotNil(t, rm)
	})
}

func TestRelayManager_Connect(t *testing.T) {
	t.Run("connect to invalid url", func(t *testing.T) {
		rm := client.NewRelayManager("wss://fake-01")
		err := rm.Connect(context.Background())
		assert.ErrorContains(t, err, "failed to WebSocket dial: failed to send handshake request: Get \"https://fake-01\": dial tcp: lookup fake-01: no such host\nwebsocket connection")
	})
}
