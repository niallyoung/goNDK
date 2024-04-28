package identity

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Identity struct {
	//Profile
	Pubkey string
	NPub   string // convenience, derived from Pubkey
}

func NewIdentity(pubkey string, npub string) Identity {
	return Identity{
		Pubkey: pubkey,
		NPub:   npub,
	}
}

func (i Identity) Validate() error {
	return validation.ValidateStruct(&i,
		validation.Field(&i.Pubkey, // hex, secp256k1 schnorr public key
			validation.Required, is.Hexadecimal, validation.Length(64, 64)),
		validation.Field(&i.NPub,
			validation.Required),
	)
}
