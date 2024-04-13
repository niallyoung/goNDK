package event

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
)

// Sign signs an event with a given privateKey.
func (e *Event) Sign(privateKey string, signOpts ...schnorr.SignOption) error {
	s, err := hex.DecodeString(privateKey)
	if err != nil {
		return fmt.Errorf("sign called with invalid private key '%s': %w", privateKey, err)
	}

	if e.Tags == nil {
		e.Tags = make(Tags, 0)
	}

	sk, pk := btcec.PrivKeyFromBytes(s)
	pkBytes := pk.SerializeCompressed()
	*e.PubKey = hex.EncodeToString(pkBytes[1:])

	h := sha256.Sum256(e.Serialize())
	sig, err := schnorr.Sign(sk, h[:], signOpts...)
	if err != nil {
		return err
	}

	*e.ID = hex.EncodeToString(h[:])
	*e.Sig = hex.EncodeToString(sig.Serialize())

	return nil
}

// ValidateSignature checks if the signature is valid for the id
func (e Event) ValidateSignature() (bool, error) {
	// read and check pubkey
	pk, err := hex.DecodeString(*e.PubKey)
	if err != nil {
		return false, fmt.Errorf("event pubkey '%s' is invalid hex: %w", *e.PubKey, err)
	}

	pubkey, err := schnorr.ParsePubKey(pk)
	if err != nil {
		return false, fmt.Errorf("event has invalid pubkey '%s': %w", *e.PubKey, err)
	}

	// read signature
	s, err := hex.DecodeString(*e.Sig)
	if err != nil {
		return false, fmt.Errorf("signature '%s' is invalid hex: %w", *e.Sig, err)
	}
	sig, err := schnorr.ParseSignature(s)
	if err != nil {
		return false, fmt.Errorf("failed to parse signature: %w", err)
	}

	// check signature
	hash := sha256.Sum256(e.Serialize())
	return sig.Verify(hash[:], pubkey), nil
}
