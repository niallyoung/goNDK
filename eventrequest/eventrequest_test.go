package eventrequest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventRequest(t *testing.T) {
	t.Run("construct", func(t *testing.T) {
		req := EventRequest{
			ID:        "id",
			PubKey:    "pubkey",
			CreatedAt: Now(),
			Kind:      1,
			Tags:      nil,
			Content:   "content",
			Sig:       "sig",
		}

		assert.NotNil(t, req)
	})
}
