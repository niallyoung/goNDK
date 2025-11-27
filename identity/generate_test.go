package identity

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// RED: Test identity generation
func TestGenerate(t *testing.T) {
	id, err := Generate()
	require.NoError(t, err)
	assert.NotNil(t, id)
	assert.Len(t, id.PrivKeyHex, 64)
	assert.Len(t, id.PubKeyHex, 64)
	assert.NotEmpty(t, id.NPub)
	assert.NotEmpty(t, id.Nsec)
}

// RED: Test from nsec
func TestFromNsec(t *testing.T) {
	// Generate first
	id1, err := Generate()
	require.NoError(t, err)

	// Load from nsec
	id2, err := FromNsec(id1.Nsec)
	require.NoError(t, err)

	assert.Equal(t, id1.PrivKeyHex, id2.PrivKeyHex)
	assert.Equal(t, id1.PubKeyHex, id2.PubKeyHex)
	assert.Equal(t, id1.NPub, id2.NPub)
	assert.Equal(t, id1.Nsec, id2.Nsec)
}

// RED: Test from hex
func TestFromHex(t *testing.T) {
	privKeyHex := "0000000000000000000000000000000000000000000000000000000000000001"

	id, err := FromHex(privKeyHex)
	require.NoError(t, err)
	assert.Equal(t, privKeyHex, id.PrivKeyHex)
	assert.NotEmpty(t, id.PubKeyHex)
	assert.NotEmpty(t, id.NPub)
	assert.NotEmpty(t, id.Nsec)
}
