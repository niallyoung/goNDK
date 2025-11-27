package identity

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// RED: Test npub to hex conversion
func TestNpubToHex(t *testing.T) {
	npub := "npub180cvv07tjdrrgpa0j7j7tmnyl2yr6yr7l8j4s3evf6u64th6gkwsyjh6w6"
	expected := "3bf0c63fcb93463407af97a5e5ee64fa883d107ef9e558472c4eb9aaaefa459d"

	hex, err := NpubToHex(npub)
	require.NoError(t, err)
	assert.Equal(t, expected, hex)
}

// RED: Test hex to npub conversion
func TestHexToNpub(t *testing.T) {
	hex := "3bf0c63fcb93463407af97a5e5ee64fa883d107ef9e558472c4eb9aaaefa459d"
	expected := "npub180cvv07tjdrrgpa0j7j7tmnyl2yr6yr7l8j4s3evf6u64th6gkwsyjh6w6"

	npub, err := HexToNpub(hex)
	require.NoError(t, err)
	assert.Equal(t, expected, npub)
}

// RED: Test nsec round-trip
func TestNsecRoundTrip(t *testing.T) {
	originalHex := "0000000000000000000000000000000000000000000000000000000000000001"

	nsec, err := HexToNsec(originalHex)
	require.NoError(t, err)

	hex, err := NsecToHex(nsec)
	require.NoError(t, err)

	assert.Equal(t, originalHex, hex)
}

// RED: Test round-trip
func TestNip19RoundTrip(t *testing.T) {
	originalHex := "3bf0c63fcb93463407af97a5e5ee64fa883d107ef9e558472c4eb9aaaefa459d"

	npub, err := HexToNpub(originalHex)
	require.NoError(t, err)

	hex, err := NpubToHex(npub)
	require.NoError(t, err)

	assert.Equal(t, originalHex, hex)
}
