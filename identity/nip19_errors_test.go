package identity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNpubToHex_Errors(t *testing.T) {
	t.Run("invalid bech32", func(t *testing.T) {
		_, err := NpubToHex("invalid")
		assert.Error(t, err)
	})

	t.Run("wrong hrp", func(t *testing.T) {
		_, err := NpubToHex("nsec180cvv07tjdrrgpa0j7j7tmnyl2yr6yr7l8j4s3evf6u64th6gkwsrf5vwm")
		assert.Error(t, err)
	})
}

func TestHexToNpub_Errors(t *testing.T) {
	t.Run("invalid hex", func(t *testing.T) {
		_, err := HexToNpub("invalid")
		assert.Error(t, err)
	})
}

func TestNsecToHex_Errors(t *testing.T) {
	t.Run("invalid bech32", func(t *testing.T) {
		_, err := NsecToHex("invalid")
		assert.Error(t, err)
	})

	t.Run("wrong hrp", func(t *testing.T) {
		_, err := NsecToHex("npub180cvv07tjdrrgpa0j7j7tmnyl2yr6yr7l8j4s3evf6u64th6gkwsyjh6w6")
		assert.Error(t, err)
	})
}

func TestHexToNsec_Errors(t *testing.T) {
	t.Run("invalid hex", func(t *testing.T) {
		_, err := HexToNsec("invalid")
		assert.Error(t, err)
	})
}

func TestFromHex_Errors(t *testing.T) {
	t.Run("invalid hex", func(t *testing.T) {
		_, err := FromHex("invalid")
		assert.Error(t, err)
	})
}

func TestFromNsec_Errors(t *testing.T) {
	t.Run("invalid nsec", func(t *testing.T) {
		_, err := FromNsec("invalid")
		assert.Error(t, err)
	})
}
