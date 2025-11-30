package nip59

import (
	"encoding/json"
	"testing"

	"github.com/niallyoung/goNDK/event"
	"github.com/niallyoung/goNDK/identity"
	"github.com/stretchr/testify/require"
)

// TestNIP59Compliance verifies gift wrap creates NIP-59 compliant events
func TestNIP59Compliance(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping NIP-59 compliance test in short mode")
	}

	// Recipient: npub1ydxa9ss3xkps49s2gck7lk6pptpx79uvh78p87ly8zg0setwaxps3edd7d
	recipientNpub := "npub1ydxa9ss3xkps49s2gck7lk6pptpx79uvh78p87ly8zg0setwaxps3edd7d"
	recipientPubKey, err := identity.NpubToHex(recipientNpub)
	require.NoError(t, err)

	// Create sender identity
	sender, err := identity.Generate()
	require.NoError(t, err)

	// Create inner event (kind 1 text note)
	innerContent := "NIP-59 compliance test from goNDK"
	inner := event.NewEvent(1, innerContent, nil, nil, nil, nil, nil)
	
	// Sign inner event
	err = inner.Sign(sender.PrivKeyHex)
	require.NoError(t, err)

	// Wrap the event
	wrapped, err := Wrap(inner, sender.PrivKeyHex, recipientPubKey)
	require.NoError(t, err)

	// Verify NIP-59 compliance
	t.Run("envelope_is_kind_1059", func(t *testing.T) {
		require.Equal(t, 1059, wrapped.Kind, "Envelope must be kind 1059")
	})

	t.Run("envelope_has_p_tag", func(t *testing.T) {
		require.NotNil(t, wrapped.Tags, "Envelope must have tags")
		require.Greater(t, len(wrapped.Tags), 0, "Envelope must have at least one tag")
		
		foundPTag := false
		for _, tag := range wrapped.Tags {
			if len(tag) >= 2 && tag[0] == "p" && tag[1] == recipientPubKey {
				foundPTag = true
				break
			}
		}
		require.True(t, foundPTag, "Envelope must have p tag with recipient pubkey")
	})

	t.Run("envelope_is_signed", func(t *testing.T) {
		require.NotNil(t, wrapped.Sig, "Envelope must be signed")
		require.NotEmpty(t, *wrapped.Sig, "Envelope signature must not be empty")
	})

	t.Run("envelope_content_is_encrypted", func(t *testing.T) {
		require.NotEmpty(t, wrapped.Content, "Envelope content must not be empty")
		require.NotEqual(t, innerContent, wrapped.Content, "Content must be encrypted")
		// Content is base64 encoded, so we just verify it's not plaintext
	})

	// Note: Unwrap requires the recipient's private key, which we don't have in this test
	// This test verifies the envelope structure is correct

	// Print the event for manual verification
	t.Run("print_event_json", func(t *testing.T) {
		eventJSON, err := json.MarshalIndent(wrapped, "", "  ")
		require.NoError(t, err)
		t.Logf("NIP-59 Gift Wrap Event:\n%s", string(eventJSON))
		t.Logf("\nSender npub: %s", sender.NPub)
		t.Logf("Sender pubkey: %s", sender.PubKeyHex)
	})
}

// TestCreateForNiall creates a gift wrap for Niall's npub for manual verification
func TestCreateForNiall(t *testing.T) {
	// Decode npub1ydxa9ss3xkps49s2gck7lk6pptpx79uvh78p87ly8zg0setwaxps3edd7d
	recipientNpub := "npub1ydxa9ss3xkps49s2gck7lk6pptpx79uvh78p87ly8zg0setwaxps3edd7d"
	recipientPubKey, err := identity.NpubToHex(recipientNpub)
	require.NoError(t, err)

	// Create sender identity
	sender, err := identity.Generate()
	require.NoError(t, err)

	// Create inner event
	innerContent := "Hello Niall! This is a NIP-59 compliance test from goNDK. If you can decrypt this, the implementation is correct!"
	inner := event.NewEvent(1, innerContent, nil, nil, nil, nil, nil)
	
	// Sign inner event
	err = inner.Sign(sender.PrivKeyHex)
	require.NoError(t, err)

	// Wrap the event
	wrapped, err := Wrap(inner, sender.PrivKeyHex, recipientPubKey)
	require.NoError(t, err)

	// Print details for manual publishing and verification
	eventJSON, _ := json.MarshalIndent(wrapped, "", "  ")
	t.Logf("\n=== NIP-59 GIFT WRAP FOR NIALL ===")
	t.Logf("Event ID: %s", *wrapped.ID)
	t.Logf("Sender npub: %s", sender.NPub)
	t.Logf("Sender nsec: %s", sender.Nsec)
	t.Logf("Sender pubkey: %s", sender.PubKeyHex)
	t.Logf("Recipient npub: %s", recipientNpub)
	t.Logf("Recipient pubkey: %s", recipientPubKey)
	t.Logf("\nFull Event JSON:\n%s", string(eventJSON))
	t.Logf("\n=== TO PUBLISH ===")
	t.Logf("Use: echo '%s' | websocat wss://relay.damus.io", string(eventJSON))
	t.Logf("\n=== FOR DECRYPTION ===")
	t.Logf("Event ID: %s", *wrapped.ID)
	t.Logf("Encrypted Content: %s", wrapped.Content)
	t.Logf("Sender pubkey (for decrypt): %s", sender.PubKeyHex)
}
