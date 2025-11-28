package identity

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2"
)

// ExtendedIdentity includes private key information
type ExtendedIdentity struct {
	Identity
	PrivKeyHex string
	PubKeyHex  string
	Nsec       string
}

// Generate creates a new random identity
func Generate() (*ExtendedIdentity, error) {
	// Generate random private key
	privKeyBytes := make([]byte, 32)
	if _, err := rand.Read(privKeyBytes); err != nil {
		return nil, fmt.Errorf("generate random key: %w", err)
	}

	privKey, pubKey := btcec.PrivKeyFromBytes(privKeyBytes)

	// Get hex representations
	privKeyHex := hex.EncodeToString(privKey.Serialize())
	pubKeyHex := hex.EncodeToString(pubKey.SerializeCompressed()[1:]) // Skip 0x02 prefix

	// Convert to bech32
	npub, err := HexToNpub(pubKeyHex)
	if err != nil {
		return nil, fmt.Errorf("convert to npub: %w", err)
	}

	nsec, err := HexToNsec(privKeyHex)
	if err != nil {
		return nil, fmt.Errorf("convert to nsec: %w", err)
	}

	return &ExtendedIdentity{
		Identity: Identity{
			Pubkey: pubKeyHex,
			NPub:   npub,
		},
		PrivKeyHex: privKeyHex,
		PubKeyHex:  pubKeyHex,
		Nsec:       nsec,
	}, nil
}

// FromNsec loads identity from nsec string
func FromNsec(nsec string) (*ExtendedIdentity, error) {
	privKeyHex, err := NsecToHex(nsec)
	if err != nil {
		return nil, fmt.Errorf("convert nsec: %w", err)
	}

	return FromHex(privKeyHex)
}

// FromHex loads identity from hex private key
func FromHex(privKeyHex string) (*ExtendedIdentity, error) {
	privKeyBytes, err := hex.DecodeString(privKeyHex)
	if err != nil {
		return nil, fmt.Errorf("decode hex: %w", err)
	}

	privKey, pubKey := btcec.PrivKeyFromBytes(privKeyBytes)

	pubKeyHex := hex.EncodeToString(pubKey.SerializeCompressed()[1:])

	npub, err := HexToNpub(pubKeyHex)
	if err != nil {
		return nil, fmt.Errorf("convert to npub: %w", err)
	}

	nsec, err := HexToNsec(hex.EncodeToString(privKey.Serialize()))
	if err != nil {
		return nil, fmt.Errorf("convert to nsec: %w", err)
	}

	return &ExtendedIdentity{
		Identity: Identity{
			Pubkey: pubKeyHex,
			NPub:   npub,
		},
		PrivKeyHex: privKeyHex,
		PubKeyHex:  pubKeyHex,
		Nsec:       nsec,
	}, nil
}
