package event

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"github.com/aws/smithy-go/ptr"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
)

// Sign signs an event with the given privateKey
func (e *Event) Sign(privateKey string, signOpts ...schnorr.SignOption) error {
	s, err := hex.DecodeString(privateKey)
	if err != nil {
		return errors.Join(err, errors.New("cannot sign with invalid private key"))
	}

	if e.Tags == nil {
		e.Tags = make(Tags, 0)
	}

	sk, pk := btcec.PrivKeyFromBytes(s)
	pkBytes := pk.SerializeCompressed()
	e.Pubkey = ptr.String(hex.EncodeToString(pkBytes[1:]))

	h := sha256.Sum256(e.Serialize())
	sig, err := schnorr.Sign(sk, h[:], signOpts...)
	if err != nil {
		return err
	}

	e.ID = ptr.String(hex.EncodeToString(h[:]))
	e.Sig = ptr.String(hex.EncodeToString(sig.Serialize()))

	return nil
}

// ValidateSignature checks if the signature is valid for the id.
func (e *Event) ValidateSignature() (bool, error) {
	if e.Pubkey == nil || e.Sig == nil || e.ID == nil {
		return false, errors.New("unsigned event")
	}

	// decode and parse pubkey
	pk, err := hex.DecodeString(*e.Pubkey)
	if err != nil {
		return false, errors.Join(err, errors.New("pubkey hex error"))
	}

	pubkey, err := schnorr.ParsePubKey(pk)
	if err != nil {
		return false, errors.Join(err, errors.New("invalid pubkey"))
	}

	// decode and parse signature
	s, err := hex.DecodeString(*e.Sig)
	if err != nil {
		return false, errors.Join(err, errors.New("signature hex error"))
	}

	sig, err := schnorr.ParseSignature(s)
	if err != nil {
		return false, errors.Join(err, errors.New("failed to parse signature"))
	}

	// check signature
	hash := sha256.Sum256(e.Serialize())
	return sig.Verify(hash[:], pubkey), nil
}
