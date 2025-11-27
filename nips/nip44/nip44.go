package nip44

// NIP-44: Encrypted Direct Messages using XChaCha20-Poly1305 and HKDF

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2"
	"golang.org/x/crypto/chacha20poly1305"
	"golang.org/x/crypto/hkdf"
)

// Encrypt encrypts plaintext using NIP-44 (XChaCha20-Poly1305 + HKDF)
func Encrypt(plaintext, senderPrivKeyHex, recipientPubKeyHex string) (string, error) {
	// Derive shared secret using ECDH
	sharedSecret, err := deriveSharedSecret(senderPrivKeyHex, recipientPubKeyHex)
	if err != nil {
		return "", fmt.Errorf("derive shared secret: %w", err)
	}

	// Derive encryption key using HKDF
	encKey, err := deriveEncryptionKey(sharedSecret)
	if err != nil {
		return "", fmt.Errorf("derive encryption key: %w", err)
	}

	// Create XChaCha20-Poly1305 cipher
	cipher, err := chacha20poly1305.NewX(encKey)
	if err != nil {
		return "", fmt.Errorf("create cipher: %w", err)
	}

	// Generate random nonce
	nonce := make([]byte, cipher.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return "", fmt.Errorf("generate nonce: %w", err)
	}

	// Encrypt
	ciphertext := cipher.Seal(nonce, nonce, []byte(plaintext), nil)

	// Encode to base64
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt decrypts ciphertext using NIP-44
func Decrypt(ciphertext, recipientPrivKeyHex, senderPubKeyHex string) (string, error) {
	// Decode from base64
	data, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", fmt.Errorf("decode base64: %w", err)
	}

	// Derive shared secret using ECDH
	sharedSecret, err := deriveSharedSecret(recipientPrivKeyHex, senderPubKeyHex)
	if err != nil {
		return "", fmt.Errorf("derive shared secret: %w", err)
	}

	// Derive encryption key using HKDF
	encKey, err := deriveEncryptionKey(sharedSecret)
	if err != nil {
		return "", fmt.Errorf("derive encryption key: %w", err)
	}

	// Create XChaCha20-Poly1305 cipher
	cipher, err := chacha20poly1305.NewX(encKey)
	if err != nil {
		return "", fmt.Errorf("create cipher: %w", err)
	}

	// Extract nonce
	nonceSize := cipher.NonceSize()
	if len(data) < nonceSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	nonce, cipherData := data[:nonceSize], data[nonceSize:]

	// Decrypt
	plaintext, err := cipher.Open(nil, nonce, cipherData, nil)
	if err != nil {
		return "", fmt.Errorf("decrypt: %w", err)
	}

	return string(plaintext), nil
}

// deriveSharedSecret derives ECDH shared secret using secp256k1
func deriveSharedSecret(privKeyHex, pubKeyHex string) ([]byte, error) {
	// Decode private key
	privKeyBytes, err := hex.DecodeString(privKeyHex)
	if err != nil {
		return nil, fmt.Errorf("decode private key: %w", err)
	}

	privKey, _ := btcec.PrivKeyFromBytes(privKeyBytes)

	// Decode public key
	pubKeyBytes, err := hex.DecodeString(pubKeyHex)
	if err != nil {
		return nil, fmt.Errorf("decode public key: %w", err)
	}

	// Parse as compressed public key (33 bytes with prefix)
	var pubKeyFull []byte
	if len(pubKeyBytes) == 32 {
		// Add 0x02 prefix for even y-coordinate
		pubKeyFull = append([]byte{0x02}, pubKeyBytes...)
	} else {
		pubKeyFull = pubKeyBytes
	}

	pubKey, err := btcec.ParsePubKey(pubKeyFull)
	if err != nil {
		return nil, fmt.Errorf("parse public key: %w", err)
	}

	// Compute ECDH shared secret
	sharedX, _ := btcec.S256().ScalarMult(pubKey.X(), pubKey.Y(), privKey.Serialize())
	shared := sharedX.Bytes()

	// Hash for additional security
	hash := sha256.Sum256(shared[:])
	return hash[:], nil
}

// deriveEncryptionKey derives encryption key using HKDF as per NIP-44
func deriveEncryptionKey(sharedSecret []byte) ([]byte, error) {
	// Use HKDF with SHA-256 as per NIP-44 spec
	h := hkdf.New(sha256.New, sharedSecret, nil, []byte("nip44-v2"))

	key := make([]byte, 32) // 256-bit key for XChaCha20-Poly1305
	if _, err := h.Read(key); err != nil {
		return nil, fmt.Errorf("hkdf read: %w", err)
	}

	return key, nil
}
