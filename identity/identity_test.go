package identity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/niallyoung/goNDK/identity"
)

func TestIdentity_NewIdentity(t *testing.T) {
	t.Run("constructor", func(t *testing.T) {
		i := identity.NewIdentity("pubkey", "npub")
		assert.NotNil(t, i)
	})
}

func TestIdentity_Validate(t *testing.T) {
	t.Run("validate", func(t *testing.T) {
		i := identity.NewIdentity(
			"234dd2c21135830a960a462defdb410ac26f178cbf8e13fbe43890f8656ee983",
			"npub",
		)
		err := i.Validate()
		assert.NoError(t, err)
	})
}
