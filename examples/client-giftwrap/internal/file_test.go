package internal

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncodeTextMessage(t *testing.T) {
	msg := EncodeTextMessage("Hello World")
	
	assert.Equal(t, "text", msg.Type)
	assert.Equal(t, "Hello World", msg.Content)
	assert.Nil(t, msg.Metadata)
}

func TestEncodeFileMessage(t *testing.T) {
	data := []byte("test file content")
	msg := EncodeFileMessage(data, "test.txt", "text/plain")
	
	assert.Equal(t, "file", msg.Type)
	assert.NotEmpty(t, msg.Content)
	require.NotNil(t, msg.Metadata)
	assert.Equal(t, "test.txt", msg.Metadata.Filename)
	assert.Equal(t, "text/plain", msg.Metadata.MimeType)
	assert.Equal(t, int64(17), msg.Metadata.Size)
	
	decoded, err := base64.StdEncoding.DecodeString(msg.Content)
	require.NoError(t, err)
	assert.Equal(t, data, decoded)
}

func TestDecodeMessage_Text(t *testing.T) {
	msg := Message{
		Type:    "text",
		Content: "Hello World",
	}
	
	msgType, content, filename, err := DecodeMessage(msg)
	require.NoError(t, err)
	assert.Equal(t, "text", msgType)
	assert.Equal(t, []byte("Hello World"), content)
	assert.Empty(t, filename)
}

func TestDecodeMessage_File(t *testing.T) {
	data := []byte("test file content")
	encoded := base64.StdEncoding.EncodeToString(data)
	
	msg := Message{
		Type:    "file",
		Content: encoded,
		Metadata: &FileMetadata{
			Filename: "test.txt",
			MimeType: "text/plain",
			Size:     17,
		},
	}
	
	msgType, content, filename, err := DecodeMessage(msg)
	require.NoError(t, err)
	assert.Equal(t, "file", msgType)
	assert.Equal(t, data, content)
	assert.Equal(t, "test.txt", filename)
}
