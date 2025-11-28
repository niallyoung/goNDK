package identity

import (
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/btcutil/bech32"
)

// NpubToHex converts npub (bech32) to hex public key
func NpubToHex(npub string) (string, error) {
	hrp, data, err := bech32.Decode(npub)
	if err != nil {
		return "", fmt.Errorf("decode bech32: %w", err)
	}

	if hrp != "npub" {
		return "", fmt.Errorf("invalid hrp: expected npub, got %s", hrp)
	}

	decoded, err := bech32.ConvertBits(data, 5, 8, false)
	if err != nil {
		return "", fmt.Errorf("convert bits: %w", err)
	}

	return hex.EncodeToString(decoded), nil
}

// HexToNpub converts hex public key to npub (bech32)
func HexToNpub(hexPubKey string) (string, error) {
	pubKeyBytes, err := hex.DecodeString(hexPubKey)
	if err != nil {
		return "", fmt.Errorf("decode hex: %w", err)
	}

	converted, err := bech32.ConvertBits(pubKeyBytes, 8, 5, true)
	if err != nil {
		return "", fmt.Errorf("convert bits: %w", err)
	}

	npub, err := bech32.Encode("npub", converted)
	if err != nil {
		return "", fmt.Errorf("encode bech32: %w", err)
	}

	return npub, nil
}

// NsecToHex converts nsec (bech32) to hex private key
func NsecToHex(nsec string) (string, error) {
	hrp, data, err := bech32.Decode(nsec)
	if err != nil {
		return "", fmt.Errorf("decode bech32: %w", err)
	}

	if hrp != "nsec" {
		return "", fmt.Errorf("invalid hrp: expected nsec, got %s", hrp)
	}

	decoded, err := bech32.ConvertBits(data, 5, 8, false)
	if err != nil {
		return "", fmt.Errorf("convert bits: %w", err)
	}

	return hex.EncodeToString(decoded), nil
}

// HexToNsec converts hex private key to nsec (bech32)
func HexToNsec(hexPrivKey string) (string, error) {
	privKeyBytes, err := hex.DecodeString(hexPrivKey)
	if err != nil {
		return "", fmt.Errorf("decode hex: %w", err)
	}

	converted, err := bech32.ConvertBits(privKeyBytes, 8, 5, true)
	if err != nil {
		return "", fmt.Errorf("convert bits: %w", err)
	}

	nsec, err := bech32.Encode("nsec", converted)
	if err != nil {
		return "", fmt.Errorf("encode bech32: %w", err)
	}

	return nsec, nil
}
