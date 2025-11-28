package nip44

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncryptInvalidKeys(t *testing.T) {
	_, err := Encrypt("test", "invalid", "pub")
	assert.Error(t, err)
}

func TestDecryptInvalidKeys(t *testing.T) {
	_, err := Decrypt("test", "invalid", "pub")
	assert.Error(t, err)
}

func TestDecryptInvalidCiphertext(t *testing.T) {
	_, err := Decrypt("not-base64!", "0000000000000000000000000000000000000000000000000000000000000001", "79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798")
	assert.Error(t, err)
}

func TestDecryptWrongKey(t *testing.T) {
	plaintext := "secret"
	ciphertext, err := Encrypt(plaintext, "0000000000000000000000000000000000000000000000000000000000000001", "c6047f9441ed7d6d3045406e95c07cd85c778e4b8cef3ca7abac09b95c709ee5")
	require.NoError(t, err)
	
	// Try to decrypt with wrong key
	_, err = Decrypt(ciphertext, "0000000000000000000000000000000000000000000000000000000000000003", "79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798")
	assert.Error(t, err)
}

func TestEncryptSpecialCharacters(t *testing.T) {
	plaintext := "!@#$%^&*()_+-=[]{}|;':\",./<>?"
	sec1 := "0000000000000000000000000000000000000000000000000000000000000001"
	pub2 := "c6047f9441ed7d6d3045406e95c07cd85c778e4b8cef3ca7abac09b95c709ee5"

	ciphertext, err := Encrypt(plaintext, sec1, pub2)
	require.NoError(t, err)

	decrypted, err := Decrypt(ciphertext, sec1, pub2)
	require.NoError(t, err)
	assert.Equal(t, plaintext, decrypted)
}

func TestEncryptNewlines(t *testing.T) {
	plaintext := "line1\nline2\nline3"
	sec1 := "0000000000000000000000000000000000000000000000000000000000000001"
	pub2 := "c6047f9441ed7d6d3045406e95c07cd85c778e4b8cef3ca7abac09b95c709ee5"

	ciphertext, err := Encrypt(plaintext, sec1, pub2)
	require.NoError(t, err)

	decrypted, err := Decrypt(ciphertext, sec1, pub2)
	require.NoError(t, err)
	assert.Equal(t, plaintext, decrypted)
}

func TestEncryptVeryLongMessage(t *testing.T) {
	plaintext := ""
	for i := 0; i < 1000; i++ {
		plaintext += "Lorem ipsum dolor sit amet. "
	}
	
	sec1 := "0000000000000000000000000000000000000000000000000000000000000001"
	pub2 := "c6047f9441ed7d6d3045406e95c07cd85c778e4b8cef3ca7abac09b95c709ee5"

	ciphertext, err := Encrypt(plaintext, sec1, pub2)
	require.NoError(t, err)

	decrypted, err := Decrypt(ciphertext, sec1, pub2)
	require.NoError(t, err)
	assert.Equal(t, plaintext, decrypted)
}
