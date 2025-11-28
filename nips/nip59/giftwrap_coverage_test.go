package nip59

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/niallyoung/goNDK/event"
)

func TestWrapNilEvent(t *testing.T) {
	_, err := Wrap(nil, "key", "pub")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot be nil")
}

func TestUnwrapNilEvent(t *testing.T) {
	_, err := Unwrap(nil, "key", "pub")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot be nil")
}

func TestUnwrapWrongKind(t *testing.T) {
	ev := event.NewEvent(1, "test", nil, nil, nil, nil, nil)
	_, err := Unwrap(ev, "key", "pub")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not a gift wrap")
}

func TestWrapInvalidKeys(t *testing.T) {
	ev := event.NewEvent(1, "test", nil, nil, nil, nil, nil)
	_, err := Wrap(ev, "invalid", "pub")
	assert.Error(t, err)
}

func TestUnwrapInvalidDecryption(t *testing.T) {
	ev := event.NewEvent(1, "test", nil, nil, nil, nil, nil)
	ev.Sign("0000000000000000000000000000000000000000000000000000000000000001")
	
	wrapped, err := Wrap(ev, "0000000000000000000000000000000000000000000000000000000000000001", "c6047f9441ed7d6d3045406e95c07cd85c778e4b8cef3ca7abac09b95c709ee5")
	require.NoError(t, err)
	
	// Try to decrypt with wrong key
	_, err = Unwrap(wrapped, "0000000000000000000000000000000000000000000000000000000000000003", "79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798")
	assert.Error(t, err)
}

func TestWrapWithTags(t *testing.T) {
	tags := event.Tags{{"e", "event123"}, {"p", "pubkey456"}}
	ev := event.NewEvent(1, "test", tags, nil, nil, nil, nil)
	ev.Sign("0000000000000000000000000000000000000000000000000000000000000001")
	
	wrapped, err := Wrap(ev, "0000000000000000000000000000000000000000000000000000000000000001", "c6047f9441ed7d6d3045406e95c07cd85c778e4b8cef3ca7abac09b95c709ee5")
	require.NoError(t, err)
	
	assert.Equal(t, KindGiftWrap, wrapped.Kind)
	assert.NotEmpty(t, wrapped.Content)
	
	// Verify p tag exists
	found := false
	for _, tag := range wrapped.Tags {
		if len(tag) > 0 && tag[0] == "p" {
			found = true
			break
		}
	}
	assert.True(t, found)
}

func TestGenerateID(t *testing.T) {
	pubkey := "79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798"
	ev := event.NewEvent(1, "test content", event.Tags{{"p", "test"}}, nil, nil, &pubkey, nil)
	id := generateID(ev)
	assert.NotEmpty(t, id)
	assert.Len(t, id, 64) // SHA256 hex is 64 chars
}

func TestUnwrapInvalidJSON(t *testing.T) {
	pubkey := "79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798"
	envelope := event.NewEvent(KindGiftWrap, "invalid encrypted content", event.Tags{{"p", pubkey}}, nil, nil, &pubkey, nil)
	_, err := Unwrap(envelope, "0000000000000000000000000000000000000000000000000000000000000001", pubkey)
	assert.Error(t, err)
}
