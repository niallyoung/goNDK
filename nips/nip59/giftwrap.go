package nip59

// NIP-59: Gift Wrap for encrypted event envelopes

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/niallyoung/goNDK/event"
	"github.com/niallyoung/goNDK/nips/nip44"
)

const KindGiftWrap = 1059

// Wrap creates a NIP-59 gift wrap envelope around an event
func Wrap(innerEvent *event.Event, senderPrivKeyHex, recipientPubKeyHex string) (*event.Event, error) {
	if innerEvent == nil {
		return nil, fmt.Errorf("innerEvent cannot be nil")
	}

	// Serialize inner event
	innerJSON, err := json.Marshal(innerEvent)
	if err != nil {
		return nil, fmt.Errorf("marshal inner event: %w", err)
	}

	// Encrypt using NIP-44
	encryptedContent, err := nip44.Encrypt(string(innerJSON), senderPrivKeyHex, recipientPubKeyHex)
	if err != nil {
		return nil, fmt.Errorf("encrypt: %w", err)
	}

	// Derive sender public key
	senderPrivBytes, _ := hex.DecodeString(senderPrivKeyHex)
	senderPrivKey, senderPubKey := btcec.PrivKeyFromBytes(senderPrivBytes)
	senderPubKeyHex := hex.EncodeToString(senderPubKey.SerializeCompressed()[1:])

	// Create envelope
	envelope := event.NewEvent(
		KindGiftWrap,
		encryptedContent,
		event.Tags{{"p", recipientPubKeyHex}},
		nil,
		nil,
		&senderPubKeyHex,
		nil,
	)

	// Sign envelope
	if err := envelope.Sign(hex.EncodeToString(senderPrivKey.Serialize())); err != nil {
		return nil, fmt.Errorf("sign envelope: %w", err)
	}

	return envelope, nil
}

// Unwrap decrypts a NIP-59 gift wrap envelope
func Unwrap(envelope *event.Event, recipientPrivKeyHex, senderPubKeyHex string) (*event.Event, error) {
	if envelope == nil {
		return nil, fmt.Errorf("envelope cannot be nil")
	}

	if envelope.Kind != KindGiftWrap {
		return nil, fmt.Errorf("not a gift wrap envelope (kind %d)", envelope.Kind)
	}

	// Decrypt content using NIP-44
	decryptedJSON, err := nip44.Decrypt(envelope.Content, recipientPrivKeyHex, senderPubKeyHex)
	if err != nil {
		return nil, fmt.Errorf("decrypt: %w", err)
	}

	// Unmarshal inner event
	var innerEvent event.Event
	if err := json.Unmarshal([]byte(decryptedJSON), &innerEvent); err != nil {
		return nil, fmt.Errorf("unmarshal inner event: %w", err)
	}

	return &innerEvent, nil
}

// generateID creates NOSTR event ID
func generateID(e *event.Event) string {
	tagsJSON, _ := json.Marshal(e.Tags)
	data := fmt.Sprintf(`[0,"%s",%d,%d,%s,"%s"]`,
		*e.Pubkey, e.CreatedAt, e.Kind, string(tagsJSON), e.Content)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
