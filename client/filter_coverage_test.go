package client

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFilterMarshalJSON(t *testing.T) {
	tests := []struct {
		name   string
		filter Filter
	}{
		{
			name: "empty filter",
			filter: Filter{},
		},
		{
			name: "with IDs",
			filter: Filter{
				IDs: []string{"abc123", "def456"},
			},
		},
		{
			name: "with authors",
			filter: Filter{
				Authors: []string{"pubkey1", "pubkey2"},
			},
		},
		{
			name: "with kinds",
			filter: Filter{
				Kinds: []int{1, 3, 7},
			},
		},
		{
			name: "with limit",
			filter: Filter{
				Limit: 10,
			},
		},
		{
			name: "complete filter",
			filter: Filter{
				IDs:     []string{"id1"},
				Authors: []string{"author1"},
				Kinds:   []int{1},
				Limit:   5,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.filter)
			require.NoError(t, err)
			assert.NotEmpty(t, data)

			var decoded Filter
			err = json.Unmarshal(data, &decoded)
			require.NoError(t, err)
		})
	}
}
