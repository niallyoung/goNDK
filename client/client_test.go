package client_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/niallyoung/goNDK/client"
)

func TestNewClient(t *testing.T) {
	t.Run("constructor", func(t *testing.T) {
		client := client.NewClient()
		assert.NotNil(t, client)
	})
}

func TestClient_Validate(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		client := client.NewClient()
		err := client.Validate()
		assert.NoError(t, err)
	})
}
