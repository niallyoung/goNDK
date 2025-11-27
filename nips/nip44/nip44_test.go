package nip44

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// RED: Test NIP-44 encryption
func TestEncrypt(t *testing.T) {
	plaintext := "Hello NOSTR"
	senderPrivKey := "0000000000000000000000000000000000000000000000000000000000000001"
	recipientPubKey := "79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798"

	ciphertext, err := Encrypt(plaintext, senderPrivKey, recipientPubKey)
	require.NoError(t, err)
	assert.NotEmpty(t, ciphertext)
	assert.NotEqual(t, plaintext, ciphertext)
}

// RED: Test NIP-44 decryption
func TestDecrypt(t *testing.T) {
	plaintext := "Hello NOSTR"
	senderPrivKey := "0000000000000000000000000000000000000000000000000000000000000001"
	recipientPrivKey := "0000000000000000000000000000000000000000000000000000000000000002"
	senderPubKey := "79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798"
	recipientPubKey := "c6047f9441ed7d6d3045406e95c07cd85c778e4b8cef3ca7abac09b95c709ee5"

	ciphertext, err := Encrypt(plaintext, senderPrivKey, recipientPubKey)
	require.NoError(t, err)

	decrypted, err := Decrypt(ciphertext, recipientPrivKey, senderPubKey)
	require.NoError(t, err)
	assert.Equal(t, plaintext, decrypted)
}

// RED: Test round-trip
func TestRoundTrip(t *testing.T) {
	plaintext := "The quick brown fox jumps over the lazy dog"
	
	// Generate keys
	senderPriv, _ := hex.DecodeString("0000000000000000000000000000000000000000000000000000000000000001")
	recipientPriv, _ := hex.DecodeString("0000000000000000000000000000000000000000000000000000000000000002")
	
	senderPrivHex := hex.EncodeToString(senderPriv)
	recipientPrivHex := hex.EncodeToString(recipientPriv)
	
	// Derive public keys
	senderPubHex := "79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798"
	recipientPubHex := "c6047f9441ed7d6d3045406e95c07cd85c778e4b8cef3ca7abac09b95c709ee5"

	// Encrypt
	ciphertext, err := Encrypt(plaintext, senderPrivHex, recipientPubHex)
	require.NoError(t, err)

	// Decrypt
	decrypted, err := Decrypt(ciphertext, recipientPrivHex, senderPubHex)
	require.NoError(t, err)
	
	assert.Equal(t, plaintext, decrypted)
}
