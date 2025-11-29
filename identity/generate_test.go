package identity

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExtendedIdentity_Sign(t *testing.T) {
	// Generate identity
	id, err := Generate()
	require.NoError(t, err)

	// Create a test event ID (64 hex chars)
	eventID := "0000000000000000000000000000000000000000000000000000000000000001"

	// Sign the event ID
	sig, err := id.Sign(eventID)
	require.NoError(t, err)
	assert.Len(t, sig, 128, "Schnorr signature should be 128 hex chars (64 bytes)")

	// Sign again - should be deterministic
	sig2, err := id.Sign(eventID)
	require.NoError(t, err)
	assert.Equal(t, sig, sig2, "Signatures should be deterministic")
}

func TestExtendedIdentity_Sign_InvalidEventID(t *testing.T) {
	id, err := Generate()
	require.NoError(t, err)

	// Invalid hex
	_, err = id.Sign("not-hex")
	assert.Error(t, err)

	// Wrong length
	_, err = id.Sign("abc123")
	assert.Error(t, err)
}
