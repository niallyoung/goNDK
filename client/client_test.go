package client_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/niallyoung/goNDK/client"
)

func TestNewClient(t *testing.T) {
	t.Run("constructor, no relay urls", func(t *testing.T) {
		c := client.NewClient()
		assert.NotNil(t, c)
	})

	t.Run("with relay urls", func(t *testing.T) {
		c := client.NewClient("wss://fake-001", "wss://fake-002")
		assert.NotNil(t, c)
		assert.Equal(t, 2, len(c.RelayManager))
	})
}

func TestClient_Validate(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		c := client.NewClient()
		err := c.Validate()
		assert.NoError(t, err)
	})
}
