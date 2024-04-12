package event

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvent_escapeString(t *testing.T) {
	scenarios := []struct {
		Name          string
		givenDst      []byte
		givenContent  string
		expectedBytes []byte
	}{
		{
			"escaped quotations",
			[]byte("destination \"quoted\""),
			"content \"quoted\"",
			[]byte("destination \"quoted\"\"content \\\"quoted\\\"\""),
		},
		{
			"escapted back-slash",
			[]byte("destination \\back-slashed\\"),
			"content \\\\double-slashed\\\\",
			[]byte("destination \\back-slashed\\\"content \\\\\\\\double-slashed\\\\\\\\\""),
		},
		{
			"0x21",
			[]byte{0x21},
			hex.EncodeToString([]byte{0x21}),
			[]byte{0x21, 0x22, 0x32, 0x31, 0x22},
		},
		{
			"0x08",
			[]byte{0x08},
			hex.EncodeToString([]byte{0x08}),
			[]byte{0x8, 0x22, 0x30, 0x38, 0x22},
		},
		{
			"0x09",
			[]byte{0x09},
			hex.EncodeToString([]byte{0x09}),
			[]byte{0x9, 0x22, 0x30, 0x39, 0x22},
		},
		{
			"0x0a",
			[]byte{0x0a},
			hex.EncodeToString([]byte{0x0a}),
			[]byte{0xa, 0x22, 0x30, 0x61, 0x22},
		},
		{
			"0x0c",
			[]byte{0x0c},
			hex.EncodeToString([]byte{0x0c}),
			[]byte{0xc, 0x22, 0x30, 0x63, 0x22},
		},
		{
			"0x0d",
			[]byte{0x0d},
			hex.EncodeToString([]byte{0x0d}),
			[]byte{0xd, 0x22, 0x30, 0x64, 0x22},
		},
		{
			"0x10",
			[]byte{0x10},
			hex.EncodeToString([]byte{0x10}),
			[]byte{0x10, 0x22, 0x31, 0x30, 0x22},
		},
		{
			"0x1a",
			[]byte{0x1a},
			hex.EncodeToString([]byte{0x1a}),
			[]byte{0x1a, 0x22, 0x31, 0x61, 0x22},
		},
		{
			"0x20",
			[]byte{0x20},
			hex.EncodeToString([]byte{0x20}),
			[]byte{0x20, 0x22, 0x32, 0x30, 0x22},
		},
	}

	for _, s := range scenarios {
		t.Run(s.Name, func(t *testing.T) {
			assert.Equal(t, s.expectedBytes, EscapeString(s.givenDst, s.givenContent))
			assert.Equal(t, string(s.expectedBytes), string(EscapeString(s.givenDst, s.givenContent)))
		})
	}
}
