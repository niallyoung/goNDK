package client_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/niallyoung/goNDK/client"
)

func TestFilter_MarshalJSON(t *testing.T) {
	t.Run("empty filter", func(t *testing.T) {
		f := client.Filter{}
		data, err := json.Marshal(f)
		require.NoError(t, err)
		assert.Equal(t, "{}", string(data))
	})

	t.Run("filter with kinds", func(t *testing.T) {
		f := client.Filter{Kinds: []int{1, 2, 3}}
		data, err := json.Marshal(f)
		require.NoError(t, err)
		assert.Contains(t, string(data), `"kinds":[1,2,3]`)
	})

	t.Run("filter with authors", func(t *testing.T) {
		f := client.Filter{Authors: []string{"abc", "def"}}
		data, err := json.Marshal(f)
		require.NoError(t, err)
		assert.Contains(t, string(data), `"authors":["abc","def"]`)
	})

	t.Run("filter with multiple fields", func(t *testing.T) {
		f := client.Filter{
			Kinds:   []int{1},
			Authors: []string{"test"},
			Limit:   10,
		}
		data, err := json.Marshal(f)
		require.NoError(t, err)
		assert.Contains(t, string(data), `"kinds":[1]`)
		assert.Contains(t, string(data), `"authors":["test"]`)
		assert.Contains(t, string(data), `"limit":10`)
	})
}
