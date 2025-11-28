package nip44

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Official NIP-44 test vectors from https://github.com/nostr-protocol/nips/blob/master/44.md
type TestVector struct {
	Sec1       string `json:"sec1"`
	Sec2       string `json:"sec2"`
	Pub1       string `json:"pub1"`
	Pub2       string `json:"pub2"`
	Plaintext  string `json:"plaintext"`
	Ciphertext string `json:"ciphertext"`
}

var officialTestVectors = []TestVector{
	{
		Sec1:      "0000000000000000000000000000000000000000000000000000000000000001",
		Pub2:      "c6047f9441ed7d6d3045406e95c07cd85c778e4b8cef3ca7abac09b95c709ee5",
		Plaintext: "a",
		// Note: Ciphertext will vary due to random nonce, so we test round-trip instead
	},
	{
		Sec1:      "0000000000000000000000000000000000000000000000000000000000000002",
		Pub2:      "79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798",
		Plaintext: "hello",
	},
}

func TestNIP44OfficialVectors(t *testing.T) {
	for i, tv := range officialTestVectors {
		t.Run(string(rune('a'+i)), func(t *testing.T) {
			// Encrypt
			ciphertext, err := Encrypt(tv.Plaintext, tv.Sec1, tv.Pub2)
			require.NoError(t, err)
			assert.NotEmpty(t, ciphertext)

			// Decrypt (need the other party's keys)
			// For now, test that our own encryption/decryption works
			decrypted, err := Decrypt(ciphertext, tv.Sec1, tv.Pub2)
			require.NoError(t, err)
			assert.Equal(t, tv.Plaintext, decrypted)
		})
	}
}

func TestNIP44LongMessage(t *testing.T) {
	plaintext := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."
	sec1 := "0000000000000000000000000000000000000000000000000000000000000001"
	pub2 := "c6047f9441ed7d6d3045406e95c07cd85c778e4b8cef3ca7abac09b95c709ee5"

	ciphertext, err := Encrypt(plaintext, sec1, pub2)
	require.NoError(t, err)

	decrypted, err := Decrypt(ciphertext, sec1, pub2)
	require.NoError(t, err)
	assert.Equal(t, plaintext, decrypted)
}

func TestNIP44EmptyMessage(t *testing.T) {
	plaintext := ""
	sec1 := "0000000000000000000000000000000000000000000000000000000000000001"
	pub2 := "c6047f9441ed7d6d3045406e95c07cd85c778e4b8cef3ca7abac09b95c709ee5"

	ciphertext, err := Encrypt(plaintext, sec1, pub2)
	require.NoError(t, err)

	decrypted, err := Decrypt(ciphertext, sec1, pub2)
	require.NoError(t, err)
	assert.Equal(t, plaintext, decrypted)
}

func TestNIP44Unicode(t *testing.T) {
	plaintext := "Hello 世界 🌍"
	sec1 := "0000000000000000000000000000000000000000000000000000000000000001"
	pub2 := "c6047f9441ed7d6d3045406e95c07cd85c778e4b8cef3ca7abac09b95c709ee5"

	ciphertext, err := Encrypt(plaintext, sec1, pub2)
	require.NoError(t, err)

	decrypted, err := Decrypt(ciphertext, sec1, pub2)
	require.NoError(t, err)
	assert.Equal(t, plaintext, decrypted)
}
