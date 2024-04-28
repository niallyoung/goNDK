package client_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/niallyoung/goNDK/client"
)

func TestNewClient(t *testing.T) {
	t.Run("constructor, no relay urls", func(t *testing.T) {
		client := client.NewClient()
		assert.NotNil(t, client)
	})

	t.Run("with relay urls", func(t *testing.T) {
		client := client.NewClient("wss://fake-001", "wss://fake-002")
		assert.NotNil(t, client)
		assert.Equal(t, 2, len(client.RelayManager))
	})
}

func TestClient_Validate(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		client := client.NewClient()
		err := client.Validate()
		assert.NoError(t, err)
	})
}
