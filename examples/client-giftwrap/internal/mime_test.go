package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetectMimeType(t *testing.T) {
	tests := []struct {
		filename string
		data     []byte
		expected string
	}{
		{"test.txt", []byte("hello"), "text/plain"},
		{"test.json", []byte(`{"key":"value"}`), "application/json"},
		{"test.png", []byte{0x89, 0x50, 0x4E, 0x47}, "image/png"},
		{"test.jpg", []byte{0xFF, 0xD8, 0xFF}, "image/jpeg"},
		{"test.pdf", []byte{0x25, 0x50, 0x44, 0x46}, "application/pdf"},
		{"unknown", []byte{0x00, 0x01, 0x02}, "application/octet-stream"},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			mime := DetectMimeType(tt.filename, tt.data)
			assert.Equal(t, tt.expected, mime)
		})
	}
}
