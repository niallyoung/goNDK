package nip59

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/niallyoung/goNDK/event"
)

// RED: Test gift wrap creation
func TestWrap(t *testing.T) {
	innerEvent := event.NewEvent(1, "secret message", nil, nil, nil, nil, nil)
	senderPrivKey := "0000000000000000000000000000000000000000000000000000000000000001"
	recipientPubKey := "c6047f9441ed7d6d3045406e95c07cd85c778e4b8cef3ca7abac09b95c709ee5"

	envelope, err := Wrap(innerEvent, senderPrivKey, recipientPubKey)
	require.NoError(t, err)
	assert.NotNil(t, envelope)
	assert.Equal(t, 1059, envelope.Kind)
	assert.NotEmpty(t, envelope.Content)
}

// RED: Test gift wrap unwrap
func TestUnwrap(t *testing.T) {
	innerEvent := event.NewEvent(1, "secret message", nil, nil, nil, nil, nil)
	senderPrivKey := "0000000000000000000000000000000000000000000000000000000000000001"
	recipientPrivKey := "0000000000000000000000000000000000000000000000000000000000000002"
	recipientPubKey := "c6047f9441ed7d6d3045406e95c07cd85c778e4b8cef3ca7abac09b95c709ee5"
	senderPubKey := "79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798"

	envelope, err := Wrap(innerEvent, senderPrivKey, recipientPubKey)
	require.NoError(t, err)

	unwrapped, err := Unwrap(envelope, recipientPrivKey, senderPubKey)
	require.NoError(t, err)
	assert.Equal(t, innerEvent.Content, unwrapped.Content)
	assert.Equal(t, innerEvent.Kind, unwrapped.Kind)
}

// RED: Test round-trip
func TestGiftWrapRoundTrip(t *testing.T) {
	innerEvent := event.NewEvent(1, "The quick brown fox", nil, nil, nil, nil, nil)
	require.NoError(t, innerEvent.Sign("0000000000000000000000000000000000000000000000000000000000000001"))

	senderPrivKey := "0000000000000000000000000000000000000000000000000000000000000001"
	recipientPrivKey := "0000000000000000000000000000000000000000000000000000000000000002"
	recipientPubKey := "c6047f9441ed7d6d3045406e95c07cd85c778e4b8cef3ca7abac09b95c709ee5"
	senderPubKey := "79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798"

	// Wrap
	envelope, err := Wrap(innerEvent, senderPrivKey, recipientPubKey)
	require.NoError(t, err)

	// Unwrap
	unwrapped, err := Unwrap(envelope, recipientPrivKey, senderPubKey)
	require.NoError(t, err)

	assert.Equal(t, innerEvent.Content, unwrapped.Content)
	assert.Equal(t, innerEvent.Kind, unwrapped.Kind)
	assert.Equal(t, *innerEvent.ID, *unwrapped.ID)
}
